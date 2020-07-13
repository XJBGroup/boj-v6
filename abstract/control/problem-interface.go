package control

import (
	"github.com/Myriad-Dreamin/minimum-lib/controller"
)

type ProblemService interface {
	ProblemServiceSignatureXXX() interface{}
	ListProblems(c controller.MContext)
	CountProblem(c controller.MContext)
	PostProblem(c controller.MContext)
	ChangeDescriptionRef(c controller.MContext)
	PostDesc(c controller.MContext)
	GetDesc(c controller.MContext)
	PutDesc(c controller.MContext)
	DeleteDesc(c controller.MContext)
	GetProblem(c controller.MContext)
	PutProblem(c controller.MContext)
	DeleteProblem(c controller.MContext)
}
