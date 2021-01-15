package generator

import (
	"go/ast"
	"strings"
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
	ft.FileName = strings.ToLower(targetInterface)
	ft.Package = pkg.Name
	ft.TargetInterface = targetInterface

	for _, v := range ast {
		toAdd := MethodTemplate{}
		toAdd.Name = v.Name.Name
		for i, arg := range v.Type.Params.List {

		}
	}

	return ft
}
