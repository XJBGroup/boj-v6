package comment

import (
	"github.com/Myriad-Dreamin/boj-v6/abstract/comment"
	"github.com/jinzhu/gorm"
)

func (db *DBImpl) ApplyFilter(f *comment.Filter) *gorm.DB {
	engine := db.Page(f.Page, f.PageSize)

	if f.Ref != 0 {
		engine = engine.Where("ref_t = ? and ref = ?", f.RefType, f.Ref)
	}

	if f.NoReply {
		engine = engine.Where("rid = 0")
	}

	return engine
}

func (db *DBImpl) Filter(f *comment.Filter) (comments []comment.Comment, _ error) {
	return comments, db.ApplyFilter(f).Find(&comments).Error
}

func (db *DBImpl) FilterCount(f *comment.Filter) (cnt int64, _ error) {
	return cnt, db.ApplyFilter(f).Count(&cnt).Error
}
