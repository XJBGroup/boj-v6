package api

import (
	"github.com/Myriad-Dreamin/boj-v6/abstract/announcement"
	"github.com/Myriad-Dreamin/go-model-traits/gorm-crud-dao"
)

type ListAnnouncementsRequest = gorm_crud_dao.Filter

type ListAnnouncementsReply struct {
	Code int                         `json:"code" form:"code"`
	Data []announcement.Announcement `form:"data" json:"data"`
}

type CountAnnouncementsRequest = gorm_crud_dao.Filter

type CountAnnouncementReply struct {
	Code int   `json:"code" form:"code"`
	Data int64 `form:"data" json:"data"`
}

type PostAnnouncementRequest struct {
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
}

type PostAnnouncementReply struct {
	Code         int                        `json:"code" form:"code"`
	Announcement *announcement.Announcement `json:"announcement" form:"announcement"`
}

type GetAnnouncementReply struct {
	Code         int                        `json:"code" form:"code"`
	Announcement *announcement.Announcement `json:"announcement" form:"announcement"`
}

type PutAnnouncementRequest struct {
	Title   string `json:"title" form:"title"`
	Content string `form:"content" json:"content"`
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
func PSerializePostAnnouncementReply(_code int, _announcement *announcement.Announcement) *PostAnnouncementReply {

	return &PostAnnouncementReply{
		Code:         _code,
		Announcement: _announcement,
	}
}
func SerializePostAnnouncementReply(_code int, _announcement *announcement.Announcement) PostAnnouncementReply {

	return PostAnnouncementReply{
		Code:         _code,
		Announcement: _announcement,
	}
}
func _packSerializePostAnnouncementReply(_code int, _announcement *announcement.Announcement) PostAnnouncementReply {

	return PostAnnouncementReply{
		Code:         _code,
		Announcement: _announcement,
	}
}
func PackSerializePostAnnouncementReply(_code []int, _announcement []*announcement.Announcement) (pack []PostAnnouncementReply) {
	for i := range _code {
		pack = append(pack, _packSerializePostAnnouncementReply(_code[i], _announcement[i]))
	}
	return
}
func PSerializeGetAnnouncementReply(_code int, _announcement *announcement.Announcement) *GetAnnouncementReply {

	return &GetAnnouncementReply{
		Code:         _code,
		Announcement: _announcement,
	}
}
func SerializeGetAnnouncementReply(_code int, _announcement *announcement.Announcement) GetAnnouncementReply {

	return GetAnnouncementReply{
		Code:         _code,
		Announcement: _announcement,
	}
}
func _packSerializeGetAnnouncementReply(_code int, _announcement *announcement.Announcement) GetAnnouncementReply {

	return GetAnnouncementReply{
		Code:         _code,
		Announcement: _announcement,
	}
}
func PackSerializeGetAnnouncementReply(_code []int, _announcement []*announcement.Announcement) (pack []GetAnnouncementReply) {
	for i := range _code {
		pack = append(pack, _packSerializeGetAnnouncementReply(_code[i], _announcement[i]))
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
