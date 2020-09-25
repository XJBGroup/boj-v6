package tests

import (
	unwrap_func_sqlite "github.com/Myriad-Dreamin/boj-v6/lib/unwrap_func/sqlite"
	"github.com/Myriad-Dreamin/boj-v6/pkg/server"
	"github.com/Myriad-Dreamin/boj-v6/test/tester"
	"github.com/Myriad-Dreamin/minimum-lib/crypto"
	"github.com/Myriad-Dreamin/minimum-lib/module"
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
	"golang.org/x/crypto/bcrypt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	crypto.SetCost(bcrypt.MinCost)

	serverModule := make(module.Module)
	unwrap_func_sqlite.Register(serverModule)

	sugar.WithFile(func(logFile *os.File) {
		var options = []server.Option{
			server.OptionRouterLoggerWriter{
				Writer: logFile,
			},
			server.CopyOptionModule{
				Module: serverModule,
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
