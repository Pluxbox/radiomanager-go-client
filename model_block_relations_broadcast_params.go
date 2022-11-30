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

// BlockRelationsBroadcastParams struct for BlockRelationsBroadcastParams
type BlockRelationsBroadcastParams struct {
	Id *int32 `json:"id,omitempty"`
}

// NewBlockRelationsBroadcastParams instantiates a new BlockRelationsBroadcastParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlockRelationsBroadcastParams() *BlockRelationsBroadcastParams {
	this := BlockRelationsBroadcastParams{}
	return &this
}

// NewBlockRelationsBroadcastParamsWithDefaults instantiates a new BlockRelationsBroadcastParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlockRelationsBroadcastParamsWithDefaults() *BlockRelationsBroadcastParams {
	this := BlockRelationsBroadcastParams{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BlockRelationsBroadcastParams) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockRelationsBroadcastParams) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BlockRelationsBroadcastParams) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *BlockRelationsBroadcastParams) SetId(v int32) {
	o.Id = &v
}

func (o BlockRelationsBroadcastParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableBlockRelationsBroadcastParams struct {
	value *BlockRelationsBroadcastParams
	isSet bool
}

func (v NullableBlockRelationsBroadcastParams) Get() *BlockRelationsBroadcastParams {
	return v.value
}

func (v *NullableBlockRelationsBroadcastParams) Set(val *BlockRelationsBroadcastParams) {
	v.value = val
	v.isSet = true
}

func (v NullableBlockRelationsBroadcastParams) IsSet() bool {
	return v.isSet
}

func (v *NullableBlockRelationsBroadcastParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlockRelationsBroadcastParams(val *BlockRelationsBroadcastParams) *NullableBlockRelationsBroadcastParams {
	return &NullableBlockRelationsBroadcastParams{value: val, isSet: true}
}

func (v NullableBlockRelationsBroadcastParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlockRelationsBroadcastParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


