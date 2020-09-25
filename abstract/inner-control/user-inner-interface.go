package inner_control

import (
	"github.com/Myriad-Dreamin/boj-v6/api"
	"github.com/Myriad-Dreamin/minimum-lib/controller"
)

type InnerUserService interface {
	UserServiceSignatureXXX() interface{}
	ListUser(c controller.MContext, req *api.ListUserRequest) (*api.ListUserReply, error)
	CountUser(c controller.MContext, req *api.CountUserRequest) (*api.CountUserReply, error)
	Register(c controller.MContext, req *api.RegisterRequest) (*api.RegisterReply, error)
	LoginUser(c controller.MContext, req *api.LoginUserRequest) (*api.LoginUserReply, error)
	RefreshToken(c controller.MContext, req *api.RefreshTokenRequest) (*api.RefreshTokenReply, error)
	BindEmail(c controller.MContext, req *api.BindEmailRequest) (*api.BindEmailReply, error)
	ChangePassword(c controller.MContext, req *api.ChangePasswordRequest) (*api.ChangePasswordReply, error)
	InspectUser(c controller.MContext, req *api.InspectUserRequest) (*api.InspectUserReply, error)
	GetUser(c controller.MContext, req *api.GetUserRequest) (*api.GetUserReply, error)
	PutUser(c controller.MContext, req *api.PutUserRequest) (*api.PutUserReply, error)
	DeleteUser(c controller.MContext, req *api.DeleteUserRequest) (*api.DeleteUserReply, error)
}
