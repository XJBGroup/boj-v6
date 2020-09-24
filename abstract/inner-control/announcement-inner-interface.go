package inner_control

import (
	"github.com/Myriad-Dreamin/boj-v6/api"
	"github.com/Myriad-Dreamin/minimum-lib/controller"
)

type InnerAnnouncementService interface {
	AnnouncementServiceSignatureXXX() interface{}
	ListAnnouncement(c controller.MContext, req *api.ListAnnouncementRequest) (*api.ListAnnouncementReply, error)
	CountAnnouncement(c controller.MContext, req *api.CountAnnouncementRequest) (*api.CountAnnouncementReply, error)
	PostAnnouncement(c controller.MContext, req *api.PostAnnouncementRequest) (*api.PostAnnouncementReply, error)
	GetAnnouncement(c controller.MContext, req *api.GetAnnouncementRequest) (*api.GetAnnouncementReply, error)
	PutAnnouncement(c controller.MContext, req *api.PutAnnouncementRequest) (*api.PutAnnouncementReply, error)
	DeleteAnnouncement(c controller.MContext, req *api.DeleteAnnouncementRequest) (*api.DeleteAnnouncementReply, error)
}
