
# Api Group Documentation (ContestController)

<!--beg l desc_ContestController -->

<!--end l-->

## Apis


### CountContest

The uri/restful key of this method is `/contest-count@GET`

<!--beg l desc_CountContest -->

<!--end l-->


+ `before_id`: [integer](#integer): 
    <!--beg l desc_CountContest_before_id -->
    
    <!--end l-->


+ `order`: [string](#string): 
    <!--beg l desc_CountContest_order -->
    
    <!--end l-->


+ `page`: [integer](#integer): 
    <!--beg l desc_CountContest_page -->
    
    <!--end l-->


+ `page_size`: [integer](#integer): 
    <!--beg l desc_CountContest_page_size -->
    
    <!--end l-->



### ListContest

The uri/restful key of this method is `/contest-list@GET`

<!--beg l desc_ListContest -->

<!--end l-->


+ `order`: [string](#string): 
    <!--beg l desc_ListContest_order -->
    
    <!--end l-->


+ `page`: [integer](#integer): 
    <!--beg l desc_ListContest_page -->
    
    <!--end l-->


+ `page_size`: [integer](#integer): 
    <!--beg l desc_ListContest_page_size -->
    
    <!--end l-->


+ `before_id`: [integer](#integer): 
    <!--beg l desc_ListContest_before_id -->
    
    <!--end l-->



### CountContestProblem

The uri/restful key of this method is `/contest/{cid}/problem-count@GET`

<!--beg l desc_CountContestProblem -->

<!--end l-->


+ `cid`: [string](#string) (required): 
    <!--beg l desc_CountContestProblem_cid -->
    
    <!--end l-->


+ `before_id`: [integer](#integer): 
    <!--beg l desc_CountContestProblem_before_id -->
    
    <!--end l-->


+ `order`: [string](#string): 
    <!--beg l desc_CountContestProblem_order -->
    
    <!--end l-->


+ `page`: [integer](#integer): 
    <!--beg l desc_CountContestProblem_page -->
    
    <!--end l-->


+ `page_size`: [integer](#integer): 
    <!--beg l desc_CountContestProblem_page_size -->
    
    <!--end l-->



### ListContestProblem

The uri/restful key of this method is `/contest/{cid}/problem-list@GET`

<!--beg l desc_ListContestProblem -->

<!--end l-->


+ `cid`: [string](#string) (required): 
    <!--beg l desc_ListContestProblem_cid -->
    
    <!--end l-->


+ `order`: [string](#string): 
    <!--beg l desc_ListContestProblem_order -->
    
    <!--end l-->


+ `page`: [integer](#integer): 
    <!--beg l desc_ListContestProblem_page -->
    
    <!--end l-->


+ `page_size`: [integer](#integer): 
    <!--beg l desc_ListContestProblem_page_size -->
    
    <!--end l-->


+ `before_id`: [integer](#integer): 
    <!--beg l desc_ListContestProblem_before_id -->
    
    <!--end l-->



### CountContestProblemDesc

The uri/restful key of this method is `/contest/{cid}/problem/{pid}/desc-count@GET`

<!--beg l desc_CountContestProblemDesc -->

<!--end l-->


+ `cid`: [string](#string) (required): 
    <!--beg l desc_CountContestProblemDesc_cid -->
    
    <!--end l-->


+ `pid`: [string](#string) (required): 
    <!--beg l desc_CountContestProblemDesc_pid -->
    
    <!--end l-->


+ `order`: [string](#string): 
    <!--beg l desc_CountContestProblemDesc_order -->
    
    <!--end l-->


+ `page`: [integer](#integer): 
    <!--beg l desc_CountContestProblemDesc_page -->
    
    <!--end l-->


+ `page_size`: [integer](#integer): 
    <!--beg l desc_CountContestProblemDesc_page_size -->
    
    <!--end l-->


+ `before_id`: [integer](#integer): 
    <!--beg l desc_CountContestProblemDesc_before_id -->
    
    <!--end l-->



### ListContestProblemDesc

The uri/restful key of this method is `/contest/{cid}/problem/{pid}/desc-list@GET`

<!--beg l desc_ListContestProblemDesc -->

<!--end l-->


+ `cid`: [string](#string) (required): 
    <!--beg l desc_ListContestProblemDesc_cid -->
    
    <!--end l-->


+ `pid`: [string](#string) (required): 
    <!--beg l desc_ListContestProblemDesc_pid -->
    
    <!--end l-->


+ `order`: [string](#string): 
    <!--beg l desc_ListContestProblemDesc_order -->
    
    <!--end l-->


+ `page`: [integer](#integer): 
    <!--beg l desc_ListContestProblemDesc_page -->
    
    <!--end l-->


+ `page_size`: [integer](#integer): 
    <!--beg l desc_ListContestProblemDesc_page_size -->
    
    <!--end l-->


+ `before_id`: [integer](#integer): 
    <!--beg l desc_ListContestProblemDesc_before_id -->
    
    <!--end l-->



### ChangeContestProblemDescriptionRef

The uri/restful key of this method is `/contest/{cid}/problem/{pid}/desc/ref@POST`

<!--beg l desc_ChangeContestProblemDescriptionRef -->

<!--end l-->


+ `cid`: [string](#string) (required): 
    <!--beg l desc_ChangeContestProblemDescriptionRef_cid -->
    
    <!--end l-->


+ `pid`: [string](#string) (required): 
    <!--beg l desc_ChangeContestProblemDescriptionRef_pid -->
    
    <!--end l-->


+ `ChangeContestProblemDescriptionRefRequest`: [ChangeContestProblemDescriptionRefRequest](#ChangeContestProblemDescriptionRefRequest): 
    <!--beg l desc_ChangeContestProblemDescriptionRef_ChangeContestProblemDescriptionRefRequest -->
    
    <!--end l-->



### DeleteContestProblemDesc

The uri/restful key of this method is `/contest/{cid}/problem/{pid}/desc@DELETE`

<!--beg l desc_DeleteContestProblemDesc -->

<!--end l-->


+ `cid`: [string](#string) (required): 
    <!--beg l desc_DeleteContestProblemDesc_cid -->
    
    <!--end l-->


+ `pid`: [string](#string) (required): 
    <!--beg l desc_DeleteContestProblemDesc_pid -->
    
    <!--end l-->


+ `DeleteContestProblemDescRequest`: [DeleteContestProblemDescRequest](#DeleteContestProblemDescRequest): 
    <!--beg l desc_DeleteContestProblemDesc_DeleteContestProblemDescRequest -->
    
    <!--end l-->



### GetContestProblemDesc

The uri/restful key of this method is `/contest/{cid}/problem/{pid}/desc@GET`

<!--beg l desc_GetContestProblemDesc -->

<!--end l-->


+ `cid`: [string](#string) (required): 
    <!--beg l desc_GetContestProblemDesc_cid -->
    
    <!--end l-->


+ `pid`: [string](#string) (required): 
    <!--beg l desc_GetContestProblemDesc_pid -->
    
    <!--end l-->


+ `name`: [string](#string): 
    <!--beg l desc_GetContestProblemDesc_name -->
    
    <!--end l-->



### PostContestProblemDesc

The uri/restful key of this method is `/contest/{cid}/problem/{pid}/desc@POST`

<!--beg l desc_PostContestProblemDesc -->

<!--end l-->


+ `cid`: [string](#string) (required): 
    <!--beg l desc_PostContestProblemDesc_cid -->
    
    <!--end l-->


+ `pid`: [string](#string) (required): 
    <!--beg l desc_PostContestProblemDesc_pid -->
    
    <!--end l-->


+ `PostContestProblemDescRequest`: [PostContestProblemDescRequest](#PostContestProblemDescRequest): 
    <!--beg l desc_PostContestProblemDesc_PostContestProblemDescRequest -->
    
    <!--end l-->



### PutContestProblemDesc

The uri/restful key of this method is `/contest/{cid}/problem/{pid}/desc@PUT`

<!--beg l desc_PutContestProblemDesc -->

<!--end l-->


+ `cid`: [string](#string) (required): 
    <!--beg l desc_PutContestProblemDesc_cid -->
    
    <!--end l-->


+ `pid`: [string](#string) (required): 
    <!--beg l desc_PutContestProblemDesc_pid -->
    
    <!--end l-->


+ `PutContestProblemDescRequest`: [PutContestProblemDescRequest](#PutContestProblemDescRequest): 
    <!--beg l desc_PutContestProblemDesc_PutContestProblemDescRequest -->
    
    <!--end l-->



### DeleteContestProblem

The uri/restful key of this method is `/contest/{cid}/problem/{pid}@DELETE`

<!--beg l desc_DeleteContestProblem -->

<!--end l-->


+ `cid`: [string](#string) (required): 
    <!--beg l desc_DeleteContestProblem_cid -->
    
    <!--end l-->


+ `pid`: [string](#string) (required): 
    <!--beg l desc_DeleteContestProblem_pid -->
    
    <!--end l-->


+ `DeleteContestProblemRequest`: [DeleteContestProblemRequest](#DeleteContestProblemRequest): 
    <!--beg l desc_DeleteContestProblem_DeleteContestProblemRequest -->
    
    <!--end l-->



### GetContestProblem

The uri/restful key of this method is `/contest/{cid}/problem/{pid}@GET`

<!--beg l desc_GetContestProblem -->

<!--end l-->


+ `cid`: [string](#string) (required): 
    <!--beg l desc_GetContestProblem_cid -->
    
    <!--end l-->


+ `pid`: [string](#string) (required): 
    <!--beg l desc_GetContestProblem_pid -->
    
    <!--end l-->



### PutContestProblem

The uri/restful key of this method is `/contest/{cid}/problem/{pid}@PUT`

<!--beg l desc_PutContestProblem -->

<!--end l-->


+ `cid`: [string](#string) (required): 
    <!--beg l desc_PutContestProblem_cid -->
    
    <!--end l-->


+ `pid`: [string](#string) (required): 
    <!--beg l desc_PutContestProblem_pid -->
    
    <!--end l-->


+ `PutContestProblemRequest`: [PutContestProblemRequest](#PutContestProblemRequest): 
    <!--beg l desc_PutContestProblem_PutContestProblemRequest -->
    
    <!--end l-->



### PostContestProblem

The uri/restful key of this method is `/contest/{cid}/problem@POST`

<!--beg l desc_PostContestProblem -->

<!--end l-->


+ `cid`: [string](#string) (required): 
    <!--beg l desc_PostContestProblem_cid -->
    
    <!--end l-->


+ `PostContestProblemRequest`: [PostContestProblemRequest](#PostContestProblemRequest): 
    <!--beg l desc_PostContestProblem_PostContestProblemRequest -->
    
    <!--end l-->



### ListContestUsers

The uri/restful key of this method is `/contest/{cid}/user-list@GET`

<!--beg l desc_ListContestUsers -->

<!--end l-->


+ `cid`: [string](#string) (required): 
    <!--beg l desc_ListContestUsers_cid -->
    
    <!--end l-->



### DeleteContest

The uri/restful key of this method is `/contest/{cid}@DELETE`

<!--beg l desc_DeleteContest -->

<!--end l-->


+ `cid`: [string](#string) (required): 
    <!--beg l desc_DeleteContest_cid -->
    
    <!--end l-->


+ `DeleteContestRequest`: [DeleteContestRequest](#DeleteContestRequest): 
    <!--beg l desc_DeleteContest_DeleteContestRequest -->
    
    <!--end l-->



### GetContest

The uri/restful key of this method is `/contest/{cid}@GET`

<!--beg l desc_GetContest -->

<!--end l-->


+ `cid`: [string](#string) (required): 
    <!--beg l desc_GetContest_cid -->
    
    <!--end l-->



### PutContest

The uri/restful key of this method is `/contest/{cid}@PUT`

<!--beg l desc_PutContest -->

<!--end l-->


+ `cid`: [string](#string) (required): 
    <!--beg l desc_PutContest_cid -->
    
    <!--end l-->


+ `PutContestRequest`: [PutContestRequest](#PutContestRequest): 
    <!--beg l desc_PutContest_PutContestRequest -->
    
    <!--end l-->



### PostContest

The uri/restful key of this method is `/contest@POST`

<!--beg l desc_PostContest -->

<!--end l-->


+ `PostContestRequest`: [PostContestRequest](#PostContestRequest): 
    <!--beg l desc_PostContest_PostContestRequest -->
    
    <!--end l-->



## Local Object Reference




### [ChangeContestProblemDescriptionRefRequest](./ObjectModelSpec.md#ChangeContestProblemDescriptionRefRequest)

+ type: [object](#ChangeContestProblemDescriptionRefRequest)

+ fields:
    
    + `name`: string: 
        <!--beg l desc_{{object_name}}_name -->
        
        <!--end l-->

    + `new_name`: string: 
        <!--beg l desc_{{object_name}}_new_name -->
        
        <!--end l-->

    
### [DeleteContestProblemDescRequest](./ObjectModelSpec.md#DeleteContestProblemDescRequest)

+ type: [object](#DeleteContestProblemDescRequest)

+ fields:
    
    + `name`: string: 
        <!--beg l desc_{{object_name}}_name -->
        
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

    
### [PutContestProblemDescRequest](./ObjectModelSpec.md#PutContestProblemDescRequest)

+ type: [object](#PutContestProblemDescRequest)

+ fields:
    
    + `content`: string: 
        <!--beg l desc_{{object_name}}_content -->
        
        <!--end l-->

    + `name`: string: 
        <!--beg l desc_{{object_name}}_name -->
        
        <!--end l-->

    
### [DeleteContestProblemRequest](./ObjectModelSpec.md#DeleteContestProblemRequest)

+ type: [object](#DeleteContestProblemRequest)

+ fields:
    
    
### [PutContestProblemRequest](./ObjectModelSpec.md#PutContestProblemRequest)

+ type: [object](#PutContestProblemRequest)

+ fields:
    
    + `description_ref`: string: 
        <!--beg l desc_{{object_name}}_description_ref -->
        
        <!--end l-->

    + `title`: string: 
        <!--beg l desc_{{object_name}}_title -->
        
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

    
### [DeleteContestRequest](./ObjectModelSpec.md#DeleteContestRequest)

+ type: [object](#DeleteContestRequest)

+ fields:
    
    
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

    

