package inner_control

import (
	"github.com/Myriad-Dreamin/boj-v6/api"
)

type InnerProblemService interface {
	ProblemServiceSignatureXXX() interface{}
	ListProblem(req *api.ListProblemRequest) (*api.ListProblemReply, error)
	CountProblem(req *api.CountProblemRequest) (*api.CountProblemReply, error)
	PostProblem(req *api.PostProblemRequest) (*api.PostProblemReply, error)
	ChangeProblemDescriptionRef(req *api.ChangeProblemDescriptionRefRequest) (*api.ChangeProblemDescriptionRefReply, error)
	PostProblemDesc(req *api.PostProblemDescRequest) (*api.PostProblemDescReply, error)
	GetProblemDesc(req *api.GetProblemDescRequest) (*api.GetProblemDescReply, error)
	PutProblemDesc(req *api.PutProblemDescRequest) (*api.PutProblemDescReply, error)
	DeleteProblemDesc(req *api.DeleteProblemDescRequest) (*api.DeleteProblemDescReply, error)
	GetProblem(req *api.GetProblemRequest) (*api.GetProblemReply, error)
	PutProblem(req *api.PutProblemRequest) (*api.PutProblemReply, error)
	DeleteProblem(req *api.DeleteProblemRequest) (*api.DeleteProblemReply, error)
}
