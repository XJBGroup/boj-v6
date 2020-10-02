package main

import (
	"fmt"
	"go/ast"
	"reflect"
	"strings"
)

func (g *generator) autoBindController(ctrl interface{}) {
	defer func() {
		if err := recover(); err != nil {
			err := err.(*GenerateError)
			err.FileSet = g.fileSet
			panic(err)
		}
	}()

	_, t := reflect.ValueOf(ctrl).Elem(), reflect.TypeOf(ctrl).Elem()

	if len(t.PkgPath()) == 0 {
		panic("invalid ctrl, must be pointer of struct")
	}

	parsedPackage, ok := g.packages[t.PkgPath()]
	if !ok {
		panic("package not found")
	}

	n := t.Name()

	obj := g.findInScopes(parsedPackage, n)
	if obj == nil {
		panic("object " + t.String() + " not found")
	}

	// obj.Decl for auto injection

	var on = func(stmt ast.Stmt, err error) {
		if err == nil {
			return
		}

		if isNotParsedStmt(err) {
			g.appendStmt(stmt)
			//panic(err)
			return
		}

		if _, ok := err.(*GenerateError); !ok {
			panic(NewGenerateError(err.Error(), stmt))
		} else {
			panic(err)
		}
	}

	for _, method := range g.getMethods(parsedPackage, obj) {
		g.methodScope.reset()
		for _, stmt := range method.Body.List {
			g.methodScope.methodParsingStmt = stmt
			g.stmtScope.reset()
			//g.stmtScope.stmtParsingNode = stmt
			switch stmt := stmt.(type) {
			case *ast.AssignStmt:
				on(stmt, g.tryParseLRShapeStmt(obj, stmt.Lhs, stmt.Rhs))
			case *ast.DeclStmt:
				switch decl := stmt.Decl.(type) {
				case *ast.GenDecl:
					for _, spec := range decl.Specs {
						switch spec := spec.(type) {
						case *ast.ValueSpec:
							if len(spec.Values) == 0 {
								for _, ident := range spec.Names {
									g.addLocal(ident)
								}
								g.appendStmt(stmt)
								break
							}

							on(stmt, g.tryParseLRShapeStmt_(obj, spec.Names, spec.Values))
						default:
							g.printAstNode(stmt)
							panicGenerateError("not parsed stmt", stmt)
						}
					}
				}
			case *ast.ExprStmt:
				on(stmt, g.tryParseLRShapeStmt_(obj, nil, []ast.Expr{stmt.X}))
			default:
				g.printAstNode(stmt)
				panicGenerateError("not parsed stmt", stmt)
			}
		}
		fmt.Printf("var ctx = new(")
		g.printNode(method.Name)
		fmt.Printf("Context)\n")
		fmt.Println(strings.Join(g.methodStmts, "\n"))
		//fmt.Println(strings.Join())
		break
	}
}
