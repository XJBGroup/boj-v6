package contest

import "github.com/Myriad-Dreamin/boj-v6/abstract/db"

type DB interface {
	db.BasicDB

	Create(a *Contest) (int64, error)
	Update(a *Contest) (int64, error)
	Delete(a *Contest) (int64, error)
	UpdateFields(a *Contest, fields []string) (int64, error)
	Find(page, pageSize int) ([]Contest, error)
	Count() (int64, error)

	ID(id uint) (*Contest, error)
}
