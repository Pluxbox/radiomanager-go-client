# ProgramDataInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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
**Tags** | Pointer to **[]int32** |  | [optional] 
**Presenters** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewProgramDataInput

`func NewProgramDataInput(modelTypeId int64, title string, ) *ProgramDataInput`

NewProgramDataInput instantiates a new ProgramDataInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProgramDataInputWithDefaults

`func NewProgramDataInputWithDefaults() *ProgramDataInput`

NewProgramDataInputWithDefaults instantiates a new ProgramDataInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModelTypeId

`func (o *ProgramDataInput) GetModelTypeId() int64`

GetModelTypeId returns the ModelTypeId field if non-nil, zero value otherwise.

### GetModelTypeIdOk

`func (o *ProgramDataInput) GetModelTypeIdOk() (*int64, bool)`

GetModelTypeIdOk returns a tuple with the ModelTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelTypeId

`func (o *ProgramDataInput) SetModelTypeId(v int64)`

SetModelTypeId sets ModelTypeId field to given value.


### GetFieldValues

`func (o *ProgramDataInput) GetFieldValues() map[string]interface{}`

GetFieldValues returns the FieldValues field if non-nil, zero value otherwise.

### GetFieldValuesOk

`func (o *ProgramDataInput) GetFieldValuesOk() (*map[string]interface{}, bool)`

GetFieldValuesOk returns a tuple with the FieldValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldValues

`func (o *ProgramDataInput) SetFieldValues(v map[string]interface{})`

SetFieldValues sets FieldValues field to given value.

### HasFieldValues

`func (o *ProgramDataInput) HasFieldValues() bool`

HasFieldValues returns a boolean if a field has been set.

### GetTitle

`func (o *ProgramDataInput) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ProgramDataInput) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ProgramDataInput) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDisabled

`func (o *ProgramDataInput) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *ProgramDataInput) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *ProgramDataInput) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *ProgramDataInput) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetGenreId

`func (o *ProgramDataInput) GetGenreId() int64`

GetGenreId returns the GenreId field if non-nil, zero value otherwise.

### GetGenreIdOk

`func (o *ProgramDataInput) GetGenreIdOk() (*int64, bool)`

GetGenreIdOk returns a tuple with the GenreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenreId

`func (o *ProgramDataInput) SetGenreId(v int64)`

SetGenreId sets GenreId field to given value.

### HasGenreId

`func (o *ProgramDataInput) HasGenreId() bool`

HasGenreId returns a boolean if a field has been set.

### GetGroupId

`func (o *ProgramDataInput) GetGroupId() int64`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ProgramDataInput) GetGroupIdOk() (*int64, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ProgramDataInput) SetGroupId(v int64)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *ProgramDataInput) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetDescription

`func (o *ProgramDataInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProgramDataInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProgramDataInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProgramDataInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetShortName

`func (o *ProgramDataInput) GetShortName() string`

GetShortName returns the ShortName field if non-nil, zero value otherwise.

### GetShortNameOk

`func (o *ProgramDataInput) GetShortNameOk() (*string, bool)`

GetShortNameOk returns a tuple with the ShortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortName

`func (o *ProgramDataInput) SetShortName(v string)`

SetShortName sets ShortName field to given value.

### HasShortName

`func (o *ProgramDataInput) HasShortName() bool`

HasShortName returns a boolean if a field has been set.

### GetMediumName

`func (o *ProgramDataInput) GetMediumName() string`

GetMediumName returns the MediumName field if non-nil, zero value otherwise.

### GetMediumNameOk

`func (o *ProgramDataInput) GetMediumNameOk() (*string, bool)`

GetMediumNameOk returns a tuple with the MediumName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediumName

`func (o *ProgramDataInput) SetMediumName(v string)`

SetMediumName sets MediumName field to given value.

### HasMediumName

`func (o *ProgramDataInput) HasMediumName() bool`

HasMediumName returns a boolean if a field has been set.

### GetWebsite

`func (o *ProgramDataInput) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *ProgramDataInput) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *ProgramDataInput) SetWebsite(v string)`

SetWebsite sets Website field to given value.

### HasWebsite

`func (o *ProgramDataInput) HasWebsite() bool`

HasWebsite returns a boolean if a field has been set.

### GetEmail

`func (o *ProgramDataInput) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ProgramDataInput) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ProgramDataInput) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ProgramDataInput) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetRecommended

`func (o *ProgramDataInput) GetRecommended() bool`

GetRecommended returns the Recommended field if non-nil, zero value otherwise.

### GetRecommendedOk

`func (o *ProgramDataInput) GetRecommendedOk() (*bool, bool)`

GetRecommendedOk returns a tuple with the Recommended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommended

`func (o *ProgramDataInput) SetRecommended(v bool)`

SetRecommended sets Recommended field to given value.

### HasRecommended

`func (o *ProgramDataInput) HasRecommended() bool`

HasRecommended returns a boolean if a field has been set.

### GetLanguage

`func (o *ProgramDataInput) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *ProgramDataInput) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *ProgramDataInput) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *ProgramDataInput) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetPtyCodeId

`func (o *ProgramDataInput) GetPtyCodeId() int64`

GetPtyCodeId returns the PtyCodeId field if non-nil, zero value otherwise.

### GetPtyCodeIdOk

`func (o *ProgramDataInput) GetPtyCodeIdOk() (*int64, bool)`

GetPtyCodeIdOk returns a tuple with the PtyCodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPtyCodeId

`func (o *ProgramDataInput) SetPtyCodeId(v int64)`

SetPtyCodeId sets PtyCodeId field to given value.

### HasPtyCodeId

`func (o *ProgramDataInput) HasPtyCodeId() bool`

HasPtyCodeId returns a boolean if a field has been set.

### GetTags

`func (o *ProgramDataInput) GetTags() []int32`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ProgramDataInput) GetTagsOk() (*[]int32, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ProgramDataInput) SetTags(v []int32)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ProgramDataInput) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetPresenters

`func (o *ProgramDataInput) GetPresenters() []int32`

GetPresenters returns the Presenters field if non-nil, zero value otherwise.

### GetPresentersOk

`func (o *ProgramDataInput) GetPresentersOk() (*[]int32, bool)`

GetPresentersOk returns a tuple with the Presenters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresenters

`func (o *ProgramDataInput) SetPresenters(v []int32)`

SetPresenters sets Presenters field to given value.

### HasPresenters

`func (o *ProgramDataInput) HasPresenters() bool`

HasPresenters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


