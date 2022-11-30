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

// ProgramRelationsTags struct for ProgramRelationsTags
type ProgramRelationsTags struct {
	Href *string `json:"href,omitempty"`
	Model *string `json:"model,omitempty"`
	Operation *string `json:"operation,omitempty"`
	Params *ProgramRelationsItemsParams `json:"params,omitempty"`
}

// NewProgramRelationsTags instantiates a new ProgramRelationsTags object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProgramRelationsTags() *ProgramRelationsTags {
	this := ProgramRelationsTags{}
	return &this
}

// NewProgramRelationsTagsWithDefaults instantiates a new ProgramRelationsTags object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProgramRelationsTagsWithDefaults() *ProgramRelationsTags {
	this := ProgramRelationsTags{}
	return &this
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *ProgramRelationsTags) GetHref() string {
	if o == nil || isNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramRelationsTags) GetHrefOk() (*string, bool) {
	if o == nil || isNil(o.Href) {
    return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *ProgramRelationsTags) HasHref() bool {
	if o != nil && !isNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *ProgramRelationsTags) SetHref(v string) {
	o.Href = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *ProgramRelationsTags) GetModel() string {
	if o == nil || isNil(o.Model) {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramRelationsTags) GetModelOk() (*string, bool) {
	if o == nil || isNil(o.Model) {
    return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *ProgramRelationsTags) HasModel() bool {
	if o != nil && !isNil(o.Model) {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *ProgramRelationsTags) SetModel(v string) {
	o.Model = &v
}

// GetOperation returns the Operation field value if set, zero value otherwise.
func (o *ProgramRelationsTags) GetOperation() string {
	if o == nil || isNil(o.Operation) {
		var ret string
		return ret
	}
	return *o.Operation
}

// GetOperationOk returns a tuple with the Operation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramRelationsTags) GetOperationOk() (*string, bool) {
	if o == nil || isNil(o.Operation) {
    return nil, false
	}
	return o.Operation, true
}

// HasOperation returns a boolean if a field has been set.
func (o *ProgramRelationsTags) HasOperation() bool {
	if o != nil && !isNil(o.Operation) {
		return true
	}

	return false
}

// SetOperation gets a reference to the given string and assigns it to the Operation field.
func (o *ProgramRelationsTags) SetOperation(v string) {
	o.Operation = &v
}

// GetParams returns the Params field value if set, zero value otherwise.
func (o *ProgramRelationsTags) GetParams() ProgramRelationsItemsParams {
	if o == nil || isNil(o.Params) {
		var ret ProgramRelationsItemsParams
		return ret
	}
	return *o.Params
}

// GetParamsOk returns a tuple with the Params field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramRelationsTags) GetParamsOk() (*ProgramRelationsItemsParams, bool) {
	if o == nil || isNil(o.Params) {
    return nil, false
	}
	return o.Params, true
}

// HasParams returns a boolean if a field has been set.
func (o *ProgramRelationsTags) HasParams() bool {
	if o != nil && !isNil(o.Params) {
		return true
	}

	return false
}

// SetParams gets a reference to the given ProgramRelationsItemsParams and assigns it to the Params field.
func (o *ProgramRelationsTags) SetParams(v ProgramRelationsItemsParams) {
	o.Params = &v
}

func (o ProgramRelationsTags) MarshalJSON() ([]byte, error) {
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

type NullableProgramRelationsTags struct {
	value *ProgramRelationsTags
	isSet bool
}

func (v NullableProgramRelationsTags) Get() *ProgramRelationsTags {
	return v.value
}

func (v *NullableProgramRelationsTags) Set(val *ProgramRelationsTags) {
	v.value = val
	v.isSet = true
}

func (v NullableProgramRelationsTags) IsSet() bool {
	return v.isSet
}

func (v *NullableProgramRelationsTags) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProgramRelationsTags(val *ProgramRelationsTags) *NullableProgramRelationsTags {
	return &NullableProgramRelationsTags{value: val, isSet: true}
}

func (v NullableProgramRelationsTags) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProgramRelationsTags) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


