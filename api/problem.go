package api

import (
	"github.com/Myriad-Dreamin/boj-v6/abstract/problem"
	"github.com/Myriad-Dreamin/boj-v6/types/problem-config"
	"github.com/Myriad-Dreamin/go-model-traits/gorm-crud-dao"
)

type ListProblemRequest = gorm_crud_dao.Filter

type ListProblemReply struct {
	Code int               `form:"code" json:"code"`
	Data []problem.Problem `json:"data" form:"data"`
}

type CountProblemRequest = gorm_crud_dao.Filter

type CountProblemReply struct {
	Code int   `form:"code" json:"code"`
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
	Name    string `json:"name" form:"name" binding:"required"`
	NewName string `json:"new_name" form:"new_name" binding:"required"`
}

type ChangeProblemDescriptionRefReply struct {
	Code int `form:"code" json:"code"`
}

type PostProblemDescRequest struct {
	Name    string `json:"name" form:"name" binding:"required"`
	Content string `json:"content" form:"content"`
}

type PostProblemDescReply struct {
	Code int `form:"code" json:"code"`
}

type GetProblemDescRequest struct {
	Name string `form:"name" json:"name"`
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
	Name string `json:"name" form:"name"`
}

type DeleteProblemDescReply struct {
	Code int `json:"code" form:"code"`
}

type GetProblemRequest struct {
}

type GetProblemReply struct {
	Code int              `json:"code" form:"code"`
	Data *problem.Problem `json:"data" form:"data"`
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
