package provider

import (
	"fmt"
	"github.com/Myriad-Dreamin/boj-v6/abstract/announcement"
	"github.com/Myriad-Dreamin/boj-v6/abstract/auth"
	"github.com/Myriad-Dreamin/boj-v6/abstract/comment"
	"github.com/Myriad-Dreamin/boj-v6/abstract/contest"
	"github.com/Myriad-Dreamin/boj-v6/abstract/group"
	"github.com/Myriad-Dreamin/boj-v6/abstract/problem"
	"github.com/Myriad-Dreamin/boj-v6/abstract/submission"
	"github.com/Myriad-Dreamin/boj-v6/abstract/user"
	"github.com/Myriad-Dreamin/minimum-lib/module"
	"path"
)

type Service struct {
	module.BaseModuler
	authService         auth.Service
	announcementService announcement.Service
	commentService      comment.Service
	submissionService   submission.Service
	contestService      contest.Service
	problemService      problem.Service
	userService         user.Service
	groupService        group.Service
	//contestFAQService   *contestcomment.Service
	//contestSubmissionService   *contestsubmission.Service
	//classService        *class.Service
}

func NewService(namespace string) *Service {
	return &Service{
		BaseModuler: module.BaseModuler{
			Namespace: namespace,
		},
		//subControllers: []SubController{JustProvide(control.AdditionControllerDescs...)},
	}
}

func (s *Service) Register(name string, service interface{}) {
	if err := s.Provide(path.Join(s.Namespace, name), service); err != nil {
		panic(fmt.Errorf("unknown/registered service %T, err %v", service, err))
	}

	switch ss := service.(type) {
	case announcement.Service:
		s.announcementService = ss
	case submission.Service:
		s.submissionService = ss
	case contest.Service:
		s.contestService = ss
	case comment.Service:
		s.commentService = ss
	case problem.Service:
		s.problemService = ss
	case user.Service:
		s.userService = ss
	case auth.Service:
		s.authService = ss
	case group.Service:
		s.groupService = ss
	//case *contestsubmission.Service:
	//	s.contestSubmissionService = ss
	//case *contestFAQ.Service:
	//	s.contestFAQService = ss
	//case *class.Service:
	//	s.classService = ss
	default:
		panic(fmt.Errorf("unknown service %T", service))
	}
}

//func (s *Service) ContestSubmissionService() contestsubmission.Service {
//	return s.contestSubmissionService
//}

//func (s *Service) ContestFAQService() contestFAQ.Service {
//	return s.contestFAQService
//}

//func (s *Service) ClassService() class.Service {
//	return s.classService
//}

func (s *Service) ContestService() contest.Service {
	return s.contestService
}

func (s *Service) ProblemService() problem.Service {
	return s.problemService
}

func (s *Service) CommentService() comment.Service {
	return s.commentService
}

func (s *Service) UserService() user.Service {
	return s.userService
}

func (s *Service) SubmissionService() submission.Service {
	return s.submissionService
}

func (s *Service) AuthService() auth.Service {
	return s.authService
}

func (s *Service) GroupService() group.Service {
	return s.groupService
}

func (s *Service) AnnouncementService() announcement.Service {
	return s.announcementService
}
