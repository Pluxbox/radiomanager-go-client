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

// ModelTypeRelationsPresenters struct for ModelTypeRelationsPresenters
type ModelTypeRelationsPresenters struct {
	Href *string `json:"href,omitempty"`
	Model *string `json:"model,omitempty"`
	Operation *string `json:"operation,omitempty"`
	Params *ModelTypeRelationsCampaignsParams `json:"params,omitempty"`
}

// NewModelTypeRelationsPresenters instantiates a new ModelTypeRelationsPresenters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelTypeRelationsPresenters() *ModelTypeRelationsPresenters {
	this := ModelTypeRelationsPresenters{}
	return &this
}

// NewModelTypeRelationsPresentersWithDefaults instantiates a new ModelTypeRelationsPresenters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelTypeRelationsPresentersWithDefaults() *ModelTypeRelationsPresenters {
	this := ModelTypeRelationsPresenters{}
	return &this
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *ModelTypeRelationsPresenters) GetHref() string {
	if o == nil || isNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelTypeRelationsPresenters) GetHrefOk() (*string, bool) {
	if o == nil || isNil(o.Href) {
    return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *ModelTypeRelationsPresenters) HasHref() bool {
	if o != nil && !isNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *ModelTypeRelationsPresenters) SetHref(v string) {
	o.Href = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *ModelTypeRelationsPresenters) GetModel() string {
	if o == nil || isNil(o.Model) {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelTypeRelationsPresenters) GetModelOk() (*string, bool) {
	if o == nil || isNil(o.Model) {
    return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *ModelTypeRelationsPresenters) HasModel() bool {
	if o != nil && !isNil(o.Model) {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *ModelTypeRelationsPresenters) SetModel(v string) {
	o.Model = &v
}

// GetOperation returns the Operation field value if set, zero value otherwise.
func (o *ModelTypeRelationsPresenters) GetOperation() string {
	if o == nil || isNil(o.Operation) {
		var ret string
		return ret
	}
	return *o.Operation
}

// GetOperationOk returns a tuple with the Operation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelTypeRelationsPresenters) GetOperationOk() (*string, bool) {
	if o == nil || isNil(o.Operation) {
    return nil, false
	}
	return o.Operation, true
}

// HasOperation returns a boolean if a field has been set.
func (o *ModelTypeRelationsPresenters) HasOperation() bool {
	if o != nil && !isNil(o.Operation) {
		return true
	}

	return false
}

// SetOperation gets a reference to the given string and assigns it to the Operation field.
func (o *ModelTypeRelationsPresenters) SetOperation(v string) {
	o.Operation = &v
}

// GetParams returns the Params field value if set, zero value otherwise.
func (o *ModelTypeRelationsPresenters) GetParams() ModelTypeRelationsCampaignsParams {
	if o == nil || isNil(o.Params) {
		var ret ModelTypeRelationsCampaignsParams
		return ret
	}
	return *o.Params
}

// GetParamsOk returns a tuple with the Params field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelTypeRelationsPresenters) GetParamsOk() (*ModelTypeRelationsCampaignsParams, bool) {
	if o == nil || isNil(o.Params) {
    return nil, false
	}
	return o.Params, true
}

// HasParams returns a boolean if a field has been set.
func (o *ModelTypeRelationsPresenters) HasParams() bool {
	if o != nil && !isNil(o.Params) {
		return true
	}

	return false
}

// SetParams gets a reference to the given ModelTypeRelationsCampaignsParams and assigns it to the Params field.
func (o *ModelTypeRelationsPresenters) SetParams(v ModelTypeRelationsCampaignsParams) {
	o.Params = &v
}

func (o ModelTypeRelationsPresenters) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	if !isNil(o.Model) {
		toSerialize["model"] = o.Model
	}
	if !isNil(o.Operation) {
		toSerialize["operation"] = o.Operation
	}
	if !isNil(o.Params) {
		toSerialize["params"] = o.Params
	}
	return json.Marshal(toSerialize)
}

type NullableModelTypeRelationsPresenters struct {
	value *ModelTypeRelationsPresenters
	isSet bool
}

func (v NullableModelTypeRelationsPresenters) Get() *ModelTypeRelationsPresenters {
	return v.value
}

func (v *NullableModelTypeRelationsPresenters) Set(val *ModelTypeRelationsPresenters) {
	v.value = val
	v.isSet = true
}

func (v NullableModelTypeRelationsPresenters) IsSet() bool {
	return v.isSet
}

func (v *NullableModelTypeRelationsPresenters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelTypeRelationsPresenters(val *ModelTypeRelationsPresenters) *NullableModelTypeRelationsPresenters {
	return &NullableModelTypeRelationsPresenters{value: val, isSet: true}
}

func (v NullableModelTypeRelationsPresenters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelTypeRelationsPresenters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

