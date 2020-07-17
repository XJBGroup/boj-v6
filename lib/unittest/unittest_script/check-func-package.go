package unittest_script

import (
	"fmt"
	"github.com/Myriad-Dreamin/boj-v6/lib/unittest/unittest_types"
	"math"
	"reflect"
	"strconv"
	"strings"
)

func composed(s string) (string, []string) {
	if len(s) > 0 && s[len(s)-1] == ')' {
		i := strings.IndexByte(s, '(')
		if i == -1 {
			panic(fmt.Errorf("invalid form of field %v", s))
		}
		f, s := s[:i], s[i+1:len(s)-1]
		s, fs := composed(s)
		return s, append(fs, f)
	}
	return s, nil
}

var ft = reflect.TypeOf(float64(1))

func convertValue(ref interface{}, v interface{}) (wv interface{}, err error) {
	switch ref.(type) {
	case float64:
		switch v := v.(type) {
		case float64:
			wv = v
		case int, int64, int32, int16, int8, uint, uint64, uint32, uint16, uint8:
			wv = reflect.ValueOf(v).Convert(ft).Interface()
		case string:
			wv, err = strconv.ParseFloat(v, 64)
			if err != nil {
				return nil, fmt.Errorf("convert error: %v", err)
			}
		default:
			err = fmt.Errorf("unknown convert value from %T to %T", v, ref)
		}
	case bool:
		switch v := v.(type) {
		case bool:
			wv = v
		case string:
			wv, err = strconv.ParseBool(v)
			if err != nil {
				return nil, fmt.Errorf("convert error: %v", err)
			}
		default:
			err = fmt.Errorf("unknown convert value from %T to %T", v, ref)
		}
	case string:
		wv = fmt.Sprintf("%v", v)
	case nil:
		switch v := v.(type) {
		case nil:
			wv = nil
		case string:
			if v != "nil" {
				wv = struct{}{}
			} else {
				wv = nil
			}
		default:
			err = fmt.Errorf("unknown convert value from %T to %T", v, ref)
		}
	default:
		return nil, fmt.Errorf("bad assertion type: %T", ref)
	}
	return
}

var EQFunctions = map[reflect.Type]func(u, v interface{}) bool{
	reflect.TypeOf(float64(1)): func(u, v interface{}) bool {
		return math.Abs(v.(float64)-u.(float64)) <= 1e-6
	},
	reflect.TypeOf(""):   func(u, v interface{}) bool { return u == v },
	reflect.TypeOf(true): func(u, v interface{}) bool { return u == v },
	reflect.TypeOf(nil):  func(u, v interface{}) bool { return u == v },
}

func applyFunc(value interface{}, fs []string) (interface{}, error) {
	for i := range fs {
		switch fs[i] {
		case "len":
			v := reflect.ValueOf(value)
			switch v.Kind() {
			case reflect.Array, reflect.Slice, reflect.Map, reflect.String:
				value = float64(v.Len())
			default:
				return nil, fmt.Errorf("could not perform len on %v(%T)", value, value)
			}
		}
	}
	return value, nil
}

func assertJSONEQ(st *unittest_types.State, s2 ...interface{}) (s bool, err error) {
	ensureVarLength(s2, 2, &err)
	field, fs := composed(s2[0].(string))
	if body := ensureJSONBody(st.Res, &err); err == nil {
		k, err := applyFunc(body.Get(field).Value(), fs)
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
	field, fs := composed(s2[0].(string))
	if body := ensureJSONBody(st.Res, &err); err == nil {
		k, err := applyFunc(body.Get(field).Value(), fs)
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
