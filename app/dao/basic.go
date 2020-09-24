package dao

import (
	"errors"
	"github.com/Myriad-Dreamin/boj-v6/types"
	"github.com/Myriad-Dreamin/minimum-lib/module"
	"github.com/jinzhu/gorm"
	"path"
)

type UnwrapErrorFunction = func(error) types.ServiceCode

type GORMDBImpl struct {
	DB          *gorm.DB
	unwrapError UnwrapErrorFunction
}

func NewGORMBasic(db *gorm.DB, m module.Module) GORMDBImpl {
	return GORMDBImpl{DB: db, unwrapError: detectDB(m, db)}
}

func (db *GORMDBImpl) UnwrapError(err error) types.ServiceCode {
	return db.unwrapError(err)
}

func detectDB(m module.Module, db *gorm.DB) UnwrapErrorFunction {
	return m.RequireNamedImpl(path.Join(
		"database_unwrap_functions", db.Dialect().GetName()), new(UnwrapErrorFunction)).(UnwrapErrorFunction)
}

var DBErrorNotFound = errors.New("db error not found")

func (db *GORMDBImpl) Migrate(obj interface{}) error {
	err := db.DB.AutoMigrate(obj).Error
	if err != nil {
		return err
	}

	//db.AddIndex()
	//model, err := db.DormDB.Model(db.ObjectFactory().(dorm.ORMObject))
	//if err != nil {
	//	return err
	//}
	//*db.DormModel = *model
	return nil
}

func (db GORMDBImpl) ID(id uint, obj interface{}) (err error) {
	rdb := db.DB.First(obj, id)
	err = rdb.Error
	if err == nil && rdb.RecordNotFound() {
		return DBErrorNotFound
	}
	return
}

func (db GORMDBImpl) Query(obj interface{}, tmpl string, args ...interface{}) (err error) {
	rdb := db.DB.Where(tmpl, args...).First(obj)
	err = rdb.Error
	if err == nil && rdb.RecordNotFound() {
		return DBErrorNotFound
	}
	return
}

func (db GORMDBImpl) Has(obj interface{}, tmpl string, args ...interface{}) (exists bool, err error) {
	rdb := db.DB.Where(tmpl, args).First(obj)
	err = rdb.Error
	if rdb.RecordNotFound() {
		exists = false
		err = nil
	} else {
		exists = err == nil
	}
	return
}

func (db GORMDBImpl) Create(obj interface{}) (int64, error) {
	rdb := db.DB.Create(obj)
	return rdb.RowsAffected, rdb.Error
}

func (db GORMDBImpl) Update(obj interface{}) (int64, error) {
	rdb := db.DB.Update(obj)
	return rdb.RowsAffected, rdb.Error
}

func (db GORMDBImpl) Delete(obj interface{}) (int64, error) {
	rdb := db.DB.Delete(obj)
	return rdb.RowsAffected, rdb.Error
}

func (db GORMDBImpl) UpdateFields(obj interface{}, fields []string) (int64, error) {
	rdb := db.DB.Model(obj).Select(fields).Updates(obj)
	return rdb.RowsAffected, rdb.Error
}

func (db GORMDBImpl) Count(tb string) (c int64, err error) {
	err = db.DB.Table(tb).Count(&c).Error
	return
}

func (db GORMDBImpl) CountW(tb string, tmpl string, args ...interface{}) (c int64, err error) {
	err = db.DB.Table(tb).Where(tmpl, args...).Count(&c).Error
	return
}

func (db GORMDBImpl) Find(page, pageSize int, obj interface{}) error {
	return db.Page(page, pageSize).Find(obj).Error
}

func (db GORMDBImpl) Page(page, pageSize int) *gorm.DB {
	if pageSize == 0 {
		return db.DB
	}
	return db.DB.Limit(pageSize).Offset((page - 1) * pageSize)
}
