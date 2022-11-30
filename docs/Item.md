# Item

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ModelTypeId** | **int64** |  | 
**BlockId** | Pointer to **int64** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**FieldValues** | Pointer to **map[string]interface{}** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Duration** | Pointer to **int64** |  | [optional] 
**Start** | Pointer to **time.Time** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Import** | Pointer to **int64** |  | [optional] 
**CampaignId** | Pointer to **int64** |  | [optional] 
**Recommended** | Pointer to **bool** |  | [optional] 
**StationDraftId** | Pointer to **int64** |  | [optional] 
**ProgramDraftId** | Pointer to **int64** |  | [optional] 
**UserDraftId** | Pointer to **int64** |  | [optional] 
**StaticStart** | Pointer to **bool** |  | [optional] 
**Details** | Pointer to **string** |  | [optional] 

## Methods

### NewItem

`func NewItem(modelTypeId int64, ) *Item`

NewItem instantiates a new Item object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemWithDefaults

`func NewItemWithDefaults() *Item`

NewItemWithDefaults instantiates a new Item object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModelTypeId

`func (o *Item) GetModelTypeId() int64`

GetModelTypeId returns the ModelTypeId field if non-nil, zero value otherwise.

### GetModelTypeIdOk

`func (o *Item) GetModelTypeIdOk() (*int64, bool)`

GetModelTypeIdOk returns a tuple with the ModelTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelTypeId

`func (o *Item) SetModelTypeId(v int64)`

SetModelTypeId sets ModelTypeId field to given value.


### GetBlockId

`func (o *Item) GetBlockId() int64`

GetBlockId returns the BlockId field if non-nil, zero value otherwise.

### GetBlockIdOk

`func (o *Item) GetBlockIdOk() (*int64, bool)`

GetBlockIdOk returns a tuple with the BlockId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockId

`func (o *Item) SetBlockId(v int64)`

SetBlockId sets BlockId field to given value.

### HasBlockId

`func (o *Item) HasBlockId() bool`

HasBlockId returns a boolean if a field has been set.

### GetExternalId

`func (o *Item) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *Item) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *Item) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *Item) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetFieldValues

`func (o *Item) GetFieldValues() map[string]interface{}`

GetFieldValues returns the FieldValues field if non-nil, zero value otherwise.

### GetFieldValuesOk

`func (o *Item) GetFieldValuesOk() (*map[string]interface{}, bool)`

GetFieldValuesOk returns a tuple with the FieldValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldValues

`func (o *Item) SetFieldValues(v map[string]interface{})`

SetFieldValues sets FieldValues field to given value.

### HasFieldValues

`func (o *Item) HasFieldValues() bool`

HasFieldValues returns a boolean if a field has been set.

### GetTitle

`func (o *Item) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Item) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Item) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Item) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDuration

`func (o *Item) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *Item) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *Item) SetDuration(v int64)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *Item) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetStart

`func (o *Item) GetStart() time.Time`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *Item) GetStartOk() (*time.Time, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *Item) SetStart(v time.Time)`

SetStart sets Start field to given value.

### HasStart

`func (o *Item) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetStatus

`func (o *Item) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Item) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Item) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Item) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetImport

`func (o *Item) GetImport() int64`

GetImport returns the Import field if non-nil, zero value otherwise.

### GetImportOk

`func (o *Item) GetImportOk() (*int64, bool)`

GetImportOk returns a tuple with the Import field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImport

`func (o *Item) SetImport(v int64)`

SetImport sets Import field to given value.

### HasImport

`func (o *Item) HasImport() bool`

HasImport returns a boolean if a field has been set.

### GetCampaignId

`func (o *Item) GetCampaignId() int64`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *Item) GetCampaignIdOk() (*int64, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *Item) SetCampaignId(v int64)`

SetCampaignId sets CampaignId field to given value.

### HasCampaignId

`func (o *Item) HasCampaignId() bool`

HasCampaignId returns a boolean if a field has been set.

### GetRecommended

`func (o *Item) GetRecommended() bool`

GetRecommended returns the Recommended field if non-nil, zero value otherwise.

### GetRecommendedOk

`func (o *Item) GetRecommendedOk() (*bool, bool)`

GetRecommendedOk returns a tuple with the Recommended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommended

`func (o *Item) SetRecommended(v bool)`

SetRecommended sets Recommended field to given value.

### HasRecommended

`func (o *Item) HasRecommended() bool`

HasRecommended returns a boolean if a field has been set.

### GetStationDraftId

`func (o *Item) GetStationDraftId() int64`

GetStationDraftId returns the StationDraftId field if non-nil, zero value otherwise.

### GetStationDraftIdOk

`func (o *Item) GetStationDraftIdOk() (*int64, bool)`

GetStationDraftIdOk returns a tuple with the StationDraftId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStationDraftId

`func (o *Item) SetStationDraftId(v int64)`

SetStationDraftId sets StationDraftId field to given value.

### HasStationDraftId

`func (o *Item) HasStationDraftId() bool`

HasStationDraftId returns a boolean if a field has been set.

### GetProgramDraftId

`func (o *Item) GetProgramDraftId() int64`

GetProgramDraftId returns the ProgramDraftId field if non-nil, zero value otherwise.

### GetProgramDraftIdOk

`func (o *Item) GetProgramDraftIdOk() (*int64, bool)`

GetProgramDraftIdOk returns a tuple with the ProgramDraftId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgramDraftId

`func (o *Item) SetProgramDraftId(v int64)`

SetProgramDraftId sets ProgramDraftId field to given value.

### HasProgramDraftId

`func (o *Item) HasProgramDraftId() bool`

HasProgramDraftId returns a boolean if a field has been set.

### GetUserDraftId

`func (o *Item) GetUserDraftId() int64`

GetUserDraftId returns the UserDraftId field if non-nil, zero value otherwise.

### GetUserDraftIdOk

`func (o *Item) GetUserDraftIdOk() (*int64, bool)`

GetUserDraftIdOk returns a tuple with the UserDraftId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDraftId

`func (o *Item) SetUserDraftId(v int64)`

SetUserDraftId sets UserDraftId field to given value.

### HasUserDraftId

`func (o *Item) HasUserDraftId() bool`

HasUserDraftId returns a boolean if a field has been set.

### GetStaticStart

`func (o *Item) GetStaticStart() bool`

GetStaticStart returns the StaticStart field if non-nil, zero value otherwise.

### GetStaticStartOk

`func (o *Item) GetStaticStartOk() (*bool, bool)`

GetStaticStartOk returns a tuple with the StaticStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticStart

`func (o *Item) SetStaticStart(v bool)`

SetStaticStart sets StaticStart field to given value.

### HasStaticStart

`func (o *Item) HasStaticStart() bool`

HasStaticStart returns a boolean if a field has been set.

### GetDetails

`func (o *Item) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *Item) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *Item) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *Item) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


