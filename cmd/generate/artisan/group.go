package main

import (
	"github.com/Myriad-Dreamin/artisan"
	"github.com/Myriad-Dreamin/boj-v6/abstract/group"
	"github.com/Myriad-Dreamin/boj-v6/abstract/user"
	"github.com/Myriad-Dreamin/go-model-traits/example-traits"
)

type GroupCategories struct {
	artisan.VirtualService
	List       artisan.Category
	Count      artisan.Category
	Post       artisan.Category
	GetContent artisan.Category
	IdGroup    artisan.Category
}

func DescribeGroupService() artisan.ProposingService {
	var groupModel = new(group.Group)
	var _groupModel = new(group.Group)
	var valueUserModel user.User

	svc := &GroupCategories{
		List: artisan.Ink().
			Path("group-list").
			Method(artisan.GET, "ListGroups",
				artisan.QT("ListGroupsRequest", mytraits.Filter{}),
				artisan.Reply(
					codeField,
					artisan.ArrayParam(artisan.Param("data", _groupModel)),
				),
			),
		Count: artisan.Ink().
			Path("group-count").
			Method(artisan.GET, "CountGroup",
				artisan.QT("CountGroupsRequest", mytraits.Filter{}),
				artisan.Reply(
					codeField,
					artisan.Param("data", artisan.Int64),
				),
			),
		Post: artisan.Ink().
			Path("group").
			Method(artisan.POST, "PostGroup", artisan.AuthMeta("~"),
				artisan.Request(
					artisan.SnakeParam(&groupModel.Name, required),
					artisan.SnakeParam(&groupModel.Description, required),
					artisan.Param("owner_name", artisan.String),
					artisan.SnakeParam(&groupModel.OwnerID),
				),
				artisan.Reply(
					codeField,
					artisan.Param("data", &groupModel.ID),
				),
			),
		IdGroup: artisan.Ink().
			Path("group/:gid").Meta(&Meta{artisan.RouterMeta{
			RuntimeRouterMeta: "group:gid",
		}}).
			Method(artisan.GET, "GetGroup",
				artisan.Reply(
					codeField,
					artisan.Param("data", &groupModel),
				)).
			Method(artisan.PUT, "PutGroup",
				artisan.Request(
					artisan.SnakeParam(&groupModel.Name),
					artisan.SnakeParam(&groupModel.Description),
				)).
			Method(artisan.DELETE, "DeleteGroup").
			SubCate("/owner", artisan.Ink().WithName("Owner").
				Method(artisan.PUT, "PutGroupOwner",
					artisan.Request(
						artisan.SnakeParam(&groupModel.OwnerID, required)),
				),
			).
			SubCate("/user-list", artisan.Ink().WithName("UserList").
				Method(artisan.GET, "GetGroupMembers",
					artisan.QT("GroupUserListRequest", mytraits.Filter{}),
					artisan.Reply(
						codeField,
						artisan.ArrayParam(artisan.Param("data",
							artisan.Object("ListGroupUserReply", artisan.SPsC(
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
			SubCate("user/:id", artisan.Ink().WithName("User").Meta(&Meta{artisan.RouterMeta{
				RuntimeRouterMeta: "user:id",
			}}).Method(artisan.POST, "PostGroupMember",
				artisan.Request(),
				artisan.Reply(
					codeField),
			)),
		// todo: post user by name
	}
	svc.Name("GroupService").
		UseModel(artisan.Model(artisan.Name("group"), &groupModel),
			artisan.Model(artisan.Name("valueUser"), &valueUserModel))
	return svc
}
