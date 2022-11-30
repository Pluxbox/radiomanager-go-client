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

// GroupDataInput struct for GroupDataInput
type GroupDataInput struct {
	Id int64 `json:"id"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
	Title *string `json:"title,omitempty"`
}

// NewGroupDataInput instantiates a new GroupDataInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupDataInput(id int64, updatedAt time.Time, createdAt time.Time) *GroupDataInput {
	this := GroupDataInput{}
	this.Id = id
	this.UpdatedAt = updatedAt
	this.CreatedAt = createdAt
	return &this
}

// NewGroupDataInputWithDefaults instantiates a new GroupDataInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupDataInputWithDefaults() *GroupDataInput {
	this := GroupDataInput{}
	return &this
}

// GetId returns the Id field value
func (o *GroupDataInput) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GroupDataInput) GetIdOk() (*int64, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GroupDataInput) SetId(v int64) {
	o.Id = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *GroupDataInput) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *GroupDataInput) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
    return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *GroupDataInput) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *GroupDataInput) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *GroupDataInput) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
    return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *GroupDataInput) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *GroupDataInput) GetTitle() string {
	if o == nil || isNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupDataInput) GetTitleOk() (*string, bool) {
	if o == nil || isNil(o.Title) {
    return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *GroupDataInput) HasTitle() bool {
	if o != nil && !isNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *GroupDataInput) SetTitle(v string) {
	o.Title = &v
}

func (o GroupDataInput) MarshalJSON() ([]byte, error) {
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
	if !isNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	return json.Marshal(toSerialize)
}

type NullableGroupDataInput struct {
	value *GroupDataInput
	isSet bool
}

func (v NullableGroupDataInput) Get() *GroupDataInput {
	return v.value
}

func (v *NullableGroupDataInput) Set(val *GroupDataInput) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupDataInput) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupDataInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupDataInput(val *GroupDataInput) *NullableGroupDataInput {
	return &NullableGroupDataInput{value: val, isSet: true}
}

func (v NullableGroupDataInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupDataInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


