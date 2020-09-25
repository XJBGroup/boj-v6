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
	Code int   `form:"code" json:"code"`
	Data []int `json:"data" form:"data"`
}

type PostProblemRequest struct {
	Title       string                       `form:"title" binding:"required" json:"title"`
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

type ChangeProblemDescriptionRefReply struct {
	Code int `form:"code" json:"code"`
}

type PostProblemDescRequest struct {
	Name    string `form:"name" binding:"required" json:"name"`
	Content string `json:"content" form:"content"`
}

type PostProblemDescReply struct {
	Code int `json:"code" form:"code"`
}

type GetProblemDescRequest struct {
	Name string `json:"name" form:"name"`
}

type GetProblemDescReply struct {
	Code int    `json:"code" form:"code"`
	Data string `json:"data" form:"data"`
}

type PutProblemDescRequest struct {
	Name    string `form:"name" binding:"required" json:"name"`
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

type ProblemFSReadConfigRequest struct {
	Path string `json:"path" form:"path" binding:"required"`
}

type ProblemFSReadConfigReply struct {
	Code int                           `json:"code" form:"code"`
	Data ProblemFSReadConfigInnerReply `json:"data" form:"data"`
}

type ProblemFSReadConfigInnerReply struct {
}

type ProblemFSWriteConfigRequest struct {
	Path string `form:"path" binding:"required" json:"path"`
}

type ProblemFSWriteConfigReply struct {
	Code int                            `json:"code" form:"code"`
	Data ProblemFSWriteConfigInnerReply `form:"data" json:"data"`
}

type ProblemFSWriteConfigInnerReply struct {
}

type ProblemFSPutConfigRequest struct {
	Path string `json:"path" form:"path" binding:"required"`
}

type ProblemFSPutConfigReply struct {
	Code int                          `json:"code" form:"code"`
	Data ProblemFSPutConfigInnerReply `form:"data" json:"data"`
}

type ProblemFSPutConfigInnerReply struct {
}

type ProblemFSReadRequest struct {
	Path string `json:"path" form:"path" binding:"required"`
}

type ProblemFSReadReply struct {
	Code int                     `json:"code" form:"code"`
	Data ProblemFSReadInnerReply `json:"data" form:"data"`
}

type ProblemFSReadInnerReply struct {
}

type ProblemFSStatRequest struct {
	Path string `json:"path" form:"path" binding:"required"`
}

type ProblemFSStatReply struct {
	Code int                     `json:"code" form:"code"`
	Data ProblemFSStatInnerReply `json:"data" form:"data"`
}

type ProblemFSStatInnerReply struct {
	Name    string    `json:"name" form:"name"`
	Size    int64     `json:"size" form:"size"`
	IsDir   bool      `json:"is_dir" form:"is_dir"`
	ModTime time.Time `json:"mod_time" form:"mod_time"`
}

type ProblemFSWriteRequest struct {
	Path string `json:"path" form:"path" binding:"required"`
}

type ProblemFSWriteReply struct {
	Code int                      `json:"code" form:"code"`
	Data ProblemFSWriteInnerReply `json:"data" form:"data"`
}

type ProblemFSWriteInnerReply struct {
}

type ProblemFSRemoveRequest struct {
	Path string `binding:"required" json:"path" form:"path"`
}

type ProblemFSRemoveReply struct {
	Code int                       `json:"code" form:"code"`
	Data ProblemFSRemoveInnerReply `json:"data" form:"data"`
}

type ProblemFSRemoveInnerReply struct {
}

type ProblemFSZipReadRequest struct {
	Path string `json:"path" form:"path" binding:"required"`
}

type ProblemFSZipReadReply struct {
	Code int                        `json:"code" form:"code"`
	Data ProblemFSZipReadInnerReply `json:"data" form:"data"`
}

type ProblemFSZipReadInnerReply struct {
}

type ProblemFSLSRequest struct {
	Path string `json:"path" form:"path" binding:"required"`
}

type ProblemFSLSReply struct {
	Code int                     `form:"code" json:"code"`
	Data []ProblemFSLSInnerReply `json:"data" form:"data"`
}

type ProblemFSLSInnerReply struct {
	Name    string    `json:"name" form:"name"`
	Size    int64     `json:"size" form:"size"`
	IsDir   bool      `json:"is_dir" form:"is_dir"`
	ModTime time.Time `json:"mod_time" form:"mod_time"`
}

type ProblemFSWritesRequest struct {
	Path string `json:"path" form:"path" binding:"required"`
}

type ProblemFSWritesReply struct {
	Code int                       `json:"code" form:"code"`
	Data ProblemFSWritesInnerReply `json:"data" form:"data"`
}

type ProblemFSWritesInnerReply struct {
}

type ProblemFSMkdirRequest struct {
	Path string `json:"path" form:"path" binding:"required"`
}

type ProblemFSMkdirReply struct {
	Code int                      `json:"code" form:"code"`
	Data ProblemFSMkdirInnerReply `json:"data" form:"data"`
}

type ProblemFSMkdirInnerReply struct {
}

type ProblemFSRemoveAllRequest struct {
	Path string `json:"path" form:"path" binding:"required"`
}

type ProblemFSRemoveAllReply struct {
	Code int                          `json:"code" form:"code"`
	Data ProblemFSRemoveAllInnerReply `json:"data" form:"data"`
}

type ProblemFSRemoveAllInnerReply struct {
}

type ListProblemDescRequest = gorm_crud_dao.Filter

type ListProblemDescReply struct {
	Code int               `form:"code" json:"code"`
	Data []ProblemDescData `json:"data" form:"data"`
}

type ProblemDescData struct {
	Name      string    `json:"name" form:"name"`
	UpdatedAt time.Time `json:"updated_at" form:"updated_at"`
}

type CountProblemDescRequest = gorm_crud_dao.Filter

type CountProblemDescReply struct {
	Code int   `json:"code" form:"code"`
	Data int64 `form:"data" json:"data"`
}

type GetProblemRequest struct {
}

type GetProblemReply struct {
	Code int            `json:"code" form:"code"`
	Data GetProblemData `json:"data" form:"data"`
}

type GetProblemData struct {
	Id              uint                 `form:"id" json:"id"`
	CreatedAt       time.Time            `form:"created_at" json:"created_at"`
	UpdatedAt       time.Time            `json:"updated_at" form:"updated_at"`
	IsSpj           bool                 `json:"is_spj" form:"is_spj"`
	Title           string               `json:"title" form:"title"`
	Description     string               `json:"description" form:"description"`
	DescriptionRef  string               `json:"description_ref" form:"description_ref"`
	TimeLimit       int64                `json:"time_limit" form:"time_limit"`
	MemoryLimit     int64                `json:"memory_limit" form:"memory_limit"`
	CodeLengthLimit int64                `json:"code_length_limit" form:"code_length_limit"`
	Author          GetProblemAuthorData `json:"author" form:"author"`
}

type GetProblemAuthorData struct {
	Id       uint   `json:"id" form:"id"`
	NickName string `json:"nick_name" form:"nick_name"`
}

type PutProblemRequest struct {
	Title          string `json:"title" form:"title"`
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
func PSerializeProblemFSReadConfigRequest(_path string) *ProblemFSReadConfigRequest {

	return &ProblemFSReadConfigRequest{
		Path: _path,
	}
}
func SerializeProblemFSReadConfigRequest(_path string) ProblemFSReadConfigRequest {

	return ProblemFSReadConfigRequest{
		Path: _path,
	}
}
func _packSerializeProblemFSReadConfigRequest(_path string) ProblemFSReadConfigRequest {

	return ProblemFSReadConfigRequest{
		Path: _path,
	}
}
func PackSerializeProblemFSReadConfigRequest(_path []string) (pack []ProblemFSReadConfigRequest) {
	for i := range _path {
		pack = append(pack, _packSerializeProblemFSReadConfigRequest(_path[i]))
	}
	return
}
func PSerializeProblemFSReadConfigReply(_code int, _data ProblemFSReadConfigInnerReply) *ProblemFSReadConfigReply {

	return &ProblemFSReadConfigReply{
		Code: _code,
		Data: _data,
	}
}
func SerializeProblemFSReadConfigReply(_code int, _data ProblemFSReadConfigInnerReply) ProblemFSReadConfigReply {

	return ProblemFSReadConfigReply{
		Code: _code,
		Data: _data,
	}
}
func _packSerializeProblemFSReadConfigReply(_code int, _data ProblemFSReadConfigInnerReply) ProblemFSReadConfigReply {

	return ProblemFSReadConfigReply{
		Code: _code,
		Data: _data,
	}
}
func PackSerializeProblemFSReadConfigReply(_code []int, _data []ProblemFSReadConfigInnerReply) (pack []ProblemFSReadConfigReply) {
	for i := range _code {
		pack = append(pack, _packSerializeProblemFSReadConfigReply(_code[i], _data[i]))
	}
	return
}
func PSerializeProblemFSReadConfigInnerReply() *ProblemFSReadConfigInnerReply {

	return &ProblemFSReadConfigInnerReply{}
}
func SerializeProblemFSReadConfigInnerReply() ProblemFSReadConfigInnerReply {

	return ProblemFSReadConfigInnerReply{}
}
func _packSerializeProblemFSReadConfigInnerReply() ProblemFSReadConfigInnerReply {

	return ProblemFSReadConfigInnerReply{}
}
func PackSerializeProblemFSReadConfigInnerReply() (pack []ProblemFSReadConfigInnerReply) {
	return
}
func PSerializeProblemFSWriteConfigRequest(_path string) *ProblemFSWriteConfigRequest {

	return &ProblemFSWriteConfigRequest{
		Path: _path,
	}
}
func SerializeProblemFSWriteConfigRequest(_path string) ProblemFSWriteConfigRequest {

	return ProblemFSWriteConfigRequest{
		Path: _path,
	}
}
func _packSerializeProblemFSWriteConfigRequest(_path string) ProblemFSWriteConfigRequest {

	return ProblemFSWriteConfigRequest{
		Path: _path,
	}
}
func PackSerializeProblemFSWriteConfigRequest(_path []string) (pack []ProblemFSWriteConfigRequest) {
	for i := range _path {
		pack = append(pack, _packSerializeProblemFSWriteConfigRequest(_path[i]))
	}
	return
}
func PSerializeProblemFSWriteConfigReply(_code int, _data ProblemFSWriteConfigInnerReply) *ProblemFSWriteConfigReply {

	return &ProblemFSWriteConfigReply{
		Code: _code,
		Data: _data,
	}
}
func SerializeProblemFSWriteConfigReply(_code int, _data ProblemFSWriteConfigInnerReply) ProblemFSWriteConfigReply {

	return ProblemFSWriteConfigReply{
		Code: _code,
		Data: _data,
	}
}
func _packSerializeProblemFSWriteConfigReply(_code int, _data ProblemFSWriteConfigInnerReply) ProblemFSWriteConfigReply {

	return ProblemFSWriteConfigReply{
		Code: _code,
		Data: _data,
	}
}
func PackSerializeProblemFSWriteConfigReply(_code []int, _data []ProblemFSWriteConfigInnerReply) (pack []ProblemFSWriteConfigReply) {
	for i := range _code {
		pack = append(pack, _packSerializeProblemFSWriteConfigReply(_code[i], _data[i]))
	}
	return
}
func PSerializeProblemFSWriteConfigInnerReply() *ProblemFSWriteConfigInnerReply {

	return &ProblemFSWriteConfigInnerReply{}
}
func SerializeProblemFSWriteConfigInnerReply() ProblemFSWriteConfigInnerReply {

	return ProblemFSWriteConfigInnerReply{}
}
func _packSerializeProblemFSWriteConfigInnerReply() ProblemFSWriteConfigInnerReply {

	return ProblemFSWriteConfigInnerReply{}
}
func PackSerializeProblemFSWriteConfigInnerReply() (pack []ProblemFSWriteConfigInnerReply) {
	return
}
func PSerializeProblemFSPutConfigRequest(_path string) *ProblemFSPutConfigRequest {

	return &ProblemFSPutConfigRequest{
		Path: _path,
	}
}
func SerializeProblemFSPutConfigRequest(_path string) ProblemFSPutConfigRequest {

	return ProblemFSPutConfigRequest{
		Path: _path,
	}
}
func _packSerializeProblemFSPutConfigRequest(_path string) ProblemFSPutConfigRequest {

	return ProblemFSPutConfigRequest{
		Path: _path,
	}
}
func PackSerializeProblemFSPutConfigRequest(_path []string) (pack []ProblemFSPutConfigRequest) {
	for i := range _path {
		pack = append(pack, _packSerializeProblemFSPutConfigRequest(_path[i]))
	}
	return
}
func PSerializeProblemFSPutConfigReply(_code int, _data ProblemFSPutConfigInnerReply) *ProblemFSPutConfigReply {

	return &ProblemFSPutConfigReply{
		Code: _code,
		Data: _data,
	}
}
func SerializeProblemFSPutConfigReply(_code int, _data ProblemFSPutConfigInnerReply) ProblemFSPutConfigReply {

	return ProblemFSPutConfigReply{
		Code: _code,
		Data: _data,
	}
}
func _packSerializeProblemFSPutConfigReply(_code int, _data ProblemFSPutConfigInnerReply) ProblemFSPutConfigReply {

	return ProblemFSPutConfigReply{
		Code: _code,
		Data: _data,
	}
}
func PackSerializeProblemFSPutConfigReply(_code []int, _data []ProblemFSPutConfigInnerReply) (pack []ProblemFSPutConfigReply) {
	for i := range _code {
		pack = append(pack, _packSerializeProblemFSPutConfigReply(_code[i], _data[i]))
	}
	return
}
func PSerializeProblemFSPutConfigInnerReply() *ProblemFSPutConfigInnerReply {

	return &ProblemFSPutConfigInnerReply{}
}
func SerializeProblemFSPutConfigInnerReply() ProblemFSPutConfigInnerReply {

	return ProblemFSPutConfigInnerReply{}
}
func _packSerializeProblemFSPutConfigInnerReply() ProblemFSPutConfigInnerReply {

	return ProblemFSPutConfigInnerReply{}
}
func PackSerializeProblemFSPutConfigInnerReply() (pack []ProblemFSPutConfigInnerReply) {
	return
}
func PSerializeProblemFSReadRequest(_path string) *ProblemFSReadRequest {

	return &ProblemFSReadRequest{
		Path: _path,
	}
}
func SerializeProblemFSReadRequest(_path string) ProblemFSReadRequest {

	return ProblemFSReadRequest{
		Path: _path,
	}
}
func _packSerializeProblemFSReadRequest(_path string) ProblemFSReadRequest {

	return ProblemFSReadRequest{
		Path: _path,
	}
}
func PackSerializeProblemFSReadRequest(_path []string) (pack []ProblemFSReadRequest) {
	for i := range _path {
		pack = append(pack, _packSerializeProblemFSReadRequest(_path[i]))
	}
	return
}
func PSerializeProblemFSReadReply(_code int, _data ProblemFSReadInnerReply) *ProblemFSReadReply {

	return &ProblemFSReadReply{
		Code: _code,
		Data: _data,
	}
}
func SerializeProblemFSReadReply(_code int, _data ProblemFSReadInnerReply) ProblemFSReadReply {

	return ProblemFSReadReply{
		Code: _code,
		Data: _data,
	}
}
func _packSerializeProblemFSReadReply(_code int, _data ProblemFSReadInnerReply) ProblemFSReadReply {

	return ProblemFSReadReply{
		Code: _code,
		Data: _data,
	}
}
func PackSerializeProblemFSReadReply(_code []int, _data []ProblemFSReadInnerReply) (pack []ProblemFSReadReply) {
	for i := range _code {
		pack = append(pack, _packSerializeProblemFSReadReply(_code[i], _data[i]))
	}
	return
}
func PSerializeProblemFSReadInnerReply() *ProblemFSReadInnerReply {

	return &ProblemFSReadInnerReply{}
}
func SerializeProblemFSReadInnerReply() ProblemFSReadInnerReply {

	return ProblemFSReadInnerReply{}
}
func _packSerializeProblemFSReadInnerReply() ProblemFSReadInnerReply {

	return ProblemFSReadInnerReply{}
}
func PackSerializeProblemFSReadInnerReply() (pack []ProblemFSReadInnerReply) {
	return
}
func PSerializeProblemFSStatRequest(_path string) *ProblemFSStatRequest {

	return &ProblemFSStatRequest{
		Path: _path,
	}
}
func SerializeProblemFSStatRequest(_path string) ProblemFSStatRequest {

	return ProblemFSStatRequest{
		Path: _path,
	}
}
func _packSerializeProblemFSStatRequest(_path string) ProblemFSStatRequest {

	return ProblemFSStatRequest{
		Path: _path,
	}
}
func PackSerializeProblemFSStatRequest(_path []string) (pack []ProblemFSStatRequest) {
	for i := range _path {
		pack = append(pack, _packSerializeProblemFSStatRequest(_path[i]))
	}
	return
}
func PSerializeProblemFSStatReply(_code int, _data ProblemFSStatInnerReply) *ProblemFSStatReply {

	return &ProblemFSStatReply{
		Code: _code,
		Data: _data,
	}
}
func SerializeProblemFSStatReply(_code int, _data ProblemFSStatInnerReply) ProblemFSStatReply {

	return ProblemFSStatReply{
		Code: _code,
		Data: _data,
	}
}
func _packSerializeProblemFSStatReply(_code int, _data ProblemFSStatInnerReply) ProblemFSStatReply {

	return ProblemFSStatReply{
		Code: _code,
		Data: _data,
	}
}
func PackSerializeProblemFSStatReply(_code []int, _data []ProblemFSStatInnerReply) (pack []ProblemFSStatReply) {
	for i := range _code {
		pack = append(pack, _packSerializeProblemFSStatReply(_code[i], _data[i]))
	}
	return
}
func PSerializeProblemFSStatInnerReply(_name string, _size int64, _isDir bool, _modTime time.Time) *ProblemFSStatInnerReply {

	return &ProblemFSStatInnerReply{
		Name:    _name,
		Size:    _size,
		IsDir:   _isDir,
		ModTime: _modTime,
	}
}
func SerializeProblemFSStatInnerReply(_name string, _size int64, _isDir bool, _modTime time.Time) ProblemFSStatInnerReply {

	return ProblemFSStatInnerReply{
		Name:    _name,
		Size:    _size,
		IsDir:   _isDir,
		ModTime: _modTime,
	}
}
func _packSerializeProblemFSStatInnerReply(_name string, _size int64, _isDir bool, _modTime time.Time) ProblemFSStatInnerReply {

	return ProblemFSStatInnerReply{
		Name:    _name,
		Size:    _size,
		IsDir:   _isDir,
		ModTime: _modTime,
	}
}
func PackSerializeProblemFSStatInnerReply(_name []string, _size []int64, _isDir []bool, _modTime []time.Time) (pack []ProblemFSStatInnerReply) {
	for i := range _name {
		pack = append(pack, _packSerializeProblemFSStatInnerReply(_name[i], _size[i], _isDir[i], _modTime[i]))
	}
	return
}
func PSerializeProblemFSWriteRequest(_path string) *ProblemFSWriteRequest {

	return &ProblemFSWriteRequest{
		Path: _path,
	}
}
func SerializeProblemFSWriteRequest(_path string) ProblemFSWriteRequest {

	return ProblemFSWriteRequest{
		Path: _path,
	}
}
func _packSerializeProblemFSWriteRequest(_path string) ProblemFSWriteRequest {

	return ProblemFSWriteRequest{
		Path: _path,
	}
}
func PackSerializeProblemFSWriteRequest(_path []string) (pack []ProblemFSWriteRequest) {
	for i := range _path {
		pack = append(pack, _packSerializeProblemFSWriteRequest(_path[i]))
	}
	return
}
func PSerializeProblemFSWriteReply(_code int, _data ProblemFSWriteInnerReply) *ProblemFSWriteReply {

	return &ProblemFSWriteReply{
		Code: _code,
		Data: _data,
	}
}
func SerializeProblemFSWriteReply(_code int, _data ProblemFSWriteInnerReply) ProblemFSWriteReply {

	return ProblemFSWriteReply{
		Code: _code,
		Data: _data,
	}
}
func _packSerializeProblemFSWriteReply(_code int, _data ProblemFSWriteInnerReply) ProblemFSWriteReply {

	return ProblemFSWriteReply{
		Code: _code,
		Data: _data,
	}
}
func PackSerializeProblemFSWriteReply(_code []int, _data []ProblemFSWriteInnerReply) (pack []ProblemFSWriteReply) {
	for i := range _code {
		pack = append(pack, _packSerializeProblemFSWriteReply(_code[i], _data[i]))
	}
	return
}
func PSerializeProblemFSWriteInnerReply() *ProblemFSWriteInnerReply {

	return &ProblemFSWriteInnerReply{}
}
func SerializeProblemFSWriteInnerReply() ProblemFSWriteInnerReply {

	return ProblemFSWriteInnerReply{}
}
func _packSerializeProblemFSWriteInnerReply() ProblemFSWriteInnerReply {

	return ProblemFSWriteInnerReply{}
}
func PackSerializeProblemFSWriteInnerReply() (pack []ProblemFSWriteInnerReply) {
	return
}
func PSerializeProblemFSRemoveRequest(_path string) *ProblemFSRemoveRequest {

	return &ProblemFSRemoveRequest{
		Path: _path,
	}
}
func SerializeProblemFSRemoveRequest(_path string) ProblemFSRemoveRequest {

	return ProblemFSRemoveRequest{
		Path: _path,
	}
}
func _packSerializeProblemFSRemoveRequest(_path string) ProblemFSRemoveRequest {

	return ProblemFSRemoveRequest{
		Path: _path,
	}
}
func PackSerializeProblemFSRemoveRequest(_path []string) (pack []ProblemFSRemoveRequest) {
	for i := range _path {
		pack = append(pack, _packSerializeProblemFSRemoveRequest(_path[i]))
	}
	return
}
func PSerializeProblemFSRemoveReply(_code int, _data ProblemFSRemoveInnerReply) *ProblemFSRemoveReply {

	return &ProblemFSRemoveReply{
		Code: _code,
		Data: _data,
	}
}
func SerializeProblemFSRemoveReply(_code int, _data ProblemFSRemoveInnerReply) ProblemFSRemoveReply {

	return ProblemFSRemoveReply{
		Code: _code,
		Data: _data,
	}
}
func _packSerializeProblemFSRemoveReply(_code int, _data ProblemFSRemoveInnerReply) ProblemFSRemoveReply {

	return ProblemFSRemoveReply{
		Code: _code,
		Data: _data,
	}
}
func PackSerializeProblemFSRemoveReply(_code []int, _data []ProblemFSRemoveInnerReply) (pack []ProblemFSRemoveReply) {
	for i := range _code {
		pack = append(pack, _packSerializeProblemFSRemoveReply(_code[i], _data[i]))
	}
	return
}
func PSerializeProblemFSRemoveInnerReply() *ProblemFSRemoveInnerReply {

	return &ProblemFSRemoveInnerReply{}
}
func SerializeProblemFSRemoveInnerReply() ProblemFSRemoveInnerReply {

	return ProblemFSRemoveInnerReply{}
}
func _packSerializeProblemFSRemoveInnerReply() ProblemFSRemoveInnerReply {

	return ProblemFSRemoveInnerReply{}
}
func PackSerializeProblemFSRemoveInnerReply() (pack []ProblemFSRemoveInnerReply) {
	return
}
func PSerializeProblemFSZipReadRequest(_path string) *ProblemFSZipReadRequest {

	return &ProblemFSZipReadRequest{
		Path: _path,
	}
}
func SerializeProblemFSZipReadRequest(_path string) ProblemFSZipReadRequest {

	return ProblemFSZipReadRequest{
		Path: _path,
	}
}
func _packSerializeProblemFSZipReadRequest(_path string) ProblemFSZipReadRequest {

	return ProblemFSZipReadRequest{
		Path: _path,
	}
}
func PackSerializeProblemFSZipReadRequest(_path []string) (pack []ProblemFSZipReadRequest) {
	for i := range _path {
		pack = append(pack, _packSerializeProblemFSZipReadRequest(_path[i]))
	}
	return
}
func PSerializeProblemFSZipReadReply(_code int, _data ProblemFSZipReadInnerReply) *ProblemFSZipReadReply {

	return &ProblemFSZipReadReply{
		Code: _code,
		Data: _data,
	}
}
func SerializeProblemFSZipReadReply(_code int, _data ProblemFSZipReadInnerReply) ProblemFSZipReadReply {

	return ProblemFSZipReadReply{
		Code: _code,
		Data: _data,
	}
}
func _packSerializeProblemFSZipReadReply(_code int, _data ProblemFSZipReadInnerReply) ProblemFSZipReadReply {

	return ProblemFSZipReadReply{
		Code: _code,
		Data: _data,
	}
}
func PackSerializeProblemFSZipReadReply(_code []int, _data []ProblemFSZipReadInnerReply) (pack []ProblemFSZipReadReply) {
	for i := range _code {
		pack = append(pack, _packSerializeProblemFSZipReadReply(_code[i], _data[i]))
	}
	return
}
func PSerializeProblemFSZipReadInnerReply() *ProblemFSZipReadInnerReply {

	return &ProblemFSZipReadInnerReply{}
}
func SerializeProblemFSZipReadInnerReply() ProblemFSZipReadInnerReply {

	return ProblemFSZipReadInnerReply{}
}
func _packSerializeProblemFSZipReadInnerReply() ProblemFSZipReadInnerReply {

	return ProblemFSZipReadInnerReply{}
}
func PackSerializeProblemFSZipReadInnerReply() (pack []ProblemFSZipReadInnerReply) {
	return
}
func PSerializeProblemFSLSRequest(_path string) *ProblemFSLSRequest {

	return &ProblemFSLSRequest{
		Path: _path,
	}
}
func SerializeProblemFSLSRequest(_path string) ProblemFSLSRequest {

	return ProblemFSLSRequest{
		Path: _path,
	}
}
func _packSerializeProblemFSLSRequest(_path string) ProblemFSLSRequest {

	return ProblemFSLSRequest{
		Path: _path,
	}
}
func PackSerializeProblemFSLSRequest(_path []string) (pack []ProblemFSLSRequest) {
	for i := range _path {
		pack = append(pack, _packSerializeProblemFSLSRequest(_path[i]))
	}
	return
}
func PSerializeProblemFSLSReply(_code int, _data []ProblemFSLSInnerReply) *ProblemFSLSReply {

	return &ProblemFSLSReply{
		Code: _code,
		Data: _data,
	}
}
func SerializeProblemFSLSReply(_code int, _data []ProblemFSLSInnerReply) ProblemFSLSReply {

	return ProblemFSLSReply{
		Code: _code,
		Data: _data,
	}
}
func _packSerializeProblemFSLSReply(_code int, _data []ProblemFSLSInnerReply) ProblemFSLSReply {

	return ProblemFSLSReply{
		Code: _code,
		Data: _data,
	}
}
func PackSerializeProblemFSLSReply(_code []int, _data [][]ProblemFSLSInnerReply) (pack []ProblemFSLSReply) {
	for i := range _code {
		pack = append(pack, _packSerializeProblemFSLSReply(_code[i], _data[i]))
	}
	return
}
func PSerializeProblemFSLSInnerReply(_name string, _size int64, _isDir bool, _modTime time.Time) *ProblemFSLSInnerReply {

	return &ProblemFSLSInnerReply{
		Name:    _name,
		Size:    _size,
		IsDir:   _isDir,
		ModTime: _modTime,
	}
}
func SerializeProblemFSLSInnerReply(_name string, _size int64, _isDir bool, _modTime time.Time) ProblemFSLSInnerReply {

	return ProblemFSLSInnerReply{
		Name:    _name,
		Size:    _size,
		IsDir:   _isDir,
		ModTime: _modTime,
	}
}
func _packSerializeProblemFSLSInnerReply(_name string, _size int64, _isDir bool, _modTime time.Time) ProblemFSLSInnerReply {

	return ProblemFSLSInnerReply{
		Name:    _name,
		Size:    _size,
		IsDir:   _isDir,
		ModTime: _modTime,
	}
}
func PackSerializeProblemFSLSInnerReply(_name []string, _size []int64, _isDir []bool, _modTime []time.Time) (pack []ProblemFSLSInnerReply) {
	for i := range _name {
		pack = append(pack, _packSerializeProblemFSLSInnerReply(_name[i], _size[i], _isDir[i], _modTime[i]))
	}
	return
}
func PSerializeProblemFSWritesRequest(_path string) *ProblemFSWritesRequest {

	return &ProblemFSWritesRequest{
		Path: _path,
	}
}
func SerializeProblemFSWritesRequest(_path string) ProblemFSWritesRequest {

	return ProblemFSWritesRequest{
		Path: _path,
	}
}
func _packSerializeProblemFSWritesRequest(_path string) ProblemFSWritesRequest {

	return ProblemFSWritesRequest{
		Path: _path,
	}
}
func PackSerializeProblemFSWritesRequest(_path []string) (pack []ProblemFSWritesRequest) {
	for i := range _path {
		pack = append(pack, _packSerializeProblemFSWritesRequest(_path[i]))
	}
	return
}
func PSerializeProblemFSWritesReply(_code int, _data ProblemFSWritesInnerReply) *ProblemFSWritesReply {

	return &ProblemFSWritesReply{
		Code: _code,
		Data: _data,
	}
}
func SerializeProblemFSWritesReply(_code int, _data ProblemFSWritesInnerReply) ProblemFSWritesReply {

	return ProblemFSWritesReply{
		Code: _code,
		Data: _data,
	}
}
func _packSerializeProblemFSWritesReply(_code int, _data ProblemFSWritesInnerReply) ProblemFSWritesReply {

	return ProblemFSWritesReply{
		Code: _code,
		Data: _data,
	}
}
func PackSerializeProblemFSWritesReply(_code []int, _data []ProblemFSWritesInnerReply) (pack []ProblemFSWritesReply) {
	for i := range _code {
		pack = append(pack, _packSerializeProblemFSWritesReply(_code[i], _data[i]))
	}
	return
}
func PSerializeProblemFSWritesInnerReply() *ProblemFSWritesInnerReply {

	return &ProblemFSWritesInnerReply{}
}
func SerializeProblemFSWritesInnerReply() ProblemFSWritesInnerReply {

	return ProblemFSWritesInnerReply{}
}
func _packSerializeProblemFSWritesInnerReply() ProblemFSWritesInnerReply {

	return ProblemFSWritesInnerReply{}
}
func PackSerializeProblemFSWritesInnerReply() (pack []ProblemFSWritesInnerReply) {
	return
}
func PSerializeProblemFSMkdirRequest(_path string) *ProblemFSMkdirRequest {

	return &ProblemFSMkdirRequest{
		Path: _path,
	}
}
func SerializeProblemFSMkdirRequest(_path string) ProblemFSMkdirRequest {

	return ProblemFSMkdirRequest{
		Path: _path,
	}
}
func _packSerializeProblemFSMkdirRequest(_path string) ProblemFSMkdirRequest {

	return ProblemFSMkdirRequest{
		Path: _path,
	}
}
func PackSerializeProblemFSMkdirRequest(_path []string) (pack []ProblemFSMkdirRequest) {
	for i := range _path {
		pack = append(pack, _packSerializeProblemFSMkdirRequest(_path[i]))
	}
	return
}
func PSerializeProblemFSMkdirReply(_code int, _data ProblemFSMkdirInnerReply) *ProblemFSMkdirReply {

	return &ProblemFSMkdirReply{
		Code: _code,
		Data: _data,
	}
}
func SerializeProblemFSMkdirReply(_code int, _data ProblemFSMkdirInnerReply) ProblemFSMkdirReply {

	return ProblemFSMkdirReply{
		Code: _code,
		Data: _data,
	}
}
func _packSerializeProblemFSMkdirReply(_code int, _data ProblemFSMkdirInnerReply) ProblemFSMkdirReply {

	return ProblemFSMkdirReply{
		Code: _code,
		Data: _data,
	}
}
func PackSerializeProblemFSMkdirReply(_code []int, _data []ProblemFSMkdirInnerReply) (pack []ProblemFSMkdirReply) {
	for i := range _code {
		pack = append(pack, _packSerializeProblemFSMkdirReply(_code[i], _data[i]))
	}
	return
}
func PSerializeProblemFSMkdirInnerReply() *ProblemFSMkdirInnerReply {

	return &ProblemFSMkdirInnerReply{}
}
func SerializeProblemFSMkdirInnerReply() ProblemFSMkdirInnerReply {

	return ProblemFSMkdirInnerReply{}
}
func _packSerializeProblemFSMkdirInnerReply() ProblemFSMkdirInnerReply {

	return ProblemFSMkdirInnerReply{}
}
func PackSerializeProblemFSMkdirInnerReply() (pack []ProblemFSMkdirInnerReply) {
	return
}
func PSerializeProblemFSRemoveAllRequest(_path string) *ProblemFSRemoveAllRequest {

	return &ProblemFSRemoveAllRequest{
		Path: _path,
	}
}
func SerializeProblemFSRemoveAllRequest(_path string) ProblemFSRemoveAllRequest {

	return ProblemFSRemoveAllRequest{
		Path: _path,
	}
}
func _packSerializeProblemFSRemoveAllRequest(_path string) ProblemFSRemoveAllRequest {

	return ProblemFSRemoveAllRequest{
		Path: _path,
	}
}
func PackSerializeProblemFSRemoveAllRequest(_path []string) (pack []ProblemFSRemoveAllRequest) {
	for i := range _path {
		pack = append(pack, _packSerializeProblemFSRemoveAllRequest(_path[i]))
	}
	return
}
func PSerializeProblemFSRemoveAllReply(_code int, _data ProblemFSRemoveAllInnerReply) *ProblemFSRemoveAllReply {

	return &ProblemFSRemoveAllReply{
		Code: _code,
		Data: _data,
	}
}
func SerializeProblemFSRemoveAllReply(_code int, _data ProblemFSRemoveAllInnerReply) ProblemFSRemoveAllReply {

	return ProblemFSRemoveAllReply{
		Code: _code,
		Data: _data,
	}
}
func _packSerializeProblemFSRemoveAllReply(_code int, _data ProblemFSRemoveAllInnerReply) ProblemFSRemoveAllReply {

	return ProblemFSRemoveAllReply{
		Code: _code,
		Data: _data,
	}
}
func PackSerializeProblemFSRemoveAllReply(_code []int, _data []ProblemFSRemoveAllInnerReply) (pack []ProblemFSRemoveAllReply) {
	for i := range _code {
		pack = append(pack, _packSerializeProblemFSRemoveAllReply(_code[i], _data[i]))
	}
	return
}
func PSerializeProblemFSRemoveAllInnerReply() *ProblemFSRemoveAllInnerReply {

	return &ProblemFSRemoveAllInnerReply{}
}
func SerializeProblemFSRemoveAllInnerReply() ProblemFSRemoveAllInnerReply {

	return ProblemFSRemoveAllInnerReply{}
}
func _packSerializeProblemFSRemoveAllInnerReply() ProblemFSRemoveAllInnerReply {

	return ProblemFSRemoveAllInnerReply{}
}
func PackSerializeProblemFSRemoveAllInnerReply() (pack []ProblemFSRemoveAllInnerReply) {
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
		DescriptionRef: problem.DescriptionRef,
	}
}
func SerializePutProblemRequest(problem *problem.Problem) PutProblemRequest {

	return PutProblemRequest{
		Title:          problem.Title,
		DescriptionRef: problem.DescriptionRef,
	}
}
func _packSerializePutProblemRequest(problem *problem.Problem) PutProblemRequest {

	return PutProblemRequest{
		Title:          problem.Title,
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
