package api

import (
	"github.com/Myriad-Dreamin/boj-v6/abstract/contest"
	"github.com/Myriad-Dreamin/go-model-traits/gorm-crud-dao"
)

type ListContestsRequest = gorm_crud_dao.Filter

type ListContestsReply struct {
	Code int               `json:"code" form:"code"`
	Data []contest.Contest `json:"data" form:"data"`
}

type CountContestsRequest = gorm_crud_dao.Filter

type CountContestReply struct {
	Code int   `json:"code" form:"code"`
	Data []int `json:"data" form:"data"`
}

type PostContestRequest struct {
}

type PostContestReply struct {
	Code    int              `json:"code" form:"code"`
	Contest *contest.Contest `json:"contest" form:"contest"`
}

type GetContestReply struct {
	Code    int              `json:"code" form:"code"`
	Contest *contest.Contest `json:"contest" form:"contest"`
}

type PutContestRequest struct {
}

func PSerializeListContestsReply(_code int, _data []contest.Contest) *ListContestsReply {

	return &ListContestsReply{
		Code: _code,
		Data: _data,
	}
}
func SerializeListContestsReply(_code int, _data []contest.Contest) ListContestsReply {

	return ListContestsReply{
		Code: _code,
		Data: _data,
	}
}
func _packSerializeListContestsReply(_code int, _data []contest.Contest) ListContestsReply {

	return ListContestsReply{
		Code: _code,
		Data: _data,
	}
}
func PackSerializeListContestsReply(_code []int, _data [][]contest.Contest) (pack []ListContestsReply) {
	for i := range _code {
		pack = append(pack, _packSerializeListContestsReply(_code[i], _data[i]))
	}
	return
}
func PSerializeCountContestReply(_code int, _data []int) *CountContestReply {

	return &CountContestReply{
		Code: _code,
		Data: _data,
	}
}
func SerializeCountContestReply(_code int, _data []int) CountContestReply {

	return CountContestReply{
		Code: _code,
		Data: _data,
	}
}
func _packSerializeCountContestReply(_code int, _data []int) CountContestReply {

	return CountContestReply{
		Code: _code,
		Data: _data,
	}
}
func PackSerializeCountContestReply(_code []int, _data [][]int) (pack []CountContestReply) {
	for i := range _code {
		pack = append(pack, _packSerializeCountContestReply(_code[i], _data[i]))
	}
	return
}
func PSerializePostContestRequest() *PostContestRequest {

	return &PostContestRequest{}
}
func SerializePostContestRequest() PostContestRequest {

	return PostContestRequest{}
}
func _packSerializePostContestRequest() PostContestRequest {

	return PostContestRequest{}
}
func PackSerializePostContestRequest() (pack []PostContestRequest) {
	return
}
func PSerializePostContestReply(_code int, _contest *contest.Contest) *PostContestReply {

	return &PostContestReply{
		Code:    _code,
		Contest: _contest,
	}
}
func SerializePostContestReply(_code int, _contest *contest.Contest) PostContestReply {

	return PostContestReply{
		Code:    _code,
		Contest: _contest,
	}
}
func _packSerializePostContestReply(_code int, _contest *contest.Contest) PostContestReply {

	return PostContestReply{
		Code:    _code,
		Contest: _contest,
	}
}
func PackSerializePostContestReply(_code []int, _contest []*contest.Contest) (pack []PostContestReply) {
	for i := range _code {
		pack = append(pack, _packSerializePostContestReply(_code[i], _contest[i]))
	}
	return
}
func PSerializeGetContestReply(_code int, _contest *contest.Contest) *GetContestReply {

	return &GetContestReply{
		Code:    _code,
		Contest: _contest,
	}
}
func SerializeGetContestReply(_code int, _contest *contest.Contest) GetContestReply {

	return GetContestReply{
		Code:    _code,
		Contest: _contest,
	}
}
func _packSerializeGetContestReply(_code int, _contest *contest.Contest) GetContestReply {

	return GetContestReply{
		Code:    _code,
		Contest: _contest,
	}
}
func PackSerializeGetContestReply(_code []int, _contest []*contest.Contest) (pack []GetContestReply) {
	for i := range _code {
		pack = append(pack, _packSerializeGetContestReply(_code[i], _contest[i]))
	}
	return
}
func PSerializePutContestRequest() *PutContestRequest {

	return &PutContestRequest{}
}
func SerializePutContestRequest() PutContestRequest {

	return PutContestRequest{}
}
func _packSerializePutContestRequest() PutContestRequest {

	return PutContestRequest{}
}
func PackSerializePutContestRequest() (pack []PutContestRequest) {
	return
}
