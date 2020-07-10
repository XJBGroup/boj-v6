package user

type DB interface {
	Create(a *User) (int64, error)
	Update(a *User) (int64, error)
	Delete(a *User) (int64, error)
	UpdateFields(a *User, fields []string) (int64, error)

	ID(id uint) (*User, error)
}
