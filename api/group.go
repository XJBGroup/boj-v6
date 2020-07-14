package api

import (
	"github.com/Myriad-Dreamin/boj-v6/abstract/group"
	"github.com/Myriad-Dreamin/boj-v6/abstract/user"
	"github.com/Myriad-Dreamin/go-model-traits/gorm-crud-dao"
	"time"
)

type ListGroupsRequest = gorm_crud_dao.Filter

type ListGroupsReply struct {
	Code int           `json:"code" form:"code"`
	Data []group.Group `json:"data" form:"data"`
}

type CountGroupsRequest = gorm_crud_dao.Filter

type CountGroupReply struct {
	Code int   `json:"code" form:"code"`
	Data int64 `json:"data" form:"data"`
}

type PostGroupRequest struct {
	Name        string `binding:"required" json:"name" form:"name"`
	Description string `json:"description" form:"description" binding:"required"`
	OwnerId     uint   `json:"owner_id" form:"owner_id" binding:"required"`
}

type PostGroupReply struct {
	Code int  `json:"code" form:"code"`
	Data uint `json:"data" form:"data"`
}

type PutGroupOwnerRequest struct {
	OwnerId uint `binding:"required" json:"owner_id" form:"owner_id"`
}

type GroupUserListRequest = gorm_crud_dao.Filter

type GetGroupMembersReply struct {
	Code int                  `json:"code" form:"code"`
	Data []ListGroupUserReply `form:"data" json:"data"`
}

type ListGroupUserReply struct {
	Id                  uint      `json:"id" form:"id"`
	Gender              uint8     `json:"gender" form:"gender"`
	LastLogin           time.Time `json:"last_login" form:"last_login"`
	UserName            string    `json:"user_name" form:"user_name"`
	NickName            string    `json:"nick_name" form:"nick_name"`
	Email               string    `json:"email" form:"email"`
	Motto               string    `json:"motto" form:"motto"`
	SolvedProblemsCount int64     `form:"solved_problems_count" json:"solved_problems_count"`
	TriedProblemsCount  int64     `json:"tried_problems_count" form:"tried_problems_count"`
}

type PostGroupMemberRequest struct {
}

type PostGroupMemberReply struct {
	Code int `json:"code" form:"code"`
}

type GetGroupReply struct {
	Code  int          `json:"code" form:"code"`
	Group *group.Group `json:"group" form:"group"`
}

type PutGroupRequest struct {
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
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
func PSerializeCountGroupReply(_code int, _data int64) *CountGroupReply {

	return &CountGroupReply{
		Code: _code,
		Data: _data,
	}
}
func SerializeCountGroupReply(_code int, _data int64) CountGroupReply {

	return CountGroupReply{
		Code: _code,
		Data: _data,
	}
}
func _packSerializeCountGroupReply(_code int, _data int64) CountGroupReply {

	return CountGroupReply{
		Code: _code,
		Data: _data,
	}
}
func PackSerializeCountGroupReply(_code []int, _data []int64) (pack []CountGroupReply) {
	for i := range _code {
		pack = append(pack, _packSerializeCountGroupReply(_code[i], _data[i]))
	}
	return
}
func PSerializePostGroupRequest(group *group.Group) *PostGroupRequest {

	return &PostGroupRequest{
		Name:        group.Name,
		Description: group.Description,
		OwnerId:     group.OwnerID,
	}
}
func SerializePostGroupRequest(group *group.Group) PostGroupRequest {

	return PostGroupRequest{
		Name:        group.Name,
		Description: group.Description,
		OwnerId:     group.OwnerID,
	}
}
func _packSerializePostGroupRequest(group *group.Group) PostGroupRequest {

	return PostGroupRequest{
		Name:        group.Name,
		Description: group.Description,
		OwnerId:     group.OwnerID,
	}
}
func PackSerializePostGroupRequest(group []*group.Group) (pack []PostGroupRequest) {
	for i := range group {
		pack = append(pack, _packSerializePostGroupRequest(group[i]))
	}
	return
}
func PSerializePostGroupReply(_code int, group *group.Group) *PostGroupReply {

	return &PostGroupReply{
		Code: _code,
		Data: group.ID,
	}
}
func SerializePostGroupReply(_code int, group *group.Group) PostGroupReply {

	return PostGroupReply{
		Code: _code,
		Data: group.ID,
	}
}
func _packSerializePostGroupReply(_code int, group *group.Group) PostGroupReply {

	return PostGroupReply{
		Code: _code,
		Data: group.ID,
	}
}
func PackSerializePostGroupReply(_code []int, group []*group.Group) (pack []PostGroupReply) {
	for i := range _code {
		pack = append(pack, _packSerializePostGroupReply(_code[i], group[i]))
	}
	return
}
func PSerializePutGroupOwnerRequest(group *group.Group) *PutGroupOwnerRequest {

	return &PutGroupOwnerRequest{
		OwnerId: group.OwnerID,
	}
}
func SerializePutGroupOwnerRequest(group *group.Group) PutGroupOwnerRequest {

	return PutGroupOwnerRequest{
		OwnerId: group.OwnerID,
	}
}
func _packSerializePutGroupOwnerRequest(group *group.Group) PutGroupOwnerRequest {

	return PutGroupOwnerRequest{
		OwnerId: group.OwnerID,
	}
}
func PackSerializePutGroupOwnerRequest(group []*group.Group) (pack []PutGroupOwnerRequest) {
	for i := range group {
		pack = append(pack, _packSerializePutGroupOwnerRequest(group[i]))
	}
	return
}
func PSerializeGetGroupMembersReply(_code int, _data []ListGroupUserReply) *GetGroupMembersReply {

	return &GetGroupMembersReply{
		Code: _code,
		Data: _data,
	}
}
func SerializeGetGroupMembersReply(_code int, _data []ListGroupUserReply) GetGroupMembersReply {

	return GetGroupMembersReply{
		Code: _code,
		Data: _data,
	}
}
func _packSerializeGetGroupMembersReply(_code int, _data []ListGroupUserReply) GetGroupMembersReply {

	return GetGroupMembersReply{
		Code: _code,
		Data: _data,
	}
}
func PackSerializeGetGroupMembersReply(_code []int, _data [][]ListGroupUserReply) (pack []GetGroupMembersReply) {
	for i := range _code {
		pack = append(pack, _packSerializeGetGroupMembersReply(_code[i], _data[i]))
	}
	return
}
func PSerializeListGroupUserReply(valueUser user.User) *ListGroupUserReply {

	return &ListGroupUserReply{
		Id:                  valueUser.ID,
		Gender:              valueUser.Gender,
		LastLogin:           valueUser.LastLogin,
		UserName:            valueUser.UserName,
		NickName:            valueUser.NickName,
		Email:               valueUser.Email,
		Motto:               valueUser.Motto,
		SolvedProblemsCount: valueUser.SolvedProblemsCount,
		TriedProblemsCount:  valueUser.TriedProblemsCount,
	}
}
func SerializeListGroupUserReply(valueUser user.User) ListGroupUserReply {

	return ListGroupUserReply{
		Id:                  valueUser.ID,
		Gender:              valueUser.Gender,
		LastLogin:           valueUser.LastLogin,
		UserName:            valueUser.UserName,
		NickName:            valueUser.NickName,
		Email:               valueUser.Email,
		Motto:               valueUser.Motto,
		SolvedProblemsCount: valueUser.SolvedProblemsCount,
		TriedProblemsCount:  valueUser.TriedProblemsCount,
	}
}
func _packSerializeListGroupUserReply(valueUser user.User) ListGroupUserReply {

	return ListGroupUserReply{
		Id:                  valueUser.ID,
		Gender:              valueUser.Gender,
		LastLogin:           valueUser.LastLogin,
		UserName:            valueUser.UserName,
		NickName:            valueUser.NickName,
		Email:               valueUser.Email,
		Motto:               valueUser.Motto,
		SolvedProblemsCount: valueUser.SolvedProblemsCount,
		TriedProblemsCount:  valueUser.TriedProblemsCount,
	}
}
func PackSerializeListGroupUserReply(valueUser []user.User) (pack []ListGroupUserReply) {
	for i := range valueUser {
		pack = append(pack, _packSerializeListGroupUserReply(valueUser[i]))
	}
	return
}
func PSerializePostGroupMemberRequest() *PostGroupMemberRequest {

	return &PostGroupMemberRequest{}
}
func SerializePostGroupMemberRequest() PostGroupMemberRequest {

	return PostGroupMemberRequest{}
}
func _packSerializePostGroupMemberRequest() PostGroupMemberRequest {

	return PostGroupMemberRequest{}
}
func PackSerializePostGroupMemberRequest() (pack []PostGroupMemberRequest) {
	return
}
func PSerializePostGroupMemberReply(_code int) *PostGroupMemberReply {

	return &PostGroupMemberReply{
		Code: _code,
	}
}
func SerializePostGroupMemberReply(_code int) PostGroupMemberReply {

	return PostGroupMemberReply{
		Code: _code,
	}
}
func _packSerializePostGroupMemberReply(_code int) PostGroupMemberReply {

	return PostGroupMemberReply{
		Code: _code,
	}
}
func PackSerializePostGroupMemberReply(_code []int) (pack []PostGroupMemberReply) {
	for i := range _code {
		pack = append(pack, _packSerializePostGroupMemberReply(_code[i]))
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
func PSerializePutGroupRequest(group *group.Group) *PutGroupRequest {

	return &PutGroupRequest{
		Name:        group.Name,
		Description: group.Description,
	}
}
func SerializePutGroupRequest(group *group.Group) PutGroupRequest {

	return PutGroupRequest{
		Name:        group.Name,
		Description: group.Description,
	}
}
func _packSerializePutGroupRequest(group *group.Group) PutGroupRequest {

	return PutGroupRequest{
		Name:        group.Name,
		Description: group.Description,
	}
}
func PackSerializePutGroupRequest(group []*group.Group) (pack []PutGroupRequest) {
	for i := range group {
		pack = append(pack, _packSerializePutGroupRequest(group[i]))
	}
	return
}
