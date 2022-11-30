# BroadcastEPGResult

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
**Items** | Pointer to [**BroadcastRelationsItems**](BroadcastRelationsItems.md) |  | [optional] 
**Blocks** | Pointer to [**BroadcastRelationsBlocks**](BroadcastRelationsBlocks.md) |  | [optional] 
**Program** | Pointer to [**BlockRelationsProgram**](BlockRelationsProgram.md) |  | [optional] 
**Tags** | Pointer to [**BroadcastRelationsTags**](BroadcastRelationsTags.md) |  | [optional] 
**Presenters** | Pointer to [**[]PresenterEPGResult**](PresenterEPGResult.md) |  | [optional] 
**ModelType** | Pointer to [**BroadcastRelationsModelType**](BroadcastRelationsModelType.md) |  | [optional] 

## Methods

### NewBroadcastEPGResult

`func NewBroadcastEPGResult(id int64, ) *BroadcastEPGResult`

NewBroadcastEPGResult instantiates a new BroadcastEPGResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBroadcastEPGResultWithDefaults

`func NewBroadcastEPGResultWithDefaults() *BroadcastEPGResult`

NewBroadcastEPGResultWithDefaults instantiates a new BroadcastEPGResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BroadcastEPGResult) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BroadcastEPGResult) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BroadcastEPGResult) SetId(v int64)`

SetId sets Id field to given value.


### GetUpdatedAt

`func (o *BroadcastEPGResult) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BroadcastEPGResult) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BroadcastEPGResult) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *BroadcastEPGResult) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *BroadcastEPGResult) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BroadcastEPGResult) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BroadcastEPGResult) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BroadcastEPGResult) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *BroadcastEPGResult) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *BroadcastEPGResult) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *BroadcastEPGResult) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *BroadcastEPGResult) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetExternalStationId

`func (o *BroadcastEPGResult) GetExternalStationId() int64`

GetExternalStationId returns the ExternalStationId field if non-nil, zero value otherwise.

### GetExternalStationIdOk

`func (o *BroadcastEPGResult) GetExternalStationIdOk() (*int64, bool)`

GetExternalStationIdOk returns a tuple with the ExternalStationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalStationId

`func (o *BroadcastEPGResult) SetExternalStationId(v int64)`

SetExternalStationId sets ExternalStationId field to given value.

### HasExternalStationId

`func (o *BroadcastEPGResult) HasExternalStationId() bool`

HasExternalStationId returns a boolean if a field has been set.

### GetProgramId

`func (o *BroadcastEPGResult) GetProgramId() int64`

GetProgramId returns the ProgramId field if non-nil, zero value otherwise.

### GetProgramIdOk

`func (o *BroadcastEPGResult) GetProgramIdOk() (*int64, bool)`

GetProgramIdOk returns a tuple with the ProgramId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgramId

`func (o *BroadcastEPGResult) SetProgramId(v int64)`

SetProgramId sets ProgramId field to given value.

### HasProgramId

`func (o *BroadcastEPGResult) HasProgramId() bool`

HasProgramId returns a boolean if a field has been set.

### GetModelTypeId

`func (o *BroadcastEPGResult) GetModelTypeId() int64`

GetModelTypeId returns the ModelTypeId field if non-nil, zero value otherwise.

### GetModelTypeIdOk

`func (o *BroadcastEPGResult) GetModelTypeIdOk() (*int64, bool)`

GetModelTypeIdOk returns a tuple with the ModelTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelTypeId

`func (o *BroadcastEPGResult) SetModelTypeId(v int64)`

SetModelTypeId sets ModelTypeId field to given value.

### HasModelTypeId

`func (o *BroadcastEPGResult) HasModelTypeId() bool`

HasModelTypeId returns a boolean if a field has been set.

### GetStationId

`func (o *BroadcastEPGResult) GetStationId() int64`

GetStationId returns the StationId field if non-nil, zero value otherwise.

### GetStationIdOk

`func (o *BroadcastEPGResult) GetStationIdOk() (*int64, bool)`

GetStationIdOk returns a tuple with the StationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStationId

`func (o *BroadcastEPGResult) SetStationId(v int64)`

SetStationId sets StationId field to given value.

### HasStationId

`func (o *BroadcastEPGResult) HasStationId() bool`

HasStationId returns a boolean if a field has been set.

### GetFieldValues

`func (o *BroadcastEPGResult) GetFieldValues() map[string]interface{}`

GetFieldValues returns the FieldValues field if non-nil, zero value otherwise.

### GetFieldValuesOk

`func (o *BroadcastEPGResult) GetFieldValuesOk() (*map[string]interface{}, bool)`

GetFieldValuesOk returns a tuple with the FieldValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldValues

`func (o *BroadcastEPGResult) SetFieldValues(v map[string]interface{})`

SetFieldValues sets FieldValues field to given value.

### HasFieldValues

`func (o *BroadcastEPGResult) HasFieldValues() bool`

HasFieldValues returns a boolean if a field has been set.

### GetTitle

`func (o *BroadcastEPGResult) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *BroadcastEPGResult) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *BroadcastEPGResult) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *BroadcastEPGResult) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetStart

`func (o *BroadcastEPGResult) GetStart() time.Time`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *BroadcastEPGResult) GetStartOk() (*time.Time, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *BroadcastEPGResult) SetStart(v time.Time)`

SetStart sets Start field to given value.

### HasStart

`func (o *BroadcastEPGResult) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetStop

`func (o *BroadcastEPGResult) GetStop() time.Time`

GetStop returns the Stop field if non-nil, zero value otherwise.

### GetStopOk

`func (o *BroadcastEPGResult) GetStopOk() (*time.Time, bool)`

GetStopOk returns a tuple with the Stop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStop

`func (o *BroadcastEPGResult) SetStop(v time.Time)`

SetStop sets Stop field to given value.

### HasStop

`func (o *BroadcastEPGResult) HasStop() bool`

HasStop returns a boolean if a field has been set.

### GetGenreId

`func (o *BroadcastEPGResult) GetGenreId() int64`

GetGenreId returns the GenreId field if non-nil, zero value otherwise.

### GetGenreIdOk

`func (o *BroadcastEPGResult) GetGenreIdOk() (*int64, bool)`

GetGenreIdOk returns a tuple with the GenreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenreId

`func (o *BroadcastEPGResult) SetGenreId(v int64)`

SetGenreId sets GenreId field to given value.

### HasGenreId

`func (o *BroadcastEPGResult) HasGenreId() bool`

HasGenreId returns a boolean if a field has been set.

### GetDescription

`func (o *BroadcastEPGResult) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BroadcastEPGResult) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BroadcastEPGResult) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BroadcastEPGResult) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetShortName

`func (o *BroadcastEPGResult) GetShortName() string`

GetShortName returns the ShortName field if non-nil, zero value otherwise.

### GetShortNameOk

`func (o *BroadcastEPGResult) GetShortNameOk() (*string, bool)`

GetShortNameOk returns a tuple with the ShortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortName

`func (o *BroadcastEPGResult) SetShortName(v string)`

SetShortName sets ShortName field to given value.

### HasShortName

`func (o *BroadcastEPGResult) HasShortName() bool`

HasShortName returns a boolean if a field has been set.

### GetMediumName

`func (o *BroadcastEPGResult) GetMediumName() string`

GetMediumName returns the MediumName field if non-nil, zero value otherwise.

### GetMediumNameOk

`func (o *BroadcastEPGResult) GetMediumNameOk() (*string, bool)`

GetMediumNameOk returns a tuple with the MediumName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediumName

`func (o *BroadcastEPGResult) SetMediumName(v string)`

SetMediumName sets MediumName field to given value.

### HasMediumName

`func (o *BroadcastEPGResult) HasMediumName() bool`

HasMediumName returns a boolean if a field has been set.

### GetWebsite

`func (o *BroadcastEPGResult) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *BroadcastEPGResult) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *BroadcastEPGResult) SetWebsite(v string)`

SetWebsite sets Website field to given value.

### HasWebsite

`func (o *BroadcastEPGResult) HasWebsite() bool`

HasWebsite returns a boolean if a field has been set.

### GetEmail

`func (o *BroadcastEPGResult) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *BroadcastEPGResult) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *BroadcastEPGResult) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *BroadcastEPGResult) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetRecommended

`func (o *BroadcastEPGResult) GetRecommended() bool`

GetRecommended returns the Recommended field if non-nil, zero value otherwise.

### GetRecommendedOk

`func (o *BroadcastEPGResult) GetRecommendedOk() (*bool, bool)`

GetRecommendedOk returns a tuple with the Recommended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommended

`func (o *BroadcastEPGResult) SetRecommended(v bool)`

SetRecommended sets Recommended field to given value.

### HasRecommended

`func (o *BroadcastEPGResult) HasRecommended() bool`

HasRecommended returns a boolean if a field has been set.

### GetLanguage

`func (o *BroadcastEPGResult) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *BroadcastEPGResult) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *BroadcastEPGResult) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *BroadcastEPGResult) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetPublished

`func (o *BroadcastEPGResult) GetPublished() bool`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *BroadcastEPGResult) GetPublishedOk() (*bool, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *BroadcastEPGResult) SetPublished(v bool)`

SetPublished sets Published field to given value.

### HasPublished

`func (o *BroadcastEPGResult) HasPublished() bool`

HasPublished returns a boolean if a field has been set.

### GetRepetitionUid

`func (o *BroadcastEPGResult) GetRepetitionUid() string`

GetRepetitionUid returns the RepetitionUid field if non-nil, zero value otherwise.

### GetRepetitionUidOk

`func (o *BroadcastEPGResult) GetRepetitionUidOk() (*string, bool)`

GetRepetitionUidOk returns a tuple with the RepetitionUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepetitionUid

`func (o *BroadcastEPGResult) SetRepetitionUid(v string)`

SetRepetitionUid sets RepetitionUid field to given value.

### HasRepetitionUid

`func (o *BroadcastEPGResult) HasRepetitionUid() bool`

HasRepetitionUid returns a boolean if a field has been set.

### GetRepetitionType

`func (o *BroadcastEPGResult) GetRepetitionType() string`

GetRepetitionType returns the RepetitionType field if non-nil, zero value otherwise.

### GetRepetitionTypeOk

`func (o *BroadcastEPGResult) GetRepetitionTypeOk() (*string, bool)`

GetRepetitionTypeOk returns a tuple with the RepetitionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepetitionType

`func (o *BroadcastEPGResult) SetRepetitionType(v string)`

SetRepetitionType sets RepetitionType field to given value.

### HasRepetitionType

`func (o *BroadcastEPGResult) HasRepetitionType() bool`

HasRepetitionType returns a boolean if a field has been set.

### GetRepetitionEnd

`func (o *BroadcastEPGResult) GetRepetitionEnd() time.Time`

GetRepetitionEnd returns the RepetitionEnd field if non-nil, zero value otherwise.

### GetRepetitionEndOk

`func (o *BroadcastEPGResult) GetRepetitionEndOk() (*time.Time, bool)`

GetRepetitionEndOk returns a tuple with the RepetitionEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepetitionEnd

`func (o *BroadcastEPGResult) SetRepetitionEnd(v time.Time)`

SetRepetitionEnd sets RepetitionEnd field to given value.

### HasRepetitionEnd

`func (o *BroadcastEPGResult) HasRepetitionEnd() bool`

HasRepetitionEnd returns a boolean if a field has been set.

### GetRepetitionStart

`func (o *BroadcastEPGResult) GetRepetitionStart() time.Time`

GetRepetitionStart returns the RepetitionStart field if non-nil, zero value otherwise.

### GetRepetitionStartOk

`func (o *BroadcastEPGResult) GetRepetitionStartOk() (*time.Time, bool)`

GetRepetitionStartOk returns a tuple with the RepetitionStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepetitionStart

`func (o *BroadcastEPGResult) SetRepetitionStart(v time.Time)`

SetRepetitionStart sets RepetitionStart field to given value.

### HasRepetitionStart

`func (o *BroadcastEPGResult) HasRepetitionStart() bool`

HasRepetitionStart returns a boolean if a field has been set.

### GetRepetitionDays

`func (o *BroadcastEPGResult) GetRepetitionDays() string`

GetRepetitionDays returns the RepetitionDays field if non-nil, zero value otherwise.

### GetRepetitionDaysOk

`func (o *BroadcastEPGResult) GetRepetitionDaysOk() (*string, bool)`

GetRepetitionDaysOk returns a tuple with the RepetitionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepetitionDays

`func (o *BroadcastEPGResult) SetRepetitionDays(v string)`

SetRepetitionDays sets RepetitionDays field to given value.

### HasRepetitionDays

`func (o *BroadcastEPGResult) HasRepetitionDays() bool`

HasRepetitionDays returns a boolean if a field has been set.

### GetPtyCodeId

`func (o *BroadcastEPGResult) GetPtyCodeId() int64`

GetPtyCodeId returns the PtyCodeId field if non-nil, zero value otherwise.

### GetPtyCodeIdOk

`func (o *BroadcastEPGResult) GetPtyCodeIdOk() (*int64, bool)`

GetPtyCodeIdOk returns a tuple with the PtyCodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPtyCodeId

`func (o *BroadcastEPGResult) SetPtyCodeId(v int64)`

SetPtyCodeId sets PtyCodeId field to given value.

### HasPtyCodeId

`func (o *BroadcastEPGResult) HasPtyCodeId() bool`

HasPtyCodeId returns a boolean if a field has been set.

### GetPlannedInEpg

`func (o *BroadcastEPGResult) GetPlannedInEpg() int32`

GetPlannedInEpg returns the PlannedInEpg field if non-nil, zero value otherwise.

### GetPlannedInEpgOk

`func (o *BroadcastEPGResult) GetPlannedInEpgOk() (*int32, bool)`

GetPlannedInEpgOk returns a tuple with the PlannedInEpg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedInEpg

`func (o *BroadcastEPGResult) SetPlannedInEpg(v int32)`

SetPlannedInEpg sets PlannedInEpg field to given value.

### HasPlannedInEpg

`func (o *BroadcastEPGResult) HasPlannedInEpg() bool`

HasPlannedInEpg returns a boolean if a field has been set.

### GetGroupId

`func (o *BroadcastEPGResult) GetGroupId() int64`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *BroadcastEPGResult) GetGroupIdOk() (*int64, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *BroadcastEPGResult) SetGroupId(v int64)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *BroadcastEPGResult) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetItems

`func (o *BroadcastEPGResult) GetItems() BroadcastRelationsItems`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *BroadcastEPGResult) GetItemsOk() (*BroadcastRelationsItems, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *BroadcastEPGResult) SetItems(v BroadcastRelationsItems)`

SetItems sets Items field to given value.

### HasItems

`func (o *BroadcastEPGResult) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetBlocks

`func (o *BroadcastEPGResult) GetBlocks() BroadcastRelationsBlocks`

GetBlocks returns the Blocks field if non-nil, zero value otherwise.

### GetBlocksOk

`func (o *BroadcastEPGResult) GetBlocksOk() (*BroadcastRelationsBlocks, bool)`

GetBlocksOk returns a tuple with the Blocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocks

`func (o *BroadcastEPGResult) SetBlocks(v BroadcastRelationsBlocks)`

SetBlocks sets Blocks field to given value.

### HasBlocks

`func (o *BroadcastEPGResult) HasBlocks() bool`

HasBlocks returns a boolean if a field has been set.

### GetProgram

`func (o *BroadcastEPGResult) GetProgram() BlockRelationsProgram`

GetProgram returns the Program field if non-nil, zero value otherwise.

### GetProgramOk

`func (o *BroadcastEPGResult) GetProgramOk() (*BlockRelationsProgram, bool)`

GetProgramOk returns a tuple with the Program field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgram

`func (o *BroadcastEPGResult) SetProgram(v BlockRelationsProgram)`

SetProgram sets Program field to given value.

### HasProgram

`func (o *BroadcastEPGResult) HasProgram() bool`

HasProgram returns a boolean if a field has been set.

### GetTags

`func (o *BroadcastEPGResult) GetTags() BroadcastRelationsTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BroadcastEPGResult) GetTagsOk() (*BroadcastRelationsTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BroadcastEPGResult) SetTags(v BroadcastRelationsTags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BroadcastEPGResult) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetPresenters

`func (o *BroadcastEPGResult) GetPresenters() []PresenterEPGResult`

GetPresenters returns the Presenters field if non-nil, zero value otherwise.

### GetPresentersOk

`func (o *BroadcastEPGResult) GetPresentersOk() (*[]PresenterEPGResult, bool)`

GetPresentersOk returns a tuple with the Presenters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresenters

`func (o *BroadcastEPGResult) SetPresenters(v []PresenterEPGResult)`

SetPresenters sets Presenters field to given value.

### HasPresenters

`func (o *BroadcastEPGResult) HasPresenters() bool`

HasPresenters returns a boolean if a field has been set.

### GetModelType

`func (o *BroadcastEPGResult) GetModelType() BroadcastRelationsModelType`

GetModelType returns the ModelType field if non-nil, zero value otherwise.

### GetModelTypeOk

`func (o *BroadcastEPGResult) GetModelTypeOk() (*BroadcastRelationsModelType, bool)`

GetModelTypeOk returns a tuple with the ModelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelType

`func (o *BroadcastEPGResult) SetModelType(v BroadcastRelationsModelType)`

SetModelType sets ModelType field to given value.

### HasModelType

`func (o *BroadcastEPGResult) HasModelType() bool`

HasModelType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


