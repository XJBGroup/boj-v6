package announcement

import (
	"github.com/Myriad-Dreamin/boj-v6/abstract/announcement"
	"github.com/Myriad-Dreamin/boj-v6/abstract/user"
	"github.com/Myriad-Dreamin/boj-v6/api"
	"github.com/Myriad-Dreamin/boj-v6/app/provider"
	"github.com/Myriad-Dreamin/boj-v6/app/snippet"
	"github.com/Myriad-Dreamin/boj-v6/config"
	"github.com/Myriad-Dreamin/boj-v6/types"
	"github.com/Myriad-Dreamin/core-oj/log"
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	"github.com/Myriad-Dreamin/minimum-lib/module"
	"net/http"
)

type Service struct {
	db     announcement.DB
	userDB user.DB
	logger log.TendermintLogger
	key    string
}

func NewService(m module.Module) (*Service, error) {
	s := new(Service)
	s.db = m.Require(config.ModulePath.Provider.Model).(*provider.DB).AnnouncementDB()
	s.userDB = m.Require(config.ModulePath.Provider.Model).(*provider.DB).UserDB()
	s.key = "aid"
	return s, nil
}

func (svc *Service) AnnouncementServiceSignatureXXX() interface{} {
	return svc
}

func (svc *Service) ListAnnouncements(c controller.MContext) {
	page, pageSize, ok := snippet.RosolvePageVariable(c)
	if !ok {
		return
	}

	announcements, err := svc.db.Find(page, pageSize)
	if snippet.MaybeSelectError(c, announcements, err) {
		return
	}

	c.JSON(http.StatusOK, api.SerializeListAnnouncementsReply(types.CodeOK, announcements))
	return
}

func (svc *Service) CountAnnouncement(c controller.MContext) {
	count, err := svc.db.Count()
	if snippet.MaybeCountError(c, err) {
		return
	}

	c.JSON(http.StatusOK, api.CountAnnouncementReply{
		Code: types.CodeOK,
		Data: count,
	})
}

/**
PostAnnouncement v1/announcement POST
requiring for the identity of administrator.
which means that the request to this method must be with header
"Authorization": "Bearer {your token}"
and, the operating user must be in the group of admin.
*/
func (svc *Service) PostAnnouncement(c controller.MContext) {
	var req = new(api.PostAnnouncementRequest)
	if !snippet.BindRequest(c, req) {
		return
	}

	var obj = new(announcement.Announcement)
	obj.Title = req.Title
	obj.Content = req.Content

	cc := snippet.GetCustomFields(c)
	obj.Author = cc.UID
	obj.LastUpdateUser = cc.UID

	a, e := svc.db.Create(obj)
	if snippet.CreateObj(c, a, e) {
		c.JSON(http.StatusOK, api.SerializePostAnnouncementReply(types.CodeOK, obj))
	}
}

/**
GetAnnouncement v1/announcement/:aid GET
requiring nothing, so anyone is ok.
*/
func (svc *Service) GetAnnouncement(c controller.MContext) {
	id, ok := snippet.ParseUint(c, svc.key)
	if !ok {
		return
	}
	obj, err := svc.db.ID(id)
	if snippet.MaybeSelectErrorWithTip(c, obj, err, "announcement") {
		return
	}

	//author, err := svc.userDB.ID(obj.Author)
	//if snippet.MaybeSelectErrorWithTip(c, obj, err, "author") {
	//	return
	//}
	//
	//luu, err := svc.userDB.ID(obj.LastUpdateUser)
	//if snippet.MaybeSelectErrorWithTip(c, obj, err, "last update user") {
	//	return
	//}

	c.JSON(http.StatusOK, api.SerializeGetAnnouncementReply(types.CodeOK, obj))
}

func (svc *Service) PutAnnouncement(c controller.MContext) {
	var req = new(api.PutAnnouncementRequest)
	id, ok := snippet.ParseUintAndBind(c, svc.key, req)
	if !ok {
		return
	}

	obj, err := svc.db.ID(id)
	if snippet.MaybeSelectError(c, obj, err) {
		return
	}

	cc := snippet.GetCustomFields(c)
	obj.LastUpdateUser = cc.UID

	_, err = svc.db.UpdateFields(obj, svc.FillPutFields(obj, req))
	if snippet.UpdateFields(c, err) {
		c.JSON(http.StatusOK, &snippet.ResponseOK)
	}
}

/**
DeleteAnnouncement v1/announcement/:aid DELETE
requiring the aiming announcement's write privilege
*/
func (svc *Service) Delete(c controller.MContext) {
	obj := new(announcement.Announcement)
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

func (svc *Service) FillPutFields(obj *announcement.Announcement, req *api.PutAnnouncementRequest) (fields []string) {
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
