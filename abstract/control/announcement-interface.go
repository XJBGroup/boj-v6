package control

import (
	"github.com/Myriad-Dreamin/minimum-lib/controller"
)

type AnnouncementController interface {
	AnnouncementControllerSignatureXXX() interface{}
	ListAnnouncement(c controller.MContext)
	CountAnnouncement(c controller.MContext)
	PostAnnouncement(c controller.MContext)
	GetAnnouncement(c controller.MContext)
	PutAnnouncement(c controller.MContext)
	DeleteAnnouncement(c controller.MContext)
}
