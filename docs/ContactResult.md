# ContactResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**DeletedAt** | Pointer to **time.Time** |  | [optional] 
**ExternalStationId** | Pointer to **int64** |  | [optional] 
**ModelTypeId** | **int64** |  | 
**FieldValues** | Pointer to **map[string]interface{}** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Firstname** | **string** |  | 
**Lastname** | **string** |  | 
**Phone** | Pointer to **string** |  | [optional] 
**Tags** | [**ContactRelationsTags**](ContactRelationsTags.md) |  | 
**Items** | Pointer to [**ContactRelationsItems**](ContactRelationsItems.md) |  | [optional] 
**ModelType** | Pointer to [**BroadcastRelationsModelType**](BroadcastRelationsModelType.md) |  | [optional] 

## Methods

### NewContactResult

`func NewContactResult(modelTypeId int64, firstname string, lastname string, tags ContactRelationsTags, ) *ContactResult`

NewContactResult instantiates a new ContactResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactResultWithDefaults

`func NewContactResultWithDefaults() *ContactResult`

NewContactResultWithDefaults instantiates a new ContactResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ContactResult) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ContactResult) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ContactResult) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ContactResult) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ContactResult) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ContactResult) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ContactResult) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ContactResult) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ContactResult) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ContactResult) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ContactResult) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ContactResult) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *ContactResult) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *ContactResult) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *ContactResult) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *ContactResult) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetExternalStationId

`func (o *ContactResult) GetExternalStationId() int64`

GetExternalStationId returns the ExternalStationId field if non-nil, zero value otherwise.

### GetExternalStationIdOk

`func (o *ContactResult) GetExternalStationIdOk() (*int64, bool)`

GetExternalStationIdOk returns a tuple with the ExternalStationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalStationId

`func (o *ContactResult) SetExternalStationId(v int64)`

SetExternalStationId sets ExternalStationId field to given value.

### HasExternalStationId

`func (o *ContactResult) HasExternalStationId() bool`

HasExternalStationId returns a boolean if a field has been set.

### GetModelTypeId

`func (o *ContactResult) GetModelTypeId() int64`

GetModelTypeId returns the ModelTypeId field if non-nil, zero value otherwise.

### GetModelTypeIdOk

`func (o *ContactResult) GetModelTypeIdOk() (*int64, bool)`

GetModelTypeIdOk returns a tuple with the ModelTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelTypeId

`func (o *ContactResult) SetModelTypeId(v int64)`

SetModelTypeId sets ModelTypeId field to given value.


### GetFieldValues

`func (o *ContactResult) GetFieldValues() map[string]interface{}`

GetFieldValues returns the FieldValues field if non-nil, zero value otherwise.

### GetFieldValuesOk

`func (o *ContactResult) GetFieldValuesOk() (*map[string]interface{}, bool)`

GetFieldValuesOk returns a tuple with the FieldValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldValues

`func (o *ContactResult) SetFieldValues(v map[string]interface{})`

SetFieldValues sets FieldValues field to given value.

### HasFieldValues

`func (o *ContactResult) HasFieldValues() bool`

HasFieldValues returns a boolean if a field has been set.

### GetEmail

`func (o *ContactResult) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ContactResult) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ContactResult) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ContactResult) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetFirstname

`func (o *ContactResult) GetFirstname() string`

GetFirstname returns the Firstname field if non-nil, zero value otherwise.

### GetFirstnameOk

`func (o *ContactResult) GetFirstnameOk() (*string, bool)`

GetFirstnameOk returns a tuple with the Firstname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstname

`func (o *ContactResult) SetFirstname(v string)`

SetFirstname sets Firstname field to given value.


### GetLastname

`func (o *ContactResult) GetLastname() string`

GetLastname returns the Lastname field if non-nil, zero value otherwise.

### GetLastnameOk

`func (o *ContactResult) GetLastnameOk() (*string, bool)`

GetLastnameOk returns a tuple with the Lastname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastname

`func (o *ContactResult) SetLastname(v string)`

SetLastname sets Lastname field to given value.


### GetPhone

`func (o *ContactResult) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *ContactResult) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *ContactResult) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *ContactResult) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetTags

`func (o *ContactResult) GetTags() ContactRelationsTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ContactResult) GetTagsOk() (*ContactRelationsTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ContactResult) SetTags(v ContactRelationsTags)`

SetTags sets Tags field to given value.


### GetItems

`func (o *ContactResult) GetItems() ContactRelationsItems`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ContactResult) GetItemsOk() (*ContactRelationsItems, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ContactResult) SetItems(v ContactRelationsItems)`

SetItems sets Items field to given value.

### HasItems

`func (o *ContactResult) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetModelType

`func (o *ContactResult) GetModelType() BroadcastRelationsModelType`

GetModelType returns the ModelType field if non-nil, zero value otherwise.

### GetModelTypeOk

`func (o *ContactResult) GetModelTypeOk() (*BroadcastRelationsModelType, bool)`

GetModelTypeOk returns a tuple with the ModelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelType

`func (o *ContactResult) SetModelType(v BroadcastRelationsModelType)`

SetModelType sets ModelType field to given value.

### HasModelType

`func (o *ContactResult) HasModelType() bool`

HasModelType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


