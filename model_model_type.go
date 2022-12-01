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

// ModelType struct for ModelType
type ModelType struct {
	Id int64 `json:"id"`
	Name *string `json:"name,omitempty"`
	Model *string `json:"model,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	Options *ModelTypeOptions `json:"options,omitempty"`
	Order *int64 `json:"order,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

// NewModelType instantiates a new ModelType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelType(id int64) *ModelType {
	this := ModelType{}
	this.Id = id
	return &this
}

// NewModelTypeWithDefaults instantiates a new ModelType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelTypeWithDefaults() *ModelType {
	this := ModelType{}
	return &this
}

// GetId returns the Id field value
func (o *ModelType) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ModelType) GetIdOk() (*int64, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ModelType) SetId(v int64) {
	o.Id = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ModelType) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelType) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ModelType) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ModelType) SetName(v string) {
	o.Name = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *ModelType) GetModel() string {
	if o == nil || isNil(o.Model) {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelType) GetModelOk() (*string, bool) {
	if o == nil || isNil(o.Model) {
    return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *ModelType) HasModel() bool {
	if o != nil && !isNil(o.Model) {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *ModelType) SetModel(v string) {
	o.Model = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ModelType) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelType) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ModelType) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ModelType) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ModelType) GetUpdatedAt() time.Time {
	if o == nil || isNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelType) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.UpdatedAt) {
    return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ModelType) HasUpdatedAt() bool {
	if o != nil && !isNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *ModelType) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *ModelType) GetOptions() ModelTypeOptions {
	if o == nil || isNil(o.Options) {
		var ret ModelTypeOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelType) GetOptionsOk() (*ModelTypeOptions, bool) {
	if o == nil || isNil(o.Options) {
    return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *ModelType) HasOptions() bool {
	if o != nil && !isNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given ModelTypeOptions and assigns it to the Options field.
func (o *ModelType) SetOptions(v ModelTypeOptions) {
	o.Options = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *ModelType) GetOrder() int64 {
	if o == nil || isNil(o.Order) {
		var ret int64
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelType) GetOrderOk() (*int64, bool) {
	if o == nil || isNil(o.Order) {
    return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *ModelType) HasOrder() bool {
	if o != nil && !isNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given int64 and assigns it to the Order field.
func (o *ModelType) SetOrder(v int64) {
	o.Order = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *ModelType) GetDeletedAt() time.Time {
	if o == nil || isNil(o.DeletedAt) {
		var ret time.Time
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelType) GetDeletedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.DeletedAt) {
    return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *ModelType) HasDeletedAt() bool {
	if o != nil && !isNil(o.DeletedAt) {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given time.Time and assigns it to the DeletedAt field.
func (o *ModelType) SetDeletedAt(v time.Time) {
	o.DeletedAt = &v
}

func (o ModelType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Model) {
		toSerialize["model"] = o.Model
	}
	if !isNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !isNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !isNil(o.Options) {
		toSerialize["options"] = o.Options
	}
	if !isNil(o.Order) {
		toSerialize["order"] = o.Order
	}
	if !isNil(o.DeletedAt) {
		toSerialize["deleted_at"] = o.DeletedAt
	}
	return json.Marshal(toSerialize)
}

type NullableModelType struct {
	value *ModelType
	isSet bool
}

func (v NullableModelType) Get() *ModelType {
	return v.value
}

func (v *NullableModelType) Set(val *ModelType) {
	v.value = val
	v.isSet = true
}

func (v NullableModelType) IsSet() bool {
	return v.isSet
}

func (v *NullableModelType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelType(val *ModelType) *NullableModelType {
	return &NullableModelType{value: val, isSet: true}
}

func (v NullableModelType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

