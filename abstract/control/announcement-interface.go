
package control

import (
    "github.com/Myriad-Dreamin/minimum-lib/controller"

)

type AnnouncementService interface {
    AnnouncementServiceSignatureXXX() interface{}
    ListAnnouncements(c controller.MContext)
    CountAnnouncement(c controller.MContext)
    PostAnnouncement(c controller.MContext)
    GetAnnouncement(c controller.MContext)
    PutAnnouncement(c controller.MContext)
    Delete(c controller.MContext)

}