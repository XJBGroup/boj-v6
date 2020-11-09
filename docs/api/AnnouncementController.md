
# Api Group Documentation (AnnouncementController)

<!--beg l desc_AnnouncementController -->

<!--end l-->

## Apis


### CountAnnouncement

The uri/restful key of this method is `/announcement-count@GET`

<!--beg l desc_CountAnnouncement -->

<!--end l-->



### ListAnnouncement

The uri/restful key of this method is `/announcement-list@GET`

<!--beg l desc_ListAnnouncement -->

<!--end l-->


+ `page`: [integer](#integer): 
    <!--beg l desc_ListAnnouncement_page -->
    
    <!--end l-->


+ `page_size`: [integer](#integer): 
    <!--beg l desc_ListAnnouncement_page_size -->
    
    <!--end l-->



### DeleteAnnouncement

The uri/restful key of this method is `/announcement/{aid}@DELETE`

<!--beg l desc_DeleteAnnouncement -->

<!--end l-->


+ `aid`: [string](#string) (required): 
    <!--beg l desc_DeleteAnnouncement_aid -->
    
    <!--end l-->


+ `DeleteAnnouncementRequest`: [DeleteAnnouncementRequest](#DeleteAnnouncementRequest): 
    <!--beg l desc_DeleteAnnouncement_DeleteAnnouncementRequest -->
    
    <!--end l-->



### GetAnnouncement

The uri/restful key of this method is `/announcement/{aid}@GET`

<!--beg l desc_GetAnnouncement -->

<!--end l-->


+ `aid`: [string](#string) (required): 
    <!--beg l desc_GetAnnouncement_aid -->
    
    <!--end l-->



### PutAnnouncement

The uri/restful key of this method is `/announcement/{aid}@PUT`

<!--beg l desc_PutAnnouncement -->

<!--end l-->


+ `aid`: [string](#string) (required): 
    <!--beg l desc_PutAnnouncement_aid -->
    
    <!--end l-->


+ `PutAnnouncementRequest`: [PutAnnouncementRequest](#PutAnnouncementRequest): 
    <!--beg l desc_PutAnnouncement_PutAnnouncementRequest -->
    
    <!--end l-->



### PostAnnouncement

The uri/restful key of this method is `/announcement@POST`

<!--beg l desc_PostAnnouncement -->

<!--end l-->


+ `PostAnnouncementRequest`: [PostAnnouncementRequest](#PostAnnouncementRequest): 
    <!--beg l desc_PostAnnouncement_PostAnnouncementRequest -->
    
    <!--end l-->



## Local Object Reference




### [DeleteAnnouncementRequest](./ObjectModelSpec.md#DeleteAnnouncementRequest)

+ type: [object](#DeleteAnnouncementRequest)

+ fields:
    
    
### [PutAnnouncementRequest](./ObjectModelSpec.md#PutAnnouncementRequest)

+ type: [object](#PutAnnouncementRequest)

+ fields:
    
    + `content`: string: 
        <!--beg l desc_{{object_name}}_content -->
        
        <!--end l-->

    + `title`: string: 
        <!--beg l desc_{{object_name}}_title -->
        
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

    

