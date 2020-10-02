package main

import (
	"github.com/Myriad-Dreamin/artisan/artisan-core"
	"github.com/Myriad-Dreamin/boj-v6/abstract/comment"
)

type CommentCategories struct {
	artisan_core.VirtualService
	List       artisan_core.Category
	Count      artisan_core.Category
	Post       artisan_core.Category
	GetContent artisan_core.Category
	IdGroup    artisan_core.Category
}

//
//type Filter struct {
//	Page         int
//	PageSize     int
//	RefType uint8
//	Ref uint
//	NoReply bool
//}

func DescribeCommentController() artisan_core.ProposingService {
	var commentModel = new(comment.Comment)
	var _commentModel = new(comment.Comment)

	controller := &CommentCategories{
		List: artisan_core.Ink().
			Path("comment-list").
			Method(artisan_core.GET, "ListComment",
				artisan_core.QT("ListCommentRequest", comment.Filter{}),
				artisan_core.Reply(
					codeField,
					artisan_core.ArrayParam(artisan_core.Param("data", _commentModel)),
				),
			),
		Count: artisan_core.Ink().
			Path("comment-count").
			Method(artisan_core.GET, "CountComment",
				artisan_core.QT("CountCommentRequest", comment.Filter{}),
				artisan_core.Reply(
					codeField,
					artisan_core.Param("data", artisan_core.Int64),
				),
			),
		Post: artisan_core.Ink().
			Path("comment").
			Method(artisan_core.POST, "PostComment",
				artisan_core.Request(
					artisan_core.SPsC(&commentModel.Title, &commentModel.Content),
				),
				artisan_core.Reply(
					codeField,
					artisan_core.Param("comment", &commentModel),
				),
			),
		IdGroup: artisan_core.Ink().
			Path("comment/:cmid").Meta(&Meta{artisan_core.RouterMeta{
			RuntimeRouterMeta: "comment:cmid",
		}}).
			Method(artisan_core.GET, "GetComment",
				artisan_core.Request(),
				artisan_core.Reply(
					codeField,
					artisan_core.Param("data", &commentModel),
				)).
			Method(artisan_core.PUT, "PutComment",
				artisan_core.Request(
					artisan_core.SPsC(&commentModel.Title, &commentModel.Content),
				),
				artisan_core.Reply(codeField),
			).
			Method(artisan_core.DELETE, "DeleteComment",
				artisan_core.Request(),
				artisan_core.Reply(codeField),
			),
	}
	controller.Name("CommentController").
		UseModel(artisan_core.Model(artisan_core.Name("comment"), &commentModel))
	return controller
}
