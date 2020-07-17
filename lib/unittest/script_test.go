package unittest

import (
	"fmt"
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
	"github.com/alecthomas/participle"
	"testing"
)

func TestParseScript(t *testing.T) {
	parser, err := participle.Build(&ScriptLine{})
	if err != nil {
		panic(err)
	}
	e := &ScriptLine{}
	sugar.HandlerError0(parser.ParseString("a", e))
	fmt.Println(e)
	sugar.HandlerError0(parser.ParseString("a.a", e))
	fmt.Println(e)
	sugar.HandlerError0(parser.ParseString("a.a(x)", e))
	fmt.Println(e)
	sugar.HandlerError0(parser.ParseString("a.a(x($a.a))", e))
	fmt.Println(e)
	sugar.HandlerError0(parser.ParseString("$a.a", e))
	fmt.Println(e)
	sugar.HandlerError0(parser.ParseString("$a[0]", e))
	fmt.Println(e)
	sugar.HandlerError0(parser.ParseString("$a[0, 1]", e))
	fmt.Println(e)
	sugar.HandlerError0(parser.ParseString("$req.a = $a", e))
	fmt.Println(e)
}
