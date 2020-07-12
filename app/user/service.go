package user

import (
	"github.com/Myriad-Dreamin/boj-v6/abstract/user"
	"github.com/Myriad-Dreamin/boj-v6/api"
	"github.com/Myriad-Dreamin/boj-v6/app/provider"
	"github.com/Myriad-Dreamin/boj-v6/config"
	"github.com/Myriad-Dreamin/boj-v6/external"
	ginhelper "github.com/Myriad-Dreamin/boj-v6/lib/gin-helper"
	"github.com/Myriad-Dreamin/boj-v6/lib/jwt"
	"github.com/Myriad-Dreamin/boj-v6/lib/serial"
	"github.com/Myriad-Dreamin/boj-v6/types"
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	"github.com/Myriad-Dreamin/minimum-lib/module"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type Service struct {
	db         user.DB
	middleware *jwt.Middleware
	enforcer   *provider.Enforcer
	logger     external.Logger
	key        string
}

func NewService(m module.Module) (*Service, error) {
	s := new(Service)
	s.db = m.Require(config.ModulePath.Provider.Model).(*provider.DB).UserDB()
	s.middleware = m.Require(config.ModulePath.Middleware.JWT).(*jwt.Middleware)
	s.enforcer = m.Require(config.ModulePath.Provider.Model).(*provider.DB).Enforcer()
	s.logger = m.Require(config.ModulePath.Global.Logger).(external.Logger)

	s.key = "id"
	return s, nil
}

func (svc *Service) UserServiceSignatureXXX() interface{} {
	return svc
}

func (svc *Service) ListUsers(c controller.MContext) {
	page, pageSize, ok := ginhelper.RosolvePageVariable(c)
	if !ok {
		return
	}

	users, err := svc.db.Find(page, pageSize)
	if ginhelper.MaybeSelectError(c, users, err) {
		return
	}

	c.JSON(http.StatusOK, api.SerializeListUsersReply(types.CodeOK,
		api.PackSerializeListUserReply(users)))
	return
}

func (svc *Service) CountUser(c controller.MContext) {
	count, err := svc.db.Count()
	if ginhelper.MaybeCountError(c, err) {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":  types.CodeOK,
		"count": count,
	})
}

func (svc *Service) DoRegister(c controller.MContext) (r *api.RegisterReply) {
	var req = new(api.RegisterRequest)
	if !ginhelper.BindRequest(c, req) {
		return
	}

	var usr = new(user.User)
	usr.UserName = req.UserName
	usr.Password = req.Password
	usr.NickName = req.NickName
	usr.Gender = req.Gender

	// check default value
	aff, err := svc.db.Create(usr)
	if err != nil {
		if ginhelper.CheckInsertError(c, err) {
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, &serial.ErrorSerializer{
			Code:  types.CodeInsertError,
			Error: err.Error(),
		})
		return
	} else if aff == 0 {
		c.JSON(http.StatusOK, &serial.ErrorSerializer{
			Code:  types.CodeInsertError,
			Error: "existed",
		})
		return
	}

	r = &api.RegisterReply{
		Code: types.CodeOK,
		Id:   usr.ID,
	}
	c.JSON(http.StatusOK, r)
	return
}

func (svc *Service) Register(c controller.MContext) {
	svc.DoRegister(c)
}

func (svc *Service) RegisterAdmin(c controller.MContext) {
	resp := svc.DoRegister(c)
	if resp == nil {
		return
	}

	_, err := svc.enforcer.AddGroupingPolicy("user:"+strconv.Itoa(int(resp.Id)), "admin")
	if err != nil {
		svc.logger.Debug("update group error", "error", err)
	}
}

func (svc *Service) LoginUser(c controller.MContext) {
	var req = new(api.LoginUserRequest)
	if !ginhelper.BindRequest(c, req) {
		return
	}

	var usr *user.User
	var err error
	if req.Id != 0 {
		usr, err = svc.db.ID(req.Id)
	} else if len(req.UserName) != 0 {
		usr, err = svc.db.QueryName(req.UserName)
	} else if len(req.Email) != 0 {
		usr, err = svc.db.QueryEmail(req.Email)
		//} else if len(req.Phone) != 0 {
		//	user, err = us.db.QueryPhone(req.Phone)
	} else {
		c.JSON(http.StatusOK, &serial.ErrorSerializer{
			Code: types.CodeUserIDMissing,
		})
		return
	}
	if ginhelper.MaybeSelectError(c, usr, err) {
		return
	}

	ok, err := svc.db.AuthenticatePassword(usr, req.Password)
	if !ginhelper.AuthenticatePassword(c, ok, err) {
		return
	}

	if token, refreshToken, err := svc.middleware.GenerateTokenWithRefreshToken(&types.CustomFields{UID: usr.ID}); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &serial.ErrorSerializer{
			Code:  types.CodeAuthGenerateTokenError,
			Error: err.Error(),
		})
		return
	} else {
		usr.LastLogin = time.Now()

		var identities []string
		if svc.enforcer.HasGroupingPolicy("user:"+strconv.Itoa(int(usr.ID)), "admin") {
			identities = append(identities, "admin")
		}

		c.JSON(http.StatusOK, api.LoginUserReply{
			Code:         types.CodeOK,
			User:         usr,
			RefreshToken: refreshToken,
			Token:        token,
			Identities:   identities,
		})

		aff, err := svc.db.UpdateFields(usr, []string{"last_login"})
		if err != nil || aff == 0 {
			svc.logger.Debug("update last login failed", "error", err, "affected", aff)
		}

		return
	}
}

func (svc *Service) RefreshToken(c controller.MContext) {
	newToken, err := svc.middleware.RefreshToken(c)
	if err != nil {
		_ = c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	c.JSON(http.StatusOK, api.RefreshTokenReply{
		Code:  types.CodeOK,
		Token: newToken,
	})
}

func (svc *Service) BindEmail(c controller.MContext) {
	var req = new(api.BindEmailRequest)
	id, ok := ginhelper.ParseUintAndBind(c, svc.key, req)
	if !ok {
		return
	}

	var usr = new(user.User)
	usr.ID = id
	usr.Email = req.Email

	// check default value
	_, err := svc.db.UpdateFields(usr, []string{"email"})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &serial.ErrorSerializer{
			Code:  types.CodeUpdateError,
			Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &ginhelper.ResponseOK)
}

func (svc *Service) InspectUser(c controller.MContext) {
	// todo
	panic("implement me")
}

func (svc *Service) GetUser(c controller.MContext) {
	id, ok := ginhelper.ParseUint(c, svc.key)
	if !ok {
		return
	}
	obj, err := svc.db.ID(id)
	if ginhelper.MaybeSelectError(c, obj, err) {
		return
	}

	c.JSON(http.StatusOK, api.GetUserReply{Code: types.CodeOK, User: obj})
}

func (svc *Service) PutUser(c controller.MContext) {
	var req = new(api.PutUserRequest)
	id, ok := ginhelper.ParseUintAndBind(c, svc.key, req)
	if !ok {
		return
	}

	obj, err := svc.db.ID(id)
	if ginhelper.MaybeSelectError(c, obj, err) {
		return
	}

	_, err = svc.db.UpdateFields(obj, svc.FillPutFields(obj, req))
	if ginhelper.UpdateFields(c, err) {
		c.JSON(http.StatusOK, &ginhelper.ResponseOK)
	}
}

func (svc *Service) Delete(c controller.MContext) {
	obj := new(user.User)
	var ok bool
	obj.ID, ok = ginhelper.ParseUint(c, svc.key)
	if !ok {
		return
	}

	a, e := svc.db.Delete(obj)
	if ginhelper.DeleteObj(c, a, e) {
		c.JSON(http.StatusOK, &ginhelper.ResponseOK)
	}
}

func (svc *Service) FillPutFields(obj *user.User, req *api.PutUserRequest) (fields []string) {
	if req.Gender != 255 {
		obj.Motto = req.Motto
		fields = append(fields, "gender")
	}

	if len(req.NickName) != 0 {
		obj.NickName = req.NickName
		fields = append(fields, "nick_name")
	}

	if len(req.Motto) != 0 {
		obj.Motto = req.Motto
		fields = append(fields, "motto")
	}

	return
}
