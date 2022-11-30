# PlaylistTimingBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Start** | Pointer to **time.Time** |  | [optional] 
**AllowPlaylistPast** | Pointer to **int32** |  | [optional] 
**BroadcastId** | Pointer to **int64** |  | [optional] 
**Items** | Pointer to [**[]ImportItem**](ImportItem.md) |  | [optional] 

## Methods

### NewPlaylistTimingBody

`func NewPlaylistTimingBody() *PlaylistTimingBody`

NewPlaylistTimingBody instantiates a new PlaylistTimingBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlaylistTimingBodyWithDefaults

`func NewPlaylistTimingBodyWithDefaults() *PlaylistTimingBody`

NewPlaylistTimingBodyWithDefaults instantiates a new PlaylistTimingBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStart

`func (o *PlaylistTimingBody) GetStart() time.Time`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *PlaylistTimingBody) GetStartOk() (*time.Time, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *PlaylistTimingBody) SetStart(v time.Time)`

SetStart sets Start field to given value.

### HasStart

`func (o *PlaylistTimingBody) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetAllowPlaylistPast

`func (o *PlaylistTimingBody) GetAllowPlaylistPast() int32`

GetAllowPlaylistPast returns the AllowPlaylistPast field if non-nil, zero value otherwise.

### GetAllowPlaylistPastOk

`func (o *PlaylistTimingBody) GetAllowPlaylistPastOk() (*int32, bool)`

GetAllowPlaylistPastOk returns a tuple with the AllowPlaylistPast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowPlaylistPast

`func (o *PlaylistTimingBody) SetAllowPlaylistPast(v int32)`

SetAllowPlaylistPast sets AllowPlaylistPast field to given value.

### HasAllowPlaylistPast

`func (o *PlaylistTimingBody) HasAllowPlaylistPast() bool`

HasAllowPlaylistPast returns a boolean if a field has been set.

### GetBroadcastId

`func (o *PlaylistTimingBody) GetBroadcastId() int64`

GetBroadcastId returns the BroadcastId field if non-nil, zero value otherwise.

### GetBroadcastIdOk

`func (o *PlaylistTimingBody) GetBroadcastIdOk() (*int64, bool)`

GetBroadcastIdOk returns a tuple with the BroadcastId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBroadcastId

`func (o *PlaylistTimingBody) SetBroadcastId(v int64)`

SetBroadcastId sets BroadcastId field to given value.

### HasBroadcastId

`func (o *PlaylistTimingBody) HasBroadcastId() bool`

HasBroadcastId returns a boolean if a field has been set.

### GetItems

`func (o *PlaylistTimingBody) GetItems() []ImportItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *PlaylistTimingBody) GetItemsOk() (*[]ImportItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *PlaylistTimingBody) SetItems(v []ImportItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *PlaylistTimingBody) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


