package main

import (
	"encoding/json"
	"github.com/Myriad-Dreamin/artisan/artisan-core"
	"github.com/Myriad-Dreamin/boj-v6/abstract/problem"
	problem_desc "github.com/Myriad-Dreamin/boj-v6/abstract/problem-desc"
	"github.com/Myriad-Dreamin/boj-v6/abstract/user"
	problemconfig "github.com/Myriad-Dreamin/boj-v6/types/problem-config"
	"github.com/Myriad-Dreamin/go-model-traits/example-traits"
)

type ProblemCategories struct {
	artisan_core.VirtualService
	List       artisan_core.Category
	Count      artisan_core.Category
	Post       artisan_core.Category
	GetContent artisan_core.Category
	IdGroup    artisan_core.Category
}

var problemModel = new(problem.Problem)
var _problemModel = new(problem.Problem)
var problemUserModel = new(user.User)
var problemDescModel problem_desc.ProblemDesc

func getListProblemCate(prefix string) artisan_core.Category {
	return artisan_core.Ink().
		Path("problem-list").
		Method(artisan_core.GET, "List"+prefix+"Problem",
			artisan_core.QT("List"+prefix+"ProblemRequest", mytraits.Filter{}),
			artisan_core.Reply(
				codeField,
				artisan_core.ArrayParam(artisan_core.Param("data", _problemModel)),
			),
		)
}

func getCountProblemCate(prefix string) artisan_core.Category {
	return artisan_core.Ink().
		Path("problem-count").
		Method(artisan_core.GET, "Count"+prefix+"Problem",
			artisan_core.QT("Count"+prefix+"ProblemRequest", mytraits.Filter{}),
			artisan_core.Reply(
				codeField,
				artisan_core.ArrayParam(artisan_core.Param("data", new(int))),
			),
		)
}

func getPostProblemCate(prefix string) artisan_core.Category {
	return artisan_core.Ink().
		Path("problem").
		Method(artisan_core.POST, "Post"+prefix+"Problem", artisan_core.AuthMeta("~"),
			artisan_core.Request(
				artisan_core.SnakeParam(&problemModel.Title, required),
				artisan_core.SnakeParam(&problemModel.Description),
				artisan_core.Param("config", new(*problemconfig.ProblemConfig)),
			),
			artisan_core.Reply(
				codeField,
				StdReply(artisan_core.Object(
					"Post"+prefix+"ProblemData",
					artisan_core.SnakeParam(&problemModel.ID),
				)),
			),
		)
}

func wrapProblemDescToCate(cat artisan_core.Category, prefix string) artisan_core.Category {
	return cat.
		SubCate("/desc", artisan_core.Ink().WithName("ProblemDesc").
			Method(artisan_core.POST, "Post"+prefix+"ProblemDesc",
				artisan_core.Request(
					artisan_core.Param("name", artisan_core.String, required),
					artisan_core.Param("content", artisan_core.String),
				),
				artisan_core.Reply(codeField),
			).
			Method(artisan_core.GET, "Get"+prefix+"ProblemDesc",
				artisan_core.Request(
					artisan_core.Param("name", artisan_core.String),
				),
				artisan_core.Reply(
					codeField,
					artisan_core.Param("data", artisan_core.String),
				)).
			Method(artisan_core.PUT, "Put"+prefix+"ProblemDesc",
				artisan_core.Request(
					artisan_core.Param("name", artisan_core.String, required),
					artisan_core.Param("content", artisan_core.String),
				),
				artisan_core.Reply(codeField),
			).
			SubCate("/ref", artisan_core.Ink().WithName("ProblemDesc").
				Method(artisan_core.POST, "Change"+prefix+"ProblemDescriptionRef",
					artisan_core.Request(
						artisan_core.Param("name", artisan_core.String, required),
						artisan_core.Param("new_name", artisan_core.String, required),
					),
					artisan_core.Reply(codeField),
				),
			).
			Method(artisan_core.DELETE, "Delete"+prefix+"ProblemDesc",
				artisan_core.Request(
					artisan_core.Param("name", artisan_core.String),
				),
				artisan_core.Reply(codeField),
			),
		)
}

func wrapProblemFSToCate(cate artisan_core.Category, prefix string) artisan_core.Category {
	// todo: problem fs boj/blob/master/server/router/problem-router.go#L134
	return cate.
		SubCate("/fs", artisan_core.Ink().WithName("ProblemFS").
			SubCate("file", artisan_core.Ink().WithName(prefix+"ProblemFSFileOperation").
				Method(artisan_core.GET, prefix+"ProblemFSStat",
					artisan_core.Request(
						artisan_core.Param("path", artisan_core.String, required),
					),
					artisan_core.Reply(
						codeField,
						artisan_core.Param("data", artisan_core.Object(prefix+"ProblemFSStatInnerReply",
							artisan_core.Param("name", artisan_core.String),
							artisan_core.Param("size", artisan_core.Int64),
							artisan_core.Param("is_dir", artisan_core.Bool),
							artisan_core.Param("mod_time", artisan_core.Time),
						)),
					),
				).
				Method(artisan_core.POST, prefix+"ProblemFSWrite",
					artisan_core.Request(
						artisan_core.Param("path", artisan_core.String, required),
					),
					artisan_core.Reply(
						codeField,
					),
				).
				Method(artisan_core.DELETE, prefix+"ProblemFSRemove",
					artisan_core.Request(
						artisan_core.Param("path", artisan_core.String, required),
					),
					artisan_core.Reply(
						codeField,
						artisan_core.Param("data", artisan_core.Object(prefix+"ProblemFSRemoveInnerReply")),
					),
				).
				SubCate("content", artisan_core.Ink().WithName(prefix+"ProblemFSRead").
					Method(artisan_core.GET, prefix+"ProblemFSRead",
						artisan_core.Request(
							artisan_core.Param("path", artisan_core.String, required),
						),
						// return file
						artisan_core.Reply(),
					)),
			).
			SubCate("directory", artisan_core.Ink().WithName(prefix+"ProblemFSDirectoryOperation").
				Method(artisan_core.GET, prefix+"ProblemFSLS",
					artisan_core.Request(
						artisan_core.Param("path", artisan_core.String, required),
					),
					artisan_core.Reply(
						codeField,
						artisan_core.ArrayParam(artisan_core.Param("data", artisan_core.Object(prefix+"ProblemFSLSInnerReply",
							artisan_core.Param("name", artisan_core.String),
							artisan_core.Param("size", artisan_core.Int64),
							artisan_core.Param("is_dir", artisan_core.Bool),
							artisan_core.Param("mod_time", artisan_core.Time),
						))),
					),
				).
				Method(artisan_core.POST, prefix+"ProblemFSWrites",
					artisan_core.Request(
						artisan_core.Param("path", artisan_core.String, required),
					),
					artisan_core.Reply(
						codeField,
					),
				).
				Method(artisan_core.PUT, prefix+"ProblemFSMkdir",
					artisan_core.Request(
						artisan_core.Param("path", artisan_core.String, required),
					),
					artisan_core.Reply(
						codeField,
					),
				).
				Method(artisan_core.DELETE, prefix+"ProblemFSRemoveAll",
					artisan_core.Request(
						artisan_core.Param("path", artisan_core.String, required),
					),
					artisan_core.Reply(
						codeField,
					),
				).
				SubCate("zip", artisan_core.Ink().WithName(prefix+"ProblemFSZipOperation").
					Method(artisan_core.POST, prefix+"ProblemFSZipWrite",
						artisan_core.Request(
							artisan_core.Param("path", artisan_core.String, required),
							// zip file in request stream
						),
						artisan_core.Reply(
							codeField,
						),
					).
					Method(artisan_core.GET, prefix+"ProblemFSZipRead",
						artisan_core.Request(
							artisan_core.Param("path", artisan_core.String, required),
						),
						artisan_core.Reply(
							codeField,
						),
					),
				),
			).
			SubCate("config", artisan_core.Ink().WithName(prefix+"ProblemFSConfigOperation").
				Method(artisan_core.GET, prefix+"ProblemFSReadConfig",
					artisan_core.Request(
						artisan_core.Param("path", artisan_core.String),
					),
					artisan_core.Reply(
						codeField,
						artisan_core.Param("data", new(*problemconfig.ProblemConfig)),
					),
				).
				Method(artisan_core.POST, prefix+"ProblemFSWriteConfig",
					artisan_core.Request(
						artisan_core.Param("path", artisan_core.String),
						// upload: form file
					),
					artisan_core.Reply(
						codeField,
					),
				).
				Method(artisan_core.PUT, prefix+"ProblemFSPutConfig",
					artisan_core.Request(
						artisan_core.Param("path", artisan_core.String),
						artisan_core.Param("key", artisan_core.String),
						artisan_core.Param("value", new(json.RawMessage)),
					),
					artisan_core.Reply(
						codeField,
						artisan_core.Param("data", new(*problemconfig.ProblemConfig)),
					),
				),
			),
		// todo testcases
		//.
		//SubCate("write-testcases", artisan_core.Ink().WithName(prefix+"ProblemFSWriteTestCases").
		//	Method(artisan_core.POST, prefix+"ProblemFSWriteTestCases",
		//		artisan_core.Request(
		//			artisan_core.Param("path", artisan_core.String, required),
		//		),
		//		artisan_core.Reply(
		//			codeField,
		//			artisan_core.Param("data", artisan_core.Object(prefix+"ProblemFSWriteTestCasesInnerReply",
		//			)),
		//		),
		//	)).
		)
}

func getProblemIDCate(prefix string) artisan_core.Category {

	var problemDescObject = artisan_core.Object(prefix+"ProblemDescData",
		artisan_core.SnakeParam(&problemDescModel.Name),
		artisan_core.SnakeParam(&problemDescModel.UpdatedAt),
	)

	cate := artisan_core.Ink().
		Path("problem/:pid").Meta(&Meta{artisan_core.RouterMeta{
		RuntimeRouterMeta: "problem:pid",
	}}).
		Method(artisan_core.GET, "Get"+prefix+"Problem",
			artisan_core.Request(),
			StdReply(artisan_core.Object(
				"Get"+prefix+"ProblemData",
				artisan_core.SnakeParam(&problemModel.ID),
				artisan_core.SnakeParam(&problemModel.CreatedAt),
				artisan_core.SnakeParam(&problemModel.UpdatedAt),
				artisan_core.SnakeParam(&problemModel.IsSpj),
				artisan_core.SnakeParam(&problemModel.Title),
				artisan_core.SnakeParam(&problemModel.Description),
				artisan_core.SnakeParam(&problemModel.DescriptionRef),
				artisan_core.SnakeParam(&problemModel.TimeLimit),
				artisan_core.SnakeParam(&problemModel.MemoryLimit),
				artisan_core.SnakeParam(&problemModel.CodeLengthLimit),
				artisan_core.Param("author",
					artisan_core.Object(
						"Get"+prefix+"ProblemAuthorData",
						artisan_core.SnakeParam(&problemUserModel.ID),
						artisan_core.SnakeParam(&problemUserModel.NickName),
					),
				),
			)),
		).
		Method(artisan_core.PUT, "Put"+prefix+"Problem",
			artisan_core.Request(
				artisan_core.SPsC(&problemModel.Title, &problemModel.DescriptionRef),
			),
			artisan_core.Reply(codeField),
		).
		Method(artisan_core.DELETE, "Delete"+prefix+"Problem",
			artisan_core.Request(),
			artisan_core.Reply(codeField),
		).
		SubCate("/desc-list", artisan_core.Ink().WithName("List"+prefix+"ProblemDesc").
			Method(artisan_core.GET, "List"+prefix+"ProblemDesc",
				artisan_core.QT("List"+prefix+"ProblemDescRequest", mytraits.Filter{}),
				artisan_core.Reply(
					codeField,
					artisan_core.ArrayParam(artisan_core.Param("data", problemDescObject))),
			),
		).
		SubCate("/desc-count", artisan_core.Ink().WithName("Count"+prefix+"ProblemDesc").
			Method(artisan_core.GET, "Count"+prefix+"ProblemDesc",
				artisan_core.QT("Count"+prefix+"ProblemDescRequest", mytraits.Filter{}),
				artisan_core.Reply(
					codeField,
					artisan_core.Param("data", artisan_core.Int64)),
			),
		)
	return wrapProblemDescToCate(cate, prefix)
}

func DescribeProblemController() artisan_core.ProposingService {

	controller := &ProblemCategories{
		List:  getListProblemCate(""),
		Count: getCountProblemCate(""),
		Post:  getPostProblemCate(""),
	}

	controller.IdGroup = getProblemIDCate("")
	controller.IdGroup = wrapProblemFSToCate(controller.IdGroup, "")

	controller.Name("ProblemController").
		UseModel(
			artisan_core.Model(artisan_core.Name("problem"), &problemModel),
			artisan_core.Model(artisan_core.Name("problemUser"), &problemUserModel),
			artisan_core.Model(artisan_core.Name("problemDesc"), &problemDescModel),
		)
	return controller
}

func DescribeProblemCategory(c artisan_core.Category, prefix string) artisan_core.Category {

	var ListCate = getListProblemCate(prefix).WithName("List")
	var CountCate = getCountProblemCate(prefix).WithName("Count")
	var PostCate = getPostProblemCate(prefix).WithName("Post")
	var IdGroupCate = getProblemIDCate(prefix).WithName("IdGroup")
	return c.SubCate(ListCate.GetPath(), ListCate).
		SubCate(CountCate.GetPath(), CountCate).
		SubCate(PostCate.GetPath(), PostCate).
		SubCate(IdGroupCate.GetPath(), IdGroupCate)
}
