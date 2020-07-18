package unittest

import (
	"fmt"
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
	"github.com/alecthomas/participle"
	"github.com/go-playground/assert/v2"
	"math"
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

func TestEvalScript(t *testing.T) {
	parser, err := participle.Build(&ScriptLine{})
	if err != nil {
		panic(err)
	}
	var v interface{}
	e := &ScriptLine{}
	//sugar.HandlerError0(parser.ParseString("a", e))
	//fmt.Println(e)

	e = &ScriptLine{}
	//sugar.HandlerError0(parser.ParseString("a.a", e))
	//fmt.Println(e)

	e = &ScriptLine{}
	//sugar.HandlerError0(parser.ParseString("a.a(x)", e))
	//fmt.Println(e)

	e = &ScriptLine{}
	//sugar.HandlerError0(parser.ParseString("a.a(x($a.a))", e))
	//fmt.Println(e)

	e = &ScriptLine{}
	//sugar.HandlerError0(parser.ParseString("$a.a", e))
	//fmt.Println(e)

	e = &ScriptLine{}
	//sugar.HandlerError0(parser.ParseString("$a[0]", e))
	//fmt.Println(e)

	e = &ScriptLine{}
	//sugar.HandlerError0(parser.ParseString("$a[0, 1]", e))
	//fmt.Println(e)

	e = &ScriptLine{}
	//sugar.HandlerError0(parser.ParseString("$req.a = $a", e))
	//fmt.Println(e)

	e = &ScriptLine{}
	sugar.HandlerError0(parser.ParseString("1 + 1", e))
	v, err = e.Expression.Eval(nil)
	assert.Equal(t, nil, err)
	assert.Equal(t, true, math.Abs(2-v.(float64)) <= 1e-6)

	e = &ScriptLine{}
	sugar.HandlerError0(parser.ParseString("3 * 6", e))
	v, err = e.Expression.Eval(nil)
	assert.Equal(t, nil, err)
	assert.Equal(t, true, math.Abs(18-v.(float64)) <= 1e-6)

	e = &ScriptLine{}
	sugar.HandlerError0(parser.ParseString("1 + 2 * 3", e))
	v, err = e.Expression.Eval(nil)
	assert.Equal(t, nil, err)
	assert.Equal(t, true, math.Abs(7-v.(float64)) <= 1e-6)

	e = &ScriptLine{}
	sugar.HandlerError0(parser.ParseString("1 < 2", e))
	v, err = e.Expression.Eval(nil)
	assert.Equal(t, nil, err)
	assert.Equal(t, true, v)
}
