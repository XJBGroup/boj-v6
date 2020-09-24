package user_problem

import "github.com/Myriad-Dreamin/boj-v6/abstract/db"

type TriedDB interface {
	db.BasicDB

	Create(r *UserTriedProblemRelationship) (int64, error)
	Delete(r *UserTriedProblemRelationship) (int64, error)
	FindProblems(userID uint, page, pageSize int) ([]uint, error)
	FindUsers(problemID uint, page, pageSize int) ([]uint, error)

	// if r.UserID == 0 then Count == CountProblem
	// if r.ProblemID == 0 then Count == CountUser
	// if both are not zero then Count == QueryIfUserSolvedProblem
	Count(r *UserTriedProblemRelationship) (int64, error)
}

type SolvedDB interface {
	db.BasicDB

	Create(r *UserSolvedProblemRelationship) (int64, error)
	Delete(r *UserSolvedProblemRelationship) (int64, error)
	FindProblems(userID uint, page, pageSize int) ([]uint, error)
	FindUsers(problemID uint, page, pageSize int) ([]uint, error)

	// if r.UserID == 0 then Count == CountProblem
	// if r.ProblemID == 0 then Count == CountUser
	// if both are not zero then Count == QueryIfUserSolvedProblem
	Count(r *UserSolvedProblemRelationship) (int64, error)
}
