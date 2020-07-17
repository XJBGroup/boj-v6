package unittest_script

import (
	"github.com/Myriad-Dreamin/boj-v6/lib/unittest/unittest_types"
	"github.com/tidwall/gjson"
)

func ensureVarLength(s []interface{}, wantLen int, err *error) {
	if *err != nil {
		return
	}
	if len(s) != wantLen {
		panic("var length error")
	}
}

func ensureJSONBody(req *unittest_types.Response, err *error) (body gjson.Result) {
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
