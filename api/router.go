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
	UserService         *UserServiceRouter
	AnnouncementService *AnnouncementServiceRouter
	CommentService      *CommentServiceRouter
	SubmissionService   *SubmissionServiceRouter
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

	r.UserService = NewUserServiceRouter(traits, r.H)
	r.AnnouncementService = NewAnnouncementServiceRouter(traits, r.H)
	r.CommentService = NewCommentServiceRouter(traits, r.H)
	r.SubmissionService = NewSubmissionServiceRouter(traits, r.H)

	traits.AfterBuild(r)
	traits.ApplyAuth(r)
	return
}

type UserServiceRouter struct {
	H
	List         *UserServiceListRouter
	Count        *UserServiceCountRouter
	Register     *UserServiceRegisterRouter
	Login        *UserServiceLoginRouter
	RefreshToken *UserServiceRefreshTokenRouter
	IdGroup      *UserServiceIdGroupRouter
}

func NewUserServiceRouter(traits GenerateRouterTraits, h H) (r *UserServiceRouter) {
	r = &UserServiceRouter{
		H: &BaseH{
			Router:     h.GetRouter().Extend("UserService"),
			AuthRouter: h.GetAuthRouter().Extend("UserService"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), ""),
		},
	}

	r.List = NewUserServiceListRouter(traits, r.H)
	r.Count = NewUserServiceCountRouter(traits, r.H)
	r.Register = NewUserServiceRegisterRouter(traits, r.H)
	r.Login = NewUserServiceLoginRouter(traits, r.H)
	r.RefreshToken = NewUserServiceRefreshTokenRouter(traits, r.H)
	r.IdGroup = NewUserServiceIdGroupRouter(traits, r.H)

	return
}

type UserServiceListRouter struct {
	H

	ListUsers *LeafRouter
}

func NewUserServiceListRouter(traits GenerateRouterTraits, h H) (r *UserServiceListRouter) {
	r = &UserServiceListRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("user-list"),
			AuthRouter: h.GetAuthRouter().Group("user-list"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), ""),
		},
	}

	r.ListUsers = r.GetRouter().GET("", traits.GetServiceInstance("UserService").(UserService).ListUsers)

	return
}

type UserServiceCountRouter struct {
	H

	CountUser *LeafRouter
}

func NewUserServiceCountRouter(traits GenerateRouterTraits, h H) (r *UserServiceCountRouter) {
	r = &UserServiceCountRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("user-count"),
			AuthRouter: h.GetAuthRouter().Group("user-count"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), ""),
		},
	}

	r.CountUser = r.GetRouter().GET("", traits.GetServiceInstance("UserService").(UserService).CountUser)

	return
}

type UserServiceRegisterRouter struct {
	H

	Register *LeafRouter
}

func NewUserServiceRegisterRouter(traits GenerateRouterTraits, h H) (r *UserServiceRegisterRouter) {
	r = &UserServiceRegisterRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("user/register"),
			AuthRouter: h.GetAuthRouter().Group("user/register"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), ""),
		},
	}

	r.Register = r.GetRouter().POST("", traits.GetServiceInstance("UserService").(UserService).Register)

	return
}

type UserServiceLoginRouter struct {
	H

	LoginUser *LeafRouter
}

func NewUserServiceLoginRouter(traits GenerateRouterTraits, h H) (r *UserServiceLoginRouter) {
	r = &UserServiceLoginRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("user/login"),
			AuthRouter: h.GetAuthRouter().Group("user/login"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), ""),
		},
	}

	r.LoginUser = r.GetRouter().POST("", traits.GetServiceInstance("UserService").(UserService).LoginUser)

	return
}

type UserServiceRefreshTokenRouter struct {
	H

	RefreshToken *LeafRouter
}

func NewUserServiceRefreshTokenRouter(traits GenerateRouterTraits, h H) (r *UserServiceRefreshTokenRouter) {
	r = &UserServiceRefreshTokenRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("user-token"),
			AuthRouter: h.GetAuthRouter().Group("user-token"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), ""),
		},
	}

	r.RefreshToken = r.GetRouter().GET("", traits.GetServiceInstance("UserService").(UserService).RefreshToken)

	return
}

type UserServiceIdGroupRouter struct {
	H
	Email   *IdGroupEmailRouter
	Inspect *IdGroupInspectRouter

	GetUser *LeafRouter
	PutUser *LeafRouter
	Delete  *LeafRouter
}

func NewUserServiceIdGroupRouter(traits GenerateRouterTraits, h H) (r *UserServiceIdGroupRouter) {
	r = &UserServiceIdGroupRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("user/:id"),
			AuthRouter: h.GetAuthRouter().Group("user/:id"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), "user:id"),
		},
	}

	r.Email = NewIdGroupEmailRouter(traits, r.H)
	r.Inspect = NewIdGroupInspectRouter(traits, r.H)

	r.GetUser = r.GetRouter().GET("", traits.GetServiceInstance("UserService").(UserService).GetUser)
	r.PutUser = r.GetRouter().PUT("", traits.GetServiceInstance("UserService").(UserService).PutUser)
	r.Delete = r.GetRouter().DELETE("", traits.GetServiceInstance("UserService").(UserService).Delete)

	return
}

type IdGroupEmailRouter struct {
	H

	BindEmail *LeafRouter
}

func NewIdGroupEmailRouter(traits GenerateRouterTraits, h H) (r *IdGroupEmailRouter) {
	r = &IdGroupEmailRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("/email"),
			AuthRouter: h.GetAuthRouter().Group("/email"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), ""),
		},
	}

	r.BindEmail = r.GetRouter().PUT("", traits.GetServiceInstance("UserService").(UserService).BindEmail)

	return
}

type IdGroupInspectRouter struct {
	H

	InspectUser *LeafRouter
}

func NewIdGroupInspectRouter(traits GenerateRouterTraits, h H) (r *IdGroupInspectRouter) {
	r = &IdGroupInspectRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("/inspect"),
			AuthRouter: h.GetAuthRouter().Group("/inspect"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), ""),
		},
	}

	r.InspectUser = r.GetRouter().GET("", traits.GetServiceInstance("UserService").(UserService).InspectUser)

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

	r.PostAnnouncement = r.GetAuthRouter().POST("", traits.GetServiceInstance("AnnouncementService").(AnnouncementService).PostAnnouncement)
	r.PostAnnouncement = traits.ApplyAuthOnMethod(r.PostAnnouncement, "~")

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

type CommentServiceRouter struct {
	H
	List    *CommentServiceListRouter
	Count   *CommentServiceCountRouter
	Post    *CommentServicePostRouter
	IdGroup *CommentServiceIdGroupRouter
}

func NewCommentServiceRouter(traits GenerateRouterTraits, h H) (r *CommentServiceRouter) {
	r = &CommentServiceRouter{
		H: &BaseH{
			Router:     h.GetRouter().Extend("CommentService"),
			AuthRouter: h.GetAuthRouter().Extend("CommentService"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), ""),
		},
	}

	r.List = NewCommentServiceListRouter(traits, r.H)
	r.Count = NewCommentServiceCountRouter(traits, r.H)
	r.Post = NewCommentServicePostRouter(traits, r.H)
	r.IdGroup = NewCommentServiceIdGroupRouter(traits, r.H)

	return
}

type CommentServiceListRouter struct {
	H

	ListComments *LeafRouter
}

func NewCommentServiceListRouter(traits GenerateRouterTraits, h H) (r *CommentServiceListRouter) {
	r = &CommentServiceListRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("comment-list"),
			AuthRouter: h.GetAuthRouter().Group("comment-list"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), ""),
		},
	}

	r.ListComments = r.GetRouter().GET("", traits.GetServiceInstance("CommentService").(CommentService).ListComments)

	return
}

type CommentServiceCountRouter struct {
	H

	CountComment *LeafRouter
}

func NewCommentServiceCountRouter(traits GenerateRouterTraits, h H) (r *CommentServiceCountRouter) {
	r = &CommentServiceCountRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("comment-count"),
			AuthRouter: h.GetAuthRouter().Group("comment-count"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), ""),
		},
	}

	r.CountComment = r.GetRouter().GET("", traits.GetServiceInstance("CommentService").(CommentService).CountComment)

	return
}

type CommentServicePostRouter struct {
	H

	PostComment *LeafRouter
}

func NewCommentServicePostRouter(traits GenerateRouterTraits, h H) (r *CommentServicePostRouter) {
	r = &CommentServicePostRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("comment"),
			AuthRouter: h.GetAuthRouter().Group("comment"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), ""),
		},
	}

	r.PostComment = r.GetRouter().POST("", traits.GetServiceInstance("CommentService").(CommentService).PostComment)

	return
}

type CommentServiceIdGroupRouter struct {
	H

	GetComment *LeafRouter
	PutComment *LeafRouter
	Delete     *LeafRouter
}

func NewCommentServiceIdGroupRouter(traits GenerateRouterTraits, h H) (r *CommentServiceIdGroupRouter) {
	r = &CommentServiceIdGroupRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("comment/:cmid"),
			AuthRouter: h.GetAuthRouter().Group("comment/:cmid"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), "comment:cmid"),
		},
	}

	r.GetComment = r.GetRouter().GET("", traits.GetServiceInstance("CommentService").(CommentService).GetComment)
	r.PutComment = r.GetRouter().PUT("", traits.GetServiceInstance("CommentService").(CommentService).PutComment)
	r.Delete = r.GetRouter().DELETE("", traits.GetServiceInstance("CommentService").(CommentService).Delete)

	return
}

type SubmissionServiceRouter struct {
	H
	List    *SubmissionServiceListRouter
	Count   *SubmissionServiceCountRouter
	Post    *SubmissionServicePostRouter
	IdGroup *SubmissionServiceIdGroupRouter
}

func NewSubmissionServiceRouter(traits GenerateRouterTraits, h H) (r *SubmissionServiceRouter) {
	r = &SubmissionServiceRouter{
		H: &BaseH{
			Router:     h.GetRouter().Extend("SubmissionService"),
			AuthRouter: h.GetAuthRouter().Extend("SubmissionService"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), ""),
		},
	}

	r.List = NewSubmissionServiceListRouter(traits, r.H)
	r.Count = NewSubmissionServiceCountRouter(traits, r.H)
	r.Post = NewSubmissionServicePostRouter(traits, r.H)
	r.IdGroup = NewSubmissionServiceIdGroupRouter(traits, r.H)

	return
}

type SubmissionServiceListRouter struct {
	H

	ListSubmissions *LeafRouter
}

func NewSubmissionServiceListRouter(traits GenerateRouterTraits, h H) (r *SubmissionServiceListRouter) {
	r = &SubmissionServiceListRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("submission-list"),
			AuthRouter: h.GetAuthRouter().Group("submission-list"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), ""),
		},
	}

	r.ListSubmissions = r.GetRouter().GET("", traits.GetServiceInstance("SubmissionService").(SubmissionService).ListSubmissions)

	return
}

type SubmissionServiceCountRouter struct {
	H

	CountSubmissions *LeafRouter
}

func NewSubmissionServiceCountRouter(traits GenerateRouterTraits, h H) (r *SubmissionServiceCountRouter) {
	r = &SubmissionServiceCountRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("submission-count"),
			AuthRouter: h.GetAuthRouter().Group("submission-count"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), ""),
		},
	}

	r.CountSubmissions = r.GetRouter().GET("", traits.GetServiceInstance("SubmissionService").(SubmissionService).CountSubmissions)

	return
}

type SubmissionServicePostRouter struct {
	H

	PostSubmission *LeafRouter
}

func NewSubmissionServicePostRouter(traits GenerateRouterTraits, h H) (r *SubmissionServicePostRouter) {
	r = &SubmissionServicePostRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("/problem/:pid/submission"),
			AuthRouter: h.GetAuthRouter().Group("/problem/:pid/submission"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), "problem:pid"),
		},
	}

	r.PostSubmission = r.GetAuthRouter().POST("", traits.GetServiceInstance("SubmissionService").(SubmissionService).PostSubmission)
	r.PostSubmission = traits.ApplyAuthOnMethod(r.PostSubmission, "~")

	return
}

type SubmissionServiceIdGroupRouter struct {
	H
	GetContent *IdGroupGetContentRouter

	GetSubmission *LeafRouter
	Delete        *LeafRouter
}

func NewSubmissionServiceIdGroupRouter(traits GenerateRouterTraits, h H) (r *SubmissionServiceIdGroupRouter) {
	r = &SubmissionServiceIdGroupRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("submission/:sid"),
			AuthRouter: h.GetAuthRouter().Group("submission/:sid"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), "submission:sid"),
		},
	}

	r.GetContent = NewIdGroupGetContentRouter(traits, r.H)

	r.GetSubmission = r.GetRouter().GET("", traits.GetServiceInstance("SubmissionService").(SubmissionService).GetSubmission)
	r.Delete = r.GetRouter().DELETE("", traits.GetServiceInstance("SubmissionService").(SubmissionService).Delete)

	return
}

type IdGroupGetContentRouter struct {
	H

	GetContent *LeafRouter
}

func NewIdGroupGetContentRouter(traits GenerateRouterTraits, h H) (r *IdGroupGetContentRouter) {
	r = &IdGroupGetContentRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("/content"),
			AuthRouter: h.GetAuthRouter().Group("/content"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), ""),
		},
	}

	r.GetContent = r.GetRouter().GET("", traits.GetServiceInstance("SubmissionService").(SubmissionService).GetContent)

	return
}
