package group

import (
	"github.com/Myriad-Dreamin/boj-v6/abstract/group"
	"github.com/Myriad-Dreamin/boj-v6/abstract/user"
	"github.com/Myriad-Dreamin/boj-v6/api"
	"github.com/Myriad-Dreamin/boj-v6/app/snippet"
	"github.com/Myriad-Dreamin/boj-v6/external"
	"github.com/Myriad-Dreamin/boj-v6/lib/serial"
	"github.com/Myriad-Dreamin/boj-v6/types"
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	"github.com/Myriad-Dreamin/minimum-lib/module"
	"net/http"
)

type Service struct {
	db       group.DB
	userDB   user.DB
	enforcer *external.Enforcer
	logger   external.Logger
	key      string
}

func (svc Service) GroupServiceSignatureXXX() interface{} {
	return svc
}

func NewService(m module.Module) (*Service, error) {
	s := new(Service)
	s.enforcer = m.RequireImpl(new(*external.Enforcer)).(*external.Enforcer)
	s.db = m.RequireImpl(new(group.DB)).(group.DB)
	s.userDB = m.RequireImpl(new(user.DB)).(user.DB)
	s.logger = m.RequireImpl(new(external.Logger)).(external.Logger)

	s.key = "gid"
	return s, nil
}

func (svc Service) ListGroup(c controller.MContext) {
	page, pageSize, ok := snippet.RosolvePageVariable(c)
	if !ok {
		return
	}

	groups, err := svc.db.Find(page, pageSize)
	if snippet.MaybeSelectError(c, groups, err) {
		return
	}

	c.JSON(http.StatusOK, api.SerializeListGroupReply(types.CodeOK, groups))
}

func (svc Service) CountGroup(c controller.MContext) {
	count, err := svc.db.Count()
	if snippet.MaybeCountError(c, err) {
		return
	}

	c.JSON(http.StatusOK, api.CountGroupReply{
		Code: types.CodeOK,
		Data: count,
	})
}

func (svc Service) OnGroupCreate(id uint, id2 uint) error {
	return nil
}

func (svc Service) PostGroup(c controller.MContext) {
	var req = new(api.PostGroupRequest)
	if !snippet.BindRequest(c, req) {
		return
	}

	var g = new(group.Group)
	g.Name = req.Name
	g.Description = req.Description

	var u *user.User
	var err error
	if req.OwnerId != 0 {
		u, err = svc.userDB.ID(req.OwnerId)
	} else if len(req.OwnerName) != 0 {
		u, err = svc.userDB.QueryUserName(req.OwnerName)
	} else {
		c.JSON(http.StatusOK, serial.ErrorSerializer{
			Code:   types.CodeInvalidParameters,
			ErrorS: "miss the id or name of group's owner",
		})
	}

	if snippet.MaybeSelectError(c, u, err) {
		return
	}
	g.OwnerID = u.ID

	aff, err := svc.db.Create(g)
	if !snippet.CreateObj(c, aff, err) {
		return
	}

	if err := svc.OnGroupCreate(g.ID, u.ID); err != nil {
		if aff, err := svc.db.Delete(g); aff == 0 || err != nil {
			svc.logger.Debug("group create failed, and delete error...",
				"affected", aff, "error", snippet.ConvertErrorToString(err))
		}
		c.JSON(http.StatusOK, &serial.ErrorSerializer{
			Code:   types.CodeGroupCreateError,
			ErrorS: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &api.PostGroupReply{
		Code: types.CodeOK,
		Data: g.ID,
	})
}

func (svc Service) GetGroup(c controller.MContext) {
	id, ok := snippet.ParseUint(c, svc.key)
	if !ok {
		return
	}
	obj, err := svc.db.ID(id)
	if snippet.MaybeSelectError(c, obj, err) {
		return
	}
	//user, err := srv.userDB.ID(obj.OwnerID)
	//if ginhelper.MaybeSelectError(c, user, err) {
	//	return
	//}
	//obj.Owner = *user

	c.JSON(http.StatusOK, api.SerializeGetGroupReply(types.CodeOK, obj)) // GroupToGetReply(obj))
}

func (svc Service) DeleteGroup(c controller.MContext) {
	id, ok := snippet.ParseUint(c, svc.key)
	if !ok {
		return
	}
	obj, err := svc.db.ID(id)
	if snippet.MaybeSelectError(c, obj, err) {
		return
	}
	// todo hook delete
	//if err := srv.ReportGroupDeleted(obj.ID, obj.OwnerID, nil); err != nil {
	//	return
	//}

	aff, err := svc.db.Delete(obj)
	if snippet.DeleteObj(c, aff, err) {
		c.JSON(http.StatusOK, &snippet.ResponseOK)
	}
}

func (svc Service) PutGroup(c controller.MContext) {
	var req = new(api.PutGroupRequest)
	id, ok := snippet.ParseUintAndBind(c, svc.key, req)
	if !ok {
		return
	}

	g, err := svc.db.ID(id)
	if snippet.MaybeSelectError(c, g, err) {
		return
	}

	_, err = svc.db.UpdateFields(g, svc.FillPutFields(g, req))
	if snippet.UpdateFields(c, err) {
		c.JSON(http.StatusOK, &snippet.ResponseOK)
	}
}
func (svc *Service) FillPutFields(group *group.Group, req *api.PutGroupRequest) (fields []string) {
	if len(req.Name) != 0 {
		group.Name = req.Name
		fields = append(fields, "name")
	}

	if len(req.Description) != 0 {
		group.Description = req.Description
		fields = append(fields, "description")
	}

	return
}
