package api

import (
	"github.com/Myriad-Dreamin/boj-v6/abstract/contest"
	"github.com/Myriad-Dreamin/boj-v6/abstract/problem"
	"github.com/Myriad-Dreamin/boj-v6/abstract/problem-desc"
	"github.com/Myriad-Dreamin/boj-v6/abstract/user"
	"github.com/Myriad-Dreamin/boj-v6/types/problem-config"
	"github.com/Myriad-Dreamin/go-model-traits/gorm-crud-dao"
	"time"
)

type ListContestRequest = gorm_crud_dao.Filter

type ListContestReply struct {
	Code int               `json:"code" form:"code"`
	Data []contest.Contest `json:"data" form:"data"`
}

type CountContestRequest = gorm_crud_dao.Filter

type CountContestReply struct {
	Code int   `json:"code" form:"code"`
	Data []int `json:"data" form:"data"`
}

type PostContestRequest struct {
	Title               string        `json:"title" form:"title" binding:"required"`
	Description         string        `json:"description" form:"description" binding:"required"`
	StartAt             *time.Time    `json:"start_at" form:"start_at" binding:"required"`
	EndDuration         time.Duration `json:"end_duration" form:"end_duration" binding:"required"`
	BoardFrozenDuration time.Duration `json:"board_frozen_duration" form:"board_frozen_duration" binding:"required"`
}

type PostContestReply struct {
	Code int              `json:"code" form:"code"`
	Data *contest.Contest `json:"data" form:"data"`
}

type PostContestProblemRequest struct {
	Title       string                       `json:"title" form:"title" binding:"required"`
	Description string                       `form:"description" json:"description"`
	Config      *problemconfig.ProblemConfig `json:"config" form:"config"`
}

type PostContestProblemReply struct {
	Code int                    `json:"code" form:"code"`
	Data PostContestProblemData `form:"data" json:"data"`
}

type PostContestProblemData struct {
	Id uint `json:"id" form:"id"`
}

type CountContestProblemDescRequest = gorm_crud_dao.Filter

type CountContestProblemDescReply struct {
	Code int   `json:"code" form:"code"`
	Data int64 `json:"data" form:"data"`
}

type ChangeContestProblemDescriptionRefRequest struct {
	Name    string `json:"name" form:"name" binding:"required"`
	NewName string `json:"new_name" form:"new_name" binding:"required"`
}

type ChangeContestProblemDescriptionRefReply struct {
	Code int `json:"code" form:"code"`
}

type PostContestProblemDescRequest struct {
	Name    string `json:"name" form:"name" binding:"required"`
	Content string `json:"content" form:"content"`
}

type PostContestProblemDescReply struct {
	Code int `json:"code" form:"code"`
}

type GetContestProblemDescRequest struct {
	Name string `form:"name" json:"name"`
}

type GetContestProblemDescReply struct {
	Code int    `json:"code" form:"code"`
	Data string `json:"data" form:"data"`
}

type PutContestProblemDescRequest struct {
	Name    string `json:"name" form:"name" binding:"required"`
	Content string `json:"content" form:"content"`
}

type PutContestProblemDescReply struct {
	Code int `json:"code" form:"code"`
}

type DeleteContestProblemDescRequest struct {
	Name string `form:"name" json:"name"`
}

type DeleteContestProblemDescReply struct {
	Code int `json:"code" form:"code"`
}

type ListContestProblemDescRequest = gorm_crud_dao.Filter

type ListContestProblemDescReply struct {
	Code int                      `json:"code" form:"code"`
	Data []ContestProblemDescData `json:"data" form:"data"`
}

type ContestProblemDescData struct {
	Name      string    `json:"name" form:"name"`
	UpdatedAt time.Time `json:"updated_at" form:"updated_at"`
}

type GetContestProblemRequest struct {
}

type GetContestProblemReply struct {
	Code int                   `json:"code" form:"code"`
	Data GetContestProblemData `json:"data" form:"data"`
}

type GetContestProblemData struct {
	Id              uint                        `json:"id" form:"id"`
	CreatedAt       time.Time                   `form:"created_at" json:"created_at"`
	UpdatedAt       time.Time                   `json:"updated_at" form:"updated_at"`
	IsSpj           bool                        `json:"is_spj" form:"is_spj"`
	Title           string                      `json:"title" form:"title"`
	Description     string                      `json:"description" form:"description"`
	DescriptionRef  string                      `form:"description_ref" json:"description_ref"`
	TimeLimit       int64                       `json:"time_limit" form:"time_limit"`
	MemoryLimit     int64                       `json:"memory_limit" form:"memory_limit"`
	CodeLengthLimit int64                       `form:"code_length_limit" json:"code_length_limit"`
	Author          GetContestProblemAuthorData `json:"author" form:"author"`
}

type GetContestProblemAuthorData struct {
	Id       uint   `json:"id" form:"id"`
	NickName string `json:"nick_name" form:"nick_name"`
}

type PutContestProblemRequest struct {
	Title          string `json:"title" form:"title"`
	DescriptionRef string `json:"description_ref" form:"description_ref"`
}

type PutContestProblemReply struct {
	Code int `json:"code" form:"code"`
}

type DeleteContestProblemRequest struct {
}

type DeleteContestProblemReply struct {
	Code int `form:"code" json:"code"`
}

type ListContestUsersRequest struct {
}

type ListContestUsersReply struct {
	Code int         `json:"code" form:"code"`
	Data []user.User `json:"data" form:"data"`
}

type ListContestProblemRequest = gorm_crud_dao.Filter

type ListContestProblemReply struct {
	Code int               `json:"code" form:"code"`
	Data []problem.Problem `form:"data" json:"data"`
}

type CountContestProblemRequest = gorm_crud_dao.Filter

type CountContestProblemReply struct {
	Code int   `json:"code" form:"code"`
	Data []int `form:"data" json:"data"`
}

type GetContestRequest struct {
}

type GetContestReply struct {
	Code int                  `json:"code" form:"code"`
	Data GetContestInnerReply `json:"data" form:"data"`
}

type GetContestInnerReply struct {
	Id                  uint          `form:"id" json:"id"`
	Title               string        `json:"title" form:"title"`
	StartAt             *time.Time    `form:"start_at" json:"start_at"`
	CreatedAt           time.Time     `json:"created_at" form:"created_at"`
	BoardFrozenDuration time.Duration `form:"board_frozen_duration" json:"board_frozen_duration"`
	EndDuration         time.Duration `json:"end_duration" form:"end_duration"`
	Description         string        `json:"description" form:"description"`
	AuthorId            uint          `json:"author_id" form:"author_id"`
	ContestType         string        `json:"contest_type" form:"contest_type"`
}

type PutContestRequest struct {
	Title               string        `json:"title" form:"title"`
	Description         string        `json:"description" form:"description"`
	StartAt             *time.Time    `json:"start_at" form:"start_at"`
	EndDuration         time.Duration `form:"end_duration" json:"end_duration"`
	BoardFrozenDuration time.Duration `json:"board_frozen_duration" form:"board_frozen_duration"`
	ConfigPath          string        `json:"config_path" form:"config_path"`
	RolePath            string        `json:"role_path" form:"role_path"`
}

type PutContestReply struct {
	Code int `json:"code" form:"code"`
}

type DeleteContestRequest struct {
}

type DeleteContestReply struct {
	Code int `json:"code" form:"code"`
}

func PSerializeListContestReply(_code int, _data []contest.Contest) *ListContestReply {

	return &ListContestReply{
		Code: _code,
		Data: _data,
	}
}
func SerializeListContestReply(_code int, _data []contest.Contest) ListContestReply {

	return ListContestReply{
		Code: _code,
		Data: _data,
	}
}
func _packSerializeListContestReply(_code int, _data []contest.Contest) ListContestReply {

	return ListContestReply{
		Code: _code,
		Data: _data,
	}
}
func PackSerializeListContestReply(_code []int, _data [][]contest.Contest) (pack []ListContestReply) {
	for i := range _code {
		pack = append(pack, _packSerializeListContestReply(_code[i], _data[i]))
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
func PSerializePostContestRequest(contest *contest.Contest) *PostContestRequest {

	return &PostContestRequest{
		Title:               contest.Title,
		Description:         contest.Description,
		StartAt:             contest.StartAt,
		EndDuration:         contest.EndDuration,
		BoardFrozenDuration: contest.BoardFrozenDuration,
	}
}
func SerializePostContestRequest(contest *contest.Contest) PostContestRequest {

	return PostContestRequest{
		Title:               contest.Title,
		Description:         contest.Description,
		StartAt:             contest.StartAt,
		EndDuration:         contest.EndDuration,
		BoardFrozenDuration: contest.BoardFrozenDuration,
	}
}
func _packSerializePostContestRequest(contest *contest.Contest) PostContestRequest {

	return PostContestRequest{
		Title:               contest.Title,
		Description:         contest.Description,
		StartAt:             contest.StartAt,
		EndDuration:         contest.EndDuration,
		BoardFrozenDuration: contest.BoardFrozenDuration,
	}
}
func PackSerializePostContestRequest(contest []*contest.Contest) (pack []PostContestRequest) {
	for i := range contest {
		pack = append(pack, _packSerializePostContestRequest(contest[i]))
	}
	return
}
func PSerializePostContestReply(_code int, _data *contest.Contest) *PostContestReply {

	return &PostContestReply{
		Code: _code,
		Data: _data,
	}
}
func SerializePostContestReply(_code int, _data *contest.Contest) PostContestReply {

	return PostContestReply{
		Code: _code,
		Data: _data,
	}
}
func _packSerializePostContestReply(_code int, _data *contest.Contest) PostContestReply {

	return PostContestReply{
		Code: _code,
		Data: _data,
	}
}
func PackSerializePostContestReply(_code []int, _data []*contest.Contest) (pack []PostContestReply) {
	for i := range _code {
		pack = append(pack, _packSerializePostContestReply(_code[i], _data[i]))
	}
	return
}
func PSerializePostContestProblemRequest(problem *problem.Problem, _config *problemconfig.ProblemConfig) *PostContestProblemRequest {

	return &PostContestProblemRequest{
		Title:       problem.Title,
		Description: problem.Description,
		Config:      _config,
	}
}
func SerializePostContestProblemRequest(problem *problem.Problem, _config *problemconfig.ProblemConfig) PostContestProblemRequest {

	return PostContestProblemRequest{
		Title:       problem.Title,
		Description: problem.Description,
		Config:      _config,
	}
}
func _packSerializePostContestProblemRequest(problem *problem.Problem, _config *problemconfig.ProblemConfig) PostContestProblemRequest {

	return PostContestProblemRequest{
		Title:       problem.Title,
		Description: problem.Description,
		Config:      _config,
	}
}
func PackSerializePostContestProblemRequest(problem []*problem.Problem, _config []*problemconfig.ProblemConfig) (pack []PostContestProblemRequest) {
	for i := range problem {
		pack = append(pack, _packSerializePostContestProblemRequest(problem[i], _config[i]))
	}
	return
}
func PSerializePostContestProblemReply(_code int, _data PostContestProblemData) *PostContestProblemReply {

	return &PostContestProblemReply{
		Code: _code,
		Data: _data,
	}
}
func SerializePostContestProblemReply(_code int, _data PostContestProblemData) PostContestProblemReply {

	return PostContestProblemReply{
		Code: _code,
		Data: _data,
	}
}
func _packSerializePostContestProblemReply(_code int, _data PostContestProblemData) PostContestProblemReply {

	return PostContestProblemReply{
		Code: _code,
		Data: _data,
	}
}
func PackSerializePostContestProblemReply(_code []int, _data []PostContestProblemData) (pack []PostContestProblemReply) {
	for i := range _code {
		pack = append(pack, _packSerializePostContestProblemReply(_code[i], _data[i]))
	}
	return
}
func PSerializePostContestProblemData(problem *problem.Problem) *PostContestProblemData {

	return &PostContestProblemData{
		Id: problem.ID,
	}
}
func SerializePostContestProblemData(problem *problem.Problem) PostContestProblemData {

	return PostContestProblemData{
		Id: problem.ID,
	}
}
func _packSerializePostContestProblemData(problem *problem.Problem) PostContestProblemData {

	return PostContestProblemData{
		Id: problem.ID,
	}
}
func PackSerializePostContestProblemData(problem []*problem.Problem) (pack []PostContestProblemData) {
	for i := range problem {
		pack = append(pack, _packSerializePostContestProblemData(problem[i]))
	}
	return
}
func PSerializeCountContestProblemDescReply(_code int, _data int64) *CountContestProblemDescReply {

	return &CountContestProblemDescReply{
		Code: _code,
		Data: _data,
	}
}
func SerializeCountContestProblemDescReply(_code int, _data int64) CountContestProblemDescReply {

	return CountContestProblemDescReply{
		Code: _code,
		Data: _data,
	}
}
func _packSerializeCountContestProblemDescReply(_code int, _data int64) CountContestProblemDescReply {

	return CountContestProblemDescReply{
		Code: _code,
		Data: _data,
	}
}
func PackSerializeCountContestProblemDescReply(_code []int, _data []int64) (pack []CountContestProblemDescReply) {
	for i := range _code {
		pack = append(pack, _packSerializeCountContestProblemDescReply(_code[i], _data[i]))
	}
	return
}
func PSerializeChangeContestProblemDescriptionRefRequest(_name string, _newName string) *ChangeContestProblemDescriptionRefRequest {

	return &ChangeContestProblemDescriptionRefRequest{
		Name:    _name,
		NewName: _newName,
	}
}
func SerializeChangeContestProblemDescriptionRefRequest(_name string, _newName string) ChangeContestProblemDescriptionRefRequest {

	return ChangeContestProblemDescriptionRefRequest{
		Name:    _name,
		NewName: _newName,
	}
}
func _packSerializeChangeContestProblemDescriptionRefRequest(_name string, _newName string) ChangeContestProblemDescriptionRefRequest {

	return ChangeContestProblemDescriptionRefRequest{
		Name:    _name,
		NewName: _newName,
	}
}
func PackSerializeChangeContestProblemDescriptionRefRequest(_name []string, _newName []string) (pack []ChangeContestProblemDescriptionRefRequest) {
	for i := range _name {
		pack = append(pack, _packSerializeChangeContestProblemDescriptionRefRequest(_name[i], _newName[i]))
	}
	return
}
func PSerializeChangeContestProblemDescriptionRefReply(_code int) *ChangeContestProblemDescriptionRefReply {

	return &ChangeContestProblemDescriptionRefReply{
		Code: _code,
	}
}
func SerializeChangeContestProblemDescriptionRefReply(_code int) ChangeContestProblemDescriptionRefReply {

	return ChangeContestProblemDescriptionRefReply{
		Code: _code,
	}
}
func _packSerializeChangeContestProblemDescriptionRefReply(_code int) ChangeContestProblemDescriptionRefReply {

	return ChangeContestProblemDescriptionRefReply{
		Code: _code,
	}
}
func PackSerializeChangeContestProblemDescriptionRefReply(_code []int) (pack []ChangeContestProblemDescriptionRefReply) {
	for i := range _code {
		pack = append(pack, _packSerializeChangeContestProblemDescriptionRefReply(_code[i]))
	}
	return
}
func PSerializePostContestProblemDescRequest(_name string, _content string) *PostContestProblemDescRequest {

	return &PostContestProblemDescRequest{
		Name:    _name,
		Content: _content,
	}
}
func SerializePostContestProblemDescRequest(_name string, _content string) PostContestProblemDescRequest {

	return PostContestProblemDescRequest{
		Name:    _name,
		Content: _content,
	}
}
func _packSerializePostContestProblemDescRequest(_name string, _content string) PostContestProblemDescRequest {

	return PostContestProblemDescRequest{
		Name:    _name,
		Content: _content,
	}
}
func PackSerializePostContestProblemDescRequest(_name []string, _content []string) (pack []PostContestProblemDescRequest) {
	for i := range _name {
		pack = append(pack, _packSerializePostContestProblemDescRequest(_name[i], _content[i]))
	}
	return
}
func PSerializePostContestProblemDescReply(_code int) *PostContestProblemDescReply {

	return &PostContestProblemDescReply{
		Code: _code,
	}
}
func SerializePostContestProblemDescReply(_code int) PostContestProblemDescReply {

	return PostContestProblemDescReply{
		Code: _code,
	}
}
func _packSerializePostContestProblemDescReply(_code int) PostContestProblemDescReply {

	return PostContestProblemDescReply{
		Code: _code,
	}
}
func PackSerializePostContestProblemDescReply(_code []int) (pack []PostContestProblemDescReply) {
	for i := range _code {
		pack = append(pack, _packSerializePostContestProblemDescReply(_code[i]))
	}
	return
}
func PSerializeGetContestProblemDescRequest(_name string) *GetContestProblemDescRequest {

	return &GetContestProblemDescRequest{
		Name: _name,
	}
}
func SerializeGetContestProblemDescRequest(_name string) GetContestProblemDescRequest {

	return GetContestProblemDescRequest{
		Name: _name,
	}
}
func _packSerializeGetContestProblemDescRequest(_name string) GetContestProblemDescRequest {

	return GetContestProblemDescRequest{
		Name: _name,
	}
}
func PackSerializeGetContestProblemDescRequest(_name []string) (pack []GetContestProblemDescRequest) {
	for i := range _name {
		pack = append(pack, _packSerializeGetContestProblemDescRequest(_name[i]))
	}
	return
}
func PSerializeGetContestProblemDescReply(_code int, _data string) *GetContestProblemDescReply {

	return &GetContestProblemDescReply{
		Code: _code,
		Data: _data,
	}
}
func SerializeGetContestProblemDescReply(_code int, _data string) GetContestProblemDescReply {

	return GetContestProblemDescReply{
		Code: _code,
		Data: _data,
	}
}
func _packSerializeGetContestProblemDescReply(_code int, _data string) GetContestProblemDescReply {

	return GetContestProblemDescReply{
		Code: _code,
		Data: _data,
	}
}
func PackSerializeGetContestProblemDescReply(_code []int, _data []string) (pack []GetContestProblemDescReply) {
	for i := range _code {
		pack = append(pack, _packSerializeGetContestProblemDescReply(_code[i], _data[i]))
	}
	return
}
func PSerializePutContestProblemDescRequest(_name string, _content string) *PutContestProblemDescRequest {

	return &PutContestProblemDescRequest{
		Name:    _name,
		Content: _content,
	}
}
func SerializePutContestProblemDescRequest(_name string, _content string) PutContestProblemDescRequest {

	return PutContestProblemDescRequest{
		Name:    _name,
		Content: _content,
	}
}
func _packSerializePutContestProblemDescRequest(_name string, _content string) PutContestProblemDescRequest {

	return PutContestProblemDescRequest{
		Name:    _name,
		Content: _content,
	}
}
func PackSerializePutContestProblemDescRequest(_name []string, _content []string) (pack []PutContestProblemDescRequest) {
	for i := range _name {
		pack = append(pack, _packSerializePutContestProblemDescRequest(_name[i], _content[i]))
	}
	return
}
func PSerializePutContestProblemDescReply(_code int) *PutContestProblemDescReply {

	return &PutContestProblemDescReply{
		Code: _code,
	}
}
func SerializePutContestProblemDescReply(_code int) PutContestProblemDescReply {

	return PutContestProblemDescReply{
		Code: _code,
	}
}
func _packSerializePutContestProblemDescReply(_code int) PutContestProblemDescReply {

	return PutContestProblemDescReply{
		Code: _code,
	}
}
func PackSerializePutContestProblemDescReply(_code []int) (pack []PutContestProblemDescReply) {
	for i := range _code {
		pack = append(pack, _packSerializePutContestProblemDescReply(_code[i]))
	}
	return
}
func PSerializeDeleteContestProblemDescRequest(_name string) *DeleteContestProblemDescRequest {

	return &DeleteContestProblemDescRequest{
		Name: _name,
	}
}
func SerializeDeleteContestProblemDescRequest(_name string) DeleteContestProblemDescRequest {

	return DeleteContestProblemDescRequest{
		Name: _name,
	}
}
func _packSerializeDeleteContestProblemDescRequest(_name string) DeleteContestProblemDescRequest {

	return DeleteContestProblemDescRequest{
		Name: _name,
	}
}
func PackSerializeDeleteContestProblemDescRequest(_name []string) (pack []DeleteContestProblemDescRequest) {
	for i := range _name {
		pack = append(pack, _packSerializeDeleteContestProblemDescRequest(_name[i]))
	}
	return
}
func PSerializeDeleteContestProblemDescReply(_code int) *DeleteContestProblemDescReply {

	return &DeleteContestProblemDescReply{
		Code: _code,
	}
}
func SerializeDeleteContestProblemDescReply(_code int) DeleteContestProblemDescReply {

	return DeleteContestProblemDescReply{
		Code: _code,
	}
}
func _packSerializeDeleteContestProblemDescReply(_code int) DeleteContestProblemDescReply {

	return DeleteContestProblemDescReply{
		Code: _code,
	}
}
func PackSerializeDeleteContestProblemDescReply(_code []int) (pack []DeleteContestProblemDescReply) {
	for i := range _code {
		pack = append(pack, _packSerializeDeleteContestProblemDescReply(_code[i]))
	}
	return
}
func PSerializeListContestProblemDescReply(_code int, _data []ContestProblemDescData) *ListContestProblemDescReply {

	return &ListContestProblemDescReply{
		Code: _code,
		Data: _data,
	}
}
func SerializeListContestProblemDescReply(_code int, _data []ContestProblemDescData) ListContestProblemDescReply {

	return ListContestProblemDescReply{
		Code: _code,
		Data: _data,
	}
}
func _packSerializeListContestProblemDescReply(_code int, _data []ContestProblemDescData) ListContestProblemDescReply {

	return ListContestProblemDescReply{
		Code: _code,
		Data: _data,
	}
}
func PackSerializeListContestProblemDescReply(_code []int, _data [][]ContestProblemDescData) (pack []ListContestProblemDescReply) {
	for i := range _code {
		pack = append(pack, _packSerializeListContestProblemDescReply(_code[i], _data[i]))
	}
	return
}
func PSerializeContestProblemDescData(problemDesc problem_desc.ProblemDesc) *ContestProblemDescData {

	return &ContestProblemDescData{
		Name:      problemDesc.Name,
		UpdatedAt: problemDesc.UpdatedAt,
	}
}
func SerializeContestProblemDescData(problemDesc problem_desc.ProblemDesc) ContestProblemDescData {

	return ContestProblemDescData{
		Name:      problemDesc.Name,
		UpdatedAt: problemDesc.UpdatedAt,
	}
}
func _packSerializeContestProblemDescData(problemDesc problem_desc.ProblemDesc) ContestProblemDescData {

	return ContestProblemDescData{
		Name:      problemDesc.Name,
		UpdatedAt: problemDesc.UpdatedAt,
	}
}
func PackSerializeContestProblemDescData(problemDesc []problem_desc.ProblemDesc) (pack []ContestProblemDescData) {
	for i := range problemDesc {
		pack = append(pack, _packSerializeContestProblemDescData(problemDesc[i]))
	}
	return
}
func PSerializeGetContestProblemRequest() *GetContestProblemRequest {

	return &GetContestProblemRequest{}
}
func SerializeGetContestProblemRequest() GetContestProblemRequest {

	return GetContestProblemRequest{}
}
func _packSerializeGetContestProblemRequest() GetContestProblemRequest {

	return GetContestProblemRequest{}
}
func PackSerializeGetContestProblemRequest() (pack []GetContestProblemRequest) {
	return
}
func PSerializeGetContestProblemReply(_code int, _data GetContestProblemData) *GetContestProblemReply {

	return &GetContestProblemReply{
		Code: _code,
		Data: _data,
	}
}
func SerializeGetContestProblemReply(_code int, _data GetContestProblemData) GetContestProblemReply {

	return GetContestProblemReply{
		Code: _code,
		Data: _data,
	}
}
func _packSerializeGetContestProblemReply(_code int, _data GetContestProblemData) GetContestProblemReply {

	return GetContestProblemReply{
		Code: _code,
		Data: _data,
	}
}
func PackSerializeGetContestProblemReply(_code []int, _data []GetContestProblemData) (pack []GetContestProblemReply) {
	for i := range _code {
		pack = append(pack, _packSerializeGetContestProblemReply(_code[i], _data[i]))
	}
	return
}
func PSerializeGetContestProblemData(problem *problem.Problem, _author GetContestProblemAuthorData) *GetContestProblemData {

	return &GetContestProblemData{
		Id:              problem.ID,
		CreatedAt:       problem.CreatedAt,
		UpdatedAt:       problem.UpdatedAt,
		IsSpj:           problem.IsSpj,
		Title:           problem.Title,
		Description:     problem.Description,
		DescriptionRef:  problem.DescriptionRef,
		TimeLimit:       problem.TimeLimit,
		MemoryLimit:     problem.MemoryLimit,
		CodeLengthLimit: problem.CodeLengthLimit,
		Author:          _author,
	}
}
func SerializeGetContestProblemData(problem *problem.Problem, _author GetContestProblemAuthorData) GetContestProblemData {

	return GetContestProblemData{
		Id:              problem.ID,
		CreatedAt:       problem.CreatedAt,
		UpdatedAt:       problem.UpdatedAt,
		IsSpj:           problem.IsSpj,
		Title:           problem.Title,
		Description:     problem.Description,
		DescriptionRef:  problem.DescriptionRef,
		TimeLimit:       problem.TimeLimit,
		MemoryLimit:     problem.MemoryLimit,
		CodeLengthLimit: problem.CodeLengthLimit,
		Author:          _author,
	}
}
func _packSerializeGetContestProblemData(problem *problem.Problem, _author GetContestProblemAuthorData) GetContestProblemData {

	return GetContestProblemData{
		Id:              problem.ID,
		CreatedAt:       problem.CreatedAt,
		UpdatedAt:       problem.UpdatedAt,
		IsSpj:           problem.IsSpj,
		Title:           problem.Title,
		Description:     problem.Description,
		DescriptionRef:  problem.DescriptionRef,
		TimeLimit:       problem.TimeLimit,
		MemoryLimit:     problem.MemoryLimit,
		CodeLengthLimit: problem.CodeLengthLimit,
		Author:          _author,
	}
}
func PackSerializeGetContestProblemData(problem []*problem.Problem, _author []GetContestProblemAuthorData) (pack []GetContestProblemData) {
	for i := range problem {
		pack = append(pack, _packSerializeGetContestProblemData(problem[i], _author[i]))
	}
	return
}
func PSerializeGetContestProblemAuthorData(problemUser *user.User) *GetContestProblemAuthorData {

	return &GetContestProblemAuthorData{
		Id:       problemUser.ID,
		NickName: problemUser.NickName,
	}
}
func SerializeGetContestProblemAuthorData(problemUser *user.User) GetContestProblemAuthorData {

	return GetContestProblemAuthorData{
		Id:       problemUser.ID,
		NickName: problemUser.NickName,
	}
}
func _packSerializeGetContestProblemAuthorData(problemUser *user.User) GetContestProblemAuthorData {

	return GetContestProblemAuthorData{
		Id:       problemUser.ID,
		NickName: problemUser.NickName,
	}
}
func PackSerializeGetContestProblemAuthorData(problemUser []*user.User) (pack []GetContestProblemAuthorData) {
	for i := range problemUser {
		pack = append(pack, _packSerializeGetContestProblemAuthorData(problemUser[i]))
	}
	return
}
func PSerializePutContestProblemRequest(problem *problem.Problem) *PutContestProblemRequest {

	return &PutContestProblemRequest{
		Title:          problem.Title,
		DescriptionRef: problem.DescriptionRef,
	}
}
func SerializePutContestProblemRequest(problem *problem.Problem) PutContestProblemRequest {

	return PutContestProblemRequest{
		Title:          problem.Title,
		DescriptionRef: problem.DescriptionRef,
	}
}
func _packSerializePutContestProblemRequest(problem *problem.Problem) PutContestProblemRequest {

	return PutContestProblemRequest{
		Title:          problem.Title,
		DescriptionRef: problem.DescriptionRef,
	}
}
func PackSerializePutContestProblemRequest(problem []*problem.Problem) (pack []PutContestProblemRequest) {
	for i := range problem {
		pack = append(pack, _packSerializePutContestProblemRequest(problem[i]))
	}
	return
}
func PSerializePutContestProblemReply(_code int) *PutContestProblemReply {

	return &PutContestProblemReply{
		Code: _code,
	}
}
func SerializePutContestProblemReply(_code int) PutContestProblemReply {

	return PutContestProblemReply{
		Code: _code,
	}
}
func _packSerializePutContestProblemReply(_code int) PutContestProblemReply {

	return PutContestProblemReply{
		Code: _code,
	}
}
func PackSerializePutContestProblemReply(_code []int) (pack []PutContestProblemReply) {
	for i := range _code {
		pack = append(pack, _packSerializePutContestProblemReply(_code[i]))
	}
	return
}
func PSerializeDeleteContestProblemRequest() *DeleteContestProblemRequest {

	return &DeleteContestProblemRequest{}
}
func SerializeDeleteContestProblemRequest() DeleteContestProblemRequest {

	return DeleteContestProblemRequest{}
}
func _packSerializeDeleteContestProblemRequest() DeleteContestProblemRequest {

	return DeleteContestProblemRequest{}
}
func PackSerializeDeleteContestProblemRequest() (pack []DeleteContestProblemRequest) {
	return
}
func PSerializeDeleteContestProblemReply(_code int) *DeleteContestProblemReply {

	return &DeleteContestProblemReply{
		Code: _code,
	}
}
func SerializeDeleteContestProblemReply(_code int) DeleteContestProblemReply {

	return DeleteContestProblemReply{
		Code: _code,
	}
}
func _packSerializeDeleteContestProblemReply(_code int) DeleteContestProblemReply {

	return DeleteContestProblemReply{
		Code: _code,
	}
}
func PackSerializeDeleteContestProblemReply(_code []int) (pack []DeleteContestProblemReply) {
	for i := range _code {
		pack = append(pack, _packSerializeDeleteContestProblemReply(_code[i]))
	}
	return
}
func PSerializeListContestUsersRequest() *ListContestUsersRequest {

	return &ListContestUsersRequest{}
}
func SerializeListContestUsersRequest() ListContestUsersRequest {

	return ListContestUsersRequest{}
}
func _packSerializeListContestUsersRequest() ListContestUsersRequest {

	return ListContestUsersRequest{}
}
func PackSerializeListContestUsersRequest() (pack []ListContestUsersRequest) {
	return
}
func PSerializeListContestUsersReply(_code int, _data []user.User) *ListContestUsersReply {

	return &ListContestUsersReply{
		Code: _code,
		Data: _data,
	}
}
func SerializeListContestUsersReply(_code int, _data []user.User) ListContestUsersReply {

	return ListContestUsersReply{
		Code: _code,
		Data: _data,
	}
}
func _packSerializeListContestUsersReply(_code int, _data []user.User) ListContestUsersReply {

	return ListContestUsersReply{
		Code: _code,
		Data: _data,
	}
}
func PackSerializeListContestUsersReply(_code []int, _data [][]user.User) (pack []ListContestUsersReply) {
	for i := range _code {
		pack = append(pack, _packSerializeListContestUsersReply(_code[i], _data[i]))
	}
	return
}
func PSerializeListContestProblemReply(_code int, _data []problem.Problem) *ListContestProblemReply {

	return &ListContestProblemReply{
		Code: _code,
		Data: _data,
	}
}
func SerializeListContestProblemReply(_code int, _data []problem.Problem) ListContestProblemReply {

	return ListContestProblemReply{
		Code: _code,
		Data: _data,
	}
}
func _packSerializeListContestProblemReply(_code int, _data []problem.Problem) ListContestProblemReply {

	return ListContestProblemReply{
		Code: _code,
		Data: _data,
	}
}
func PackSerializeListContestProblemReply(_code []int, _data [][]problem.Problem) (pack []ListContestProblemReply) {
	for i := range _code {
		pack = append(pack, _packSerializeListContestProblemReply(_code[i], _data[i]))
	}
	return
}
func PSerializeCountContestProblemReply(_code int, _data []int) *CountContestProblemReply {

	return &CountContestProblemReply{
		Code: _code,
		Data: _data,
	}
}
func SerializeCountContestProblemReply(_code int, _data []int) CountContestProblemReply {

	return CountContestProblemReply{
		Code: _code,
		Data: _data,
	}
}
func _packSerializeCountContestProblemReply(_code int, _data []int) CountContestProblemReply {

	return CountContestProblemReply{
		Code: _code,
		Data: _data,
	}
}
func PackSerializeCountContestProblemReply(_code []int, _data [][]int) (pack []CountContestProblemReply) {
	for i := range _code {
		pack = append(pack, _packSerializeCountContestProblemReply(_code[i], _data[i]))
	}
	return
}
func PSerializeGetContestRequest() *GetContestRequest {

	return &GetContestRequest{}
}
func SerializeGetContestRequest() GetContestRequest {

	return GetContestRequest{}
}
func _packSerializeGetContestRequest() GetContestRequest {

	return GetContestRequest{}
}
func PackSerializeGetContestRequest() (pack []GetContestRequest) {
	return
}
func PSerializeGetContestReply(_code int, _data GetContestInnerReply) *GetContestReply {

	return &GetContestReply{
		Code: _code,
		Data: _data,
	}
}
func SerializeGetContestReply(_code int, _data GetContestInnerReply) GetContestReply {

	return GetContestReply{
		Code: _code,
		Data: _data,
	}
}
func _packSerializeGetContestReply(_code int, _data GetContestInnerReply) GetContestReply {

	return GetContestReply{
		Code: _code,
		Data: _data,
	}
}
func PackSerializeGetContestReply(_code []int, _data []GetContestInnerReply) (pack []GetContestReply) {
	for i := range _code {
		pack = append(pack, _packSerializeGetContestReply(_code[i], _data[i]))
	}
	return
}
func PSerializeGetContestInnerReply(contest *contest.Contest) *GetContestInnerReply {

	return &GetContestInnerReply{
		Id:                  contest.ID,
		Title:               contest.Title,
		StartAt:             contest.StartAt,
		CreatedAt:           contest.CreatedAt,
		BoardFrozenDuration: contest.BoardFrozenDuration,
		EndDuration:         contest.EndDuration,
		Description:         contest.Description,
		AuthorId:            contest.AuthorID,
		ContestType:         contest.ContestType,
	}
}
func SerializeGetContestInnerReply(contest *contest.Contest) GetContestInnerReply {

	return GetContestInnerReply{
		Id:                  contest.ID,
		Title:               contest.Title,
		StartAt:             contest.StartAt,
		CreatedAt:           contest.CreatedAt,
		BoardFrozenDuration: contest.BoardFrozenDuration,
		EndDuration:         contest.EndDuration,
		Description:         contest.Description,
		AuthorId:            contest.AuthorID,
		ContestType:         contest.ContestType,
	}
}
func _packSerializeGetContestInnerReply(contest *contest.Contest) GetContestInnerReply {

	return GetContestInnerReply{
		Id:                  contest.ID,
		Title:               contest.Title,
		StartAt:             contest.StartAt,
		CreatedAt:           contest.CreatedAt,
		BoardFrozenDuration: contest.BoardFrozenDuration,
		EndDuration:         contest.EndDuration,
		Description:         contest.Description,
		AuthorId:            contest.AuthorID,
		ContestType:         contest.ContestType,
	}
}
func PackSerializeGetContestInnerReply(contest []*contest.Contest) (pack []GetContestInnerReply) {
	for i := range contest {
		pack = append(pack, _packSerializeGetContestInnerReply(contest[i]))
	}
	return
}
func PSerializePutContestRequest(contest *contest.Contest) *PutContestRequest {

	return &PutContestRequest{
		Title:               contest.Title,
		Description:         contest.Description,
		StartAt:             contest.StartAt,
		EndDuration:         contest.EndDuration,
		BoardFrozenDuration: contest.BoardFrozenDuration,
		ConfigPath:          contest.ConfigPath,
		RolePath:            contest.RolePath,
	}
}
func SerializePutContestRequest(contest *contest.Contest) PutContestRequest {

	return PutContestRequest{
		Title:               contest.Title,
		Description:         contest.Description,
		StartAt:             contest.StartAt,
		EndDuration:         contest.EndDuration,
		BoardFrozenDuration: contest.BoardFrozenDuration,
		ConfigPath:          contest.ConfigPath,
		RolePath:            contest.RolePath,
	}
}
func _packSerializePutContestRequest(contest *contest.Contest) PutContestRequest {

	return PutContestRequest{
		Title:               contest.Title,
		Description:         contest.Description,
		StartAt:             contest.StartAt,
		EndDuration:         contest.EndDuration,
		BoardFrozenDuration: contest.BoardFrozenDuration,
		ConfigPath:          contest.ConfigPath,
		RolePath:            contest.RolePath,
	}
}
func PackSerializePutContestRequest(contest []*contest.Contest) (pack []PutContestRequest) {
	for i := range contest {
		pack = append(pack, _packSerializePutContestRequest(contest[i]))
	}
	return
}
func PSerializePutContestReply(_code int) *PutContestReply {

	return &PutContestReply{
		Code: _code,
	}
}
func SerializePutContestReply(_code int) PutContestReply {

	return PutContestReply{
		Code: _code,
	}
}
func _packSerializePutContestReply(_code int) PutContestReply {

	return PutContestReply{
		Code: _code,
	}
}
func PackSerializePutContestReply(_code []int) (pack []PutContestReply) {
	for i := range _code {
		pack = append(pack, _packSerializePutContestReply(_code[i]))
	}
	return
}
func PSerializeDeleteContestRequest() *DeleteContestRequest {

	return &DeleteContestRequest{}
}
func SerializeDeleteContestRequest() DeleteContestRequest {

	return DeleteContestRequest{}
}
func _packSerializeDeleteContestRequest() DeleteContestRequest {

	return DeleteContestRequest{}
}
func PackSerializeDeleteContestRequest() (pack []DeleteContestRequest) {
	return
}
func PSerializeDeleteContestReply(_code int) *DeleteContestReply {

	return &DeleteContestReply{
		Code: _code,
	}
}
func SerializeDeleteContestReply(_code int) DeleteContestReply {

	return DeleteContestReply{
		Code: _code,
	}
}
func _packSerializeDeleteContestReply(_code int) DeleteContestReply {

	return DeleteContestReply{
		Code: _code,
	}
}
func PackSerializeDeleteContestReply(_code []int) (pack []DeleteContestReply) {
	for i := range _code {
		pack = append(pack, _packSerializeDeleteContestReply(_code[i]))
	}
	return
}
