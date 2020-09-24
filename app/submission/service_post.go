package submission

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"github.com/Myriad-Dreamin/boj-v6/abstract/submission"
	"github.com/Myriad-Dreamin/boj-v6/abstract/user_problem"
	"github.com/Myriad-Dreamin/boj-v6/api"
	"github.com/Myriad-Dreamin/boj-v6/app/snippet"
	"github.com/Myriad-Dreamin/boj-v6/lib/serial"
	"github.com/Myriad-Dreamin/boj-v6/types"
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func (svc Service) SaveCodeToFileSystem(s *submission.Submission, code string) (err error) {

	// Step: Compute Hash

	codeHash := sha256.New()
	_, err = codeHash.Write([]byte(code))
	if err != nil {
		return &serial.ErrorSerializer{
			Code:   types.CodeSubmissionComputeCodeHashError,
			ErrorS: err.Error(),
		}
	}

	s.Hash = hex.EncodeToString(codeHash.Sum(nil))

	// Step: Validate Existence

	if has, err := svc.db.HasHash(s.Hash); err != nil {
		return &serial.ErrorSerializer{
			Code:   types.CodeSelectError,
			ErrorS: err.Error(),
		}
	} else if has {
		return &serial.ErrorSerializer{
			Code:   types.CodeSubmissionUploaded,
			ErrorS: s.Hash,
		}
	}

	// Step: Write To Filesystem
	// because the code hash only depends on the code content, we does not need to lock mutex and do atom update

	var codePath = filepath.Join(svc.cfg.PathConfig.CodePath, s.Hash)
	var fullPath = filepath.Join(codePath, "main")

	err = WriteToFileSystem(codePath, fullPath, code)
	if err != nil {
		return &serial.ErrorSerializer{
			Code:   types.CodeSubmissionSaveCodeError,
			ErrorS: err.Error(),
			Params: []interface{}{s.Hash},
		}
	}

	// Option Step: Create Link to Submitted Code

	err = os.Link(fullPath, fullPath+types.LanguageSuffixMapping[s.Language])
	if err != nil {
		svc.logger.Error("create link to submitted code convention", "error", err)
	}

	return nil
}

func (svc Service) PostSubmission(c controller.MContext) {
	if c.IsAborted() {
		return
	}

	// Step: Bind Request

	var req = new(api.PostSubmissionRequest)
	pid, ok := snippet.ParseUintAndBind(c, svc.problemKey, req)
	if !ok {
		return
	}

	req.Pid = pid

	// Step: Validate

	p, err := svc.problemDB.ID(pid)
	if snippet.MaybeSelectErrorWithTip(c, p, err, "problem") {
		return
	}

	var s submission.Submission

	if s.Language, ok = types.LanguageTypeMapping[req.Language]; !ok {
		c.AbortWithStatusJSON(http.StatusOK, serial.ErrorSerializer{
			Code:   types.CodeSubmissionCodeTypeUnknownError,
			ErrorS: "code type unknown",
		})
		return
	}

	// Step: Save Code to Filesystem

	err = svc.SaveCodeToFileSystem(&s, req.Code)
	if snippet.DoReport(c, err) {
		return
	}

	// Step: Fill Submission Status

	cc := snippet.GetCustomFields(c)
	s.UserID = cc.UID

	s.Status = types.StatusWaitingForJudge
	s.ProblemID = pid
	s.CodeLength = len(req.Code)
	s.Shared = req.Shared
	s.Information = req.Information

	// Step: Update Database Table `submission`

	aff, err := svc.db.Create(&s)
	if !snippet.CreateObjWithTip(c, svc.db.UnwrapError, aff, err, "submission") {
		return
	}

	var rollback = func() {
		aff, err := svc.db.Delete(&s)
		if err != nil || aff == 0 {
			svc.logger.Error("create rollback error", "affect", aff,
				"error", snippet.ConvertErrorToString(err))
		}
	}

	// Step: Update Database Table `user_tried_problem`

	var rel = user_problem.UserTriedProblemRelationship{
		UserID:    s.UserID,
		ProblemID: s.ProblemID,
	}

	aff, err = svc.userTriedProblemDB.Create(&rel)
	if err != nil {
		svcCode := svc.userTriedProblemDB.UnwrapError(err)
		switch svcCode {
		case types.CodeDuplicatePrimaryKey, types.CodeUniqueConstraintFailed:
			// ignore
		default:
			rollback()
			c.AbortWithStatusJSON(http.StatusOK, serial.ErrorSerializer{
				Code:   svcCode,
				ErrorS: "create user tried problem relationship error",
			})
			return
		}
	}

	c.JSON(http.StatusOK, api.PostSubmissionReply{
		Code: types.CodeOK,
		Data: api.SerializePostSubmissionData(&s),
	})

	// step: Fire Event
	svc.dispatcher.HandlePostSubmission(context.Background(), submission.PostEvent{
		S:    s,
		Code: strings.NewReader(req.Code),
	})
}

func WriteToFileSystem(directory string, fullPath string, code string) (err error) {

	if _, err = os.Stat(directory); err != nil && !os.IsExist(err) {
		err = os.Mkdir(directory, 0777)
		if err != nil {
			return err
		}
		err = os.Chmod(directory, 0777)
		if err != nil {
			return err
		}
	}

	if _, err = os.Stat(fullPath); err != nil && !os.IsExist(err) {
		var f *os.File
		f, err = os.Create(fullPath)
		if err != nil {
			return err
		}
		_, err = f.WriteString(code)
		if err != nil {
			return err
		}
		err = f.Close()
		if err != nil {
			return err
		}
	}
	return
}
