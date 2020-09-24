package contest

import (
	"time"
)

type Contest struct {
	ID                  uint `gorm:"primary_key" json:"id"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DeletedAt           *time.Time    `sql:"index"`
	ContestType         string        `gorm:"column:contest_type;type:char(5);not null" json:"contest_type"`
	Title               string        `gorm:"type:varchar(128);column:title;default:'Untitled';not null" json:"title"`
	Description         string        `gorm:"column:description;type:text;not null" json:"description"`
	AuthorID            uint          `gorm:"column:author_id;not null" json:"author_id"`
	LastUpdateUserID    uint          `gorm:"column:last_update_user_id;not null" json:"last_update_id"`
	StartAt             *time.Time    `gorm:"column:start_at;not null;default:CURRENT_TIMESTAMP;not null" json:"start_at"`
	EndDuration         time.Duration `gorm:"column:end_duration;not null" json:"end_duration"`
	BoardFrozenDuration time.Duration `gorm:"column:board_frozen_duration;not null" json:"board_frozen_duration"`
	ConfigPath          string        `gorm:"column:config_path;not null" json:"config_path"`
	RolePath            string        `gorm:"column:role_path" json:"role_path"`
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
