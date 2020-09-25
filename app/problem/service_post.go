package problem

import (
	"github.com/Myriad-Dreamin/boj-v6/abstract/problem"
	"github.com/Myriad-Dreamin/boj-v6/api"
	"github.com/Myriad-Dreamin/boj-v6/app/snippet"
	"github.com/Myriad-Dreamin/boj-v6/lib/serial"
	"github.com/Myriad-Dreamin/boj-v6/types"
	problemconfig "github.com/Myriad-Dreamin/boj-v6/types/problem-config"
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

func (svc *Service) SaveConfigurationToFileSystem(p *problem.Problem, c *problemconfig.ProblemConfig) error {

	// Step: Create Problem Directory

	var path = filepath.Join(svc.cfg.PathConfig.ProblemPath, strconv.Itoa(int(p.ID)))
	if _, err := svc.filesystem.Stat(path); err != nil {
		if !os.IsNotExist(err) {
			return err
		}

		// todo: 处于安全考虑，需要合理商量控制一下perm
		err := svc.filesystem.Mkdir(path, 0770)
		if err != nil {
			return err
		}
	} // else { } // already exists

	// Step: Save Configuration
	return problemconfig.SaveFS(svc.filesystem, c, filepath.Join(path, "problem-config"))
}

func (svc *Service) PostProblem(c controller.MContext) {

	// Step: Bind Request

	var req = new(api.PostProblemRequest)
	req.Config = problemconfig.DefaultProblemConfig()
	if !snippet.BindRequest(c, req) {
		return
	}

	// Step: Initialize Fields

	cc := snippet.GetCustomFields(c)

	var p = &problem.Problem{
		Title:          req.Title,
		AuthorID:       cc.UID,
		DescriptionRef: "default",
	}

	// Step: Update Database

	aff, err := svc.db.Create(p)
	if !snippet.CreateObjWithTip(c, svc.db.UnwrapError, aff, err, "problem") {
		return
	}

	var pd = svc.descDB.NewProblemDesc(p.ID, p.DescriptionRef, []byte(req.Description))

	aff, err = svc.descDB.Create(pd)
	if !snippet.CreateObjWithTip(c, svc.descDB.UnwrapError, aff, err, "problem_desc") {
		return
	}

	if err = svc.descDB.SaveDesc(pd); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, serial.ErrorSerializer{
			Code:   types.CodeInsertError,
			ErrorS: err.Error(),
		})
		return
	}

	// Step: Create Problem Directory and Save Problem Configuration

	if err = svc.SaveConfigurationToFileSystem(p, req.Config); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, serial.ErrorSerializer{
			Code:   types.CodeProblemSaveConfigurationError,
			ErrorS: err.Error(),
		})
		return
	}

	// Step: Reply

	c.JSON(http.StatusOK, &api.PostProblemReply{
		Code: types.CodeOK,
		Data: api.SerializePostProblemData(p),
	})
}
