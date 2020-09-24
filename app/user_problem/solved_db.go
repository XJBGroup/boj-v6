package user_problem

import (
	"errors"
	"github.com/Myriad-Dreamin/boj-v6/abstract/user_problem"
	"github.com/Myriad-Dreamin/boj-v6/app/dao"
	"github.com/Myriad-Dreamin/minimum-lib/module"
	"github.com/jinzhu/gorm"
)

type SolvedDBImpl struct {
	dao.GORMDBImpl

	idleObject user_problem.UserSolvedProblemRelationship
	tableName  string
}

func (D SolvedDBImpl) Create(r *user_problem.UserSolvedProblemRelationship) (int64, error) {
	if r.ProblemID == 0 || r.UserID == 0 {
		return 0, errors.New("invalid zero relationship")
	}
	return D.GORMDBImpl.Create(r)
}

func (D SolvedDBImpl) Delete(r *user_problem.UserSolvedProblemRelationship) (int64, error) {
	if r.ProblemID == 0 || r.UserID == 0 {
		return 0, errors.New("invalid zero relationship")
	}
	return D.GORMDBImpl.Delete(r)
}

func (D SolvedDBImpl) FindProblems(userID uint, page, pageSize int) (c []uint, err error) {
	if userID == 0 {
		return nil, errors.New("invalid zero value query")
	}
	err = D.GORMDBImpl.Page(page, pageSize).Model(D.idleObject).
		Where("user_id = ?", userID).Select("problem_id").Scan(&c).Error
	return
}

func (D SolvedDBImpl) FindUsers(problemID uint, page, pageSize int) (c []uint, err error) {
	if problemID == 0 {
		return nil, errors.New("invalid zero value query")
	}
	err = D.GORMDBImpl.Page(page, pageSize).
		Where("problem_id = ?", problemID).Select("user_id").Scan(&c).Error
	return
}

func (D SolvedDBImpl) Count(r *user_problem.UserSolvedProblemRelationship) (int64, error) {
	if (r.ProblemID | r.UserID) == 0 {
		return 0, errors.New("invalid zero value query")
	}

	if r.ProblemID == 0 {
		return D.GORMDBImpl.CountW(r.TableName(), "user_id = ?", r.UserID)
	} else if r.UserID == 0 {
		return D.GORMDBImpl.CountW(r.TableName(), "problem_id = ?", r.ProblemID)
	} else {
		return D.GORMDBImpl.CountW(r.TableName(), "user_id = ? and problem_id = ?", r.UserID, r.ProblemID)
	}
}

func (D SolvedDBImpl) Migrate() error {
	e := D.GORMDBImpl.Migrate(&D.idleObject)
	if e != nil {
		return e
	}
	return D.GORMDBImpl.DB.Model(
		D.idleObject).AddUniqueIndex("user_solved_problem_pk", "user_id", "problem_id").Error
}

func NewSolvedDB(m module.Module) (*SolvedDBImpl, error) {
	return &SolvedDBImpl{
		GORMDBImpl: dao.NewGORMBasic(m.RequireImpl(new(*gorm.DB)).(*gorm.DB), m),
	}, nil
}
