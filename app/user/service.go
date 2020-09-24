package user

import (
	"github.com/Myriad-Dreamin/boj-v6/abstract/user"
	"github.com/Myriad-Dreamin/boj-v6/api"
	"github.com/Myriad-Dreamin/boj-v6/app/snippet"
	"github.com/Myriad-Dreamin/boj-v6/external"
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
	enforcer   *external.Enforcer
	logger     external.Logger
	key        string
}

func NewService(m module.Module) (*Service, error) {
	s := new(Service)
	s.db = m.RequireImpl(new(user.DB)).(user.DB)
	s.middleware = m.RequireImpl(new(*jwt.Middleware)).(*jwt.Middleware)
	s.enforcer = m.RequireImpl(new(*external.Enforcer)).(*external.Enforcer)
	s.logger = m.RequireImpl(new(external.Logger)).(external.Logger)

	s.key = "id"
	return s, nil
}

func (svc *Service) UserServiceSignatureXXX() interface{} {
	return svc
}

func (svc *Service) ListUser(c controller.MContext) {
	page, pageSize, ok := snippet.RosolvePageVariable(c)
	if !ok {
		return
	}

	users, err := svc.db.Find(page, pageSize)
	if snippet.MaybeSelectError(c, users, err) {
		return
	}

	c.JSON(http.StatusOK, api.SerializeListUserReply(types.CodeOK,
		api.PackSerializeListUserInnerReply(users)))
	return
}

func (svc *Service) CountUser(c controller.MContext) {
	count, err := svc.db.Count()
	if snippet.MaybeCountError(c, err) {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":  types.CodeOK,
		"count": count,
	})
}

func (svc *Service) DoRegister(c controller.MContext) (r *api.RegisterReply) {
	var req = new(api.RegisterRequest)
	if !snippet.BindRequest(c, req) {
		return
	}

	var usr = new(user.User)
	usr.UserName = req.UserName
	usr.NickName = req.NickName
	usr.Gender = req.Gender

	err := svc.db.RecalculatePassword(usr, req.Password)
	if err != nil {
		c.JSON(http.StatusOK, &serial.ErrorSerializer{
			Code:  types.CodeGeneratePasswordError,
			Error: err.Error(),
		})
	}

	// todo: bind email option

	// check default value
	aff, err := svc.db.Create(usr)
	if err != nil {
		if snippet.CheckInsertError(c, err) {
			return
		}
		c.AbortWithStatusJSON(http.StatusOK, &serial.ErrorSerializer{
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
		Data: api.SerializeUserRegisterData(usr),
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

	_, err := svc.enforcer.AddGroupingPolicy("user:"+strconv.Itoa(int(resp.Data.Id)), "admin")
	if err != nil {
		svc.logger.Debug("update group error", "error", err)
	}
}

func (svc *Service) GetUserIdentities(id uint) (identities []string) {

	// todo: all identities
	if svc.enforcer.HasGroupingPolicy("user:"+strconv.Itoa(int(id)), "admin") {
		identities = append(identities, "admin")
	}
	return
}

func (svc *Service) LoginUser(c controller.MContext) {
	var req = new(api.LoginUserRequest)
	if !snippet.BindRequest(c, req) {
		return
	}

	var usr *user.User
	var err error
	if req.Id != 0 {
		usr, err = svc.db.ID(req.Id)
	} else if len(req.UserName) != 0 {
		usr, err = svc.db.QueryUserName(req.UserName)
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
	if snippet.MaybeSelectError(c, usr, err) {
		return
	}

	ok, err := svc.db.AuthenticatePassword(usr, req.Password)
	if !snippet.AuthenticatePassword(c, ok, err) {
		return
	}

	if token, refreshToken, err := svc.middleware.GenerateTokenWithRefreshToken(&types.CustomFields{UID: usr.ID}); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, &serial.ErrorSerializer{
			Code:  types.CodeAuthGenerateTokenError,
			Error: err.Error(),
		})
		return
	} else {
		usr.LastLogin = time.Now()

		c.JSON(http.StatusOK, api.LoginUserReply{
			Code: types.CodeOK,
			Data: api.SerializeUserLoginData(usr, refreshToken, token,
				svc.GetUserIdentities(usr.ID)),
		})

		aff, err := svc.db.UpdateFields(usr, []string{"last_login"})
		if err != nil || aff == 0 {
			svc.logger.Debug("update last login failed", "error", snippet.ConvertErrorToString(err), "affected", aff)
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
		Code: types.CodeOK,
		Data: api.UserRefreshTokenData{Token: newToken},
	})
}

func (svc *Service) BindEmail(c controller.MContext) {
	var req = new(api.BindEmailRequest)
	id, ok := snippet.ParseUintAndBind(c, svc.key, req)
	if !ok {
		return
	}

	var usr = new(user.User)
	usr.ID = id
	usr.Email = req.Email

	// todo: send email

	// check default value
	_, err := svc.db.UpdateFields(usr, []string{"email"})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, &serial.ErrorSerializer{
			Code:  types.CodeUpdateError,
			Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &snippet.ResponseOK)
}

func (svc *Service) ChangePassword(c controller.MContext) {
	var req = new(api.ChangePasswordRequest)
	id, ok := snippet.ParseUintAndBind(c, svc.key, req)
	if !ok {
		return
	}

	var usr = new(user.User)
	usr.ID = id

	verified, err := svc.db.AuthenticatePassword(usr, req.OldPassword)
	if !snippet.AuthenticatePassword(c, verified, err) {
		return
	}

	err = svc.db.RecalculatePassword(usr, req.NewPassword)
	if err != nil {
		c.JSON(http.StatusOK, &serial.ErrorSerializer{
			Code:  types.CodeGeneratePasswordError,
			Error: err.Error(),
		})
	}

	// check default value
	_, err = svc.db.UpdateFields(usr, []string{"password"})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, &serial.ErrorSerializer{
			Code:  types.CodeUpdateError,
			Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &snippet.ResponseOK)
}

func (svc *Service) InspectUser(c controller.MContext) {
	id, ok := snippet.ParseUint(c, svc.key)
	if !ok {
		return
	}
	obj, err := svc.db.ID(id)
	if snippet.MaybeSelectError(c, obj, err) {
		return
	}

	// todo: find all user related problems
	var successProblems []uint
	var triedProblems []uint

	c.JSON(http.StatusOK, api.SerializeInspectUserReply(types.CodeOK,
		api.SerializeInspectUserInnerReply(
			obj, svc.GetUserIdentities(obj.ID), successProblems, triedProblems)))
}

func (svc *Service) GetUser(c controller.MContext) {
	id, ok := snippet.ParseUint(c, svc.key)
	if !ok {
		return
	}
	obj, err := svc.db.ID(id)
	if snippet.MaybeSelectError(c, obj, err) {
		return
	}

	c.JSON(http.StatusOK, api.SerializeGetUserReply(types.CodeOK,
		api.SerializeGetUserInnerReply(obj)))
}

func (svc *Service) PutUser(c controller.MContext) {
	var req = new(api.PutUserRequest)
	id, ok := snippet.ParseUintAndBind(c, svc.key, req)
	if !ok {
		return
	}

	obj, err := svc.db.ID(id)
	if snippet.MaybeSelectError(c, obj, err) {
		return
	}

	_, err = svc.db.UpdateFields(obj, svc.FillPutFields(obj, req))
	if snippet.UpdateFields(c, err) {
		c.JSON(http.StatusOK, &snippet.ResponseOK)
	}
}

func (svc *Service) DeleteUser(c controller.MContext) {
	obj := new(user.User)
	var ok bool
	obj.ID, ok = snippet.ParseUint(c, svc.key)
	if !ok {
		return
	}

	a, e := svc.db.Delete(obj)
	if snippet.DeleteObj(c, a, e) {
		c.JSON(http.StatusOK, &snippet.ResponseOK)
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
