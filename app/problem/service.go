package problem

import (
	"github.com/Myriad-Dreamin/boj-v6/abstract/problem"
	problem_desc "github.com/Myriad-Dreamin/boj-v6/abstract/problem-desc"
	"github.com/Myriad-Dreamin/boj-v6/abstract/user"
	"github.com/Myriad-Dreamin/boj-v6/api"
	"github.com/Myriad-Dreamin/boj-v6/app/snippet"
	"github.com/Myriad-Dreamin/boj-v6/config"
	"github.com/Myriad-Dreamin/boj-v6/external"
	"github.com/Myriad-Dreamin/boj-v6/lib/serial"
	"github.com/Myriad-Dreamin/boj-v6/types"
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	"github.com/Myriad-Dreamin/minimum-lib/module"
	"net/http"
	"strconv"
)

type Service struct {
	db       problem.DB
	descDB   problem_desc.DB
	userDB   user.DB
	enforcer *external.Enforcer
	logger   external.Logger
	cfg      *config.ServerConfig
	key      string
}

func NewService(m module.Module) (*Service, error) {
	s := new(Service)
	s.db = m.RequireImpl(new(problem.DB)).(problem.DB)
	s.descDB = m.RequireImpl(new(problem_desc.DB)).(problem_desc.DB)
	s.userDB = m.RequireImpl(new(user.DB)).(user.DB)
	s.enforcer = m.RequireImpl(new(*external.Enforcer)).(*external.Enforcer)
	s.logger = m.RequireImpl(new(external.Logger)).(external.Logger)
	s.cfg = m.RequireImpl(new(*config.ServerConfig)).(*config.ServerConfig)

	s.key = "pid"
	return s, nil
}

func (svc Service) ProblemServiceSignatureXXX() interface{} {
	return svc
}

func (svc Service) ListProblem(c controller.MContext) {
	page, pageSize, ok := snippet.RosolvePageVariable(c)
	if !ok {
		return
	}

	problems, err := svc.db.Find(page, pageSize)
	if snippet.MaybeSelectError(c, problems, err) {
		return
	}

	c.JSON(http.StatusOK, api.SerializeListProblemReply(types.CodeOK, problems))
}

func (svc Service) CountProblem(c controller.MContext) {
	// todo: problem filter
	cnt, err := svc.db.Count()
	if snippet.MaybeCountError(c, err) {
		return
	}

	c.JSON(http.StatusOK, api.CountSubmissionReply{
		Code: types.CodeOK,
		Data: cnt,
	})
}

func (svc Service) PutProblem(c controller.MContext) {
	var req = new(api.PutProblemRequest)
	id, ok := snippet.ParseUintAndBind(c, svc.key, req)
	if !ok {
		return
	}

	obj, err := svc.db.ID(id)
	if snippet.MaybeSelectError(c, obj, err) {
		return
	}

	_, err = svc.db.UpdateFields(obj, svc.FillPutFields(obj, req))
	if snippet.UpdateFields(c, err) {
		c.JSON(http.StatusOK, &snippet.ResponseOK)
	}
}

func (svc *Service) FillPutFields(problem *problem.Problem, req *api.PutProblemRequest) (fields []string) {
	// todo: update request fields
	if len(req.Title) != 0 {
		problem.Title = req.Title
		fields = append(fields, "title")
	}
	//if len(req.Description) != 0 {
	//	problem.Description = req.Description
	//	fields = append(fields, "description")
	//}
	if len(req.DescriptionRef) != 0 {
		problem.DescriptionRef = req.DescriptionRef
		fields = append(fields, "description_ref")
	}

	//if req.TimeLimit != 0 {
	//	problem.TimeLimit = req.TimeLimit
	//	fields = append(fields, "time-limit")
	//}
	//
	//if req.MemoryLimit != 0 {
	//	problem.MemoryLimit = req.MemoryLimit
	//	fields = append(fields, "memory-limit")
	//}
	//
	//if req.CodeLengthLimit != 0 {
	//	problem.CodeLengthLimit = req.CodeLengthLimit
	//	fields = append(fields, "code_length-limit")
	//}
	//
	//if req.UpdateSpj {
	//	problem.IsSpj = req.IsSpj
	//	fields = append(fields, "is-spj")
	//}

	return
}

func (svc Service) GetProblem(c controller.MContext) {
	id, ok := snippet.ParseUint(c, svc.key)
	if !ok {
		return
	}
	p, err := svc.db.ID(id)
	if snippet.MaybeSelectErrorWithTip(c, p, err, "problem") {
		return
	}

	author, err := svc.userDB.ID(p.AuthorID)
	if snippet.MaybeSelectErrorWithTip(c, author, err, "author") {
		return
	}

	if len(p.Description) == 0 {
		pd, err := svc.descDB.QueryByKey(id, p.DescriptionRef)
		if snippet.MaybeSelectErrorWithTip(c, pd, err, "problem_desc") {
			return
		}
		err = svc.descDB.LoadDesc(pd)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusOK, serial.ErrorSerializer{
				Code:   types.CodeSelectError,
				ErrorS: err.Error(),
			})
			return
		}

		p.Description = string(pd.Content)
		err = svc.descDB.ReleaseDesc(pd)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusOK, serial.ErrorSerializer{
				Code:   types.CodeSelectError,
				ErrorS: err.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusOK, api.SerializeGetProblemReply(types.CodeOK,
		api.SerializeGetProblemData(p, api.SerializeGetProblemAuthorData(author))))
}

func (svc Service) DeleteProblem(c controller.MContext) {
	obj := new(problem.Problem)
	var ok bool
	obj.ID, ok = snippet.ParseUint(c, svc.key)
	if !ok {
		return
	}

	ds, err := svc.descDB.QueryByPID(obj.ID)
	if snippet.MaybeSelectError(c, ds, err) {
		return
	}

	for i := range ds {
		err = svc.descDB.DeleteDesc(&ds[i])
		if err != nil {
			c.JSON(http.StatusOK, serial.ErrorSerializer{
				Code:   types.CodeDeleteError,
				ErrorS: strconv.FormatInt(int64(i), 10) + " not deleted oss source: " + err.Error(),
			})
		}

		aff, err := svc.descDB.Delete(&ds[i])
		if !snippet.DeleteObj(c, aff, err) {
			c.JSON(http.StatusOK, serial.ErrorSerializer{
				Code:   types.CodeDeleteError,
				ErrorS: strconv.FormatInt(int64(i), 10) + " not deleted relationship: " + err.Error(),
			})
		}
	}

	aff, err := svc.db.Delete(obj)
	if snippet.DeleteObj(c, aff, err) {
		c.JSON(http.StatusOK, snippet.ResponseOK)
	}
}
