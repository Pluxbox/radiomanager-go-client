# ModelTypeResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**Options** | Pointer to [**ModelTypeOptions**](ModelTypeOptions.md) |  | [optional] 
**Order** | Pointer to **int64** |  | [optional] 
**DeletedAt** | Pointer to **time.Time** |  | [optional] 
**Campaigns** | Pointer to [**ModelTypeRelationsCampaigns**](ModelTypeRelationsCampaigns.md) |  | [optional] 
**Broadcasts** | Pointer to [**ModelTypeRelationsBroadcasts**](ModelTypeRelationsBroadcasts.md) |  | [optional] 
**Programs** | Pointer to [**ModelTypeRelationsPrograms**](ModelTypeRelationsPrograms.md) |  | [optional] 
**Contacts** | Pointer to [**ModelTypeRelationsContacts**](ModelTypeRelationsContacts.md) |  | [optional] 
**Presenters** | Pointer to [**ModelTypeRelationsPresenters**](ModelTypeRelationsPresenters.md) |  | [optional] 
**Items** | Pointer to [**ModelTypeRelationsItems**](ModelTypeRelationsItems.md) |  | [optional] 
**ExternalStationId** | Pointer to **int64** |  | [optional] 

## Methods

### NewModelTypeResult

`func NewModelTypeResult(id int64, ) *ModelTypeResult`

NewModelTypeResult instantiates a new ModelTypeResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelTypeResultWithDefaults

`func NewModelTypeResultWithDefaults() *ModelTypeResult`

NewModelTypeResultWithDefaults instantiates a new ModelTypeResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelTypeResult) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelTypeResult) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelTypeResult) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *ModelTypeResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelTypeResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelTypeResult) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelTypeResult) HasName() bool`

HasName returns a boolean if a field has been set.

### GetModel

`func (o *ModelTypeResult) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ModelTypeResult) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ModelTypeResult) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ModelTypeResult) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ModelTypeResult) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ModelTypeResult) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ModelTypeResult) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ModelTypeResult) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ModelTypeResult) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ModelTypeResult) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ModelTypeResult) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ModelTypeResult) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetOptions

`func (o *ModelTypeResult) GetOptions() ModelTypeOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *ModelTypeResult) GetOptionsOk() (*ModelTypeOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *ModelTypeResult) SetOptions(v ModelTypeOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *ModelTypeResult) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetOrder

`func (o *ModelTypeResult) GetOrder() int64`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *ModelTypeResult) GetOrderOk() (*int64, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *ModelTypeResult) SetOrder(v int64)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *ModelTypeResult) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetDeletedAt

`func (o *ModelTypeResult) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *ModelTypeResult) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *ModelTypeResult) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *ModelTypeResult) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetCampaigns

`func (o *ModelTypeResult) GetCampaigns() ModelTypeRelationsCampaigns`

GetCampaigns returns the Campaigns field if non-nil, zero value otherwise.

### GetCampaignsOk

`func (o *ModelTypeResult) GetCampaignsOk() (*ModelTypeRelationsCampaigns, bool)`

GetCampaignsOk returns a tuple with the Campaigns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaigns

`func (o *ModelTypeResult) SetCampaigns(v ModelTypeRelationsCampaigns)`

SetCampaigns sets Campaigns field to given value.

### HasCampaigns

`func (o *ModelTypeResult) HasCampaigns() bool`

HasCampaigns returns a boolean if a field has been set.

### GetBroadcasts

`func (o *ModelTypeResult) GetBroadcasts() ModelTypeRelationsBroadcasts`

GetBroadcasts returns the Broadcasts field if non-nil, zero value otherwise.

### GetBroadcastsOk

`func (o *ModelTypeResult) GetBroadcastsOk() (*ModelTypeRelationsBroadcasts, bool)`

GetBroadcastsOk returns a tuple with the Broadcasts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBroadcasts

`func (o *ModelTypeResult) SetBroadcasts(v ModelTypeRelationsBroadcasts)`

SetBroadcasts sets Broadcasts field to given value.

### HasBroadcasts

`func (o *ModelTypeResult) HasBroadcasts() bool`

HasBroadcasts returns a boolean if a field has been set.

### GetPrograms

`func (o *ModelTypeResult) GetPrograms() ModelTypeRelationsPrograms`

GetPrograms returns the Programs field if non-nil, zero value otherwise.

### GetProgramsOk

`func (o *ModelTypeResult) GetProgramsOk() (*ModelTypeRelationsPrograms, bool)`

GetProgramsOk returns a tuple with the Programs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrograms

`func (o *ModelTypeResult) SetPrograms(v ModelTypeRelationsPrograms)`

SetPrograms sets Programs field to given value.

### HasPrograms

`func (o *ModelTypeResult) HasPrograms() bool`

HasPrograms returns a boolean if a field has been set.

### GetContacts

`func (o *ModelTypeResult) GetContacts() ModelTypeRelationsContacts`

GetContacts returns the Contacts field if non-nil, zero value otherwise.

### GetContactsOk

`func (o *ModelTypeResult) GetContactsOk() (*ModelTypeRelationsContacts, bool)`

GetContactsOk returns a tuple with the Contacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContacts

`func (o *ModelTypeResult) SetContacts(v ModelTypeRelationsContacts)`

SetContacts sets Contacts field to given value.

### HasContacts

`func (o *ModelTypeResult) HasContacts() bool`

HasContacts returns a boolean if a field has been set.

### GetPresenters

`func (o *ModelTypeResult) GetPresenters() ModelTypeRelationsPresenters`

GetPresenters returns the Presenters field if non-nil, zero value otherwise.

### GetPresentersOk

`func (o *ModelTypeResult) GetPresentersOk() (*ModelTypeRelationsPresenters, bool)`

GetPresentersOk returns a tuple with the Presenters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresenters

`func (o *ModelTypeResult) SetPresenters(v ModelTypeRelationsPresenters)`

SetPresenters sets Presenters field to given value.

### HasPresenters

`func (o *ModelTypeResult) HasPresenters() bool`

HasPresenters returns a boolean if a field has been set.

### GetItems

`func (o *ModelTypeResult) GetItems() ModelTypeRelationsItems`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ModelTypeResult) GetItemsOk() (*ModelTypeRelationsItems, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ModelTypeResult) SetItems(v ModelTypeRelationsItems)`

SetItems sets Items field to given value.

### HasItems

`func (o *ModelTypeResult) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetExternalStationId

`func (o *ModelTypeResult) GetExternalStationId() int64`

GetExternalStationId returns the ExternalStationId field if non-nil, zero value otherwise.

### GetExternalStationIdOk

`func (o *ModelTypeResult) GetExternalStationIdOk() (*int64, bool)`

GetExternalStationIdOk returns a tuple with the ExternalStationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalStationId

`func (o *ModelTypeResult) SetExternalStationId(v int64)`

SetExternalStationId sets ExternalStationId field to given value.

### HasExternalStationId

`func (o *ModelTypeResult) HasExternalStationId() bool`

HasExternalStationId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


