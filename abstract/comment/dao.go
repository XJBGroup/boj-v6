package comment

type DB interface {
	Create(a *Comment) (int64, error)
	Update(a *Comment) (int64, error)
	Delete(a *Comment) (int64, error)
	UpdateFields(a *Comment, fields []string) (int64, error)
	Find(page, pageSize int) ([]Comment, error)
	Count() (int64, error)

	ID(id uint) (*Comment, error)
}
