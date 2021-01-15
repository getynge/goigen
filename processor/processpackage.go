package processor

import (
	"go/ast"
	"strings"
)

func processExpr(expr ast.Expr, name string, targetPointerCount int) (shouldAppend bool) {
	current := expr
	for i := 0; i <= targetPointerCount; i++ {
		star, ok := current.(*ast.StarExpr)
		if ok {
			current = star.X
			continue
		}
		ident, ok := current.(*ast.Ident)
		if !ok {
			return false
		}
		if ident.Name == name && i == targetPointerCount {
			return true
		}
		return false
	}
	return false
}

func ProcessPackage(pkg *ast.Package, targetStruct string) (methods []*ast.FuncDecl) {
	ast.Inspect(pkg, func(node ast.Node) bool {
		f, ok := node.(*ast.FuncDecl)
		targetPointerCount := strings.Count(targetStruct, "*")
		cleanName := strings.TrimPrefix(targetStruct, "*")
		if !ok || f.Recv == nil {
			return true
		}
		for _, v := range f.Recv.List {
			if processExpr(v.Type, cleanName, targetPointerCount) {
				methods = append(methods, f)
			}
		}
		return true
	})
	return methods
}
