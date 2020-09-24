package inner_control

import (
	"github.com/Myriad-Dreamin/boj-v6/api"
)

type InnerGroupService interface {
	GroupServiceSignatureXXX() interface{}
	ListGroup(req *api.ListGroupRequest) (*api.ListGroupReply, error)
	CountGroup(req *api.CountGroupRequest) (*api.CountGroupReply, error)
	PostGroup(req *api.PostGroupRequest) (*api.PostGroupReply, error)
	PutGroupOwner(req *api.PutGroupOwnerRequest) (*api.PutGroupOwnerReply, error)
	GetGroupMembers(req *api.GetGroupMembersRequest) (*api.GetGroupMembersReply, error)
	PostGroupMember(req *api.PostGroupMemberRequest) (*api.PostGroupMemberReply, error)
	GetGroup(req *api.GetGroupRequest) (*api.GetGroupReply, error)
	PutGroup(req *api.PutGroupRequest) (*api.PutGroupReply, error)
	DeleteGroup(req *api.DeleteGroupRequest) (*api.DeleteGroupReply, error)
}
