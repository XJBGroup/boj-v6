package submission

import "github.com/Myriad-Dreamin/boj-v6/abstract/db"

type DB interface {
	db.BasicDB

	Create(a *Submission) (int64, error)
	Update(a *Submission) (int64, error)
	Delete(a *Submission) (int64, error)
	UpdateFields(a *Submission, fields []string) (int64, error)
	Find(page, pageSize int) ([]Submission, error)
	Filter(filter *Filter) ([]Submission, error)
	FilterCount(filter *Filter) (int64, error)
	Count() (int64, error)

	ID(id uint) (*Submission, error)

	HasHash(hash string) (exists bool, err error)
}

type Filter struct {
	Page         int
	PageSize     int
	Order        string
	ByUser       uint
	OnProblem    uint
	WithLanguage uint8
	HasStatus    int64
}
