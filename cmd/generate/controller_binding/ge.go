package main

import (
	"fmt"
	"go/ast"
	"go/token"
)

type GenerateError struct {
	E       string
	Node    ast.Node
	FileSet *token.FileSet
}

func NewGenerateError(err string, node ast.Node) *GenerateError {
	return &GenerateError{
		E:    err,
		Node: node,
	}
}

func (g *GenerateError) Error() string {
	if g.FileSet != nil {
		return fmt.Sprintf("%v at pos %v", g.E, g.FileSet.Position(g.Node.Pos()))
	}
	return fmt.Sprintf("%v at pos %v", g.E, g.Node.Pos())
}

func panicUnknownAt(node ast.Node) {
	panic(NewGenerateError("unknown func type, maybe it is a bug", node))
}

func unknownAt(node ast.Node) error {
	panicUnknownAt(node)
	return NewGenerateError("unknown func type, maybe it is a bug", node)
}

func panicGenerateError(err string, node ast.Node) {
	panic(NewGenerateError(err, node))
}

func panicGenerateError_(err string) {
	panic(NewGenerateError(err, nil))
}
