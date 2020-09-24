package api

import (
	"github.com/Myriad-Dreamin/boj-v6/abstract/submission"
	"time"
)

type ListSubmissionRequest struct {
	Page         int   `json:"page" form:"page"`
	PageSize     int   `form:"page_size" json:"page_size"`
	MemOrder     *bool `json:"mem_order" form:"mem_order"`
	TimeOrder    *bool `json:"time_order" form:"time_order"`
	IdOrder      *bool `json:"id_order" form:"id_order"`
	ByUser       uint  `json:"by_user" form:"by_user"`
	OnProblem    uint  `json:"on_problem" form:"on_problem"`
	WithLanguage uint8 `json:"with_language" form:"with_language"`
	HasStatus    int64 `form:"has_status" json:"has_status"`
}

type ListSubmissionReply struct {
	Code int                        `json:"code" form:"code"`
	Data []ListSubmissionInnerReply `json:"data" form:"data"`
}

type ListSubmissionInnerReply struct {
	Id         uint      `json:"id" form:"id"`
	CreatedAt  time.Time `json:"created_at" form:"created_at"`
	ProblemId  uint      `json:"problem_id" form:"problem_id"`
	UserId     uint      `json:"user_id" form:"user_id"`
	Score      int64     `json:"score" form:"score"`
	Status     int64     `json:"status" form:"status"`
	RunTime    int64     `json:"run_time" form:"run_time"`
	RunMemory  int64     `json:"run_memory" form:"run_memory"`
	CodeLength int       `json:"code_length" form:"code_length"`
	Language   uint8     `form:"language" json:"language"`
	Shared     uint8     `json:"shared" form:"shared"`
}

type CountSubmissionRequest struct {
	Page         int   `json:"page" form:"page"`
	PageSize     int   `json:"page_size" form:"page_size"`
	MemOrder     *bool `json:"mem_order" form:"mem_order"`
	TimeOrder    *bool `json:"time_order" form:"time_order"`
	IdOrder      *bool `json:"id_order" form:"id_order"`
	ByUser       uint  `json:"by_user" form:"by_user"`
	OnProblem    uint  `json:"on_problem" form:"on_problem"`
	WithLanguage uint8 `json:"with_language" form:"with_language"`
	HasStatus    int64 `json:"has_status" form:"has_status"`
}

type CountSubmissionReply struct {
	Code int   `json:"code" form:"code"`
	Data int64 `json:"data" form:"data"`
}

type PostSubmissionRequest struct {
	Pid         uint   `json:"pid" form:"pid" route-param:"-"`
	Information string `json:"information" form:"information"`
	Shared      uint8  `json:"shared" form:"shared"`
	Language    string `json:"language" form:"language" binding:"required"`
	Code        string `json:"code" form:"code" binding:"required"`
}

type PostSubmissionReply struct {
	Code int                `form:"code" json:"code"`
	Data PostSubmissionData `json:"data" form:"data"`
}

type PostSubmissionData struct {
	Id uint `json:"id" form:"id"`
}

type GetSubmissionContentRequest struct {
}

type GetSubmissionContentReply struct {
	Code int `json:"code" form:"code"`
}

type GetSubmissionRequest struct {
}

type GetSubmissionReply struct {
	Code int                     `json:"code" form:"code"`
	Data GetSubmissionInnerReply `json:"data" form:"data"`
}

type GetSubmissionInnerReply struct {
	Id         uint      `json:"id" form:"id"`
	CreatedAt  time.Time `json:"created_at" form:"created_at"`
	ProblemId  uint      `json:"problem_id" form:"problem_id"`
	UserId     uint      `form:"user_id" json:"user_id"`
	Score      int64     `json:"score" form:"score"`
	Status     int64     `json:"status" form:"status"`
	RunTime    int64     `json:"run_time" form:"run_time"`
	RunMemory  int64     `json:"run_memory" form:"run_memory"`
	CodeLength int       `json:"code_length" form:"code_length"`
	Language   uint8     `json:"language" form:"language"`
	Shared     uint8     `json:"shared" form:"shared"`
}

type DeleteSubmissionRequest struct {
}

type DeleteSubmissionReply struct {
	Code int `json:"code" form:"code"`
}

func PSerializeListSubmissionRequest(_page int, _pageSize int, _memOrder *bool, _timeOrder *bool, _idOrder *bool, _byUser uint, _onProblem uint, _withLanguage uint8, _hasStatus int64) *ListSubmissionRequest {

	return &ListSubmissionRequest{
		Page:         _page,
		PageSize:     _pageSize,
		MemOrder:     _memOrder,
		TimeOrder:    _timeOrder,
		IdOrder:      _idOrder,
		ByUser:       _byUser,
		OnProblem:    _onProblem,
		WithLanguage: _withLanguage,
		HasStatus:    _hasStatus,
	}
}
func SerializeListSubmissionRequest(_page int, _pageSize int, _memOrder *bool, _timeOrder *bool, _idOrder *bool, _byUser uint, _onProblem uint, _withLanguage uint8, _hasStatus int64) ListSubmissionRequest {

	return ListSubmissionRequest{
		Page:         _page,
		PageSize:     _pageSize,
		MemOrder:     _memOrder,
		TimeOrder:    _timeOrder,
		IdOrder:      _idOrder,
		ByUser:       _byUser,
		OnProblem:    _onProblem,
		WithLanguage: _withLanguage,
		HasStatus:    _hasStatus,
	}
}
func _packSerializeListSubmissionRequest(_page int, _pageSize int, _memOrder *bool, _timeOrder *bool, _idOrder *bool, _byUser uint, _onProblem uint, _withLanguage uint8, _hasStatus int64) ListSubmissionRequest {

	return ListSubmissionRequest{
		Page:         _page,
		PageSize:     _pageSize,
		MemOrder:     _memOrder,
		TimeOrder:    _timeOrder,
		IdOrder:      _idOrder,
		ByUser:       _byUser,
		OnProblem:    _onProblem,
		WithLanguage: _withLanguage,
		HasStatus:    _hasStatus,
	}
}
func PackSerializeListSubmissionRequest(_page []int, _pageSize []int, _memOrder []*bool, _timeOrder []*bool, _idOrder []*bool, _byUser []uint, _onProblem []uint, _withLanguage []uint8, _hasStatus []int64) (pack []ListSubmissionRequest) {
	for i := range _page {
		pack = append(pack, _packSerializeListSubmissionRequest(_page[i], _pageSize[i], _memOrder[i], _timeOrder[i], _idOrder[i], _byUser[i], _onProblem[i], _withLanguage[i], _hasStatus[i]))
	}
	return
}
func PSerializeListSubmissionReply(_code int, _data []ListSubmissionInnerReply) *ListSubmissionReply {

	return &ListSubmissionReply{
		Code: _code,
		Data: _data,
	}
}
func SerializeListSubmissionReply(_code int, _data []ListSubmissionInnerReply) ListSubmissionReply {

	return ListSubmissionReply{
		Code: _code,
		Data: _data,
	}
}
func _packSerializeListSubmissionReply(_code int, _data []ListSubmissionInnerReply) ListSubmissionReply {

	return ListSubmissionReply{
		Code: _code,
		Data: _data,
	}
}
func PackSerializeListSubmissionReply(_code []int, _data [][]ListSubmissionInnerReply) (pack []ListSubmissionReply) {
	for i := range _code {
		pack = append(pack, _packSerializeListSubmissionReply(_code[i], _data[i]))
	}
	return
}
func PSerializeListSubmissionInnerReply(valueSubmission submission.Submission) *ListSubmissionInnerReply {

	return &ListSubmissionInnerReply{
		Id:         valueSubmission.ID,
		CreatedAt:  valueSubmission.CreatedAt,
		ProblemId:  valueSubmission.ProblemID,
		UserId:     valueSubmission.UserID,
		Score:      valueSubmission.Score,
		Status:     valueSubmission.Status,
		RunTime:    valueSubmission.RunTime,
		RunMemory:  valueSubmission.RunMemory,
		CodeLength: valueSubmission.CodeLength,
		Language:   valueSubmission.Language,
		Shared:     valueSubmission.Shared,
	}
}
func SerializeListSubmissionInnerReply(valueSubmission submission.Submission) ListSubmissionInnerReply {

	return ListSubmissionInnerReply{
		Id:         valueSubmission.ID,
		CreatedAt:  valueSubmission.CreatedAt,
		ProblemId:  valueSubmission.ProblemID,
		UserId:     valueSubmission.UserID,
		Score:      valueSubmission.Score,
		Status:     valueSubmission.Status,
		RunTime:    valueSubmission.RunTime,
		RunMemory:  valueSubmission.RunMemory,
		CodeLength: valueSubmission.CodeLength,
		Language:   valueSubmission.Language,
		Shared:     valueSubmission.Shared,
	}
}
func _packSerializeListSubmissionInnerReply(valueSubmission submission.Submission) ListSubmissionInnerReply {

	return ListSubmissionInnerReply{
		Id:         valueSubmission.ID,
		CreatedAt:  valueSubmission.CreatedAt,
		ProblemId:  valueSubmission.ProblemID,
		UserId:     valueSubmission.UserID,
		Score:      valueSubmission.Score,
		Status:     valueSubmission.Status,
		RunTime:    valueSubmission.RunTime,
		RunMemory:  valueSubmission.RunMemory,
		CodeLength: valueSubmission.CodeLength,
		Language:   valueSubmission.Language,
		Shared:     valueSubmission.Shared,
	}
}
func PackSerializeListSubmissionInnerReply(valueSubmission []submission.Submission) (pack []ListSubmissionInnerReply) {
	for i := range valueSubmission {
		pack = append(pack, _packSerializeListSubmissionInnerReply(valueSubmission[i]))
	}
	return
}
func PSerializeCountSubmissionRequest(_page int, _pageSize int, _memOrder *bool, _timeOrder *bool, _idOrder *bool, _byUser uint, _onProblem uint, _withLanguage uint8, _hasStatus int64) *CountSubmissionRequest {

	return &CountSubmissionRequest{
		Page:         _page,
		PageSize:     _pageSize,
		MemOrder:     _memOrder,
		TimeOrder:    _timeOrder,
		IdOrder:      _idOrder,
		ByUser:       _byUser,
		OnProblem:    _onProblem,
		WithLanguage: _withLanguage,
		HasStatus:    _hasStatus,
	}
}
func SerializeCountSubmissionRequest(_page int, _pageSize int, _memOrder *bool, _timeOrder *bool, _idOrder *bool, _byUser uint, _onProblem uint, _withLanguage uint8, _hasStatus int64) CountSubmissionRequest {

	return CountSubmissionRequest{
		Page:         _page,
		PageSize:     _pageSize,
		MemOrder:     _memOrder,
		TimeOrder:    _timeOrder,
		IdOrder:      _idOrder,
		ByUser:       _byUser,
		OnProblem:    _onProblem,
		WithLanguage: _withLanguage,
		HasStatus:    _hasStatus,
	}
}
func _packSerializeCountSubmissionRequest(_page int, _pageSize int, _memOrder *bool, _timeOrder *bool, _idOrder *bool, _byUser uint, _onProblem uint, _withLanguage uint8, _hasStatus int64) CountSubmissionRequest {

	return CountSubmissionRequest{
		Page:         _page,
		PageSize:     _pageSize,
		MemOrder:     _memOrder,
		TimeOrder:    _timeOrder,
		IdOrder:      _idOrder,
		ByUser:       _byUser,
		OnProblem:    _onProblem,
		WithLanguage: _withLanguage,
		HasStatus:    _hasStatus,
	}
}
func PackSerializeCountSubmissionRequest(_page []int, _pageSize []int, _memOrder []*bool, _timeOrder []*bool, _idOrder []*bool, _byUser []uint, _onProblem []uint, _withLanguage []uint8, _hasStatus []int64) (pack []CountSubmissionRequest) {
	for i := range _page {
		pack = append(pack, _packSerializeCountSubmissionRequest(_page[i], _pageSize[i], _memOrder[i], _timeOrder[i], _idOrder[i], _byUser[i], _onProblem[i], _withLanguage[i], _hasStatus[i]))
	}
	return
}
func PSerializeCountSubmissionReply(_code int, _data int64) *CountSubmissionReply {

	return &CountSubmissionReply{
		Code: _code,
		Data: _data,
	}
}
func SerializeCountSubmissionReply(_code int, _data int64) CountSubmissionReply {

	return CountSubmissionReply{
		Code: _code,
		Data: _data,
	}
}
func _packSerializeCountSubmissionReply(_code int, _data int64) CountSubmissionReply {

	return CountSubmissionReply{
		Code: _code,
		Data: _data,
	}
}
func PackSerializeCountSubmissionReply(_code []int, _data []int64) (pack []CountSubmissionReply) {
	for i := range _code {
		pack = append(pack, _packSerializeCountSubmissionReply(_code[i], _data[i]))
	}
	return
}
func PSerializePostSubmissionRequest(_pid uint, submission *submission.Submission, _language string, _code string) *PostSubmissionRequest {

	return &PostSubmissionRequest{
		Pid:         _pid,
		Information: submission.Information,
		Shared:      submission.Shared,
		Language:    _language,
		Code:        _code,
	}
}
func SerializePostSubmissionRequest(_pid uint, submission *submission.Submission, _language string, _code string) PostSubmissionRequest {

	return PostSubmissionRequest{
		Pid:         _pid,
		Information: submission.Information,
		Shared:      submission.Shared,
		Language:    _language,
		Code:        _code,
	}
}
func _packSerializePostSubmissionRequest(_pid uint, submission *submission.Submission, _language string, _code string) PostSubmissionRequest {

	return PostSubmissionRequest{
		Pid:         _pid,
		Information: submission.Information,
		Shared:      submission.Shared,
		Language:    _language,
		Code:        _code,
	}
}
func PackSerializePostSubmissionRequest(_pid []uint, submission []*submission.Submission, _language []string, _code []string) (pack []PostSubmissionRequest) {
	for i := range _pid {
		pack = append(pack, _packSerializePostSubmissionRequest(_pid[i], submission[i], _language[i], _code[i]))
	}
	return
}
func PSerializePostSubmissionReply(_code int, _data PostSubmissionData) *PostSubmissionReply {

	return &PostSubmissionReply{
		Code: _code,
		Data: _data,
	}
}
func SerializePostSubmissionReply(_code int, _data PostSubmissionData) PostSubmissionReply {

	return PostSubmissionReply{
		Code: _code,
		Data: _data,
	}
}
func _packSerializePostSubmissionReply(_code int, _data PostSubmissionData) PostSubmissionReply {

	return PostSubmissionReply{
		Code: _code,
		Data: _data,
	}
}
func PackSerializePostSubmissionReply(_code []int, _data []PostSubmissionData) (pack []PostSubmissionReply) {
	for i := range _code {
		pack = append(pack, _packSerializePostSubmissionReply(_code[i], _data[i]))
	}
	return
}
func PSerializePostSubmissionData(submission *submission.Submission) *PostSubmissionData {

	return &PostSubmissionData{
		Id: submission.ID,
	}
}
func SerializePostSubmissionData(submission *submission.Submission) PostSubmissionData {

	return PostSubmissionData{
		Id: submission.ID,
	}
}
func _packSerializePostSubmissionData(submission *submission.Submission) PostSubmissionData {

	return PostSubmissionData{
		Id: submission.ID,
	}
}
func PackSerializePostSubmissionData(submission []*submission.Submission) (pack []PostSubmissionData) {
	for i := range submission {
		pack = append(pack, _packSerializePostSubmissionData(submission[i]))
	}
	return
}
func PSerializeGetSubmissionContentRequest() *GetSubmissionContentRequest {

	return &GetSubmissionContentRequest{}
}
func SerializeGetSubmissionContentRequest() GetSubmissionContentRequest {

	return GetSubmissionContentRequest{}
}
func _packSerializeGetSubmissionContentRequest() GetSubmissionContentRequest {

	return GetSubmissionContentRequest{}
}
func PackSerializeGetSubmissionContentRequest() (pack []GetSubmissionContentRequest) {
	return
}
func PSerializeGetSubmissionContentReply(_code int) *GetSubmissionContentReply {

	return &GetSubmissionContentReply{
		Code: _code,
	}
}
func SerializeGetSubmissionContentReply(_code int) GetSubmissionContentReply {

	return GetSubmissionContentReply{
		Code: _code,
	}
}
func _packSerializeGetSubmissionContentReply(_code int) GetSubmissionContentReply {

	return GetSubmissionContentReply{
		Code: _code,
	}
}
func PackSerializeGetSubmissionContentReply(_code []int) (pack []GetSubmissionContentReply) {
	for i := range _code {
		pack = append(pack, _packSerializeGetSubmissionContentReply(_code[i]))
	}
	return
}
func PSerializeGetSubmissionRequest() *GetSubmissionRequest {

	return &GetSubmissionRequest{}
}
func SerializeGetSubmissionRequest() GetSubmissionRequest {

	return GetSubmissionRequest{}
}
func _packSerializeGetSubmissionRequest() GetSubmissionRequest {

	return GetSubmissionRequest{}
}
func PackSerializeGetSubmissionRequest() (pack []GetSubmissionRequest) {
	return
}
func PSerializeGetSubmissionReply(_code int, _data GetSubmissionInnerReply) *GetSubmissionReply {

	return &GetSubmissionReply{
		Code: _code,
		Data: _data,
	}
}
func SerializeGetSubmissionReply(_code int, _data GetSubmissionInnerReply) GetSubmissionReply {

	return GetSubmissionReply{
		Code: _code,
		Data: _data,
	}
}
func _packSerializeGetSubmissionReply(_code int, _data GetSubmissionInnerReply) GetSubmissionReply {

	return GetSubmissionReply{
		Code: _code,
		Data: _data,
	}
}
func PackSerializeGetSubmissionReply(_code []int, _data []GetSubmissionInnerReply) (pack []GetSubmissionReply) {
	for i := range _code {
		pack = append(pack, _packSerializeGetSubmissionReply(_code[i], _data[i]))
	}
	return
}
func PSerializeGetSubmissionInnerReply(submission *submission.Submission) *GetSubmissionInnerReply {

	return &GetSubmissionInnerReply{
		Id:         submission.ID,
		CreatedAt:  submission.CreatedAt,
		ProblemId:  submission.ProblemID,
		UserId:     submission.UserID,
		Score:      submission.Score,
		Status:     submission.Status,
		RunTime:    submission.RunTime,
		RunMemory:  submission.RunMemory,
		CodeLength: submission.CodeLength,
		Language:   submission.Language,
		Shared:     submission.Shared,
	}
}
func SerializeGetSubmissionInnerReply(submission *submission.Submission) GetSubmissionInnerReply {

	return GetSubmissionInnerReply{
		Id:         submission.ID,
		CreatedAt:  submission.CreatedAt,
		ProblemId:  submission.ProblemID,
		UserId:     submission.UserID,
		Score:      submission.Score,
		Status:     submission.Status,
		RunTime:    submission.RunTime,
		RunMemory:  submission.RunMemory,
		CodeLength: submission.CodeLength,
		Language:   submission.Language,
		Shared:     submission.Shared,
	}
}
func _packSerializeGetSubmissionInnerReply(submission *submission.Submission) GetSubmissionInnerReply {

	return GetSubmissionInnerReply{
		Id:         submission.ID,
		CreatedAt:  submission.CreatedAt,
		ProblemId:  submission.ProblemID,
		UserId:     submission.UserID,
		Score:      submission.Score,
		Status:     submission.Status,
		RunTime:    submission.RunTime,
		RunMemory:  submission.RunMemory,
		CodeLength: submission.CodeLength,
		Language:   submission.Language,
		Shared:     submission.Shared,
	}
}
func PackSerializeGetSubmissionInnerReply(submission []*submission.Submission) (pack []GetSubmissionInnerReply) {
	for i := range submission {
		pack = append(pack, _packSerializeGetSubmissionInnerReply(submission[i]))
	}
	return
}
func PSerializeDeleteSubmissionRequest() *DeleteSubmissionRequest {

	return &DeleteSubmissionRequest{}
}
func SerializeDeleteSubmissionRequest() DeleteSubmissionRequest {

	return DeleteSubmissionRequest{}
}
func _packSerializeDeleteSubmissionRequest() DeleteSubmissionRequest {

	return DeleteSubmissionRequest{}
}
func PackSerializeDeleteSubmissionRequest() (pack []DeleteSubmissionRequest) {
	return
}
func PSerializeDeleteSubmissionReply(_code int) *DeleteSubmissionReply {

	return &DeleteSubmissionReply{
		Code: _code,
	}
}
func SerializeDeleteSubmissionReply(_code int) DeleteSubmissionReply {

	return DeleteSubmissionReply{
		Code: _code,
	}
}
func _packSerializeDeleteSubmissionReply(_code int) DeleteSubmissionReply {

	return DeleteSubmissionReply{
		Code: _code,
	}
}
func PackSerializeDeleteSubmissionReply(_code []int) (pack []DeleteSubmissionReply) {
	for i := range _code {
		pack = append(pack, _packSerializeDeleteSubmissionReply(_code[i]))
	}
	return
}
