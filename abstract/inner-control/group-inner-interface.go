package inner_control

import (
	"github.com/Myriad-Dreamin/boj-v6/api"
	"github.com/Myriad-Dreamin/minimum-lib/controller"
)

type InnerGroupController interface {
	GroupControllerSignatureXXX() interface{}
	ListGroup(c controller.MContext, req *api.ListGroupRequest) (*api.ListGroupReply, error)
	CountGroup(c controller.MContext, req *api.CountGroupRequest) (*api.CountGroupReply, error)
	PostGroup(c controller.MContext, req *api.PostGroupRequest) (*api.PostGroupReply, error)
	PutGroupOwner(c controller.MContext, req *api.PutGroupOwnerRequest) (*api.PutGroupOwnerReply, error)
	GetGroupMembers(c controller.MContext, req *api.GetGroupMembersRequest) (*api.GetGroupMembersReply, error)
	PostGroupMember(c controller.MContext, req *api.PostGroupMemberRequest) (*api.PostGroupMemberReply, error)
	GetGroup(c controller.MContext, req *api.GetGroupRequest) (*api.GetGroupReply, error)
	PutGroup(c controller.MContext, req *api.PutGroupRequest) (*api.PutGroupReply, error)
	DeleteGroup(c controller.MContext, req *api.DeleteGroupRequest) (*api.DeleteGroupReply, error)
}
