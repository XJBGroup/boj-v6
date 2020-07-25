package problem_desc

import (
	"time"
)

type ProblemDesc struct {
	ID         uint `gorm:"primary_key" json:"id"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time `sql:"index"`
	ProblemID  uint       `dorm:"pid;not null" gorm:"column:pid" json:"pid"`
	Name       string     `dorm:"name;not null" gorm:"column:name;unique" json:"name"`
	Key        []byte     `gorm:"-" json:"-"`
	Content    []byte     `gorm:"-" json:"-"`
	FreeHandle func()     `gorm:"-" json:"-"`
}

func (ProblemDesc) TableName() string {
	return "problem_desc"
}

func (p ProblemDesc) GetID() uint {
	return p.ID
}
