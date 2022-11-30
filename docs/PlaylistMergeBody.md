# PlaylistMergeBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Start** | Pointer to **time.Time** |  | [optional] 
**AllowPlaylistPast** | Pointer to **int32** |  | [optional] 
**BroadcastId** | Pointer to **int64** |  | [optional] 
**Items** | Pointer to [**[]ImportItem**](ImportItem.md) |  | [optional] 

## Methods

### NewPlaylistMergeBody

`func NewPlaylistMergeBody() *PlaylistMergeBody`

NewPlaylistMergeBody instantiates a new PlaylistMergeBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlaylistMergeBodyWithDefaults

`func NewPlaylistMergeBodyWithDefaults() *PlaylistMergeBody`

NewPlaylistMergeBodyWithDefaults instantiates a new PlaylistMergeBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStart

`func (o *PlaylistMergeBody) GetStart() time.Time`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *PlaylistMergeBody) GetStartOk() (*time.Time, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *PlaylistMergeBody) SetStart(v time.Time)`

SetStart sets Start field to given value.

### HasStart

`func (o *PlaylistMergeBody) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetAllowPlaylistPast

`func (o *PlaylistMergeBody) GetAllowPlaylistPast() int32`

GetAllowPlaylistPast returns the AllowPlaylistPast field if non-nil, zero value otherwise.

### GetAllowPlaylistPastOk

`func (o *PlaylistMergeBody) GetAllowPlaylistPastOk() (*int32, bool)`

GetAllowPlaylistPastOk returns a tuple with the AllowPlaylistPast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowPlaylistPast

`func (o *PlaylistMergeBody) SetAllowPlaylistPast(v int32)`

SetAllowPlaylistPast sets AllowPlaylistPast field to given value.

### HasAllowPlaylistPast

`func (o *PlaylistMergeBody) HasAllowPlaylistPast() bool`

HasAllowPlaylistPast returns a boolean if a field has been set.

### GetBroadcastId

`func (o *PlaylistMergeBody) GetBroadcastId() int64`

GetBroadcastId returns the BroadcastId field if non-nil, zero value otherwise.

### GetBroadcastIdOk

`func (o *PlaylistMergeBody) GetBroadcastIdOk() (*int64, bool)`

GetBroadcastIdOk returns a tuple with the BroadcastId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBroadcastId

`func (o *PlaylistMergeBody) SetBroadcastId(v int64)`

SetBroadcastId sets BroadcastId field to given value.

### HasBroadcastId

`func (o *PlaylistMergeBody) HasBroadcastId() bool`

HasBroadcastId returns a boolean if a field has been set.

### GetItems

`func (o *PlaylistMergeBody) GetItems() []ImportItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *PlaylistMergeBody) GetItemsOk() (*[]ImportItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *PlaylistMergeBody) SetItems(v []ImportItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *PlaylistMergeBody) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


