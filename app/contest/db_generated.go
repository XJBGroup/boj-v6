package contest

import (
	"github.com/Myriad-Dreamin/boj-v6/abstract/contest"
	"github.com/Myriad-Dreamin/boj-v6/app/dao"
	"github.com/Myriad-Dreamin/minimum-lib/module"
	"github.com/jinzhu/gorm"
)

func NewDB(m module.Module) (*DBImpl, error) {
	d, e := newDB(m)
	return &DBImpl{
		db: d,
	}, e
}

type DBImpl struct {
	db
}

func newDB(m module.Module) (db, error) {
	return db{
		GORMDBImpl: dao.NewGORMBasic(m.RequireImpl(new(*gorm.DB)).(*gorm.DB), m),
	}, nil
}

type db struct {
	dao.GORMDBImpl

	idleObject contest.Contest
}

func (d db) Migrate() error {
	return d.GORMDBImpl.Migrate(&d.idleObject)
}

func (d db) ID(id uint) (usr *contest.Contest, err error) {
	usr = new(contest.Contest)
	err = d.GORMDBImpl.ID(id, usr)
	if err == dao.DBErrorNotFound {
		usr = nil
		err = nil
	}
	return
}

func (d db) Create(obj *contest.Contest) (int64, error) {
	return d.GORMDBImpl.Create(obj)
}

func (d db) Update(obj *contest.Contest) (int64, error) {
	return d.GORMDBImpl.Update(obj)
}

func (d db) Delete(obj *contest.Contest) (int64, error) {
	return d.GORMDBImpl.Delete(obj)
}

func (d db) Find(page, pageSize int) (objs []contest.Contest, err error) {
	return objs, d.GORMDBImpl.Find(page, pageSize, &objs)
}

func (d db) Count() (c int64, err error) {
	return d.GORMDBImpl.Count(d.idleObject.TableName())
}

func (d db) UpdateFields(obj *contest.Contest, fields []string) (int64, error) {
	return d.GORMDBImpl.UpdateFields(obj, fields)
}
