package tests

import (
	"fmt"
	"github.com/Myriad-Dreamin/boj-v6/lib/unittest"
	"github.com/Myriad-Dreamin/boj-v6/pkg/server"
	"github.com/Myriad-Dreamin/minimum-lib/mock"
	"hash/crc32"
	"strings"
	"testing"
)

func TestUnit(t *testing.T) {
	g := unittest.Load("test.yaml", false)
	runUnitTest(t, g.TestCases)
}

func runUnitTest(t *testing.T, ts []*unittest.TestCase) {
	for _, tt := range ts {
		if tt.Abstract {
			continue
		}
		if tt.Meta[unittest.MetaUri] == nil {
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
				mockResponse = ctx.Get(tt.Meta[unittest.MetaUri].(string), tt.Meta[unittest.MetaData], server.Header(header))
			case "POST", "PUT", "DELETE":
				mockResponse = ctx.Method(method.(string),
					tt.Meta[unittest.MetaUri].(string), tt.Meta[unittest.MetaData], server.Header(header))
			default:
				panic(fmt.Sprintf("%v", tt))
			}

			if mockResponse == nil {
				panic("nil response")
			}
			for _, assertion := range tt.Assertions {
				ok, err := assertion.F(&unittest.Request{Body: mockResponse.Body().Bytes()}, assertion.VArgs...)
				if err != nil {
					t.Errorf("%v(%v): url>> %v@%v, err>> %v, test id>> %v", assertion.FN, strings.Join(assertion.VArgs, ", "), tt.Meta[unittest.MetaUri].(string), method, err, characteristic)
				} else if !ok {
					t.Errorf("%v(%v) == false: url %v@%v, test id>> %v", assertion.FN, strings.Join(assertion.VArgs, ", "), tt.Meta[unittest.MetaUri].(string), method, characteristic)
				}
			}
		})
	}
}
