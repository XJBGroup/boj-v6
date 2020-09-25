package problem

import (
	"github.com/Myriad-Dreamin/boj-v6/abstract/problem"
)

func (db *DBImpl) HasByPID(pid uint) (exists bool, err error) {
	return db.GORMDBImpl.Has(new(problem.Problem), "id = ?", pid)
}
