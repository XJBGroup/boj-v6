package user

import (
	"github.com/Myriad-Dreamin/boj-v6/abstract/user"
	"github.com/Myriad-Dreamin/boj-v6/app/dao"
	"github.com/Myriad-Dreamin/minimum-lib/crypto"
)

func (db DBImpl) Create(obj *user.User) (int64, error) {
	var err error
	obj.Password, err = crypto.NewPasswordString(obj.Password)
	if err != nil {
		return 0, err
	}

	return db.db.Create(obj)
}

func (db DBImpl) QueryName(name string) (usr *user.User, err error) {
	usr = new(user.User)
	err = db.GORMDBImpl.QueryOne("name = ?", name, usr)
	if err == dao.DBErrorNotFound {
		usr = nil
		err = nil
	}
	return
}

func (db DBImpl) QueryEmail(email string) (usr *user.User, err error) {
	usr = new(user.User)
	err = db.GORMDBImpl.QueryOne("email = ?", email, usr)
	if err == dao.DBErrorNotFound {
		usr = nil
		err = nil
	}
	return
}

func (db DBImpl) AuthenticatePassword(user *user.User, password string) (verified bool, err error) {
	return crypto.CheckPasswordString(password, user.Password)
}
