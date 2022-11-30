# InlineResponse2008

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentPage** | Pointer to **int64** |  | [optional] [readonly] 
**From** | Pointer to **int64** |  | [optional] [readonly] 
**LastPage** | Pointer to **int64** |  | [optional] [readonly] 
**NextPageUrl** | Pointer to **string** |  | [optional] [readonly] 
**Path** | Pointer to **string** |  | [optional] [readonly] 
**PerPage** | Pointer to **int64** |  | [optional] [readonly] 
**To** | Pointer to **int64** |  | [optional] [readonly] 
**Total** | Pointer to **int64** |  | [optional] [readonly] 
**Results** | Pointer to [**[]ItemResult**](ItemResult.md) |  | [optional] 

## Methods

### NewInlineResponse2008

`func NewInlineResponse2008() *InlineResponse2008`

NewInlineResponse2008 instantiates a new InlineResponse2008 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2008WithDefaults

`func NewInlineResponse2008WithDefaults() *InlineResponse2008`

NewInlineResponse2008WithDefaults instantiates a new InlineResponse2008 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentPage

`func (o *InlineResponse2008) GetCurrentPage() int64`

GetCurrentPage returns the CurrentPage field if non-nil, zero value otherwise.

### GetCurrentPageOk

`func (o *InlineResponse2008) GetCurrentPageOk() (*int64, bool)`

GetCurrentPageOk returns a tuple with the CurrentPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPage

`func (o *InlineResponse2008) SetCurrentPage(v int64)`

SetCurrentPage sets CurrentPage field to given value.

### HasCurrentPage

`func (o *InlineResponse2008) HasCurrentPage() bool`

HasCurrentPage returns a boolean if a field has been set.

### GetFrom

`func (o *InlineResponse2008) GetFrom() int64`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *InlineResponse2008) GetFromOk() (*int64, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *InlineResponse2008) SetFrom(v int64)`

SetFrom sets From field to given value.

### HasFrom

`func (o *InlineResponse2008) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetLastPage

`func (o *InlineResponse2008) GetLastPage() int64`

GetLastPage returns the LastPage field if non-nil, zero value otherwise.

### GetLastPageOk

`func (o *InlineResponse2008) GetLastPageOk() (*int64, bool)`

GetLastPageOk returns a tuple with the LastPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPage

`func (o *InlineResponse2008) SetLastPage(v int64)`

SetLastPage sets LastPage field to given value.

### HasLastPage

`func (o *InlineResponse2008) HasLastPage() bool`

HasLastPage returns a boolean if a field has been set.

### GetNextPageUrl

`func (o *InlineResponse2008) GetNextPageUrl() string`

GetNextPageUrl returns the NextPageUrl field if non-nil, zero value otherwise.

### GetNextPageUrlOk

`func (o *InlineResponse2008) GetNextPageUrlOk() (*string, bool)`

GetNextPageUrlOk returns a tuple with the NextPageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageUrl

`func (o *InlineResponse2008) SetNextPageUrl(v string)`

SetNextPageUrl sets NextPageUrl field to given value.

### HasNextPageUrl

`func (o *InlineResponse2008) HasNextPageUrl() bool`

HasNextPageUrl returns a boolean if a field has been set.

### GetPath

`func (o *InlineResponse2008) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *InlineResponse2008) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *InlineResponse2008) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *InlineResponse2008) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetPerPage

`func (o *InlineResponse2008) GetPerPage() int64`

GetPerPage returns the PerPage field if non-nil, zero value otherwise.

### GetPerPageOk

`func (o *InlineResponse2008) GetPerPageOk() (*int64, bool)`

GetPerPageOk returns a tuple with the PerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerPage

`func (o *InlineResponse2008) SetPerPage(v int64)`

SetPerPage sets PerPage field to given value.

### HasPerPage

`func (o *InlineResponse2008) HasPerPage() bool`

HasPerPage returns a boolean if a field has been set.

### GetTo

`func (o *InlineResponse2008) GetTo() int64`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *InlineResponse2008) GetToOk() (*int64, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *InlineResponse2008) SetTo(v int64)`

SetTo sets To field to given value.

### HasTo

`func (o *InlineResponse2008) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetTotal

`func (o *InlineResponse2008) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *InlineResponse2008) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *InlineResponse2008) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *InlineResponse2008) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetResults

`func (o *InlineResponse2008) GetResults() []ItemResult`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *InlineResponse2008) GetResultsOk() (*[]ItemResult, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *InlineResponse2008) SetResults(v []ItemResult)`

SetResults sets Results field to given value.

### HasResults

`func (o *InlineResponse2008) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


