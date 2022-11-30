# ItemResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**DeletedAt** | Pointer to **time.Time** |  | [optional] 
**DataModified** | Pointer to **int64** |  | [optional] 
**Order** | Pointer to **int64** |  | [optional] 
**TemplateblockId** | Pointer to **int64** |  | [optional] 
**TemplateitemId** | Pointer to **int64** |  | [optional] 
**ExternalStationId** | Pointer to **int64** |  | [optional] 
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
**Block** | Pointer to [**ItemRelationsBlock**](ItemRelationsBlock.md) |  | [optional] 
**Broadcast** | Pointer to [**BlockRelationsBroadcast**](BlockRelationsBroadcast.md) |  | [optional] 
**Program** | Pointer to [**ItemRelationsProgram**](ItemRelationsProgram.md) |  | [optional] 
**Contacts** | Pointer to [**ItemRelationsContacts**](ItemRelationsContacts.md) |  | [optional] 
**Tags** | Pointer to [**ItemRelationsTags**](ItemRelationsTags.md) |  | [optional] 
**Campaign** | Pointer to [**ItemRelationsCampaign**](ItemRelationsCampaign.md) |  | [optional] 
**ModelType** | Pointer to [**BroadcastRelationsModelType**](BroadcastRelationsModelType.md) |  | [optional] 

## Methods

### NewItemResult

`func NewItemResult(modelTypeId int64, ) *ItemResult`

NewItemResult instantiates a new ItemResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemResultWithDefaults

`func NewItemResultWithDefaults() *ItemResult`

NewItemResultWithDefaults instantiates a new ItemResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ItemResult) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ItemResult) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ItemResult) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ItemResult) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ItemResult) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ItemResult) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ItemResult) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ItemResult) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ItemResult) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ItemResult) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ItemResult) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ItemResult) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *ItemResult) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *ItemResult) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *ItemResult) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *ItemResult) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetDataModified

`func (o *ItemResult) GetDataModified() int64`

GetDataModified returns the DataModified field if non-nil, zero value otherwise.

### GetDataModifiedOk

`func (o *ItemResult) GetDataModifiedOk() (*int64, bool)`

GetDataModifiedOk returns a tuple with the DataModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataModified

`func (o *ItemResult) SetDataModified(v int64)`

SetDataModified sets DataModified field to given value.

### HasDataModified

`func (o *ItemResult) HasDataModified() bool`

HasDataModified returns a boolean if a field has been set.

### GetOrder

`func (o *ItemResult) GetOrder() int64`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *ItemResult) GetOrderOk() (*int64, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *ItemResult) SetOrder(v int64)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *ItemResult) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetTemplateblockId

`func (o *ItemResult) GetTemplateblockId() int64`

GetTemplateblockId returns the TemplateblockId field if non-nil, zero value otherwise.

### GetTemplateblockIdOk

`func (o *ItemResult) GetTemplateblockIdOk() (*int64, bool)`

GetTemplateblockIdOk returns a tuple with the TemplateblockId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateblockId

`func (o *ItemResult) SetTemplateblockId(v int64)`

SetTemplateblockId sets TemplateblockId field to given value.

### HasTemplateblockId

`func (o *ItemResult) HasTemplateblockId() bool`

HasTemplateblockId returns a boolean if a field has been set.

### GetTemplateitemId

`func (o *ItemResult) GetTemplateitemId() int64`

GetTemplateitemId returns the TemplateitemId field if non-nil, zero value otherwise.

### GetTemplateitemIdOk

`func (o *ItemResult) GetTemplateitemIdOk() (*int64, bool)`

GetTemplateitemIdOk returns a tuple with the TemplateitemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateitemId

`func (o *ItemResult) SetTemplateitemId(v int64)`

SetTemplateitemId sets TemplateitemId field to given value.

### HasTemplateitemId

`func (o *ItemResult) HasTemplateitemId() bool`

HasTemplateitemId returns a boolean if a field has been set.

### GetExternalStationId

`func (o *ItemResult) GetExternalStationId() int64`

GetExternalStationId returns the ExternalStationId field if non-nil, zero value otherwise.

### GetExternalStationIdOk

`func (o *ItemResult) GetExternalStationIdOk() (*int64, bool)`

GetExternalStationIdOk returns a tuple with the ExternalStationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalStationId

`func (o *ItemResult) SetExternalStationId(v int64)`

SetExternalStationId sets ExternalStationId field to given value.

### HasExternalStationId

`func (o *ItemResult) HasExternalStationId() bool`

HasExternalStationId returns a boolean if a field has been set.

### GetModelTypeId

`func (o *ItemResult) GetModelTypeId() int64`

GetModelTypeId returns the ModelTypeId field if non-nil, zero value otherwise.

### GetModelTypeIdOk

`func (o *ItemResult) GetModelTypeIdOk() (*int64, bool)`

GetModelTypeIdOk returns a tuple with the ModelTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelTypeId

`func (o *ItemResult) SetModelTypeId(v int64)`

SetModelTypeId sets ModelTypeId field to given value.


### GetBlockId

`func (o *ItemResult) GetBlockId() int64`

GetBlockId returns the BlockId field if non-nil, zero value otherwise.

### GetBlockIdOk

`func (o *ItemResult) GetBlockIdOk() (*int64, bool)`

GetBlockIdOk returns a tuple with the BlockId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockId

`func (o *ItemResult) SetBlockId(v int64)`

SetBlockId sets BlockId field to given value.

### HasBlockId

`func (o *ItemResult) HasBlockId() bool`

HasBlockId returns a boolean if a field has been set.

### GetExternalId

`func (o *ItemResult) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ItemResult) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ItemResult) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ItemResult) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetFieldValues

`func (o *ItemResult) GetFieldValues() map[string]interface{}`

GetFieldValues returns the FieldValues field if non-nil, zero value otherwise.

### GetFieldValuesOk

`func (o *ItemResult) GetFieldValuesOk() (*map[string]interface{}, bool)`

GetFieldValuesOk returns a tuple with the FieldValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldValues

`func (o *ItemResult) SetFieldValues(v map[string]interface{})`

SetFieldValues sets FieldValues field to given value.

### HasFieldValues

`func (o *ItemResult) HasFieldValues() bool`

HasFieldValues returns a boolean if a field has been set.

### GetTitle

`func (o *ItemResult) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ItemResult) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ItemResult) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ItemResult) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDuration

`func (o *ItemResult) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *ItemResult) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *ItemResult) SetDuration(v int64)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *ItemResult) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetStart

`func (o *ItemResult) GetStart() time.Time`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *ItemResult) GetStartOk() (*time.Time, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *ItemResult) SetStart(v time.Time)`

SetStart sets Start field to given value.

### HasStart

`func (o *ItemResult) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetStatus

`func (o *ItemResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ItemResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ItemResult) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ItemResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetImport

`func (o *ItemResult) GetImport() int64`

GetImport returns the Import field if non-nil, zero value otherwise.

### GetImportOk

`func (o *ItemResult) GetImportOk() (*int64, bool)`

GetImportOk returns a tuple with the Import field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImport

`func (o *ItemResult) SetImport(v int64)`

SetImport sets Import field to given value.

### HasImport

`func (o *ItemResult) HasImport() bool`

HasImport returns a boolean if a field has been set.

### GetCampaignId

`func (o *ItemResult) GetCampaignId() int64`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *ItemResult) GetCampaignIdOk() (*int64, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *ItemResult) SetCampaignId(v int64)`

SetCampaignId sets CampaignId field to given value.

### HasCampaignId

`func (o *ItemResult) HasCampaignId() bool`

HasCampaignId returns a boolean if a field has been set.

### GetRecommended

`func (o *ItemResult) GetRecommended() bool`

GetRecommended returns the Recommended field if non-nil, zero value otherwise.

### GetRecommendedOk

`func (o *ItemResult) GetRecommendedOk() (*bool, bool)`

GetRecommendedOk returns a tuple with the Recommended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommended

`func (o *ItemResult) SetRecommended(v bool)`

SetRecommended sets Recommended field to given value.

### HasRecommended

`func (o *ItemResult) HasRecommended() bool`

HasRecommended returns a boolean if a field has been set.

### GetStationDraftId

`func (o *ItemResult) GetStationDraftId() int64`

GetStationDraftId returns the StationDraftId field if non-nil, zero value otherwise.

### GetStationDraftIdOk

`func (o *ItemResult) GetStationDraftIdOk() (*int64, bool)`

GetStationDraftIdOk returns a tuple with the StationDraftId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStationDraftId

`func (o *ItemResult) SetStationDraftId(v int64)`

SetStationDraftId sets StationDraftId field to given value.

### HasStationDraftId

`func (o *ItemResult) HasStationDraftId() bool`

HasStationDraftId returns a boolean if a field has been set.

### GetProgramDraftId

`func (o *ItemResult) GetProgramDraftId() int64`

GetProgramDraftId returns the ProgramDraftId field if non-nil, zero value otherwise.

### GetProgramDraftIdOk

`func (o *ItemResult) GetProgramDraftIdOk() (*int64, bool)`

GetProgramDraftIdOk returns a tuple with the ProgramDraftId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgramDraftId

`func (o *ItemResult) SetProgramDraftId(v int64)`

SetProgramDraftId sets ProgramDraftId field to given value.

### HasProgramDraftId

`func (o *ItemResult) HasProgramDraftId() bool`

HasProgramDraftId returns a boolean if a field has been set.

### GetUserDraftId

`func (o *ItemResult) GetUserDraftId() int64`

GetUserDraftId returns the UserDraftId field if non-nil, zero value otherwise.

### GetUserDraftIdOk

`func (o *ItemResult) GetUserDraftIdOk() (*int64, bool)`

GetUserDraftIdOk returns a tuple with the UserDraftId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDraftId

`func (o *ItemResult) SetUserDraftId(v int64)`

SetUserDraftId sets UserDraftId field to given value.

### HasUserDraftId

`func (o *ItemResult) HasUserDraftId() bool`

HasUserDraftId returns a boolean if a field has been set.

### GetStaticStart

`func (o *ItemResult) GetStaticStart() bool`

GetStaticStart returns the StaticStart field if non-nil, zero value otherwise.

### GetStaticStartOk

`func (o *ItemResult) GetStaticStartOk() (*bool, bool)`

GetStaticStartOk returns a tuple with the StaticStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticStart

`func (o *ItemResult) SetStaticStart(v bool)`

SetStaticStart sets StaticStart field to given value.

### HasStaticStart

`func (o *ItemResult) HasStaticStart() bool`

HasStaticStart returns a boolean if a field has been set.

### GetDetails

`func (o *ItemResult) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *ItemResult) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *ItemResult) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *ItemResult) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetBlock

`func (o *ItemResult) GetBlock() ItemRelationsBlock`

GetBlock returns the Block field if non-nil, zero value otherwise.

### GetBlockOk

`func (o *ItemResult) GetBlockOk() (*ItemRelationsBlock, bool)`

GetBlockOk returns a tuple with the Block field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlock

`func (o *ItemResult) SetBlock(v ItemRelationsBlock)`

SetBlock sets Block field to given value.

### HasBlock

`func (o *ItemResult) HasBlock() bool`

HasBlock returns a boolean if a field has been set.

### GetBroadcast

`func (o *ItemResult) GetBroadcast() BlockRelationsBroadcast`

GetBroadcast returns the Broadcast field if non-nil, zero value otherwise.

### GetBroadcastOk

`func (o *ItemResult) GetBroadcastOk() (*BlockRelationsBroadcast, bool)`

GetBroadcastOk returns a tuple with the Broadcast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBroadcast

`func (o *ItemResult) SetBroadcast(v BlockRelationsBroadcast)`

SetBroadcast sets Broadcast field to given value.

### HasBroadcast

`func (o *ItemResult) HasBroadcast() bool`

HasBroadcast returns a boolean if a field has been set.

### GetProgram

`func (o *ItemResult) GetProgram() ItemRelationsProgram`

GetProgram returns the Program field if non-nil, zero value otherwise.

### GetProgramOk

`func (o *ItemResult) GetProgramOk() (*ItemRelationsProgram, bool)`

GetProgramOk returns a tuple with the Program field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgram

`func (o *ItemResult) SetProgram(v ItemRelationsProgram)`

SetProgram sets Program field to given value.

### HasProgram

`func (o *ItemResult) HasProgram() bool`

HasProgram returns a boolean if a field has been set.

### GetContacts

`func (o *ItemResult) GetContacts() ItemRelationsContacts`

GetContacts returns the Contacts field if non-nil, zero value otherwise.

### GetContactsOk

`func (o *ItemResult) GetContactsOk() (*ItemRelationsContacts, bool)`

GetContactsOk returns a tuple with the Contacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContacts

`func (o *ItemResult) SetContacts(v ItemRelationsContacts)`

SetContacts sets Contacts field to given value.

### HasContacts

`func (o *ItemResult) HasContacts() bool`

HasContacts returns a boolean if a field has been set.

### GetTags

`func (o *ItemResult) GetTags() ItemRelationsTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ItemResult) GetTagsOk() (*ItemRelationsTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ItemResult) SetTags(v ItemRelationsTags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ItemResult) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCampaign

`func (o *ItemResult) GetCampaign() ItemRelationsCampaign`

GetCampaign returns the Campaign field if non-nil, zero value otherwise.

### GetCampaignOk

`func (o *ItemResult) GetCampaignOk() (*ItemRelationsCampaign, bool)`

GetCampaignOk returns a tuple with the Campaign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaign

`func (o *ItemResult) SetCampaign(v ItemRelationsCampaign)`

SetCampaign sets Campaign field to given value.

### HasCampaign

`func (o *ItemResult) HasCampaign() bool`

HasCampaign returns a boolean if a field has been set.

### GetModelType

`func (o *ItemResult) GetModelType() BroadcastRelationsModelType`

GetModelType returns the ModelType field if non-nil, zero value otherwise.

### GetModelTypeOk

`func (o *ItemResult) GetModelTypeOk() (*BroadcastRelationsModelType, bool)`

GetModelTypeOk returns a tuple with the ModelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelType

`func (o *ItemResult) SetModelType(v BroadcastRelationsModelType)`

SetModelType sets ModelType field to given value.

### HasModelType

`func (o *ItemResult) HasModelType() bool`

HasModelType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


