package contest

import (
	"github.com/Myriad-Dreamin/boj-v6/app/provider"
	"github.com/Myriad-Dreamin/boj-v6/config"
	"github.com/Myriad-Dreamin/boj-v6/external"
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	"github.com/Myriad-Dreamin/minimum-lib/module"
)

type Service struct {
	enforcer *provider.Enforcer
	logger   external.Logger
	key      string
}

func (s Service) ChangeContestDescriptionRef(c controller.MContext) {
	panic("implement me")
}

func (s Service) PostContestDesc(c controller.MContext) {
	panic("implement me")
}

func (s Service) GetContestDesc(c controller.MContext) {
	panic("implement me")
}

func (s Service) PutContestDesc(c controller.MContext) {
	panic("implement me")
}

func (s Service) DeleteContestDesc(c controller.MContext) {
	panic("implement me")
}

func (s Service) GetContestProblem(c controller.MContext) {
	panic("implement me")
}

func (s Service) PutContestProblem(c controller.MContext) {
	panic("implement me")
}

func (s Service) DeleteContestProblem(c controller.MContext) {
	panic("implement me")
}

func (s Service) ListContestProblems(c controller.MContext) {
	panic("implement me")
}

func (s Service) CountContestProblem(c controller.MContext) {
	panic("implement me")
}

func (s Service) PostContestProblem(c controller.MContext) {
	panic("implement me")
}

func (s Service) ContestServiceSignatureXXX() interface{} {
	panic("implement me")
}

func (s Service) ListContests(c controller.MContext) {
	panic("implement me")
}

func (s Service) CountContest(c controller.MContext) {
	panic("implement me")
}

func (s Service) PostContest(c controller.MContext) {
	panic("implement me")
}

func (s Service) GetContest(c controller.MContext) {
	panic("implement me")
}

func (s Service) PutContest(c controller.MContext) {
	panic("implement me")
}

func (s Service) Delete(c controller.MContext) {
	panic("implement me")
}

func NewService(m module.Module) (*Service, error) {
	s := new(Service)
	s.enforcer = m.Require(config.ModulePath.Provider.Model).(*provider.DB).Enforcer()
	s.logger = m.Require(config.ModulePath.Global.Logger).(external.Logger)

	s.key = "id"
	return s, nil
}
