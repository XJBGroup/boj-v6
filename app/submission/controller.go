package submission

import (
	"github.com/Myriad-Dreamin/boj-v6/abstract/problem"
	"github.com/Myriad-Dreamin/boj-v6/abstract/submission"
	"github.com/Myriad-Dreamin/boj-v6/abstract/user"
	user_problem2 "github.com/Myriad-Dreamin/boj-v6/abstract/user_problem"
	"github.com/Myriad-Dreamin/boj-v6/api"
	"github.com/Myriad-Dreamin/boj-v6/app/snippet"
	"github.com/Myriad-Dreamin/boj-v6/config"
	"github.com/Myriad-Dreamin/boj-v6/external"
	"github.com/Myriad-Dreamin/boj-v6/types"
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	"github.com/Myriad-Dreamin/minimum-lib/module"
	"net/http"
	"os"
	"path/filepath"
)

type Controller struct {
	db                 submission.DB
	problemDB          problem.DB
	userDB             user.DB
	userTriedProblemDB user_problem2.TriedDB
	logger             external.Logger
	cfg                *config.ServerConfig
	key                string
	problemKey         string
	dispatcher         submission.Dispatcher

	//inner inner_control.InnerSubmissionService
}

func NewController(m module.Module) (*Controller, error) {
	s := new(Controller)
	s.db = m.RequireImpl(new(submission.DB)).(submission.DB)
	s.problemDB = m.RequireImpl(new(problem.DB)).(problem.DB)
	s.userDB = m.RequireImpl(new(user.DB)).(user.DB)
	s.logger = m.RequireImpl(new(external.Logger)).(external.Logger)
	s.cfg = m.RequireImpl(new(*config.ServerConfig)).(*config.ServerConfig)
	s.userTriedProblemDB = m.RequireImpl(new(user_problem2.TriedDB)).(user_problem2.TriedDB)
	s.dispatcher = m.RequireImpl(new(submission.Dispatcher)).(submission.Dispatcher)
	s.problemKey = "pid"
	s.key = "sid"

	//s.inner = m.RequireImpl(new(inner_control.InnerSubmissionController)).(inner_control.InnerSubmissionController)
	return s, nil
}

func (svc *Controller) SubmissionControllerSignatureXXX() interface{} {
	return svc
}

func (svc *Controller) ListSubmission(c controller.MContext) {
	f := svc.ResolveFilter(c)
	if c.IsAborted() {
		return
	}

	ss, err := svc.db.Filter(f)
	if snippet.MaybeSelectError(c, ss, err) {
		return
	}

	c.JSON(http.StatusOK, api.SerializeListSubmissionReply(types.CodeOK,
		api.PackSerializeListSubmissionInnerReply(ss)))

	return
}

func (svc *Controller) CountSubmission(c controller.MContext) {
	f := svc.ResolveFilter(c)
	if c.IsAborted() {
		return
	}

	cnt, err := svc.db.FilterCount(f)
	if snippet.MaybeCountError(c, err) {
		return
	}

	c.JSON(http.StatusOK, api.CountSubmissionReply{
		Code: types.CodeOK,
		Data: cnt,
	})
}

func (svc *Controller) GetSubmission(c controller.MContext) {
	id, ok := snippet.ParseUint(c, svc.key)
	if !ok {
		return
	}
	obj, err := svc.db.ID(id)
	if snippet.MaybeSelectError(c, obj, err) {
		return
	}

	c.JSON(http.StatusOK, api.SerializeGetSubmissionReply(types.CodeOK, api.SerializeGetSubmissionInnerReply(obj)))
}

func (svc *Controller) DeleteSubmission(c controller.MContext) {
	id, ok := snippet.ParseUint(c, svc.key)
	if !ok {
		return
	}
	obj, err := svc.db.ID(id)
	if snippet.MaybeSelectError(c, obj, err) {
		return
	}

	// todo: check remove all files

	var path = filepath.Join(svc.cfg.PathConfig.CodePath, obj.Hash)
	if _, err = os.Stat(path); err == nil {
		err = os.RemoveAll(path)
		if err != nil {
			_ = c.AbortWithError(http.StatusOK, err)
			return
		}
	} else {
		_ = c.AbortWithError(http.StatusOK, err)
	}

	aff, err := svc.db.Delete(obj)
	if snippet.DeleteObj(c, aff, err) {
		c.JSON(http.StatusOK, &snippet.ResponseOK)
	}
}

func (svc *Controller) GetSubmissionContent(c controller.MContext) {
	if c.IsAborted() {
		return
	}
	id, ok := snippet.ParseUint(c, svc.key)
	if !ok {
		return
	}

	s, err := svc.db.ID(id)
	if snippet.MaybeSelectError(c, s, err) {
		return
	}

	// if s.Shared != 1
	if s.Shared == 1 {
		c.File(filepath.Join(svc.cfg.PathConfig.CodePath, s.Hash, "main"))
	}
}

func WriteToFileSystem(directory string, fullPath string, code string) (err error) {

	if _, err = os.Stat(directory); err != nil {
		if !os.IsNotExist(err) {
			return nil
		}

		// todo: 处于安全考虑，需要合理商量控制一下perm
		err = os.Mkdir(directory, 0777)
		if err != nil {
			return err
		}
	}

	if _, err = os.Stat(fullPath); err != nil {
		if !os.IsNotExist(err) {
			return nil
		}

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
