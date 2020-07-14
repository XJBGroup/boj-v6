package submission

import (
	"github.com/Myriad-Dreamin/boj-v6/abstract/submission"
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
	db     submission.DB
	userDB user.DB
	logger log.TendermintLogger
	key    string
}

func NewService(m module.Module) (*Service, error) {
	s := new(Service)
	s.db = m.Require(config.ModulePath.Provider.Model).(*provider.DB).SubmissionDB()
	s.userDB = m.Require(config.ModulePath.Provider.Model).(*provider.DB).UserDB()
	s.key = "sid"
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
	var req = new(api.SubmissionFilter)
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

func (svc Service) ListSubmissions(c controller.MContext) {
	f := svc.ResolveFilter(c)
	if c.IsAborted() {
		return
	}

	ss, err := svc.db.Filter(f)
	if snippet.MaybeSelectError(c, ss, err) {
		return
	}

	c.JSON(http.StatusOK, api.SerializeListSubmissionsReply(types.CodeOK,
		api.PackSerializeListSubmissionReply(ss)))

	return
}

func (svc Service) CountSubmissions(c controller.MContext) {
	f := svc.ResolveFilter(c)
	if c.IsAborted() {
		return
	}

	cnt, err := svc.db.FilterCount(f)
	if snippet.MaybeCountError(c, err) {
		return
	}

	c.JSON(http.StatusOK, api.CountSubmissionsReply{
		Code: types.CodeOK,
		Data: cnt,
	})
}

func (svc Service) PostSubmission(c controller.MContext) {
	panic("implement me")
}

func (svc Service) GetContent(c controller.MContext) {
	panic("implement me")
}

func (svc Service) GetSubmission(c controller.MContext) {
	panic("implement me")
}

func (svc Service) Delete(c controller.MContext) {
	panic("implement me")
}
