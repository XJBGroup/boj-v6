package comment

import (
	"github.com/Myriad-Dreamin/boj-v6/abstract/announcement"
	"github.com/Myriad-Dreamin/boj-v6/abstract/comment"
	"github.com/Myriad-Dreamin/boj-v6/api"
	"github.com/Myriad-Dreamin/boj-v6/app/provider"
	"github.com/Myriad-Dreamin/boj-v6/config"
	ginhelper "github.com/Myriad-Dreamin/boj-v6/lib/gin-helper"
	"github.com/Myriad-Dreamin/core-oj/log"
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	"github.com/Myriad-Dreamin/minimum-lib/module"
	"net/http"
)

type Service struct {
	db     comment.DB
	logger log.TendermintLogger
	key    string
}

func (srv *Service) CommentServiceSignatureXXX() interface{} {
	panic("implement me")
}

func (srv *Service) ListComments(c controller.MContext) {
	panic("implement me")
}

func (srv *Service) CountComment(c controller.MContext) {
	panic("implement me")
}

func (srv *Service) PostComment(c controller.MContext) {
	panic("implement me")
}

func (srv *Service) GetComment(c controller.MContext) {
	panic("implement me")
}

func (srv *Service) PutComment(c controller.MContext) {
	panic("implement me")
}

func NewService(m module.Module) (*Service, error) {
	s := new(Service)
	s.db = m.Require(config.ModulePath.Provider.Model).(*provider.DB).CommentDB()
	return s, nil
}

func (srv *Service) Post(c controller.MContext) {
	var req = new(api.PostCommentRequest)
	if !ginhelper.BindRequest(c, req) {
		return
	}

	//var obj = new(announcement.Announcement)
	//obj.Title = req.Title
	//obj.Content = req.Content
	//
	//cc := ginhelper.GetCustomFields(c)
	//obj.Author = cc.UID
	//obj.LastUpdateUser = cc.UID
	//
	//a, e := srv.db.Create(obj)
	//if ginhelper.CreateObj(c, a, e) {
	//	c.JSON(http.StatusOK, &api.AnnouncementPostReply{
	//		Code:         types.CodeOK,
	//		Announcement: obj,
	//	})
	//}
}

/**
GetAnnouncement v1/announcement/:anid GET
requiring nothing, so anyone is ok.

params:
- `id` uint: the number on placeholder :anid, represent one announcement in
the database


returns:
- `code` int: the operation results
	- types.CodeBindError(1): wrong input, description will be
	attached to the segment of `error`
	- types.CodeInvalidParameters(3): wrong input, description will be
	attached to the segment of `error`. this error might be caused by
	negative id.
	- types.CodeNotFound(102): select error, description will be
	attached to the segment of `error`. this error might be caused by
	the operating contest is not found in database
- `error` string: options description of bad code
- `announcement` GetAnnouncementReply: the selected comment
- `GetAnnouncementReply.id` uint: the announcement's id
- `GetAnnouncementReply.created_at` time.Time: when the announcement is created
- `GetAnnouncementReply.updated_at` time.Time: when the announcement data is last updated
- `GetAnnouncementReply.title` string: the announcement's title
- `GetAnnouncementReply.content` string: the announcement's content
- `GetAnnouncementReply.author` Comment: the announcement's author
- `GetAnnouncementReply.last_update_comment` Comment: the last comment edited this announcement
- `GetAnnouncementReply.is_sticked` Comment: if the announcement is sticked
- `Comment.id` uint: the announcement's comment's id
- `Comment.comment_name` string: the announcement's comment's comment name
- `Comment.nick_name` string: the announcement's comment's nick name
- `Comment.email` string: the announcement's comment's email

Internal Error:
- types.CodeSelectError(101): query error, description will be
attached to the segment of `error`
*/
func (srv *Service) Get(c controller.MContext) {
	//id, ok := ginhelper.ParseUint(c, "anid")
	//if !ok {
	//	return
	//}
	//obj, err := srv.db.ID(id)
	//if ginhelper.MaybeSelectErrorWithTip(c, obj, err, "announcement") {
	//	return
	//}
	//
	//author, err := srv.commentDB.ID(obj.Author)
	//if ginhelper.MaybeSelectErrorWithTip(c, obj, err, "author") {
	//	return
	//}
	//
	//luu, err := srv.commentDB.ID(obj.LastUpdateUser)
	//if ginhelper.MaybeSelectErrorWithTip(c, obj, err, "last update comment") {
	//	return
	//}
	//
	//c.JSON(http.StatusOK, api.AnnouncementToGetReply(obj, author, luu))
}

func (srv *Service) Put(c controller.MContext) {
	//var req = new(api.AnnouncementPutRequest)
	//id, ok := ginhelper.ParseUintAndBind(c, srv.key, req)
	//if !ok {
	//	return
	//}
	//
	//obj, err := srv.db.ID(id)
	//if ginhelper.MaybeSelectError(c, obj, err) {
	//	return
	//}
	//
	//cc := ginhelper.GetCustomFields(c)
	//obj.LastUpdateUser = cc.UID
	//
	//_, err = srv.db.UpdateFields(obj, srv.FillPutFields(obj, req))
	//if ginhelper.UpdateFields(c, err) {
	//	c.JSON(http.StatusOK, &ginhelper.ResponseOK)
	//}
}

/**
DeleteAnnouncement v1/announcement/:anid DELETE
requiring the aiming announcement's write privilege

params:
- `id` uint: the number on placeholder :anid, represent one announcement in
the database

returns:
- `code` int: the operation results
	- types.CodeBindError(1): wrong input, description will be
	attached to the segment of `error`
	- types.CodeInvalidParameters(3): wrong input, description will be
	attached to the segment of `error`. this error might be caused by
	non positive page number or page size.
	- types.CodeNotFound(102): query error, nothing was selected
	- types.CodeDeleteNoEffect(103): delete has no effect
- `error` string: options description of bad code

Internal Error:
- types.CodeDeleteError(106): delete error, description will be
attached to the segment of `error`
*/
func (srv *Service) Delete(c controller.MContext) {
	obj := new(comment.Comment)
	var ok bool
	obj.ID, ok = ginhelper.ParseUint(c, srv.key)
	if !ok {
		return
	}

	a, e := srv.db.Delete(obj)
	if ginhelper.DeleteObj(c, a, e) {
		c.JSON(http.StatusOK, &ginhelper.ResponseOK)
	}
}

func (srv *Service) FillPutFields(obj *announcement.Announcement, req *api.PutUserRequest) (fields []string) {
	//if len(req.Title) != 0 {
	//	obj.Title = req.Title
	//	fields = append(fields, "title")
	//}
	//
	//if len(req.Content) != 0 {
	//	obj.Content = req.Content
	//	fields = append(fields, "content")
	//}

	return
}
