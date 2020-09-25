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
	ChangeProblemDescriptionRef(c controller.MContext, req *api.ChangeProblemDescriptionRefRequest) (*api.ChangeProblemDescriptionRefReply, error)
	PostProblemDesc(c controller.MContext, req *api.PostProblemDescRequest) (*api.PostProblemDescReply, error)
	GetProblemDesc(c controller.MContext, req *api.GetProblemDescRequest) (*api.GetProblemDescReply, error)
	PutProblemDesc(c controller.MContext, req *api.PutProblemDescRequest) (*api.PutProblemDescReply, error)
	DeleteProblemDesc(c controller.MContext, req *api.DeleteProblemDescRequest) (*api.DeleteProblemDescReply, error)
	ProblemFSReadConfig(c controller.MContext, req *api.ProblemFSReadConfigRequest) (*api.ProblemFSReadConfigReply, error)
	ProblemFSWriteConfig(c controller.MContext, req *api.ProblemFSWriteConfigRequest) (*api.ProblemFSWriteConfigReply, error)
	ProblemFSPutConfig(c controller.MContext, req *api.ProblemFSPutConfigRequest) (*api.ProblemFSPutConfigReply, error)
	ProblemFSRead(c controller.MContext, req *api.ProblemFSReadRequest) (*api.ProblemFSReadReply, error)
	ProblemFSStat(c controller.MContext, req *api.ProblemFSStatRequest) (*api.ProblemFSStatReply, error)
	ProblemFSWrite(c controller.MContext, req *api.ProblemFSWriteRequest) (*api.ProblemFSWriteReply, error)
	ProblemFSRemove(c controller.MContext, req *api.ProblemFSRemoveRequest) (*api.ProblemFSRemoveReply, error)
	ProblemFSZipRead(c controller.MContext, req *api.ProblemFSZipReadRequest) (*api.ProblemFSZipReadReply, error)
	ProblemFSLS(c controller.MContext, req *api.ProblemFSLSRequest) (*api.ProblemFSLSReply, error)
	ProblemFSWrites(c controller.MContext, req *api.ProblemFSWritesRequest) (*api.ProblemFSWritesReply, error)
	ProblemFSMkdir(c controller.MContext, req *api.ProblemFSMkdirRequest) (*api.ProblemFSMkdirReply, error)
	ProblemFSRemoveAll(c controller.MContext, req *api.ProblemFSRemoveAllRequest) (*api.ProblemFSRemoveAllReply, error)
	ListProblemDesc(c controller.MContext, req *api.ListProblemDescRequest) (*api.ListProblemDescReply, error)
	CountProblemDesc(c controller.MContext, req *api.CountProblemDescRequest) (*api.CountProblemDescReply, error)
	GetProblem(c controller.MContext, req *api.GetProblemRequest) (*api.GetProblemReply, error)
	PutProblem(c controller.MContext, req *api.PutProblemRequest) (*api.PutProblemReply, error)
	DeleteProblem(c controller.MContext, req *api.DeleteProblemRequest) (*api.DeleteProblemReply, error)
}
