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
	"github.com/Myriad-Dreamin/boj-v6/app/user_problem"
	"github.com/Myriad-Dreamin/boj-v6/deployment/database"
	"github.com/Myriad-Dreamin/boj-v6/deployment/oss"
	"github.com/Myriad-Dreamin/boj-v6/external"
	"github.com/Myriad-Dreamin/minimum-lib/rbac"
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
	"log"
	"path"
	"reflect"
)

var mdt = reflect.TypeOf(new(migrateDB)).Elem()

type migrateDB interface {
	Migrate() error
}

func RegisterDatabase(dbName string, f interface{}) InitializeAction {
	return func(srv *Server) error {
		return reflectCallInitDB(srv, dbName, ModuleInjectFunc(f))
	}
}

func reflectCallInitDB(srv *Server, dbName string, rf reflect.Value) error {

	res := rf.Call([]reflect.Value{reflect.ValueOf(srv.Module)})
	if len(res) != 2 {
		log.Fatalf("%v.Call return not length 2, got %v", rf.Type(), len(res))
	}
	db, err := res[0], res[1].Interface()
	if err != nil {
		srv.Logger.Debug(fmt.Sprintf("init %T DB error", db), "error", err)
		return err.(error)
	}

	if db.Type().Implements(mdt) {
		mg := db.Interface().(migrateDB)
		if err := mg.Migrate(); err != nil {
			srv.Logger.Debug(fmt.Sprintf("migrate %T DB error", mg), "error", err)
			return err
		}
	}

	err = srv.Module.ProvideWithCheck(
		path.Join("minimum", dbName), db.Interface())
	if err != nil {
		srv.Logger.Debug("provide database error", "name", dbName)
		return err.(error)
	}
	return nil
}

func InitDatabaseModule(mock bool) InitializeAction {
	return func(srv *Server) error {
		var m = database.NewModule()
		srv.DatabaseModule = &m

		if mock {
			if !m.InstallMock(srv.Module) {
				return fmt.Errorf("mock database initialize failed")
			}
		} else {
			srv.Cfg.DatabaseConfig.Debug(srv.Logger)
			if !m.Install(srv.Module) {
				return fmt.Errorf("initialize database with configuration failed")
			}
		}
		return nil
	}
}

func InitRedisModule(mock bool) InitializeAction {
	return func(srv *Server) error {
		//cfg:=
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
		return nil
	}
}

func InitRBACDatabase() InitializeAction {
	return func(srv *Server) error {
		err := rbac.InitGorm(srv.DatabaseModule.GormDB)
		if err != nil {
			srv.Logger.Debug("rbac to database error", "error", err)
			return err
		}
		err = srv.Module.ProvideImpl(new(*external.Enforcer), rbac.GetEnforcer())
		if err != nil {
			srv.Logger.Debug("provide enforcer error", "error", err)
			return err
		}
		return nil
	}
}

func InitOSSLevelDB(mock bool) InitializeAction {
	return func(srv *Server) error {
		_ = mock
		engine, err := oss.NewMemLevelDB(nil)
		if err != nil {
			srv.Logger.Debug("init mem mock oss", "error", err)
			return err
		}
		sugar.HandlerError0(srv.Module.ProvideImpl(new(*external.OSSEngine), engine))
		return nil
	}
}

func PrepareDatabase(mock bool) InitializeAction {
	return func(srv *Server) error {
		return srv.applyInitializeAction([]InitializeAction{
			InitDatabaseModule(mock),
			InitRBACDatabase(),
			InitOSSLevelDB(mock),
			RegisterDatabase("AnnouncementDB", announcement.NewDB),
			RegisterDatabase("UserDB", user.NewDB),
			RegisterDatabase("CommentDB", comment.NewDB),
			RegisterDatabase("SubmissionDB", submission.NewDB),
			RegisterDatabase("ProblemDB", problem.NewDB),
			RegisterDatabase("ProblemDescDB", problem_desc.NewDB),
			RegisterDatabase("ContestDB", contest.NewDB),
			RegisterDatabase("GroupDB", group.NewDB),
			RegisterDatabase("UserProblemRelationshipDB", user_problem.NewDB),
		})
	}
}
