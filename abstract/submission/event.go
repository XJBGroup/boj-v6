package submission

import (
	"context"
	"io"
)

type PostEvent struct {
	S    Submission
	Code io.Reader
}

type PostSubmissionEventHandler interface {
	HandlePostSubmission(ctx context.Context, e PostEvent)
}

type Subscriber interface {
	AddPostSubmissionHandler(handler PostSubmissionEventHandler)
	RemovePostSubmissionHandler(handler PostSubmissionEventHandler)
}

type Dispatcher interface {
	PostSubmissionEventHandler
}
