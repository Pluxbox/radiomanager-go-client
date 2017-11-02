# Go API client for radiomanagerclient

Pluxbox RadioManager gives you the power, flexibility and speed you always wanted in a lightweight and easy-to-use web-based radio solution. With Pluxbox RadioManager you can organise your radio workflow and automate your omnichannel communication with your listeners. We offer wide range specialised services for the radio and connections like Hybrid Radio, Visual Radio, your website and social media without losing focus on your broadcast. For more information visit: pluxbox.com

## Overview
- API version: 2.0
- Package version: 1.1.2
- Build package: io.swagger.codegen.languages.GoClientCodegen
For more information, please visit [https://pluxbox.com](https://pluxbox.com)

## Installation
Put the package under your project folder and add the following in import:
```
    "./radiomanagerclient"
```

## Documentation for API Endpoints

All URIs are relative to *https://radiomanager.io/api/v2*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*BlockApi* | [**GetBlockById**](docs/BlockApi.md#getblockbyid) | **Get** /blocks/{id} | Get block by id
*BlockApi* | [**GetCurrentBlock**](docs/BlockApi.md#getcurrentblock) | **Get** /blocks/current | Get current Block
*BlockApi* | [**GetNextBlock**](docs/BlockApi.md#getnextblock) | **Get** /blocks/next | Get next Block
*BlockApi* | [**ListBlocks**](docs/BlockApi.md#listblocks) | **Get** /blocks | Get a list of all blocks currently in your station.
*BroadcastApi* | [**CreateBroadcast**](docs/BroadcastApi.md#createbroadcast) | **Post** /broadcasts | Create broadcast.
*BroadcastApi* | [**DeleteBroadcastById**](docs/BroadcastApi.md#deletebroadcastbyid) | **Delete** /broadcasts/{id} | Delete broadcast by id
*BroadcastApi* | [**GetBroadcastById**](docs/BroadcastApi.md#getbroadcastbyid) | **Get** /broadcasts/{id} | Get broadcast by id
*BroadcastApi* | [**GetCurrentBroadcast**](docs/BroadcastApi.md#getcurrentbroadcast) | **Get** /broadcasts/current | Get current Broadcast
*BroadcastApi* | [**GetDailyEPG**](docs/BroadcastApi.md#getdailyepg) | **Get** /broadcasts/epg/daily | Get daily EPG
*BroadcastApi* | [**GetEPGByDate**](docs/BroadcastApi.md#getepgbydate) | **Get** /broadcasts/epg | Get EPG by date
*BroadcastApi* | [**GetNextBroadcast**](docs/BroadcastApi.md#getnextbroadcast) | **Get** /broadcasts/next | Get next Broadcast
*BroadcastApi* | [**GetWeeklyEPG**](docs/BroadcastApi.md#getweeklyepg) | **Get** /broadcasts/epg/weekly | Get weekly EPG
*BroadcastApi* | [**ListBroadcasts**](docs/BroadcastApi.md#listbroadcasts) | **Get** /broadcasts | Get all broadcasts.
*BroadcastApi* | [**PrintBroadcastById**](docs/BroadcastApi.md#printbroadcastbyid) | **Get** /broadcasts/print/{id} | Print Broadcast by id
*BroadcastApi* | [**UpdateBroadcastByID**](docs/BroadcastApi.md#updatebroadcastbyid) | **Patch** /broadcasts/{id} | Update broadcast by id
*CampaignApi* | [**CreateCampaign**](docs/CampaignApi.md#createcampaign) | **Post** /campaigns | Create campaign.
*CampaignApi* | [**DeleteCampaignById**](docs/CampaignApi.md#deletecampaignbyid) | **Delete** /campaigns/{id} | Delete campaign by id
*CampaignApi* | [**GetCampaignById**](docs/CampaignApi.md#getcampaignbyid) | **Get** /campaigns/{id} | Get campaign by id
*CampaignApi* | [**ListCampaigns**](docs/CampaignApi.md#listcampaigns) | **Get** /campaigns | Get all campaigns.
*CampaignApi* | [**UpdateCampaignByID**](docs/CampaignApi.md#updatecampaignbyid) | **Patch** /campaigns/{id} | Update campaign by id
*ContactApi* | [**CreateContact**](docs/ContactApi.md#createcontact) | **Post** /contacts | Create contact.
*ContactApi* | [**DeleteContactById**](docs/ContactApi.md#deletecontactbyid) | **Delete** /contacts/{id} | Delete contact by id
*ContactApi* | [**GetContactById**](docs/ContactApi.md#getcontactbyid) | **Get** /contacts/{id} | Get contact by id
*ContactApi* | [**ListContacts**](docs/ContactApi.md#listcontacts) | **Get** /contacts | Get all contacts.
*ContactApi* | [**UpdateContactByID**](docs/ContactApi.md#updatecontactbyid) | **Patch** /contacts/{id} | Update contact by id
*GenreApi* | [**GetGenreById**](docs/GenreApi.md#getgenrebyid) | **Get** /genres/{id} | Get genre by id
*GenreApi* | [**ListGenres**](docs/GenreApi.md#listgenres) | **Get** /genres | List all genres.
*ItemApi* | [**CreateItem**](docs/ItemApi.md#createitem) | **Post** /items | Create an new item.
*ItemApi* | [**CurrentItemPostStructure**](docs/ItemApi.md#currentitempoststructure) | **Post** /items/current/structure | Post a current playing item, keep structure
*ItemApi* | [**CurrentItemPostTiming**](docs/ItemApi.md#currentitemposttiming) | **Post** /items/current/timing | Post a current playing item
*ItemApi* | [**DeleteItemById**](docs/ItemApi.md#deleteitembyid) | **Delete** /items/{id} | Delete item by ID.
*ItemApi* | [**GetItemById**](docs/ItemApi.md#getitembyid) | **Get** /items/{id} | Get extended item details by ID.
*ItemApi* | [**ListItems**](docs/ItemApi.md#listitems) | **Get** /items | Get a list of all the items currently in your station.
*ItemApi* | [**PlaylistPostStructure**](docs/ItemApi.md#playlistpoststructure) | **Post** /items/playlist/structure | Post a playlist, keep current structure
*ItemApi* | [**PlaylistPostTiming**](docs/ItemApi.md#playlistposttiming) | **Post** /items/playlist/timing | Post a playlist
*ItemApi* | [**UpdateItemById**](docs/ItemApi.md#updateitembyid) | **Patch** /items/{id} | Update extended item details by ID.
*ModelTypeApi* | [**GetModelTypeById**](docs/ModelTypeApi.md#getmodeltypebyid) | **Get** /model_types/{id} | Get modelType by id
*ModelTypeApi* | [**ListModelTypes**](docs/ModelTypeApi.md#listmodeltypes) | **Get** /model_types | Get all modelTypes.
*PresenterApi* | [**CreatePresenter**](docs/PresenterApi.md#createpresenter) | **Post** /presenters | Create presenter.
*PresenterApi* | [**DeletePresenterById**](docs/PresenterApi.md#deletepresenterbyid) | **Delete** /presenters/{id} | Delete presenter by id
*PresenterApi* | [**GetPresenterById**](docs/PresenterApi.md#getpresenterbyid) | **Get** /presenters/{id} | Get presenter by id
*PresenterApi* | [**ListPresenters**](docs/PresenterApi.md#listpresenters) | **Get** /presenters | Get all presenters.
*PresenterApi* | [**UpdatePresenterByID**](docs/PresenterApi.md#updatepresenterbyid) | **Patch** /presenters/{id} | Update presenter by id
*ProgramApi* | [**CreateProgram**](docs/ProgramApi.md#createprogram) | **Post** /programs | Create program.
*ProgramApi* | [**DeleteProgramById**](docs/ProgramApi.md#deleteprogrambyid) | **Delete** /programs/{id} | Delete program by id
*ProgramApi* | [**GetProgramById**](docs/ProgramApi.md#getprogrambyid) | **Get** /programs/{id} | Get program by id
*ProgramApi* | [**ListPrograms**](docs/ProgramApi.md#listprograms) | **Get** /programs | Get all programs.
*ProgramApi* | [**UpdateProgramByID**](docs/ProgramApi.md#updateprogrambyid) | **Patch** /programs/{id} | Update program by id
*StationApi* | [**GetStation**](docs/StationApi.md#getstation) | **Get** /station | Get own station only
*StoryApi* | [**CreateStory**](docs/StoryApi.md#createstory) | **Post** /stories | Create story.
*StoryApi* | [**DeleteStoryById**](docs/StoryApi.md#deletestorybyid) | **Delete** /stories/{id} | Delete story by id
*StoryApi* | [**GetStoryById**](docs/StoryApi.md#getstorybyid) | **Get** /stories/{id} | Get story by id
*StoryApi* | [**ListStories**](docs/StoryApi.md#liststories) | **Get** /stories | Get all stories.
*StoryApi* | [**UpdateStoryByID**](docs/StoryApi.md#updatestorybyid) | **Patch** /stories/{id} | Update story by id
*StringApi* | [**GetStringsByName**](docs/StringApi.md#getstringsbyname) | **Get** /strings/{name} | Get Strings (formatted)
*TagApi* | [**CreateTag**](docs/TagApi.md#createtag) | **Post** /tags | Create tag.
*TagApi* | [**DeleteTagById**](docs/TagApi.md#deletetagbyid) | **Delete** /tags/{id} | Delete tag by id
*TagApi* | [**GetTagById**](docs/TagApi.md#gettagbyid) | **Get** /tags/{id} | Get tags by id
*TagApi* | [**ListTags**](docs/TagApi.md#listtags) | **Get** /tags | Get a list of all the tags currently in your station.
*TagApi* | [**UpdateTagByID**](docs/TagApi.md#updatetagbyid) | **Patch** /tags/{id} | Update tag by id
*UserApi* | [**DeleteUserById**](docs/UserApi.md#deleteuserbyid) | **Delete** /users/{id} | Remove user from station by Id
*UserApi* | [**GetUserById**](docs/UserApi.md#getuserbyid) | **Get** /users/{id} | Get user by id
*UserApi* | [**InviteUserByMail**](docs/UserApi.md#inviteuserbymail) | **Post** /users/invite | Invite user by mail
*UserApi* | [**ListUsers**](docs/UserApi.md#listusers) | **Get** /users | Get all users.
*VisualSlideApi* | [**GetVisualSlide**](docs/VisualSlideApi.md#getvisualslide) | **Get** /visual | Get Visual Slide Image as Base64


## Documentation For Models

 - [Block](docs/Block.md)
 - [BlockRelations](docs/BlockRelations.md)
 - [BlockRelationsBroadcast](docs/BlockRelationsBroadcast.md)
 - [BlockRelationsBroadcastParams](docs/BlockRelationsBroadcastParams.md)
 - [BlockRelationsItems](docs/BlockRelationsItems.md)
 - [BlockRelationsItemsParams](docs/BlockRelationsItemsParams.md)
 - [BlockRelationsProgram](docs/BlockRelationsProgram.md)
 - [BlockResults](docs/BlockResults.md)
 - [Broadcast](docs/Broadcast.md)
 - [BroadcastEpgDay](docs/BroadcastEpgDay.md)
 - [BroadcastEpgRelations](docs/BroadcastEpgRelations.md)
 - [BroadcastInputOnly](docs/BroadcastInputOnly.md)
 - [BroadcastOutputOnly](docs/BroadcastOutputOnly.md)
 - [BroadcastRelations](docs/BroadcastRelations.md)
 - [BroadcastRelationsBlocks](docs/BroadcastRelationsBlocks.md)
 - [BroadcastRelationsItems](docs/BroadcastRelationsItems.md)
 - [BroadcastRelationsItemsParams](docs/BroadcastRelationsItemsParams.md)
 - [BroadcastRelationsModelType](docs/BroadcastRelationsModelType.md)
 - [BroadcastRelationsPresenters](docs/BroadcastRelationsPresenters.md)
 - [BroadcastRelationsTags](docs/BroadcastRelationsTags.md)
 - [BroadcastResults](docs/BroadcastResults.md)
 - [Campaign](docs/Campaign.md)
 - [CampaignOutputOnly](docs/CampaignOutputOnly.md)
 - [CampaignRelations](docs/CampaignRelations.md)
 - [CampaignRelationsItems](docs/CampaignRelationsItems.md)
 - [CampaignRelationsItemsParams](docs/CampaignRelationsItemsParams.md)
 - [CampaignResults](docs/CampaignResults.md)
 - [Contact](docs/Contact.md)
 - [ContactOutputOnly](docs/ContactOutputOnly.md)
 - [ContactRelations](docs/ContactRelations.md)
 - [ContactRelationsItems](docs/ContactRelationsItems.md)
 - [ContactRelationsTags](docs/ContactRelationsTags.md)
 - [ContactRelationsTagsParams](docs/ContactRelationsTagsParams.md)
 - [ContactResults](docs/ContactResults.md)
 - [Data](docs/Data.md)
 - [Data1](docs/Data1.md)
 - [EpgResults](docs/EpgResults.md)
 - [Forbidden](docs/Forbidden.md)
 - [Genre](docs/Genre.md)
 - [GenreOutputOnly](docs/GenreOutputOnly.md)
 - [GenreRelations](docs/GenreRelations.md)
 - [GenreRelationsBroadcasts](docs/GenreRelationsBroadcasts.md)
 - [GenreRelationsBroadcastsParams](docs/GenreRelationsBroadcastsParams.md)
 - [GenreRelationsPrograms](docs/GenreRelationsPrograms.md)
 - [GenreResults](docs/GenreResults.md)
 - [ImportItem](docs/ImportItem.md)
 - [InlineResponse202](docs/InlineResponse202.md)
 - [InternalServerError](docs/InternalServerError.md)
 - [InviteUserData](docs/InviteUserData.md)
 - [InviteUserSuccess](docs/InviteUserSuccess.md)
 - [Item](docs/Item.md)
 - [ItemInputOnly](docs/ItemInputOnly.md)
 - [ItemOutputOnly](docs/ItemOutputOnly.md)
 - [ItemRelations](docs/ItemRelations.md)
 - [ItemRelationsBlock](docs/ItemRelationsBlock.md)
 - [ItemRelationsCampaign](docs/ItemRelationsCampaign.md)
 - [ItemRelationsContacts](docs/ItemRelationsContacts.md)
 - [ItemRelationsContactsParams](docs/ItemRelationsContactsParams.md)
 - [ItemRelationsProgram](docs/ItemRelationsProgram.md)
 - [ItemRelationsTags](docs/ItemRelationsTags.md)
 - [ItemResults](docs/ItemResults.md)
 - [ModelType](docs/ModelType.md)
 - [ModelTypeOptions](docs/ModelTypeOptions.md)
 - [ModelTypeOutputOnly](docs/ModelTypeOutputOnly.md)
 - [ModelTypeRelations](docs/ModelTypeRelations.md)
 - [ModelTypeRelationsBroadcasts](docs/ModelTypeRelationsBroadcasts.md)
 - [ModelTypeRelationsCampaigns](docs/ModelTypeRelationsCampaigns.md)
 - [ModelTypeRelationsCampaignsParams](docs/ModelTypeRelationsCampaignsParams.md)
 - [ModelTypeRelationsContacts](docs/ModelTypeRelationsContacts.md)
 - [ModelTypeRelationsItems](docs/ModelTypeRelationsItems.md)
 - [ModelTypeRelationsPresenters](docs/ModelTypeRelationsPresenters.md)
 - [ModelTypeRelationsPrograms](docs/ModelTypeRelationsPrograms.md)
 - [ModelTypeResults](docs/ModelTypeResults.md)
 - [NotFound](docs/NotFound.md)
 - [PostSuccess](docs/PostSuccess.md)
 - [Presenter](docs/Presenter.md)
 - [PresenterOutputOnly](docs/PresenterOutputOnly.md)
 - [PresenterRelations](docs/PresenterRelations.md)
 - [PresenterRelationsBroadcasts](docs/PresenterRelationsBroadcasts.md)
 - [PresenterRelationsPrograms](docs/PresenterRelationsPrograms.md)
 - [PresenterRelationsProgramsParams](docs/PresenterRelationsProgramsParams.md)
 - [PresenterResults](docs/PresenterResults.md)
 - [Program](docs/Program.md)
 - [ProgramInputOnly](docs/ProgramInputOnly.md)
 - [ProgramOutputOnly](docs/ProgramOutputOnly.md)
 - [ProgramRelations](docs/ProgramRelations.md)
 - [ProgramRelationsBlocks](docs/ProgramRelationsBlocks.md)
 - [ProgramRelationsBroadcasts](docs/ProgramRelationsBroadcasts.md)
 - [ProgramRelationsItems](docs/ProgramRelationsItems.md)
 - [ProgramRelationsItemsParams](docs/ProgramRelationsItemsParams.md)
 - [ProgramRelationsPresenters](docs/ProgramRelationsPresenters.md)
 - [ProgramRelationsTags](docs/ProgramRelationsTags.md)
 - [ProgramResults](docs/ProgramResults.md)
 - [ReadOnly](docs/ReadOnly.md)
 - [RelationsPlaceholder](docs/RelationsPlaceholder.md)
 - [StationResult](docs/StationResult.md)
 - [StationResultStation](docs/StationResultStation.md)
 - [Story](docs/Story.md)
 - [StoryInputOnly](docs/StoryInputOnly.md)
 - [StoryOutputOnly](docs/StoryOutputOnly.md)
 - [StoryRelations](docs/StoryRelations.md)
 - [StoryRelationsItems](docs/StoryRelationsItems.md)
 - [StoryRelationsTags](docs/StoryRelationsTags.md)
 - [StoryRelationsTagsParams](docs/StoryRelationsTagsParams.md)
 - [StoryResults](docs/StoryResults.md)
 - [Success](docs/Success.md)
 - [Tag](docs/Tag.md)
 - [TagOutputOnly](docs/TagOutputOnly.md)
 - [TagRelations](docs/TagRelations.md)
 - [TagRelationsBroadcasts](docs/TagRelationsBroadcasts.md)
 - [TagRelationsBroadcastsParams](docs/TagRelationsBroadcastsParams.md)
 - [TagRelationsContacts](docs/TagRelationsContacts.md)
 - [TagRelationsItems](docs/TagRelationsItems.md)
 - [TagRelationsPrograms](docs/TagRelationsPrograms.md)
 - [TagResults](docs/TagResults.md)
 - [TextString](docs/TextString.md)
 - [TooManyRequests](docs/TooManyRequests.md)
 - [UnprocessableEntity](docs/UnprocessableEntity.md)
 - [UserResult](docs/UserResult.md)
 - [UserResultSettings](docs/UserResultSettings.md)
 - [UserResults](docs/UserResults.md)
 - [VisualResult](docs/VisualResult.md)
 - [BlockResult](docs/BlockResult.md)
 - [BroadcastDataInput](docs/BroadcastDataInput.md)
 - [BroadcastEpgResult](docs/BroadcastEpgResult.md)
 - [BroadcastResult](docs/BroadcastResult.md)
 - [CampaignDataInput](docs/CampaignDataInput.md)
 - [CampaignResult](docs/CampaignResult.md)
 - [ContactDataInput](docs/ContactDataInput.md)
 - [ContactResult](docs/ContactResult.md)
 - [GenreResult](docs/GenreResult.md)
 - [ItemDataInput](docs/ItemDataInput.md)
 - [ItemResult](docs/ItemResult.md)
 - [ModelTypeResult](docs/ModelTypeResult.md)
 - [PresenterDataInput](docs/PresenterDataInput.md)
 - [PresenterEpgResult](docs/PresenterEpgResult.md)
 - [PresenterResult](docs/PresenterResult.md)
 - [ProgramDataInput](docs/ProgramDataInput.md)
 - [ProgramResult](docs/ProgramResult.md)
 - [StoryDataInput](docs/StoryDataInput.md)
 - [StoryResult](docs/StoryResult.md)
 - [TagDataInput](docs/TagDataInput.md)
 - [TagResult](docs/TagResult.md)


## Documentation For Authorization


## API Key

- **Type**: API key 
- **API key parameter name**: api-key
- **Location**: HTTP header


## Author

support@pluxbox.com

