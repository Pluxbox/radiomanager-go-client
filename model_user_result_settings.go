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

// UserResultSettings struct for UserResultSettings
type UserResultSettings struct {
	ShowSideBar *bool `json:"showSideBar,omitempty"`
	ShowSocialBar *bool `json:"showSocialBar,omitempty"`
	ShowCheckboxColumn *bool `json:"showCheckboxColumn,omitempty"`
	ShowTimeColumn *bool `json:"showTimeColumn,omitempty"`
	ShowSpeechTime *bool `json:"showSpeechTime,omitempty"`
	ZoomFactor *int64 `json:"zoomFactor,omitempty"`
}

// NewUserResultSettings instantiates a new UserResultSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserResultSettings() *UserResultSettings {
	this := UserResultSettings{}
	return &this
}

// NewUserResultSettingsWithDefaults instantiates a new UserResultSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserResultSettingsWithDefaults() *UserResultSettings {
	this := UserResultSettings{}
	return &this
}

// GetShowSideBar returns the ShowSideBar field value if set, zero value otherwise.
func (o *UserResultSettings) GetShowSideBar() bool {
	if o == nil || isNil(o.ShowSideBar) {
		var ret bool
		return ret
	}
	return *o.ShowSideBar
}

// GetShowSideBarOk returns a tuple with the ShowSideBar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResultSettings) GetShowSideBarOk() (*bool, bool) {
	if o == nil || isNil(o.ShowSideBar) {
    return nil, false
	}
	return o.ShowSideBar, true
}

// HasShowSideBar returns a boolean if a field has been set.
func (o *UserResultSettings) HasShowSideBar() bool {
	if o != nil && !isNil(o.ShowSideBar) {
		return true
	}

	return false
}

// SetShowSideBar gets a reference to the given bool and assigns it to the ShowSideBar field.
func (o *UserResultSettings) SetShowSideBar(v bool) {
	o.ShowSideBar = &v
}

// GetShowSocialBar returns the ShowSocialBar field value if set, zero value otherwise.
func (o *UserResultSettings) GetShowSocialBar() bool {
	if o == nil || isNil(o.ShowSocialBar) {
		var ret bool
		return ret
	}
	return *o.ShowSocialBar
}

// GetShowSocialBarOk returns a tuple with the ShowSocialBar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResultSettings) GetShowSocialBarOk() (*bool, bool) {
	if o == nil || isNil(o.ShowSocialBar) {
    return nil, false
	}
	return o.ShowSocialBar, true
}

// HasShowSocialBar returns a boolean if a field has been set.
func (o *UserResultSettings) HasShowSocialBar() bool {
	if o != nil && !isNil(o.ShowSocialBar) {
		return true
	}

	return false
}

// SetShowSocialBar gets a reference to the given bool and assigns it to the ShowSocialBar field.
func (o *UserResultSettings) SetShowSocialBar(v bool) {
	o.ShowSocialBar = &v
}

// GetShowCheckboxColumn returns the ShowCheckboxColumn field value if set, zero value otherwise.
func (o *UserResultSettings) GetShowCheckboxColumn() bool {
	if o == nil || isNil(o.ShowCheckboxColumn) {
		var ret bool
		return ret
	}
	return *o.ShowCheckboxColumn
}

// GetShowCheckboxColumnOk returns a tuple with the ShowCheckboxColumn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResultSettings) GetShowCheckboxColumnOk() (*bool, bool) {
	if o == nil || isNil(o.ShowCheckboxColumn) {
    return nil, false
	}
	return o.ShowCheckboxColumn, true
}

// HasShowCheckboxColumn returns a boolean if a field has been set.
func (o *UserResultSettings) HasShowCheckboxColumn() bool {
	if o != nil && !isNil(o.ShowCheckboxColumn) {
		return true
	}

	return false
}

// SetShowCheckboxColumn gets a reference to the given bool and assigns it to the ShowCheckboxColumn field.
func (o *UserResultSettings) SetShowCheckboxColumn(v bool) {
	o.ShowCheckboxColumn = &v
}

// GetShowTimeColumn returns the ShowTimeColumn field value if set, zero value otherwise.
func (o *UserResultSettings) GetShowTimeColumn() bool {
	if o == nil || isNil(o.ShowTimeColumn) {
		var ret bool
		return ret
	}
	return *o.ShowTimeColumn
}

// GetShowTimeColumnOk returns a tuple with the ShowTimeColumn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResultSettings) GetShowTimeColumnOk() (*bool, bool) {
	if o == nil || isNil(o.ShowTimeColumn) {
    return nil, false
	}
	return o.ShowTimeColumn, true
}

// HasShowTimeColumn returns a boolean if a field has been set.
func (o *UserResultSettings) HasShowTimeColumn() bool {
	if o != nil && !isNil(o.ShowTimeColumn) {
		return true
	}

	return false
}

// SetShowTimeColumn gets a reference to the given bool and assigns it to the ShowTimeColumn field.
func (o *UserResultSettings) SetShowTimeColumn(v bool) {
	o.ShowTimeColumn = &v
}

// GetShowSpeechTime returns the ShowSpeechTime field value if set, zero value otherwise.
func (o *UserResultSettings) GetShowSpeechTime() bool {
	if o == nil || isNil(o.ShowSpeechTime) {
		var ret bool
		return ret
	}
	return *o.ShowSpeechTime
}

// GetShowSpeechTimeOk returns a tuple with the ShowSpeechTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResultSettings) GetShowSpeechTimeOk() (*bool, bool) {
	if o == nil || isNil(o.ShowSpeechTime) {
    return nil, false
	}
	return o.ShowSpeechTime, true
}

// HasShowSpeechTime returns a boolean if a field has been set.
func (o *UserResultSettings) HasShowSpeechTime() bool {
	if o != nil && !isNil(o.ShowSpeechTime) {
		return true
	}

	return false
}

// SetShowSpeechTime gets a reference to the given bool and assigns it to the ShowSpeechTime field.
func (o *UserResultSettings) SetShowSpeechTime(v bool) {
	o.ShowSpeechTime = &v
}

// GetZoomFactor returns the ZoomFactor field value if set, zero value otherwise.
func (o *UserResultSettings) GetZoomFactor() int64 {
	if o == nil || isNil(o.ZoomFactor) {
		var ret int64
		return ret
	}
	return *o.ZoomFactor
}

// GetZoomFactorOk returns a tuple with the ZoomFactor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResultSettings) GetZoomFactorOk() (*int64, bool) {
	if o == nil || isNil(o.ZoomFactor) {
    return nil, false
	}
	return o.ZoomFactor, true
}

// HasZoomFactor returns a boolean if a field has been set.
func (o *UserResultSettings) HasZoomFactor() bool {
	if o != nil && !isNil(o.ZoomFactor) {
		return true
	}

	return false
}

// SetZoomFactor gets a reference to the given int64 and assigns it to the ZoomFactor field.
func (o *UserResultSettings) SetZoomFactor(v int64) {
	o.ZoomFactor = &v
}

func (o UserResultSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ShowSideBar) {
		toSerialize["showSideBar"] = o.ShowSideBar
	}
	if !isNil(o.ShowSocialBar) {
		toSerialize["showSocialBar"] = o.ShowSocialBar
	}
	if !isNil(o.ShowCheckboxColumn) {
		toSerialize["showCheckboxColumn"] = o.ShowCheckboxColumn
	}
	if !isNil(o.ShowTimeColumn) {
		toSerialize["showTimeColumn"] = o.ShowTimeColumn
	}
	if !isNil(o.ShowSpeechTime) {
		toSerialize["showSpeechTime"] = o.ShowSpeechTime
	}
	if !isNil(o.ZoomFactor) {
		toSerialize["zoomFactor"] = o.ZoomFactor
	}
	return json.Marshal(toSerialize)
}

type NullableUserResultSettings struct {
	value *UserResultSettings
	isSet bool
}

func (v NullableUserResultSettings) Get() *UserResultSettings {
	return v.value
}

func (v *NullableUserResultSettings) Set(val *UserResultSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableUserResultSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableUserResultSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserResultSettings(val *UserResultSettings) *NullableUserResultSettings {
	return &NullableUserResultSettings{value: val, isSet: true}
}

func (v NullableUserResultSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserResultSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

