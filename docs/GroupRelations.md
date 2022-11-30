# GroupRelations

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Users** | Pointer to [**GroupRelationsUsers**](GroupRelationsUsers.md) |  | [optional] 
**Broadcasts** | Pointer to [**GroupRelationsBroadcasts**](GroupRelationsBroadcasts.md) |  | [optional] 
**Programs** | Pointer to [**GroupRelationsPrograms**](GroupRelationsPrograms.md) |  | [optional] 

## Methods

### NewGroupRelations

`func NewGroupRelations() *GroupRelations`

NewGroupRelations instantiates a new GroupRelations object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupRelationsWithDefaults

`func NewGroupRelationsWithDefaults() *GroupRelations`

NewGroupRelationsWithDefaults instantiates a new GroupRelations object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsers

`func (o *GroupRelations) GetUsers() GroupRelationsUsers`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *GroupRelations) GetUsersOk() (*GroupRelationsUsers, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *GroupRelations) SetUsers(v GroupRelationsUsers)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *GroupRelations) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetBroadcasts

`func (o *GroupRelations) GetBroadcasts() GroupRelationsBroadcasts`

GetBroadcasts returns the Broadcasts field if non-nil, zero value otherwise.

### GetBroadcastsOk

`func (o *GroupRelations) GetBroadcastsOk() (*GroupRelationsBroadcasts, bool)`

GetBroadcastsOk returns a tuple with the Broadcasts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBroadcasts

`func (o *GroupRelations) SetBroadcasts(v GroupRelationsBroadcasts)`

SetBroadcasts sets Broadcasts field to given value.

### HasBroadcasts

`func (o *GroupRelations) HasBroadcasts() bool`

HasBroadcasts returns a boolean if a field has been set.

### GetPrograms

`func (o *GroupRelations) GetPrograms() GroupRelationsPrograms`

GetPrograms returns the Programs field if non-nil, zero value otherwise.

### GetProgramsOk

`func (o *GroupRelations) GetProgramsOk() (*GroupRelationsPrograms, bool)`

GetProgramsOk returns a tuple with the Programs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrograms

`func (o *GroupRelations) SetPrograms(v GroupRelationsPrograms)`

SetPrograms sets Programs field to given value.

### HasPrograms

`func (o *GroupRelations) HasPrograms() bool`

HasPrograms returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


