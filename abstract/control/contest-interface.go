package control

import (
	"github.com/Myriad-Dreamin/minimum-lib/controller"
)

type ContestService interface {
	ContestServiceSignatureXXX() interface{}
	ListContests(c controller.MContext)
	CountContest(c controller.MContext)
	PostContest(c controller.MContext)
	ListContestProblems(c controller.MContext)
	CountContestProblem(c controller.MContext)
	PostContestProblem(c controller.MContext)
	ChangeContestTemplateName(c controller.MContext)
	PostContestTemplate(c controller.MContext)
	GetContestTemplate(c controller.MContext)
	PutContestTemplate(c controller.MContext)
	DeleteContestTemplate(c controller.MContext)
	GetContestProblem(c controller.MContext)
	PutContestProblem(c controller.MContext)
	DeleteContestProblem(c controller.MContext)
	GetContest(c controller.MContext)
	PutContest(c controller.MContext)
	Delete(c controller.MContext)
}
