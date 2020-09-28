package model

import (
	"github.com/Myriad-Dreamin/boj-v6/abstract/problem"
	"github.com/Myriad-Dreamin/boj-v6/abstract/submission"
	"github.com/Myriad-Dreamin/boj-v6/api"
	"github.com/Myriad-Dreamin/boj-v6/cmd/generate/stub"
	"github.com/Myriad-Dreamin/minimum-lib/controller"
)

type Sc struct {
	stub.StubVariables
	Binder stub.Stub

	key string

	problemDB problem.DB
}

func (svc Sc) PostSubmission(c controller.MContext) {

	id := svc.Binder.GetID()
	var request = new(api.PostSubmissionRequest)
	svc.Binder.Bind(request)

	var response *api.GetSubmissionReply
	var s *submission.Submission
	svc.Binder.Context(s).Serve(id, request, response)

	svc.Binder.EmitSelf(s, request.Code)
}

func (svc Sc) SubmissionControllerSignatureXXX() interface{} {
	panic("implement me")
}

func (svc Sc) ListSubmission(c controller.MContext) {
	panic("implement me")
}

func (svc Sc) CountSubmission(c controller.MContext) {
	panic("implement me")
}

func (svc Sc) GetSubmissionContent(c controller.MContext) {
	panic("implement me")
}

func (svc Sc) GetSubmission(c controller.MContext) {

	id := svc.Binder.GetID()
	var request *api.GetSubmissionRequest
	svc.Binder.Bind(request)

	var response *api.GetSubmissionReply
	svc.Binder.Serve(id, request, response)
}

func (svc Sc) DeleteSubmission(c controller.MContext) {
	panic("implement me")
}
