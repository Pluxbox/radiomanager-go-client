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

// GenreResult struct for GenreResult
type GenreResult struct {
	Id int64 `json:"id"`
	Urn *string `json:"urn,omitempty"`
	ParentId *int64 `json:"parent_id,omitempty"`
	Name string `json:"name"`
	Broadcasts *GenreRelationsBroadcasts `json:"broadcasts,omitempty"`
	Programs *GenreRelationsPrograms `json:"programs,omitempty"`
	ExternalStationId *int64 `json:"_external_station_id,omitempty"`
}

// NewGenreResult instantiates a new GenreResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenreResult(id int64, name string) *GenreResult {
	this := GenreResult{}
	this.Id = id
	this.Name = name
	return &this
}

// NewGenreResultWithDefaults instantiates a new GenreResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenreResultWithDefaults() *GenreResult {
	this := GenreResult{}
	return &this
}

// GetId returns the Id field value
func (o *GenreResult) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GenreResult) GetIdOk() (*int64, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GenreResult) SetId(v int64) {
	o.Id = v
}

// GetUrn returns the Urn field value if set, zero value otherwise.
func (o *GenreResult) GetUrn() string {
	if o == nil || isNil(o.Urn) {
		var ret string
		return ret
	}
	return *o.Urn
}

// GetUrnOk returns a tuple with the Urn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenreResult) GetUrnOk() (*string, bool) {
	if o == nil || isNil(o.Urn) {
    return nil, false
	}
	return o.Urn, true
}

// HasUrn returns a boolean if a field has been set.
func (o *GenreResult) HasUrn() bool {
	if o != nil && !isNil(o.Urn) {
		return true
	}

	return false
}

// SetUrn gets a reference to the given string and assigns it to the Urn field.
func (o *GenreResult) SetUrn(v string) {
	o.Urn = &v
}

// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *GenreResult) GetParentId() int64 {
	if o == nil || isNil(o.ParentId) {
		var ret int64
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenreResult) GetParentIdOk() (*int64, bool) {
	if o == nil || isNil(o.ParentId) {
    return nil, false
	}
	return o.ParentId, true
}

// HasParentId returns a boolean if a field has been set.
func (o *GenreResult) HasParentId() bool {
	if o != nil && !isNil(o.ParentId) {
		return true
	}

	return false
}

// SetParentId gets a reference to the given int64 and assigns it to the ParentId field.
func (o *GenreResult) SetParentId(v int64) {
	o.ParentId = &v
}

// GetName returns the Name field value
func (o *GenreResult) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GenreResult) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GenreResult) SetName(v string) {
	o.Name = v
}

// GetBroadcasts returns the Broadcasts field value if set, zero value otherwise.
func (o *GenreResult) GetBroadcasts() GenreRelationsBroadcasts {
	if o == nil || isNil(o.Broadcasts) {
		var ret GenreRelationsBroadcasts
		return ret
	}
	return *o.Broadcasts
}

// GetBroadcastsOk returns a tuple with the Broadcasts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenreResult) GetBroadcastsOk() (*GenreRelationsBroadcasts, bool) {
	if o == nil || isNil(o.Broadcasts) {
    return nil, false
	}
	return o.Broadcasts, true
}

// HasBroadcasts returns a boolean if a field has been set.
func (o *GenreResult) HasBroadcasts() bool {
	if o != nil && !isNil(o.Broadcasts) {
		return true
	}

	return false
}

// SetBroadcasts gets a reference to the given GenreRelationsBroadcasts and assigns it to the Broadcasts field.
func (o *GenreResult) SetBroadcasts(v GenreRelationsBroadcasts) {
	o.Broadcasts = &v
}

// GetPrograms returns the Programs field value if set, zero value otherwise.
func (o *GenreResult) GetPrograms() GenreRelationsPrograms {
	if o == nil || isNil(o.Programs) {
		var ret GenreRelationsPrograms
		return ret
	}
	return *o.Programs
}

// GetProgramsOk returns a tuple with the Programs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenreResult) GetProgramsOk() (*GenreRelationsPrograms, bool) {
	if o == nil || isNil(o.Programs) {
    return nil, false
	}
	return o.Programs, true
}

// HasPrograms returns a boolean if a field has been set.
func (o *GenreResult) HasPrograms() bool {
	if o != nil && !isNil(o.Programs) {
		return true
	}

	return false
}

// SetPrograms gets a reference to the given GenreRelationsPrograms and assigns it to the Programs field.
func (o *GenreResult) SetPrograms(v GenreRelationsPrograms) {
	o.Programs = &v
}

// GetExternalStationId returns the ExternalStationId field value if set, zero value otherwise.
func (o *GenreResult) GetExternalStationId() int64 {
	if o == nil || isNil(o.ExternalStationId) {
		var ret int64
		return ret
	}
	return *o.ExternalStationId
}

// GetExternalStationIdOk returns a tuple with the ExternalStationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenreResult) GetExternalStationIdOk() (*int64, bool) {
	if o == nil || isNil(o.ExternalStationId) {
    return nil, false
	}
	return o.ExternalStationId, true
}

// HasExternalStationId returns a boolean if a field has been set.
func (o *GenreResult) HasExternalStationId() bool {
	if o != nil && !isNil(o.ExternalStationId) {
		return true
	}

	return false
}

// SetExternalStationId gets a reference to the given int64 and assigns it to the ExternalStationId field.
func (o *GenreResult) SetExternalStationId(v int64) {
	o.ExternalStationId = &v
}

func (o GenreResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Urn) {
		toSerialize["urn"] = o.Urn
	}
	if !isNil(o.ParentId) {
		toSerialize["parent_id"] = o.ParentId
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Broadcasts) {
		toSerialize["broadcasts"] = o.Broadcasts
	}
	if !isNil(o.Programs) {
		toSerialize["programs"] = o.Programs
	}
	if !isNil(o.ExternalStationId) {
		toSerialize["_external_station_id"] = o.ExternalStationId
	}
	return json.Marshal(toSerialize)
}

type NullableGenreResult struct {
	value *GenreResult
	isSet bool
}

func (v NullableGenreResult) Get() *GenreResult {
	return v.value
}

func (v *NullableGenreResult) Set(val *GenreResult) {
	v.value = val
	v.isSet = true
}

func (v NullableGenreResult) IsSet() bool {
	return v.isSet
}

func (v *NullableGenreResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenreResult(val *GenreResult) *NullableGenreResult {
	return &NullableGenreResult{value: val, isSet: true}
}

func (v NullableGenreResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenreResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


