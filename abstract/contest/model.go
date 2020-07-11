package contest

import (
	"time"
)

type Contest struct {
	ID          uint `gorm:"primary_key"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time `sql:"index"`
	Title       string     `gorm:"type:varchar(128);column:title;default:'Untitled';not null" json:"title"`
	Description string     `gorm:"column:description;type:text;not null"`
	//Author              User          `gorm:"ForeignKey:AuthorID;AssociationForeignKey:ID;preload:false"` // one to many created_Contests
	AuthorID            uint          `gorm:"column:author_id" json:"author_id;not null"` // author_id
	StartAt             *time.Time    `gorm:"column:start_at;not null;default:CURRENT_TIMESTAMP;not null"`
	EndDuration         time.Duration `gorm:"column:end_duration;not null"`
	BoardFrozenDuration time.Duration `gorm:"column:board_frozen_duration;not null"`
	ConfigPath          string        `gorm:"column:config_path;not null"`
	RolePath            string        `gorm:"column:role_path"`
	//UsersBuffer         []User        `gorm:"many2many:contest_users;association_foreignkey:ID;foreignkey:ID;preload:false"`
	//ProblemsBuffer      []Problem     `gorm:"many2many:contest_problems;association_foreignkey:ID;foreignkey:ID;preload:false"`
}

// TableName specification
func (Contest) TableName() string {
	return "contest"
}

func (a Contest) GetID() uint {
	return a.ID
}
