package tests

import (
	"fmt"
	"github.com/Myriad-Dreamin/boj-v6/api"
	"github.com/Myriad-Dreamin/boj-v6/test/tester"
	"testing"
)

func TestAnnouncement(t *testing.T) {
	_ = t.Run("Post", srv.HandleTestWithOutError(testAnnouncementPost))
	//_ = t.Run("RegisterLogin", srv.HandleTestWithOutError(testUserRegisterLogin)) &&
	//	t.Run("Get", testUserGet) &&
	//	t.Run("Inspect", srv.HandleTestWithOutError(testUserInspect)) &&
	//	t.Run("List", testUserList) &&
	//	t.Run("ChangePassword", srv.HandleTestWithOutError(testUserChangePassword)) &&
	//	t.Run("Put", srv.HandleTestWithOutError(testUserPut)) &&
	//	t.Run("Delete", srv.HandleTestWithOutError(testUserDelete))

}

func testAnnouncementPost(t *tester.TesterContext) {

	resp := t.Post("/v1/announcement", api.PostAnnouncementRequest{
		Title:   "test",
		Content: "test",
	})

	reply := t.DecodeJSON(resp.Body(), new(api.PostAnnouncementReply)).(*api.PostAnnouncementReply)
	fmt.Println(reply)
}
