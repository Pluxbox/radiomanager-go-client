# BlockRelations

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**BlockRelationsItems**](BlockRelationsItems.md) |  | [optional] 
**Broadcast** | Pointer to [**BlockRelationsBroadcast**](BlockRelationsBroadcast.md) |  | [optional] 
**Program** | Pointer to [**BlockRelationsProgram**](BlockRelationsProgram.md) |  | [optional] 

## Methods

### NewBlockRelations

`func NewBlockRelations() *BlockRelations`

NewBlockRelations instantiates a new BlockRelations object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlockRelationsWithDefaults

`func NewBlockRelationsWithDefaults() *BlockRelations`

NewBlockRelationsWithDefaults instantiates a new BlockRelations object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *BlockRelations) GetItems() BlockRelationsItems`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *BlockRelations) GetItemsOk() (*BlockRelationsItems, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *BlockRelations) SetItems(v BlockRelationsItems)`

SetItems sets Items field to given value.

### HasItems

`func (o *BlockRelations) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetBroadcast

`func (o *BlockRelations) GetBroadcast() BlockRelationsBroadcast`

GetBroadcast returns the Broadcast field if non-nil, zero value otherwise.

### GetBroadcastOk

`func (o *BlockRelations) GetBroadcastOk() (*BlockRelationsBroadcast, bool)`

GetBroadcastOk returns a tuple with the Broadcast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBroadcast

`func (o *BlockRelations) SetBroadcast(v BlockRelationsBroadcast)`

SetBroadcast sets Broadcast field to given value.

### HasBroadcast

`func (o *BlockRelations) HasBroadcast() bool`

HasBroadcast returns a boolean if a field has been set.

### GetProgram

`func (o *BlockRelations) GetProgram() BlockRelationsProgram`

GetProgram returns the Program field if non-nil, zero value otherwise.

### GetProgramOk

`func (o *BlockRelations) GetProgramOk() (*BlockRelationsProgram, bool)`

GetProgramOk returns a tuple with the Program field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgram

`func (o *BlockRelations) SetProgram(v BlockRelationsProgram)`

SetProgram sets Program field to given value.

### HasProgram

`func (o *BlockRelations) HasProgram() bool`

HasProgram returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


