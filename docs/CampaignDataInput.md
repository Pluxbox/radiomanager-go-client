# CampaignDataInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ModelTypeId** | **int64** |  | 
**FieldValues** | Pointer to **map[string]interface{}** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Start** | **time.Time** |  | 
**Stop** | **time.Time** |  | 
**Recommended** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewCampaignDataInput

`func NewCampaignDataInput(modelTypeId int64, start time.Time, stop time.Time, ) *CampaignDataInput`

NewCampaignDataInput instantiates a new CampaignDataInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignDataInputWithDefaults

`func NewCampaignDataInputWithDefaults() *CampaignDataInput`

NewCampaignDataInputWithDefaults instantiates a new CampaignDataInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModelTypeId

`func (o *CampaignDataInput) GetModelTypeId() int64`

GetModelTypeId returns the ModelTypeId field if non-nil, zero value otherwise.

### GetModelTypeIdOk

`func (o *CampaignDataInput) GetModelTypeIdOk() (*int64, bool)`

GetModelTypeIdOk returns a tuple with the ModelTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelTypeId

`func (o *CampaignDataInput) SetModelTypeId(v int64)`

SetModelTypeId sets ModelTypeId field to given value.


### GetFieldValues

`func (o *CampaignDataInput) GetFieldValues() map[string]interface{}`

GetFieldValues returns the FieldValues field if non-nil, zero value otherwise.

### GetFieldValuesOk

`func (o *CampaignDataInput) GetFieldValuesOk() (*map[string]interface{}, bool)`

GetFieldValuesOk returns a tuple with the FieldValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldValues

`func (o *CampaignDataInput) SetFieldValues(v map[string]interface{})`

SetFieldValues sets FieldValues field to given value.

### HasFieldValues

`func (o *CampaignDataInput) HasFieldValues() bool`

HasFieldValues returns a boolean if a field has been set.

### GetTitle

`func (o *CampaignDataInput) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CampaignDataInput) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CampaignDataInput) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CampaignDataInput) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetStart

`func (o *CampaignDataInput) GetStart() time.Time`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *CampaignDataInput) GetStartOk() (*time.Time, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *CampaignDataInput) SetStart(v time.Time)`

SetStart sets Start field to given value.


### GetStop

`func (o *CampaignDataInput) GetStop() time.Time`

GetStop returns the Stop field if non-nil, zero value otherwise.

### GetStopOk

`func (o *CampaignDataInput) GetStopOk() (*time.Time, bool)`

GetStopOk returns a tuple with the Stop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStop

`func (o *CampaignDataInput) SetStop(v time.Time)`

SetStop sets Stop field to given value.


### GetRecommended

`func (o *CampaignDataInput) GetRecommended() bool`

GetRecommended returns the Recommended field if non-nil, zero value otherwise.

### GetRecommendedOk

`func (o *CampaignDataInput) GetRecommendedOk() (*bool, bool)`

GetRecommendedOk returns a tuple with the Recommended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommended

`func (o *CampaignDataInput) SetRecommended(v bool)`

SetRecommended sets Recommended field to given value.

### HasRecommended

`func (o *CampaignDataInput) HasRecommended() bool`

HasRecommended returns a boolean if a field has been set.

### GetDescription

`func (o *CampaignDataInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CampaignDataInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CampaignDataInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CampaignDataInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


