package deepcopy

import "reflect"

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
		switch t.Kind() {
		case reflect.Array, reflect.Slice:
			xv := reflect.ValueOf(v)
			var l = xv.Len()
			var nv = reflect.MakeSlice(t, l, l)
			var values []reflect.Value
			for i := 0; i < l; i++ {
				values = append(values, reflect.ValueOf(DeepCopy(xv.Index(i).Interface())))
			}
			nv = reflect.Append(nv, values...)
			return nv.Interface()
		case reflect.Map:
			var (
				xv   = reflect.ValueOf(v)
				mv   = reflect.MakeMap(t)
				iter = xv.MapRange()
			)
			for iter.Next() {
				k, v := iter.Key(), iter.Value()
				mv.SetMapIndex(k, reflect.ValueOf(DeepCopy(v.Interface())))
			}
			return mv.Interface()
		}

		return v
	}
}
