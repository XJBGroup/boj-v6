
# Api Group Documentation (GroupController)

<!--beg l desc_GroupController -->

<!--end l-->

## Apis


### CountGroup

The uri/restful key of this method is `/group-count@GET`

<!--beg l desc_CountGroup -->

<!--end l-->

parameters:

+ `order`: [string](#string): 
    <!--beg l desc_CountGroup_params_order -->
    
    <!--end l-->


+ `page`: [integer](#integer): 
    <!--beg l desc_CountGroup_params_page -->
    
    <!--end l-->


+ `page_size`: [integer](#integer): 
    <!--beg l desc_CountGroup_params_page_size -->
    
    <!--end l-->


+ `before_id`: [integer](#integer): 
    <!--beg l desc_CountGroup_params_before_id -->
    
    <!--end l-->


responses:

+ code: `200`, type: [genericResponse](#genericResponse)
    <!--beg l desc_CountGroup_response_200_[genericResponse](#genericResponse) -->
    
    <!--end l-->


+ code: `200`, type: [CountGroupReply](#CountGroupReply)
    <!--beg l desc_CountGroup_response_200_[CountGroupReply](#CountGroupReply) -->
    
    <!--end l-->


+ code: `500`, type: No Response
    <!--beg l desc_CountGroup_response_500_No Response -->
    
    <!--end l-->




### ListGroup

The uri/restful key of this method is `/group-list@GET`

<!--beg l desc_ListGroup -->

<!--end l-->

parameters:

+ `before_id`: [integer](#integer): 
    <!--beg l desc_ListGroup_params_before_id -->
    
    <!--end l-->


+ `order`: [string](#string): 
    <!--beg l desc_ListGroup_params_order -->
    
    <!--end l-->


+ `page`: [integer](#integer): 
    <!--beg l desc_ListGroup_params_page -->
    
    <!--end l-->


+ `page_size`: [integer](#integer): 
    <!--beg l desc_ListGroup_params_page_size -->
    
    <!--end l-->


responses:

+ code: `200`, type: [genericResponse](#genericResponse)
    <!--beg l desc_ListGroup_response_200_[genericResponse](#genericResponse) -->
    
    <!--end l-->


+ code: `200`, type: [ListGroupReply](#ListGroupReply)
    <!--beg l desc_ListGroup_response_200_[ListGroupReply](#ListGroupReply) -->
    
    <!--end l-->


+ code: `500`, type: No Response
    <!--beg l desc_ListGroup_response_500_No Response -->
    
    <!--end l-->




### PutGroupOwner

The uri/restful key of this method is `/group/{gid}/owner@PUT`

<!--beg l desc_PutGroupOwner -->

<!--end l-->

parameters:

+ `gid`: [string](#string) (required): 
    <!--beg l desc_PutGroupOwner_params_gid -->
    
    <!--end l-->


+ `PutGroupOwnerRequest`: [PutGroupOwnerRequest](#PutGroupOwnerRequest): 
    <!--beg l desc_PutGroupOwner_params_PutGroupOwnerRequest -->
    
    <!--end l-->


responses:

+ code: `200`, type: [genericResponse](#genericResponse)
    <!--beg l desc_PutGroupOwner_response_200_[genericResponse](#genericResponse) -->
    
    <!--end l-->


+ code: `200`, type: [PutGroupOwnerReply](#PutGroupOwnerReply)
    <!--beg l desc_PutGroupOwner_response_200_[PutGroupOwnerReply](#PutGroupOwnerReply) -->
    
    <!--end l-->


+ code: `500`, type: No Response
    <!--beg l desc_PutGroupOwner_response_500_No Response -->
    
    <!--end l-->




### GetGroupMembers

The uri/restful key of this method is `/group/{gid}/user-list@GET`

<!--beg l desc_GetGroupMembers -->

<!--end l-->

parameters:

+ `gid`: [string](#string) (required): 
    <!--beg l desc_GetGroupMembers_params_gid -->
    
    <!--end l-->


+ `order`: [string](#string): 
    <!--beg l desc_GetGroupMembers_params_order -->
    
    <!--end l-->


+ `page`: [integer](#integer): 
    <!--beg l desc_GetGroupMembers_params_page -->
    
    <!--end l-->


+ `page_size`: [integer](#integer): 
    <!--beg l desc_GetGroupMembers_params_page_size -->
    
    <!--end l-->


+ `before_id`: [integer](#integer): 
    <!--beg l desc_GetGroupMembers_params_before_id -->
    
    <!--end l-->


responses:

+ code: `200`, type: [genericResponse](#genericResponse)
    <!--beg l desc_GetGroupMembers_response_200_[genericResponse](#genericResponse) -->
    
    <!--end l-->


+ code: `200`, type: [GetGroupMembersReply](#GetGroupMembersReply)
    <!--beg l desc_GetGroupMembers_response_200_[GetGroupMembersReply](#GetGroupMembersReply) -->
    
    <!--end l-->


+ code: `500`, type: No Response
    <!--beg l desc_GetGroupMembers_response_500_No Response -->
    
    <!--end l-->




### PostGroupMember

The uri/restful key of this method is `/group/{gid}/user/{id}@POST`

<!--beg l desc_PostGroupMember -->

<!--end l-->

parameters:

+ `gid`: [string](#string) (required): 
    <!--beg l desc_PostGroupMember_params_gid -->
    
    <!--end l-->


+ `id`: [string](#string) (required): 
    <!--beg l desc_PostGroupMember_params_id -->
    
    <!--end l-->


+ `PostGroupMemberRequest`: [PostGroupMemberRequest](#PostGroupMemberRequest): 
    <!--beg l desc_PostGroupMember_params_PostGroupMemberRequest -->
    
    <!--end l-->


responses:

+ code: `200`, type: [genericResponse](#genericResponse)
    <!--beg l desc_PostGroupMember_response_200_[genericResponse](#genericResponse) -->
    
    <!--end l-->


+ code: `200`, type: [PostGroupMemberReply](#PostGroupMemberReply)
    <!--beg l desc_PostGroupMember_response_200_[PostGroupMemberReply](#PostGroupMemberReply) -->
    
    <!--end l-->


+ code: `500`, type: No Response
    <!--beg l desc_PostGroupMember_response_500_No Response -->
    
    <!--end l-->




### DeleteGroup

The uri/restful key of this method is `/group/{gid}@DELETE`

<!--beg l desc_DeleteGroup -->

<!--end l-->

parameters:

+ `gid`: [string](#string) (required): 
    <!--beg l desc_DeleteGroup_params_gid -->
    
    <!--end l-->


+ `DeleteGroupRequest`: [DeleteGroupRequest](#DeleteGroupRequest): 
    <!--beg l desc_DeleteGroup_params_DeleteGroupRequest -->
    
    <!--end l-->


responses:

+ code: `200`, type: [genericResponse](#genericResponse)
    <!--beg l desc_DeleteGroup_response_200_[genericResponse](#genericResponse) -->
    
    <!--end l-->


+ code: `200`, type: [DeleteGroupReply](#DeleteGroupReply)
    <!--beg l desc_DeleteGroup_response_200_[DeleteGroupReply](#DeleteGroupReply) -->
    
    <!--end l-->


+ code: `500`, type: No Response
    <!--beg l desc_DeleteGroup_response_500_No Response -->
    
    <!--end l-->




### GetGroup

The uri/restful key of this method is `/group/{gid}@GET`

<!--beg l desc_GetGroup -->

<!--end l-->

parameters:

+ `gid`: [string](#string) (required): 
    <!--beg l desc_GetGroup_params_gid -->
    
    <!--end l-->


responses:

+ code: `200`, type: [genericResponse](#genericResponse)
    <!--beg l desc_GetGroup_response_200_[genericResponse](#genericResponse) -->
    
    <!--end l-->


+ code: `200`, type: [GetGroupReply](#GetGroupReply)
    <!--beg l desc_GetGroup_response_200_[GetGroupReply](#GetGroupReply) -->
    
    <!--end l-->


+ code: `500`, type: No Response
    <!--beg l desc_GetGroup_response_500_No Response -->
    
    <!--end l-->




### PutGroup

The uri/restful key of this method is `/group/{gid}@PUT`

<!--beg l desc_PutGroup -->

<!--end l-->

parameters:

+ `gid`: [string](#string) (required): 
    <!--beg l desc_PutGroup_params_gid -->
    
    <!--end l-->


+ `PutGroupRequest`: [PutGroupRequest](#PutGroupRequest): 
    <!--beg l desc_PutGroup_params_PutGroupRequest -->
    
    <!--end l-->


responses:

+ code: `200`, type: [genericResponse](#genericResponse)
    <!--beg l desc_PutGroup_response_200_[genericResponse](#genericResponse) -->
    
    <!--end l-->


+ code: `200`, type: [PutGroupReply](#PutGroupReply)
    <!--beg l desc_PutGroup_response_200_[PutGroupReply](#PutGroupReply) -->
    
    <!--end l-->


+ code: `500`, type: No Response
    <!--beg l desc_PutGroup_response_500_No Response -->
    
    <!--end l-->




### PostGroup

The uri/restful key of this method is `/group@POST`

<!--beg l desc_PostGroup -->

<!--end l-->

parameters:

+ `PostGroupRequest`: [PostGroupRequest](#PostGroupRequest): 
    <!--beg l desc_PostGroup_params_PostGroupRequest -->
    
    <!--end l-->


responses:

+ code: `200`, type: [genericResponse](#genericResponse)
    <!--beg l desc_PostGroup_response_200_[genericResponse](#genericResponse) -->
    
    <!--end l-->


+ code: `200`, type: [PostGroupReply](#PostGroupReply)
    <!--beg l desc_PostGroup_response_200_[PostGroupReply](#PostGroupReply) -->
    
    <!--end l-->


+ code: `500`, type: No Response
    <!--beg l desc_PostGroup_response_500_No Response -->
    
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

    
### [CountGroupReply](./ObjectModelSpec.md#CountGroupReply)

+ type: [object](#CountGroupReply)

+ fields:
    
    + `code`: integer: 
        <!--beg l desc_{{object_name}}_code -->
        
        <!--end l-->

    + `data`: integer: 
        <!--beg l desc_{{object_name}}_data -->
        
        <!--end l-->

    
### [ListGroupReply](./ObjectModelSpec.md#ListGroupReply)

+ type: [object](#ListGroupReply)

+ fields:
    
    + `code`: integer: 
        <!--beg l desc_{{object_name}}_code -->
        
        <!--end l-->

    + `data`: array: 
        <!--beg l desc_{{object_name}}_data -->
        
        <!--end l-->

    
### [PutGroupOwnerRequest](./ObjectModelSpec.md#PutGroupOwnerRequest)

+ type: [object](#PutGroupOwnerRequest)

+ fields:
    
    + `owner_id`: integer: 
        <!--beg l desc_{{object_name}}_owner_id -->
        
        <!--end l-->

    
### [PutGroupOwnerReply](./ObjectModelSpec.md#PutGroupOwnerReply)

+ type: [object](#PutGroupOwnerReply)

+ fields:
    
    + `code`: integer: 
        <!--beg l desc_{{object_name}}_code -->
        
        <!--end l-->

    
### [GetGroupMembersReply](./ObjectModelSpec.md#GetGroupMembersReply)

+ type: [object](#GetGroupMembersReply)

+ fields:
    
    + `code`: integer: 
        <!--beg l desc_{{object_name}}_code -->
        
        <!--end l-->

    + `data`: array: 
        <!--beg l desc_{{object_name}}_data -->
        
        <!--end l-->

    
### [PostGroupMemberRequest](./ObjectModelSpec.md#PostGroupMemberRequest)

+ type: [object](#PostGroupMemberRequest)

+ fields:
    
    
### [PostGroupMemberReply](./ObjectModelSpec.md#PostGroupMemberReply)

+ type: [object](#PostGroupMemberReply)

+ fields:
    
    + `code`: integer: 
        <!--beg l desc_{{object_name}}_code -->
        
        <!--end l-->

    
### [DeleteGroupRequest](./ObjectModelSpec.md#DeleteGroupRequest)

+ type: [object](#DeleteGroupRequest)

+ fields:
    
    
### [DeleteGroupReply](./ObjectModelSpec.md#DeleteGroupReply)

+ type: [object](#DeleteGroupReply)

+ fields:
    
    + `code`: integer: 
        <!--beg l desc_{{object_name}}_code -->
        
        <!--end l-->

    
### [GetGroupReply](./ObjectModelSpec.md#GetGroupReply)

+ type: [object](#GetGroupReply)

+ fields:
    
    + `code`: integer: 
        <!--beg l desc_{{object_name}}_code -->
        
        <!--end l-->

    + `data`: : 
        <!--beg l desc_{{object_name}}_data -->
        
        <!--end l-->

    
### [PutGroupRequest](./ObjectModelSpec.md#PutGroupRequest)

+ type: [object](#PutGroupRequest)

+ fields:
    
    + `description`: string: 
        <!--beg l desc_{{object_name}}_description -->
        
        <!--end l-->

    + `name`: string: 
        <!--beg l desc_{{object_name}}_name -->
        
        <!--end l-->

    
### [PutGroupReply](./ObjectModelSpec.md#PutGroupReply)

+ type: [object](#PutGroupReply)

+ fields:
    
    + `code`: integer: 
        <!--beg l desc_{{object_name}}_code -->
        
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

    
### [PostGroupReply](./ObjectModelSpec.md#PostGroupReply)

+ type: [object](#PostGroupReply)

+ fields:
    
    + `code`: integer: 
        <!--beg l desc_{{object_name}}_code -->
        
        <!--end l-->

    + `data`: integer: 
        <!--beg l desc_{{object_name}}_data -->
        
        <!--end l-->

    

