package api

import (
	"github.com/Myriad-Dreamin/boj-v6/abstract/announcement"
)

type ListAnnouncementRequest struct {
	Page     int `json:"page" form:"page"`
	PageSize int `json:"page_size" form:"page_size"`
}

type ListAnnouncementReply struct {
	Code int                         `json:"code" form:"code"`
	Data []announcement.Announcement `json:"data" form:"data"`
}

type CountAnnouncementRequest struct {
}

type CountAnnouncementReply struct {
	Code int   `json:"code" form:"code"`
	Data int64 `json:"data" form:"data"`
}

type PostAnnouncementRequest struct {
	Title   string `json:"title" form:"title" binding:"required"`
	Content string `binding:"required" json:"content" form:"content"`
}

type PostAnnouncementReply struct {
	Code int                        `json:"code" form:"code"`
	Data *announcement.Announcement `form:"data" json:"data"`
}

type GetAnnouncementRequest struct {
}

type GetAnnouncementReply struct {
	Code int                        `json:"code" form:"code"`
	Data *announcement.Announcement `json:"data" form:"data"`
}

type PutAnnouncementRequest struct {
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
}

type PutAnnouncementReply struct {
	Code int `form:"code" json:"code"`
}

type DeleteAnnouncementRequest struct {
}

type DeleteAnnouncementReply struct {
	Code int `json:"code" form:"code"`
}

func PSerializeListAnnouncementRequest(_page int, _pageSize int) *ListAnnouncementRequest {

	return &ListAnnouncementRequest{
		Page:     _page,
		PageSize: _pageSize,
	}
}
func SerializeListAnnouncementRequest(_page int, _pageSize int) ListAnnouncementRequest {

	return ListAnnouncementRequest{
		Page:     _page,
		PageSize: _pageSize,
	}
}
func _packSerializeListAnnouncementRequest(_page int, _pageSize int) ListAnnouncementRequest {

	return ListAnnouncementRequest{
		Page:     _page,
		PageSize: _pageSize,
	}
}
func PackSerializeListAnnouncementRequest(_page []int, _pageSize []int) (pack []ListAnnouncementRequest) {
	for i := range _page {
		pack = append(pack, _packSerializeListAnnouncementRequest(_page[i], _pageSize[i]))
	}
	return
}
func PSerializeListAnnouncementReply(_code int, _data []announcement.Announcement) *ListAnnouncementReply {

	return &ListAnnouncementReply{
		Code: _code,
		Data: _data,
	}
}
func SerializeListAnnouncementReply(_code int, _data []announcement.Announcement) ListAnnouncementReply {

	return ListAnnouncementReply{
		Code: _code,
		Data: _data,
	}
}
func _packSerializeListAnnouncementReply(_code int, _data []announcement.Announcement) ListAnnouncementReply {

	return ListAnnouncementReply{
		Code: _code,
		Data: _data,
	}
}
func PackSerializeListAnnouncementReply(_code []int, _data [][]announcement.Announcement) (pack []ListAnnouncementReply) {
	for i := range _code {
		pack = append(pack, _packSerializeListAnnouncementReply(_code[i], _data[i]))
	}
	return
}
func PSerializeCountAnnouncementRequest() *CountAnnouncementRequest {

	return &CountAnnouncementRequest{}
}
func SerializeCountAnnouncementRequest() CountAnnouncementRequest {

	return CountAnnouncementRequest{}
}
func _packSerializeCountAnnouncementRequest() CountAnnouncementRequest {

	return CountAnnouncementRequest{}
}
func PackSerializeCountAnnouncementRequest() (pack []CountAnnouncementRequest) {
	return
}
func PSerializeCountAnnouncementReply(_code int, _data int64) *CountAnnouncementReply {

	return &CountAnnouncementReply{
		Code: _code,
		Data: _data,
	}
}
func SerializeCountAnnouncementReply(_code int, _data int64) CountAnnouncementReply {

	return CountAnnouncementReply{
		Code: _code,
		Data: _data,
	}
}
func _packSerializeCountAnnouncementReply(_code int, _data int64) CountAnnouncementReply {

	return CountAnnouncementReply{
		Code: _code,
		Data: _data,
	}
}
func PackSerializeCountAnnouncementReply(_code []int, _data []int64) (pack []CountAnnouncementReply) {
	for i := range _code {
		pack = append(pack, _packSerializeCountAnnouncementReply(_code[i], _data[i]))
	}
	return
}
func PSerializePostAnnouncementRequest(announcement *announcement.Announcement) *PostAnnouncementRequest {

	return &PostAnnouncementRequest{
		Title:   announcement.Title,
		Content: announcement.Content,
	}
}
func SerializePostAnnouncementRequest(announcement *announcement.Announcement) PostAnnouncementRequest {

	return PostAnnouncementRequest{
		Title:   announcement.Title,
		Content: announcement.Content,
	}
}
func _packSerializePostAnnouncementRequest(announcement *announcement.Announcement) PostAnnouncementRequest {

	return PostAnnouncementRequest{
		Title:   announcement.Title,
		Content: announcement.Content,
	}
}
func PackSerializePostAnnouncementRequest(announcement []*announcement.Announcement) (pack []PostAnnouncementRequest) {
	for i := range announcement {
		pack = append(pack, _packSerializePostAnnouncementRequest(announcement[i]))
	}
	return
}
func PSerializePostAnnouncementReply(_code int, _data *announcement.Announcement) *PostAnnouncementReply {

	return &PostAnnouncementReply{
		Code: _code,
		Data: _data,
	}
}
func SerializePostAnnouncementReply(_code int, _data *announcement.Announcement) PostAnnouncementReply {

	return PostAnnouncementReply{
		Code: _code,
		Data: _data,
	}
}
func _packSerializePostAnnouncementReply(_code int, _data *announcement.Announcement) PostAnnouncementReply {

	return PostAnnouncementReply{
		Code: _code,
		Data: _data,
	}
}
func PackSerializePostAnnouncementReply(_code []int, _data []*announcement.Announcement) (pack []PostAnnouncementReply) {
	for i := range _code {
		pack = append(pack, _packSerializePostAnnouncementReply(_code[i], _data[i]))
	}
	return
}
func PSerializeGetAnnouncementRequest() *GetAnnouncementRequest {

	return &GetAnnouncementRequest{}
}
func SerializeGetAnnouncementRequest() GetAnnouncementRequest {

	return GetAnnouncementRequest{}
}
func _packSerializeGetAnnouncementRequest() GetAnnouncementRequest {

	return GetAnnouncementRequest{}
}
func PackSerializeGetAnnouncementRequest() (pack []GetAnnouncementRequest) {
	return
}
func PSerializeGetAnnouncementReply(_code int, _data *announcement.Announcement) *GetAnnouncementReply {

	return &GetAnnouncementReply{
		Code: _code,
		Data: _data,
	}
}
func SerializeGetAnnouncementReply(_code int, _data *announcement.Announcement) GetAnnouncementReply {

	return GetAnnouncementReply{
		Code: _code,
		Data: _data,
	}
}
func _packSerializeGetAnnouncementReply(_code int, _data *announcement.Announcement) GetAnnouncementReply {

	return GetAnnouncementReply{
		Code: _code,
		Data: _data,
	}
}
func PackSerializeGetAnnouncementReply(_code []int, _data []*announcement.Announcement) (pack []GetAnnouncementReply) {
	for i := range _code {
		pack = append(pack, _packSerializeGetAnnouncementReply(_code[i], _data[i]))
	}
	return
}
func PSerializePutAnnouncementRequest(announcement *announcement.Announcement) *PutAnnouncementRequest {

	return &PutAnnouncementRequest{
		Title:   announcement.Title,
		Content: announcement.Content,
	}
}
func SerializePutAnnouncementRequest(announcement *announcement.Announcement) PutAnnouncementRequest {

	return PutAnnouncementRequest{
		Title:   announcement.Title,
		Content: announcement.Content,
	}
}
func _packSerializePutAnnouncementRequest(announcement *announcement.Announcement) PutAnnouncementRequest {

	return PutAnnouncementRequest{
		Title:   announcement.Title,
		Content: announcement.Content,
	}
}
func PackSerializePutAnnouncementRequest(announcement []*announcement.Announcement) (pack []PutAnnouncementRequest) {
	for i := range announcement {
		pack = append(pack, _packSerializePutAnnouncementRequest(announcement[i]))
	}
	return
}
func PSerializePutAnnouncementReply(_code int) *PutAnnouncementReply {

	return &PutAnnouncementReply{
		Code: _code,
	}
}
func SerializePutAnnouncementReply(_code int) PutAnnouncementReply {

	return PutAnnouncementReply{
		Code: _code,
	}
}
func _packSerializePutAnnouncementReply(_code int) PutAnnouncementReply {

	return PutAnnouncementReply{
		Code: _code,
	}
}
func PackSerializePutAnnouncementReply(_code []int) (pack []PutAnnouncementReply) {
	for i := range _code {
		pack = append(pack, _packSerializePutAnnouncementReply(_code[i]))
	}
	return
}
func PSerializeDeleteAnnouncementRequest() *DeleteAnnouncementRequest {

	return &DeleteAnnouncementRequest{}
}
func SerializeDeleteAnnouncementRequest() DeleteAnnouncementRequest {

	return DeleteAnnouncementRequest{}
}
func _packSerializeDeleteAnnouncementRequest() DeleteAnnouncementRequest {

	return DeleteAnnouncementRequest{}
}
func PackSerializeDeleteAnnouncementRequest() (pack []DeleteAnnouncementRequest) {
	return
}
func PSerializeDeleteAnnouncementReply(_code int) *DeleteAnnouncementReply {

	return &DeleteAnnouncementReply{
		Code: _code,
	}
}
func SerializeDeleteAnnouncementReply(_code int) DeleteAnnouncementReply {

	return DeleteAnnouncementReply{
		Code: _code,
	}
}
func _packSerializeDeleteAnnouncementReply(_code int) DeleteAnnouncementReply {

	return DeleteAnnouncementReply{
		Code: _code,
	}
}
func PackSerializeDeleteAnnouncementReply(_code []int) (pack []DeleteAnnouncementReply) {
	for i := range _code {
		pack = append(pack, _packSerializeDeleteAnnouncementReply(_code[i]))
	}
	return
}
