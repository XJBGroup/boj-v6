package announcement

import (
	"time"
)

type Announcement struct {
	ID             uint       `gorm:"primary_key" json:"id"`
	CreatedAt      time.Time  `json:"created_at" form:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at" form:"updated_at"`
	DeletedAt      *time.Time `sql:"index" json:"deleted_at" form:"deleted_at"`
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
