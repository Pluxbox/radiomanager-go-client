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

// PlaylistTimingBody struct for PlaylistTimingBody
type PlaylistTimingBody struct {
	Start *time.Time `json:"start,omitempty"`
	AllowPlaylistPast *int32 `json:"allow_playlist_past,omitempty"`
	BroadcastId *int64 `json:"broadcast_id,omitempty"`
	Items []ImportItem `json:"items,omitempty"`
}

// NewPlaylistTimingBody instantiates a new PlaylistTimingBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlaylistTimingBody() *PlaylistTimingBody {
	this := PlaylistTimingBody{}
	return &this
}

// NewPlaylistTimingBodyWithDefaults instantiates a new PlaylistTimingBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlaylistTimingBodyWithDefaults() *PlaylistTimingBody {
	this := PlaylistTimingBody{}
	return &this
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *PlaylistTimingBody) GetStart() time.Time {
	if o == nil || isNil(o.Start) {
		var ret time.Time
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaylistTimingBody) GetStartOk() (*time.Time, bool) {
	if o == nil || isNil(o.Start) {
    return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *PlaylistTimingBody) HasStart() bool {
	if o != nil && !isNil(o.Start) {
		return true
	}

	return false
}

// SetStart gets a reference to the given time.Time and assigns it to the Start field.
func (o *PlaylistTimingBody) SetStart(v time.Time) {
	o.Start = &v
}

// GetAllowPlaylistPast returns the AllowPlaylistPast field value if set, zero value otherwise.
func (o *PlaylistTimingBody) GetAllowPlaylistPast() int32 {
	if o == nil || isNil(o.AllowPlaylistPast) {
		var ret int32
		return ret
	}
	return *o.AllowPlaylistPast
}

// GetAllowPlaylistPastOk returns a tuple with the AllowPlaylistPast field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaylistTimingBody) GetAllowPlaylistPastOk() (*int32, bool) {
	if o == nil || isNil(o.AllowPlaylistPast) {
    return nil, false
	}
	return o.AllowPlaylistPast, true
}

// HasAllowPlaylistPast returns a boolean if a field has been set.
func (o *PlaylistTimingBody) HasAllowPlaylistPast() bool {
	if o != nil && !isNil(o.AllowPlaylistPast) {
		return true
	}

	return false
}

// SetAllowPlaylistPast gets a reference to the given int32 and assigns it to the AllowPlaylistPast field.
func (o *PlaylistTimingBody) SetAllowPlaylistPast(v int32) {
	o.AllowPlaylistPast = &v
}

// GetBroadcastId returns the BroadcastId field value if set, zero value otherwise.
func (o *PlaylistTimingBody) GetBroadcastId() int64 {
	if o == nil || isNil(o.BroadcastId) {
		var ret int64
		return ret
	}
	return *o.BroadcastId
}

// GetBroadcastIdOk returns a tuple with the BroadcastId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaylistTimingBody) GetBroadcastIdOk() (*int64, bool) {
	if o == nil || isNil(o.BroadcastId) {
    return nil, false
	}
	return o.BroadcastId, true
}

// HasBroadcastId returns a boolean if a field has been set.
func (o *PlaylistTimingBody) HasBroadcastId() bool {
	if o != nil && !isNil(o.BroadcastId) {
		return true
	}

	return false
}

// SetBroadcastId gets a reference to the given int64 and assigns it to the BroadcastId field.
func (o *PlaylistTimingBody) SetBroadcastId(v int64) {
	o.BroadcastId = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *PlaylistTimingBody) GetItems() []ImportItem {
	if o == nil || isNil(o.Items) {
		var ret []ImportItem
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaylistTimingBody) GetItemsOk() ([]ImportItem, bool) {
	if o == nil || isNil(o.Items) {
    return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *PlaylistTimingBody) HasItems() bool {
	if o != nil && !isNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []ImportItem and assigns it to the Items field.
func (o *PlaylistTimingBody) SetItems(v []ImportItem) {
	o.Items = v
}

func (o PlaylistTimingBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Start) {
		toSerialize["start"] = o.Start
	}
	if !isNil(o.AllowPlaylistPast) {
		toSerialize["allow_playlist_past"] = o.AllowPlaylistPast
	}
	if !isNil(o.BroadcastId) {
		toSerialize["broadcast_id"] = o.BroadcastId
	}
	if !isNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullablePlaylistTimingBody struct {
	value *PlaylistTimingBody
	isSet bool
}

func (v NullablePlaylistTimingBody) Get() *PlaylistTimingBody {
	return v.value
}

func (v *NullablePlaylistTimingBody) Set(val *PlaylistTimingBody) {
	v.value = val
	v.isSet = true
}

func (v NullablePlaylistTimingBody) IsSet() bool {
	return v.isSet
}

func (v *NullablePlaylistTimingBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlaylistTimingBody(val *PlaylistTimingBody) *NullablePlaylistTimingBody {
	return &NullablePlaylistTimingBody{value: val, isSet: true}
}

func (v NullablePlaylistTimingBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlaylistTimingBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


