package submission

import (
	"github.com/Myriad-Dreamin/boj-v6/abstract/problem"
	"github.com/Myriad-Dreamin/boj-v6/abstract/submission"
	"github.com/Myriad-Dreamin/boj-v6/abstract/user"
	user_problem2 "github.com/Myriad-Dreamin/boj-v6/abstract/user_problem"
	"github.com/Myriad-Dreamin/boj-v6/api"
	"github.com/Myriad-Dreamin/boj-v6/config"
	"github.com/Myriad-Dreamin/boj-v6/external"
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	"github.com/Myriad-Dreamin/minimum-lib/module"
)

type InnerService struct {
	db                 submission.DB
	problemDB          problem.DB
	userDB             user.DB
	userTriedProblemDB user_problem2.TriedDB
	logger             external.Logger
	dispatcher         submission.Dispatcher
	cfg                *config.ServerConfig
	key                string
	problemKey         string
}

func NewInnerService(m module.Module) (*InnerService, error) {
	s := new(InnerService)
	s.db = m.RequireImpl(new(submission.DB)).(submission.DB)
	s.problemDB = m.RequireImpl(new(problem.DB)).(problem.DB)
	s.userDB = m.RequireImpl(new(user.DB)).(user.DB)
	s.logger = m.RequireImpl(new(external.Logger)).(external.Logger)
	s.cfg = m.RequireImpl(new(*config.ServerConfig)).(*config.ServerConfig)
	s.userTriedProblemDB = m.RequireImpl(new(user_problem2.TriedDB)).(user_problem2.TriedDB)
	s.dispatcher = m.RequireImpl(new(submission.Dispatcher)).(submission.Dispatcher)

	s.problemKey = "pid"
	s.key = "sid"
	return s, nil
}

func (i InnerService) SubmissionServiceSignatureXXX() interface{} {
	panic("implement me")
}

func (i InnerService) ListSubmission(c controller.MContext, req *api.ListSubmissionRequest) (*api.ListSubmissionReply, error) {
	panic("implement me")
}

func (i InnerService) CountSubmission(c controller.MContext, req *api.CountSubmissionRequest) (*api.CountSubmissionReply, error) {
	panic("implement me")
}

func (i InnerService) PostSubmission(c controller.MContext, req *api.PostSubmissionRequest) (*api.PostSubmissionReply, error) {
	panic("implement me")
}

func (i InnerService) GetContent(c controller.MContext, req *api.GetContentRequest) (*api.GetContentReply, error) {
	panic("implement me")
}

func (i InnerService) GetSubmission(c controller.MContext, req *api.GetSubmissionRequest) (*api.GetSubmissionReply, error) {
	panic("implement me")
}

func (i InnerService) DeleteSubmission(c controller.MContext, req *api.DeleteSubmissionRequest) (*api.DeleteSubmissionReply, error) {
	panic("implement me")
}
