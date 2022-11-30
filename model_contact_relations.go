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

// ContactRelations struct for ContactRelations
type ContactRelations struct {
	Tags ContactRelationsTags `json:"tags"`
	Items *ContactRelationsItems `json:"items,omitempty"`
	ModelType *BroadcastRelationsModelType `json:"model_type,omitempty"`
}

// NewContactRelations instantiates a new ContactRelations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContactRelations(tags ContactRelationsTags) *ContactRelations {
	this := ContactRelations{}
	this.Tags = tags
	return &this
}

// NewContactRelationsWithDefaults instantiates a new ContactRelations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContactRelationsWithDefaults() *ContactRelations {
	this := ContactRelations{}
	return &this
}

// GetTags returns the Tags field value
func (o *ContactRelations) GetTags() ContactRelationsTags {
	if o == nil {
		var ret ContactRelationsTags
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *ContactRelations) GetTagsOk() (*ContactRelationsTags, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Tags, true
}

// SetTags sets field value
func (o *ContactRelations) SetTags(v ContactRelationsTags) {
	o.Tags = v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ContactRelations) GetItems() ContactRelationsItems {
	if o == nil || isNil(o.Items) {
		var ret ContactRelationsItems
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactRelations) GetItemsOk() (*ContactRelationsItems, bool) {
	if o == nil || isNil(o.Items) {
    return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ContactRelations) HasItems() bool {
	if o != nil && !isNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given ContactRelationsItems and assigns it to the Items field.
func (o *ContactRelations) SetItems(v ContactRelationsItems) {
	o.Items = &v
}

// GetModelType returns the ModelType field value if set, zero value otherwise.
func (o *ContactRelations) GetModelType() BroadcastRelationsModelType {
	if o == nil || isNil(o.ModelType) {
		var ret BroadcastRelationsModelType
		return ret
	}
	return *o.ModelType
}

// GetModelTypeOk returns a tuple with the ModelType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactRelations) GetModelTypeOk() (*BroadcastRelationsModelType, bool) {
	if o == nil || isNil(o.ModelType) {
    return nil, false
	}
	return o.ModelType, true
}

// HasModelType returns a boolean if a field has been set.
func (o *ContactRelations) HasModelType() bool {
	if o != nil && !isNil(o.ModelType) {
		return true
	}

	return false
}

// SetModelType gets a reference to the given BroadcastRelationsModelType and assigns it to the ModelType field.
func (o *ContactRelations) SetModelType(v BroadcastRelationsModelType) {
	o.ModelType = &v
}

func (o ContactRelations) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["tags"] = o.Tags
	}
	if !isNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !isNil(o.ModelType) {
		toSerialize["model_type"] = o.ModelType
	}
	return json.Marshal(toSerialize)
}

type NullableContactRelations struct {
	value *ContactRelations
	isSet bool
}

func (v NullableContactRelations) Get() *ContactRelations {
	return v.value
}

func (v *NullableContactRelations) Set(val *ContactRelations) {
	v.value = val
	v.isSet = true
}

func (v NullableContactRelations) IsSet() bool {
	return v.isSet
}

func (v *NullableContactRelations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContactRelations(val *ContactRelations) *NullableContactRelations {
	return &NullableContactRelations{value: val, isSet: true}
}

func (v NullableContactRelations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContactRelations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


