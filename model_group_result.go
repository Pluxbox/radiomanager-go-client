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

// GroupResult struct for GroupResult
type GroupResult struct {
	Id int64 `json:"id"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
	DeletedAt time.Time `json:"deleted_at"`
	Title *string `json:"title,omitempty"`
	Users *GroupRelationsUsers `json:"users,omitempty"`
	Broadcasts *GroupRelationsBroadcasts `json:"broadcasts,omitempty"`
	Programs *GroupRelationsPrograms `json:"programs,omitempty"`
}

// NewGroupResult instantiates a new GroupResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupResult(id int64, updatedAt time.Time, createdAt time.Time, deletedAt time.Time) *GroupResult {
	this := GroupResult{}
	this.Id = id
	this.UpdatedAt = updatedAt
	this.CreatedAt = createdAt
	this.DeletedAt = deletedAt
	return &this
}

// NewGroupResultWithDefaults instantiates a new GroupResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupResultWithDefaults() *GroupResult {
	this := GroupResult{}
	return &this
}

// GetId returns the Id field value
func (o *GroupResult) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GroupResult) GetIdOk() (*int64, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GroupResult) SetId(v int64) {
	o.Id = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *GroupResult) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *GroupResult) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
    return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *GroupResult) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *GroupResult) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *GroupResult) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
    return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *GroupResult) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetDeletedAt returns the DeletedAt field value
func (o *GroupResult) GetDeletedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value
// and a boolean to check if the value has been set.
func (o *GroupResult) GetDeletedAtOk() (*time.Time, bool) {
	if o == nil {
    return nil, false
	}
	return &o.DeletedAt, true
}

// SetDeletedAt sets field value
func (o *GroupResult) SetDeletedAt(v time.Time) {
	o.DeletedAt = v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *GroupResult) GetTitle() string {
	if o == nil || isNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupResult) GetTitleOk() (*string, bool) {
	if o == nil || isNil(o.Title) {
    return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *GroupResult) HasTitle() bool {
	if o != nil && !isNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *GroupResult) SetTitle(v string) {
	o.Title = &v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *GroupResult) GetUsers() GroupRelationsUsers {
	if o == nil || isNil(o.Users) {
		var ret GroupRelationsUsers
		return ret
	}
	return *o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupResult) GetUsersOk() (*GroupRelationsUsers, bool) {
	if o == nil || isNil(o.Users) {
    return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *GroupResult) HasUsers() bool {
	if o != nil && !isNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given GroupRelationsUsers and assigns it to the Users field.
func (o *GroupResult) SetUsers(v GroupRelationsUsers) {
	o.Users = &v
}

// GetBroadcasts returns the Broadcasts field value if set, zero value otherwise.
func (o *GroupResult) GetBroadcasts() GroupRelationsBroadcasts {
	if o == nil || isNil(o.Broadcasts) {
		var ret GroupRelationsBroadcasts
		return ret
	}
	return *o.Broadcasts
}

// GetBroadcastsOk returns a tuple with the Broadcasts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupResult) GetBroadcastsOk() (*GroupRelationsBroadcasts, bool) {
	if o == nil || isNil(o.Broadcasts) {
    return nil, false
	}
	return o.Broadcasts, true
}

// HasBroadcasts returns a boolean if a field has been set.
func (o *GroupResult) HasBroadcasts() bool {
	if o != nil && !isNil(o.Broadcasts) {
		return true
	}

	return false
}

// SetBroadcasts gets a reference to the given GroupRelationsBroadcasts and assigns it to the Broadcasts field.
func (o *GroupResult) SetBroadcasts(v GroupRelationsBroadcasts) {
	o.Broadcasts = &v
}

// GetPrograms returns the Programs field value if set, zero value otherwise.
func (o *GroupResult) GetPrograms() GroupRelationsPrograms {
	if o == nil || isNil(o.Programs) {
		var ret GroupRelationsPrograms
		return ret
	}
	return *o.Programs
}

// GetProgramsOk returns a tuple with the Programs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupResult) GetProgramsOk() (*GroupRelationsPrograms, bool) {
	if o == nil || isNil(o.Programs) {
    return nil, false
	}
	return o.Programs, true
}

// HasPrograms returns a boolean if a field has been set.
func (o *GroupResult) HasPrograms() bool {
	if o != nil && !isNil(o.Programs) {
		return true
	}

	return false
}

// SetPrograms gets a reference to the given GroupRelationsPrograms and assigns it to the Programs field.
func (o *GroupResult) SetPrograms(v GroupRelationsPrograms) {
	o.Programs = &v
}

func (o GroupResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["deleted_at"] = o.DeletedAt
	}
	if !isNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !isNil(o.Users) {
		toSerialize["users"] = o.Users
	}
	if !isNil(o.Broadcasts) {
		toSerialize["broadcasts"] = o.Broadcasts
	}
	if !isNil(o.Programs) {
		toSerialize["programs"] = o.Programs
	}
	return json.Marshal(toSerialize)
}

type NullableGroupResult struct {
	value *GroupResult
	isSet bool
}

func (v NullableGroupResult) Get() *GroupResult {
	return v.value
}

func (v *NullableGroupResult) Set(val *GroupResult) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupResult) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupResult(val *GroupResult) *NullableGroupResult {
	return &NullableGroupResult{value: val, isSet: true}
}

func (v NullableGroupResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


