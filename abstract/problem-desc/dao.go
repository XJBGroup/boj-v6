package problem_desc

type DB interface {
	Create(a *ProblemDesc) (int64, error)
	Update(a *ProblemDesc) (int64, error)
	Delete(a *ProblemDesc) (int64, error)
	UpdateFields(a *ProblemDesc, fields []string) (int64, error)
	Find(page, pageSize int) ([]ProblemDesc, error)
	Count() (int64, error)

	ID(id uint) (*ProblemDesc, error)

	LoadDesc(a *ProblemDesc) error
	SaveDesc(a *ProblemDesc) error
	ReleaseDesc(a *ProblemDesc) error
	DeleteDesc(a *ProblemDesc) error
	RenameDesc(a *ProblemDesc, newName string) (int64, error)
}
