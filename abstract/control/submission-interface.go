package control

import (
	"github.com/Myriad-Dreamin/minimum-lib/controller"
)

type SubmissionService interface {
	SubmissionServiceSignatureXXX() interface{}
	ListSubmission(c controller.MContext)
	CountSubmission(c controller.MContext)
	PostSubmission(c controller.MContext)
	GetContent(c controller.MContext)
	GetSubmission(c controller.MContext)
	DeleteSubmission(c controller.MContext)
}
