package model

import (
	problem "github.com/Myriad-Dreamin/boj-v6/abstract/problem"
	"github.com/Myriad-Dreamin/boj-v6/abstract/submission"
	"github.com/Myriad-Dreamin/boj-v6/api"
	"github.com/Myriad-Dreamin/boj-v6/app/snippet"
	"github.com/Myriad-Dreamin/boj-v6/cmd/generate/stub"
	"github.com/Myriad-Dreamin/boj-v6/external"
	"github.com/Myriad-Dreamin/minimum-lib/controller"
)

var _ = snippet.AuthenticatePassword

type Sc struct {
	stub.StubVariables
	Binder stub.Stub

	key string

	problemDB problem.DB
	logStash  external.Logger
}

func (svc Sc) PostSubmission(c controller.MContext) {

	id := svc.Binder.GetID()
	var request = new(api.PostSubmissionRequest)
	svc.Binder.Bind(request)

	var response *api.GetSubmissionReply
	var s = new(submission.Submission)

	svc.Binder.Context(s).Serve(id, request, response).Catch(func() {
		svc.logStash.Info("Serve request failed", "context", c, "data", request)
	})

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
