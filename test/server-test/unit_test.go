package server_test

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/Myriad-Dreamin/boj-v6/lib/unittest"
	"github.com/Myriad-Dreamin/boj-v6/lib/unittest/unittest_types"
	"github.com/Myriad-Dreamin/boj-v6/pkg/server"
	"github.com/Myriad-Dreamin/boj-v6/test/tester"
	"github.com/Myriad-Dreamin/minimum-lib/mock"
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
	"hash/crc32"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func runUnitTest(ctx *tester.Context, ts []*unittest.TestCase) {
	runUnitTestCB(ctx, func() {}, ts)
}

type dataRecord struct {
	Comment        string      `json:"comment"`
	RequestBody    []byte      `json:"request_body"`
	RequestHeader  http.Header `json:"request_header"`
	ResponseCode   int         `json:"response_code"`
	ResponseBody   []byte      `json:"response_body"`
	ResponseHeader http.Header `json:"response_header"`
}

type dataResult struct {
	Handler string       `json:"handler"`
	Method  string       `json:"method"`
	Path    string       `json:"path"`
	Records []dataRecord `json:"data_records"`
}

func runUnitTestFileIsolatedWithSaveTestCases(t *testing.T, fileName string, options ...interface{}) {
	dataContext := runUnitTestFileIsolated(t, fileName, options...)
	if dataContext == nil {
		return
	}

	results := dataContext.DumpResults()

	dataResults := make([]dataResult, len(results))
	for i, result := range results {
		dataResults[i].Handler = result.GetHandler()
		dataResults[i].Method = result.GetMethod()
		dataResults[i].Path = result.GetPath()

		records := result.GetRecords()
		dataResults[i].Records = make([]dataRecord, len(records))
		for j, record := range records {
			dataResults[i].Records[j].Comment = record.GetComment()
			dataResults[i].Records[j].ResponseCode = record.GetResponseCode()
			dataResults[i].Records[j].RequestBody = record.GetRequestBody()
			dataResults[i].Records[j].RequestHeader = record.GetRequestHeader()
			dataResults[i].Records[j].ResponseBody = record.GetResponseBody()
			dataResults[i].Records[j].ResponseHeader = record.GetResponseHeader()
		}
	}

	sugar.WithWriteFile(func(file *os.File) {
		encoder := json.NewEncoder(file)
		sugar.HandlerError0(encoder.Encode(dataResults))
	}, fmt.Sprintf("../../docs/test_cases/%s.json",
		strings.TrimSuffix(strings.TrimSuffix(filepath.Base(fileName), ".yaml"), ".yml")))
}

func runUnitTestFileIsolated(t *testing.T, fileName string, options ...interface{}) *tester.Context {
	var (
		optionCallback func()
		optionInitFunc func(ctx *tester.Context)
	)

	for _, o := range options {
		switch o := o.(type) {
		case func():
			optionCallback = o
		case func(ctx *tester.Context):
			optionInitFunc = o
		}
	}

	instance := srv.Context(t).ResetServerInstance()
	g := unittest.Load(fileName, false, unittest.V1Opt)
	instance.Main(func(ctx *tester.Context) {
		if optionInitFunc != nil {
			optionInitFunc(ctx)
		}

		if optionCallback != nil {
			runUnitTestCB(ctx, optionCallback, g.TestCases)
		} else {
			runUnitTest(ctx, g.TestCases)
		}
	})
	return instance
}

func TestCRUDUnit(t *testing.T)    { runUnitTestFileIsolatedWithSaveTestCases(t, "test.crud.yaml") }
func TestProblemUnit(t *testing.T) { runUnitTestFileIsolatedWithSaveTestCases(t, "problem_test.yaml") }
func TestUnit(t *testing.T)        { runUnitTestFileIsolatedWithSaveTestCases(t, "test.yaml") }

func mapConvertString(f func(interface{}) string, x []interface{}) (s []string) {
	s = make([]string, len(x))
	for i := range x {
		s[i] = f(x[i])
	}
	return
}

func runUnitTestCB(ctx *tester.Context, cb func(), ts []*unittest.TestCase) {
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
			ctx.Error(err)
		}
		characteristic := s.Sum32()
		ctx.T.Run(tt.Name+"."+tt.Path, func(t *testing.T) {
			characteristic := characteristic
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

			b, err := json.Marshal(tt)
			sugar.HandlerError0(err)
			header["UnitTestMeta"] = base64.StdEncoding.EncodeToString(b)

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
