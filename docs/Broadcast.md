# Broadcast

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

## Methods

### NewBroadcast

`func NewBroadcast() *Broadcast`

NewBroadcast instantiates a new Broadcast object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBroadcastWithDefaults

`func NewBroadcastWithDefaults() *Broadcast`

NewBroadcastWithDefaults instantiates a new Broadcast object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProgramId

`func (o *Broadcast) GetProgramId() int64`

GetProgramId returns the ProgramId field if non-nil, zero value otherwise.

### GetProgramIdOk

`func (o *Broadcast) GetProgramIdOk() (*int64, bool)`

GetProgramIdOk returns a tuple with the ProgramId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgramId

`func (o *Broadcast) SetProgramId(v int64)`

SetProgramId sets ProgramId field to given value.

### HasProgramId

`func (o *Broadcast) HasProgramId() bool`

HasProgramId returns a boolean if a field has been set.

### GetModelTypeId

`func (o *Broadcast) GetModelTypeId() int64`

GetModelTypeId returns the ModelTypeId field if non-nil, zero value otherwise.

### GetModelTypeIdOk

`func (o *Broadcast) GetModelTypeIdOk() (*int64, bool)`

GetModelTypeIdOk returns a tuple with the ModelTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelTypeId

`func (o *Broadcast) SetModelTypeId(v int64)`

SetModelTypeId sets ModelTypeId field to given value.

### HasModelTypeId

`func (o *Broadcast) HasModelTypeId() bool`

HasModelTypeId returns a boolean if a field has been set.

### GetStationId

`func (o *Broadcast) GetStationId() int64`

GetStationId returns the StationId field if non-nil, zero value otherwise.

### GetStationIdOk

`func (o *Broadcast) GetStationIdOk() (*int64, bool)`

GetStationIdOk returns a tuple with the StationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStationId

`func (o *Broadcast) SetStationId(v int64)`

SetStationId sets StationId field to given value.

### HasStationId

`func (o *Broadcast) HasStationId() bool`

HasStationId returns a boolean if a field has been set.

### GetFieldValues

`func (o *Broadcast) GetFieldValues() map[string]interface{}`

GetFieldValues returns the FieldValues field if non-nil, zero value otherwise.

### GetFieldValuesOk

`func (o *Broadcast) GetFieldValuesOk() (*map[string]interface{}, bool)`

GetFieldValuesOk returns a tuple with the FieldValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldValues

`func (o *Broadcast) SetFieldValues(v map[string]interface{})`

SetFieldValues sets FieldValues field to given value.

### HasFieldValues

`func (o *Broadcast) HasFieldValues() bool`

HasFieldValues returns a boolean if a field has been set.

### GetTitle

`func (o *Broadcast) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Broadcast) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Broadcast) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Broadcast) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetStart

`func (o *Broadcast) GetStart() time.Time`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *Broadcast) GetStartOk() (*time.Time, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *Broadcast) SetStart(v time.Time)`

SetStart sets Start field to given value.

### HasStart

`func (o *Broadcast) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetStop

`func (o *Broadcast) GetStop() time.Time`

GetStop returns the Stop field if non-nil, zero value otherwise.

### GetStopOk

`func (o *Broadcast) GetStopOk() (*time.Time, bool)`

GetStopOk returns a tuple with the Stop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStop

`func (o *Broadcast) SetStop(v time.Time)`

SetStop sets Stop field to given value.

### HasStop

`func (o *Broadcast) HasStop() bool`

HasStop returns a boolean if a field has been set.

### GetGenreId

`func (o *Broadcast) GetGenreId() int64`

GetGenreId returns the GenreId field if non-nil, zero value otherwise.

### GetGenreIdOk

`func (o *Broadcast) GetGenreIdOk() (*int64, bool)`

GetGenreIdOk returns a tuple with the GenreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenreId

`func (o *Broadcast) SetGenreId(v int64)`

SetGenreId sets GenreId field to given value.

### HasGenreId

`func (o *Broadcast) HasGenreId() bool`

HasGenreId returns a boolean if a field has been set.

### GetDescription

`func (o *Broadcast) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Broadcast) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Broadcast) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Broadcast) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetShortName

`func (o *Broadcast) GetShortName() string`

GetShortName returns the ShortName field if non-nil, zero value otherwise.

### GetShortNameOk

`func (o *Broadcast) GetShortNameOk() (*string, bool)`

GetShortNameOk returns a tuple with the ShortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortName

`func (o *Broadcast) SetShortName(v string)`

SetShortName sets ShortName field to given value.

### HasShortName

`func (o *Broadcast) HasShortName() bool`

HasShortName returns a boolean if a field has been set.

### GetMediumName

`func (o *Broadcast) GetMediumName() string`

GetMediumName returns the MediumName field if non-nil, zero value otherwise.

### GetMediumNameOk

`func (o *Broadcast) GetMediumNameOk() (*string, bool)`

GetMediumNameOk returns a tuple with the MediumName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediumName

`func (o *Broadcast) SetMediumName(v string)`

SetMediumName sets MediumName field to given value.

### HasMediumName

`func (o *Broadcast) HasMediumName() bool`

HasMediumName returns a boolean if a field has been set.

### GetWebsite

`func (o *Broadcast) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *Broadcast) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *Broadcast) SetWebsite(v string)`

SetWebsite sets Website field to given value.

### HasWebsite

`func (o *Broadcast) HasWebsite() bool`

HasWebsite returns a boolean if a field has been set.

### GetEmail

`func (o *Broadcast) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Broadcast) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Broadcast) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Broadcast) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetRecommended

`func (o *Broadcast) GetRecommended() bool`

GetRecommended returns the Recommended field if non-nil, zero value otherwise.

### GetRecommendedOk

`func (o *Broadcast) GetRecommendedOk() (*bool, bool)`

GetRecommendedOk returns a tuple with the Recommended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommended

`func (o *Broadcast) SetRecommended(v bool)`

SetRecommended sets Recommended field to given value.

### HasRecommended

`func (o *Broadcast) HasRecommended() bool`

HasRecommended returns a boolean if a field has been set.

### GetLanguage

`func (o *Broadcast) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *Broadcast) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *Broadcast) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *Broadcast) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetPublished

`func (o *Broadcast) GetPublished() bool`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *Broadcast) GetPublishedOk() (*bool, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *Broadcast) SetPublished(v bool)`

SetPublished sets Published field to given value.

### HasPublished

`func (o *Broadcast) HasPublished() bool`

HasPublished returns a boolean if a field has been set.

### GetRepetitionUid

`func (o *Broadcast) GetRepetitionUid() string`

GetRepetitionUid returns the RepetitionUid field if non-nil, zero value otherwise.

### GetRepetitionUidOk

`func (o *Broadcast) GetRepetitionUidOk() (*string, bool)`

GetRepetitionUidOk returns a tuple with the RepetitionUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepetitionUid

`func (o *Broadcast) SetRepetitionUid(v string)`

SetRepetitionUid sets RepetitionUid field to given value.

### HasRepetitionUid

`func (o *Broadcast) HasRepetitionUid() bool`

HasRepetitionUid returns a boolean if a field has been set.

### GetRepetitionType

`func (o *Broadcast) GetRepetitionType() string`

GetRepetitionType returns the RepetitionType field if non-nil, zero value otherwise.

### GetRepetitionTypeOk

`func (o *Broadcast) GetRepetitionTypeOk() (*string, bool)`

GetRepetitionTypeOk returns a tuple with the RepetitionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepetitionType

`func (o *Broadcast) SetRepetitionType(v string)`

SetRepetitionType sets RepetitionType field to given value.

### HasRepetitionType

`func (o *Broadcast) HasRepetitionType() bool`

HasRepetitionType returns a boolean if a field has been set.

### GetRepetitionEnd

`func (o *Broadcast) GetRepetitionEnd() time.Time`

GetRepetitionEnd returns the RepetitionEnd field if non-nil, zero value otherwise.

### GetRepetitionEndOk

`func (o *Broadcast) GetRepetitionEndOk() (*time.Time, bool)`

GetRepetitionEndOk returns a tuple with the RepetitionEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepetitionEnd

`func (o *Broadcast) SetRepetitionEnd(v time.Time)`

SetRepetitionEnd sets RepetitionEnd field to given value.

### HasRepetitionEnd

`func (o *Broadcast) HasRepetitionEnd() bool`

HasRepetitionEnd returns a boolean if a field has been set.

### GetRepetitionStart

`func (o *Broadcast) GetRepetitionStart() time.Time`

GetRepetitionStart returns the RepetitionStart field if non-nil, zero value otherwise.

### GetRepetitionStartOk

`func (o *Broadcast) GetRepetitionStartOk() (*time.Time, bool)`

GetRepetitionStartOk returns a tuple with the RepetitionStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepetitionStart

`func (o *Broadcast) SetRepetitionStart(v time.Time)`

SetRepetitionStart sets RepetitionStart field to given value.

### HasRepetitionStart

`func (o *Broadcast) HasRepetitionStart() bool`

HasRepetitionStart returns a boolean if a field has been set.

### GetRepetitionDays

`func (o *Broadcast) GetRepetitionDays() string`

GetRepetitionDays returns the RepetitionDays field if non-nil, zero value otherwise.

### GetRepetitionDaysOk

`func (o *Broadcast) GetRepetitionDaysOk() (*string, bool)`

GetRepetitionDaysOk returns a tuple with the RepetitionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepetitionDays

`func (o *Broadcast) SetRepetitionDays(v string)`

SetRepetitionDays sets RepetitionDays field to given value.

### HasRepetitionDays

`func (o *Broadcast) HasRepetitionDays() bool`

HasRepetitionDays returns a boolean if a field has been set.

### GetPtyCodeId

`func (o *Broadcast) GetPtyCodeId() int64`

GetPtyCodeId returns the PtyCodeId field if non-nil, zero value otherwise.

### GetPtyCodeIdOk

`func (o *Broadcast) GetPtyCodeIdOk() (*int64, bool)`

GetPtyCodeIdOk returns a tuple with the PtyCodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPtyCodeId

`func (o *Broadcast) SetPtyCodeId(v int64)`

SetPtyCodeId sets PtyCodeId field to given value.

### HasPtyCodeId

`func (o *Broadcast) HasPtyCodeId() bool`

HasPtyCodeId returns a boolean if a field has been set.

### GetPlannedInEpg

`func (o *Broadcast) GetPlannedInEpg() int32`

GetPlannedInEpg returns the PlannedInEpg field if non-nil, zero value otherwise.

### GetPlannedInEpgOk

`func (o *Broadcast) GetPlannedInEpgOk() (*int32, bool)`

GetPlannedInEpgOk returns a tuple with the PlannedInEpg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedInEpg

`func (o *Broadcast) SetPlannedInEpg(v int32)`

SetPlannedInEpg sets PlannedInEpg field to given value.

### HasPlannedInEpg

`func (o *Broadcast) HasPlannedInEpg() bool`

HasPlannedInEpg returns a boolean if a field has been set.

### GetGroupId

`func (o *Broadcast) GetGroupId() int64`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *Broadcast) GetGroupIdOk() (*int64, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *Broadcast) SetGroupId(v int64)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *Broadcast) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


