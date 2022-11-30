# ImportItemAllOf

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

### NewImportItemAllOf

`func NewImportItemAllOf(modelTypeId int64, externalId string, ) *ImportItemAllOf`

NewImportItemAllOf instantiates a new ImportItemAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportItemAllOfWithDefaults

`func NewImportItemAllOfWithDefaults() *ImportItemAllOf`

NewImportItemAllOfWithDefaults instantiates a new ImportItemAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModelTypeId

`func (o *ImportItemAllOf) GetModelTypeId() int64`

GetModelTypeId returns the ModelTypeId field if non-nil, zero value otherwise.

### GetModelTypeIdOk

`func (o *ImportItemAllOf) GetModelTypeIdOk() (*int64, bool)`

GetModelTypeIdOk returns a tuple with the ModelTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelTypeId

`func (o *ImportItemAllOf) SetModelTypeId(v int64)`

SetModelTypeId sets ModelTypeId field to given value.


### GetExternalId

`func (o *ImportItemAllOf) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ImportItemAllOf) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ImportItemAllOf) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetFieldValues

`func (o *ImportItemAllOf) GetFieldValues() map[string]interface{}`

GetFieldValues returns the FieldValues field if non-nil, zero value otherwise.

### GetFieldValuesOk

`func (o *ImportItemAllOf) GetFieldValuesOk() (*map[string]interface{}, bool)`

GetFieldValuesOk returns a tuple with the FieldValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldValues

`func (o *ImportItemAllOf) SetFieldValues(v map[string]interface{})`

SetFieldValues sets FieldValues field to given value.

### HasFieldValues

`func (o *ImportItemAllOf) HasFieldValues() bool`

HasFieldValues returns a boolean if a field has been set.

### GetTitle

`func (o *ImportItemAllOf) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ImportItemAllOf) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ImportItemAllOf) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ImportItemAllOf) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDuration

`func (o *ImportItemAllOf) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *ImportItemAllOf) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *ImportItemAllOf) SetDuration(v int64)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *ImportItemAllOf) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetStart

`func (o *ImportItemAllOf) GetStart() time.Time`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *ImportItemAllOf) GetStartOk() (*time.Time, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *ImportItemAllOf) SetStart(v time.Time)`

SetStart sets Start field to given value.

### HasStart

`func (o *ImportItemAllOf) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetRecommended

`func (o *ImportItemAllOf) GetRecommended() bool`

GetRecommended returns the Recommended field if non-nil, zero value otherwise.

### GetRecommendedOk

`func (o *ImportItemAllOf) GetRecommendedOk() (*bool, bool)`

GetRecommendedOk returns a tuple with the Recommended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommended

`func (o *ImportItemAllOf) SetRecommended(v bool)`

SetRecommended sets Recommended field to given value.

### HasRecommended

`func (o *ImportItemAllOf) HasRecommended() bool`

HasRecommended returns a boolean if a field has been set.

### GetStaticStart

`func (o *ImportItemAllOf) GetStaticStart() bool`

GetStaticStart returns the StaticStart field if non-nil, zero value otherwise.

### GetStaticStartOk

`func (o *ImportItemAllOf) GetStaticStartOk() (*bool, bool)`

GetStaticStartOk returns a tuple with the StaticStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticStart

`func (o *ImportItemAllOf) SetStaticStart(v bool)`

SetStaticStart sets StaticStart field to given value.

### HasStaticStart

`func (o *ImportItemAllOf) HasStaticStart() bool`

HasStaticStart returns a boolean if a field has been set.

### GetDetails

`func (o *ImportItemAllOf) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *ImportItemAllOf) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *ImportItemAllOf) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *ImportItemAllOf) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetContacts

`func (o *ImportItemAllOf) GetContacts() []int32`

GetContacts returns the Contacts field if non-nil, zero value otherwise.

### GetContactsOk

`func (o *ImportItemAllOf) GetContactsOk() (*[]int32, bool)`

GetContactsOk returns a tuple with the Contacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContacts

`func (o *ImportItemAllOf) SetContacts(v []int32)`

SetContacts sets Contacts field to given value.

### HasContacts

`func (o *ImportItemAllOf) HasContacts() bool`

HasContacts returns a boolean if a field has been set.

### GetTags

`func (o *ImportItemAllOf) GetTags() []int32`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ImportItemAllOf) GetTagsOk() (*[]int32, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ImportItemAllOf) SetTags(v []int32)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ImportItemAllOf) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetBroadcastId

`func (o *ImportItemAllOf) GetBroadcastId() int32`

GetBroadcastId returns the BroadcastId field if non-nil, zero value otherwise.

### GetBroadcastIdOk

`func (o *ImportItemAllOf) GetBroadcastIdOk() (*int32, bool)`

GetBroadcastIdOk returns a tuple with the BroadcastId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBroadcastId

`func (o *ImportItemAllOf) SetBroadcastId(v int32)`

SetBroadcastId sets BroadcastId field to given value.

### HasBroadcastId

`func (o *ImportItemAllOf) HasBroadcastId() bool`

HasBroadcastId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


