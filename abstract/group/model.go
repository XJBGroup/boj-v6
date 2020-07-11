package group

import (
	"time"
)

type Group struct {
	ID             uint `gorm:"primary_key"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time `sql:"index"`
	Title          string     `gorm:"column:title;default:'Untitled';not null" json:"title"`
	Content        string     `gorm:"column:content;not null" json:"content"`
	Author         uint       `gorm:"ForeignKey:AuthorID;AssociationForeignKey:ID;auto_preload:true"` // one to many created_announcements
	LastUpdateUser uint       `gorm:"ForeignKey:LastUpdateUserID;AssociationForeignKey:ID;auto_preload:true"`
	IsSticky       bool       `gorm:"column:is_sticky;default:false;not null"`
}

// TableName specification
func (Group) TableName() string {
	return "group"
}

func (a Group) GetID() uint {
	return a.ID
}
