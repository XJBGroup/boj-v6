package provider

import (
	"fmt"
	"github.com/Myriad-Dreamin/boj-v6/abstract/announcement"
	"github.com/Myriad-Dreamin/boj-v6/abstract/comment"
	"github.com/Myriad-Dreamin/boj-v6/abstract/contest"
	"github.com/Myriad-Dreamin/boj-v6/abstract/group"
	"github.com/Myriad-Dreamin/boj-v6/abstract/problem"
	"github.com/Myriad-Dreamin/boj-v6/abstract/submission"
	"github.com/Myriad-Dreamin/boj-v6/abstract/user"
	"github.com/Myriad-Dreamin/minimum-lib/module"
	"github.com/casbin/casbin/v2"
	"path"
)

type DB struct {
	module.BaseModuler
	announcementDB announcement.DB
	userDB         user.DB
	commentDB      comment.DB
	submissionDB   submission.DB
	problemDB      problem.DB
	contestDB      contest.DB
	groupDB        group.DB
	enforcer       *Enforcer
}

func NewDB(namespace string) *DB {
	return &DB{
		BaseModuler: module.BaseModuler{
			Namespace: namespace,
		},
	}
}

func (s *DB) AnnouncementDB() announcement.DB {
	return s.announcementDB
}

func (s *DB) SubmissionDB() submission.DB {
	return s.submissionDB
}

func (s *DB) ProblemDB() problem.DB {
	return s.problemDB
}

func (s *DB) ContestDB() contest.DB {
	return s.contestDB
}

func (s *DB) GroupDB() group.DB {
	return s.groupDB
}

func (s *DB) UserDB() user.DB {
	return s.userDB
}
func (s *DB) CommentDB() comment.DB {
	return s.commentDB
}

type Enforcer = casbin.SyncedEnforcer

func (s *DB) Enforcer() *Enforcer {
	return s.enforcer
}

func (s *DB) Register(name string, database interface{}) {
	if err := s.Provide(path.Join(s.Namespace, name), database); err != nil {
		panic(fmt.Errorf("unknown database %T", database))
	}

	switch ss := database.(type) {
	case announcement.DB:
		s.announcementDB = ss
	case user.DB:
		s.userDB = ss
	case comment.DB:
		s.commentDB = ss
	case submission.DB:
		s.submissionDB = ss
	case problem.DB:
		s.problemDB = ss
	case contest.DB:
		s.contestDB = ss
	case group.DB:
		s.groupDB = ss
	case *Enforcer:
		s.enforcer = ss
	default:
		if mm, ok := ss.(module.Moduler); ok {
			// todo:
			_ = mm
		}
	}
}
