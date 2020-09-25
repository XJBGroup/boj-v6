package api

import (
	"github.com/Myriad-Dreamin/boj-v6/abstract/user"
	"time"
)

type ListUserRequest struct {
	Page     int `json:"page" form:"page"`
	PageSize int `json:"page_size" form:"page_size"`
}

type ListUserReply struct {
	Code int                  `json:"code" form:"code"`
	Data []ListUserInnerReply `json:"data" form:"data"`
}

type ListUserInnerReply struct {
	Id                  uint      `json:"id" form:"id"`
	Gender              uint8     `json:"gender" form:"gender"`
	LastLogin           time.Time `form:"last_login" json:"last_login"`
	UserName            string    `json:"user_name" form:"user_name"`
	NickName            string    `json:"nick_name" form:"nick_name"`
	Email               string    `json:"email" form:"email"`
	Motto               string    `json:"motto" form:"motto"`
	SolvedProblemsCount int64     `json:"solved_problems_count" form:"solved_problems_count"`
	TriedProblemsCount  int64     `json:"tried_problems_count" form:"tried_problems_count"`
}

type CountUserRequest struct {
	Page     int `json:"page" form:"page"`
	PageSize int `json:"page_size" form:"page_size"`
}

type CountUserReply struct {
	Code int   `json:"code" form:"code"`
	Data int64 `json:"data" form:"data"`
}

type RegisterRequest struct {
	UserName string `json:"user_name" form:"user_name" binding:"required"`
	Password string `form:"password" binding:"required" json:"password"`
	NickName string `json:"nick_name" form:"nick_name" binding:"required"`
	Gender   uint8  `json:"gender" form:"gender"`
}

type RegisterReply struct {
	Code int              `json:"code" form:"code"`
	Data UserRegisterData `json:"data" form:"data"`
}

type UserRegisterData struct {
	Id uint `json:"id" form:"id"`
}

type LoginUserRequest struct {
	Id       uint   `json:"id" form:"id"`
	UserName string `form:"user_name" json:"user_name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password" binding:"required"`
}

type LoginUserReply struct {
	Code int           `json:"code" form:"code"`
	Data UserLoginData `json:"data" form:"data"`
}

type UserLoginData struct {
	Id           uint     `form:"id" json:"id"`
	Email        string   `json:"email" form:"email"`
	UserName     string   `json:"user_name" form:"user_name"`
	NickName     string   `json:"nick_name" form:"nick_name"`
	RefreshToken string   `json:"refresh_token" form:"refresh_token"`
	Token        string   `json:"token" form:"token"`
	Identities   []string `json:"identities" form:"identities"`
}

type RefreshTokenRequest struct {
}

type RefreshTokenReply struct {
	Code int                  `json:"code" form:"code"`
	Data UserRefreshTokenData `json:"data" form:"data"`
}

type UserRefreshTokenData struct {
	Token string `form:"token" json:"token"`
}

type BindEmailRequest struct {
	Email string `json:"email" form:"email" binding:"email"`
}

type BindEmailReply struct {
	Code int `json:"code" form:"code"`
}

type ChangePasswordRequest struct {
	OldPassword string `json:"old_password" form:"old_password" binding:"required"`
	NewPassword string `form:"new_password" binding:"required" json:"new_password"`
}

type ChangePasswordReply struct {
	Code int `json:"code" form:"code"`
}

type InspectUserRequest struct {
}

type InspectUserReply struct {
	Code int                   `json:"code" form:"code"`
	Data InspectUserInnerReply `form:"data" json:"data"`
}

type InspectUserInnerReply struct {
	Id              uint      `json:"id" form:"id"`
	NickName        string    `json:"nick_name" form:"nick_name"`
	UserName        string    `json:"user_name" form:"user_name"`
	LastLogin       time.Time `json:"last_login" form:"last_login"`
	Email           string    `json:"email" form:"email"`
	Motto           string    `json:"motto" form:"motto"`
	Gender          uint8     `json:"gender" form:"gender"`
	Identities      []string  `form:"identities" json:"identities"`
	SuccessProblems []uint    `json:"success_problems" form:"success_problems"`
	TriedProblems   []uint    `json:"tried_problems" form:"tried_problems"`
}

type GetUserRequest struct {
}

type GetUserReply struct {
	Code int               `json:"code" form:"code"`
	Data GetUserInnerReply `json:"data" form:"data"`
}

type GetUserInnerReply struct {
	Id        uint      `json:"id" form:"id"`
	NickName  string    `json:"nick_name" form:"nick_name"`
	LastLogin time.Time `json:"last_login" form:"last_login"`
	Motto     string    `json:"motto" form:"motto"`
	Gender    uint8     `form:"gender" json:"gender"`
}

type PutUserRequest struct {
	Gender   uint8  `json:"gender" form:"gender"`
	NickName string `json:"nick_name" form:"nick_name"`
	Motto    string `json:"motto" form:"motto"`
}

type PutUserReply struct {
	Code int `json:"code" form:"code"`
}

type DeleteUserRequest struct {
}

type DeleteUserReply struct {
	Code int `json:"code" form:"code"`
}

func PSerializeListUserRequest(_page int, _pageSize int) *ListUserRequest {

	return &ListUserRequest{
		Page:     _page,
		PageSize: _pageSize,
	}
}
func SerializeListUserRequest(_page int, _pageSize int) ListUserRequest {

	return ListUserRequest{
		Page:     _page,
		PageSize: _pageSize,
	}
}
func _packSerializeListUserRequest(_page int, _pageSize int) ListUserRequest {

	return ListUserRequest{
		Page:     _page,
		PageSize: _pageSize,
	}
}
func PackSerializeListUserRequest(_page []int, _pageSize []int) (pack []ListUserRequest) {
	for i := range _page {
		pack = append(pack, _packSerializeListUserRequest(_page[i], _pageSize[i]))
	}
	return
}
func PSerializeListUserReply(_code int, _data []ListUserInnerReply) *ListUserReply {

	return &ListUserReply{
		Code: _code,
		Data: _data,
	}
}
func SerializeListUserReply(_code int, _data []ListUserInnerReply) ListUserReply {

	return ListUserReply{
		Code: _code,
		Data: _data,
	}
}
func _packSerializeListUserReply(_code int, _data []ListUserInnerReply) ListUserReply {

	return ListUserReply{
		Code: _code,
		Data: _data,
	}
}
func PackSerializeListUserReply(_code []int, _data [][]ListUserInnerReply) (pack []ListUserReply) {
	for i := range _code {
		pack = append(pack, _packSerializeListUserReply(_code[i], _data[i]))
	}
	return
}
func PSerializeListUserInnerReply(valueUser user.User) *ListUserInnerReply {

	return &ListUserInnerReply{
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
func SerializeListUserInnerReply(valueUser user.User) ListUserInnerReply {

	return ListUserInnerReply{
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
func _packSerializeListUserInnerReply(valueUser user.User) ListUserInnerReply {

	return ListUserInnerReply{
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
func PackSerializeListUserInnerReply(valueUser []user.User) (pack []ListUserInnerReply) {
	for i := range valueUser {
		pack = append(pack, _packSerializeListUserInnerReply(valueUser[i]))
	}
	return
}
func PSerializeCountUserRequest(_page int, _pageSize int) *CountUserRequest {

	return &CountUserRequest{
		Page:     _page,
		PageSize: _pageSize,
	}
}
func SerializeCountUserRequest(_page int, _pageSize int) CountUserRequest {

	return CountUserRequest{
		Page:     _page,
		PageSize: _pageSize,
	}
}
func _packSerializeCountUserRequest(_page int, _pageSize int) CountUserRequest {

	return CountUserRequest{
		Page:     _page,
		PageSize: _pageSize,
	}
}
func PackSerializeCountUserRequest(_page []int, _pageSize []int) (pack []CountUserRequest) {
	for i := range _page {
		pack = append(pack, _packSerializeCountUserRequest(_page[i], _pageSize[i]))
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
		Email:        user.Email,
		UserName:     user.UserName,
		NickName:     user.NickName,
		RefreshToken: _refreshToken,
		Token:        _token,
		Identities:   _identities,
	}
}
func SerializeUserLoginData(user *user.User, _refreshToken string, _token string, _identities []string) UserLoginData {

	return UserLoginData{
		Id:           user.ID,
		Email:        user.Email,
		UserName:     user.UserName,
		NickName:     user.NickName,
		RefreshToken: _refreshToken,
		Token:        _token,
		Identities:   _identities,
	}
}
func _packSerializeUserLoginData(user *user.User, _refreshToken string, _token string, _identities []string) UserLoginData {

	return UserLoginData{
		Id:           user.ID,
		Email:        user.Email,
		UserName:     user.UserName,
		NickName:     user.NickName,
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
func PSerializeRefreshTokenRequest() *RefreshTokenRequest {

	return &RefreshTokenRequest{}
}
func SerializeRefreshTokenRequest() RefreshTokenRequest {

	return RefreshTokenRequest{}
}
func _packSerializeRefreshTokenRequest() RefreshTokenRequest {

	return RefreshTokenRequest{}
}
func PackSerializeRefreshTokenRequest() (pack []RefreshTokenRequest) {
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
func PSerializeBindEmailReply(_code int) *BindEmailReply {

	return &BindEmailReply{
		Code: _code,
	}
}
func SerializeBindEmailReply(_code int) BindEmailReply {

	return BindEmailReply{
		Code: _code,
	}
}
func _packSerializeBindEmailReply(_code int) BindEmailReply {

	return BindEmailReply{
		Code: _code,
	}
}
func PackSerializeBindEmailReply(_code []int) (pack []BindEmailReply) {
	for i := range _code {
		pack = append(pack, _packSerializeBindEmailReply(_code[i]))
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
func PSerializeChangePasswordReply(_code int) *ChangePasswordReply {

	return &ChangePasswordReply{
		Code: _code,
	}
}
func SerializeChangePasswordReply(_code int) ChangePasswordReply {

	return ChangePasswordReply{
		Code: _code,
	}
}
func _packSerializeChangePasswordReply(_code int) ChangePasswordReply {

	return ChangePasswordReply{
		Code: _code,
	}
}
func PackSerializeChangePasswordReply(_code []int) (pack []ChangePasswordReply) {
	for i := range _code {
		pack = append(pack, _packSerializeChangePasswordReply(_code[i]))
	}
	return
}
func PSerializeInspectUserRequest() *InspectUserRequest {

	return &InspectUserRequest{}
}
func SerializeInspectUserRequest() InspectUserRequest {

	return InspectUserRequest{}
}
func _packSerializeInspectUserRequest() InspectUserRequest {

	return InspectUserRequest{}
}
func PackSerializeInspectUserRequest() (pack []InspectUserRequest) {
	return
}
func PSerializeInspectUserReply(_code int, _data InspectUserInnerReply) *InspectUserReply {

	return &InspectUserReply{
		Code: _code,
		Data: _data,
	}
}
func SerializeInspectUserReply(_code int, _data InspectUserInnerReply) InspectUserReply {

	return InspectUserReply{
		Code: _code,
		Data: _data,
	}
}
func _packSerializeInspectUserReply(_code int, _data InspectUserInnerReply) InspectUserReply {

	return InspectUserReply{
		Code: _code,
		Data: _data,
	}
}
func PackSerializeInspectUserReply(_code []int, _data []InspectUserInnerReply) (pack []InspectUserReply) {
	for i := range _code {
		pack = append(pack, _packSerializeInspectUserReply(_code[i], _data[i]))
	}
	return
}
func PSerializeInspectUserInnerReply(user *user.User, _identities []string, _successProblems []uint, _triedProblems []uint) *InspectUserInnerReply {

	return &InspectUserInnerReply{
		Id:              user.ID,
		NickName:        user.NickName,
		UserName:        user.UserName,
		LastLogin:       user.LastLogin,
		Email:           user.Email,
		Motto:           user.Motto,
		Gender:          user.Gender,
		Identities:      _identities,
		SuccessProblems: _successProblems,
		TriedProblems:   _triedProblems,
	}
}
func SerializeInspectUserInnerReply(user *user.User, _identities []string, _successProblems []uint, _triedProblems []uint) InspectUserInnerReply {

	return InspectUserInnerReply{
		Id:              user.ID,
		NickName:        user.NickName,
		UserName:        user.UserName,
		LastLogin:       user.LastLogin,
		Email:           user.Email,
		Motto:           user.Motto,
		Gender:          user.Gender,
		Identities:      _identities,
		SuccessProblems: _successProblems,
		TriedProblems:   _triedProblems,
	}
}
func _packSerializeInspectUserInnerReply(user *user.User, _identities []string, _successProblems []uint, _triedProblems []uint) InspectUserInnerReply {

	return InspectUserInnerReply{
		Id:              user.ID,
		NickName:        user.NickName,
		UserName:        user.UserName,
		LastLogin:       user.LastLogin,
		Email:           user.Email,
		Motto:           user.Motto,
		Gender:          user.Gender,
		Identities:      _identities,
		SuccessProblems: _successProblems,
		TriedProblems:   _triedProblems,
	}
}
func PackSerializeInspectUserInnerReply(user []*user.User, _identities [][]string, _successProblems [][]uint, _triedProblems [][]uint) (pack []InspectUserInnerReply) {
	for i := range user {
		pack = append(pack, _packSerializeInspectUserInnerReply(user[i], _identities[i], _successProblems[i], _triedProblems[i]))
	}
	return
}
func PSerializeGetUserRequest() *GetUserRequest {

	return &GetUserRequest{}
}
func SerializeGetUserRequest() GetUserRequest {

	return GetUserRequest{}
}
func _packSerializeGetUserRequest() GetUserRequest {

	return GetUserRequest{}
}
func PackSerializeGetUserRequest() (pack []GetUserRequest) {
	return
}
func PSerializeGetUserReply(_code int, _data GetUserInnerReply) *GetUserReply {

	return &GetUserReply{
		Code: _code,
		Data: _data,
	}
}
func SerializeGetUserReply(_code int, _data GetUserInnerReply) GetUserReply {

	return GetUserReply{
		Code: _code,
		Data: _data,
	}
}
func _packSerializeGetUserReply(_code int, _data GetUserInnerReply) GetUserReply {

	return GetUserReply{
		Code: _code,
		Data: _data,
	}
}
func PackSerializeGetUserReply(_code []int, _data []GetUserInnerReply) (pack []GetUserReply) {
	for i := range _code {
		pack = append(pack, _packSerializeGetUserReply(_code[i], _data[i]))
	}
	return
}
func PSerializeGetUserInnerReply(user *user.User) *GetUserInnerReply {

	return &GetUserInnerReply{
		Id:        user.ID,
		NickName:  user.NickName,
		LastLogin: user.LastLogin,
		Motto:     user.Motto,
		Gender:    user.Gender,
	}
}
func SerializeGetUserInnerReply(user *user.User) GetUserInnerReply {

	return GetUserInnerReply{
		Id:        user.ID,
		NickName:  user.NickName,
		LastLogin: user.LastLogin,
		Motto:     user.Motto,
		Gender:    user.Gender,
	}
}
func _packSerializeGetUserInnerReply(user *user.User) GetUserInnerReply {

	return GetUserInnerReply{
		Id:        user.ID,
		NickName:  user.NickName,
		LastLogin: user.LastLogin,
		Motto:     user.Motto,
		Gender:    user.Gender,
	}
}
func PackSerializeGetUserInnerReply(user []*user.User) (pack []GetUserInnerReply) {
	for i := range user {
		pack = append(pack, _packSerializeGetUserInnerReply(user[i]))
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
func PSerializePutUserReply(_code int) *PutUserReply {

	return &PutUserReply{
		Code: _code,
	}
}
func SerializePutUserReply(_code int) PutUserReply {

	return PutUserReply{
		Code: _code,
	}
}
func _packSerializePutUserReply(_code int) PutUserReply {

	return PutUserReply{
		Code: _code,
	}
}
func PackSerializePutUserReply(_code []int) (pack []PutUserReply) {
	for i := range _code {
		pack = append(pack, _packSerializePutUserReply(_code[i]))
	}
	return
}
func PSerializeDeleteUserRequest() *DeleteUserRequest {

	return &DeleteUserRequest{}
}
func SerializeDeleteUserRequest() DeleteUserRequest {

	return DeleteUserRequest{}
}
func _packSerializeDeleteUserRequest() DeleteUserRequest {

	return DeleteUserRequest{}
}
func PackSerializeDeleteUserRequest() (pack []DeleteUserRequest) {
	return
}
func PSerializeDeleteUserReply(_code int) *DeleteUserReply {

	return &DeleteUserReply{
		Code: _code,
	}
}
func SerializeDeleteUserReply(_code int) DeleteUserReply {

	return DeleteUserReply{
		Code: _code,
	}
}
func _packSerializeDeleteUserReply(_code int) DeleteUserReply {

	return DeleteUserReply{
		Code: _code,
	}
}
func PackSerializeDeleteUserReply(_code []int) (pack []DeleteUserReply) {
	for i := range _code {
		pack = append(pack, _packSerializeDeleteUserReply(_code[i]))
	}
	return
}
