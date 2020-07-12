package control

import (
	"github.com/Myriad-Dreamin/minimum-lib/controller"
)

type ContestService interface {
	ContestServiceSignatureXXX() interface{}
	ListContests(c controller.MContext)
	CountContest(c controller.MContext)
	PostContest(c controller.MContext)
	GetContest(c controller.MContext)
	PutContest(c controller.MContext)
	Delete(c controller.MContext)
}
