# InviteUserData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** |  | 
**RoleIds** | **[]int32** |  | 

## Methods

### NewInviteUserData

`func NewInviteUserData(email string, roleIds []int32, ) *InviteUserData`

NewInviteUserData instantiates a new InviteUserData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInviteUserDataWithDefaults

`func NewInviteUserDataWithDefaults() *InviteUserData`

NewInviteUserDataWithDefaults instantiates a new InviteUserData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *InviteUserData) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *InviteUserData) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *InviteUserData) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetRoleIds

`func (o *InviteUserData) GetRoleIds() []int32`

GetRoleIds returns the RoleIds field if non-nil, zero value otherwise.

### GetRoleIdsOk

`func (o *InviteUserData) GetRoleIdsOk() (*[]int32, bool)`

GetRoleIdsOk returns a tuple with the RoleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleIds

`func (o *InviteUserData) SetRoleIds(v []int32)`

SetRoleIds sets RoleIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


