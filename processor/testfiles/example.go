//go:generate goigen . *Example IExample
package testfiles

import "go/ast"

type Example struct{}

func (e Example) NotIncluded()                               {}
func (e *Example) Included()                                 {}
func (e *Example) VeryIncluded() int                         { return 0 }
func (e *Example) EvenWorksWithPointers() (*int, error)      { result := 0; return &result, nil }
func (e *Example) WorksWithArgumentsToo(i int)               {}
func (e *Example) WorksWithManyArguments(i, j int)           {}
func (e *Example) WorksWithAutomaticImports(a *ast.StarExpr) {}
func (*Example) AlsoIncluded()                               {}
