package server

import (
	"context"
	"github.com/Myriad-Dreamin/boj-v6/abstract/submission"
)

// todo: security consideration

type SubmissionEventEmitter struct {
	postHandlers map[submission.PostSubmissionEventHandler]struct{}
}

func newSubmissionEventEmitter() *SubmissionEventEmitter {
	return &SubmissionEventEmitter{postHandlers: make(map[submission.PostSubmissionEventHandler]struct{})}
}

func (s2 SubmissionEventEmitter) HandlePostSubmission(ctx context.Context, e submission.PostEvent) {
	for h := range s2.postHandlers {
		go h.HandlePostSubmission(ctx, e)
	}
}

func (s2 SubmissionEventEmitter) AddPostSubmissionHandler(handler submission.PostSubmissionEventHandler) {
	s2.postHandlers[handler] = struct{}{}
}

func (s2 SubmissionEventEmitter) RemovePostSubmissionHandler(handler submission.PostSubmissionEventHandler) {
	delete(s2.postHandlers, handler)
}

func AddEvent(srv *Server) (err error) {
	var doRegister = func(protocolPtr, impl interface{}) {
		if err != nil {
			return
		}
		err = srv.Module.ProvideImpl(protocolPtr, impl)
	}

	doRegister(new(*SubmissionEventEmitter), newSubmissionEventEmitter())

	return
}
