# EPGResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Days** | [**[]BroadcastEPGDay**](BroadcastEPGDay.md) |  | 
**NextPageUrl** | **string** |  | 
**PrevPageUrl** | **string** |  | 

## Methods

### NewEPGResults

`func NewEPGResults(days []BroadcastEPGDay, nextPageUrl string, prevPageUrl string, ) *EPGResults`

NewEPGResults instantiates a new EPGResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEPGResultsWithDefaults

`func NewEPGResultsWithDefaults() *EPGResults`

NewEPGResultsWithDefaults instantiates a new EPGResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDays

`func (o *EPGResults) GetDays() []BroadcastEPGDay`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *EPGResults) GetDaysOk() (*[]BroadcastEPGDay, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *EPGResults) SetDays(v []BroadcastEPGDay)`

SetDays sets Days field to given value.


### GetNextPageUrl

`func (o *EPGResults) GetNextPageUrl() string`

GetNextPageUrl returns the NextPageUrl field if non-nil, zero value otherwise.

### GetNextPageUrlOk

`func (o *EPGResults) GetNextPageUrlOk() (*string, bool)`

GetNextPageUrlOk returns a tuple with the NextPageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageUrl

`func (o *EPGResults) SetNextPageUrl(v string)`

SetNextPageUrl sets NextPageUrl field to given value.


### GetPrevPageUrl

`func (o *EPGResults) GetPrevPageUrl() string`

GetPrevPageUrl returns the PrevPageUrl field if non-nil, zero value otherwise.

### GetPrevPageUrlOk

`func (o *EPGResults) GetPrevPageUrlOk() (*string, bool)`

GetPrevPageUrlOk returns a tuple with the PrevPageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevPageUrl

`func (o *EPGResults) SetPrevPageUrl(v string)`

SetPrevPageUrl sets PrevPageUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


