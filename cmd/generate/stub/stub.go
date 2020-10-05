package stub

type InvokingStub interface {
	Context(...interface{}) InvokingStub
	Serve(args ...interface{}) Promise
	ServeKeyed(key string, args ...interface{}) Promise
}

type Stub interface {
	InvokingStub

	GetID() uint
	GetIDKeyed(string) uint

	AbortIf(cond bool, args ...interface{})
	AbortIfHint(cond bool, hint int, args ...interface{})
	Bind(request interface{}) Promise

	OnErr(err error, handler func(err error) error, capturing ...interface{})
	//Next() Promise
	//Await(Promise) func(func())
	//Emit(name string, eventArgs ...interface{}) Promise
	//EmitSelf(eventArgs ...interface{}) Promise
}

type StubVariables struct {
	Ok  bool
	Err error

	Int64 int64
	Int int
	Uint64 uint64
	Uint uint
}

type Promise interface {
	Then(func()) Promise
	Catch(func()) Promise
	Finally(func()) Promise
	ThenRef(referableFunc interface{}) Promise
	CatchRef(referableFunc interface{}) Promise
	FinallyRef(referableFunc interface{}) Promise
}
