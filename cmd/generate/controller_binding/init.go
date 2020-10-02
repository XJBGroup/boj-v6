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

			g.addContext(lhs[0])

			var stmt = fmt.Sprintf(`
ctx.%v, ok = snippet.ParseUint(c, ctrl.key)
if !ok {
	return
}`, lhs[0].Name)

			if !g.hasOkDeclared {
				stmt = "var ok bool\n" + stmt
				g.hasOkDeclared = true
			}
			g.methodStmts = append(g.methodStmts, stmt)

			return nil, nil
		},
		"Bind": func(g *generator, lhs []*ast.Ident, fnExpr *ast.CallExpr) (ParseChainedMap, error) {
			panicOnInvoking(lhs, 0, fnExpr, 1)
			node := g.formatNode(fnExpr.Args[0])

			g.addContext(fnExpr.Args[0].(*ast.Ident))

			var stmt = fmt.Sprintf(`
ctx.%v = %v
if !snippet.BindRequest(c, ctx.%v) {
	return
}`, node, node, node)
			g.methodStmts = append(g.methodStmts, stmt)

			return nil, nil
		},
		"Context": func(g *generator, lhs []*ast.Ident, fnExpr *ast.CallExpr) (ParseChainedMap, error) {
			for _, v := range fnExpr.Args {
				switch exp := v.(type) {
				case *ast.Ident:

					g.addContext(exp)

					var b = bytes.NewBuffer(make([]byte, 15))

					b.WriteString("ctx.")
					g.printNode_(b, v)
					b.WriteString(" = ")
					g.printNode_(b, v)
					g.methodStmts = append(g.methodStmts, b.String())
				case *ast.StarExpr:
					return nil, unknownAt(v)
				default:
					return nil, unknownAt(v)
				}
			}
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
				g.hasErrorDeclared = true
			}

			g.printNode_(b, output)
			b.WriteString(", err = svc.service.Do(ctx, ")
			for _, input := range inputs {
				// todo: context ?
				if _, ok := g.contextVars[input.(*ast.Ident).Name]; !ok {
					panicGenerateError("todo", input)
				}
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
