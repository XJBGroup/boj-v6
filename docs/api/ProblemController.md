
# Api Group Documentation (ProblemController)

<!--beg l desc_ProblemController -->

<!--end l-->

## Apis


### CountProblem

The uri/restful key of this method is `/problem-count@GET`

<!--beg l desc_CountProblem -->

<!--end l-->


+ `page`: [integer](#integer): 
    <!--beg l desc_CountProblem_page -->
    
    <!--end l-->


+ `page_size`: [integer](#integer): 
    <!--beg l desc_CountProblem_page_size -->
    
    <!--end l-->


+ `before_id`: [integer](#integer): 
    <!--beg l desc_CountProblem_before_id -->
    
    <!--end l-->


+ `order`: [string](#string): 
    <!--beg l desc_CountProblem_order -->
    
    <!--end l-->



### ListProblem

The uri/restful key of this method is `/problem-list@GET`

<!--beg l desc_ListProblem -->

<!--end l-->


+ `order`: [string](#string): 
    <!--beg l desc_ListProblem_order -->
    
    <!--end l-->


+ `page`: [integer](#integer): 
    <!--beg l desc_ListProblem_page -->
    
    <!--end l-->


+ `page_size`: [integer](#integer): 
    <!--beg l desc_ListProblem_page_size -->
    
    <!--end l-->


+ `before_id`: [integer](#integer): 
    <!--beg l desc_ListProblem_before_id -->
    
    <!--end l-->



### CountProblemDesc

The uri/restful key of this method is `/problem/{pid}/desc-count@GET`

<!--beg l desc_CountProblemDesc -->

<!--end l-->


+ `pid`: [string](#string) (required): 
    <!--beg l desc_CountProblemDesc_pid -->
    
    <!--end l-->


+ `before_id`: [integer](#integer): 
    <!--beg l desc_CountProblemDesc_before_id -->
    
    <!--end l-->


+ `order`: [string](#string): 
    <!--beg l desc_CountProblemDesc_order -->
    
    <!--end l-->


+ `page`: [integer](#integer): 
    <!--beg l desc_CountProblemDesc_page -->
    
    <!--end l-->


+ `page_size`: [integer](#integer): 
    <!--beg l desc_CountProblemDesc_page_size -->
    
    <!--end l-->



### ListProblemDesc

The uri/restful key of this method is `/problem/{pid}/desc-list@GET`

<!--beg l desc_ListProblemDesc -->

<!--end l-->


+ `pid`: [string](#string) (required): 
    <!--beg l desc_ListProblemDesc_pid -->
    
    <!--end l-->


+ `order`: [string](#string): 
    <!--beg l desc_ListProblemDesc_order -->
    
    <!--end l-->


+ `page`: [integer](#integer): 
    <!--beg l desc_ListProblemDesc_page -->
    
    <!--end l-->


+ `page_size`: [integer](#integer): 
    <!--beg l desc_ListProblemDesc_page_size -->
    
    <!--end l-->


+ `before_id`: [integer](#integer): 
    <!--beg l desc_ListProblemDesc_before_id -->
    
    <!--end l-->



### ChangeProblemDescriptionRef

The uri/restful key of this method is `/problem/{pid}/desc/ref@POST`

<!--beg l desc_ChangeProblemDescriptionRef -->

<!--end l-->


+ `pid`: [string](#string) (required): 
    <!--beg l desc_ChangeProblemDescriptionRef_pid -->
    
    <!--end l-->


+ `ChangeProblemDescriptionRefRequest`: [ChangeProblemDescriptionRefRequest](#ChangeProblemDescriptionRefRequest): 
    <!--beg l desc_ChangeProblemDescriptionRef_ChangeProblemDescriptionRefRequest -->
    
    <!--end l-->



### DeleteProblemDesc

The uri/restful key of this method is `/problem/{pid}/desc@DELETE`

<!--beg l desc_DeleteProblemDesc -->

<!--end l-->


+ `pid`: [string](#string) (required): 
    <!--beg l desc_DeleteProblemDesc_pid -->
    
    <!--end l-->


+ `DeleteProblemDescRequest`: [DeleteProblemDescRequest](#DeleteProblemDescRequest): 
    <!--beg l desc_DeleteProblemDesc_DeleteProblemDescRequest -->
    
    <!--end l-->



### GetProblemDesc

The uri/restful key of this method is `/problem/{pid}/desc@GET`

<!--beg l desc_GetProblemDesc -->

<!--end l-->


+ `pid`: [string](#string) (required): 
    <!--beg l desc_GetProblemDesc_pid -->
    
    <!--end l-->


+ `name`: [string](#string): 
    <!--beg l desc_GetProblemDesc_name -->
    
    <!--end l-->



### PostProblemDesc

The uri/restful key of this method is `/problem/{pid}/desc@POST`

<!--beg l desc_PostProblemDesc -->

<!--end l-->


+ `pid`: [string](#string) (required): 
    <!--beg l desc_PostProblemDesc_pid -->
    
    <!--end l-->


+ `PostProblemDescRequest`: [PostProblemDescRequest](#PostProblemDescRequest): 
    <!--beg l desc_PostProblemDesc_PostProblemDescRequest -->
    
    <!--end l-->



### PutProblemDesc

The uri/restful key of this method is `/problem/{pid}/desc@PUT`

<!--beg l desc_PutProblemDesc -->

<!--end l-->


+ `pid`: [string](#string) (required): 
    <!--beg l desc_PutProblemDesc_pid -->
    
    <!--end l-->


+ `PutProblemDescRequest`: [PutProblemDescRequest](#PutProblemDescRequest): 
    <!--beg l desc_PutProblemDesc_PutProblemDescRequest -->
    
    <!--end l-->



### ProblemFSReadConfig

The uri/restful key of this method is `/problem/{pid}/fs/config@GET`

<!--beg l desc_ProblemFSReadConfig -->

<!--end l-->


+ `pid`: [string](#string) (required): 
    <!--beg l desc_ProblemFSReadConfig_pid -->
    
    <!--end l-->


+ `path`: [string](#string): 
    <!--beg l desc_ProblemFSReadConfig_path -->
    
    <!--end l-->



### ProblemFSWriteConfig

The uri/restful key of this method is `/problem/{pid}/fs/config@POST`

<!--beg l desc_ProblemFSWriteConfig -->

<!--end l-->


+ `pid`: [string](#string) (required): 
    <!--beg l desc_ProblemFSWriteConfig_pid -->
    
    <!--end l-->


+ `ProblemFSWriteConfigRequest`: [ProblemFSWriteConfigRequest](#ProblemFSWriteConfigRequest): 
    <!--beg l desc_ProblemFSWriteConfig_ProblemFSWriteConfigRequest -->
    
    <!--end l-->



### ProblemFSPutConfig

The uri/restful key of this method is `/problem/{pid}/fs/config@PUT`

<!--beg l desc_ProblemFSPutConfig -->

<!--end l-->


+ `pid`: [string](#string) (required): 
    <!--beg l desc_ProblemFSPutConfig_pid -->
    
    <!--end l-->


+ `ProblemFSPutConfigRequest`: [ProblemFSPutConfigRequest](#ProblemFSPutConfigRequest): 
    <!--beg l desc_ProblemFSPutConfig_ProblemFSPutConfigRequest -->
    
    <!--end l-->



### ProblemFSZipRead

The uri/restful key of this method is `/problem/{pid}/fs/directory/zip@GET`

<!--beg l desc_ProblemFSZipRead -->

<!--end l-->


+ `pid`: [string](#string) (required): 
    <!--beg l desc_ProblemFSZipRead_pid -->
    
    <!--end l-->


+ `path`: [string](#string): 
    <!--beg l desc_ProblemFSZipRead_path -->
    
    <!--end l-->



### ProblemFSZipWrite

The uri/restful key of this method is `/problem/{pid}/fs/directory/zip@POST`

<!--beg l desc_ProblemFSZipWrite -->

<!--end l-->


+ `pid`: [string](#string) (required): 
    <!--beg l desc_ProblemFSZipWrite_pid -->
    
    <!--end l-->


+ `ProblemFSZipWriteRequest`: [ProblemFSZipWriteRequest](#ProblemFSZipWriteRequest): 
    <!--beg l desc_ProblemFSZipWrite_ProblemFSZipWriteRequest -->
    
    <!--end l-->



### ProblemFSRemoveAll

The uri/restful key of this method is `/problem/{pid}/fs/directory@DELETE`

<!--beg l desc_ProblemFSRemoveAll -->

<!--end l-->


+ `pid`: [string](#string) (required): 
    <!--beg l desc_ProblemFSRemoveAll_pid -->
    
    <!--end l-->


+ `ProblemFSRemoveAllRequest`: [ProblemFSRemoveAllRequest](#ProblemFSRemoveAllRequest): 
    <!--beg l desc_ProblemFSRemoveAll_ProblemFSRemoveAllRequest -->
    
    <!--end l-->



### ProblemFSLS

The uri/restful key of this method is `/problem/{pid}/fs/directory@GET`

<!--beg l desc_ProblemFSLS -->

<!--end l-->


+ `pid`: [string](#string) (required): 
    <!--beg l desc_ProblemFSLS_pid -->
    
    <!--end l-->


+ `path`: [string](#string): 
    <!--beg l desc_ProblemFSLS_path -->
    
    <!--end l-->



### ProblemFSWrites

The uri/restful key of this method is `/problem/{pid}/fs/directory@POST`

<!--beg l desc_ProblemFSWrites -->

<!--end l-->


+ `pid`: [string](#string) (required): 
    <!--beg l desc_ProblemFSWrites_pid -->
    
    <!--end l-->


+ `ProblemFSWritesRequest`: [ProblemFSWritesRequest](#ProblemFSWritesRequest): 
    <!--beg l desc_ProblemFSWrites_ProblemFSWritesRequest -->
    
    <!--end l-->



### ProblemFSMkdir

The uri/restful key of this method is `/problem/{pid}/fs/directory@PUT`

<!--beg l desc_ProblemFSMkdir -->

<!--end l-->


+ `pid`: [string](#string) (required): 
    <!--beg l desc_ProblemFSMkdir_pid -->
    
    <!--end l-->


+ `ProblemFSMkdirRequest`: [ProblemFSMkdirRequest](#ProblemFSMkdirRequest): 
    <!--beg l desc_ProblemFSMkdir_ProblemFSMkdirRequest -->
    
    <!--end l-->



### ProblemFSRead

The uri/restful key of this method is `/problem/{pid}/fs/file/content@GET`

<!--beg l desc_ProblemFSRead -->

<!--end l-->


+ `pid`: [string](#string) (required): 
    <!--beg l desc_ProblemFSRead_pid -->
    
    <!--end l-->


+ `path`: [string](#string): 
    <!--beg l desc_ProblemFSRead_path -->
    
    <!--end l-->



### ProblemFSRemove

The uri/restful key of this method is `/problem/{pid}/fs/file@DELETE`

<!--beg l desc_ProblemFSRemove -->

<!--end l-->


+ `pid`: [string](#string) (required): 
    <!--beg l desc_ProblemFSRemove_pid -->
    
    <!--end l-->


+ `ProblemFSRemoveRequest`: [ProblemFSRemoveRequest](#ProblemFSRemoveRequest): 
    <!--beg l desc_ProblemFSRemove_ProblemFSRemoveRequest -->
    
    <!--end l-->



### ProblemFSStat

The uri/restful key of this method is `/problem/{pid}/fs/file@GET`

<!--beg l desc_ProblemFSStat -->

<!--end l-->


+ `pid`: [string](#string) (required): 
    <!--beg l desc_ProblemFSStat_pid -->
    
    <!--end l-->


+ `path`: [string](#string): 
    <!--beg l desc_ProblemFSStat_path -->
    
    <!--end l-->



### ProblemFSWrite

The uri/restful key of this method is `/problem/{pid}/fs/file@POST`

<!--beg l desc_ProblemFSWrite -->

<!--end l-->


+ `pid`: [string](#string) (required): 
    <!--beg l desc_ProblemFSWrite_pid -->
    
    <!--end l-->


+ `ProblemFSWriteRequest`: [ProblemFSWriteRequest](#ProblemFSWriteRequest): 
    <!--beg l desc_ProblemFSWrite_ProblemFSWriteRequest -->
    
    <!--end l-->



### DeleteProblem

The uri/restful key of this method is `/problem/{pid}@DELETE`

<!--beg l desc_DeleteProblem -->

<!--end l-->


+ `pid`: [string](#string) (required): 
    <!--beg l desc_DeleteProblem_pid -->
    
    <!--end l-->


+ `DeleteProblemRequest`: [DeleteProblemRequest](#DeleteProblemRequest): 
    <!--beg l desc_DeleteProblem_DeleteProblemRequest -->
    
    <!--end l-->



### GetProblem

The uri/restful key of this method is `/problem/{pid}@GET`

<!--beg l desc_GetProblem -->

<!--end l-->


+ `pid`: [string](#string) (required): 
    <!--beg l desc_GetProblem_pid -->
    
    <!--end l-->



### PutProblem

The uri/restful key of this method is `/problem/{pid}@PUT`

<!--beg l desc_PutProblem -->

<!--end l-->


+ `pid`: [string](#string) (required): 
    <!--beg l desc_PutProblem_pid -->
    
    <!--end l-->


+ `PutProblemRequest`: [PutProblemRequest](#PutProblemRequest): 
    <!--beg l desc_PutProblem_PutProblemRequest -->
    
    <!--end l-->



### PostProblem

The uri/restful key of this method is `/problem@POST`

<!--beg l desc_PostProblem -->

<!--end l-->


+ `PostProblemRequest`: [PostProblemRequest](#PostProblemRequest): 
    <!--beg l desc_PostProblem_PostProblemRequest -->
    
    <!--end l-->



## Local Object Reference




### [ChangeProblemDescriptionRefRequest](./ObjectModelSpec.md#ChangeProblemDescriptionRefRequest)

+ type: [object](#ChangeProblemDescriptionRefRequest)

+ fields:
    
    + `name`: string: 
        <!--beg l desc_{{object_name}}_name -->
        
        <!--end l-->

    + `new_name`: string: 
        <!--beg l desc_{{object_name}}_new_name -->
        
        <!--end l-->

    
### [DeleteProblemDescRequest](./ObjectModelSpec.md#DeleteProblemDescRequest)

+ type: [object](#DeleteProblemDescRequest)

+ fields:
    
    + `name`: string: 
        <!--beg l desc_{{object_name}}_name -->
        
        <!--end l-->

    
### [PostProblemDescRequest](./ObjectModelSpec.md#PostProblemDescRequest)

+ type: [object](#PostProblemDescRequest)

+ fields:
    
    + `content`: string: 
        <!--beg l desc_{{object_name}}_content -->
        
        <!--end l-->

    + `name`: string: 
        <!--beg l desc_{{object_name}}_name -->
        
        <!--end l-->

    
### [PutProblemDescRequest](./ObjectModelSpec.md#PutProblemDescRequest)

+ type: [object](#PutProblemDescRequest)

+ fields:
    
    + `content`: string: 
        <!--beg l desc_{{object_name}}_content -->
        
        <!--end l-->

    + `name`: string: 
        <!--beg l desc_{{object_name}}_name -->
        
        <!--end l-->

    
### [ProblemFSWriteConfigRequest](./ObjectModelSpec.md#ProblemFSWriteConfigRequest)

+ type: [object](#ProblemFSWriteConfigRequest)

+ fields:
    
    + `path`: string: 
        <!--beg l desc_{{object_name}}_path -->
        
        <!--end l-->

    
### [ProblemFSPutConfigRequest](./ObjectModelSpec.md#ProblemFSPutConfigRequest)

+ type: [object](#ProblemFSPutConfigRequest)

+ fields:
    
    + `key`: string: 
        <!--beg l desc_{{object_name}}_key -->
        
        <!--end l-->

    + `path`: string: 
        <!--beg l desc_{{object_name}}_path -->
        
        <!--end l-->

    + `value`: array: 
        <!--beg l desc_{{object_name}}_value -->
        
        <!--end l-->

    
### [ProblemFSZipWriteRequest](./ObjectModelSpec.md#ProblemFSZipWriteRequest)

+ type: [object](#ProblemFSZipWriteRequest)

+ fields:
    
    + `path`: string: 
        <!--beg l desc_{{object_name}}_path -->
        
        <!--end l-->

    
### [ProblemFSRemoveAllRequest](./ObjectModelSpec.md#ProblemFSRemoveAllRequest)

+ type: [object](#ProblemFSRemoveAllRequest)

+ fields:
    
    + `path`: string: 
        <!--beg l desc_{{object_name}}_path -->
        
        <!--end l-->

    
### [ProblemFSWritesRequest](./ObjectModelSpec.md#ProblemFSWritesRequest)

+ type: [object](#ProblemFSWritesRequest)

+ fields:
    
    + `path`: string: 
        <!--beg l desc_{{object_name}}_path -->
        
        <!--end l-->

    
### [ProblemFSMkdirRequest](./ObjectModelSpec.md#ProblemFSMkdirRequest)

+ type: [object](#ProblemFSMkdirRequest)

+ fields:
    
    + `path`: string: 
        <!--beg l desc_{{object_name}}_path -->
        
        <!--end l-->

    
### [ProblemFSRemoveRequest](./ObjectModelSpec.md#ProblemFSRemoveRequest)

+ type: [object](#ProblemFSRemoveRequest)

+ fields:
    
    + `path`: string: 
        <!--beg l desc_{{object_name}}_path -->
        
        <!--end l-->

    
### [ProblemFSWriteRequest](./ObjectModelSpec.md#ProblemFSWriteRequest)

+ type: [object](#ProblemFSWriteRequest)

+ fields:
    
    + `path`: string: 
        <!--beg l desc_{{object_name}}_path -->
        
        <!--end l-->

    
### [DeleteProblemRequest](./ObjectModelSpec.md#DeleteProblemRequest)

+ type: [object](#DeleteProblemRequest)

+ fields:
    
    
### [PutProblemRequest](./ObjectModelSpec.md#PutProblemRequest)

+ type: [object](#PutProblemRequest)

+ fields:
    
    + `description_ref`: string: 
        <!--beg l desc_{{object_name}}_description_ref -->
        
        <!--end l-->

    + `title`: string: 
        <!--beg l desc_{{object_name}}_title -->
        
        <!--end l-->

    
### [PostProblemRequest](./ObjectModelSpec.md#PostProblemRequest)

+ type: [object](#PostProblemRequest)

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

    

