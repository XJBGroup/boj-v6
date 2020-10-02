package integrate_test

import (
	"github.com/Myriad-Dreamin/boj-v6/lib/unwrap_func/sqlite"
	"github.com/Myriad-Dreamin/boj-v6/pkg/server"
	"github.com/Myriad-Dreamin/minimum-lib/crypto"
	"github.com/Myriad-Dreamin/minimum-lib/module"
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
	_ "net/http/pprof"
	"os"
	"testing"
)

func TestIntegrate(t *testing.T) {
	//isDebug := true

	//srv.Inject(sphinxcore.New())

	//if isDebug {
	//	srv.ServeWithPProf(":23336")
	//} else {
	//}
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

		sugar.HandlerError0(server.Mock(options...).HttpEngine.Run(":23336"))
	}, "./test/integrate-test/test.log")
}
