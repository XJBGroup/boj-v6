
# Api Group Documentation (ContestController)

<!--beg l desc_ContestController -->

<!--end l-->

## Apis


### CountContest

restful key: `/contest-count@GET`

<!--beg l desc_CountContest -->

<!--end l-->

parameters:

+ `before_id`: [integer](#integer): 
    <!--beg l desc_CountContest_params_before_id -->
    
    <!--end l-->


+ `order`: [string](#string): 
    <!--beg l desc_CountContest_params_order -->
    
    <!--end l-->


+ `page`: [integer](#integer): 
    <!--beg l desc_CountContest_params_page -->
    
    <!--end l-->


+ `page_size`: [integer](#integer): 
    <!--beg l desc_CountContest_params_page_size -->
    
    <!--end l-->


responses:

+ code: `200`, type: [genericResponse](#genericResponse)
    <!--beg l desc_CountContest_response_200_[genericResponse](#genericResponse) -->
    
    <!--end l-->


+ code: `200`, type: [CountContestReply](#CountContestReply)
    <!--beg l desc_CountContest_response_200_[CountContestReply](#CountContestReply) -->
    
    <!--end l-->


+ code: `500`, type: No Response
    <!--beg l desc_CountContest_response_500_No Response -->
    
    <!--end l-->




### ListContest

restful key: `/contest-list@GET`

<!--beg l desc_ListContest -->

<!--end l-->

parameters:

+ `order`: [string](#string): 
    <!--beg l desc_ListContest_params_order -->
    
    <!--end l-->


+ `page`: [integer](#integer): 
    <!--beg l desc_ListContest_params_page -->
    
    <!--end l-->


+ `page_size`: [integer](#integer): 
    <!--beg l desc_ListContest_params_page_size -->
    
    <!--end l-->


+ `before_id`: [integer](#integer): 
    <!--beg l desc_ListContest_params_before_id -->
    
    <!--end l-->


responses:

+ code: `200`, type: [genericResponse](#genericResponse)
    <!--beg l desc_ListContest_response_200_[genericResponse](#genericResponse) -->
    
    <!--end l-->


+ code: `200`, type: [ListContestReply](#ListContestReply)
    <!--beg l desc_ListContest_response_200_[ListContestReply](#ListContestReply) -->
    
    <!--end l-->


+ code: `500`, type: No Response
    <!--beg l desc_ListContest_response_500_No Response -->
    
    <!--end l-->




### CountContestProblem

restful key: `/contest/{cid}/problem-count@GET`

<!--beg l desc_CountContestProblem -->

<!--end l-->

parameters:

+ `cid`: [string](#string) (required): 
    <!--beg l desc_CountContestProblem_params_cid -->
    
    <!--end l-->


+ `before_id`: [integer](#integer): 
    <!--beg l desc_CountContestProblem_params_before_id -->
    
    <!--end l-->


+ `order`: [string](#string): 
    <!--beg l desc_CountContestProblem_params_order -->
    
    <!--end l-->


+ `page`: [integer](#integer): 
    <!--beg l desc_CountContestProblem_params_page -->
    
    <!--end l-->


+ `page_size`: [integer](#integer): 
    <!--beg l desc_CountContestProblem_params_page_size -->
    
    <!--end l-->


responses:

+ code: `200`, type: [genericResponse](#genericResponse)
    <!--beg l desc_CountContestProblem_response_200_[genericResponse](#genericResponse) -->
    
    <!--end l-->


+ code: `200`, type: [CountContestProblemReply](#CountContestProblemReply)
    <!--beg l desc_CountContestProblem_response_200_[CountContestProblemReply](#CountContestProblemReply) -->
    
    <!--end l-->


+ code: `500`, type: No Response
    <!--beg l desc_CountContestProblem_response_500_No Response -->
    
    <!--end l-->




### ListContestProblem

restful key: `/contest/{cid}/problem-list@GET`

<!--beg l desc_ListContestProblem -->

<!--end l-->

parameters:

+ `cid`: [string](#string) (required): 
    <!--beg l desc_ListContestProblem_params_cid -->
    
    <!--end l-->


+ `order`: [string](#string): 
    <!--beg l desc_ListContestProblem_params_order -->
    
    <!--end l-->


+ `page`: [integer](#integer): 
    <!--beg l desc_ListContestProblem_params_page -->
    
    <!--end l-->


+ `page_size`: [integer](#integer): 
    <!--beg l desc_ListContestProblem_params_page_size -->
    
    <!--end l-->


+ `before_id`: [integer](#integer): 
    <!--beg l desc_ListContestProblem_params_before_id -->
    
    <!--end l-->


responses:

+ code: `200`, type: [genericResponse](#genericResponse)
    <!--beg l desc_ListContestProblem_response_200_[genericResponse](#genericResponse) -->
    
    <!--end l-->


+ code: `200`, type: [ListContestProblemReply](#ListContestProblemReply)
    <!--beg l desc_ListContestProblem_response_200_[ListContestProblemReply](#ListContestProblemReply) -->
    
    <!--end l-->


+ code: `500`, type: No Response
    <!--beg l desc_ListContestProblem_response_500_No Response -->
    
    <!--end l-->




### CountContestProblemDesc

restful key: `/contest/{cid}/problem/{pid}/desc-count@GET`

<!--beg l desc_CountContestProblemDesc -->

<!--end l-->

parameters:

+ `cid`: [string](#string) (required): 
    <!--beg l desc_CountContestProblemDesc_params_cid -->
    
    <!--end l-->


+ `pid`: [string](#string) (required): 
    <!--beg l desc_CountContestProblemDesc_params_pid -->
    
    <!--end l-->


+ `order`: [string](#string): 
    <!--beg l desc_CountContestProblemDesc_params_order -->
    
    <!--end l-->


+ `page`: [integer](#integer): 
    <!--beg l desc_CountContestProblemDesc_params_page -->
    
    <!--end l-->


+ `page_size`: [integer](#integer): 
    <!--beg l desc_CountContestProblemDesc_params_page_size -->
    
    <!--end l-->


+ `before_id`: [integer](#integer): 
    <!--beg l desc_CountContestProblemDesc_params_before_id -->
    
    <!--end l-->


responses:

+ code: `200`, type: [genericResponse](#genericResponse)
    <!--beg l desc_CountContestProblemDesc_response_200_[genericResponse](#genericResponse) -->
    
    <!--end l-->


+ code: `200`, type: [CountContestProblemDescReply](#CountContestProblemDescReply)
    <!--beg l desc_CountContestProblemDesc_response_200_[CountContestProblemDescReply](#CountContestProblemDescReply) -->
    
    <!--end l-->


+ code: `500`, type: No Response
    <!--beg l desc_CountContestProblemDesc_response_500_No Response -->
    
    <!--end l-->




### ListContestProblemDesc

restful key: `/contest/{cid}/problem/{pid}/desc-list@GET`

<!--beg l desc_ListContestProblemDesc -->

<!--end l-->

parameters:

+ `cid`: [string](#string) (required): 
    <!--beg l desc_ListContestProblemDesc_params_cid -->
    
    <!--end l-->


+ `pid`: [string](#string) (required): 
    <!--beg l desc_ListContestProblemDesc_params_pid -->
    
    <!--end l-->


+ `order`: [string](#string): 
    <!--beg l desc_ListContestProblemDesc_params_order -->
    
    <!--end l-->


+ `page`: [integer](#integer): 
    <!--beg l desc_ListContestProblemDesc_params_page -->
    
    <!--end l-->


+ `page_size`: [integer](#integer): 
    <!--beg l desc_ListContestProblemDesc_params_page_size -->
    
    <!--end l-->


+ `before_id`: [integer](#integer): 
    <!--beg l desc_ListContestProblemDesc_params_before_id -->
    
    <!--end l-->


responses:

+ code: `200`, type: [genericResponse](#genericResponse)
    <!--beg l desc_ListContestProblemDesc_response_200_[genericResponse](#genericResponse) -->
    
    <!--end l-->


+ code: `200`, type: [ListContestProblemDescReply](#ListContestProblemDescReply)
    <!--beg l desc_ListContestProblemDesc_response_200_[ListContestProblemDescReply](#ListContestProblemDescReply) -->
    
    <!--end l-->


+ code: `500`, type: No Response
    <!--beg l desc_ListContestProblemDesc_response_500_No Response -->
    
    <!--end l-->




### ChangeContestProblemDescriptionRef

restful key: `/contest/{cid}/problem/{pid}/desc/ref@POST`

<!--beg l desc_ChangeContestProblemDescriptionRef -->

<!--end l-->

parameters:

+ `cid`: [string](#string) (required): 
    <!--beg l desc_ChangeContestProblemDescriptionRef_params_cid -->
    
    <!--end l-->


+ `pid`: [string](#string) (required): 
    <!--beg l desc_ChangeContestProblemDescriptionRef_params_pid -->
    
    <!--end l-->


+ `ChangeContestProblemDescriptionRefRequest`: [ChangeContestProblemDescriptionRefRequest](#ChangeContestProblemDescriptionRefRequest): 
    <!--beg l desc_ChangeContestProblemDescriptionRef_params_ChangeContestProblemDescriptionRefRequest -->
    
    <!--end l-->


responses:

+ code: `200`, type: [genericResponse](#genericResponse)
    <!--beg l desc_ChangeContestProblemDescriptionRef_response_200_[genericResponse](#genericResponse) -->
    
    <!--end l-->


+ code: `200`, type: [ChangeContestProblemDescriptionRefReply](#ChangeContestProblemDescriptionRefReply)
    <!--beg l desc_ChangeContestProblemDescriptionRef_response_200_[ChangeContestProblemDescriptionRefReply](#ChangeContestProblemDescriptionRefReply) -->
    
    <!--end l-->


+ code: `500`, type: No Response
    <!--beg l desc_ChangeContestProblemDescriptionRef_response_500_No Response -->
    
    <!--end l-->




### DeleteContestProblemDesc

restful key: `/contest/{cid}/problem/{pid}/desc@DELETE`

<!--beg l desc_DeleteContestProblemDesc -->

<!--end l-->

parameters:

+ `cid`: [string](#string) (required): 
    <!--beg l desc_DeleteContestProblemDesc_params_cid -->
    
    <!--end l-->


+ `pid`: [string](#string) (required): 
    <!--beg l desc_DeleteContestProblemDesc_params_pid -->
    
    <!--end l-->


+ `DeleteContestProblemDescRequest`: [DeleteContestProblemDescRequest](#DeleteContestProblemDescRequest): 
    <!--beg l desc_DeleteContestProblemDesc_params_DeleteContestProblemDescRequest -->
    
    <!--end l-->


responses:

+ code: `200`, type: [genericResponse](#genericResponse)
    <!--beg l desc_DeleteContestProblemDesc_response_200_[genericResponse](#genericResponse) -->
    
    <!--end l-->


+ code: `200`, type: [DeleteContestProblemDescReply](#DeleteContestProblemDescReply)
    <!--beg l desc_DeleteContestProblemDesc_response_200_[DeleteContestProblemDescReply](#DeleteContestProblemDescReply) -->
    
    <!--end l-->


+ code: `500`, type: No Response
    <!--beg l desc_DeleteContestProblemDesc_response_500_No Response -->
    
    <!--end l-->




### GetContestProblemDesc

restful key: `/contest/{cid}/problem/{pid}/desc@GET`

<!--beg l desc_GetContestProblemDesc -->

<!--end l-->

parameters:

+ `cid`: [string](#string) (required): 
    <!--beg l desc_GetContestProblemDesc_params_cid -->
    
    <!--end l-->


+ `pid`: [string](#string) (required): 
    <!--beg l desc_GetContestProblemDesc_params_pid -->
    
    <!--end l-->


+ `name`: [string](#string): 
    <!--beg l desc_GetContestProblemDesc_params_name -->
    
    <!--end l-->


responses:

+ code: `200`, type: [genericResponse](#genericResponse)
    <!--beg l desc_GetContestProblemDesc_response_200_[genericResponse](#genericResponse) -->
    
    <!--end l-->


+ code: `200`, type: [GetContestProblemDescReply](#GetContestProblemDescReply)
    <!--beg l desc_GetContestProblemDesc_response_200_[GetContestProblemDescReply](#GetContestProblemDescReply) -->
    
    <!--end l-->


+ code: `500`, type: No Response
    <!--beg l desc_GetContestProblemDesc_response_500_No Response -->
    
    <!--end l-->




### PostContestProblemDesc

restful key: `/contest/{cid}/problem/{pid}/desc@POST`

<!--beg l desc_PostContestProblemDesc -->

<!--end l-->

parameters:

+ `cid`: [string](#string) (required): 
    <!--beg l desc_PostContestProblemDesc_params_cid -->
    
    <!--end l-->


+ `pid`: [string](#string) (required): 
    <!--beg l desc_PostContestProblemDesc_params_pid -->
    
    <!--end l-->


+ `PostContestProblemDescRequest`: [PostContestProblemDescRequest](#PostContestProblemDescRequest): 
    <!--beg l desc_PostContestProblemDesc_params_PostContestProblemDescRequest -->
    
    <!--end l-->


responses:

+ code: `200`, type: [genericResponse](#genericResponse)
    <!--beg l desc_PostContestProblemDesc_response_200_[genericResponse](#genericResponse) -->
    
    <!--end l-->


+ code: `200`, type: [PostContestProblemDescReply](#PostContestProblemDescReply)
    <!--beg l desc_PostContestProblemDesc_response_200_[PostContestProblemDescReply](#PostContestProblemDescReply) -->
    
    <!--end l-->


+ code: `500`, type: No Response
    <!--beg l desc_PostContestProblemDesc_response_500_No Response -->
    
    <!--end l-->




### PutContestProblemDesc

restful key: `/contest/{cid}/problem/{pid}/desc@PUT`

<!--beg l desc_PutContestProblemDesc -->

<!--end l-->

parameters:

+ `cid`: [string](#string) (required): 
    <!--beg l desc_PutContestProblemDesc_params_cid -->
    
    <!--end l-->


+ `pid`: [string](#string) (required): 
    <!--beg l desc_PutContestProblemDesc_params_pid -->
    
    <!--end l-->


+ `PutContestProblemDescRequest`: [PutContestProblemDescRequest](#PutContestProblemDescRequest): 
    <!--beg l desc_PutContestProblemDesc_params_PutContestProblemDescRequest -->
    
    <!--end l-->


responses:

+ code: `200`, type: [genericResponse](#genericResponse)
    <!--beg l desc_PutContestProblemDesc_response_200_[genericResponse](#genericResponse) -->
    
    <!--end l-->


+ code: `200`, type: [PutContestProblemDescReply](#PutContestProblemDescReply)
    <!--beg l desc_PutContestProblemDesc_response_200_[PutContestProblemDescReply](#PutContestProblemDescReply) -->
    
    <!--end l-->


+ code: `500`, type: No Response
    <!--beg l desc_PutContestProblemDesc_response_500_No Response -->
    
    <!--end l-->




### DeleteContestProblem

restful key: `/contest/{cid}/problem/{pid}@DELETE`

<!--beg l desc_DeleteContestProblem -->

<!--end l-->

parameters:

+ `cid`: [string](#string) (required): 
    <!--beg l desc_DeleteContestProblem_params_cid -->
    
    <!--end l-->


+ `pid`: [string](#string) (required): 
    <!--beg l desc_DeleteContestProblem_params_pid -->
    
    <!--end l-->


+ `DeleteContestProblemRequest`: [DeleteContestProblemRequest](#DeleteContestProblemRequest): 
    <!--beg l desc_DeleteContestProblem_params_DeleteContestProblemRequest -->
    
    <!--end l-->


responses:

+ code: `200`, type: [genericResponse](#genericResponse)
    <!--beg l desc_DeleteContestProblem_response_200_[genericResponse](#genericResponse) -->
    
    <!--end l-->


+ code: `200`, type: [DeleteContestProblemReply](#DeleteContestProblemReply)
    <!--beg l desc_DeleteContestProblem_response_200_[DeleteContestProblemReply](#DeleteContestProblemReply) -->
    
    <!--end l-->


+ code: `500`, type: No Response
    <!--beg l desc_DeleteContestProblem_response_500_No Response -->
    
    <!--end l-->




### GetContestProblem

restful key: `/contest/{cid}/problem/{pid}@GET`

<!--beg l desc_GetContestProblem -->

<!--end l-->

parameters:

+ `cid`: [string](#string) (required): 
    <!--beg l desc_GetContestProblem_params_cid -->
    
    <!--end l-->


+ `pid`: [string](#string) (required): 
    <!--beg l desc_GetContestProblem_params_pid -->
    
    <!--end l-->


responses:

+ code: `200`, type: [genericResponse](#genericResponse)
    <!--beg l desc_GetContestProblem_response_200_[genericResponse](#genericResponse) -->
    
    <!--end l-->


+ code: `200`, type: [GetContestProblemReply](#GetContestProblemReply)
    <!--beg l desc_GetContestProblem_response_200_[GetContestProblemReply](#GetContestProblemReply) -->
    
    <!--end l-->


+ code: `500`, type: No Response
    <!--beg l desc_GetContestProblem_response_500_No Response -->
    
    <!--end l-->




### PutContestProblem

restful key: `/contest/{cid}/problem/{pid}@PUT`

<!--beg l desc_PutContestProblem -->

<!--end l-->

parameters:

+ `cid`: [string](#string) (required): 
    <!--beg l desc_PutContestProblem_params_cid -->
    
    <!--end l-->


+ `pid`: [string](#string) (required): 
    <!--beg l desc_PutContestProblem_params_pid -->
    
    <!--end l-->


+ `PutContestProblemRequest`: [PutContestProblemRequest](#PutContestProblemRequest): 
    <!--beg l desc_PutContestProblem_params_PutContestProblemRequest -->
    
    <!--end l-->


responses:

+ code: `200`, type: [genericResponse](#genericResponse)
    <!--beg l desc_PutContestProblem_response_200_[genericResponse](#genericResponse) -->
    
    <!--end l-->


+ code: `200`, type: [PutContestProblemReply](#PutContestProblemReply)
    <!--beg l desc_PutContestProblem_response_200_[PutContestProblemReply](#PutContestProblemReply) -->
    
    <!--end l-->


+ code: `500`, type: No Response
    <!--beg l desc_PutContestProblem_response_500_No Response -->
    
    <!--end l-->




### PostContestProblem

restful key: `/contest/{cid}/problem@POST`

<!--beg l desc_PostContestProblem -->

<!--end l-->

parameters:

+ `cid`: [string](#string) (required): 
    <!--beg l desc_PostContestProblem_params_cid -->
    
    <!--end l-->


+ `PostContestProblemRequest`: [PostContestProblemRequest](#PostContestProblemRequest): 
    <!--beg l desc_PostContestProblem_params_PostContestProblemRequest -->
    
    <!--end l-->


responses:

+ code: `200`, type: [genericResponse](#genericResponse)
    <!--beg l desc_PostContestProblem_response_200_[genericResponse](#genericResponse) -->
    
    <!--end l-->


+ code: `200`, type: [PostContestProblemReply](#PostContestProblemReply)
    <!--beg l desc_PostContestProblem_response_200_[PostContestProblemReply](#PostContestProblemReply) -->
    
    <!--end l-->


+ code: `500`, type: No Response
    <!--beg l desc_PostContestProblem_response_500_No Response -->
    
    <!--end l-->




### ListContestUsers

restful key: `/contest/{cid}/user-list@GET`

<!--beg l desc_ListContestUsers -->

<!--end l-->

parameters:

+ `cid`: [string](#string) (required): 
    <!--beg l desc_ListContestUsers_params_cid -->
    
    <!--end l-->


responses:

+ code: `200`, type: [genericResponse](#genericResponse)
    <!--beg l desc_ListContestUsers_response_200_[genericResponse](#genericResponse) -->
    
    <!--end l-->


+ code: `200`, type: [ListContestUsersReply](#ListContestUsersReply)
    <!--beg l desc_ListContestUsers_response_200_[ListContestUsersReply](#ListContestUsersReply) -->
    
    <!--end l-->


+ code: `500`, type: No Response
    <!--beg l desc_ListContestUsers_response_500_No Response -->
    
    <!--end l-->




### DeleteContest

restful key: `/contest/{cid}@DELETE`

<!--beg l desc_DeleteContest -->

<!--end l-->

parameters:

+ `cid`: [string](#string) (required): 
    <!--beg l desc_DeleteContest_params_cid -->
    
    <!--end l-->


+ `DeleteContestRequest`: [DeleteContestRequest](#DeleteContestRequest): 
    <!--beg l desc_DeleteContest_params_DeleteContestRequest -->
    
    <!--end l-->


responses:

+ code: `200`, type: [genericResponse](#genericResponse)
    <!--beg l desc_DeleteContest_response_200_[genericResponse](#genericResponse) -->
    
    <!--end l-->


+ code: `200`, type: [DeleteContestReply](#DeleteContestReply)
    <!--beg l desc_DeleteContest_response_200_[DeleteContestReply](#DeleteContestReply) -->
    
    <!--end l-->


+ code: `500`, type: No Response
    <!--beg l desc_DeleteContest_response_500_No Response -->
    
    <!--end l-->




### GetContest

restful key: `/contest/{cid}@GET`

<!--beg l desc_GetContest -->

<!--end l-->

parameters:

+ `cid`: [string](#string) (required): 
    <!--beg l desc_GetContest_params_cid -->
    
    <!--end l-->


responses:

+ code: `200`, type: [genericResponse](#genericResponse)
    <!--beg l desc_GetContest_response_200_[genericResponse](#genericResponse) -->
    
    <!--end l-->


+ code: `200`, type: [GetContestReply](#GetContestReply)
    <!--beg l desc_GetContest_response_200_[GetContestReply](#GetContestReply) -->
    
    <!--end l-->


+ code: `500`, type: No Response
    <!--beg l desc_GetContest_response_500_No Response -->
    
    <!--end l-->




### PutContest

restful key: `/contest/{cid}@PUT`

<!--beg l desc_PutContest -->

<!--end l-->

parameters:

+ `cid`: [string](#string) (required): 
    <!--beg l desc_PutContest_params_cid -->
    
    <!--end l-->


+ `PutContestRequest`: [PutContestRequest](#PutContestRequest): 
    <!--beg l desc_PutContest_params_PutContestRequest -->
    
    <!--end l-->


responses:

+ code: `200`, type: [genericResponse](#genericResponse)
    <!--beg l desc_PutContest_response_200_[genericResponse](#genericResponse) -->
    
    <!--end l-->


+ code: `200`, type: [PutContestReply](#PutContestReply)
    <!--beg l desc_PutContest_response_200_[PutContestReply](#PutContestReply) -->
    
    <!--end l-->


+ code: `500`, type: No Response
    <!--beg l desc_PutContest_response_500_No Response -->
    
    <!--end l-->




### PostContest

restful key: `/contest@POST`

<!--beg l desc_PostContest -->

<!--end l-->

parameters:

+ `PostContestRequest`: [PostContestRequest](#PostContestRequest): 
    <!--beg l desc_PostContest_params_PostContestRequest -->
    
    <!--end l-->


responses:

+ code: `200`, type: [genericResponse](#genericResponse)
    <!--beg l desc_PostContest_response_200_[genericResponse](#genericResponse) -->
    
    <!--end l-->


+ code: `200`, type: [PostContestReply](#PostContestReply)
    <!--beg l desc_PostContest_response_200_[PostContestReply](#PostContestReply) -->
    
    <!--end l-->


+ code: `500`, type: No Response
    <!--beg l desc_PostContest_response_500_No Response -->
    
    <!--end l-->




## Local Object Reference




### [genericResponse](./ObjectModelSpec.md#genericResponse)

+ type: [object](#genericResponse)

+ fields:
    
    + `code`: integer: 
        <!--beg l desc_{{object_name}}_code -->
        
        <!--end l-->

    + `error`: object: 
        <!--beg l desc_{{object_name}}_error -->
        
        <!--end l-->

    + `params`: object: 
        <!--beg l desc_{{object_name}}_params -->
        
        <!--end l-->

    
### [CountContestReply](./ObjectModelSpec.md#CountContestReply)

+ type: [object](#CountContestReply)

+ fields:
    
    + `code`: integer: 
        <!--beg l desc_{{object_name}}_code -->
        
        <!--end l-->

    + `data`: array: 
        <!--beg l desc_{{object_name}}_data -->
        
        <!--end l-->

    
### [ListContestReply](./ObjectModelSpec.md#ListContestReply)

+ type: [object](#ListContestReply)

+ fields:
    
    + `code`: integer: 
        <!--beg l desc_{{object_name}}_code -->
        
        <!--end l-->

    + `data`: array: 
        <!--beg l desc_{{object_name}}_data -->
        
        <!--end l-->

    
### [CountContestProblemReply](./ObjectModelSpec.md#CountContestProblemReply)

+ type: [object](#CountContestProblemReply)

+ fields:
    
    + `code`: integer: 
        <!--beg l desc_{{object_name}}_code -->
        
        <!--end l-->

    + `data`: array: 
        <!--beg l desc_{{object_name}}_data -->
        
        <!--end l-->

    
### [ListContestProblemReply](./ObjectModelSpec.md#ListContestProblemReply)

+ type: [object](#ListContestProblemReply)

+ fields:
    
    + `code`: integer: 
        <!--beg l desc_{{object_name}}_code -->
        
        <!--end l-->

    + `data`: array: 
        <!--beg l desc_{{object_name}}_data -->
        
        <!--end l-->

    
### [CountContestProblemDescReply](./ObjectModelSpec.md#CountContestProblemDescReply)

+ type: [object](#CountContestProblemDescReply)

+ fields:
    
    + `code`: integer: 
        <!--beg l desc_{{object_name}}_code -->
        
        <!--end l-->

    + `data`: integer: 
        <!--beg l desc_{{object_name}}_data -->
        
        <!--end l-->

    
### [ListContestProblemDescReply](./ObjectModelSpec.md#ListContestProblemDescReply)

+ type: [object](#ListContestProblemDescReply)

+ fields:
    
    + `code`: integer: 
        <!--beg l desc_{{object_name}}_code -->
        
        <!--end l-->

    + `data`: array: 
        <!--beg l desc_{{object_name}}_data -->
        
        <!--end l-->

    
### [ChangeContestProblemDescriptionRefRequest](./ObjectModelSpec.md#ChangeContestProblemDescriptionRefRequest)

+ type: [object](#ChangeContestProblemDescriptionRefRequest)

+ fields:
    
    + `name`: string: 
        <!--beg l desc_{{object_name}}_name -->
        
        <!--end l-->

    + `new_name`: string: 
        <!--beg l desc_{{object_name}}_new_name -->
        
        <!--end l-->

    
### [ChangeContestProblemDescriptionRefReply](./ObjectModelSpec.md#ChangeContestProblemDescriptionRefReply)

+ type: [object](#ChangeContestProblemDescriptionRefReply)

+ fields:
    
    + `code`: integer: 
        <!--beg l desc_{{object_name}}_code -->
        
        <!--end l-->

    
### [DeleteContestProblemDescRequest](./ObjectModelSpec.md#DeleteContestProblemDescRequest)

+ type: [object](#DeleteContestProblemDescRequest)

+ fields:
    
    + `name`: string: 
        <!--beg l desc_{{object_name}}_name -->
        
        <!--end l-->

    
### [DeleteContestProblemDescReply](./ObjectModelSpec.md#DeleteContestProblemDescReply)

+ type: [object](#DeleteContestProblemDescReply)

+ fields:
    
    + `code`: integer: 
        <!--beg l desc_{{object_name}}_code -->
        
        <!--end l-->

    
### [GetContestProblemDescReply](./ObjectModelSpec.md#GetContestProblemDescReply)

+ type: [object](#GetContestProblemDescReply)

+ fields:
    
    + `code`: integer: 
        <!--beg l desc_{{object_name}}_code -->
        
        <!--end l-->

    + `data`: string: 
        <!--beg l desc_{{object_name}}_data -->
        
        <!--end l-->

    
### [PostContestProblemDescRequest](./ObjectModelSpec.md#PostContestProblemDescRequest)

+ type: [object](#PostContestProblemDescRequest)

+ fields:
    
    + `content`: string: 
        <!--beg l desc_{{object_name}}_content -->
        
        <!--end l-->

    + `name`: string: 
        <!--beg l desc_{{object_name}}_name -->
        
        <!--end l-->

    
### [PostContestProblemDescReply](./ObjectModelSpec.md#PostContestProblemDescReply)

+ type: [object](#PostContestProblemDescReply)

+ fields:
    
    + `code`: integer: 
        <!--beg l desc_{{object_name}}_code -->
        
        <!--end l-->

    
### [PutContestProblemDescRequest](./ObjectModelSpec.md#PutContestProblemDescRequest)

+ type: [object](#PutContestProblemDescRequest)

+ fields:
    
    + `content`: string: 
        <!--beg l desc_{{object_name}}_content -->
        
        <!--end l-->

    + `name`: string: 
        <!--beg l desc_{{object_name}}_name -->
        
        <!--end l-->

    
### [PutContestProblemDescReply](./ObjectModelSpec.md#PutContestProblemDescReply)

+ type: [object](#PutContestProblemDescReply)

+ fields:
    
    + `code`: integer: 
        <!--beg l desc_{{object_name}}_code -->
        
        <!--end l-->

    
### [DeleteContestProblemRequest](./ObjectModelSpec.md#DeleteContestProblemRequest)

+ type: [object](#DeleteContestProblemRequest)

+ fields:
    
    
### [DeleteContestProblemReply](./ObjectModelSpec.md#DeleteContestProblemReply)

+ type: [object](#DeleteContestProblemReply)

+ fields:
    
    + `code`: integer: 
        <!--beg l desc_{{object_name}}_code -->
        
        <!--end l-->

    
### [GetContestProblemReply](./ObjectModelSpec.md#GetContestProblemReply)

+ type: [object](#GetContestProblemReply)

+ fields:
    
    + `code`: integer: 
        <!--beg l desc_{{object_name}}_code -->
        
        <!--end l-->

    + `data`: : 
        <!--beg l desc_{{object_name}}_data -->
        
        <!--end l-->

    
### [PutContestProblemRequest](./ObjectModelSpec.md#PutContestProblemRequest)

+ type: [object](#PutContestProblemRequest)

+ fields:
    
    + `description_ref`: string: 
        <!--beg l desc_{{object_name}}_description_ref -->
        
        <!--end l-->

    + `title`: string: 
        <!--beg l desc_{{object_name}}_title -->
        
        <!--end l-->

    
### [PutContestProblemReply](./ObjectModelSpec.md#PutContestProblemReply)

+ type: [object](#PutContestProblemReply)

+ fields:
    
    + `code`: integer: 
        <!--beg l desc_{{object_name}}_code -->
        
        <!--end l-->

    
### [PostContestProblemRequest](./ObjectModelSpec.md#PostContestProblemRequest)

+ type: [object](#PostContestProblemRequest)

+ fields:
    
    + `config`: : 
        <!--beg l desc_{{object_name}}_config -->
        
        <!--end l-->

    + `description`: string: 
        <!--beg l desc_{{object_name}}_description -->
        
        <!--end l-->

    + `title`: string: 
        <!--beg l desc_{{object_name}}_title -->
        
        <!--end l-->

    
### [PostContestProblemReply](./ObjectModelSpec.md#PostContestProblemReply)

+ type: [object](#PostContestProblemReply)

+ fields:
    
    + `code`: integer: 
        <!--beg l desc_{{object_name}}_code -->
        
        <!--end l-->

    + `data`: : 
        <!--beg l desc_{{object_name}}_data -->
        
        <!--end l-->

    
### [ListContestUsersReply](./ObjectModelSpec.md#ListContestUsersReply)

+ type: [object](#ListContestUsersReply)

+ fields:
    
    + `code`: integer: 
        <!--beg l desc_{{object_name}}_code -->
        
        <!--end l-->

    + `data`: array: 
        <!--beg l desc_{{object_name}}_data -->
        
        <!--end l-->

    
### [DeleteContestRequest](./ObjectModelSpec.md#DeleteContestRequest)

+ type: [object](#DeleteContestRequest)

+ fields:
    
    
### [DeleteContestReply](./ObjectModelSpec.md#DeleteContestReply)

+ type: [object](#DeleteContestReply)

+ fields:
    
    + `code`: integer: 
        <!--beg l desc_{{object_name}}_code -->
        
        <!--end l-->

    
### [GetContestReply](./ObjectModelSpec.md#GetContestReply)

+ type: [object](#GetContestReply)

+ fields:
    
    + `code`: integer: 
        <!--beg l desc_{{object_name}}_code -->
        
        <!--end l-->

    + `data`: : 
        <!--beg l desc_{{object_name}}_data -->
        
        <!--end l-->

    
### [PutContestRequest](./ObjectModelSpec.md#PutContestRequest)

+ type: [object](#PutContestRequest)

+ fields:
    
    + `board_frozen_duration`: integer: 
        <!--beg l desc_{{object_name}}_board_frozen_duration -->
        
        <!--end l-->

    + `config_path`: string: 
        <!--beg l desc_{{object_name}}_config_path -->
        
        <!--end l-->

    + `description`: string: 
        <!--beg l desc_{{object_name}}_description -->
        
        <!--end l-->

    + `end_duration`: integer: 
        <!--beg l desc_{{object_name}}_end_duration -->
        
        <!--end l-->

    + `role_path`: string: 
        <!--beg l desc_{{object_name}}_role_path -->
        
        <!--end l-->

    + `start_at`: string: 
        <!--beg l desc_{{object_name}}_start_at -->
        
        <!--end l-->

    + `title`: string: 
        <!--beg l desc_{{object_name}}_title -->
        
        <!--end l-->

    
### [PutContestReply](./ObjectModelSpec.md#PutContestReply)

+ type: [object](#PutContestReply)

+ fields:
    
    + `code`: integer: 
        <!--beg l desc_{{object_name}}_code -->
        
        <!--end l-->

    
### [PostContestRequest](./ObjectModelSpec.md#PostContestRequest)

+ type: [object](#PostContestRequest)

+ fields:
    
    + `board_frozen_duration`: integer: 
        <!--beg l desc_{{object_name}}_board_frozen_duration -->
        
        <!--end l-->

    + `description`: string: 
        <!--beg l desc_{{object_name}}_description -->
        
        <!--end l-->

    + `end_duration`: integer: 
        <!--beg l desc_{{object_name}}_end_duration -->
        
        <!--end l-->

    + `start_at`: string: 
        <!--beg l desc_{{object_name}}_start_at -->
        
        <!--end l-->

    + `title`: string: 
        <!--beg l desc_{{object_name}}_title -->
        
        <!--end l-->

    
### [PostContestReply](./ObjectModelSpec.md#PostContestReply)

+ type: [object](#PostContestReply)

+ fields:
    
    + `code`: integer: 
        <!--beg l desc_{{object_name}}_code -->
        
        <!--end l-->

    + `data`: : 
        <!--beg l desc_{{object_name}}_data -->
        
        <!--end l-->

    

