# ContactRelations

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | [**ContactRelationsTags**](ContactRelationsTags.md) |  | 
**Items** | Pointer to [**ContactRelationsItems**](ContactRelationsItems.md) |  | [optional] 
**ModelType** | Pointer to [**BroadcastRelationsModelType**](BroadcastRelationsModelType.md) |  | [optional] 

## Methods

### NewContactRelations

`func NewContactRelations(tags ContactRelationsTags, ) *ContactRelations`

NewContactRelations instantiates a new ContactRelations object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactRelationsWithDefaults

`func NewContactRelationsWithDefaults() *ContactRelations`

NewContactRelationsWithDefaults instantiates a new ContactRelations object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *ContactRelations) GetTags() ContactRelationsTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ContactRelations) GetTagsOk() (*ContactRelationsTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ContactRelations) SetTags(v ContactRelationsTags)`

SetTags sets Tags field to given value.


### GetItems

`func (o *ContactRelations) GetItems() ContactRelationsItems`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ContactRelations) GetItemsOk() (*ContactRelationsItems, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ContactRelations) SetItems(v ContactRelationsItems)`

SetItems sets Items field to given value.

### HasItems

`func (o *ContactRelations) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetModelType

`func (o *ContactRelations) GetModelType() BroadcastRelationsModelType`

GetModelType returns the ModelType field if non-nil, zero value otherwise.

### GetModelTypeOk

`func (o *ContactRelations) GetModelTypeOk() (*BroadcastRelationsModelType, bool)`

GetModelTypeOk returns a tuple with the ModelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelType

`func (o *ContactRelations) SetModelType(v BroadcastRelationsModelType)`

SetModelType sets ModelType field to given value.

### HasModelType

`func (o *ContactRelations) HasModelType() bool`

HasModelType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


