package inner_control

import (
	"github.com/Myriad-Dreamin/boj-v6/api"
	"github.com/Myriad-Dreamin/minimum-lib/controller"
)

type InnerCommentService interface {
	CommentServiceSignatureXXX() interface{}
	ListComment(c controller.MContext, req *api.ListCommentRequest) (*api.ListCommentReply, error)
	CountComment(c controller.MContext, req *api.CountCommentRequest) (*api.CountCommentReply, error)
	PostComment(c controller.MContext, req *api.PostCommentRequest) (*api.PostCommentReply, error)
	GetComment(c controller.MContext, req *api.GetCommentRequest) (*api.GetCommentReply, error)
	PutComment(c controller.MContext, req *api.PutCommentRequest) (*api.PutCommentReply, error)
	DeleteComment(c controller.MContext, req *api.DeleteCommentRequest) (*api.DeleteCommentReply, error)
}
