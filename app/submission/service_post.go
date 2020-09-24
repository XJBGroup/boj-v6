package submission

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/Myriad-Dreamin/boj-v6/abstract/submission"
	"github.com/Myriad-Dreamin/boj-v6/api"
	"github.com/Myriad-Dreamin/boj-v6/app/snippet"
	"github.com/Myriad-Dreamin/boj-v6/lib/serial"
	"github.com/Myriad-Dreamin/boj-v6/types"
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	"net/http"
	"os"
	"path/filepath"
)

func (svc Service) PostSubmission(c controller.MContext) {
	if c.IsAborted() {
		return
	}

	var req = new(api.PostSubmissionRequest)
	pid, ok := snippet.ParseUintAndBind(c, svc.problemKey, req)
	if !ok {
		return
	}

	p, err := svc.problemDB.ID(pid)
	if snippet.MaybeSelectErrorWithTip(c, p, err, "problem") {
		return
	}

	//if code.SubmissionType, ok = morm.SubmissionTypeMap[codeType]; !ok {
	//	c.JSON(http.StatusOK, gin.H{
	//		"code": SubmissionSubmissionTypeUnknown,
	//	})
	//	return
	//}

	codeHash := sha256.New()
	_, err = codeHash.Write([]byte(req.Code))
	if err != nil {
		_ = c.AbortWithError(http.StatusOK, err)
		return
	}

	s := new(submission.Submission)
	s.Hash = hex.EncodeToString(codeHash.Sum(nil))

	if has, err := svc.db.HasHash(s.Hash); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, &serial.ErrorSerializer{
			Code:  types.CodeSelectError,
			Error: err.Error(),
		})
		return
	} else if has {
		c.AbortWithStatusJSON(http.StatusOK, &serial.ErrorSerializer{
			Code:  types.CodeSubmissionUploaded,
			Error: s.Hash,
		})
		return
	}

	var path = filepath.Join(svc.cfg.PathConfig.CodePath, s.Hash)
	if _, err = os.Stat(path); err != nil && !os.IsExist(err) {
		err = os.Mkdir(path, 0777)
		if err != nil {
			_ = c.AbortWithError(http.StatusOK, err)
			return
		}
		err = os.Chmod(path, 0777)
		if err != nil {
			_ = c.AbortWithError(http.StatusOK, err)
			return
		}
	}
	path = filepath.Join(path, "main")
	if _, err = os.Stat(path); err != nil && !os.IsExist(err) {
		f, err := os.Create(path)
		if err != nil {
			_ = c.AbortWithError(http.StatusOK, err)
			return
		}
		_, err = f.WriteString(req.Code)
		_ = f.Close()
		if err != nil {
			_ = c.AbortWithError(http.StatusOK, err)
			return
		}
	}
	err = nil

	s.Status = types.StatusWaitingForJudge
	cc := snippet.GetCustomFields(c)
	s.UserID = cc.UID
	s.ProblemID = pid
	s.CodeLength = len(req.Code)
	s.Language = req.Language
	s.Shared = req.Shared
	s.Information = req.Information

	aff, err := svc.db.Create(s)
	if snippet.CreateObj(c, aff, err) {
		//cr.Submissionr.PushTask(code)
		c.JSON(http.StatusOK, api.PostSubmissionReply{
			Code: types.CodeOK,
			Data: api.SerializePostSubmissionData(s),
		})
	}

	// todo append
	//user := new(user.User)
	//user.ID = s.UserID
	//if _, err = user.TriedProblems().Append(p); err != nil {
	//	srv.logger.Debug("update s failed", "error", err)
	//}

	c.Set("s", s)
	c.Set("p", p)
	c.Set("c", req.Code)
}
