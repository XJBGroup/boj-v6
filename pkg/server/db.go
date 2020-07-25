package server

import (
	"fmt"
	"github.com/Myriad-Dreamin/boj-v6/app/announcement"
	"github.com/Myriad-Dreamin/boj-v6/app/comment"
	"github.com/Myriad-Dreamin/boj-v6/app/contest"
	"github.com/Myriad-Dreamin/boj-v6/app/group"
	"github.com/Myriad-Dreamin/boj-v6/app/problem"
	problem_desc "github.com/Myriad-Dreamin/boj-v6/app/problem-desc"
	"github.com/Myriad-Dreamin/boj-v6/app/submission"
	"github.com/Myriad-Dreamin/boj-v6/app/user"
	"github.com/Myriad-Dreamin/boj-v6/deployment/database"
	"github.com/Myriad-Dreamin/boj-v6/deployment/oss"
	"github.com/Myriad-Dreamin/boj-v6/external"
	"github.com/Myriad-Dreamin/functional-go"
	"github.com/Myriad-Dreamin/minimum-lib/rbac"
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
	"path"
)

type dbResult struct {
	dbName string
	proto  interface{}
	functional.DecayResult
}

func (srv *Server) registerDatabaseService() bool {

	for _, dbResult := range []dbResult{
		{"AnnouncementDB", new(*announcement.DBImpl), functional.Decay(announcement.NewDB(srv.Module))},
		{"UserDB", new(*user.DBImpl), functional.Decay(user.NewDB(srv.Module))},
		{"CommentDB", new(*comment.DBImpl), functional.Decay(comment.NewDB(srv.Module))},
		{"SubmissionDB", new(*submission.DBImpl), functional.Decay(submission.NewDB(srv.Module))},
		{"ProblemDB", new(*problem.DBImpl), functional.Decay(problem.NewDB(srv.Module))},
		{"ProblemDescDB", new(*problem_desc.DBImpl), functional.Decay(problem_desc.NewDB(srv.Module))},
		{"ContestDB", new(*contest.DBImpl), functional.Decay(contest.NewDB(srv.Module))},
		{"GroupDB", new(*group.DBImpl), functional.Decay(group.NewDB(srv.Module))},
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
		err := srv.Module.ProvideNamedImpl(
			path.Join("minimum", dbResult.dbName), dbResult.proto, dbResult.First)
		if err != nil {
			srv.Logger.Debug("provide database error", "name", dbResult.dbName)
			return false
		}
	}
	return true
}

func (srv *Server) PrepareDatabase() bool {
	srv.Cfg.DatabaseConfig.Debug(srv.Logger)

	var m = database.NewModule()
	srv.DatabaseModule = &m

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
	err = srv.Module.ProvideImpl(new(*external.Enforcer), rbac.GetEnforcer())
	if err != nil {
		srv.Logger.Debug("provide enforcer error", "error", err)
		return false
	}

	engine, err := oss.NewMemLevelDB(nil)
	if err != nil {
		srv.Logger.Debug("init mem mock oss", "error", err)
		return false
	}
	sugar.HandlerError0(srv.Module.ProvideImpl(new(*external.OSSEngine), engine))

	return srv.registerDatabaseService()
}

func (srv *Server) MockDatabase() bool {
	srv.Cfg.DatabaseConfig.Debug(srv.Logger)

	var m = database.NewModule()
	srv.DatabaseModule = &m

	if !m.InstallMock(srv.Module) {
		return false
	}

	err := rbac.InitGorm(m.GormDB)
	if err != nil {
		srv.Logger.Debug("rbac to database error", "error", err)
		return false
	}
	err = srv.Module.ProvideImpl(new(*external.Enforcer), rbac.GetEnforcer())
	if err != nil {
		srv.Logger.Debug("provide enforcer error", "error", err)
		return false
	}

	engine, err := oss.NewMemLevelDB(nil)
	if err != nil {
		srv.Logger.Debug("init mem mock oss", "error", err)
		return false
	}
	sugar.HandlerError0(srv.Module.ProvideImpl(new(*external.OSSEngine), engine))

	return srv.registerDatabaseService()
}
