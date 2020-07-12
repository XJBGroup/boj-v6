package main

import (
	"github.com/Myriad-Dreamin/artisan"
	"github.com/Myriad-Dreamin/boj-v6/abstract/group"
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
					artisan.ArrayParam(artisan.Param("data", new(int))),
				),
			),
		Post: artisan.Ink().
			Path("group").
			Method(artisan.POST, "PostGroup", artisan.AuthMeta("~"),
				artisan.Request(
				//artisan.SPsC(&groupModel.Title, &groupModel.Content),
				),
				artisan.Reply(
					codeField,
					artisan.Param("group", &groupModel),
				),
			),
		IdGroup: artisan.Ink().
			Path("group/:gid").Meta(&Meta{artisan.RouterMeta{
			RuntimeRouterMeta: "group:gid",
		}}).
			Method(artisan.GET, "GetGroup",
				artisan.Reply(
					codeField,
					artisan.Param("group", &groupModel),
				)).
			Method(artisan.PUT, "PutGroup",
				artisan.Request(
				//artisan.SPsC(&groupModel.Title, &groupModel.Content),
				)).
			Method(artisan.DELETE, "Delete"),
	}
	svc.Name("GroupService").
		UseModel(artisan.Model(artisan.Name("group"), &groupModel))
	return svc
}
