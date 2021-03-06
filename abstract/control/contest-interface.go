
package control

import (
    "github.com/Myriad-Dreamin/minimum-lib/controller"

)

type ContestController interface {
    ContestControllerSignatureXXX() interface{}
    ListContest(c controller.MContext)
    CountContest(c controller.MContext)
    PostContest(c controller.MContext)
    CountContestProblem(c controller.MContext)
    PostContestProblem(c controller.MContext)
    CountContestProblemDesc(c controller.MContext)
    ChangeContestProblemDescriptionRef(c controller.MContext)
    PostContestProblemDesc(c controller.MContext)
    GetContestProblemDesc(c controller.MContext)
    PutContestProblemDesc(c controller.MContext)
    DeleteContestProblemDesc(c controller.MContext)
    ListContestProblemDesc(c controller.MContext)
    GetContestProblem(c controller.MContext)
    PutContestProblem(c controller.MContext)
    DeleteContestProblem(c controller.MContext)
    ListContestUsers(c controller.MContext)
    ListContestProblem(c controller.MContext)
    GetContest(c controller.MContext)
    PutContest(c controller.MContext)
    DeleteContest(c controller.MContext)

}