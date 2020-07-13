package api

import (
	"github.com/Myriad-Dreamin/boj-v6/abstract/group"
	"github.com/Myriad-Dreamin/go-model-traits/gorm-crud-dao"
)

type ListGroupsRequest = gorm_crud_dao.Filter

type ListGroupsReply struct {
	Code int           `json:"code" form:"code"`
	Data []group.Group `json:"data" form:"data"`
}

type CountGroupsRequest = gorm_crud_dao.Filter

type CountGroupReply struct {
	Code int   `json:"code" form:"code"`
	Data []int `json:"data" form:"data"`
}

type PostGroupRequest struct {
}

type PostGroupReply struct {
	Code  int          `json:"code" form:"code"`
	Group *group.Group `form:"group" json:"group"`
}

type GetGroupReply struct {
	Code  int          `json:"code" form:"code"`
	Group *group.Group `json:"group" form:"group"`
}

type PutGroupRequest struct {
}

func PSerializeListGroupsReply(_code int, _data []group.Group) *ListGroupsReply {

	return &ListGroupsReply{
		Code: _code,
		Data: _data,
	}
}
func SerializeListGroupsReply(_code int, _data []group.Group) ListGroupsReply {

	return ListGroupsReply{
		Code: _code,
		Data: _data,
	}
}
func _packSerializeListGroupsReply(_code int, _data []group.Group) ListGroupsReply {

	return ListGroupsReply{
		Code: _code,
		Data: _data,
	}
}
func PackSerializeListGroupsReply(_code []int, _data [][]group.Group) (pack []ListGroupsReply) {
	for i := range _code {
		pack = append(pack, _packSerializeListGroupsReply(_code[i], _data[i]))
	}
	return
}
func PSerializeCountGroupReply(_code int, _data []int) *CountGroupReply {

	return &CountGroupReply{
		Code: _code,
		Data: _data,
	}
}
func SerializeCountGroupReply(_code int, _data []int) CountGroupReply {

	return CountGroupReply{
		Code: _code,
		Data: _data,
	}
}
func _packSerializeCountGroupReply(_code int, _data []int) CountGroupReply {

	return CountGroupReply{
		Code: _code,
		Data: _data,
	}
}
func PackSerializeCountGroupReply(_code []int, _data [][]int) (pack []CountGroupReply) {
	for i := range _code {
		pack = append(pack, _packSerializeCountGroupReply(_code[i], _data[i]))
	}
	return
}
func PSerializePostGroupRequest() *PostGroupRequest {

	return &PostGroupRequest{}
}
func SerializePostGroupRequest() PostGroupRequest {

	return PostGroupRequest{}
}
func _packSerializePostGroupRequest() PostGroupRequest {

	return PostGroupRequest{}
}
func PackSerializePostGroupRequest() (pack []PostGroupRequest) {
	return
}
func PSerializePostGroupReply(_code int, _group *group.Group) *PostGroupReply {

	return &PostGroupReply{
		Code:  _code,
		Group: _group,
	}
}
func SerializePostGroupReply(_code int, _group *group.Group) PostGroupReply {

	return PostGroupReply{
		Code:  _code,
		Group: _group,
	}
}
func _packSerializePostGroupReply(_code int, _group *group.Group) PostGroupReply {

	return PostGroupReply{
		Code:  _code,
		Group: _group,
	}
}
func PackSerializePostGroupReply(_code []int, _group []*group.Group) (pack []PostGroupReply) {
	for i := range _code {
		pack = append(pack, _packSerializePostGroupReply(_code[i], _group[i]))
	}
	return
}
func PSerializeGetGroupReply(_code int, _group *group.Group) *GetGroupReply {

	return &GetGroupReply{
		Code:  _code,
		Group: _group,
	}
}
func SerializeGetGroupReply(_code int, _group *group.Group) GetGroupReply {

	return GetGroupReply{
		Code:  _code,
		Group: _group,
	}
}
func _packSerializeGetGroupReply(_code int, _group *group.Group) GetGroupReply {

	return GetGroupReply{
		Code:  _code,
		Group: _group,
	}
}
func PackSerializeGetGroupReply(_code []int, _group []*group.Group) (pack []GetGroupReply) {
	for i := range _code {
		pack = append(pack, _packSerializeGetGroupReply(_code[i], _group[i]))
	}
	return
}
func PSerializePutGroupRequest() *PutGroupRequest {

	return &PutGroupRequest{}
}
func SerializePutGroupRequest() PutGroupRequest {

	return PutGroupRequest{}
}
func _packSerializePutGroupRequest() PutGroupRequest {

	return PutGroupRequest{}
}
func PackSerializePutGroupRequest() (pack []PutGroupRequest) {
	return
}
