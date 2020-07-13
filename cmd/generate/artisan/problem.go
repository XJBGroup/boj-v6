package main

import (
	"github.com/Myriad-Dreamin/artisan"
	"github.com/Myriad-Dreamin/boj-v6/abstract/problem"
	problemconfig "github.com/Myriad-Dreamin/boj-v6/types/problem-config"
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

var problemModel = new(problem.Problem)
var _problemModel = new(problem.Problem)

func getListProblemCate(prefix string) artisan.Category {
	return artisan.Ink().
		Path("problem-list").
		Method(artisan.GET, "List"+prefix+"Problems",
			artisan.QT("List"+prefix+"ProblemsRequest", mytraits.Filter{}),
			artisan.Reply(
				codeField,
				artisan.ArrayParam(artisan.Param("data", _problemModel)),
			),
		)
}

func getCountProblemCate(prefix string) artisan.Category {
	return artisan.Ink().
		Path("problem-count").
		Method(artisan.GET, "Count"+prefix+"Problem",
			artisan.QT("Count"+prefix+"ProblemsRequest", mytraits.Filter{}),
			artisan.Reply(
				codeField,
				artisan.ArrayParam(artisan.Param("data", new(int))),
			),
		)
}

func getPostProblemCate(prefix string) artisan.Category {
	return artisan.Ink().
		Path("problem").
		Method(artisan.POST, "Post"+prefix+"Problem", artisan.AuthMeta("~"),
			artisan.Request(
				artisan.SnakeParam(&problemModel.Title, required),
				artisan.SnakeParam(&problemModel.Description),
				artisan.Param("config", new(*problemconfig.ProblemConfig)),
			),
			artisan.Reply(
				codeField,
				artisan.SnakeParam(&problemModel.ID),
			),
		)
}

func getProblemIDCate(prefix string) artisan.Category {

	var problemDescObject = artisan.Object(prefix+"ProblemDesc",
		artisan.Param("name", artisan.String),
		artisan.Param("content", artisan.String),
	)

	return artisan.Ink().
		Path("problem/:pid").Meta(&Meta{artisan.RouterMeta{
		RuntimeRouterMeta: "problem:pid",
	}}).
		Method(artisan.GET, "Get"+prefix+"Problem",
			artisan.Reply(
				codeField,
				artisan.Param("problem", &problemModel),
			)).
		Method(artisan.PUT, "Put"+prefix+"Problem",
			artisan.Request(
				artisan.SPsC(&problemModel.Title, &problemModel.Description, &problemModel.DescriptionRef),
			)).
		Method(artisan.DELETE, "Delete"+prefix+"Problem").
		SubCate("/template-list", artisan.Ink().WithName("Desc").
			Method(artisan.GET, "List"+prefix+"ProblemDescs",
				artisan.Reply(
					codeField,
					artisan.ArrayParam(artisan.Param("data", problemDescObject))),
			),
		).
		SubCate("/template", artisan.Ink().WithName("Desc").
			Method(artisan.POST, "Post"+prefix+"Desc",
				artisan.Request(
					artisan.Param("name", artisan.String, required),
					artisan.Param("content", artisan.String),
				)).
			Method(artisan.GET, "Get"+prefix+"Desc",
				artisan.Request(
					artisan.Param("name", artisan.String),
				),
				artisan.Reply(
					codeField,
					artisan.Param("data", problemDescObject),
				)).
			Method(artisan.PUT, "Put"+prefix+"Desc",
				artisan.Request(
					artisan.Param("name", artisan.String, required),
					artisan.Param("content", artisan.String),
				)).
			SubCate("/template", artisan.Ink().WithName("Desc").
				Method(artisan.POST, "Change"+prefix+"DescriptionRef",
					artisan.Request(
						artisan.Param("name", artisan.String, required),
						artisan.Param("new_name", artisan.String, required),
					)),
			).
			Method(artisan.DELETE, "Delete"+prefix+"Desc"),
		)
}

func DescribeProblemService() artisan.ProposingService {

	svc := &ProblemCategories{
		List:    getListProblemCate(""),
		Count:   getCountProblemCate(""),
		Post:    getPostProblemCate(""),
		IdGroup: getProblemIDCate(""),
	}
	svc.Name("ProblemService").
		UseModel(artisan.Model(artisan.Name("problem"), &problemModel))
	return svc
}

func DescribeProblemCategory(c artisan.Category, prefix string) artisan.Category {

	var ListCate = getListProblemCate(prefix).WithName("List")
	var CountCate = getCountProblemCate(prefix).WithName("Count")
	var PostCate = getPostProblemCate(prefix).WithName("Post")
	var IdGroupCate = getProblemIDCate(prefix).WithName("IdGroup")
	return c.SubCate(ListCate.GetPath(), ListCate).
		SubCate(CountCate.GetPath(), CountCate).
		SubCate(PostCate.GetPath(), PostCate).
		SubCate(IdGroupCate.GetPath(), IdGroupCate)
}
