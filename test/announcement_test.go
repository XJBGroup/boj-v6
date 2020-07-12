package tests

import (
	"github.com/Myriad-Dreamin/boj-v6/api"
	"github.com/Myriad-Dreamin/boj-v6/test/tester"
	"github.com/Myriad-Dreamin/boj-v6/types"
	"strconv"
	"testing"
)

func TestAnnouncement(t *testing.T) {
	_ = t.Run("Post", srv.HandleTestWithoutError(testAnnouncementPost)) &&
		t.Run("Get", srv.HandleTestWithoutError(testAnnouncementGet))
	//	t.Run("Inspect", srv.HandleTestWithoutError(testAnnouncementInspect)) &&
	//	t.Run("List", testAnnouncementList) &&
	//	t.Run("ChangePassword", srv.HandleTestWithoutError(testAnnouncementChangePassword)) &&
	//	t.Run("Put", srv.HandleTestWithoutError(testAnnouncementPut)) &&
	//	t.Run("Delete", srv.HandleTestWithoutError(testAnnouncementDelete))

}

func testAnnouncementGet(t *tester.Context) {

	resp := t.Get("/v1/announcement/" + strconv.Itoa(int(srv.Get(AnnouncementIdKey).(uint))))

	reply := t.DecodeJSON(resp.Body(), new(api.GetAnnouncementReply)).(*api.GetAnnouncementReply)
	t.Equal(types.CodeOK, reply.Code)
	t.Equal(srv.Get(AnnouncementIdKey).(uint), reply.Announcement.ID)
	t.Equal("title", reply.Announcement.Title)
	t.Equal("content", reply.Announcement.Content)
}

func testAnnouncementPost(t *tester.Context) {

	resp := t.Post("/v1/announcement", api.PostAnnouncementRequest{
		Title:   "title",
		Content: "content",
	})

	reply := t.DecodeJSON(resp.Body(), new(api.PostAnnouncementReply)).(*api.PostAnnouncementReply)
	t.Equal(types.CodeOK, reply.Code)
	srv.Set(AnnouncementIdKey, reply.Announcement.ID)
}
