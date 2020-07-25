package problem_desc

import (
	"github.com/Myriad-Dreamin/boj-v6/abstract/problem-desc"
	"github.com/Myriad-Dreamin/boj-v6/app/dao"
	"github.com/Myriad-Dreamin/minimum-lib/module"
	"github.com/jinzhu/gorm"
)

func newDB(m module.Module) (db, error) {
	return db{
		GORMDBImpl: dao.NewGORMBasic(m.RequireImpl(new(*gorm.DB)).(*gorm.DB)),
	}, nil
}

type db struct {
	dao.GORMDBImpl

	idleObject problem_desc.ProblemDesc
}

func (d db) Migrate() error {
	return d.GORMDBImpl.Migrate(&d.idleObject)
}

func (d db) ID(id uint) (usr *problem_desc.ProblemDesc, err error) {
	usr = new(problem_desc.ProblemDesc)
	err = d.GORMDBImpl.ID(id, usr)
	if err == dao.DBErrorNotFound {
		usr = nil
		err = nil
	}
	return
}

func (d db) Create(obj *problem_desc.ProblemDesc) (int64, error) {
	return d.GORMDBImpl.Create(obj)
}

func (d db) Update(obj *problem_desc.ProblemDesc) (int64, error) {
	return d.GORMDBImpl.Update(obj)
}

func (d db) Delete(obj *problem_desc.ProblemDesc) (int64, error) {
	return d.GORMDBImpl.Delete(obj)
}

func (d db) Find(page, pageSize int) (objs []problem_desc.ProblemDesc, err error) {
	return objs, d.GORMDBImpl.Find(page, pageSize, &objs)
}

func (d db) Count() (c int64, err error) {
	return d.GORMDBImpl.Count(d.idleObject.TableName())
}

func (d db) UpdateFields(obj *problem_desc.ProblemDesc, fields []string) (int64, error) {
	return d.GORMDBImpl.UpdateFields(obj, fields)
}
