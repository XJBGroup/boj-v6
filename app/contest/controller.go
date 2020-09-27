package contest

import (
	"github.com/Myriad-Dreamin/boj-v6/abstract/contest"
	"github.com/Myriad-Dreamin/boj-v6/api"
	"github.com/Myriad-Dreamin/boj-v6/app/snippet"
	"github.com/Myriad-Dreamin/boj-v6/external"
	"github.com/Myriad-Dreamin/boj-v6/types"
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	"github.com/Myriad-Dreamin/minimum-lib/module"
	"net/http"
)

type Controller struct {
	enforcer *external.Enforcer
	db       contest.DB
	logger   external.Logger
	key      string
}

func (svc Controller) ContestControllerSignatureXXX() interface{} {
	return svc
}

func NewController(m module.Module) (*Controller, error) {
	s := new(Controller)
	s.enforcer = m.RequireImpl(new(*external.Enforcer)).(*external.Enforcer)
	s.db = m.RequireImpl(new(contest.DB)).(contest.DB)
	s.logger = m.RequireImpl(new(external.Logger)).(external.Logger)

	s.key = "cid"
	return s, nil
}

func (svc Controller) ListContest(c controller.MContext) {
	page, pageSize, ok := snippet.RosolvePageVariable(c)
	if !ok {
		return
	}

	contests, err := svc.db.Find(page, pageSize)
	if snippet.MaybeSelectError(c, contests, err) {
		return
	}

	c.JSON(http.StatusOK, api.SerializeListContestReply(types.CodeOK, contests))
	return
}

func (svc Controller) CountContest(c controller.MContext) {
	count, err := svc.db.Count()
	if snippet.MaybeCountError(c, err) {
		return
	}

	c.JSON(http.StatusOK, api.CountAnnouncementReply{
		Code: types.CodeOK,
		Data: count,
	})
}

func (svc Controller) PostContest(c controller.MContext) {
	var req = new(api.PostContestRequest)
	if !snippet.BindRequest(c, req) {
		return
	}

	var con = new(contest.Contest)
	con.Title = req.Title
	con.Description = req.Description
	con.StartAt = req.StartAt
	con.EndDuration = req.EndDuration
	con.BoardFrozenDuration = req.BoardFrozenDuration
	//c.ConfigPath = req.ConfigPath
	//c.RolePath = req.RolePath

	cc := snippet.GetCustomFields(c)
	con.AuthorID = cc.UID

	aff, err := svc.db.Create(con)
	if snippet.CreateObj(c, svc.db.UnwrapError, aff, err) {
		c.JSON(http.StatusOK, api.SerializePostContestReply(types.CodeOK, con))
	}
}

func (svc Controller) GetContest(c controller.MContext) {
	id, ok := snippet.ParseUint(c, svc.key)
	if !ok {
		return
	}
	obj, err := svc.db.ID(id)
	if snippet.MaybeSelectErrorWithTip(c, obj, err, "contest") {
		return
	}

	c.JSON(http.StatusOK, api.SerializeGetContestReply(types.CodeOK,
		api.SerializeGetContestInnerReply(obj)))
}

func (svc Controller) DeleteContest(c controller.MContext) {
	obj := new(contest.Contest)
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

func (svc Controller) PutContest(c controller.MContext) {
	var req = new(api.PutContestRequest)
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

func (svc *Controller) FillPutFields(contest *contest.Contest, req *api.PutContestRequest) (fields []string) {
	if len(req.Title) != 0 {
		contest.Title = req.Title
		fields = append(fields, "title")
	}

	if len(req.Description) != 0 {
		contest.Description = req.Description
		fields = append(fields, "description")
	}

	if req.StartAt != nil {
		contest.StartAt = req.StartAt
		fields = append(fields, "start_at")
	}

	if req.EndDuration != 0 {
		contest.EndDuration = req.EndDuration
		fields = append(fields, "end_duration")
	}

	if req.BoardFrozenDuration != 0 {
		contest.BoardFrozenDuration = req.BoardFrozenDuration
		fields = append(fields, "board_frozen_duration")
	}

	if len(req.ConfigPath) != 0 {
		contest.ConfigPath = req.ConfigPath
		fields = append(fields, "config_path")
	}

	if len(req.RolePath) != 0 {
		contest.RolePath = req.RolePath
		fields = append(fields, "role_path")
	}

	return
}
