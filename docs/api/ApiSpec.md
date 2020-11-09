
# Api Documentation

<!--beg l desc -->

<!--end l-->

## Api Reference

+ [AnnouncementController](./AnnouncementController.md)
    + [CountAnnouncement](./AnnouncementController.md#CountAnnouncement): The uri/restful key of this method is `/announcement-count@GET`
    + [ListAnnouncement](./AnnouncementController.md#ListAnnouncement): The uri/restful key of this method is `/announcement-list@GET`
    
        params: page ([integer]()), page_size ([integer]())
    + [DeleteAnnouncement](./AnnouncementController.md#DeleteAnnouncement): The uri/restful key of this method is `/announcement/{aid}@DELETE`
    
        params: aid ([!string]()), DeleteAnnouncementRequest ([any]())
    + [GetAnnouncement](./AnnouncementController.md#GetAnnouncement): The uri/restful key of this method is `/announcement/{aid}@GET`
    
        params: aid ([!string]())
    + [PutAnnouncement](./AnnouncementController.md#PutAnnouncement): The uri/restful key of this method is `/announcement/{aid}@PUT`
    
        params: aid ([!string]()), PutAnnouncementRequest ([any]())
    + [PostAnnouncement](./AnnouncementController.md#PostAnnouncement): The uri/restful key of this method is `/announcement@POST`
    
        params: PostAnnouncementRequest ([any]())
    
    <!--beg l desc_AnnouncementController -->
    
    <!--end l-->

+ [CommentController](./CommentController.md)
    + [CountComment](./CommentController.md#CountComment): The uri/restful key of this method is `/comment-count@GET`
    
        params: ref ([integer]()), no_reply ([boolean]()), page ([integer]()), page_size ([integer]()), ref_type ([integer]())
    + [ListComment](./CommentController.md#ListComment): The uri/restful key of this method is `/comment-list@GET`
    
        params: page_size ([integer]()), ref_type ([integer]()), ref ([integer]()), no_reply ([boolean]()), page ([integer]())
    + [DeleteComment](./CommentController.md#DeleteComment): The uri/restful key of this method is `/comment/{cmid}@DELETE`
    
        params: cmid ([!string]()), DeleteCommentRequest ([any]())
    + [GetComment](./CommentController.md#GetComment): The uri/restful key of this method is `/comment/{cmid}@GET`
    
        params: cmid ([!string]())
    + [PutComment](./CommentController.md#PutComment): The uri/restful key of this method is `/comment/{cmid}@PUT`
    
        params: cmid ([!string]()), PutCommentRequest ([any]())
    + [PostComment](./CommentController.md#PostComment): The uri/restful key of this method is `/comment@POST`
    
        params: PostCommentRequest ([any]())
    
    <!--beg l desc_CommentController -->
    
    <!--end l-->

+ [ContestController](./ContestController.md)
    + [CountContest](./ContestController.md#CountContest): The uri/restful key of this method is `/contest-count@GET`
    
        params: before_id ([integer]()), order ([string]()), page ([integer]()), page_size ([integer]())
    + [ListContest](./ContestController.md#ListContest): The uri/restful key of this method is `/contest-list@GET`
    
        params: order ([string]()), page ([integer]()), page_size ([integer]()), before_id ([integer]())
    + [CountContestProblem](./ContestController.md#CountContestProblem): The uri/restful key of this method is `/contest/{cid}/problem-count@GET`
    
        params: cid ([!string]()), before_id ([integer]()), order ([string]()), page ([integer]()), page_size ([integer]())
    + [ListContestProblem](./ContestController.md#ListContestProblem): The uri/restful key of this method is `/contest/{cid}/problem-list@GET`
    
        params: cid ([!string]()), order ([string]()), page ([integer]()), page_size ([integer]()), before_id ([integer]())
    + [CountContestProblemDesc](./ContestController.md#CountContestProblemDesc): The uri/restful key of this method is `/contest/{cid}/problem/{pid}/desc-count@GET`
    
        params: cid ([!string]()), pid ([!string]()), order ([string]()), page ([integer]()), page_size ([integer]()), before_id ([integer]())
    + [ListContestProblemDesc](./ContestController.md#ListContestProblemDesc): The uri/restful key of this method is `/contest/{cid}/problem/{pid}/desc-list@GET`
    
        params: cid ([!string]()), pid ([!string]()), order ([string]()), page ([integer]()), page_size ([integer]()), before_id ([integer]())
    + [ChangeContestProblemDescriptionRef](./ContestController.md#ChangeContestProblemDescriptionRef): The uri/restful key of this method is `/contest/{cid}/problem/{pid}/desc/ref@POST`
    
        params: cid ([!string]()), pid ([!string]()), ChangeContestProblemDescriptionRefRequest ([any]())
    + [DeleteContestProblemDesc](./ContestController.md#DeleteContestProblemDesc): The uri/restful key of this method is `/contest/{cid}/problem/{pid}/desc@DELETE`
    
        params: cid ([!string]()), pid ([!string]()), DeleteContestProblemDescRequest ([any]())
    + [GetContestProblemDesc](./ContestController.md#GetContestProblemDesc): The uri/restful key of this method is `/contest/{cid}/problem/{pid}/desc@GET`
    
        params: cid ([!string]()), pid ([!string]()), name ([string]())
    + [PostContestProblemDesc](./ContestController.md#PostContestProblemDesc): The uri/restful key of this method is `/contest/{cid}/problem/{pid}/desc@POST`
    
        params: cid ([!string]()), pid ([!string]()), PostContestProblemDescRequest ([any]())
    + [PutContestProblemDesc](./ContestController.md#PutContestProblemDesc): The uri/restful key of this method is `/contest/{cid}/problem/{pid}/desc@PUT`
    
        params: cid ([!string]()), pid ([!string]()), PutContestProblemDescRequest ([any]())
    + [DeleteContestProblem](./ContestController.md#DeleteContestProblem): The uri/restful key of this method is `/contest/{cid}/problem/{pid}@DELETE`
    
        params: cid ([!string]()), pid ([!string]()), DeleteContestProblemRequest ([any]())
    + [GetContestProblem](./ContestController.md#GetContestProblem): The uri/restful key of this method is `/contest/{cid}/problem/{pid}@GET`
    
        params: cid ([!string]()), pid ([!string]())
    + [PutContestProblem](./ContestController.md#PutContestProblem): The uri/restful key of this method is `/contest/{cid}/problem/{pid}@PUT`
    
        params: cid ([!string]()), pid ([!string]()), PutContestProblemRequest ([any]())
    + [PostContestProblem](./ContestController.md#PostContestProblem): The uri/restful key of this method is `/contest/{cid}/problem@POST`
    
        params: cid ([!string]()), PostContestProblemRequest ([any]())
    + [ListContestUsers](./ContestController.md#ListContestUsers): The uri/restful key of this method is `/contest/{cid}/user-list@GET`
    
        params: cid ([!string]())
    + [DeleteContest](./ContestController.md#DeleteContest): The uri/restful key of this method is `/contest/{cid}@DELETE`
    
        params: cid ([!string]()), DeleteContestRequest ([any]())
    + [GetContest](./ContestController.md#GetContest): The uri/restful key of this method is `/contest/{cid}@GET`
    
        params: cid ([!string]())
    + [PutContest](./ContestController.md#PutContest): The uri/restful key of this method is `/contest/{cid}@PUT`
    
        params: cid ([!string]()), PutContestRequest ([any]())
    + [PostContest](./ContestController.md#PostContest): The uri/restful key of this method is `/contest@POST`
    
        params: PostContestRequest ([any]())
    
    <!--beg l desc_ContestController -->
    
    <!--end l-->

+ [GroupController](./GroupController.md)
    + [CountGroup](./GroupController.md#CountGroup): The uri/restful key of this method is `/group-count@GET`
    
        params: order ([string]()), page ([integer]()), page_size ([integer]()), before_id ([integer]())
    + [ListGroup](./GroupController.md#ListGroup): The uri/restful key of this method is `/group-list@GET`
    
        params: before_id ([integer]()), order ([string]()), page ([integer]()), page_size ([integer]())
    + [PutGroupOwner](./GroupController.md#PutGroupOwner): The uri/restful key of this method is `/group/{gid}/owner@PUT`
    
        params: gid ([!string]()), PutGroupOwnerRequest ([any]())
    + [GetGroupMembers](./GroupController.md#GetGroupMembers): The uri/restful key of this method is `/group/{gid}/user-list@GET`
    
        params: gid ([!string]()), order ([string]()), page ([integer]()), page_size ([integer]()), before_id ([integer]())
    + [PostGroupMember](./GroupController.md#PostGroupMember): The uri/restful key of this method is `/group/{gid}/user/{id}@POST`
    
        params: gid ([!string]()), id ([!string]()), PostGroupMemberRequest ([any]())
    + [DeleteGroup](./GroupController.md#DeleteGroup): The uri/restful key of this method is `/group/{gid}@DELETE`
    
        params: gid ([!string]()), DeleteGroupRequest ([any]())
    + [GetGroup](./GroupController.md#GetGroup): The uri/restful key of this method is `/group/{gid}@GET`
    
        params: gid ([!string]())
    + [PutGroup](./GroupController.md#PutGroup): The uri/restful key of this method is `/group/{gid}@PUT`
    
        params: gid ([!string]()), PutGroupRequest ([any]())
    + [PostGroup](./GroupController.md#PostGroup): The uri/restful key of this method is `/group@POST`
    
        params: PostGroupRequest ([any]())
    
    <!--beg l desc_GroupController -->
    
    <!--end l-->

+ [AuthController](./AuthController.md)
    + [RemoveGroupingPolicy](./AuthController.md#RemoveGroupingPolicy): The uri/restful key of this method is `/policy/group@DELETE`
    
        params: RemoveGroupingPolicyRequest ([any]())
    + [HasGroupingPolicy](./AuthController.md#HasGroupingPolicy): The uri/restful key of this method is `/policy/group@GET`
    
        params: subject ([string]()), group ([string]())
    + [AddGroupingPolicy](./AuthController.md#AddGroupingPolicy): The uri/restful key of this method is `/policy/group@POST`
    
        params: AddGroupingPolicyRequest ([any]())
    + [RemovePolicy](./AuthController.md#RemovePolicy): The uri/restful key of this method is `/policy@DELETE`
    
        params: RemovePolicyRequest ([any]())
    + [HasPolicy](./AuthController.md#HasPolicy): The uri/restful key of this method is `/policy@GET`
    
        params: subject ([string]()), object ([string]()), action ([string]())
    + [AddPolicy](./AuthController.md#AddPolicy): The uri/restful key of this method is `/policy@POST`
    
        params: AddPolicyRequest ([any]())
    
    <!--beg l desc_AuthController -->
    
    <!--end l-->

+ [ProblemController](./ProblemController.md)
    + [CountProblem](./ProblemController.md#CountProblem): The uri/restful key of this method is `/problem-count@GET`
    
        params: page ([integer]()), page_size ([integer]()), before_id ([integer]()), order ([string]())
    + [ListProblem](./ProblemController.md#ListProblem): The uri/restful key of this method is `/problem-list@GET`
    
        params: order ([string]()), page ([integer]()), page_size ([integer]()), before_id ([integer]())
    + [CountProblemDesc](./ProblemController.md#CountProblemDesc): The uri/restful key of this method is `/problem/{pid}/desc-count@GET`
    
        params: pid ([!string]()), before_id ([integer]()), order ([string]()), page ([integer]()), page_size ([integer]())
    + [ListProblemDesc](./ProblemController.md#ListProblemDesc): The uri/restful key of this method is `/problem/{pid}/desc-list@GET`
    
        params: pid ([!string]()), order ([string]()), page ([integer]()), page_size ([integer]()), before_id ([integer]())
    + [ChangeProblemDescriptionRef](./ProblemController.md#ChangeProblemDescriptionRef): The uri/restful key of this method is `/problem/{pid}/desc/ref@POST`
    
        params: pid ([!string]()), ChangeProblemDescriptionRefRequest ([any]())
    + [DeleteProblemDesc](./ProblemController.md#DeleteProblemDesc): The uri/restful key of this method is `/problem/{pid}/desc@DELETE`
    
        params: pid ([!string]()), DeleteProblemDescRequest ([any]())
    + [GetProblemDesc](./ProblemController.md#GetProblemDesc): The uri/restful key of this method is `/problem/{pid}/desc@GET`
    
        params: pid ([!string]()), name ([string]())
    + [PostProblemDesc](./ProblemController.md#PostProblemDesc): The uri/restful key of this method is `/problem/{pid}/desc@POST`
    
        params: pid ([!string]()), PostProblemDescRequest ([any]())
    + [PutProblemDesc](./ProblemController.md#PutProblemDesc): The uri/restful key of this method is `/problem/{pid}/desc@PUT`
    
        params: pid ([!string]()), PutProblemDescRequest ([any]())
    + [ProblemFSReadConfig](./ProblemController.md#ProblemFSReadConfig): The uri/restful key of this method is `/problem/{pid}/fs/config@GET`
    
        params: pid ([!string]()), path ([string]())
    + [ProblemFSWriteConfig](./ProblemController.md#ProblemFSWriteConfig): The uri/restful key of this method is `/problem/{pid}/fs/config@POST`
    
        params: pid ([!string]()), ProblemFSWriteConfigRequest ([any]())
    + [ProblemFSPutConfig](./ProblemController.md#ProblemFSPutConfig): The uri/restful key of this method is `/problem/{pid}/fs/config@PUT`
    
        params: pid ([!string]()), ProblemFSPutConfigRequest ([any]())
    + [ProblemFSZipRead](./ProblemController.md#ProblemFSZipRead): The uri/restful key of this method is `/problem/{pid}/fs/directory/zip@GET`
    
        params: pid ([!string]()), path ([string]())
    + [ProblemFSZipWrite](./ProblemController.md#ProblemFSZipWrite): The uri/restful key of this method is `/problem/{pid}/fs/directory/zip@POST`
    
        params: pid ([!string]()), ProblemFSZipWriteRequest ([any]())
    + [ProblemFSRemoveAll](./ProblemController.md#ProblemFSRemoveAll): The uri/restful key of this method is `/problem/{pid}/fs/directory@DELETE`
    
        params: pid ([!string]()), ProblemFSRemoveAllRequest ([any]())
    + [ProblemFSLS](./ProblemController.md#ProblemFSLS): The uri/restful key of this method is `/problem/{pid}/fs/directory@GET`
    
        params: pid ([!string]()), path ([string]())
    + [ProblemFSWrites](./ProblemController.md#ProblemFSWrites): The uri/restful key of this method is `/problem/{pid}/fs/directory@POST`
    
        params: pid ([!string]()), ProblemFSWritesRequest ([any]())
    + [ProblemFSMkdir](./ProblemController.md#ProblemFSMkdir): The uri/restful key of this method is `/problem/{pid}/fs/directory@PUT`
    
        params: pid ([!string]()), ProblemFSMkdirRequest ([any]())
    + [ProblemFSRead](./ProblemController.md#ProblemFSRead): The uri/restful key of this method is `/problem/{pid}/fs/file/content@GET`
    
        params: pid ([!string]()), path ([string]())
    + [ProblemFSRemove](./ProblemController.md#ProblemFSRemove): The uri/restful key of this method is `/problem/{pid}/fs/file@DELETE`
    
        params: pid ([!string]()), ProblemFSRemoveRequest ([any]())
    + [ProblemFSStat](./ProblemController.md#ProblemFSStat): The uri/restful key of this method is `/problem/{pid}/fs/file@GET`
    
        params: pid ([!string]()), path ([string]())
    + [ProblemFSWrite](./ProblemController.md#ProblemFSWrite): The uri/restful key of this method is `/problem/{pid}/fs/file@POST`
    
        params: pid ([!string]()), ProblemFSWriteRequest ([any]())
    + [DeleteProblem](./ProblemController.md#DeleteProblem): The uri/restful key of this method is `/problem/{pid}@DELETE`
    
        params: pid ([!string]()), DeleteProblemRequest ([any]())
    + [GetProblem](./ProblemController.md#GetProblem): The uri/restful key of this method is `/problem/{pid}@GET`
    
        params: pid ([!string]())
    + [PutProblem](./ProblemController.md#PutProblem): The uri/restful key of this method is `/problem/{pid}@PUT`
    
        params: pid ([!string]()), PutProblemRequest ([any]())
    + [PostProblem](./ProblemController.md#PostProblem): The uri/restful key of this method is `/problem@POST`
    
        params: PostProblemRequest ([any]())
    
    <!--beg l desc_ProblemController -->
    
    <!--end l-->

+ [SubmissionController](./SubmissionController.md)
    + [PostSubmission](./SubmissionController.md#PostSubmission): The uri/restful key of this method is `/problem/{pid}/submission@POST`
    
        params: pid ([!string]()), PostSubmissionRequest ([any]())
    + [CountSubmission](./SubmissionController.md#CountSubmission): The uri/restful key of this method is `/submission-count@GET`
    
        params: page ([integer]()), page_size ([integer]()), mem_order ([boolean]()), time_order ([boolean]()), id_order ([boolean]()), by_user ([integer]()), on_problem ([integer]()), with_language ([integer]()), has_status ([integer]())
    + [ListSubmission](./SubmissionController.md#ListSubmission): The uri/restful key of this method is `/submission-list@GET`
    
        params: page ([integer]()), page_size ([integer]()), mem_order ([boolean]()), time_order ([boolean]()), id_order ([boolean]()), by_user ([integer]()), on_problem ([integer]()), with_language ([integer]()), has_status ([integer]())
    + [GetSubmissionContent](./SubmissionController.md#GetSubmissionContent): The uri/restful key of this method is `/submission/{sid}/content@GET`
    
        params: sid ([!string]())
    + [DeleteSubmission](./SubmissionController.md#DeleteSubmission): The uri/restful key of this method is `/submission/{sid}@DELETE`
    
        params: sid ([!string]()), DeleteSubmissionRequest ([any]())
    + [GetSubmission](./SubmissionController.md#GetSubmission): The uri/restful key of this method is `/submission/{sid}@GET`
    
        params: sid ([!string]())
    
    <!--beg l desc_SubmissionController -->
    
    <!--end l-->

+ [UserController](./UserController.md)
    + [CountUser](./UserController.md#CountUser): The uri/restful key of this method is `/user-count@GET`
    
        params: page ([integer]()), page_size ([integer]())
    + [ListUser](./UserController.md#ListUser): The uri/restful key of this method is `/user-list@GET`
    
        params: page ([integer]()), page_size ([integer]())
    + [RefreshToken](./UserController.md#RefreshToken): The uri/restful key of this method is `/user-token@GET`
    + [LoginUser](./UserController.md#LoginUser): The uri/restful key of this method is `/user/login@POST`
    
        params: LoginUserRequest ([any]())
    + [Register](./UserController.md#Register): The uri/restful key of this method is `/user/register@POST`
    
        params: RegisterRequest ([any]())
    + [BindEmail](./UserController.md#BindEmail): The uri/restful key of this method is `/user/{id}/email@PUT`
    
        params: id ([!string]()), BindEmailRequest ([any]())
    + [InspectUser](./UserController.md#InspectUser): The uri/restful key of this method is `/user/{id}/inspect@GET`
    
        params: id ([!string]())
    + [ChangePassword](./UserController.md#ChangePassword): The uri/restful key of this method is `/user/{id}/password@PUT`
    
        params: id ([!string]()), ChangePasswordRequest ([any]())
    + [DeleteUser](./UserController.md#DeleteUser): The uri/restful key of this method is `/user/{id}@DELETE`
    
        params: id ([!string]()), DeleteUserRequest ([any]())
    + [GetUser](./UserController.md#GetUser): The uri/restful key of this method is `/user/{id}@GET`
    
        params: id ([!string]())
    + [PutUser](./UserController.md#PutUser): The uri/restful key of this method is `/user/{id}@PUT`
    
        params: id ([!string]()), PutUserRequest ([any]())
    
    <!--beg l desc_UserController -->
    
    <!--end l-->

