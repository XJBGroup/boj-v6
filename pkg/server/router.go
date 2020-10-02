package server

import (
	"github.com/Myriad-Dreamin/boj-v6/api"
	"github.com/Myriad-Dreamin/boj-v6/config"
	"github.com/Myriad-Dreamin/boj-v6/lib/jwt"
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	"github.com/Myriad-Dreamin/minimum-lib/module"
	"github.com/gin-gonic/gin"
	"path"
	"strings"
)

type routerTraits struct {
	module.Module
}

func (rt routerTraits) GetJWTMiddleware() api.HandlerFunc {
	return rt.Module.RequireImpl(new(*jwt.Middleware)).(*jwt.Middleware).Build()
}

func (rt routerTraits) GetAuthMiddleware() *api.Middleware {
	return rt.Module.RequireImpl(new(*controller.Middleware)).(*controller.Middleware)
}

func (rt routerTraits) AfterBuild(_ *api.RootRouter) {
}

func (rt routerTraits) ApplyAuth(r *api.RootRouter) {
	as := r.AnnouncementController
	as.Post.PostAnnouncement.Use(
		as.Post.GetAuth().AdminOnly())
	asi := as.IdGroup
	asi.PutAnnouncement.Use(asi.GetAuth().AdminOnly())
	asi.DeleteAnnouncement.Use(asi.GetAuth().AdminOnly())
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
	case "user":
		return m.MustGroup("user", rm[1])
	case "comment":
		return m.MustGroup("comment", rm[1])
	case "problem":
		return m.MustGroup("problem", rm[1])
	case "contest":
		return m.MustGroup("contest", rm[1])
	case "submission":
		return m.MustGroup("submission", rm[1])
	case "group":
		return m.MustGroup("group", rm[1])
	case "~":
		fallthrough
	case "":
		return m.Copy()
	default:
		panic("unknown meta " + routeMeta)
	}
}

func (rt routerTraits) GetServiceInstance(svcName string) interface{} {
	return rt.Module.Require(path.Join("minimum", svcName))
}

func newTraitsHelper(m module.Module) routerTraits {
	return routerTraits{m}
}

func (srv *Server) BuildRouter(mock bool) bool {
	gin.DefaultErrorWriter = srv.LoggerWriter
	gin.DefaultWriter = srv.LoggerWriter
	srv.HttpEngine = gin.New()
	srv.HttpEngine.Use(gin.LoggerWithConfig(gin.LoggerConfig{
		Output: srv.LoggerWriter,
	}), gin.Recovery())
	//if !mock {

	srv.HttpEngine.Use(srv.corsMW)
	//}

	srv.Router = api.NewRootRouter(newTraitsHelper(srv.Module))

	srv.Router.GetRouter().StaticFile("/swagger/main_spec.json", "docs/main_spec.json")

	srv.Module.Provide(config.ModulePath.Global.HttpEngine, srv.HttpEngine)
	return true
}
