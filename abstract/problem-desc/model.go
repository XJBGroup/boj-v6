package problem_desc

import (
	"github.com/Myriad-Dreamin/boj-v6/deployment/oss"
	"time"
)

type ProblemDesc struct {
	ID        uint `gorm:"primary_key" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time      `sql:"index"`
	ProblemID uint            `dorm:"pid;not null" gorm:"column:pid" json:"pid"`
	Name      string          `dorm:"name;not null" gorm:"column:name;unique" json:"name"`
	desc      oss.ProblemDesc `gorm:"-" json:"-"`
}

func NewProblemDesc(pid uint, name string, content []byte) *ProblemDesc {
	return &ProblemDesc{
		ProblemID: pid,
		Name:      name,
		desc:      oss.NewProblemDesc(pid, name, content),
	}
}
