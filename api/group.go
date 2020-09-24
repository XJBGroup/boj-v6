package api

import (
	"github.com/Myriad-Dreamin/boj-v6/abstract/group"
	"github.com/Myriad-Dreamin/boj-v6/abstract/user"
	"github.com/Myriad-Dreamin/go-model-traits/gorm-crud-dao"
	"time"
)

type ListGroupRequest = gorm_crud_dao.Filter

type ListGroupReply struct {
	Code int           `json:"code" form:"code"`
	Data []group.Group `json:"data" form:"data"`
}

type CountGroupRequest = gorm_crud_dao.Filter

type CountGroupReply struct {
	Code int   `json:"code" form:"code"`
	Data int64 `json:"data" form:"data"`
}

type PostGroupRequest struct {
	Name        string `json:"name" form:"name" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	OwnerName   string `json:"owner_name" form:"owner_name"`
	OwnerId     uint   `json:"owner_id" form:"owner_id"`
}

type PostGroupReply struct {
	Code int  `json:"code" form:"code"`
	Data uint `json:"data" form:"data"`
}

type PutGroupOwnerRequest struct {
	OwnerId uint `json:"owner_id" form:"owner_id" binding:"required"`
}

type PutGroupOwnerReply struct {
	Code int `json:"code" form:"code"`
}

type GetGroupMembersRequest = gorm_crud_dao.Filter

type GetGroupMembersReply struct {
	Code int                         `json:"code" form:"code"`
	Data []GetGroupMembersInnerReply `json:"data" form:"data"`
}

type GetGroupMembersInnerReply struct {
	Id                  uint      `json:"id" form:"id"`
	Gender              uint8     `json:"gender" form:"gender"`
	LastLogin           time.Time `json:"last_login" form:"last_login"`
	UserName            string    `json:"user_name" form:"user_name"`
	NickName            string    `form:"nick_name" json:"nick_name"`
	Email               string    `json:"email" form:"email"`
	Motto               string    `json:"motto" form:"motto"`
	SolvedProblemsCount int64     `json:"solved_problems_count" form:"solved_problems_count"`
	TriedProblemsCount  int64     `json:"tried_problems_count" form:"tried_problems_count"`
}

type PostGroupMemberRequest struct {
}

type PostGroupMemberReply struct {
	Code int `json:"code" form:"code"`
}

type GetGroupRequest struct {
}

type GetGroupReply struct {
	Code int          `json:"code" form:"code"`
	Data *group.Group `json:"data" form:"data"`
}

type PutGroupRequest struct {
	Name        string `json:"name" form:"name"`
	Description string `form:"description" json:"description"`
}

type PutGroupReply struct {
	Code int `json:"code" form:"code"`
}

type DeleteGroupRequest struct {
}

type DeleteGroupReply struct {
	Code int `json:"code" form:"code"`
}

func PSerializeListGroupReply(_code int, _data []group.Group) *ListGroupReply {

	return &ListGroupReply{
		Code: _code,
		Data: _data,
	}
}
func SerializeListGroupReply(_code int, _data []group.Group) ListGroupReply {

	return ListGroupReply{
		Code: _code,
		Data: _data,
	}
}
func _packSerializeListGroupReply(_code int, _data []group.Group) ListGroupReply {

	return ListGroupReply{
		Code: _code,
		Data: _data,
	}
}
func PackSerializeListGroupReply(_code []int, _data [][]group.Group) (pack []ListGroupReply) {
	for i := range _code {
		pack = append(pack, _packSerializeListGroupReply(_code[i], _data[i]))
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
func PSerializePostGroupRequest(group *group.Group, _ownerName string) *PostGroupRequest {

	return &PostGroupRequest{
		Name:        group.Name,
		Description: group.Description,
		OwnerName:   _ownerName,
		OwnerId:     group.OwnerID,
	}
}
func SerializePostGroupRequest(group *group.Group, _ownerName string) PostGroupRequest {

	return PostGroupRequest{
		Name:        group.Name,
		Description: group.Description,
		OwnerName:   _ownerName,
		OwnerId:     group.OwnerID,
	}
}
func _packSerializePostGroupRequest(group *group.Group, _ownerName string) PostGroupRequest {

	return PostGroupRequest{
		Name:        group.Name,
		Description: group.Description,
		OwnerName:   _ownerName,
		OwnerId:     group.OwnerID,
	}
}
func PackSerializePostGroupRequest(group []*group.Group, _ownerName []string) (pack []PostGroupRequest) {
	for i := range group {
		pack = append(pack, _packSerializePostGroupRequest(group[i], _ownerName[i]))
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
func PSerializePutGroupOwnerReply(_code int) *PutGroupOwnerReply {

	return &PutGroupOwnerReply{
		Code: _code,
	}
}
func SerializePutGroupOwnerReply(_code int) PutGroupOwnerReply {

	return PutGroupOwnerReply{
		Code: _code,
	}
}
func _packSerializePutGroupOwnerReply(_code int) PutGroupOwnerReply {

	return PutGroupOwnerReply{
		Code: _code,
	}
}
func PackSerializePutGroupOwnerReply(_code []int) (pack []PutGroupOwnerReply) {
	for i := range _code {
		pack = append(pack, _packSerializePutGroupOwnerReply(_code[i]))
	}
	return
}
func PSerializeGetGroupMembersReply(_code int, _data []GetGroupMembersInnerReply) *GetGroupMembersReply {

	return &GetGroupMembersReply{
		Code: _code,
		Data: _data,
	}
}
func SerializeGetGroupMembersReply(_code int, _data []GetGroupMembersInnerReply) GetGroupMembersReply {

	return GetGroupMembersReply{
		Code: _code,
		Data: _data,
	}
}
func _packSerializeGetGroupMembersReply(_code int, _data []GetGroupMembersInnerReply) GetGroupMembersReply {

	return GetGroupMembersReply{
		Code: _code,
		Data: _data,
	}
}
func PackSerializeGetGroupMembersReply(_code []int, _data [][]GetGroupMembersInnerReply) (pack []GetGroupMembersReply) {
	for i := range _code {
		pack = append(pack, _packSerializeGetGroupMembersReply(_code[i], _data[i]))
	}
	return
}
func PSerializeGetGroupMembersInnerReply(valueUser user.User) *GetGroupMembersInnerReply {

	return &GetGroupMembersInnerReply{
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
func SerializeGetGroupMembersInnerReply(valueUser user.User) GetGroupMembersInnerReply {

	return GetGroupMembersInnerReply{
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
func _packSerializeGetGroupMembersInnerReply(valueUser user.User) GetGroupMembersInnerReply {

	return GetGroupMembersInnerReply{
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
func PackSerializeGetGroupMembersInnerReply(valueUser []user.User) (pack []GetGroupMembersInnerReply) {
	for i := range valueUser {
		pack = append(pack, _packSerializeGetGroupMembersInnerReply(valueUser[i]))
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
func PSerializeGetGroupRequest() *GetGroupRequest {

	return &GetGroupRequest{}
}
func SerializeGetGroupRequest() GetGroupRequest {

	return GetGroupRequest{}
}
func _packSerializeGetGroupRequest() GetGroupRequest {

	return GetGroupRequest{}
}
func PackSerializeGetGroupRequest() (pack []GetGroupRequest) {
	return
}
func PSerializeGetGroupReply(_code int, _data *group.Group) *GetGroupReply {

	return &GetGroupReply{
		Code: _code,
		Data: _data,
	}
}
func SerializeGetGroupReply(_code int, _data *group.Group) GetGroupReply {

	return GetGroupReply{
		Code: _code,
		Data: _data,
	}
}
func _packSerializeGetGroupReply(_code int, _data *group.Group) GetGroupReply {

	return GetGroupReply{
		Code: _code,
		Data: _data,
	}
}
func PackSerializeGetGroupReply(_code []int, _data []*group.Group) (pack []GetGroupReply) {
	for i := range _code {
		pack = append(pack, _packSerializeGetGroupReply(_code[i], _data[i]))
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
func PSerializePutGroupReply(_code int) *PutGroupReply {

	return &PutGroupReply{
		Code: _code,
	}
}
func SerializePutGroupReply(_code int) PutGroupReply {

	return PutGroupReply{
		Code: _code,
	}
}
func _packSerializePutGroupReply(_code int) PutGroupReply {

	return PutGroupReply{
		Code: _code,
	}
}
func PackSerializePutGroupReply(_code []int) (pack []PutGroupReply) {
	for i := range _code {
		pack = append(pack, _packSerializePutGroupReply(_code[i]))
	}
	return
}
func PSerializeDeleteGroupRequest() *DeleteGroupRequest {

	return &DeleteGroupRequest{}
}
func SerializeDeleteGroupRequest() DeleteGroupRequest {

	return DeleteGroupRequest{}
}
func _packSerializeDeleteGroupRequest() DeleteGroupRequest {

	return DeleteGroupRequest{}
}
func PackSerializeDeleteGroupRequest() (pack []DeleteGroupRequest) {
	return
}
func PSerializeDeleteGroupReply(_code int) *DeleteGroupReply {

	return &DeleteGroupReply{
		Code: _code,
	}
}
func SerializeDeleteGroupReply(_code int) DeleteGroupReply {

	return DeleteGroupReply{
		Code: _code,
	}
}
func _packSerializeDeleteGroupReply(_code int) DeleteGroupReply {

	return DeleteGroupReply{
		Code: _code,
	}
}
func PackSerializeDeleteGroupReply(_code []int) (pack []DeleteGroupReply) {
	for i := range _code {
		pack = append(pack, _packSerializeDeleteGroupReply(_code[i]))
	}
	return
}
