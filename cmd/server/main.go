package main

import (
	"github.com/Myriad-Dreamin/boj-v6/pkg/server"
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
	_ "net/http/pprof"
	"os"
)

func main() {
	//isDebug := true

	//srv.Inject(sphinxcore.New())

	//if isDebug {
	//	srv.ServeWithPProf(":23336")
	//} else {
	//}
	sugar.WithFile(func(logFile *os.File) {
		var options = []server.Option{
			server.OptionRouterLoggerWriter{
				Writer: logFile,
			},
		}

		srv := server.New("config", options...)
		if srv == nil {
			return
		}

		srv.Serve(":23336")
	}, "test.log")

}
