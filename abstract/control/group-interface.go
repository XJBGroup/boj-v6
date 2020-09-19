package control

import (
	"github.com/Myriad-Dreamin/minimum-lib/controller"
)

type GroupService interface {
	GroupServiceSignatureXXX() interface{}
	ListGroups(c controller.MContext)
	CountGroup(c controller.MContext)
	PostGroup(c controller.MContext)
	PutGroupOwner(c controller.MContext)
	GetGroupMembers(c controller.MContext)
	PostGroupMember(c controller.MContext)
	GetGroup(c controller.MContext)
	PutGroup(c controller.MContext)
	DeleteGroup(c controller.MContext)
}
