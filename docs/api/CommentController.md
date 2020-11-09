
# Api Group Documentation (CommentController)

<!--beg l desc_CommentController -->

<!--end l-->

## Apis


### CountComment

restful key: `/comment-count@GET`

<!--beg l desc_CountComment -->

<!--end l-->

parameters:

+ `ref`: [integer](#integer): 
    <!--beg l desc_CountComment_params_ref -->
    
    <!--end l-->


+ `no_reply`: [boolean](#boolean): 
    <!--beg l desc_CountComment_params_no_reply -->
    
    <!--end l-->


+ `page`: [integer](#integer): 
    <!--beg l desc_CountComment_params_page -->
    
    <!--end l-->


+ `page_size`: [integer](#integer): 
    <!--beg l desc_CountComment_params_page_size -->
    
    <!--end l-->


+ `ref_type`: [integer](#integer): 
    <!--beg l desc_CountComment_params_ref_type -->
    
    <!--end l-->


responses:

+ code: `200`, type: [genericResponse](#genericResponse)
    <!--beg l desc_CountComment_response_200_[genericResponse](#genericResponse) -->
    
    <!--end l-->


+ code: `200`, type: [CountCommentReply](#CountCommentReply)
    <!--beg l desc_CountComment_response_200_[CountCommentReply](#CountCommentReply) -->
    
    <!--end l-->


+ code: `500`, type: No Response
    <!--beg l desc_CountComment_response_500_No Response -->
    
    <!--end l-->




### ListComment

restful key: `/comment-list@GET`

<!--beg l desc_ListComment -->

<!--end l-->

parameters:

+ `page_size`: [integer](#integer): 
    <!--beg l desc_ListComment_params_page_size -->
    
    <!--end l-->


+ `ref_type`: [integer](#integer): 
    <!--beg l desc_ListComment_params_ref_type -->
    
    <!--end l-->


+ `ref`: [integer](#integer): 
    <!--beg l desc_ListComment_params_ref -->
    
    <!--end l-->


+ `no_reply`: [boolean](#boolean): 
    <!--beg l desc_ListComment_params_no_reply -->
    
    <!--end l-->


+ `page`: [integer](#integer): 
    <!--beg l desc_ListComment_params_page -->
    
    <!--end l-->


responses:

+ code: `200`, type: [genericResponse](#genericResponse)
    <!--beg l desc_ListComment_response_200_[genericResponse](#genericResponse) -->
    
    <!--end l-->


+ code: `200`, type: [ListCommentReply](#ListCommentReply)
    <!--beg l desc_ListComment_response_200_[ListCommentReply](#ListCommentReply) -->
    
    <!--end l-->


+ code: `500`, type: No Response
    <!--beg l desc_ListComment_response_500_No Response -->
    
    <!--end l-->




### DeleteComment

restful key: `/comment/{cmid}@DELETE`

<!--beg l desc_DeleteComment -->

<!--end l-->

parameters:

+ `cmid`: [string](#string) (required): 
    <!--beg l desc_DeleteComment_params_cmid -->
    
    <!--end l-->


+ `DeleteCommentRequest`: [DeleteCommentRequest](#DeleteCommentRequest): 
    <!--beg l desc_DeleteComment_params_DeleteCommentRequest -->
    
    <!--end l-->


responses:

+ code: `200`, type: [genericResponse](#genericResponse)
    <!--beg l desc_DeleteComment_response_200_[genericResponse](#genericResponse) -->
    
    <!--end l-->


+ code: `200`, type: [DeleteCommentReply](#DeleteCommentReply)
    <!--beg l desc_DeleteComment_response_200_[DeleteCommentReply](#DeleteCommentReply) -->
    
    <!--end l-->


+ code: `500`, type: No Response
    <!--beg l desc_DeleteComment_response_500_No Response -->
    
    <!--end l-->




### GetComment

restful key: `/comment/{cmid}@GET`

<!--beg l desc_GetComment -->

<!--end l-->

parameters:

+ `cmid`: [string](#string) (required): 
    <!--beg l desc_GetComment_params_cmid -->
    
    <!--end l-->


responses:

+ code: `200`, type: [genericResponse](#genericResponse)
    <!--beg l desc_GetComment_response_200_[genericResponse](#genericResponse) -->
    
    <!--end l-->


+ code: `200`, type: [GetCommentReply](#GetCommentReply)
    <!--beg l desc_GetComment_response_200_[GetCommentReply](#GetCommentReply) -->
    
    <!--end l-->


+ code: `500`, type: No Response
    <!--beg l desc_GetComment_response_500_No Response -->
    
    <!--end l-->




### PutComment

restful key: `/comment/{cmid}@PUT`

<!--beg l desc_PutComment -->

<!--end l-->

parameters:

+ `cmid`: [string](#string) (required): 
    <!--beg l desc_PutComment_params_cmid -->
    
    <!--end l-->


+ `PutCommentRequest`: [PutCommentRequest](#PutCommentRequest): 
    <!--beg l desc_PutComment_params_PutCommentRequest -->
    
    <!--end l-->


responses:

+ code: `200`, type: [genericResponse](#genericResponse)
    <!--beg l desc_PutComment_response_200_[genericResponse](#genericResponse) -->
    
    <!--end l-->


+ code: `200`, type: [PutCommentReply](#PutCommentReply)
    <!--beg l desc_PutComment_response_200_[PutCommentReply](#PutCommentReply) -->
    
    <!--end l-->


+ code: `500`, type: No Response
    <!--beg l desc_PutComment_response_500_No Response -->
    
    <!--end l-->




### PostComment

restful key: `/comment@POST`

<!--beg l desc_PostComment -->

<!--end l-->

parameters:

+ `PostCommentRequest`: [PostCommentRequest](#PostCommentRequest): 
    <!--beg l desc_PostComment_params_PostCommentRequest -->
    
    <!--end l-->


responses:

+ code: `200`, type: [genericResponse](#genericResponse)
    <!--beg l desc_PostComment_response_200_[genericResponse](#genericResponse) -->
    
    <!--end l-->


+ code: `200`, type: [PostCommentReply](#PostCommentReply)
    <!--beg l desc_PostComment_response_200_[PostCommentReply](#PostCommentReply) -->
    
    <!--end l-->


+ code: `500`, type: No Response
    <!--beg l desc_PostComment_response_500_No Response -->
    
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

    
### [CountCommentReply](./ObjectModelSpec.md#CountCommentReply)

+ type: [object](#CountCommentReply)

+ fields:
    
    + `code`: integer: 
        <!--beg l desc_{{object_name}}_code -->
        
        <!--end l-->

    + `data`: integer: 
        <!--beg l desc_{{object_name}}_data -->
        
        <!--end l-->

    
### [ListCommentReply](./ObjectModelSpec.md#ListCommentReply)

+ type: [object](#ListCommentReply)

+ fields:
    
    + `code`: integer: 
        <!--beg l desc_{{object_name}}_code -->
        
        <!--end l-->

    + `data`: array: 
        <!--beg l desc_{{object_name}}_data -->
        
        <!--end l-->

    
### [DeleteCommentRequest](./ObjectModelSpec.md#DeleteCommentRequest)

+ type: [object](#DeleteCommentRequest)

+ fields:
    
    
### [DeleteCommentReply](./ObjectModelSpec.md#DeleteCommentReply)

+ type: [object](#DeleteCommentReply)

+ fields:
    
    + `code`: integer: 
        <!--beg l desc_{{object_name}}_code -->
        
        <!--end l-->

    
### [GetCommentReply](./ObjectModelSpec.md#GetCommentReply)

+ type: [object](#GetCommentReply)

+ fields:
    
    + `code`: integer: 
        <!--beg l desc_{{object_name}}_code -->
        
        <!--end l-->

    + `data`: : 
        <!--beg l desc_{{object_name}}_data -->
        
        <!--end l-->

    
### [PutCommentRequest](./ObjectModelSpec.md#PutCommentRequest)

+ type: [object](#PutCommentRequest)

+ fields:
    
    + `content`: string: 
        <!--beg l desc_{{object_name}}_content -->
        
        <!--end l-->

    + `title`: string: 
        <!--beg l desc_{{object_name}}_title -->
        
        <!--end l-->

    
### [PutCommentReply](./ObjectModelSpec.md#PutCommentReply)

+ type: [object](#PutCommentReply)

+ fields:
    
    + `code`: integer: 
        <!--beg l desc_{{object_name}}_code -->
        
        <!--end l-->

    
### [PostCommentRequest](./ObjectModelSpec.md#PostCommentRequest)

+ type: [object](#PostCommentRequest)

+ fields:
    
    + `content`: string: 
        <!--beg l desc_{{object_name}}_content -->
        
        <!--end l-->

    + `title`: string: 
        <!--beg l desc_{{object_name}}_title -->
        
        <!--end l-->

    
### [PostCommentReply](./ObjectModelSpec.md#PostCommentReply)

+ type: [object](#PostCommentReply)

+ fields:
    
    + `code`: integer: 
        <!--beg l desc_{{object_name}}_code -->
        
        <!--end l-->

    + `comment`: : 
        <!--beg l desc_{{object_name}}_comment -->
        
        <!--end l-->

    

