package group

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

func (s Service) GroupServiceSignatureXXX() interface{} {
	panic("implement me")
}

func (s Service) ListGroups(c controller.MContext) {
	panic("implement me")
}

func (s Service) CountGroup(c controller.MContext) {
	panic("implement me")
}

func (s Service) PostGroup(c controller.MContext) {
	panic("implement me")
}

func (s Service) GetGroup(c controller.MContext) {
	panic("implement me")
}

func (s Service) PutGroup(c controller.MContext) {
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
