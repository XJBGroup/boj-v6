package inner_control

import (
	"github.com/Myriad-Dreamin/boj-v6/api"
)

type InnerContestService interface {
	ContestServiceSignatureXXX() interface{}
	ListContest(req *api.ListContestRequest) (*api.ListContestReply, error)
	CountContest(req *api.CountContestRequest) (*api.CountContestReply, error)
	PostContest(req *api.PostContestRequest) (*api.PostContestReply, error)
	ChangeContestProblemDescriptionRef(req *api.ChangeContestProblemDescriptionRefRequest) (*api.ChangeContestProblemDescriptionRefReply, error)
	PostContestProblemDesc(req *api.PostContestProblemDescRequest) (*api.PostContestProblemDescReply, error)
	GetContestProblemDesc(req *api.GetContestProblemDescRequest) (*api.GetContestProblemDescReply, error)
	PutContestProblemDesc(req *api.PutContestProblemDescRequest) (*api.PutContestProblemDescReply, error)
	DeleteContestProblemDesc(req *api.DeleteContestProblemDescRequest) (*api.DeleteContestProblemDescReply, error)
	GetContestProblem(req *api.GetContestProblemRequest) (*api.GetContestProblemReply, error)
	PutContestProblem(req *api.PutContestProblemRequest) (*api.PutContestProblemReply, error)
	DeleteContestProblem(req *api.DeleteContestProblemRequest) (*api.DeleteContestProblemReply, error)
	ListContestUsers(req *api.ListContestUsersRequest) (*api.ListContestUsersReply, error)
	ListContestProblem(req *api.ListContestProblemRequest) (*api.ListContestProblemReply, error)
	CountContestProblem(req *api.CountContestProblemRequest) (*api.CountContestProblemReply, error)
	PostContestProblem(req *api.PostContestProblemRequest) (*api.PostContestProblemReply, error)
	GetContest(req *api.GetContestRequest) (*api.GetContestReply, error)
	PutContest(req *api.PutContestRequest) (*api.PutContestReply, error)
	DeleteContest(req *api.DeleteContestRequest) (*api.DeleteContestReply, error)
}
