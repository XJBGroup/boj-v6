
# Api Group Documentation (ProblemController)

<!--beg l desc_ProblemController -->

<!--end l-->


## CountProblem

The uri/restful key of this method is `/problem-count@GET`

<!--beg l desc_CountProblem -->

<!--end l-->


+ `page`: `integer`: 
    <!--beg l desc_CountProblem_page -->
    
    <!--end l-->

+ `page_size`: `integer`: 
    <!--beg l desc_CountProblem_page_size -->
    
    <!--end l-->

+ `before_id`: `integer`: 
    <!--beg l desc_CountProblem_before_id -->
    
    <!--end l-->

+ `order`: `string`: 
    <!--beg l desc_CountProblem_order -->
    
    <!--end l-->



## ListProblem

The uri/restful key of this method is `/problem-list@GET`

<!--beg l desc_ListProblem -->

<!--end l-->


+ `order`: `string`: 
    <!--beg l desc_ListProblem_order -->
    
    <!--end l-->

+ `page`: `integer`: 
    <!--beg l desc_ListProblem_page -->
    
    <!--end l-->

+ `page_size`: `integer`: 
    <!--beg l desc_ListProblem_page_size -->
    
    <!--end l-->

+ `before_id`: `integer`: 
    <!--beg l desc_ListProblem_before_id -->
    
    <!--end l-->



## CountProblemDesc

The uri/restful key of this method is `/problem/{pid}/desc-count@GET`

<!--beg l desc_CountProblemDesc -->

<!--end l-->


+ `pid`: `string` (required): 
    <!--beg l desc_CountProblemDesc_pid -->
    
    <!--end l-->

+ `before_id`: `integer`: 
    <!--beg l desc_CountProblemDesc_before_id -->
    
    <!--end l-->

+ `order`: `string`: 
    <!--beg l desc_CountProblemDesc_order -->
    
    <!--end l-->

+ `page`: `integer`: 
    <!--beg l desc_CountProblemDesc_page -->
    
    <!--end l-->

+ `page_size`: `integer`: 
    <!--beg l desc_CountProblemDesc_page_size -->
    
    <!--end l-->



## ListProblemDesc

The uri/restful key of this method is `/problem/{pid}/desc-list@GET`

<!--beg l desc_ListProblemDesc -->

<!--end l-->


+ `pid`: `string` (required): 
    <!--beg l desc_ListProblemDesc_pid -->
    
    <!--end l-->

+ `order`: `string`: 
    <!--beg l desc_ListProblemDesc_order -->
    
    <!--end l-->

+ `page`: `integer`: 
    <!--beg l desc_ListProblemDesc_page -->
    
    <!--end l-->

+ `page_size`: `integer`: 
    <!--beg l desc_ListProblemDesc_page_size -->
    
    <!--end l-->

+ `before_id`: `integer`: 
    <!--beg l desc_ListProblemDesc_before_id -->
    
    <!--end l-->



## ChangeProblemDescriptionRef

The uri/restful key of this method is `/problem/{pid}/desc/ref@POST`

<!--beg l desc_ChangeProblemDescriptionRef -->

<!--end l-->


+ `pid`: `string` (required): 
    <!--beg l desc_ChangeProblemDescriptionRef_pid -->
    
    <!--end l-->

+ `ChangeProblemDescriptionRefRequest`: `any`: 
    <!--beg l desc_ChangeProblemDescriptionRef_ChangeProblemDescriptionRefRequest -->
    
    <!--end l-->



## DeleteProblemDesc

The uri/restful key of this method is `/problem/{pid}/desc@DELETE`

<!--beg l desc_DeleteProblemDesc -->

<!--end l-->


+ `pid`: `string` (required): 
    <!--beg l desc_DeleteProblemDesc_pid -->
    
    <!--end l-->

+ `DeleteProblemDescRequest`: `any`: 
    <!--beg l desc_DeleteProblemDesc_DeleteProblemDescRequest -->
    
    <!--end l-->



## GetProblemDesc

The uri/restful key of this method is `/problem/{pid}/desc@GET`

<!--beg l desc_GetProblemDesc -->

<!--end l-->


+ `pid`: `string` (required): 
    <!--beg l desc_GetProblemDesc_pid -->
    
    <!--end l-->

+ `name`: `string`: 
    <!--beg l desc_GetProblemDesc_name -->
    
    <!--end l-->



## PostProblemDesc

The uri/restful key of this method is `/problem/{pid}/desc@POST`

<!--beg l desc_PostProblemDesc -->

<!--end l-->


+ `pid`: `string` (required): 
    <!--beg l desc_PostProblemDesc_pid -->
    
    <!--end l-->

+ `PostProblemDescRequest`: `any`: 
    <!--beg l desc_PostProblemDesc_PostProblemDescRequest -->
    
    <!--end l-->



## PutProblemDesc

The uri/restful key of this method is `/problem/{pid}/desc@PUT`

<!--beg l desc_PutProblemDesc -->

<!--end l-->


+ `pid`: `string` (required): 
    <!--beg l desc_PutProblemDesc_pid -->
    
    <!--end l-->

+ `PutProblemDescRequest`: `any`: 
    <!--beg l desc_PutProblemDesc_PutProblemDescRequest -->
    
    <!--end l-->



## ProblemFSReadConfig

The uri/restful key of this method is `/problem/{pid}/fs/config@GET`

<!--beg l desc_ProblemFSReadConfig -->

<!--end l-->


+ `pid`: `string` (required): 
    <!--beg l desc_ProblemFSReadConfig_pid -->
    
    <!--end l-->

+ `path`: `string`: 
    <!--beg l desc_ProblemFSReadConfig_path -->
    
    <!--end l-->



## ProblemFSWriteConfig

The uri/restful key of this method is `/problem/{pid}/fs/config@POST`

<!--beg l desc_ProblemFSWriteConfig -->

<!--end l-->


+ `pid`: `string` (required): 
    <!--beg l desc_ProblemFSWriteConfig_pid -->
    
    <!--end l-->

+ `ProblemFSWriteConfigRequest`: `any`: 
    <!--beg l desc_ProblemFSWriteConfig_ProblemFSWriteConfigRequest -->
    
    <!--end l-->



## ProblemFSPutConfig

The uri/restful key of this method is `/problem/{pid}/fs/config@PUT`

<!--beg l desc_ProblemFSPutConfig -->

<!--end l-->


+ `pid`: `string` (required): 
    <!--beg l desc_ProblemFSPutConfig_pid -->
    
    <!--end l-->

+ `ProblemFSPutConfigRequest`: `any`: 
    <!--beg l desc_ProblemFSPutConfig_ProblemFSPutConfigRequest -->
    
    <!--end l-->



## ProblemFSZipRead

The uri/restful key of this method is `/problem/{pid}/fs/directory/zip@GET`

<!--beg l desc_ProblemFSZipRead -->

<!--end l-->


+ `pid`: `string` (required): 
    <!--beg l desc_ProblemFSZipRead_pid -->
    
    <!--end l-->

+ `path`: `string`: 
    <!--beg l desc_ProblemFSZipRead_path -->
    
    <!--end l-->



## ProblemFSZipWrite

The uri/restful key of this method is `/problem/{pid}/fs/directory/zip@POST`

<!--beg l desc_ProblemFSZipWrite -->

<!--end l-->


+ `pid`: `string` (required): 
    <!--beg l desc_ProblemFSZipWrite_pid -->
    
    <!--end l-->

+ `ProblemFSZipWriteRequest`: `any`: 
    <!--beg l desc_ProblemFSZipWrite_ProblemFSZipWriteRequest -->
    
    <!--end l-->



## ProblemFSRemoveAll

The uri/restful key of this method is `/problem/{pid}/fs/directory@DELETE`

<!--beg l desc_ProblemFSRemoveAll -->

<!--end l-->


+ `pid`: `string` (required): 
    <!--beg l desc_ProblemFSRemoveAll_pid -->
    
    <!--end l-->

+ `ProblemFSRemoveAllRequest`: `any`: 
    <!--beg l desc_ProblemFSRemoveAll_ProblemFSRemoveAllRequest -->
    
    <!--end l-->



## ProblemFSLS

The uri/restful key of this method is `/problem/{pid}/fs/directory@GET`

<!--beg l desc_ProblemFSLS -->

<!--end l-->


+ `pid`: `string` (required): 
    <!--beg l desc_ProblemFSLS_pid -->
    
    <!--end l-->

+ `path`: `string`: 
    <!--beg l desc_ProblemFSLS_path -->
    
    <!--end l-->



## ProblemFSWrites

The uri/restful key of this method is `/problem/{pid}/fs/directory@POST`

<!--beg l desc_ProblemFSWrites -->

<!--end l-->


+ `pid`: `string` (required): 
    <!--beg l desc_ProblemFSWrites_pid -->
    
    <!--end l-->

+ `ProblemFSWritesRequest`: `any`: 
    <!--beg l desc_ProblemFSWrites_ProblemFSWritesRequest -->
    
    <!--end l-->



## ProblemFSMkdir

The uri/restful key of this method is `/problem/{pid}/fs/directory@PUT`

<!--beg l desc_ProblemFSMkdir -->

<!--end l-->


+ `pid`: `string` (required): 
    <!--beg l desc_ProblemFSMkdir_pid -->
    
    <!--end l-->

+ `ProblemFSMkdirRequest`: `any`: 
    <!--beg l desc_ProblemFSMkdir_ProblemFSMkdirRequest -->
    
    <!--end l-->



## ProblemFSRead

The uri/restful key of this method is `/problem/{pid}/fs/file/content@GET`

<!--beg l desc_ProblemFSRead -->

<!--end l-->


+ `pid`: `string` (required): 
    <!--beg l desc_ProblemFSRead_pid -->
    
    <!--end l-->

+ `path`: `string`: 
    <!--beg l desc_ProblemFSRead_path -->
    
    <!--end l-->



## ProblemFSRemove

The uri/restful key of this method is `/problem/{pid}/fs/file@DELETE`

<!--beg l desc_ProblemFSRemove -->

<!--end l-->


+ `pid`: `string` (required): 
    <!--beg l desc_ProblemFSRemove_pid -->
    
    <!--end l-->

+ `ProblemFSRemoveRequest`: `any`: 
    <!--beg l desc_ProblemFSRemove_ProblemFSRemoveRequest -->
    
    <!--end l-->



## ProblemFSStat

The uri/restful key of this method is `/problem/{pid}/fs/file@GET`

<!--beg l desc_ProblemFSStat -->

<!--end l-->


+ `pid`: `string` (required): 
    <!--beg l desc_ProblemFSStat_pid -->
    
    <!--end l-->

+ `path`: `string`: 
    <!--beg l desc_ProblemFSStat_path -->
    
    <!--end l-->



## ProblemFSWrite

The uri/restful key of this method is `/problem/{pid}/fs/file@POST`

<!--beg l desc_ProblemFSWrite -->

<!--end l-->


+ `pid`: `string` (required): 
    <!--beg l desc_ProblemFSWrite_pid -->
    
    <!--end l-->

+ `ProblemFSWriteRequest`: `any`: 
    <!--beg l desc_ProblemFSWrite_ProblemFSWriteRequest -->
    
    <!--end l-->



## DeleteProblem

The uri/restful key of this method is `/problem/{pid}@DELETE`

<!--beg l desc_DeleteProblem -->

<!--end l-->


+ `pid`: `string` (required): 
    <!--beg l desc_DeleteProblem_pid -->
    
    <!--end l-->

+ `DeleteProblemRequest`: `any`: 
    <!--beg l desc_DeleteProblem_DeleteProblemRequest -->
    
    <!--end l-->



## GetProblem

The uri/restful key of this method is `/problem/{pid}@GET`

<!--beg l desc_GetProblem -->

<!--end l-->


+ `pid`: `string` (required): 
    <!--beg l desc_GetProblem_pid -->
    
    <!--end l-->



## PutProblem

The uri/restful key of this method is `/problem/{pid}@PUT`

<!--beg l desc_PutProblem -->

<!--end l-->


+ `pid`: `string` (required): 
    <!--beg l desc_PutProblem_pid -->
    
    <!--end l-->

+ `PutProblemRequest`: `any`: 
    <!--beg l desc_PutProblem_PutProblemRequest -->
    
    <!--end l-->



## PostProblem

The uri/restful key of this method is `/problem@POST`

<!--beg l desc_PostProblem -->

<!--end l-->


+ `PostProblemRequest`: `any`: 
    <!--beg l desc_PostProblem_PostProblemRequest -->
    
    <!--end l-->




