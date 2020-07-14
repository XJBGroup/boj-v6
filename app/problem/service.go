package problem

import (
	"github.com/Myriad-Dreamin/boj-v6/abstract/problem"
	"github.com/Myriad-Dreamin/boj-v6/api"
	"github.com/Myriad-Dreamin/boj-v6/app/provider"
	"github.com/Myriad-Dreamin/boj-v6/app/snippet"
	"github.com/Myriad-Dreamin/boj-v6/config"
	"github.com/Myriad-Dreamin/boj-v6/external"
	"github.com/Myriad-Dreamin/boj-v6/types"
	problemconfig "github.com/Myriad-Dreamin/boj-v6/types/problem-config"
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	"github.com/Myriad-Dreamin/minimum-lib/module"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
)

type Service struct {
	db       problem.DB
	enforcer *provider.Enforcer
	logger   external.Logger
	cfg      *config.ServerConfig
	key      string
}

func NewService(m module.Module) (*Service, error) {
	s := new(Service)
	s.db = m.Require(config.ModulePath.Provider.Model).(*provider.DB).ProblemDB()
	s.enforcer = m.Require(config.ModulePath.Provider.Model).(*provider.DB).Enforcer()
	s.logger = m.Require(config.ModulePath.Global.Logger).(external.Logger)
	s.cfg = m.Require(config.ModulePath.Global.Configuration).(*config.ServerConfig)

	s.key = "pid"
	return s, nil
}

func (svc Service) ProblemServiceSignatureXXX() interface{} {
	return svc
}

func (svc Service) ListProblems(c controller.MContext) {
	page, pageSize, ok := snippet.RosolvePageVariable(c)
	if !ok {
		return
	}

	problems, err := svc.db.Find(page, pageSize)
	if snippet.MaybeSelectError(c, problems, err) {
		return
	}

	c.JSON(http.StatusOK, api.SerializeListProblemsReply(types.CodeOK, problems))
}

func (svc Service) CountProblem(c controller.MContext) {
	cnt, err := svc.db.Count()
	if snippet.MaybeCountError(c, err) {
		return
	}

	c.JSON(http.StatusOK, api.CountSubmissionsReply{
		Code: types.CodeOK,
		Data: cnt,
	})
}

func (svc Service) PostProblem(c controller.MContext) {
	var req = new(api.PostProblemRequest)
	req.Config = problemconfig.DefaultProblemConfig()
	if !snippet.BindRequest(c, req) {
		return
	}

	var p = new(problem.Problem)
	p.Title = req.Title
	p.DescriptionRef = "default"

	cc := snippet.GetCustomFields(c)
	p.AuthorID = cc.UID

	aff, err := svc.db.Create(p)
	if !snippet.CreateObjWithTip(c, aff, err, "problem") {
		return
	}

	// todo: problem desc
	//var problemDesc = model.NewProblemDesc(problem.ID, "default", []byte(req.Description))
	//
	//
	//if !snippet.CreateObjWithTip(c, problemDesc) {
	//	return
	//}

	//if err := problemDesc.Save(); err != nil {
	//	c.AbortWithStatusJSON(http.StatusInternalServerError, snippet.ErrorSerializer{
	//		Code:  types.CodeInsertError,
	//		Error: err.Error(),
	//	})
	//}

	var path = svc.cfg.PathConfig.ProblemPath + strconv.Itoa(int(p.ID))
	if _, err := os.Stat(path); err != nil {
		if !os.IsExist(err) {
			err = os.Mkdir(path, 0770)
			if err != nil {
				_ = c.AbortWithError(http.StatusInternalServerError, err)
				return
			}
		} else {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{
				"code": types.CodeFSExecError,
				"err":  err.Error(),
			})
			return
		}
	}

	configPath := path + "/problem-config"

	err = problemconfig.Save(req.Config, configPath)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": types.CodeFSExecError,
			"err":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &api.PostProblemReply{
		Code: types.CodeOK,
		Id:   p.ID,
	})
}

func (svc Service) GetProblem(c controller.MContext) {
	id, ok := snippet.ParseUint(c, svc.key)
	if !ok {
		return
	}
	p, err := svc.db.ID(id)
	if snippet.MaybeSelectError(c, p, err) {
		return
	}

	// todo: get problem desc
	//user, err := srv.userDB.ID(problem.AuthorID)
	//if snippet.MaybeSelectError(c, user, err) {
	//	return
	//}
	//problem.Author = *user

	//problemDesc, err := srv.problemDescDB.QueryTemplate(id, problem.Description)
	//if snippet.MaybeSelectError(c, problemDesc, err) {
	//	return
	//}
	//err = problemDesc.Load()
	//if err != nil {
	//	c.AbortWithStatusJSON(http.StatusInternalServerError, snippet.ErrorSerializer{
	//		Code:  types.CodeSelectError,
	//		Error: err.Error(),
	//	})
	//	return
	//}

	c.JSON(http.StatusOK, api.SerializeGetProblemReply(types.CodeOK, p)) //ProblemToGetReply(problem, problemDesc))
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

func (svc Service) DeleteProblem(c controller.MContext) {
	obj := new(problem.Problem)
	var ok bool
	obj.ID, ok = snippet.ParseUint(c, svc.key)
	if !ok {
		return
	}

	aff, err := svc.db.Delete(obj)
	if snippet.DeleteObj(c, aff, err) {
		c.JSON(http.StatusOK, snippet.ResponseOK)
	}

	// todo: delete problem desc
}

func (svc *Service) FillPutFields(problem *problem.Problem, req *api.PutProblemRequest) (fields []string) {
	if len(req.Title) != 0 {
		problem.Title = req.Title
		fields = append(fields, "title")
	}
	if len(req.Description) != 0 {
		problem.Description = req.Description
		fields = append(fields, "description")
	}
	if len(req.DescriptionRef) != 0 {
		problem.DescriptionRef = req.DescriptionRef
		fields = append(fields, "desc_ref_name")
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
