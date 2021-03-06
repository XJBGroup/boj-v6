package problem_desc

import "github.com/Myriad-Dreamin/boj-v6/abstract/db"

type DB interface {
	db.BasicDB

	NewProblemDesc(pid uint, name string, content []byte) *ProblemDesc
	Create(a *ProblemDesc) (int64, error)
	Update(a *ProblemDesc) (int64, error)
	Delete(a *ProblemDesc) (int64, error)
	UpdateFields(a *ProblemDesc, fields []string) (int64, error)
	Find(problemId uint, page, pageSize int) ([]ProblemDesc, error)
	Count(problemId uint) (int64, error)
	QueryByKey(pid uint, pdName string) (pd *ProblemDesc, err error)
	QueryByPID(pid uint) (pd []ProblemDesc, err error)
	QueryExistenceByKey(pid uint, pdName string) (exists bool, err error)

	ID(id uint) (*ProblemDesc, error)

	LoadDesc(a *ProblemDesc) error
	SaveDesc(a *ProblemDesc) error
	ReleaseDesc(a *ProblemDesc) error
	DeleteDesc(a *ProblemDesc) error
	RenameDesc(a *ProblemDesc, newName string) (int64, error)
	InvalidateDescCache(a *ProblemDesc) error
}
