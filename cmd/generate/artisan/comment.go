package main

import (
	"github.com/Myriad-Dreamin/artisan"
	"github.com/Myriad-Dreamin/boj-v6/abstract/comment"
	"github.com/Myriad-Dreamin/go-model-traits/example-traits"
)

type CommentCategories struct {
	artisan.VirtualService
	List       artisan.Category
	Count      artisan.Category
	Post       artisan.Category
	GetContent artisan.Category
	IdGroup    artisan.Category
}

func DescribeCommentService() artisan.ProposingService {
	var commentModel = new(comment.Comment)
	var _commentModel = new(comment.Comment)

	svc := &CommentCategories{
		List: artisan.Ink().
			Path("comment-list").
			Method(artisan.GET, "ListComments",
				artisan.QT("ListCommentsRequest", mytraits.Filter{}),
				artisan.Reply(
					codeField,
					artisan.ArrayParam(artisan.Param("data", _commentModel)),
				),
			),
		Count: artisan.Ink().
			Path("comment-count").
			Method(artisan.GET, "CountComment",
				artisan.QT("CountCommentsRequest", mytraits.Filter{}),
				artisan.Reply(
					codeField,
					artisan.ArrayParam(artisan.Param("data", new(int))),
				),
			),
		Post: artisan.Ink().
			Path("comment").
			Method(artisan.POST, "PostComment",
				artisan.Request(
					artisan.SPsC(&commentModel.Title, &commentModel.Content),
				),
				artisan.Reply(
					codeField,
					artisan.Param("comment", &commentModel),
				),
			),
		IdGroup: artisan.Ink().
			Path("comment/:cmid").Meta(&Meta{artisan.RouterMeta{
			RuntimeRouterMeta: "comment:cmid",
		}}).
			Method(artisan.GET, "GetComment",
				artisan.Reply(
					codeField,
					artisan.Param("comment", &commentModel),
				)).
			Method(artisan.PUT, "PutComment",
				artisan.Request(
					artisan.SPsC(&commentModel.Title, &commentModel.Content),
				)).
			Method(artisan.DELETE, "Delete"),
	}
	svc.Name("CommentService").
		UseModel(artisan.Model(artisan.Name("comment"), &commentModel))
	return svc
}
