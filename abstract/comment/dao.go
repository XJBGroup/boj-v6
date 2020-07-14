package comment

type DB interface {
	Create(a *Comment) (int64, error)
	Update(a *Comment) (int64, error)
	Delete(a *Comment) (int64, error)
	UpdateFields(a *Comment, fields []string) (int64, error)
	Find(page, pageSize int) ([]Comment, error)
	Filter(filter *Filter) ([]Comment, error)
	FilterCount(f *Filter) (int64, error)
	Count() (int64, error)

	ID(id uint) (*Comment, error)
}

type Filter struct {
	Page     int   `json:"page" form:"page"`
	PageSize int   `json:"page_size" form:"page_size"`
	RefType  uint8 `json:"ref_type" form:"ref_type"`
	Ref      uint  `json:"ref" form:"ref"`
	NoReply  bool  `json:"no_reply" form:"no_reply"`
}
