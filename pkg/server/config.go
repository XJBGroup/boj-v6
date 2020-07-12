package server

import (
	"github.com/Myriad-Dreamin/boj-v6/config"
	"github.com/Myriad-Dreamin/boj-v6/lib/errorc"
	"github.com/Myriad-Dreamin/boj-v6/types"
	"github.com/go-sql-driver/mysql"
	"strconv"
)

func (srv *Server) LoadConfig(cfgPath string) bool {
	srv.Cfg = config.Default()
	err := config.Load(srv.Cfg, cfgPath)
	if err != nil {
		srv.Logger.Debug("parse config error", "error", err)
		return false
	}
	srv.Module.Provide(config.ModulePath.Global.Configuration, srv.Cfg)
	return true
}
func (srv *Server) UseDefaultConfig() bool {
	srv.Cfg = config.Default()
	srv.Module.Provide(config.ModulePath.Global.Configuration, srv.Cfg)
	return true
}

func (srv *Server) FetchConfig(cfg interface{}, cfgPath string) bool {
	err := config.LoadStatic(cfg, cfgPath)
	if err != nil {
		srv.Logger.Debug("parse config error", "error", err)
		return false
	}
	return true
}

func init() {

	errorc.RegisterCheckInsertError(func(err error) (code errorc.Code, s string) {
		if mysqlError, ok := err.(*mysql.MySQLError); ok {
			switch mysqlError.Number {
			case 1062:
				return types.CodeDuplicatePrimaryKey, ""
			case 1366:
				return types.CodeDatabaseIncorrectStringValue, ""
			default:
				return types.CodeInsertError, strconv.Itoa(int(mysqlError.Number))
			}
		}
		return types.CodeOK, ""
	})
}
