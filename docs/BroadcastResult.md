# BroadcastResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**DeletedAt** | Pointer to **time.Time** |  | [optional] 
**ExternalStationId** | Pointer to **int64** |  | [optional] 
**ProgramId** | Pointer to **int64** |  | [optional] 
**ModelTypeId** | Pointer to **int64** |  | [optional] 
**StationId** | Pointer to **int64** |  | [optional] 
**FieldValues** | Pointer to **map[string]interface{}** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Start** | Pointer to **time.Time** |  | [optional] 
**Stop** | Pointer to **time.Time** |  | [optional] 
**GenreId** | Pointer to **int64** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ShortName** | Pointer to **string** |  | [optional] 
**MediumName** | Pointer to **string** |  | [optional] 
**Website** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Recommended** | Pointer to **bool** |  | [optional] 
**Language** | Pointer to **string** |  | [optional] 
**Published** | Pointer to **bool** |  | [optional] 
**RepetitionUid** | Pointer to **string** |  | [optional] 
**RepetitionType** | Pointer to **string** |  | [optional] 
**RepetitionEnd** | Pointer to **time.Time** |  | [optional] 
**RepetitionStart** | Pointer to **time.Time** |  | [optional] 
**RepetitionDays** | Pointer to **string** |  | [optional] 
**PtyCodeId** | Pointer to **int64** |  | [optional] 
**PlannedInEpg** | Pointer to **int32** |  | [optional] 
**GroupId** | Pointer to **int64** |  | [optional] 
**Genre** | Pointer to [**BroadcastRelationsGenre**](BroadcastRelationsGenre.md) |  | [optional] 
**Items** | Pointer to [**BroadcastRelationsItems**](BroadcastRelationsItems.md) |  | [optional] 
**Blocks** | Pointer to [**BroadcastRelationsBlocks**](BroadcastRelationsBlocks.md) |  | [optional] 
**Program** | Pointer to [**BlockRelationsProgram**](BlockRelationsProgram.md) |  | [optional] 
**Tags** | Pointer to [**BroadcastRelationsTags**](BroadcastRelationsTags.md) |  | [optional] 
**Presenters** | Pointer to [**BroadcastRelationsPresenters**](BroadcastRelationsPresenters.md) |  | [optional] 
**ModelType** | Pointer to [**BroadcastRelationsModelType**](BroadcastRelationsModelType.md) |  | [optional] 
**Group** | Pointer to [**BroadcastRelationsGroup**](BroadcastRelationsGroup.md) |  | [optional] 

## Methods

### NewBroadcastResult

`func NewBroadcastResult(id int64, ) *BroadcastResult`

NewBroadcastResult instantiates a new BroadcastResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBroadcastResultWithDefaults

`func NewBroadcastResultWithDefaults() *BroadcastResult`

NewBroadcastResultWithDefaults instantiates a new BroadcastResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BroadcastResult) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BroadcastResult) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BroadcastResult) SetId(v int64)`

SetId sets Id field to given value.


### GetUpdatedAt

`func (o *BroadcastResult) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BroadcastResult) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BroadcastResult) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *BroadcastResult) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *BroadcastResult) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BroadcastResult) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BroadcastResult) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BroadcastResult) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *BroadcastResult) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *BroadcastResult) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *BroadcastResult) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *BroadcastResult) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetExternalStationId

`func (o *BroadcastResult) GetExternalStationId() int64`

GetExternalStationId returns the ExternalStationId field if non-nil, zero value otherwise.

### GetExternalStationIdOk

`func (o *BroadcastResult) GetExternalStationIdOk() (*int64, bool)`

GetExternalStationIdOk returns a tuple with the ExternalStationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalStationId

`func (o *BroadcastResult) SetExternalStationId(v int64)`

SetExternalStationId sets ExternalStationId field to given value.

### HasExternalStationId

`func (o *BroadcastResult) HasExternalStationId() bool`

HasExternalStationId returns a boolean if a field has been set.

### GetProgramId

`func (o *BroadcastResult) GetProgramId() int64`

GetProgramId returns the ProgramId field if non-nil, zero value otherwise.

### GetProgramIdOk

`func (o *BroadcastResult) GetProgramIdOk() (*int64, bool)`

GetProgramIdOk returns a tuple with the ProgramId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgramId

`func (o *BroadcastResult) SetProgramId(v int64)`

SetProgramId sets ProgramId field to given value.

### HasProgramId

`func (o *BroadcastResult) HasProgramId() bool`

HasProgramId returns a boolean if a field has been set.

### GetModelTypeId

`func (o *BroadcastResult) GetModelTypeId() int64`

GetModelTypeId returns the ModelTypeId field if non-nil, zero value otherwise.

### GetModelTypeIdOk

`func (o *BroadcastResult) GetModelTypeIdOk() (*int64, bool)`

GetModelTypeIdOk returns a tuple with the ModelTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelTypeId

`func (o *BroadcastResult) SetModelTypeId(v int64)`

SetModelTypeId sets ModelTypeId field to given value.

### HasModelTypeId

`func (o *BroadcastResult) HasModelTypeId() bool`

HasModelTypeId returns a boolean if a field has been set.

### GetStationId

`func (o *BroadcastResult) GetStationId() int64`

GetStationId returns the StationId field if non-nil, zero value otherwise.

### GetStationIdOk

`func (o *BroadcastResult) GetStationIdOk() (*int64, bool)`

GetStationIdOk returns a tuple with the StationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStationId

`func (o *BroadcastResult) SetStationId(v int64)`

SetStationId sets StationId field to given value.

### HasStationId

`func (o *BroadcastResult) HasStationId() bool`

HasStationId returns a boolean if a field has been set.

### GetFieldValues

`func (o *BroadcastResult) GetFieldValues() map[string]interface{}`

GetFieldValues returns the FieldValues field if non-nil, zero value otherwise.

### GetFieldValuesOk

`func (o *BroadcastResult) GetFieldValuesOk() (*map[string]interface{}, bool)`

GetFieldValuesOk returns a tuple with the FieldValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldValues

`func (o *BroadcastResult) SetFieldValues(v map[string]interface{})`

SetFieldValues sets FieldValues field to given value.

### HasFieldValues

`func (o *BroadcastResult) HasFieldValues() bool`

HasFieldValues returns a boolean if a field has been set.

### GetTitle

`func (o *BroadcastResult) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *BroadcastResult) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *BroadcastResult) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *BroadcastResult) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetStart

`func (o *BroadcastResult) GetStart() time.Time`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *BroadcastResult) GetStartOk() (*time.Time, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *BroadcastResult) SetStart(v time.Time)`

SetStart sets Start field to given value.

### HasStart

`func (o *BroadcastResult) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetStop

`func (o *BroadcastResult) GetStop() time.Time`

GetStop returns the Stop field if non-nil, zero value otherwise.

### GetStopOk

`func (o *BroadcastResult) GetStopOk() (*time.Time, bool)`

GetStopOk returns a tuple with the Stop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStop

`func (o *BroadcastResult) SetStop(v time.Time)`

SetStop sets Stop field to given value.

### HasStop

`func (o *BroadcastResult) HasStop() bool`

HasStop returns a boolean if a field has been set.

### GetGenreId

`func (o *BroadcastResult) GetGenreId() int64`

GetGenreId returns the GenreId field if non-nil, zero value otherwise.

### GetGenreIdOk

`func (o *BroadcastResult) GetGenreIdOk() (*int64, bool)`

GetGenreIdOk returns a tuple with the GenreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenreId

`func (o *BroadcastResult) SetGenreId(v int64)`

SetGenreId sets GenreId field to given value.

### HasGenreId

`func (o *BroadcastResult) HasGenreId() bool`

HasGenreId returns a boolean if a field has been set.

### GetDescription

`func (o *BroadcastResult) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BroadcastResult) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BroadcastResult) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BroadcastResult) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetShortName

`func (o *BroadcastResult) GetShortName() string`

GetShortName returns the ShortName field if non-nil, zero value otherwise.

### GetShortNameOk

`func (o *BroadcastResult) GetShortNameOk() (*string, bool)`

GetShortNameOk returns a tuple with the ShortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortName

`func (o *BroadcastResult) SetShortName(v string)`

SetShortName sets ShortName field to given value.

### HasShortName

`func (o *BroadcastResult) HasShortName() bool`

HasShortName returns a boolean if a field has been set.

### GetMediumName

`func (o *BroadcastResult) GetMediumName() string`

GetMediumName returns the MediumName field if non-nil, zero value otherwise.

### GetMediumNameOk

`func (o *BroadcastResult) GetMediumNameOk() (*string, bool)`

GetMediumNameOk returns a tuple with the MediumName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediumName

`func (o *BroadcastResult) SetMediumName(v string)`

SetMediumName sets MediumName field to given value.

### HasMediumName

`func (o *BroadcastResult) HasMediumName() bool`

HasMediumName returns a boolean if a field has been set.

### GetWebsite

`func (o *BroadcastResult) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *BroadcastResult) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *BroadcastResult) SetWebsite(v string)`

SetWebsite sets Website field to given value.

### HasWebsite

`func (o *BroadcastResult) HasWebsite() bool`

HasWebsite returns a boolean if a field has been set.

### GetEmail

`func (o *BroadcastResult) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *BroadcastResult) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *BroadcastResult) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *BroadcastResult) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetRecommended

`func (o *BroadcastResult) GetRecommended() bool`

GetRecommended returns the Recommended field if non-nil, zero value otherwise.

### GetRecommendedOk

`func (o *BroadcastResult) GetRecommendedOk() (*bool, bool)`

GetRecommendedOk returns a tuple with the Recommended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommended

`func (o *BroadcastResult) SetRecommended(v bool)`

SetRecommended sets Recommended field to given value.

### HasRecommended

`func (o *BroadcastResult) HasRecommended() bool`

HasRecommended returns a boolean if a field has been set.

### GetLanguage

`func (o *BroadcastResult) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *BroadcastResult) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *BroadcastResult) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *BroadcastResult) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetPublished

`func (o *BroadcastResult) GetPublished() bool`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *BroadcastResult) GetPublishedOk() (*bool, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *BroadcastResult) SetPublished(v bool)`

SetPublished sets Published field to given value.

### HasPublished

`func (o *BroadcastResult) HasPublished() bool`

HasPublished returns a boolean if a field has been set.

### GetRepetitionUid

`func (o *BroadcastResult) GetRepetitionUid() string`

GetRepetitionUid returns the RepetitionUid field if non-nil, zero value otherwise.

### GetRepetitionUidOk

`func (o *BroadcastResult) GetRepetitionUidOk() (*string, bool)`

GetRepetitionUidOk returns a tuple with the RepetitionUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepetitionUid

`func (o *BroadcastResult) SetRepetitionUid(v string)`

SetRepetitionUid sets RepetitionUid field to given value.

### HasRepetitionUid

`func (o *BroadcastResult) HasRepetitionUid() bool`

HasRepetitionUid returns a boolean if a field has been set.

### GetRepetitionType

`func (o *BroadcastResult) GetRepetitionType() string`

GetRepetitionType returns the RepetitionType field if non-nil, zero value otherwise.

### GetRepetitionTypeOk

`func (o *BroadcastResult) GetRepetitionTypeOk() (*string, bool)`

GetRepetitionTypeOk returns a tuple with the RepetitionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepetitionType

`func (o *BroadcastResult) SetRepetitionType(v string)`

SetRepetitionType sets RepetitionType field to given value.

### HasRepetitionType

`func (o *BroadcastResult) HasRepetitionType() bool`

HasRepetitionType returns a boolean if a field has been set.

### GetRepetitionEnd

`func (o *BroadcastResult) GetRepetitionEnd() time.Time`

GetRepetitionEnd returns the RepetitionEnd field if non-nil, zero value otherwise.

### GetRepetitionEndOk

`func (o *BroadcastResult) GetRepetitionEndOk() (*time.Time, bool)`

GetRepetitionEndOk returns a tuple with the RepetitionEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepetitionEnd

`func (o *BroadcastResult) SetRepetitionEnd(v time.Time)`

SetRepetitionEnd sets RepetitionEnd field to given value.

### HasRepetitionEnd

`func (o *BroadcastResult) HasRepetitionEnd() bool`

HasRepetitionEnd returns a boolean if a field has been set.

### GetRepetitionStart

`func (o *BroadcastResult) GetRepetitionStart() time.Time`

GetRepetitionStart returns the RepetitionStart field if non-nil, zero value otherwise.

### GetRepetitionStartOk

`func (o *BroadcastResult) GetRepetitionStartOk() (*time.Time, bool)`

GetRepetitionStartOk returns a tuple with the RepetitionStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepetitionStart

`func (o *BroadcastResult) SetRepetitionStart(v time.Time)`

SetRepetitionStart sets RepetitionStart field to given value.

### HasRepetitionStart

`func (o *BroadcastResult) HasRepetitionStart() bool`

HasRepetitionStart returns a boolean if a field has been set.

### GetRepetitionDays

`func (o *BroadcastResult) GetRepetitionDays() string`

GetRepetitionDays returns the RepetitionDays field if non-nil, zero value otherwise.

### GetRepetitionDaysOk

`func (o *BroadcastResult) GetRepetitionDaysOk() (*string, bool)`

GetRepetitionDaysOk returns a tuple with the RepetitionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepetitionDays

`func (o *BroadcastResult) SetRepetitionDays(v string)`

SetRepetitionDays sets RepetitionDays field to given value.

### HasRepetitionDays

`func (o *BroadcastResult) HasRepetitionDays() bool`

HasRepetitionDays returns a boolean if a field has been set.

### GetPtyCodeId

`func (o *BroadcastResult) GetPtyCodeId() int64`

GetPtyCodeId returns the PtyCodeId field if non-nil, zero value otherwise.

### GetPtyCodeIdOk

`func (o *BroadcastResult) GetPtyCodeIdOk() (*int64, bool)`

GetPtyCodeIdOk returns a tuple with the PtyCodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPtyCodeId

`func (o *BroadcastResult) SetPtyCodeId(v int64)`

SetPtyCodeId sets PtyCodeId field to given value.

### HasPtyCodeId

`func (o *BroadcastResult) HasPtyCodeId() bool`

HasPtyCodeId returns a boolean if a field has been set.

### GetPlannedInEpg

`func (o *BroadcastResult) GetPlannedInEpg() int32`

GetPlannedInEpg returns the PlannedInEpg field if non-nil, zero value otherwise.

### GetPlannedInEpgOk

`func (o *BroadcastResult) GetPlannedInEpgOk() (*int32, bool)`

GetPlannedInEpgOk returns a tuple with the PlannedInEpg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedInEpg

`func (o *BroadcastResult) SetPlannedInEpg(v int32)`

SetPlannedInEpg sets PlannedInEpg field to given value.

### HasPlannedInEpg

`func (o *BroadcastResult) HasPlannedInEpg() bool`

HasPlannedInEpg returns a boolean if a field has been set.

### GetGroupId

`func (o *BroadcastResult) GetGroupId() int64`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *BroadcastResult) GetGroupIdOk() (*int64, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *BroadcastResult) SetGroupId(v int64)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *BroadcastResult) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetGenre

`func (o *BroadcastResult) GetGenre() BroadcastRelationsGenre`

GetGenre returns the Genre field if non-nil, zero value otherwise.

### GetGenreOk

`func (o *BroadcastResult) GetGenreOk() (*BroadcastRelationsGenre, bool)`

GetGenreOk returns a tuple with the Genre field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenre

`func (o *BroadcastResult) SetGenre(v BroadcastRelationsGenre)`

SetGenre sets Genre field to given value.

### HasGenre

`func (o *BroadcastResult) HasGenre() bool`

HasGenre returns a boolean if a field has been set.

### GetItems

`func (o *BroadcastResult) GetItems() BroadcastRelationsItems`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *BroadcastResult) GetItemsOk() (*BroadcastRelationsItems, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *BroadcastResult) SetItems(v BroadcastRelationsItems)`

SetItems sets Items field to given value.

### HasItems

`func (o *BroadcastResult) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetBlocks

`func (o *BroadcastResult) GetBlocks() BroadcastRelationsBlocks`

GetBlocks returns the Blocks field if non-nil, zero value otherwise.

### GetBlocksOk

`func (o *BroadcastResult) GetBlocksOk() (*BroadcastRelationsBlocks, bool)`

GetBlocksOk returns a tuple with the Blocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocks

`func (o *BroadcastResult) SetBlocks(v BroadcastRelationsBlocks)`

SetBlocks sets Blocks field to given value.

### HasBlocks

`func (o *BroadcastResult) HasBlocks() bool`

HasBlocks returns a boolean if a field has been set.

### GetProgram

`func (o *BroadcastResult) GetProgram() BlockRelationsProgram`

GetProgram returns the Program field if non-nil, zero value otherwise.

### GetProgramOk

`func (o *BroadcastResult) GetProgramOk() (*BlockRelationsProgram, bool)`

GetProgramOk returns a tuple with the Program field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgram

`func (o *BroadcastResult) SetProgram(v BlockRelationsProgram)`

SetProgram sets Program field to given value.

### HasProgram

`func (o *BroadcastResult) HasProgram() bool`

HasProgram returns a boolean if a field has been set.

### GetTags

`func (o *BroadcastResult) GetTags() BroadcastRelationsTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BroadcastResult) GetTagsOk() (*BroadcastRelationsTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BroadcastResult) SetTags(v BroadcastRelationsTags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BroadcastResult) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetPresenters

`func (o *BroadcastResult) GetPresenters() BroadcastRelationsPresenters`

GetPresenters returns the Presenters field if non-nil, zero value otherwise.

### GetPresentersOk

`func (o *BroadcastResult) GetPresentersOk() (*BroadcastRelationsPresenters, bool)`

GetPresentersOk returns a tuple with the Presenters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresenters

`func (o *BroadcastResult) SetPresenters(v BroadcastRelationsPresenters)`

SetPresenters sets Presenters field to given value.

### HasPresenters

`func (o *BroadcastResult) HasPresenters() bool`

HasPresenters returns a boolean if a field has been set.

### GetModelType

`func (o *BroadcastResult) GetModelType() BroadcastRelationsModelType`

GetModelType returns the ModelType field if non-nil, zero value otherwise.

### GetModelTypeOk

`func (o *BroadcastResult) GetModelTypeOk() (*BroadcastRelationsModelType, bool)`

GetModelTypeOk returns a tuple with the ModelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelType

`func (o *BroadcastResult) SetModelType(v BroadcastRelationsModelType)`

SetModelType sets ModelType field to given value.

### HasModelType

`func (o *BroadcastResult) HasModelType() bool`

HasModelType returns a boolean if a field has been set.

### GetGroup

`func (o *BroadcastResult) GetGroup() BroadcastRelationsGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *BroadcastResult) GetGroupOk() (*BroadcastRelationsGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *BroadcastResult) SetGroup(v BroadcastRelationsGroup)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *BroadcastResult) HasGroup() bool`

HasGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


