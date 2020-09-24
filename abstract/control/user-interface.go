package control

import (
	"github.com/Myriad-Dreamin/minimum-lib/controller"
)

type UserService interface {
	UserServiceSignatureXXX() interface{}
	ListUser(c controller.MContext)
	CountUser(c controller.MContext)
	Register(c controller.MContext)
	LoginUser(c controller.MContext)
	RefreshToken(c controller.MContext)
	BindEmail(c controller.MContext)
	ChangePassword(c controller.MContext)
	InspectUser(c controller.MContext)
	GetUser(c controller.MContext)
	PutUser(c controller.MContext)
	DeleteUser(c controller.MContext)
}
