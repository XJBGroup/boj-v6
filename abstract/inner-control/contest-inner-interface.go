package inner_control

import (
	"github.com/Myriad-Dreamin/boj-v6/api"
	"github.com/Myriad-Dreamin/minimum-lib/controller"
)

type InnerContestController interface {
	ContestControllerSignatureXXX() interface{}
	ListContest(c controller.MContext, req *api.ListContestRequest) (*api.ListContestReply, error)
	CountContest(c controller.MContext, req *api.CountContestRequest) (*api.CountContestReply, error)
	PostContest(c controller.MContext, req *api.PostContestRequest) (*api.PostContestReply, error)
	ListContestProblem(c controller.MContext, req *api.ListContestProblemRequest) (*api.ListContestProblemReply, error)
	CountContestProblem(c controller.MContext, req *api.CountContestProblemRequest) (*api.CountContestProblemReply, error)
	PostContestProblem(c controller.MContext, req *api.PostContestProblemRequest) (*api.PostContestProblemReply, error)
	ListContestProblemDesc(c controller.MContext, req *api.ListContestProblemDescRequest) (*api.ListContestProblemDescReply, error)
	CountContestProblemDesc(c controller.MContext, req *api.CountContestProblemDescRequest) (*api.CountContestProblemDescReply, error)
	ChangeContestProblemDescriptionRef(c controller.MContext, req *api.ChangeContestProblemDescriptionRefRequest) (*api.ChangeContestProblemDescriptionRefReply, error)
	PostContestProblemDesc(c controller.MContext, req *api.PostContestProblemDescRequest) (*api.PostContestProblemDescReply, error)
	GetContestProblemDesc(c controller.MContext, req *api.GetContestProblemDescRequest) (*api.GetContestProblemDescReply, error)
	PutContestProblemDesc(c controller.MContext, req *api.PutContestProblemDescRequest) (*api.PutContestProblemDescReply, error)
	DeleteContestProblemDesc(c controller.MContext, req *api.DeleteContestProblemDescRequest) (*api.DeleteContestProblemDescReply, error)
	GetContestProblem(c controller.MContext, req *api.GetContestProblemRequest) (*api.GetContestProblemReply, error)
	PutContestProblem(c controller.MContext, req *api.PutContestProblemRequest) (*api.PutContestProblemReply, error)
	DeleteContestProblem(c controller.MContext, req *api.DeleteContestProblemRequest) (*api.DeleteContestProblemReply, error)
	ListContestUsers(c controller.MContext, req *api.ListContestUsersRequest) (*api.ListContestUsersReply, error)
	GetContest(c controller.MContext, req *api.GetContestRequest) (*api.GetContestReply, error)
	PutContest(c controller.MContext, req *api.PutContestRequest) (*api.PutContestReply, error)
	DeleteContest(c controller.MContext, req *api.DeleteContestRequest) (*api.DeleteContestReply, error)
}
