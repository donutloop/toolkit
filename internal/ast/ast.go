package ast

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"log"
	"strings"

	"github.com/fatih/astrewrite"
)

const (
	DebugMode string = "DEV"
)

func RenamePackage(packageName string) func(file *ast.File) *ast.File {
	return func(file *ast.File) *ast.File {
		file.Name = &ast.Ident{Name: packageName}
		return file
	}
}

func ChangeType(typeName string, newType string, debugMode string) func(file *ast.File) *ast.File {
	return func(file *ast.File) *ast.File {
		rewriteFunc := func(n ast.Node) (ast.Node, bool) {
			switch x := n.(type) {
			case *ast.Ident:
				if typeName == x.Name {
					x = &ast.Ident{Name: newType}
				}
				return x, true
			case *ast.CallExpr:
				for i := 0; i < len(x.Args); i++ {
					v, ok := x.Args[i].(*ast.Ident)
					if ok {
						if strings.EqualFold(typeName, v.Name) {
							x.Args[i] = &ast.Ident{Name: fmt.Sprintf("%s.(%s)", v.Name, newType)}
						}
					}
				}
				return x, true
			default:
				if debugMode == DebugMode {
					log.Println("ast node:")
					log.Println(fmt.Sprintf("verbose value: %#v", x))
					log.Println(fmt.Sprintf("type: %T", x))
					log.Println(fmt.Sprintf("value: %v", x))
				}
			}
			return n, true
		}

		astrewrite.Walk(file, rewriteFunc)

		return file
	}
}

func ModifyAst(dest []byte, fns ...func(*ast.File) *ast.File) ([]byte, error) {
	destFset := token.NewFileSet()
	destF, err := parser.ParseFile(destFset, "", dest, 0)
	if err != nil {
		return nil, err
	}

	for _, fn := range fns {
		destF = fn(destF)
	}

	var buf bytes.Buffer
	if err := format.Node(&buf, destFset, destF); err != nil {
		return nil, &BadFormattedCode{Err: err}
	}

	return buf.Bytes(), nil
}

type BadFormattedCode struct {
	Err error
}

func (e BadFormattedCode) Error() string {
	return fmt.Sprintf("couldn't format package code (%v)", e.Err)
}
