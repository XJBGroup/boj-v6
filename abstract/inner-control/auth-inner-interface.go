package inner_control

import (
	"github.com/Myriad-Dreamin/boj-v6/api"
)

type InnerAuthService interface {
	AuthServiceSignatureXXX() interface{}
	AddPolicy(req *api.AddPolicyRequest) (*api.AddPolicyReply, error)
	RemovePolicy(req *api.RemovePolicyRequest) (*api.RemovePolicyReply, error)
	HasPolicy(req *api.HasPolicyRequest) (*api.HasPolicyReply, error)
	AddGroupingPolicy(req *api.AddGroupingPolicyRequest) (*api.AddGroupingPolicyReply, error)
	RemoveGroupingPolicy(req *api.RemoveGroupingPolicyRequest) (*api.RemoveGroupingPolicyReply, error)
	HasGroupingPolicy(req *api.HasGroupingPolicyRequest) (*api.HasGroupingPolicyReply, error)
}
