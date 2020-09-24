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
		Method(artisan.GET, "List"+prefix+"Problem",
			artisan.QT("List"+prefix+"ProblemRequest", mytraits.Filter{}),
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
			artisan.QT("Count"+prefix+"ProblemRequest", mytraits.Filter{}),
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
				StdReply(artisan.Object(
					"Post"+prefix+"ProblemData",
					artisan.SnakeParam(&problemModel.ID),
				)),
			),
		)
}

func getProblemIDCate(prefix string) artisan.Category {

	var problemDescObject = artisan.Object(prefix+"ProblemDesc",
		artisan.Param("name", artisan.String),
		artisan.Param("content", artisan.String),
	)

	// todo: problem fs boj/blob/master/server/router/problem-router.go#L134
	return artisan.Ink().
		Path("problem/:pid").Meta(&Meta{artisan.RouterMeta{
		RuntimeRouterMeta: "problem:pid",
	}}).
		Method(artisan.GET, "Get"+prefix+"Problem",
			artisan.Request(),
			artisan.Reply(
				codeField,
				artisan.Param("data", &problemModel),
			)).
		Method(artisan.PUT, "Put"+prefix+"Problem",
			artisan.Request(
				artisan.SPsC(&problemModel.Title, &problemModel.Description, &problemModel.DescriptionRef),
			),
			artisan.Reply(codeField),
		).
		Method(artisan.DELETE, "Delete"+prefix+"Problem",
			artisan.Request(),
			artisan.Reply(codeField),
		).
		SubCate("/desc-list", artisan.Ink().WithName("ProblemDesc").
			Method(artisan.GET, "List"+prefix+"ProblemDescs",
				artisan.Request(),
				artisan.Reply(
					codeField,
					artisan.ArrayParam(artisan.Param("data", problemDescObject))),
			),
		).
		SubCate("/desc", artisan.Ink().WithName("ProblemDesc").
			Method(artisan.POST, "Post"+prefix+"ProblemDesc",
				artisan.Request(
					artisan.Param("name", artisan.String, required),
					artisan.Param("content", artisan.String),
				),
				artisan.Reply(codeField),
			).
			Method(artisan.GET, "Get"+prefix+"ProblemDesc",
				artisan.Request(
					artisan.Param("name", artisan.String),
				),
				artisan.Reply(
					codeField,
					artisan.Param("data", artisan.String),
				)).
			Method(artisan.PUT, "Put"+prefix+"ProblemDesc",
				artisan.Request(
					artisan.Param("name", artisan.String, required),
					artisan.Param("content", artisan.String),
				),
				artisan.Reply(codeField),
			).
			SubCate("/ref", artisan.Ink().WithName("ProblemDesc").
				Method(artisan.POST, "Change"+prefix+"ProblemDescriptionRef",
					artisan.Request(
						artisan.Param("name", artisan.String, required),
						artisan.Param("new_name", artisan.String, required),
					),
					artisan.Reply(codeField),
				),
			).
			Method(artisan.DELETE, "Delete"+prefix+"ProblemDesc",
				artisan.Request(
					artisan.Param("name", artisan.String),
				),
				artisan.Reply(codeField),
			),
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
