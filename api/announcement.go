package api

import (
	"github.com/Myriad-Dreamin/boj-v6/abstract/announcement"
)

type AnnouncementFilter struct {
	Page     int `json:"page" form:"page"`
	PageSize int `json:"page_size" form:"page_size"`
}

type ListAnnouncementsReply struct {
	Code int                         `form:"code" json:"code"`
	Data []announcement.Announcement `json:"data" form:"data"`
}

type CountAnnouncementReply struct {
	Code int   `form:"code" json:"code"`
	Data int64 `json:"data" form:"data"`
}

type PostAnnouncementRequest struct {
	Title   string `json:"title" form:"title" binding:"required"`
	Content string `json:"content" form:"content" binding:"required"`
}

type PostAnnouncementReply struct {
	Code int                        `form:"code" json:"code"`
	Data *announcement.Announcement `json:"data" form:"data"`
}

type GetAnnouncementReply struct {
	Code int                        `json:"code" form:"code"`
	Data *announcement.Announcement `json:"data" form:"data"`
}

type PutAnnouncementRequest struct {
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
}

func PSerializeAnnouncementFilter(_page int, _pageSize int) *AnnouncementFilter {

	return &AnnouncementFilter{
		Page:     _page,
		PageSize: _pageSize,
	}
}
func SerializeAnnouncementFilter(_page int, _pageSize int) AnnouncementFilter {

	return AnnouncementFilter{
		Page:     _page,
		PageSize: _pageSize,
	}
}
func _packSerializeAnnouncementFilter(_page int, _pageSize int) AnnouncementFilter {

	return AnnouncementFilter{
		Page:     _page,
		PageSize: _pageSize,
	}
}
func PackSerializeAnnouncementFilter(_page []int, _pageSize []int) (pack []AnnouncementFilter) {
	for i := range _page {
		pack = append(pack, _packSerializeAnnouncementFilter(_page[i], _pageSize[i]))
	}
	return
}
func PSerializeListAnnouncementsReply(_code int, _data []announcement.Announcement) *ListAnnouncementsReply {

	return &ListAnnouncementsReply{
		Code: _code,
		Data: _data,
	}
}
func SerializeListAnnouncementsReply(_code int, _data []announcement.Announcement) ListAnnouncementsReply {

	return ListAnnouncementsReply{
		Code: _code,
		Data: _data,
	}
}
func _packSerializeListAnnouncementsReply(_code int, _data []announcement.Announcement) ListAnnouncementsReply {

	return ListAnnouncementsReply{
		Code: _code,
		Data: _data,
	}
}
func PackSerializeListAnnouncementsReply(_code []int, _data [][]announcement.Announcement) (pack []ListAnnouncementsReply) {
	for i := range _code {
		pack = append(pack, _packSerializeListAnnouncementsReply(_code[i], _data[i]))
	}
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
