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

// UserResult struct for UserResult
type UserResult struct {
	Id int64 `json:"id"`
	Email string `json:"email"`
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
	Phone *string `json:"phone,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	Active *bool `json:"active,omitempty"`
	Settings *UserResultSettings `json:"settings,omitempty"`
	Language *bool `json:"language,omitempty"`
	ActiveExternalStationId *int64 `json:"active_external_station_id,omitempty"`
	Roles []UserResultRoles `json:"roles,omitempty"`
	Groups []UserResultGroups `json:"groups,omitempty"`
}

// NewUserResult instantiates a new UserResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserResult(id int64, email string, firstname string, lastname string) *UserResult {
	this := UserResult{}
	this.Id = id
	this.Email = email
	this.Firstname = firstname
	this.Lastname = lastname
	return &this
}

// NewUserResultWithDefaults instantiates a new UserResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserResultWithDefaults() *UserResult {
	this := UserResult{}
	return &this
}

// GetId returns the Id field value
func (o *UserResult) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UserResult) GetIdOk() (*int64, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UserResult) SetId(v int64) {
	o.Id = v
}

// GetEmail returns the Email field value
func (o *UserResult) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *UserResult) GetEmailOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *UserResult) SetEmail(v string) {
	o.Email = v
}

// GetFirstname returns the Firstname field value
func (o *UserResult) GetFirstname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Firstname
}

// GetFirstnameOk returns a tuple with the Firstname field value
// and a boolean to check if the value has been set.
func (o *UserResult) GetFirstnameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Firstname, true
}

// SetFirstname sets field value
func (o *UserResult) SetFirstname(v string) {
	o.Firstname = v
}

// GetLastname returns the Lastname field value
func (o *UserResult) GetLastname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Lastname
}

// GetLastnameOk returns a tuple with the Lastname field value
// and a boolean to check if the value has been set.
func (o *UserResult) GetLastnameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Lastname, true
}

// SetLastname sets field value
func (o *UserResult) SetLastname(v string) {
	o.Lastname = v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *UserResult) GetPhone() string {
	if o == nil || isNil(o.Phone) {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResult) GetPhoneOk() (*string, bool) {
	if o == nil || isNil(o.Phone) {
    return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *UserResult) HasPhone() bool {
	if o != nil && !isNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *UserResult) SetPhone(v string) {
	o.Phone = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *UserResult) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResult) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *UserResult) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *UserResult) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *UserResult) GetUpdatedAt() time.Time {
	if o == nil || isNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResult) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.UpdatedAt) {
    return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *UserResult) HasUpdatedAt() bool {
	if o != nil && !isNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *UserResult) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *UserResult) GetActive() bool {
	if o == nil || isNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResult) GetActiveOk() (*bool, bool) {
	if o == nil || isNil(o.Active) {
    return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *UserResult) HasActive() bool {
	if o != nil && !isNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *UserResult) SetActive(v bool) {
	o.Active = &v
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *UserResult) GetSettings() UserResultSettings {
	if o == nil || isNil(o.Settings) {
		var ret UserResultSettings
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResult) GetSettingsOk() (*UserResultSettings, bool) {
	if o == nil || isNil(o.Settings) {
    return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *UserResult) HasSettings() bool {
	if o != nil && !isNil(o.Settings) {
		return true
	}

	return false
}

// SetSettings gets a reference to the given UserResultSettings and assigns it to the Settings field.
func (o *UserResult) SetSettings(v UserResultSettings) {
	o.Settings = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *UserResult) GetLanguage() bool {
	if o == nil || isNil(o.Language) {
		var ret bool
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResult) GetLanguageOk() (*bool, bool) {
	if o == nil || isNil(o.Language) {
    return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *UserResult) HasLanguage() bool {
	if o != nil && !isNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given bool and assigns it to the Language field.
func (o *UserResult) SetLanguage(v bool) {
	o.Language = &v
}

// GetActiveExternalStationId returns the ActiveExternalStationId field value if set, zero value otherwise.
func (o *UserResult) GetActiveExternalStationId() int64 {
	if o == nil || isNil(o.ActiveExternalStationId) {
		var ret int64
		return ret
	}
	return *o.ActiveExternalStationId
}

// GetActiveExternalStationIdOk returns a tuple with the ActiveExternalStationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResult) GetActiveExternalStationIdOk() (*int64, bool) {
	if o == nil || isNil(o.ActiveExternalStationId) {
    return nil, false
	}
	return o.ActiveExternalStationId, true
}

// HasActiveExternalStationId returns a boolean if a field has been set.
func (o *UserResult) HasActiveExternalStationId() bool {
	if o != nil && !isNil(o.ActiveExternalStationId) {
		return true
	}

	return false
}

// SetActiveExternalStationId gets a reference to the given int64 and assigns it to the ActiveExternalStationId field.
func (o *UserResult) SetActiveExternalStationId(v int64) {
	o.ActiveExternalStationId = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *UserResult) GetRoles() []UserResultRoles {
	if o == nil || isNil(o.Roles) {
		var ret []UserResultRoles
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResult) GetRolesOk() ([]UserResultRoles, bool) {
	if o == nil || isNil(o.Roles) {
    return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *UserResult) HasRoles() bool {
	if o != nil && !isNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []UserResultRoles and assigns it to the Roles field.
func (o *UserResult) SetRoles(v []UserResultRoles) {
	o.Roles = v
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *UserResult) GetGroups() []UserResultGroups {
	if o == nil || isNil(o.Groups) {
		var ret []UserResultGroups
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResult) GetGroupsOk() ([]UserResultGroups, bool) {
	if o == nil || isNil(o.Groups) {
    return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *UserResult) HasGroups() bool {
	if o != nil && !isNil(o.Groups) {
		return true
	}

	return false
}

// SetGroups gets a reference to the given []UserResultGroups and assigns it to the Groups field.
func (o *UserResult) SetGroups(v []UserResultGroups) {
	o.Groups = v
}

func (o UserResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["email"] = o.Email
	}
	if true {
		toSerialize["firstname"] = o.Firstname
	}
	if true {
		toSerialize["lastname"] = o.Lastname
	}
	if !isNil(o.Phone) {
		toSerialize["phone"] = o.Phone
	}
	if !isNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !isNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !isNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !isNil(o.Settings) {
		toSerialize["settings"] = o.Settings
	}
	if !isNil(o.Language) {
		toSerialize["language"] = o.Language
	}
	if !isNil(o.ActiveExternalStationId) {
		toSerialize["active_external_station_id"] = o.ActiveExternalStationId
	}
	if !isNil(o.Roles) {
		toSerialize["roles"] = o.Roles
	}
	if !isNil(o.Groups) {
		toSerialize["groups"] = o.Groups
	}
	return json.Marshal(toSerialize)
}

type NullableUserResult struct {
	value *UserResult
	isSet bool
}

func (v NullableUserResult) Get() *UserResult {
	return v.value
}

func (v *NullableUserResult) Set(val *UserResult) {
	v.value = val
	v.isSet = true
}

func (v NullableUserResult) IsSet() bool {
	return v.isSet
}

func (v *NullableUserResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserResult(val *UserResult) *NullableUserResult {
	return &NullableUserResult{value: val, isSet: true}
}

func (v NullableUserResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


