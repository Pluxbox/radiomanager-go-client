# BroadcastRelations

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Genre** | Pointer to [**BroadcastRelationsGenre**](BroadcastRelationsGenre.md) |  | [optional] 
**Items** | Pointer to [**BroadcastRelationsItems**](BroadcastRelationsItems.md) |  | [optional] 
**Blocks** | Pointer to [**BroadcastRelationsBlocks**](BroadcastRelationsBlocks.md) |  | [optional] 
**Program** | Pointer to [**BlockRelationsProgram**](BlockRelationsProgram.md) |  | [optional] 
**Tags** | Pointer to [**BroadcastRelationsTags**](BroadcastRelationsTags.md) |  | [optional] 
**Presenters** | Pointer to [**BroadcastRelationsPresenters**](BroadcastRelationsPresenters.md) |  | [optional] 
**ModelType** | Pointer to [**BroadcastRelationsModelType**](BroadcastRelationsModelType.md) |  | [optional] 
**Group** | Pointer to [**BroadcastRelationsGroup**](BroadcastRelationsGroup.md) |  | [optional] 

## Methods

### NewBroadcastRelations

`func NewBroadcastRelations() *BroadcastRelations`

NewBroadcastRelations instantiates a new BroadcastRelations object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBroadcastRelationsWithDefaults

`func NewBroadcastRelationsWithDefaults() *BroadcastRelations`

NewBroadcastRelationsWithDefaults instantiates a new BroadcastRelations object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGenre

`func (o *BroadcastRelations) GetGenre() BroadcastRelationsGenre`

GetGenre returns the Genre field if non-nil, zero value otherwise.

### GetGenreOk

`func (o *BroadcastRelations) GetGenreOk() (*BroadcastRelationsGenre, bool)`

GetGenreOk returns a tuple with the Genre field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenre

`func (o *BroadcastRelations) SetGenre(v BroadcastRelationsGenre)`

SetGenre sets Genre field to given value.

### HasGenre

`func (o *BroadcastRelations) HasGenre() bool`

HasGenre returns a boolean if a field has been set.

### GetItems

`func (o *BroadcastRelations) GetItems() BroadcastRelationsItems`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *BroadcastRelations) GetItemsOk() (*BroadcastRelationsItems, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *BroadcastRelations) SetItems(v BroadcastRelationsItems)`

SetItems sets Items field to given value.

### HasItems

`func (o *BroadcastRelations) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetBlocks

`func (o *BroadcastRelations) GetBlocks() BroadcastRelationsBlocks`

GetBlocks returns the Blocks field if non-nil, zero value otherwise.

### GetBlocksOk

`func (o *BroadcastRelations) GetBlocksOk() (*BroadcastRelationsBlocks, bool)`

GetBlocksOk returns a tuple with the Blocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocks

`func (o *BroadcastRelations) SetBlocks(v BroadcastRelationsBlocks)`

SetBlocks sets Blocks field to given value.

### HasBlocks

`func (o *BroadcastRelations) HasBlocks() bool`

HasBlocks returns a boolean if a field has been set.

### GetProgram

`func (o *BroadcastRelations) GetProgram() BlockRelationsProgram`

GetProgram returns the Program field if non-nil, zero value otherwise.

### GetProgramOk

`func (o *BroadcastRelations) GetProgramOk() (*BlockRelationsProgram, bool)`

GetProgramOk returns a tuple with the Program field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgram

`func (o *BroadcastRelations) SetProgram(v BlockRelationsProgram)`

SetProgram sets Program field to given value.

### HasProgram

`func (o *BroadcastRelations) HasProgram() bool`

HasProgram returns a boolean if a field has been set.

### GetTags

`func (o *BroadcastRelations) GetTags() BroadcastRelationsTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BroadcastRelations) GetTagsOk() (*BroadcastRelationsTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BroadcastRelations) SetTags(v BroadcastRelationsTags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BroadcastRelations) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetPresenters

`func (o *BroadcastRelations) GetPresenters() BroadcastRelationsPresenters`

GetPresenters returns the Presenters field if non-nil, zero value otherwise.

### GetPresentersOk

`func (o *BroadcastRelations) GetPresentersOk() (*BroadcastRelationsPresenters, bool)`

GetPresentersOk returns a tuple with the Presenters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresenters

`func (o *BroadcastRelations) SetPresenters(v BroadcastRelationsPresenters)`

SetPresenters sets Presenters field to given value.

### HasPresenters

`func (o *BroadcastRelations) HasPresenters() bool`

HasPresenters returns a boolean if a field has been set.

### GetModelType

`func (o *BroadcastRelations) GetModelType() BroadcastRelationsModelType`

GetModelType returns the ModelType field if non-nil, zero value otherwise.

### GetModelTypeOk

`func (o *BroadcastRelations) GetModelTypeOk() (*BroadcastRelationsModelType, bool)`

GetModelTypeOk returns a tuple with the ModelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelType

`func (o *BroadcastRelations) SetModelType(v BroadcastRelationsModelType)`

SetModelType sets ModelType field to given value.

### HasModelType

`func (o *BroadcastRelations) HasModelType() bool`

HasModelType returns a boolean if a field has been set.

### GetGroup

`func (o *BroadcastRelations) GetGroup() BroadcastRelationsGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *BroadcastRelations) GetGroupOk() (*BroadcastRelationsGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *BroadcastRelations) SetGroup(v BroadcastRelationsGroup)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *BroadcastRelations) HasGroup() bool`

HasGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


