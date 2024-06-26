//go:build ignore

package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path"
	"strings"
	"text/template"
	"time"

	"golang.org/x/tools/go/ast/inspector"
)

var headerAliasTpl = template.Must(template.New("").Parse(`
package {{ .BasePackageName }}

import (
	{{- range .Imports }}
	"github.com/niklak/binance_connector/{{ . }}"
	{{- end }}
)

`))

var headerServiceTpl = template.Must(template.New("").Parse(`
package {{ .BasePackageName }}

`))

var serviceTpl = template.Must(template.New("").Parse(`
func (c *Client) {{ .FuncName }}() *{{ .EntityName }} {
	return &{{ .EntityName }}{C: c.Connector}
}
`))

var aliasTpl = template.Must(template.New("").Parse(`type {{ .EntityName }} = {{ .PackageName }}.{{ .EntityName }}
`))

type service struct {
	EntityName string
	FuncName   string
}

type fileAliases struct {
	FileName string
	Aliases  []alias
}

type alias struct {
	EntityName  string
	PackageName string
}

type aliasHeader struct {
	BasePackageName string
	Timestamp       time.Time
	Imports         []string
}

func getSrcFiles(src string) (allFiles []string, err error) {
	files, err := os.ReadDir(src)
	if err != nil {
		return
	}
	for _, f := range files {
		if f.IsDir() {
			children, err := getSrcFiles(path.Join(src, f.Name()))
			if err != nil {
				return nil, err
			}
			allFiles = append(allFiles, children...)
			continue
		}
		if !strings.HasSuffix(f.Name(), ".go") {
			continue
		}
		if strings.HasSuffix(f.Name(), "_test.go") {
			continue
		}
		allFiles = append(allFiles, path.Join(src, f.Name()))

	}

	return
}

func genFuncName(typeName string) string {
	funcNameBuilder := strings.Builder{}
	funcNameBuilder.WriteString("New")
	funcNameBuilder.WriteString(typeName)
	if !strings.HasSuffix(typeName, "Service") {
		funcNameBuilder.WriteString("Service")
	}
	return funcNameBuilder.String()
}

func generateServices(basePackageName string, services []service) (buf *bytes.Buffer, err error) {
	buf = bytes.NewBuffer(nil)

	err = headerServiceTpl.Execute(buf, aliasHeader{
		BasePackageName: basePackageName,
		Timestamp:       time.Now(),
	})

	for _, s := range services {
		err = serviceTpl.Execute(buf, s)
		if err != nil {
			return
		}
		buf.WriteString("\n")
	}
	return

}

func generateAliases(basePackageName string, fAliases []fileAliases, allFiles []string) (buf *bytes.Buffer, err error) {
	buf = bytes.NewBuffer(nil)

	var packagesMap = make(map[string]struct{})

	for _, fp := range allFiles {
		dir, _ := path.Split(fp)
		dir = strings.TrimSuffix(dir, "/")
		packagesMap[dir] = struct{}{}
	}

	sourcePackages := make([]string, 0, len(packagesMap))

	for k := range packagesMap {
		sourcePackages = append(sourcePackages, k)
	}

	err = headerAliasTpl.Execute(buf, aliasHeader{
		BasePackageName: basePackageName,
		Timestamp:       time.Now(),
		Imports:         sourcePackages,
	})

	for _, fA := range fAliases {
		if len(fA.Aliases) == 0 {
			continue
		}
		buf.WriteString("//file: " + fA.FileName + "\n")
		for _, al := range fA.Aliases {
			err = aliasTpl.Execute(buf, al)
			if err != nil {
				return
			}
		}
		buf.WriteString("\n")
	}
	return
}

func generate(basePackageName string, files []string, outAliasFp string, outServiceFp string) (err error) {

	fAliases := make([]fileAliases, 0)
	services := make([]service, 0)
	for _, fp := range files {
		aliases := make([]alias, 0)
		astFile, err := parser.ParseFile(
			token.NewFileSet(),
			fp,
			nil,
			parser.ParseComments,
		)
		if err != nil {
			return err
		}

		packageName := astFile.Name.Name

		inspektr := inspector.New([]*ast.File{astFile})

		filtr := []ast.Node{
			&ast.GenDecl{},
		}
		inspektr.Nodes(filtr, func(node ast.Node, push bool) (proceed bool) {
			nodeDecl := node.(*ast.GenDecl)

			typeSpec, ok := nodeDecl.Specs[0].(*ast.TypeSpec)
			if !ok {
				return false
			}

			_, ok = typeSpec.Type.(*ast.StructType)
			if !ok {
				return false
			}

			aliases = append(aliases, alias{
				EntityName:  typeSpec.Name.Name,
				PackageName: packageName,
			})

			if nodeDecl.Doc == nil {
				return false
			}

			for _, comment := range nodeDecl.Doc.List {
				switch comment.Text {
				case "//gen:new_service":
					services = append(services, service{
						EntityName: typeSpec.Name.Name,
						FuncName:   genFuncName(typeSpec.Name.Name),
					})
				}
			}

			return false
		})
		fAliases = append(fAliases, fileAliases{
			FileName: fp,
			Aliases:  aliases,
		})
	}

	bufAlias, err := generateAliases(basePackageName, fAliases, files)
	if err != nil {
		return
	}

	if err = writeResultFile(outAliasFp, bufAlias); err != nil {
		return
	}

	bufService, err := generateServices(basePackageName, services)
	if err != nil {
		return
	}
	err = writeResultFile(outServiceFp, bufService)

	return
}

func writeResultFile(fp string, buf *bytes.Buffer) (err error) {
	outFile, err := os.Create(fp)
	if err != nil {
		return
	}
	defer outFile.Close()

	_, err = fmt.Fprint(outFile, buf)

	return
}

func main() {

	sources := flag.String("src", "", "comma-separated list of source files or directories to apply")
	dstAlias := flag.String("dst-alias", "", "file to write output")
	dstService := flag.String("dst-service", "", "file to write output")
	flag.Parse()

	var srcList []string
	if len(*sources) > 0 {
		srcList = strings.Split(*sources, ",")
	}

	fp := os.Getenv("GOFILE")
	if fp == "" {
		log.Fatal("GOFILE is missing")
	}

	var outAliasFp string
	if len(*dstAlias) > 0 {
		outAliasFp = *dstAlias
	} else {
		outAliasFp = strings.TrimSuffix(fp, ".go") + "_alias_gen.go"
	}

	var outServiceFp string
	if len(*dstService) > 0 {
		outServiceFp = *dstService
	} else {
		outServiceFp = strings.TrimSuffix(fp, ".go") + "_service_gen.go"
	}

	var allFiles []string
	for _, src := range srcList {
		files, err := getSrcFiles(src)
		if err != nil {
			log.Fatal(err)
		}
		allFiles = append(allFiles, files...)
	}

	baseAstFile, err := parser.ParseFile(
		token.NewFileSet(),
		fp,
		nil,
		parser.ParseComments,
	)
	if err != nil {
		log.Fatalf("parse file: %v", err)
	}

	err = generate(baseAstFile.Name.Name, allFiles, outAliasFp, outServiceFp)

	if err != nil {
		log.Fatal(err)
	}

}
