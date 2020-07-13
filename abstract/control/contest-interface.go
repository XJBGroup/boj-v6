package control

import (
	"github.com/Myriad-Dreamin/minimum-lib/controller"
)

type ContestService interface {
	ContestServiceSignatureXXX() interface{}
	ListContests(c controller.MContext)
	CountContest(c controller.MContext)
	PostContest(c controller.MContext)
	ChangeContestDescriptionRef(c controller.MContext)
	PostContestDesc(c controller.MContext)
	GetContestDesc(c controller.MContext)
	PutContestDesc(c controller.MContext)
	DeleteContestDesc(c controller.MContext)
	GetContestProblem(c controller.MContext)
	PutContestProblem(c controller.MContext)
	DeleteContestProblem(c controller.MContext)
	ListContestProblems(c controller.MContext)
	CountContestProblem(c controller.MContext)
	PostContestProblem(c controller.MContext)
	GetContest(c controller.MContext)
	PutContest(c controller.MContext)
	Delete(c controller.MContext)
}
