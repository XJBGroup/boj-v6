package problem

import (
	"time"
)

type Problem struct {
	ID             uint       `gorm:"primary_key" json:"id"`
	CreatedAt      time.Time  `json:"created_at" form:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at" form:"updated_at"`
	DeletedAt      *time.Time `sql:"index" json:"deleted_at" form:"deleted_at"`
	Title          string     `dorm:"title" gorm:"type:varchar(50);column:title;default:'Untitled';not null" json:"title"`
	DescriptionRef string     `dorm:"description_ref" gorm:"column:description_ref;type:varchar(100);default:'default';not null" json:"tmpl_name"`
	AuthorID       uint       `dorm:"author_id" gorm:"column:author_id;not null" json:"author_id"`

	// legacy columns
	TimeLimit       int64  `dorm:"time_limit" gorm:"column:time_limit;default:1000;not null" json:"time_limit"`
	MemoryLimit     int64  `dorm:"memory_limit" gorm:"column:memory_limit;default:65536;not null" json:"memory_limit"`
	CodeLengthLimit int64  `dorm:"code_length_limit" gorm:"column:code_length_limit;default:65536;not null" json:"code_length_limit"`
	IsSpj           bool   `dorm:"is_spj" gorm:"column:is_spj;default:false;not null" json:"is_spj"`
	Description     string `dorm:"description" gorm:"column:description;type:varchar(200);not null" json:"description"`

	//Tags            []ProblemTag `gorm:"many2many:problem_problemtags;association_foreignkey:ID;foreignkey:ID;preload:false" json:"-"`
}

// TableName specification
func (Problem) TableName() string {
	return "problem"
}

func (a Problem) GetID() uint {
	return a.ID
}
