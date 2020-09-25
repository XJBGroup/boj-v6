package problem

import (
	"archive/zip"
	"encoding/json"
	"github.com/Myriad-Dreamin/boj-v6/abstract/problem"
	"github.com/Myriad-Dreamin/boj-v6/app/snippet"
	"github.com/Myriad-Dreamin/boj-v6/lib/serial"
	"github.com/Myriad-Dreamin/boj-v6/types"
	problemconfig "github.com/Myriad-Dreamin/boj-v6/types/problem-config"
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

type Adaptor struct {
	Name    string    `json:"name"`
	Size    int64     `json:"size"`
	IsDir   bool      `json:"is_dir"`
	ModTime time.Time `json:"modtime"`
}

func adaptToJson(stat os.FileInfo) *Adaptor {
	return &Adaptor{
		Name:    stat.Name(),
		Size:    stat.Size(),
		IsDir:   stat.IsDir(),
		ModTime: stat.ModTime(),
	}
}

type StatRequest struct {
	Path string `form:"path" json:"path" binding:"required"`
}

func (svc *Service) Stat(c controller.MContext) {
	var req = new(StatRequest)
	id, ok := snippet.ParseUintAndBind(c, "pid", req)
	if !ok {
		return
	}

	ok, err := svc.db.HasByPID(id)
	if snippet.MaybeMissingError(c, ok, err) {
		return
	}

	path := svc.cfg.PathConfig.ProblemPath + strconv.Itoa(int(id)) + "/" + req.Path
	var stat os.FileInfo
	if stat, err = os.Stat(path); err != nil {
		c.JSON(http.StatusOK, serial.ErrorSerializer{
			Code:   types.CodeStatError,
			ErrorS: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":   types.CodeOK,
		"status": adaptToJson(stat),
	})
}

type MkdirRequest struct {
	Path string `form:"path" json:"path" binding:"required"`
}

func (svc *Service) Mkdir(c controller.MContext) {
	var req = new(MkdirRequest)
	id, ok := snippet.ParseUintAndBind(c, "pid", req)
	if !ok {
		return
	}

	ok, err := svc.db.HasByPID(id)
	if snippet.MaybeMissingError(c, ok, err) {
		return
	}

	path := svc.cfg.PathConfig.ProblemPath + strconv.Itoa(int(id)) + "/" + req.Path

	if err = os.MkdirAll(path, 0755); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": types.CodeFSExecError,
			"err":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": types.CodeOK,
	})
}

type LsRequest struct {
	Path string `form:"path" json:"path" binding:"required"`
}

func (svc *Service) Ls(c controller.MContext) {
	var req = new(LsRequest)
	id, ok := snippet.ParseUintAndBind(c, "pid", req)
	if !ok {
		return
	}

	ok, err := svc.db.HasByPID(id)
	if snippet.MaybeMissingError(c, ok, err) {
		return
	}

	path := svc.cfg.PathConfig.ProblemPath + strconv.Itoa(int(id)) + "/" + req.Path

	var files []os.FileInfo
	if files, err = ioutil.ReadDir(path); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": types.CodeFSExecError,
			"err":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": types.CodeOK,
		"result": func() (ret []*Adaptor) {
			ret = make([]*Adaptor, 0, len(files))
			for _, stat := range files {
				ret = append(ret, adaptToJson(stat))
			}
			return
		}(),
	})
}

type RmRequest struct {
	Path string `form:"path" json:"path" binding:"required"`
}

func (svc *Service) Rm(c controller.MContext) {
	var req = new(LsRequest)
	id, ok := snippet.ParseUintAndBind(c, "pid", req)
	if !ok {
		return
	}

	ok, err := svc.db.HasByPID(id)
	if snippet.MaybeMissingError(c, ok, err) {
		return
	}

	path := svc.cfg.PathConfig.ProblemPath + strconv.Itoa(int(id)) + "/" + req.Path

	if err = os.Remove(path); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": types.CodeFSExecError,
			"err":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, snippet.ResponseOK)
}

type RmAllRRequest struct {
	Path string `form:"path" json:"path" binding:"required"`
}

func (svc *Service) RmAll(c controller.MContext) {
	var req = new(LsRequest)
	id, ok := snippet.ParseUintAndBind(c, "pid", req)
	if !ok {
		return
	}

	ok, err := svc.db.HasByPID(id)
	if snippet.MaybeMissingError(c, ok, err) {
		return
	}

	path := svc.cfg.PathConfig.ProblemPath + strconv.Itoa(int(id)) + "/" + req.Path

	if err = os.RemoveAll(path); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": types.CodeFSExecError,
			"err":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, snippet.ResponseOK)
}

type ReadRequest struct {
	ID   uint   `form:"id" json:"id" binding:"required"`
	Path string `form:"path" json:"path" binding:"required"`
}

func (svc *Service) Read(c controller.MContext) {
	var req = new(ReadRequest)
	if err := c.ShouldBind(req); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":  types.CodeBindError,
			"error": err.Error(),
		})
		return
	}

	c.File(svc.cfg.PathConfig.ProblemPath + strconv.Itoa(int(req.ID)) + req.Path)
}

func (svc *Service) ReadConfig(c controller.MContext) {
	id, ok := snippet.ParseUint(c, "pid")
	if !ok {
		return
	}

	ok, err := svc.db.HasByPID(id)
	if snippet.MaybeMissingError(c, ok, err) {
		return
	}

	path := svc.cfg.PathConfig.ProblemPath + strconv.Itoa(int(id)) + "/problem-config"
	var cfg problemconfig.ProblemConfig

	err = problemconfig.Load(&cfg, path)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": types.CodeFSExecError,
			"err":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":   types.CodeOK,
		"config": &cfg,
	})
}

func maxInt64(a, b int64) int64 {
	if a < b {
		return b
	}
	return a
}

func modifyProblem(cfg *problemconfig.ProblemConfig, problem *problem.Problem) {
	var s int64
	for _, task := range cfg.JudgeConfig.Tasks {
		s = maxInt64(task.TimeLimit, s)
	}
	problem.TimeLimit = s
	s = 0
	for _, task := range cfg.JudgeConfig.Tasks {
		s = maxInt64(task.MemoryLimit, s)
	}
	problem.MemoryLimit = s

}

type PutConfigRequest struct {
	Key   string          `form:"key" json:"key"`
	Value json.RawMessage `form:"config" json:"config" binding:"required"`
}

func (svc *Service) PutConfig(c controller.MContext) {
	var req = new(PutConfigRequest)
	id, ok := snippet.ParseUintAndBind(c, "pid", req)
	if !ok {
		return
	}

	ok, err := svc.db.HasByPID(id)
	if snippet.MaybeMissingError(c, ok, err) {
		return
	}

	path := svc.cfg.PathConfig.ProblemPath + strconv.Itoa(int(id)) + "/problem-config"
	var cfg = new(problemconfig.ProblemConfig)
	err = problemconfig.Load(cfg, path)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": types.CodeFSExecError,
			"err":  err.Error(),
		})
		return
	}
	err = cfg.Modify(req.Key, req.Value)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": types.CodeConfigModifyError,
			"err":  err.Error(),
		})
		return
	}
	err = problemconfig.Save(cfg, path)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": types.CodeFSExecError,
			"err":  err.Error(),
		})
		return
	}

	var s = new(problem.Problem)
	modifyProblem(cfg, s)

	_, err = svc.db.UpdateFields(s, []string{"time_limit", "memory_limit"})
	if !snippet.UpdateFields(c, err) {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":   types.CodeOK,
		"config": cfg,
	})
}

type PutFullConfigRequest struct {
	Value string `form:"value" json:"value" binding:"required"`
}

func (svc *Service) PutFullConfig(c controller.MContext) {
	var req = new(PutFullConfigRequest)
	id, ok := snippet.ParseUintAndBind(c, "pid", req)
	if !ok {
		return
	}

	ok, err := svc.db.HasByPID(id)
	if snippet.MaybeMissingError(c, ok, err) {
		return
	}

	path := svc.cfg.PathConfig.ProblemPath + strconv.Itoa(int(id)) + "/problem-config"
	var cfg = new(problemconfig.ProblemConfig)
	err = problemconfig.Load(cfg, path)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": types.CodeFSExecError,
			"err":  err.Error(),
		})
		return
	}
	err = cfg.Modify("", json.RawMessage(req.Value))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": types.CodeConfigModifyError,
			"err":  err.Error(),
		})
		return
	}
	err = problemconfig.Save(cfg, path)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": types.CodeFSExecError,
			"err":  err.Error(),
		})
		return
	}

	var s = new(problem.Problem)
	modifyProblem(cfg, s)

	_, err = svc.db.UpdateFields(s, []string{"time_limit", "memory_limit"})
	if !snippet.UpdateFields(c, err) {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":   types.CodeOK,
		"config": cfg,
	})
}

//
//type PutFullConfigRequest struct {
//	Value string `form:"value" json:"value" binding:"required"`
//}
//
//func (srv *Service) PutFullConfig(c controller.MContext) {
//	var req =  new(PutFullConfigRequest)
//	id, ok := snippet.ParseUintAndBind(c, "pid", req)
//	if !ok {
//		return
//	}
//
//	ok, err := srv.db.HasByPID(id)
//	if snippet.MaybeMissingError(c, ok, err) {
//		return
//	}
//
//	path := srv.problemPath + strconv.Itoa(int(id)) + "/problem-config"
//	var cfg = new(problemconfig.ProblemConfig)
//	err = problemconfig.Load(cfg, path)
//	if err != nil {
//		c.JSON(http.StatusOK, gin.H{
//			"code": types.CodeFSExecError,
//			"err":  err.Error(),
//		})
//		return
//	}
//	err = cfg.Modify("", json.RawMessage(req.Value))
//	if err != nil {
//		c.JSON(http.StatusOK, gin.H{
//			"code": types.CodeConfigModifyError,
//			"err":  err.Error(),
//		})
//		return
//	}
//	err = problemconfig.Save(cfg, path)
//	if err != nil {
//		c.JSON(http.StatusOK, gin.H{
//			"code": types.CodeFSExecError,
//			"err":  err.Error(),
//		})
//		return
//	}
//
//	c.JSON(http.StatusOK, gin.H{
//		"code":   types.CodeOK,
//		"config": cfg,
//	})
//}

type WriteRequest struct {
	Path string `form:"path" json:"path" binding:"required"`
}

func (svc *Service) Write(c controller.MContext) {
	var req = new(WriteRequest)
	id, ok := snippet.ParseUintAndBind(c, "pid", req)
	if !ok {
		return
	}

	ok, err := svc.db.HasByPID(id)
	if snippet.MaybeMissingError(c, ok, err) {
		return
	}

	file, err := c.FormFile("upload")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": types.CodeUploadFileError,
			"err":  err.Error(),
		})
		return
	}

	if err = c.SaveUploadedFile(file, svc.cfg.PathConfig.ProblemPath+strconv.Itoa(int(id))+"/"+req.Path); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": types.CodeFSExecError,
			"err":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": types.CodeOK,
	})
}

func (svc *Service) WriteConfig(c controller.MContext) {
	id, ok := snippet.ParseUint(c, "pid")
	if !ok {
		return
	}

	ok, err := svc.db.HasByPID(id)
	if snippet.MaybeMissingError(c, ok, err) {
		return
	}

	path := svc.cfg.PathConfig.ProblemPath + strconv.Itoa(int(id)) + "/problem-config"

	file, err := c.FormFile("upload")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": types.CodeUploadFileError,
			"err":  err.Error(),
		})
		return
	}

	if err = c.SaveUploadedFile(file, path); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": types.CodeFSExecError,
			"err":  err.Error(),
		})
		return
	}

	var cfg = new(problemconfig.ProblemConfig)
	err = problemconfig.Load(cfg, path)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": types.CodeUploadFileError,
			"err":  err.Error(),
		})
		return
	}

	var s = new(problem.Problem)
	modifyProblem(cfg, s)

	_, err = svc.db.UpdateFields(s, []string{"time_limit", "memory_limit"})
	if !snippet.UpdateFields(c, err) {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": types.CodeOK,
	})
}

type WritesRequest struct {
	Path string `form:"path" json:"path" binding:"required"`
}

func (svc *Service) Writes(c controller.MContext) {
	var req = new(WritesRequest)
	id, ok := snippet.ParseUintAndBind(c, "pid", req)
	if !ok {
		return
	}

	ok, err := svc.db.HasByPID(id)
	if snippet.MaybeMissingError(c, ok, err) {
		return
	}

	path := svc.cfg.PathConfig.ProblemPath + strconv.Itoa(int(id)) + "/" + req.Path + "/"

	if err = os.MkdirAll(path, 0770); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": types.CodeFSExecError,
			"err":  err.Error(),
		})
		return
	}

	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": types.CodeUploadFileError,
			"err":  err.Error(),
		})
		return
	}
	files := form.File["upload"]
	for _, file := range files {
		if err = c.SaveUploadedFile(file, path+file.Filename); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": types.CodeFSExecError,
				"err":  err.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": types.CodeOK,
	})
}

func (svc *Service) WriteTestCases(c controller.MContext) {
	id, ok := snippet.ParseUint(c, "pid")
	if !ok {
		return
	}

	ok, err := svc.db.HasByPID(id)
	if snippet.MaybeMissingError(c, ok, err) {
		return
	}

	path := svc.cfg.PathConfig.ProblemPath + strconv.Itoa(int(id)) + "/test/"

	err = os.MkdirAll(path, 0770)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": types.CodeFSExecError,
			"err":  err.Error(),
		})
		return
	}

	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": types.CodeUploadFileError,
			"err":  err.Error(),
		})
		return
	}
	files := form.File["upload"]
	for _, file := range files {
		if err = c.SaveUploadedFile(file, path+file.Filename); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": types.CodeFSExecError,
				"err":  err.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": types.CodeOK,
	})
}

type ZipRequest struct {
	Path string `form:"path" json:"path" binding:"required"`
}

func (svc *Service) Zip(c controller.MContext) {
	var req = new(ZipRequest)
	id, ok := snippet.ParseUintAndBind(c, "pid", req)
	if !ok {
		return
	}

	ok, err := svc.db.HasByPID(id)
	if snippet.MaybeMissingError(c, ok, err) {
		return
	}

	path := svc.cfg.PathConfig.ProblemPath + strconv.Itoa(int(id)) + "/" + req.Path

	file, err := c.FormFile("upload")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": types.CodeUploadFileError,
			"err":  err.Error(),
		})
		return
	}

	zipName := path + file.Filename
	if err = c.SaveUploadedFile(file, zipName); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": types.CodeFSExecError,
			"err":  err.Error(),
		})
		return
	}

	r, err := zip.OpenReader(zipName)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": types.CodeFSExecError,
			"err":  err.Error(),
		})
		return
	}

	var release = func() {
		err := r.Close()
		if err != nil {
			svc.logger.Debug("error occurs", "error", err)
		}
		err = os.Remove(zipName)
		if err != nil {
			svc.logger.Debug("error occurs", "error", err)
			return
		}
	}
	for _, file := range r.File {
		rc, err := file.Open()
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": types.CodeFSExecError,
				"err":  err.Error(),
			})
			release()
			return
		}
		filename := path + file.Name
		err = os.MkdirAll(filepath.Dir(filename), 0770)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": types.CodeFSExecError,
				"err":  err.Error(),
			})
			_ = rc.Close()
			release()
			return
		}
		w, err := os.Create(filename)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": types.CodeFSExecError,
				"err":  err.Error(),
			})
			_ = rc.Close()
			release()
			return
		}
		_, err = io.Copy(w, rc)
		_ = rc.Close()
		_ = w.Close()
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": types.CodeFSExecError,
				"err":  err.Error(),
			})
			release()
			return
		}

	}

	c.JSON(http.StatusOK, gin.H{
		"code": types.CodeOK,
	})
	release()
}
