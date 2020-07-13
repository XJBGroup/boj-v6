package api

import (
	"github.com/Myriad-Dreamin/boj-v6/abstract/problem"
	"github.com/Myriad-Dreamin/boj-v6/types/problem-config"
	"github.com/Myriad-Dreamin/go-model-traits/gorm-crud-dao"
)

type ListProblemsRequest = gorm_crud_dao.Filter

type ListProblemsReply struct {
	Code int               `json:"code" form:"code"`
	Data []problem.Problem `json:"data" form:"data"`
}

type CountProblemsRequest = gorm_crud_dao.Filter

type CountProblemReply struct {
	Code int   `json:"code" form:"code"`
	Data []int `json:"data" form:"data"`
}

type PostProblemRequest struct {
	Title       string                       `binding:"required" json:"title" form:"title"`
	Description string                       `json:"description" form:"description"`
	Config      *problemconfig.ProblemConfig `json:"config" form:"config"`
}

type PostProblemReply struct {
	Code int  `json:"code" form:"code"`
	Id   uint `json:"id" form:"id"`
}

type ChangeTemplateNameRequest struct {
	Name    string `json:"name" form:"name" binding:"required"`
	NewName string `json:"new_name" form:"new_name" binding:"required"`
}

type PostTemplateRequest struct {
	Name    string `json:"name" form:"name" binding:"required"`
	Content string `form:"content" json:"content"`
}

type GetTemplateRequest struct {
	Name string `json:"name" form:"name"`
}

type GetTemplateReply struct {
	Code int             `form:"code" json:"code"`
	Data ProblemTemplate `json:"data" form:"data"`
}

type ProblemTemplate struct {
	Name    string `json:"name" form:"name"`
	Content string `json:"content" form:"content"`
}

type PutTemplateRequest struct {
	Name    string `json:"name" form:"name" binding:"required"`
	Content string `json:"content" form:"content"`
}

type GetProblemReply struct {
	Code    int              `form:"code" json:"code"`
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
func PSerializePostProblemReply(_code int, problem *problem.Problem) *PostProblemReply {

	return &PostProblemReply{
		Code: _code,
		Id:   problem.ID,
	}
}
func SerializePostProblemReply(_code int, problem *problem.Problem) PostProblemReply {

	return PostProblemReply{
		Code: _code,
		Id:   problem.ID,
	}
}
func _packSerializePostProblemReply(_code int, problem *problem.Problem) PostProblemReply {

	return PostProblemReply{
		Code: _code,
		Id:   problem.ID,
	}
}
func PackSerializePostProblemReply(_code []int, problem []*problem.Problem) (pack []PostProblemReply) {
	for i := range _code {
		pack = append(pack, _packSerializePostProblemReply(_code[i], problem[i]))
	}
	return
}
func PSerializeChangeTemplateNameRequest(_name string, _newName string) *ChangeTemplateNameRequest {

	return &ChangeTemplateNameRequest{
		Name:    _name,
		NewName: _newName,
	}
}
func SerializeChangeTemplateNameRequest(_name string, _newName string) ChangeTemplateNameRequest {

	return ChangeTemplateNameRequest{
		Name:    _name,
		NewName: _newName,
	}
}
func _packSerializeChangeTemplateNameRequest(_name string, _newName string) ChangeTemplateNameRequest {

	return ChangeTemplateNameRequest{
		Name:    _name,
		NewName: _newName,
	}
}
func PackSerializeChangeTemplateNameRequest(_name []string, _newName []string) (pack []ChangeTemplateNameRequest) {
	for i := range _name {
		pack = append(pack, _packSerializeChangeTemplateNameRequest(_name[i], _newName[i]))
	}
	return
}
func PSerializePostTemplateRequest(_name string, _content string) *PostTemplateRequest {

	return &PostTemplateRequest{
		Name:    _name,
		Content: _content,
	}
}
func SerializePostTemplateRequest(_name string, _content string) PostTemplateRequest {

	return PostTemplateRequest{
		Name:    _name,
		Content: _content,
	}
}
func _packSerializePostTemplateRequest(_name string, _content string) PostTemplateRequest {

	return PostTemplateRequest{
		Name:    _name,
		Content: _content,
	}
}
func PackSerializePostTemplateRequest(_name []string, _content []string) (pack []PostTemplateRequest) {
	for i := range _name {
		pack = append(pack, _packSerializePostTemplateRequest(_name[i], _content[i]))
	}
	return
}
func PSerializeGetTemplateRequest(_name string) *GetTemplateRequest {

	return &GetTemplateRequest{
		Name: _name,
	}
}
func SerializeGetTemplateRequest(_name string) GetTemplateRequest {

	return GetTemplateRequest{
		Name: _name,
	}
}
func _packSerializeGetTemplateRequest(_name string) GetTemplateRequest {

	return GetTemplateRequest{
		Name: _name,
	}
}
func PackSerializeGetTemplateRequest(_name []string) (pack []GetTemplateRequest) {
	for i := range _name {
		pack = append(pack, _packSerializeGetTemplateRequest(_name[i]))
	}
	return
}
func PSerializeGetTemplateReply(_code int, _data ProblemTemplate) *GetTemplateReply {

	return &GetTemplateReply{
		Code: _code,
		Data: _data,
	}
}
func SerializeGetTemplateReply(_code int, _data ProblemTemplate) GetTemplateReply {

	return GetTemplateReply{
		Code: _code,
		Data: _data,
	}
}
func _packSerializeGetTemplateReply(_code int, _data ProblemTemplate) GetTemplateReply {

	return GetTemplateReply{
		Code: _code,
		Data: _data,
	}
}
func PackSerializeGetTemplateReply(_code []int, _data []ProblemTemplate) (pack []GetTemplateReply) {
	for i := range _code {
		pack = append(pack, _packSerializeGetTemplateReply(_code[i], _data[i]))
	}
	return
}
func PSerializeProblemTemplate(_name string, _content string) *ProblemTemplate {

	return &ProblemTemplate{
		Name:    _name,
		Content: _content,
	}
}
func SerializeProblemTemplate(_name string, _content string) ProblemTemplate {

	return ProblemTemplate{
		Name:    _name,
		Content: _content,
	}
}
func _packSerializeProblemTemplate(_name string, _content string) ProblemTemplate {

	return ProblemTemplate{
		Name:    _name,
		Content: _content,
	}
}
func PackSerializeProblemTemplate(_name []string, _content []string) (pack []ProblemTemplate) {
	for i := range _name {
		pack = append(pack, _packSerializeProblemTemplate(_name[i], _content[i]))
	}
	return
}
func PSerializePutTemplateRequest(_name string, _content string) *PutTemplateRequest {

	return &PutTemplateRequest{
		Name:    _name,
		Content: _content,
	}
}
func SerializePutTemplateRequest(_name string, _content string) PutTemplateRequest {

	return PutTemplateRequest{
		Name:    _name,
		Content: _content,
	}
}
func _packSerializePutTemplateRequest(_name string, _content string) PutTemplateRequest {

	return PutTemplateRequest{
		Name:    _name,
		Content: _content,
	}
}
func PackSerializePutTemplateRequest(_name []string, _content []string) (pack []PutTemplateRequest) {
	for i := range _name {
		pack = append(pack, _packSerializePutTemplateRequest(_name[i], _content[i]))
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
