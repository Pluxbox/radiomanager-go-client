# InlineResponse429

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** | Rate limit response | [optional] 
**StatusCode** | Pointer to **int64** | Corresponding HTTP Response Status Code | [optional] 

## Methods

### NewInlineResponse429

`func NewInlineResponse429() *InlineResponse429`

NewInlineResponse429 instantiates a new InlineResponse429 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse429WithDefaults

`func NewInlineResponse429WithDefaults() *InlineResponse429`

NewInlineResponse429WithDefaults instantiates a new InlineResponse429 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *InlineResponse429) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *InlineResponse429) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *InlineResponse429) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *InlineResponse429) HasError() bool`

HasError returns a boolean if a field has been set.

### GetStatusCode

`func (o *InlineResponse429) GetStatusCode() int64`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *InlineResponse429) GetStatusCodeOk() (*int64, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *InlineResponse429) SetStatusCode(v int64)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *InlineResponse429) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


