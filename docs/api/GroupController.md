
# Api Group Documentation (GroupController)

<!--beg l desc_GroupController -->

<!--end l-->


## CountGroup

The uri/restful key of this method is `/group-count@GET`

<!--beg l desc_CountGroup -->

<!--end l-->


+ `order`: `string`: 
    <!--beg l desc_CountGroup_order -->
    
    <!--end l-->

+ `page`: `integer`: 
    <!--beg l desc_CountGroup_page -->
    
    <!--end l-->

+ `page_size`: `integer`: 
    <!--beg l desc_CountGroup_page_size -->
    
    <!--end l-->

+ `before_id`: `integer`: 
    <!--beg l desc_CountGroup_before_id -->
    
    <!--end l-->



## ListGroup

The uri/restful key of this method is `/group-list@GET`

<!--beg l desc_ListGroup -->

<!--end l-->


+ `before_id`: `integer`: 
    <!--beg l desc_ListGroup_before_id -->
    
    <!--end l-->

+ `order`: `string`: 
    <!--beg l desc_ListGroup_order -->
    
    <!--end l-->

+ `page`: `integer`: 
    <!--beg l desc_ListGroup_page -->
    
    <!--end l-->

+ `page_size`: `integer`: 
    <!--beg l desc_ListGroup_page_size -->
    
    <!--end l-->



## PutGroupOwner

The uri/restful key of this method is `/group/{gid}/owner@PUT`

<!--beg l desc_PutGroupOwner -->

<!--end l-->


+ `gid`: `string` (required): 
    <!--beg l desc_PutGroupOwner_gid -->
    
    <!--end l-->

+ `PutGroupOwnerRequest`: `any`: 
    <!--beg l desc_PutGroupOwner_PutGroupOwnerRequest -->
    
    <!--end l-->



## GetGroupMembers

The uri/restful key of this method is `/group/{gid}/user-list@GET`

<!--beg l desc_GetGroupMembers -->

<!--end l-->


+ `gid`: `string` (required): 
    <!--beg l desc_GetGroupMembers_gid -->
    
    <!--end l-->

+ `order`: `string`: 
    <!--beg l desc_GetGroupMembers_order -->
    
    <!--end l-->

+ `page`: `integer`: 
    <!--beg l desc_GetGroupMembers_page -->
    
    <!--end l-->

+ `page_size`: `integer`: 
    <!--beg l desc_GetGroupMembers_page_size -->
    
    <!--end l-->

+ `before_id`: `integer`: 
    <!--beg l desc_GetGroupMembers_before_id -->
    
    <!--end l-->



## PostGroupMember

The uri/restful key of this method is `/group/{gid}/user/{id}@POST`

<!--beg l desc_PostGroupMember -->

<!--end l-->


+ `gid`: `string` (required): 
    <!--beg l desc_PostGroupMember_gid -->
    
    <!--end l-->

+ `id`: `string` (required): 
    <!--beg l desc_PostGroupMember_id -->
    
    <!--end l-->

+ `PostGroupMemberRequest`: `any`: 
    <!--beg l desc_PostGroupMember_PostGroupMemberRequest -->
    
    <!--end l-->



## DeleteGroup

The uri/restful key of this method is `/group/{gid}@DELETE`

<!--beg l desc_DeleteGroup -->

<!--end l-->


+ `gid`: `string` (required): 
    <!--beg l desc_DeleteGroup_gid -->
    
    <!--end l-->

+ `DeleteGroupRequest`: `any`: 
    <!--beg l desc_DeleteGroup_DeleteGroupRequest -->
    
    <!--end l-->



## GetGroup

The uri/restful key of this method is `/group/{gid}@GET`

<!--beg l desc_GetGroup -->

<!--end l-->


+ `gid`: `string` (required): 
    <!--beg l desc_GetGroup_gid -->
    
    <!--end l-->



## PutGroup

The uri/restful key of this method is `/group/{gid}@PUT`

<!--beg l desc_PutGroup -->

<!--end l-->


+ `gid`: `string` (required): 
    <!--beg l desc_PutGroup_gid -->
    
    <!--end l-->

+ `PutGroupRequest`: `any`: 
    <!--beg l desc_PutGroup_PutGroupRequest -->
    
    <!--end l-->



## PostGroup

The uri/restful key of this method is `/group@POST`

<!--beg l desc_PostGroup -->

<!--end l-->


+ `PostGroupRequest`: `any`: 
    <!--beg l desc_PostGroup_PostGroupRequest -->
    
    <!--end l-->




