package main

import (
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
	"github.com/stretchr/testify/assert"
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func Test_generator_parseImports(t *testing.T) {
	fileSet := token.NewFileSet()
	pkg, err := parseDir(".", fileSet, parser.ParseComments|parser.DeclarationErrors)
	sugar.HandlerError0(err)

	pkg.Files = map[string]*ast.File{
		"parse_imports_test_target_test.go": pkg.Files["parse_imports_test_target_test.go"],
	}

	g := &generator{fileSet: fileSet, packages: map[string]*ast.Package{
		"main": pkg,
	}}

	g.parseAllImports([]interface{}{g.resolveBuiltIn()})

	for _, pkg := range g.packages {
		for _, f := range pkg.Files {
			assert.Len(t, f.Unresolved, 0)
		}
	}
}
