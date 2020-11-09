
# Api Group Documentation (AuthController)

<!--beg l desc_AuthController -->

<!--end l-->

## Apis


### RemoveGroupingPolicy

The uri/restful key of this method is `/policy/group@DELETE`

<!--beg l desc_RemoveGroupingPolicy -->

<!--end l-->


+ `RemoveGroupingPolicyRequest`: [RemoveGroupingPolicyRequest](#RemoveGroupingPolicyRequest): 
    <!--beg l desc_RemoveGroupingPolicy_RemoveGroupingPolicyRequest -->
    
    <!--end l-->



### HasGroupingPolicy

The uri/restful key of this method is `/policy/group@GET`

<!--beg l desc_HasGroupingPolicy -->

<!--end l-->


+ `subject`: [string](#string): 
    <!--beg l desc_HasGroupingPolicy_subject -->
    
    <!--end l-->


+ `group`: [string](#string): 
    <!--beg l desc_HasGroupingPolicy_group -->
    
    <!--end l-->



### AddGroupingPolicy

The uri/restful key of this method is `/policy/group@POST`

<!--beg l desc_AddGroupingPolicy -->

<!--end l-->


+ `AddGroupingPolicyRequest`: [AddGroupingPolicyRequest](#AddGroupingPolicyRequest): 
    <!--beg l desc_AddGroupingPolicy_AddGroupingPolicyRequest -->
    
    <!--end l-->



### RemovePolicy

The uri/restful key of this method is `/policy@DELETE`

<!--beg l desc_RemovePolicy -->

<!--end l-->


+ `RemovePolicyRequest`: [RemovePolicyRequest](#RemovePolicyRequest): 
    <!--beg l desc_RemovePolicy_RemovePolicyRequest -->
    
    <!--end l-->



### HasPolicy

The uri/restful key of this method is `/policy@GET`

<!--beg l desc_HasPolicy -->

<!--end l-->


+ `subject`: [string](#string): 
    <!--beg l desc_HasPolicy_subject -->
    
    <!--end l-->


+ `object`: [string](#string): 
    <!--beg l desc_HasPolicy_object -->
    
    <!--end l-->


+ `action`: [string](#string): 
    <!--beg l desc_HasPolicy_action -->
    
    <!--end l-->



### AddPolicy

The uri/restful key of this method is `/policy@POST`

<!--beg l desc_AddPolicy -->

<!--end l-->


+ `AddPolicyRequest`: [AddPolicyRequest](#AddPolicyRequest): 
    <!--beg l desc_AddPolicy_AddPolicyRequest -->
    
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

    
### [AddGroupingPolicyRequest](./ObjectModelSpec.md#AddGroupingPolicyRequest)

+ type: [object](#AddGroupingPolicyRequest)

+ fields:
    
    + `group`: string: 
        <!--beg l desc_{{object_name}}_group -->
        
        <!--end l-->

    + `subject`: string: 
        <!--beg l desc_{{object_name}}_subject -->
        
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

    

