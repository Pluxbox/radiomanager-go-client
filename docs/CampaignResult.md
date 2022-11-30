# CampaignResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**UpdatedAt** | **time.Time** |  | 
**CreatedAt** | **time.Time** |  | 
**DeletedAt** | **time.Time** |  | 
**Item** | Pointer to [**CampaignTemplateItem**](CampaignTemplateItem.md) |  | [optional] 
**ExternalStationId** | Pointer to **int64** |  | [optional] 
**ModelTypeId** | **int64** |  | 
**FieldValues** | Pointer to **map[string]interface{}** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Start** | **time.Time** |  | 
**Stop** | **time.Time** |  | 
**Recommended** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Items** | Pointer to [**CampaignRelationsItems**](CampaignRelationsItems.md) |  | [optional] 
**ModelType** | Pointer to [**BroadcastRelationsModelType**](BroadcastRelationsModelType.md) |  | [optional] 

## Methods

### NewCampaignResult

`func NewCampaignResult(id int64, updatedAt time.Time, createdAt time.Time, deletedAt time.Time, modelTypeId int64, start time.Time, stop time.Time, ) *CampaignResult`

NewCampaignResult instantiates a new CampaignResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignResultWithDefaults

`func NewCampaignResultWithDefaults() *CampaignResult`

NewCampaignResultWithDefaults instantiates a new CampaignResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CampaignResult) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CampaignResult) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CampaignResult) SetId(v int64)`

SetId sets Id field to given value.


### GetUpdatedAt

`func (o *CampaignResult) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CampaignResult) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CampaignResult) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetCreatedAt

`func (o *CampaignResult) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CampaignResult) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CampaignResult) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDeletedAt

`func (o *CampaignResult) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *CampaignResult) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *CampaignResult) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.


### GetItem

`func (o *CampaignResult) GetItem() CampaignTemplateItem`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *CampaignResult) GetItemOk() (*CampaignTemplateItem, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *CampaignResult) SetItem(v CampaignTemplateItem)`

SetItem sets Item field to given value.

### HasItem

`func (o *CampaignResult) HasItem() bool`

HasItem returns a boolean if a field has been set.

### GetExternalStationId

`func (o *CampaignResult) GetExternalStationId() int64`

GetExternalStationId returns the ExternalStationId field if non-nil, zero value otherwise.

### GetExternalStationIdOk

`func (o *CampaignResult) GetExternalStationIdOk() (*int64, bool)`

GetExternalStationIdOk returns a tuple with the ExternalStationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalStationId

`func (o *CampaignResult) SetExternalStationId(v int64)`

SetExternalStationId sets ExternalStationId field to given value.

### HasExternalStationId

`func (o *CampaignResult) HasExternalStationId() bool`

HasExternalStationId returns a boolean if a field has been set.

### GetModelTypeId

`func (o *CampaignResult) GetModelTypeId() int64`

GetModelTypeId returns the ModelTypeId field if non-nil, zero value otherwise.

### GetModelTypeIdOk

`func (o *CampaignResult) GetModelTypeIdOk() (*int64, bool)`

GetModelTypeIdOk returns a tuple with the ModelTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelTypeId

`func (o *CampaignResult) SetModelTypeId(v int64)`

SetModelTypeId sets ModelTypeId field to given value.


### GetFieldValues

`func (o *CampaignResult) GetFieldValues() map[string]interface{}`

GetFieldValues returns the FieldValues field if non-nil, zero value otherwise.

### GetFieldValuesOk

`func (o *CampaignResult) GetFieldValuesOk() (*map[string]interface{}, bool)`

GetFieldValuesOk returns a tuple with the FieldValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldValues

`func (o *CampaignResult) SetFieldValues(v map[string]interface{})`

SetFieldValues sets FieldValues field to given value.

### HasFieldValues

`func (o *CampaignResult) HasFieldValues() bool`

HasFieldValues returns a boolean if a field has been set.

### GetTitle

`func (o *CampaignResult) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CampaignResult) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CampaignResult) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CampaignResult) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetStart

`func (o *CampaignResult) GetStart() time.Time`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *CampaignResult) GetStartOk() (*time.Time, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *CampaignResult) SetStart(v time.Time)`

SetStart sets Start field to given value.


### GetStop

`func (o *CampaignResult) GetStop() time.Time`

GetStop returns the Stop field if non-nil, zero value otherwise.

### GetStopOk

`func (o *CampaignResult) GetStopOk() (*time.Time, bool)`

GetStopOk returns a tuple with the Stop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStop

`func (o *CampaignResult) SetStop(v time.Time)`

SetStop sets Stop field to given value.


### GetRecommended

`func (o *CampaignResult) GetRecommended() bool`

GetRecommended returns the Recommended field if non-nil, zero value otherwise.

### GetRecommendedOk

`func (o *CampaignResult) GetRecommendedOk() (*bool, bool)`

GetRecommendedOk returns a tuple with the Recommended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommended

`func (o *CampaignResult) SetRecommended(v bool)`

SetRecommended sets Recommended field to given value.

### HasRecommended

`func (o *CampaignResult) HasRecommended() bool`

HasRecommended returns a boolean if a field has been set.

### GetDescription

`func (o *CampaignResult) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CampaignResult) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CampaignResult) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CampaignResult) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetItems

`func (o *CampaignResult) GetItems() CampaignRelationsItems`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *CampaignResult) GetItemsOk() (*CampaignRelationsItems, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *CampaignResult) SetItems(v CampaignRelationsItems)`

SetItems sets Items field to given value.

### HasItems

`func (o *CampaignResult) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetModelType

`func (o *CampaignResult) GetModelType() BroadcastRelationsModelType`

GetModelType returns the ModelType field if non-nil, zero value otherwise.

### GetModelTypeOk

`func (o *CampaignResult) GetModelTypeOk() (*BroadcastRelationsModelType, bool)`

GetModelTypeOk returns a tuple with the ModelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelType

`func (o *CampaignResult) SetModelType(v BroadcastRelationsModelType)`

SetModelType sets ModelType field to given value.

### HasModelType

`func (o *CampaignResult) HasModelType() bool`

HasModelType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


