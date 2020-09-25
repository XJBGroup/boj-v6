package problem

import (
	"github.com/Myriad-Dreamin/boj-v6/api"
	"github.com/Myriad-Dreamin/boj-v6/app/snippet"
	"github.com/Myriad-Dreamin/boj-v6/lib/serial"
	"github.com/Myriad-Dreamin/boj-v6/types"
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	"net/http"
	"path/filepath"
	"strconv"
)

func (svc Service) ProblemFSStat(c controller.MContext) {
	var req api.ProblemFSStatRequest
	id, ok := svc.BindProblemFSRequest(c, &req)
	if !ok {
		return
	}

	path := filepath.Join(svc.cfg.PathConfig.ProblemPath, strconv.Itoa(int(id)), req.Path)

	if stat, err := svc.filesystem.Stat(path); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, &serial.ErrorSerializer{
			Code:   types.CodeStatError,
			ErrorS: err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, api.SerializeProblemFSStatReply(types.CodeOK,
			api.SerializeProblemFSStatInnerReply(
				stat.Name(), stat.Size(), stat.IsDir(), stat.ModTime())))
	}
}

func (svc Service) ProblemFSRead(c controller.MContext) {
	var req api.ProblemFSReadRequest
	id, ok := svc.BindProblemFSRequest(c, &req)
	if !ok {
		return
	}

	c.File(svc.cfg.PathConfig.ProblemPath + strconv.Itoa(int(id)) + req.Path)
}

func (svc Service) ProblemFSWrite(c controller.MContext) {
	var req api.ProblemFSWriteRequest
	id, ok := svc.BindProblemFSRequest(c, &req)
	if !ok {
		return
	}
	path := filepath.Join(svc.cfg.PathConfig.ProblemPath, strconv.Itoa(int(id)), req.Path)

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
	c.JSON(http.StatusOK, api.SerializeProblemFSWriteReply(types.CodeOK))
}

func (svc Service) ProblemFSRemove(c controller.MContext) {
	var req api.ProblemFSRemoveRequest
	id, ok := svc.BindProblemFSRequest(c, &req)
	if !ok {
		return
	}

	path := filepath.Join(svc.cfg.PathConfig.ProblemPath, strconv.Itoa(int(id)), req.Path)

	if err := svc.filesystem.Remove(path); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, &serial.ErrorSerializer{
			Code:   types.CodeStatError,
			ErrorS: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, snippet.ResponseOK)
}
