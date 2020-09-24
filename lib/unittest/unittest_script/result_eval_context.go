package unittest_script

import (
	"errors"
	"github.com/tidwall/gjson"
)

type ResultEvalContext struct {
	Body *gjson.Result
}

func (c ResultEvalContext) EvalVariable(field string) (interface{}, error) {
	return c.Body.Get(field).Value(), nil
}

func (c ResultEvalContext) EvalFunc(fn string) (Func, error) {
	f, ok := functions[fn]
	if !ok {
		return nil, errors.New("function " + fn + " not found")
	}
	return f, nil
}
