# ItemRelationsBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**Operation** | Pointer to **string** |  | [optional] 
**Params** | Pointer to [**BlockRelationsBroadcastParams**](BlockRelationsBroadcastParams.md) |  | [optional] 

## Methods

### NewItemRelationsBlock

`func NewItemRelationsBlock() *ItemRelationsBlock`

NewItemRelationsBlock instantiates a new ItemRelationsBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemRelationsBlockWithDefaults

`func NewItemRelationsBlockWithDefaults() *ItemRelationsBlock`

NewItemRelationsBlockWithDefaults instantiates a new ItemRelationsBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *ItemRelationsBlock) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ItemRelationsBlock) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ItemRelationsBlock) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *ItemRelationsBlock) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetModel

`func (o *ItemRelationsBlock) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ItemRelationsBlock) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ItemRelationsBlock) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ItemRelationsBlock) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetOperation

`func (o *ItemRelationsBlock) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *ItemRelationsBlock) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *ItemRelationsBlock) SetOperation(v string)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *ItemRelationsBlock) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetParams

`func (o *ItemRelationsBlock) GetParams() BlockRelationsBroadcastParams`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *ItemRelationsBlock) GetParamsOk() (*BlockRelationsBroadcastParams, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *ItemRelationsBlock) SetParams(v BlockRelationsBroadcastParams)`

SetParams sets Params field to given value.

### HasParams

`func (o *ItemRelationsBlock) HasParams() bool`

HasParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


