package api

import (
	"github.com/Myriad-Dreamin/boj-v6/abstract/contest"
	"github.com/Myriad-Dreamin/boj-v6/abstract/problem"
	"github.com/Myriad-Dreamin/boj-v6/types/problem-config"
	"github.com/Myriad-Dreamin/go-model-traits/gorm-crud-dao"
)

type ListContestsRequest = gorm_crud_dao.Filter

type ListContestsReply struct {
	Code int               `json:"code" form:"code"`
	Data []contest.Contest `json:"data" form:"data"`
}

type CountContestsRequest = gorm_crud_dao.Filter

type CountContestReply struct {
	Code int   `json:"code" form:"code"`
	Data []int `json:"data" form:"data"`
}

type PostContestRequest struct {
}

type PostContestReply struct {
	Code    int              `json:"code" form:"code"`
	Contest *contest.Contest `form:"contest" json:"contest"`
}

type ListContestProblemsRequest = gorm_crud_dao.Filter

type ListContestProblemsReply struct {
	Code int               `json:"code" form:"code"`
	Data []problem.Problem `json:"data" form:"data"`
}

type CountContestProblemsRequest = gorm_crud_dao.Filter

type CountContestProblemReply struct {
	Code int   `json:"code" form:"code"`
	Data []int `json:"data" form:"data"`
}

type PostContestProblemRequest struct {
	Title       string                       `json:"title" form:"title" binding:"required"`
	Description string                       `json:"description" form:"description"`
	Config      *problemconfig.ProblemConfig `json:"config" form:"config"`
}

type PostContestProblemReply struct {
	Code int  `json:"code" form:"code"`
	Id   uint `json:"id" form:"id"`
}

type ChangeContestTemplateNameRequest struct {
	Name    string `binding:"required" json:"name" form:"name"`
	NewName string `json:"new_name" form:"new_name" binding:"required"`
}

type PostContestTemplateRequest struct {
	Name    string `binding:"required" json:"name" form:"name"`
	Content string `json:"content" form:"content"`
}

type GetContestTemplateRequest struct {
	Name string `json:"name" form:"name"`
}

type GetContestTemplateReply struct {
	Code int                    `form:"code" json:"code"`
	Data ContestProblemTemplate `json:"data" form:"data"`
}

type ContestProblemTemplate struct {
	Name    string `json:"name" form:"name"`
	Content string `json:"content" form:"content"`
}

type PutContestTemplateRequest struct {
	Name    string `json:"name" form:"name" binding:"required"`
	Content string `json:"content" form:"content"`
}

type GetContestProblemReply struct {
	Code    int              `json:"code" form:"code"`
	Problem *problem.Problem `form:"problem" json:"problem"`
}

type PutContestProblemRequest struct {
}

type GetContestReply struct {
	Code    int              `json:"code" form:"code"`
	Contest *contest.Contest `form:"contest" json:"contest"`
}

type PutContestRequest struct {
}

func PSerializeListContestsReply(_code int, _data []contest.Contest) *ListContestsReply {

	return &ListContestsReply{
		Code: _code,
		Data: _data,
	}
}
func SerializeListContestsReply(_code int, _data []contest.Contest) ListContestsReply {

	return ListContestsReply{
		Code: _code,
		Data: _data,
	}
}
func _packSerializeListContestsReply(_code int, _data []contest.Contest) ListContestsReply {

	return ListContestsReply{
		Code: _code,
		Data: _data,
	}
}
func PackSerializeListContestsReply(_code []int, _data [][]contest.Contest) (pack []ListContestsReply) {
	for i := range _code {
		pack = append(pack, _packSerializeListContestsReply(_code[i], _data[i]))
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
func PSerializePostContestRequest() *PostContestRequest {

	return &PostContestRequest{}
}
func SerializePostContestRequest() PostContestRequest {

	return PostContestRequest{}
}
func _packSerializePostContestRequest() PostContestRequest {

	return PostContestRequest{}
}
func PackSerializePostContestRequest() (pack []PostContestRequest) {
	return
}
func PSerializePostContestReply(_code int, _contest *contest.Contest) *PostContestReply {

	return &PostContestReply{
		Code:    _code,
		Contest: _contest,
	}
}
func SerializePostContestReply(_code int, _contest *contest.Contest) PostContestReply {

	return PostContestReply{
		Code:    _code,
		Contest: _contest,
	}
}
func _packSerializePostContestReply(_code int, _contest *contest.Contest) PostContestReply {

	return PostContestReply{
		Code:    _code,
		Contest: _contest,
	}
}
func PackSerializePostContestReply(_code []int, _contest []*contest.Contest) (pack []PostContestReply) {
	for i := range _code {
		pack = append(pack, _packSerializePostContestReply(_code[i], _contest[i]))
	}
	return
}
func PSerializeListContestProblemsReply(_code int, _data []problem.Problem) *ListContestProblemsReply {

	return &ListContestProblemsReply{
		Code: _code,
		Data: _data,
	}
}
func SerializeListContestProblemsReply(_code int, _data []problem.Problem) ListContestProblemsReply {

	return ListContestProblemsReply{
		Code: _code,
		Data: _data,
	}
}
func _packSerializeListContestProblemsReply(_code int, _data []problem.Problem) ListContestProblemsReply {

	return ListContestProblemsReply{
		Code: _code,
		Data: _data,
	}
}
func PackSerializeListContestProblemsReply(_code []int, _data [][]problem.Problem) (pack []ListContestProblemsReply) {
	for i := range _code {
		pack = append(pack, _packSerializeListContestProblemsReply(_code[i], _data[i]))
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
func PSerializePostContestProblemReply(_code int, problem *problem.Problem) *PostContestProblemReply {

	return &PostContestProblemReply{
		Code: _code,
		Id:   problem.ID,
	}
}
func SerializePostContestProblemReply(_code int, problem *problem.Problem) PostContestProblemReply {

	return PostContestProblemReply{
		Code: _code,
		Id:   problem.ID,
	}
}
func _packSerializePostContestProblemReply(_code int, problem *problem.Problem) PostContestProblemReply {

	return PostContestProblemReply{
		Code: _code,
		Id:   problem.ID,
	}
}
func PackSerializePostContestProblemReply(_code []int, problem []*problem.Problem) (pack []PostContestProblemReply) {
	for i := range _code {
		pack = append(pack, _packSerializePostContestProblemReply(_code[i], problem[i]))
	}
	return
}
func PSerializeChangeContestTemplateNameRequest(_name string, _newName string) *ChangeContestTemplateNameRequest {

	return &ChangeContestTemplateNameRequest{
		Name:    _name,
		NewName: _newName,
	}
}
func SerializeChangeContestTemplateNameRequest(_name string, _newName string) ChangeContestTemplateNameRequest {

	return ChangeContestTemplateNameRequest{
		Name:    _name,
		NewName: _newName,
	}
}
func _packSerializeChangeContestTemplateNameRequest(_name string, _newName string) ChangeContestTemplateNameRequest {

	return ChangeContestTemplateNameRequest{
		Name:    _name,
		NewName: _newName,
	}
}
func PackSerializeChangeContestTemplateNameRequest(_name []string, _newName []string) (pack []ChangeContestTemplateNameRequest) {
	for i := range _name {
		pack = append(pack, _packSerializeChangeContestTemplateNameRequest(_name[i], _newName[i]))
	}
	return
}
func PSerializePostContestTemplateRequest(_name string, _content string) *PostContestTemplateRequest {

	return &PostContestTemplateRequest{
		Name:    _name,
		Content: _content,
	}
}
func SerializePostContestTemplateRequest(_name string, _content string) PostContestTemplateRequest {

	return PostContestTemplateRequest{
		Name:    _name,
		Content: _content,
	}
}
func _packSerializePostContestTemplateRequest(_name string, _content string) PostContestTemplateRequest {

	return PostContestTemplateRequest{
		Name:    _name,
		Content: _content,
	}
}
func PackSerializePostContestTemplateRequest(_name []string, _content []string) (pack []PostContestTemplateRequest) {
	for i := range _name {
		pack = append(pack, _packSerializePostContestTemplateRequest(_name[i], _content[i]))
	}
	return
}
func PSerializeGetContestTemplateRequest(_name string) *GetContestTemplateRequest {

	return &GetContestTemplateRequest{
		Name: _name,
	}
}
func SerializeGetContestTemplateRequest(_name string) GetContestTemplateRequest {

	return GetContestTemplateRequest{
		Name: _name,
	}
}
func _packSerializeGetContestTemplateRequest(_name string) GetContestTemplateRequest {

	return GetContestTemplateRequest{
		Name: _name,
	}
}
func PackSerializeGetContestTemplateRequest(_name []string) (pack []GetContestTemplateRequest) {
	for i := range _name {
		pack = append(pack, _packSerializeGetContestTemplateRequest(_name[i]))
	}
	return
}
func PSerializeGetContestTemplateReply(_code int, _data ContestProblemTemplate) *GetContestTemplateReply {

	return &GetContestTemplateReply{
		Code: _code,
		Data: _data,
	}
}
func SerializeGetContestTemplateReply(_code int, _data ContestProblemTemplate) GetContestTemplateReply {

	return GetContestTemplateReply{
		Code: _code,
		Data: _data,
	}
}
func _packSerializeGetContestTemplateReply(_code int, _data ContestProblemTemplate) GetContestTemplateReply {

	return GetContestTemplateReply{
		Code: _code,
		Data: _data,
	}
}
func PackSerializeGetContestTemplateReply(_code []int, _data []ContestProblemTemplate) (pack []GetContestTemplateReply) {
	for i := range _code {
		pack = append(pack, _packSerializeGetContestTemplateReply(_code[i], _data[i]))
	}
	return
}
func PSerializeContestProblemTemplate(_name string, _content string) *ContestProblemTemplate {

	return &ContestProblemTemplate{
		Name:    _name,
		Content: _content,
	}
}
func SerializeContestProblemTemplate(_name string, _content string) ContestProblemTemplate {

	return ContestProblemTemplate{
		Name:    _name,
		Content: _content,
	}
}
func _packSerializeContestProblemTemplate(_name string, _content string) ContestProblemTemplate {

	return ContestProblemTemplate{
		Name:    _name,
		Content: _content,
	}
}
func PackSerializeContestProblemTemplate(_name []string, _content []string) (pack []ContestProblemTemplate) {
	for i := range _name {
		pack = append(pack, _packSerializeContestProblemTemplate(_name[i], _content[i]))
	}
	return
}
func PSerializePutContestTemplateRequest(_name string, _content string) *PutContestTemplateRequest {

	return &PutContestTemplateRequest{
		Name:    _name,
		Content: _content,
	}
}
func SerializePutContestTemplateRequest(_name string, _content string) PutContestTemplateRequest {

	return PutContestTemplateRequest{
		Name:    _name,
		Content: _content,
	}
}
func _packSerializePutContestTemplateRequest(_name string, _content string) PutContestTemplateRequest {

	return PutContestTemplateRequest{
		Name:    _name,
		Content: _content,
	}
}
func PackSerializePutContestTemplateRequest(_name []string, _content []string) (pack []PutContestTemplateRequest) {
	for i := range _name {
		pack = append(pack, _packSerializePutContestTemplateRequest(_name[i], _content[i]))
	}
	return
}
func PSerializeGetContestProblemReply(_code int, _problem *problem.Problem) *GetContestProblemReply {

	return &GetContestProblemReply{
		Code:    _code,
		Problem: _problem,
	}
}
func SerializeGetContestProblemReply(_code int, _problem *problem.Problem) GetContestProblemReply {

	return GetContestProblemReply{
		Code:    _code,
		Problem: _problem,
	}
}
func _packSerializeGetContestProblemReply(_code int, _problem *problem.Problem) GetContestProblemReply {

	return GetContestProblemReply{
		Code:    _code,
		Problem: _problem,
	}
}
func PackSerializeGetContestProblemReply(_code []int, _problem []*problem.Problem) (pack []GetContestProblemReply) {
	for i := range _code {
		pack = append(pack, _packSerializeGetContestProblemReply(_code[i], _problem[i]))
	}
	return
}
func PSerializePutContestProblemRequest() *PutContestProblemRequest {

	return &PutContestProblemRequest{}
}
func SerializePutContestProblemRequest() PutContestProblemRequest {

	return PutContestProblemRequest{}
}
func _packSerializePutContestProblemRequest() PutContestProblemRequest {

	return PutContestProblemRequest{}
}
func PackSerializePutContestProblemRequest() (pack []PutContestProblemRequest) {
	return
}
func PSerializeGetContestReply(_code int, _contest *contest.Contest) *GetContestReply {

	return &GetContestReply{
		Code:    _code,
		Contest: _contest,
	}
}
func SerializeGetContestReply(_code int, _contest *contest.Contest) GetContestReply {

	return GetContestReply{
		Code:    _code,
		Contest: _contest,
	}
}
func _packSerializeGetContestReply(_code int, _contest *contest.Contest) GetContestReply {

	return GetContestReply{
		Code:    _code,
		Contest: _contest,
	}
}
func PackSerializeGetContestReply(_code []int, _contest []*contest.Contest) (pack []GetContestReply) {
	for i := range _code {
		pack = append(pack, _packSerializeGetContestReply(_code[i], _contest[i]))
	}
	return
}
func PSerializePutContestRequest() *PutContestRequest {

	return &PutContestRequest{}
}
func SerializePutContestRequest() PutContestRequest {

	return PutContestRequest{}
}
func _packSerializePutContestRequest() PutContestRequest {

	return PutContestRequest{}
}
func PackSerializePutContestRequest() (pack []PutContestRequest) {
	return
}