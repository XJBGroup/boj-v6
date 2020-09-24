package server

import (
	"fmt"
	"github.com/Myriad-Dreamin/boj-v6/config"
	"github.com/Myriad-Dreamin/minimum-lib/logger"
	"go.uber.org/zap/zapcore"
)

func  InstantiateLogger() InitializeAction {
	return func  (srv *Server) error {
		var err error
		srv.Logger, err = logger.NewZapLogger(logger.NewZapDevelopmentSugarOption(), zapcore.DebugLevel)
		if err != nil {
			fmt.Println("Initialize Server Logger Error", err)
			return err
		}
		srv.Module.Provide(config.ModulePath.Global.Logger, srv.Logger)
		return nil
	}
}

func (srv *Server) println(a ...interface{}) {
	_, err := fmt.Fprintln(srv.LoggerWriter, a...)
	if err != nil {
		fmt.Println(err)
	}
}

func (srv *Server) printf(format string, a ...interface{}) {
	_, err := fmt.Fprintf(srv.LoggerWriter, format, a...)
	if err != nil {
		fmt.Println(err)
	}
}
