package user_problem

type UserTriedProblemRelationship struct {
	UserID    uint `dorm:"user_id" gorm:"column:user_id;not null" json:"user_id"`
	ProblemID uint `dorm:"problem_id" gorm:"column:problem_id;not null" json:"problem_id"`
}

// TableName specification
func (UserTriedProblemRelationship) TableName() string {
	return "user_tried_problem_rel"
}

type UserSolvedProblemRelationship struct {
	UserID    uint `dorm:"user_id" gorm:"column:user_id;not null" json:"user_id"`
	ProblemID uint `dorm:"problem_id" gorm:"column:problem_id;not null" json:"problem_id"`
}

// TableName specification
func (UserSolvedProblemRelationship) TableName() string {
	return "user_solved_problem_rel"
}
