package inner_control

import (
	"github.com/Myriad-Dreamin/boj-v6/api"
	"github.com/Myriad-Dreamin/minimum-lib/controller"
)

type InnerAuthService interface {
	AuthServiceSignatureXXX() interface{}
	AddPolicy(c controller.MContext, req *api.AddPolicyRequest) (*api.AddPolicyReply, error)
	RemovePolicy(c controller.MContext, req *api.RemovePolicyRequest) (*api.RemovePolicyReply, error)
	HasPolicy(c controller.MContext, req *api.HasPolicyRequest) (*api.HasPolicyReply, error)
	AddGroupingPolicy(c controller.MContext, req *api.AddGroupingPolicyRequest) (*api.AddGroupingPolicyReply, error)
	RemoveGroupingPolicy(c controller.MContext, req *api.RemoveGroupingPolicyRequest) (*api.RemoveGroupingPolicyReply, error)
	HasGroupingPolicy(c controller.MContext, req *api.HasGroupingPolicyRequest) (*api.HasGroupingPolicyReply, error)
}
