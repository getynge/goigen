package generator

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/types"
	"golang.org/x/tools/imports"
	"path"
	"strings"
	"text/template"
)

type IdentifierTemplate struct {
	Name    string
	Type    string
	HasNext bool
}

type MethodTemplate struct {
	Name      string
	Arguments []IdentifierTemplate
	Returns   []IdentifierTemplate
}

type FileTemplate struct {
	FileName        string // name of file without extension or directory
	Package         string
	TargetInterface string
	Methods         []MethodTemplate
}

func NewFileTemplate(ast []*ast.FuncDecl, pkg *ast.Package, targetInterface string) (ft *FileTemplate) {
	ft = &FileTemplate{}
	ft.FileName = fmt.Sprintf("generated_%s", strings.ToLower(targetInterface))
	ft.Package = pkg.Name
	ft.TargetInterface = targetInterface

	for _, v := range ast {
		methodTemplate := MethodTemplate{}
		methodTemplate.Name = v.Name.Name
		if v.Type.Params != nil {
			for i, arg := range v.Type.Params.List {
				identifierTemplate := IdentifierTemplate{}
				identifierTemplate.HasNext = (i + 1) != len(v.Type.Params.List)
				identifierTemplate.Type = types.ExprString(arg.Type)
				if len(arg.Names) > 0 {
					identifierTemplate.Name = arg.Names[0].Name
					for _, name := range arg.Names[1:] {
						identifierTemplate.Name += "," + name.Name
					}
				}
				methodTemplate.Arguments = append(methodTemplate.Arguments, identifierTemplate)
			}
		}
		if v.Type.Results != nil {
			for i, arg := range v.Type.Results.List {
				identifierTemplate := IdentifierTemplate{}
				identifierTemplate.HasNext = (i + 1) != len(v.Type.Results.List)
				identifierTemplate.Type = types.ExprString(arg.Type)
				if len(arg.Names) > 0 {
					identifierTemplate.Name = arg.Names[0].Name
					for _, name := range arg.Names[1:] {
						identifierTemplate.Name += "," + name.Name
					}
				}
				methodTemplate.Returns = append(methodTemplate.Returns, identifierTemplate)
			}
		}
		ft.Methods = append(ft.Methods, methodTemplate)
	}

	return ft
}

func (f *FileTemplate) Generate(targetDirectory string) (string, error) {
	var buf bytes.Buffer
	tmpl, err := template.New(f.FileName).Parse(source)

	if err != nil {
		return "", err
	}

	err = tmpl.Execute(&buf, f)

	if err != nil {
		return "", err
	}

	b := buf.Bytes()

	result, err := imports.Process(path.Join(targetDirectory, f.FileName+".go"), b, nil)

	return string(result), err
}
