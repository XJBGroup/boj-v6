package server

import (
	"context"
	"github.com/DeanThompson/ginpprof"
	"github.com/Myriad-Dreamin/boj-v6/api"
	"github.com/Myriad-Dreamin/boj-v6/config"
	"github.com/Myriad-Dreamin/boj-v6/deployment/database"
	"github.com/Myriad-Dreamin/boj-v6/external"
	"github.com/Myriad-Dreamin/boj-v6/lib/control"
	"github.com/Myriad-Dreamin/boj-v6/lib/deepcopy"
	"github.com/Myriad-Dreamin/boj-v6/lib/jwt"
	"github.com/Myriad-Dreamin/boj-v6/pkg/plugin"
	"github.com/Myriad-Dreamin/boj-v6/types"
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	"github.com/Myriad-Dreamin/minimum-lib/module"
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"io"
	"net/http"
	"os"
	"sync"
	"syscall"
)

type Server struct {
	Name         string
	Cfg          *config.ServerConfig
	Logger       external.Logger
	LoggerWriter io.Writer
	Module       module.Module

	Filesystem types.Filesystem
	HttpEngine *control.HttpEngine

	DatabaseModule *database.Module
	RedisPool      *redis.Pool
	Router         *api.RootRouter

	jwtMW        *jwt.Middleware
	routerAuthMW *controller.Middleware
	corsMW       gin.HandlerFunc

	plugins []plugin.Plugin
}

func NewServer() *Server {
	return &Server{}
}

func (srv *Server) Terminate() {
	//model.Close(srv.Module)
	syscall.Exit(0)
}

type Option interface {
	MinimumServerOption() bool
}

type OptionImpl struct{}

func (OptionImpl) MinimumServerOption() bool { return false }

type OptionRouterLoggerWriter struct {
	OptionImpl
	Writer io.Writer
}

type OptionModule struct {
	OptionImpl
	Module module.Module
}

type OptionName struct {
	OptionImpl
	Name string
}

type OptionFilesystem struct {
	OptionImpl
	Filesystem types.Filesystem
}

type CopyOptionModule OptionModule

func newServer(options []Option) *Server {
	srv := NewServer()

	for i := range options {
		switch option := options[i].(type) {
		case OptionRouterLoggerWriter:
			srv.LoggerWriter = option.Writer
		case *OptionRouterLoggerWriter:
			srv.LoggerWriter = option.Writer
		case OptionModule:
			srv.Module = option.Module
		case *OptionModule:
			srv.Module = option.Module
		case CopyOptionModule:
			srv.Module = deepcopy.DeepCopy(option.Module).(module.Module)
		case *CopyOptionModule:
			srv.Module = deepcopy.DeepCopy(option.Module).(module.Module)
		case OptionName:
			srv.Name = option.Name
		case *OptionName:
			srv.Name = option.Name
		case OptionFilesystem:
			srv.Filesystem = option.Filesystem
		case *OptionFilesystem:
			srv.Filesystem = option.Filesystem
		}
	}

	if srv.Module == nil {
		srv.Module = make(module.Module)
	}

	if srv.LoggerWriter == nil {
		srv.LoggerWriter = os.Stdout
	}

	if len(srv.Name) == 0 {
		srv.Name = "BOJ-v6"
	}

	//srv.RouterProvider = router.NewProvider(config.ModulePath.Provider.Router)

	//srv.Module.Provide(config.ModulePath.Provider.Router, srv.RouterProvider)
	return srv
}

func InitServer(cfgPath string, mock bool) InitializeAction {
	return func(srv *Server) error {
		return srv.applyInitializeAction([]InitializeAction{
			InstantiateLogger(),
			LoadConfig(cfgPath),
			PrepareFileSystem(mock),
			PrepareDatabase(mock),
			AddEvent,
		})
	}

}

func New(cfgPath string, options ...Option) (srv *Server) {
	srv = newServer(options)
	err := InitServer(cfgPath, false)(srv)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := recover(); err != nil {
			sugar.PrintStack()
			srv.Logger.Error("panic error", "error", err)
			srv.Terminate()
		} else if srv == nil {
			srv.Terminate()
		}
	}()

	if !(srv.PrepareMiddleware() &&
		srv.PrepareService() &&
		srv.BuildRouter(false)) {
		srv = nil
		return
	}

	//if err := srv.Module.Install(srv.RouterProvider); err != nil {
	//	srv.println("install router provider error", err)
	//}

	//
	//if !PreparePlugin(cfg) {
	//	srv = nil
	//return
	//}

	// Pressure()
	return
}

func (srv *Server) Inject(plugins ...plugin.Plugin) (injectSuccess bool) {
	defer func() {
		if err := recover(); err != nil {
			sugar.PrintStack()
			srv.Logger.Error("panic error", "error", err)
			srv.Terminate()
		} else if injectSuccess == false {
			srv.Terminate()
		}
	}()

	for _, plg := range plugins {
		plg = plg.Configuration(srv.Logger, srv.FetchConfig, srv.Cfg)
		if plg == nil {
			return false
		}
		plg = plg.Inject(srv.Module)
		if plg == nil {
			return false
		}
		srv.plugins = append(srv.plugins, plg)
	}
	return true
}

func (srv *Server) Serve(port string) {
	defer func() {
		if err := recover(); err != nil {
			sugar.PrintStack()
			srv.Logger.Error("panic error", "error", err)
			srv.Terminate()
		}
	}()

	control.BuildHttp(srv.Router.Root, srv.HttpEngine)
	srv.Module.Debug(srv.Logger)

	ctx, cancel := context.WithCancel(context.Background())
	defer func() {
		cancel()
	}()

	for _, plg := range srv.plugins {
		go plg.Work(ctx)
	}

	if err := srv.DatabaseModule.GetRawSQLInstance().Ping(); err != nil {
		srv.Logger.Debug("database died", "error", err)
		return
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		if err := srv.HttpEngine.Run(port); err != nil {
			srv.Logger.Debug("IRouter run error", "error", err)
		}
		wg.Done()
	}()

	//do something
	wg.Wait()
}

func (srv *Server) ServeTLS(port, crtFile, privateKeyFile string) {
	defer func() {
		if err := recover(); err != nil {
			sugar.PrintStack()
			srv.Logger.Error("panic error", "error", err)
			srv.Terminate()
		}
	}()

	control.BuildHttp(srv.Router.Root, srv.HttpEngine)
	srv.Module.Debug(srv.Logger)

	ctx, cancel := context.WithCancel(context.Background())
	defer func() {
		cancel()
	}()

	for _, plg := range srv.plugins {
		go plg.Work(ctx)
	}

	if err := srv.DatabaseModule.GetRawSQLInstance().Ping(); err != nil {
		srv.Logger.Debug("database died", "error", err)
		return
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		if err := http.ListenAndServeTLS(port, crtFile, privateKeyFile, srv.HttpEngine); err != nil {
			srv.Logger.Debug("IRouter run error", "error", err)
		}
		wg.Done()
	}()

	//do something
	wg.Wait()
}

func (srv *Server) WithPProf() *Server {
	ginpprof.Wrap(srv.HttpEngine)
	return srv
}
