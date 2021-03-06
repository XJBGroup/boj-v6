package problem

import "github.com/Myriad-Dreamin/boj-v6/abstract/db"

type DB interface {
	db.BasicDB

	Create(a *Problem) (int64, error)
	Update(a *Problem) (int64, error)
	Delete(a *Problem) (int64, error)
	UpdateFields(a *Problem, fields []string) (int64, error)
	Find(page, pageSize int) ([]Problem, error)
	Count() (int64, error)

	ID(id uint) (*Problem, error)
	HasByPID(pid uint) (exists bool, err error)
}
