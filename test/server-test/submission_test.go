package server_test

import (
	"context"
	"github.com/Myriad-Dreamin/boj-v6/abstract/submission"
	"github.com/Myriad-Dreamin/boj-v6/test/tester"
	"github.com/Myriad-Dreamin/boj-v6/types"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type handler struct {
	db submission.DB
	ch chan *submission.Submission
	t  *testing.T
}

func (h *handler) HandlePostSubmission(ctx context.Context, e submission.PostEvent) {
	h.ch <- &e.S
	return
}

func TestSubmissionUnit(t *testing.T) {
	var (
		db submission.DB
		ch chan *submission.Submission
	)

	runUnitTestFileIsolatedWithSaveTestCases(t, "submission_test.yaml",
		func(ctx *tester.Context) {

			subscriber := ctx.Module.RequireImpl(new(submission.Subscriber)).(submission.Subscriber)

			ch = make(chan *submission.Submission, 5)
			db = ctx.Module.RequireImpl(new(submission.DB)).(submission.DB)
			subscriber.AddPostSubmissionHandler(&handler{
				ch: ch, t: t})
		},
		func() {
			time.Sleep(time.Millisecond)
			select {
			case s := <-ch:
				if s.ID == 1 {
					s.Status = types.StatusAccepted

					aff, err := db.UpdateFields(s, []string{"status"})
					assert.Equal(t, int64(1), aff)
					assert.NoError(t, err, aff)
				}
			default:

			}
		})
}
