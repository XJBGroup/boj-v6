package api

import (
	"github.com/Myriad-Dreamin/boj-v6/abstract/problem"
	"github.com/Myriad-Dreamin/go-model-traits/gorm-crud-dao"
)

type ListProblemsRequest = gorm_crud_dao.Filter

type ListProblemsReply struct {
	Code int               `form:"code" json:"code"`
	Data []problem.Problem `json:"data" form:"data"`
}

type CountProblemsRequest = gorm_crud_dao.Filter

type CountProblemReply struct {
	Code int   `json:"code" form:"code"`
	Data []int `json:"data" form:"data"`
}

type PostProblemRequest struct {
}

type PostProblemReply struct {
	Code    int              `json:"code" form:"code"`
	Problem *problem.Problem `json:"problem" form:"problem"`
}

type GetProblemReply struct {
	Code    int              `json:"code" form:"code"`
	Problem *problem.Problem `json:"problem" form:"problem"`
}

type PutProblemRequest struct {
}

func PSerializeListProblemsReply(_code int, _data []problem.Problem) *ListProblemsReply {

	return &ListProblemsReply{
		Code: _code,
		Data: _data,
	}
}
func SerializeListProblemsReply(_code int, _data []problem.Problem) ListProblemsReply {

	return ListProblemsReply{
		Code: _code,
		Data: _data,
	}
}
func _packSerializeListProblemsReply(_code int, _data []problem.Problem) ListProblemsReply {

	return ListProblemsReply{
		Code: _code,
		Data: _data,
	}
}
func PackSerializeListProblemsReply(_code []int, _data [][]problem.Problem) (pack []ListProblemsReply) {
	for i := range _code {
		pack = append(pack, _packSerializeListProblemsReply(_code[i], _data[i]))
	}
	return
}
func PSerializeCountProblemReply(_code int, _data []int) *CountProblemReply {

	return &CountProblemReply{
		Code: _code,
		Data: _data,
	}
}
func SerializeCountProblemReply(_code int, _data []int) CountProblemReply {

	return CountProblemReply{
		Code: _code,
		Data: _data,
	}
}
func _packSerializeCountProblemReply(_code int, _data []int) CountProblemReply {

	return CountProblemReply{
		Code: _code,
		Data: _data,
	}
}
func PackSerializeCountProblemReply(_code []int, _data [][]int) (pack []CountProblemReply) {
	for i := range _code {
		pack = append(pack, _packSerializeCountProblemReply(_code[i], _data[i]))
	}
	return
}
func PSerializePostProblemRequest() *PostProblemRequest {

	return &PostProblemRequest{}
}
func SerializePostProblemRequest() PostProblemRequest {

	return PostProblemRequest{}
}
func _packSerializePostProblemRequest() PostProblemRequest {

	return PostProblemRequest{}
}
func PackSerializePostProblemRequest() (pack []PostProblemRequest) {
	return
}
func PSerializePostProblemReply(_code int, _problem *problem.Problem) *PostProblemReply {

	return &PostProblemReply{
		Code:    _code,
		Problem: _problem,
	}
}
func SerializePostProblemReply(_code int, _problem *problem.Problem) PostProblemReply {

	return PostProblemReply{
		Code:    _code,
		Problem: _problem,
	}
}
func _packSerializePostProblemReply(_code int, _problem *problem.Problem) PostProblemReply {

	return PostProblemReply{
		Code:    _code,
		Problem: _problem,
	}
}
func PackSerializePostProblemReply(_code []int, _problem []*problem.Problem) (pack []PostProblemReply) {
	for i := range _code {
		pack = append(pack, _packSerializePostProblemReply(_code[i], _problem[i]))
	}
	return
}
func PSerializeGetProblemReply(_code int, _problem *problem.Problem) *GetProblemReply {

	return &GetProblemReply{
		Code:    _code,
		Problem: _problem,
	}
}
func SerializeGetProblemReply(_code int, _problem *problem.Problem) GetProblemReply {

	return GetProblemReply{
		Code:    _code,
		Problem: _problem,
	}
}
func _packSerializeGetProblemReply(_code int, _problem *problem.Problem) GetProblemReply {

	return GetProblemReply{
		Code:    _code,
		Problem: _problem,
	}
}
func PackSerializeGetProblemReply(_code []int, _problem []*problem.Problem) (pack []GetProblemReply) {
	for i := range _code {
		pack = append(pack, _packSerializeGetProblemReply(_code[i], _problem[i]))
	}
	return
}
func PSerializePutProblemRequest() *PutProblemRequest {

	return &PutProblemRequest{}
}
func SerializePutProblemRequest() PutProblemRequest {

	return PutProblemRequest{}
}
func _packSerializePutProblemRequest() PutProblemRequest {

	return PutProblemRequest{}
}
func PackSerializePutProblemRequest() (pack []PutProblemRequest) {
	return
}
