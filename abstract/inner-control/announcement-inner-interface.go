package inner_control

import (
	"github.com/Myriad-Dreamin/boj-v6/api"
)

type InnerAnnouncementService interface {
	AnnouncementServiceSignatureXXX() interface{}
	ListAnnouncement(req *api.ListAnnouncementRequest) (*api.ListAnnouncementReply, error)
	CountAnnouncement(req *api.CountAnnouncementRequest) (*api.CountAnnouncementReply, error)
	PostAnnouncement(req *api.PostAnnouncementRequest) (*api.PostAnnouncementReply, error)
	GetAnnouncement(req *api.GetAnnouncementRequest) (*api.GetAnnouncementReply, error)
	PutAnnouncement(req *api.PutAnnouncementRequest) (*api.PutAnnouncementReply, error)
	DeleteAnnouncement(req *api.DeleteAnnouncementRequest) (*api.DeleteAnnouncementReply, error)
}
