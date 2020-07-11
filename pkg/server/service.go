package server

import (
	"fmt"
	"github.com/Myriad-Dreamin/boj-v6/app/announcement"
	"github.com/Myriad-Dreamin/boj-v6/app/comment"
	"github.com/Myriad-Dreamin/boj-v6/app/user"
	"github.com/Myriad-Dreamin/functional-go"
)

type serviceResult struct {
	serviceName string
	functional.DecayResult
}

func (srv *Server) PrepareService() bool {
	for _, serviceResult := range []serviceResult{
		{"announcementService",
			functional.Decay(announcement.NewService(srv.Module))},
		{"userService",
			functional.Decay(user.NewService(srv.Module))},
		{"commentService",
			functional.Decay(comment.NewService(srv.Module))},
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
