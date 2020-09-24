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

		{"InnerSubmissionService", new(*submission.InnerService),
			functional.Decay(submission.NewInnerService(srv.Module))},
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
		{"AuthService", new(*auth.Service),
			functional.Decay(auth.NewService(srv.Module))},
		{"AnnouncementService", new(*announcement.Service),
			functional.Decay(announcement.NewService(srv.Module))},
		{"CommentService", new(*comment.Service),
			functional.Decay(comment.NewService(srv.Module))},
		{"SubmissionService", new(*submission.Service),
			functional.Decay(submission.NewService(srv.Module))},
		{"UserService", new(*user.Service),
			functional.Decay(user.NewService(srv.Module))},
		{"ProblemService", new(*problem.Service),
			functional.Decay(problem.NewService(srv.Module))},
		{"ContestService", new(*contest.Service),
			functional.Decay(contest.NewService(srv.Module))},
		{"GroupService", new(*group.Service),
			functional.Decay(group.NewService(srv.Module))},
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
