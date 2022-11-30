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

// ProgramInputOnly struct for ProgramInputOnly
type ProgramInputOnly struct {
	Tags []int32 `json:"tags,omitempty"`
	Presenters []int32 `json:"presenters,omitempty"`
}

// NewProgramInputOnly instantiates a new ProgramInputOnly object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProgramInputOnly() *ProgramInputOnly {
	this := ProgramInputOnly{}
	return &this
}

// NewProgramInputOnlyWithDefaults instantiates a new ProgramInputOnly object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProgramInputOnlyWithDefaults() *ProgramInputOnly {
	this := ProgramInputOnly{}
	return &this
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ProgramInputOnly) GetTags() []int32 {
	if o == nil || isNil(o.Tags) {
		var ret []int32
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramInputOnly) GetTagsOk() ([]int32, bool) {
	if o == nil || isNil(o.Tags) {
    return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ProgramInputOnly) HasTags() bool {
	if o != nil && !isNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []int32 and assigns it to the Tags field.
func (o *ProgramInputOnly) SetTags(v []int32) {
	o.Tags = v
}

// GetPresenters returns the Presenters field value if set, zero value otherwise.
func (o *ProgramInputOnly) GetPresenters() []int32 {
	if o == nil || isNil(o.Presenters) {
		var ret []int32
		return ret
	}
	return o.Presenters
}

// GetPresentersOk returns a tuple with the Presenters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramInputOnly) GetPresentersOk() ([]int32, bool) {
	if o == nil || isNil(o.Presenters) {
    return nil, false
	}
	return o.Presenters, true
}

// HasPresenters returns a boolean if a field has been set.
func (o *ProgramInputOnly) HasPresenters() bool {
	if o != nil && !isNil(o.Presenters) {
		return true
	}

	return false
}

// SetPresenters gets a reference to the given []int32 and assigns it to the Presenters field.
func (o *ProgramInputOnly) SetPresenters(v []int32) {
	o.Presenters = v
}

func (o ProgramInputOnly) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !isNil(o.Presenters) {
		toSerialize["presenters"] = o.Presenters
	}
	return json.Marshal(toSerialize)
}

type NullableProgramInputOnly struct {
	value *ProgramInputOnly
	isSet bool
}

func (v NullableProgramInputOnly) Get() *ProgramInputOnly {
	return v.value
}

func (v *NullableProgramInputOnly) Set(val *ProgramInputOnly) {
	v.value = val
	v.isSet = true
}

func (v NullableProgramInputOnly) IsSet() bool {
	return v.isSet
}

func (v *NullableProgramInputOnly) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProgramInputOnly(val *ProgramInputOnly) *NullableProgramInputOnly {
	return &NullableProgramInputOnly{value: val, isSet: true}
}

func (v NullableProgramInputOnly) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProgramInputOnly) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


