
package control

import (
    "github.com/Myriad-Dreamin/minimum-lib/controller"

)

type SubmissionService interface {
    SubmissionServiceSignatureXXX() interface{}
    ListSubmissions(c controller.MContext)
    CountSubmissions(c controller.MContext)
    PostSubmission(c controller.MContext)
    GetContent(c controller.MContext)
    GetSubmission(c controller.MContext)
    Delete(c controller.MContext)

}