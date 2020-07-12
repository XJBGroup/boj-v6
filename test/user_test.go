package tests

import (
	"github.com/Myriad-Dreamin/boj-v6/api"
	"github.com/Myriad-Dreamin/boj-v6/test/tester"
	"github.com/Myriad-Dreamin/boj-v6/types"
	"strconv"
	"testing"
)

type referUserService = api.UserService

func TestUser(t *testing.T) {
	_ = t.Run("Register", srv.HandlePureTest(testUserRegister)) &&
		t.Run("Get", srv.HandlePureTest(testUserGet))
	//	t.Run("Inspect", srv.HandleTestWithoutError(testUserInspect)) &&
	//	t.Run("List", testUserList) &&
	//	t.Run("ChangePassword", srv.HandleTestWithoutError(testUserChangePassword)) &&
	//	t.Run("Put", srv.HandleTestWithoutError(testUserPut)) &&
	//	t.Run("Delete", srv.HandleTestWithoutError(testUserDelete))

}

func testUserGet(t *tester.Context) {

	getFunc := func(id int) *api.GetUserReply {
		return t.DecodeJSON(t.Get("/v1/user/"+strconv.Itoa(id)).Body(), new(api.GetUserReply)).(*api.GetUserReply)
	}

	reply := getFunc(1)
	t.Equal(types.CodeOK, reply.Code)
}

func testUserRegister(t *tester.Context) {

	registerFunc := func(rr *api.RegisterRequest) *api.RegisterReply {
		return t.DecodeJSON(t.Post("/v1/user/register", rr).Body(), new(api.RegisterReply)).(*api.RegisterReply)
	}

	t.Equal(types.CodeOK, registerFunc(&api.RegisterRequest{
		UserName: "test0",
		Password: "test0",
		NickName: "test0",
		Gender:   0,
	}).Code)

	t.Equal(types.CodeUniqueConstraintFailed, registerFunc(&api.RegisterRequest{
		UserName: "test0",
		Password: "test0",
		NickName: "test0",
		Gender:   0,
	}).Code)

	t.Equal(types.CodeUniqueConstraintFailed, registerFunc(&api.RegisterRequest{
		UserName: "test1",
		Password: "test0",
		NickName: "test0",
		Gender:   0,
	}).Code)

	t.Equal(types.CodeUniqueConstraintFailed, registerFunc(&api.RegisterRequest{
		UserName: "test0",
		Password: "test0",
		NickName: "test1",
		Gender:   0,
	}).Code)

}
