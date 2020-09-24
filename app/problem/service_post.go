package problem

import (
	"github.com/Myriad-Dreamin/boj-v6/abstract/problem"
	"github.com/Myriad-Dreamin/boj-v6/api"
	"github.com/Myriad-Dreamin/boj-v6/app/snippet"
	"github.com/Myriad-Dreamin/boj-v6/lib/serial"
	"github.com/Myriad-Dreamin/boj-v6/types"
	problemconfig "github.com/Myriad-Dreamin/boj-v6/types/problem-config"
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

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
	if !snippet.CreateObjWithTip(c, svc.db.UnwrapError, aff, err, "problem") {
		return
	}

	var pd = svc.descDB.NewProblemDesc(p.ID, p.DescriptionRef, []byte(req.Description))

	aff, err = svc.descDB.Create(pd)
	if !snippet.CreateObjWithTip(c, svc.descDB.UnwrapError, aff, err, "problem_desc") {
		return
	}

	if err := svc.descDB.SaveDesc(pd); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, serial.ErrorSerializer{
			Code:   types.CodeInsertError,
			ErrorS: err.Error(),
		})
	}

	var path = filepath.Join(svc.cfg.PathConfig.ProblemPath, strconv.Itoa(int(p.ID)))
	if _, err := os.Stat(path); err != nil {
		if !os.IsExist(err) {
			err = os.Mkdir(path, 0770)
			if err != nil {
				_ = c.AbortWithError(http.StatusInternalServerError, err)
				return
			}
		} else {
			c.AbortWithStatusJSON(http.StatusOK, serial.ErrorSerializer{
				Code:   types.CodeFSExecError,
				ErrorS: err.Error(),
			})
			return
		}
	}

	configPath := filepath.Join(path, "/problem-config")

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
		Data: api.PostProblemData{Id: p.ID},
	})
}
