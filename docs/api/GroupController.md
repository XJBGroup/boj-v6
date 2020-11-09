
# Api Group Documentation (GroupController)

<!--beg l desc_GroupController -->

<!--end l-->

## Apis


### CountGroup

The uri/restful key of this method is `/group-count@GET`

<!--beg l desc_CountGroup -->

<!--end l-->


+ `order`: [string](#string): 
    <!--beg l desc_CountGroup_order -->
    
    <!--end l-->


+ `page`: [integer](#integer): 
    <!--beg l desc_CountGroup_page -->
    
    <!--end l-->


+ `page_size`: [integer](#integer): 
    <!--beg l desc_CountGroup_page_size -->
    
    <!--end l-->


+ `before_id`: [integer](#integer): 
    <!--beg l desc_CountGroup_before_id -->
    
    <!--end l-->



### ListGroup

The uri/restful key of this method is `/group-list@GET`

<!--beg l desc_ListGroup -->

<!--end l-->


+ `before_id`: [integer](#integer): 
    <!--beg l desc_ListGroup_before_id -->
    
    <!--end l-->


+ `order`: [string](#string): 
    <!--beg l desc_ListGroup_order -->
    
    <!--end l-->


+ `page`: [integer](#integer): 
    <!--beg l desc_ListGroup_page -->
    
    <!--end l-->


+ `page_size`: [integer](#integer): 
    <!--beg l desc_ListGroup_page_size -->
    
    <!--end l-->



### PutGroupOwner

The uri/restful key of this method is `/group/{gid}/owner@PUT`

<!--beg l desc_PutGroupOwner -->

<!--end l-->


+ `gid`: [string](#string) (required): 
    <!--beg l desc_PutGroupOwner_gid -->
    
    <!--end l-->


+ `PutGroupOwnerRequest`: [PutGroupOwnerRequest](#PutGroupOwnerRequest): 
    <!--beg l desc_PutGroupOwner_PutGroupOwnerRequest -->
    
    <!--end l-->



### GetGroupMembers

The uri/restful key of this method is `/group/{gid}/user-list@GET`

<!--beg l desc_GetGroupMembers -->

<!--end l-->


+ `gid`: [string](#string) (required): 
    <!--beg l desc_GetGroupMembers_gid -->
    
    <!--end l-->


+ `order`: [string](#string): 
    <!--beg l desc_GetGroupMembers_order -->
    
    <!--end l-->


+ `page`: [integer](#integer): 
    <!--beg l desc_GetGroupMembers_page -->
    
    <!--end l-->


+ `page_size`: [integer](#integer): 
    <!--beg l desc_GetGroupMembers_page_size -->
    
    <!--end l-->


+ `before_id`: [integer](#integer): 
    <!--beg l desc_GetGroupMembers_before_id -->
    
    <!--end l-->



### PostGroupMember

The uri/restful key of this method is `/group/{gid}/user/{id}@POST`

<!--beg l desc_PostGroupMember -->

<!--end l-->


+ `gid`: [string](#string) (required): 
    <!--beg l desc_PostGroupMember_gid -->
    
    <!--end l-->


+ `id`: [string](#string) (required): 
    <!--beg l desc_PostGroupMember_id -->
    
    <!--end l-->


+ `PostGroupMemberRequest`: [PostGroupMemberRequest](#PostGroupMemberRequest): 
    <!--beg l desc_PostGroupMember_PostGroupMemberRequest -->
    
    <!--end l-->



### DeleteGroup

The uri/restful key of this method is `/group/{gid}@DELETE`

<!--beg l desc_DeleteGroup -->

<!--end l-->


+ `gid`: [string](#string) (required): 
    <!--beg l desc_DeleteGroup_gid -->
    
    <!--end l-->


+ `DeleteGroupRequest`: [DeleteGroupRequest](#DeleteGroupRequest): 
    <!--beg l desc_DeleteGroup_DeleteGroupRequest -->
    
    <!--end l-->



### GetGroup

The uri/restful key of this method is `/group/{gid}@GET`

<!--beg l desc_GetGroup -->

<!--end l-->


+ `gid`: [string](#string) (required): 
    <!--beg l desc_GetGroup_gid -->
    
    <!--end l-->



### PutGroup

The uri/restful key of this method is `/group/{gid}@PUT`

<!--beg l desc_PutGroup -->

<!--end l-->


+ `gid`: [string](#string) (required): 
    <!--beg l desc_PutGroup_gid -->
    
    <!--end l-->


+ `PutGroupRequest`: [PutGroupRequest](#PutGroupRequest): 
    <!--beg l desc_PutGroup_PutGroupRequest -->
    
    <!--end l-->



### PostGroup

The uri/restful key of this method is `/group@POST`

<!--beg l desc_PostGroup -->

<!--end l-->


+ `PostGroupRequest`: [PostGroupRequest](#PostGroupRequest): 
    <!--beg l desc_PostGroup_PostGroupRequest -->
    
    <!--end l-->



## Local Object Reference




### [PutGroupOwnerRequest](./ObjectModelSpec.md#PutGroupOwnerRequest)

+ type: [object](#PutGroupOwnerRequest)

+ fields:
    
    + `owner_id`: integer: 
        <!--beg l desc_{{object_name}}_owner_id -->
        
        <!--end l-->

    
### [PostGroupMemberRequest](./ObjectModelSpec.md#PostGroupMemberRequest)

+ type: [object](#PostGroupMemberRequest)

+ fields:
    
    
### [DeleteGroupRequest](./ObjectModelSpec.md#DeleteGroupRequest)

+ type: [object](#DeleteGroupRequest)

+ fields:
    
    
### [PutGroupRequest](./ObjectModelSpec.md#PutGroupRequest)

+ type: [object](#PutGroupRequest)

+ fields:
    
    + `description`: string: 
        <!--beg l desc_{{object_name}}_description -->
        
        <!--end l-->

    + `name`: string: 
        <!--beg l desc_{{object_name}}_name -->
        
        <!--end l-->

    
### [PostGroupRequest](./ObjectModelSpec.md#PostGroupRequest)

+ type: [object](#PostGroupRequest)

+ fields:
    
    + `description`: string: 
        <!--beg l desc_{{object_name}}_description -->
        
        <!--end l-->

    + `name`: string: 
        <!--beg l desc_{{object_name}}_name -->
        
        <!--end l-->

    + `owner_id`: integer: 
        <!--beg l desc_{{object_name}}_owner_id -->
        
        <!--end l-->

    + `owner_name`: string: 
        <!--beg l desc_{{object_name}}_owner_name -->
        
        <!--end l-->

    

