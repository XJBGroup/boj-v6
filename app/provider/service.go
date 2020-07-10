package provider

import (
	"fmt"
	"github.com/Myriad-Dreamin/boj-v6/abstract/announcement"
	"github.com/Myriad-Dreamin/minimum-lib/module"
	"path"
)

type Service struct {
	module.BaseModuler
	//authService         *authservice.Service
	announcementService announcement.Service
	//commentService      *commentservice.Service
	//contestFAQService *contestFAQservice.Service
	//submissionService   *submissionservice.Service
	//contestSubmissionService   *contestsubmissionservice.Service
	//contestService      *contestservice.Service
	//problemService      *problemservice.Service
	//userService         *userservice.Service
	//groupService        *groupservice.Service
	//classService        *classservice.Service
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
	//case *submissionservice.Service:
	//	s.submissionService = ss
	//case *contestsubmissionservice.Service:
	//	s.contestSubmissionService = ss
	//case *contestFAQservice.Service:
	//	s.contestFAQService = ss
	//case *contestservice.Service:
	//	s.contestService = ss
	//case *commentservice.Service:
	//	s.commentService = ss
	//case *problemservice.Service:
	//	s.problemService = ss
	//case *userservice.Service:
	//	s.userService = ss
	//case *authservice.Service:
	//	s.authService = ss
	//case *groupservice.Service:
	//	s.groupService = ss
	//case *classservice.Service:
	//	s.classService = ss
	default:
		panic(fmt.Errorf("unknown service %T", service))
	}
}

//func (s *Service) SubmissionService() *submissionservice.Service {
//	return s.submissionService
//}
//
//func (s *Service) ContestSubmissionService() *contestsubmissionservice.Service {
//	return s.contestSubmissionService
//}
//
//func (s *Service) ContestService() *contestservice.Service {
//	return s.contestService
//}
//
//func (s *Service) ContestFAQService() *contestFAQservice.Service {
//	return s.contestFAQService
//}
//
//func (s *Service) ProblemService() *problemservice.Service {
//	return s.problemService
//}
//
//func (s *Service) CommentService() *commentservice.Service {
//	return s.commentService
//}
//
//func (s *Service) UserService() *userservice.Service {
//	return s.userService
//}
//
//func (s *Service) AuthService() *authservice.Service {
//	return s.authService
//}
//
//func (s *Service) GroupService() *groupservice.Service {
//	return s.groupService
//}

func (s *Service) AnnouncementService() announcement.Service {
	return s.announcementService
}

//func (s *Service) ClassService() *classservice.Service {
//	return s.classService
//}
