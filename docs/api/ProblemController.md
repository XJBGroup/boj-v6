
# Api Group Documentation (ProblemController)

<!--beg l desc_ProblemController -->

<!--end l-->

## Apis


### CountProblem

The uri/restful key of this method is `/problem-count@GET`

<!--beg l desc_CountProblem -->

<!--end l-->

parameters:

+ `page`: [integer](#integer): 
    <!--beg l desc_CountProblem_params_page -->
    
    <!--end l-->


+ `page_size`: [integer](#integer): 
    <!--beg l desc_CountProblem_params_page_size -->
    
    <!--end l-->


+ `before_id`: [integer](#integer): 
    <!--beg l desc_CountProblem_params_before_id -->
    
    <!--end l-->


+ `order`: [string](#string): 
    <!--beg l desc_CountProblem_params_order -->
    
    <!--end l-->


responses:

+ code: `200`, type: [genericResponse](#genericResponse)
    <!--beg l desc_CountProblem_response_200_[genericResponse](#genericResponse) -->
    
    <!--end l-->


+ code: `200`, type: [CountProblemReply](#CountProblemReply)
    <!--beg l desc_CountProblem_response_200_[CountProblemReply](#CountProblemReply) -->
    
    <!--end l-->


+ code: `500`, type: No Response
    <!--beg l desc_CountProblem_response_500_No Response -->
    
    <!--end l-->




### ListProblem

The uri/restful key of this method is `/problem-list@GET`

<!--beg l desc_ListProblem -->

<!--end l-->

parameters:

+ `order`: [string](#string): 
    <!--beg l desc_ListProblem_params_order -->
    
    <!--end l-->


+ `page`: [integer](#integer): 
    <!--beg l desc_ListProblem_params_page -->
    
    <!--end l-->


+ `page_size`: [integer](#integer): 
    <!--beg l desc_ListProblem_params_page_size -->
    
    <!--end l-->


+ `before_id`: [integer](#integer): 
    <!--beg l desc_ListProblem_params_before_id -->
    
    <!--end l-->


responses:

+ code: `200`, type: [genericResponse](#genericResponse)
    <!--beg l desc_ListProblem_response_200_[genericResponse](#genericResponse) -->
    
    <!--end l-->


+ code: `200`, type: [ListProblemReply](#ListProblemReply)
    <!--beg l desc_ListProblem_response_200_[ListProblemReply](#ListProblemReply) -->
    
    <!--end l-->


+ code: `500`, type: No Response
    <!--beg l desc_ListProblem_response_500_No Response -->
    
    <!--end l-->




### CountProblemDesc

The uri/restful key of this method is `/problem/{pid}/desc-count@GET`

<!--beg l desc_CountProblemDesc -->

<!--end l-->

parameters:

+ `pid`: [string](#string) (required): 
    <!--beg l desc_CountProblemDesc_params_pid -->
    
    <!--end l-->


+ `before_id`: [integer](#integer): 
    <!--beg l desc_CountProblemDesc_params_before_id -->
    
    <!--end l-->


+ `order`: [string](#string): 
    <!--beg l desc_CountProblemDesc_params_order -->
    
    <!--end l-->


+ `page`: [integer](#integer): 
    <!--beg l desc_CountProblemDesc_params_page -->
    
    <!--end l-->


+ `page_size`: [integer](#integer): 
    <!--beg l desc_CountProblemDesc_params_page_size -->
    
    <!--end l-->


responses:

+ code: `200`, type: [genericResponse](#genericResponse)
    <!--beg l desc_CountProblemDesc_response_200_[genericResponse](#genericResponse) -->
    
    <!--end l-->


+ code: `200`, type: [CountProblemDescReply](#CountProblemDescReply)
    <!--beg l desc_CountProblemDesc_response_200_[CountProblemDescReply](#CountProblemDescReply) -->
    
    <!--end l-->


+ code: `500`, type: No Response
    <!--beg l desc_CountProblemDesc_response_500_No Response -->
    
    <!--end l-->




### ListProblemDesc

The uri/restful key of this method is `/problem/{pid}/desc-list@GET`

<!--beg l desc_ListProblemDesc -->

<!--end l-->

parameters:

+ `pid`: [string](#string) (required): 
    <!--beg l desc_ListProblemDesc_params_pid -->
    
    <!--end l-->


+ `order`: [string](#string): 
    <!--beg l desc_ListProblemDesc_params_order -->
    
    <!--end l-->


+ `page`: [integer](#integer): 
    <!--beg l desc_ListProblemDesc_params_page -->
    
    <!--end l-->


+ `page_size`: [integer](#integer): 
    <!--beg l desc_ListProblemDesc_params_page_size -->
    
    <!--end l-->


+ `before_id`: [integer](#integer): 
    <!--beg l desc_ListProblemDesc_params_before_id -->
    
    <!--end l-->


responses:

+ code: `200`, type: [genericResponse](#genericResponse)
    <!--beg l desc_ListProblemDesc_response_200_[genericResponse](#genericResponse) -->
    
    <!--end l-->


+ code: `200`, type: [ListProblemDescReply](#ListProblemDescReply)
    <!--beg l desc_ListProblemDesc_response_200_[ListProblemDescReply](#ListProblemDescReply) -->
    
    <!--end l-->


+ code: `500`, type: No Response
    <!--beg l desc_ListProblemDesc_response_500_No Response -->
    
    <!--end l-->




### ChangeProblemDescriptionRef

The uri/restful key of this method is `/problem/{pid}/desc/ref@POST`

<!--beg l desc_ChangeProblemDescriptionRef -->

<!--end l-->

parameters:

+ `pid`: [string](#string) (required): 
    <!--beg l desc_ChangeProblemDescriptionRef_params_pid -->
    
    <!--end l-->


+ `ChangeProblemDescriptionRefRequest`: [ChangeProblemDescriptionRefRequest](#ChangeProblemDescriptionRefRequest): 
    <!--beg l desc_ChangeProblemDescriptionRef_params_ChangeProblemDescriptionRefRequest -->
    
    <!--end l-->


responses:

+ code: `200`, type: [genericResponse](#genericResponse)
    <!--beg l desc_ChangeProblemDescriptionRef_response_200_[genericResponse](#genericResponse) -->
    
    <!--end l-->


+ code: `200`, type: [ChangeProblemDescriptionRefReply](#ChangeProblemDescriptionRefReply)
    <!--beg l desc_ChangeProblemDescriptionRef_response_200_[ChangeProblemDescriptionRefReply](#ChangeProblemDescriptionRefReply) -->
    
    <!--end l-->


+ code: `500`, type: No Response
    <!--beg l desc_ChangeProblemDescriptionRef_response_500_No Response -->
    
    <!--end l-->




### DeleteProblemDesc

The uri/restful key of this method is `/problem/{pid}/desc@DELETE`

<!--beg l desc_DeleteProblemDesc -->

<!--end l-->

parameters:

+ `pid`: [string](#string) (required): 
    <!--beg l desc_DeleteProblemDesc_params_pid -->
    
    <!--end l-->


+ `DeleteProblemDescRequest`: [DeleteProblemDescRequest](#DeleteProblemDescRequest): 
    <!--beg l desc_DeleteProblemDesc_params_DeleteProblemDescRequest -->
    
    <!--end l-->


responses:

+ code: `200`, type: [genericResponse](#genericResponse)
    <!--beg l desc_DeleteProblemDesc_response_200_[genericResponse](#genericResponse) -->
    
    <!--end l-->


+ code: `200`, type: [DeleteProblemDescReply](#DeleteProblemDescReply)
    <!--beg l desc_DeleteProblemDesc_response_200_[DeleteProblemDescReply](#DeleteProblemDescReply) -->
    
    <!--end l-->


+ code: `500`, type: No Response
    <!--beg l desc_DeleteProblemDesc_response_500_No Response -->
    
    <!--end l-->




### GetProblemDesc

The uri/restful key of this method is `/problem/{pid}/desc@GET`

<!--beg l desc_GetProblemDesc -->

<!--end l-->

parameters:

+ `pid`: [string](#string) (required): 
    <!--beg l desc_GetProblemDesc_params_pid -->
    
    <!--end l-->


+ `name`: [string](#string): 
    <!--beg l desc_GetProblemDesc_params_name -->
    
    <!--end l-->


responses:

+ code: `200`, type: [genericResponse](#genericResponse)
    <!--beg l desc_GetProblemDesc_response_200_[genericResponse](#genericResponse) -->
    
    <!--end l-->


+ code: `200`, type: [GetProblemDescReply](#GetProblemDescReply)
    <!--beg l desc_GetProblemDesc_response_200_[GetProblemDescReply](#GetProblemDescReply) -->
    
    <!--end l-->


+ code: `500`, type: No Response
    <!--beg l desc_GetProblemDesc_response_500_No Response -->
    
    <!--end l-->




### PostProblemDesc

The uri/restful key of this method is `/problem/{pid}/desc@POST`

<!--beg l desc_PostProblemDesc -->

<!--end l-->

parameters:

+ `pid`: [string](#string) (required): 
    <!--beg l desc_PostProblemDesc_params_pid -->
    
    <!--end l-->


+ `PostProblemDescRequest`: [PostProblemDescRequest](#PostProblemDescRequest): 
    <!--beg l desc_PostProblemDesc_params_PostProblemDescRequest -->
    
    <!--end l-->


responses:

+ code: `200`, type: [genericResponse](#genericResponse)
    <!--beg l desc_PostProblemDesc_response_200_[genericResponse](#genericResponse) -->
    
    <!--end l-->


+ code: `200`, type: [PostProblemDescReply](#PostProblemDescReply)
    <!--beg l desc_PostProblemDesc_response_200_[PostProblemDescReply](#PostProblemDescReply) -->
    
    <!--end l-->


+ code: `500`, type: No Response
    <!--beg l desc_PostProblemDesc_response_500_No Response -->
    
    <!--end l-->




### PutProblemDesc

The uri/restful key of this method is `/problem/{pid}/desc@PUT`

<!--beg l desc_PutProblemDesc -->

<!--end l-->

parameters:

+ `pid`: [string](#string) (required): 
    <!--beg l desc_PutProblemDesc_params_pid -->
    
    <!--end l-->


+ `PutProblemDescRequest`: [PutProblemDescRequest](#PutProblemDescRequest): 
    <!--beg l desc_PutProblemDesc_params_PutProblemDescRequest -->
    
    <!--end l-->


responses:

+ code: `200`, type: [genericResponse](#genericResponse)
    <!--beg l desc_PutProblemDesc_response_200_[genericResponse](#genericResponse) -->
    
    <!--end l-->


+ code: `200`, type: [PutProblemDescReply](#PutProblemDescReply)
    <!--beg l desc_PutProblemDesc_response_200_[PutProblemDescReply](#PutProblemDescReply) -->
    
    <!--end l-->


+ code: `500`, type: No Response
    <!--beg l desc_PutProblemDesc_response_500_No Response -->
    
    <!--end l-->




### ProblemFSReadConfig

The uri/restful key of this method is `/problem/{pid}/fs/config@GET`

<!--beg l desc_ProblemFSReadConfig -->

<!--end l-->

parameters:

+ `pid`: [string](#string) (required): 
    <!--beg l desc_ProblemFSReadConfig_params_pid -->
    
    <!--end l-->


+ `path`: [string](#string): 
    <!--beg l desc_ProblemFSReadConfig_params_path -->
    
    <!--end l-->


responses:

+ code: `200`, type: [genericResponse](#genericResponse)
    <!--beg l desc_ProblemFSReadConfig_response_200_[genericResponse](#genericResponse) -->
    
    <!--end l-->


+ code: `200`, type: [ProblemFSReadConfigReply](#ProblemFSReadConfigReply)
    <!--beg l desc_ProblemFSReadConfig_response_200_[ProblemFSReadConfigReply](#ProblemFSReadConfigReply) -->
    
    <!--end l-->


+ code: `500`, type: No Response
    <!--beg l desc_ProblemFSReadConfig_response_500_No Response -->
    
    <!--end l-->




### ProblemFSWriteConfig

The uri/restful key of this method is `/problem/{pid}/fs/config@POST`

<!--beg l desc_ProblemFSWriteConfig -->

<!--end l-->

parameters:

+ `pid`: [string](#string) (required): 
    <!--beg l desc_ProblemFSWriteConfig_params_pid -->
    
    <!--end l-->


+ `ProblemFSWriteConfigRequest`: [ProblemFSWriteConfigRequest](#ProblemFSWriteConfigRequest): 
    <!--beg l desc_ProblemFSWriteConfig_params_ProblemFSWriteConfigRequest -->
    
    <!--end l-->


responses:

+ code: `200`, type: [genericResponse](#genericResponse)
    <!--beg l desc_ProblemFSWriteConfig_response_200_[genericResponse](#genericResponse) -->
    
    <!--end l-->


+ code: `200`, type: [ProblemFSWriteConfigReply](#ProblemFSWriteConfigReply)
    <!--beg l desc_ProblemFSWriteConfig_response_200_[ProblemFSWriteConfigReply](#ProblemFSWriteConfigReply) -->
    
    <!--end l-->


+ code: `500`, type: No Response
    <!--beg l desc_ProblemFSWriteConfig_response_500_No Response -->
    
    <!--end l-->




### ProblemFSPutConfig

The uri/restful key of this method is `/problem/{pid}/fs/config@PUT`

<!--beg l desc_ProblemFSPutConfig -->

<!--end l-->

parameters:

+ `pid`: [string](#string) (required): 
    <!--beg l desc_ProblemFSPutConfig_params_pid -->
    
    <!--end l-->


+ `ProblemFSPutConfigRequest`: [ProblemFSPutConfigRequest](#ProblemFSPutConfigRequest): 
    <!--beg l desc_ProblemFSPutConfig_params_ProblemFSPutConfigRequest -->
    
    <!--end l-->


responses:

+ code: `200`, type: [genericResponse](#genericResponse)
    <!--beg l desc_ProblemFSPutConfig_response_200_[genericResponse](#genericResponse) -->
    
    <!--end l-->


+ code: `200`, type: [ProblemFSPutConfigReply](#ProblemFSPutConfigReply)
    <!--beg l desc_ProblemFSPutConfig_response_200_[ProblemFSPutConfigReply](#ProblemFSPutConfigReply) -->
    
    <!--end l-->


+ code: `500`, type: No Response
    <!--beg l desc_ProblemFSPutConfig_response_500_No Response -->
    
    <!--end l-->




### ProblemFSZipRead

The uri/restful key of this method is `/problem/{pid}/fs/directory/zip@GET`

<!--beg l desc_ProblemFSZipRead -->

<!--end l-->

parameters:

+ `pid`: [string](#string) (required): 
    <!--beg l desc_ProblemFSZipRead_params_pid -->
    
    <!--end l-->


+ `path`: [string](#string): 
    <!--beg l desc_ProblemFSZipRead_params_path -->
    
    <!--end l-->


responses:

+ code: `200`, type: [genericResponse](#genericResponse)
    <!--beg l desc_ProblemFSZipRead_response_200_[genericResponse](#genericResponse) -->
    
    <!--end l-->


+ code: `200`, type: [ProblemFSZipReadReply](#ProblemFSZipReadReply)
    <!--beg l desc_ProblemFSZipRead_response_200_[ProblemFSZipReadReply](#ProblemFSZipReadReply) -->
    
    <!--end l-->


+ code: `500`, type: No Response
    <!--beg l desc_ProblemFSZipRead_response_500_No Response -->
    
    <!--end l-->




### ProblemFSZipWrite

The uri/restful key of this method is `/problem/{pid}/fs/directory/zip@POST`

<!--beg l desc_ProblemFSZipWrite -->

<!--end l-->

parameters:

+ `pid`: [string](#string) (required): 
    <!--beg l desc_ProblemFSZipWrite_params_pid -->
    
    <!--end l-->


+ `ProblemFSZipWriteRequest`: [ProblemFSZipWriteRequest](#ProblemFSZipWriteRequest): 
    <!--beg l desc_ProblemFSZipWrite_params_ProblemFSZipWriteRequest -->
    
    <!--end l-->


responses:

+ code: `200`, type: [genericResponse](#genericResponse)
    <!--beg l desc_ProblemFSZipWrite_response_200_[genericResponse](#genericResponse) -->
    
    <!--end l-->


+ code: `200`, type: [ProblemFSZipWriteReply](#ProblemFSZipWriteReply)
    <!--beg l desc_ProblemFSZipWrite_response_200_[ProblemFSZipWriteReply](#ProblemFSZipWriteReply) -->
    
    <!--end l-->


+ code: `500`, type: No Response
    <!--beg l desc_ProblemFSZipWrite_response_500_No Response -->
    
    <!--end l-->




### ProblemFSRemoveAll

The uri/restful key of this method is `/problem/{pid}/fs/directory@DELETE`

<!--beg l desc_ProblemFSRemoveAll -->

<!--end l-->

parameters:

+ `pid`: [string](#string) (required): 
    <!--beg l desc_ProblemFSRemoveAll_params_pid -->
    
    <!--end l-->


+ `ProblemFSRemoveAllRequest`: [ProblemFSRemoveAllRequest](#ProblemFSRemoveAllRequest): 
    <!--beg l desc_ProblemFSRemoveAll_params_ProblemFSRemoveAllRequest -->
    
    <!--end l-->


responses:

+ code: `200`, type: [genericResponse](#genericResponse)
    <!--beg l desc_ProblemFSRemoveAll_response_200_[genericResponse](#genericResponse) -->
    
    <!--end l-->


+ code: `200`, type: [ProblemFSRemoveAllReply](#ProblemFSRemoveAllReply)
    <!--beg l desc_ProblemFSRemoveAll_response_200_[ProblemFSRemoveAllReply](#ProblemFSRemoveAllReply) -->
    
    <!--end l-->


+ code: `500`, type: No Response
    <!--beg l desc_ProblemFSRemoveAll_response_500_No Response -->
    
    <!--end l-->




### ProblemFSLS

The uri/restful key of this method is `/problem/{pid}/fs/directory@GET`

<!--beg l desc_ProblemFSLS -->

<!--end l-->

parameters:

+ `pid`: [string](#string) (required): 
    <!--beg l desc_ProblemFSLS_params_pid -->
    
    <!--end l-->


+ `path`: [string](#string): 
    <!--beg l desc_ProblemFSLS_params_path -->
    
    <!--end l-->


responses:

+ code: `200`, type: [genericResponse](#genericResponse)
    <!--beg l desc_ProblemFSLS_response_200_[genericResponse](#genericResponse) -->
    
    <!--end l-->


+ code: `200`, type: [ProblemFSLSReply](#ProblemFSLSReply)
    <!--beg l desc_ProblemFSLS_response_200_[ProblemFSLSReply](#ProblemFSLSReply) -->
    
    <!--end l-->


+ code: `500`, type: No Response
    <!--beg l desc_ProblemFSLS_response_500_No Response -->
    
    <!--end l-->




### ProblemFSWrites

The uri/restful key of this method is `/problem/{pid}/fs/directory@POST`

<!--beg l desc_ProblemFSWrites -->

<!--end l-->

parameters:

+ `pid`: [string](#string) (required): 
    <!--beg l desc_ProblemFSWrites_params_pid -->
    
    <!--end l-->


+ `ProblemFSWritesRequest`: [ProblemFSWritesRequest](#ProblemFSWritesRequest): 
    <!--beg l desc_ProblemFSWrites_params_ProblemFSWritesRequest -->
    
    <!--end l-->


responses:

+ code: `200`, type: [genericResponse](#genericResponse)
    <!--beg l desc_ProblemFSWrites_response_200_[genericResponse](#genericResponse) -->
    
    <!--end l-->


+ code: `200`, type: [ProblemFSWritesReply](#ProblemFSWritesReply)
    <!--beg l desc_ProblemFSWrites_response_200_[ProblemFSWritesReply](#ProblemFSWritesReply) -->
    
    <!--end l-->


+ code: `500`, type: No Response
    <!--beg l desc_ProblemFSWrites_response_500_No Response -->
    
    <!--end l-->




### ProblemFSMkdir

The uri/restful key of this method is `/problem/{pid}/fs/directory@PUT`

<!--beg l desc_ProblemFSMkdir -->

<!--end l-->

parameters:

+ `pid`: [string](#string) (required): 
    <!--beg l desc_ProblemFSMkdir_params_pid -->
    
    <!--end l-->


+ `ProblemFSMkdirRequest`: [ProblemFSMkdirRequest](#ProblemFSMkdirRequest): 
    <!--beg l desc_ProblemFSMkdir_params_ProblemFSMkdirRequest -->
    
    <!--end l-->


responses:

+ code: `200`, type: [genericResponse](#genericResponse)
    <!--beg l desc_ProblemFSMkdir_response_200_[genericResponse](#genericResponse) -->
    
    <!--end l-->


+ code: `200`, type: [ProblemFSMkdirReply](#ProblemFSMkdirReply)
    <!--beg l desc_ProblemFSMkdir_response_200_[ProblemFSMkdirReply](#ProblemFSMkdirReply) -->
    
    <!--end l-->


+ code: `500`, type: No Response
    <!--beg l desc_ProblemFSMkdir_response_500_No Response -->
    
    <!--end l-->




### ProblemFSRead

The uri/restful key of this method is `/problem/{pid}/fs/file/content@GET`

<!--beg l desc_ProblemFSRead -->

<!--end l-->

parameters:

+ `pid`: [string](#string) (required): 
    <!--beg l desc_ProblemFSRead_params_pid -->
    
    <!--end l-->


+ `path`: [string](#string): 
    <!--beg l desc_ProblemFSRead_params_path -->
    
    <!--end l-->


responses:

+ code: `200`, type: [genericResponse](#genericResponse)
    <!--beg l desc_ProblemFSRead_response_200_[genericResponse](#genericResponse) -->
    
    <!--end l-->


+ code: `200`, type: [ProblemFSReadReply](#ProblemFSReadReply)
    <!--beg l desc_ProblemFSRead_response_200_[ProblemFSReadReply](#ProblemFSReadReply) -->
    
    <!--end l-->


+ code: `500`, type: No Response
    <!--beg l desc_ProblemFSRead_response_500_No Response -->
    
    <!--end l-->




### ProblemFSRemove

The uri/restful key of this method is `/problem/{pid}/fs/file@DELETE`

<!--beg l desc_ProblemFSRemove -->

<!--end l-->

parameters:

+ `pid`: [string](#string) (required): 
    <!--beg l desc_ProblemFSRemove_params_pid -->
    
    <!--end l-->


+ `ProblemFSRemoveRequest`: [ProblemFSRemoveRequest](#ProblemFSRemoveRequest): 
    <!--beg l desc_ProblemFSRemove_params_ProblemFSRemoveRequest -->
    
    <!--end l-->


responses:

+ code: `200`, type: [genericResponse](#genericResponse)
    <!--beg l desc_ProblemFSRemove_response_200_[genericResponse](#genericResponse) -->
    
    <!--end l-->


+ code: `200`, type: [ProblemFSRemoveReply](#ProblemFSRemoveReply)
    <!--beg l desc_ProblemFSRemove_response_200_[ProblemFSRemoveReply](#ProblemFSRemoveReply) -->
    
    <!--end l-->


+ code: `500`, type: No Response
    <!--beg l desc_ProblemFSRemove_response_500_No Response -->
    
    <!--end l-->




### ProblemFSStat

The uri/restful key of this method is `/problem/{pid}/fs/file@GET`

<!--beg l desc_ProblemFSStat -->

<!--end l-->

parameters:

+ `pid`: [string](#string) (required): 
    <!--beg l desc_ProblemFSStat_params_pid -->
    
    <!--end l-->


+ `path`: [string](#string): 
    <!--beg l desc_ProblemFSStat_params_path -->
    
    <!--end l-->


responses:

+ code: `200`, type: [genericResponse](#genericResponse)
    <!--beg l desc_ProblemFSStat_response_200_[genericResponse](#genericResponse) -->
    
    <!--end l-->


+ code: `200`, type: [ProblemFSStatReply](#ProblemFSStatReply)
    <!--beg l desc_ProblemFSStat_response_200_[ProblemFSStatReply](#ProblemFSStatReply) -->
    
    <!--end l-->


+ code: `500`, type: No Response
    <!--beg l desc_ProblemFSStat_response_500_No Response -->
    
    <!--end l-->




### ProblemFSWrite

The uri/restful key of this method is `/problem/{pid}/fs/file@POST`

<!--beg l desc_ProblemFSWrite -->

<!--end l-->

parameters:

+ `pid`: [string](#string) (required): 
    <!--beg l desc_ProblemFSWrite_params_pid -->
    
    <!--end l-->


+ `ProblemFSWriteRequest`: [ProblemFSWriteRequest](#ProblemFSWriteRequest): 
    <!--beg l desc_ProblemFSWrite_params_ProblemFSWriteRequest -->
    
    <!--end l-->


responses:

+ code: `200`, type: [genericResponse](#genericResponse)
    <!--beg l desc_ProblemFSWrite_response_200_[genericResponse](#genericResponse) -->
    
    <!--end l-->


+ code: `200`, type: [ProblemFSWriteReply](#ProblemFSWriteReply)
    <!--beg l desc_ProblemFSWrite_response_200_[ProblemFSWriteReply](#ProblemFSWriteReply) -->
    
    <!--end l-->


+ code: `500`, type: No Response
    <!--beg l desc_ProblemFSWrite_response_500_No Response -->
    
    <!--end l-->




### DeleteProblem

The uri/restful key of this method is `/problem/{pid}@DELETE`

<!--beg l desc_DeleteProblem -->

<!--end l-->

parameters:

+ `pid`: [string](#string) (required): 
    <!--beg l desc_DeleteProblem_params_pid -->
    
    <!--end l-->


+ `DeleteProblemRequest`: [DeleteProblemRequest](#DeleteProblemRequest): 
    <!--beg l desc_DeleteProblem_params_DeleteProblemRequest -->
    
    <!--end l-->


responses:

+ code: `200`, type: [genericResponse](#genericResponse)
    <!--beg l desc_DeleteProblem_response_200_[genericResponse](#genericResponse) -->
    
    <!--end l-->


+ code: `200`, type: [DeleteProblemReply](#DeleteProblemReply)
    <!--beg l desc_DeleteProblem_response_200_[DeleteProblemReply](#DeleteProblemReply) -->
    
    <!--end l-->


+ code: `500`, type: No Response
    <!--beg l desc_DeleteProblem_response_500_No Response -->
    
    <!--end l-->




### GetProblem

The uri/restful key of this method is `/problem/{pid}@GET`

<!--beg l desc_GetProblem -->

<!--end l-->

parameters:

+ `pid`: [string](#string) (required): 
    <!--beg l desc_GetProblem_params_pid -->
    
    <!--end l-->


responses:

+ code: `200`, type: [genericResponse](#genericResponse)
    <!--beg l desc_GetProblem_response_200_[genericResponse](#genericResponse) -->
    
    <!--end l-->


+ code: `200`, type: [GetProblemReply](#GetProblemReply)
    <!--beg l desc_GetProblem_response_200_[GetProblemReply](#GetProblemReply) -->
    
    <!--end l-->


+ code: `500`, type: No Response
    <!--beg l desc_GetProblem_response_500_No Response -->
    
    <!--end l-->




### PutProblem

The uri/restful key of this method is `/problem/{pid}@PUT`

<!--beg l desc_PutProblem -->

<!--end l-->

parameters:

+ `pid`: [string](#string) (required): 
    <!--beg l desc_PutProblem_params_pid -->
    
    <!--end l-->


+ `PutProblemRequest`: [PutProblemRequest](#PutProblemRequest): 
    <!--beg l desc_PutProblem_params_PutProblemRequest -->
    
    <!--end l-->


responses:

+ code: `200`, type: [genericResponse](#genericResponse)
    <!--beg l desc_PutProblem_response_200_[genericResponse](#genericResponse) -->
    
    <!--end l-->


+ code: `200`, type: [PutProblemReply](#PutProblemReply)
    <!--beg l desc_PutProblem_response_200_[PutProblemReply](#PutProblemReply) -->
    
    <!--end l-->


+ code: `500`, type: No Response
    <!--beg l desc_PutProblem_response_500_No Response -->
    
    <!--end l-->




### PostProblem

The uri/restful key of this method is `/problem@POST`

<!--beg l desc_PostProblem -->

<!--end l-->

parameters:

+ `PostProblemRequest`: [PostProblemRequest](#PostProblemRequest): 
    <!--beg l desc_PostProblem_params_PostProblemRequest -->
    
    <!--end l-->


responses:

+ code: `200`, type: [genericResponse](#genericResponse)
    <!--beg l desc_PostProblem_response_200_[genericResponse](#genericResponse) -->
    
    <!--end l-->


+ code: `200`, type: [PostProblemReply](#PostProblemReply)
    <!--beg l desc_PostProblem_response_200_[PostProblemReply](#PostProblemReply) -->
    
    <!--end l-->


+ code: `500`, type: No Response
    <!--beg l desc_PostProblem_response_500_No Response -->
    
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

    
### [CountProblemReply](./ObjectModelSpec.md#CountProblemReply)

+ type: [object](#CountProblemReply)

+ fields:
    
    + `code`: integer: 
        <!--beg l desc_{{object_name}}_code -->
        
        <!--end l-->

    + `data`: array: 
        <!--beg l desc_{{object_name}}_data -->
        
        <!--end l-->

    
### [ListProblemReply](./ObjectModelSpec.md#ListProblemReply)

+ type: [object](#ListProblemReply)

+ fields:
    
    + `code`: integer: 
        <!--beg l desc_{{object_name}}_code -->
        
        <!--end l-->

    + `data`: array: 
        <!--beg l desc_{{object_name}}_data -->
        
        <!--end l-->

    
### [CountProblemDescReply](./ObjectModelSpec.md#CountProblemDescReply)

+ type: [object](#CountProblemDescReply)

+ fields:
    
    + `code`: integer: 
        <!--beg l desc_{{object_name}}_code -->
        
        <!--end l-->

    + `data`: integer: 
        <!--beg l desc_{{object_name}}_data -->
        
        <!--end l-->

    
### [ListProblemDescReply](./ObjectModelSpec.md#ListProblemDescReply)

+ type: [object](#ListProblemDescReply)

+ fields:
    
    + `code`: integer: 
        <!--beg l desc_{{object_name}}_code -->
        
        <!--end l-->

    + `data`: array: 
        <!--beg l desc_{{object_name}}_data -->
        
        <!--end l-->

    
### [ChangeProblemDescriptionRefRequest](./ObjectModelSpec.md#ChangeProblemDescriptionRefRequest)

+ type: [object](#ChangeProblemDescriptionRefRequest)

+ fields:
    
    + `name`: string: 
        <!--beg l desc_{{object_name}}_name -->
        
        <!--end l-->

    + `new_name`: string: 
        <!--beg l desc_{{object_name}}_new_name -->
        
        <!--end l-->

    
### [ChangeProblemDescriptionRefReply](./ObjectModelSpec.md#ChangeProblemDescriptionRefReply)

+ type: [object](#ChangeProblemDescriptionRefReply)

+ fields:
    
    + `code`: integer: 
        <!--beg l desc_{{object_name}}_code -->
        
        <!--end l-->

    
### [DeleteProblemDescRequest](./ObjectModelSpec.md#DeleteProblemDescRequest)

+ type: [object](#DeleteProblemDescRequest)

+ fields:
    
    + `name`: string: 
        <!--beg l desc_{{object_name}}_name -->
        
        <!--end l-->

    
### [DeleteProblemDescReply](./ObjectModelSpec.md#DeleteProblemDescReply)

+ type: [object](#DeleteProblemDescReply)

+ fields:
    
    + `code`: integer: 
        <!--beg l desc_{{object_name}}_code -->
        
        <!--end l-->

    
### [GetProblemDescReply](./ObjectModelSpec.md#GetProblemDescReply)

+ type: [object](#GetProblemDescReply)

+ fields:
    
    + `code`: integer: 
        <!--beg l desc_{{object_name}}_code -->
        
        <!--end l-->

    + `data`: string: 
        <!--beg l desc_{{object_name}}_data -->
        
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

    
### [PostProblemDescReply](./ObjectModelSpec.md#PostProblemDescReply)

+ type: [object](#PostProblemDescReply)

+ fields:
    
    + `code`: integer: 
        <!--beg l desc_{{object_name}}_code -->
        
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

    
### [PutProblemDescReply](./ObjectModelSpec.md#PutProblemDescReply)

+ type: [object](#PutProblemDescReply)

+ fields:
    
    + `code`: integer: 
        <!--beg l desc_{{object_name}}_code -->
        
        <!--end l-->

    
### [ProblemFSReadConfigReply](./ObjectModelSpec.md#ProblemFSReadConfigReply)

+ type: [object](#ProblemFSReadConfigReply)

+ fields:
    
    + `code`: integer: 
        <!--beg l desc_{{object_name}}_code -->
        
        <!--end l-->

    + `data`: : 
        <!--beg l desc_{{object_name}}_data -->
        
        <!--end l-->

    
### [ProblemFSWriteConfigRequest](./ObjectModelSpec.md#ProblemFSWriteConfigRequest)

+ type: [object](#ProblemFSWriteConfigRequest)

+ fields:
    
    + `path`: string: 
        <!--beg l desc_{{object_name}}_path -->
        
        <!--end l-->

    
### [ProblemFSWriteConfigReply](./ObjectModelSpec.md#ProblemFSWriteConfigReply)

+ type: [object](#ProblemFSWriteConfigReply)

+ fields:
    
    + `code`: integer: 
        <!--beg l desc_{{object_name}}_code -->
        
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

    
### [ProblemFSPutConfigReply](./ObjectModelSpec.md#ProblemFSPutConfigReply)

+ type: [object](#ProblemFSPutConfigReply)

+ fields:
    
    + `code`: integer: 
        <!--beg l desc_{{object_name}}_code -->
        
        <!--end l-->

    + `data`: : 
        <!--beg l desc_{{object_name}}_data -->
        
        <!--end l-->

    
### [ProblemFSZipReadReply](./ObjectModelSpec.md#ProblemFSZipReadReply)

+ type: [object](#ProblemFSZipReadReply)

+ fields:
    
    + `code`: integer: 
        <!--beg l desc_{{object_name}}_code -->
        
        <!--end l-->

    
### [ProblemFSZipWriteRequest](./ObjectModelSpec.md#ProblemFSZipWriteRequest)

+ type: [object](#ProblemFSZipWriteRequest)

+ fields:
    
    + `path`: string: 
        <!--beg l desc_{{object_name}}_path -->
        
        <!--end l-->

    
### [ProblemFSZipWriteReply](./ObjectModelSpec.md#ProblemFSZipWriteReply)

+ type: [object](#ProblemFSZipWriteReply)

+ fields:
    
    + `code`: integer: 
        <!--beg l desc_{{object_name}}_code -->
        
        <!--end l-->

    
### [ProblemFSRemoveAllRequest](./ObjectModelSpec.md#ProblemFSRemoveAllRequest)

+ type: [object](#ProblemFSRemoveAllRequest)

+ fields:
    
    + `path`: string: 
        <!--beg l desc_{{object_name}}_path -->
        
        <!--end l-->

    
### [ProblemFSRemoveAllReply](./ObjectModelSpec.md#ProblemFSRemoveAllReply)

+ type: [object](#ProblemFSRemoveAllReply)

+ fields:
    
    + `code`: integer: 
        <!--beg l desc_{{object_name}}_code -->
        
        <!--end l-->

    
### [ProblemFSLSReply](./ObjectModelSpec.md#ProblemFSLSReply)

+ type: [object](#ProblemFSLSReply)

+ fields:
    
    + `code`: integer: 
        <!--beg l desc_{{object_name}}_code -->
        
        <!--end l-->

    + `data`: array: 
        <!--beg l desc_{{object_name}}_data -->
        
        <!--end l-->

    
### [ProblemFSWritesRequest](./ObjectModelSpec.md#ProblemFSWritesRequest)

+ type: [object](#ProblemFSWritesRequest)

+ fields:
    
    + `path`: string: 
        <!--beg l desc_{{object_name}}_path -->
        
        <!--end l-->

    
### [ProblemFSWritesReply](./ObjectModelSpec.md#ProblemFSWritesReply)

+ type: [object](#ProblemFSWritesReply)

+ fields:
    
    + `code`: integer: 
        <!--beg l desc_{{object_name}}_code -->
        
        <!--end l-->

    
### [ProblemFSMkdirRequest](./ObjectModelSpec.md#ProblemFSMkdirRequest)

+ type: [object](#ProblemFSMkdirRequest)

+ fields:
    
    + `path`: string: 
        <!--beg l desc_{{object_name}}_path -->
        
        <!--end l-->

    
### [ProblemFSMkdirReply](./ObjectModelSpec.md#ProblemFSMkdirReply)

+ type: [object](#ProblemFSMkdirReply)

+ fields:
    
    + `code`: integer: 
        <!--beg l desc_{{object_name}}_code -->
        
        <!--end l-->

    
### [ProblemFSReadReply](./ObjectModelSpec.md#ProblemFSReadReply)

+ type: [object](#ProblemFSReadReply)

+ fields:
    
    
### [ProblemFSRemoveRequest](./ObjectModelSpec.md#ProblemFSRemoveRequest)

+ type: [object](#ProblemFSRemoveRequest)

+ fields:
    
    + `path`: string: 
        <!--beg l desc_{{object_name}}_path -->
        
        <!--end l-->

    
### [ProblemFSRemoveReply](./ObjectModelSpec.md#ProblemFSRemoveReply)

+ type: [object](#ProblemFSRemoveReply)

+ fields:
    
    + `code`: integer: 
        <!--beg l desc_{{object_name}}_code -->
        
        <!--end l-->

    + `data`: : 
        <!--beg l desc_{{object_name}}_data -->
        
        <!--end l-->

    
### [ProblemFSStatReply](./ObjectModelSpec.md#ProblemFSStatReply)

+ type: [object](#ProblemFSStatReply)

+ fields:
    
    + `code`: integer: 
        <!--beg l desc_{{object_name}}_code -->
        
        <!--end l-->

    + `data`: : 
        <!--beg l desc_{{object_name}}_data -->
        
        <!--end l-->

    
### [ProblemFSWriteRequest](./ObjectModelSpec.md#ProblemFSWriteRequest)

+ type: [object](#ProblemFSWriteRequest)

+ fields:
    
    + `path`: string: 
        <!--beg l desc_{{object_name}}_path -->
        
        <!--end l-->

    
### [ProblemFSWriteReply](./ObjectModelSpec.md#ProblemFSWriteReply)

+ type: [object](#ProblemFSWriteReply)

+ fields:
    
    + `code`: integer: 
        <!--beg l desc_{{object_name}}_code -->
        
        <!--end l-->

    
### [DeleteProblemRequest](./ObjectModelSpec.md#DeleteProblemRequest)

+ type: [object](#DeleteProblemRequest)

+ fields:
    
    
### [DeleteProblemReply](./ObjectModelSpec.md#DeleteProblemReply)

+ type: [object](#DeleteProblemReply)

+ fields:
    
    + `code`: integer: 
        <!--beg l desc_{{object_name}}_code -->
        
        <!--end l-->

    
### [GetProblemReply](./ObjectModelSpec.md#GetProblemReply)

+ type: [object](#GetProblemReply)

+ fields:
    
    + `code`: integer: 
        <!--beg l desc_{{object_name}}_code -->
        
        <!--end l-->

    + `data`: : 
        <!--beg l desc_{{object_name}}_data -->
        
        <!--end l-->

    
### [PutProblemRequest](./ObjectModelSpec.md#PutProblemRequest)

+ type: [object](#PutProblemRequest)

+ fields:
    
    + `description_ref`: string: 
        <!--beg l desc_{{object_name}}_description_ref -->
        
        <!--end l-->

    + `title`: string: 
        <!--beg l desc_{{object_name}}_title -->
        
        <!--end l-->

    
### [PutProblemReply](./ObjectModelSpec.md#PutProblemReply)

+ type: [object](#PutProblemReply)

+ fields:
    
    + `code`: integer: 
        <!--beg l desc_{{object_name}}_code -->
        
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

    
### [PostProblemReply](./ObjectModelSpec.md#PostProblemReply)

+ type: [object](#PostProblemReply)

+ fields:
    
    + `code`: integer: 
        <!--beg l desc_{{object_name}}_code -->
        
        <!--end l-->

    + `data`: : 
        <!--beg l desc_{{object_name}}_data -->
        
        <!--end l-->

    

