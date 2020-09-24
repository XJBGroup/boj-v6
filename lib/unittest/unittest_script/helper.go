package unittest_script

import (
	"fmt"
	"reflect"
	"strconv"
)

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
