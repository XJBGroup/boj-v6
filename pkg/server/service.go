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
)

type serviceResult struct {
	serviceName string
	functional.DecayResult
}

func (srv *Server) PrepareService() bool {
	for _, serviceResult := range []serviceResult{
		{"authService",
			functional.Decay(auth.NewService(srv.Module))},
		{"announcementService",
			functional.Decay(announcement.NewService(srv.Module))},
		{"commentService",
			functional.Decay(comment.NewService(srv.Module))},
		{"submissionService",
			functional.Decay(submission.NewService(srv.Module))},
		{"userService",
			functional.Decay(user.NewService(srv.Module))},
		{"problemService",
			functional.Decay(problem.NewService(srv.Module))},
		{"contestService",
			functional.Decay(contest.NewService(srv.Module))},
		{"groupService",
			functional.Decay(group.NewService(srv.Module))},
	} {
		// build Router failed when requesting service with database, report and return
		if serviceResult.Err != nil {
			srv.Logger.Debug(fmt.Sprintf("get %T service error", serviceResult.First), "error", serviceResult.Err)
			return false
		}
		srv.ServiceProvider.Register(serviceResult.serviceName, serviceResult.First)
	}
	return true
}
