package inner_control

import (
	"github.com/Myriad-Dreamin/boj-v6/api"
)

type InnerUserService interface {
	UserServiceSignatureXXX() interface{}
	ListUser(req *api.ListUserRequest) (*api.ListUserReply, error)
	CountUser(req *api.CountUserRequest) (*api.CountUserReply, error)
	Register(req *api.RegisterRequest) (*api.RegisterReply, error)
	LoginUser(req *api.LoginUserRequest) (*api.LoginUserReply, error)
	RefreshToken(req *api.RefreshTokenRequest) (*api.RefreshTokenReply, error)
	BindEmail(req *api.BindEmailRequest) (*api.BindEmailReply, error)
	ChangePassword(req *api.ChangePasswordRequest) (*api.ChangePasswordReply, error)
	InspectUser(req *api.InspectUserRequest) (*api.InspectUserReply, error)
	GetUser(req *api.GetUserRequest) (*api.GetUserReply, error)
	PutUser(req *api.PutUserRequest) (*api.PutUserReply, error)
	DeleteUser(req *api.DeleteUserRequest) (*api.DeleteUserReply, error)
}
