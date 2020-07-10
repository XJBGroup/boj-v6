package server

import (
	"github.com/Myriad-Dreamin/boj-v6/api"
	"github.com/Myriad-Dreamin/boj-v6/app/provider"
	"github.com/Myriad-Dreamin/boj-v6/config"
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	"github.com/Myriad-Dreamin/minimum-lib/module"
	"github.com/gin-gonic/gin"
	"strings"
)

type routerTraits struct {
	module.Module
	*provider.Service
}

func (rt routerTraits) GetJWTMiddleware() api.HandlerFunc {
	return rt.Module.Require(config.ModulePath.Middleware.JWT).(api.HandlerFunc)
}

func (rt routerTraits) GetAuthMiddleware() *api.Middleware {
	return rt.Module.Require(
		config.ModulePath.Middleware.RouteAuth).(*controller.Middleware)
}

func (rt routerTraits) AfterBuild(_ *api.RootRouter) {
}

func (rt routerTraits) ApplyAuth(r *api.RootRouter) {
	as := r.AnnouncementService
	as.Post.PostAnnouncement.Use(
		as.Post.GetAuth().AdminOnly())
	asi := as.IdGroup
	asi.PutAnnouncement.Use(asi.GetAuth().AdminOnly())
	asi.Delete.Use(asi.GetAuth().AdminOnly())
}

func (rt routerTraits) ApplyAuthOnMethod(r *api.LeafRouter, authMeta string) *api.LeafRouter {
	print(authMeta)
	return r
}

func (rt routerTraits) ApplyRouteMeta(m *api.Middleware, routeMeta string) *api.Middleware {
	var rm = strings.Split(routeMeta, ":")
	switch rm[0] {
	case "announcement":
		return m.MustGroup("announcement", rm[1])
	case "":
		return m.Copy()
	default:
		panic("unknown meta " + routeMeta)
	}
}

func (rt routerTraits) GetServiceInstance(svcName string) interface{} {
	switch svcName {
	case "AnnouncementService":
		return rt.Service.AnnouncementService()
	default:
		panic(svcName + " not found")
	}
}

func newTraitsHelper(m module.Module, s *provider.Service) routerTraits {
	return routerTraits{m, s}
}

func (srv *Server) BuildRouter() bool {
	gin.DefaultErrorWriter = srv.LoggerWriter
	gin.DefaultWriter = srv.LoggerWriter
	srv.HttpEngine = gin.New()
	srv.HttpEngine.Use(gin.LoggerWithConfig(gin.LoggerConfig{
		Output: srv.LoggerWriter,
	}), gin.Recovery())
	srv.HttpEngine.Use(srv.corsMW)

	srv.Router = api.NewRootRouter(newTraitsHelper(srv.Module, srv.ServiceProvider))
	srv.Module.Provide(config.ModulePath.Global.HttpEngine, srv.HttpEngine)
	return true
}
