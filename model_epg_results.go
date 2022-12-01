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

// EPGResults struct for EPGResults
type EPGResults struct {
	Days []BroadcastEPGDay `json:"days"`
	NextPageUrl string `json:"next_page_url"`
	PrevPageUrl string `json:"prev_page_url"`
}

// NewEPGResults instantiates a new EPGResults object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEPGResults(days []BroadcastEPGDay, nextPageUrl string, prevPageUrl string) *EPGResults {
	this := EPGResults{}
	this.Days = days
	this.NextPageUrl = nextPageUrl
	this.PrevPageUrl = prevPageUrl
	return &this
}

// NewEPGResultsWithDefaults instantiates a new EPGResults object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEPGResultsWithDefaults() *EPGResults {
	this := EPGResults{}
	return &this
}

// GetDays returns the Days field value
func (o *EPGResults) GetDays() []BroadcastEPGDay {
	if o == nil {
		var ret []BroadcastEPGDay
		return ret
	}

	return o.Days
}

// GetDaysOk returns a tuple with the Days field value
// and a boolean to check if the value has been set.
func (o *EPGResults) GetDaysOk() ([]BroadcastEPGDay, bool) {
	if o == nil {
    return nil, false
	}
	return o.Days, true
}

// SetDays sets field value
func (o *EPGResults) SetDays(v []BroadcastEPGDay) {
	o.Days = v
}

// GetNextPageUrl returns the NextPageUrl field value
func (o *EPGResults) GetNextPageUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NextPageUrl
}

// GetNextPageUrlOk returns a tuple with the NextPageUrl field value
// and a boolean to check if the value has been set.
func (o *EPGResults) GetNextPageUrlOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.NextPageUrl, true
}

// SetNextPageUrl sets field value
func (o *EPGResults) SetNextPageUrl(v string) {
	o.NextPageUrl = v
}

// GetPrevPageUrl returns the PrevPageUrl field value
func (o *EPGResults) GetPrevPageUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PrevPageUrl
}

// GetPrevPageUrlOk returns a tuple with the PrevPageUrl field value
// and a boolean to check if the value has been set.
func (o *EPGResults) GetPrevPageUrlOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.PrevPageUrl, true
}

// SetPrevPageUrl sets field value
func (o *EPGResults) SetPrevPageUrl(v string) {
	o.PrevPageUrl = v
}

func (o EPGResults) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["days"] = o.Days
	}
	if true {
		toSerialize["next_page_url"] = o.NextPageUrl
	}
	if true {
		toSerialize["prev_page_url"] = o.PrevPageUrl
	}
	return json.Marshal(toSerialize)
}

type NullableEPGResults struct {
	value *EPGResults
	isSet bool
}

func (v NullableEPGResults) Get() *EPGResults {
	return v.value
}

func (v *NullableEPGResults) Set(val *EPGResults) {
	v.value = val
	v.isSet = true
}

func (v NullableEPGResults) IsSet() bool {
	return v.isSet
}

func (v *NullableEPGResults) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEPGResults(val *EPGResults) *NullableEPGResults {
	return &NullableEPGResults{value: val, isSet: true}
}

func (v NullableEPGResults) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEPGResults) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

