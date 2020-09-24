package inner_control

import (
	"github.com/Myriad-Dreamin/boj-v6/api"
)

type InnerSubmissionService interface {
	SubmissionServiceSignatureXXX() interface{}
	ListSubmission(req *api.ListSubmissionRequest) (*api.ListSubmissionReply, error)
	CountSubmission(req *api.CountSubmissionRequest) (*api.CountSubmissionReply, error)
	PostSubmission(req *api.PostSubmissionRequest) (*api.PostSubmissionReply, error)
	GetContent(req *api.GetContentRequest) (*api.GetContentReply, error)
	GetSubmission(req *api.GetSubmissionRequest) (*api.GetSubmissionReply, error)
	DeleteSubmission(req *api.DeleteSubmissionRequest) (*api.DeleteSubmissionReply, error)
}
