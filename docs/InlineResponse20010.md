# InlineResponse20010

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentPage** | Pointer to **int64** |  | [optional] 
**From** | Pointer to **int64** |  | [optional] 
**LastPage** | Pointer to **int64** |  | [optional] 
**NextPageUrl** | Pointer to **string** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**PerPage** | Pointer to **int64** |  | [optional] 
**PrevPageUrl** | Pointer to **string** |  | [optional] 
**To** | Pointer to **int64** |  | [optional] 
**Total** | Pointer to **int64** |  | [optional] 
**Results** | Pointer to [**[]PresenterResult**](PresenterResult.md) |  | [optional] 

## Methods

### NewInlineResponse20010

`func NewInlineResponse20010() *InlineResponse20010`

NewInlineResponse20010 instantiates a new InlineResponse20010 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20010WithDefaults

`func NewInlineResponse20010WithDefaults() *InlineResponse20010`

NewInlineResponse20010WithDefaults instantiates a new InlineResponse20010 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentPage

`func (o *InlineResponse20010) GetCurrentPage() int64`

GetCurrentPage returns the CurrentPage field if non-nil, zero value otherwise.

### GetCurrentPageOk

`func (o *InlineResponse20010) GetCurrentPageOk() (*int64, bool)`

GetCurrentPageOk returns a tuple with the CurrentPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPage

`func (o *InlineResponse20010) SetCurrentPage(v int64)`

SetCurrentPage sets CurrentPage field to given value.

### HasCurrentPage

`func (o *InlineResponse20010) HasCurrentPage() bool`

HasCurrentPage returns a boolean if a field has been set.

### GetFrom

`func (o *InlineResponse20010) GetFrom() int64`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *InlineResponse20010) GetFromOk() (*int64, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *InlineResponse20010) SetFrom(v int64)`

SetFrom sets From field to given value.

### HasFrom

`func (o *InlineResponse20010) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetLastPage

`func (o *InlineResponse20010) GetLastPage() int64`

GetLastPage returns the LastPage field if non-nil, zero value otherwise.

### GetLastPageOk

`func (o *InlineResponse20010) GetLastPageOk() (*int64, bool)`

GetLastPageOk returns a tuple with the LastPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPage

`func (o *InlineResponse20010) SetLastPage(v int64)`

SetLastPage sets LastPage field to given value.

### HasLastPage

`func (o *InlineResponse20010) HasLastPage() bool`

HasLastPage returns a boolean if a field has been set.

### GetNextPageUrl

`func (o *InlineResponse20010) GetNextPageUrl() string`

GetNextPageUrl returns the NextPageUrl field if non-nil, zero value otherwise.

### GetNextPageUrlOk

`func (o *InlineResponse20010) GetNextPageUrlOk() (*string, bool)`

GetNextPageUrlOk returns a tuple with the NextPageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageUrl

`func (o *InlineResponse20010) SetNextPageUrl(v string)`

SetNextPageUrl sets NextPageUrl field to given value.

### HasNextPageUrl

`func (o *InlineResponse20010) HasNextPageUrl() bool`

HasNextPageUrl returns a boolean if a field has been set.

### GetPath

`func (o *InlineResponse20010) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *InlineResponse20010) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *InlineResponse20010) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *InlineResponse20010) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetPerPage

`func (o *InlineResponse20010) GetPerPage() int64`

GetPerPage returns the PerPage field if non-nil, zero value otherwise.

### GetPerPageOk

`func (o *InlineResponse20010) GetPerPageOk() (*int64, bool)`

GetPerPageOk returns a tuple with the PerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerPage

`func (o *InlineResponse20010) SetPerPage(v int64)`

SetPerPage sets PerPage field to given value.

### HasPerPage

`func (o *InlineResponse20010) HasPerPage() bool`

HasPerPage returns a boolean if a field has been set.

### GetPrevPageUrl

`func (o *InlineResponse20010) GetPrevPageUrl() string`

GetPrevPageUrl returns the PrevPageUrl field if non-nil, zero value otherwise.

### GetPrevPageUrlOk

`func (o *InlineResponse20010) GetPrevPageUrlOk() (*string, bool)`

GetPrevPageUrlOk returns a tuple with the PrevPageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevPageUrl

`func (o *InlineResponse20010) SetPrevPageUrl(v string)`

SetPrevPageUrl sets PrevPageUrl field to given value.

### HasPrevPageUrl

`func (o *InlineResponse20010) HasPrevPageUrl() bool`

HasPrevPageUrl returns a boolean if a field has been set.

### GetTo

`func (o *InlineResponse20010) GetTo() int64`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *InlineResponse20010) GetToOk() (*int64, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *InlineResponse20010) SetTo(v int64)`

SetTo sets To field to given value.

### HasTo

`func (o *InlineResponse20010) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetTotal

`func (o *InlineResponse20010) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *InlineResponse20010) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *InlineResponse20010) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *InlineResponse20010) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetResults

`func (o *InlineResponse20010) GetResults() []PresenterResult`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *InlineResponse20010) GetResultsOk() (*[]PresenterResult, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *InlineResponse20010) SetResults(v []PresenterResult)`

SetResults sets Results field to given value.

### HasResults

`func (o *InlineResponse20010) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


