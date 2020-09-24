package user

import (
	"github.com/Myriad-Dreamin/boj-v6/abstract/user"
	"github.com/Myriad-Dreamin/boj-v6/app/dao"
	"github.com/Myriad-Dreamin/minimum-lib/crypto"
)

func (db DBImpl) RecalculatePassword(obj *user.User, password string) (err error) {
	obj.Password, err = crypto.NewPasswordString(password)
	return
}

func (db DBImpl) QueryUserName(name string) (usr *user.User, err error) {
	usr = new(user.User)
	err = db.GORMDBImpl.Query(usr, "user_name = ?", name)
	if err == dao.DBErrorNotFound {
		usr = nil
		err = nil
	}
	return
}

func (db DBImpl) QueryEmail(email string) (usr *user.User, err error) {
	usr = new(user.User)
	err = db.GORMDBImpl.Query(usr, "email = ?", email)
	if err == dao.DBErrorNotFound {
		usr = nil
		err = nil
	}
	return
}

func (db DBImpl) AuthenticatePassword(user *user.User, password string) (verified bool, err error) {
	return crypto.CheckPasswordString(password, user.Password)
}
