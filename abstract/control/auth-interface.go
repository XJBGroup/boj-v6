package control

import (
	"github.com/Myriad-Dreamin/minimum-lib/controller"
)

type AuthService interface {
	AuthServiceSignatureXXX() interface{}
	AddPolicy(c controller.MContext)
	RemovePolicy(c controller.MContext)
	HasPolicy(c controller.MContext)
	AddGroupingPolicy(c controller.MContext)
	RemoveGroupingPolicy(c controller.MContext)
	HasGroupingPolicy(c controller.MContext)
}
