package tests

import (
	"context"
	"fmt"
	"github.com/Myriad-Dreamin/boj-v6/abstract/submission"
	"github.com/Myriad-Dreamin/boj-v6/lib/unittest"
	"github.com/Myriad-Dreamin/boj-v6/lib/unittest/unittest_types"
	"github.com/Myriad-Dreamin/boj-v6/pkg/server"
	"github.com/Myriad-Dreamin/boj-v6/types"
	"github.com/Myriad-Dreamin/minimum-lib/mock"
	"github.com/stretchr/testify/assert"
	"hash/crc32"
	"strings"
	"testing"
)

func TestCRUDUnit(t *testing.T) {
	g := unittest.Load("test.crud.yaml", false, unittest.V1Opt)
	runUnitTest(t, g.TestCases)
}

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
	subscriber := srv.Module.RequireImpl(new(submission.Subscriber)).(submission.Subscriber)

	ch := make(chan *submission.Submission, 5)
	db := srv.Module.RequireImpl(new(submission.DB)).(submission.DB)
	subscriber.AddPostSubmissionHandler(&handler{
		ch: ch, t: t})

	g := unittest.Load("submission_test.yaml", false, unittest.V1Opt)
	runUnitTestCB(t, func() {
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

	}, g.TestCases)
}

func TestUnit(t *testing.T) {
	g := unittest.Load("test.yaml", false, unittest.V1Opt)
	runUnitTest(t, g.TestCases)
}

func mapConvertString(f func(interface{}) string, x []interface{}) (s []string) {
	s = make([]string, len(x))
	for i := range x {
		s[i] = f(x[i])
	}
	return
}

func runUnitTest(t *testing.T, ts []*unittest.TestCase) {
	runUnitTestCB(t, func() {}, ts)
}

func runUnitTestCB(t *testing.T, cb func(), ts []*unittest.TestCase) {
	for _, tt := range ts {
		if tt.Abstract {
			continue
		}
		if tt.Meta[unittest.MetaUrl] == nil {
			continue
		}
		s := crc32.NewIEEE()
		_, err := s.Write([]byte(tt.Name + "." + tt.Path))
		if err != nil {
			t.Error(err)
		}
		characteristic := s.Sum32()
		t.Run(tt.Name+"."+tt.Path, func(t *testing.T) {
			characteristic := characteristic
			ctx := srv.Context(t)
			method, ok := tt.Meta[unittest.MetaHTTPMethod]
			if !ok {
				method, ok = tt.Meta[unittest.MetaMethod]
			}
			if !ok {
				return
			}
			xheader, ok := tt.Meta[unittest.MetaHTTPHeader]
			if !ok {
				xheader, ok = tt.Meta[unittest.MetaHeader]
			}
			var header = map[string]string{}
			if ok {
				header = xheader.(map[string]string)
			}

			var mockResponse mock.ResponseI
			switch method {
			case "GET":
				mockResponse = ctx.Get(tt.Meta[unittest.MetaUrl].(string), tt.Meta[unittest.MetaData], server.Header(header))
			case "POST", "PUT", "DELETE":
				mockResponse = ctx.Method(method.(string),
					tt.Meta[unittest.MetaUrl].(string), tt.Meta[unittest.MetaData], server.Header(header))
			default:
				panic(fmt.Sprintf("%v", tt))
			}

			if mockResponse == nil {
				panic("nil response")
			}
			req := &unittest_types.Response{Body: mockResponse.Body().Bytes()}
			st := &unittest_types.State{Res: req}
			for _, assertion := range tt.Script {
				ok, err := assertion.F(st, assertion.VArgs...)
				if err != nil {
					t.Errorf("%v(%v): url>> %v@%v, err>> %v, test id>> %v", assertion.FN,
						strings.Join(mapConvertString(func(i interface{}) string {
							return fmt.Sprintf("%v", i)
						}, assertion.VArgs), ", "),
						tt.Meta[unittest.MetaUrl].(string), method, err, characteristic)
				} else if !ok {
					t.Errorf("%v(%v) == false: url %v@%v, test id>> %v", assertion.FN,
						strings.Join(mapConvertString(func(i interface{}) string {
							return fmt.Sprintf("%v", i)
						}, assertion.VArgs), ", "),
						tt.Meta[unittest.MetaUrl].(string), method, characteristic)
				}
			}
		})
		cb()
	}
}
