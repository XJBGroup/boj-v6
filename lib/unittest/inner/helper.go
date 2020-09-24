package inner

import (
	"github.com/Myriad-Dreamin/boj-v6/lib/unittest/unittest_script"
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
	"reflect"
)

func DotJoin(u, v string) string {
	if len(u) == 0 {
		return v
	}
	if len(v) == 0 {
		return u
	}
	return u + "." + v
}

func DeepCopy(v interface{}) interface{} {
	switch v := v.(type) {
	case map[string]interface{}:
		var nv = make(map[string]interface{})
		for k, vv := range v {
			nv[k] = DeepCopy(vv)
		}
		return nv
	default:
		t := reflect.TypeOf(v)
		if t.Kind() == reflect.Slice || t.Kind() == reflect.Array {
			xv := reflect.ValueOf(v)
			var l = xv.Len()
			var nv = reflect.MakeSlice(t, l, l)
			var values []reflect.Value
			for i := 0; i < l; i++ {
				values = append(values, reflect.ValueOf(DeepCopy(xv.Index(i).Interface())))
			}
			nv = reflect.Append(nv, values...)
			return nv.Interface()
		}
		return v
	}
}

func InheritMapType(dst map[string]interface{}, src map[string]interface{}) {
	for k, v := range src {
		switch v := v.(type) {
		case map[string]interface{}:
			if dv, ok := dst[k]; !ok {
				dst[k] = DeepCopy(v)
			} else if dv, ok := dv.(map[string]interface{}); ok {
				InheritMapType(dv, v)
			}
		default:
			if reflect.TypeOf(v).Kind() == reflect.Array {
				dst[k] = DeepCopy(v)
			}
			if _, ok := dst[k]; !ok {
				dst[k] = v
			}
		}
	}
}

func ToDataBody(v interface{}) interface{} {
	switch v := v.(type) {
	case map[interface{}]interface{}:
		var nv = make(map[string]interface{})
		for kk, vv := range v {
			k := kk.(string)
			if len(k) != 0 && k[0] == '$' {
				k = k[1:]
				if vvv, ok := vv.(string); ok {
					nv[k] = sugar.HandlerError(unittest_script.Eval(
						unittest_script.ResultEvalContext{}, vvv))
				}
				continue
			}

			nv[k] = ToDataBody(vv)
		}
		return nv
	case map[string]interface{}:
		var nv = make(map[string]interface{})
		for k, vv := range v {
			if len(k) != 0 && k[0] == '$' {
				k = k[1:]
				if vvv, ok := vv.(string); ok {
					nv[k] = sugar.HandlerError(unittest_script.Eval(
						unittest_script.ResultEvalContext{}, vvv))
				}
			}

			nv[k] = ToDataBody(vv)
		}
		return nv
	default:
		return v
	}
}
