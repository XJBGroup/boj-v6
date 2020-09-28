package main

import (
	"bytes"
	"fmt"
	"go/ast"
)

type ParseChainedMap map[string]func(g *generator, lhs []*ast.Ident, fnExpr *ast.CallExpr) (ParseChainedMap, error)

var (
	selfStubFilter         ParseChainedMap
	selfInvokingStubFilter ParseChainedMap
)

func init() {
	selfStubFilter = ParseChainedMap{
		"GetID": func(g *generator, lhs []*ast.Ident, fnExpr *ast.CallExpr) (ParseChainedMap, error) {

			panicOnInvoking(lhs, 1, fnExpr, 0)

			var stmt = fmt.Sprintf(`
%v, ok := snippet.ParseUint(c, ctrl.key)
if !ok {
	return
}`, lhs[0].Name)
			g.methodStmts = append(g.methodStmts, stmt)

			return nil, nil
		},
		"Bind": func(g *generator, lhs []*ast.Ident, fnExpr *ast.CallExpr) (ParseChainedMap, error) {
			panicOnInvoking(lhs, 0, fnExpr, 1)
			node := g.formatNode(fnExpr.Args[0])

			var stmt = fmt.Sprintf(`
if !snippet.BindRequest(c, %v) {
	return
}`, node)
			g.methodStmts = append(g.methodStmts, stmt)

			return nil, nil
		},
		"Context": func(g *generator, lhs []*ast.Ident, fnExpr *ast.CallExpr) (ParseChainedMap, error) {
			var args []*ast.Ident
			for _, v := range fnExpr.Args {
				switch exp := v.(type) {
				case *ast.Ident:
					args = append(args, exp)
				case *ast.StarExpr:
					return nil, unknownAt(v)
				default:
					return nil, unknownAt(v)
				}
			}

			g.contextVars = append(g.contextVars, args...)
			return selfInvokingStubFilter, nil
		},

		"EmitSelf": func(g *generator, lhs []*ast.Ident, fnExpr *ast.CallExpr) (ParseChainedMap, error) {
			// todo
			return nil, notParsedStmt(fnExpr)
		},
	}

	selfInvokingStubFilter = ParseChainedMap{
		"Serve": func(g *generator, lhs []*ast.Ident, fnExpr *ast.CallExpr) (ParseChainedMap, error) {
			panicOnInvoking(lhs, 0, fnExpr, -1)
			if len(fnExpr.Args) == 0 {
				panicUnknownAt(fnExpr)
			}

			inputs, output := fnExpr.Args[:len(fnExpr.Args)-1], fnExpr.Args[len(fnExpr.Args)-1]

			var b = bytes.NewBuffer(make([]byte, 30))
			if !g.hasErrorDeclared {
				b.WriteString("var err error\n")
			}

			g.printNode_(b, output)
			b.WriteString(", err = svc.service.Do(")
			for _, input := range inputs {
				// todo: context ?
				g.printNode_(b, input)
				b.WriteString(", ")
			}
			b.WriteString(")\nif err != nil {\n    snippet.DoReport(err)\n}")

			g.methodStmts = append(g.methodStmts, b.String())

			return nil, nil
		},
	}

	for k, v := range selfInvokingStubFilter {
		if _, ok := selfStubFilter[k]; ok {
			panic("merge error on " + k)
		}
		selfStubFilter[k] = v
	}
}
