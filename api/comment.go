package api

import (
	"github.com/Myriad-Dreamin/boj-v6/abstract/comment"
	"github.com/Myriad-Dreamin/go-model-traits/gorm-crud-dao"
)

type ListCommentsRequest = gorm_crud_dao.Filter

type ListCommentsReply struct {
	Code int               `json:"code" form:"code"`
	Data []comment.Comment `json:"data" form:"data"`
}

type CountCommentsRequest = gorm_crud_dao.Filter

type CountCommentReply struct {
	Code int   `json:"code" form:"code"`
	Data []int `form:"data" json:"data"`
}

type PostCommentRequest struct {
	Title   string `json:"title" form:"title"`
	Content string `form:"content" json:"content"`
}

type PostCommentReply struct {
	Code    int              `json:"code" form:"code"`
	Comment *comment.Comment `json:"comment" form:"comment"`
}

type GetCommentReply struct {
	Code    int              `json:"code" form:"code"`
	Comment *comment.Comment `json:"comment" form:"comment"`
}

type PutCommentRequest struct {
	Title   string `form:"title" json:"title"`
	Content string `json:"content" form:"content"`
}

func PSerializeListCommentsReply(_code int, _data []comment.Comment) *ListCommentsReply {

	return &ListCommentsReply{
		Code: _code,
		Data: _data,
	}
}
func SerializeListCommentsReply(_code int, _data []comment.Comment) ListCommentsReply {

	return ListCommentsReply{
		Code: _code,
		Data: _data,
	}
}
func _packSerializeListCommentsReply(_code int, _data []comment.Comment) ListCommentsReply {

	return ListCommentsReply{
		Code: _code,
		Data: _data,
	}
}
func PackSerializeListCommentsReply(_code []int, _data [][]comment.Comment) (pack []ListCommentsReply) {
	for i := range _code {
		pack = append(pack, _packSerializeListCommentsReply(_code[i], _data[i]))
	}
	return
}
func PSerializeCountCommentReply(_code int, _data []int) *CountCommentReply {

	return &CountCommentReply{
		Code: _code,
		Data: _data,
	}
}
func SerializeCountCommentReply(_code int, _data []int) CountCommentReply {

	return CountCommentReply{
		Code: _code,
		Data: _data,
	}
}
func _packSerializeCountCommentReply(_code int, _data []int) CountCommentReply {

	return CountCommentReply{
		Code: _code,
		Data: _data,
	}
}
func PackSerializeCountCommentReply(_code []int, _data [][]int) (pack []CountCommentReply) {
	for i := range _code {
		pack = append(pack, _packSerializeCountCommentReply(_code[i], _data[i]))
	}
	return
}
func PSerializePostCommentRequest(comment *comment.Comment) *PostCommentRequest {

	return &PostCommentRequest{
		Title:   comment.Title,
		Content: comment.Content,
	}
}
func SerializePostCommentRequest(comment *comment.Comment) PostCommentRequest {

	return PostCommentRequest{
		Title:   comment.Title,
		Content: comment.Content,
	}
}
func _packSerializePostCommentRequest(comment *comment.Comment) PostCommentRequest {

	return PostCommentRequest{
		Title:   comment.Title,
		Content: comment.Content,
	}
}
func PackSerializePostCommentRequest(comment []*comment.Comment) (pack []PostCommentRequest) {
	for i := range comment {
		pack = append(pack, _packSerializePostCommentRequest(comment[i]))
	}
	return
}
func PSerializePostCommentReply(_code int, _comment *comment.Comment) *PostCommentReply {

	return &PostCommentReply{
		Code:    _code,
		Comment: _comment,
	}
}
func SerializePostCommentReply(_code int, _comment *comment.Comment) PostCommentReply {

	return PostCommentReply{
		Code:    _code,
		Comment: _comment,
	}
}
func _packSerializePostCommentReply(_code int, _comment *comment.Comment) PostCommentReply {

	return PostCommentReply{
		Code:    _code,
		Comment: _comment,
	}
}
func PackSerializePostCommentReply(_code []int, _comment []*comment.Comment) (pack []PostCommentReply) {
	for i := range _code {
		pack = append(pack, _packSerializePostCommentReply(_code[i], _comment[i]))
	}
	return
}
func PSerializeGetCommentReply(_code int, _comment *comment.Comment) *GetCommentReply {

	return &GetCommentReply{
		Code:    _code,
		Comment: _comment,
	}
}
func SerializeGetCommentReply(_code int, _comment *comment.Comment) GetCommentReply {

	return GetCommentReply{
		Code:    _code,
		Comment: _comment,
	}
}
func _packSerializeGetCommentReply(_code int, _comment *comment.Comment) GetCommentReply {

	return GetCommentReply{
		Code:    _code,
		Comment: _comment,
	}
}
func PackSerializeGetCommentReply(_code []int, _comment []*comment.Comment) (pack []GetCommentReply) {
	for i := range _code {
		pack = append(pack, _packSerializeGetCommentReply(_code[i], _comment[i]))
	}
	return
}
func PSerializePutCommentRequest(comment *comment.Comment) *PutCommentRequest {

	return &PutCommentRequest{
		Title:   comment.Title,
		Content: comment.Content,
	}
}
func SerializePutCommentRequest(comment *comment.Comment) PutCommentRequest {

	return PutCommentRequest{
		Title:   comment.Title,
		Content: comment.Content,
	}
}
func _packSerializePutCommentRequest(comment *comment.Comment) PutCommentRequest {

	return PutCommentRequest{
		Title:   comment.Title,
		Content: comment.Content,
	}
}
func PackSerializePutCommentRequest(comment []*comment.Comment) (pack []PutCommentRequest) {
	for i := range comment {
		pack = append(pack, _packSerializePutCommentRequest(comment[i]))
	}
	return
}
