package inner_control

import (
	"github.com/Myriad-Dreamin/boj-v6/api"
	"github.com/Myriad-Dreamin/minimum-lib/controller"
)

type InnerProblemService interface {
	ProblemServiceSignatureXXX() interface{}
	ListProblem(c controller.MContext, req *api.ListProblemRequest) (*api.ListProblemReply, error)
	CountProblem(c controller.MContext, req *api.CountProblemRequest) (*api.CountProblemReply, error)
	PostProblem(c controller.MContext, req *api.PostProblemRequest) (*api.PostProblemReply, error)
	CountProblemDesc(c controller.MContext, req *api.CountProblemDescRequest) (*api.CountProblemDescReply, error)
	ChangeProblemDescriptionRef(c controller.MContext, req *api.ChangeProblemDescriptionRefRequest) (*api.ChangeProblemDescriptionRefReply, error)
	PostProblemDesc(c controller.MContext, req *api.PostProblemDescRequest) (*api.PostProblemDescReply, error)
	GetProblemDesc(c controller.MContext, req *api.GetProblemDescRequest) (*api.GetProblemDescReply, error)
	PutProblemDesc(c controller.MContext, req *api.PutProblemDescRequest) (*api.PutProblemDescReply, error)
	DeleteProblemDesc(c controller.MContext, req *api.DeleteProblemDescRequest) (*api.DeleteProblemDescReply, error)
	ListProblemDesc(c controller.MContext, req *api.ListProblemDescRequest) (*api.ListProblemDescReply, error)
	GetProblem(c controller.MContext, req *api.GetProblemRequest) (*api.GetProblemReply, error)
	PutProblem(c controller.MContext, req *api.PutProblemRequest) (*api.PutProblemReply, error)
	DeleteProblem(c controller.MContext, req *api.DeleteProblemRequest) (*api.DeleteProblemReply, error)
}
