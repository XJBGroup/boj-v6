package tests

import (
	user_problem2 "github.com/Myriad-Dreamin/boj-v6/abstract/user_problem"
	"github.com/Myriad-Dreamin/boj-v6/app/user_problem"
	"testing"
)

func TestUserProblemRelationshipModel(t *testing.T) {
	ctx := srv.Context(t)

	db, err := user_problem.NewTriedDB(ctx.Module)
	ctx.Nil(err)

	aff, err := db.Create(&user_problem2.UserTriedProblemRelationship{
		UserID:    1,
		ProblemID: 1,
	})
	ctx.Nil(err)
	ctx.Equal(int64(1), aff)

	aff, err = db.Create(&user_problem2.UserTriedProblemRelationship{
		UserID:    1,
		ProblemID: 2,
	})
	ctx.Nil(err)
	ctx.Equal(int64(1), aff)

	aff, err = db.Create(&user_problem2.UserTriedProblemRelationship{
		UserID:    1,
		ProblemID: 3,
	})
	ctx.Nil(err)
	ctx.Equal(int64(1), aff)

	aff, err = db.Create(&user_problem2.UserTriedProblemRelationship{
		UserID:    2,
		ProblemID: 2,
	})
	ctx.Nil(err)
	ctx.Equal(int64(1), aff)

	aff, err = db.Create(&user_problem2.UserTriedProblemRelationship{
		UserID:    2,
		ProblemID: 2,
	})
	ctx.NotNil(err)
	ctx.Equal(int64(0), aff)

	cnt, err := db.Count(&user_problem2.UserTriedProblemRelationship{ProblemID: 1})
	ctx.Nil(err)
	ctx.Equal(int64(1), cnt)

	cnt, err = db.Count(&user_problem2.UserTriedProblemRelationship{ProblemID: 2})
	ctx.Nil(err)
	ctx.Equal(int64(2), cnt)

	cnt, err = db.Count(&user_problem2.UserTriedProblemRelationship{ProblemID: 3})
	ctx.Nil(err)
	ctx.Equal(int64(1), cnt)

	cnt, err = db.Count(&user_problem2.UserTriedProblemRelationship{UserID: 1})
	ctx.Nil(err)
	ctx.Equal(int64(3), cnt)

	cnt, err = db.Count(&user_problem2.UserTriedProblemRelationship{UserID: 2, ProblemID: 2})
	ctx.Nil(err)
	ctx.Equal(int64(1), cnt)

	cnt, err = db.Count(&user_problem2.UserTriedProblemRelationship{UserID: 2, ProblemID: 3})
	ctx.Nil(err)
	ctx.Equal(int64(0), cnt)

	solved, err := db.FindProblems(1, 1, 0)
	ctx.Nil(err)
	ctx.Equal(3, len(solved))

	solved, err = db.FindProblems(1, 2, 2)
	ctx.Nil(err)
	ctx.Equal(1, len(solved))
}
