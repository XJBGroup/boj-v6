package user

import (
	"github.com/Myriad-Dreamin/boj-v6/abstract/user"
	"github.com/Myriad-Dreamin/boj-v6/config"
	"github.com/Myriad-Dreamin/boj-v6/lib/traits"
	"github.com/Myriad-Dreamin/minimum-lib/crypto"
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

	idleObject user.User
}

func (db DBImpl) Migrate() error {
	return db.GORMDBImpl.Migrate(&db.idleObject)
}

func (db DBImpl) Find(page, pageSize int) (as []user.User, err error) {
	return as, db.GORMDBImpl.Find(page, pageSize, &as)
}

func (db DBImpl) Count() (c int64, err error) {
	return db.GORMDBImpl.Count(db.idleObject.TableName())
}

func (db DBImpl) UpdateFields(obj *user.User, fields []string) (int64, error) {
	return db.GORMDBImpl.UpdateFields(obj, fields)
}

func (db DBImpl) ID(id uint) (usr *user.User, err error) {
	usr = new(user.User)
	err = db.GORMDBImpl.ID(id, usr)
	if err == traits.DBErrorNotFound {
		usr = nil
		err = nil
	}
	return
}

func (db DBImpl) QueryName(name string) (usr *user.User, err error) {
	usr = new(user.User)
	err = db.GORMDBImpl.QueryOne("name = ?", name, usr)
	if err == traits.DBErrorNotFound {
		usr = nil
		err = nil
	}
	return
}

func (db DBImpl) QueryEmail(email string) (usr *user.User, err error) {
	usr = new(user.User)
	err = db.GORMDBImpl.QueryOne("email = ?", email, usr)
	if err == traits.DBErrorNotFound {
		usr = nil
		err = nil
	}
	return
}

func (db DBImpl) AuthenticatePassword(user *user.User, password string) (verified bool, err error) {
	return crypto.CheckPasswordString(password, user.Password)
}

func (db DBImpl) Create(obj *user.User) (int64, error) {
	var err error
	obj.Password, err = crypto.NewPasswordString(obj.Password)
	if err != nil {
		return 0, err
	}

	return db.GORMDBImpl.Create(obj)
}

func (db DBImpl) Update(obj *user.User) (int64, error) {
	return db.GORMDBImpl.Update(obj)
}

func (db DBImpl) Delete(obj *user.User) (int64, error) {
	return db.GORMDBImpl.Delete(obj)
}
