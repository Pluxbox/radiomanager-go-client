# InlineResponse401

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** | Unauthorized | [optional] 
**ErrorDescription** | Pointer to **string** | Given error, describes the problem. | [optional] 
**StatusCode** | Pointer to **int32** | Corresponding HTTP Response Status Code | [optional] 

## Methods

### NewInlineResponse401

`func NewInlineResponse401() *InlineResponse401`

NewInlineResponse401 instantiates a new InlineResponse401 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse401WithDefaults

`func NewInlineResponse401WithDefaults() *InlineResponse401`

NewInlineResponse401WithDefaults instantiates a new InlineResponse401 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *InlineResponse401) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *InlineResponse401) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *InlineResponse401) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *InlineResponse401) HasError() bool`

HasError returns a boolean if a field has been set.

### GetErrorDescription

`func (o *InlineResponse401) GetErrorDescription() string`

GetErrorDescription returns the ErrorDescription field if non-nil, zero value otherwise.

### GetErrorDescriptionOk

`func (o *InlineResponse401) GetErrorDescriptionOk() (*string, bool)`

GetErrorDescriptionOk returns a tuple with the ErrorDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDescription

`func (o *InlineResponse401) SetErrorDescription(v string)`

SetErrorDescription sets ErrorDescription field to given value.

### HasErrorDescription

`func (o *InlineResponse401) HasErrorDescription() bool`

HasErrorDescription returns a boolean if a field has been set.

### GetStatusCode

`func (o *InlineResponse401) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *InlineResponse401) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *InlineResponse401) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *InlineResponse401) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


