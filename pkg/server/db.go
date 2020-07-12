package server

import (
	"fmt"
	"github.com/Myriad-Dreamin/boj-v6/app/announcement"
	"github.com/Myriad-Dreamin/boj-v6/app/comment"
	"github.com/Myriad-Dreamin/boj-v6/app/submission"
	"github.com/Myriad-Dreamin/boj-v6/app/user"
	"github.com/Myriad-Dreamin/boj-v6/deployment/database"
	"github.com/Myriad-Dreamin/functional-go"
	"github.com/Myriad-Dreamin/minimum-lib/rbac"
)

type dbResult struct {
	dbName string
	functional.DecayResult
}

func (srv *Server) registerDatabaseService() bool {

	for _, dbResult := range []dbResult{
		{"announcementDB", functional.Decay(announcement.NewDB(srv.Module))},
		{"userDB", functional.Decay(user.NewDB(srv.Module))},
		{"commentDB", functional.Decay(comment.NewDB(srv.Module))},
		{"submissionDB", functional.Decay(submission.NewDB(srv.Module))},
	} {
		if dbResult.Err != nil {
			srv.Logger.Debug(fmt.Sprintf("init %T DB error", dbResult.First), "error", dbResult.Err)
			return false
		}

		if migratingDB, ok := dbResult.First.(interface {
			Migrate() error
		}); ok {
			if err := migratingDB.Migrate(); err != nil {
				srv.Logger.Debug(fmt.Sprintf("migrate %T DB error", migratingDB), "error", dbResult.Err)
				return false
			}
		}

		srv.ModelProvider.Register(dbResult.dbName, dbResult.First)
	}
	return true
}

func (srv *Server) PrepareDatabase() bool {
	srv.Cfg.DatabaseConfig.Debug(srv.Logger)

	var m = database.NewModule()
	srv.databaseModule = &m

	if !m.Install(srv.Module) {
		return false
	}

	//srv.RedisPool, err = model.OpenRedis(cfg)
	//if err != nil {
	//	srv.Logger.Debug("create redis pool error", "error", err)
	//	return false
	//}
	//
	//srv.Logger.Info("connected to redis",
	//	"connection-type", cfg.RedisConfig.ConnectionType,
	//	"host", cfg.RedisConfig.Host,
	//	"connection-timeout", cfg.RedisConfig.ConnectionTimeout,
	//	"database", cfg.RedisConfig.Database,
	//	"read-timeout", cfg.RedisConfig.ReadTimeout,
	//	"write-timeout", cfg.RedisConfig.WriteTimeout,
	//	"idle-timeout", cfg.RedisConfig.IdleTimeout,
	//	"wait", cfg.RedisConfig.Wait,
	//	"max-active", cfg.RedisConfig.MaxActive,
	//	"max-idle", cfg.RedisConfig.MaxIdle,
	//)
	//err = model.RegisterRedis(srv.RedisPool, srv.Logger)
	//if err != nil {
	//	srv.Logger.Debug("register redis error", "error", err)
	//	return false
	//}
	err := rbac.InitGorm(m.GormDB)
	if err != nil {
		srv.Logger.Debug("rbac to database error", "error", err)
		return false
	}
	srv.ModelProvider.Register("enforcer", rbac.GetEnforcer())

	return srv.registerDatabaseService()
}

func (srv *Server) MockDatabase() bool {
	srv.Cfg.DatabaseConfig.Debug(srv.Logger)

	var m = database.NewModule()
	srv.databaseModule = &m

	if !m.InstallMock(srv.Module) {
		return false
	}

	err := rbac.InitGorm(m.GormDB)
	if err != nil {
		srv.Logger.Debug("rbac to database error", "error", err)
		return false
	}
	srv.ModelProvider.Register("enforcer", rbac.GetEnforcer())

	return srv.registerDatabaseService()
}
