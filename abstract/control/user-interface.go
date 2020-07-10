package control

import (
	"github.com/Myriad-Dreamin/minimum-lib/controller"
)

type UserService interface {
	UserServiceSignatureXXX() interface{}
	ListUsers(c controller.MContext)
	CountUser(c controller.MContext)
	ListUsersNameLike(c controller.MContext)
	PostUser(c controller.MContext)
	LoginUser(c controller.MContext)
	PutUserContent(c controller.MContext)
	InspectUser(c controller.MContext)
	GetUser(c controller.MContext)
	PutUser(c controller.MContext)
	Delete(c controller.MContext)
}
