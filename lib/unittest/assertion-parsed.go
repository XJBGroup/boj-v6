package unittest

import "github.com/Myriad-Dreamin/boj-v6/lib/unittest/unittest_types"

type Statement struct {
	FN    string                   `json:"func_name"`
	F     unittest_types.CheckFunc `json:"-"`
	VArgs []interface{}            `json:"args"`
}

type TestCase struct {
	Abstract bool                   `json:"abstract"`
	Path     string                 `json:"path"`
	Name     string                 `json:"name"`
	Meta     map[string]interface{} `json:"meta"`
	Script   []Statement            `json:"scripts"`
}

func (t *TestCase) GetMeta(k string) (v interface{}) {
	return t.Meta[k]
}

type GoDynamicTestData struct {
	TestCases []*TestCase
	Cache     string
}
