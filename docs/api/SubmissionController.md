
# Api Group Documentation (SubmissionController)


## PostSubmission

The uri/restful key of this method is `/problem/{pid}/submission@POST`

<!--beg l desc_PostSubmission -->

<!--end l-->


+ `pid`: `string` (required): 
    <!--beg l desc_PostSubmission_pid -->
    
    <!--end l-->

+ `PostSubmissionRequest`: `any`: 
    <!--beg l desc_PostSubmission_PostSubmissionRequest -->
    
    <!--end l-->



## CountSubmission

The uri/restful key of this method is `/submission-count@GET`

<!--beg l desc_CountSubmission -->

<!--end l-->


+ `page`: `integer`: 
    <!--beg l desc_CountSubmission_page -->
    
    <!--end l-->

+ `page_size`: `integer`: 
    <!--beg l desc_CountSubmission_page_size -->
    
    <!--end l-->

+ `mem_order`: `boolean`: 
    <!--beg l desc_CountSubmission_mem_order -->
    
    <!--end l-->

+ `time_order`: `boolean`: 
    <!--beg l desc_CountSubmission_time_order -->
    
    <!--end l-->

+ `id_order`: `boolean`: 
    <!--beg l desc_CountSubmission_id_order -->
    
    <!--end l-->

+ `by_user`: `integer`: 
    <!--beg l desc_CountSubmission_by_user -->
    
    <!--end l-->

+ `on_problem`: `integer`: 
    <!--beg l desc_CountSubmission_on_problem -->
    
    <!--end l-->

+ `with_language`: `integer`: 
    <!--beg l desc_CountSubmission_with_language -->
    
    <!--end l-->

+ `has_status`: `integer`: 
    <!--beg l desc_CountSubmission_has_status -->
    
    <!--end l-->



## ListSubmission

The uri/restful key of this method is `/submission-list@GET`

<!--beg l desc_ListSubmission -->

<!--end l-->


+ `page`: `integer`: 
    <!--beg l desc_ListSubmission_page -->
    
    <!--end l-->

+ `page_size`: `integer`: 
    <!--beg l desc_ListSubmission_page_size -->
    
    <!--end l-->

+ `mem_order`: `boolean`: 
    <!--beg l desc_ListSubmission_mem_order -->
    
    <!--end l-->

+ `time_order`: `boolean`: 
    <!--beg l desc_ListSubmission_time_order -->
    
    <!--end l-->

+ `id_order`: `boolean`: 
    <!--beg l desc_ListSubmission_id_order -->
    
    <!--end l-->

+ `by_user`: `integer`: 
    <!--beg l desc_ListSubmission_by_user -->
    
    <!--end l-->

+ `on_problem`: `integer`: 
    <!--beg l desc_ListSubmission_on_problem -->
    
    <!--end l-->

+ `with_language`: `integer`: 
    <!--beg l desc_ListSubmission_with_language -->
    
    <!--end l-->

+ `has_status`: `integer`: 
    <!--beg l desc_ListSubmission_has_status -->
    
    <!--end l-->



## GetSubmissionContent

The uri/restful key of this method is `/submission/{sid}/content@GET`

<!--beg l desc_GetSubmissionContent -->

<!--end l-->


+ `sid`: `string` (required): 
    <!--beg l desc_GetSubmissionContent_sid -->
    
    <!--end l-->



## DeleteSubmission

The uri/restful key of this method is `/submission/{sid}@DELETE`

<!--beg l desc_DeleteSubmission -->

<!--end l-->


+ `sid`: `string` (required): 
    <!--beg l desc_DeleteSubmission_sid -->
    
    <!--end l-->

+ `DeleteSubmissionRequest`: `any`: 
    <!--beg l desc_DeleteSubmission_DeleteSubmissionRequest -->
    
    <!--end l-->



## GetSubmission

The uri/restful key of this method is `/submission/{sid}@GET`

<!--beg l desc_GetSubmission -->

<!--end l-->


+ `sid`: `string` (required): 
    <!--beg l desc_GetSubmission_sid -->
    
    <!--end l-->



