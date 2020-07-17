package unittest

type DataBody = map[string]interface{}

type Request struct {
	Body      []byte
	CacheBody interface{}
}

type Assertion struct {
	FN    string
	F     CheckFunc
	VArgs []interface{}
}

type TestCase struct {
	Abstract bool
	Path     string
	Name     string
	Meta     map[string]interface{}
	Script   []Assertion
}

func (t *TestCase) GetMeta(k string) (v interface{}) {
	return t.Meta[k]
}

type GoDynamicTestData struct {
	TestCases []*TestCase
	Cache     string
}
