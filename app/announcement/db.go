package announcement

import (
	"github.com/Myriad-Dreamin/boj-v6/abstract/announcement"
	"github.com/Myriad-Dreamin/boj-v6/config"
	"github.com/Myriad-Dreamin/boj-v6/lib/traits"
	"github.com/Myriad-Dreamin/minimum-lib/module"
	"github.com/jinzhu/gorm"
)

func NewDB(m module.Module) (*DBImpl, error) {
	return &DBImpl{
		GORMDBImpl: traits.NewGORMTraits(m.Require(config.ModulePath.DBInstance.GormDB).(*gorm.DB)),
	}, nil
}

type DBImpl struct {
	traits.GORMDBImpl
	idleObject announcement.Announcement
}

func (db DBImpl) Migrate() error {
	return db.GORMDBImpl.Migrate(&db.idleObject)
}

func (db DBImpl) Find(page, pageSize int) (as []announcement.Announcement, err error) {
	return as, db.GORMDBImpl.Find(page, pageSize, &as)
}

func (db DBImpl) Count() (c int64, err error) {
	return db.GORMDBImpl.Count(db.idleObject.TableName())
}

func (db DBImpl) UpdateFields(obj *announcement.Announcement, fields []string) (int64, error) {
	return db.GORMDBImpl.UpdateFields(obj, fields)
}

func (db DBImpl) ID(id uint) (ann *announcement.Announcement, err error) {
	ann = new(announcement.Announcement)
	err = db.GORMDBImpl.ID(id, ann)
	if err == traits.DBErrorNotFound {
		ann = nil
		err = nil
	}
	return
}

func (db DBImpl) Create(obj *announcement.Announcement) (int64, error) {
	return db.GORMDBImpl.Create(obj)
}

func (db DBImpl) Update(obj *announcement.Announcement) (int64, error) {
	return db.GORMDBImpl.Update(obj)
}

func (db DBImpl) Delete(obj *announcement.Announcement) (int64, error) {
	return db.GORMDBImpl.Delete(obj)
}
