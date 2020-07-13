package auth

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
}

func (svc Service) AuthServiceSignatureXXX() interface{} {
	return svc
}

func NewService(m module.Module) (*Service, error) {
	s := new(Service)
	s.enforcer = m.Require(config.ModulePath.Provider.Model).(*provider.DB).Enforcer()
	s.logger = m.Require(config.ModulePath.Global.Logger).(external.Logger)

	return s, nil
}

func (svc Service) AddPolicy(c controller.MContext) {
	panic("implement me")
}

func (svc Service) RemovePolicy(c controller.MContext) {
	panic("implement me")
}

func (svc Service) HasPolicy(c controller.MContext) {
	panic("implement me")
}

func (svc Service) AddGroupingPolicy(c controller.MContext) {
	panic("implement me")
}

func (svc Service) RemoveGroupingPolicy(c controller.MContext) {
	panic("implement me")
}

func (svc Service) HasGroupingPolicy(c controller.MContext) {
	panic("implement me")
}
