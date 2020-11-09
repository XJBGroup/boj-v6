
# Api Group Documentation (SubmissionController)

<!--beg l desc_SubmissionController -->

<!--end l-->

## Apis


### PostSubmission

restful key: `/problem/{pid}/submission@POST`

<!--beg l desc_PostSubmission -->

<!--end l-->

parameters:

+ `pid`: [string](#string) (required): 
    <!--beg l desc_PostSubmission_params_pid -->
    
    <!--end l-->


+ `PostSubmissionRequest`: [PostSubmissionRequest](#PostSubmissionRequest): 
    <!--beg l desc_PostSubmission_params_PostSubmissionRequest -->
    
    <!--end l-->


responses:

+ code: `200`, type: [genericResponse](#genericResponse)
    <!--beg l desc_PostSubmission_response_200_[genericResponse](#genericResponse) -->
    
    <!--end l-->


+ code: `200`, type: [PostSubmissionReply](#PostSubmissionReply)
    <!--beg l desc_PostSubmission_response_200_[PostSubmissionReply](#PostSubmissionReply) -->
    
    <!--end l-->


+ code: `500`, type: No Response
    <!--beg l desc_PostSubmission_response_500_No Response -->
    
    <!--end l-->




### CountSubmission

restful key: `/submission-count@GET`

<!--beg l desc_CountSubmission -->

<!--end l-->

parameters:

+ `page`: [integer](#integer): 
    <!--beg l desc_CountSubmission_params_page -->
    
    <!--end l-->


+ `page_size`: [integer](#integer): 
    <!--beg l desc_CountSubmission_params_page_size -->
    
    <!--end l-->


+ `mem_order`: [boolean](#boolean): 
    <!--beg l desc_CountSubmission_params_mem_order -->
    
    <!--end l-->


+ `time_order`: [boolean](#boolean): 
    <!--beg l desc_CountSubmission_params_time_order -->
    
    <!--end l-->


+ `id_order`: [boolean](#boolean): 
    <!--beg l desc_CountSubmission_params_id_order -->
    
    <!--end l-->


+ `by_user`: [integer](#integer): 
    <!--beg l desc_CountSubmission_params_by_user -->
    
    <!--end l-->


+ `on_problem`: [integer](#integer): 
    <!--beg l desc_CountSubmission_params_on_problem -->
    
    <!--end l-->


+ `with_language`: [integer](#integer): 
    <!--beg l desc_CountSubmission_params_with_language -->
    
    <!--end l-->


+ `has_status`: [integer](#integer): 
    <!--beg l desc_CountSubmission_params_has_status -->
    
    <!--end l-->


responses:

+ code: `200`, type: [genericResponse](#genericResponse)
    <!--beg l desc_CountSubmission_response_200_[genericResponse](#genericResponse) -->
    
    <!--end l-->


+ code: `200`, type: [CountSubmissionReply](#CountSubmissionReply)
    <!--beg l desc_CountSubmission_response_200_[CountSubmissionReply](#CountSubmissionReply) -->
    
    <!--end l-->


+ code: `500`, type: No Response
    <!--beg l desc_CountSubmission_response_500_No Response -->
    
    <!--end l-->




### ListSubmission

restful key: `/submission-list@GET`

<!--beg l desc_ListSubmission -->

<!--end l-->

parameters:

+ `page`: [integer](#integer): 
    <!--beg l desc_ListSubmission_params_page -->
    
    <!--end l-->


+ `page_size`: [integer](#integer): 
    <!--beg l desc_ListSubmission_params_page_size -->
    
    <!--end l-->


+ `mem_order`: [boolean](#boolean): 
    <!--beg l desc_ListSubmission_params_mem_order -->
    
    <!--end l-->


+ `time_order`: [boolean](#boolean): 
    <!--beg l desc_ListSubmission_params_time_order -->
    
    <!--end l-->


+ `id_order`: [boolean](#boolean): 
    <!--beg l desc_ListSubmission_params_id_order -->
    
    <!--end l-->


+ `by_user`: [integer](#integer): 
    <!--beg l desc_ListSubmission_params_by_user -->
    
    <!--end l-->


+ `on_problem`: [integer](#integer): 
    <!--beg l desc_ListSubmission_params_on_problem -->
    
    <!--end l-->


+ `with_language`: [integer](#integer): 
    <!--beg l desc_ListSubmission_params_with_language -->
    
    <!--end l-->


+ `has_status`: [integer](#integer): 
    <!--beg l desc_ListSubmission_params_has_status -->
    
    <!--end l-->


responses:

+ code: `200`, type: [genericResponse](#genericResponse)
    <!--beg l desc_ListSubmission_response_200_[genericResponse](#genericResponse) -->
    
    <!--end l-->


+ code: `200`, type: [ListSubmissionReply](#ListSubmissionReply)
    <!--beg l desc_ListSubmission_response_200_[ListSubmissionReply](#ListSubmissionReply) -->
    
    <!--end l-->


+ code: `500`, type: No Response
    <!--beg l desc_ListSubmission_response_500_No Response -->
    
    <!--end l-->




### GetSubmissionContent

restful key: `/submission/{sid}/content@GET`

<!--beg l desc_GetSubmissionContent -->

<!--end l-->

parameters:

+ `sid`: [string](#string) (required): 
    <!--beg l desc_GetSubmissionContent_params_sid -->
    
    <!--end l-->


responses:

+ code: `200`, type: [genericResponse](#genericResponse)
    <!--beg l desc_GetSubmissionContent_response_200_[genericResponse](#genericResponse) -->
    
    <!--end l-->


+ code: `200`, type: [GetSubmissionContentReply](#GetSubmissionContentReply)
    <!--beg l desc_GetSubmissionContent_response_200_[GetSubmissionContentReply](#GetSubmissionContentReply) -->
    
    <!--end l-->


+ code: `500`, type: No Response
    <!--beg l desc_GetSubmissionContent_response_500_No Response -->
    
    <!--end l-->




### DeleteSubmission

restful key: `/submission/{sid}@DELETE`

<!--beg l desc_DeleteSubmission -->

<!--end l-->

parameters:

+ `sid`: [string](#string) (required): 
    <!--beg l desc_DeleteSubmission_params_sid -->
    
    <!--end l-->


+ `DeleteSubmissionRequest`: [DeleteSubmissionRequest](#DeleteSubmissionRequest): 
    <!--beg l desc_DeleteSubmission_params_DeleteSubmissionRequest -->
    
    <!--end l-->


responses:

+ code: `200`, type: [genericResponse](#genericResponse)
    <!--beg l desc_DeleteSubmission_response_200_[genericResponse](#genericResponse) -->
    
    <!--end l-->


+ code: `200`, type: [DeleteSubmissionReply](#DeleteSubmissionReply)
    <!--beg l desc_DeleteSubmission_response_200_[DeleteSubmissionReply](#DeleteSubmissionReply) -->
    
    <!--end l-->


+ code: `500`, type: No Response
    <!--beg l desc_DeleteSubmission_response_500_No Response -->
    
    <!--end l-->




### GetSubmission

restful key: `/submission/{sid}@GET`

<!--beg l desc_GetSubmission -->

<!--end l-->

parameters:

+ `sid`: [string](#string) (required): 
    <!--beg l desc_GetSubmission_params_sid -->
    
    <!--end l-->


responses:

+ code: `200`, type: [genericResponse](#genericResponse)
    <!--beg l desc_GetSubmission_response_200_[genericResponse](#genericResponse) -->
    
    <!--end l-->


+ code: `200`, type: [GetSubmissionReply](#GetSubmissionReply)
    <!--beg l desc_GetSubmission_response_200_[GetSubmissionReply](#GetSubmissionReply) -->
    
    <!--end l-->


+ code: `500`, type: No Response
    <!--beg l desc_GetSubmission_response_500_No Response -->
    
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

    
### [PostSubmissionReply](./ObjectModelSpec.md#PostSubmissionReply)

+ type: [object](#PostSubmissionReply)

+ fields:
    
    + `code`: integer: 
        <!--beg l desc_{{object_name}}_code -->
        
        <!--end l-->

    + `data`: : 
        <!--beg l desc_{{object_name}}_data -->
        
        <!--end l-->

    
### [CountSubmissionReply](./ObjectModelSpec.md#CountSubmissionReply)

+ type: [object](#CountSubmissionReply)

+ fields:
    
    + `code`: integer: 
        <!--beg l desc_{{object_name}}_code -->
        
        <!--end l-->

    + `data`: integer: 
        <!--beg l desc_{{object_name}}_data -->
        
        <!--end l-->

    
### [ListSubmissionReply](./ObjectModelSpec.md#ListSubmissionReply)

+ type: [object](#ListSubmissionReply)

+ fields:
    
    + `code`: integer: 
        <!--beg l desc_{{object_name}}_code -->
        
        <!--end l-->

    + `data`: array: 
        <!--beg l desc_{{object_name}}_data -->
        
        <!--end l-->

    
### [GetSubmissionContentReply](./ObjectModelSpec.md#GetSubmissionContentReply)

+ type: [object](#GetSubmissionContentReply)

+ fields:
    
    + `code`: integer: 
        <!--beg l desc_{{object_name}}_code -->
        
        <!--end l-->

    
### [DeleteSubmissionRequest](./ObjectModelSpec.md#DeleteSubmissionRequest)

+ type: [object](#DeleteSubmissionRequest)

+ fields:
    
    
### [DeleteSubmissionReply](./ObjectModelSpec.md#DeleteSubmissionReply)

+ type: [object](#DeleteSubmissionReply)

+ fields:
    
    + `code`: integer: 
        <!--beg l desc_{{object_name}}_code -->
        
        <!--end l-->

    
### [GetSubmissionReply](./ObjectModelSpec.md#GetSubmissionReply)

+ type: [object](#GetSubmissionReply)

+ fields:
    
    + `code`: integer: 
        <!--beg l desc_{{object_name}}_code -->
        
        <!--end l-->

    + `data`: : 
        <!--beg l desc_{{object_name}}_data -->
        
        <!--end l-->

    

