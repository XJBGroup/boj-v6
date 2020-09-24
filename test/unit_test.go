package tests

import (
	"fmt"
	"github.com/Myriad-Dreamin/boj-v6/lib/unittest"
	"github.com/Myriad-Dreamin/boj-v6/lib/unittest/unittest_types"
	"github.com/Myriad-Dreamin/boj-v6/pkg/server"
	"github.com/Myriad-Dreamin/minimum-lib/mock"
	"hash/crc32"
	"strings"
	"testing"
)

func TestCRUDUnit(t *testing.T) {
	g := unittest.Load("test.crud.yaml", false, unittest.V1Opt)
	runUnitTest(t, g.TestCases)
}

func TestSubmissionUnit(t *testing.T) {
	g := unittest.Load("submission_test.yaml", false, unittest.V1Opt)
	runUnitTest(t, g.TestCases)
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
	}
}
