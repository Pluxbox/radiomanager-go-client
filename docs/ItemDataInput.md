# ItemDataInput

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
**PreviousId** | Pointer to **int64** |  | [optional] 
**BeforeId** | Pointer to **int64** |  | [optional] 
**Contacts** | Pointer to **[]int32** |  | [optional] 
**Tags** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewItemDataInput

`func NewItemDataInput(modelTypeId int64, ) *ItemDataInput`

NewItemDataInput instantiates a new ItemDataInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemDataInputWithDefaults

`func NewItemDataInputWithDefaults() *ItemDataInput`

NewItemDataInputWithDefaults instantiates a new ItemDataInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModelTypeId

`func (o *ItemDataInput) GetModelTypeId() int64`

GetModelTypeId returns the ModelTypeId field if non-nil, zero value otherwise.

### GetModelTypeIdOk

`func (o *ItemDataInput) GetModelTypeIdOk() (*int64, bool)`

GetModelTypeIdOk returns a tuple with the ModelTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelTypeId

`func (o *ItemDataInput) SetModelTypeId(v int64)`

SetModelTypeId sets ModelTypeId field to given value.


### GetBlockId

`func (o *ItemDataInput) GetBlockId() int64`

GetBlockId returns the BlockId field if non-nil, zero value otherwise.

### GetBlockIdOk

`func (o *ItemDataInput) GetBlockIdOk() (*int64, bool)`

GetBlockIdOk returns a tuple with the BlockId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockId

`func (o *ItemDataInput) SetBlockId(v int64)`

SetBlockId sets BlockId field to given value.

### HasBlockId

`func (o *ItemDataInput) HasBlockId() bool`

HasBlockId returns a boolean if a field has been set.

### GetExternalId

`func (o *ItemDataInput) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ItemDataInput) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ItemDataInput) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ItemDataInput) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetFieldValues

`func (o *ItemDataInput) GetFieldValues() map[string]interface{}`

GetFieldValues returns the FieldValues field if non-nil, zero value otherwise.

### GetFieldValuesOk

`func (o *ItemDataInput) GetFieldValuesOk() (*map[string]interface{}, bool)`

GetFieldValuesOk returns a tuple with the FieldValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldValues

`func (o *ItemDataInput) SetFieldValues(v map[string]interface{})`

SetFieldValues sets FieldValues field to given value.

### HasFieldValues

`func (o *ItemDataInput) HasFieldValues() bool`

HasFieldValues returns a boolean if a field has been set.

### GetTitle

`func (o *ItemDataInput) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ItemDataInput) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ItemDataInput) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ItemDataInput) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDuration

`func (o *ItemDataInput) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *ItemDataInput) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *ItemDataInput) SetDuration(v int64)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *ItemDataInput) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetStart

`func (o *ItemDataInput) GetStart() time.Time`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *ItemDataInput) GetStartOk() (*time.Time, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *ItemDataInput) SetStart(v time.Time)`

SetStart sets Start field to given value.

### HasStart

`func (o *ItemDataInput) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetStatus

`func (o *ItemDataInput) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ItemDataInput) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ItemDataInput) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ItemDataInput) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetImport

`func (o *ItemDataInput) GetImport() int64`

GetImport returns the Import field if non-nil, zero value otherwise.

### GetImportOk

`func (o *ItemDataInput) GetImportOk() (*int64, bool)`

GetImportOk returns a tuple with the Import field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImport

`func (o *ItemDataInput) SetImport(v int64)`

SetImport sets Import field to given value.

### HasImport

`func (o *ItemDataInput) HasImport() bool`

HasImport returns a boolean if a field has been set.

### GetCampaignId

`func (o *ItemDataInput) GetCampaignId() int64`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *ItemDataInput) GetCampaignIdOk() (*int64, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *ItemDataInput) SetCampaignId(v int64)`

SetCampaignId sets CampaignId field to given value.

### HasCampaignId

`func (o *ItemDataInput) HasCampaignId() bool`

HasCampaignId returns a boolean if a field has been set.

### GetRecommended

`func (o *ItemDataInput) GetRecommended() bool`

GetRecommended returns the Recommended field if non-nil, zero value otherwise.

### GetRecommendedOk

`func (o *ItemDataInput) GetRecommendedOk() (*bool, bool)`

GetRecommendedOk returns a tuple with the Recommended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommended

`func (o *ItemDataInput) SetRecommended(v bool)`

SetRecommended sets Recommended field to given value.

### HasRecommended

`func (o *ItemDataInput) HasRecommended() bool`

HasRecommended returns a boolean if a field has been set.

### GetStationDraftId

`func (o *ItemDataInput) GetStationDraftId() int64`

GetStationDraftId returns the StationDraftId field if non-nil, zero value otherwise.

### GetStationDraftIdOk

`func (o *ItemDataInput) GetStationDraftIdOk() (*int64, bool)`

GetStationDraftIdOk returns a tuple with the StationDraftId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStationDraftId

`func (o *ItemDataInput) SetStationDraftId(v int64)`

SetStationDraftId sets StationDraftId field to given value.

### HasStationDraftId

`func (o *ItemDataInput) HasStationDraftId() bool`

HasStationDraftId returns a boolean if a field has been set.

### GetProgramDraftId

`func (o *ItemDataInput) GetProgramDraftId() int64`

GetProgramDraftId returns the ProgramDraftId field if non-nil, zero value otherwise.

### GetProgramDraftIdOk

`func (o *ItemDataInput) GetProgramDraftIdOk() (*int64, bool)`

GetProgramDraftIdOk returns a tuple with the ProgramDraftId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgramDraftId

`func (o *ItemDataInput) SetProgramDraftId(v int64)`

SetProgramDraftId sets ProgramDraftId field to given value.

### HasProgramDraftId

`func (o *ItemDataInput) HasProgramDraftId() bool`

HasProgramDraftId returns a boolean if a field has been set.

### GetUserDraftId

`func (o *ItemDataInput) GetUserDraftId() int64`

GetUserDraftId returns the UserDraftId field if non-nil, zero value otherwise.

### GetUserDraftIdOk

`func (o *ItemDataInput) GetUserDraftIdOk() (*int64, bool)`

GetUserDraftIdOk returns a tuple with the UserDraftId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDraftId

`func (o *ItemDataInput) SetUserDraftId(v int64)`

SetUserDraftId sets UserDraftId field to given value.

### HasUserDraftId

`func (o *ItemDataInput) HasUserDraftId() bool`

HasUserDraftId returns a boolean if a field has been set.

### GetStaticStart

`func (o *ItemDataInput) GetStaticStart() bool`

GetStaticStart returns the StaticStart field if non-nil, zero value otherwise.

### GetStaticStartOk

`func (o *ItemDataInput) GetStaticStartOk() (*bool, bool)`

GetStaticStartOk returns a tuple with the StaticStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticStart

`func (o *ItemDataInput) SetStaticStart(v bool)`

SetStaticStart sets StaticStart field to given value.

### HasStaticStart

`func (o *ItemDataInput) HasStaticStart() bool`

HasStaticStart returns a boolean if a field has been set.

### GetDetails

`func (o *ItemDataInput) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *ItemDataInput) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *ItemDataInput) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *ItemDataInput) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetPreviousId

`func (o *ItemDataInput) GetPreviousId() int64`

GetPreviousId returns the PreviousId field if non-nil, zero value otherwise.

### GetPreviousIdOk

`func (o *ItemDataInput) GetPreviousIdOk() (*int64, bool)`

GetPreviousIdOk returns a tuple with the PreviousId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousId

`func (o *ItemDataInput) SetPreviousId(v int64)`

SetPreviousId sets PreviousId field to given value.

### HasPreviousId

`func (o *ItemDataInput) HasPreviousId() bool`

HasPreviousId returns a boolean if a field has been set.

### GetBeforeId

`func (o *ItemDataInput) GetBeforeId() int64`

GetBeforeId returns the BeforeId field if non-nil, zero value otherwise.

### GetBeforeIdOk

`func (o *ItemDataInput) GetBeforeIdOk() (*int64, bool)`

GetBeforeIdOk returns a tuple with the BeforeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeforeId

`func (o *ItemDataInput) SetBeforeId(v int64)`

SetBeforeId sets BeforeId field to given value.

### HasBeforeId

`func (o *ItemDataInput) HasBeforeId() bool`

HasBeforeId returns a boolean if a field has been set.

### GetContacts

`func (o *ItemDataInput) GetContacts() []int32`

GetContacts returns the Contacts field if non-nil, zero value otherwise.

### GetContactsOk

`func (o *ItemDataInput) GetContactsOk() (*[]int32, bool)`

GetContactsOk returns a tuple with the Contacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContacts

`func (o *ItemDataInput) SetContacts(v []int32)`

SetContacts sets Contacts field to given value.

### HasContacts

`func (o *ItemDataInput) HasContacts() bool`

HasContacts returns a boolean if a field has been set.

### GetTags

`func (o *ItemDataInput) GetTags() []int32`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ItemDataInput) GetTagsOk() (*[]int32, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ItemDataInput) SetTags(v []int32)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ItemDataInput) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


