package problem

import (
	"github.com/Myriad-Dreamin/boj-v6/abstract/problem"
	"github.com/Myriad-Dreamin/boj-v6/api"
	"github.com/Myriad-Dreamin/boj-v6/app/provider"
	"github.com/Myriad-Dreamin/boj-v6/config"
	"github.com/Myriad-Dreamin/boj-v6/external"
	ginhelper "github.com/Myriad-Dreamin/boj-v6/lib/gin-helper"
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

	s.key = "id"
	return s, nil
}

func (svc Service) ProblemServiceSignatureXXX() interface{} {
	return svc
}

func (svc Service) ListProblems(c controller.MContext) {
	page, pageSize, ok := ginhelper.RosolvePageVariable(c)
	if !ok {
		return
	}

	problems, err := svc.db.Find(page, pageSize)
	if ginhelper.MaybeSelectError(c, problems, err) {
		return
	}

	c.JSON(http.StatusOK, api.SerializeListProblemsReply(types.CodeOK, problems))
}

func (svc Service) CountProblem(c controller.MContext) {
	cnt, err := svc.db.Count()
	if ginhelper.MaybeSelectError(c, &cnt, err) {
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
	if !ginhelper.BindRequest(c, req) {
		return
	}

	var p = new(problem.Problem)
	p.Title = req.Title

	cc := ginhelper.GetCustomFields(c)
	p.AuthorID = uint(cc.UID)
	p.Description = "default"

	aff, err := svc.db.Create(p)
	if !ginhelper.CreateObjWithTip(c, aff, err, "problem") {
		return
	}

	// todo: problem desc
	//var problemDesc = model.NewProblemDesc(problem.ID, "default", []byte(req.Description))
	//
	//
	//if !ginhelper.CreateObjWithTip(c, problemDesc) {
	//	return
	//}

	//if err := problemDesc.Save(); err != nil {
	//	c.AbortWithStatusJSON(http.StatusInternalServerError, ginhelper.ErrorSerializer{
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
	panic("implement me")
}

func (svc Service) PutProblem(c controller.MContext) {
	panic("implement me")
}

func (svc Service) ChangeTemplateName(c controller.MContext) {
	panic("implement me")
}

func (svc Service) PostTemplate(c controller.MContext) {
	panic("implement me")
}

func (svc Service) GetTemplate(c controller.MContext) {
	panic("implement me")
}

func (svc Service) PutTemplate(c controller.MContext) {
	panic("implement me")
}

func (svc Service) DeleteTemplate(c controller.MContext) {
	panic("implement me")
}

func (svc Service) DeleteProblem(c controller.MContext) {
	panic("implement me")
}
