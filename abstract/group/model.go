package group

import (
	"time"
)

type Group struct {
	ID          uint `gorm:"primary_key"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time `sql:"index"`
	Name        string     `gorm:"type:varchar(128);column:name;default:'anonymous group';unique;not null" json:"name"`
	Description string     `gorm:"column:description;type:text;not null"`
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
