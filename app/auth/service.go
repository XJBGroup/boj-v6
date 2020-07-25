package auth

import (
	"github.com/Myriad-Dreamin/boj-v6/api"
	ginhelper "github.com/Myriad-Dreamin/boj-v6/app/snippet"
	"github.com/Myriad-Dreamin/boj-v6/external"
	"github.com/Myriad-Dreamin/boj-v6/lib/serial"
	"github.com/Myriad-Dreamin/boj-v6/types"
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	"github.com/Myriad-Dreamin/minimum-lib/module"
	"net/http"
)

type Service struct {
	enforcer *external.Enforcer
	logger   external.Logger
}

func (svc Service) AuthServiceSignatureXXX() interface{} {
	return svc
}

func NewService(m module.Module) (*Service, error) {
	s := new(Service)
	s.enforcer = m.RequireImpl(new(*external.Enforcer)).(*external.Enforcer)
	s.logger = m.RequireImpl(new(external.Logger)).(external.Logger)

	return s, nil
}

func (svc Service) AddPolicy(c controller.MContext) {
	var req = new(api.AddPolicyRequest)
	if !ginhelper.BindRequest(c, req) {
		return
	}

	added, err := svc.enforcer.AddPolicy(req.Subject, req.Object, req.Action)

	if err != nil {
		c.JSON(http.StatusOK, &serial.ErrorSerializer{
			Code:  types.CodeInsertError,
			Error: err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, api.AddPolicyReply{Code: types.CodeOK, Data: added})
	}

	return
}

func (svc Service) RemovePolicy(c controller.MContext) {
	var req = new(api.RemovePolicyRequest)
	if !ginhelper.BindRequest(c, req) {
		return
	}

	deleted, err := svc.enforcer.RemovePolicy(req.Subject, req.Object, req.Action)

	if err != nil {
		c.JSON(http.StatusOK, &serial.ErrorSerializer{
			Code:  types.CodeDeleteError,
			Error: err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, api.RemovePolicyReply{Code: types.CodeOK, Data: deleted})
	}

	return
}

func (svc Service) HasPolicy(c controller.MContext) {
	var req = new(api.HasPolicyRequest)
	if !ginhelper.BindRequest(c, req) {
		return
	}
	c.JSON(http.StatusOK, api.HasPolicyReply{
		Code: types.CodeOK,
		Data: svc.enforcer.HasPolicy(req.Subject, req.Object, req.Action)})
}

func (svc Service) AddGroupingPolicy(c controller.MContext) {
	var req = new(api.AddGroupingPolicyRequest)
	if !ginhelper.BindRequest(c, req) {
		return
	}

	added, err := svc.enforcer.AddGroupingPolicy(req.Subject, req.Group)

	if err != nil {
		c.JSON(http.StatusOK, &serial.ErrorSerializer{
			Code:  types.CodeInsertError,
			Error: err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, api.AddGroupingPolicyReply{Code: types.CodeOK, Data: added})
	}
}

func (svc Service) RemoveGroupingPolicy(c controller.MContext) {
	var req = new(api.RemoveGroupingPolicyRequest)
	if !ginhelper.BindRequest(c, req) {
		return
	}

	deleted, err := svc.enforcer.RemoveGroupingPolicy(req.Subject, req.Group)

	if err != nil {
		c.JSON(http.StatusOK, &serial.ErrorSerializer{
			Code:  types.CodeDeleteError,
			Error: err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, api.RemoveGroupingPolicyReply{Code: types.CodeOK, Data: deleted})
	}

}

func (svc Service) HasGroupingPolicy(c controller.MContext) {
	var req = new(api.HasGroupingPolicyRequest)
	if !ginhelper.BindRequest(c, req) {
		return
	}
	c.JSON(http.StatusOK, api.HasGroupingPolicyReply{
		Code: types.CodeOK,
		Data: svc.enforcer.HasGroupingPolicy(req.Subject, req.Group)})
}
