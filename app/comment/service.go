package comment

import (
	"github.com/Myriad-Dreamin/boj-v6/abstract/comment"
	"github.com/Myriad-Dreamin/boj-v6/api"
	"github.com/Myriad-Dreamin/boj-v6/app/snippet"
	"github.com/Myriad-Dreamin/boj-v6/types"
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

func (svc *Service) CommentServiceSignatureXXX() interface{} {
	return svc
}

func NewService(m module.Module) (*Service, error) {
	s := new(Service)
	s.db = m.RequireImpl(new(comment.DB)).(comment.DB)
	return s, nil
}

func (svc *Service) ListComment(c controller.MContext) {
	var req = new(api.ListCommentRequest)
	if !snippet.BindRequest(c, req) {
		return
	}

	ss, err := svc.db.Filter(req)
	if snippet.MaybeSelectError(c, ss, err) {
		return
	}

	c.JSON(http.StatusOK, api.SerializeListCommentReply(types.CodeOK, ss))
	// api.PackSerializeListCommentReply(ss)))

	return
}

func (svc *Service) CountComment(c controller.MContext) {
	var req = new(api.CountCommentRequest)
	if !snippet.BindRequest(c, req) {
		return
	}

	count, err := svc.db.FilterCount(req)
	if snippet.MaybeCountError(c, err) {
		return
	}

	c.JSON(http.StatusOK, api.CountCommentReply{
		Code: types.CodeOK,
		Data: count,
	})
}

func (svc *Service) PostComment(c controller.MContext) {
	var req = new(api.PostCommentRequest)
	if !snippet.BindRequest(c, req) {
		return
	}

	var obj = new(comment.Comment)
	obj.Title = req.Title
	obj.Content = req.Content

	cc := snippet.GetCustomFields(c)
	obj.AuthorID = cc.UID
	obj.LastUpdateUserID = cc.UID

	a, e := svc.db.Create(obj)
	if snippet.CreateObj(c, a, e) {
		c.JSON(http.StatusOK, api.SerializePostCommentReply(types.CodeOK, obj))
	}
}

func (svc *Service) GetComment(c controller.MContext) {
	id, ok := snippet.ParseUint(c, svc.key)
	if !ok {
		return
	}
	obj, err := svc.db.ID(id)
	if snippet.MaybeSelectErrorWithTip(c, obj, err, "announcement") {
		return
	}

	//author, err := svc.commentDB.ID(obj.Author)
	//if snippet.MaybeSelectErrorWithTip(c, obj, err, "author") {
	//	return
	//}
	//
	//luu, err := svc.commentDB.ID(obj.LastUpdateUser)
	//if snippet.MaybeSelectErrorWithTip(c, obj, err, "last update comment") {
	//	return
	//}

	c.JSON(http.StatusOK, api.SerializeGetCommentReply(types.CodeOK, obj)) // api.AnnouncementToGetReply(obj, author, luu))
}

func (svc *Service) DeleteComment(c controller.MContext) {
	obj := new(comment.Comment)
	var ok bool
	obj.ID, ok = snippet.ParseUint(c, svc.key)
	if !ok {
		return
	}

	a, e := svc.db.Delete(obj)
	if snippet.DeleteObj(c, a, e) {
		c.JSON(http.StatusOK, &snippet.ResponseOK)
	}
}

func (svc *Service) PutComment(c controller.MContext) {
	var req = new(api.PutCommentRequest)
	id, ok := snippet.ParseUintAndBind(c, svc.key, req)
	if !ok {
		return
	}

	obj, err := svc.db.ID(id)
	if snippet.MaybeSelectError(c, obj, err) {
		return
	}

	cc := snippet.GetCustomFields(c)
	obj.LastUpdateUserID = cc.UID

	_, err = svc.db.UpdateFields(obj, svc.FillPutFields(obj, req))
	if snippet.UpdateFields(c, err) {
		c.JSON(http.StatusOK, &snippet.ResponseOK)
	}
}

func (svc *Service) FillPutFields(obj *comment.Comment, req *api.PutCommentRequest) (fields []string) {
	if len(req.Title) != 0 {
		obj.Title = req.Title
		fields = append(fields, "title")
	}

	if len(req.Content) != 0 {
		obj.Content = req.Content
		fields = append(fields, "content")
	}

	return
}
