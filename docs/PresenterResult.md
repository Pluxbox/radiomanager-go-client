# PresenterResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**UpdatedAt** | **time.Time** |  | 
**CreatedAt** | **time.Time** |  | 
**DeletedAt** | **time.Time** |  | 
**ExternalStationId** | Pointer to **int64** |  | [optional] 
**ModelTypeId** | **int64** |  | 
**FieldValues** | Pointer to **map[string]interface{}** |  | [optional] 
**Firstname** | Pointer to **string** |  | [optional] 
**Lastname** | Pointer to **string** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Programs** | Pointer to [**PresenterRelationsPrograms**](PresenterRelationsPrograms.md) |  | [optional] 
**Broadcasts** | Pointer to [**PresenterRelationsBroadcasts**](PresenterRelationsBroadcasts.md) |  | [optional] 
**ModelType** | Pointer to [**PresenterRelationsModelType**](PresenterRelationsModelType.md) |  | [optional] 

## Methods

### NewPresenterResult

`func NewPresenterResult(id int64, updatedAt time.Time, createdAt time.Time, deletedAt time.Time, modelTypeId int64, ) *PresenterResult`

NewPresenterResult instantiates a new PresenterResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPresenterResultWithDefaults

`func NewPresenterResultWithDefaults() *PresenterResult`

NewPresenterResultWithDefaults instantiates a new PresenterResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PresenterResult) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PresenterResult) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PresenterResult) SetId(v int64)`

SetId sets Id field to given value.


### GetUpdatedAt

`func (o *PresenterResult) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PresenterResult) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PresenterResult) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetCreatedAt

`func (o *PresenterResult) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PresenterResult) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PresenterResult) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDeletedAt

`func (o *PresenterResult) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *PresenterResult) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *PresenterResult) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.


### GetExternalStationId

`func (o *PresenterResult) GetExternalStationId() int64`

GetExternalStationId returns the ExternalStationId field if non-nil, zero value otherwise.

### GetExternalStationIdOk

`func (o *PresenterResult) GetExternalStationIdOk() (*int64, bool)`

GetExternalStationIdOk returns a tuple with the ExternalStationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalStationId

`func (o *PresenterResult) SetExternalStationId(v int64)`

SetExternalStationId sets ExternalStationId field to given value.

### HasExternalStationId

`func (o *PresenterResult) HasExternalStationId() bool`

HasExternalStationId returns a boolean if a field has been set.

### GetModelTypeId

`func (o *PresenterResult) GetModelTypeId() int64`

GetModelTypeId returns the ModelTypeId field if non-nil, zero value otherwise.

### GetModelTypeIdOk

`func (o *PresenterResult) GetModelTypeIdOk() (*int64, bool)`

GetModelTypeIdOk returns a tuple with the ModelTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelTypeId

`func (o *PresenterResult) SetModelTypeId(v int64)`

SetModelTypeId sets ModelTypeId field to given value.


### GetFieldValues

`func (o *PresenterResult) GetFieldValues() map[string]interface{}`

GetFieldValues returns the FieldValues field if non-nil, zero value otherwise.

### GetFieldValuesOk

`func (o *PresenterResult) GetFieldValuesOk() (*map[string]interface{}, bool)`

GetFieldValuesOk returns a tuple with the FieldValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldValues

`func (o *PresenterResult) SetFieldValues(v map[string]interface{})`

SetFieldValues sets FieldValues field to given value.

### HasFieldValues

`func (o *PresenterResult) HasFieldValues() bool`

HasFieldValues returns a boolean if a field has been set.

### GetFirstname

`func (o *PresenterResult) GetFirstname() string`

GetFirstname returns the Firstname field if non-nil, zero value otherwise.

### GetFirstnameOk

`func (o *PresenterResult) GetFirstnameOk() (*string, bool)`

GetFirstnameOk returns a tuple with the Firstname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstname

`func (o *PresenterResult) SetFirstname(v string)`

SetFirstname sets Firstname field to given value.

### HasFirstname

`func (o *PresenterResult) HasFirstname() bool`

HasFirstname returns a boolean if a field has been set.

### GetLastname

`func (o *PresenterResult) GetLastname() string`

GetLastname returns the Lastname field if non-nil, zero value otherwise.

### GetLastnameOk

`func (o *PresenterResult) GetLastnameOk() (*string, bool)`

GetLastnameOk returns a tuple with the Lastname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastname

`func (o *PresenterResult) SetLastname(v string)`

SetLastname sets Lastname field to given value.

### HasLastname

`func (o *PresenterResult) HasLastname() bool`

HasLastname returns a boolean if a field has been set.

### GetActive

`func (o *PresenterResult) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *PresenterResult) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *PresenterResult) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *PresenterResult) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetName

`func (o *PresenterResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PresenterResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PresenterResult) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PresenterResult) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrograms

`func (o *PresenterResult) GetPrograms() PresenterRelationsPrograms`

GetPrograms returns the Programs field if non-nil, zero value otherwise.

### GetProgramsOk

`func (o *PresenterResult) GetProgramsOk() (*PresenterRelationsPrograms, bool)`

GetProgramsOk returns a tuple with the Programs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrograms

`func (o *PresenterResult) SetPrograms(v PresenterRelationsPrograms)`

SetPrograms sets Programs field to given value.

### HasPrograms

`func (o *PresenterResult) HasPrograms() bool`

HasPrograms returns a boolean if a field has been set.

### GetBroadcasts

`func (o *PresenterResult) GetBroadcasts() PresenterRelationsBroadcasts`

GetBroadcasts returns the Broadcasts field if non-nil, zero value otherwise.

### GetBroadcastsOk

`func (o *PresenterResult) GetBroadcastsOk() (*PresenterRelationsBroadcasts, bool)`

GetBroadcastsOk returns a tuple with the Broadcasts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBroadcasts

`func (o *PresenterResult) SetBroadcasts(v PresenterRelationsBroadcasts)`

SetBroadcasts sets Broadcasts field to given value.

### HasBroadcasts

`func (o *PresenterResult) HasBroadcasts() bool`

HasBroadcasts returns a boolean if a field has been set.

### GetModelType

`func (o *PresenterResult) GetModelType() PresenterRelationsModelType`

GetModelType returns the ModelType field if non-nil, zero value otherwise.

### GetModelTypeOk

`func (o *PresenterResult) GetModelTypeOk() (*PresenterRelationsModelType, bool)`

GetModelTypeOk returns a tuple with the ModelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelType

`func (o *PresenterResult) SetModelType(v PresenterRelationsModelType)`

SetModelType sets ModelType field to given value.

### HasModelType

`func (o *PresenterResult) HasModelType() bool`

HasModelType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


