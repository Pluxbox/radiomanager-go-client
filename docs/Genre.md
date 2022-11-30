# Genre

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**Urn** | Pointer to **string** |  | [optional] 
**ParentId** | Pointer to **int64** |  | [optional] 
**Name** | **string** |  | 

## Methods

### NewGenre

`func NewGenre(id int64, name string, ) *Genre`

NewGenre instantiates a new Genre object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenreWithDefaults

`func NewGenreWithDefaults() *Genre`

NewGenreWithDefaults instantiates a new Genre object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Genre) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Genre) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Genre) SetId(v int64)`

SetId sets Id field to given value.


### GetUrn

`func (o *Genre) GetUrn() string`

GetUrn returns the Urn field if non-nil, zero value otherwise.

### GetUrnOk

`func (o *Genre) GetUrnOk() (*string, bool)`

GetUrnOk returns a tuple with the Urn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrn

`func (o *Genre) SetUrn(v string)`

SetUrn sets Urn field to given value.

### HasUrn

`func (o *Genre) HasUrn() bool`

HasUrn returns a boolean if a field has been set.

### GetParentId

`func (o *Genre) GetParentId() int64`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *Genre) GetParentIdOk() (*int64, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *Genre) SetParentId(v int64)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *Genre) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetName

`func (o *Genre) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Genre) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Genre) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


