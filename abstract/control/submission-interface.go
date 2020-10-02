
package control

import (
    "github.com/Myriad-Dreamin/minimum-lib/controller"

)

type SubmissionController interface {
    SubmissionControllerSignatureXXX() interface{}
    ListSubmission(c controller.MContext)
    CountSubmission(c controller.MContext)
    PostSubmission(c controller.MContext)
    GetSubmissionContent(c controller.MContext)
    GetSubmission(c controller.MContext)
    DeleteSubmission(c controller.MContext)

}