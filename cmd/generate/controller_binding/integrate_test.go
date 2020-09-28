package main

import (
	"github.com/Myriad-Dreamin/boj-v6/cmd/generate/controller_binding/inner/model"
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestInnerTest(t *testing.T) {
	var fileSet = token.NewFileSet()

	pkgName := "github.com/Myriad-Dreamin/boj-v6/cmd/generate/controller_binding/inner/model"

	pkg, err := parsePkg(pkgName, fileSet, parser.ParseComments|parser.DeclarationErrors)
	sugar.HandlerError0(err)

	g := &generator{fileSet: fileSet, packages: map[string]*ast.Package{
		pkgName: pkg,
	}}

	g.parse(g.resolveBuiltIn(), StubPackage("github.com/Myriad-Dreamin/boj-v6/cmd/generate/stub"))

	g.autoBindController(new(model.Sc))
}
