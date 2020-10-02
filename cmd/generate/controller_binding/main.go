package main

import (
	"go/ast"
	"go/token"
	"sync"
)

//type invokeParams struct {
//	invokeName string
//	params     []interface{}
//}

//type stub struct {
//	sequence []invokeParams
//}

type generator struct {
	mutex     sync.Mutex
	parseFlag ParseFlag

	stubObject         *ast.Object
	stubVariableObject *ast.Object

	fileSet  *token.FileSet
	packages map[string]*ast.Package

	stubImportSpecsMapping map[string]map[*ast.File]string

	objectScope
	methodScope
	stmtScope
}

func (g *generator) findInScopes(p *ast.Package, name string) *ast.Object {
	g.stubFieldInfo = make(map[string]*ast.Object)

	for _, file := range p.Files {
		obj := file.Scope.Lookup(name)
		if obj == nil {
			continue
		}

		return g.checkFillStubFields(obj)
	}
	return nil
}

func (g *generator) getMethods(p *ast.Package, obj *ast.Object) (methods []*ast.FuncDecl) {
	methods = make([]*ast.FuncDecl, 0, 5)
	for _, file := range p.Files {
		for _, decl := range file.Decls {
			if decl, ok := decl.(*ast.FuncDecl); ok && decl.Recv != nil {
				for _, recv := range decl.Recv.List {
					if recv.Type == nil {
						panicGenerateError("null type", recv)
					}

					var t = recv.Type
					for t != nil {
						switch st := t.(type) {
						case *ast.Ident:
							if obj == st.Obj {
								methods = append(methods, decl)
							}
							t = nil
						case *ast.StarExpr:
							t = st.X
						default:
							panicGenerateError("todo field", t)
						}
					}
				}
			}
		}
	}
	return
}

func (g *generator) tryParseLRShapeStmt_(obj *ast.Object, identifiers []*ast.Ident, rhs []ast.Expr) error {
	for _, ident := range identifiers {
		g.addLocal(ident)
	}

	if len(rhs) == 1 {
		switch rhs := rhs[0].(type) {
		case *ast.CallExpr:
			return g.tryParseCallExpr(obj, identifiers, rhs)
		}
	}

	panicGenerateError_("not parsed stmt")
	return notParsedStmt(nil)
}

func (g *generator) tryParseLRShapeStmt(obj *ast.Object, lhs, rhs []ast.Expr) error {
	if len(rhs) == 1 {
		var identifiers []*ast.Ident

		for _, lhs := range lhs {
			switch lhs := lhs.(type) {
			case *ast.Ident:
				identifiers = append(identifiers, lhs)
			default:
				panicGenerateError("broken ident", lhs)
			}
		}

		return g.tryParseLRShapeStmt_(obj, identifiers, rhs)
	}

	panicGenerateError_("not parsed stmt")
	return notParsedStmt(nil)
}

func (g *generator) tryParseCallExpr(obj *ast.Object, lhs []*ast.Ident, rhs *ast.CallExpr) error {

	switch fnExpr := rhs.Fun.(type) {
	case *ast.Ident:
		// todo find local function
		switch fnExpr.Name {
		case "new":

			if len(lhs) != 1 {
				panicGenerateError("want single variable assigning 1 variable", rhs)
				return notParsedStmt(rhs)
			}

			g.appendStmt(g.methodParsingStmt)
			return nil
		}

		return notParsedStmt(rhs)
	case *ast.SelectorExpr:
		if err := g.functionIsMethodOfType(obj, fnExpr); err != nil {
			return err
		}
		return g.tryParseSelfCalling(lhs, rhs)
	default:
	}
	panicUnknownAt(rhs.Fun)
	return unknownAt(rhs.Fun)
}

func (g *generator) tryParseSelfCalling(lhs []*ast.Ident, fnExpr *ast.CallExpr) (err error) {
	g.stmtParsingNode = fnExpr

	chain, calleeChain := g.unwrapSelectorCallChain(fnExpr.Fun.(*ast.SelectorExpr), []*ast.CallExpr{fnExpr})

	if len(chain) < 2 {
		panic("invalid call expr")
	}

	if len(chain) == 3 {
		var parserChainMap ParseChainedMap
		parserChainMap, err = g.evalStubParseChain(chain[1].Name)
		if err != nil {
			return err
		}

		for ptr := len(calleeChain) - 1; ptr >= 0; ptr-- {
			if parserChainMap == nil {
				panicGenerateError("not consumed callee chain pointer", fnExpr)
			}

			if parserFunc, ok := parserChainMap[calleeChain[ptr].Fun.(*ast.SelectorExpr).Sel.Name]; !ok {
				panicUnknownAt(fnExpr)
			} else {
				// todo: check struct type
				parserChainMap, err = parserFunc(g, lhs, calleeChain[ptr])
				if err != nil {
					return err
				}
			}
		}
		return nil
	}

	return unknownAt(fnExpr)
}

func (g *generator) getOwnerOfSelector(expr *ast.SelectorExpr) *ast.Ident {
	chain := g.unwrapSelectorChain(expr)
	return chain[len(chain)-1]
}

func (g *generator) functionIsMethodOfType(obj *ast.Object, fnExpr *ast.SelectorExpr) error {
	ident := g.getOwnerOfSelector(fnExpr)
	if ident == nil || ident.Obj.Decl == nil {
		return notParsedStmt(fnExpr)
	}

	if decl, ok := ident.Obj.Decl.(*ast.Field); ok {
		switch t := decl.Type.(type) {
		case *ast.Ident:
			if obj != t.Obj && t.Name == obj.Name {
				panicGenerateError("not parsed but name is same", fnExpr)
			}
			return nil
		}
	}
	return unknownAt(fnExpr)
}

func main() {
}
