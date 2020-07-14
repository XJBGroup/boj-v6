package submission

import (
	"github.com/Myriad-Dreamin/boj-v6/abstract/submission"
	"github.com/jinzhu/gorm"
)

func (db *DBImpl) ApplyFilter(f *submission.Filter) *gorm.DB {
	engine := db.Page(f.Page, f.PageSize)

	if len(f.Order) != 0 {
		engine = engine.Order(f.Order, true)
	}
	if f.ByUser != 0 {
		engine = engine.Where("user_id = ?", f.ByUser)
	}
	if f.OnProblem != 0 {
		engine = engine.Where("problem_id = ?", f.OnProblem)
	}
	if f.WithLanguage != 0 {
		engine = engine.Where("language = ?", f.WithLanguage)
	}
	if f.HasStatus != 0 {
		engine = engine.Where("status = ?", f.HasStatus)
	}

	return engine
}

func (db *DBImpl) Filter(f *submission.Filter) (submissions []submission.Submission, _ error) {
	return submissions, db.ApplyFilter(f).Find(&submissions).Error
}

func (db *DBImpl) FilterCount(f *submission.Filter) (cnt int64, _ error) {
	return cnt, db.ApplyFilter(f).Count(&cnt).Error
}

func (db *DBImpl) HasHash(hash string) (exists bool, err error) {
	return db.HasOne("hash = ?", hash, &db.idleObject)
}
