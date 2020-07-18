package unittest

import (
	"fmt"
	"github.com/Myriad-Dreamin/boj-v6/lib/unittest/unittest_types"
	"github.com/alecthomas/participle/lexer"
)

func assignLValue(dst, src interface{}) error {
	panic("todo")
}

//noinspection GoStructTag
type ScriptLine struct {
	Expression Expression `@@`
}

type Operator int

const (
	OpMul Operator = iota
	OpDiv
	OpAdd
	OpSub

	OpGT
	OpLT
	OpGE
	OpLE
	OpEQ
	OpNEQ
	OpLAnd
	OpLOr

	OpAssign
)

var operatorMap = map[string]Operator{"+": OpAdd, "-": OpSub, "*": OpMul, "/": OpDiv,
	">=": OpGE, "<=": OpLE, ">": OpGT, "<": OpLT, "==": OpEQ, "!=": OpNEQ,
	"&&": OpLAnd, "||": OpLOr,
	"=": OpAssign}

func (o *Operator) Capture(s []string) error {
	*o = operatorMap[s[0]]
	return nil
}

//noinspection GoStructTag
type FuncCall struct {
	Pos    lexer.Position
	Callee []string     `[ @Ident "." ] @Ident`
	Args   []Expression `[ "(" @@ [ "," @@ ] ")" ]`
}

func (c FuncCall) Eval(st *unittest_types.State) (interface{}, error) {
	panic("todo")
}

//noinspection GoStructTag
type Ident struct {
	Pos     lexer.Position
	Number  []string     `"$" [ @Ident "." ] @Ident`
	Indices []Expression `[ "[" @@ [ "," @@ ] "]" ]`
}

func (i Ident) Eval(st *unittest_types.State) (interface{}, error) {
	panic("todo")
}

func (i Ident) GetLValue(st *unittest_types.State) (interface{}, error) {
	panic("todo")
}

//noinspection GoStructTag
type Number struct {
	Pos    lexer.Position
	Number float64 `@(Float|Int)`
}

//noinspection GoStructTag
type String struct {
	Pos    lexer.Position
	String string `@String`
}

//noinspection GoStructTag
type Boolean struct {
	Pos     lexer.Position
	Boolean bool `@ ("true" | "false")`
}

//noinspection GoStructTag
type Value struct {
	Number        *Number     `@@`
	String        *String     `| @@`
	Boolean       *Boolean    `| @@`
	Variable      *Ident      ` | @@`
	Subexpression *Expression `| "(" @@ ")"`
	FuncCall      *FuncCall   `| @@`
}

func (exp *Value) GetLValue(st *unittest_types.State) (v interface{}, err error) {
	switch {
	case exp.Number != nil:
		fallthrough
	case exp.String != nil:
		fallthrough
	case exp.Boolean != nil:
		fallthrough
	case exp.FuncCall != nil:
		err = fmt.Errorf("has no lvalue")
	case exp.Variable != nil:
		v, err = exp.Variable.GetLValue(st)
	case exp.Subexpression != nil:
		v, err = exp.Subexpression.GetLValue(st)
	default:
		err = fmt.Errorf("bad switch value case")
	}
	return
}

func (exp *Value) Eval(st *unittest_types.State) (v interface{}, err error) {
	switch {
	case exp.Number != nil:
		v = exp.Number.Number
	case exp.String != nil:
		v = exp.String.String
	case exp.Boolean != nil:
		v = exp.Boolean.Boolean
	case exp.Variable != nil:
		v, err = exp.Variable.Eval(st)
	case exp.Subexpression != nil:
		v, err = exp.Subexpression.Eval(st)
	case exp.FuncCall != nil:
		v, err = exp.FuncCall.Eval(st)
	default:
		err = fmt.Errorf("bad switch value case")
	}
	return
}

//noinspection GoStructTag
type OpV1 struct {
	Pos      lexer.Position
	Operator Operator `@("*" | "/")`
	Value    Value    `@@`
}

//noinspection GoStructTag
type V2 struct {
	Pos   lexer.Position
	Left  Value  `@@`
	Right []OpV1 `{ @@ }`
}

//noinspection GoStructTag
type OpV2 struct {
	Pos      lexer.Position
	Operator Operator `@("+" | "-")`
	Value    V2       `@@`
}

func (exp *V2) GetLValue(st *unittest_types.State) (interface{}, error) {
	if len(exp.Right) != 0 {
		return nil, fmt.Errorf("not factor")
	}
	return exp.Left.GetLValue(st)
}

func (exp *V2) Eval(st *unittest_types.State) (interface{}, error) {
	lv, err := exp.Left.Eval(st)
	if err != nil {
		return nil, err
	}
	for _, rExp := range exp.Right {
		rv, err := rExp.Value.Eval(st)
		if err != nil {
			return nil, err
		}
		_ = rv
		switch rExp.Operator {
		case OpMul:
			return lv.(float64) * rv.(float64), nil
		case OpDiv:
			return lv.(float64) / rv.(float64), nil
		default:
			return nil, fmt.Errorf("unknown operator")
		}
	}
	return lv, nil
}

//noinspection GoStructTag
type V3 struct {
	Pos   lexer.Position
	Left  V2     `@@`
	Right []OpV2 `{ @@ }`
}

//noinspection GoStructTag
type OpV3 struct {
	Pos      lexer.Position
	Operator Operator `@(">=" | "<=" | "==" | "!=" | "<" | ">")`
	Value    V3       `@@`
}

func (exp *V3) GetLValue(st *unittest_types.State) (interface{}, error) {
	if len(exp.Right) != 0 {
		return nil, fmt.Errorf("not factor")
	}
	return exp.Left.GetLValue(st)
}

func (exp *V3) Eval(st *unittest_types.State) (interface{}, error) {
	lv, err := exp.Left.Eval(st)
	if err != nil {
		return nil, err
	}
	for _, rExp := range exp.Right {
		rv, err := rExp.Value.Eval(st)
		if err != nil {
			return nil, err
		}
		switch rExp.Operator {
		case OpAdd:
			return lv.(float64) + rv.(float64), nil
		case OpSub:
			return lv.(float64) - rv.(float64), nil
		default:
			return nil, fmt.Errorf("unknown operator")
		}
	}
	return lv, nil
}

//noinspection GoStructTag
type V4 struct {
	Pos   lexer.Position
	Left  V3     `@@`
	Right []OpV3 `{ @@ }`
}

//noinspection GoStructTag
type OpV4 struct {
	Pos      lexer.Position
	Operator Operator `@("&&" | "||")`
	Value    V4       `@@`
}

func (exp *V4) GetLValue(st *unittest_types.State) (interface{}, error) {
	if len(exp.Right) != 0 {
		return nil, fmt.Errorf("not factor")
	}
	return exp.Left.GetLValue(st)
}

func (exp *V4) Eval(st *unittest_types.State) (interface{}, error) {
	lv, err := exp.Left.Eval(st)
	if err != nil {
		return nil, err
	}
	for _, rExp := range exp.Right {
		rv, err := rExp.Value.Eval(st)
		if err != nil {
			return nil, err
		}
		_ = rv
		switch rExp.Operator {
		case OpLE:
			return lv.(float64) <= rv.(float64), nil
		case OpGE:
			return lv.(float64) >= rv.(float64), nil
		case OpLT:
			return lv.(float64) < rv.(float64), nil
		case OpGT:
			return lv.(float64) > rv.(float64), nil
		case OpEQ:
			return lv.(float64) == rv.(float64), nil
		case OpNEQ:
			return lv.(float64) != rv.(float64), nil
		default:
			return nil, fmt.Errorf("unknown operator")
		}
	}
	return lv, nil
}

//noinspection GoStructTag
type V5 struct {
	Pos   lexer.Position
	Left  V4     `@@`
	Right []OpV4 `{ @@ }`
}

//noinspection GoStructTag
type OpV5 struct {
	Pos      lexer.Position
	Operator Operator `@("=")`
	Value    V5       `@@`
}

func (exp *V5) GetLValue(st *unittest_types.State) (interface{}, error) {
	if len(exp.Right) != 0 {
		return nil, fmt.Errorf("not factor")
	}
	return exp.Left.GetLValue(st)
}

func (exp *V5) Eval(st *unittest_types.State) (interface{}, error) {
	lv, err := exp.Left.Eval(st)
	if err != nil {
		return nil, err
	}
	for _, rExp := range exp.Right {
		rv, err := rExp.Value.Eval(st)
		if err != nil {
			return nil, err
		}
		_ = rv
		switch rExp.Operator {
		case OpLAnd:
			panic("todo")
		case OpLOr:
			panic("todo")
		default:
			return nil, fmt.Errorf("unknown operator")
		}
	}
	return lv, nil
}

//noinspection GoStructTag
type Expression struct {
	Pos   lexer.Position
	Left  V5     `@@`
	Right []OpV5 `{ @@ }`
}

func (exp *Expression) GetLValue(st *unittest_types.State) (interface{}, error) {
	if len(exp.Right) != 0 {
		return nil, fmt.Errorf("not factor")
	}
	return exp.Left.GetLValue(st)
}

func (exp *Expression) Eval(st *unittest_types.State) (interface{}, error) {
	if len(exp.Right) == 0 {
		return exp.Left.Eval(st)
	}
	lv, err := exp.Left.GetLValue(st)
	if err != nil {
		return nil, err
	}
	for _, rExp := range exp.Right {
		switch rExp.Operator {
		case OpAssign:
		default:
			return nil, fmt.Errorf("unknown operator")
		}
	}
	x, y := exp.Right[:len(exp.Right)-1], exp.Right[len(exp.Right)-1]
	rv, err := y.Value.Eval(st)
	if err != nil {
		return nil, err
	}
	err = assignLValue(lv, rv)
	if err != nil {
		return nil, err
	}
	for _, rExp := range x {
		lv, err := rExp.Value.GetLValue(st)
		if err != nil {
			return nil, err
		}
		err = assignLValue(lv, rv)
		if err != nil {
			return nil, err
		}
	}
	return rv, nil
}
