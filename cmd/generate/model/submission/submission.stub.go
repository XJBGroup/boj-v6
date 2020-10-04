package submission

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	problem "github.com/Myriad-Dreamin/boj-v6/abstract/problem"
	"github.com/Myriad-Dreamin/boj-v6/abstract/submission"
	"github.com/Myriad-Dreamin/boj-v6/abstract/user"
	"github.com/Myriad-Dreamin/boj-v6/abstract/user_problem"
	"github.com/Myriad-Dreamin/boj-v6/api"
	"github.com/Myriad-Dreamin/boj-v6/app/snippet"
	"github.com/Myriad-Dreamin/boj-v6/cmd/generate/stub"
	"github.com/Myriad-Dreamin/boj-v6/config"
	"github.com/Myriad-Dreamin/boj-v6/external"
	"github.com/Myriad-Dreamin/boj-v6/lib/serial"
	"github.com/Myriad-Dreamin/boj-v6/types"
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	"github.com/Myriad-Dreamin/minimum-lib/module"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var _ = snippet.AuthenticatePassword

type Controller struct {
	stub.StubVariables
	stub.Stub

	db                 submission.DB
	problemDB          problem.DB
	userDB             user.DB
	userTriedProblemDB user_problem.TriedDB
	logger             external.Logger
	cfg                *config.ServerConfig
	key                string
	problemKey         string
	dispatcher         submission.Dispatcher

	resolver module.Module
}

func (ctrl *Controller) PostSubmission(c controller.MContext) {
	ctrl.AbortIf(c.IsAborted())

	// Step: Bind Request

	var request = new(api.PostSubmissionRequest)
	pid := ctrl.GetID()
	ctrl.Bind(request)

	// Step: Validate
	_, ctrl.Err = ctrl.problemDB.ID(pid)

	var s submission.Submission

	s.Language, ctrl.Ok = types.LanguageTypeMapping[request.Language]

	// Step: Save Code to Filesystem

	ctrl.Err = ctrl.SaveCodeToFileSystem(&s, request.Code)

	// Step: Fill Submission Status

	cc := snippet.GetCustomFields(c)
	s.UserID = cc.UID

	s.Status = types.StatusWaitingForJudge
	s.ProblemID = pid
	s.CodeLength = len(request.Code)
	s.Shared = request.Shared
	s.Information = request.Information

	// Step: Update Database Table `submission`

	ctrl.Int64, ctrl.Err = ctrl.db.Create(&s)
	ctrl.AbortIf(ctrl.Int64 == 0, types.CodeInsertError)

	// Step: Update Database Table `user_tried_problem`

	var rel = user_problem.UserTriedProblemRelationship{
		UserID:    s.UserID,
		ProblemID: s.ProblemID,
	}

	ctrl.Int64, ctrl.Err = ctrl.userTriedProblemDB.Create(&rel)
	ctrl.OnErr(ctrl.Err, func(err error) error {
		svcCode := ctrl.userTriedProblemDB.UnwrapError(err)
		switch svcCode {
		case types.CodeDuplicatePrimaryKey, types.CodeUniqueConstraintFailed:
			return nil
		default:
			aff, err2 := ctrl.db.Delete(&s)
			if err2 != nil || aff == 0 {
				ctrl.logger.Error("create rollback error", "affect", aff,
					"error", snippet.ConvertErrorToString(err))
			}

			return err
		}
	})

	c.JSON(http.StatusOK, api.PostSubmissionReply{
		Code: types.CodeOK,
		Data: api.SerializePostSubmissionData(&s),
	})

	// step: Fire Event
	ctrl.dispatcher.HandlePostSubmission(context.Background(), submission.PostEvent{
		S:    s,
		Code: strings.NewReader(request.Code),
	})
}


func (ctrl *Controller) SaveCodeToFileSystem(s *submission.Submission, code string) (err error) {

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

	if has, err := ctrl.db.HasHash(s.Hash); err != nil {
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

	var codePath = filepath.Join(ctrl.cfg.PathConfig.CodePath, s.Hash)
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
		ctrl.logger.Error("create link to submitted code convention", "error", err)
	}

	return nil
}

func (ctrl *Controller) SubmissionControllerSignatureXXX() interface{} {
	panic("implement me")
}

func (ctrl *Controller) ListSubmission(c controller.MContext) {
	panic("implement me")
}

func (ctrl *Controller) CountSubmission(c controller.MContext) {
	panic("implement me")
}

func (ctrl *Controller) GetSubmissionContent(c controller.MContext) {
	panic("implement me")
}

func (ctrl *Controller) GetSubmission(c controller.MContext) {

	id := ctrl.GetID()
	var request *api.GetSubmissionRequest
	ctrl.Bind(request)

	var response *api.GetSubmissionReply
	ctrl.Serve(id, request, response)
	c.JSON(http.StatusOK, response)
}

func (ctrl *Controller) DeleteSubmission(c controller.MContext) {
	panic("implement me")
}
