package announcement

import "github.com/Myriad-Dreamin/boj-v6/abstract/db"

type DB interface {
	db.BasicDB

	Create(a *Announcement) (int64, error)
	Update(a *Announcement) (int64, error)
	Delete(a *Announcement) (int64, error)
	UpdateFields(a *Announcement, fields []string) (int64, error)
	Find(page, pageSize int) ([]Announcement, error)
	Count() (int64, error)

	ID(id uint) (*Announcement, error)
}
