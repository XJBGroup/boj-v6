package main

import (
	"github.com/Myriad-Dreamin/artisan/artisan-core"
	"github.com/Myriad-Dreamin/boj-v6/abstract/submission"
)

type SubmissionCategories struct {
	artisan_core.VirtualService
	List       artisan_core.Category
	Count      artisan_core.Category
	Post       artisan_core.Category
	GetContent artisan_core.Category
	IdGroup    artisan_core.Category
}

func DescribeSubmissionController() artisan_core.ProposingService {
	var submissionModel = new(submission.Submission)
	var valueSubmissionModel submission.Submission

	var listParams = []interface{}{
		artisan_core.Param("page", artisan_core.Int),
		artisan_core.Param("page_size", artisan_core.Int),
		artisan_core.Param("mem_order", new(*bool)),
		artisan_core.Param("time_order", new(*bool)),
		artisan_core.Param("id_order", new(*bool)),
		artisan_core.Param("by_user", artisan_core.Uint),
		artisan_core.Param("on_problem", artisan_core.Uint),
		artisan_core.Param("with_language", artisan_core.Uint8),
		artisan_core.Param("has_status", artisan_core.Int64),
	}

	var submissionFilter = func(name string) artisan_core.SerializeObject {
		return artisan_core.Object(
			append(listParams, name)...)
	}

	controller := &SubmissionCategories{
		List: artisan_core.Ink().
			Path("submission-list").
			Method(artisan_core.GET, "ListSubmission",
				artisan_core.Request(submissionFilter("ListSubmissionRequest")),
				artisan_core.Reply(
					codeField,
					artisan_core.ArrayParam(artisan_core.Param("data", artisan_core.Object(
						"ListSubmissionInnerReply",
						artisan_core.SPsC(
							&valueSubmissionModel.ID, &valueSubmissionModel.CreatedAt, &valueSubmissionModel.ProblemID,
							&valueSubmissionModel.UserID, &valueSubmissionModel.Score, &valueSubmissionModel.Status,
							&valueSubmissionModel.RunTime, &valueSubmissionModel.RunMemory, &valueSubmissionModel.CodeLength,
							&valueSubmissionModel.Language, &valueSubmissionModel.Shared)))),
				),
			),
		Count: artisan_core.Ink().
			Path("submission-count").
			Method(artisan_core.GET, "CountSubmission",
				artisan_core.Request(submissionFilter("CountSubmissionRequest")),
				artisan_core.Reply(
					codeField,
					artisan_core.Param("data", artisan_core.Int64),
				),
			),
		Post: artisan_core.Ink().
			Path("/problem/:pid/submission").Meta(&Meta{artisan_core.RouterMeta{
			RuntimeRouterMeta: "problem:pid",
		}}).
			Method(artisan_core.POST, "PostSubmission", artisan_core.AuthMeta("~"),
				artisan_core.Request(
					artisan_core.Param("pid", artisan_core.Uint, routeParam),
					artisan_core.SPsC(&submissionModel.Information, &submissionModel.Shared),
					artisan_core.Param("language", artisan_core.String, required),
					artisan_core.Param("code", artisan_core.String, required),
				),
				artisan_core.Reply(
					codeField,
					StdReply(artisan_core.Object(
						"PostSubmissionData",
						artisan_core.SnakeParam(&submissionModel.ID),
					)),
				),
			),
		IdGroup: artisan_core.Ink().
			Path("submission/:sid").Meta(&Meta{artisan_core.RouterMeta{
			RuntimeRouterMeta: "submission:sid",
		}}).
			Method(artisan_core.GET, "GetSubmission",
				artisan_core.Request(),
				artisan_core.Reply(
					codeField,
					artisan_core.Param("data", artisan_core.Object("GetSubmissionInnerReply",
						artisan_core.SPsC(
							&submissionModel.ID, &submissionModel.CreatedAt, &submissionModel.ProblemID,
							&submissionModel.UserID, &submissionModel.Score, &submissionModel.Status,
							&submissionModel.RunTime, &submissionModel.RunMemory, &submissionModel.CodeLength,
							&submissionModel.Language, &submissionModel.Shared)),
					),
				)).
			SubCate("/content", artisan_core.Ink().WithName("GetSubmissionContent").
				Method(artisan_core.GET, "GetSubmissionContent",
					artisan_core.Request(),
					artisan_core.Reply(codeField),
				),
			).
			Method(artisan_core.DELETE, "DeleteSubmission",
				artisan_core.Request(),
				artisan_core.Reply(codeField),
			),
	}
	controller.Name("Controller").
		UseModel(artisan_core.Model(artisan_core.Name("submission"), &submissionModel),
			artisan_core.Model(artisan_core.Name("valueSubmission"), &valueSubmissionModel))
	return controller
}
