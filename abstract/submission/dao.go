package submission

type DB interface {
	Create(a *Submission) (int64, error)
	Update(a *Submission) (int64, error)
	Delete(a *Submission) (int64, error)
	UpdateFields(a *Submission, fields []string) (int64, error)
	Find(page, pageSize int) ([]Submission, error)
	Count() (int64, error)

	ID(id uint) (*Submission, error)
}
