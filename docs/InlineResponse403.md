# InlineResponse403

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** | Describes what is forbiden. | [optional] 
**StatusCode** | Pointer to **int32** | Corresponding HTTP Response Status Code | [optional] 

## Methods

### NewInlineResponse403

`func NewInlineResponse403() *InlineResponse403`

NewInlineResponse403 instantiates a new InlineResponse403 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse403WithDefaults

`func NewInlineResponse403WithDefaults() *InlineResponse403`

NewInlineResponse403WithDefaults instantiates a new InlineResponse403 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *InlineResponse403) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *InlineResponse403) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *InlineResponse403) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *InlineResponse403) HasError() bool`

HasError returns a boolean if a field has been set.

### GetStatusCode

`func (o *InlineResponse403) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *InlineResponse403) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *InlineResponse403) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *InlineResponse403) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


