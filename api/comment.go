
package api

import (
    "github.com/Myriad-Dreamin/boj-v6/abstract/comment"

)

type ListCommentRequest = comment.Filter

type ListCommentReply struct {
    Code int `json:"code" form:"code"`
    Data []comment.Comment `form:"data" json:"data"`
}

type CountCommentRequest = comment.Filter

type CountCommentReply struct {
    Code int `json:"code" form:"code"`
    Data int64 `form:"data" json:"data"`
}

type PostCommentRequest struct {
    Title string `json:"title" form:"title"`
    Content string `json:"content" form:"content"`
}

type PostCommentReply struct {
    Code int `json:"code" form:"code"`
    Comment *comment.Comment `json:"comment" form:"comment"`
}

type GetCommentRequest struct {

}

type GetCommentReply struct {
    Code int `json:"code" form:"code"`
    Data *comment.Comment `json:"data" form:"data"`
}

type PutCommentRequest struct {
    Title string `json:"title" form:"title"`
    Content string `form:"content" json:"content"`
}

type PutCommentReply struct {
    Code int `json:"code" form:"code"`
}

type DeleteCommentRequest struct {

}

type DeleteCommentReply struct {
    Code int `json:"code" form:"code"`
}
func PSerializeListCommentReply(_code int, _data []comment.Comment) *ListCommentReply {

    return &ListCommentReply{
        Code: _code,
        Data: _data,
    }
}
func SerializeListCommentReply(_code int, _data []comment.Comment) ListCommentReply {

    return ListCommentReply{
        Code: _code,
        Data: _data,
    }
}
func _packSerializeListCommentReply(_code int, _data []comment.Comment) ListCommentReply {

    return ListCommentReply{
        Code: _code,
        Data: _data,
    }
}
func PackSerializeListCommentReply(_code []int, _data [][]comment.Comment) (pack []ListCommentReply) {
	for i := range _code {
		pack = append(pack, _packSerializeListCommentReply(_code[i], _data[i]))
	}
	return
}
func PSerializeCountCommentReply(_code int, _data int64) *CountCommentReply {

    return &CountCommentReply{
        Code: _code,
        Data: _data,
    }
}
func SerializeCountCommentReply(_code int, _data int64) CountCommentReply {

    return CountCommentReply{
        Code: _code,
        Data: _data,
    }
}
func _packSerializeCountCommentReply(_code int, _data int64) CountCommentReply {

    return CountCommentReply{
        Code: _code,
        Data: _data,
    }
}
func PackSerializeCountCommentReply(_code []int, _data []int64) (pack []CountCommentReply) {
	for i := range _code {
		pack = append(pack, _packSerializeCountCommentReply(_code[i], _data[i]))
	}
	return
}
func PSerializePostCommentRequest(comment *comment.Comment) *PostCommentRequest {

    return &PostCommentRequest{
        Title: comment.Title,
        Content: comment.Content,
    }
}
func SerializePostCommentRequest(comment *comment.Comment) PostCommentRequest {

    return PostCommentRequest{
        Title: comment.Title,
        Content: comment.Content,
    }
}
func _packSerializePostCommentRequest(comment *comment.Comment) PostCommentRequest {

    return PostCommentRequest{
        Title: comment.Title,
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
        Code: _code,
        Comment: _comment,
    }
}
func SerializePostCommentReply(_code int, _comment *comment.Comment) PostCommentReply {

    return PostCommentReply{
        Code: _code,
        Comment: _comment,
    }
}
func _packSerializePostCommentReply(_code int, _comment *comment.Comment) PostCommentReply {

    return PostCommentReply{
        Code: _code,
        Comment: _comment,
    }
}
func PackSerializePostCommentReply(_code []int, _comment []*comment.Comment) (pack []PostCommentReply) {
	for i := range _code {
		pack = append(pack, _packSerializePostCommentReply(_code[i], _comment[i]))
	}
	return
}
func PSerializeGetCommentRequest() *GetCommentRequest {

    return &GetCommentRequest{

    }
}
func SerializeGetCommentRequest() GetCommentRequest {

    return GetCommentRequest{

    }
}
func _packSerializeGetCommentRequest() GetCommentRequest {

    return GetCommentRequest{

    }
}
func PackSerializeGetCommentRequest() (pack []GetCommentRequest) {
	return
}
func PSerializeGetCommentReply(_code int, _data *comment.Comment) *GetCommentReply {

    return &GetCommentReply{
        Code: _code,
        Data: _data,
    }
}
func SerializeGetCommentReply(_code int, _data *comment.Comment) GetCommentReply {

    return GetCommentReply{
        Code: _code,
        Data: _data,
    }
}
func _packSerializeGetCommentReply(_code int, _data *comment.Comment) GetCommentReply {

    return GetCommentReply{
        Code: _code,
        Data: _data,
    }
}
func PackSerializeGetCommentReply(_code []int, _data []*comment.Comment) (pack []GetCommentReply) {
	for i := range _code {
		pack = append(pack, _packSerializeGetCommentReply(_code[i], _data[i]))
	}
	return
}
func PSerializePutCommentRequest(comment *comment.Comment) *PutCommentRequest {

    return &PutCommentRequest{
        Title: comment.Title,
        Content: comment.Content,
    }
}
func SerializePutCommentRequest(comment *comment.Comment) PutCommentRequest {

    return PutCommentRequest{
        Title: comment.Title,
        Content: comment.Content,
    }
}
func _packSerializePutCommentRequest(comment *comment.Comment) PutCommentRequest {

    return PutCommentRequest{
        Title: comment.Title,
        Content: comment.Content,
    }
}
func PackSerializePutCommentRequest(comment []*comment.Comment) (pack []PutCommentRequest) {
	for i := range comment {
		pack = append(pack, _packSerializePutCommentRequest(comment[i]))
	}
	return
}
func PSerializePutCommentReply(_code int) *PutCommentReply {

    return &PutCommentReply{
        Code: _code,
    }
}
func SerializePutCommentReply(_code int) PutCommentReply {

    return PutCommentReply{
        Code: _code,
    }
}
func _packSerializePutCommentReply(_code int) PutCommentReply {

    return PutCommentReply{
        Code: _code,
    }
}
func PackSerializePutCommentReply(_code []int) (pack []PutCommentReply) {
	for i := range _code {
		pack = append(pack, _packSerializePutCommentReply(_code[i]))
	}
	return
}
func PSerializeDeleteCommentRequest() *DeleteCommentRequest {

    return &DeleteCommentRequest{

    }
}
func SerializeDeleteCommentRequest() DeleteCommentRequest {

    return DeleteCommentRequest{

    }
}
func _packSerializeDeleteCommentRequest() DeleteCommentRequest {

    return DeleteCommentRequest{

    }
}
func PackSerializeDeleteCommentRequest() (pack []DeleteCommentRequest) {
	return
}
func PSerializeDeleteCommentReply(_code int) *DeleteCommentReply {

    return &DeleteCommentReply{
        Code: _code,
    }
}
func SerializeDeleteCommentReply(_code int) DeleteCommentReply {

    return DeleteCommentReply{
        Code: _code,
    }
}
func _packSerializeDeleteCommentReply(_code int) DeleteCommentReply {

    return DeleteCommentReply{
        Code: _code,
    }
}
func PackSerializeDeleteCommentReply(_code []int) (pack []DeleteCommentReply) {
	for i := range _code {
		pack = append(pack, _packSerializeDeleteCommentReply(_code[i]))
	}
	return
}