package main

import "go/ast"

func (g *generator) unwrapSelectorCallChain(selectExpr *ast.SelectorExpr, calleeChain []*ast.CallExpr) (selectChain []*ast.Ident, outCalleeChain []*ast.CallExpr) {
	outCalleeChain = calleeChain
	selectChain = append(selectChain, selectExpr.Sel)
	for selectExpr != nil {
		var exp = selectExpr.X
	reSelect:
		switch xExpr := exp.(type) {
		case *ast.SelectorExpr:
			selectExpr = xExpr
			selectChain = append(selectChain, xExpr.Sel)
		case *ast.Ident:
			selectChain = append(selectChain, xExpr)
			return
		case *ast.CallExpr:
			selectChain = selectChain[:0]
			exp = xExpr.Fun
			outCalleeChain = append(outCalleeChain, xExpr)
			goto reSelect
		default:
			panicGenerateError("not parsed selector", xExpr)
		}
	}
	return
}

func (g *generator) unwrapSelectorChain(selectExpr *ast.SelectorExpr) (selectChain []*ast.Ident) {
	selectChain, _ = g.unwrapSelectorCallChain(selectExpr, nil)
	return
}
