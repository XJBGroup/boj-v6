package main

import (
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
	"go/ast"
	"go/parser"
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
		if obj != nil {
			if obj.Decl == nil {
				panic("todo")
			}

			if ts, ok := obj.Decl.(*ast.TypeSpec); !ok {
				panic("todo")
			} else {
				if st, ok := ts.Type.(*ast.StructType); !ok {
					panic("todo")
				} else {
					for _, f := range st.Fields.List {

						if sel, ok := f.Type.(*ast.SelectorExpr); ok {
							chain := g.unwrapSelectorChain(sel)

							p := chain[len(chain)-1].Obj.Decl.(*ast.ImportSpec).Path.Value
							p = p[1 : len(p)-1]
							if p != "github.com/Myriad-Dreamin/boj-v6/cmd/generate/stub" {
								continue
							}

							var fillCheck = func(object *ast.Object) {

								if chain[0].Obj == nil {
									chain[0].Obj = object
								}
								if chain[0].Obj != object {
									panicGenerateError("not equal...", chain[0])
								}
							}

							switch chain[0].Name {
							case "Stub":
								fillCheck(g.stubObject)
							case "StubVariables":
								fillCheck(g.stubVariableObject)
							default:
								panic("todo")
							}

							for _, name := range f.Names {
								g.stubFieldInfo[name.Name] = chain[0].Obj
							}
							//if len(chain) == 2 {
							//	var stubType = StubTypeLength
							//	switch chain[0].Name {
							//	case "Stub":
							//		stubType = StubTypeStub
							//	case "StubVariables":
							//		stubType = StubTypeStubVariables
							//	default:
							//		panic("todo")
							//	}
							//	if len(f.Names) == 0 {
							//		stubTypes[stubType] = append(stubTypes[stubType], chain[0].Name)
							//	} else {
							//		for _, n := range f.Names {
							//			stubTypes[stubType] = append(stubTypes[stubType], n.Name)
							//		}
							//	}
							//
							//} else {
							//	panic("todo")
							//}
						}
					}
				}
			}

			return obj
		}
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

func (g *generator) addLocal(ident *ast.Ident) {
	if g.localName == nil {
		g.localName = make(map[string]*ast.Ident)
	}
	g.localName[ident.Name] = ident
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
			// todo: type insert
			//g.printAstNode(lhs)
			//g.printAstNode(rhs.Args)
			return notParsedStmt(rhs)
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

func (g *generator) evalStubParseChain(fieldName string) (parserChainMap ParseChainedMap, _ error) {
	if obj := g.getStubFieldByName(fieldName); obj != nil {
		switch obj {
		case g.stubObject:
			parserChainMap = selfStubFilter
		default:
			panicUnknownAt(g.stmtParsingNode)
			return nil, unknownAt(g.stmtParsingNode)
		}
	} else {
		return nil, notParsedStmt(g.stmtParsingNode)
	}
	return parserChainMap, nil
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

func (g *generator) resolveBuiltIn() *ast.Package {

	pkgName := "builtin"
	pkg, err := parsePkg(pkgName, g.fileSet, parser.ParseComments|parser.DeclarationErrors)
	sugar.HandlerError0(err)
	return pkg
}

func (g *generator) resolvePackageObject(file *ast.File) {
	i := 0
	for _, ident := range file.Unresolved {
		var resolved = false

		for _, importSpec := range file.Imports {
			if importSpec.Name != nil && importSpec.Name.Name == ident.Name {
				ident.Obj = ast.NewObj(ast.Pkg, importSpec.Name.Name)
				ident.Obj.Decl = importSpec
				resolved = true
				break
			}
		}

		if !resolved {
			file.Unresolved[i] = ident
			i++
		}
	}
	file.Unresolved = file.Unresolved[0:i]
}

func main() {
}
