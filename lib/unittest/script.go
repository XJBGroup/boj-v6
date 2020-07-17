package unittest

import "github.com/alecthomas/participle/lexer"

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

	OpTransfer
)

var operatorMap = map[string]Operator{"+": OpAdd, "-": OpSub, "*": OpMul, "/": OpDiv,
	">=": OpGE, "<=": OpLE, ">": OpGT, "<": OpLT, "==": OpEQ, "!=": OpNEQ,
	"=": OpTransfer}

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

//noinspection GoStructTag
type Ident struct {
	Pos     lexer.Position
	Number  []string     `"$" [ @Ident "." ] @Ident`
	Indices []Expression `[ "[" @@ [ "," @@ ] "]" ]`
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
	Number        Number      `@@`
	String        String      `| @@`
	Boolean       Boolean     `| @@`
	Variable      Ident       ` | @@`
	Subexpression *Expression `| "(" @@ ")"`
	FuncCall      FuncCall    `| @@`
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

//noinspection GoStructTag
type V3 struct {
	Pos   lexer.Position
	Left  Value  `@@`
	Right []OpV2 `{ @@ }`
}

//noinspection GoStructTag
type OpV3 struct {
	Pos      lexer.Position
	Operator Operator `@(">=" | "<=" | "==" | "!=" | "<" | ">")`
	Value    V3       `@@`
}

//noinspection GoStructTag
type V4 struct {
	Pos   lexer.Position
	Left  Value  `@@`
	Right []OpV3 `{ @@ }`
}

//noinspection GoStructTag
type OpV4 struct {
	Pos      lexer.Position
	Operator Operator `@("=")`
	Value    V4       `@@`
}

//noinspection GoStructTag
type Expression struct {
	Pos   lexer.Position
	Left  V4     `@@`
	Right []OpV4 `{ @@ }`
}
