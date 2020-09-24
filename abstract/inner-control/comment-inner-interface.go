package inner_control

import (
	"github.com/Myriad-Dreamin/boj-v6/api"
)

type InnerCommentService interface {
	CommentServiceSignatureXXX() interface{}
	ListComment(req *api.ListCommentRequest) (*api.ListCommentReply, error)
	CountComment(req *api.CountCommentRequest) (*api.CountCommentReply, error)
	PostComment(req *api.PostCommentRequest) (*api.PostCommentReply, error)
	GetComment(req *api.GetCommentRequest) (*api.GetCommentReply, error)
	PutComment(req *api.PutCommentRequest) (*api.PutCommentReply, error)
	DeleteComment(req *api.DeleteCommentRequest) (*api.DeleteCommentReply, error)
}
