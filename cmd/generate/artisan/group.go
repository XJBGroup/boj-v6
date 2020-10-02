package main

import (
	"github.com/Myriad-Dreamin/artisan/artisan-core"
	"github.com/Myriad-Dreamin/boj-v6/abstract/group"
	"github.com/Myriad-Dreamin/boj-v6/abstract/user"
	"github.com/Myriad-Dreamin/go-model-traits/example-traits"
)

type GroupCategories struct {
	artisan_core.VirtualService
	List       artisan_core.Category
	Count      artisan_core.Category
	Post       artisan_core.Category
	GetContent artisan_core.Category
	IdGroup    artisan_core.Category
}

func DescribeGroupController() artisan_core.ProposingService {
	var groupModel = new(group.Group)
	var _groupModel = new(group.Group)
	var valueUserModel user.User

	controller := &GroupCategories{
		List: artisan_core.Ink().
			Path("group-list").
			Method(artisan_core.GET, "ListGroup",
				artisan_core.QT("ListGroupRequest", mytraits.Filter{}),
				artisan_core.Reply(
					codeField,
					artisan_core.ArrayParam(artisan_core.Param("data", _groupModel)),
				),
			),
		Count: artisan_core.Ink().
			Path("group-count").
			Method(artisan_core.GET, "CountGroup",
				artisan_core.QT("CountGroupRequest", mytraits.Filter{}),
				artisan_core.Reply(
					codeField,
					artisan_core.Param("data", artisan_core.Int64),
				),
			),
		Post: artisan_core.Ink().
			Path("group").
			Method(artisan_core.POST, "PostGroup", artisan_core.AuthMeta("~"),
				artisan_core.Request(
					artisan_core.SnakeParam(&groupModel.Name, required),
					artisan_core.SnakeParam(&groupModel.Description, required),
					artisan_core.Param("owner_name", artisan_core.String),
					artisan_core.SnakeParam(&groupModel.OwnerID),
				),
				artisan_core.Reply(
					codeField,
					artisan_core.Param("data", &groupModel.ID),
				),
			),
		IdGroup: artisan_core.Ink().
			Path("group/:gid").Meta(&Meta{artisan_core.RouterMeta{
			RuntimeRouterMeta: "group:gid",
		}}).
			Method(artisan_core.GET, "GetGroup",
				artisan_core.Request(),
				artisan_core.Reply(
					codeField,
					artisan_core.Param("data", &groupModel),
				)).
			Method(artisan_core.PUT, "PutGroup",
				artisan_core.Request(
					artisan_core.SnakeParam(&groupModel.Name),
					artisan_core.SnakeParam(&groupModel.Description),
				),
				artisan_core.Reply(codeField),
			).
			Method(artisan_core.DELETE, "DeleteGroup",
				artisan_core.Request(),
				artisan_core.Reply(codeField),
			).
			SubCate("/owner", artisan_core.Ink().WithName("Owner").
				Method(artisan_core.PUT, "PutGroupOwner",
					artisan_core.Request(
						artisan_core.SnakeParam(&groupModel.OwnerID, required)),
					artisan_core.Reply(codeField),
				),
			).
			SubCate("/user-list", artisan_core.Ink().WithName("UserList").
				Method(artisan_core.GET, "GetGroupMembers",
					artisan_core.QT("GetGroupMembersRequest", mytraits.Filter{}),
					artisan_core.Reply(
						codeField,
						artisan_core.ArrayParam(artisan_core.Param("data",
							artisan_core.Object("GetGroupMembersInnerReply", artisan_core.SPsC(
								&valueUserModel.ID,
								&valueUserModel.Gender,
								&valueUserModel.LastLogin,
								&valueUserModel.UserName,
								&valueUserModel.NickName,
								&valueUserModel.Email,
								&valueUserModel.Motto,
								&valueUserModel.SolvedProblemsCount,
								&valueUserModel.TriedProblemsCount,
							)))),
					),
				),
			).
			SubCate("user/:id", artisan_core.Ink().WithName("User").Meta(&Meta{artisan_core.RouterMeta{
				RuntimeRouterMeta: "user:id",
			}}).Method(artisan_core.POST, "PostGroupMember",
				artisan_core.Request(),
				artisan_core.Reply(codeField),
			)),
		// todo: post user by name
	}
	controller.Name("GroupController").
		UseModel(artisan_core.Model(artisan_core.Name("group"), &groupModel),
			artisan_core.Model(artisan_core.Name("valueUser"), &valueUserModel))
	return controller
}
