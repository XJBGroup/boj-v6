package unittest_script

import (
	"errors"
	fmt "fmt"
	"strings"
)

type Func func([]interface{}) (interface{}, error)
type EvalContext interface {
	EvalVariable(field string) (interface{}, error)
	EvalFunc(functionName string) (Func, error)
}

func Eval(ctx EvalContext, strVal string) (interface{}, error) {
	if len(strVal) == 0 {
		return nil, errors.New("empty eval string")
	}

	if strVal[len(strVal)-1] == ')' { // maybe function
		i := strings.IndexByte(strVal, '(')
		if i == -1 {
			return nil, fmt.Errorf("invalid form of field %v", strVal)
		}
		if i == 0 {
			return Eval(ctx, strVal[1:len(strVal)-1])
		}

		fn, strVal := strings.TrimSpace(strVal[:i]), strings.TrimSpace(strVal[i+1:len(strVal)-1])

		f, err := ctx.EvalFunc(fn)
		if err != nil {
			return nil, err
		}

		if len(strVal) == 0 {
			return f(nil)
		} else { // unary function
			// todo: escape ,
			var valList = strings.Split(strVal, ",")
			var values []interface{}
			for i := range valList {
				value, err := Eval(ctx, strings.TrimSpace(valList[i]))
				if err != nil {
					return nil, err
				}

				values = append(values, value)
			}

			return f(values)
		}
	}

	// todo: not variable

	v, err := ctx.EvalVariable(strVal)
	if err != nil {
		return nil, err
	}

	return v, nil
}
