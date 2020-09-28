package stub

type InvokingStub interface {
	Context(...interface{}) InvokingStub
	Serve(args ...interface{}) Promise
	ServeKeyed(args ...interface{}) Promise
}

type Stub interface {
	InvokingStub

	GetID() *uint
	GetIDKeyed(string) *uint
	AbortIf(bool)
	Bind(request interface{}) Promise

	Await(Promise) func(func())
	Emit(name string, eventArgs ...interface{})
	EmitSelf(eventArgs ...interface{})
}

type StubVariables struct {
	Ok  bool
	Err error
}

type Promise interface {
	Stub

	Then(func()) Promise
	Catch(func()) Promise
	ThenDo(f interface{}) Promise
	CatchDo(f interface{}) Promise
}
