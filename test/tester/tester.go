package tester

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Myriad-Dreamin/boj-v6/api"
	"github.com/Myriad-Dreamin/boj-v6/pkg/server"
	"github.com/Myriad-Dreamin/minimum-lib/mock"
	"github.com/Myriad-Dreamin/minimum-lib/rbac"
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
	"io"
	"log"
	"strconv"
	"testing"
)

type Tester struct {
	*server.Mocker

	ContextVars   map[string]interface{}
	identityToken map[string]string
}

type Context struct {
	*server.MockerContext
	t *testing.T
	sugar.HandlerErrorLogger
}

func (tester Tester) Set(k string, v interface{}) interface{} {
	res, _ := tester.ContextVars[k]
	tester.ContextVars[k] = v
	return res
}

func (tester Tester) Get(k string) interface{} {
	return tester.ContextVars[k]
}

func (tester Tester) ShouldGet(k string) (v interface{}, ok bool) {
	v, ok = tester.ContextVars[k]
	return
}

func (tester Tester) MustGet(k string) interface{} {
	v, ok := tester.ContextVars[k]
	if !ok {
		panic(fmt.Errorf("could not get %v from context", k))
	}
	return v
}

func StartTester(serverOptions []server.Option) (tester *Tester) {
	tester = new(Tester)
	tester.ContextVars = make(map[string]interface{})
	tester.identityToken = make(map[string]string)
	tester.Mocker = server.Mock(serverOptions...)
	if tester.Mocker == nil {
		panic(errors.New("req mocker error"))
	}
	return tester
}

func (tester *Tester) Context(tt *testing.T) (s *Context) {
	return &Context{
		MockerContext:      tester.Mocker.Context(tt),
		t:                  tt,
		HandlerErrorLogger: sugar.NewHandlerErrorLogger(tt),
	}
}

func (t *Context) AssertNoError(noErr bool) *Context {
	t.MockerContext = t.MockerContext.AssertNoError(noErr)
	return t
}

type ErrorObject struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

func (t *Context) DecodeJSON(body io.Reader, req interface{}) interface{} {
	if err := json.NewDecoder(body).Decode(req); err != nil {
		t.t.Fatal(err)
	}
	return req
}

func (tester *Tester) Release() {
	tester.Mocker.ReleaseMock()
}

func (tester *Tester) MakeAdminContext() bool {
	resp := tester.Post("/v1/user/register", api.RegisterRequest{
		UserName: "admin_context",
		Password: "Admin12345678",
		NickName: "admin_context",
		//Phone:    "1234567891011",
	}, mock.Comment("admin register for test"))
	if !tester.NoErr(resp) {
		return false
	}

	var r api.RegisterReply
	err := resp.JSON(&r)
	if err != nil {
		log.Fatal(err)
		return false
	}
	resp = tester.Post("/v1/user/login",
		api.LoginUserRequest{
			Id:       r.Data.Id,
			Password: "Admin12345678",
		}, mock.Comment("admin login for test"))
	if !tester.NoErr(resp) {
		return false
	}

	var r2 api.LoginUserReply
	err = resp.JSON(&r2)
	if err != nil {
		log.Fatal(err)
		return false
	}

	//fmt.Println(r2)
	//r2.RefreshToken
	_, err = rbac.AddGroupingPolicy("user:"+strconv.Itoa(int(r2.Data.User.ID)), "admin")
	if err != nil {
		tester.Logger.Debug("update group error", "error", err)
	}
	fmt.Println("QAQQQ", rbac.GetPolicy())
	fmt.Println("QAQQQ", rbac.GetGroupingPolicy())
	tester.UseToken(r2.Data.Token)
	tester.identityToken["admin"] = r2.Data.Token
	return true
}

func (tester *Tester) MainM(m *testing.M) {
	tester.Main(func() {
		m.Run()
	})
}

func (tester *Tester) Main(doSomething func()) {
	defer func() {
		if err := recover(); err != nil {
			sugar.PrintStack()
			tester.Logger.Error("panic", "error", err)
		}
		tester.Release()
	}()
	if !tester.MakeAdminContext() {
		return
	}
	doSomething()
}

type GoStyleTestFunc func(*testing.T)
type MinimumStyleTestFunc func(ctx *Context)

type Identity = string

const (
	IdentifyDefault = ""
	IdentityNoAuth  = "no-auth"
	IdentityNormal  = "normal"
	IdentityAdmin   = "admin"
)

type Option struct {
	NoError  bool
	Identity Identity
}

func NewOption() *Option {
	return &Option{NoError: true}
}

func (o *Option) AssertWithoutError(e bool) *Option {
	o.NoError = e
	return o
}

func (o *Option) WithIdentity(identity string) *Option {
	o.Identity = identity
	return o
}

func (*Tester) NewOption() *Option {
	return &Option{NoError: true}
}

func (tester *Tester) HandleTest(testFunc MinimumStyleTestFunc, option *Option) GoStyleTestFunc {
	return func(t *testing.T) {
		ctx := tester.Context(t)
		ctx = ctx.AssertNoError(option.NoError)
		var tok string
		var ok bool
		if option.Identity != IdentifyDefault {
			tok, ok = tester.GetToken()
			if option.Identity == IdentityNoAuth {
				tester.RemoveToken()
			} else {
				tester.UseToken(tester.identityToken[option.Identity])
			}
		}

		testFunc(ctx)

		if option.Identity != IdentifyDefault {
			if ok {
				tester.UseToken(tok)
			} else {
				tester.RemoveToken()
			}
		}
	}
}

func (tester *Tester) HandleTestWithoutError(testFunc MinimumStyleTestFunc) GoStyleTestFunc {
	return tester.HandleTest(testFunc, &Option{NoError: true})
}

func (tester *Tester) HandlePureTest(testFunc MinimumStyleTestFunc) GoStyleTestFunc {
	return tester.HandleTest(testFunc, &Option{NoError: false, Identity: IdentityNoAuth})
}
