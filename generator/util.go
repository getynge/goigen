package generator

import "go/ast"

func buildTypeFromExpr(expr ast.Expr) (name string) {
	current := expr
	for {
		star, ok := current.(*ast.StarExpr)
		if ok {
			current = star.X
			name += "*"
			continue
		}
		ident, ok := current.(*ast.Ident)
		if !ok {
			return name
		}
		name += ident.Name
		return name
	}
}
