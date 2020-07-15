package unittest

import (
	"github.com/tidwall/gjson"
	"reflect"
)

func dotJoin(u, v string) string {
	if len(u) == 0 {
		return v
	}
	if len(v) == 0 {
		return u
	}
	return u + "." + v
}

func deepCopy(v interface{}) interface{} {
	switch v := v.(type) {
	case map[string]interface{}:
		var nv = make(map[string]interface{})
		for k, vv := range v {
			nv[k] = deepCopy(vv)
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
				values = append(values, reflect.ValueOf(deepCopy(xv.Index(i).Interface())))
			}
			nv = reflect.Append(nv, values...)
			return nv.Interface()
		}
		return v
	}
}

func inheritMapType(dst map[string]interface{}, src map[string]interface{}) {
	for k, v := range src {
		switch v := v.(type) {
		case map[string]interface{}:
			if dv, ok := dst[k]; !ok {
				dst[k] = deepCopy(v)
			} else if dv, ok := dv.(map[string]interface{}); ok {
				inheritMapType(dv, v)
			}
		default:
			if reflect.TypeOf(v).Kind() == reflect.Array {
				dst[k] = deepCopy(v)
			}
			if _, ok := dst[k]; !ok {
				dst[k] = v
			}
		}
	}
}

func toDataBody(v interface{}) interface{} {
	switch v := v.(type) {
	case map[interface{}]interface{}:
		var nv = make(map[string]interface{})
		for k, vv := range v {
			nv[k.(string)] = toDataBody(vv)
		}
		return nv
	default:
		return v
	}
}

func ensureJSONBody(req *Request, err *error) (body gjson.Result) {
	if *err != nil {
		return
	}
	if req.CacheBody != nil {
		body = req.CacheBody.(gjson.Result)
	} else {
		body = gjson.ParseBytes(req.Body)
		if *err != nil {
			req.CacheBody = body
		}
	}
	return
}

func ensureVarLength(s []string, wantLen int, err *error) {
	if *err != nil {
		return
	}
	if len(s) != wantLen {
		panic("var length error")
	}
}
