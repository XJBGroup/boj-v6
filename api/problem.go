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
	Code int             `json:"code" form:"code"`
	Data PostProblemData `json:"data" form:"data"`
}

type PostProblemData struct {
	Id uint `json:"id" form:"id"`
}

type ChangeProblemDescriptionRefRequest struct {
	Name    string `form:"name" binding:"required" json:"name"`
	NewName string `json:"new_name" form:"new_name" binding:"required"`
}

type PostProblemDescRequest struct {
	Name    string `json:"name" form:"name" binding:"required"`
	Content string `form:"content" json:"content"`
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

type DeleteProblemDescRequest struct {
	Name string `json:"name" form:"name"`
}

type GetProblemReply struct {
	Code int              `form:"code" json:"code"`
	Data *problem.Problem `form:"data" json:"data"`
}

type PutProblemRequest struct {
	Title          string `form:"title" json:"title"`
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
func PSerializeGetProblemReply(_code int, _data *problem.Problem) *GetProblemReply {

	return &GetProblemReply{
		Code: _code,
		Data: _data,
	}
}
func SerializeGetProblemReply(_code int, _data *problem.Problem) GetProblemReply {

	return GetProblemReply{
		Code: _code,
		Data: _data,
	}
}
func _packSerializeGetProblemReply(_code int, _data *problem.Problem) GetProblemReply {

	return GetProblemReply{
		Code: _code,
		Data: _data,
	}
}
func PackSerializeGetProblemReply(_code []int, _data []*problem.Problem) (pack []GetProblemReply) {
	for i := range _code {
		pack = append(pack, _packSerializeGetProblemReply(_code[i], _data[i]))
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
