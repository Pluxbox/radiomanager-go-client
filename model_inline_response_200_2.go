/*
RadioManager

This OpenAPI 3 Document describes the functionality of the API v2 of RadioManager. Note that no rights can be derived from this Document and the true functionality of the API might differ.

API version: 2.0
Contact: support@pluxbox.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package radiomanagerclient

import (
	"encoding/json"
)

// InlineResponse2002 struct for InlineResponse2002
type InlineResponse2002 struct {
	// Describes whatever has succeeded
	Success *string `json:"success,omitempty"`
	// ID of the newly posted object
	Id *int32 `json:"id,omitempty"`
}

// NewInlineResponse2002 instantiates a new InlineResponse2002 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2002() *InlineResponse2002 {
	this := InlineResponse2002{}
	return &this
}

// NewInlineResponse2002WithDefaults instantiates a new InlineResponse2002 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2002WithDefaults() *InlineResponse2002 {
	this := InlineResponse2002{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *InlineResponse2002) GetSuccess() string {
	if o == nil || isNil(o.Success) {
		var ret string
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2002) GetSuccessOk() (*string, bool) {
	if o == nil || isNil(o.Success) {
    return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *InlineResponse2002) HasSuccess() bool {
	if o != nil && !isNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given string and assigns it to the Success field.
func (o *InlineResponse2002) SetSuccess(v string) {
	o.Success = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse2002) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2002) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse2002) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *InlineResponse2002) SetId(v int32) {
	o.Id = &v
}

func (o InlineResponse2002) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2002 struct {
	value *InlineResponse2002
	isSet bool
}

func (v NullableInlineResponse2002) Get() *InlineResponse2002 {
	return v.value
}

func (v *NullableInlineResponse2002) Set(val *InlineResponse2002) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2002) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2002) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2002(val *InlineResponse2002) *NullableInlineResponse2002 {
	return &NullableInlineResponse2002{value: val, isSet: true}
}

func (v NullableInlineResponse2002) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2002) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


