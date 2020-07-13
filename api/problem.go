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
	Title       string                       `json:"title" form:"title" binding:"required"`
	Description string                       `json:"description" form:"description"`
	Config      *problemconfig.ProblemConfig `json:"config" form:"config"`
}

type PostProblemReply struct {
	Code int  `json:"code" form:"code"`
	Id   uint `json:"id" form:"id"`
}

type ChangeDescriptionRefRequest struct {
	Name    string `form:"name" binding:"required" json:"name"`
	NewName string `form:"new_name" binding:"required" json:"new_name"`
}

type PostDescRequest struct {
	Name    string `json:"name" form:"name" binding:"required"`
	Content string `json:"content" form:"content"`
}

type GetDescRequest struct {
	Name string `json:"name" form:"name"`
}

type GetDescReply struct {
	Code int         `json:"code" form:"code"`
	Data ProblemDesc `json:"data" form:"data"`
}

type ProblemDesc struct {
	Name    string `json:"name" form:"name"`
	Content string `json:"content" form:"content"`
}

type PutDescRequest struct {
	Name    string `form:"name" binding:"required" json:"name"`
	Content string `json:"content" form:"content"`
}

type GetProblemReply struct {
	Code    int              `json:"code" form:"code"`
	Problem *problem.Problem `json:"problem" form:"problem"`
}

type PutProblemRequest struct {
	Title          string `json:"title" form:"title"`
	Description    string `json:"description" form:"description"`
	DescriptionRef string `json:"description_ref" form:"description_ref"`
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
func PSerializeChangeDescriptionRefRequest(_name string, _newName string) *ChangeDescriptionRefRequest {

	return &ChangeDescriptionRefRequest{
		Name:    _name,
		NewName: _newName,
	}
}
func SerializeChangeDescriptionRefRequest(_name string, _newName string) ChangeDescriptionRefRequest {

	return ChangeDescriptionRefRequest{
		Name:    _name,
		NewName: _newName,
	}
}
func _packSerializeChangeDescriptionRefRequest(_name string, _newName string) ChangeDescriptionRefRequest {

	return ChangeDescriptionRefRequest{
		Name:    _name,
		NewName: _newName,
	}
}
func PackSerializeChangeDescriptionRefRequest(_name []string, _newName []string) (pack []ChangeDescriptionRefRequest) {
	for i := range _name {
		pack = append(pack, _packSerializeChangeDescriptionRefRequest(_name[i], _newName[i]))
	}
	return
}
func PSerializePostDescRequest(_name string, _content string) *PostDescRequest {

	return &PostDescRequest{
		Name:    _name,
		Content: _content,
	}
}
func SerializePostDescRequest(_name string, _content string) PostDescRequest {

	return PostDescRequest{
		Name:    _name,
		Content: _content,
	}
}
func _packSerializePostDescRequest(_name string, _content string) PostDescRequest {

	return PostDescRequest{
		Name:    _name,
		Content: _content,
	}
}
func PackSerializePostDescRequest(_name []string, _content []string) (pack []PostDescRequest) {
	for i := range _name {
		pack = append(pack, _packSerializePostDescRequest(_name[i], _content[i]))
	}
	return
}
func PSerializeGetDescRequest(_name string) *GetDescRequest {

	return &GetDescRequest{
		Name: _name,
	}
}
func SerializeGetDescRequest(_name string) GetDescRequest {

	return GetDescRequest{
		Name: _name,
	}
}
func _packSerializeGetDescRequest(_name string) GetDescRequest {

	return GetDescRequest{
		Name: _name,
	}
}
func PackSerializeGetDescRequest(_name []string) (pack []GetDescRequest) {
	for i := range _name {
		pack = append(pack, _packSerializeGetDescRequest(_name[i]))
	}
	return
}
func PSerializeGetDescReply(_code int, _data ProblemDesc) *GetDescReply {

	return &GetDescReply{
		Code: _code,
		Data: _data,
	}
}
func SerializeGetDescReply(_code int, _data ProblemDesc) GetDescReply {

	return GetDescReply{
		Code: _code,
		Data: _data,
	}
}
func _packSerializeGetDescReply(_code int, _data ProblemDesc) GetDescReply {

	return GetDescReply{
		Code: _code,
		Data: _data,
	}
}
func PackSerializeGetDescReply(_code []int, _data []ProblemDesc) (pack []GetDescReply) {
	for i := range _code {
		pack = append(pack, _packSerializeGetDescReply(_code[i], _data[i]))
	}
	return
}
func PSerializeProblemDesc(_name string, _content string) *ProblemDesc {

	return &ProblemDesc{
		Name:    _name,
		Content: _content,
	}
}
func SerializeProblemDesc(_name string, _content string) ProblemDesc {

	return ProblemDesc{
		Name:    _name,
		Content: _content,
	}
}
func _packSerializeProblemDesc(_name string, _content string) ProblemDesc {

	return ProblemDesc{
		Name:    _name,
		Content: _content,
	}
}
func PackSerializeProblemDesc(_name []string, _content []string) (pack []ProblemDesc) {
	for i := range _name {
		pack = append(pack, _packSerializeProblemDesc(_name[i], _content[i]))
	}
	return
}
func PSerializePutDescRequest(_name string, _content string) *PutDescRequest {

	return &PutDescRequest{
		Name:    _name,
		Content: _content,
	}
}
func SerializePutDescRequest(_name string, _content string) PutDescRequest {

	return PutDescRequest{
		Name:    _name,
		Content: _content,
	}
}
func _packSerializePutDescRequest(_name string, _content string) PutDescRequest {

	return PutDescRequest{
		Name:    _name,
		Content: _content,
	}
}
func PackSerializePutDescRequest(_name []string, _content []string) (pack []PutDescRequest) {
	for i := range _name {
		pack = append(pack, _packSerializePutDescRequest(_name[i], _content[i]))
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
