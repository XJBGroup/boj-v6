package user

import (
	"errors"
	"github.com/Myriad-Dreamin/boj-v6/abstract/user"
	"github.com/Myriad-Dreamin/boj-v6/config"
	"github.com/Myriad-Dreamin/minimum-lib/module"
	"github.com/jinzhu/gorm"
)

func NewDB(m module.Module) (*DBImpl, error) {
	return &DBImpl{
		GORMDBImpl{
			m.Require(config.ModulePath.DBInstance.GormDB).(*gorm.DB)},
	}, nil
}

type GORMDBImpl struct {
	*gorm.DB
}

var DBErrorNotFound = errors.New("db error not found")

func (db GORMDBImpl) ID(id uint, obj interface{}) (err error) {
	rdb := db.First(obj, id)
	err = rdb.Error
	if err == nil && rdb.RecordNotFound() {
		return DBErrorNotFound
	}
	return
}

func (db GORMDBImpl) Create(obj *user.User) (int64, error) {
	rdb := db.DB.Create(obj)
	return rdb.RowsAffected, rdb.Error
}

func (db GORMDBImpl) Update(obj *user.User) (int64, error) {
	rdb := db.DB.Update(obj)
	return rdb.RowsAffected, rdb.Error
}

func (db GORMDBImpl) Delete(obj *user.User) (int64, error) {
	rdb := db.DB.Delete(obj)
	return rdb.RowsAffected, rdb.Error
}

type DBImpl struct {
	GORMDBImpl
}

func (db DBImpl) UpdateFields(obj *user.User, fields []string) (int64, error) {
	rdb := db.Model(obj)
	for _, field := range fields {
		rdb = rdb.Select(field)
	}
	rdb = rdb.Updates(obj)
	return rdb.RowsAffected, rdb.Error
}

func (db DBImpl) ID(id uint) (ann *user.User, err error) {
	ann = new(user.User)
	err = db.GORMDBImpl.ID(id, ann)
	if err == DBErrorNotFound {
		ann = nil
		err = nil
	}
	return
}

func (db DBImpl) Create(obj *user.User) (int64, error) {
	return db.GORMDBImpl.Create(obj)
}

func (db DBImpl) Update(obj *user.User) (int64, error) {
	return db.GORMDBImpl.Update(obj)
}

func (db DBImpl) Delete(obj *user.User) (int64, error) {
	return db.GORMDBImpl.Delete(obj)
}
