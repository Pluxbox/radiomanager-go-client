# BlockResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | ID of the current Block. | 
**BroadcastId** | **int64** | Currently assigned Broadcast connected to the current Block, identified by the Broadcast ID. | 
**Start** | **time.Time** | Start of the Block (formatted as a DateTime object), saved with an TimeZone. | 
**Stop** | **time.Time** | End of the Block (formatted as a DateTime object), saved with an TimeZone. | 
**CreatedAt** | **time.Time** | Time of the creation of the Block (formatted as a DateTime object), saved with an TimeZone. | 
**UpdatedAt** | **time.Time** | Time of the last update of the Block (formatted as a DateTime object), saved with an TimeZone. | 
**DeletedAt** | **time.Time** | Moment when the Block got deleted (formatted as a DateTime object), saved with an TimeZone. | 
**ExternalStationId** | Pointer to **int64** |  | [optional] 
**Items** | Pointer to [**BlockRelationsItems**](BlockRelationsItems.md) |  | [optional] 
**Broadcast** | Pointer to [**BlockRelationsBroadcast**](BlockRelationsBroadcast.md) |  | [optional] 
**Program** | Pointer to [**BlockRelationsProgram**](BlockRelationsProgram.md) |  | [optional] 

## Methods

### NewBlockResult

`func NewBlockResult(id int64, broadcastId int64, start time.Time, stop time.Time, createdAt time.Time, updatedAt time.Time, deletedAt time.Time, ) *BlockResult`

NewBlockResult instantiates a new BlockResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlockResultWithDefaults

`func NewBlockResultWithDefaults() *BlockResult`

NewBlockResultWithDefaults instantiates a new BlockResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BlockResult) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BlockResult) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BlockResult) SetId(v int64)`

SetId sets Id field to given value.


### GetBroadcastId

`func (o *BlockResult) GetBroadcastId() int64`

GetBroadcastId returns the BroadcastId field if non-nil, zero value otherwise.

### GetBroadcastIdOk

`func (o *BlockResult) GetBroadcastIdOk() (*int64, bool)`

GetBroadcastIdOk returns a tuple with the BroadcastId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBroadcastId

`func (o *BlockResult) SetBroadcastId(v int64)`

SetBroadcastId sets BroadcastId field to given value.


### GetStart

`func (o *BlockResult) GetStart() time.Time`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *BlockResult) GetStartOk() (*time.Time, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *BlockResult) SetStart(v time.Time)`

SetStart sets Start field to given value.


### GetStop

`func (o *BlockResult) GetStop() time.Time`

GetStop returns the Stop field if non-nil, zero value otherwise.

### GetStopOk

`func (o *BlockResult) GetStopOk() (*time.Time, bool)`

GetStopOk returns a tuple with the Stop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStop

`func (o *BlockResult) SetStop(v time.Time)`

SetStop sets Stop field to given value.


### GetCreatedAt

`func (o *BlockResult) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BlockResult) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BlockResult) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *BlockResult) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BlockResult) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BlockResult) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDeletedAt

`func (o *BlockResult) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *BlockResult) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *BlockResult) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.


### GetExternalStationId

`func (o *BlockResult) GetExternalStationId() int64`

GetExternalStationId returns the ExternalStationId field if non-nil, zero value otherwise.

### GetExternalStationIdOk

`func (o *BlockResult) GetExternalStationIdOk() (*int64, bool)`

GetExternalStationIdOk returns a tuple with the ExternalStationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalStationId

`func (o *BlockResult) SetExternalStationId(v int64)`

SetExternalStationId sets ExternalStationId field to given value.

### HasExternalStationId

`func (o *BlockResult) HasExternalStationId() bool`

HasExternalStationId returns a boolean if a field has been set.

### GetItems

`func (o *BlockResult) GetItems() BlockRelationsItems`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *BlockResult) GetItemsOk() (*BlockRelationsItems, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *BlockResult) SetItems(v BlockRelationsItems)`

SetItems sets Items field to given value.

### HasItems

`func (o *BlockResult) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetBroadcast

`func (o *BlockResult) GetBroadcast() BlockRelationsBroadcast`

GetBroadcast returns the Broadcast field if non-nil, zero value otherwise.

### GetBroadcastOk

`func (o *BlockResult) GetBroadcastOk() (*BlockRelationsBroadcast, bool)`

GetBroadcastOk returns a tuple with the Broadcast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBroadcast

`func (o *BlockResult) SetBroadcast(v BlockRelationsBroadcast)`

SetBroadcast sets Broadcast field to given value.

### HasBroadcast

`func (o *BlockResult) HasBroadcast() bool`

HasBroadcast returns a boolean if a field has been set.

### GetProgram

`func (o *BlockResult) GetProgram() BlockRelationsProgram`

GetProgram returns the Program field if non-nil, zero value otherwise.

### GetProgramOk

`func (o *BlockResult) GetProgramOk() (*BlockRelationsProgram, bool)`

GetProgramOk returns a tuple with the Program field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgram

`func (o *BlockResult) SetProgram(v BlockRelationsProgram)`

SetProgram sets Program field to given value.

### HasProgram

`func (o *BlockResult) HasProgram() bool`

HasProgram returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


