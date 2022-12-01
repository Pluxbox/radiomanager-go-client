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

// ProgramDataInput struct for ProgramDataInput
type ProgramDataInput struct {
	ModelTypeId int64 `json:"model_type_id"`
	FieldValues map[string]interface{} `json:"field_values,omitempty"`
	Title string `json:"title"`
	Disabled *bool `json:"disabled,omitempty"`
	GenreId *int64 `json:"genre_id,omitempty"`
	GroupId *int64 `json:"group_id,omitempty"`
	Description *string `json:"description,omitempty"`
	ShortName *string `json:"short_name,omitempty"`
	MediumName *string `json:"medium_name,omitempty"`
	Website *string `json:"website,omitempty"`
	Email *string `json:"email,omitempty"`
	Recommended *bool `json:"recommended,omitempty"`
	Language *string `json:"language,omitempty"`
	PtyCodeId *int64 `json:"pty_code_id,omitempty"`
	Tags []int32 `json:"tags,omitempty"`
	Presenters []int32 `json:"presenters,omitempty"`
}

// NewProgramDataInput instantiates a new ProgramDataInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProgramDataInput(modelTypeId int64, title string) *ProgramDataInput {
	this := ProgramDataInput{}
	this.ModelTypeId = modelTypeId
	this.Title = title
	return &this
}

// NewProgramDataInputWithDefaults instantiates a new ProgramDataInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProgramDataInputWithDefaults() *ProgramDataInput {
	this := ProgramDataInput{}
	return &this
}

// GetModelTypeId returns the ModelTypeId field value
func (o *ProgramDataInput) GetModelTypeId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ModelTypeId
}

// GetModelTypeIdOk returns a tuple with the ModelTypeId field value
// and a boolean to check if the value has been set.
func (o *ProgramDataInput) GetModelTypeIdOk() (*int64, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ModelTypeId, true
}

// SetModelTypeId sets field value
func (o *ProgramDataInput) SetModelTypeId(v int64) {
	o.ModelTypeId = v
}

// GetFieldValues returns the FieldValues field value if set, zero value otherwise.
func (o *ProgramDataInput) GetFieldValues() map[string]interface{} {
	if o == nil || isNil(o.FieldValues) {
		var ret map[string]interface{}
		return ret
	}
	return o.FieldValues
}

// GetFieldValuesOk returns a tuple with the FieldValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramDataInput) GetFieldValuesOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.FieldValues) {
    return map[string]interface{}{}, false
	}
	return o.FieldValues, true
}

// HasFieldValues returns a boolean if a field has been set.
func (o *ProgramDataInput) HasFieldValues() bool {
	if o != nil && !isNil(o.FieldValues) {
		return true
	}

	return false
}

// SetFieldValues gets a reference to the given map[string]interface{} and assigns it to the FieldValues field.
func (o *ProgramDataInput) SetFieldValues(v map[string]interface{}) {
	o.FieldValues = v
}

// GetTitle returns the Title field value
func (o *ProgramDataInput) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *ProgramDataInput) GetTitleOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *ProgramDataInput) SetTitle(v string) {
	o.Title = v
}

// GetDisabled returns the Disabled field value if set, zero value otherwise.
func (o *ProgramDataInput) GetDisabled() bool {
	if o == nil || isNil(o.Disabled) {
		var ret bool
		return ret
	}
	return *o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramDataInput) GetDisabledOk() (*bool, bool) {
	if o == nil || isNil(o.Disabled) {
    return nil, false
	}
	return o.Disabled, true
}

// HasDisabled returns a boolean if a field has been set.
func (o *ProgramDataInput) HasDisabled() bool {
	if o != nil && !isNil(o.Disabled) {
		return true
	}

	return false
}

// SetDisabled gets a reference to the given bool and assigns it to the Disabled field.
func (o *ProgramDataInput) SetDisabled(v bool) {
	o.Disabled = &v
}

// GetGenreId returns the GenreId field value if set, zero value otherwise.
func (o *ProgramDataInput) GetGenreId() int64 {
	if o == nil || isNil(o.GenreId) {
		var ret int64
		return ret
	}
	return *o.GenreId
}

// GetGenreIdOk returns a tuple with the GenreId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramDataInput) GetGenreIdOk() (*int64, bool) {
	if o == nil || isNil(o.GenreId) {
    return nil, false
	}
	return o.GenreId, true
}

// HasGenreId returns a boolean if a field has been set.
func (o *ProgramDataInput) HasGenreId() bool {
	if o != nil && !isNil(o.GenreId) {
		return true
	}

	return false
}

// SetGenreId gets a reference to the given int64 and assigns it to the GenreId field.
func (o *ProgramDataInput) SetGenreId(v int64) {
	o.GenreId = &v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *ProgramDataInput) GetGroupId() int64 {
	if o == nil || isNil(o.GroupId) {
		var ret int64
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramDataInput) GetGroupIdOk() (*int64, bool) {
	if o == nil || isNil(o.GroupId) {
    return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *ProgramDataInput) HasGroupId() bool {
	if o != nil && !isNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given int64 and assigns it to the GroupId field.
func (o *ProgramDataInput) SetGroupId(v int64) {
	o.GroupId = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ProgramDataInput) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramDataInput) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ProgramDataInput) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ProgramDataInput) SetDescription(v string) {
	o.Description = &v
}

// GetShortName returns the ShortName field value if set, zero value otherwise.
func (o *ProgramDataInput) GetShortName() string {
	if o == nil || isNil(o.ShortName) {
		var ret string
		return ret
	}
	return *o.ShortName
}

// GetShortNameOk returns a tuple with the ShortName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramDataInput) GetShortNameOk() (*string, bool) {
	if o == nil || isNil(o.ShortName) {
    return nil, false
	}
	return o.ShortName, true
}

// HasShortName returns a boolean if a field has been set.
func (o *ProgramDataInput) HasShortName() bool {
	if o != nil && !isNil(o.ShortName) {
		return true
	}

	return false
}

// SetShortName gets a reference to the given string and assigns it to the ShortName field.
func (o *ProgramDataInput) SetShortName(v string) {
	o.ShortName = &v
}

// GetMediumName returns the MediumName field value if set, zero value otherwise.
func (o *ProgramDataInput) GetMediumName() string {
	if o == nil || isNil(o.MediumName) {
		var ret string
		return ret
	}
	return *o.MediumName
}

// GetMediumNameOk returns a tuple with the MediumName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramDataInput) GetMediumNameOk() (*string, bool) {
	if o == nil || isNil(o.MediumName) {
    return nil, false
	}
	return o.MediumName, true
}

// HasMediumName returns a boolean if a field has been set.
func (o *ProgramDataInput) HasMediumName() bool {
	if o != nil && !isNil(o.MediumName) {
		return true
	}

	return false
}

// SetMediumName gets a reference to the given string and assigns it to the MediumName field.
func (o *ProgramDataInput) SetMediumName(v string) {
	o.MediumName = &v
}

// GetWebsite returns the Website field value if set, zero value otherwise.
func (o *ProgramDataInput) GetWebsite() string {
	if o == nil || isNil(o.Website) {
		var ret string
		return ret
	}
	return *o.Website
}

// GetWebsiteOk returns a tuple with the Website field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramDataInput) GetWebsiteOk() (*string, bool) {
	if o == nil || isNil(o.Website) {
    return nil, false
	}
	return o.Website, true
}

// HasWebsite returns a boolean if a field has been set.
func (o *ProgramDataInput) HasWebsite() bool {
	if o != nil && !isNil(o.Website) {
		return true
	}

	return false
}

// SetWebsite gets a reference to the given string and assigns it to the Website field.
func (o *ProgramDataInput) SetWebsite(v string) {
	o.Website = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *ProgramDataInput) GetEmail() string {
	if o == nil || isNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramDataInput) GetEmailOk() (*string, bool) {
	if o == nil || isNil(o.Email) {
    return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *ProgramDataInput) HasEmail() bool {
	if o != nil && !isNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *ProgramDataInput) SetEmail(v string) {
	o.Email = &v
}

// GetRecommended returns the Recommended field value if set, zero value otherwise.
func (o *ProgramDataInput) GetRecommended() bool {
	if o == nil || isNil(o.Recommended) {
		var ret bool
		return ret
	}
	return *o.Recommended
}

// GetRecommendedOk returns a tuple with the Recommended field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramDataInput) GetRecommendedOk() (*bool, bool) {
	if o == nil || isNil(o.Recommended) {
    return nil, false
	}
	return o.Recommended, true
}

// HasRecommended returns a boolean if a field has been set.
func (o *ProgramDataInput) HasRecommended() bool {
	if o != nil && !isNil(o.Recommended) {
		return true
	}

	return false
}

// SetRecommended gets a reference to the given bool and assigns it to the Recommended field.
func (o *ProgramDataInput) SetRecommended(v bool) {
	o.Recommended = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *ProgramDataInput) GetLanguage() string {
	if o == nil || isNil(o.Language) {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramDataInput) GetLanguageOk() (*string, bool) {
	if o == nil || isNil(o.Language) {
    return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *ProgramDataInput) HasLanguage() bool {
	if o != nil && !isNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *ProgramDataInput) SetLanguage(v string) {
	o.Language = &v
}

// GetPtyCodeId returns the PtyCodeId field value if set, zero value otherwise.
func (o *ProgramDataInput) GetPtyCodeId() int64 {
	if o == nil || isNil(o.PtyCodeId) {
		var ret int64
		return ret
	}
	return *o.PtyCodeId
}

// GetPtyCodeIdOk returns a tuple with the PtyCodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramDataInput) GetPtyCodeIdOk() (*int64, bool) {
	if o == nil || isNil(o.PtyCodeId) {
    return nil, false
	}
	return o.PtyCodeId, true
}

// HasPtyCodeId returns a boolean if a field has been set.
func (o *ProgramDataInput) HasPtyCodeId() bool {
	if o != nil && !isNil(o.PtyCodeId) {
		return true
	}

	return false
}

// SetPtyCodeId gets a reference to the given int64 and assigns it to the PtyCodeId field.
func (o *ProgramDataInput) SetPtyCodeId(v int64) {
	o.PtyCodeId = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ProgramDataInput) GetTags() []int32 {
	if o == nil || isNil(o.Tags) {
		var ret []int32
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramDataInput) GetTagsOk() ([]int32, bool) {
	if o == nil || isNil(o.Tags) {
    return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ProgramDataInput) HasTags() bool {
	if o != nil && !isNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []int32 and assigns it to the Tags field.
func (o *ProgramDataInput) SetTags(v []int32) {
	o.Tags = v
}

// GetPresenters returns the Presenters field value if set, zero value otherwise.
func (o *ProgramDataInput) GetPresenters() []int32 {
	if o == nil || isNil(o.Presenters) {
		var ret []int32
		return ret
	}
	return o.Presenters
}

// GetPresentersOk returns a tuple with the Presenters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramDataInput) GetPresentersOk() ([]int32, bool) {
	if o == nil || isNil(o.Presenters) {
    return nil, false
	}
	return o.Presenters, true
}

// HasPresenters returns a boolean if a field has been set.
func (o *ProgramDataInput) HasPresenters() bool {
	if o != nil && !isNil(o.Presenters) {
		return true
	}

	return false
}

// SetPresenters gets a reference to the given []int32 and assigns it to the Presenters field.
func (o *ProgramDataInput) SetPresenters(v []int32) {
	o.Presenters = v
}

func (o ProgramDataInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["model_type_id"] = o.ModelTypeId
	}
	if !isNil(o.FieldValues) {
		toSerialize["field_values"] = o.FieldValues
	}
	if true {
		toSerialize["title"] = o.Title
	}
	if !isNil(o.Disabled) {
		toSerialize["disabled"] = o.Disabled
	}
	if !isNil(o.GenreId) {
		toSerialize["genre_id"] = o.GenreId
	}
	if !isNil(o.GroupId) {
		toSerialize["group_id"] = o.GroupId
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.ShortName) {
		toSerialize["short_name"] = o.ShortName
	}
	if !isNil(o.MediumName) {
		toSerialize["medium_name"] = o.MediumName
	}
	if !isNil(o.Website) {
		toSerialize["website"] = o.Website
	}
	if !isNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !isNil(o.Recommended) {
		toSerialize["recommended"] = o.Recommended
	}
	if !isNil(o.Language) {
		toSerialize["language"] = o.Language
	}
	if !isNil(o.PtyCodeId) {
		toSerialize["pty_code_id"] = o.PtyCodeId
	}
	if !isNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !isNil(o.Presenters) {
		toSerialize["presenters"] = o.Presenters
	}
	return json.Marshal(toSerialize)
}

type NullableProgramDataInput struct {
	value *ProgramDataInput
	isSet bool
}

func (v NullableProgramDataInput) Get() *ProgramDataInput {
	return v.value
}

func (v *NullableProgramDataInput) Set(val *ProgramDataInput) {
	v.value = val
	v.isSet = true
}

func (v NullableProgramDataInput) IsSet() bool {
	return v.isSet
}

func (v *NullableProgramDataInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProgramDataInput(val *ProgramDataInput) *NullableProgramDataInput {
	return &NullableProgramDataInput{value: val, isSet: true}
}

func (v NullableProgramDataInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProgramDataInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

