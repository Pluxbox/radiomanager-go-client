# ProgramResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**DeletedAt** | Pointer to **time.Time** |  | [optional] 
**ExternalStationId** | Pointer to **int64** |  | [optional] 
**ModelTypeId** | **int64** |  | 
**FieldValues** | Pointer to **map[string]interface{}** |  | [optional] 
**Title** | **string** |  | 
**Disabled** | Pointer to **bool** |  | [optional] 
**GenreId** | Pointer to **int64** |  | [optional] 
**GroupId** | Pointer to **int64** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ShortName** | Pointer to **string** |  | [optional] 
**MediumName** | Pointer to **string** |  | [optional] 
**Website** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Recommended** | Pointer to **bool** |  | [optional] 
**Language** | Pointer to **string** |  | [optional] 
**PtyCodeId** | Pointer to **int64** |  | [optional] 
**Genre** | Pointer to [**BroadcastRelationsGenre**](BroadcastRelationsGenre.md) |  | [optional] 
**Items** | Pointer to [**ProgramRelationsItems**](ProgramRelationsItems.md) |  | [optional] 
**Blocks** | Pointer to [**ProgramRelationsBlocks**](ProgramRelationsBlocks.md) |  | [optional] 
**Broadcasts** | Pointer to [**ProgramRelationsBroadcasts**](ProgramRelationsBroadcasts.md) |  | [optional] 
**Presenters** | Pointer to [**ProgramRelationsPresenters**](ProgramRelationsPresenters.md) |  | [optional] 
**Tags** | Pointer to [**ProgramRelationsTags**](ProgramRelationsTags.md) |  | [optional] 
**ModelType** | Pointer to [**BroadcastRelationsModelType**](BroadcastRelationsModelType.md) |  | [optional] 
**Group** | Pointer to [**BroadcastRelationsGroup**](BroadcastRelationsGroup.md) |  | [optional] 

## Methods

### NewProgramResult

`func NewProgramResult(id int64, modelTypeId int64, title string, ) *ProgramResult`

NewProgramResult instantiates a new ProgramResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProgramResultWithDefaults

`func NewProgramResultWithDefaults() *ProgramResult`

NewProgramResultWithDefaults instantiates a new ProgramResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProgramResult) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProgramResult) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProgramResult) SetId(v int64)`

SetId sets Id field to given value.


### GetUpdatedAt

`func (o *ProgramResult) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ProgramResult) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ProgramResult) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ProgramResult) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ProgramResult) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ProgramResult) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ProgramResult) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ProgramResult) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *ProgramResult) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *ProgramResult) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *ProgramResult) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *ProgramResult) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetExternalStationId

`func (o *ProgramResult) GetExternalStationId() int64`

GetExternalStationId returns the ExternalStationId field if non-nil, zero value otherwise.

### GetExternalStationIdOk

`func (o *ProgramResult) GetExternalStationIdOk() (*int64, bool)`

GetExternalStationIdOk returns a tuple with the ExternalStationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalStationId

`func (o *ProgramResult) SetExternalStationId(v int64)`

SetExternalStationId sets ExternalStationId field to given value.

### HasExternalStationId

`func (o *ProgramResult) HasExternalStationId() bool`

HasExternalStationId returns a boolean if a field has been set.

### GetModelTypeId

`func (o *ProgramResult) GetModelTypeId() int64`

GetModelTypeId returns the ModelTypeId field if non-nil, zero value otherwise.

### GetModelTypeIdOk

`func (o *ProgramResult) GetModelTypeIdOk() (*int64, bool)`

GetModelTypeIdOk returns a tuple with the ModelTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelTypeId

`func (o *ProgramResult) SetModelTypeId(v int64)`

SetModelTypeId sets ModelTypeId field to given value.


### GetFieldValues

`func (o *ProgramResult) GetFieldValues() map[string]interface{}`

GetFieldValues returns the FieldValues field if non-nil, zero value otherwise.

### GetFieldValuesOk

`func (o *ProgramResult) GetFieldValuesOk() (*map[string]interface{}, bool)`

GetFieldValuesOk returns a tuple with the FieldValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldValues

`func (o *ProgramResult) SetFieldValues(v map[string]interface{})`

SetFieldValues sets FieldValues field to given value.

### HasFieldValues

`func (o *ProgramResult) HasFieldValues() bool`

HasFieldValues returns a boolean if a field has been set.

### GetTitle

`func (o *ProgramResult) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ProgramResult) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ProgramResult) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDisabled

`func (o *ProgramResult) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *ProgramResult) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *ProgramResult) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *ProgramResult) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetGenreId

`func (o *ProgramResult) GetGenreId() int64`

GetGenreId returns the GenreId field if non-nil, zero value otherwise.

### GetGenreIdOk

`func (o *ProgramResult) GetGenreIdOk() (*int64, bool)`

GetGenreIdOk returns a tuple with the GenreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenreId

`func (o *ProgramResult) SetGenreId(v int64)`

SetGenreId sets GenreId field to given value.

### HasGenreId

`func (o *ProgramResult) HasGenreId() bool`

HasGenreId returns a boolean if a field has been set.

### GetGroupId

`func (o *ProgramResult) GetGroupId() int64`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ProgramResult) GetGroupIdOk() (*int64, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ProgramResult) SetGroupId(v int64)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *ProgramResult) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetDescription

`func (o *ProgramResult) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProgramResult) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProgramResult) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProgramResult) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetShortName

`func (o *ProgramResult) GetShortName() string`

GetShortName returns the ShortName field if non-nil, zero value otherwise.

### GetShortNameOk

`func (o *ProgramResult) GetShortNameOk() (*string, bool)`

GetShortNameOk returns a tuple with the ShortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortName

`func (o *ProgramResult) SetShortName(v string)`

SetShortName sets ShortName field to given value.

### HasShortName

`func (o *ProgramResult) HasShortName() bool`

HasShortName returns a boolean if a field has been set.

### GetMediumName

`func (o *ProgramResult) GetMediumName() string`

GetMediumName returns the MediumName field if non-nil, zero value otherwise.

### GetMediumNameOk

`func (o *ProgramResult) GetMediumNameOk() (*string, bool)`

GetMediumNameOk returns a tuple with the MediumName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediumName

`func (o *ProgramResult) SetMediumName(v string)`

SetMediumName sets MediumName field to given value.

### HasMediumName

`func (o *ProgramResult) HasMediumName() bool`

HasMediumName returns a boolean if a field has been set.

### GetWebsite

`func (o *ProgramResult) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *ProgramResult) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *ProgramResult) SetWebsite(v string)`

SetWebsite sets Website field to given value.

### HasWebsite

`func (o *ProgramResult) HasWebsite() bool`

HasWebsite returns a boolean if a field has been set.

### GetEmail

`func (o *ProgramResult) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ProgramResult) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ProgramResult) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ProgramResult) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetRecommended

`func (o *ProgramResult) GetRecommended() bool`

GetRecommended returns the Recommended field if non-nil, zero value otherwise.

### GetRecommendedOk

`func (o *ProgramResult) GetRecommendedOk() (*bool, bool)`

GetRecommendedOk returns a tuple with the Recommended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommended

`func (o *ProgramResult) SetRecommended(v bool)`

SetRecommended sets Recommended field to given value.

### HasRecommended

`func (o *ProgramResult) HasRecommended() bool`

HasRecommended returns a boolean if a field has been set.

### GetLanguage

`func (o *ProgramResult) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *ProgramResult) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *ProgramResult) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *ProgramResult) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetPtyCodeId

`func (o *ProgramResult) GetPtyCodeId() int64`

GetPtyCodeId returns the PtyCodeId field if non-nil, zero value otherwise.

### GetPtyCodeIdOk

`func (o *ProgramResult) GetPtyCodeIdOk() (*int64, bool)`

GetPtyCodeIdOk returns a tuple with the PtyCodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPtyCodeId

`func (o *ProgramResult) SetPtyCodeId(v int64)`

SetPtyCodeId sets PtyCodeId field to given value.

### HasPtyCodeId

`func (o *ProgramResult) HasPtyCodeId() bool`

HasPtyCodeId returns a boolean if a field has been set.

### GetGenre

`func (o *ProgramResult) GetGenre() BroadcastRelationsGenre`

GetGenre returns the Genre field if non-nil, zero value otherwise.

### GetGenreOk

`func (o *ProgramResult) GetGenreOk() (*BroadcastRelationsGenre, bool)`

GetGenreOk returns a tuple with the Genre field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenre

`func (o *ProgramResult) SetGenre(v BroadcastRelationsGenre)`

SetGenre sets Genre field to given value.

### HasGenre

`func (o *ProgramResult) HasGenre() bool`

HasGenre returns a boolean if a field has been set.

### GetItems

`func (o *ProgramResult) GetItems() ProgramRelationsItems`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ProgramResult) GetItemsOk() (*ProgramRelationsItems, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ProgramResult) SetItems(v ProgramRelationsItems)`

SetItems sets Items field to given value.

### HasItems

`func (o *ProgramResult) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetBlocks

`func (o *ProgramResult) GetBlocks() ProgramRelationsBlocks`

GetBlocks returns the Blocks field if non-nil, zero value otherwise.

### GetBlocksOk

`func (o *ProgramResult) GetBlocksOk() (*ProgramRelationsBlocks, bool)`

GetBlocksOk returns a tuple with the Blocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocks

`func (o *ProgramResult) SetBlocks(v ProgramRelationsBlocks)`

SetBlocks sets Blocks field to given value.

### HasBlocks

`func (o *ProgramResult) HasBlocks() bool`

HasBlocks returns a boolean if a field has been set.

### GetBroadcasts

`func (o *ProgramResult) GetBroadcasts() ProgramRelationsBroadcasts`

GetBroadcasts returns the Broadcasts field if non-nil, zero value otherwise.

### GetBroadcastsOk

`func (o *ProgramResult) GetBroadcastsOk() (*ProgramRelationsBroadcasts, bool)`

GetBroadcastsOk returns a tuple with the Broadcasts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBroadcasts

`func (o *ProgramResult) SetBroadcasts(v ProgramRelationsBroadcasts)`

SetBroadcasts sets Broadcasts field to given value.

### HasBroadcasts

`func (o *ProgramResult) HasBroadcasts() bool`

HasBroadcasts returns a boolean if a field has been set.

### GetPresenters

`func (o *ProgramResult) GetPresenters() ProgramRelationsPresenters`

GetPresenters returns the Presenters field if non-nil, zero value otherwise.

### GetPresentersOk

`func (o *ProgramResult) GetPresentersOk() (*ProgramRelationsPresenters, bool)`

GetPresentersOk returns a tuple with the Presenters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresenters

`func (o *ProgramResult) SetPresenters(v ProgramRelationsPresenters)`

SetPresenters sets Presenters field to given value.

### HasPresenters

`func (o *ProgramResult) HasPresenters() bool`

HasPresenters returns a boolean if a field has been set.

### GetTags

`func (o *ProgramResult) GetTags() ProgramRelationsTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ProgramResult) GetTagsOk() (*ProgramRelationsTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ProgramResult) SetTags(v ProgramRelationsTags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ProgramResult) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetModelType

`func (o *ProgramResult) GetModelType() BroadcastRelationsModelType`

GetModelType returns the ModelType field if non-nil, zero value otherwise.

### GetModelTypeOk

`func (o *ProgramResult) GetModelTypeOk() (*BroadcastRelationsModelType, bool)`

GetModelTypeOk returns a tuple with the ModelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelType

`func (o *ProgramResult) SetModelType(v BroadcastRelationsModelType)`

SetModelType sets ModelType field to given value.

### HasModelType

`func (o *ProgramResult) HasModelType() bool`

HasModelType returns a boolean if a field has been set.

### GetGroup

`func (o *ProgramResult) GetGroup() BroadcastRelationsGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *ProgramResult) GetGroupOk() (*BroadcastRelationsGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *ProgramResult) SetGroup(v BroadcastRelationsGroup)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *ProgramResult) HasGroup() bool`

HasGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


