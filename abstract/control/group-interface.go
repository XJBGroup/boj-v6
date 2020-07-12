package control

import (
	"github.com/Myriad-Dreamin/minimum-lib/controller"
)

type GroupService interface {
	GroupServiceSignatureXXX() interface{}
	ListGroups(c controller.MContext)
	CountGroup(c controller.MContext)
	PostGroup(c controller.MContext)
	GetGroup(c controller.MContext)
	PutGroup(c controller.MContext)
	Delete(c controller.MContext)
}
