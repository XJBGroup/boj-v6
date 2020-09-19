package user_problem

type DB interface {
	Create(r *UserProblemRelationship) (int64, error)
	Delete(r *UserProblemRelationship) (int64, error)
	FindProblems(userID uint, page, pageSize int) ([]uint, error)
	FindUsers(problemID uint, page, pageSize int) ([]uint, error)

	// if r.UserID == 0 then Count == CountProblem
	// if r.ProblemID == 0 then Count == CountUser
	// if both are not zero then Count == QueryIfUserSolvedProblem
	Count(r *UserProblemRelationship) (int64, error)
}
