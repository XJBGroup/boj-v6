
# Api Group Documentation (AnnouncementController)

<!--beg l desc_AnnouncementController -->

<!--end l-->

## Apis


### CountAnnouncement

The uri/restful key of this method is `/announcement-count@GET`

<!--beg l desc_CountAnnouncement -->

<!--end l-->

parameters:
+ the api does not have any input.

responses:

+ code: `200`, type: [genericResponse](#genericResponse)
    <!--beg l desc_CountAnnouncement_response_200_[genericResponse](#genericResponse) -->
    
    <!--end l-->


+ code: `200`, type: [CountAnnouncementReply](#CountAnnouncementReply)
    <!--beg l desc_CountAnnouncement_response_200_[CountAnnouncementReply](#CountAnnouncementReply) -->
    
    <!--end l-->


+ code: `500`, type: No Response
    <!--beg l desc_CountAnnouncement_response_500_No Response -->
    
    <!--end l-->




### ListAnnouncement

The uri/restful key of this method is `/announcement-list@GET`

<!--beg l desc_ListAnnouncement -->

<!--end l-->

parameters:

+ `page`: [integer](#integer): 
    <!--beg l desc_ListAnnouncement_params_page -->
    
    <!--end l-->


+ `page_size`: [integer](#integer): 
    <!--beg l desc_ListAnnouncement_params_page_size -->
    
    <!--end l-->


responses:

+ code: `200`, type: [genericResponse](#genericResponse)
    <!--beg l desc_ListAnnouncement_response_200_[genericResponse](#genericResponse) -->
    
    <!--end l-->


+ code: `200`, type: [ListAnnouncementReply](#ListAnnouncementReply)
    <!--beg l desc_ListAnnouncement_response_200_[ListAnnouncementReply](#ListAnnouncementReply) -->
    
    <!--end l-->


+ code: `500`, type: No Response
    <!--beg l desc_ListAnnouncement_response_500_No Response -->
    
    <!--end l-->




### DeleteAnnouncement

The uri/restful key of this method is `/announcement/{aid}@DELETE`

<!--beg l desc_DeleteAnnouncement -->

<!--end l-->

parameters:

+ `aid`: [string](#string) (required): 
    <!--beg l desc_DeleteAnnouncement_params_aid -->
    
    <!--end l-->


+ `DeleteAnnouncementRequest`: [DeleteAnnouncementRequest](#DeleteAnnouncementRequest): 
    <!--beg l desc_DeleteAnnouncement_params_DeleteAnnouncementRequest -->
    
    <!--end l-->


responses:

+ code: `200`, type: [genericResponse](#genericResponse)
    <!--beg l desc_DeleteAnnouncement_response_200_[genericResponse](#genericResponse) -->
    
    <!--end l-->


+ code: `200`, type: [DeleteAnnouncementReply](#DeleteAnnouncementReply)
    <!--beg l desc_DeleteAnnouncement_response_200_[DeleteAnnouncementReply](#DeleteAnnouncementReply) -->
    
    <!--end l-->


+ code: `500`, type: No Response
    <!--beg l desc_DeleteAnnouncement_response_500_No Response -->
    
    <!--end l-->




### GetAnnouncement

The uri/restful key of this method is `/announcement/{aid}@GET`

<!--beg l desc_GetAnnouncement -->

<!--end l-->

parameters:

+ `aid`: [string](#string) (required): 
    <!--beg l desc_GetAnnouncement_params_aid -->
    
    <!--end l-->


responses:

+ code: `200`, type: [genericResponse](#genericResponse)
    <!--beg l desc_GetAnnouncement_response_200_[genericResponse](#genericResponse) -->
    
    <!--end l-->


+ code: `200`, type: [GetAnnouncementReply](#GetAnnouncementReply)
    <!--beg l desc_GetAnnouncement_response_200_[GetAnnouncementReply](#GetAnnouncementReply) -->
    
    <!--end l-->


+ code: `500`, type: No Response
    <!--beg l desc_GetAnnouncement_response_500_No Response -->
    
    <!--end l-->




### PutAnnouncement

The uri/restful key of this method is `/announcement/{aid}@PUT`

<!--beg l desc_PutAnnouncement -->

<!--end l-->

parameters:

+ `aid`: [string](#string) (required): 
    <!--beg l desc_PutAnnouncement_params_aid -->
    
    <!--end l-->


+ `PutAnnouncementRequest`: [PutAnnouncementRequest](#PutAnnouncementRequest): 
    <!--beg l desc_PutAnnouncement_params_PutAnnouncementRequest -->
    
    <!--end l-->


responses:

+ code: `200`, type: [genericResponse](#genericResponse)
    <!--beg l desc_PutAnnouncement_response_200_[genericResponse](#genericResponse) -->
    
    <!--end l-->


+ code: `200`, type: [PutAnnouncementReply](#PutAnnouncementReply)
    <!--beg l desc_PutAnnouncement_response_200_[PutAnnouncementReply](#PutAnnouncementReply) -->
    
    <!--end l-->


+ code: `500`, type: No Response
    <!--beg l desc_PutAnnouncement_response_500_No Response -->
    
    <!--end l-->




### PostAnnouncement

The uri/restful key of this method is `/announcement@POST`

<!--beg l desc_PostAnnouncement -->

<!--end l-->

parameters:

+ `PostAnnouncementRequest`: [PostAnnouncementRequest](#PostAnnouncementRequest): 
    <!--beg l desc_PostAnnouncement_params_PostAnnouncementRequest -->
    
    <!--end l-->


responses:

+ code: `200`, type: [genericResponse](#genericResponse)
    <!--beg l desc_PostAnnouncement_response_200_[genericResponse](#genericResponse) -->
    
    <!--end l-->


+ code: `200`, type: [PostAnnouncementReply](#PostAnnouncementReply)
    <!--beg l desc_PostAnnouncement_response_200_[PostAnnouncementReply](#PostAnnouncementReply) -->
    
    <!--end l-->


+ code: `500`, type: No Response
    <!--beg l desc_PostAnnouncement_response_500_No Response -->
    
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

    
### [CountAnnouncementReply](./ObjectModelSpec.md#CountAnnouncementReply)

+ type: [object](#CountAnnouncementReply)

+ fields:
    
    + `code`: integer: 
        <!--beg l desc_{{object_name}}_code -->
        
        <!--end l-->

    + `data`: integer: 
        <!--beg l desc_{{object_name}}_data -->
        
        <!--end l-->

    
### [ListAnnouncementReply](./ObjectModelSpec.md#ListAnnouncementReply)

+ type: [object](#ListAnnouncementReply)

+ fields:
    
    + `code`: integer: 
        <!--beg l desc_{{object_name}}_code -->
        
        <!--end l-->

    + `data`: array: 
        <!--beg l desc_{{object_name}}_data -->
        
        <!--end l-->

    
### [DeleteAnnouncementRequest](./ObjectModelSpec.md#DeleteAnnouncementRequest)

+ type: [object](#DeleteAnnouncementRequest)

+ fields:
    
    
### [DeleteAnnouncementReply](./ObjectModelSpec.md#DeleteAnnouncementReply)

+ type: [object](#DeleteAnnouncementReply)

+ fields:
    
    + `code`: integer: 
        <!--beg l desc_{{object_name}}_code -->
        
        <!--end l-->

    
### [GetAnnouncementReply](./ObjectModelSpec.md#GetAnnouncementReply)

+ type: [object](#GetAnnouncementReply)

+ fields:
    
    + `code`: integer: 
        <!--beg l desc_{{object_name}}_code -->
        
        <!--end l-->

    + `data`: : 
        <!--beg l desc_{{object_name}}_data -->
        
        <!--end l-->

    
### [PutAnnouncementRequest](./ObjectModelSpec.md#PutAnnouncementRequest)

+ type: [object](#PutAnnouncementRequest)

+ fields:
    
    + `content`: string: 
        <!--beg l desc_{{object_name}}_content -->
        
        <!--end l-->

    + `title`: string: 
        <!--beg l desc_{{object_name}}_title -->
        
        <!--end l-->

    
### [PutAnnouncementReply](./ObjectModelSpec.md#PutAnnouncementReply)

+ type: [object](#PutAnnouncementReply)

+ fields:
    
    + `code`: integer: 
        <!--beg l desc_{{object_name}}_code -->
        
        <!--end l-->

    
### [PostAnnouncementRequest](./ObjectModelSpec.md#PostAnnouncementRequest)

+ type: [object](#PostAnnouncementRequest)

+ fields:
    
    + `content`: string: 
        <!--beg l desc_{{object_name}}_content -->
        
        <!--end l-->

    + `title`: string: 
        <!--beg l desc_{{object_name}}_title -->
        
        <!--end l-->

    
### [PostAnnouncementReply](./ObjectModelSpec.md#PostAnnouncementReply)

+ type: [object](#PostAnnouncementReply)

+ fields:
    
    + `code`: integer: 
        <!--beg l desc_{{object_name}}_code -->
        
        <!--end l-->

    + `data`: : 
        <!--beg l desc_{{object_name}}_data -->
        
        <!--end l-->

    

