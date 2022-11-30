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

// InlineResponse2009 struct for InlineResponse2009
type InlineResponse2009 struct {
	CurrentPage *int64 `json:"current_page,omitempty"`
	From *int64 `json:"from,omitempty"`
	LastPage *int64 `json:"last_page,omitempty"`
	NextPageUrl *string `json:"next_page_url,omitempty"`
	Path *string `json:"path,omitempty"`
	PerPage *int64 `json:"per_page,omitempty"`
	PrevPageUrl *string `json:"prev_page_url,omitempty"`
	To *int64 `json:"to,omitempty"`
	Total *int64 `json:"total,omitempty"`
	Results []ModelTypeResult `json:"results,omitempty"`
}

// NewInlineResponse2009 instantiates a new InlineResponse2009 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2009() *InlineResponse2009 {
	this := InlineResponse2009{}
	return &this
}

// NewInlineResponse2009WithDefaults instantiates a new InlineResponse2009 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2009WithDefaults() *InlineResponse2009 {
	this := InlineResponse2009{}
	return &this
}

// GetCurrentPage returns the CurrentPage field value if set, zero value otherwise.
func (o *InlineResponse2009) GetCurrentPage() int64 {
	if o == nil || isNil(o.CurrentPage) {
		var ret int64
		return ret
	}
	return *o.CurrentPage
}

// GetCurrentPageOk returns a tuple with the CurrentPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2009) GetCurrentPageOk() (*int64, bool) {
	if o == nil || isNil(o.CurrentPage) {
    return nil, false
	}
	return o.CurrentPage, true
}

// HasCurrentPage returns a boolean if a field has been set.
func (o *InlineResponse2009) HasCurrentPage() bool {
	if o != nil && !isNil(o.CurrentPage) {
		return true
	}

	return false
}

// SetCurrentPage gets a reference to the given int64 and assigns it to the CurrentPage field.
func (o *InlineResponse2009) SetCurrentPage(v int64) {
	o.CurrentPage = &v
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *InlineResponse2009) GetFrom() int64 {
	if o == nil || isNil(o.From) {
		var ret int64
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2009) GetFromOk() (*int64, bool) {
	if o == nil || isNil(o.From) {
    return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *InlineResponse2009) HasFrom() bool {
	if o != nil && !isNil(o.From) {
		return true
	}

	return false
}

// SetFrom gets a reference to the given int64 and assigns it to the From field.
func (o *InlineResponse2009) SetFrom(v int64) {
	o.From = &v
}

// GetLastPage returns the LastPage field value if set, zero value otherwise.
func (o *InlineResponse2009) GetLastPage() int64 {
	if o == nil || isNil(o.LastPage) {
		var ret int64
		return ret
	}
	return *o.LastPage
}

// GetLastPageOk returns a tuple with the LastPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2009) GetLastPageOk() (*int64, bool) {
	if o == nil || isNil(o.LastPage) {
    return nil, false
	}
	return o.LastPage, true
}

// HasLastPage returns a boolean if a field has been set.
func (o *InlineResponse2009) HasLastPage() bool {
	if o != nil && !isNil(o.LastPage) {
		return true
	}

	return false
}

// SetLastPage gets a reference to the given int64 and assigns it to the LastPage field.
func (o *InlineResponse2009) SetLastPage(v int64) {
	o.LastPage = &v
}

// GetNextPageUrl returns the NextPageUrl field value if set, zero value otherwise.
func (o *InlineResponse2009) GetNextPageUrl() string {
	if o == nil || isNil(o.NextPageUrl) {
		var ret string
		return ret
	}
	return *o.NextPageUrl
}

// GetNextPageUrlOk returns a tuple with the NextPageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2009) GetNextPageUrlOk() (*string, bool) {
	if o == nil || isNil(o.NextPageUrl) {
    return nil, false
	}
	return o.NextPageUrl, true
}

// HasNextPageUrl returns a boolean if a field has been set.
func (o *InlineResponse2009) HasNextPageUrl() bool {
	if o != nil && !isNil(o.NextPageUrl) {
		return true
	}

	return false
}

// SetNextPageUrl gets a reference to the given string and assigns it to the NextPageUrl field.
func (o *InlineResponse2009) SetNextPageUrl(v string) {
	o.NextPageUrl = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *InlineResponse2009) GetPath() string {
	if o == nil || isNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2009) GetPathOk() (*string, bool) {
	if o == nil || isNil(o.Path) {
    return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *InlineResponse2009) HasPath() bool {
	if o != nil && !isNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *InlineResponse2009) SetPath(v string) {
	o.Path = &v
}

// GetPerPage returns the PerPage field value if set, zero value otherwise.
func (o *InlineResponse2009) GetPerPage() int64 {
	if o == nil || isNil(o.PerPage) {
		var ret int64
		return ret
	}
	return *o.PerPage
}

// GetPerPageOk returns a tuple with the PerPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2009) GetPerPageOk() (*int64, bool) {
	if o == nil || isNil(o.PerPage) {
    return nil, false
	}
	return o.PerPage, true
}

// HasPerPage returns a boolean if a field has been set.
func (o *InlineResponse2009) HasPerPage() bool {
	if o != nil && !isNil(o.PerPage) {
		return true
	}

	return false
}

// SetPerPage gets a reference to the given int64 and assigns it to the PerPage field.
func (o *InlineResponse2009) SetPerPage(v int64) {
	o.PerPage = &v
}

// GetPrevPageUrl returns the PrevPageUrl field value if set, zero value otherwise.
func (o *InlineResponse2009) GetPrevPageUrl() string {
	if o == nil || isNil(o.PrevPageUrl) {
		var ret string
		return ret
	}
	return *o.PrevPageUrl
}

// GetPrevPageUrlOk returns a tuple with the PrevPageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2009) GetPrevPageUrlOk() (*string, bool) {
	if o == nil || isNil(o.PrevPageUrl) {
    return nil, false
	}
	return o.PrevPageUrl, true
}

// HasPrevPageUrl returns a boolean if a field has been set.
func (o *InlineResponse2009) HasPrevPageUrl() bool {
	if o != nil && !isNil(o.PrevPageUrl) {
		return true
	}

	return false
}

// SetPrevPageUrl gets a reference to the given string and assigns it to the PrevPageUrl field.
func (o *InlineResponse2009) SetPrevPageUrl(v string) {
	o.PrevPageUrl = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *InlineResponse2009) GetTo() int64 {
	if o == nil || isNil(o.To) {
		var ret int64
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2009) GetToOk() (*int64, bool) {
	if o == nil || isNil(o.To) {
    return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *InlineResponse2009) HasTo() bool {
	if o != nil && !isNil(o.To) {
		return true
	}

	return false
}

// SetTo gets a reference to the given int64 and assigns it to the To field.
func (o *InlineResponse2009) SetTo(v int64) {
	o.To = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *InlineResponse2009) GetTotal() int64 {
	if o == nil || isNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2009) GetTotalOk() (*int64, bool) {
	if o == nil || isNil(o.Total) {
    return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *InlineResponse2009) HasTotal() bool {
	if o != nil && !isNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *InlineResponse2009) SetTotal(v int64) {
	o.Total = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *InlineResponse2009) GetResults() []ModelTypeResult {
	if o == nil || isNil(o.Results) {
		var ret []ModelTypeResult
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2009) GetResultsOk() ([]ModelTypeResult, bool) {
	if o == nil || isNil(o.Results) {
    return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *InlineResponse2009) HasResults() bool {
	if o != nil && !isNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []ModelTypeResult and assigns it to the Results field.
func (o *InlineResponse2009) SetResults(v []ModelTypeResult) {
	o.Results = v
}

func (o InlineResponse2009) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.CurrentPage) {
		toSerialize["current_page"] = o.CurrentPage
	}
	if !isNil(o.From) {
		toSerialize["from"] = o.From
	}
	if !isNil(o.LastPage) {
		toSerialize["last_page"] = o.LastPage
	}
	if !isNil(o.NextPageUrl) {
		toSerialize["next_page_url"] = o.NextPageUrl
	}
	if !isNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if !isNil(o.PerPage) {
		toSerialize["per_page"] = o.PerPage
	}
	if !isNil(o.PrevPageUrl) {
		toSerialize["prev_page_url"] = o.PrevPageUrl
	}
	if !isNil(o.To) {
		toSerialize["to"] = o.To
	}
	if !isNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !isNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2009 struct {
	value *InlineResponse2009
	isSet bool
}

func (v NullableInlineResponse2009) Get() *InlineResponse2009 {
	return v.value
}

func (v *NullableInlineResponse2009) Set(val *InlineResponse2009) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2009) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2009) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2009(val *InlineResponse2009) *NullableInlineResponse2009 {
	return &NullableInlineResponse2009{value: val, isSet: true}
}

func (v NullableInlineResponse2009) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2009) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


