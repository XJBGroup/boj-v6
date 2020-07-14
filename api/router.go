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
	CommentService      *CommentServiceRouter
	SubmissionService   *SubmissionServiceRouter
	ProblemService      *ProblemServiceRouter
	ContestService      *ContestServiceRouter
	GroupService        *GroupServiceRouter
	UserService         *UserServiceRouter
	AuthService         *AuthServiceRouter
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

	r.CommentService = NewCommentServiceRouter(traits, r.H)
	r.SubmissionService = NewSubmissionServiceRouter(traits, r.H)
	r.ProblemService = NewProblemServiceRouter(traits, r.H)
	r.ContestService = NewContestServiceRouter(traits, r.H)
	r.GroupService = NewGroupServiceRouter(traits, r.H)
	r.UserService = NewUserServiceRouter(traits, r.H)
	r.AuthService = NewAuthServiceRouter(traits, r.H)
	r.AnnouncementService = NewAnnouncementServiceRouter(traits, r.H)

	traits.AfterBuild(r)
	traits.ApplyAuth(r)
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
	GetContent *SubmissionServiceIdGroupGetContentRouter

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

	r.GetContent = NewSubmissionServiceIdGroupGetContentRouter(traits, r.H)

	r.GetSubmission = r.GetRouter().GET("", traits.GetServiceInstance("SubmissionService").(SubmissionService).GetSubmission)
	r.Delete = r.GetRouter().DELETE("", traits.GetServiceInstance("SubmissionService").(SubmissionService).Delete)

	return
}

type SubmissionServiceIdGroupGetContentRouter struct {
	H

	GetContent *LeafRouter
}

func NewSubmissionServiceIdGroupGetContentRouter(traits GenerateRouterTraits, h H) (r *SubmissionServiceIdGroupGetContentRouter) {
	r = &SubmissionServiceIdGroupGetContentRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("/content"),
			AuthRouter: h.GetAuthRouter().Group("/content"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), ""),
		},
	}

	r.GetContent = r.GetRouter().GET("", traits.GetServiceInstance("SubmissionService").(SubmissionService).GetContent)

	return
}

type ProblemServiceRouter struct {
	H
	List    *ProblemServiceListRouter
	Count   *ProblemServiceCountRouter
	Post    *ProblemServicePostRouter
	IdGroup *ProblemServiceIdGroupRouter
}

func NewProblemServiceRouter(traits GenerateRouterTraits, h H) (r *ProblemServiceRouter) {
	r = &ProblemServiceRouter{
		H: &BaseH{
			Router:     h.GetRouter().Extend("ProblemService"),
			AuthRouter: h.GetAuthRouter().Extend("ProblemService"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), ""),
		},
	}

	r.List = NewProblemServiceListRouter(traits, r.H)
	r.Count = NewProblemServiceCountRouter(traits, r.H)
	r.Post = NewProblemServicePostRouter(traits, r.H)
	r.IdGroup = NewProblemServiceIdGroupRouter(traits, r.H)

	return
}

type ProblemServiceListRouter struct {
	H

	ListProblems *LeafRouter
}

func NewProblemServiceListRouter(traits GenerateRouterTraits, h H) (r *ProblemServiceListRouter) {
	r = &ProblemServiceListRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("problem-list"),
			AuthRouter: h.GetAuthRouter().Group("problem-list"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), ""),
		},
	}

	r.ListProblems = r.GetRouter().GET("", traits.GetServiceInstance("ProblemService").(ProblemService).ListProblems)

	return
}

type ProblemServiceCountRouter struct {
	H

	CountProblem *LeafRouter
}

func NewProblemServiceCountRouter(traits GenerateRouterTraits, h H) (r *ProblemServiceCountRouter) {
	r = &ProblemServiceCountRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("problem-count"),
			AuthRouter: h.GetAuthRouter().Group("problem-count"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), ""),
		},
	}

	r.CountProblem = r.GetRouter().GET("", traits.GetServiceInstance("ProblemService").(ProblemService).CountProblem)

	return
}

type ProblemServicePostRouter struct {
	H

	PostProblem *LeafRouter
}

func NewProblemServicePostRouter(traits GenerateRouterTraits, h H) (r *ProblemServicePostRouter) {
	r = &ProblemServicePostRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("problem"),
			AuthRouter: h.GetAuthRouter().Group("problem"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), ""),
		},
	}

	r.PostProblem = r.GetAuthRouter().POST("", traits.GetServiceInstance("ProblemService").(ProblemService).PostProblem)
	r.PostProblem = traits.ApplyAuthOnMethod(r.PostProblem, "~")

	return
}

type ProblemServiceIdGroupRouter struct {
	H
	ProblemDesc *ProblemServiceIdGroupProblemDescRouter

	GetProblem    *LeafRouter
	PutProblem    *LeafRouter
	DeleteProblem *LeafRouter
}

func NewProblemServiceIdGroupRouter(traits GenerateRouterTraits, h H) (r *ProblemServiceIdGroupRouter) {
	r = &ProblemServiceIdGroupRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("problem/:pid"),
			AuthRouter: h.GetAuthRouter().Group("problem/:pid"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), "problem:pid"),
		},
	}

	r.ProblemDesc = NewProblemServiceIdGroupProblemDescRouter(traits, r.H)

	r.GetProblem = r.GetRouter().GET("", traits.GetServiceInstance("ProblemService").(ProblemService).GetProblem)
	r.PutProblem = r.GetRouter().PUT("", traits.GetServiceInstance("ProblemService").(ProblemService).PutProblem)
	r.DeleteProblem = r.GetRouter().DELETE("", traits.GetServiceInstance("ProblemService").(ProblemService).DeleteProblem)

	return
}

type ProblemServiceIdGroupProblemDescRouter struct {
	H
	ProblemDesc *ProblemServiceIdGroupProblemDescProblemDescRouter

	PostProblemDesc   *LeafRouter
	GetProblemDesc    *LeafRouter
	PutProblemDesc    *LeafRouter
	DeleteProblemDesc *LeafRouter
}

func NewProblemServiceIdGroupProblemDescRouter(traits GenerateRouterTraits, h H) (r *ProblemServiceIdGroupProblemDescRouter) {
	r = &ProblemServiceIdGroupProblemDescRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("/desc"),
			AuthRouter: h.GetAuthRouter().Group("/desc"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), ""),
		},
	}

	r.ProblemDesc = NewProblemServiceIdGroupProblemDescProblemDescRouter(traits, r.H)

	r.PostProblemDesc = r.GetRouter().POST("", traits.GetServiceInstance("ProblemService").(ProblemService).PostProblemDesc)
	r.GetProblemDesc = r.GetRouter().GET("", traits.GetServiceInstance("ProblemService").(ProblemService).GetProblemDesc)
	r.PutProblemDesc = r.GetRouter().PUT("", traits.GetServiceInstance("ProblemService").(ProblemService).PutProblemDesc)
	r.DeleteProblemDesc = r.GetRouter().DELETE("", traits.GetServiceInstance("ProblemService").(ProblemService).DeleteProblemDesc)

	return
}

type ProblemServiceIdGroupProblemDescProblemDescRouter struct {
	H

	ChangeProblemDescriptionRef *LeafRouter
}

func NewProblemServiceIdGroupProblemDescProblemDescRouter(traits GenerateRouterTraits, h H) (r *ProblemServiceIdGroupProblemDescProblemDescRouter) {
	r = &ProblemServiceIdGroupProblemDescProblemDescRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("/desc"),
			AuthRouter: h.GetAuthRouter().Group("/desc"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), ""),
		},
	}

	r.ChangeProblemDescriptionRef = r.GetRouter().POST("", traits.GetServiceInstance("ProblemService").(ProblemService).ChangeProblemDescriptionRef)

	return
}

type ContestServiceRouter struct {
	H
	List    *ContestServiceListRouter
	Count   *ContestServiceCountRouter
	Post    *ContestServicePostRouter
	IdGroup *ContestServiceIdGroupRouter
}

func NewContestServiceRouter(traits GenerateRouterTraits, h H) (r *ContestServiceRouter) {
	r = &ContestServiceRouter{
		H: &BaseH{
			Router:     h.GetRouter().Extend("ContestService"),
			AuthRouter: h.GetAuthRouter().Extend("ContestService"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), ""),
		},
	}

	r.List = NewContestServiceListRouter(traits, r.H)
	r.Count = NewContestServiceCountRouter(traits, r.H)
	r.Post = NewContestServicePostRouter(traits, r.H)
	r.IdGroup = NewContestServiceIdGroupRouter(traits, r.H)

	return
}

type ContestServiceListRouter struct {
	H

	ListContests *LeafRouter
}

func NewContestServiceListRouter(traits GenerateRouterTraits, h H) (r *ContestServiceListRouter) {
	r = &ContestServiceListRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("contest-list"),
			AuthRouter: h.GetAuthRouter().Group("contest-list"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), ""),
		},
	}

	r.ListContests = r.GetRouter().GET("", traits.GetServiceInstance("ContestService").(ContestService).ListContests)

	return
}

type ContestServiceCountRouter struct {
	H

	CountContest *LeafRouter
}

func NewContestServiceCountRouter(traits GenerateRouterTraits, h H) (r *ContestServiceCountRouter) {
	r = &ContestServiceCountRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("contest-count"),
			AuthRouter: h.GetAuthRouter().Group("contest-count"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), ""),
		},
	}

	r.CountContest = r.GetRouter().GET("", traits.GetServiceInstance("ContestService").(ContestService).CountContest)

	return
}

type ContestServicePostRouter struct {
	H

	PostContest *LeafRouter
}

func NewContestServicePostRouter(traits GenerateRouterTraits, h H) (r *ContestServicePostRouter) {
	r = &ContestServicePostRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("contest"),
			AuthRouter: h.GetAuthRouter().Group("contest"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), ""),
		},
	}

	r.PostContest = r.GetAuthRouter().POST("", traits.GetServiceInstance("ContestService").(ContestService).PostContest)
	r.PostContest = traits.ApplyAuthOnMethod(r.PostContest, "~")

	return
}

type ContestServiceIdGroupRouter struct {
	H
	Count   *ContestServiceIdGroupCountRouter
	Post    *ContestServiceIdGroupPostRouter
	IdGroup *ContestServiceIdGroupIdGroupRouter
	List    *ContestServiceIdGroupListRouter

	GetContest *LeafRouter
	PutContest *LeafRouter
	Delete     *LeafRouter
}

func NewContestServiceIdGroupRouter(traits GenerateRouterTraits, h H) (r *ContestServiceIdGroupRouter) {
	r = &ContestServiceIdGroupRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("contest/:cid"),
			AuthRouter: h.GetAuthRouter().Group("contest/:cid"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), "contest:cid"),
		},
	}

	r.Count = NewContestServiceIdGroupCountRouter(traits, r.H)
	r.Post = NewContestServiceIdGroupPostRouter(traits, r.H)
	r.IdGroup = NewContestServiceIdGroupIdGroupRouter(traits, r.H)
	r.List = NewContestServiceIdGroupListRouter(traits, r.H)

	r.GetContest = r.GetRouter().GET("", traits.GetServiceInstance("ContestService").(ContestService).GetContest)
	r.PutContest = r.GetRouter().PUT("", traits.GetServiceInstance("ContestService").(ContestService).PutContest)
	r.Delete = r.GetRouter().DELETE("", traits.GetServiceInstance("ContestService").(ContestService).Delete)

	return
}

type ContestServiceIdGroupCountRouter struct {
	H

	CountContestProblem *LeafRouter
}

func NewContestServiceIdGroupCountRouter(traits GenerateRouterTraits, h H) (r *ContestServiceIdGroupCountRouter) {
	r = &ContestServiceIdGroupCountRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("problem-count"),
			AuthRouter: h.GetAuthRouter().Group("problem-count"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), ""),
		},
	}

	r.CountContestProblem = r.GetRouter().GET("", traits.GetServiceInstance("ContestService").(ContestService).CountContestProblem)

	return
}

type ContestServiceIdGroupPostRouter struct {
	H

	PostContestProblem *LeafRouter
}

func NewContestServiceIdGroupPostRouter(traits GenerateRouterTraits, h H) (r *ContestServiceIdGroupPostRouter) {
	r = &ContestServiceIdGroupPostRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("problem"),
			AuthRouter: h.GetAuthRouter().Group("problem"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), ""),
		},
	}

	r.PostContestProblem = r.GetAuthRouter().POST("", traits.GetServiceInstance("ContestService").(ContestService).PostContestProblem)
	r.PostContestProblem = traits.ApplyAuthOnMethod(r.PostContestProblem, "~")

	return
}

type ContestServiceIdGroupIdGroupRouter struct {
	H
	ProblemDesc *ContestServiceIdGroupIdGroupProblemDescRouter

	GetContestProblem    *LeafRouter
	PutContestProblem    *LeafRouter
	DeleteContestProblem *LeafRouter
}

func NewContestServiceIdGroupIdGroupRouter(traits GenerateRouterTraits, h H) (r *ContestServiceIdGroupIdGroupRouter) {
	r = &ContestServiceIdGroupIdGroupRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("problem/:pid"),
			AuthRouter: h.GetAuthRouter().Group("problem/:pid"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), "problem:pid"),
		},
	}

	r.ProblemDesc = NewContestServiceIdGroupIdGroupProblemDescRouter(traits, r.H)

	r.GetContestProblem = r.GetRouter().GET("", traits.GetServiceInstance("ContestService").(ContestService).GetContestProblem)
	r.PutContestProblem = r.GetRouter().PUT("", traits.GetServiceInstance("ContestService").(ContestService).PutContestProblem)
	r.DeleteContestProblem = r.GetRouter().DELETE("", traits.GetServiceInstance("ContestService").(ContestService).DeleteContestProblem)

	return
}

type ContestServiceIdGroupIdGroupProblemDescRouter struct {
	H
	ProblemDesc *ContestServiceIdGroupIdGroupProblemDescProblemDescRouter

	PostContestProblemDesc   *LeafRouter
	GetContestProblemDesc    *LeafRouter
	PutContestProblemDesc    *LeafRouter
	DeleteContestProblemDesc *LeafRouter
}

func NewContestServiceIdGroupIdGroupProblemDescRouter(traits GenerateRouterTraits, h H) (r *ContestServiceIdGroupIdGroupProblemDescRouter) {
	r = &ContestServiceIdGroupIdGroupProblemDescRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("/desc"),
			AuthRouter: h.GetAuthRouter().Group("/desc"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), ""),
		},
	}

	r.ProblemDesc = NewContestServiceIdGroupIdGroupProblemDescProblemDescRouter(traits, r.H)

	r.PostContestProblemDesc = r.GetRouter().POST("", traits.GetServiceInstance("ContestService").(ContestService).PostContestProblemDesc)
	r.GetContestProblemDesc = r.GetRouter().GET("", traits.GetServiceInstance("ContestService").(ContestService).GetContestProblemDesc)
	r.PutContestProblemDesc = r.GetRouter().PUT("", traits.GetServiceInstance("ContestService").(ContestService).PutContestProblemDesc)
	r.DeleteContestProblemDesc = r.GetRouter().DELETE("", traits.GetServiceInstance("ContestService").(ContestService).DeleteContestProblemDesc)

	return
}

type ContestServiceIdGroupIdGroupProblemDescProblemDescRouter struct {
	H

	ChangeContestProblemDescriptionRef *LeafRouter
}

func NewContestServiceIdGroupIdGroupProblemDescProblemDescRouter(traits GenerateRouterTraits, h H) (r *ContestServiceIdGroupIdGroupProblemDescProblemDescRouter) {
	r = &ContestServiceIdGroupIdGroupProblemDescProblemDescRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("/desc"),
			AuthRouter: h.GetAuthRouter().Group("/desc"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), ""),
		},
	}

	r.ChangeContestProblemDescriptionRef = r.GetRouter().POST("", traits.GetServiceInstance("ContestService").(ContestService).ChangeContestProblemDescriptionRef)

	return
}

type ContestServiceIdGroupListRouter struct {
	H

	ListContestProblems *LeafRouter
}

func NewContestServiceIdGroupListRouter(traits GenerateRouterTraits, h H) (r *ContestServiceIdGroupListRouter) {
	r = &ContestServiceIdGroupListRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("problem-list"),
			AuthRouter: h.GetAuthRouter().Group("problem-list"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), ""),
		},
	}

	r.ListContestProblems = r.GetRouter().GET("", traits.GetServiceInstance("ContestService").(ContestService).ListContestProblems)

	return
}

type GroupServiceRouter struct {
	H
	List    *GroupServiceListRouter
	Count   *GroupServiceCountRouter
	Post    *GroupServicePostRouter
	IdGroup *GroupServiceIdGroupRouter
}

func NewGroupServiceRouter(traits GenerateRouterTraits, h H) (r *GroupServiceRouter) {
	r = &GroupServiceRouter{
		H: &BaseH{
			Router:     h.GetRouter().Extend("GroupService"),
			AuthRouter: h.GetAuthRouter().Extend("GroupService"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), ""),
		},
	}

	r.List = NewGroupServiceListRouter(traits, r.H)
	r.Count = NewGroupServiceCountRouter(traits, r.H)
	r.Post = NewGroupServicePostRouter(traits, r.H)
	r.IdGroup = NewGroupServiceIdGroupRouter(traits, r.H)

	return
}

type GroupServiceListRouter struct {
	H

	ListGroups *LeafRouter
}

func NewGroupServiceListRouter(traits GenerateRouterTraits, h H) (r *GroupServiceListRouter) {
	r = &GroupServiceListRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("group-list"),
			AuthRouter: h.GetAuthRouter().Group("group-list"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), ""),
		},
	}

	r.ListGroups = r.GetRouter().GET("", traits.GetServiceInstance("GroupService").(GroupService).ListGroups)

	return
}

type GroupServiceCountRouter struct {
	H

	CountGroup *LeafRouter
}

func NewGroupServiceCountRouter(traits GenerateRouterTraits, h H) (r *GroupServiceCountRouter) {
	r = &GroupServiceCountRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("group-count"),
			AuthRouter: h.GetAuthRouter().Group("group-count"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), ""),
		},
	}

	r.CountGroup = r.GetRouter().GET("", traits.GetServiceInstance("GroupService").(GroupService).CountGroup)

	return
}

type GroupServicePostRouter struct {
	H

	PostGroup *LeafRouter
}

func NewGroupServicePostRouter(traits GenerateRouterTraits, h H) (r *GroupServicePostRouter) {
	r = &GroupServicePostRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("group"),
			AuthRouter: h.GetAuthRouter().Group("group"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), ""),
		},
	}

	r.PostGroup = r.GetAuthRouter().POST("", traits.GetServiceInstance("GroupService").(GroupService).PostGroup)
	r.PostGroup = traits.ApplyAuthOnMethod(r.PostGroup, "~")

	return
}

type GroupServiceIdGroupRouter struct {
	H
	Owner    *GroupServiceIdGroupOwnerRouter
	UserList *GroupServiceIdGroupUserListRouter
	User     *GroupServiceIdGroupUserRouter

	GetGroup    *LeafRouter
	PutGroup    *LeafRouter
	DeleteGroup *LeafRouter
}

func NewGroupServiceIdGroupRouter(traits GenerateRouterTraits, h H) (r *GroupServiceIdGroupRouter) {
	r = &GroupServiceIdGroupRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("group/:gid"),
			AuthRouter: h.GetAuthRouter().Group("group/:gid"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), "group:gid"),
		},
	}

	r.Owner = NewGroupServiceIdGroupOwnerRouter(traits, r.H)
	r.UserList = NewGroupServiceIdGroupUserListRouter(traits, r.H)
	r.User = NewGroupServiceIdGroupUserRouter(traits, r.H)

	r.GetGroup = r.GetRouter().GET("", traits.GetServiceInstance("GroupService").(GroupService).GetGroup)
	r.PutGroup = r.GetRouter().PUT("", traits.GetServiceInstance("GroupService").(GroupService).PutGroup)
	r.DeleteGroup = r.GetRouter().DELETE("", traits.GetServiceInstance("GroupService").(GroupService).DeleteGroup)

	return
}

type GroupServiceIdGroupOwnerRouter struct {
	H

	PutGroupOwner *LeafRouter
}

func NewGroupServiceIdGroupOwnerRouter(traits GenerateRouterTraits, h H) (r *GroupServiceIdGroupOwnerRouter) {
	r = &GroupServiceIdGroupOwnerRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("/owner"),
			AuthRouter: h.GetAuthRouter().Group("/owner"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), ""),
		},
	}

	r.PutGroupOwner = r.GetRouter().PUT("", traits.GetServiceInstance("GroupService").(GroupService).PutGroupOwner)

	return
}

type GroupServiceIdGroupUserListRouter struct {
	H

	GetGroupMembers *LeafRouter
}

func NewGroupServiceIdGroupUserListRouter(traits GenerateRouterTraits, h H) (r *GroupServiceIdGroupUserListRouter) {
	r = &GroupServiceIdGroupUserListRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("/user-list"),
			AuthRouter: h.GetAuthRouter().Group("/user-list"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), ""),
		},
	}

	r.GetGroupMembers = r.GetRouter().GET("", traits.GetServiceInstance("GroupService").(GroupService).GetGroupMembers)

	return
}

type GroupServiceIdGroupUserRouter struct {
	H

	PostGroupMember *LeafRouter
}

func NewGroupServiceIdGroupUserRouter(traits GenerateRouterTraits, h H) (r *GroupServiceIdGroupUserRouter) {
	r = &GroupServiceIdGroupUserRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("user/:id"),
			AuthRouter: h.GetAuthRouter().Group("user/:id"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), "user:id"),
		},
	}

	r.PostGroupMember = r.GetRouter().POST("", traits.GetServiceInstance("GroupService").(GroupService).PostGroupMember)

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
	Inspect *UserServiceIdGroupInspectRouter
	Email   *UserServiceIdGroupEmailRouter

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

	r.Inspect = NewUserServiceIdGroupInspectRouter(traits, r.H)
	r.Email = NewUserServiceIdGroupEmailRouter(traits, r.H)

	r.GetUser = r.GetRouter().GET("", traits.GetServiceInstance("UserService").(UserService).GetUser)
	r.PutUser = r.GetRouter().PUT("", traits.GetServiceInstance("UserService").(UserService).PutUser)
	r.Delete = r.GetRouter().DELETE("", traits.GetServiceInstance("UserService").(UserService).Delete)

	return
}

type UserServiceIdGroupInspectRouter struct {
	H

	InspectUser *LeafRouter
}

func NewUserServiceIdGroupInspectRouter(traits GenerateRouterTraits, h H) (r *UserServiceIdGroupInspectRouter) {
	r = &UserServiceIdGroupInspectRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("/inspect"),
			AuthRouter: h.GetAuthRouter().Group("/inspect"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), ""),
		},
	}

	r.InspectUser = r.GetRouter().GET("", traits.GetServiceInstance("UserService").(UserService).InspectUser)

	return
}

type UserServiceIdGroupEmailRouter struct {
	H

	BindEmail *LeafRouter
}

func NewUserServiceIdGroupEmailRouter(traits GenerateRouterTraits, h H) (r *UserServiceIdGroupEmailRouter) {
	r = &UserServiceIdGroupEmailRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("/email"),
			AuthRouter: h.GetAuthRouter().Group("/email"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), ""),
		},
	}

	r.BindEmail = r.GetRouter().PUT("", traits.GetServiceInstance("UserService").(UserService).BindEmail)

	return
}

type AuthServiceRouter struct {
	H
	Policy         *AuthServicePolicyRouter
	GroupingPolicy *AuthServiceGroupingPolicyRouter
}

func NewAuthServiceRouter(traits GenerateRouterTraits, h H) (r *AuthServiceRouter) {
	r = &AuthServiceRouter{
		H: &BaseH{
			Router:     h.GetRouter().Extend("AuthService"),
			AuthRouter: h.GetAuthRouter().Extend("AuthService"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), ""),
		},
	}

	r.Policy = NewAuthServicePolicyRouter(traits, r.H)
	r.GroupingPolicy = NewAuthServiceGroupingPolicyRouter(traits, r.H)

	return
}

type AuthServicePolicyRouter struct {
	H

	AddPolicy    *LeafRouter
	RemovePolicy *LeafRouter
	HasPolicy    *LeafRouter
}

func NewAuthServicePolicyRouter(traits GenerateRouterTraits, h H) (r *AuthServicePolicyRouter) {
	r = &AuthServicePolicyRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("/policy"),
			AuthRouter: h.GetAuthRouter().Group("/policy"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), ""),
		},
	}

	r.AddPolicy = r.GetRouter().POST("", traits.GetServiceInstance("AuthService").(AuthService).AddPolicy)
	r.RemovePolicy = r.GetRouter().DELETE("", traits.GetServiceInstance("AuthService").(AuthService).RemovePolicy)
	r.HasPolicy = r.GetRouter().GET("", traits.GetServiceInstance("AuthService").(AuthService).HasPolicy)

	return
}

type AuthServiceGroupingPolicyRouter struct {
	H

	AddGroupingPolicy    *LeafRouter
	RemoveGroupingPolicy *LeafRouter
	HasGroupingPolicy    *LeafRouter
}

func NewAuthServiceGroupingPolicyRouter(traits GenerateRouterTraits, h H) (r *AuthServiceGroupingPolicyRouter) {
	r = &AuthServiceGroupingPolicyRouter{
		H: &BaseH{
			Router:     h.GetRouter().Group("/policy/group"),
			AuthRouter: h.GetAuthRouter().Group("/policy/group"),
			Auth:       traits.ApplyRouteMeta(h.GetAuth(), ""),
		},
	}

	r.AddGroupingPolicy = r.GetRouter().POST("", traits.GetServiceInstance("AuthService").(AuthService).AddGroupingPolicy)
	r.RemoveGroupingPolicy = r.GetRouter().DELETE("", traits.GetServiceInstance("AuthService").(AuthService).RemoveGroupingPolicy)
	r.HasGroupingPolicy = r.GetRouter().GET("", traits.GetServiceInstance("AuthService").(AuthService).HasGroupingPolicy)

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
