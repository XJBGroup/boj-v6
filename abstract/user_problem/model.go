package user_problem

type UserProblemRelationship struct {

	UserID uint `dorm:"user_id" gorm:"column:user_id;not null" json:"user_id"`
	ProblemID uint `dorm:"problem_id" gorm:"column:problem_id;not null" json:"problem_id"`
}

// TableName specification
func (UserProblemRelationship) TableName() string {
	return "user_problem_rel"
}