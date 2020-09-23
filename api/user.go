package api

import (
	"github.com/Myriad-Dreamin/boj-v6/abstract/user"
	"time"
)

type UserFilter struct {
	Page     int `json:"page" form:"page"`
	PageSize int `json:"page_size" form:"page_size"`
}

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
	Email               string    `form:"email" json:"email"`
	Motto               string    `json:"motto" form:"motto"`
	SolvedProblemsCount int64     `form:"solved_problems_count" json:"solved_problems_count"`
	TriedProblemsCount  int64     `json:"tried_problems_count" form:"tried_problems_count"`
}

type CountUserReply struct {
	Code int   `json:"code" form:"code"`
	Data int64 `json:"data" form:"data"`
}

type RegisterRequest struct {
	UserName string `binding:"required" json:"user_name" form:"user_name"`
	Password string `form:"password" binding:"required" json:"password"`
	NickName string `json:"nick_name" form:"nick_name" binding:"required"`
	Gender   uint8  `json:"gender" form:"gender"`
}

type RegisterReply struct {
	Code int              `json:"code" form:"code"`
	Data UserRegisterData `form:"data" json:"data"`
}

type UserRegisterData struct {
	Id uint `json:"id" form:"id"`
}

type LoginUserRequest struct {
	Id       uint   `json:"id" form:"id"`
	UserName string `json:"user_name" form:"user_name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password" binding:"required"`
}

type LoginUserReply struct {
	Code int           `json:"code" form:"code"`
	Data UserLoginData `json:"data" form:"data"`
}

type UserLoginData struct {
	Id           uint     `form:"id" json:"id"`
	RefreshToken string   `json:"refresh_token" form:"refresh_token"`
	Token        string   `json:"token" form:"token"`
	Identities   []string `json:"identities" form:"identities"`
}

type RefreshTokenReply struct {
	Code int                  `json:"code" form:"code"`
	Data UserRefreshTokenData `json:"data" form:"data"`
}

type UserRefreshTokenData struct {
	Token string `form:"token" json:"token"`
}

type ChangePasswordRequest struct {
	OldPassword string `binding:"required" json:"old_password" form:"old_password"`
	NewPassword string `json:"new_password" form:"new_password" binding:"required"`
}

type InspectUserReply struct {
	Code int        `json:"code" form:"code"`
	Data *user.User `json:"data" form:"data"`
}

type BindEmailRequest struct {
	Email string `json:"email" form:"email" binding:"email"`
}

type GetUserReply struct {
	Code int        `json:"code" form:"code"`
	Data *user.User `json:"data" form:"data"`
}

type PutUserRequest struct {
	Gender   uint8  `json:"gender" form:"gender"`
	NickName string `json:"nick_name" form:"nick_name"`
	Motto    string `json:"motto" form:"motto"`
}

func PSerializeUserFilter(_page int, _pageSize int) *UserFilter {

	return &UserFilter{
		Page:     _page,
		PageSize: _pageSize,
	}
}
func SerializeUserFilter(_page int, _pageSize int) UserFilter {

	return UserFilter{
		Page:     _page,
		PageSize: _pageSize,
	}
}
func _packSerializeUserFilter(_page int, _pageSize int) UserFilter {

	return UserFilter{
		Page:     _page,
		PageSize: _pageSize,
	}
}
func PackSerializeUserFilter(_page []int, _pageSize []int) (pack []UserFilter) {
	for i := range _page {
		pack = append(pack, _packSerializeUserFilter(_page[i], _pageSize[i]))
	}
	return
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
func PSerializeCountUserReply(_code int, _data int64) *CountUserReply {

	return &CountUserReply{
		Code: _code,
		Data: _data,
	}
}
func SerializeCountUserReply(_code int, _data int64) CountUserReply {

	return CountUserReply{
		Code: _code,
		Data: _data,
	}
}
func _packSerializeCountUserReply(_code int, _data int64) CountUserReply {

	return CountUserReply{
		Code: _code,
		Data: _data,
	}
}
func PackSerializeCountUserReply(_code []int, _data []int64) (pack []CountUserReply) {
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
func PSerializeRegisterReply(_code int, _data UserRegisterData) *RegisterReply {

	return &RegisterReply{
		Code: _code,
		Data: _data,
	}
}
func SerializeRegisterReply(_code int, _data UserRegisterData) RegisterReply {

	return RegisterReply{
		Code: _code,
		Data: _data,
	}
}
func _packSerializeRegisterReply(_code int, _data UserRegisterData) RegisterReply {

	return RegisterReply{
		Code: _code,
		Data: _data,
	}
}
func PackSerializeRegisterReply(_code []int, _data []UserRegisterData) (pack []RegisterReply) {
	for i := range _code {
		pack = append(pack, _packSerializeRegisterReply(_code[i], _data[i]))
	}
	return
}
func PSerializeUserRegisterData(user *user.User) *UserRegisterData {

	return &UserRegisterData{
		Id: user.ID,
	}
}
func SerializeUserRegisterData(user *user.User) UserRegisterData {

	return UserRegisterData{
		Id: user.ID,
	}
}
func _packSerializeUserRegisterData(user *user.User) UserRegisterData {

	return UserRegisterData{
		Id: user.ID,
	}
}
func PackSerializeUserRegisterData(user []*user.User) (pack []UserRegisterData) {
	for i := range user {
		pack = append(pack, _packSerializeUserRegisterData(user[i]))
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
func PSerializeLoginUserReply(_code int, _data UserLoginData) *LoginUserReply {

	return &LoginUserReply{
		Code: _code,
		Data: _data,
	}
}
func SerializeLoginUserReply(_code int, _data UserLoginData) LoginUserReply {

	return LoginUserReply{
		Code: _code,
		Data: _data,
	}
}
func _packSerializeLoginUserReply(_code int, _data UserLoginData) LoginUserReply {

	return LoginUserReply{
		Code: _code,
		Data: _data,
	}
}
func PackSerializeLoginUserReply(_code []int, _data []UserLoginData) (pack []LoginUserReply) {
	for i := range _code {
		pack = append(pack, _packSerializeLoginUserReply(_code[i], _data[i]))
	}
	return
}
func PSerializeUserLoginData(user *user.User, _refreshToken string, _token string, _identities []string) *UserLoginData {

	return &UserLoginData{
		Id:           user.ID,
		RefreshToken: _refreshToken,
		Token:        _token,
		Identities:   _identities,
	}
}
func SerializeUserLoginData(user *user.User, _refreshToken string, _token string, _identities []string) UserLoginData {

	return UserLoginData{
		Id:           user.ID,
		RefreshToken: _refreshToken,
		Token:        _token,
		Identities:   _identities,
	}
}
func _packSerializeUserLoginData(user *user.User, _refreshToken string, _token string, _identities []string) UserLoginData {

	return UserLoginData{
		Id:           user.ID,
		RefreshToken: _refreshToken,
		Token:        _token,
		Identities:   _identities,
	}
}
func PackSerializeUserLoginData(user []*user.User, _refreshToken []string, _token []string, _identities [][]string) (pack []UserLoginData) {
	for i := range user {
		pack = append(pack, _packSerializeUserLoginData(user[i], _refreshToken[i], _token[i], _identities[i]))
	}
	return
}
func PSerializeRefreshTokenReply(_code int, _data UserRefreshTokenData) *RefreshTokenReply {

	return &RefreshTokenReply{
		Code: _code,
		Data: _data,
	}
}
func SerializeRefreshTokenReply(_code int, _data UserRefreshTokenData) RefreshTokenReply {

	return RefreshTokenReply{
		Code: _code,
		Data: _data,
	}
}
func _packSerializeRefreshTokenReply(_code int, _data UserRefreshTokenData) RefreshTokenReply {

	return RefreshTokenReply{
		Code: _code,
		Data: _data,
	}
}
func PackSerializeRefreshTokenReply(_code []int, _data []UserRefreshTokenData) (pack []RefreshTokenReply) {
	for i := range _code {
		pack = append(pack, _packSerializeRefreshTokenReply(_code[i], _data[i]))
	}
	return
}
func PSerializeUserRefreshTokenData(_token string) *UserRefreshTokenData {

	return &UserRefreshTokenData{
		Token: _token,
	}
}
func SerializeUserRefreshTokenData(_token string) UserRefreshTokenData {

	return UserRefreshTokenData{
		Token: _token,
	}
}
func _packSerializeUserRefreshTokenData(_token string) UserRefreshTokenData {

	return UserRefreshTokenData{
		Token: _token,
	}
}
func PackSerializeUserRefreshTokenData(_token []string) (pack []UserRefreshTokenData) {
	for i := range _token {
		pack = append(pack, _packSerializeUserRefreshTokenData(_token[i]))
	}
	return
}
func PSerializeChangePasswordRequest(_oldPassword string, _newPassword string) *ChangePasswordRequest {

	return &ChangePasswordRequest{
		OldPassword: _oldPassword,
		NewPassword: _newPassword,
	}
}
func SerializeChangePasswordRequest(_oldPassword string, _newPassword string) ChangePasswordRequest {

	return ChangePasswordRequest{
		OldPassword: _oldPassword,
		NewPassword: _newPassword,
	}
}
func _packSerializeChangePasswordRequest(_oldPassword string, _newPassword string) ChangePasswordRequest {

	return ChangePasswordRequest{
		OldPassword: _oldPassword,
		NewPassword: _newPassword,
	}
}
func PackSerializeChangePasswordRequest(_oldPassword []string, _newPassword []string) (pack []ChangePasswordRequest) {
	for i := range _oldPassword {
		pack = append(pack, _packSerializeChangePasswordRequest(_oldPassword[i], _newPassword[i]))
	}
	return
}
func PSerializeInspectUserReply(_code int, _data *user.User) *InspectUserReply {

	return &InspectUserReply{
		Code: _code,
		Data: _data,
	}
}
func SerializeInspectUserReply(_code int, _data *user.User) InspectUserReply {

	return InspectUserReply{
		Code: _code,
		Data: _data,
	}
}
func _packSerializeInspectUserReply(_code int, _data *user.User) InspectUserReply {

	return InspectUserReply{
		Code: _code,
		Data: _data,
	}
}
func PackSerializeInspectUserReply(_code []int, _data []*user.User) (pack []InspectUserReply) {
	for i := range _code {
		pack = append(pack, _packSerializeInspectUserReply(_code[i], _data[i]))
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
func PSerializeGetUserReply(_code int, _data *user.User) *GetUserReply {

	return &GetUserReply{
		Code: _code,
		Data: _data,
	}
}
func SerializeGetUserReply(_code int, _data *user.User) GetUserReply {

	return GetUserReply{
		Code: _code,
		Data: _data,
	}
}
func _packSerializeGetUserReply(_code int, _data *user.User) GetUserReply {

	return GetUserReply{
		Code: _code,
		Data: _data,
	}
}
func PackSerializeGetUserReply(_code []int, _data []*user.User) (pack []GetUserReply) {
	for i := range _code {
		pack = append(pack, _packSerializeGetUserReply(_code[i], _data[i]))
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
