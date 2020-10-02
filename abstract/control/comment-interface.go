
package control

import (
    "github.com/Myriad-Dreamin/minimum-lib/controller"

)

type CommentController interface {
    CommentControllerSignatureXXX() interface{}
    ListComment(c controller.MContext)
    CountComment(c controller.MContext)
    PostComment(c controller.MContext)
    GetComment(c controller.MContext)
    PutComment(c controller.MContext)
    DeleteComment(c controller.MContext)

}