package submission

import (
	"time"
)

type Submission struct {
	ID           uint       `gorm:"primary_key" json:"id"`
	CreatedAt    time.Time  `json:"created_at" form:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at" form:"updated_at"`
	DeletedAt    *time.Time `sql:"index" json:"deleted_at" form:"deleted_at"`
	ProblemID    uint       `gorm:"column:problem_id;default:0;not null;force" json:"problem_id"`
	UserID       uint       `gorm:"column:user_id;not null;force" form:"user_id" json:"user_id" binding:"required"`
	Score        int64      `gorm:"column:score;default:0;not null;force" json:"score"`
	Status       int64      `gorm:"column:status;default:0;not null;force"`
	RunTime      int64      `gorm:"column:running_time;default:0;not null;force"`
	RunMemory    int64      `gorm:"column:running_memory;default:0;not null;force"`
	LastTestCase int64      `gorm:"column:last_test_case;default:0;not null;force"`
	CodeLength   int        `gorm:"column:length;default:0;not null;force"`
	Information  string     `gorm:"column:info;type:text;not null;force"`
	Language     uint8      `gorm:"column:language;not null;force"`
	Hash         string     `gorm:"column:hash;not null;force"`
	SubmitIP     string     `gorm:"submit_ip;not null;force"`
	Shared       uint8      `gorm:"shared;default:0;not null;force"`
}

// TableName specification
func (Submission) TableName() string {
	return "submission"
}

func (a Submission) GetID() uint {
	return a.ID
}
