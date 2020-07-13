package main

import (
	"github.com/Myriad-Dreamin/artisan"
	"github.com/Myriad-Dreamin/boj-v6/abstract/contest"
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
			Method(artisan.GET, "ListContests",
				artisan.QT("ListContestsRequest", mytraits.Filter{}),
				artisan.Reply(
					codeField,
					artisan.ArrayParam(artisan.Param("data", _contestModel)),
				),
			),
		Count: artisan.Ink().
			Path("contest-count").
			Method(artisan.GET, "CountContest",
				artisan.QT("CountContestsRequest", mytraits.Filter{}),
				artisan.Reply(
					codeField,
					artisan.ArrayParam(artisan.Param("data", new(int))),
				),
			),
		Post: artisan.Ink().
			Path("contest").
			Method(artisan.POST, "PostContest", artisan.AuthMeta("~"),
				artisan.Request(
				//artisan.SPsC(&contestModel.Title, &contestModel.Content),
				),
				artisan.Reply(
					codeField,
					artisan.Param("contest", &contestModel),
				),
			),
		IdGroup: artisan.Ink().
			Path("contest/:cid").Meta(&Meta{artisan.RouterMeta{
			RuntimeRouterMeta: "contest:cid",
		}}).
			Method(artisan.GET, "GetContest",
				artisan.Reply(
					codeField,
					artisan.Param("contest", &contestModel),
				)).
			Method(artisan.PUT, "PutContest",
				artisan.Request(
				//artisan.SPsC(&contestModel.Title, &contestModel.Content),
				)).
			Method(artisan.DELETE, "Delete"),
	}

	svc.IdGroup = DescribeProblemCategory(svc.IdGroup, "Contest")

	svc.Name("ContestService").
		UseModel(artisan.Model(artisan.Name("contest"), &contestModel),
			artisan.Model(artisan.Name("problem"), &problemModel))
	return svc
}
