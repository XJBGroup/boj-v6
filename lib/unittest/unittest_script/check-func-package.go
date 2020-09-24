package unittest_script

import (
	"fmt"
	"github.com/Myriad-Dreamin/boj-v6/lib/unittest/unittest_types"
	"math"
	"reflect"
	"time"
)

var ft = reflect.TypeOf(float64(1))
var EQFunctions = map[reflect.Type]func(u, v interface{}) bool{
	reflect.TypeOf(float64(1)): func(u, v interface{}) bool {
		return math.Abs(v.(float64)-u.(float64)) <= 1e-6
	},
	reflect.TypeOf(""):   func(u, v interface{}) bool { return u == v },
	reflect.TypeOf(true): func(u, v interface{}) bool { return u == v },
	reflect.TypeOf(nil):  func(u, v interface{}) bool { return u == v },
}

var functions = map[string]Func{
	"len": func(values []interface{}) (value interface{}, err error) {
		ensureVarLength(values, 1, &err)
		if err != nil {
			return nil, err
		}
		value = values[0]
		v := reflect.ValueOf(value)
		switch v.Kind() {
		case reflect.Array, reflect.Slice, reflect.Map, reflect.String:
			value = float64(v.Len())
		default:
			return nil, fmt.Errorf("could not perform len on %v(%T)", value, value)
		}
		return value, nil
	},

	"date.now": func(values []interface{}) (value interface{}, err error) {
		ensureVarLength(values, 0, &err)
		if err != nil {
			return nil, err
		}
		value = time.Now()
		return
	},
}

func assertJSONEQ(st *unittest_types.State, s2 ...interface{}) (s bool, err error) {
	ensureVarLength(s2, 2, &err)
	if body := ensureJSONBody(st.Res, &err); err == nil {
		k, err := Eval(ResultEvalContext{&body}, s2[0].(string))
		if err != nil {
			return false, err
		}
		wv, err := convertValue(k, s2[1])
		if err != nil {
			return false, err
		}
		if EQFunctions[reflect.TypeOf(k)](k, wv) == false {
			return false, fmt.Errorf("float assertion equal error: want %v, got %v", wv, k)
		}
		return true, nil
	}
	return
}

func assertJSONNEQ(st *unittest_types.State, s2 ...interface{}) (s bool, err error) {
	ensureVarLength(s2, 2, &err)
	if body := ensureJSONBody(st.Res, &err); err == nil {
		k, err := Eval(ResultEvalContext{&body}, s2[0].(string))
		if err != nil {
			return false, err
		}
		wv, err := convertValue(k, s2[1])
		if err != nil {
			return false, err
		}
		if EQFunctions[reflect.TypeOf(k)](k, wv) == true {
			return false, fmt.Errorf("float assertion not equal error: want %v, got %v", wv, k)
		}
		return true, nil
	}
	return
}

var NamespaceStd = unittest_types.Package{
	"Assert":    assertJSONEQ,
	"AssertEQ":  assertJSONEQ,
	"AssertNEQ": assertJSONNEQ,
	"AssertZeroValue": func(st *unittest_types.State, s2 ...interface{}) (s bool, err error) {
		ensureVarLength(s2, 1, &err)
		if body := ensureJSONBody(st.Res, &err); err == nil {
			fmt.Println("asserting", body)
		}
		return
	},
}

var NamespaceJSON = unittest_types.Package{
	"Assert":    assertJSONEQ,
	"AssertEQ":  assertJSONEQ,
	"AssertNEQ": assertJSONNEQ,
	"AssertZeroValue": func(st *unittest_types.State, s2 ...interface{}) (s bool, err error) {
		ensureVarLength(s2, 1, &err)
		if body := ensureJSONBody(st.Res, &err); err == nil {
			fmt.Println("asserting", body)
		}
		return
	},
}
