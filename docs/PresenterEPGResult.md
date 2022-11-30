# PresenterEPGResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ModelTypeId** | **int64** |  | 
**FieldValues** | Pointer to **map[string]interface{}** |  | [optional] 
**Firstname** | Pointer to **string** |  | [optional] 
**Lastname** | Pointer to **string** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Id** | **int64** |  | 
**UpdatedAt** | **time.Time** |  | 
**CreatedAt** | **time.Time** |  | 
**DeletedAt** | **time.Time** |  | 
**ExternalStationId** | Pointer to **int64** |  | [optional] 

## Methods

### NewPresenterEPGResult

`func NewPresenterEPGResult(modelTypeId int64, id int64, updatedAt time.Time, createdAt time.Time, deletedAt time.Time, ) *PresenterEPGResult`

NewPresenterEPGResult instantiates a new PresenterEPGResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPresenterEPGResultWithDefaults

`func NewPresenterEPGResultWithDefaults() *PresenterEPGResult`

NewPresenterEPGResultWithDefaults instantiates a new PresenterEPGResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModelTypeId

`func (o *PresenterEPGResult) GetModelTypeId() int64`

GetModelTypeId returns the ModelTypeId field if non-nil, zero value otherwise.

### GetModelTypeIdOk

`func (o *PresenterEPGResult) GetModelTypeIdOk() (*int64, bool)`

GetModelTypeIdOk returns a tuple with the ModelTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelTypeId

`func (o *PresenterEPGResult) SetModelTypeId(v int64)`

SetModelTypeId sets ModelTypeId field to given value.


### GetFieldValues

`func (o *PresenterEPGResult) GetFieldValues() map[string]interface{}`

GetFieldValues returns the FieldValues field if non-nil, zero value otherwise.

### GetFieldValuesOk

`func (o *PresenterEPGResult) GetFieldValuesOk() (*map[string]interface{}, bool)`

GetFieldValuesOk returns a tuple with the FieldValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldValues

`func (o *PresenterEPGResult) SetFieldValues(v map[string]interface{})`

SetFieldValues sets FieldValues field to given value.

### HasFieldValues

`func (o *PresenterEPGResult) HasFieldValues() bool`

HasFieldValues returns a boolean if a field has been set.

### GetFirstname

`func (o *PresenterEPGResult) GetFirstname() string`

GetFirstname returns the Firstname field if non-nil, zero value otherwise.

### GetFirstnameOk

`func (o *PresenterEPGResult) GetFirstnameOk() (*string, bool)`

GetFirstnameOk returns a tuple with the Firstname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstname

`func (o *PresenterEPGResult) SetFirstname(v string)`

SetFirstname sets Firstname field to given value.

### HasFirstname

`func (o *PresenterEPGResult) HasFirstname() bool`

HasFirstname returns a boolean if a field has been set.

### GetLastname

`func (o *PresenterEPGResult) GetLastname() string`

GetLastname returns the Lastname field if non-nil, zero value otherwise.

### GetLastnameOk

`func (o *PresenterEPGResult) GetLastnameOk() (*string, bool)`

GetLastnameOk returns a tuple with the Lastname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastname

`func (o *PresenterEPGResult) SetLastname(v string)`

SetLastname sets Lastname field to given value.

### HasLastname

`func (o *PresenterEPGResult) HasLastname() bool`

HasLastname returns a boolean if a field has been set.

### GetActive

`func (o *PresenterEPGResult) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *PresenterEPGResult) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *PresenterEPGResult) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *PresenterEPGResult) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetName

`func (o *PresenterEPGResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PresenterEPGResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PresenterEPGResult) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PresenterEPGResult) HasName() bool`

HasName returns a boolean if a field has been set.

### GetId

`func (o *PresenterEPGResult) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PresenterEPGResult) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PresenterEPGResult) SetId(v int64)`

SetId sets Id field to given value.


### GetUpdatedAt

`func (o *PresenterEPGResult) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PresenterEPGResult) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PresenterEPGResult) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetCreatedAt

`func (o *PresenterEPGResult) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PresenterEPGResult) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PresenterEPGResult) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDeletedAt

`func (o *PresenterEPGResult) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *PresenterEPGResult) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *PresenterEPGResult) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.


### GetExternalStationId

`func (o *PresenterEPGResult) GetExternalStationId() int64`

GetExternalStationId returns the ExternalStationId field if non-nil, zero value otherwise.

### GetExternalStationIdOk

`func (o *PresenterEPGResult) GetExternalStationIdOk() (*int64, bool)`

GetExternalStationIdOk returns a tuple with the ExternalStationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalStationId

`func (o *PresenterEPGResult) SetExternalStationId(v int64)`

SetExternalStationId sets ExternalStationId field to given value.

### HasExternalStationId

`func (o *PresenterEPGResult) HasExternalStationId() bool`

HasExternalStationId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


