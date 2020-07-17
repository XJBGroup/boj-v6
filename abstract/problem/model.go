package problem

import (
	"time"
)

type Problem struct {
	ID              uint `gorm:"primary_key" json:"id"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       *time.Time `sql:"index"`
	Title           string     `dorm:"title" gorm:"type:varchar(50);column:title;default:'Untitled';not null" json:"title"`
	TimeLimit       int64      `dorm:"time_limit" gorm:"column:time_limit;default:1000;not null" json:"time-limit"`
	MemoryLimit     int64      `dorm:"memory_limit" gorm:"column:memory_limit;default:65536;not null" json:"memory-limit"`
	CodeLengthLimit int64      `dorm:"code_length_limit" gorm:"column:code_length_limit;default:65536;not null" json:"code_length-limit"`
	Description     string     `dorm:"description" gorm:"column:description;type:varchar(200);not null" json:"description"`
	DescriptionRef  string     `dorm:"desc_ref_name" gorm:"column:desc_ref_name;type:varchar(100);not null" json:"tmpl_name"`
	//Author          User         `gorm:"ForeignKey:AuthorID;AssociationForeignKey:ID;preload:false" json:"author"` // one to many created_announcements
	AuthorID uint `dorm:"author_id" gorm:"column:author_id;not null" json:"author_id"` // author_id
	IsSpj    bool `dorm:"is_spj" gorm:"column:is_spj;default:false;not null" json:"is-spj"`
	//Tags            []ProblemTag `gorm:"many2many:problem_problemtags;association_foreignkey:ID;foreignkey:ID;preload:false" json:"-"`
}

// TableName specification
func (Problem) TableName() string {
	return "problem"
}

func (a Problem) GetID() uint {
	return a.ID
}
