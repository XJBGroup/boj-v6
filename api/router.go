package api

import (
	"github.com/Myriad-Dreamin/minimum-lib/controller"
)

type Router = controller.Router
type Middleware = controller.Middleware
type LeafRouter = controller.LeafRouter
type HandlerFunc = controller.HandlerFunc

type H interface {
	GetRouter() *Router
	GetAuthRouter() *Router
	GetAuth() *Middleware
}

type BaseH struct {
	*Router
	AuthRouter *Router
	Auth       *Middleware
}

func (r *BaseH) GetRouter() *Router {
	return r.Router
}

func (r *BaseH) GetAuthRouter() *Router {
	return r.AuthRouter
}

func (r *BaseH) GetAuth() *Middleware {
	return r.Auth
}

type GenerateRouterTraits interface {
	GetJWTMiddleware() HandlerFunc
	GetAuthMiddleware() *Middleware
	AfterBuild(r *RootRouter)
	ApplyAuth(r *RootRouter)
	ApplyAuthOnMethod(r *LeafRouter, authMeta string) *LeafRouter

	ApplyRouteMeta(m *Middleware, routeMeta string) *Middleware
	GetServiceInstance(svcName string) interface{}
}

type RootRouter struct {
	H
	Root                *Router
	AnnouncementService *AnnouncementServiceRouter
	Ping                *LeafRouter
	//Images   *LeafRouter
	//Musics   *LeafRouter
	//Articles *LeafRouter
}

// @title Ping
// @description result
func PingFunc(c controller.MContext) {
	c.JSON(200, map[string]interface{}{
		"message": "pong",
	})
}

func NewRootRouter(traits GenerateRouterTraits) (r *RootRouter) {
	rr := controller.NewRouterGroup()
	apiRouterV1 := rr.Group("/v1")
	authRouterV1 := apiRouterV1.Group("", traits.GetJWTMiddleware())

	r = &RootRouter{
		Root: rr,
		H: &BaseH{
			Router:     apiRouterV1,
			AuthRouter: authRouterV1,
			Auth:       traits.GetAuthMiddleware(),
		},
	}

	r.Ping = r.Root.GET("/ping", PingFunc)

	r.AnnouncementService = NewAnnouncementServiceRouter(traits, r.H)

	traits.AfterBuild(r)
	traits.ApplyAuth(r)
	return
}

type AnnouncementServiceRouter struct {
	H
	List    *AnnouncementServiceListRouter
	Count   *AnnouncementServiceCountRouter
	Post    *AnnouncementServicePostRouter
	IdGroup *AnnouncementServiceIdGroupRouter
}

func NewAnnouncementServiceRouter(traits GenerateRouterTraits, h H) (r *AnnouncementServiceRouter) {
	r = &AnnouncementServiceRouter{
		H: &BaseH{
			Router:     h.GetRouter().Extend("AnnouncementService"),
			AuthRouter: h.GetAuthRouter().Extend("AnnouncementService"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), ""),
		},
	}

	r.List = NewAnnouncementServiceListRouter(traits, r.H)
	r.Count = NewAnnouncementServiceCountRouter(traits, r.H)
	r.Post = NewAnnouncementServicePostRouter(traits, r.H)
	r.IdGroup = NewAnnouncementServiceIdGroupRouter(traits, r.H)

	return
}

type AnnouncementServiceListRouter struct {
	H

	ListAnnouncements *LeafRouter
}

func NewAnnouncementServiceListRouter(traits GenerateRouterTraits, h H) (r *AnnouncementServiceListRouter) {
	r = &AnnouncementServiceListRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("announcement-list"),
			AuthRouter: h.GetAuthRouter().Group("announcement-list"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), ""),
		},
	}

	r.ListAnnouncements = r.GetRouter().GET("", traits.GetServiceInstance("AnnouncementService").(AnnouncementService).ListAnnouncements)

	return
}

type AnnouncementServiceCountRouter struct {
	H

	CountAnnouncement *LeafRouter
}

func NewAnnouncementServiceCountRouter(traits GenerateRouterTraits, h H) (r *AnnouncementServiceCountRouter) {
	r = &AnnouncementServiceCountRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("announcement-count"),
			AuthRouter: h.GetAuthRouter().Group("announcement-count"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), ""),
		},
	}

	r.CountAnnouncement = r.GetRouter().GET("", traits.GetServiceInstance("AnnouncementService").(AnnouncementService).CountAnnouncement)

	return
}

type AnnouncementServicePostRouter struct {
	H

	PostAnnouncement *LeafRouter
}

func NewAnnouncementServicePostRouter(traits GenerateRouterTraits, h H) (r *AnnouncementServicePostRouter) {
	r = &AnnouncementServicePostRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("announcement"),
			AuthRouter: h.GetAuthRouter().Group("announcement"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), ""),
		},
	}

	r.PostAnnouncement = r.GetRouter().POST("", traits.GetServiceInstance("AnnouncementService").(AnnouncementService).PostAnnouncement)

	return
}

type AnnouncementServiceIdGroupRouter struct {
	H

	GetAnnouncement *LeafRouter
	PutAnnouncement *LeafRouter
	Delete          *LeafRouter
}

func NewAnnouncementServiceIdGroupRouter(traits GenerateRouterTraits, h H) (r *AnnouncementServiceIdGroupRouter) {
	r = &AnnouncementServiceIdGroupRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("announcement/:aid"),
			AuthRouter: h.GetAuthRouter().Group("announcement/:aid"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), "announcement:aid"),
		},
	}

	r.GetAnnouncement = r.GetRouter().GET("", traits.GetServiceInstance("AnnouncementService").(AnnouncementService).GetAnnouncement)
	r.PutAnnouncement = r.GetRouter().PUT("", traits.GetServiceInstance("AnnouncementService").(AnnouncementService).PutAnnouncement)
	r.Delete = r.GetRouter().DELETE("", traits.GetServiceInstance("AnnouncementService").(AnnouncementService).Delete)

	return
}
