package problem_desc

import (
	"time"
)

type ProblemDesc struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at" form:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" form:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at" form:"deleted_at"`

	// todo: add unique index
	ProblemID  uint   `dorm:"pid;not null" gorm:"column:pid" json:"pid"`
	Name       string `dorm:"name;not null" gorm:"column:name" json:"name"`
	Key        []byte `gorm:"-" json:"-"`
	Content    []byte `gorm:"-" json:"-"`
	FreeHandle func() `gorm:"-" json:"-"`
}

func (ProblemDesc) TableName() string {
	return "problem_desc"
}

func (p ProblemDesc) GetID() uint {
	return p.ID
}
