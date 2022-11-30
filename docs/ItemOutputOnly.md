# ItemOutputOnly

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**DeletedAt** | Pointer to **time.Time** |  | [optional] 
**DataModified** | Pointer to **int64** |  | [optional] 
**Order** | Pointer to **int64** |  | [optional] 
**TemplateblockId** | Pointer to **int64** |  | [optional] 
**TemplateitemId** | Pointer to **int64** |  | [optional] 
**ExternalStationId** | Pointer to **int64** |  | [optional] 

## Methods

### NewItemOutputOnly

`func NewItemOutputOnly() *ItemOutputOnly`

NewItemOutputOnly instantiates a new ItemOutputOnly object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemOutputOnlyWithDefaults

`func NewItemOutputOnlyWithDefaults() *ItemOutputOnly`

NewItemOutputOnlyWithDefaults instantiates a new ItemOutputOnly object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ItemOutputOnly) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ItemOutputOnly) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ItemOutputOnly) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ItemOutputOnly) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ItemOutputOnly) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ItemOutputOnly) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ItemOutputOnly) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ItemOutputOnly) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ItemOutputOnly) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ItemOutputOnly) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ItemOutputOnly) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ItemOutputOnly) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *ItemOutputOnly) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *ItemOutputOnly) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *ItemOutputOnly) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *ItemOutputOnly) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetDataModified

`func (o *ItemOutputOnly) GetDataModified() int64`

GetDataModified returns the DataModified field if non-nil, zero value otherwise.

### GetDataModifiedOk

`func (o *ItemOutputOnly) GetDataModifiedOk() (*int64, bool)`

GetDataModifiedOk returns a tuple with the DataModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataModified

`func (o *ItemOutputOnly) SetDataModified(v int64)`

SetDataModified sets DataModified field to given value.

### HasDataModified

`func (o *ItemOutputOnly) HasDataModified() bool`

HasDataModified returns a boolean if a field has been set.

### GetOrder

`func (o *ItemOutputOnly) GetOrder() int64`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *ItemOutputOnly) GetOrderOk() (*int64, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *ItemOutputOnly) SetOrder(v int64)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *ItemOutputOnly) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetTemplateblockId

`func (o *ItemOutputOnly) GetTemplateblockId() int64`

GetTemplateblockId returns the TemplateblockId field if non-nil, zero value otherwise.

### GetTemplateblockIdOk

`func (o *ItemOutputOnly) GetTemplateblockIdOk() (*int64, bool)`

GetTemplateblockIdOk returns a tuple with the TemplateblockId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateblockId

`func (o *ItemOutputOnly) SetTemplateblockId(v int64)`

SetTemplateblockId sets TemplateblockId field to given value.

### HasTemplateblockId

`func (o *ItemOutputOnly) HasTemplateblockId() bool`

HasTemplateblockId returns a boolean if a field has been set.

### GetTemplateitemId

`func (o *ItemOutputOnly) GetTemplateitemId() int64`

GetTemplateitemId returns the TemplateitemId field if non-nil, zero value otherwise.

### GetTemplateitemIdOk

`func (o *ItemOutputOnly) GetTemplateitemIdOk() (*int64, bool)`

GetTemplateitemIdOk returns a tuple with the TemplateitemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateitemId

`func (o *ItemOutputOnly) SetTemplateitemId(v int64)`

SetTemplateitemId sets TemplateitemId field to given value.

### HasTemplateitemId

`func (o *ItemOutputOnly) HasTemplateitemId() bool`

HasTemplateitemId returns a boolean if a field has been set.

### GetExternalStationId

`func (o *ItemOutputOnly) GetExternalStationId() int64`

GetExternalStationId returns the ExternalStationId field if non-nil, zero value otherwise.

### GetExternalStationIdOk

`func (o *ItemOutputOnly) GetExternalStationIdOk() (*int64, bool)`

GetExternalStationIdOk returns a tuple with the ExternalStationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalStationId

`func (o *ItemOutputOnly) SetExternalStationId(v int64)`

SetExternalStationId sets ExternalStationId field to given value.

### HasExternalStationId

`func (o *ItemOutputOnly) HasExternalStationId() bool`

HasExternalStationId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


