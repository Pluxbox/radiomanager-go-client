# GenreResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**Urn** | Pointer to **string** |  | [optional] 
**ParentId** | Pointer to **int64** |  | [optional] 
**Name** | **string** |  | 
**Broadcasts** | Pointer to [**GenreRelationsBroadcasts**](GenreRelationsBroadcasts.md) |  | [optional] 
**Programs** | Pointer to [**GenreRelationsPrograms**](GenreRelationsPrograms.md) |  | [optional] 
**ExternalStationId** | Pointer to **int64** |  | [optional] 

## Methods

### NewGenreResult

`func NewGenreResult(id int64, name string, ) *GenreResult`

NewGenreResult instantiates a new GenreResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenreResultWithDefaults

`func NewGenreResultWithDefaults() *GenreResult`

NewGenreResultWithDefaults instantiates a new GenreResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GenreResult) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GenreResult) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GenreResult) SetId(v int64)`

SetId sets Id field to given value.


### GetUrn

`func (o *GenreResult) GetUrn() string`

GetUrn returns the Urn field if non-nil, zero value otherwise.

### GetUrnOk

`func (o *GenreResult) GetUrnOk() (*string, bool)`

GetUrnOk returns a tuple with the Urn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrn

`func (o *GenreResult) SetUrn(v string)`

SetUrn sets Urn field to given value.

### HasUrn

`func (o *GenreResult) HasUrn() bool`

HasUrn returns a boolean if a field has been set.

### GetParentId

`func (o *GenreResult) GetParentId() int64`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *GenreResult) GetParentIdOk() (*int64, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *GenreResult) SetParentId(v int64)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *GenreResult) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetName

`func (o *GenreResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GenreResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GenreResult) SetName(v string)`

SetName sets Name field to given value.


### GetBroadcasts

`func (o *GenreResult) GetBroadcasts() GenreRelationsBroadcasts`

GetBroadcasts returns the Broadcasts field if non-nil, zero value otherwise.

### GetBroadcastsOk

`func (o *GenreResult) GetBroadcastsOk() (*GenreRelationsBroadcasts, bool)`

GetBroadcastsOk returns a tuple with the Broadcasts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBroadcasts

`func (o *GenreResult) SetBroadcasts(v GenreRelationsBroadcasts)`

SetBroadcasts sets Broadcasts field to given value.

### HasBroadcasts

`func (o *GenreResult) HasBroadcasts() bool`

HasBroadcasts returns a boolean if a field has been set.

### GetPrograms

`func (o *GenreResult) GetPrograms() GenreRelationsPrograms`

GetPrograms returns the Programs field if non-nil, zero value otherwise.

### GetProgramsOk

`func (o *GenreResult) GetProgramsOk() (*GenreRelationsPrograms, bool)`

GetProgramsOk returns a tuple with the Programs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrograms

`func (o *GenreResult) SetPrograms(v GenreRelationsPrograms)`

SetPrograms sets Programs field to given value.

### HasPrograms

`func (o *GenreResult) HasPrograms() bool`

HasPrograms returns a boolean if a field has been set.

### GetExternalStationId

`func (o *GenreResult) GetExternalStationId() int64`

GetExternalStationId returns the ExternalStationId field if non-nil, zero value otherwise.

### GetExternalStationIdOk

`func (o *GenreResult) GetExternalStationIdOk() (*int64, bool)`

GetExternalStationIdOk returns a tuple with the ExternalStationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalStationId

`func (o *GenreResult) SetExternalStationId(v int64)`

SetExternalStationId sets ExternalStationId field to given value.

### HasExternalStationId

`func (o *GenreResult) HasExternalStationId() bool`

HasExternalStationId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


