package user

import "time"

type User struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at" form:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" form:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at" form:"deleted_at"`

	UserName string `dorm:"user_name" gorm:"type:varchar(30);column:user_name;not null;unique" json:"user_name"` // todo: regex
	NickName string `dorm:"nick_name" gorm:"type:varchar(30);column:nick_name;not null;unique" json:"nick_name"` // todo: regex

	Email    string `dorm:"email" gorm:"column:email;unique;default:NULL" json:"email" binding:"email"` // todo: email
	Motto    string `dorm:"motto" gorm:"column:motto" json:"motto"`
	Password string `dorm:"password" gorm:"type:varchar(128);column:password" json:"-"`
	Gender   uint8  `dorm:"gender" gorm:"type:varchar(128);column:gender" json:"gender"`

	LastLogin           time.Time `dorm:"last_login" gorm:"column:last_login;default:CURRENT_TIMESTAMP" json:"last_login"`
	SolvedProblemsCount int64     `dorm:"solved_problems" gorm:"column:solved_problems" json:"solved_problems"`
	TriedProblemsCount  int64     `dorm:"tried_problems" gorm:"column:tried_problems" json:"tried_problems"`
}

// TableName specification
func (User) TableName() string {
	return "user"
}

func (d User) GetID() uint {
	return d.ID
}
