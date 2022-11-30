# BroadcastEPGDay

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Day** | Pointer to **string** |  | [optional] 
**Results** | [**[]BroadcastEPGResult**](BroadcastEPGResult.md) |  | 

## Methods

### NewBroadcastEPGDay

`func NewBroadcastEPGDay(results []BroadcastEPGResult, ) *BroadcastEPGDay`

NewBroadcastEPGDay instantiates a new BroadcastEPGDay object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBroadcastEPGDayWithDefaults

`func NewBroadcastEPGDayWithDefaults() *BroadcastEPGDay`

NewBroadcastEPGDayWithDefaults instantiates a new BroadcastEPGDay object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDay

`func (o *BroadcastEPGDay) GetDay() string`

GetDay returns the Day field if non-nil, zero value otherwise.

### GetDayOk

`func (o *BroadcastEPGDay) GetDayOk() (*string, bool)`

GetDayOk returns a tuple with the Day field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDay

`func (o *BroadcastEPGDay) SetDay(v string)`

SetDay sets Day field to given value.

### HasDay

`func (o *BroadcastEPGDay) HasDay() bool`

HasDay returns a boolean if a field has been set.

### GetResults

`func (o *BroadcastEPGDay) GetResults() []BroadcastEPGResult`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *BroadcastEPGDay) GetResultsOk() (*[]BroadcastEPGResult, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *BroadcastEPGDay) SetResults(v []BroadcastEPGResult)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


