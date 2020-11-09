package server

import (
	"fmt"
	"github.com/Myriad-Dreamin/boj-v6/app/announcement"
	"github.com/Myriad-Dreamin/boj-v6/app/auth"
	"github.com/Myriad-Dreamin/boj-v6/app/comment"
	"github.com/Myriad-Dreamin/boj-v6/app/contest"
	"github.com/Myriad-Dreamin/boj-v6/app/group"
	"github.com/Myriad-Dreamin/boj-v6/app/problem"
	"github.com/Myriad-Dreamin/boj-v6/app/submission"
	"github.com/Myriad-Dreamin/boj-v6/app/user"
	"github.com/Myriad-Dreamin/functional-go"
	"path"
)

type serviceResult struct {
	serviceName string
	proto       interface{}
	functional.DecayResult
}

func (srv *Server) PrepareService() bool {
	for _, serviceResult := range []serviceResult{

		//{"InnerSubmissionController", new(*submission.InnerService),
		//	functional.Decay(submission.NewInnerService(srv.Module))},
	} {
		// build Router failed when requesting service with database, report and return
		if serviceResult.Err != nil {
			srv.Logger.Debug(fmt.Sprintf("get %T service error", serviceResult.First), "error", serviceResult.Err)
			return false
		}
		err := srv.Module.ProvideNamedImpl(
			path.Join("minimum", serviceResult.serviceName), serviceResult.proto, serviceResult.First)
		if err != nil {
			srv.Logger.Debug("provide service error", "name", serviceResult.First)
			return false
		}
	}

	for _, serviceResult := range []serviceResult{
		{"AuthController", new(*auth.Controller),
			functional.Decay(auth.NewController(srv.Module))},
		{"AnnouncementController", new(*announcement.Controller),
			functional.Decay(announcement.NewController(srv.Module))},
		{"CommentController", new(*comment.Controller),
			functional.Decay(comment.NewController(srv.Module))},
		{"SubmissionController", new(*submission.Controller),
			functional.Decay(submission.NewController(srv.Module))},
		{"UserController", new(*user.Controller),
			functional.Decay(user.NewController(srv.Module))},
		{"ProblemController", new(*problem.Controller),
			functional.Decay(problem.NewController(srv.Module))},
		{"ContestController", new(*contest.Controller),
			functional.Decay(contest.NewController(srv.Module))},
		{"GroupController", new(*group.Controller),
			functional.Decay(group.NewController(srv.Module))},
	} {
		// build Router failed when requesting service with database, report and return
		if serviceResult.Err != nil {
			srv.Logger.Debug(fmt.Sprintf("get %T service error", serviceResult.First), "error", serviceResult.Err)
			return false
		}
		err := srv.Module.ProvideNamedImpl(
			path.Join("minimum", serviceResult.serviceName), serviceResult.proto, serviceResult.First)
		if err != nil {
			srv.Logger.Debug("provide service error", "name", serviceResult.First)
			return false
		}
	}
	return true
}
