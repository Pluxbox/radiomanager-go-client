# UserResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**Email** | **string** |  | 
**Firstname** | **string** |  | 
**Lastname** | **string** |  | 
**Phone** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**Settings** | Pointer to [**UserResultSettings**](UserResultSettings.md) |  | [optional] 
**Language** | Pointer to **bool** |  | [optional] 
**ActiveExternalStationId** | Pointer to **int64** |  | [optional] 
**Roles** | Pointer to [**[]UserResultRoles**](UserResultRoles.md) |  | [optional] 
**Groups** | Pointer to [**[]UserResultGroups**](UserResultGroups.md) |  | [optional] 

## Methods

### NewUserResult

`func NewUserResult(id int64, email string, firstname string, lastname string, ) *UserResult`

NewUserResult instantiates a new UserResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserResultWithDefaults

`func NewUserResultWithDefaults() *UserResult`

NewUserResultWithDefaults instantiates a new UserResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserResult) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserResult) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserResult) SetId(v int64)`

SetId sets Id field to given value.


### GetEmail

`func (o *UserResult) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserResult) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserResult) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetFirstname

`func (o *UserResult) GetFirstname() string`

GetFirstname returns the Firstname field if non-nil, zero value otherwise.

### GetFirstnameOk

`func (o *UserResult) GetFirstnameOk() (*string, bool)`

GetFirstnameOk returns a tuple with the Firstname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstname

`func (o *UserResult) SetFirstname(v string)`

SetFirstname sets Firstname field to given value.


### GetLastname

`func (o *UserResult) GetLastname() string`

GetLastname returns the Lastname field if non-nil, zero value otherwise.

### GetLastnameOk

`func (o *UserResult) GetLastnameOk() (*string, bool)`

GetLastnameOk returns a tuple with the Lastname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastname

`func (o *UserResult) SetLastname(v string)`

SetLastname sets Lastname field to given value.


### GetPhone

`func (o *UserResult) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *UserResult) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *UserResult) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *UserResult) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetCreatedAt

`func (o *UserResult) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UserResult) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UserResult) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *UserResult) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *UserResult) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *UserResult) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *UserResult) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *UserResult) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetActive

`func (o *UserResult) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *UserResult) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *UserResult) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *UserResult) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetSettings

`func (o *UserResult) GetSettings() UserResultSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *UserResult) GetSettingsOk() (*UserResultSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *UserResult) SetSettings(v UserResultSettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *UserResult) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetLanguage

`func (o *UserResult) GetLanguage() bool`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *UserResult) GetLanguageOk() (*bool, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *UserResult) SetLanguage(v bool)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *UserResult) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetActiveExternalStationId

`func (o *UserResult) GetActiveExternalStationId() int64`

GetActiveExternalStationId returns the ActiveExternalStationId field if non-nil, zero value otherwise.

### GetActiveExternalStationIdOk

`func (o *UserResult) GetActiveExternalStationIdOk() (*int64, bool)`

GetActiveExternalStationIdOk returns a tuple with the ActiveExternalStationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveExternalStationId

`func (o *UserResult) SetActiveExternalStationId(v int64)`

SetActiveExternalStationId sets ActiveExternalStationId field to given value.

### HasActiveExternalStationId

`func (o *UserResult) HasActiveExternalStationId() bool`

HasActiveExternalStationId returns a boolean if a field has been set.

### GetRoles

`func (o *UserResult) GetRoles() []UserResultRoles`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *UserResult) GetRolesOk() (*[]UserResultRoles, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *UserResult) SetRoles(v []UserResultRoles)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *UserResult) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetGroups

`func (o *UserResult) GetGroups() []UserResultGroups`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *UserResult) GetGroupsOk() (*[]UserResultGroups, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *UserResult) SetGroups(v []UserResultGroups)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *UserResult) HasGroups() bool`

HasGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


