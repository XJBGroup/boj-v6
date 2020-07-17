package submission

import (
	"time"
)

type Submission struct {
	ID           uint `gorm:"primary_key" json:"id"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time `sql:"index"`
	ProblemID    uint       `gorm:"column:problem_id;default:0;not null" json:"problem_id"`
	UserID       uint       `gorm:"column:user_id;not null" form:"user_id" json:"user_id" binding:"required"`
	Score        int64      `gorm:"column:score;default:0;not null" json:"score"`
	Status       int64      `gorm:"column:status;default:0;not null"`
	RunTime      int64      `gorm:"column:running_time;default:0;not null"`
	RunMemory    int64      `gorm:"column:running_memory;default:0;not null"`
	LastTestCase int64      `gorm:"column:last_test_case;default:0;not null"`
	CodeLength   int        `gorm:"column:length;default:0;not null"`
	Information  string     `gorm:"column:info;type:text;not null"`
	Language     uint8      `gorm:"column:language;not null"`
	Hash         string     `gorm:"column:hash;not null"`
	SubmitIP     string     `gorm:"submit_ip;not null"`
	Shared       uint8      `gorm:"shared;default:0;not null"`
}

// TableName specification
func (Submission) TableName() string {
	return "submission"
}

func (a Submission) GetID() uint {
	return a.ID
}
