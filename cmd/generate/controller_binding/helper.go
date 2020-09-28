package main

import (
	"bytes"
	"fmt"
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
	"go/ast"
	"go/build"
	"go/parser"
	"go/printer"
	"go/token"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func (g *generator) printAstNode(i interface{}) {
	sugar.HandlerError0(ast.Print(g.fileSet, i))
}

func (g *generator) printNode(i interface{}) {
	g.printNode_(os.Stdout, i)
}

func (g *generator) printNode_(w io.Writer, i interface{}) {
	sugar.HandlerError0(printer.Fprint(w, g.fileSet, i))
}

func (g *generator) appendStmt(stmt ast.Stmt) {
	var b = bytes.NewBuffer(nil)
	g.printNode_(b, stmt)
	g.methodStmts = append(g.methodStmts, b.String())
}

func (g *generator) formatNode(expr interface{}) string {
	var b = bytes.NewBuffer(nil)
	g.printNode_(b, expr)
	return b.String()
}

func parseDir(fileDir string, fset *token.FileSet, mode parser.Mode) (_ *ast.Package, err error) {
	if fset == nil {
		fset = token.NewFileSet()
	}

	var pkgs map[string]*ast.Package
	pkgs, err = parser.ParseDir(fset, fileDir, nil, mode)
	if err != nil {
		return
	}

	switch len(pkgs) {
	case 3:
		for _, v := range pkgs {
			if strings.HasSuffix(v.Name, "_test") || v.Name == "main" {
				continue
			}
			return v, nil
		}
	case 2:
		for _, v := range pkgs {
			if strings.HasSuffix(v.Name, "_test") {
				continue
			}
			return v, nil
		}
	case 1:
		for _, v := range pkgs {
			return v, nil
		}
	default:
		err = fmt.Errorf("invalid parsed package result: %v", pkgs)
		return
	}

	err = fmt.Errorf("invalid unwrap package result: %v", pkgs)
	return nil, nil
}

func parsePkg(pkgName string, fset *token.FileSet, mode parser.Mode) (_ *ast.Package, err error) {

	srcDir, err := filepath.Abs(pkgName)
	if err != nil {
		return
	}
	pkg, err := build.Import(pkgName, srcDir, 0)
	if err != nil {
		return
	}

	return parseDir(pkg.Dir, fset, mode)
}

func panicNotParsedStmt(node ast.Node) {
	panic(NewGenerateError("np-stmt", node))
}

func notParsedStmt(node ast.Node) error {
	return NewGenerateError("np-stmt", node)
}

func isNotParsedStmt(err error) bool {
	if err, ok := err.(*GenerateError); ok {
		return err.E == "np-stmt"
	}
	return false
}

func panicOnInvoking(lhs []*ast.Ident, numOut int, fnExpr *ast.CallExpr, numIn int) {
	if numIn != -1 && len(fnExpr.Args) != numIn {
		panicGenerateError("invalid get id invoking, want args is "+strconv.Itoa(numIn), fnExpr)
	}

	if numOut != -1 && len(lhs) != numOut {
		panicGenerateError("invalid get id invoking, want lhs is "+strconv.Itoa(numOut), fnExpr)
	}
}
