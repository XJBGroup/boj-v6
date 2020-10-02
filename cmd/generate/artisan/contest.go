package main

import (
	"github.com/Myriad-Dreamin/artisan/artisan-core"
	"github.com/Myriad-Dreamin/boj-v6/abstract/contest"
	"github.com/Myriad-Dreamin/boj-v6/abstract/user"
	"github.com/Myriad-Dreamin/go-model-traits/example-traits"
)

type ContestCategories struct {
	artisan_core.VirtualService
	List       artisan_core.Category
	Count      artisan_core.Category
	Post       artisan_core.Category
	GetContent artisan_core.Category
	IdGroup    artisan_core.Category
}

func DescribeContestController() artisan_core.ProposingService {
	var contestModel = new(contest.Contest)
	var _contestModel = new(contest.Contest)

	controller := &ContestCategories{
		List: artisan_core.Ink().
			Path("contest-list").
			Method(artisan_core.GET, "ListContest",
				artisan_core.QT("ListContestRequest", mytraits.Filter{}),
				artisan_core.Reply(
					codeField,
					artisan_core.ArrayParam(artisan_core.Param("data", _contestModel)),
				),
			),
		Count: artisan_core.Ink().
			Path("contest-count").
			Method(artisan_core.GET, "CountContest",
				artisan_core.QT("CountContestRequest", mytraits.Filter{}),
				artisan_core.Reply(
					codeField,
					artisan_core.ArrayParam(artisan_core.Param("data", new(int))),
				),
			),
		Post: artisan_core.Ink().
			Path("contest").
			Method(artisan_core.POST, "PostContest", artisan_core.AuthMeta("~"),
				artisan_core.Request(
					artisan_core.SnakeParam(&contestModel.Title, required),
					artisan_core.SnakeParam(&contestModel.Description, required),
					artisan_core.SnakeParam(&contestModel.StartAt, required),
					artisan_core.SnakeParam(&contestModel.EndDuration, required),
					artisan_core.SnakeParam(&contestModel.BoardFrozenDuration, required),
				),
				artisan_core.Reply(
					codeField,
					artisan_core.Param("data", &contestModel),
				),
			),
		IdGroup: artisan_core.Ink().
			Path("contest/:cid").Meta(&Meta{artisan_core.RouterMeta{
			RuntimeRouterMeta: "contest:cid",
		}}).
			Method(artisan_core.GET, "GetContest",
				artisan_core.Request(),
				artisan_core.Reply(
					codeField,
					artisan_core.Param("data", artisan_core.Object("GetContestInnerReply",
						artisan_core.SPsC(
							&contestModel.ID, &contestModel.Title, &contestModel.StartAt, &contestModel.CreatedAt,
							&contestModel.BoardFrozenDuration, &contestModel.EndDuration, &contestModel.Description,
							&contestModel.AuthorID, &contestModel.ContestType),
					)),
				),
			).
			SubCate("/user-list", artisan_core.Ink().WithName("ListContestUsers").
				Method(artisan_core.GET, "ListContestUsers",
					artisan_core.Request(),
					artisan_core.Reply(
						codeField,
						artisan_core.ArrayParam(artisan_core.Param("data", new(user.User))),
					),
				),
			).
			// todo: user management
			Method(artisan_core.PUT, "PutContest", artisan_core.AuthMeta("~"),
				artisan_core.Request(
					artisan_core.SnakeParam(&contestModel.Title),
					artisan_core.SnakeParam(&contestModel.Description),
					artisan_core.SnakeParam(&contestModel.StartAt),
					artisan_core.SnakeParam(&contestModel.EndDuration),
					artisan_core.SnakeParam(&contestModel.BoardFrozenDuration),
					artisan_core.SnakeParam(&contestModel.ConfigPath),
					artisan_core.SnakeParam(&contestModel.RolePath),
				),
				artisan_core.Reply(codeField),
			).
			Method(artisan_core.DELETE, "DeleteContest",
				artisan_core.Request(),
				artisan_core.Reply(codeField),
			),
	}

	controller.IdGroup = DescribeProblemCategory(controller.IdGroup, "Contest")

	controller.Name("ContestController").
		UseModel(
			artisan_core.Model(artisan_core.Name("contest"), &contestModel),
			artisan_core.Model(artisan_core.Name("problem"), &problemModel),
			artisan_core.Model(artisan_core.Name("problemUser"), &problemUserModel),
			artisan_core.Model(artisan_core.Name("problemDesc"), &problemDescModel),
		)
	return controller
}
