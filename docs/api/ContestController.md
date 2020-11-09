
# Api Group Documentation (ContestController)


## CountContest

The uri/restful key of this method is `/contest-count@GET`

<!--beg l desc_CountContest -->

<!--end l-->


+ `before_id`: `integer`: 
    <!--beg l desc_CountContest_before_id -->
    
    <!--end l-->

+ `order`: `string`: 
    <!--beg l desc_CountContest_order -->
    
    <!--end l-->

+ `page`: `integer`: 
    <!--beg l desc_CountContest_page -->
    
    <!--end l-->

+ `page_size`: `integer`: 
    <!--beg l desc_CountContest_page_size -->
    
    <!--end l-->



## ListContest

The uri/restful key of this method is `/contest-list@GET`

<!--beg l desc_ListContest -->

<!--end l-->


+ `order`: `string`: 
    <!--beg l desc_ListContest_order -->
    
    <!--end l-->

+ `page`: `integer`: 
    <!--beg l desc_ListContest_page -->
    
    <!--end l-->

+ `page_size`: `integer`: 
    <!--beg l desc_ListContest_page_size -->
    
    <!--end l-->

+ `before_id`: `integer`: 
    <!--beg l desc_ListContest_before_id -->
    
    <!--end l-->



## CountContestProblem

The uri/restful key of this method is `/contest/{cid}/problem-count@GET`

<!--beg l desc_CountContestProblem -->

<!--end l-->


+ `cid`: `string` (required): 
    <!--beg l desc_CountContestProblem_cid -->
    
    <!--end l-->

+ `before_id`: `integer`: 
    <!--beg l desc_CountContestProblem_before_id -->
    
    <!--end l-->

+ `order`: `string`: 
    <!--beg l desc_CountContestProblem_order -->
    
    <!--end l-->

+ `page`: `integer`: 
    <!--beg l desc_CountContestProblem_page -->
    
    <!--end l-->

+ `page_size`: `integer`: 
    <!--beg l desc_CountContestProblem_page_size -->
    
    <!--end l-->



## ListContestProblem

The uri/restful key of this method is `/contest/{cid}/problem-list@GET`

<!--beg l desc_ListContestProblem -->

<!--end l-->


+ `cid`: `string` (required): 
    <!--beg l desc_ListContestProblem_cid -->
    
    <!--end l-->

+ `order`: `string`: 
    <!--beg l desc_ListContestProblem_order -->
    
    <!--end l-->

+ `page`: `integer`: 
    <!--beg l desc_ListContestProblem_page -->
    
    <!--end l-->

+ `page_size`: `integer`: 
    <!--beg l desc_ListContestProblem_page_size -->
    
    <!--end l-->

+ `before_id`: `integer`: 
    <!--beg l desc_ListContestProblem_before_id -->
    
    <!--end l-->



## CountContestProblemDesc

The uri/restful key of this method is `/contest/{cid}/problem/{pid}/desc-count@GET`

<!--beg l desc_CountContestProblemDesc -->

<!--end l-->


+ `cid`: `string` (required): 
    <!--beg l desc_CountContestProblemDesc_cid -->
    
    <!--end l-->

+ `pid`: `string` (required): 
    <!--beg l desc_CountContestProblemDesc_pid -->
    
    <!--end l-->

+ `order`: `string`: 
    <!--beg l desc_CountContestProblemDesc_order -->
    
    <!--end l-->

+ `page`: `integer`: 
    <!--beg l desc_CountContestProblemDesc_page -->
    
    <!--end l-->

+ `page_size`: `integer`: 
    <!--beg l desc_CountContestProblemDesc_page_size -->
    
    <!--end l-->

+ `before_id`: `integer`: 
    <!--beg l desc_CountContestProblemDesc_before_id -->
    
    <!--end l-->



## ListContestProblemDesc

The uri/restful key of this method is `/contest/{cid}/problem/{pid}/desc-list@GET`

<!--beg l desc_ListContestProblemDesc -->

<!--end l-->


+ `cid`: `string` (required): 
    <!--beg l desc_ListContestProblemDesc_cid -->
    
    <!--end l-->

+ `pid`: `string` (required): 
    <!--beg l desc_ListContestProblemDesc_pid -->
    
    <!--end l-->

+ `order`: `string`: 
    <!--beg l desc_ListContestProblemDesc_order -->
    
    <!--end l-->

+ `page`: `integer`: 
    <!--beg l desc_ListContestProblemDesc_page -->
    
    <!--end l-->

+ `page_size`: `integer`: 
    <!--beg l desc_ListContestProblemDesc_page_size -->
    
    <!--end l-->

+ `before_id`: `integer`: 
    <!--beg l desc_ListContestProblemDesc_before_id -->
    
    <!--end l-->



## ChangeContestProblemDescriptionRef

The uri/restful key of this method is `/contest/{cid}/problem/{pid}/desc/ref@POST`

<!--beg l desc_ChangeContestProblemDescriptionRef -->

<!--end l-->


+ `cid`: `string` (required): 
    <!--beg l desc_ChangeContestProblemDescriptionRef_cid -->
    
    <!--end l-->

+ `pid`: `string` (required): 
    <!--beg l desc_ChangeContestProblemDescriptionRef_pid -->
    
    <!--end l-->

+ `ChangeContestProblemDescriptionRefRequest`: `any`: 
    <!--beg l desc_ChangeContestProblemDescriptionRef_ChangeContestProblemDescriptionRefRequest -->
    
    <!--end l-->



## DeleteContestProblemDesc

The uri/restful key of this method is `/contest/{cid}/problem/{pid}/desc@DELETE`

<!--beg l desc_DeleteContestProblemDesc -->

<!--end l-->


+ `cid`: `string` (required): 
    <!--beg l desc_DeleteContestProblemDesc_cid -->
    
    <!--end l-->

+ `pid`: `string` (required): 
    <!--beg l desc_DeleteContestProblemDesc_pid -->
    
    <!--end l-->

+ `DeleteContestProblemDescRequest`: `any`: 
    <!--beg l desc_DeleteContestProblemDesc_DeleteContestProblemDescRequest -->
    
    <!--end l-->



## GetContestProblemDesc

The uri/restful key of this method is `/contest/{cid}/problem/{pid}/desc@GET`

<!--beg l desc_GetContestProblemDesc -->

<!--end l-->


+ `cid`: `string` (required): 
    <!--beg l desc_GetContestProblemDesc_cid -->
    
    <!--end l-->

+ `pid`: `string` (required): 
    <!--beg l desc_GetContestProblemDesc_pid -->
    
    <!--end l-->

+ `name`: `string`: 
    <!--beg l desc_GetContestProblemDesc_name -->
    
    <!--end l-->



## PostContestProblemDesc

The uri/restful key of this method is `/contest/{cid}/problem/{pid}/desc@POST`

<!--beg l desc_PostContestProblemDesc -->

<!--end l-->


+ `cid`: `string` (required): 
    <!--beg l desc_PostContestProblemDesc_cid -->
    
    <!--end l-->

+ `pid`: `string` (required): 
    <!--beg l desc_PostContestProblemDesc_pid -->
    
    <!--end l-->

+ `PostContestProblemDescRequest`: `any`: 
    <!--beg l desc_PostContestProblemDesc_PostContestProblemDescRequest -->
    
    <!--end l-->



## PutContestProblemDesc

The uri/restful key of this method is `/contest/{cid}/problem/{pid}/desc@PUT`

<!--beg l desc_PutContestProblemDesc -->

<!--end l-->


+ `cid`: `string` (required): 
    <!--beg l desc_PutContestProblemDesc_cid -->
    
    <!--end l-->

+ `pid`: `string` (required): 
    <!--beg l desc_PutContestProblemDesc_pid -->
    
    <!--end l-->

+ `PutContestProblemDescRequest`: `any`: 
    <!--beg l desc_PutContestProblemDesc_PutContestProblemDescRequest -->
    
    <!--end l-->



## DeleteContestProblem

The uri/restful key of this method is `/contest/{cid}/problem/{pid}@DELETE`

<!--beg l desc_DeleteContestProblem -->

<!--end l-->


+ `cid`: `string` (required): 
    <!--beg l desc_DeleteContestProblem_cid -->
    
    <!--end l-->

+ `pid`: `string` (required): 
    <!--beg l desc_DeleteContestProblem_pid -->
    
    <!--end l-->

+ `DeleteContestProblemRequest`: `any`: 
    <!--beg l desc_DeleteContestProblem_DeleteContestProblemRequest -->
    
    <!--end l-->



## GetContestProblem

The uri/restful key of this method is `/contest/{cid}/problem/{pid}@GET`

<!--beg l desc_GetContestProblem -->

<!--end l-->


+ `cid`: `string` (required): 
    <!--beg l desc_GetContestProblem_cid -->
    
    <!--end l-->

+ `pid`: `string` (required): 
    <!--beg l desc_GetContestProblem_pid -->
    
    <!--end l-->



## PutContestProblem

The uri/restful key of this method is `/contest/{cid}/problem/{pid}@PUT`

<!--beg l desc_PutContestProblem -->

<!--end l-->


+ `cid`: `string` (required): 
    <!--beg l desc_PutContestProblem_cid -->
    
    <!--end l-->

+ `pid`: `string` (required): 
    <!--beg l desc_PutContestProblem_pid -->
    
    <!--end l-->

+ `PutContestProblemRequest`: `any`: 
    <!--beg l desc_PutContestProblem_PutContestProblemRequest -->
    
    <!--end l-->



## PostContestProblem

The uri/restful key of this method is `/contest/{cid}/problem@POST`

<!--beg l desc_PostContestProblem -->

<!--end l-->


+ `cid`: `string` (required): 
    <!--beg l desc_PostContestProblem_cid -->
    
    <!--end l-->

+ `PostContestProblemRequest`: `any`: 
    <!--beg l desc_PostContestProblem_PostContestProblemRequest -->
    
    <!--end l-->



## ListContestUsers

The uri/restful key of this method is `/contest/{cid}/user-list@GET`

<!--beg l desc_ListContestUsers -->

<!--end l-->


+ `cid`: `string` (required): 
    <!--beg l desc_ListContestUsers_cid -->
    
    <!--end l-->



## DeleteContest

The uri/restful key of this method is `/contest/{cid}@DELETE`

<!--beg l desc_DeleteContest -->

<!--end l-->


+ `cid`: `string` (required): 
    <!--beg l desc_DeleteContest_cid -->
    
    <!--end l-->

+ `DeleteContestRequest`: `any`: 
    <!--beg l desc_DeleteContest_DeleteContestRequest -->
    
    <!--end l-->



## GetContest

The uri/restful key of this method is `/contest/{cid}@GET`

<!--beg l desc_GetContest -->

<!--end l-->


+ `cid`: `string` (required): 
    <!--beg l desc_GetContest_cid -->
    
    <!--end l-->



## PutContest

The uri/restful key of this method is `/contest/{cid}@PUT`

<!--beg l desc_PutContest -->

<!--end l-->


+ `cid`: `string` (required): 
    <!--beg l desc_PutContest_cid -->
    
    <!--end l-->

+ `PutContestRequest`: `any`: 
    <!--beg l desc_PutContest_PutContestRequest -->
    
    <!--end l-->



## PostContest

The uri/restful key of this method is `/contest@POST`

<!--beg l desc_PostContest -->

<!--end l-->


+ `PostContestRequest`: `any`: 
    <!--beg l desc_PostContest_PostContestRequest -->
    
    <!--end l-->



