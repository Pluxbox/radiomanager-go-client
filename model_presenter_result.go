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

// PresenterResult struct for PresenterResult
type PresenterResult struct {
	Id int64 `json:"id"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
	DeletedAt time.Time `json:"deleted_at"`
	ExternalStationId *int64 `json:"_external_station_id,omitempty"`
	ModelTypeId int64 `json:"model_type_id"`
	FieldValues map[string]interface{} `json:"field_values,omitempty"`
	Firstname *string `json:"firstname,omitempty"`
	Lastname *string `json:"lastname,omitempty"`
	Active *bool `json:"active,omitempty"`
	Name *string `json:"name,omitempty"`
	Programs *PresenterRelationsPrograms `json:"programs,omitempty"`
	Broadcasts *PresenterRelationsBroadcasts `json:"broadcasts,omitempty"`
	ModelType *PresenterRelationsModelType `json:"model_type,omitempty"`
}

// NewPresenterResult instantiates a new PresenterResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPresenterResult(id int64, updatedAt time.Time, createdAt time.Time, deletedAt time.Time, modelTypeId int64) *PresenterResult {
	this := PresenterResult{}
	this.Id = id
	this.UpdatedAt = updatedAt
	this.CreatedAt = createdAt
	this.DeletedAt = deletedAt
	this.ModelTypeId = modelTypeId
	return &this
}

// NewPresenterResultWithDefaults instantiates a new PresenterResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPresenterResultWithDefaults() *PresenterResult {
	this := PresenterResult{}
	return &this
}

// GetId returns the Id field value
func (o *PresenterResult) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PresenterResult) GetIdOk() (*int64, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PresenterResult) SetId(v int64) {
	o.Id = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *PresenterResult) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *PresenterResult) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
    return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *PresenterResult) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *PresenterResult) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *PresenterResult) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
    return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *PresenterResult) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetDeletedAt returns the DeletedAt field value
func (o *PresenterResult) GetDeletedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value
// and a boolean to check if the value has been set.
func (o *PresenterResult) GetDeletedAtOk() (*time.Time, bool) {
	if o == nil {
    return nil, false
	}
	return &o.DeletedAt, true
}

// SetDeletedAt sets field value
func (o *PresenterResult) SetDeletedAt(v time.Time) {
	o.DeletedAt = v
}

// GetExternalStationId returns the ExternalStationId field value if set, zero value otherwise.
func (o *PresenterResult) GetExternalStationId() int64 {
	if o == nil || isNil(o.ExternalStationId) {
		var ret int64
		return ret
	}
	return *o.ExternalStationId
}

// GetExternalStationIdOk returns a tuple with the ExternalStationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PresenterResult) GetExternalStationIdOk() (*int64, bool) {
	if o == nil || isNil(o.ExternalStationId) {
    return nil, false
	}
	return o.ExternalStationId, true
}

// HasExternalStationId returns a boolean if a field has been set.
func (o *PresenterResult) HasExternalStationId() bool {
	if o != nil && !isNil(o.ExternalStationId) {
		return true
	}

	return false
}

// SetExternalStationId gets a reference to the given int64 and assigns it to the ExternalStationId field.
func (o *PresenterResult) SetExternalStationId(v int64) {
	o.ExternalStationId = &v
}

// GetModelTypeId returns the ModelTypeId field value
func (o *PresenterResult) GetModelTypeId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ModelTypeId
}

// GetModelTypeIdOk returns a tuple with the ModelTypeId field value
// and a boolean to check if the value has been set.
func (o *PresenterResult) GetModelTypeIdOk() (*int64, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ModelTypeId, true
}

// SetModelTypeId sets field value
func (o *PresenterResult) SetModelTypeId(v int64) {
	o.ModelTypeId = v
}

// GetFieldValues returns the FieldValues field value if set, zero value otherwise.
func (o *PresenterResult) GetFieldValues() map[string]interface{} {
	if o == nil || isNil(o.FieldValues) {
		var ret map[string]interface{}
		return ret
	}
	return o.FieldValues
}

// GetFieldValuesOk returns a tuple with the FieldValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PresenterResult) GetFieldValuesOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.FieldValues) {
    return map[string]interface{}{}, false
	}
	return o.FieldValues, true
}

// HasFieldValues returns a boolean if a field has been set.
func (o *PresenterResult) HasFieldValues() bool {
	if o != nil && !isNil(o.FieldValues) {
		return true
	}

	return false
}

// SetFieldValues gets a reference to the given map[string]interface{} and assigns it to the FieldValues field.
func (o *PresenterResult) SetFieldValues(v map[string]interface{}) {
	o.FieldValues = v
}

// GetFirstname returns the Firstname field value if set, zero value otherwise.
func (o *PresenterResult) GetFirstname() string {
	if o == nil || isNil(o.Firstname) {
		var ret string
		return ret
	}
	return *o.Firstname
}

// GetFirstnameOk returns a tuple with the Firstname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PresenterResult) GetFirstnameOk() (*string, bool) {
	if o == nil || isNil(o.Firstname) {
    return nil, false
	}
	return o.Firstname, true
}

// HasFirstname returns a boolean if a field has been set.
func (o *PresenterResult) HasFirstname() bool {
	if o != nil && !isNil(o.Firstname) {
		return true
	}

	return false
}

// SetFirstname gets a reference to the given string and assigns it to the Firstname field.
func (o *PresenterResult) SetFirstname(v string) {
	o.Firstname = &v
}

// GetLastname returns the Lastname field value if set, zero value otherwise.
func (o *PresenterResult) GetLastname() string {
	if o == nil || isNil(o.Lastname) {
		var ret string
		return ret
	}
	return *o.Lastname
}

// GetLastnameOk returns a tuple with the Lastname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PresenterResult) GetLastnameOk() (*string, bool) {
	if o == nil || isNil(o.Lastname) {
    return nil, false
	}
	return o.Lastname, true
}

// HasLastname returns a boolean if a field has been set.
func (o *PresenterResult) HasLastname() bool {
	if o != nil && !isNil(o.Lastname) {
		return true
	}

	return false
}

// SetLastname gets a reference to the given string and assigns it to the Lastname field.
func (o *PresenterResult) SetLastname(v string) {
	o.Lastname = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *PresenterResult) GetActive() bool {
	if o == nil || isNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PresenterResult) GetActiveOk() (*bool, bool) {
	if o == nil || isNil(o.Active) {
    return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *PresenterResult) HasActive() bool {
	if o != nil && !isNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *PresenterResult) SetActive(v bool) {
	o.Active = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PresenterResult) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PresenterResult) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PresenterResult) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PresenterResult) SetName(v string) {
	o.Name = &v
}

// GetPrograms returns the Programs field value if set, zero value otherwise.
func (o *PresenterResult) GetPrograms() PresenterRelationsPrograms {
	if o == nil || isNil(o.Programs) {
		var ret PresenterRelationsPrograms
		return ret
	}
	return *o.Programs
}

// GetProgramsOk returns a tuple with the Programs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PresenterResult) GetProgramsOk() (*PresenterRelationsPrograms, bool) {
	if o == nil || isNil(o.Programs) {
    return nil, false
	}
	return o.Programs, true
}

// HasPrograms returns a boolean if a field has been set.
func (o *PresenterResult) HasPrograms() bool {
	if o != nil && !isNil(o.Programs) {
		return true
	}

	return false
}

// SetPrograms gets a reference to the given PresenterRelationsPrograms and assigns it to the Programs field.
func (o *PresenterResult) SetPrograms(v PresenterRelationsPrograms) {
	o.Programs = &v
}

// GetBroadcasts returns the Broadcasts field value if set, zero value otherwise.
func (o *PresenterResult) GetBroadcasts() PresenterRelationsBroadcasts {
	if o == nil || isNil(o.Broadcasts) {
		var ret PresenterRelationsBroadcasts
		return ret
	}
	return *o.Broadcasts
}

// GetBroadcastsOk returns a tuple with the Broadcasts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PresenterResult) GetBroadcastsOk() (*PresenterRelationsBroadcasts, bool) {
	if o == nil || isNil(o.Broadcasts) {
    return nil, false
	}
	return o.Broadcasts, true
}

// HasBroadcasts returns a boolean if a field has been set.
func (o *PresenterResult) HasBroadcasts() bool {
	if o != nil && !isNil(o.Broadcasts) {
		return true
	}

	return false
}

// SetBroadcasts gets a reference to the given PresenterRelationsBroadcasts and assigns it to the Broadcasts field.
func (o *PresenterResult) SetBroadcasts(v PresenterRelationsBroadcasts) {
	o.Broadcasts = &v
}

// GetModelType returns the ModelType field value if set, zero value otherwise.
func (o *PresenterResult) GetModelType() PresenterRelationsModelType {
	if o == nil || isNil(o.ModelType) {
		var ret PresenterRelationsModelType
		return ret
	}
	return *o.ModelType
}

// GetModelTypeOk returns a tuple with the ModelType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PresenterResult) GetModelTypeOk() (*PresenterRelationsModelType, bool) {
	if o == nil || isNil(o.ModelType) {
    return nil, false
	}
	return o.ModelType, true
}

// HasModelType returns a boolean if a field has been set.
func (o *PresenterResult) HasModelType() bool {
	if o != nil && !isNil(o.ModelType) {
		return true
	}

	return false
}

// SetModelType gets a reference to the given PresenterRelationsModelType and assigns it to the ModelType field.
func (o *PresenterResult) SetModelType(v PresenterRelationsModelType) {
	o.ModelType = &v
}

func (o PresenterResult) MarshalJSON() ([]byte, error) {
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
	if !isNil(o.ExternalStationId) {
		toSerialize["_external_station_id"] = o.ExternalStationId
	}
	if true {
		toSerialize["model_type_id"] = o.ModelTypeId
	}
	if !isNil(o.FieldValues) {
		toSerialize["field_values"] = o.FieldValues
	}
	if !isNil(o.Firstname) {
		toSerialize["firstname"] = o.Firstname
	}
	if !isNil(o.Lastname) {
		toSerialize["lastname"] = o.Lastname
	}
	if !isNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Programs) {
		toSerialize["programs"] = o.Programs
	}
	if !isNil(o.Broadcasts) {
		toSerialize["broadcasts"] = o.Broadcasts
	}
	if !isNil(o.ModelType) {
		toSerialize["model_type"] = o.ModelType
	}
	return json.Marshal(toSerialize)
}

type NullablePresenterResult struct {
	value *PresenterResult
	isSet bool
}

func (v NullablePresenterResult) Get() *PresenterResult {
	return v.value
}

func (v *NullablePresenterResult) Set(val *PresenterResult) {
	v.value = val
	v.isSet = true
}

func (v NullablePresenterResult) IsSet() bool {
	return v.isSet
}

func (v *NullablePresenterResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePresenterResult(val *PresenterResult) *NullablePresenterResult {
	return &NullablePresenterResult{value: val, isSet: true}
}

func (v NullablePresenterResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePresenterResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

