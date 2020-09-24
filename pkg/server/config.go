package server

import (
	"github.com/Myriad-Dreamin/boj-v6/config"
)

func LoadConfig(cfgPath string) InitializeAction {
	if len(cfgPath) == 0 {
		return UseDefaultConfig()
	}

	return func(srv *Server) error {
		srv.Cfg = config.Default()
		err := config.Load(srv.Cfg, cfgPath)
		if err != nil {
			srv.Logger.Debug("parse config error", "error", err)
			return err
		}
		srv.Module.Provide(config.ModulePath.Global.Configuration, srv.Cfg)
		return nil
	}
}

func UseDefaultConfig() InitializeAction {
	return func(srv *Server) error {
		srv.Cfg = config.Default()
		srv.Module.Provide(config.ModulePath.Global.Configuration, srv.Cfg)
		return nil
	}
}

func (srv *Server) FetchConfig(cfg interface{}, cfgPath string) bool {
	err := config.LoadStatic(cfg, cfgPath)
	if err != nil {
		srv.Logger.Debug("parse config error", "error", err)
		return false
	}
	return true
}
