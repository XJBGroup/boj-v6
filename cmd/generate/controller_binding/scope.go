package main

import "go/ast"

type objectScope struct {
	stubFieldInfo map[string]*ast.Object
}

func (scope *objectScope) getStubFieldByName(fieldName string) *ast.Object {
	return scope.stubFieldInfo[fieldName]
}

type methodScope struct {
	contextVars []*ast.Ident

	methodStmts      []string
	localName        map[string]*ast.Ident
	hasErrorDeclared bool
}

func newStmtScope() *stmtScope {
	return &stmtScope{}
}
