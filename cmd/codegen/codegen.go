package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"html/template"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/Masterminds/sprig"
)

var defaultOutputPath = "../../pkg/nr/nr-rrc-definitions"

type Data struct {
	InterfaceName   string
	ContainerName   string
	FieldName       string
	Implementations []Implementation
}

type Implementation struct {
	Name string
}

func main() {
	pbFile := flag.String("pbFile", path.Join(defaultOutputPath, "nr-rrc-definitions.pb.go"), "pb file name")

	flag.Parse()

	datas := parseForInterfacess(*pbFile)
	processInterfaces("interface.tmpl", *pbFile, datas)
}

func processTemplate(tplFileName, pbFile string, data *Data) {
	data.Implementations = parseForImplementations(pbFile, data.ContainerName)
	tmpl := template.Must(template.New("").Funcs(sprig.FuncMap()).ParseFiles(tplFileName))
	var processed bytes.Buffer
	err := tmpl.ExecuteTemplate(&processed, tplFileName, data)
	if err != nil {
		log.Fatalf("Unable to parse data into template: %v\n", err)
	}
	formatted, err := format.Source(processed.Bytes())
	if err != nil {
		log.Fatalf("Could not format processed template: %v\n", err)
	}
	outputPath := filepath.Dir(pbFile)
	outputFile := path.Join(outputPath, data.ContainerName+".go")
	fmt.Println("Writing file: ", outputFile)
	f, _ := os.Create(outputFile)
	w := bufio.NewWriter(f)
	w.WriteString(string(formatted))
	w.Flush()
}

func processInterfaces(tplFileName, pbFile string, datas []*Data) {
	for _, data := range datas {
		processTemplate(tplFileName, pbFile, data)
	}
}

func parseForImplementations(filename, containerName string) []Implementation {
	imps := make([]Implementation, 0, 8)
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
	if err != nil {
		fmt.Println("Error parsing file: " + filename)
		os.Exit(255)
	}

	// traverse all tokens
	ast.Inspect(node, func(n ast.Node) bool {
		switch t := n.(type) {
		case *ast.TypeSpec:
			if strings.HasPrefix(t.Name.Name, containerName+"_") {
				switch t.Type.(type) {
				case *ast.StructType:
					imps = append(imps, Implementation{Name: t.Name.Name})
				}
			}
		}
		return true
	})

	return imps
}

func parseForInterfacess(filename string) []*Data {
	inters := make([]*Data, 0, 8)
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
	if err != nil {
		fmt.Println("Error parsing file: " + filename)
		os.Exit(255)
	}

	// traverse all tokens
	ast.Inspect(node, func(n ast.Node) bool {
		switch t := n.(type) {
		case *ast.TypeSpec:
			switch t.Type.(type) {
			case *ast.InterfaceType:
				iscontainerName := strings.Split(t.Name.Name, "_")[0]
				fieldName := strings.Split(t.Name.Name, "_")[1]
				containerName := iscontainerName[2:len(iscontainerName)]
				inters = append(inters, &Data{
					InterfaceName: t.Name.Name,
					ContainerName: containerName,
					FieldName:     fieldName,
				})
			}
		}
		return true
	})

	return inters
}
