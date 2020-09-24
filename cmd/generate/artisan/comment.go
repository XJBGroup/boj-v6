package main

import (
	"github.com/Myriad-Dreamin/artisan"
	"github.com/Myriad-Dreamin/boj-v6/abstract/comment"
)

type CommentCategories struct {
	artisan.VirtualService
	List       artisan.Category
	Count      artisan.Category
	Post       artisan.Category
	GetContent artisan.Category
	IdGroup    artisan.Category
}

//
//type Filter struct {
//	Page         int
//	PageSize     int
//	RefType uint8
//	Ref uint
//	NoReply bool
//}

func DescribeCommentService() artisan.ProposingService {
	var commentModel = new(comment.Comment)
	var _commentModel = new(comment.Comment)

	svc := &CommentCategories{
		List: artisan.Ink().
			Path("comment-list").
			Method(artisan.GET, "ListComment",
				artisan.QT("ListCommentRequest", comment.Filter{}),
				artisan.Reply(
					codeField,
					artisan.ArrayParam(artisan.Param("data", _commentModel)),
				),
			),
		Count: artisan.Ink().
			Path("comment-count").
			Method(artisan.GET, "CountComment",
				artisan.QT("CountCommentRequest", comment.Filter{}),
				artisan.Reply(
					codeField,
					artisan.Param("data", artisan.Int64),
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
				artisan.Request(),
				artisan.Reply(
					codeField,
					artisan.Param("data", &commentModel),
				)).
			Method(artisan.PUT, "PutComment",
				artisan.Request(
					artisan.SPsC(&commentModel.Title, &commentModel.Content),
				),
				artisan.Reply(codeField),
			).
			Method(artisan.DELETE, "DeleteComment",
				artisan.Request(),
				artisan.Reply(codeField),
			),
	}
	svc.Name("CommentService").
		UseModel(artisan.Model(artisan.Name("comment"), &commentModel))
	return svc
}
