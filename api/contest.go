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
	Contest *contest.Contest `json:"contest" form:"contest"`
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

type ChangeContestProblemDescriptionRefRequest struct {
	Name    string `form:"name" binding:"required" json:"name"`
	NewName string `form:"new_name" binding:"required" json:"new_name"`
}

type PostContestProblemDescRequest struct {
	Name    string `json:"name" form:"name" binding:"required"`
	Content string `json:"content" form:"content"`
}

type GetContestProblemDescRequest struct {
	Name string `json:"name" form:"name"`
}

type GetContestProblemDescReply struct {
	Code int                `json:"code" form:"code"`
	Data ContestProblemDesc `json:"data" form:"data"`
}

type ContestProblemDesc struct {
	Name    string `json:"name" form:"name"`
	Content string `json:"content" form:"content"`
}

type PutContestProblemDescRequest struct {
	Name    string `json:"name" form:"name" binding:"required"`
	Content string `form:"content" json:"content"`
}

type GetContestProblemReply struct {
	Code    int              `json:"code" form:"code"`
	Problem *problem.Problem `json:"problem" form:"problem"`
}

type PutContestProblemRequest struct {
	Title          string `json:"title" form:"title"`
	Description    string `json:"description" form:"description"`
	DescriptionRef string `json:"description_ref" form:"description_ref"`
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

type GetContestReply struct {
	Code    int              `json:"code" form:"code"`
	Contest *contest.Contest `json:"contest" form:"contest"`
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
func PSerializeGetContestProblemDescReply(_code int, _data ContestProblemDesc) *GetContestProblemDescReply {

	return &GetContestProblemDescReply{
		Code: _code,
		Data: _data,
	}
}
func SerializeGetContestProblemDescReply(_code int, _data ContestProblemDesc) GetContestProblemDescReply {

	return GetContestProblemDescReply{
		Code: _code,
		Data: _data,
	}
}
func _packSerializeGetContestProblemDescReply(_code int, _data ContestProblemDesc) GetContestProblemDescReply {

	return GetContestProblemDescReply{
		Code: _code,
		Data: _data,
	}
}
func PackSerializeGetContestProblemDescReply(_code []int, _data []ContestProblemDesc) (pack []GetContestProblemDescReply) {
	for i := range _code {
		pack = append(pack, _packSerializeGetContestProblemDescReply(_code[i], _data[i]))
	}
	return
}
func PSerializeContestProblemDesc(_name string, _content string) *ContestProblemDesc {

	return &ContestProblemDesc{
		Name:    _name,
		Content: _content,
	}
}
func SerializeContestProblemDesc(_name string, _content string) ContestProblemDesc {

	return ContestProblemDesc{
		Name:    _name,
		Content: _content,
	}
}
func _packSerializeContestProblemDesc(_name string, _content string) ContestProblemDesc {

	return ContestProblemDesc{
		Name:    _name,
		Content: _content,
	}
}
func PackSerializeContestProblemDesc(_name []string, _content []string) (pack []ContestProblemDesc) {
	for i := range _name {
		pack = append(pack, _packSerializeContestProblemDesc(_name[i], _content[i]))
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
func PSerializePutContestProblemRequest(problem *problem.Problem) *PutContestProblemRequest {

	return &PutContestProblemRequest{
		Title:          problem.Title,
		Description:    problem.Description,
		DescriptionRef: problem.DescriptionRef,
	}
}
func SerializePutContestProblemRequest(problem *problem.Problem) PutContestProblemRequest {

	return PutContestProblemRequest{
		Title:          problem.Title,
		Description:    problem.Description,
		DescriptionRef: problem.DescriptionRef,
	}
}
func _packSerializePutContestProblemRequest(problem *problem.Problem) PutContestProblemRequest {

	return PutContestProblemRequest{
		Title:          problem.Title,
		Description:    problem.Description,
		DescriptionRef: problem.DescriptionRef,
	}
}
func PackSerializePutContestProblemRequest(problem []*problem.Problem) (pack []PutContestProblemRequest) {
	for i := range problem {
		pack = append(pack, _packSerializePutContestProblemRequest(problem[i]))
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
