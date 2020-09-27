package submission

import (
	"github.com/Myriad-Dreamin/boj-v6/abstract/submission"
	"github.com/Myriad-Dreamin/boj-v6/api"
	"github.com/Myriad-Dreamin/boj-v6/app/snippet"
	"github.com/Myriad-Dreamin/minimum-lib/controller"
)

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

func (svc *Controller) ResolveFilter(c controller.MContext) *submission.Filter {
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
