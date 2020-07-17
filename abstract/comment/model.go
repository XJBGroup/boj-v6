package comment

import (
	"time"
)

type Comment struct {
	ID            uint `gorm:"primary_key" json:"id"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time `sql:"index"`
	ReferenceType uint8      `gorm:"column:ref_t;not null"`
	Reference     uint       `gorm:"column:ref;not null"`
	ReplyID       uint       `gorm:"column:rid;default:0;not null"`
	Title         string     `gorm:"column:title;default:'Untitled';not null"`
	Content       string     `gorm:"column:content;not null"`
	//Author           User   `gorm:"ForeignKey:AuthorID;AssociationForeignKey:ID"` // one to many created_comments
	AuthorID uint `gorm:"column:author_id" json:"author_id;not null"` // author_id
	//LastUpdateUser   User   `gorm:"ForeignKey:LastUpdateUserID;AssociationForeignKey:ID"`
	LastUpdateUserID uint   `gorm:"column:last_update_user_id;not null" json:"last_update_id"` // last_update_user_id
	IsSticky         bool   `gorm:"column:is_sticky;default:false;not null"`
	IP               string `gorm:"ip"`
}

// TableName specification
func (Comment) TableName() string {
	return "comment"
}

func (a Comment) GetID() uint {
	return a.ID
}
