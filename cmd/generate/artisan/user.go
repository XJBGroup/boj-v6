package main

import (
	"github.com/Myriad-Dreamin/artisan/artisan-core"
	"github.com/Myriad-Dreamin/boj-v6/abstract/user"
)

type UserCategories struct {
	artisan_core.VirtualService
	List  artisan_core.Category
	Count artisan_core.Category
	//ListNameLike artisan_core.Category
	Register     artisan_core.Category
	Login        artisan_core.Category
	RefreshToken artisan_core.Category
	GetContent   artisan_core.Category
	Inspect      artisan_core.Category
	IdGroup      artisan_core.Category
}

func DescribeUserController() artisan_core.ProposingService {
	var userModel = new(user.User)
	var _valueUserModel user.User

	var listParams = []interface{}{
		artisan_core.Param("page", artisan_core.Int),
		artisan_core.Param("page_size", artisan_core.Int),
	}

	var userFilter = func(name string) artisan_core.SerializeObject {
		return artisan_core.Object(
			append(listParams, name)...)
	}

	controller := &UserCategories{
		List: artisan_core.Ink().
			Path("user-list").
			Method(artisan_core.GET, "ListUser",
				artisan_core.Request(userFilter("ListUserRequest")),
				artisan_core.Reply(
					codeField,
					artisan_core.ArrayParam(artisan_core.Param("data",
						artisan_core.Object("ListUserInnerReply", artisan_core.SPsC(
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
		Count: artisan_core.Ink().
			Path("user-count").
			Method(artisan_core.GET, "CountUser",
				artisan_core.Request(userFilter("CountUserRequest")),
				artisan_core.Reply(
					codeField,
					artisan_core.Param("data", artisan_core.Int64),
				),
			),
		Register: artisan_core.Ink().
			Path("user/register").
			Method(artisan_core.POST, "Register",
				artisan_core.Request(
					// UserName: 注册用户的名字
					artisan_core.SnakeParam(&userModel.UserName, required),
					// Password: 密码
					artisan_core.SnakeParam(&userModel.Password, required),
					// NickName: 昵称
					artisan_core.SnakeParam(&userModel.NickName, required),
					// Gender: 0表示保密, 1表示女, 2表示男, 3~255表示其他
					artisan_core.SnakeParam(&userModel.Gender),
				),
				StdReply(artisan_core.Object(
					"UserRegisterData",
					artisan_core.SnakeParam(&userModel.ID),
				)),
			),
		Login: artisan_core.Ink().
			Path("user/login").
			Method(artisan_core.POST, "LoginUser",
				artisan_core.Request(
					artisan_core.SPsC(
						&userModel.ID, &userModel.UserName, &userModel.Email,
					),
					artisan_core.SnakeParam(&userModel.Password, required),
				),
				StdReply(artisan_core.Object(
					"UserLoginData",
					artisan_core.SnakeParam(&userModel.ID),
					artisan_core.SnakeParam(&userModel.Email),
					artisan_core.SnakeParam(&userModel.UserName),
					artisan_core.SnakeParam(&userModel.NickName),
					artisan_core.Param("refresh_token", artisan_core.String),
					artisan_core.Param("token", artisan_core.String),
					artisan_core.Param("identities", artisan_core.Strings),
				)),
			),
		RefreshToken: artisan_core.Ink().
			Path("user-token").
			Method(artisan_core.GET, "RefreshToken",
				artisan_core.Request(),
				StdReply(artisan_core.Object(
					"UserRefreshTokenData",
					artisan_core.Param("token", artisan_core.String),
				)),
			),

		IdGroup: artisan_core.Ink().
			Path("user/:id").Meta(&Meta{artisan_core.RouterMeta{
			RuntimeRouterMeta: "user:id",
		}}).
			Method(artisan_core.GET, "GetUser",
				artisan_core.Request(),
				artisan_core.Reply(
					codeField,
					artisan_core.Param("data", artisan_core.Object("GetUserInnerReply",
						artisan_core.SPsC(
							&userModel.ID, &userModel.NickName,
							&userModel.LastLogin, &userModel.Motto,
							&userModel.Gender),
					),
					))).
			Method(artisan_core.PUT, "PutUser",
				artisan_core.Request(
					artisan_core.SPsC(
						// Gender: 0表示保密, 1表示女, 2表示男, 255表示不修改
						&userModel.Gender, &userModel.NickName, &userModel.Motto)),
				artisan_core.Reply(codeField),
			).
			SubCate("/email", artisan_core.Ink().WithName("Email").
				Method(artisan_core.PUT, "BindEmail",
					artisan_core.Request(
						// Email: 邮箱
						artisan_core.SnakeParam(&userModel.Email, artisan_core.Tag("binding", "email"))),
					artisan_core.Reply(codeField),
				),
			).
			SubCate("/password", artisan_core.Ink().WithName("ChangePassword").
				Method(artisan_core.PUT, "ChangePassword",
					artisan_core.Request(
						// Old Password: 旧密码
						artisan_core.Param("old_password", artisan_core.String, required),
						// New Password: 新密码
						artisan_core.Param("new_password", artisan_core.String, required),
					),
					artisan_core.Reply(codeField),
				),
			).
			SubCate("/inspect", artisan_core.Ink().WithName("Inspect").
				Method(artisan_core.GET, "InspectUser",
					artisan_core.Request(),
					artisan_core.Reply(
						codeField,
						artisan_core.Param("data", artisan_core.Object("InspectUserInnerReply",
							artisan_core.SPsC(
								&userModel.ID, &userModel.NickName, &userModel.UserName,
								&userModel.LastLogin, &userModel.Email, &userModel.Motto,
								&userModel.Gender),
							artisan_core.Param("identities", artisan_core.Strings),
							artisan_core.Param("success_problems", new([]uint)),
							artisan_core.Param("tried_problems", new([]uint)),
						),
						),
					)),
			).
			Method(artisan_core.DELETE, "DeleteUser",
				artisan_core.Request(),
				artisan_core.Reply(codeField),
			),
	}
	controller.Name("UserController").
		UseModel(artisan_core.Model(artisan_core.Name("user"), &userModel),
			artisan_core.Model(artisan_core.Name("valueUser"), &_valueUserModel))
	return controller
}
