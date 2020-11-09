
# Api Group Documentation (UserController)

<!--beg l desc_UserController -->

<!--end l-->

## Apis


### CountUser

restful key: `/user-count@GET`

<!--beg l desc_CountUser -->

<!--end l-->

parameters:

+ `page`: [integer](#integer): 
    <!--beg l desc_CountUser_params_page -->
    
    <!--end l-->


+ `page_size`: [integer](#integer): 
    <!--beg l desc_CountUser_params_page_size -->
    
    <!--end l-->


responses:

+ code: `200`, type: [genericResponse](#genericResponse)
    <!--beg l desc_CountUser_response_200_[genericResponse](#genericResponse) -->
    
    <!--end l-->


+ code: `200`, type: [CountUserReply](#CountUserReply)
    <!--beg l desc_CountUser_response_200_[CountUserReply](#CountUserReply) -->
    
    <!--end l-->


+ code: `500`, type: No Response
    <!--beg l desc_CountUser_response_500_No Response -->
    
    <!--end l-->




### ListUser

restful key: `/user-list@GET`

<!--beg l desc_ListUser -->

<!--end l-->

parameters:

+ `page`: [integer](#integer): 
    <!--beg l desc_ListUser_params_page -->
    
    <!--end l-->


+ `page_size`: [integer](#integer): 
    <!--beg l desc_ListUser_params_page_size -->
    
    <!--end l-->


responses:

+ code: `200`, type: [genericResponse](#genericResponse)
    <!--beg l desc_ListUser_response_200_[genericResponse](#genericResponse) -->
    
    <!--end l-->


+ code: `200`, type: [ListUserReply](#ListUserReply)
    <!--beg l desc_ListUser_response_200_[ListUserReply](#ListUserReply) -->
    
    <!--end l-->


+ code: `500`, type: No Response
    <!--beg l desc_ListUser_response_500_No Response -->
    
    <!--end l-->




### RefreshToken

restful key: `/user-token@GET`

<!--beg l desc_RefreshToken -->

<!--end l-->

parameters:
+ the api does not have any input.

responses:

+ code: `200`, type: [genericResponse](#genericResponse)
    <!--beg l desc_RefreshToken_response_200_[genericResponse](#genericResponse) -->
    
    <!--end l-->


+ code: `200`, type: [RefreshTokenReply](#RefreshTokenReply)
    <!--beg l desc_RefreshToken_response_200_[RefreshTokenReply](#RefreshTokenReply) -->
    
    <!--end l-->


+ code: `500`, type: No Response
    <!--beg l desc_RefreshToken_response_500_No Response -->
    
    <!--end l-->




### LoginUser

restful key: `/user/login@POST`

<!--beg l desc_LoginUser -->

<!--end l-->

parameters:

+ `LoginUserRequest`: [LoginUserRequest](#LoginUserRequest): 
    <!--beg l desc_LoginUser_params_LoginUserRequest -->
    
    <!--end l-->


responses:

+ code: `200`, type: [genericResponse](#genericResponse)
    <!--beg l desc_LoginUser_response_200_[genericResponse](#genericResponse) -->
    
    <!--end l-->


+ code: `200`, type: [LoginUserReply](#LoginUserReply)
    <!--beg l desc_LoginUser_response_200_[LoginUserReply](#LoginUserReply) -->
    
    <!--end l-->


+ code: `500`, type: No Response
    <!--beg l desc_LoginUser_response_500_No Response -->
    
    <!--end l-->




### Register

restful key: `/user/register@POST`

<!--beg l desc_Register -->

<!--end l-->

parameters:

+ `RegisterRequest`: [RegisterRequest](#RegisterRequest): 
    <!--beg l desc_Register_params_RegisterRequest -->
    
    <!--end l-->


responses:

+ code: `200`, type: [genericResponse](#genericResponse)
    <!--beg l desc_Register_response_200_[genericResponse](#genericResponse) -->
    
    <!--end l-->


+ code: `200`, type: [RegisterReply](#RegisterReply)
    <!--beg l desc_Register_response_200_[RegisterReply](#RegisterReply) -->
    
    <!--end l-->


+ code: `500`, type: No Response
    <!--beg l desc_Register_response_500_No Response -->
    
    <!--end l-->




### BindEmail

restful key: `/user/{id}/email@PUT`

<!--beg l desc_BindEmail -->

<!--end l-->

parameters:

+ `id`: [string](#string) (required): 
    <!--beg l desc_BindEmail_params_id -->
    
    <!--end l-->


+ `BindEmailRequest`: [BindEmailRequest](#BindEmailRequest): 
    <!--beg l desc_BindEmail_params_BindEmailRequest -->
    
    <!--end l-->


responses:

+ code: `200`, type: [genericResponse](#genericResponse)
    <!--beg l desc_BindEmail_response_200_[genericResponse](#genericResponse) -->
    
    <!--end l-->


+ code: `200`, type: [BindEmailReply](#BindEmailReply)
    <!--beg l desc_BindEmail_response_200_[BindEmailReply](#BindEmailReply) -->
    
    <!--end l-->


+ code: `500`, type: No Response
    <!--beg l desc_BindEmail_response_500_No Response -->
    
    <!--end l-->




### InspectUser

restful key: `/user/{id}/inspect@GET`

<!--beg l desc_InspectUser -->

<!--end l-->

parameters:

+ `id`: [string](#string) (required): 
    <!--beg l desc_InspectUser_params_id -->
    
    <!--end l-->


responses:

+ code: `200`, type: [genericResponse](#genericResponse)
    <!--beg l desc_InspectUser_response_200_[genericResponse](#genericResponse) -->
    
    <!--end l-->


+ code: `200`, type: [InspectUserReply](#InspectUserReply)
    <!--beg l desc_InspectUser_response_200_[InspectUserReply](#InspectUserReply) -->
    
    <!--end l-->


+ code: `500`, type: No Response
    <!--beg l desc_InspectUser_response_500_No Response -->
    
    <!--end l-->




### ChangePassword

restful key: `/user/{id}/password@PUT`

<!--beg l desc_ChangePassword -->

<!--end l-->

parameters:

+ `id`: [string](#string) (required): 
    <!--beg l desc_ChangePassword_params_id -->
    
    <!--end l-->


+ `ChangePasswordRequest`: [ChangePasswordRequest](#ChangePasswordRequest): 
    <!--beg l desc_ChangePassword_params_ChangePasswordRequest -->
    
    <!--end l-->


responses:

+ code: `200`, type: [genericResponse](#genericResponse)
    <!--beg l desc_ChangePassword_response_200_[genericResponse](#genericResponse) -->
    
    <!--end l-->


+ code: `200`, type: [ChangePasswordReply](#ChangePasswordReply)
    <!--beg l desc_ChangePassword_response_200_[ChangePasswordReply](#ChangePasswordReply) -->
    
    <!--end l-->


+ code: `500`, type: No Response
    <!--beg l desc_ChangePassword_response_500_No Response -->
    
    <!--end l-->




### DeleteUser

restful key: `/user/{id}@DELETE`

<!--beg l desc_DeleteUser -->

<!--end l-->

parameters:

+ `id`: [string](#string) (required): 
    <!--beg l desc_DeleteUser_params_id -->
    
    <!--end l-->


+ `DeleteUserRequest`: [DeleteUserRequest](#DeleteUserRequest): 
    <!--beg l desc_DeleteUser_params_DeleteUserRequest -->
    
    <!--end l-->


responses:

+ code: `200`, type: [genericResponse](#genericResponse)
    <!--beg l desc_DeleteUser_response_200_[genericResponse](#genericResponse) -->
    
    <!--end l-->


+ code: `200`, type: [DeleteUserReply](#DeleteUserReply)
    <!--beg l desc_DeleteUser_response_200_[DeleteUserReply](#DeleteUserReply) -->
    
    <!--end l-->


+ code: `500`, type: No Response
    <!--beg l desc_DeleteUser_response_500_No Response -->
    
    <!--end l-->




### GetUser

restful key: `/user/{id}@GET`

<!--beg l desc_GetUser -->

<!--end l-->

parameters:

+ `id`: [string](#string) (required): 
    <!--beg l desc_GetUser_params_id -->
    
    <!--end l-->


responses:

+ code: `200`, type: [genericResponse](#genericResponse)
    <!--beg l desc_GetUser_response_200_[genericResponse](#genericResponse) -->
    
    <!--end l-->


+ code: `200`, type: [GetUserReply](#GetUserReply)
    <!--beg l desc_GetUser_response_200_[GetUserReply](#GetUserReply) -->
    
    <!--end l-->


+ code: `500`, type: No Response
    <!--beg l desc_GetUser_response_500_No Response -->
    
    <!--end l-->




### PutUser

restful key: `/user/{id}@PUT`

<!--beg l desc_PutUser -->

<!--end l-->

parameters:

+ `id`: [string](#string) (required): 
    <!--beg l desc_PutUser_params_id -->
    
    <!--end l-->


+ `PutUserRequest`: [PutUserRequest](#PutUserRequest): 
    <!--beg l desc_PutUser_params_PutUserRequest -->
    
    <!--end l-->


responses:

+ code: `200`, type: [genericResponse](#genericResponse)
    <!--beg l desc_PutUser_response_200_[genericResponse](#genericResponse) -->
    
    <!--end l-->


+ code: `200`, type: [PutUserReply](#PutUserReply)
    <!--beg l desc_PutUser_response_200_[PutUserReply](#PutUserReply) -->
    
    <!--end l-->


+ code: `500`, type: No Response
    <!--beg l desc_PutUser_response_500_No Response -->
    
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

    
### [CountUserReply](./ObjectModelSpec.md#CountUserReply)

+ type: [object](#CountUserReply)

+ fields:
    
    + `code`: integer: 
        <!--beg l desc_{{object_name}}_code -->
        
        <!--end l-->

    + `data`: integer: 
        <!--beg l desc_{{object_name}}_data -->
        
        <!--end l-->

    
### [ListUserReply](./ObjectModelSpec.md#ListUserReply)

+ type: [object](#ListUserReply)

+ fields:
    
    + `code`: integer: 
        <!--beg l desc_{{object_name}}_code -->
        
        <!--end l-->

    + `data`: array: 
        <!--beg l desc_{{object_name}}_data -->
        
        <!--end l-->

    
### [RefreshTokenReply](./ObjectModelSpec.md#RefreshTokenReply)

+ type: [object](#RefreshTokenReply)

+ fields:
    
    + `code`: integer: 
        <!--beg l desc_{{object_name}}_code -->
        
        <!--end l-->

    + `data`: : 
        <!--beg l desc_{{object_name}}_data -->
        
        <!--end l-->

    
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

    
### [LoginUserReply](./ObjectModelSpec.md#LoginUserReply)

+ type: [object](#LoginUserReply)

+ fields:
    
    + `code`: integer: 
        <!--beg l desc_{{object_name}}_code -->
        
        <!--end l-->

    + `data`: : 
        <!--beg l desc_{{object_name}}_data -->
        
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

    
### [RegisterReply](./ObjectModelSpec.md#RegisterReply)

+ type: [object](#RegisterReply)

+ fields:
    
    + `code`: integer: 
        <!--beg l desc_{{object_name}}_code -->
        
        <!--end l-->

    + `data`: : 
        <!--beg l desc_{{object_name}}_data -->
        
        <!--end l-->

    
### [BindEmailRequest](./ObjectModelSpec.md#BindEmailRequest)

+ type: [object](#BindEmailRequest)

+ fields:
    
    + `email`: string: 
        <!--beg l desc_{{object_name}}_email -->
        
        <!--end l-->

    
### [BindEmailReply](./ObjectModelSpec.md#BindEmailReply)

+ type: [object](#BindEmailReply)

+ fields:
    
    + `code`: integer: 
        <!--beg l desc_{{object_name}}_code -->
        
        <!--end l-->

    
### [InspectUserReply](./ObjectModelSpec.md#InspectUserReply)

+ type: [object](#InspectUserReply)

+ fields:
    
    + `code`: integer: 
        <!--beg l desc_{{object_name}}_code -->
        
        <!--end l-->

    + `data`: : 
        <!--beg l desc_{{object_name}}_data -->
        
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

    
### [ChangePasswordReply](./ObjectModelSpec.md#ChangePasswordReply)

+ type: [object](#ChangePasswordReply)

+ fields:
    
    + `code`: integer: 
        <!--beg l desc_{{object_name}}_code -->
        
        <!--end l-->

    
### [DeleteUserRequest](./ObjectModelSpec.md#DeleteUserRequest)

+ type: [object](#DeleteUserRequest)

+ fields:
    
    
### [DeleteUserReply](./ObjectModelSpec.md#DeleteUserReply)

+ type: [object](#DeleteUserReply)

+ fields:
    
    + `code`: integer: 
        <!--beg l desc_{{object_name}}_code -->
        
        <!--end l-->

    
### [GetUserReply](./ObjectModelSpec.md#GetUserReply)

+ type: [object](#GetUserReply)

+ fields:
    
    + `code`: integer: 
        <!--beg l desc_{{object_name}}_code -->
        
        <!--end l-->

    + `data`: : 
        <!--beg l desc_{{object_name}}_data -->
        
        <!--end l-->

    
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

    
### [PutUserReply](./ObjectModelSpec.md#PutUserReply)

+ type: [object](#PutUserReply)

+ fields:
    
    + `code`: integer: 
        <!--beg l desc_{{object_name}}_code -->
        
        <!--end l-->

    

