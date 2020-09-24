package main

import (
	"github.com/Myriad-Dreamin/artisan"
	"github.com/Myriad-Dreamin/boj-v6/abstract/user"
)

type UserCategories struct {
	artisan.VirtualService
	List  artisan.Category
	Count artisan.Category
	//ListNameLike artisan.Category
	Register     artisan.Category
	Login        artisan.Category
	RefreshToken artisan.Category
	GetContent   artisan.Category
	Inspect      artisan.Category
	IdGroup      artisan.Category
}

func DescribeUserService() artisan.ProposingService {
	var userModel = new(user.User)
	var _valueUserModel user.User

	var listParams = []interface{}{
		artisan.Param("page", artisan.Int),
		artisan.Param("page_size", artisan.Int),
	}

	var userFilter = func(name string) artisan.SerializeObject {
		return artisan.Object(
			append(listParams, name)...)
	}

	svc := &UserCategories{
		List: artisan.Ink().
			Path("user-list").
			Method(artisan.GET, "ListUser",
				artisan.Request(userFilter("ListUserRequest")),
				artisan.Reply(
					codeField,
					artisan.ArrayParam(artisan.Param("data",
						artisan.Object("ListUserInnerReply", artisan.SPsC(
							&_valueUserModel.ID,
							&_valueUserModel.Gender,
							&_valueUserModel.LastLogin,
							&_valueUserModel.UserName,
							&_valueUserModel.NickName,
							&_valueUserModel.Email,
							&_valueUserModel.Motto,
							&_valueUserModel.SolvedProblemsCount,
							&_valueUserModel.TriedProblemsCount,
						)))),
				),
			),
		Count: artisan.Ink().
			Path("user-count").
			Method(artisan.GET, "CountUser",
				artisan.Request(userFilter("CountUserRequest")),
				artisan.Reply(
					codeField,
					artisan.Param("data", artisan.Int64),
				),
			),
		Register: artisan.Ink().
			Path("user/register").
			Method(artisan.POST, "Register",
				artisan.Request(
					// UserName: 注册用户的名字
					artisan.SnakeParam(&userModel.UserName, required),
					// Password: 密码
					artisan.SnakeParam(&userModel.Password, required),
					// NickName: 昵称
					artisan.SnakeParam(&userModel.NickName, required),
					// Gender: 0表示保密, 1表示女, 2表示男, 3~255表示其他
					artisan.SnakeParam(&userModel.Gender),
				),
				StdReply(artisan.Object(
					"UserRegisterData",
					artisan.SnakeParam(&userModel.ID),
				)),
			),
		Login: artisan.Ink().
			Path("user/login").
			Method(artisan.POST, "LoginUser",
				artisan.Request(
					artisan.SPsC(
						&userModel.ID, &userModel.UserName, &userModel.Email,
					),
					artisan.SnakeParam(&userModel.Password, required),
				),
				StdReply(artisan.Object(
					"UserLoginData",
					artisan.SnakeParam(&userModel.ID),
					artisan.SnakeParam(&userModel.Email),
					artisan.SnakeParam(&userModel.UserName),
					artisan.SnakeParam(&userModel.NickName),
					artisan.Param("refresh_token", artisan.String),
					artisan.Param("token", artisan.String),
					artisan.Param("identities", artisan.Strings),
				)),
			),
		RefreshToken: artisan.Ink().
			Path("user-token").
			Method(artisan.GET, "RefreshToken",
				artisan.Request(),
				StdReply(artisan.Object(
					"UserRefreshTokenData",
					artisan.Param("token", artisan.String),
				)),
			),

		IdGroup: artisan.Ink().
			Path("user/:id").Meta(&Meta{artisan.RouterMeta{
			RuntimeRouterMeta: "user:id",
		}}).
			Method(artisan.GET, "GetUser",
				artisan.Request(),
				artisan.Reply(
					codeField,
					artisan.Param("data", artisan.Object("GetUserInnerReply",
						artisan.SPsC(
							&userModel.ID, &userModel.NickName,
							&userModel.LastLogin, &userModel.Motto,
							&userModel.Gender),
					),
					))).
			Method(artisan.PUT, "PutUser",
				artisan.Request(
					artisan.SPsC(
						// Gender: 0表示保密, 1表示女, 2表示男, 255表示不修改
						&userModel.Gender, &userModel.NickName, &userModel.Motto)),
				artisan.Reply(codeField),
			).
			SubCate("/email", artisan.Ink().WithName("Email").
				Method(artisan.PUT, "BindEmail",
					artisan.Request(
						// Email: 邮箱
						artisan.SnakeParam(&userModel.Email, artisan.Tag("binding", "email"))),
					artisan.Reply(codeField),
				),
			).
			SubCate("/password", artisan.Ink().WithName("ChangePassword").
				Method(artisan.PUT, "ChangePassword",
					artisan.Request(
						// Old Password: 旧密码
						artisan.Param("old_password", artisan.String, required),
						// New Password: 新密码
						artisan.Param("new_password", artisan.String, required),
					),
					artisan.Reply(codeField),
				),
			).
			SubCate("/inspect", artisan.Ink().WithName("Inspect").
				Method(artisan.GET, "InspectUser",
					artisan.Request(),
					artisan.Reply(
						codeField,
						artisan.Param("data", artisan.Object("InspectUserInnerReply",
							artisan.SPsC(
								&userModel.ID, &userModel.NickName, &userModel.UserName,
								&userModel.LastLogin, &userModel.Email, &userModel.Motto,
								&userModel.Gender),
							artisan.Param("identities", artisan.Strings),
							artisan.Param("success_problems", new([]uint)),
							artisan.Param("tried_problems", new([]uint)),
						),
						),
					)),
			).
			Method(artisan.DELETE, "DeleteUser",
				artisan.Request(),
				artisan.Reply(codeField),
			),
	}
	svc.Name("UserService").
		UseModel(artisan.Model(artisan.Name("user"), &userModel),
			artisan.Model(artisan.Name("valueUser"), &_valueUserModel))
	return svc
}
