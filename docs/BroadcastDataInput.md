# BroadcastDataInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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
**Tags** | Pointer to **[]int32** |  | [optional] 
**Presenters** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewBroadcastDataInput

`func NewBroadcastDataInput() *BroadcastDataInput`

NewBroadcastDataInput instantiates a new BroadcastDataInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBroadcastDataInputWithDefaults

`func NewBroadcastDataInputWithDefaults() *BroadcastDataInput`

NewBroadcastDataInputWithDefaults instantiates a new BroadcastDataInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProgramId

`func (o *BroadcastDataInput) GetProgramId() int64`

GetProgramId returns the ProgramId field if non-nil, zero value otherwise.

### GetProgramIdOk

`func (o *BroadcastDataInput) GetProgramIdOk() (*int64, bool)`

GetProgramIdOk returns a tuple with the ProgramId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgramId

`func (o *BroadcastDataInput) SetProgramId(v int64)`

SetProgramId sets ProgramId field to given value.

### HasProgramId

`func (o *BroadcastDataInput) HasProgramId() bool`

HasProgramId returns a boolean if a field has been set.

### GetModelTypeId

`func (o *BroadcastDataInput) GetModelTypeId() int64`

GetModelTypeId returns the ModelTypeId field if non-nil, zero value otherwise.

### GetModelTypeIdOk

`func (o *BroadcastDataInput) GetModelTypeIdOk() (*int64, bool)`

GetModelTypeIdOk returns a tuple with the ModelTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelTypeId

`func (o *BroadcastDataInput) SetModelTypeId(v int64)`

SetModelTypeId sets ModelTypeId field to given value.

### HasModelTypeId

`func (o *BroadcastDataInput) HasModelTypeId() bool`

HasModelTypeId returns a boolean if a field has been set.

### GetStationId

`func (o *BroadcastDataInput) GetStationId() int64`

GetStationId returns the StationId field if non-nil, zero value otherwise.

### GetStationIdOk

`func (o *BroadcastDataInput) GetStationIdOk() (*int64, bool)`

GetStationIdOk returns a tuple with the StationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStationId

`func (o *BroadcastDataInput) SetStationId(v int64)`

SetStationId sets StationId field to given value.

### HasStationId

`func (o *BroadcastDataInput) HasStationId() bool`

HasStationId returns a boolean if a field has been set.

### GetFieldValues

`func (o *BroadcastDataInput) GetFieldValues() map[string]interface{}`

GetFieldValues returns the FieldValues field if non-nil, zero value otherwise.

### GetFieldValuesOk

`func (o *BroadcastDataInput) GetFieldValuesOk() (*map[string]interface{}, bool)`

GetFieldValuesOk returns a tuple with the FieldValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldValues

`func (o *BroadcastDataInput) SetFieldValues(v map[string]interface{})`

SetFieldValues sets FieldValues field to given value.

### HasFieldValues

`func (o *BroadcastDataInput) HasFieldValues() bool`

HasFieldValues returns a boolean if a field has been set.

### GetTitle

`func (o *BroadcastDataInput) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *BroadcastDataInput) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *BroadcastDataInput) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *BroadcastDataInput) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetStart

`func (o *BroadcastDataInput) GetStart() time.Time`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *BroadcastDataInput) GetStartOk() (*time.Time, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *BroadcastDataInput) SetStart(v time.Time)`

SetStart sets Start field to given value.

### HasStart

`func (o *BroadcastDataInput) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetStop

`func (o *BroadcastDataInput) GetStop() time.Time`

GetStop returns the Stop field if non-nil, zero value otherwise.

### GetStopOk

`func (o *BroadcastDataInput) GetStopOk() (*time.Time, bool)`

GetStopOk returns a tuple with the Stop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStop

`func (o *BroadcastDataInput) SetStop(v time.Time)`

SetStop sets Stop field to given value.

### HasStop

`func (o *BroadcastDataInput) HasStop() bool`

HasStop returns a boolean if a field has been set.

### GetGenreId

`func (o *BroadcastDataInput) GetGenreId() int64`

GetGenreId returns the GenreId field if non-nil, zero value otherwise.

### GetGenreIdOk

`func (o *BroadcastDataInput) GetGenreIdOk() (*int64, bool)`

GetGenreIdOk returns a tuple with the GenreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenreId

`func (o *BroadcastDataInput) SetGenreId(v int64)`

SetGenreId sets GenreId field to given value.

### HasGenreId

`func (o *BroadcastDataInput) HasGenreId() bool`

HasGenreId returns a boolean if a field has been set.

### GetDescription

`func (o *BroadcastDataInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BroadcastDataInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BroadcastDataInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BroadcastDataInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetShortName

`func (o *BroadcastDataInput) GetShortName() string`

GetShortName returns the ShortName field if non-nil, zero value otherwise.

### GetShortNameOk

`func (o *BroadcastDataInput) GetShortNameOk() (*string, bool)`

GetShortNameOk returns a tuple with the ShortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortName

`func (o *BroadcastDataInput) SetShortName(v string)`

SetShortName sets ShortName field to given value.

### HasShortName

`func (o *BroadcastDataInput) HasShortName() bool`

HasShortName returns a boolean if a field has been set.

### GetMediumName

`func (o *BroadcastDataInput) GetMediumName() string`

GetMediumName returns the MediumName field if non-nil, zero value otherwise.

### GetMediumNameOk

`func (o *BroadcastDataInput) GetMediumNameOk() (*string, bool)`

GetMediumNameOk returns a tuple with the MediumName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediumName

`func (o *BroadcastDataInput) SetMediumName(v string)`

SetMediumName sets MediumName field to given value.

### HasMediumName

`func (o *BroadcastDataInput) HasMediumName() bool`

HasMediumName returns a boolean if a field has been set.

### GetWebsite

`func (o *BroadcastDataInput) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *BroadcastDataInput) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *BroadcastDataInput) SetWebsite(v string)`

SetWebsite sets Website field to given value.

### HasWebsite

`func (o *BroadcastDataInput) HasWebsite() bool`

HasWebsite returns a boolean if a field has been set.

### GetEmail

`func (o *BroadcastDataInput) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *BroadcastDataInput) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *BroadcastDataInput) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *BroadcastDataInput) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetRecommended

`func (o *BroadcastDataInput) GetRecommended() bool`

GetRecommended returns the Recommended field if non-nil, zero value otherwise.

### GetRecommendedOk

`func (o *BroadcastDataInput) GetRecommendedOk() (*bool, bool)`

GetRecommendedOk returns a tuple with the Recommended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommended

`func (o *BroadcastDataInput) SetRecommended(v bool)`

SetRecommended sets Recommended field to given value.

### HasRecommended

`func (o *BroadcastDataInput) HasRecommended() bool`

HasRecommended returns a boolean if a field has been set.

### GetLanguage

`func (o *BroadcastDataInput) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *BroadcastDataInput) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *BroadcastDataInput) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *BroadcastDataInput) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetPublished

`func (o *BroadcastDataInput) GetPublished() bool`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *BroadcastDataInput) GetPublishedOk() (*bool, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *BroadcastDataInput) SetPublished(v bool)`

SetPublished sets Published field to given value.

### HasPublished

`func (o *BroadcastDataInput) HasPublished() bool`

HasPublished returns a boolean if a field has been set.

### GetRepetitionUid

`func (o *BroadcastDataInput) GetRepetitionUid() string`

GetRepetitionUid returns the RepetitionUid field if non-nil, zero value otherwise.

### GetRepetitionUidOk

`func (o *BroadcastDataInput) GetRepetitionUidOk() (*string, bool)`

GetRepetitionUidOk returns a tuple with the RepetitionUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepetitionUid

`func (o *BroadcastDataInput) SetRepetitionUid(v string)`

SetRepetitionUid sets RepetitionUid field to given value.

### HasRepetitionUid

`func (o *BroadcastDataInput) HasRepetitionUid() bool`

HasRepetitionUid returns a boolean if a field has been set.

### GetRepetitionType

`func (o *BroadcastDataInput) GetRepetitionType() string`

GetRepetitionType returns the RepetitionType field if non-nil, zero value otherwise.

### GetRepetitionTypeOk

`func (o *BroadcastDataInput) GetRepetitionTypeOk() (*string, bool)`

GetRepetitionTypeOk returns a tuple with the RepetitionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepetitionType

`func (o *BroadcastDataInput) SetRepetitionType(v string)`

SetRepetitionType sets RepetitionType field to given value.

### HasRepetitionType

`func (o *BroadcastDataInput) HasRepetitionType() bool`

HasRepetitionType returns a boolean if a field has been set.

### GetRepetitionEnd

`func (o *BroadcastDataInput) GetRepetitionEnd() time.Time`

GetRepetitionEnd returns the RepetitionEnd field if non-nil, zero value otherwise.

### GetRepetitionEndOk

`func (o *BroadcastDataInput) GetRepetitionEndOk() (*time.Time, bool)`

GetRepetitionEndOk returns a tuple with the RepetitionEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepetitionEnd

`func (o *BroadcastDataInput) SetRepetitionEnd(v time.Time)`

SetRepetitionEnd sets RepetitionEnd field to given value.

### HasRepetitionEnd

`func (o *BroadcastDataInput) HasRepetitionEnd() bool`

HasRepetitionEnd returns a boolean if a field has been set.

### GetRepetitionStart

`func (o *BroadcastDataInput) GetRepetitionStart() time.Time`

GetRepetitionStart returns the RepetitionStart field if non-nil, zero value otherwise.

### GetRepetitionStartOk

`func (o *BroadcastDataInput) GetRepetitionStartOk() (*time.Time, bool)`

GetRepetitionStartOk returns a tuple with the RepetitionStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepetitionStart

`func (o *BroadcastDataInput) SetRepetitionStart(v time.Time)`

SetRepetitionStart sets RepetitionStart field to given value.

### HasRepetitionStart

`func (o *BroadcastDataInput) HasRepetitionStart() bool`

HasRepetitionStart returns a boolean if a field has been set.

### GetRepetitionDays

`func (o *BroadcastDataInput) GetRepetitionDays() string`

GetRepetitionDays returns the RepetitionDays field if non-nil, zero value otherwise.

### GetRepetitionDaysOk

`func (o *BroadcastDataInput) GetRepetitionDaysOk() (*string, bool)`

GetRepetitionDaysOk returns a tuple with the RepetitionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepetitionDays

`func (o *BroadcastDataInput) SetRepetitionDays(v string)`

SetRepetitionDays sets RepetitionDays field to given value.

### HasRepetitionDays

`func (o *BroadcastDataInput) HasRepetitionDays() bool`

HasRepetitionDays returns a boolean if a field has been set.

### GetPtyCodeId

`func (o *BroadcastDataInput) GetPtyCodeId() int64`

GetPtyCodeId returns the PtyCodeId field if non-nil, zero value otherwise.

### GetPtyCodeIdOk

`func (o *BroadcastDataInput) GetPtyCodeIdOk() (*int64, bool)`

GetPtyCodeIdOk returns a tuple with the PtyCodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPtyCodeId

`func (o *BroadcastDataInput) SetPtyCodeId(v int64)`

SetPtyCodeId sets PtyCodeId field to given value.

### HasPtyCodeId

`func (o *BroadcastDataInput) HasPtyCodeId() bool`

HasPtyCodeId returns a boolean if a field has been set.

### GetPlannedInEpg

`func (o *BroadcastDataInput) GetPlannedInEpg() int32`

GetPlannedInEpg returns the PlannedInEpg field if non-nil, zero value otherwise.

### GetPlannedInEpgOk

`func (o *BroadcastDataInput) GetPlannedInEpgOk() (*int32, bool)`

GetPlannedInEpgOk returns a tuple with the PlannedInEpg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedInEpg

`func (o *BroadcastDataInput) SetPlannedInEpg(v int32)`

SetPlannedInEpg sets PlannedInEpg field to given value.

### HasPlannedInEpg

`func (o *BroadcastDataInput) HasPlannedInEpg() bool`

HasPlannedInEpg returns a boolean if a field has been set.

### GetGroupId

`func (o *BroadcastDataInput) GetGroupId() int64`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *BroadcastDataInput) GetGroupIdOk() (*int64, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *BroadcastDataInput) SetGroupId(v int64)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *BroadcastDataInput) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetTags

`func (o *BroadcastDataInput) GetTags() []int32`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BroadcastDataInput) GetTagsOk() (*[]int32, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BroadcastDataInput) SetTags(v []int32)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BroadcastDataInput) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetPresenters

`func (o *BroadcastDataInput) GetPresenters() []int32`

GetPresenters returns the Presenters field if non-nil, zero value otherwise.

### GetPresentersOk

`func (o *BroadcastDataInput) GetPresentersOk() (*[]int32, bool)`

GetPresentersOk returns a tuple with the Presenters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresenters

`func (o *BroadcastDataInput) SetPresenters(v []int32)`

SetPresenters sets Presenters field to given value.

### HasPresenters

`func (o *BroadcastDataInput) HasPresenters() bool`

HasPresenters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


