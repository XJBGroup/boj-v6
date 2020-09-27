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

type Controller struct {
	enforcer *external.Enforcer
	logger   external.Logger
}

func (svc Controller) AuthControllerSignatureXXX() interface{} {
	return svc
}

func NewController(m module.Module) (*Controller, error) {
	s := new(Controller)
	s.enforcer = m.RequireImpl(new(*external.Enforcer)).(*external.Enforcer)
	s.logger = m.RequireImpl(new(external.Logger)).(external.Logger)

	return s, nil
}

func (svc Controller) AddPolicy(c controller.MContext) {
	var req = new(api.AddPolicyRequest)
	if !ginhelper.BindRequest(c, req) {
		return
	}

	added, err := svc.enforcer.AddPolicy(req.Subject, req.Object, req.Action)

	if err != nil {
		c.JSON(http.StatusOK, &serial.ErrorSerializer{
			Code:   types.CodeInsertError,
			ErrorS: err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, api.AddPolicyReply{Code: types.CodeOK, Data: added})
	}

	return
}

func (svc Controller) RemovePolicy(c controller.MContext) {
	var req = new(api.RemovePolicyRequest)
	if !ginhelper.BindRequest(c, req) {
		return
	}

	deleted, err := svc.enforcer.RemovePolicy(req.Subject, req.Object, req.Action)

	if err != nil {
		c.JSON(http.StatusOK, &serial.ErrorSerializer{
			Code:   types.CodeDeleteError,
			ErrorS: err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, api.RemovePolicyReply{Code: types.CodeOK, Data: deleted})
	}

	return
}

func (svc Controller) HasPolicy(c controller.MContext) {
	var req = new(api.HasPolicyRequest)
	if !ginhelper.BindRequest(c, req) {
		return
	}
	c.JSON(http.StatusOK, api.HasPolicyReply{
		Code: types.CodeOK,
		Data: svc.enforcer.HasPolicy(req.Subject, req.Object, req.Action)})
}

func (svc Controller) AddGroupingPolicy(c controller.MContext) {
	var req = new(api.AddGroupingPolicyRequest)
	if !ginhelper.BindRequest(c, req) {
		return
	}

	added, err := svc.enforcer.AddGroupingPolicy(req.Subject, req.Group)

	if err != nil {
		c.JSON(http.StatusOK, &serial.ErrorSerializer{
			Code:   types.CodeInsertError,
			ErrorS: err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, api.AddGroupingPolicyReply{Code: types.CodeOK, Data: added})
	}
}

func (svc Controller) RemoveGroupingPolicy(c controller.MContext) {
	var req = new(api.RemoveGroupingPolicyRequest)
	if !ginhelper.BindRequest(c, req) {
		return
	}

	deleted, err := svc.enforcer.RemoveGroupingPolicy(req.Subject, req.Group)

	if err != nil {
		c.JSON(http.StatusOK, &serial.ErrorSerializer{
			Code:   types.CodeDeleteError,
			ErrorS: err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, api.RemoveGroupingPolicyReply{Code: types.CodeOK, Data: deleted})
	}

}

func (svc Controller) HasGroupingPolicy(c controller.MContext) {
	var req = new(api.HasGroupingPolicyRequest)
	if !ginhelper.BindRequest(c, req) {
		return
	}
	c.JSON(http.StatusOK, api.HasGroupingPolicyReply{
		Code: types.CodeOK,
		Data: svc.enforcer.HasGroupingPolicy(req.Subject, req.Group)})
}
