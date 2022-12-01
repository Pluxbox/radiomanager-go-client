# InlineResponse500

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** | Given error (not formatted), describes the problem. | [optional] 
**StatusCode** | Pointer to **int64** | Corresponding HTTP Response Status Code | [optional] 

## Methods

### NewInlineResponse500

`func NewInlineResponse500() *InlineResponse500`

NewInlineResponse500 instantiates a new InlineResponse500 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse500WithDefaults

`func NewInlineResponse500WithDefaults() *InlineResponse500`

NewInlineResponse500WithDefaults instantiates a new InlineResponse500 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *InlineResponse500) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *InlineResponse500) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *InlineResponse500) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *InlineResponse500) HasError() bool`

HasError returns a boolean if a field has been set.

### GetStatusCode

`func (o *InlineResponse500) GetStatusCode() int64`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *InlineResponse500) GetStatusCodeOk() (*int64, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *InlineResponse500) SetStatusCode(v int64)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *InlineResponse500) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

