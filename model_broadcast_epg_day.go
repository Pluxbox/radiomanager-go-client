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

// BroadcastEPGDay struct for BroadcastEPGDay
type BroadcastEPGDay struct {
	Day *string `json:"day,omitempty"`
	Results []BroadcastEPGResult `json:"results"`
}

// NewBroadcastEPGDay instantiates a new BroadcastEPGDay object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBroadcastEPGDay(results []BroadcastEPGResult) *BroadcastEPGDay {
	this := BroadcastEPGDay{}
	this.Results = results
	return &this
}

// NewBroadcastEPGDayWithDefaults instantiates a new BroadcastEPGDay object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBroadcastEPGDayWithDefaults() *BroadcastEPGDay {
	this := BroadcastEPGDay{}
	return &this
}

// GetDay returns the Day field value if set, zero value otherwise.
func (o *BroadcastEPGDay) GetDay() string {
	if o == nil || isNil(o.Day) {
		var ret string
		return ret
	}
	return *o.Day
}

// GetDayOk returns a tuple with the Day field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BroadcastEPGDay) GetDayOk() (*string, bool) {
	if o == nil || isNil(o.Day) {
    return nil, false
	}
	return o.Day, true
}

// HasDay returns a boolean if a field has been set.
func (o *BroadcastEPGDay) HasDay() bool {
	if o != nil && !isNil(o.Day) {
		return true
	}

	return false
}

// SetDay gets a reference to the given string and assigns it to the Day field.
func (o *BroadcastEPGDay) SetDay(v string) {
	o.Day = &v
}

// GetResults returns the Results field value
func (o *BroadcastEPGDay) GetResults() []BroadcastEPGResult {
	if o == nil {
		var ret []BroadcastEPGResult
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *BroadcastEPGDay) GetResultsOk() ([]BroadcastEPGResult, bool) {
	if o == nil {
    return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *BroadcastEPGDay) SetResults(v []BroadcastEPGResult) {
	o.Results = v
}

func (o BroadcastEPGDay) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Day) {
		toSerialize["day"] = o.Day
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullableBroadcastEPGDay struct {
	value *BroadcastEPGDay
	isSet bool
}

func (v NullableBroadcastEPGDay) Get() *BroadcastEPGDay {
	return v.value
}

func (v *NullableBroadcastEPGDay) Set(val *BroadcastEPGDay) {
	v.value = val
	v.isSet = true
}

func (v NullableBroadcastEPGDay) IsSet() bool {
	return v.isSet
}

func (v *NullableBroadcastEPGDay) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBroadcastEPGDay(val *BroadcastEPGDay) *NullableBroadcastEPGDay {
	return &NullableBroadcastEPGDay{value: val, isSet: true}
}

func (v NullableBroadcastEPGDay) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBroadcastEPGDay) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


