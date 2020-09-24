package control

import (
	"github.com/Myriad-Dreamin/minimum-lib/controller"
)

type ProblemService interface {
	ProblemServiceSignatureXXX() interface{}
	ListProblem(c controller.MContext)
	CountProblem(c controller.MContext)
	PostProblem(c controller.MContext)
	ListProblemDesc(c controller.MContext)
	CountProblemDesc(c controller.MContext)
	ChangeProblemDescriptionRef(c controller.MContext)
	PostProblemDesc(c controller.MContext)
	GetProblemDesc(c controller.MContext)
	PutProblemDesc(c controller.MContext)
	DeleteProblemDesc(c controller.MContext)
	GetProblem(c controller.MContext)
	PutProblem(c controller.MContext)
	DeleteProblem(c controller.MContext)
}
