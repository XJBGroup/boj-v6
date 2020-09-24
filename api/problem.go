package api

import (
	"github.com/Myriad-Dreamin/boj-v6/abstract/problem"
	"github.com/Myriad-Dreamin/boj-v6/abstract/problem-desc"
	"github.com/Myriad-Dreamin/boj-v6/abstract/user"
	"github.com/Myriad-Dreamin/boj-v6/types/problem-config"
	"github.com/Myriad-Dreamin/go-model-traits/gorm-crud-dao"
	"time"
)

type ListProblemRequest = gorm_crud_dao.Filter

type ListProblemReply struct {
	Code int               `json:"code" form:"code"`
	Data []problem.Problem `json:"data" form:"data"`
}

type CountProblemRequest = gorm_crud_dao.Filter

type CountProblemReply struct {
	Code int   `json:"code" form:"code"`
	Data []int `json:"data" form:"data"`
}

type PostProblemRequest struct {
	Title       string                       `json:"title" form:"title" binding:"required"`
	Description string                       `json:"description" form:"description"`
	Config      *problemconfig.ProblemConfig `json:"config" form:"config"`
}

type PostProblemReply struct {
	Code int             `json:"code" form:"code"`
	Data PostProblemData `json:"data" form:"data"`
}

type PostProblemData struct {
	Id uint `json:"id" form:"id"`
}

type CountProblemDescRequest = gorm_crud_dao.Filter

type CountProblemDescReply struct {
	Code int   `json:"code" form:"code"`
	Data int64 `json:"data" form:"data"`
}

type ChangeProblemDescriptionRefRequest struct {
	Name    string `form:"name" binding:"required" json:"name"`
	NewName string `binding:"required" json:"new_name" form:"new_name"`
}

type ChangeProblemDescriptionRefReply struct {
	Code int `json:"code" form:"code"`
}

type PostProblemDescRequest struct {
	Name    string `json:"name" form:"name" binding:"required"`
	Content string `form:"content" json:"content"`
}

type PostProblemDescReply struct {
	Code int `form:"code" json:"code"`
}

type GetProblemDescRequest struct {
	Name string `json:"name" form:"name"`
}

type GetProblemDescReply struct {
	Code int    `json:"code" form:"code"`
	Data string `json:"data" form:"data"`
}

type PutProblemDescRequest struct {
	Name    string `json:"name" form:"name" binding:"required"`
	Content string `json:"content" form:"content"`
}

type PutProblemDescReply struct {
	Code int `json:"code" form:"code"`
}

type DeleteProblemDescRequest struct {
	Name string `form:"name" json:"name"`
}

type DeleteProblemDescReply struct {
	Code int `json:"code" form:"code"`
}

type ListProblemDescRequest = gorm_crud_dao.Filter

type ListProblemDescReply struct {
	Code int               `json:"code" form:"code"`
	Data []ProblemDescData `json:"data" form:"data"`
}

type ProblemDescData struct {
	Name      string    `json:"name" form:"name"`
	UpdatedAt time.Time `json:"updated_at" form:"updated_at"`
}

type GetProblemRequest struct {
}

type GetProblemReply struct {
	Code int            `form:"code" json:"code"`
	Data GetProblemData `json:"data" form:"data"`
}

type GetProblemData struct {
	Id              uint                 `form:"id" json:"id"`
	CreatedAt       time.Time            `json:"created_at" form:"created_at"`
	UpdatedAt       time.Time            `json:"updated_at" form:"updated_at"`
	IsSpj           bool                 `json:"is_spj" form:"is_spj"`
	Title           string               `form:"title" json:"title"`
	Description     string               `json:"description" form:"description"`
	DescriptionRef  string               `json:"description_ref" form:"description_ref"`
	TimeLimit       int64                `form:"time_limit" json:"time_limit"`
	MemoryLimit     int64                `json:"memory_limit" form:"memory_limit"`
	CodeLengthLimit int64                `json:"code_length_limit" form:"code_length_limit"`
	Author          GetProblemAuthorData `json:"author" form:"author"`
}

type GetProblemAuthorData struct {
	Id       uint   `json:"id" form:"id"`
	NickName string `form:"nick_name" json:"nick_name"`
}

type PutProblemRequest struct {
	Title          string `json:"title" form:"title"`
	Description    string `json:"description" form:"description"`
	DescriptionRef string `json:"description_ref" form:"description_ref"`
}

type PutProblemReply struct {
	Code int `json:"code" form:"code"`
}

type DeleteProblemRequest struct {
}

type DeleteProblemReply struct {
	Code int `json:"code" form:"code"`
}

func PSerializeListProblemReply(_code int, _data []problem.Problem) *ListProblemReply {

	return &ListProblemReply{
		Code: _code,
		Data: _data,
	}
}
func SerializeListProblemReply(_code int, _data []problem.Problem) ListProblemReply {

	return ListProblemReply{
		Code: _code,
		Data: _data,
	}
}
func _packSerializeListProblemReply(_code int, _data []problem.Problem) ListProblemReply {

	return ListProblemReply{
		Code: _code,
		Data: _data,
	}
}
func PackSerializeListProblemReply(_code []int, _data [][]problem.Problem) (pack []ListProblemReply) {
	for i := range _code {
		pack = append(pack, _packSerializeListProblemReply(_code[i], _data[i]))
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
func PSerializePostProblemRequest(problem *problem.Problem, _config *problemconfig.ProblemConfig) *PostProblemRequest {

	return &PostProblemRequest{
		Title:       problem.Title,
		Description: problem.Description,
		Config:      _config,
	}
}
func SerializePostProblemRequest(problem *problem.Problem, _config *problemconfig.ProblemConfig) PostProblemRequest {

	return PostProblemRequest{
		Title:       problem.Title,
		Description: problem.Description,
		Config:      _config,
	}
}
func _packSerializePostProblemRequest(problem *problem.Problem, _config *problemconfig.ProblemConfig) PostProblemRequest {

	return PostProblemRequest{
		Title:       problem.Title,
		Description: problem.Description,
		Config:      _config,
	}
}
func PackSerializePostProblemRequest(problem []*problem.Problem, _config []*problemconfig.ProblemConfig) (pack []PostProblemRequest) {
	for i := range problem {
		pack = append(pack, _packSerializePostProblemRequest(problem[i], _config[i]))
	}
	return
}
func PSerializePostProblemReply(_code int, _data PostProblemData) *PostProblemReply {

	return &PostProblemReply{
		Code: _code,
		Data: _data,
	}
}
func SerializePostProblemReply(_code int, _data PostProblemData) PostProblemReply {

	return PostProblemReply{
		Code: _code,
		Data: _data,
	}
}
func _packSerializePostProblemReply(_code int, _data PostProblemData) PostProblemReply {

	return PostProblemReply{
		Code: _code,
		Data: _data,
	}
}
func PackSerializePostProblemReply(_code []int, _data []PostProblemData) (pack []PostProblemReply) {
	for i := range _code {
		pack = append(pack, _packSerializePostProblemReply(_code[i], _data[i]))
	}
	return
}
func PSerializePostProblemData(problem *problem.Problem) *PostProblemData {

	return &PostProblemData{
		Id: problem.ID,
	}
}
func SerializePostProblemData(problem *problem.Problem) PostProblemData {

	return PostProblemData{
		Id: problem.ID,
	}
}
func _packSerializePostProblemData(problem *problem.Problem) PostProblemData {

	return PostProblemData{
		Id: problem.ID,
	}
}
func PackSerializePostProblemData(problem []*problem.Problem) (pack []PostProblemData) {
	for i := range problem {
		pack = append(pack, _packSerializePostProblemData(problem[i]))
	}
	return
}
func PSerializeCountProblemDescReply(_code int, _data int64) *CountProblemDescReply {

	return &CountProblemDescReply{
		Code: _code,
		Data: _data,
	}
}
func SerializeCountProblemDescReply(_code int, _data int64) CountProblemDescReply {

	return CountProblemDescReply{
		Code: _code,
		Data: _data,
	}
}
func _packSerializeCountProblemDescReply(_code int, _data int64) CountProblemDescReply {

	return CountProblemDescReply{
		Code: _code,
		Data: _data,
	}
}
func PackSerializeCountProblemDescReply(_code []int, _data []int64) (pack []CountProblemDescReply) {
	for i := range _code {
		pack = append(pack, _packSerializeCountProblemDescReply(_code[i], _data[i]))
	}
	return
}
func PSerializeChangeProblemDescriptionRefRequest(_name string, _newName string) *ChangeProblemDescriptionRefRequest {

	return &ChangeProblemDescriptionRefRequest{
		Name:    _name,
		NewName: _newName,
	}
}
func SerializeChangeProblemDescriptionRefRequest(_name string, _newName string) ChangeProblemDescriptionRefRequest {

	return ChangeProblemDescriptionRefRequest{
		Name:    _name,
		NewName: _newName,
	}
}
func _packSerializeChangeProblemDescriptionRefRequest(_name string, _newName string) ChangeProblemDescriptionRefRequest {

	return ChangeProblemDescriptionRefRequest{
		Name:    _name,
		NewName: _newName,
	}
}
func PackSerializeChangeProblemDescriptionRefRequest(_name []string, _newName []string) (pack []ChangeProblemDescriptionRefRequest) {
	for i := range _name {
		pack = append(pack, _packSerializeChangeProblemDescriptionRefRequest(_name[i], _newName[i]))
	}
	return
}
func PSerializeChangeProblemDescriptionRefReply(_code int) *ChangeProblemDescriptionRefReply {

	return &ChangeProblemDescriptionRefReply{
		Code: _code,
	}
}
func SerializeChangeProblemDescriptionRefReply(_code int) ChangeProblemDescriptionRefReply {

	return ChangeProblemDescriptionRefReply{
		Code: _code,
	}
}
func _packSerializeChangeProblemDescriptionRefReply(_code int) ChangeProblemDescriptionRefReply {

	return ChangeProblemDescriptionRefReply{
		Code: _code,
	}
}
func PackSerializeChangeProblemDescriptionRefReply(_code []int) (pack []ChangeProblemDescriptionRefReply) {
	for i := range _code {
		pack = append(pack, _packSerializeChangeProblemDescriptionRefReply(_code[i]))
	}
	return
}
func PSerializePostProblemDescRequest(_name string, _content string) *PostProblemDescRequest {

	return &PostProblemDescRequest{
		Name:    _name,
		Content: _content,
	}
}
func SerializePostProblemDescRequest(_name string, _content string) PostProblemDescRequest {

	return PostProblemDescRequest{
		Name:    _name,
		Content: _content,
	}
}
func _packSerializePostProblemDescRequest(_name string, _content string) PostProblemDescRequest {

	return PostProblemDescRequest{
		Name:    _name,
		Content: _content,
	}
}
func PackSerializePostProblemDescRequest(_name []string, _content []string) (pack []PostProblemDescRequest) {
	for i := range _name {
		pack = append(pack, _packSerializePostProblemDescRequest(_name[i], _content[i]))
	}
	return
}
func PSerializePostProblemDescReply(_code int) *PostProblemDescReply {

	return &PostProblemDescReply{
		Code: _code,
	}
}
func SerializePostProblemDescReply(_code int) PostProblemDescReply {

	return PostProblemDescReply{
		Code: _code,
	}
}
func _packSerializePostProblemDescReply(_code int) PostProblemDescReply {

	return PostProblemDescReply{
		Code: _code,
	}
}
func PackSerializePostProblemDescReply(_code []int) (pack []PostProblemDescReply) {
	for i := range _code {
		pack = append(pack, _packSerializePostProblemDescReply(_code[i]))
	}
	return
}
func PSerializeGetProblemDescRequest(_name string) *GetProblemDescRequest {

	return &GetProblemDescRequest{
		Name: _name,
	}
}
func SerializeGetProblemDescRequest(_name string) GetProblemDescRequest {

	return GetProblemDescRequest{
		Name: _name,
	}
}
func _packSerializeGetProblemDescRequest(_name string) GetProblemDescRequest {

	return GetProblemDescRequest{
		Name: _name,
	}
}
func PackSerializeGetProblemDescRequest(_name []string) (pack []GetProblemDescRequest) {
	for i := range _name {
		pack = append(pack, _packSerializeGetProblemDescRequest(_name[i]))
	}
	return
}
func PSerializeGetProblemDescReply(_code int, _data string) *GetProblemDescReply {

	return &GetProblemDescReply{
		Code: _code,
		Data: _data,
	}
}
func SerializeGetProblemDescReply(_code int, _data string) GetProblemDescReply {

	return GetProblemDescReply{
		Code: _code,
		Data: _data,
	}
}
func _packSerializeGetProblemDescReply(_code int, _data string) GetProblemDescReply {

	return GetProblemDescReply{
		Code: _code,
		Data: _data,
	}
}
func PackSerializeGetProblemDescReply(_code []int, _data []string) (pack []GetProblemDescReply) {
	for i := range _code {
		pack = append(pack, _packSerializeGetProblemDescReply(_code[i], _data[i]))
	}
	return
}
func PSerializePutProblemDescRequest(_name string, _content string) *PutProblemDescRequest {

	return &PutProblemDescRequest{
		Name:    _name,
		Content: _content,
	}
}
func SerializePutProblemDescRequest(_name string, _content string) PutProblemDescRequest {

	return PutProblemDescRequest{
		Name:    _name,
		Content: _content,
	}
}
func _packSerializePutProblemDescRequest(_name string, _content string) PutProblemDescRequest {

	return PutProblemDescRequest{
		Name:    _name,
		Content: _content,
	}
}
func PackSerializePutProblemDescRequest(_name []string, _content []string) (pack []PutProblemDescRequest) {
	for i := range _name {
		pack = append(pack, _packSerializePutProblemDescRequest(_name[i], _content[i]))
	}
	return
}
func PSerializePutProblemDescReply(_code int) *PutProblemDescReply {

	return &PutProblemDescReply{
		Code: _code,
	}
}
func SerializePutProblemDescReply(_code int) PutProblemDescReply {

	return PutProblemDescReply{
		Code: _code,
	}
}
func _packSerializePutProblemDescReply(_code int) PutProblemDescReply {

	return PutProblemDescReply{
		Code: _code,
	}
}
func PackSerializePutProblemDescReply(_code []int) (pack []PutProblemDescReply) {
	for i := range _code {
		pack = append(pack, _packSerializePutProblemDescReply(_code[i]))
	}
	return
}
func PSerializeDeleteProblemDescRequest(_name string) *DeleteProblemDescRequest {

	return &DeleteProblemDescRequest{
		Name: _name,
	}
}
func SerializeDeleteProblemDescRequest(_name string) DeleteProblemDescRequest {

	return DeleteProblemDescRequest{
		Name: _name,
	}
}
func _packSerializeDeleteProblemDescRequest(_name string) DeleteProblemDescRequest {

	return DeleteProblemDescRequest{
		Name: _name,
	}
}
func PackSerializeDeleteProblemDescRequest(_name []string) (pack []DeleteProblemDescRequest) {
	for i := range _name {
		pack = append(pack, _packSerializeDeleteProblemDescRequest(_name[i]))
	}
	return
}
func PSerializeDeleteProblemDescReply(_code int) *DeleteProblemDescReply {

	return &DeleteProblemDescReply{
		Code: _code,
	}
}
func SerializeDeleteProblemDescReply(_code int) DeleteProblemDescReply {

	return DeleteProblemDescReply{
		Code: _code,
	}
}
func _packSerializeDeleteProblemDescReply(_code int) DeleteProblemDescReply {

	return DeleteProblemDescReply{
		Code: _code,
	}
}
func PackSerializeDeleteProblemDescReply(_code []int) (pack []DeleteProblemDescReply) {
	for i := range _code {
		pack = append(pack, _packSerializeDeleteProblemDescReply(_code[i]))
	}
	return
}
func PSerializeListProblemDescReply(_code int, _data []ProblemDescData) *ListProblemDescReply {

	return &ListProblemDescReply{
		Code: _code,
		Data: _data,
	}
}
func SerializeListProblemDescReply(_code int, _data []ProblemDescData) ListProblemDescReply {

	return ListProblemDescReply{
		Code: _code,
		Data: _data,
	}
}
func _packSerializeListProblemDescReply(_code int, _data []ProblemDescData) ListProblemDescReply {

	return ListProblemDescReply{
		Code: _code,
		Data: _data,
	}
}
func PackSerializeListProblemDescReply(_code []int, _data [][]ProblemDescData) (pack []ListProblemDescReply) {
	for i := range _code {
		pack = append(pack, _packSerializeListProblemDescReply(_code[i], _data[i]))
	}
	return
}
func PSerializeProblemDescData(problemDesc problem_desc.ProblemDesc) *ProblemDescData {

	return &ProblemDescData{
		Name:      problemDesc.Name,
		UpdatedAt: problemDesc.UpdatedAt,
	}
}
func SerializeProblemDescData(problemDesc problem_desc.ProblemDesc) ProblemDescData {

	return ProblemDescData{
		Name:      problemDesc.Name,
		UpdatedAt: problemDesc.UpdatedAt,
	}
}
func _packSerializeProblemDescData(problemDesc problem_desc.ProblemDesc) ProblemDescData {

	return ProblemDescData{
		Name:      problemDesc.Name,
		UpdatedAt: problemDesc.UpdatedAt,
	}
}
func PackSerializeProblemDescData(problemDesc []problem_desc.ProblemDesc) (pack []ProblemDescData) {
	for i := range problemDesc {
		pack = append(pack, _packSerializeProblemDescData(problemDesc[i]))
	}
	return
}
func PSerializeGetProblemRequest() *GetProblemRequest {

	return &GetProblemRequest{}
}
func SerializeGetProblemRequest() GetProblemRequest {

	return GetProblemRequest{}
}
func _packSerializeGetProblemRequest() GetProblemRequest {

	return GetProblemRequest{}
}
func PackSerializeGetProblemRequest() (pack []GetProblemRequest) {
	return
}
func PSerializeGetProblemReply(_code int, _data GetProblemData) *GetProblemReply {

	return &GetProblemReply{
		Code: _code,
		Data: _data,
	}
}
func SerializeGetProblemReply(_code int, _data GetProblemData) GetProblemReply {

	return GetProblemReply{
		Code: _code,
		Data: _data,
	}
}
func _packSerializeGetProblemReply(_code int, _data GetProblemData) GetProblemReply {

	return GetProblemReply{
		Code: _code,
		Data: _data,
	}
}
func PackSerializeGetProblemReply(_code []int, _data []GetProblemData) (pack []GetProblemReply) {
	for i := range _code {
		pack = append(pack, _packSerializeGetProblemReply(_code[i], _data[i]))
	}
	return
}
func PSerializeGetProblemData(problem *problem.Problem, _author GetProblemAuthorData) *GetProblemData {

	return &GetProblemData{
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
func SerializeGetProblemData(problem *problem.Problem, _author GetProblemAuthorData) GetProblemData {

	return GetProblemData{
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
func _packSerializeGetProblemData(problem *problem.Problem, _author GetProblemAuthorData) GetProblemData {

	return GetProblemData{
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
func PackSerializeGetProblemData(problem []*problem.Problem, _author []GetProblemAuthorData) (pack []GetProblemData) {
	for i := range problem {
		pack = append(pack, _packSerializeGetProblemData(problem[i], _author[i]))
	}
	return
}
func PSerializeGetProblemAuthorData(problemUser *user.User) *GetProblemAuthorData {

	return &GetProblemAuthorData{
		Id:       problemUser.ID,
		NickName: problemUser.NickName,
	}
}
func SerializeGetProblemAuthorData(problemUser *user.User) GetProblemAuthorData {

	return GetProblemAuthorData{
		Id:       problemUser.ID,
		NickName: problemUser.NickName,
	}
}
func _packSerializeGetProblemAuthorData(problemUser *user.User) GetProblemAuthorData {

	return GetProblemAuthorData{
		Id:       problemUser.ID,
		NickName: problemUser.NickName,
	}
}
func PackSerializeGetProblemAuthorData(problemUser []*user.User) (pack []GetProblemAuthorData) {
	for i := range problemUser {
		pack = append(pack, _packSerializeGetProblemAuthorData(problemUser[i]))
	}
	return
}
func PSerializePutProblemRequest(problem *problem.Problem) *PutProblemRequest {

	return &PutProblemRequest{
		Title:          problem.Title,
		Description:    problem.Description,
		DescriptionRef: problem.DescriptionRef,
	}
}
func SerializePutProblemRequest(problem *problem.Problem) PutProblemRequest {

	return PutProblemRequest{
		Title:          problem.Title,
		Description:    problem.Description,
		DescriptionRef: problem.DescriptionRef,
	}
}
func _packSerializePutProblemRequest(problem *problem.Problem) PutProblemRequest {

	return PutProblemRequest{
		Title:          problem.Title,
		Description:    problem.Description,
		DescriptionRef: problem.DescriptionRef,
	}
}
func PackSerializePutProblemRequest(problem []*problem.Problem) (pack []PutProblemRequest) {
	for i := range problem {
		pack = append(pack, _packSerializePutProblemRequest(problem[i]))
	}
	return
}
func PSerializePutProblemReply(_code int) *PutProblemReply {

	return &PutProblemReply{
		Code: _code,
	}
}
func SerializePutProblemReply(_code int) PutProblemReply {

	return PutProblemReply{
		Code: _code,
	}
}
func _packSerializePutProblemReply(_code int) PutProblemReply {

	return PutProblemReply{
		Code: _code,
	}
}
func PackSerializePutProblemReply(_code []int) (pack []PutProblemReply) {
	for i := range _code {
		pack = append(pack, _packSerializePutProblemReply(_code[i]))
	}
	return
}
func PSerializeDeleteProblemRequest() *DeleteProblemRequest {

	return &DeleteProblemRequest{}
}
func SerializeDeleteProblemRequest() DeleteProblemRequest {

	return DeleteProblemRequest{}
}
func _packSerializeDeleteProblemRequest() DeleteProblemRequest {

	return DeleteProblemRequest{}
}
func PackSerializeDeleteProblemRequest() (pack []DeleteProblemRequest) {
	return
}
func PSerializeDeleteProblemReply(_code int) *DeleteProblemReply {

	return &DeleteProblemReply{
		Code: _code,
	}
}
func SerializeDeleteProblemReply(_code int) DeleteProblemReply {

	return DeleteProblemReply{
		Code: _code,
	}
}
func _packSerializeDeleteProblemReply(_code int) DeleteProblemReply {

	return DeleteProblemReply{
		Code: _code,
	}
}
func PackSerializeDeleteProblemReply(_code []int) (pack []DeleteProblemReply) {
	for i := range _code {
		pack = append(pack, _packSerializeDeleteProblemReply(_code[i]))
	}
	return
}
