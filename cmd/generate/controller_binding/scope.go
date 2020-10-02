package main

import (
	"fmt"
	"go/ast"
)

type objectScope struct {
	stubFieldInfo map[string]*ast.Object
}

func (scope *objectScope) getStubFieldByName(fieldName string) *ast.Object {
	return scope.stubFieldInfo[fieldName]
}

type methodScope struct {
	methodParsingStmt ast.Stmt
	methodStmts       []string
	contextVars       map[string]*ast.Ident
	localName         map[string]*ast.Ident
	hasErrorDeclared  bool
	hasOkDeclared     bool
}

func (g *methodScope) reset() {
	g.methodParsingStmt = nil
	g.contextVars = nil
	g.methodStmts = nil
	g.localName = nil
	g.hasOkDeclared = false
	g.hasErrorDeclared = false
}

func (g *methodScope) addContext(ident *ast.Ident) {
	if pIdent, ok := g.contextVars[ident.Name]; ok {
		if ident.Obj == pIdent.Obj && ident.Name == pIdent.Name {
			return
		}

		panicGenerateError(fmt.Sprintf("conflict on context binding %v", ident.Name), ident)
		return
	}
	if g.contextVars == nil {
		g.contextVars = make(map[string]*ast.Ident)
	}
	g.contextVars[ident.Name] = ident
}

func (g *methodScope) addLocal(ident *ast.Ident) {
	if g.localName == nil {
		g.localName = make(map[string]*ast.Ident)
	}
	g.localName[ident.Name] = ident
}

func newStmtScope() *stmtScope {
	return &stmtScope{}
}
