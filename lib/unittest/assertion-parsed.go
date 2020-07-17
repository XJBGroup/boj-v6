package unittest

import "github.com/Myriad-Dreamin/boj-v6/lib/unittest/unittest_types"

type Statement struct {
	FN    string
	F     unittest_types.CheckFunc
	VArgs []interface{}
}

type TestCase struct {
	Abstract bool
	Path     string
	Name     string
	Meta     map[string]interface{}
	Script   []Statement
}

func (t *TestCase) GetMeta(k string) (v interface{}) {
	return t.Meta[k]
}

type GoDynamicTestData struct {
	TestCases []*TestCase
	Cache     string
}
