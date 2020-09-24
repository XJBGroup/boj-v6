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
		c.JSON(http.StatusOK, serial.ErrorSerializer{
			Code:   types.CodeSubmissionCodeTypeUnknownError,
			ErrorS: "code type unknown",
		})
		return
	}

	// Step: Save Code to Filesystem

	err = svc.SaveCodeToFileSystem(&s, req.Code)
	if err != nil {
		snippet.DoReport(c, err)
	}

	// Step: Update Database

	cc := snippet.GetCustomFields(c)
	s.UserID = cc.UID

	s.Status = types.StatusWaitingForJudge
	s.ProblemID = pid
	s.CodeLength = len(req.Code)
	s.Shared = req.Shared
	s.Information = req.Information

	aff, err := svc.db.Create(&s)
	if snippet.CreateObj(c, aff, err) {
		//cr.Submissionr.PushTask(code)
		c.JSON(http.StatusOK, api.PostSubmissionReply{
			Code: types.CodeOK,
			Data: api.SerializePostSubmissionData(&s),
		})
	}

	// todo append
	//user := new(user.User)
	//user.ID = s.UserID
	//if _, err = user.TriedProblems().Append(p); err != nil {
	//	srv.logger.Debug("update s failed", "error", err)
	//}

	// step: Fire Event
	// todo: fire event
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
