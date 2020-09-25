package main

import (
	"github.com/Myriad-Dreamin/artisan"
	"github.com/Myriad-Dreamin/boj-v6/abstract/contest"
	"github.com/Myriad-Dreamin/boj-v6/abstract/user"
	"github.com/Myriad-Dreamin/go-model-traits/example-traits"
)

type ContestCategories struct {
	artisan.VirtualService
	List       artisan.Category
	Count      artisan.Category
	Post       artisan.Category
	GetContent artisan.Category
	IdGroup    artisan.Category
}

func DescribeContestService() artisan.ProposingService {
	var contestModel = new(contest.Contest)
	var _contestModel = new(contest.Contest)

	svc := &ContestCategories{
		List: artisan.Ink().
			Path("contest-list").
			Method(artisan.GET, "ListContest",
				artisan.QT("ListContestRequest", mytraits.Filter{}),
				artisan.Reply(
					codeField,
					artisan.ArrayParam(artisan.Param("data", _contestModel)),
				),
			),
		Count: artisan.Ink().
			Path("contest-count").
			Method(artisan.GET, "CountContest",
				artisan.QT("CountContestRequest", mytraits.Filter{}),
				artisan.Reply(
					codeField,
					artisan.ArrayParam(artisan.Param("data", new(int))),
				),
			),
		Post: artisan.Ink().
			Path("contest").
			Method(artisan.POST, "PostContest", artisan.AuthMeta("~"),
				artisan.Request(
					artisan.SnakeParam(&contestModel.Title, required),
					artisan.SnakeParam(&contestModel.Description, required),
					artisan.SnakeParam(&contestModel.StartAt, required),
					artisan.SnakeParam(&contestModel.EndDuration, required),
					artisan.SnakeParam(&contestModel.BoardFrozenDuration, required),
				),
				artisan.Reply(
					codeField,
					artisan.Param("data", &contestModel),
				),
			),
		IdGroup: artisan.Ink().
			Path("contest/:cid").Meta(&Meta{artisan.RouterMeta{
			RuntimeRouterMeta: "contest:cid",
		}}).
			Method(artisan.GET, "GetContest",
				artisan.Request(),
				artisan.Reply(
					codeField,
					artisan.Param("data", artisan.Object("GetContestInnerReply",
						artisan.SPsC(
							&contestModel.ID, &contestModel.Title, &contestModel.StartAt, &contestModel.CreatedAt,
							&contestModel.BoardFrozenDuration, &contestModel.EndDuration, &contestModel.Description,
							&contestModel.AuthorID, &contestModel.ContestType),
					)),
				),
			).
			SubCate("/user-list", artisan.Ink().WithName("ListContestUsers").
				Method(artisan.GET, "ListContestUsers",
					artisan.Request(),
					artisan.Reply(
						codeField,
						artisan.ArrayParam(artisan.Param("data", new(user.User))),
					),
				),
			).
			// todo: user management
			Method(artisan.PUT, "PutContest", artisan.AuthMeta("~"),
				artisan.Request(
					artisan.SnakeParam(&contestModel.Title),
					artisan.SnakeParam(&contestModel.Description),
					artisan.SnakeParam(&contestModel.StartAt),
					artisan.SnakeParam(&contestModel.EndDuration),
					artisan.SnakeParam(&contestModel.BoardFrozenDuration),
					artisan.SnakeParam(&contestModel.ConfigPath),
					artisan.SnakeParam(&contestModel.RolePath),
				),
				artisan.Reply(codeField),
			).
			Method(artisan.DELETE, "DeleteContest",
				artisan.Request(),
				artisan.Reply(codeField),
			),
	}

	svc.IdGroup = DescribeProblemCategory(svc.IdGroup, "Contest")

	svc.Name("ContestService").
		UseModel(
			artisan.Model(artisan.Name("contest"), &contestModel),
			artisan.Model(artisan.Name("problem"), &problemModel),
			artisan.Model(artisan.Name("problemUser"), &problemUserModel),
			artisan.Model(artisan.Name("problemDesc"), &problemDescModel),
		)
	return svc
}
