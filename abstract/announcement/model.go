package announcement

import (
	"time"
)

type Announcement struct {
	ID             uint `gorm:"primary_key" json:"id"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time `sql:"index"`
	Title          string     `gorm:"column:title;default:'Untitled';not null" json:"title"`
	Content        string     `gorm:"column:content;not null" json:"content"`
	Author         uint       `gorm:"column:author_id;not null" json:"author_id"`                // author_id
	LastUpdateUser uint       `gorm:"column:last_update_user_id;not null" json:"last_update_id"` // last_update_user_id
	IsSticky       bool       `gorm:"column:is_sticky;default:false;not null"`
}

// TableName specification
func (Announcement) TableName() string {
	return "announcement"
}

func (a Announcement) GetID() uint {
	return a.ID
}
