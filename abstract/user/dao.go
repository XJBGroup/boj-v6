package user

type DB interface {
	Create(a *User) (int64, error)
	Update(a *User) (int64, error)
	Delete(a *User) (int64, error)
	UpdateFields(a *User, fields []string) (int64, error)
	Find(page, pageSize int) ([]User, error)
	Count() (int64, error)

	ID(id uint) (*User, error)
	QueryName(name string) (*User, error)
	QueryEmail(email string) (*User, error)
	AuthenticatePassword(user *User, password string) (verified bool, err error)
}
