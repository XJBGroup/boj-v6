package unittest_types

type Response struct {
	Body      []byte
	CacheBody interface{}
}

type State struct {
	Res *Response
}
type CheckFunc = func(*State, ...interface{}) (bool, error)
type Package = map[string]CheckFunc

type LinkedContext interface {
	Last() LinkedContext
	Name() string
	Get(s string) LinkedContext
	GetFunc(s string) CheckFunc
	GetFunctions(func(s string, f CheckFunc) error) error
	SetName(s string)
	SetLast(LinkedContext)
	Insert(k string, ctx LinkedContext)
	InsertFunc(s string, f CheckFunc)
}

type DataBody = map[string]interface{}
