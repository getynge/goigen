package processor

import (
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
)

func ProcessDirectory(directory, targetStruct string) (methods []*ast.FuncDecl, pkg *ast.Package, err error) {
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, directory, nil, parser.AllErrors)

	if err != nil {
		return nil, nil, err
	}

	for _, v := range pkgs {
		if !strings.HasSuffix(v.Name, "_test") {
			if pkg == nil {
				pkg = v
			}
			methods = append(methods, ProcessPackage(v, targetStruct)...)
		}
	}

	return methods, nil, nil
}
