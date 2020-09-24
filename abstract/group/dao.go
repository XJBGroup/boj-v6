package group

import "github.com/Myriad-Dreamin/boj-v6/abstract/db"

type DB interface {
	db.BasicDB

	Create(a *Group) (int64, error)
	Update(a *Group) (int64, error)
	Delete(a *Group) (int64, error)
	UpdateFields(a *Group, fields []string) (int64, error)
	Find(page, pageSize int) ([]Group, error)
	Count() (int64, error)

	ID(id uint) (*Group, error)
}
