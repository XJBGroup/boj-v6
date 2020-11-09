
# Api Group Documentation (UserController)

<!--beg l desc_UserController -->

<!--end l-->

## Apis


### CountUser

The uri/restful key of this method is `/user-count@GET`

<!--beg l desc_CountUser -->

<!--end l-->


+ `page`: [integer](#integer): 
    <!--beg l desc_CountUser_page -->
    
    <!--end l-->


+ `page_size`: [integer](#integer): 
    <!--beg l desc_CountUser_page_size -->
    
    <!--end l-->



### ListUser

The uri/restful key of this method is `/user-list@GET`

<!--beg l desc_ListUser -->

<!--end l-->


+ `page`: [integer](#integer): 
    <!--beg l desc_ListUser_page -->
    
    <!--end l-->


+ `page_size`: [integer](#integer): 
    <!--beg l desc_ListUser_page_size -->
    
    <!--end l-->



### RefreshToken

The uri/restful key of this method is `/user-token@GET`

<!--beg l desc_RefreshToken -->

<!--end l-->



### LoginUser

The uri/restful key of this method is `/user/login@POST`

<!--beg l desc_LoginUser -->

<!--end l-->


+ `LoginUserRequest`: [LoginUserRequest](#LoginUserRequest): 
    <!--beg l desc_LoginUser_LoginUserRequest -->
    
    <!--end l-->



### Register

The uri/restful key of this method is `/user/register@POST`

<!--beg l desc_Register -->

<!--end l-->


+ `RegisterRequest`: [RegisterRequest](#RegisterRequest): 
    <!--beg l desc_Register_RegisterRequest -->
    
    <!--end l-->



### BindEmail

The uri/restful key of this method is `/user/{id}/email@PUT`

<!--beg l desc_BindEmail -->

<!--end l-->


+ `id`: [string](#string) (required): 
    <!--beg l desc_BindEmail_id -->
    
    <!--end l-->


+ `BindEmailRequest`: [BindEmailRequest](#BindEmailRequest): 
    <!--beg l desc_BindEmail_BindEmailRequest -->
    
    <!--end l-->



### InspectUser

The uri/restful key of this method is `/user/{id}/inspect@GET`

<!--beg l desc_InspectUser -->

<!--end l-->


+ `id`: [string](#string) (required): 
    <!--beg l desc_InspectUser_id -->
    
    <!--end l-->



### ChangePassword

The uri/restful key of this method is `/user/{id}/password@PUT`

<!--beg l desc_ChangePassword -->

<!--end l-->


+ `id`: [string](#string) (required): 
    <!--beg l desc_ChangePassword_id -->
    
    <!--end l-->


+ `ChangePasswordRequest`: [ChangePasswordRequest](#ChangePasswordRequest): 
    <!--beg l desc_ChangePassword_ChangePasswordRequest -->
    
    <!--end l-->



### DeleteUser

The uri/restful key of this method is `/user/{id}@DELETE`

<!--beg l desc_DeleteUser -->

<!--end l-->


+ `id`: [string](#string) (required): 
    <!--beg l desc_DeleteUser_id -->
    
    <!--end l-->


+ `DeleteUserRequest`: [DeleteUserRequest](#DeleteUserRequest): 
    <!--beg l desc_DeleteUser_DeleteUserRequest -->
    
    <!--end l-->



### GetUser

The uri/restful key of this method is `/user/{id}@GET`

<!--beg l desc_GetUser -->

<!--end l-->


+ `id`: [string](#string) (required): 
    <!--beg l desc_GetUser_id -->
    
    <!--end l-->



### PutUser

The uri/restful key of this method is `/user/{id}@PUT`

<!--beg l desc_PutUser -->

<!--end l-->


+ `id`: [string](#string) (required): 
    <!--beg l desc_PutUser_id -->
    
    <!--end l-->


+ `PutUserRequest`: [PutUserRequest](#PutUserRequest): 
    <!--beg l desc_PutUser_PutUserRequest -->
    
    <!--end l-->



## Local Object Reference




### [LoginUserRequest](./ObjectModelSpec.md#LoginUserRequest)

+ type: [object](#LoginUserRequest)

+ fields:
    
    + `email`: string: 
        <!--beg l desc_{{object_name}}_email -->
        
        <!--end l-->

    + `id`: integer: 
        <!--beg l desc_{{object_name}}_id -->
        
        <!--end l-->

    + `password`: string: 
        <!--beg l desc_{{object_name}}_password -->
        
        <!--end l-->

    + `user_name`: string: 
        <!--beg l desc_{{object_name}}_user_name -->
        
        <!--end l-->

    
### [RegisterRequest](./ObjectModelSpec.md#RegisterRequest)

+ type: [object](#RegisterRequest)

+ fields:
    
    + `gender`: integer: 
        <!--beg l desc_{{object_name}}_gender -->
        
        <!--end l-->

    + `nick_name`: string: 
        <!--beg l desc_{{object_name}}_nick_name -->
        
        <!--end l-->

    + `password`: string: 
        <!--beg l desc_{{object_name}}_password -->
        
        <!--end l-->

    + `user_name`: string: 
        <!--beg l desc_{{object_name}}_user_name -->
        
        <!--end l-->

    
### [BindEmailRequest](./ObjectModelSpec.md#BindEmailRequest)

+ type: [object](#BindEmailRequest)

+ fields:
    
    + `email`: string: 
        <!--beg l desc_{{object_name}}_email -->
        
        <!--end l-->

    
### [ChangePasswordRequest](./ObjectModelSpec.md#ChangePasswordRequest)

+ type: [object](#ChangePasswordRequest)

+ fields:
    
    + `new_password`: string: 
        <!--beg l desc_{{object_name}}_new_password -->
        
        <!--end l-->

    + `old_password`: string: 
        <!--beg l desc_{{object_name}}_old_password -->
        
        <!--end l-->

    
### [DeleteUserRequest](./ObjectModelSpec.md#DeleteUserRequest)

+ type: [object](#DeleteUserRequest)

+ fields:
    
    
### [PutUserRequest](./ObjectModelSpec.md#PutUserRequest)

+ type: [object](#PutUserRequest)

+ fields:
    
    + `gender`: integer: 
        <!--beg l desc_{{object_name}}_gender -->
        
        <!--end l-->

    + `motto`: string: 
        <!--beg l desc_{{object_name}}_motto -->
        
        <!--end l-->

    + `nick_name`: string: 
        <!--beg l desc_{{object_name}}_nick_name -->
        
        <!--end l-->

    

