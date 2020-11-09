
# Api Group Documentation (AuthController)

<!--beg l desc_AuthController -->

<!--end l-->

## Apis


### RemoveGroupingPolicy

restful key: `/policy/group@DELETE`

<!--beg l desc_RemoveGroupingPolicy -->

<!--end l-->

parameters:

+ `RemoveGroupingPolicyRequest`: [RemoveGroupingPolicyRequest](#RemoveGroupingPolicyRequest): 
    <!--beg l desc_RemoveGroupingPolicy_params_RemoveGroupingPolicyRequest -->
    
    <!--end l-->


responses:

+ code: `200`, type: [genericResponse](#genericResponse)
    <!--beg l desc_RemoveGroupingPolicy_response_200_[genericResponse](#genericResponse) -->
    
    <!--end l-->


+ code: `200`, type: [RemoveGroupingPolicyReply](#RemoveGroupingPolicyReply)
    <!--beg l desc_RemoveGroupingPolicy_response_200_[RemoveGroupingPolicyReply](#RemoveGroupingPolicyReply) -->
    
    <!--end l-->


+ code: `500`, type: No Response
    <!--beg l desc_RemoveGroupingPolicy_response_500_No Response -->
    
    <!--end l-->




### HasGroupingPolicy

restful key: `/policy/group@GET`

<!--beg l desc_HasGroupingPolicy -->

<!--end l-->

parameters:

+ `subject`: [string](#string): 
    <!--beg l desc_HasGroupingPolicy_params_subject -->
    
    <!--end l-->


+ `group`: [string](#string): 
    <!--beg l desc_HasGroupingPolicy_params_group -->
    
    <!--end l-->


responses:

+ code: `200`, type: [genericResponse](#genericResponse)
    <!--beg l desc_HasGroupingPolicy_response_200_[genericResponse](#genericResponse) -->
    
    <!--end l-->


+ code: `200`, type: [HasGroupingPolicyReply](#HasGroupingPolicyReply)
    <!--beg l desc_HasGroupingPolicy_response_200_[HasGroupingPolicyReply](#HasGroupingPolicyReply) -->
    
    <!--end l-->


+ code: `500`, type: No Response
    <!--beg l desc_HasGroupingPolicy_response_500_No Response -->
    
    <!--end l-->




### AddGroupingPolicy

restful key: `/policy/group@POST`

<!--beg l desc_AddGroupingPolicy -->

<!--end l-->

parameters:

+ `AddGroupingPolicyRequest`: [AddGroupingPolicyRequest](#AddGroupingPolicyRequest): 
    <!--beg l desc_AddGroupingPolicy_params_AddGroupingPolicyRequest -->
    
    <!--end l-->


responses:

+ code: `200`, type: [genericResponse](#genericResponse)
    <!--beg l desc_AddGroupingPolicy_response_200_[genericResponse](#genericResponse) -->
    
    <!--end l-->


+ code: `200`, type: [AddGroupingPolicyReply](#AddGroupingPolicyReply)
    <!--beg l desc_AddGroupingPolicy_response_200_[AddGroupingPolicyReply](#AddGroupingPolicyReply) -->
    
    <!--end l-->


+ code: `500`, type: No Response
    <!--beg l desc_AddGroupingPolicy_response_500_No Response -->
    
    <!--end l-->




### RemovePolicy

restful key: `/policy@DELETE`

<!--beg l desc_RemovePolicy -->

<!--end l-->

parameters:

+ `RemovePolicyRequest`: [RemovePolicyRequest](#RemovePolicyRequest): 
    <!--beg l desc_RemovePolicy_params_RemovePolicyRequest -->
    
    <!--end l-->


responses:

+ code: `200`, type: [genericResponse](#genericResponse)
    <!--beg l desc_RemovePolicy_response_200_[genericResponse](#genericResponse) -->
    
    <!--end l-->


+ code: `200`, type: [RemovePolicyReply](#RemovePolicyReply)
    <!--beg l desc_RemovePolicy_response_200_[RemovePolicyReply](#RemovePolicyReply) -->
    
    <!--end l-->


+ code: `500`, type: No Response
    <!--beg l desc_RemovePolicy_response_500_No Response -->
    
    <!--end l-->




### HasPolicy

restful key: `/policy@GET`

<!--beg l desc_HasPolicy -->

<!--end l-->

parameters:

+ `subject`: [string](#string): 
    <!--beg l desc_HasPolicy_params_subject -->
    
    <!--end l-->


+ `object`: [string](#string): 
    <!--beg l desc_HasPolicy_params_object -->
    
    <!--end l-->


+ `action`: [string](#string): 
    <!--beg l desc_HasPolicy_params_action -->
    
    <!--end l-->


responses:

+ code: `200`, type: [genericResponse](#genericResponse)
    <!--beg l desc_HasPolicy_response_200_[genericResponse](#genericResponse) -->
    
    <!--end l-->


+ code: `200`, type: [HasPolicyReply](#HasPolicyReply)
    <!--beg l desc_HasPolicy_response_200_[HasPolicyReply](#HasPolicyReply) -->
    
    <!--end l-->


+ code: `500`, type: No Response
    <!--beg l desc_HasPolicy_response_500_No Response -->
    
    <!--end l-->




### AddPolicy

restful key: `/policy@POST`

<!--beg l desc_AddPolicy -->

<!--end l-->

parameters:

+ `AddPolicyRequest`: [AddPolicyRequest](#AddPolicyRequest): 
    <!--beg l desc_AddPolicy_params_AddPolicyRequest -->
    
    <!--end l-->


responses:

+ code: `200`, type: [genericResponse](#genericResponse)
    <!--beg l desc_AddPolicy_response_200_[genericResponse](#genericResponse) -->
    
    <!--end l-->


+ code: `200`, type: [AddPolicyReply](#AddPolicyReply)
    <!--beg l desc_AddPolicy_response_200_[AddPolicyReply](#AddPolicyReply) -->
    
    <!--end l-->


+ code: `500`, type: No Response
    <!--beg l desc_AddPolicy_response_500_No Response -->
    
    <!--end l-->




## Local Object Reference




### [RemoveGroupingPolicyRequest](./ObjectModelSpec.md#RemoveGroupingPolicyRequest)

+ type: [object](#RemoveGroupingPolicyRequest)

+ fields:
    
    + `group`: string: 
        <!--beg l desc_{{object_name}}_group -->
        
        <!--end l-->

    + `subject`: string: 
        <!--beg l desc_{{object_name}}_subject -->
        
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

    
### [RemoveGroupingPolicyReply](./ObjectModelSpec.md#RemoveGroupingPolicyReply)

+ type: [object](#RemoveGroupingPolicyReply)

+ fields:
    
    + `code`: integer: 
        <!--beg l desc_{{object_name}}_code -->
        
        <!--end l-->

    + `data`: boolean: 
        <!--beg l desc_{{object_name}}_data -->
        
        <!--end l-->

    
### [HasGroupingPolicyReply](./ObjectModelSpec.md#HasGroupingPolicyReply)

+ type: [object](#HasGroupingPolicyReply)

+ fields:
    
    + `code`: integer: 
        <!--beg l desc_{{object_name}}_code -->
        
        <!--end l-->

    + `data`: boolean: 
        <!--beg l desc_{{object_name}}_data -->
        
        <!--end l-->

    
### [AddGroupingPolicyRequest](./ObjectModelSpec.md#AddGroupingPolicyRequest)

+ type: [object](#AddGroupingPolicyRequest)

+ fields:
    
    + `group`: string: 
        <!--beg l desc_{{object_name}}_group -->
        
        <!--end l-->

    + `subject`: string: 
        <!--beg l desc_{{object_name}}_subject -->
        
        <!--end l-->

    
### [AddGroupingPolicyReply](./ObjectModelSpec.md#AddGroupingPolicyReply)

+ type: [object](#AddGroupingPolicyReply)

+ fields:
    
    + `code`: integer: 
        <!--beg l desc_{{object_name}}_code -->
        
        <!--end l-->

    + `data`: boolean: 
        <!--beg l desc_{{object_name}}_data -->
        
        <!--end l-->

    
### [RemovePolicyRequest](./ObjectModelSpec.md#RemovePolicyRequest)

+ type: [object](#RemovePolicyRequest)

+ fields:
    
    + `action`: string: 
        <!--beg l desc_{{object_name}}_action -->
        
        <!--end l-->

    + `object`: string: 
        <!--beg l desc_{{object_name}}_object -->
        
        <!--end l-->

    + `subject`: string: 
        <!--beg l desc_{{object_name}}_subject -->
        
        <!--end l-->

    
### [RemovePolicyReply](./ObjectModelSpec.md#RemovePolicyReply)

+ type: [object](#RemovePolicyReply)

+ fields:
    
    + `code`: integer: 
        <!--beg l desc_{{object_name}}_code -->
        
        <!--end l-->

    + `data`: boolean: 
        <!--beg l desc_{{object_name}}_data -->
        
        <!--end l-->

    
### [HasPolicyReply](./ObjectModelSpec.md#HasPolicyReply)

+ type: [object](#HasPolicyReply)

+ fields:
    
    + `code`: integer: 
        <!--beg l desc_{{object_name}}_code -->
        
        <!--end l-->

    + `data`: boolean: 
        <!--beg l desc_{{object_name}}_data -->
        
        <!--end l-->

    
### [AddPolicyRequest](./ObjectModelSpec.md#AddPolicyRequest)

+ type: [object](#AddPolicyRequest)

+ fields:
    
    + `action`: string: 
        <!--beg l desc_{{object_name}}_action -->
        
        <!--end l-->

    + `object`: string: 
        <!--beg l desc_{{object_name}}_object -->
        
        <!--end l-->

    + `subject`: string: 
        <!--beg l desc_{{object_name}}_subject -->
        
        <!--end l-->

    
### [AddPolicyReply](./ObjectModelSpec.md#AddPolicyReply)

+ type: [object](#AddPolicyReply)

+ fields:
    
    + `code`: integer: 
        <!--beg l desc_{{object_name}}_code -->
        
        <!--end l-->

    + `data`: boolean: 
        <!--beg l desc_{{object_name}}_data -->
        
        <!--end l-->

    

