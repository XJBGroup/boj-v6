package main

import "go/ast"

func (g *generator) checkFillStubFields(obj *ast.Object) *ast.Object {
	if obj.Decl == nil {
		panic("todo")
	}

	var (
		ts *ast.TypeSpec
		st *ast.StructType
		ok bool
	)

	if ts, ok = obj.Decl.(*ast.TypeSpec); !ok {
		panic("todo")
	}

	if st, ok = ts.Type.(*ast.StructType); !ok {
		panic("todo")
	}

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
		}
	}

	return obj
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
