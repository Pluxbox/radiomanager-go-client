# TagResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**DeletedAt** | Pointer to **time.Time** |  | [optional] 
**ExternalStationId** | Pointer to **int64** |  | [optional] 
**Name** | **string** |  | 
**Broadcasts** | Pointer to [**TagRelationsBroadcasts**](TagRelationsBroadcasts.md) |  | [optional] 
**Programs** | Pointer to [**TagRelationsPrograms**](TagRelationsPrograms.md) |  | [optional] 
**Contacts** | Pointer to [**TagRelationsContacts**](TagRelationsContacts.md) |  | [optional] 
**Items** | Pointer to [**TagRelationsItems**](TagRelationsItems.md) |  | [optional] 

## Methods

### NewTagResult

`func NewTagResult(id int64, name string, ) *TagResult`

NewTagResult instantiates a new TagResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagResultWithDefaults

`func NewTagResultWithDefaults() *TagResult`

NewTagResultWithDefaults instantiates a new TagResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TagResult) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TagResult) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TagResult) SetId(v int64)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *TagResult) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TagResult) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TagResult) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TagResult) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *TagResult) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TagResult) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TagResult) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *TagResult) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *TagResult) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *TagResult) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *TagResult) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *TagResult) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetExternalStationId

`func (o *TagResult) GetExternalStationId() int64`

GetExternalStationId returns the ExternalStationId field if non-nil, zero value otherwise.

### GetExternalStationIdOk

`func (o *TagResult) GetExternalStationIdOk() (*int64, bool)`

GetExternalStationIdOk returns a tuple with the ExternalStationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalStationId

`func (o *TagResult) SetExternalStationId(v int64)`

SetExternalStationId sets ExternalStationId field to given value.

### HasExternalStationId

`func (o *TagResult) HasExternalStationId() bool`

HasExternalStationId returns a boolean if a field has been set.

### GetName

`func (o *TagResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TagResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TagResult) SetName(v string)`

SetName sets Name field to given value.


### GetBroadcasts

`func (o *TagResult) GetBroadcasts() TagRelationsBroadcasts`

GetBroadcasts returns the Broadcasts field if non-nil, zero value otherwise.

### GetBroadcastsOk

`func (o *TagResult) GetBroadcastsOk() (*TagRelationsBroadcasts, bool)`

GetBroadcastsOk returns a tuple with the Broadcasts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBroadcasts

`func (o *TagResult) SetBroadcasts(v TagRelationsBroadcasts)`

SetBroadcasts sets Broadcasts field to given value.

### HasBroadcasts

`func (o *TagResult) HasBroadcasts() bool`

HasBroadcasts returns a boolean if a field has been set.

### GetPrograms

`func (o *TagResult) GetPrograms() TagRelationsPrograms`

GetPrograms returns the Programs field if non-nil, zero value otherwise.

### GetProgramsOk

`func (o *TagResult) GetProgramsOk() (*TagRelationsPrograms, bool)`

GetProgramsOk returns a tuple with the Programs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrograms

`func (o *TagResult) SetPrograms(v TagRelationsPrograms)`

SetPrograms sets Programs field to given value.

### HasPrograms

`func (o *TagResult) HasPrograms() bool`

HasPrograms returns a boolean if a field has been set.

### GetContacts

`func (o *TagResult) GetContacts() TagRelationsContacts`

GetContacts returns the Contacts field if non-nil, zero value otherwise.

### GetContactsOk

`func (o *TagResult) GetContactsOk() (*TagRelationsContacts, bool)`

GetContactsOk returns a tuple with the Contacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContacts

`func (o *TagResult) SetContacts(v TagRelationsContacts)`

SetContacts sets Contacts field to given value.

### HasContacts

`func (o *TagResult) HasContacts() bool`

HasContacts returns a boolean if a field has been set.

### GetItems

`func (o *TagResult) GetItems() TagRelationsItems`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *TagResult) GetItemsOk() (*TagRelationsItems, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *TagResult) SetItems(v TagRelationsItems)`

SetItems sets Items field to given value.

### HasItems

`func (o *TagResult) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


