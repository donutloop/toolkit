package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/tabwriter"
	"github.com/fatih/astrewrite"
)

func main() {
	log.SetFlags(log.Ldate | log.Lshortfile | log.Ltime)

	fs := flag.NewFlagSet("xcode", flag.ExitOnError)
	var (
		in                         = fs.String("in", "", "input file")
		out                        = fs.String("out", "", "output file")
		pkg                        = fs.String("pkg", "", "package name")
		typ                        = fs.String("type", "", "type")
		mode                      = fs.String("mode", "", "activate mode")
	)
	fs.Usage = usageFor(fs, "xcode [flags]")
	fs.Parse(os.Args[1:])

	if *in == "" {
		log.Fatal("input file is missing")
	}

	if *out == "" {
		log.Fatal("output file is missing")
	}

	if *pkg == "" {
		log.Fatal("package is missing")
	}

	if *typ == "" {
		log.Fatal("type is missing")
	}

	inputFile, err := ioutil.ReadFile(*in)
	if err != nil {
		log.Fatalf("could not read file (%v)", err)
	}

	rnFunc := RenamePackage(*pkg)
	ctFunc := ChangeType("GenericType", *typ, *mode)

	modifiedFile, err := modifyAst(inputFile, rnFunc, ctFunc)
	if err != nil {
		log.Fatalf("could not modify ast of file (%v)", err)
	}

	if err := ioutil.WriteFile(*out, modifiedFile, 0755); err != nil {
		log.Fatalf("could not write file (%v)", err)
	}
}

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
						if strings.ToLower(typeName) == strings.ToLower(v.Name) {
							x.Args[i] = &ast.Ident{Name: fmt.Sprintf("%s.(%s)", v.Name, newType)}
						}
					}
				}
				return x, true
			default:
				if debugMode == DebugMode {
					log.SetFlags(0)
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

func modifyAst(dest []byte, fns ...func(*ast.File) *ast.File) ([]byte, error) {
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
		return nil, fmt.Errorf("couldn't format package code (%v)", err)
	}

	return buf.Bytes(), nil
}

func usageFor(fs *flag.FlagSet, short string) func() {
	return func() {
		fmt.Fprintf(os.Stdout, "USAGE\n")
		fmt.Fprintf(os.Stdout, "  %s\n", short)
		fmt.Fprintf(os.Stdout, "\n")
		fmt.Fprintf(os.Stdout, "FLAGS\n")
		tw := tabwriter.NewWriter(os.Stdout, 0, 2, 2, ' ', 0)
		fs.VisitAll(func(f *flag.Flag) {
			if f.Name == "debug" {
				return
			}
			def := f.DefValue
			if def == "" {
				def = "..."
			}
			fmt.Fprintf(tw, "  -%s %s\t%s\n", f.Name, f.DefValue, f.Usage)
		})
		tw.Flush()
	}
}

