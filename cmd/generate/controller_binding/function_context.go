package main

import "go/ast"

type stmtScope struct {
	stmtParsingNode ast.Node
}

func (c *stmtScope) reset() {
	c.stmtParsingNode = nil
}
