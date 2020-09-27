package main

import (
	"encoding/json"
	"github.com/Myriad-Dreamin/artisan"
	"github.com/Myriad-Dreamin/boj-v6/abstract/problem"
	problem_desc "github.com/Myriad-Dreamin/boj-v6/abstract/problem-desc"
	"github.com/Myriad-Dreamin/boj-v6/abstract/user"
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
var problemUserModel = new(user.User)
var problemDescModel problem_desc.ProblemDesc

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

func wrapProblemDescToCate(cat artisan.Category, prefix string) artisan.Category {
	return cat.
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

func wrapProblemFSToCate(cate artisan.Category, prefix string) artisan.Category {
	// todo: problem fs boj/blob/master/server/router/problem-router.go#L134
	return cate.
		SubCate("/fs", artisan.Ink().WithName("ProblemFS").
			SubCate("file", artisan.Ink().WithName(prefix+"ProblemFSFileOperation").
				Method(artisan.GET, prefix+"ProblemFSStat",
					artisan.Request(
						artisan.Param("path", artisan.String, required),
					),
					artisan.Reply(
						codeField,
						artisan.Param("data", artisan.Object(prefix+"ProblemFSStatInnerReply",
							artisan.Param("name", artisan.String),
							artisan.Param("size", artisan.Int64),
							artisan.Param("is_dir", artisan.Bool),
							artisan.Param("mod_time", artisan.Time),
						)),
					),
				).
				Method(artisan.POST, prefix+"ProblemFSWrite",
					artisan.Request(
						artisan.Param("path", artisan.String, required),
					),
					artisan.Reply(
						codeField,
					),
				).
				Method(artisan.DELETE, prefix+"ProblemFSRemove",
					artisan.Request(
						artisan.Param("path", artisan.String, required),
					),
					artisan.Reply(
						codeField,
						artisan.Param("data", artisan.Object(prefix+"ProblemFSRemoveInnerReply")),
					),
				).
				SubCate("content", artisan.Ink().WithName(prefix+"ProblemFSRead").
					Method(artisan.GET, prefix+"ProblemFSRead",
						artisan.Request(
							artisan.Param("path", artisan.String, required),
						),
						// return file
						artisan.Reply(),
					)),
			).
			SubCate("directory", artisan.Ink().WithName(prefix+"ProblemFSDirectoryOperation").
				Method(artisan.GET, prefix+"ProblemFSLS",
					artisan.Request(
						artisan.Param("path", artisan.String, required),
					),
					artisan.Reply(
						codeField,
						artisan.ArrayParam(artisan.Param("data", artisan.Object(prefix+"ProblemFSLSInnerReply",
							artisan.Param("name", artisan.String),
							artisan.Param("size", artisan.Int64),
							artisan.Param("is_dir", artisan.Bool),
							artisan.Param("mod_time", artisan.Time),
						))),
					),
				).
				Method(artisan.POST, prefix+"ProblemFSWrites",
					artisan.Request(
						artisan.Param("path", artisan.String, required),
					),
					artisan.Reply(
						codeField,
					),
				).
				Method(artisan.PUT, prefix+"ProblemFSMkdir",
					artisan.Request(
						artisan.Param("path", artisan.String, required),
					),
					artisan.Reply(
						codeField,
					),
				).
				Method(artisan.DELETE, prefix+"ProblemFSRemoveAll",
					artisan.Request(
						artisan.Param("path", artisan.String, required),
					),
					artisan.Reply(
						codeField,
					),
				).
				SubCate("zip", artisan.Ink().WithName(prefix+"ProblemFSZipOperation").
					Method(artisan.POST, prefix+"ProblemFSZipWrite",
						artisan.Request(
							artisan.Param("path", artisan.String, required),
							// zip file in request stream
						),
						artisan.Reply(
							codeField,
						),
					).
					Method(artisan.GET, prefix+"ProblemFSZipRead",
						artisan.Request(
							artisan.Param("path", artisan.String, required),
						),
						artisan.Reply(
							codeField,
						),
					),
				),
			).
			SubCate("config", artisan.Ink().WithName(prefix+"ProblemFSConfigOperation").
				Method(artisan.GET, prefix+"ProblemFSReadConfig",
					artisan.Request(
						artisan.Param("path", artisan.String),
					),
					artisan.Reply(
						codeField,
						artisan.Param("data", new(*problemconfig.ProblemConfig)),
					),
				).
				Method(artisan.POST, prefix+"ProblemFSWriteConfig",
					artisan.Request(
						artisan.Param("path", artisan.String),
						// upload: form file
					),
					artisan.Reply(
						codeField,
					),
				).
				Method(artisan.PUT, prefix+"ProblemFSPutConfig",
					artisan.Request(
						artisan.Param("path", artisan.String),
						artisan.Param("key", artisan.String),
						artisan.Param("value", new(json.RawMessage)),
					),
					artisan.Reply(
						codeField,
						artisan.Param("data", new(*problemconfig.ProblemConfig)),
					),
				),
			),
		// todo testcases
		//.
		//SubCate("write-testcases", artisan.Ink().WithName(prefix+"ProblemFSWriteTestCases").
		//	Method(artisan.POST, prefix+"ProblemFSWriteTestCases",
		//		artisan.Request(
		//			artisan.Param("path", artisan.String, required),
		//		),
		//		artisan.Reply(
		//			codeField,
		//			artisan.Param("data", artisan.Object(prefix+"ProblemFSWriteTestCasesInnerReply",
		//			)),
		//		),
		//	)).
		)
}

func getProblemIDCate(prefix string) artisan.Category {

	var problemDescObject = artisan.Object(prefix+"ProblemDescData",
		artisan.SnakeParam(&problemDescModel.Name),
		artisan.SnakeParam(&problemDescModel.UpdatedAt),
	)

	cate := artisan.Ink().
		Path("problem/:pid").Meta(&Meta{artisan.RouterMeta{
		RuntimeRouterMeta: "problem:pid",
	}}).
		Method(artisan.GET, "Get"+prefix+"Problem",
			artisan.Request(),
			StdReply(artisan.Object(
				"Get"+prefix+"ProblemData",
				artisan.SnakeParam(&problemModel.ID),
				artisan.SnakeParam(&problemModel.CreatedAt),
				artisan.SnakeParam(&problemModel.UpdatedAt),
				artisan.SnakeParam(&problemModel.IsSpj),
				artisan.SnakeParam(&problemModel.Title),
				artisan.SnakeParam(&problemModel.Description),
				artisan.SnakeParam(&problemModel.DescriptionRef),
				artisan.SnakeParam(&problemModel.TimeLimit),
				artisan.SnakeParam(&problemModel.MemoryLimit),
				artisan.SnakeParam(&problemModel.CodeLengthLimit),
				artisan.Param("author",
					artisan.Object(
						"Get"+prefix+"ProblemAuthorData",
						artisan.SnakeParam(&problemUserModel.ID),
						artisan.SnakeParam(&problemUserModel.NickName),
					),
				),
			)),
		).
		Method(artisan.PUT, "Put"+prefix+"Problem",
			artisan.Request(
				artisan.SPsC(&problemModel.Title, &problemModel.DescriptionRef),
			),
			artisan.Reply(codeField),
		).
		Method(artisan.DELETE, "Delete"+prefix+"Problem",
			artisan.Request(),
			artisan.Reply(codeField),
		).
		SubCate("/desc-list", artisan.Ink().WithName("List"+prefix+"ProblemDesc").
			Method(artisan.GET, "List"+prefix+"ProblemDesc",
				artisan.QT("List"+prefix+"ProblemDescRequest", mytraits.Filter{}),
				artisan.Reply(
					codeField,
					artisan.ArrayParam(artisan.Param("data", problemDescObject))),
			),
		).
		SubCate("/desc-count", artisan.Ink().WithName("Count"+prefix+"ProblemDesc").
			Method(artisan.GET, "Count"+prefix+"ProblemDesc",
				artisan.QT("Count"+prefix+"ProblemDescRequest", mytraits.Filter{}),
				artisan.Reply(
					codeField,
					artisan.Param("data", artisan.Int64)),
			),
		)
	return wrapProblemDescToCate(cate, prefix)
}

func DescribeProblemController() artisan.ProposingService {

	controller := &ProblemCategories{
		List:  getListProblemCate(""),
		Count: getCountProblemCate(""),
		Post:  getPostProblemCate(""),
	}

	controller.IdGroup = getProblemIDCate("")
	controller.IdGroup = wrapProblemFSToCate(controller.IdGroup, "")

	controller.Name("ProblemController").
		UseModel(
			artisan.Model(artisan.Name("problem"), &problemModel),
			artisan.Model(artisan.Name("problemUser"), &problemUserModel),
			artisan.Model(artisan.Name("problemDesc"), &problemDescModel),
		)
	return controller
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
