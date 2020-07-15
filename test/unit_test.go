package tests

import (
	"fmt"
	"github.com/Myriad-Dreamin/boj-v6/lib/unittest"
	"github.com/Myriad-Dreamin/boj-v6/pkg/server"
	"github.com/Myriad-Dreamin/minimum-lib/mock"
	"testing"
)

func TestUnit(t *testing.T) {
	g := unittest.Load("test.yaml", false)
	fmt.Println(g)
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
		if tt.Meta[unittest.MetaData] == nil {
			continue
		}
		t.Run(tt.Name+"."+tt.Path, func(t *testing.T) {
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
			case "POST":
				mockResponse = ctx.Post(tt.Meta[unittest.MetaUri].(string), tt.Meta[unittest.MetaData], server.Header(header))
			default:
				panic(fmt.Sprintf("%v", tt))
			}

			if mockResponse == nil {
				panic("nil response")
			}
			for _, assertion := range tt.Assertions {
				ok, err := assertion.F(&unittest.Request{Body: mockResponse.Body().Bytes()}, assertion.VArgs...)
				if err != nil {
					panic(err)
				}
				if !ok {
					panic(assertion.FN)
				}
			}
		})
	}
}
