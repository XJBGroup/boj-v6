package tests

import (
	"github.com/Myriad-Dreamin/boj-v6/lib/errorc"
	"github.com/Myriad-Dreamin/boj-v6/pkg/server"
	"github.com/Myriad-Dreamin/boj-v6/test/tester"
	"github.com/Myriad-Dreamin/boj-v6/types"
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
	"github.com/mattn/go-sqlite3"
	"os"
	"strconv"
	"testing"
)

func TestMain(m *testing.M) {
	errorc.RegisterCheckInsertError(func(err error) (code errorc.Code, s string) {
		if sqlError, ok := err.(sqlite3.Error); ok {
			switch sqlError.ExtendedCode {
			case 1062:
				return types.CodeDuplicatePrimaryKey, ""
			case 1366:
				return types.CodeDatabaseIncorrectStringValue, ""
			case 2067:
				return types.CodeUniqueConstraintFailed, ""
			default:
				return types.CodeInsertError, strconv.Itoa(int(sqlError.ExtendedCode))
			}
		}
		return types.CodeOK, ""
	})

	sugar.WithFile(func(logFile *os.File) {
		var options = []server.Option{
			server.OptionRouterLoggerWriter{
				Writer: logFile,
			},
		}
		srv = tester.StartTester(options)
		srv.PrintRequest(true)
		srv.CollectResults(true)
		srv.MainM(m)
		srv.DropMock()
		//err := doc_gen.FromGinResults(&doc_gen.GinInfo{
		//	Result:             srv.DumpResults(),
		//	Host:               "127.0.0.1",
		//	ControllerProvider: srv.ServiceProvider,
		//})
		//if err != nil {
		//	panic(err)
		//}
	}, "test.log")
}
