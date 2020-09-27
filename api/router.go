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
	Root                   *Router
	ProblemController      *ProblemControllerRouter
	ContestController      *ContestControllerRouter
	GroupController        *GroupControllerRouter
	UserController         *UserControllerRouter
	AuthController         *AuthControllerRouter
	AnnouncementController *AnnouncementControllerRouter
	CommentController      *CommentControllerRouter
	SubmissionController   *SubmissionControllerRouter
	Ping                   *LeafRouter
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

	r.ProblemController = NewProblemControllerRouter(traits, r.H)
	r.ContestController = NewContestControllerRouter(traits, r.H)
	r.GroupController = NewGroupControllerRouter(traits, r.H)
	r.UserController = NewUserControllerRouter(traits, r.H)
	r.AuthController = NewAuthControllerRouter(traits, r.H)
	r.AnnouncementController = NewAnnouncementControllerRouter(traits, r.H)
	r.CommentController = NewCommentControllerRouter(traits, r.H)
	r.SubmissionController = NewSubmissionControllerRouter(traits, r.H)

	traits.AfterBuild(r)
	traits.ApplyAuth(r)
	return
}

type ProblemControllerRouter struct {
	H
	List    *ProblemControllerListRouter
	Count   *ProblemControllerCountRouter
	Post    *ProblemControllerPostRouter
	IdGroup *ProblemControllerIdGroupRouter
}

func NewProblemControllerRouter(traits GenerateRouterTraits, h H) (r *ProblemControllerRouter) {
	r = &ProblemControllerRouter{
		H: &BaseH{
			Router:     h.GetRouter().Extend("ProblemController"),
			AuthRouter: h.GetAuthRouter().Extend("ProblemController"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), ""),
		},
	}

	r.List = NewProblemControllerListRouter(traits, r.H)
	r.Count = NewProblemControllerCountRouter(traits, r.H)
	r.Post = NewProblemControllerPostRouter(traits, r.H)
	r.IdGroup = NewProblemControllerIdGroupRouter(traits, r.H)

	return
}

type ProblemControllerListRouter struct {
	H

	ListProblem *LeafRouter
}

func NewProblemControllerListRouter(traits GenerateRouterTraits, h H) (r *ProblemControllerListRouter) {
	r = &ProblemControllerListRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("problem-list"),
			AuthRouter: h.GetAuthRouter().Group("problem-list"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), ""),
		},
	}

	r.ListProblem = r.GetRouter().GET("", traits.GetServiceInstance("ProblemController").(ProblemController).ListProblem)

	return
}

type ProblemControllerCountRouter struct {
	H

	CountProblem *LeafRouter
}

func NewProblemControllerCountRouter(traits GenerateRouterTraits, h H) (r *ProblemControllerCountRouter) {
	r = &ProblemControllerCountRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("problem-count"),
			AuthRouter: h.GetAuthRouter().Group("problem-count"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), ""),
		},
	}

	r.CountProblem = r.GetRouter().GET("", traits.GetServiceInstance("ProblemController").(ProblemController).CountProblem)

	return
}

type ProblemControllerPostRouter struct {
	H

	PostProblem *LeafRouter
}

func NewProblemControllerPostRouter(traits GenerateRouterTraits, h H) (r *ProblemControllerPostRouter) {
	r = &ProblemControllerPostRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("problem"),
			AuthRouter: h.GetAuthRouter().Group("problem"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), ""),
		},
	}

	r.PostProblem = r.GetAuthRouter().POST("", traits.GetServiceInstance("ProblemController").(ProblemController).PostProblem)
	r.PostProblem = traits.ApplyAuthOnMethod(r.PostProblem, "~")

	return
}

type ProblemControllerIdGroupRouter struct {
	H
	CountProblemDesc *ProblemControllerIdGroupCountProblemDescRouter
	ProblemDesc      *ProblemControllerIdGroupProblemDescRouter
	ProblemFS        *ProblemControllerIdGroupProblemFSRouter
	ListProblemDesc  *ProblemControllerIdGroupListProblemDescRouter

	GetProblem    *LeafRouter
	PutProblem    *LeafRouter
	DeleteProblem *LeafRouter
}

func NewProblemControllerIdGroupRouter(traits GenerateRouterTraits, h H) (r *ProblemControllerIdGroupRouter) {
	r = &ProblemControllerIdGroupRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("problem/:pid"),
			AuthRouter: h.GetAuthRouter().Group("problem/:pid"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), "problem:pid"),
		},
	}

	r.CountProblemDesc = NewProblemControllerIdGroupCountProblemDescRouter(traits, r.H)
	r.ProblemDesc = NewProblemControllerIdGroupProblemDescRouter(traits, r.H)
	r.ProblemFS = NewProblemControllerIdGroupProblemFSRouter(traits, r.H)
	r.ListProblemDesc = NewProblemControllerIdGroupListProblemDescRouter(traits, r.H)

	r.GetProblem = r.GetRouter().GET("", traits.GetServiceInstance("ProblemController").(ProblemController).GetProblem)
	r.PutProblem = r.GetRouter().PUT("", traits.GetServiceInstance("ProblemController").(ProblemController).PutProblem)
	r.DeleteProblem = r.GetRouter().DELETE("", traits.GetServiceInstance("ProblemController").(ProblemController).DeleteProblem)

	return
}

type ProblemControllerIdGroupCountProblemDescRouter struct {
	H

	CountProblemDesc *LeafRouter
}

func NewProblemControllerIdGroupCountProblemDescRouter(traits GenerateRouterTraits, h H) (r *ProblemControllerIdGroupCountProblemDescRouter) {
	r = &ProblemControllerIdGroupCountProblemDescRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("/desc-count"),
			AuthRouter: h.GetAuthRouter().Group("/desc-count"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), ""),
		},
	}

	r.CountProblemDesc = r.GetRouter().GET("", traits.GetServiceInstance("ProblemController").(ProblemController).CountProblemDesc)

	return
}

type ProblemControllerIdGroupProblemDescRouter struct {
	H
	ProblemDesc *ProblemControllerIdGroupProblemDescProblemDescRouter

	PostProblemDesc   *LeafRouter
	GetProblemDesc    *LeafRouter
	PutProblemDesc    *LeafRouter
	DeleteProblemDesc *LeafRouter
}

func NewProblemControllerIdGroupProblemDescRouter(traits GenerateRouterTraits, h H) (r *ProblemControllerIdGroupProblemDescRouter) {
	r = &ProblemControllerIdGroupProblemDescRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("/desc"),
			AuthRouter: h.GetAuthRouter().Group("/desc"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), ""),
		},
	}

	r.ProblemDesc = NewProblemControllerIdGroupProblemDescProblemDescRouter(traits, r.H)

	r.PostProblemDesc = r.GetRouter().POST("", traits.GetServiceInstance("ProblemController").(ProblemController).PostProblemDesc)
	r.GetProblemDesc = r.GetRouter().GET("", traits.GetServiceInstance("ProblemController").(ProblemController).GetProblemDesc)
	r.PutProblemDesc = r.GetRouter().PUT("", traits.GetServiceInstance("ProblemController").(ProblemController).PutProblemDesc)
	r.DeleteProblemDesc = r.GetRouter().DELETE("", traits.GetServiceInstance("ProblemController").(ProblemController).DeleteProblemDesc)

	return
}

type ProblemControllerIdGroupProblemDescProblemDescRouter struct {
	H

	ChangeProblemDescriptionRef *LeafRouter
}

func NewProblemControllerIdGroupProblemDescProblemDescRouter(traits GenerateRouterTraits, h H) (r *ProblemControllerIdGroupProblemDescProblemDescRouter) {
	r = &ProblemControllerIdGroupProblemDescProblemDescRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("/ref"),
			AuthRouter: h.GetAuthRouter().Group("/ref"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), ""),
		},
	}

	r.ChangeProblemDescriptionRef = r.GetRouter().POST("", traits.GetServiceInstance("ProblemController").(ProblemController).ChangeProblemDescriptionRef)

	return
}

type ProblemControllerIdGroupProblemFSRouter struct {
	H
	ProblemFSFileOperation      *ProblemControllerIdGroupProblemFSProblemFSFileOperationRouter
	ProblemFSDirectoryOperation *ProblemControllerIdGroupProblemFSProblemFSDirectoryOperationRouter
	ProblemFSConfigOperation    *ProblemControllerIdGroupProblemFSProblemFSConfigOperationRouter
}

func NewProblemControllerIdGroupProblemFSRouter(traits GenerateRouterTraits, h H) (r *ProblemControllerIdGroupProblemFSRouter) {
	r = &ProblemControllerIdGroupProblemFSRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("/fs"),
			AuthRouter: h.GetAuthRouter().Group("/fs"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), ""),
		},
	}

	r.ProblemFSFileOperation = NewProblemControllerIdGroupProblemFSProblemFSFileOperationRouter(traits, r.H)
	r.ProblemFSDirectoryOperation = NewProblemControllerIdGroupProblemFSProblemFSDirectoryOperationRouter(traits, r.H)
	r.ProblemFSConfigOperation = NewProblemControllerIdGroupProblemFSProblemFSConfigOperationRouter(traits, r.H)

	return
}

type ProblemControllerIdGroupProblemFSProblemFSFileOperationRouter struct {
	H
	ProblemFSRead *ProblemControllerIdGroupProblemFSProblemFSFileOperationProblemFSReadRouter

	ProblemFSStat   *LeafRouter
	ProblemFSWrite  *LeafRouter
	ProblemFSRemove *LeafRouter
}

func NewProblemControllerIdGroupProblemFSProblemFSFileOperationRouter(traits GenerateRouterTraits, h H) (r *ProblemControllerIdGroupProblemFSProblemFSFileOperationRouter) {
	r = &ProblemControllerIdGroupProblemFSProblemFSFileOperationRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("file"),
			AuthRouter: h.GetAuthRouter().Group("file"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), ""),
		},
	}

	r.ProblemFSRead = NewProblemControllerIdGroupProblemFSProblemFSFileOperationProblemFSReadRouter(traits, r.H)

	r.ProblemFSStat = r.GetRouter().GET("", traits.GetServiceInstance("ProblemController").(ProblemController).ProblemFSStat)
	r.ProblemFSWrite = r.GetRouter().POST("", traits.GetServiceInstance("ProblemController").(ProblemController).ProblemFSWrite)
	r.ProblemFSRemove = r.GetRouter().DELETE("", traits.GetServiceInstance("ProblemController").(ProblemController).ProblemFSRemove)

	return
}

type ProblemControllerIdGroupProblemFSProblemFSFileOperationProblemFSReadRouter struct {
	H

	ProblemFSRead *LeafRouter
}

func NewProblemControllerIdGroupProblemFSProblemFSFileOperationProblemFSReadRouter(traits GenerateRouterTraits, h H) (r *ProblemControllerIdGroupProblemFSProblemFSFileOperationProblemFSReadRouter) {
	r = &ProblemControllerIdGroupProblemFSProblemFSFileOperationProblemFSReadRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("content"),
			AuthRouter: h.GetAuthRouter().Group("content"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), ""),
		},
	}

	r.ProblemFSRead = r.GetRouter().GET("", traits.GetServiceInstance("ProblemController").(ProblemController).ProblemFSRead)

	return
}

type ProblemControllerIdGroupProblemFSProblemFSDirectoryOperationRouter struct {
	H
	ProblemFSZipOperation *ProblemControllerIdGroupProblemFSProblemFSDirectoryOperationProblemFSZipOperationRouter

	ProblemFSLS        *LeafRouter
	ProblemFSWrites    *LeafRouter
	ProblemFSMkdir     *LeafRouter
	ProblemFSRemoveAll *LeafRouter
}

func NewProblemControllerIdGroupProblemFSProblemFSDirectoryOperationRouter(traits GenerateRouterTraits, h H) (r *ProblemControllerIdGroupProblemFSProblemFSDirectoryOperationRouter) {
	r = &ProblemControllerIdGroupProblemFSProblemFSDirectoryOperationRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("directory"),
			AuthRouter: h.GetAuthRouter().Group("directory"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), ""),
		},
	}

	r.ProblemFSZipOperation = NewProblemControllerIdGroupProblemFSProblemFSDirectoryOperationProblemFSZipOperationRouter(traits, r.H)

	r.ProblemFSLS = r.GetRouter().GET("", traits.GetServiceInstance("ProblemController").(ProblemController).ProblemFSLS)
	r.ProblemFSWrites = r.GetRouter().POST("", traits.GetServiceInstance("ProblemController").(ProblemController).ProblemFSWrites)
	r.ProblemFSMkdir = r.GetRouter().PUT("", traits.GetServiceInstance("ProblemController").(ProblemController).ProblemFSMkdir)
	r.ProblemFSRemoveAll = r.GetRouter().DELETE("", traits.GetServiceInstance("ProblemController").(ProblemController).ProblemFSRemoveAll)

	return
}

type ProblemControllerIdGroupProblemFSProblemFSDirectoryOperationProblemFSZipOperationRouter struct {
	H

	ProblemFSZipWrite *LeafRouter
	ProblemFSZipRead  *LeafRouter
}

func NewProblemControllerIdGroupProblemFSProblemFSDirectoryOperationProblemFSZipOperationRouter(traits GenerateRouterTraits, h H) (r *ProblemControllerIdGroupProblemFSProblemFSDirectoryOperationProblemFSZipOperationRouter) {
	r = &ProblemControllerIdGroupProblemFSProblemFSDirectoryOperationProblemFSZipOperationRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("zip"),
			AuthRouter: h.GetAuthRouter().Group("zip"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), ""),
		},
	}

	r.ProblemFSZipWrite = r.GetRouter().POST("", traits.GetServiceInstance("ProblemController").(ProblemController).ProblemFSZipWrite)
	r.ProblemFSZipRead = r.GetRouter().GET("", traits.GetServiceInstance("ProblemController").(ProblemController).ProblemFSZipRead)

	return
}

type ProblemControllerIdGroupProblemFSProblemFSConfigOperationRouter struct {
	H

	ProblemFSReadConfig  *LeafRouter
	ProblemFSWriteConfig *LeafRouter
	ProblemFSPutConfig   *LeafRouter
}

func NewProblemControllerIdGroupProblemFSProblemFSConfigOperationRouter(traits GenerateRouterTraits, h H) (r *ProblemControllerIdGroupProblemFSProblemFSConfigOperationRouter) {
	r = &ProblemControllerIdGroupProblemFSProblemFSConfigOperationRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("config"),
			AuthRouter: h.GetAuthRouter().Group("config"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), ""),
		},
	}

	r.ProblemFSReadConfig = r.GetRouter().GET("", traits.GetServiceInstance("ProblemController").(ProblemController).ProblemFSReadConfig)
	r.ProblemFSWriteConfig = r.GetRouter().POST("", traits.GetServiceInstance("ProblemController").(ProblemController).ProblemFSWriteConfig)
	r.ProblemFSPutConfig = r.GetRouter().PUT("", traits.GetServiceInstance("ProblemController").(ProblemController).ProblemFSPutConfig)

	return
}

type ProblemControllerIdGroupListProblemDescRouter struct {
	H

	ListProblemDesc *LeafRouter
}

func NewProblemControllerIdGroupListProblemDescRouter(traits GenerateRouterTraits, h H) (r *ProblemControllerIdGroupListProblemDescRouter) {
	r = &ProblemControllerIdGroupListProblemDescRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("/desc-list"),
			AuthRouter: h.GetAuthRouter().Group("/desc-list"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), ""),
		},
	}

	r.ListProblemDesc = r.GetRouter().GET("", traits.GetServiceInstance("ProblemController").(ProblemController).ListProblemDesc)

	return
}

type ContestControllerRouter struct {
	H
	List    *ContestControllerListRouter
	Count   *ContestControllerCountRouter
	Post    *ContestControllerPostRouter
	IdGroup *ContestControllerIdGroupRouter
}

func NewContestControllerRouter(traits GenerateRouterTraits, h H) (r *ContestControllerRouter) {
	r = &ContestControllerRouter{
		H: &BaseH{
			Router:     h.GetRouter().Extend("ContestController"),
			AuthRouter: h.GetAuthRouter().Extend("ContestController"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), ""),
		},
	}

	r.List = NewContestControllerListRouter(traits, r.H)
	r.Count = NewContestControllerCountRouter(traits, r.H)
	r.Post = NewContestControllerPostRouter(traits, r.H)
	r.IdGroup = NewContestControllerIdGroupRouter(traits, r.H)

	return
}

type ContestControllerListRouter struct {
	H

	ListContest *LeafRouter
}

func NewContestControllerListRouter(traits GenerateRouterTraits, h H) (r *ContestControllerListRouter) {
	r = &ContestControllerListRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("contest-list"),
			AuthRouter: h.GetAuthRouter().Group("contest-list"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), ""),
		},
	}

	r.ListContest = r.GetRouter().GET("", traits.GetServiceInstance("ContestController").(ContestController).ListContest)

	return
}

type ContestControllerCountRouter struct {
	H

	CountContest *LeafRouter
}

func NewContestControllerCountRouter(traits GenerateRouterTraits, h H) (r *ContestControllerCountRouter) {
	r = &ContestControllerCountRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("contest-count"),
			AuthRouter: h.GetAuthRouter().Group("contest-count"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), ""),
		},
	}

	r.CountContest = r.GetRouter().GET("", traits.GetServiceInstance("ContestController").(ContestController).CountContest)

	return
}

type ContestControllerPostRouter struct {
	H

	PostContest *LeafRouter
}

func NewContestControllerPostRouter(traits GenerateRouterTraits, h H) (r *ContestControllerPostRouter) {
	r = &ContestControllerPostRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("contest"),
			AuthRouter: h.GetAuthRouter().Group("contest"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), ""),
		},
	}

	r.PostContest = r.GetAuthRouter().POST("", traits.GetServiceInstance("ContestController").(ContestController).PostContest)
	r.PostContest = traits.ApplyAuthOnMethod(r.PostContest, "~")

	return
}

type ContestControllerIdGroupRouter struct {
	H
	ListContestUsers *ContestControllerIdGroupListContestUsersRouter
	List             *ContestControllerIdGroupListRouter
	Count            *ContestControllerIdGroupCountRouter
	Post             *ContestControllerIdGroupPostRouter
	IdGroup          *ContestControllerIdGroupIdGroupRouter

	GetContest    *LeafRouter
	PutContest    *LeafRouter
	DeleteContest *LeafRouter
}

func NewContestControllerIdGroupRouter(traits GenerateRouterTraits, h H) (r *ContestControllerIdGroupRouter) {
	r = &ContestControllerIdGroupRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("contest/:cid"),
			AuthRouter: h.GetAuthRouter().Group("contest/:cid"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), "contest:cid"),
		},
	}

	r.ListContestUsers = NewContestControllerIdGroupListContestUsersRouter(traits, r.H)
	r.List = NewContestControllerIdGroupListRouter(traits, r.H)
	r.Count = NewContestControllerIdGroupCountRouter(traits, r.H)
	r.Post = NewContestControllerIdGroupPostRouter(traits, r.H)
	r.IdGroup = NewContestControllerIdGroupIdGroupRouter(traits, r.H)

	r.GetContest = r.GetRouter().GET("", traits.GetServiceInstance("ContestController").(ContestController).GetContest)
	r.PutContest = r.GetAuthRouter().PUT("", traits.GetServiceInstance("ContestController").(ContestController).PutContest)
	r.PutContest = traits.ApplyAuthOnMethod(r.PutContest, "~")
	r.DeleteContest = r.GetRouter().DELETE("", traits.GetServiceInstance("ContestController").(ContestController).DeleteContest)

	return
}

type ContestControllerIdGroupListContestUsersRouter struct {
	H

	ListContestUsers *LeafRouter
}

func NewContestControllerIdGroupListContestUsersRouter(traits GenerateRouterTraits, h H) (r *ContestControllerIdGroupListContestUsersRouter) {
	r = &ContestControllerIdGroupListContestUsersRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("/user-list"),
			AuthRouter: h.GetAuthRouter().Group("/user-list"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), ""),
		},
	}

	r.ListContestUsers = r.GetRouter().GET("", traits.GetServiceInstance("ContestController").(ContestController).ListContestUsers)

	return
}

type ContestControllerIdGroupListRouter struct {
	H

	ListContestProblem *LeafRouter
}

func NewContestControllerIdGroupListRouter(traits GenerateRouterTraits, h H) (r *ContestControllerIdGroupListRouter) {
	r = &ContestControllerIdGroupListRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("problem-list"),
			AuthRouter: h.GetAuthRouter().Group("problem-list"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), ""),
		},
	}

	r.ListContestProblem = r.GetRouter().GET("", traits.GetServiceInstance("ContestController").(ContestController).ListContestProblem)

	return
}

type ContestControllerIdGroupCountRouter struct {
	H

	CountContestProblem *LeafRouter
}

func NewContestControllerIdGroupCountRouter(traits GenerateRouterTraits, h H) (r *ContestControllerIdGroupCountRouter) {
	r = &ContestControllerIdGroupCountRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("problem-count"),
			AuthRouter: h.GetAuthRouter().Group("problem-count"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), ""),
		},
	}

	r.CountContestProblem = r.GetRouter().GET("", traits.GetServiceInstance("ContestController").(ContestController).CountContestProblem)

	return
}

type ContestControllerIdGroupPostRouter struct {
	H

	PostContestProblem *LeafRouter
}

func NewContestControllerIdGroupPostRouter(traits GenerateRouterTraits, h H) (r *ContestControllerIdGroupPostRouter) {
	r = &ContestControllerIdGroupPostRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("problem"),
			AuthRouter: h.GetAuthRouter().Group("problem"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), ""),
		},
	}

	r.PostContestProblem = r.GetAuthRouter().POST("", traits.GetServiceInstance("ContestController").(ContestController).PostContestProblem)
	r.PostContestProblem = traits.ApplyAuthOnMethod(r.PostContestProblem, "~")

	return
}

type ContestControllerIdGroupIdGroupRouter struct {
	H
	ListContestProblemDesc  *ContestControllerIdGroupIdGroupListContestProblemDescRouter
	CountContestProblemDesc *ContestControllerIdGroupIdGroupCountContestProblemDescRouter
	ProblemDesc             *ContestControllerIdGroupIdGroupProblemDescRouter

	GetContestProblem    *LeafRouter
	PutContestProblem    *LeafRouter
	DeleteContestProblem *LeafRouter
}

func NewContestControllerIdGroupIdGroupRouter(traits GenerateRouterTraits, h H) (r *ContestControllerIdGroupIdGroupRouter) {
	r = &ContestControllerIdGroupIdGroupRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("problem/:pid"),
			AuthRouter: h.GetAuthRouter().Group("problem/:pid"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), "problem:pid"),
		},
	}

	r.ListContestProblemDesc = NewContestControllerIdGroupIdGroupListContestProblemDescRouter(traits, r.H)
	r.CountContestProblemDesc = NewContestControllerIdGroupIdGroupCountContestProblemDescRouter(traits, r.H)
	r.ProblemDesc = NewContestControllerIdGroupIdGroupProblemDescRouter(traits, r.H)

	r.GetContestProblem = r.GetRouter().GET("", traits.GetServiceInstance("ContestController").(ContestController).GetContestProblem)
	r.PutContestProblem = r.GetRouter().PUT("", traits.GetServiceInstance("ContestController").(ContestController).PutContestProblem)
	r.DeleteContestProblem = r.GetRouter().DELETE("", traits.GetServiceInstance("ContestController").(ContestController).DeleteContestProblem)

	return
}

type ContestControllerIdGroupIdGroupListContestProblemDescRouter struct {
	H

	ListContestProblemDesc *LeafRouter
}

func NewContestControllerIdGroupIdGroupListContestProblemDescRouter(traits GenerateRouterTraits, h H) (r *ContestControllerIdGroupIdGroupListContestProblemDescRouter) {
	r = &ContestControllerIdGroupIdGroupListContestProblemDescRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("/desc-list"),
			AuthRouter: h.GetAuthRouter().Group("/desc-list"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), ""),
		},
	}

	r.ListContestProblemDesc = r.GetRouter().GET("", traits.GetServiceInstance("ContestController").(ContestController).ListContestProblemDesc)

	return
}

type ContestControllerIdGroupIdGroupCountContestProblemDescRouter struct {
	H

	CountContestProblemDesc *LeafRouter
}

func NewContestControllerIdGroupIdGroupCountContestProblemDescRouter(traits GenerateRouterTraits, h H) (r *ContestControllerIdGroupIdGroupCountContestProblemDescRouter) {
	r = &ContestControllerIdGroupIdGroupCountContestProblemDescRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("/desc-count"),
			AuthRouter: h.GetAuthRouter().Group("/desc-count"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), ""),
		},
	}

	r.CountContestProblemDesc = r.GetRouter().GET("", traits.GetServiceInstance("ContestController").(ContestController).CountContestProblemDesc)

	return
}

type ContestControllerIdGroupIdGroupProblemDescRouter struct {
	H
	ProblemDesc *ContestControllerIdGroupIdGroupProblemDescProblemDescRouter

	PostContestProblemDesc   *LeafRouter
	GetContestProblemDesc    *LeafRouter
	PutContestProblemDesc    *LeafRouter
	DeleteContestProblemDesc *LeafRouter
}

func NewContestControllerIdGroupIdGroupProblemDescRouter(traits GenerateRouterTraits, h H) (r *ContestControllerIdGroupIdGroupProblemDescRouter) {
	r = &ContestControllerIdGroupIdGroupProblemDescRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("/desc"),
			AuthRouter: h.GetAuthRouter().Group("/desc"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), ""),
		},
	}

	r.ProblemDesc = NewContestControllerIdGroupIdGroupProblemDescProblemDescRouter(traits, r.H)

	r.PostContestProblemDesc = r.GetRouter().POST("", traits.GetServiceInstance("ContestController").(ContestController).PostContestProblemDesc)
	r.GetContestProblemDesc = r.GetRouter().GET("", traits.GetServiceInstance("ContestController").(ContestController).GetContestProblemDesc)
	r.PutContestProblemDesc = r.GetRouter().PUT("", traits.GetServiceInstance("ContestController").(ContestController).PutContestProblemDesc)
	r.DeleteContestProblemDesc = r.GetRouter().DELETE("", traits.GetServiceInstance("ContestController").(ContestController).DeleteContestProblemDesc)

	return
}

type ContestControllerIdGroupIdGroupProblemDescProblemDescRouter struct {
	H

	ChangeContestProblemDescriptionRef *LeafRouter
}

func NewContestControllerIdGroupIdGroupProblemDescProblemDescRouter(traits GenerateRouterTraits, h H) (r *ContestControllerIdGroupIdGroupProblemDescProblemDescRouter) {
	r = &ContestControllerIdGroupIdGroupProblemDescProblemDescRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("/ref"),
			AuthRouter: h.GetAuthRouter().Group("/ref"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), ""),
		},
	}

	r.ChangeContestProblemDescriptionRef = r.GetRouter().POST("", traits.GetServiceInstance("ContestController").(ContestController).ChangeContestProblemDescriptionRef)

	return
}

type GroupControllerRouter struct {
	H
	List    *GroupControllerListRouter
	Count   *GroupControllerCountRouter
	Post    *GroupControllerPostRouter
	IdGroup *GroupControllerIdGroupRouter
}

func NewGroupControllerRouter(traits GenerateRouterTraits, h H) (r *GroupControllerRouter) {
	r = &GroupControllerRouter{
		H: &BaseH{
			Router:     h.GetRouter().Extend("GroupController"),
			AuthRouter: h.GetAuthRouter().Extend("GroupController"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), ""),
		},
	}

	r.List = NewGroupControllerListRouter(traits, r.H)
	r.Count = NewGroupControllerCountRouter(traits, r.H)
	r.Post = NewGroupControllerPostRouter(traits, r.H)
	r.IdGroup = NewGroupControllerIdGroupRouter(traits, r.H)

	return
}

type GroupControllerListRouter struct {
	H

	ListGroup *LeafRouter
}

func NewGroupControllerListRouter(traits GenerateRouterTraits, h H) (r *GroupControllerListRouter) {
	r = &GroupControllerListRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("group-list"),
			AuthRouter: h.GetAuthRouter().Group("group-list"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), ""),
		},
	}

	r.ListGroup = r.GetRouter().GET("", traits.GetServiceInstance("GroupController").(GroupController).ListGroup)

	return
}

type GroupControllerCountRouter struct {
	H

	CountGroup *LeafRouter
}

func NewGroupControllerCountRouter(traits GenerateRouterTraits, h H) (r *GroupControllerCountRouter) {
	r = &GroupControllerCountRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("group-count"),
			AuthRouter: h.GetAuthRouter().Group("group-count"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), ""),
		},
	}

	r.CountGroup = r.GetRouter().GET("", traits.GetServiceInstance("GroupController").(GroupController).CountGroup)

	return
}

type GroupControllerPostRouter struct {
	H

	PostGroup *LeafRouter
}

func NewGroupControllerPostRouter(traits GenerateRouterTraits, h H) (r *GroupControllerPostRouter) {
	r = &GroupControllerPostRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("group"),
			AuthRouter: h.GetAuthRouter().Group("group"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), ""),
		},
	}

	r.PostGroup = r.GetAuthRouter().POST("", traits.GetServiceInstance("GroupController").(GroupController).PostGroup)
	r.PostGroup = traits.ApplyAuthOnMethod(r.PostGroup, "~")

	return
}

type GroupControllerIdGroupRouter struct {
	H
	Owner    *GroupControllerIdGroupOwnerRouter
	UserList *GroupControllerIdGroupUserListRouter
	User     *GroupControllerIdGroupUserRouter

	GetGroup    *LeafRouter
	PutGroup    *LeafRouter
	DeleteGroup *LeafRouter
}

func NewGroupControllerIdGroupRouter(traits GenerateRouterTraits, h H) (r *GroupControllerIdGroupRouter) {
	r = &GroupControllerIdGroupRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("group/:gid"),
			AuthRouter: h.GetAuthRouter().Group("group/:gid"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), "group:gid"),
		},
	}

	r.Owner = NewGroupControllerIdGroupOwnerRouter(traits, r.H)
	r.UserList = NewGroupControllerIdGroupUserListRouter(traits, r.H)
	r.User = NewGroupControllerIdGroupUserRouter(traits, r.H)

	r.GetGroup = r.GetRouter().GET("", traits.GetServiceInstance("GroupController").(GroupController).GetGroup)
	r.PutGroup = r.GetRouter().PUT("", traits.GetServiceInstance("GroupController").(GroupController).PutGroup)
	r.DeleteGroup = r.GetRouter().DELETE("", traits.GetServiceInstance("GroupController").(GroupController).DeleteGroup)

	return
}

type GroupControllerIdGroupOwnerRouter struct {
	H

	PutGroupOwner *LeafRouter
}

func NewGroupControllerIdGroupOwnerRouter(traits GenerateRouterTraits, h H) (r *GroupControllerIdGroupOwnerRouter) {
	r = &GroupControllerIdGroupOwnerRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("/owner"),
			AuthRouter: h.GetAuthRouter().Group("/owner"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), ""),
		},
	}

	r.PutGroupOwner = r.GetRouter().PUT("", traits.GetServiceInstance("GroupController").(GroupController).PutGroupOwner)

	return
}

type GroupControllerIdGroupUserListRouter struct {
	H

	GetGroupMembers *LeafRouter
}

func NewGroupControllerIdGroupUserListRouter(traits GenerateRouterTraits, h H) (r *GroupControllerIdGroupUserListRouter) {
	r = &GroupControllerIdGroupUserListRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("/user-list"),
			AuthRouter: h.GetAuthRouter().Group("/user-list"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), ""),
		},
	}

	r.GetGroupMembers = r.GetRouter().GET("", traits.GetServiceInstance("GroupController").(GroupController).GetGroupMembers)

	return
}

type GroupControllerIdGroupUserRouter struct {
	H

	PostGroupMember *LeafRouter
}

func NewGroupControllerIdGroupUserRouter(traits GenerateRouterTraits, h H) (r *GroupControllerIdGroupUserRouter) {
	r = &GroupControllerIdGroupUserRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("user/:id"),
			AuthRouter: h.GetAuthRouter().Group("user/:id"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), "user:id"),
		},
	}

	r.PostGroupMember = r.GetRouter().POST("", traits.GetServiceInstance("GroupController").(GroupController).PostGroupMember)

	return
}

type UserControllerRouter struct {
	H
	List         *UserControllerListRouter
	Count        *UserControllerCountRouter
	Register     *UserControllerRegisterRouter
	Login        *UserControllerLoginRouter
	RefreshToken *UserControllerRefreshTokenRouter
	IdGroup      *UserControllerIdGroupRouter
}

func NewUserControllerRouter(traits GenerateRouterTraits, h H) (r *UserControllerRouter) {
	r = &UserControllerRouter{
		H: &BaseH{
			Router:     h.GetRouter().Extend("UserController"),
			AuthRouter: h.GetAuthRouter().Extend("UserController"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), ""),
		},
	}

	r.List = NewUserControllerListRouter(traits, r.H)
	r.Count = NewUserControllerCountRouter(traits, r.H)
	r.Register = NewUserControllerRegisterRouter(traits, r.H)
	r.Login = NewUserControllerLoginRouter(traits, r.H)
	r.RefreshToken = NewUserControllerRefreshTokenRouter(traits, r.H)
	r.IdGroup = NewUserControllerIdGroupRouter(traits, r.H)

	return
}

type UserControllerListRouter struct {
	H

	ListUser *LeafRouter
}

func NewUserControllerListRouter(traits GenerateRouterTraits, h H) (r *UserControllerListRouter) {
	r = &UserControllerListRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("user-list"),
			AuthRouter: h.GetAuthRouter().Group("user-list"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), ""),
		},
	}

	r.ListUser = r.GetRouter().GET("", traits.GetServiceInstance("UserController").(UserController).ListUser)

	return
}

type UserControllerCountRouter struct {
	H

	CountUser *LeafRouter
}

func NewUserControllerCountRouter(traits GenerateRouterTraits, h H) (r *UserControllerCountRouter) {
	r = &UserControllerCountRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("user-count"),
			AuthRouter: h.GetAuthRouter().Group("user-count"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), ""),
		},
	}

	r.CountUser = r.GetRouter().GET("", traits.GetServiceInstance("UserController").(UserController).CountUser)

	return
}

type UserControllerRegisterRouter struct {
	H

	Register *LeafRouter
}

func NewUserControllerRegisterRouter(traits GenerateRouterTraits, h H) (r *UserControllerRegisterRouter) {
	r = &UserControllerRegisterRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("user/register"),
			AuthRouter: h.GetAuthRouter().Group("user/register"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), ""),
		},
	}

	r.Register = r.GetRouter().POST("", traits.GetServiceInstance("UserController").(UserController).Register)

	return
}

type UserControllerLoginRouter struct {
	H

	LoginUser *LeafRouter
}

func NewUserControllerLoginRouter(traits GenerateRouterTraits, h H) (r *UserControllerLoginRouter) {
	r = &UserControllerLoginRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("user/login"),
			AuthRouter: h.GetAuthRouter().Group("user/login"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), ""),
		},
	}

	r.LoginUser = r.GetRouter().POST("", traits.GetServiceInstance("UserController").(UserController).LoginUser)

	return
}

type UserControllerRefreshTokenRouter struct {
	H

	RefreshToken *LeafRouter
}

func NewUserControllerRefreshTokenRouter(traits GenerateRouterTraits, h H) (r *UserControllerRefreshTokenRouter) {
	r = &UserControllerRefreshTokenRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("user-token"),
			AuthRouter: h.GetAuthRouter().Group("user-token"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), ""),
		},
	}

	r.RefreshToken = r.GetRouter().GET("", traits.GetServiceInstance("UserController").(UserController).RefreshToken)

	return
}

type UserControllerIdGroupRouter struct {
	H
	Email          *UserControllerIdGroupEmailRouter
	ChangePassword *UserControllerIdGroupChangePasswordRouter
	Inspect        *UserControllerIdGroupInspectRouter

	GetUser    *LeafRouter
	PutUser    *LeafRouter
	DeleteUser *LeafRouter
}

func NewUserControllerIdGroupRouter(traits GenerateRouterTraits, h H) (r *UserControllerIdGroupRouter) {
	r = &UserControllerIdGroupRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("user/:id"),
			AuthRouter: h.GetAuthRouter().Group("user/:id"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), "user:id"),
		},
	}

	r.Email = NewUserControllerIdGroupEmailRouter(traits, r.H)
	r.ChangePassword = NewUserControllerIdGroupChangePasswordRouter(traits, r.H)
	r.Inspect = NewUserControllerIdGroupInspectRouter(traits, r.H)

	r.GetUser = r.GetRouter().GET("", traits.GetServiceInstance("UserController").(UserController).GetUser)
	r.PutUser = r.GetRouter().PUT("", traits.GetServiceInstance("UserController").(UserController).PutUser)
	r.DeleteUser = r.GetRouter().DELETE("", traits.GetServiceInstance("UserController").(UserController).DeleteUser)

	return
}

type UserControllerIdGroupEmailRouter struct {
	H

	BindEmail *LeafRouter
}

func NewUserControllerIdGroupEmailRouter(traits GenerateRouterTraits, h H) (r *UserControllerIdGroupEmailRouter) {
	r = &UserControllerIdGroupEmailRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("/email"),
			AuthRouter: h.GetAuthRouter().Group("/email"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), ""),
		},
	}

	r.BindEmail = r.GetRouter().PUT("", traits.GetServiceInstance("UserController").(UserController).BindEmail)

	return
}

type UserControllerIdGroupChangePasswordRouter struct {
	H

	ChangePassword *LeafRouter
}

func NewUserControllerIdGroupChangePasswordRouter(traits GenerateRouterTraits, h H) (r *UserControllerIdGroupChangePasswordRouter) {
	r = &UserControllerIdGroupChangePasswordRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("/password"),
			AuthRouter: h.GetAuthRouter().Group("/password"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), ""),
		},
	}

	r.ChangePassword = r.GetRouter().PUT("", traits.GetServiceInstance("UserController").(UserController).ChangePassword)

	return
}

type UserControllerIdGroupInspectRouter struct {
	H

	InspectUser *LeafRouter
}

func NewUserControllerIdGroupInspectRouter(traits GenerateRouterTraits, h H) (r *UserControllerIdGroupInspectRouter) {
	r = &UserControllerIdGroupInspectRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("/inspect"),
			AuthRouter: h.GetAuthRouter().Group("/inspect"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), ""),
		},
	}

	r.InspectUser = r.GetRouter().GET("", traits.GetServiceInstance("UserController").(UserController).InspectUser)

	return
}

type AuthControllerRouter struct {
	H
	Policy         *AuthControllerPolicyRouter
	GroupingPolicy *AuthControllerGroupingPolicyRouter
}

func NewAuthControllerRouter(traits GenerateRouterTraits, h H) (r *AuthControllerRouter) {
	r = &AuthControllerRouter{
		H: &BaseH{
			Router:     h.GetRouter().Extend("AuthController"),
			AuthRouter: h.GetAuthRouter().Extend("AuthController"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), ""),
		},
	}

	r.Policy = NewAuthControllerPolicyRouter(traits, r.H)
	r.GroupingPolicy = NewAuthControllerGroupingPolicyRouter(traits, r.H)

	return
}

type AuthControllerPolicyRouter struct {
	H

	AddPolicy    *LeafRouter
	RemovePolicy *LeafRouter
	HasPolicy    *LeafRouter
}

func NewAuthControllerPolicyRouter(traits GenerateRouterTraits, h H) (r *AuthControllerPolicyRouter) {
	r = &AuthControllerPolicyRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("/policy"),
			AuthRouter: h.GetAuthRouter().Group("/policy"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), ""),
		},
	}

	r.AddPolicy = r.GetRouter().POST("", traits.GetServiceInstance("AuthController").(AuthController).AddPolicy)
	r.RemovePolicy = r.GetRouter().DELETE("", traits.GetServiceInstance("AuthController").(AuthController).RemovePolicy)
	r.HasPolicy = r.GetRouter().GET("", traits.GetServiceInstance("AuthController").(AuthController).HasPolicy)

	return
}

type AuthControllerGroupingPolicyRouter struct {
	H

	AddGroupingPolicy    *LeafRouter
	RemoveGroupingPolicy *LeafRouter
	HasGroupingPolicy    *LeafRouter
}

func NewAuthControllerGroupingPolicyRouter(traits GenerateRouterTraits, h H) (r *AuthControllerGroupingPolicyRouter) {
	r = &AuthControllerGroupingPolicyRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("/policy/group"),
			AuthRouter: h.GetAuthRouter().Group("/policy/group"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), ""),
		},
	}

	r.AddGroupingPolicy = r.GetRouter().POST("", traits.GetServiceInstance("AuthController").(AuthController).AddGroupingPolicy)
	r.RemoveGroupingPolicy = r.GetRouter().DELETE("", traits.GetServiceInstance("AuthController").(AuthController).RemoveGroupingPolicy)
	r.HasGroupingPolicy = r.GetRouter().GET("", traits.GetServiceInstance("AuthController").(AuthController).HasGroupingPolicy)

	return
}

type AnnouncementControllerRouter struct {
	H
	List    *AnnouncementControllerListRouter
	Count   *AnnouncementControllerCountRouter
	Post    *AnnouncementControllerPostRouter
	IdGroup *AnnouncementControllerIdGroupRouter
}

func NewAnnouncementControllerRouter(traits GenerateRouterTraits, h H) (r *AnnouncementControllerRouter) {
	r = &AnnouncementControllerRouter{
		H: &BaseH{
			Router:     h.GetRouter().Extend("AnnouncementController"),
			AuthRouter: h.GetAuthRouter().Extend("AnnouncementController"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), ""),
		},
	}

	r.List = NewAnnouncementControllerListRouter(traits, r.H)
	r.Count = NewAnnouncementControllerCountRouter(traits, r.H)
	r.Post = NewAnnouncementControllerPostRouter(traits, r.H)
	r.IdGroup = NewAnnouncementControllerIdGroupRouter(traits, r.H)

	return
}

type AnnouncementControllerListRouter struct {
	H

	ListAnnouncement *LeafRouter
}

func NewAnnouncementControllerListRouter(traits GenerateRouterTraits, h H) (r *AnnouncementControllerListRouter) {
	r = &AnnouncementControllerListRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("announcement-list"),
			AuthRouter: h.GetAuthRouter().Group("announcement-list"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), ""),
		},
	}

	r.ListAnnouncement = r.GetRouter().GET("", traits.GetServiceInstance("AnnouncementController").(AnnouncementController).ListAnnouncement)

	return
}

type AnnouncementControllerCountRouter struct {
	H

	CountAnnouncement *LeafRouter
}

func NewAnnouncementControllerCountRouter(traits GenerateRouterTraits, h H) (r *AnnouncementControllerCountRouter) {
	r = &AnnouncementControllerCountRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("announcement-count"),
			AuthRouter: h.GetAuthRouter().Group("announcement-count"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), ""),
		},
	}

	r.CountAnnouncement = r.GetRouter().GET("", traits.GetServiceInstance("AnnouncementController").(AnnouncementController).CountAnnouncement)

	return
}

type AnnouncementControllerPostRouter struct {
	H

	PostAnnouncement *LeafRouter
}

func NewAnnouncementControllerPostRouter(traits GenerateRouterTraits, h H) (r *AnnouncementControllerPostRouter) {
	r = &AnnouncementControllerPostRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("announcement"),
			AuthRouter: h.GetAuthRouter().Group("announcement"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), ""),
		},
	}

	r.PostAnnouncement = r.GetAuthRouter().POST("", traits.GetServiceInstance("AnnouncementController").(AnnouncementController).PostAnnouncement)
	r.PostAnnouncement = traits.ApplyAuthOnMethod(r.PostAnnouncement, "~")

	return
}

type AnnouncementControllerIdGroupRouter struct {
	H

	GetAnnouncement    *LeafRouter
	PutAnnouncement    *LeafRouter
	DeleteAnnouncement *LeafRouter
}

func NewAnnouncementControllerIdGroupRouter(traits GenerateRouterTraits, h H) (r *AnnouncementControllerIdGroupRouter) {
	r = &AnnouncementControllerIdGroupRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("announcement/:aid"),
			AuthRouter: h.GetAuthRouter().Group("announcement/:aid"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), "announcement:aid"),
		},
	}

	r.GetAnnouncement = r.GetRouter().GET("", traits.GetServiceInstance("AnnouncementController").(AnnouncementController).GetAnnouncement)
	r.PutAnnouncement = r.GetAuthRouter().PUT("", traits.GetServiceInstance("AnnouncementController").(AnnouncementController).PutAnnouncement)
	r.PutAnnouncement = traits.ApplyAuthOnMethod(r.PutAnnouncement, "~")
	r.DeleteAnnouncement = r.GetAuthRouter().DELETE("", traits.GetServiceInstance("AnnouncementController").(AnnouncementController).DeleteAnnouncement)
	r.DeleteAnnouncement = traits.ApplyAuthOnMethod(r.DeleteAnnouncement, "~")

	return
}

type CommentControllerRouter struct {
	H
	List    *CommentControllerListRouter
	Count   *CommentControllerCountRouter
	Post    *CommentControllerPostRouter
	IdGroup *CommentControllerIdGroupRouter
}

func NewCommentControllerRouter(traits GenerateRouterTraits, h H) (r *CommentControllerRouter) {
	r = &CommentControllerRouter{
		H: &BaseH{
			Router:     h.GetRouter().Extend("CommentController"),
			AuthRouter: h.GetAuthRouter().Extend("CommentController"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), ""),
		},
	}

	r.List = NewCommentControllerListRouter(traits, r.H)
	r.Count = NewCommentControllerCountRouter(traits, r.H)
	r.Post = NewCommentControllerPostRouter(traits, r.H)
	r.IdGroup = NewCommentControllerIdGroupRouter(traits, r.H)

	return
}

type CommentControllerListRouter struct {
	H

	ListComment *LeafRouter
}

func NewCommentControllerListRouter(traits GenerateRouterTraits, h H) (r *CommentControllerListRouter) {
	r = &CommentControllerListRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("comment-list"),
			AuthRouter: h.GetAuthRouter().Group("comment-list"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), ""),
		},
	}

	r.ListComment = r.GetRouter().GET("", traits.GetServiceInstance("CommentController").(CommentController).ListComment)

	return
}

type CommentControllerCountRouter struct {
	H

	CountComment *LeafRouter
}

func NewCommentControllerCountRouter(traits GenerateRouterTraits, h H) (r *CommentControllerCountRouter) {
	r = &CommentControllerCountRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("comment-count"),
			AuthRouter: h.GetAuthRouter().Group("comment-count"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), ""),
		},
	}

	r.CountComment = r.GetRouter().GET("", traits.GetServiceInstance("CommentController").(CommentController).CountComment)

	return
}

type CommentControllerPostRouter struct {
	H

	PostComment *LeafRouter
}

func NewCommentControllerPostRouter(traits GenerateRouterTraits, h H) (r *CommentControllerPostRouter) {
	r = &CommentControllerPostRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("comment"),
			AuthRouter: h.GetAuthRouter().Group("comment"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), ""),
		},
	}

	r.PostComment = r.GetRouter().POST("", traits.GetServiceInstance("CommentController").(CommentController).PostComment)

	return
}

type CommentControllerIdGroupRouter struct {
	H

	GetComment    *LeafRouter
	PutComment    *LeafRouter
	DeleteComment *LeafRouter
}

func NewCommentControllerIdGroupRouter(traits GenerateRouterTraits, h H) (r *CommentControllerIdGroupRouter) {
	r = &CommentControllerIdGroupRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("comment/:cmid"),
			AuthRouter: h.GetAuthRouter().Group("comment/:cmid"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), "comment:cmid"),
		},
	}

	r.GetComment = r.GetRouter().GET("", traits.GetServiceInstance("CommentController").(CommentController).GetComment)
	r.PutComment = r.GetRouter().PUT("", traits.GetServiceInstance("CommentController").(CommentController).PutComment)
	r.DeleteComment = r.GetRouter().DELETE("", traits.GetServiceInstance("CommentController").(CommentController).DeleteComment)

	return
}

type SubmissionControllerRouter struct {
	H
	List    *SubmissionControllerListRouter
	Count   *SubmissionControllerCountRouter
	Post    *SubmissionControllerPostRouter
	IdGroup *SubmissionControllerIdGroupRouter
}

func NewSubmissionControllerRouter(traits GenerateRouterTraits, h H) (r *SubmissionControllerRouter) {
	r = &SubmissionControllerRouter{
		H: &BaseH{
			Router:     h.GetRouter().Extend("SubmissionController"),
			AuthRouter: h.GetAuthRouter().Extend("SubmissionController"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), ""),
		},
	}

	r.List = NewSubmissionControllerListRouter(traits, r.H)
	r.Count = NewSubmissionControllerCountRouter(traits, r.H)
	r.Post = NewSubmissionControllerPostRouter(traits, r.H)
	r.IdGroup = NewSubmissionControllerIdGroupRouter(traits, r.H)

	return
}

type SubmissionControllerListRouter struct {
	H

	ListSubmission *LeafRouter
}

func NewSubmissionControllerListRouter(traits GenerateRouterTraits, h H) (r *SubmissionControllerListRouter) {
	r = &SubmissionControllerListRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("submission-list"),
			AuthRouter: h.GetAuthRouter().Group("submission-list"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), ""),
		},
	}

	r.ListSubmission = r.GetRouter().GET("", traits.GetServiceInstance("SubmissionController").(SubmissionController).ListSubmission)

	return
}

type SubmissionControllerCountRouter struct {
	H

	CountSubmission *LeafRouter
}

func NewSubmissionControllerCountRouter(traits GenerateRouterTraits, h H) (r *SubmissionControllerCountRouter) {
	r = &SubmissionControllerCountRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("submission-count"),
			AuthRouter: h.GetAuthRouter().Group("submission-count"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), ""),
		},
	}

	r.CountSubmission = r.GetRouter().GET("", traits.GetServiceInstance("SubmissionController").(SubmissionController).CountSubmission)

	return
}

type SubmissionControllerPostRouter struct {
	H

	PostSubmission *LeafRouter
}

func NewSubmissionControllerPostRouter(traits GenerateRouterTraits, h H) (r *SubmissionControllerPostRouter) {
	r = &SubmissionControllerPostRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("/problem/:pid/submission"),
			AuthRouter: h.GetAuthRouter().Group("/problem/:pid/submission"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), "problem:pid"),
		},
	}

	r.PostSubmission = r.GetAuthRouter().POST("", traits.GetServiceInstance("SubmissionController").(SubmissionController).PostSubmission)
	r.PostSubmission = traits.ApplyAuthOnMethod(r.PostSubmission, "~")

	return
}

type SubmissionControllerIdGroupRouter struct {
	H
	GetSubmissionContent *SubmissionControllerIdGroupGetSubmissionContentRouter

	GetSubmission    *LeafRouter
	DeleteSubmission *LeafRouter
}

func NewSubmissionControllerIdGroupRouter(traits GenerateRouterTraits, h H) (r *SubmissionControllerIdGroupRouter) {
	r = &SubmissionControllerIdGroupRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("submission/:sid"),
			AuthRouter: h.GetAuthRouter().Group("submission/:sid"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), "submission:sid"),
		},
	}

	r.GetSubmissionContent = NewSubmissionControllerIdGroupGetSubmissionContentRouter(traits, r.H)

	r.GetSubmission = r.GetRouter().GET("", traits.GetServiceInstance("SubmissionController").(SubmissionController).GetSubmission)
	r.DeleteSubmission = r.GetRouter().DELETE("", traits.GetServiceInstance("SubmissionController").(SubmissionController).DeleteSubmission)

	return
}

type SubmissionControllerIdGroupGetSubmissionContentRouter struct {
	H

	GetSubmissionContent *LeafRouter
}

func NewSubmissionControllerIdGroupGetSubmissionContentRouter(traits GenerateRouterTraits, h H) (r *SubmissionControllerIdGroupGetSubmissionContentRouter) {
	r = &SubmissionControllerIdGroupGetSubmissionContentRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("/content"),
			AuthRouter: h.GetAuthRouter().Group("/content"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), ""),
		},
	}

	r.GetSubmissionContent = r.GetRouter().GET("", traits.GetServiceInstance("SubmissionController").(SubmissionController).GetSubmissionContent)

	return
}
