# CampaignOutputOnly

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**UpdatedAt** | **time.Time** |  | 
**CreatedAt** | **time.Time** |  | 
**DeletedAt** | **time.Time** |  | 
**Item** | Pointer to [**CampaignTemplateItem**](CampaignTemplateItem.md) |  | [optional] 
**ExternalStationId** | Pointer to **int64** |  | [optional] 

## Methods

### NewCampaignOutputOnly

`func NewCampaignOutputOnly(id int64, updatedAt time.Time, createdAt time.Time, deletedAt time.Time, ) *CampaignOutputOnly`

NewCampaignOutputOnly instantiates a new CampaignOutputOnly object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignOutputOnlyWithDefaults

`func NewCampaignOutputOnlyWithDefaults() *CampaignOutputOnly`

NewCampaignOutputOnlyWithDefaults instantiates a new CampaignOutputOnly object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CampaignOutputOnly) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CampaignOutputOnly) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CampaignOutputOnly) SetId(v int64)`

SetId sets Id field to given value.


### GetUpdatedAt

`func (o *CampaignOutputOnly) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CampaignOutputOnly) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CampaignOutputOnly) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetCreatedAt

`func (o *CampaignOutputOnly) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CampaignOutputOnly) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CampaignOutputOnly) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDeletedAt

`func (o *CampaignOutputOnly) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *CampaignOutputOnly) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *CampaignOutputOnly) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.


### GetItem

`func (o *CampaignOutputOnly) GetItem() CampaignTemplateItem`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *CampaignOutputOnly) GetItemOk() (*CampaignTemplateItem, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *CampaignOutputOnly) SetItem(v CampaignTemplateItem)`

SetItem sets Item field to given value.

### HasItem

`func (o *CampaignOutputOnly) HasItem() bool`

HasItem returns a boolean if a field has been set.

### GetExternalStationId

`func (o *CampaignOutputOnly) GetExternalStationId() int64`

GetExternalStationId returns the ExternalStationId field if non-nil, zero value otherwise.

### GetExternalStationIdOk

`func (o *CampaignOutputOnly) GetExternalStationIdOk() (*int64, bool)`

GetExternalStationIdOk returns a tuple with the ExternalStationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalStationId

`func (o *CampaignOutputOnly) SetExternalStationId(v int64)`

SetExternalStationId sets ExternalStationId field to given value.

### HasExternalStationId

`func (o *CampaignOutputOnly) HasExternalStationId() bool`

HasExternalStationId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


