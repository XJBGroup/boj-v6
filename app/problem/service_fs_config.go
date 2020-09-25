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
	"path/filepath"
	"strconv"
)

func (svc Service) ProblemFSReadConfig(c controller.MContext) {
	var req api.ProblemFSReadConfigRequest
	id, ok := svc.BindProblemFSRequest(c, &req)
	if !ok {
		return
	}
	path := filepath.Join(svc.cfg.PathConfig.ProblemPath, strconv.Itoa(int(id)), "problem-config")

	var cfg problemconfig.ProblemConfig
	err := problemconfig.LoadFS(svc.filesystem, &cfg, path)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, &serial.ErrorSerializer{
			Code:   types.CodeStatError,
			ErrorS: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK,
		api.SerializeProblemFSReadConfigReply(types.CodeOK, &cfg))
}

func (svc Service) ProblemFSWriteConfig(c controller.MContext) {
	var req api.ProblemFSWriteConfigRequest
	id, ok := svc.BindProblemFSRequest(c, &req)
	if !ok {
		return
	}
	path := filepath.Join(svc.cfg.PathConfig.ProblemPath, strconv.Itoa(int(id)), "problem-config")

	file, err := c.FormFile("upload")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, &serial.ErrorSerializer{
			Code:   types.CodeUploadFileError,
			ErrorS: err.Error(),
		})
		return
	}

	if err = c.SaveUploadedFile(file, path); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, &serial.ErrorSerializer{
			Code:   types.CodeStatError,
			ErrorS: err.Error(),
		})
		return
	}

	var cfg = new(problemconfig.ProblemConfig)
	err = problemconfig.LoadFS(svc.filesystem, cfg, path)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, &serial.ErrorSerializer{
			Code:   types.CodeUploadFileError,
			ErrorS: err.Error(),
		})
		return
	}

	var s = new(problem.Problem)
	modifyProblem(cfg, s)

	_, err = svc.db.UpdateFields(s, []string{"time_limit", "memory_limit"})
	if !snippet.UpdateFields(c, err) {
		return
	}

	c.JSON(http.StatusOK, api.SerializeProblemFSWriteConfigReply(types.CodeOK))
}

func (svc Service) ProblemFSPutConfig(c controller.MContext) {
	var req api.ProblemFSPutConfigRequest
	id, ok := svc.BindProblemFSRequest(c, &req)
	if !ok {
		return
	}
	path := filepath.Join(svc.cfg.PathConfig.ProblemPath, strconv.Itoa(int(id)), "problem-config")

	var cfg problemconfig.ProblemConfig
	err := problemconfig.LoadFS(svc.filesystem, &cfg, path)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, &serial.ErrorSerializer{
			Code:   types.CodeStatError,
			ErrorS: err.Error(),
		})
		return
	}
	err = cfg.Modify(req.Key, req.Value)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, &serial.ErrorSerializer{
			Code:   types.CodeConfigModifyError,
			ErrorS: err.Error(),
		})
		return
	}
	err = problemconfig.SaveFS(svc.filesystem, &cfg, path)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, &serial.ErrorSerializer{
			Code:   types.CodeStatError,
			ErrorS: err.Error(),
		})
		return
	}

	var s = new(problem.Problem)
	modifyProblem(&cfg, s)

	_, err = svc.db.UpdateFields(s, []string{"time_limit", "memory_limit"})
	if !snippet.UpdateFields(c, err) {
		return
	}

	c.JSON(http.StatusOK, api.SerializeProblemFSPutConfigReply(types.CodeOK, &cfg))
}
