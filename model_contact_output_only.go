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
	"time"
)

// ContactOutputOnly Contact
type ContactOutputOnly struct {
	Id *int64 `json:"id,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	ExternalStationId *int64 `json:"_external_station_id,omitempty"`
}

// NewContactOutputOnly instantiates a new ContactOutputOnly object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContactOutputOnly() *ContactOutputOnly {
	this := ContactOutputOnly{}
	return &this
}

// NewContactOutputOnlyWithDefaults instantiates a new ContactOutputOnly object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContactOutputOnlyWithDefaults() *ContactOutputOnly {
	this := ContactOutputOnly{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ContactOutputOnly) GetId() int64 {
	if o == nil || isNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactOutputOnly) GetIdOk() (*int64, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ContactOutputOnly) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *ContactOutputOnly) SetId(v int64) {
	o.Id = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ContactOutputOnly) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactOutputOnly) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ContactOutputOnly) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ContactOutputOnly) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ContactOutputOnly) GetUpdatedAt() time.Time {
	if o == nil || isNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactOutputOnly) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.UpdatedAt) {
    return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ContactOutputOnly) HasUpdatedAt() bool {
	if o != nil && !isNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *ContactOutputOnly) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *ContactOutputOnly) GetDeletedAt() time.Time {
	if o == nil || isNil(o.DeletedAt) {
		var ret time.Time
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactOutputOnly) GetDeletedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.DeletedAt) {
    return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *ContactOutputOnly) HasDeletedAt() bool {
	if o != nil && !isNil(o.DeletedAt) {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given time.Time and assigns it to the DeletedAt field.
func (o *ContactOutputOnly) SetDeletedAt(v time.Time) {
	o.DeletedAt = &v
}

// GetExternalStationId returns the ExternalStationId field value if set, zero value otherwise.
func (o *ContactOutputOnly) GetExternalStationId() int64 {
	if o == nil || isNil(o.ExternalStationId) {
		var ret int64
		return ret
	}
	return *o.ExternalStationId
}

// GetExternalStationIdOk returns a tuple with the ExternalStationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactOutputOnly) GetExternalStationIdOk() (*int64, bool) {
	if o == nil || isNil(o.ExternalStationId) {
    return nil, false
	}
	return o.ExternalStationId, true
}

// HasExternalStationId returns a boolean if a field has been set.
func (o *ContactOutputOnly) HasExternalStationId() bool {
	if o != nil && !isNil(o.ExternalStationId) {
		return true
	}

	return false
}

// SetExternalStationId gets a reference to the given int64 and assigns it to the ExternalStationId field.
func (o *ContactOutputOnly) SetExternalStationId(v int64) {
	o.ExternalStationId = &v
}

func (o ContactOutputOnly) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !isNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !isNil(o.DeletedAt) {
		toSerialize["deleted_at"] = o.DeletedAt
	}
	if !isNil(o.ExternalStationId) {
		toSerialize["_external_station_id"] = o.ExternalStationId
	}
	return json.Marshal(toSerialize)
}

type NullableContactOutputOnly struct {
	value *ContactOutputOnly
	isSet bool
}

func (v NullableContactOutputOnly) Get() *ContactOutputOnly {
	return v.value
}

func (v *NullableContactOutputOnly) Set(val *ContactOutputOnly) {
	v.value = val
	v.isSet = true
}

func (v NullableContactOutputOnly) IsSet() bool {
	return v.isSet
}

func (v *NullableContactOutputOnly) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContactOutputOnly(val *ContactOutputOnly) *NullableContactOutputOnly {
	return &NullableContactOutputOnly{value: val, isSet: true}
}

func (v NullableContactOutputOnly) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContactOutputOnly) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


