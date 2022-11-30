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

// Item struct for Item
type Item struct {
	ModelTypeId int64 `json:"model_type_id"`
	BlockId *int64 `json:"block_id,omitempty"`
	ExternalId *string `json:"external_id,omitempty"`
	FieldValues map[string]interface{} `json:"field_values,omitempty"`
	Title *string `json:"title,omitempty"`
	Duration *int64 `json:"duration,omitempty"`
	Start *time.Time `json:"start,omitempty"`
	Status *string `json:"status,omitempty"`
	Import *int64 `json:"import,omitempty"`
	CampaignId *int64 `json:"campaign_id,omitempty"`
	Recommended *bool `json:"recommended,omitempty"`
	StationDraftId *int64 `json:"station_draft_id,omitempty"`
	ProgramDraftId *int64 `json:"program_draft_id,omitempty"`
	UserDraftId *int64 `json:"user_draft_id,omitempty"`
	StaticStart *bool `json:"static_start,omitempty"`
	Details *string `json:"details,omitempty"`
}

// NewItem instantiates a new Item object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItem(modelTypeId int64) *Item {
	this := Item{}
	this.ModelTypeId = modelTypeId
	return &this
}

// NewItemWithDefaults instantiates a new Item object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemWithDefaults() *Item {
	this := Item{}
	return &this
}

// GetModelTypeId returns the ModelTypeId field value
func (o *Item) GetModelTypeId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ModelTypeId
}

// GetModelTypeIdOk returns a tuple with the ModelTypeId field value
// and a boolean to check if the value has been set.
func (o *Item) GetModelTypeIdOk() (*int64, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ModelTypeId, true
}

// SetModelTypeId sets field value
func (o *Item) SetModelTypeId(v int64) {
	o.ModelTypeId = v
}

// GetBlockId returns the BlockId field value if set, zero value otherwise.
func (o *Item) GetBlockId() int64 {
	if o == nil || isNil(o.BlockId) {
		var ret int64
		return ret
	}
	return *o.BlockId
}

// GetBlockIdOk returns a tuple with the BlockId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Item) GetBlockIdOk() (*int64, bool) {
	if o == nil || isNil(o.BlockId) {
    return nil, false
	}
	return o.BlockId, true
}

// HasBlockId returns a boolean if a field has been set.
func (o *Item) HasBlockId() bool {
	if o != nil && !isNil(o.BlockId) {
		return true
	}

	return false
}

// SetBlockId gets a reference to the given int64 and assigns it to the BlockId field.
func (o *Item) SetBlockId(v int64) {
	o.BlockId = &v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *Item) GetExternalId() string {
	if o == nil || isNil(o.ExternalId) {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Item) GetExternalIdOk() (*string, bool) {
	if o == nil || isNil(o.ExternalId) {
    return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *Item) HasExternalId() bool {
	if o != nil && !isNil(o.ExternalId) {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *Item) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetFieldValues returns the FieldValues field value if set, zero value otherwise.
func (o *Item) GetFieldValues() map[string]interface{} {
	if o == nil || isNil(o.FieldValues) {
		var ret map[string]interface{}
		return ret
	}
	return o.FieldValues
}

// GetFieldValuesOk returns a tuple with the FieldValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Item) GetFieldValuesOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.FieldValues) {
    return map[string]interface{}{}, false
	}
	return o.FieldValues, true
}

// HasFieldValues returns a boolean if a field has been set.
func (o *Item) HasFieldValues() bool {
	if o != nil && !isNil(o.FieldValues) {
		return true
	}

	return false
}

// SetFieldValues gets a reference to the given map[string]interface{} and assigns it to the FieldValues field.
func (o *Item) SetFieldValues(v map[string]interface{}) {
	o.FieldValues = v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *Item) GetTitle() string {
	if o == nil || isNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Item) GetTitleOk() (*string, bool) {
	if o == nil || isNil(o.Title) {
    return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *Item) HasTitle() bool {
	if o != nil && !isNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *Item) SetTitle(v string) {
	o.Title = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *Item) GetDuration() int64 {
	if o == nil || isNil(o.Duration) {
		var ret int64
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Item) GetDurationOk() (*int64, bool) {
	if o == nil || isNil(o.Duration) {
    return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *Item) HasDuration() bool {
	if o != nil && !isNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given int64 and assigns it to the Duration field.
func (o *Item) SetDuration(v int64) {
	o.Duration = &v
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *Item) GetStart() time.Time {
	if o == nil || isNil(o.Start) {
		var ret time.Time
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Item) GetStartOk() (*time.Time, bool) {
	if o == nil || isNil(o.Start) {
    return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *Item) HasStart() bool {
	if o != nil && !isNil(o.Start) {
		return true
	}

	return false
}

// SetStart gets a reference to the given time.Time and assigns it to the Start field.
func (o *Item) SetStart(v time.Time) {
	o.Start = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Item) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Item) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Item) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Item) SetStatus(v string) {
	o.Status = &v
}

// GetImport returns the Import field value if set, zero value otherwise.
func (o *Item) GetImport() int64 {
	if o == nil || isNil(o.Import) {
		var ret int64
		return ret
	}
	return *o.Import
}

// GetImportOk returns a tuple with the Import field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Item) GetImportOk() (*int64, bool) {
	if o == nil || isNil(o.Import) {
    return nil, false
	}
	return o.Import, true
}

// HasImport returns a boolean if a field has been set.
func (o *Item) HasImport() bool {
	if o != nil && !isNil(o.Import) {
		return true
	}

	return false
}

// SetImport gets a reference to the given int64 and assigns it to the Import field.
func (o *Item) SetImport(v int64) {
	o.Import = &v
}

// GetCampaignId returns the CampaignId field value if set, zero value otherwise.
func (o *Item) GetCampaignId() int64 {
	if o == nil || isNil(o.CampaignId) {
		var ret int64
		return ret
	}
	return *o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Item) GetCampaignIdOk() (*int64, bool) {
	if o == nil || isNil(o.CampaignId) {
    return nil, false
	}
	return o.CampaignId, true
}

// HasCampaignId returns a boolean if a field has been set.
func (o *Item) HasCampaignId() bool {
	if o != nil && !isNil(o.CampaignId) {
		return true
	}

	return false
}

// SetCampaignId gets a reference to the given int64 and assigns it to the CampaignId field.
func (o *Item) SetCampaignId(v int64) {
	o.CampaignId = &v
}

// GetRecommended returns the Recommended field value if set, zero value otherwise.
func (o *Item) GetRecommended() bool {
	if o == nil || isNil(o.Recommended) {
		var ret bool
		return ret
	}
	return *o.Recommended
}

// GetRecommendedOk returns a tuple with the Recommended field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Item) GetRecommendedOk() (*bool, bool) {
	if o == nil || isNil(o.Recommended) {
    return nil, false
	}
	return o.Recommended, true
}

// HasRecommended returns a boolean if a field has been set.
func (o *Item) HasRecommended() bool {
	if o != nil && !isNil(o.Recommended) {
		return true
	}

	return false
}

// SetRecommended gets a reference to the given bool and assigns it to the Recommended field.
func (o *Item) SetRecommended(v bool) {
	o.Recommended = &v
}

// GetStationDraftId returns the StationDraftId field value if set, zero value otherwise.
func (o *Item) GetStationDraftId() int64 {
	if o == nil || isNil(o.StationDraftId) {
		var ret int64
		return ret
	}
	return *o.StationDraftId
}

// GetStationDraftIdOk returns a tuple with the StationDraftId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Item) GetStationDraftIdOk() (*int64, bool) {
	if o == nil || isNil(o.StationDraftId) {
    return nil, false
	}
	return o.StationDraftId, true
}

// HasStationDraftId returns a boolean if a field has been set.
func (o *Item) HasStationDraftId() bool {
	if o != nil && !isNil(o.StationDraftId) {
		return true
	}

	return false
}

// SetStationDraftId gets a reference to the given int64 and assigns it to the StationDraftId field.
func (o *Item) SetStationDraftId(v int64) {
	o.StationDraftId = &v
}

// GetProgramDraftId returns the ProgramDraftId field value if set, zero value otherwise.
func (o *Item) GetProgramDraftId() int64 {
	if o == nil || isNil(o.ProgramDraftId) {
		var ret int64
		return ret
	}
	return *o.ProgramDraftId
}

// GetProgramDraftIdOk returns a tuple with the ProgramDraftId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Item) GetProgramDraftIdOk() (*int64, bool) {
	if o == nil || isNil(o.ProgramDraftId) {
    return nil, false
	}
	return o.ProgramDraftId, true
}

// HasProgramDraftId returns a boolean if a field has been set.
func (o *Item) HasProgramDraftId() bool {
	if o != nil && !isNil(o.ProgramDraftId) {
		return true
	}

	return false
}

// SetProgramDraftId gets a reference to the given int64 and assigns it to the ProgramDraftId field.
func (o *Item) SetProgramDraftId(v int64) {
	o.ProgramDraftId = &v
}

// GetUserDraftId returns the UserDraftId field value if set, zero value otherwise.
func (o *Item) GetUserDraftId() int64 {
	if o == nil || isNil(o.UserDraftId) {
		var ret int64
		return ret
	}
	return *o.UserDraftId
}

// GetUserDraftIdOk returns a tuple with the UserDraftId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Item) GetUserDraftIdOk() (*int64, bool) {
	if o == nil || isNil(o.UserDraftId) {
    return nil, false
	}
	return o.UserDraftId, true
}

// HasUserDraftId returns a boolean if a field has been set.
func (o *Item) HasUserDraftId() bool {
	if o != nil && !isNil(o.UserDraftId) {
		return true
	}

	return false
}

// SetUserDraftId gets a reference to the given int64 and assigns it to the UserDraftId field.
func (o *Item) SetUserDraftId(v int64) {
	o.UserDraftId = &v
}

// GetStaticStart returns the StaticStart field value if set, zero value otherwise.
func (o *Item) GetStaticStart() bool {
	if o == nil || isNil(o.StaticStart) {
		var ret bool
		return ret
	}
	return *o.StaticStart
}

// GetStaticStartOk returns a tuple with the StaticStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Item) GetStaticStartOk() (*bool, bool) {
	if o == nil || isNil(o.StaticStart) {
    return nil, false
	}
	return o.StaticStart, true
}

// HasStaticStart returns a boolean if a field has been set.
func (o *Item) HasStaticStart() bool {
	if o != nil && !isNil(o.StaticStart) {
		return true
	}

	return false
}

// SetStaticStart gets a reference to the given bool and assigns it to the StaticStart field.
func (o *Item) SetStaticStart(v bool) {
	o.StaticStart = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *Item) GetDetails() string {
	if o == nil || isNil(o.Details) {
		var ret string
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Item) GetDetailsOk() (*string, bool) {
	if o == nil || isNil(o.Details) {
    return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *Item) HasDetails() bool {
	if o != nil && !isNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given string and assigns it to the Details field.
func (o *Item) SetDetails(v string) {
	o.Details = &v
}

func (o Item) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["model_type_id"] = o.ModelTypeId
	}
	if !isNil(o.BlockId) {
		toSerialize["block_id"] = o.BlockId
	}
	if !isNil(o.ExternalId) {
		toSerialize["external_id"] = o.ExternalId
	}
	if !isNil(o.FieldValues) {
		toSerialize["field_values"] = o.FieldValues
	}
	if !isNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !isNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	if !isNil(o.Start) {
		toSerialize["start"] = o.Start
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.Import) {
		toSerialize["import"] = o.Import
	}
	if !isNil(o.CampaignId) {
		toSerialize["campaign_id"] = o.CampaignId
	}
	if !isNil(o.Recommended) {
		toSerialize["recommended"] = o.Recommended
	}
	if !isNil(o.StationDraftId) {
		toSerialize["station_draft_id"] = o.StationDraftId
	}
	if !isNil(o.ProgramDraftId) {
		toSerialize["program_draft_id"] = o.ProgramDraftId
	}
	if !isNil(o.UserDraftId) {
		toSerialize["user_draft_id"] = o.UserDraftId
	}
	if !isNil(o.StaticStart) {
		toSerialize["static_start"] = o.StaticStart
	}
	if !isNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	return json.Marshal(toSerialize)
}

type NullableItem struct {
	value *Item
	isSet bool
}

func (v NullableItem) Get() *Item {
	return v.value
}

func (v *NullableItem) Set(val *Item) {
	v.value = val
	v.isSet = true
}

func (v NullableItem) IsSet() bool {
	return v.isSet
}

func (v *NullableItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItem(val *Item) *NullableItem {
	return &NullableItem{value: val, isSet: true}
}

func (v NullableItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


