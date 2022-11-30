# ImportItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ModelTypeId** | **int64** |  | 
**ExternalId** | **string** |  | 
**FieldValues** | Pointer to **map[string]interface{}** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Duration** | Pointer to **int64** |  | [optional] 
**Start** | Pointer to **time.Time** |  | [optional] 
**Recommended** | Pointer to **bool** |  | [optional] 
**StaticStart** | Pointer to **bool** |  | [optional] 
**Details** | Pointer to **string** |  | [optional] 
**Contacts** | Pointer to **[]int32** |  | [optional] 
**Tags** | Pointer to **[]int32** |  | [optional] 
**BroadcastId** | Pointer to **int32** |  | [optional] 

## Methods

### NewImportItem

`func NewImportItem(modelTypeId int64, externalId string, ) *ImportItem`

NewImportItem instantiates a new ImportItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportItemWithDefaults

`func NewImportItemWithDefaults() *ImportItem`

NewImportItemWithDefaults instantiates a new ImportItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModelTypeId

`func (o *ImportItem) GetModelTypeId() int64`

GetModelTypeId returns the ModelTypeId field if non-nil, zero value otherwise.

### GetModelTypeIdOk

`func (o *ImportItem) GetModelTypeIdOk() (*int64, bool)`

GetModelTypeIdOk returns a tuple with the ModelTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelTypeId

`func (o *ImportItem) SetModelTypeId(v int64)`

SetModelTypeId sets ModelTypeId field to given value.


### GetExternalId

`func (o *ImportItem) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ImportItem) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ImportItem) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetFieldValues

`func (o *ImportItem) GetFieldValues() map[string]interface{}`

GetFieldValues returns the FieldValues field if non-nil, zero value otherwise.

### GetFieldValuesOk

`func (o *ImportItem) GetFieldValuesOk() (*map[string]interface{}, bool)`

GetFieldValuesOk returns a tuple with the FieldValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldValues

`func (o *ImportItem) SetFieldValues(v map[string]interface{})`

SetFieldValues sets FieldValues field to given value.

### HasFieldValues

`func (o *ImportItem) HasFieldValues() bool`

HasFieldValues returns a boolean if a field has been set.

### GetTitle

`func (o *ImportItem) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ImportItem) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ImportItem) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ImportItem) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDuration

`func (o *ImportItem) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *ImportItem) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *ImportItem) SetDuration(v int64)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *ImportItem) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetStart

`func (o *ImportItem) GetStart() time.Time`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *ImportItem) GetStartOk() (*time.Time, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *ImportItem) SetStart(v time.Time)`

SetStart sets Start field to given value.

### HasStart

`func (o *ImportItem) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetRecommended

`func (o *ImportItem) GetRecommended() bool`

GetRecommended returns the Recommended field if non-nil, zero value otherwise.

### GetRecommendedOk

`func (o *ImportItem) GetRecommendedOk() (*bool, bool)`

GetRecommendedOk returns a tuple with the Recommended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommended

`func (o *ImportItem) SetRecommended(v bool)`

SetRecommended sets Recommended field to given value.

### HasRecommended

`func (o *ImportItem) HasRecommended() bool`

HasRecommended returns a boolean if a field has been set.

### GetStaticStart

`func (o *ImportItem) GetStaticStart() bool`

GetStaticStart returns the StaticStart field if non-nil, zero value otherwise.

### GetStaticStartOk

`func (o *ImportItem) GetStaticStartOk() (*bool, bool)`

GetStaticStartOk returns a tuple with the StaticStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticStart

`func (o *ImportItem) SetStaticStart(v bool)`

SetStaticStart sets StaticStart field to given value.

### HasStaticStart

`func (o *ImportItem) HasStaticStart() bool`

HasStaticStart returns a boolean if a field has been set.

### GetDetails

`func (o *ImportItem) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *ImportItem) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *ImportItem) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *ImportItem) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetContacts

`func (o *ImportItem) GetContacts() []int32`

GetContacts returns the Contacts field if non-nil, zero value otherwise.

### GetContactsOk

`func (o *ImportItem) GetContactsOk() (*[]int32, bool)`

GetContactsOk returns a tuple with the Contacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContacts

`func (o *ImportItem) SetContacts(v []int32)`

SetContacts sets Contacts field to given value.

### HasContacts

`func (o *ImportItem) HasContacts() bool`

HasContacts returns a boolean if a field has been set.

### GetTags

`func (o *ImportItem) GetTags() []int32`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ImportItem) GetTagsOk() (*[]int32, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ImportItem) SetTags(v []int32)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ImportItem) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetBroadcastId

`func (o *ImportItem) GetBroadcastId() int32`

GetBroadcastId returns the BroadcastId field if non-nil, zero value otherwise.

### GetBroadcastIdOk

`func (o *ImportItem) GetBroadcastIdOk() (*int32, bool)`

GetBroadcastIdOk returns a tuple with the BroadcastId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBroadcastId

`func (o *ImportItem) SetBroadcastId(v int32)`

SetBroadcastId sets BroadcastId field to given value.

### HasBroadcastId

`func (o *ImportItem) HasBroadcastId() bool`

HasBroadcastId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


