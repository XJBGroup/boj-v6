
# Api Group Documentation (SubmissionController)

<!--beg l desc_SubmissionController -->

<!--end l-->

## Apis


### PostSubmission

The uri/restful key of this method is `/problem/{pid}/submission@POST`

<!--beg l desc_PostSubmission -->

<!--end l-->


+ `pid`: [string](#string) (required): 
    <!--beg l desc_PostSubmission_pid -->
    
    <!--end l-->


+ `PostSubmissionRequest`: [PostSubmissionRequest](#PostSubmissionRequest): 
    <!--beg l desc_PostSubmission_PostSubmissionRequest -->
    
    <!--end l-->



### CountSubmission

The uri/restful key of this method is `/submission-count@GET`

<!--beg l desc_CountSubmission -->

<!--end l-->


+ `page`: [integer](#integer): 
    <!--beg l desc_CountSubmission_page -->
    
    <!--end l-->


+ `page_size`: [integer](#integer): 
    <!--beg l desc_CountSubmission_page_size -->
    
    <!--end l-->


+ `mem_order`: [boolean](#boolean): 
    <!--beg l desc_CountSubmission_mem_order -->
    
    <!--end l-->


+ `time_order`: [boolean](#boolean): 
    <!--beg l desc_CountSubmission_time_order -->
    
    <!--end l-->


+ `id_order`: [boolean](#boolean): 
    <!--beg l desc_CountSubmission_id_order -->
    
    <!--end l-->


+ `by_user`: [integer](#integer): 
    <!--beg l desc_CountSubmission_by_user -->
    
    <!--end l-->


+ `on_problem`: [integer](#integer): 
    <!--beg l desc_CountSubmission_on_problem -->
    
    <!--end l-->


+ `with_language`: [integer](#integer): 
    <!--beg l desc_CountSubmission_with_language -->
    
    <!--end l-->


+ `has_status`: [integer](#integer): 
    <!--beg l desc_CountSubmission_has_status -->
    
    <!--end l-->



### ListSubmission

The uri/restful key of this method is `/submission-list@GET`

<!--beg l desc_ListSubmission -->

<!--end l-->


+ `page`: [integer](#integer): 
    <!--beg l desc_ListSubmission_page -->
    
    <!--end l-->


+ `page_size`: [integer](#integer): 
    <!--beg l desc_ListSubmission_page_size -->
    
    <!--end l-->


+ `mem_order`: [boolean](#boolean): 
    <!--beg l desc_ListSubmission_mem_order -->
    
    <!--end l-->


+ `time_order`: [boolean](#boolean): 
    <!--beg l desc_ListSubmission_time_order -->
    
    <!--end l-->


+ `id_order`: [boolean](#boolean): 
    <!--beg l desc_ListSubmission_id_order -->
    
    <!--end l-->


+ `by_user`: [integer](#integer): 
    <!--beg l desc_ListSubmission_by_user -->
    
    <!--end l-->


+ `on_problem`: [integer](#integer): 
    <!--beg l desc_ListSubmission_on_problem -->
    
    <!--end l-->


+ `with_language`: [integer](#integer): 
    <!--beg l desc_ListSubmission_with_language -->
    
    <!--end l-->


+ `has_status`: [integer](#integer): 
    <!--beg l desc_ListSubmission_has_status -->
    
    <!--end l-->



### GetSubmissionContent

The uri/restful key of this method is `/submission/{sid}/content@GET`

<!--beg l desc_GetSubmissionContent -->

<!--end l-->


+ `sid`: [string](#string) (required): 
    <!--beg l desc_GetSubmissionContent_sid -->
    
    <!--end l-->



### DeleteSubmission

The uri/restful key of this method is `/submission/{sid}@DELETE`

<!--beg l desc_DeleteSubmission -->

<!--end l-->


+ `sid`: [string](#string) (required): 
    <!--beg l desc_DeleteSubmission_sid -->
    
    <!--end l-->


+ `DeleteSubmissionRequest`: [DeleteSubmissionRequest](#DeleteSubmissionRequest): 
    <!--beg l desc_DeleteSubmission_DeleteSubmissionRequest -->
    
    <!--end l-->



### GetSubmission

The uri/restful key of this method is `/submission/{sid}@GET`

<!--beg l desc_GetSubmission -->

<!--end l-->


+ `sid`: [string](#string) (required): 
    <!--beg l desc_GetSubmission_sid -->
    
    <!--end l-->



## Local Object Reference




### [PostSubmissionRequest](./ObjectModelSpec.md#PostSubmissionRequest)

+ type: [object](#PostSubmissionRequest)

+ fields:
    
    + `code`: string: 
        <!--beg l desc_{{object_name}}_code -->
        
        <!--end l-->

    + `information`: string: 
        <!--beg l desc_{{object_name}}_information -->
        
        <!--end l-->

    + `language`: string: 
        <!--beg l desc_{{object_name}}_language -->
        
        <!--end l-->

    + `pid`: integer: 
        <!--beg l desc_{{object_name}}_pid -->
        
        <!--end l-->

    + `shared`: integer: 
        <!--beg l desc_{{object_name}}_shared -->
        
        <!--end l-->

    
### [DeleteSubmissionRequest](./ObjectModelSpec.md#DeleteSubmissionRequest)

+ type: [object](#DeleteSubmissionRequest)

+ fields:
    
    

