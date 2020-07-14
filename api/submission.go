package api

import (
	"github.com/Myriad-Dreamin/boj-v6/abstract/submission"
	"time"
)

type SubmissionFilter struct {
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

type ListSubmissionsReply struct {
	Code int                   `json:"code" form:"code"`
	Data []ListSubmissionReply `json:"data" form:"data"`
}

type ListSubmissionReply struct {
	Id         uint      `json:"id" form:"id"`
	CreatedAt  time.Time `json:"created_at" form:"created_at"`
	ProblemId  uint      `json:"problem_id" form:"problem_id"`
	UserId     uint      `json:"user_id" form:"user_id"`
	Score      int64     `json:"score" form:"score"`
	Status     int64     `json:"status" form:"status"`
	RunTime    int64     `json:"run_time" form:"run_time"`
	RunMemory  int64     `json:"run_memory" form:"run_memory"`
	CodeLength int       `json:"code_length" form:"code_length"`
	Language   uint8     `json:"language" form:"language"`
	Shared     uint8     `json:"shared" form:"shared"`
}

type CountSubmissionsReply struct {
	Code int   `json:"code" form:"code"`
	Data int64 `json:"data" form:"data"`
}

type PostSubmissionRequest struct {
	Information string `json:"information" form:"information"`
	Shared      uint8  `json:"shared" form:"shared"`
	Language    uint8  `binding:"required" json:"language" form:"language"`
	Code        string `json:"code" form:"code" binding:"required"`
}

type PostSubmissionReply struct {
	Code int  `json:"code" form:"code"`
	Id   uint `json:"id" form:"id"`
}

type GetSubmissionReply struct {
	Code       int                     `json:"code" form:"code"`
	Submission GetSubmissionInnerReply `json:"submission" form:"submission"`
}

type GetSubmissionInnerReply struct {
	Id         uint      `json:"id" form:"id"`
	CreatedAt  time.Time `json:"created_at" form:"created_at"`
	ProblemId  uint      `json:"problem_id" form:"problem_id"`
	UserId     uint      `json:"user_id" form:"user_id"`
	Score      int64     `json:"score" form:"score"`
	Status     int64     `json:"status" form:"status"`
	RunTime    int64     `json:"run_time" form:"run_time"`
	RunMemory  int64     `json:"run_memory" form:"run_memory"`
	CodeLength int       `json:"code_length" form:"code_length"`
	Language   uint8     `json:"language" form:"language"`
	Shared     uint8     `json:"shared" form:"shared"`
}

func PSerializeSubmissionFilter(_page int, _pageSize int, _memOrder *bool, _timeOrder *bool, _idOrder *bool, _byUser uint, _onProblem uint, _withLanguage uint8, _hasStatus int64) *SubmissionFilter {

	return &SubmissionFilter{
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
func SerializeSubmissionFilter(_page int, _pageSize int, _memOrder *bool, _timeOrder *bool, _idOrder *bool, _byUser uint, _onProblem uint, _withLanguage uint8, _hasStatus int64) SubmissionFilter {

	return SubmissionFilter{
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
func _packSerializeSubmissionFilter(_page int, _pageSize int, _memOrder *bool, _timeOrder *bool, _idOrder *bool, _byUser uint, _onProblem uint, _withLanguage uint8, _hasStatus int64) SubmissionFilter {

	return SubmissionFilter{
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
func PackSerializeSubmissionFilter(_page []int, _pageSize []int, _memOrder []*bool, _timeOrder []*bool, _idOrder []*bool, _byUser []uint, _onProblem []uint, _withLanguage []uint8, _hasStatus []int64) (pack []SubmissionFilter) {
	for i := range _page {
		pack = append(pack, _packSerializeSubmissionFilter(_page[i], _pageSize[i], _memOrder[i], _timeOrder[i], _idOrder[i], _byUser[i], _onProblem[i], _withLanguage[i], _hasStatus[i]))
	}
	return
}
func PSerializeListSubmissionsReply(_code int, _data []ListSubmissionReply) *ListSubmissionsReply {

	return &ListSubmissionsReply{
		Code: _code,
		Data: _data,
	}
}
func SerializeListSubmissionsReply(_code int, _data []ListSubmissionReply) ListSubmissionsReply {

	return ListSubmissionsReply{
		Code: _code,
		Data: _data,
	}
}
func _packSerializeListSubmissionsReply(_code int, _data []ListSubmissionReply) ListSubmissionsReply {

	return ListSubmissionsReply{
		Code: _code,
		Data: _data,
	}
}
func PackSerializeListSubmissionsReply(_code []int, _data [][]ListSubmissionReply) (pack []ListSubmissionsReply) {
	for i := range _code {
		pack = append(pack, _packSerializeListSubmissionsReply(_code[i], _data[i]))
	}
	return
}
func PSerializeListSubmissionReply(valueSubmission submission.Submission) *ListSubmissionReply {

	return &ListSubmissionReply{
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
func SerializeListSubmissionReply(valueSubmission submission.Submission) ListSubmissionReply {

	return ListSubmissionReply{
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
func _packSerializeListSubmissionReply(valueSubmission submission.Submission) ListSubmissionReply {

	return ListSubmissionReply{
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
func PackSerializeListSubmissionReply(valueSubmission []submission.Submission) (pack []ListSubmissionReply) {
	for i := range valueSubmission {
		pack = append(pack, _packSerializeListSubmissionReply(valueSubmission[i]))
	}
	return
}
func PSerializeCountSubmissionsReply(_code int, _data int64) *CountSubmissionsReply {

	return &CountSubmissionsReply{
		Code: _code,
		Data: _data,
	}
}
func SerializeCountSubmissionsReply(_code int, _data int64) CountSubmissionsReply {

	return CountSubmissionsReply{
		Code: _code,
		Data: _data,
	}
}
func _packSerializeCountSubmissionsReply(_code int, _data int64) CountSubmissionsReply {

	return CountSubmissionsReply{
		Code: _code,
		Data: _data,
	}
}
func PackSerializeCountSubmissionsReply(_code []int, _data []int64) (pack []CountSubmissionsReply) {
	for i := range _code {
		pack = append(pack, _packSerializeCountSubmissionsReply(_code[i], _data[i]))
	}
	return
}
func PSerializePostSubmissionRequest(submission *submission.Submission, _code string) *PostSubmissionRequest {

	return &PostSubmissionRequest{
		Information: submission.Information,
		Shared:      submission.Shared,
		Language:    submission.Language,
		Code:        _code,
	}
}
func SerializePostSubmissionRequest(submission *submission.Submission, _code string) PostSubmissionRequest {

	return PostSubmissionRequest{
		Information: submission.Information,
		Shared:      submission.Shared,
		Language:    submission.Language,
		Code:        _code,
	}
}
func _packSerializePostSubmissionRequest(submission *submission.Submission, _code string) PostSubmissionRequest {

	return PostSubmissionRequest{
		Information: submission.Information,
		Shared:      submission.Shared,
		Language:    submission.Language,
		Code:        _code,
	}
}
func PackSerializePostSubmissionRequest(submission []*submission.Submission, _code []string) (pack []PostSubmissionRequest) {
	for i := range submission {
		pack = append(pack, _packSerializePostSubmissionRequest(submission[i], _code[i]))
	}
	return
}
func PSerializePostSubmissionReply(_code int, submission *submission.Submission) *PostSubmissionReply {

	return &PostSubmissionReply{
		Code: _code,
		Id:   submission.ID,
	}
}
func SerializePostSubmissionReply(_code int, submission *submission.Submission) PostSubmissionReply {

	return PostSubmissionReply{
		Code: _code,
		Id:   submission.ID,
	}
}
func _packSerializePostSubmissionReply(_code int, submission *submission.Submission) PostSubmissionReply {

	return PostSubmissionReply{
		Code: _code,
		Id:   submission.ID,
	}
}
func PackSerializePostSubmissionReply(_code []int, submission []*submission.Submission) (pack []PostSubmissionReply) {
	for i := range _code {
		pack = append(pack, _packSerializePostSubmissionReply(_code[i], submission[i]))
	}
	return
}
func PSerializeGetSubmissionReply(_code int, _submission GetSubmissionInnerReply) *GetSubmissionReply {

	return &GetSubmissionReply{
		Code:       _code,
		Submission: _submission,
	}
}
func SerializeGetSubmissionReply(_code int, _submission GetSubmissionInnerReply) GetSubmissionReply {

	return GetSubmissionReply{
		Code:       _code,
		Submission: _submission,
	}
}
func _packSerializeGetSubmissionReply(_code int, _submission GetSubmissionInnerReply) GetSubmissionReply {

	return GetSubmissionReply{
		Code:       _code,
		Submission: _submission,
	}
}
func PackSerializeGetSubmissionReply(_code []int, _submission []GetSubmissionInnerReply) (pack []GetSubmissionReply) {
	for i := range _code {
		pack = append(pack, _packSerializeGetSubmissionReply(_code[i], _submission[i]))
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
