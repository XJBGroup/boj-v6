package control

import (
	"github.com/Myriad-Dreamin/minimum-lib/controller"
)

type AuthController interface {
	AuthControllerSignatureXXX() interface{}
	AddPolicy(c controller.MContext)
	RemovePolicy(c controller.MContext)
	HasPolicy(c controller.MContext)
	AddGroupingPolicy(c controller.MContext)
	RemoveGroupingPolicy(c controller.MContext)
	HasGroupingPolicy(c controller.MContext)
}
