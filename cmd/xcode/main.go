package main

import (
	"flag"
	"fmt"
	"github.com/donutloop/toolkit/internal/ast"
	"io/ioutil"
	"log"
	"os"
	"text/tabwriter"
)

func main() {
	log.SetFlags(0)

	fs := flag.NewFlagSet("xcode", flag.ExitOnError)
	var (
		in   = fs.String("in", "", "input file")
		out  = fs.String("out", "", "output file")
		pkg  = fs.String("pkg", "", "package name")
		typ  = fs.String("type", "", "type")
		mode = fs.String("mode", "", "activate mode")
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

	rnFunc := ast.RenamePackage(*pkg)
	ctFunc := ast.ChangeType("GenericType", *typ, *mode)

	modifiedFile, err := ast.ModifyAst(inputFile, rnFunc, ctFunc)
	if err != nil {
		log.Fatalf("could not modify ast of file (%v)", err)
	}

	if err := ioutil.WriteFile(*out, modifiedFile, 0755); err != nil {
		log.Fatalf("could not write file (%v)", err)
	}
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
