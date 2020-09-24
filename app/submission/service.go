package submission

import (
	inner_control "github.com/Myriad-Dreamin/boj-v6/abstract/inner-control"
	"github.com/Myriad-Dreamin/boj-v6/abstract/problem"
	"github.com/Myriad-Dreamin/boj-v6/abstract/submission"
	"github.com/Myriad-Dreamin/boj-v6/abstract/user"
	"github.com/Myriad-Dreamin/boj-v6/api"
	"github.com/Myriad-Dreamin/boj-v6/app/snippet"
	"github.com/Myriad-Dreamin/boj-v6/config"
	"github.com/Myriad-Dreamin/boj-v6/external"
	"github.com/Myriad-Dreamin/boj-v6/types"
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	"github.com/Myriad-Dreamin/minimum-lib/module"
	"net/http"
	"os"
	"path/filepath"
)

type Service struct {
	db         submission.DB
	problemDB  problem.DB
	userDB     user.DB
	logger     external.Logger
	cfg        *config.ServerConfig
	key        string
	problemKey string

	inner inner_control.InnerSubmissionService
}

func NewService(m module.Module) (*Service, error) {
	s := new(Service)
	s.db = m.RequireImpl(new(submission.DB)).(submission.DB)
	s.problemDB = m.RequireImpl(new(problem.DB)).(problem.DB)
	s.userDB = m.RequireImpl(new(user.DB)).(user.DB)
	s.logger = m.RequireImpl(new(external.Logger)).(external.Logger)
	s.cfg = m.RequireImpl(new(*config.ServerConfig)).(*config.ServerConfig)
	s.problemKey = "pid"
	s.key = "sid"

	s.inner = m.RequireImpl(new(inner_control.InnerSubmissionService)).(inner_control.InnerSubmissionService)
	return s, nil
}

func (svc Service) SubmissionServiceSignatureXXX() interface{} {
	return svc
}

const (
	nonOrder        = "id desc"
	revTimeOrder    = "running_time desc, id desc"
	revTimeRevOrder = "running_time desc"
	timeOrder       = "running_time, id desc"
	timeRevOrder    = "running_time"
	revNonOrder     = ""
	revMemOrder     = "running_memory desc, id desc"
	revMemRevOrder  = "running_memory desc"
	memOrder        = "running_memory, id desc"
	memRevOrder     = "running_memory"
	notContest      = 1 << 62
)

func (svc Service) ResolveFilter(c controller.MContext) *submission.Filter {
	var req = new(api.ListSubmissionRequest)
	if !snippet.BindRequest(c, req) {
		return nil
	}
	var f submission.Filter
	f.Order = nonOrder
	if req.MemOrder != nil {
		if req.IdOrder != nil {
			switch {
			case *req.MemOrder && *req.IdOrder:
				f.Order = memRevOrder
			case *req.MemOrder && !*req.IdOrder:
				f.Order = memOrder
			case !*req.MemOrder && *req.IdOrder:
				f.Order = revMemRevOrder
			default:
				f.Order = revMemOrder
			}
		} else {
			switch {
			case *req.MemOrder:
				f.Order = memOrder
			default:
				f.Order = revMemOrder
			}
		}
	} else if req.TimeOrder != nil {
		if req.IdOrder != nil {
			switch {
			case *req.TimeOrder && *req.IdOrder:
				f.Order = timeRevOrder
			case *req.TimeOrder && !*req.IdOrder:
				f.Order = timeOrder
			case !*req.TimeOrder && *req.IdOrder:
				f.Order = revTimeRevOrder
			default:
				f.Order = revTimeOrder
			}
		} else {
			switch {
			case *req.TimeOrder:
				f.Order = timeOrder
			default:
				f.Order = revTimeOrder
			}
		}
	} else if req.IdOrder != nil {
		switch {
		case *req.IdOrder:
			f.Order = revNonOrder
		default:
			f.Order = nonOrder
		}
	}

	f.Page = req.Page
	f.PageSize = req.PageSize
	f.ByUser = req.ByUser
	f.OnProblem = req.OnProblem
	f.WithLanguage = req.WithLanguage
	f.HasStatus = req.HasStatus
	return &f
}

func (svc Service) ListSubmission(c controller.MContext) {
	f := svc.ResolveFilter(c)
	if c.IsAborted() {
		return
	}

	ss, err := svc.db.Filter(f)
	if snippet.MaybeSelectError(c, ss, err) {
		return
	}

	c.JSON(http.StatusOK, api.SerializeListSubmissionReply(types.CodeOK,
		api.PackSerializeListSubmissionInnerReply(ss)))

	return
}

func (svc Service) CountSubmission(c controller.MContext) {
	f := svc.ResolveFilter(c)
	if c.IsAborted() {
		return
	}

	cnt, err := svc.db.FilterCount(f)
	if snippet.MaybeCountError(c, err) {
		return
	}

	c.JSON(http.StatusOK, api.CountSubmissionReply{
		Code: types.CodeOK,
		Data: cnt,
	})
}

func (svc Service) GetContent(c controller.MContext) {
	if c.IsAborted() {
		return
	}
	id, ok := snippet.ParseUint(c, svc.key)
	if !ok {
		return
	}

	s, err := svc.db.ID(id)
	if snippet.MaybeSelectError(c, s, err) {
		return
	}

	// if s.Shared != 1
	if s.Shared == 1 {
		c.File(filepath.Join(svc.cfg.PathConfig.CodePath, s.Hash, "main"))
	}
}

func (svc Service) GetSubmission(c controller.MContext) {
	id, ok := snippet.ParseUint(c, svc.key)
	if !ok {
		return
	}
	obj, err := svc.db.ID(id)
	if snippet.MaybeSelectError(c, obj, err) {
		return
	}

	c.JSON(http.StatusOK, api.SerializeGetSubmissionReply(types.CodeOK, api.SerializeGetSubmissionInnerReply(obj)))
}

func (svc Service) DeleteSubmission(c controller.MContext) {
	id, ok := snippet.ParseUint(c, svc.key)
	if !ok {
		return
	}
	obj, err := svc.db.ID(id)
	if snippet.MaybeSelectError(c, obj, err) {
		return
	}

	var path = filepath.Join(svc.cfg.PathConfig.CodePath, obj.Hash)
	if _, err = os.Stat(path); err == nil {
		err = os.RemoveAll(path)
		if err != nil {
			_ = c.AbortWithError(http.StatusOK, err)
			return
		}
	} else {
		_ = c.AbortWithError(http.StatusOK, err)
	}

	aff, err := svc.db.Delete(obj)
	if snippet.DeleteObj(c, aff, err) {
		c.JSON(http.StatusOK, &snippet.ResponseOK)
	}
}
