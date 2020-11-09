
# Api Group Documentation (CommentController)


## CountComment

The uri/restful key of this method is `/comment-count@GET`

<!--beg l desc_CountComment -->

<!--end l-->


+ `ref`: `integer`: 
    <!--beg l desc_CountComment_ref -->
    
    <!--end l-->

+ `no_reply`: `boolean`: 
    <!--beg l desc_CountComment_no_reply -->
    
    <!--end l-->

+ `page`: `integer`: 
    <!--beg l desc_CountComment_page -->
    
    <!--end l-->

+ `page_size`: `integer`: 
    <!--beg l desc_CountComment_page_size -->
    
    <!--end l-->

+ `ref_type`: `integer`: 
    <!--beg l desc_CountComment_ref_type -->
    
    <!--end l-->



## ListComment

The uri/restful key of this method is `/comment-list@GET`

<!--beg l desc_ListComment -->

<!--end l-->


+ `page_size`: `integer`: 
    <!--beg l desc_ListComment_page_size -->
    
    <!--end l-->

+ `ref_type`: `integer`: 
    <!--beg l desc_ListComment_ref_type -->
    
    <!--end l-->

+ `ref`: `integer`: 
    <!--beg l desc_ListComment_ref -->
    
    <!--end l-->

+ `no_reply`: `boolean`: 
    <!--beg l desc_ListComment_no_reply -->
    
    <!--end l-->

+ `page`: `integer`: 
    <!--beg l desc_ListComment_page -->
    
    <!--end l-->



## DeleteComment

The uri/restful key of this method is `/comment/{cmid}@DELETE`

<!--beg l desc_DeleteComment -->

<!--end l-->


+ `cmid`: `string` (required): 
    <!--beg l desc_DeleteComment_cmid -->
    
    <!--end l-->

+ `DeleteCommentRequest`: `any`: 
    <!--beg l desc_DeleteComment_DeleteCommentRequest -->
    
    <!--end l-->



## GetComment

The uri/restful key of this method is `/comment/{cmid}@GET`

<!--beg l desc_GetComment -->

<!--end l-->


+ `cmid`: `string` (required): 
    <!--beg l desc_GetComment_cmid -->
    
    <!--end l-->



## PutComment

The uri/restful key of this method is `/comment/{cmid}@PUT`

<!--beg l desc_PutComment -->

<!--end l-->


+ `cmid`: `string` (required): 
    <!--beg l desc_PutComment_cmid -->
    
    <!--end l-->

+ `PutCommentRequest`: `any`: 
    <!--beg l desc_PutComment_PutCommentRequest -->
    
    <!--end l-->



## PostComment

The uri/restful key of this method is `/comment@POST`

<!--beg l desc_PostComment -->

<!--end l-->


+ `PostCommentRequest`: `any`: 
    <!--beg l desc_PostComment_PostCommentRequest -->
    
    <!--end l-->



