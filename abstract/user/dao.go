package user

import "github.com/Myriad-Dreamin/boj-v6/abstract/db"

type DB interface {
	db.BasicDB

	Create(a *User) (int64, error)
	Update(a *User) (int64, error)
	Delete(a *User) (int64, error)
	UpdateFields(a *User, fields []string) (int64, error)
	Find(page, pageSize int) ([]User, error)
	Count() (int64, error)

	ID(id uint) (*User, error)
	QueryUserName(name string) (*User, error)
	QueryEmail(email string) (*User, error)

	AuthenticatePassword(user *User, password string) (verified bool, err error)
	RecalculatePassword(user *User, password string) (err error)
}
