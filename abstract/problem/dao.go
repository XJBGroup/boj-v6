package problem

type DB interface {
	Create(a *Problem) (int64, error)
	Update(a *Problem) (int64, error)
	Delete(a *Problem) (int64, error)
	UpdateFields(a *Problem, fields []string) (int64, error)
	Find(page, pageSize int) ([]Problem, error)
	Count() (int64, error)

	ID(id uint) (*Problem, error)
}
