package main

import (
	"github.com/Myriad-Dreamin/artisan"
	"github.com/Myriad-Dreamin/boj-v6/abstract/problem"
	"github.com/Myriad-Dreamin/go-model-traits/example-traits"
)

type ProblemCategories struct {
	artisan.VirtualService
	List       artisan.Category
	Count      artisan.Category
	Post       artisan.Category
	GetContent artisan.Category
	IdGroup    artisan.Category
}

func DescribeProblemService() artisan.ProposingService {
	var problemModel = new(problem.Problem)
	var _problemModel = new(problem.Problem)

	svc := &ProblemCategories{
		List: artisan.Ink().
			Path("problem-list").
			Method(artisan.GET, "ListProblems",
				artisan.QT("ListProblemsRequest", mytraits.Filter{}),
				artisan.Reply(
					codeField,
					artisan.ArrayParam(artisan.Param("data", _problemModel)),
				),
			),
		Count: artisan.Ink().
			Path("problem-count").
			Method(artisan.GET, "CountProblem",
				artisan.QT("CountProblemsRequest", mytraits.Filter{}),
				artisan.Reply(
					codeField,
					artisan.ArrayParam(artisan.Param("data", new(int))),
				),
			),
		Post: artisan.Ink().
			Path("problem").
			Method(artisan.POST, "PostProblem", artisan.AuthMeta("~"),
				artisan.Request(
				//artisan.SPsC(&problemModel.Title, &problemModel.Content),
				),
				artisan.Reply(
					codeField,
					artisan.Param("problem", &problemModel),
				),
			),
		IdGroup: artisan.Ink().
			Path("problem/:pid").Meta(&Meta{artisan.RouterMeta{
			RuntimeRouterMeta: "problem:pid",
		}}).
			Method(artisan.GET, "GetProblem",
				artisan.Reply(
					codeField,
					artisan.Param("problem", &problemModel),
				)).
			Method(artisan.PUT, "PutProblem",
				artisan.Request(
				//artisan.SPsC(&problemModel.Title, &problemModel.Content),
				)).
			Method(artisan.DELETE, "Delete"),
	}
	svc.Name("ProblemService").
		UseModel(artisan.Model(artisan.Name("problem"), &problemModel))
	return svc
}
