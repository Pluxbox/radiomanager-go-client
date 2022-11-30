# GroupResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**UpdatedAt** | **time.Time** |  | 
**CreatedAt** | **time.Time** |  | 
**DeletedAt** | **time.Time** |  | 
**Title** | Pointer to **string** |  | [optional] 
**Users** | Pointer to [**GroupRelationsUsers**](GroupRelationsUsers.md) |  | [optional] 
**Broadcasts** | Pointer to [**GroupRelationsBroadcasts**](GroupRelationsBroadcasts.md) |  | [optional] 
**Programs** | Pointer to [**GroupRelationsPrograms**](GroupRelationsPrograms.md) |  | [optional] 

## Methods

### NewGroupResult

`func NewGroupResult(id int64, updatedAt time.Time, createdAt time.Time, deletedAt time.Time, ) *GroupResult`

NewGroupResult instantiates a new GroupResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupResultWithDefaults

`func NewGroupResultWithDefaults() *GroupResult`

NewGroupResultWithDefaults instantiates a new GroupResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GroupResult) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GroupResult) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GroupResult) SetId(v int64)`

SetId sets Id field to given value.


### GetUpdatedAt

`func (o *GroupResult) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GroupResult) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GroupResult) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetCreatedAt

`func (o *GroupResult) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GroupResult) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GroupResult) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDeletedAt

`func (o *GroupResult) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *GroupResult) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *GroupResult) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.


### GetTitle

`func (o *GroupResult) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *GroupResult) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *GroupResult) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *GroupResult) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUsers

`func (o *GroupResult) GetUsers() GroupRelationsUsers`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *GroupResult) GetUsersOk() (*GroupRelationsUsers, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *GroupResult) SetUsers(v GroupRelationsUsers)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *GroupResult) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetBroadcasts

`func (o *GroupResult) GetBroadcasts() GroupRelationsBroadcasts`

GetBroadcasts returns the Broadcasts field if non-nil, zero value otherwise.

### GetBroadcastsOk

`func (o *GroupResult) GetBroadcastsOk() (*GroupRelationsBroadcasts, bool)`

GetBroadcastsOk returns a tuple with the Broadcasts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBroadcasts

`func (o *GroupResult) SetBroadcasts(v GroupRelationsBroadcasts)`

SetBroadcasts sets Broadcasts field to given value.

### HasBroadcasts

`func (o *GroupResult) HasBroadcasts() bool`

HasBroadcasts returns a boolean if a field has been set.

### GetPrograms

`func (o *GroupResult) GetPrograms() GroupRelationsPrograms`

GetPrograms returns the Programs field if non-nil, zero value otherwise.

### GetProgramsOk

`func (o *GroupResult) GetProgramsOk() (*GroupRelationsPrograms, bool)`

GetProgramsOk returns a tuple with the Programs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrograms

`func (o *GroupResult) SetPrograms(v GroupRelationsPrograms)`

SetPrograms sets Programs field to given value.

### HasPrograms

`func (o *GroupResult) HasPrograms() bool`

HasPrograms returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


