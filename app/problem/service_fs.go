package problem

import (
	"github.com/Myriad-Dreamin/boj-v6/abstract/problem"
	"github.com/Myriad-Dreamin/boj-v6/app/snippet"
	problemconfig "github.com/Myriad-Dreamin/boj-v6/types/problem-config"
	"github.com/Myriad-Dreamin/minimum-lib/controller"
)

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

func (svc *Service) BindProblemFSRequest(c controller.MContext, req interface{}) (id uint, ok bool) {
	id, ok = snippet.ParseUintAndBind(c, svc.key, req)
	if !ok {
		return
	}

	ok, err := svc.db.HasByPID(id)
	if snippet.MaybeMissingError(c, ok, err) {
		ok = false
		return
	}

	return id, ok
}

//func (svc *Service) WriteTestCases(c controller.MContext) {
//	id, ok := snippet.ParseUint(c, svc.key)
//	if !ok {
//		return
//	}
//
//	ok, err := svc.db.HasByPID(id)
//	if snippet.MaybeMissingError(c, ok, err) {
//		return
//	}
//
//	path := svc.cfg.PathConfig.ProblemPath + strconv.Itoa(int(id)) + "/test/"
//
//	err = os.MkdirAll(path, 0770)
//	if err != nil {
//		c.AbortWithStatusJSON(http.StatusOK, &serial.ErrorSerializer{
//			Code:   types.CodeStatError,
//			ErrorS: err.Error(),
//		})
//		return
//	}
//
//	form, err := c.MultipartForm()
//	if err != nil {
//		c.AbortWithStatusJSON(http.StatusOK, &serial.ErrorSerializer{
//			Code: types.CodeUploadFileError,
//			ErrorS:  err.Error(),
//		})
//		return
//	}
//	files := form.File["upload"]
//	for _, file := range files {
//		if err = c.SaveUploadedFile(file, path+file.Filename); err != nil {
//			c.AbortWithStatusJSON(http.StatusOK, &serial.ErrorSerializer{
//				Code: types.CodeFSExecError,
//				ErrorS:  err.Error(),
//			})
//			return
//		}
//	}
//
//	c.JSON(http.StatusOK, gin.H{
//		"code": types.CodeOK,
//	})
//}
