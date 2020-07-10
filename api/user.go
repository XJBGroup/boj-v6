package api

import (
	"github.com/Myriad-Dreamin/boj-v6/abstract/user"
	"github.com/Myriad-Dreamin/go-model-traits/gorm-crud-dao"
)

type ListUsersRequest = gorm_crud_dao.Filter

type ListUsersReply struct {
	Code int         `json:"code" form:"code"`
	Data []user.User `json:"data" form:"data"`
}

type CountUsersRequest = gorm_crud_dao.Filter

type CountUserReply struct {
	Code int   `json:"code" form:"code"`
	Data []int `json:"data" form:"data"`
}

type ListUsersNameLikeRequest = gorm_crud_dao.Filter

type ListUsersNameLikeReply struct {
	Code int         `json:"code" form:"code"`
	Data []user.User `json:"data" form:"data"`
}

type PostUserRequest struct {
}

type PostUserReply struct {
	Code int        `json:"code" form:"code"`
	User *user.User `json:"user" form:"user"`
}

type LoginUserRequest struct {
}

type LoginUserReply struct {
	Code int        `json:"code" form:"code"`
	User *user.User `form:"user" json:"user"`
}

type InspectUserReply struct {
	Code int        `form:"code" json:"code"`
	User *user.User `json:"user" form:"user"`
}

type GetUserReply struct {
	Code int        `json:"code" form:"code"`
	User *user.User `json:"user" form:"user"`
}

type PutUserRequest struct {
}

func PSerializeListUsersReply(_code int, _data []user.User) *ListUsersReply {

	return &ListUsersReply{
		Code: _code,
		Data: _data,
	}
}
func SerializeListUsersReply(_code int, _data []user.User) ListUsersReply {

	return ListUsersReply{
		Code: _code,
		Data: _data,
	}
}
func _packSerializeListUsersReply(_code int, _data []user.User) ListUsersReply {

	return ListUsersReply{
		Code: _code,
		Data: _data,
	}
}
func PackSerializeListUsersReply(_code []int, _data [][]user.User) (pack []ListUsersReply) {
	for i := range _code {
		pack = append(pack, _packSerializeListUsersReply(_code[i], _data[i]))
	}
	return
}
func PSerializeCountUserReply(_code int, _data []int) *CountUserReply {

	return &CountUserReply{
		Code: _code,
		Data: _data,
	}
}
func SerializeCountUserReply(_code int, _data []int) CountUserReply {

	return CountUserReply{
		Code: _code,
		Data: _data,
	}
}
func _packSerializeCountUserReply(_code int, _data []int) CountUserReply {

	return CountUserReply{
		Code: _code,
		Data: _data,
	}
}
func PackSerializeCountUserReply(_code []int, _data [][]int) (pack []CountUserReply) {
	for i := range _code {
		pack = append(pack, _packSerializeCountUserReply(_code[i], _data[i]))
	}
	return
}
func PSerializeListUsersNameLikeReply(_code int, _data []user.User) *ListUsersNameLikeReply {

	return &ListUsersNameLikeReply{
		Code: _code,
		Data: _data,
	}
}
func SerializeListUsersNameLikeReply(_code int, _data []user.User) ListUsersNameLikeReply {

	return ListUsersNameLikeReply{
		Code: _code,
		Data: _data,
	}
}
func _packSerializeListUsersNameLikeReply(_code int, _data []user.User) ListUsersNameLikeReply {

	return ListUsersNameLikeReply{
		Code: _code,
		Data: _data,
	}
}
func PackSerializeListUsersNameLikeReply(_code []int, _data [][]user.User) (pack []ListUsersNameLikeReply) {
	for i := range _code {
		pack = append(pack, _packSerializeListUsersNameLikeReply(_code[i], _data[i]))
	}
	return
}
func PSerializePostUserRequest() *PostUserRequest {

	return &PostUserRequest{}
}
func SerializePostUserRequest() PostUserRequest {

	return PostUserRequest{}
}
func _packSerializePostUserRequest() PostUserRequest {

	return PostUserRequest{}
}
func PackSerializePostUserRequest() (pack []PostUserRequest) {
	return
}
func PSerializePostUserReply(_code int, _user *user.User) *PostUserReply {

	return &PostUserReply{
		Code: _code,
		User: _user,
	}
}
func SerializePostUserReply(_code int, _user *user.User) PostUserReply {

	return PostUserReply{
		Code: _code,
		User: _user,
	}
}
func _packSerializePostUserReply(_code int, _user *user.User) PostUserReply {

	return PostUserReply{
		Code: _code,
		User: _user,
	}
}
func PackSerializePostUserReply(_code []int, _user []*user.User) (pack []PostUserReply) {
	for i := range _code {
		pack = append(pack, _packSerializePostUserReply(_code[i], _user[i]))
	}
	return
}
func PSerializeLoginUserRequest() *LoginUserRequest {

	return &LoginUserRequest{}
}
func SerializeLoginUserRequest() LoginUserRequest {

	return LoginUserRequest{}
}
func _packSerializeLoginUserRequest() LoginUserRequest {

	return LoginUserRequest{}
}
func PackSerializeLoginUserRequest() (pack []LoginUserRequest) {
	return
}
func PSerializeLoginUserReply(_code int, _user *user.User) *LoginUserReply {

	return &LoginUserReply{
		Code: _code,
		User: _user,
	}
}
func SerializeLoginUserReply(_code int, _user *user.User) LoginUserReply {

	return LoginUserReply{
		Code: _code,
		User: _user,
	}
}
func _packSerializeLoginUserReply(_code int, _user *user.User) LoginUserReply {

	return LoginUserReply{
		Code: _code,
		User: _user,
	}
}
func PackSerializeLoginUserReply(_code []int, _user []*user.User) (pack []LoginUserReply) {
	for i := range _code {
		pack = append(pack, _packSerializeLoginUserReply(_code[i], _user[i]))
	}
	return
}
func PSerializeInspectUserReply(_code int, _user *user.User) *InspectUserReply {

	return &InspectUserReply{
		Code: _code,
		User: _user,
	}
}
func SerializeInspectUserReply(_code int, _user *user.User) InspectUserReply {

	return InspectUserReply{
		Code: _code,
		User: _user,
	}
}
func _packSerializeInspectUserReply(_code int, _user *user.User) InspectUserReply {

	return InspectUserReply{
		Code: _code,
		User: _user,
	}
}
func PackSerializeInspectUserReply(_code []int, _user []*user.User) (pack []InspectUserReply) {
	for i := range _code {
		pack = append(pack, _packSerializeInspectUserReply(_code[i], _user[i]))
	}
	return
}
func PSerializeGetUserReply(_code int, _user *user.User) *GetUserReply {

	return &GetUserReply{
		Code: _code,
		User: _user,
	}
}
func SerializeGetUserReply(_code int, _user *user.User) GetUserReply {

	return GetUserReply{
		Code: _code,
		User: _user,
	}
}
func _packSerializeGetUserReply(_code int, _user *user.User) GetUserReply {

	return GetUserReply{
		Code: _code,
		User: _user,
	}
}
func PackSerializeGetUserReply(_code []int, _user []*user.User) (pack []GetUserReply) {
	for i := range _code {
		pack = append(pack, _packSerializeGetUserReply(_code[i], _user[i]))
	}
	return
}
func PSerializePutUserRequest() *PutUserRequest {

	return &PutUserRequest{}
}
func SerializePutUserRequest() PutUserRequest {

	return PutUserRequest{}
}
func _packSerializePutUserRequest() PutUserRequest {

	return PutUserRequest{}
}
func PackSerializePutUserRequest() (pack []PutUserRequest) {
	return
}
