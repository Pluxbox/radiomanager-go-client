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

// TagRelationsBroadcastsParams struct for TagRelationsBroadcastsParams
type TagRelationsBroadcastsParams struct {
	TagId *int32 `json:"tag_id,omitempty"`
}

// NewTagRelationsBroadcastsParams instantiates a new TagRelationsBroadcastsParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTagRelationsBroadcastsParams() *TagRelationsBroadcastsParams {
	this := TagRelationsBroadcastsParams{}
	return &this
}

// NewTagRelationsBroadcastsParamsWithDefaults instantiates a new TagRelationsBroadcastsParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTagRelationsBroadcastsParamsWithDefaults() *TagRelationsBroadcastsParams {
	this := TagRelationsBroadcastsParams{}
	return &this
}

// GetTagId returns the TagId field value if set, zero value otherwise.
func (o *TagRelationsBroadcastsParams) GetTagId() int32 {
	if o == nil || isNil(o.TagId) {
		var ret int32
		return ret
	}
	return *o.TagId
}

// GetTagIdOk returns a tuple with the TagId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagRelationsBroadcastsParams) GetTagIdOk() (*int32, bool) {
	if o == nil || isNil(o.TagId) {
    return nil, false
	}
	return o.TagId, true
}

// HasTagId returns a boolean if a field has been set.
func (o *TagRelationsBroadcastsParams) HasTagId() bool {
	if o != nil && !isNil(o.TagId) {
		return true
	}

	return false
}

// SetTagId gets a reference to the given int32 and assigns it to the TagId field.
func (o *TagRelationsBroadcastsParams) SetTagId(v int32) {
	o.TagId = &v
}

func (o TagRelationsBroadcastsParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.TagId) {
		toSerialize["tag_id"] = o.TagId
	}
	return json.Marshal(toSerialize)
}

type NullableTagRelationsBroadcastsParams struct {
	value *TagRelationsBroadcastsParams
	isSet bool
}

func (v NullableTagRelationsBroadcastsParams) Get() *TagRelationsBroadcastsParams {
	return v.value
}

func (v *NullableTagRelationsBroadcastsParams) Set(val *TagRelationsBroadcastsParams) {
	v.value = val
	v.isSet = true
}

func (v NullableTagRelationsBroadcastsParams) IsSet() bool {
	return v.isSet
}

func (v *NullableTagRelationsBroadcastsParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTagRelationsBroadcastsParams(val *TagRelationsBroadcastsParams) *NullableTagRelationsBroadcastsParams {
	return &NullableTagRelationsBroadcastsParams{value: val, isSet: true}
}

func (v NullableTagRelationsBroadcastsParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTagRelationsBroadcastsParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

