package group

import (
	"time"
)

type Group struct {
	ID          uint       `gorm:"primary_key" json:"id"`
	CreatedAt   time.Time  `json:"created_at" form:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at" form:"updated_at"`
	DeletedAt   *time.Time `sql:"index" json:"deleted_at" form:"deleted_at"`
	Name        string     `gorm:"type:varchar(128);column:name;default:'anonymous group';unique;not null" json:"name"`
	Description string     `gorm:"column:description;type:text;not null" json:"description"`
	//Owner       User   `gorm:"ForeignKey:OwnerID;AssociationForeignKey:ID"`
	OwnerID uint `gorm:"column:owner_id;not null" json:"owner_id"`
	//UsersBuffer []User `gorm:"many2many:group_members;association_foreignkey:ID;foreignkey:ID;preload:false"`
}

// TableName specification
func (Group) TableName() string {
	return "group"
}

func (a Group) GetID() uint {
	return a.ID
}
