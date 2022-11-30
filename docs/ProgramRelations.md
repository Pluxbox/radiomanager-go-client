# ProgramRelations

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Genre** | Pointer to [**BroadcastRelationsGenre**](BroadcastRelationsGenre.md) |  | [optional] 
**Items** | Pointer to [**ProgramRelationsItems**](ProgramRelationsItems.md) |  | [optional] 
**Blocks** | Pointer to [**ProgramRelationsBlocks**](ProgramRelationsBlocks.md) |  | [optional] 
**Broadcasts** | Pointer to [**ProgramRelationsBroadcasts**](ProgramRelationsBroadcasts.md) |  | [optional] 
**Presenters** | Pointer to [**ProgramRelationsPresenters**](ProgramRelationsPresenters.md) |  | [optional] 
**Tags** | Pointer to [**ProgramRelationsTags**](ProgramRelationsTags.md) |  | [optional] 
**ModelType** | Pointer to [**BroadcastRelationsModelType**](BroadcastRelationsModelType.md) |  | [optional] 
**Group** | Pointer to [**BroadcastRelationsGroup**](BroadcastRelationsGroup.md) |  | [optional] 

## Methods

### NewProgramRelations

`func NewProgramRelations() *ProgramRelations`

NewProgramRelations instantiates a new ProgramRelations object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProgramRelationsWithDefaults

`func NewProgramRelationsWithDefaults() *ProgramRelations`

NewProgramRelationsWithDefaults instantiates a new ProgramRelations object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGenre

`func (o *ProgramRelations) GetGenre() BroadcastRelationsGenre`

GetGenre returns the Genre field if non-nil, zero value otherwise.

### GetGenreOk

`func (o *ProgramRelations) GetGenreOk() (*BroadcastRelationsGenre, bool)`

GetGenreOk returns a tuple with the Genre field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenre

`func (o *ProgramRelations) SetGenre(v BroadcastRelationsGenre)`

SetGenre sets Genre field to given value.

### HasGenre

`func (o *ProgramRelations) HasGenre() bool`

HasGenre returns a boolean if a field has been set.

### GetItems

`func (o *ProgramRelations) GetItems() ProgramRelationsItems`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ProgramRelations) GetItemsOk() (*ProgramRelationsItems, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ProgramRelations) SetItems(v ProgramRelationsItems)`

SetItems sets Items field to given value.

### HasItems

`func (o *ProgramRelations) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetBlocks

`func (o *ProgramRelations) GetBlocks() ProgramRelationsBlocks`

GetBlocks returns the Blocks field if non-nil, zero value otherwise.

### GetBlocksOk

`func (o *ProgramRelations) GetBlocksOk() (*ProgramRelationsBlocks, bool)`

GetBlocksOk returns a tuple with the Blocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocks

`func (o *ProgramRelations) SetBlocks(v ProgramRelationsBlocks)`

SetBlocks sets Blocks field to given value.

### HasBlocks

`func (o *ProgramRelations) HasBlocks() bool`

HasBlocks returns a boolean if a field has been set.

### GetBroadcasts

`func (o *ProgramRelations) GetBroadcasts() ProgramRelationsBroadcasts`

GetBroadcasts returns the Broadcasts field if non-nil, zero value otherwise.

### GetBroadcastsOk

`func (o *ProgramRelations) GetBroadcastsOk() (*ProgramRelationsBroadcasts, bool)`

GetBroadcastsOk returns a tuple with the Broadcasts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBroadcasts

`func (o *ProgramRelations) SetBroadcasts(v ProgramRelationsBroadcasts)`

SetBroadcasts sets Broadcasts field to given value.

### HasBroadcasts

`func (o *ProgramRelations) HasBroadcasts() bool`

HasBroadcasts returns a boolean if a field has been set.

### GetPresenters

`func (o *ProgramRelations) GetPresenters() ProgramRelationsPresenters`

GetPresenters returns the Presenters field if non-nil, zero value otherwise.

### GetPresentersOk

`func (o *ProgramRelations) GetPresentersOk() (*ProgramRelationsPresenters, bool)`

GetPresentersOk returns a tuple with the Presenters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresenters

`func (o *ProgramRelations) SetPresenters(v ProgramRelationsPresenters)`

SetPresenters sets Presenters field to given value.

### HasPresenters

`func (o *ProgramRelations) HasPresenters() bool`

HasPresenters returns a boolean if a field has been set.

### GetTags

`func (o *ProgramRelations) GetTags() ProgramRelationsTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ProgramRelations) GetTagsOk() (*ProgramRelationsTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ProgramRelations) SetTags(v ProgramRelationsTags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ProgramRelations) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetModelType

`func (o *ProgramRelations) GetModelType() BroadcastRelationsModelType`

GetModelType returns the ModelType field if non-nil, zero value otherwise.

### GetModelTypeOk

`func (o *ProgramRelations) GetModelTypeOk() (*BroadcastRelationsModelType, bool)`

GetModelTypeOk returns a tuple with the ModelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelType

`func (o *ProgramRelations) SetModelType(v BroadcastRelationsModelType)`

SetModelType sets ModelType field to given value.

### HasModelType

`func (o *ProgramRelations) HasModelType() bool`

HasModelType returns a boolean if a field has been set.

### GetGroup

`func (o *ProgramRelations) GetGroup() BroadcastRelationsGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *ProgramRelations) GetGroupOk() (*BroadcastRelationsGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *ProgramRelations) SetGroup(v BroadcastRelationsGroup)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *ProgramRelations) HasGroup() bool`

HasGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


