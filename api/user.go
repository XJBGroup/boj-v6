package api

import (
	"github.com/Myriad-Dreamin/boj-v6/abstract/user"
	"github.com/Myriad-Dreamin/go-model-traits/gorm-crud-dao"
	"time"
)

type ListUsersRequest = gorm_crud_dao.Filter

type ListUsersReply struct {
	Code int             `json:"code" form:"code"`
	Data []ListUserReply `json:"data" form:"data"`
}

type ListUserReply struct {
	Id                  uint      `json:"id" form:"id"`
	Gender              uint8     `json:"gender" form:"gender"`
	LastLogin           time.Time `json:"last_login" form:"last_login"`
	UserName            string    `json:"user_name" form:"user_name"`
	NickName            string    `json:"nick_name" form:"nick_name"`
	Email               string    `json:"email" form:"email"`
	Motto               string    `json:"motto" form:"motto"`
	SolvedProblemsCount int64     `json:"solved_problems_count" form:"solved_problems_count"`
	TriedProblemsCount  int64     `json:"tried_problems_count" form:"tried_problems_count"`
}

type CountUsersRequest = gorm_crud_dao.Filter

type CountUserReply struct {
	Code int   `json:"code" form:"code"`
	Data []int `json:"data" form:"data"`
}

type RegisterRequest struct {
	UserName string `json:"user_name" form:"user_name" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
	NickName string `binding:"required" json:"nick_name" form:"nick_name"`
	Gender   uint8  `json:"gender" form:"gender"`
}

type RegisterReply struct {
	Code int  `json:"code" form:"code"`
	Id   uint `json:"id" form:"id"`
}

type LoginUserRequest struct {
	Id       uint   `json:"id" form:"id"`
	UserName string `form:"user_name" json:"user_name"`
	Email    string `json:"email" form:"email"`
	Password string `binding:"required" json:"password" form:"password"`
}

type LoginUserReply struct {
	Code         int        `form:"code" json:"code"`
	User         *user.User `json:"user" form:"user"`
	RefreshToken string     `json:"refresh_token" form:"refresh_token"`
	Token        string     `json:"token" form:"token"`
	Identities   []string   `form:"identities" json:"identities"`
}

type RefreshTokenReply struct {
	Code  int    `json:"code" form:"code"`
	Token string `json:"token" form:"token"`
}

type InspectUserReply struct {
	Code int        `json:"code" form:"code"`
	User *user.User `form:"user" json:"user"`
}

type BindEmailRequest struct {
	Email string `binding:"email" json:"email" form:"email"`
}

type GetUserReply struct {
	Code int        `json:"code" form:"code"`
	User *user.User `json:"user" form:"user"`
}

type PutUserRequest struct {
	Gender   uint8  `json:"gender" form:"gender"`
	NickName string `form:"nick_name" json:"nick_name"`
	Motto    string `json:"motto" form:"motto"`
}

func PSerializeListUsersReply(_code int, _data []ListUserReply) *ListUsersReply {

	return &ListUsersReply{
		Code: _code,
		Data: _data,
	}
}
func SerializeListUsersReply(_code int, _data []ListUserReply) ListUsersReply {

	return ListUsersReply{
		Code: _code,
		Data: _data,
	}
}
func _packSerializeListUsersReply(_code int, _data []ListUserReply) ListUsersReply {

	return ListUsersReply{
		Code: _code,
		Data: _data,
	}
}
func PackSerializeListUsersReply(_code []int, _data [][]ListUserReply) (pack []ListUsersReply) {
	for i := range _code {
		pack = append(pack, _packSerializeListUsersReply(_code[i], _data[i]))
	}
	return
}
func PSerializeListUserReply(valueUser user.User) *ListUserReply {

	return &ListUserReply{
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
func SerializeListUserReply(valueUser user.User) ListUserReply {

	return ListUserReply{
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
func _packSerializeListUserReply(valueUser user.User) ListUserReply {

	return ListUserReply{
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
func PackSerializeListUserReply(valueUser []user.User) (pack []ListUserReply) {
	for i := range valueUser {
		pack = append(pack, _packSerializeListUserReply(valueUser[i]))
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
func PSerializeRegisterRequest(user *user.User) *RegisterRequest {

	return &RegisterRequest{
		UserName: user.UserName,
		Password: user.Password,
		NickName: user.NickName,
		Gender:   user.Gender,
	}
}
func SerializeRegisterRequest(user *user.User) RegisterRequest {

	return RegisterRequest{
		UserName: user.UserName,
		Password: user.Password,
		NickName: user.NickName,
		Gender:   user.Gender,
	}
}
func _packSerializeRegisterRequest(user *user.User) RegisterRequest {

	return RegisterRequest{
		UserName: user.UserName,
		Password: user.Password,
		NickName: user.NickName,
		Gender:   user.Gender,
	}
}
func PackSerializeRegisterRequest(user []*user.User) (pack []RegisterRequest) {
	for i := range user {
		pack = append(pack, _packSerializeRegisterRequest(user[i]))
	}
	return
}
func PSerializeRegisterReply(_code int, user *user.User) *RegisterReply {

	return &RegisterReply{
		Code: _code,
		Id:   user.ID,
	}
}
func SerializeRegisterReply(_code int, user *user.User) RegisterReply {

	return RegisterReply{
		Code: _code,
		Id:   user.ID,
	}
}
func _packSerializeRegisterReply(_code int, user *user.User) RegisterReply {

	return RegisterReply{
		Code: _code,
		Id:   user.ID,
	}
}
func PackSerializeRegisterReply(_code []int, user []*user.User) (pack []RegisterReply) {
	for i := range _code {
		pack = append(pack, _packSerializeRegisterReply(_code[i], user[i]))
	}
	return
}
func PSerializeLoginUserRequest(user *user.User) *LoginUserRequest {

	return &LoginUserRequest{
		Id:       user.ID,
		UserName: user.UserName,
		Email:    user.Email,
		Password: user.Password,
	}
}
func SerializeLoginUserRequest(user *user.User) LoginUserRequest {

	return LoginUserRequest{
		Id:       user.ID,
		UserName: user.UserName,
		Email:    user.Email,
		Password: user.Password,
	}
}
func _packSerializeLoginUserRequest(user *user.User) LoginUserRequest {

	return LoginUserRequest{
		Id:       user.ID,
		UserName: user.UserName,
		Email:    user.Email,
		Password: user.Password,
	}
}
func PackSerializeLoginUserRequest(user []*user.User) (pack []LoginUserRequest) {
	for i := range user {
		pack = append(pack, _packSerializeLoginUserRequest(user[i]))
	}
	return
}
func PSerializeLoginUserReply(_code int, _user *user.User, _refreshToken string, _token string, _identities []string) *LoginUserReply {

	return &LoginUserReply{
		Code:         _code,
		User:         _user,
		RefreshToken: _refreshToken,
		Token:        _token,
		Identities:   _identities,
	}
}
func SerializeLoginUserReply(_code int, _user *user.User, _refreshToken string, _token string, _identities []string) LoginUserReply {

	return LoginUserReply{
		Code:         _code,
		User:         _user,
		RefreshToken: _refreshToken,
		Token:        _token,
		Identities:   _identities,
	}
}
func _packSerializeLoginUserReply(_code int, _user *user.User, _refreshToken string, _token string, _identities []string) LoginUserReply {

	return LoginUserReply{
		Code:         _code,
		User:         _user,
		RefreshToken: _refreshToken,
		Token:        _token,
		Identities:   _identities,
	}
}
func PackSerializeLoginUserReply(_code []int, _user []*user.User, _refreshToken []string, _token []string, _identities [][]string) (pack []LoginUserReply) {
	for i := range _code {
		pack = append(pack, _packSerializeLoginUserReply(_code[i], _user[i], _refreshToken[i], _token[i], _identities[i]))
	}
	return
}
func PSerializeRefreshTokenReply(_code int, _token string) *RefreshTokenReply {

	return &RefreshTokenReply{
		Code:  _code,
		Token: _token,
	}
}
func SerializeRefreshTokenReply(_code int, _token string) RefreshTokenReply {

	return RefreshTokenReply{
		Code:  _code,
		Token: _token,
	}
}
func _packSerializeRefreshTokenReply(_code int, _token string) RefreshTokenReply {

	return RefreshTokenReply{
		Code:  _code,
		Token: _token,
	}
}
func PackSerializeRefreshTokenReply(_code []int, _token []string) (pack []RefreshTokenReply) {
	for i := range _code {
		pack = append(pack, _packSerializeRefreshTokenReply(_code[i], _token[i]))
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
func PSerializeBindEmailRequest(user *user.User) *BindEmailRequest {

	return &BindEmailRequest{
		Email: user.Email,
	}
}
func SerializeBindEmailRequest(user *user.User) BindEmailRequest {

	return BindEmailRequest{
		Email: user.Email,
	}
}
func _packSerializeBindEmailRequest(user *user.User) BindEmailRequest {

	return BindEmailRequest{
		Email: user.Email,
	}
}
func PackSerializeBindEmailRequest(user []*user.User) (pack []BindEmailRequest) {
	for i := range user {
		pack = append(pack, _packSerializeBindEmailRequest(user[i]))
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
func PSerializePutUserRequest(user *user.User) *PutUserRequest {

	return &PutUserRequest{
		Gender:   user.Gender,
		NickName: user.NickName,
		Motto:    user.Motto,
	}
}
func SerializePutUserRequest(user *user.User) PutUserRequest {

	return PutUserRequest{
		Gender:   user.Gender,
		NickName: user.NickName,
		Motto:    user.Motto,
	}
}
func _packSerializePutUserRequest(user *user.User) PutUserRequest {

	return PutUserRequest{
		Gender:   user.Gender,
		NickName: user.NickName,
		Motto:    user.Motto,
	}
}
func PackSerializePutUserRequest(user []*user.User) (pack []PutUserRequest) {
	for i := range user {
		pack = append(pack, _packSerializePutUserRequest(user[i]))
	}
	return
}
