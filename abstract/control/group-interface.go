
package control

import (
    "github.com/Myriad-Dreamin/minimum-lib/controller"

)

type GroupController interface {
    GroupControllerSignatureXXX() interface{}
    ListGroup(c controller.MContext)
    CountGroup(c controller.MContext)
    PostGroup(c controller.MContext)
    GetGroupMembers(c controller.MContext)
    PostGroupMember(c controller.MContext)
    PutGroupOwner(c controller.MContext)
    GetGroup(c controller.MContext)
    PutGroup(c controller.MContext)
    DeleteGroup(c controller.MContext)

}