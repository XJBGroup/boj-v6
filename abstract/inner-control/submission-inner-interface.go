package inner_control

import (
	"github.com/Myriad-Dreamin/boj-v6/api"
	"github.com/Myriad-Dreamin/minimum-lib/controller"
)

type InnerSubmissionController interface {
	SubmissionControllerSignatureXXX() interface{}
	ListSubmission(c controller.MContext, req *api.ListSubmissionRequest) (*api.ListSubmissionReply, error)
	CountSubmission(c controller.MContext, req *api.CountSubmissionRequest) (*api.CountSubmissionReply, error)
	PostSubmission(c controller.MContext, req *api.PostSubmissionRequest) (*api.PostSubmissionReply, error)
	GetSubmissionContent(c controller.MContext, req *api.GetSubmissionContentRequest) (*api.GetSubmissionContentReply, error)
	GetSubmission(c controller.MContext, req *api.GetSubmissionRequest) (*api.GetSubmissionReply, error)
	DeleteSubmission(c controller.MContext, req *api.DeleteSubmissionRequest) (*api.DeleteSubmissionReply, error)
}
