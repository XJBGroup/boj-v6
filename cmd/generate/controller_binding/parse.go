package main

type ParseFlag int64

const (
	ParseTestFiles ParseFlag = 1 << iota
)

func (g *generator) parse(opts ...interface{}) {
	g.parseAllImports(opts)
}
