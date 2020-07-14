package submission

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/Myriad-Dreamin/boj-v6/abstract/problem"
	"github.com/Myriad-Dreamin/boj-v6/abstract/submission"
	"github.com/Myriad-Dreamin/boj-v6/abstract/user"
	"github.com/Myriad-Dreamin/boj-v6/api"
	"github.com/Myriad-Dreamin/boj-v6/app/provider"
	"github.com/Myriad-Dreamin/boj-v6/app/snippet"
	"github.com/Myriad-Dreamin/boj-v6/config"
	"github.com/Myriad-Dreamin/boj-v6/external"
	"github.com/Myriad-Dreamin/boj-v6/lib/serial"
	"github.com/Myriad-Dreamin/boj-v6/types"
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	"github.com/Myriad-Dreamin/minimum-lib/module"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
)

type Service struct {
	db         submission.DB
	problemDB  problem.DB
	userDB     user.DB
	logger     external.Logger
	cfg        *config.ServerConfig
	key        string
	problemKey string
}

func NewService(m module.Module) (*Service, error) {
	s := new(Service)
	s.db = m.Require(config.ModulePath.Provider.Model).(*provider.DB).SubmissionDB()
	s.problemDB = m.Require(config.ModulePath.Provider.Model).(*provider.DB).ProblemDB()
	s.userDB = m.Require(config.ModulePath.Provider.Model).(*provider.DB).UserDB()
	s.logger = m.Require(config.ModulePath.Global.Logger).(external.Logger)
	s.cfg = m.Require(config.ModulePath.Global.Configuration).(*config.ServerConfig)
	s.problemKey = "pid"
	s.key = "sid"
	return s, nil
}

func (svc Service) SubmissionServiceSignatureXXX() interface{} {
	return svc
}

const (
	nonOrder        = "id desc"
	revTimeOrder    = "running_time desc, id desc"
	revTimeRevOrder = "running_time desc"
	timeOrder       = "running_time, id desc"
	timeRevOrder    = "running_time"
	revNonOrder     = ""
	revMemOrder     = "running_memory desc, id desc"
	revMemRevOrder  = "running_memory desc"
	memOrder        = "running_memory, id desc"
	memRevOrder     = "running_memory"
	notContest      = 1 << 62
)

func (svc Service) ResolveFilter(c controller.MContext) *submission.Filter {
	var req = new(api.SubmissionFilter)
	if !snippet.BindRequest(c, req) {
		return nil
	}
	var f submission.Filter
	f.Order = nonOrder
	if req.MemOrder != nil {
		if req.IdOrder != nil {
			switch {
			case *req.MemOrder && *req.IdOrder:
				f.Order = memRevOrder
			case *req.MemOrder && !*req.IdOrder:
				f.Order = memOrder
			case !*req.MemOrder && *req.IdOrder:
				f.Order = revMemRevOrder
			default:
				f.Order = revMemOrder
			}
		} else {
			switch {
			case *req.MemOrder:
				f.Order = memOrder
			default:
				f.Order = revMemOrder
			}
		}
	} else if req.TimeOrder != nil {
		if req.IdOrder != nil {
			switch {
			case *req.TimeOrder && *req.IdOrder:
				f.Order = timeRevOrder
			case *req.TimeOrder && !*req.IdOrder:
				f.Order = timeOrder
			case !*req.TimeOrder && *req.IdOrder:
				f.Order = revTimeRevOrder
			default:
				f.Order = revTimeOrder
			}
		} else {
			switch {
			case *req.TimeOrder:
				f.Order = timeOrder
			default:
				f.Order = revTimeOrder
			}
		}
	} else if req.IdOrder != nil {
		switch {
		case *req.IdOrder:
			f.Order = revNonOrder
		default:
			f.Order = nonOrder
		}
	}

	f.Page = req.Page
	f.PageSize = req.PageSize
	f.ByUser = req.ByUser
	f.OnProblem = req.OnProblem
	f.WithLanguage = req.WithLanguage
	f.HasStatus = req.HasStatus
	return &f
}

func (svc Service) ListSubmissions(c controller.MContext) {
	f := svc.ResolveFilter(c)
	if c.IsAborted() {
		return
	}

	ss, err := svc.db.Filter(f)
	if snippet.MaybeSelectError(c, ss, err) {
		return
	}

	c.JSON(http.StatusOK, api.SerializeListSubmissionsReply(types.CodeOK,
		api.PackSerializeListSubmissionReply(ss)))

	return
}

func (svc Service) CountSubmissions(c controller.MContext) {
	f := svc.ResolveFilter(c)
	if c.IsAborted() {
		return
	}

	cnt, err := svc.db.FilterCount(f)
	if snippet.MaybeCountError(c, err) {
		return
	}

	c.JSON(http.StatusOK, api.CountSubmissionsReply{
		Code: types.CodeOK,
		Data: cnt,
	})
}

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
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	s := new(submission.Submission)
	s.Hash = hex.EncodeToString(codeHash.Sum(nil))

	if has, err := svc.db.HasHash(s.Hash); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &serial.ErrorSerializer{
			Code:  types.CodeSelectError,
			Error: err.Error(),
		})
		return
	} else if has {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": types.CodeSubmissionUploaded,
			"hash": s.Hash,
		})
		return
	}

	var path = filepath.Join(svc.cfg.PathConfig.CodePath, s.Hash)
	if _, err = os.Stat(path); err != nil && !os.IsExist(err) {
		err = os.Mkdir(path, 0777)
		if err != nil {
			_ = c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		err = os.Chmod(path, 0777)
		if err != nil {
			_ = c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	}
	path += "/main.cpp"
	if _, err = os.Stat(path); err != nil && !os.IsExist(err) {
		f, err := os.Create(path)
		if err != nil {
			_ = c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		_, err = f.WriteString(req.Code)
		_ = f.Close()
		if err != nil {
			_ = c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	}
	err = nil

	s.Status = types.StatusWaitingForJudge
	cc := snippet.GetCustomFields(c)
	s.UserID = uint(cc.UID)
	s.ProblemID = pid
	s.CodeLength = len(req.Code)
	s.Language = req.Language
	s.Shared = req.Shared
	s.Information = req.Information

	aff, err := svc.db.Create(s)
	if snippet.CreateObj(c, aff, err) {
		//cr.Submissionr.PushTask(code)
		c.JSON(http.StatusOK, api.SerializePostSubmissionReply(types.CodeOK, s))
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

func (svc Service) GetContent(c controller.MContext) {
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
		c.File(filepath.Join(svc.cfg.PathConfig.CodePath, s.Hash, "main.cpp"))
	}
}

func (svc Service) GetSubmission(c controller.MContext) {
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

func (svc Service) Delete(c controller.MContext) {
	id, ok := snippet.ParseUint(c, svc.key)
	if !ok {
		return
	}
	obj, err := svc.db.ID(id)
	if snippet.MaybeSelectError(c, obj, err) {
		return
	}

	var path = filepath.Join(svc.cfg.PathConfig.CodePath, obj.Hash)
	if _, err = os.Stat(path); err == nil {
		err = os.RemoveAll(path)
		if err != nil {
			_ = c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	} else {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
	}

	aff, err := svc.db.Delete(obj)
	if snippet.DeleteObj(c, aff, err) {
		c.JSON(http.StatusOK, &snippet.ResponseOK)
	}
}
