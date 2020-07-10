package main

import (
	"github.com/Myriad-Dreamin/artisan"
	"github.com/Myriad-Dreamin/boj-v6/abstract/user"

	"github.com/Myriad-Dreamin/go-model-traits/example-traits"
)

type UserCategories struct {
	artisan.VirtualService
	List         artisan.Category
	Count        artisan.Category
	ListNameLike artisan.Category
	Post         artisan.Category
	Login        artisan.Category
	GetUserToken artisan.Category
	GetContent   artisan.Category
	Inspect      artisan.Category
	IdGroup      artisan.Category
}

func DescribeUserService() artisan.ProposingService {
	var userModel = new(user.User)
	var _userModel = new(user.User)

	svc := &UserCategories{
		List: artisan.Ink().
			Path("user-list").
			Method(artisan.GET, "ListUsers",
				artisan.QT("ListUsersRequest", mytraits.Filter{}),
				artisan.Reply(
					codeField,
					artisan.ArrayParam(artisan.Param("data", _userModel)),
				),
			),
		Count: artisan.Ink().
			Path("user-count").
			Method(artisan.GET, "CountUser",
				artisan.QT("CountUsersRequest", mytraits.Filter{}),
				artisan.Reply(
					codeField,
					artisan.ArrayParam(artisan.Param("data", new(int))),
				),
			),
		ListNameLike: artisan.Ink().
			Path("user-list-name-like").
			Method(artisan.GET, "ListUsersNameLike",
				artisan.QT("ListUsersNameLikeRequest", mytraits.Filter{}),
				artisan.Reply(
					codeField,
					artisan.ArrayParam(artisan.Param("data", _userModel)),
				),
			),
		Post: artisan.Ink().
			Path("user/register").
			Method(artisan.POST, "PostUser",
				artisan.Request(
				//todo
				),
				artisan.Reply(
					codeField,
					artisan.Param("user", &userModel),
				),
			),
		Login: artisan.Ink().
			Path("user/login").
			Method(artisan.POST, "LoginUser",
				artisan.Request(
				//todo
				),
				artisan.Reply(
					codeField,
					artisan.Param("user", &userModel),
				),
			),
		GetUserToken: artisan.Ink().
			Path("user-token").
			Method(artisan.GET, "PutUserContent"),

		//todo
		Inspect: artisan.Ink().Path("user/:aid/inspect").
			Method(artisan.GET, "InspectUser",
				artisan.Reply(
					codeField,
					artisan.Param("user", &userModel),
				),
			),
		IdGroup: artisan.Ink().
			Path("user/:aid").
			Method(artisan.GET, "GetUser",
				artisan.Reply(
					codeField,
					artisan.Param("user", &userModel),
				)).
			Method(artisan.PUT, "PutUser",
				artisan.Request()).
			Method(artisan.DELETE, "Delete"),
	}
	svc.Name("UserService").
		UseModel(artisan.Model(artisan.Name("user"), &userModel))
	return svc
}
