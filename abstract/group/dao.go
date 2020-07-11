package group

type DB interface {
	Create(a *Group) (int64, error)
	Update(a *Group) (int64, error)
	Delete(a *Group) (int64, error)
	UpdateFields(a *Group, fields []string) (int64, error)
	Find(page, pageSize int) ([]Group, error)
	Count() (int64, error)

	ID(id uint) (*Group, error)
}
