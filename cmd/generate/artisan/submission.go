package main

import (
	"github.com/Myriad-Dreamin/artisan"
	"github.com/Myriad-Dreamin/boj-v6/abstract/submission"
)

type SubmissionCategories struct {
	artisan.VirtualService
	List       artisan.Category
	Count      artisan.Category
	Post       artisan.Category
	GetContent artisan.Category
	IdGroup    artisan.Category
}

func DescribeSubmissionService() artisan.ProposingService {
	var submissionModel = new(submission.Submission)
	var valueSubmissionModel submission.Submission

	var listParams = []interface{}{
		artisan.Param("page", artisan.Int),
		artisan.Param("page_size", artisan.Int),
		artisan.Param("mem_order", new(*bool)),
		artisan.Param("time_order", new(*bool)),
		artisan.Param("id_order", new(*bool)),
		artisan.Param("by_user", artisan.Uint),
		artisan.Param("on_problem", artisan.Uint),
		artisan.Param("with_language", artisan.Uint8),
		artisan.Param("has_status", artisan.Int64),
	}

	var submissionFilter = artisan.Object(
		append(listParams, "SubmissionFilter")...)

	svc := &SubmissionCategories{
		List: artisan.Ink().
			Path("submission-list").
			Method(artisan.GET, "ListSubmissions",
				artisan.Request(submissionFilter),
				artisan.Reply(
					codeField,
					artisan.ArrayParam(artisan.Param("data", artisan.Object(
						"ListSubmissionReply",
						artisan.SPsC(
							&valueSubmissionModel.ID, &valueSubmissionModel.CreatedAt, &valueSubmissionModel.ProblemID,
							&valueSubmissionModel.UserID, &valueSubmissionModel.Score, &valueSubmissionModel.Status,
							&valueSubmissionModel.RunTime, &valueSubmissionModel.RunMemory, &valueSubmissionModel.CodeLength,
							&valueSubmissionModel.Language, &valueSubmissionModel.Shared)))),
				),
			),
		Count: artisan.Ink().
			Path("submission-count").
			Method(artisan.GET, "CountSubmissions",
				artisan.Request(submissionFilter),
				artisan.Reply(
					codeField,
					artisan.Param("data", artisan.Int64),
				),
			),
		Post: artisan.Ink().
			Path("/problem/:pid/submission").Meta(&Meta{artisan.RouterMeta{
			RuntimeRouterMeta: "problem:pid",
		}}).
			Method(artisan.POST, "PostSubmission", artisan.AuthMeta("~"),
				artisan.Request(
					artisan.SPsC(&submissionModel.Information, &submissionModel.Shared),
					artisan.SnakeParam(&submissionModel.Language, required),
					artisan.Param("code", artisan.String, required),
				),
				artisan.Reply(
					codeField,
					StdReply(artisan.Object(
						"PostSubmissionData",
						artisan.SnakeParam(&submissionModel.ID),
					)),
				),
			),
		IdGroup: artisan.Ink().
			Path("submission/:sid").Meta(&Meta{artisan.RouterMeta{
			RuntimeRouterMeta: "submission:sid",
		}}).
			Method(artisan.GET, "GetSubmission",
				artisan.Reply(
					codeField,
					artisan.Param("submission", artisan.Object("GetSubmissionInnerReply",
						artisan.SPsC(
							&submissionModel.ID, &submissionModel.CreatedAt, &submissionModel.ProblemID,
							&submissionModel.UserID, &submissionModel.Score, &submissionModel.Status,
							&submissionModel.RunTime, &submissionModel.RunMemory, &submissionModel.CodeLength,
							&submissionModel.Language, &submissionModel.Shared)),
					),
				)).
			SubCate("/content", artisan.Ink().WithName("GetContent").
				Method(artisan.GET, "GetContent"),
			).
			Method(artisan.DELETE, "Delete"),
	}
	svc.Name("SubmissionService").
		UseModel(artisan.Model(artisan.Name("submission"), &submissionModel),
			artisan.Model(artisan.Name("valueSubmission"), &valueSubmissionModel))
	return svc
}
