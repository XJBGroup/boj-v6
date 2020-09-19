
package control

import (
    "github.com/Myriad-Dreamin/minimum-lib/controller"

)

type CommentService interface {
    CommentServiceSignatureXXX() interface{}
    ListComments(c controller.MContext)
    CountComment(c controller.MContext)
    PostComment(c controller.MContext)
    GetComment(c controller.MContext)
    PutComment(c controller.MContext)
    Delete(c controller.MContext)

}