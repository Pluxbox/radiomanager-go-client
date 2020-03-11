# \ItemApi

All URIs are relative to *https://radiomanager.io/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateItem**](ItemApi.md#CreateItem) | **Post** /items | Create an new item.
[**CurrentItemPostStructure**](ItemApi.md#CurrentItemPostStructure) | **Post** /items/current/structure | Post a current playing item, keep structure
[**CurrentItemPostTiming**](ItemApi.md#CurrentItemPostTiming) | **Post** /items/current/timing | Post a current playing item
[**DeleteItemById**](ItemApi.md#DeleteItemById) | **Delete** /items/{id} | Delete item by ID.
[**GetCurrentItem**](ItemApi.md#GetCurrentItem) | **Get** /items/current | Get current Item
[**GetItemById**](ItemApi.md#GetItemById) | **Get** /items/{id} | Get extended item details by ID.
[**ListItems**](ItemApi.md#ListItems) | **Get** /items | Get a list of all the items currently in your station.
[**PlaylistPostMerge**](ItemApi.md#PlaylistPostMerge) | **Post** /items/playlist/merge | Post a playlist, do not remove previously imported items
[**PlaylistPostStructure**](ItemApi.md#PlaylistPostStructure) | **Post** /items/playlist/structure | Post a playlist, keep current structure
[**PlaylistPostTiming**](ItemApi.md#PlaylistPostTiming) | **Post** /items/playlist/timing | Post a playlist
[**StopCurrentItem**](ItemApi.md#StopCurrentItem) | **Post** /items/stopcurrent | Stop an Item
[**UpdateItemById**](ItemApi.md#UpdateItemById) | **Patch** /items/{id} | Update extended item details by ID.



## CreateItem

> PostSuccess CreateItem(ctx, optional)

Create an new item.

Create item.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateItemOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateItemOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data** | [**optional.Interface of ItemDataInput**](ItemDataInput.md)| Data *(Optional)* | 

### Return type

[**PostSuccess**](PostSuccess.md)

### Authorization

[API-Key](../README.md#API-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CurrentItemPostStructure

> Success CurrentItemPostStructure(ctx, optional)

Post a current playing item, keep structure

Post a current playing item, keep structure

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CurrentItemPostStructureOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CurrentItemPostStructureOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data** | [**optional.Interface of ImportItem**](ImportItem.md)| Data *(Optional)* | 

### Return type

[**Success**](Success.md)

### Authorization

[API-Key](../README.md#API-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CurrentItemPostTiming

> Success CurrentItemPostTiming(ctx, optional)

Post a current playing item

Post a current playing item

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CurrentItemPostTimingOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CurrentItemPostTimingOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data** | [**optional.Interface of ImportItem**](ImportItem.md)| Data *(Optional)* | 

### Return type

[**Success**](Success.md)

### Authorization

[API-Key](../README.md#API-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteItemById

> Success DeleteItemById(ctx, id)

Delete item by ID.

Delete item by id.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64**| ID of Item **(Required)** | [default to 0]

### Return type

[**Success**](Success.md)

### Authorization

[API-Key](../README.md#API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCurrentItem

> ItemResult GetCurrentItem(ctx, optional)

Get current Item

Get current Item

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetCurrentItemOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetCurrentItemOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lastplayed** | **optional.Bool**| Show last played item if there is no current item*(Optional)* | 

### Return type

[**ItemResult**](ItemResult.md)

### Authorization

[API-Key](../README.md#API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItemById

> ItemResult GetItemById(ctx, id, optional)

Get extended item details by ID.

Read item by id.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64**| ID of Item **(Required)** | [default to 0]
 **optional** | ***GetItemByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetItemByIdOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **externalStationId** | **optional.Int64**| Query on a different (content providing) station *(Optional)* | 

### Return type

[**ItemResult**](ItemResult.md)

### Authorization

[API-Key](../README.md#API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListItems

> ItemResults ListItems(ctx, optional)

Get a list of all the items currently in your station.

Get a list of all the items currently in your station. This feature supports pagination and will give a maximum results of 50 items back.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListItemsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListItemsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int64**| Current page *(Optional)* | 
 **blockId** | **optional.Int64**| Search on Block ID *(Optional)* &#x60;(Relation)&#x60; | 
 **broadcastId** | **optional.Int64**| Search on Broadcast ID *(Optional)* &#x60;(Relation)&#x60; | 
 **modelTypeId** | **optional.Int64**| Search on ModelType ID *(Optional)* &#x60;(Relation)&#x60; | 
 **tagId** | **optional.Int64**| Search on Tag ID *(Optional)* &#x60;(Relation)&#x60; | 
 **campaignId** | **optional.Int64**| Search on Campaign ID *(Optional)* &#x60;(Relation)&#x60; | 
 **contactId** | **optional.Int64**| Search on Contact ID *(Optional)* &#x60;(Relation)&#x60; | 
 **programDraftId** | **optional.Int64**| Search on Program Draft ID *(Optional)* | 
 **userDraftId** | **optional.Int64**| Search on User Draft ID *(Optional)* | 
 **stationDraftId** | **optional.Int64**| Search on Station Draft ID *(Optional)* | 
 **programId** | **optional.Int64**| Search on Program ID *(Optional)* &#x60;(Relation)&#x60; | 
 **externalId** | **optional.String**| Search on External ID *(Optional)* | 
 **startMin** | **optional.Time**| Minimum start date *(Optional)* | 
 **startMax** | **optional.Time**| Maximum start date *(Optional)* | 
 **durationMin** | **optional.Int32**| Minimum duration (seconds) *(Optional)* | 
 **durationMax** | **optional.Int32**| Maximum duration (seconds) *(Optional)* | 
 **status** | **optional.String**| Play Status of item *(Optional)* | 
 **limit** | **optional.Int64**| Results per page *(Optional)* | 
 **orderBy** | **optional.String**| Field to order the results *(Optional)* | 
 **orderDirection** | **optional.String**| Direction of ordering *(Optional)* | 
 **externalStationId** | **optional.Int64**| Query on a different (content providing) station *(Optional)* | 

### Return type

[**ItemResults**](ItemResults.md)

### Authorization

[API-Key](../README.md#API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlaylistPostMerge

> InlineResponse202 PlaylistPostMerge(ctx, optional)

Post a playlist, do not remove previously imported items

Post a playlist, do not remove previously imported items

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PlaylistPostMergeOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PlaylistPostMergeOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data** | [**optional.Interface of InlineObject2**](InlineObject2.md)|  | 

### Return type

[**InlineResponse202**](inline_response_202.md)

### Authorization

[API-Key](../README.md#API-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlaylistPostStructure

> InlineResponse202 PlaylistPostStructure(ctx, optional)

Post a playlist, keep current structure

Post a playlist, keep current structure

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PlaylistPostStructureOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PlaylistPostStructureOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data** | [**optional.Interface of InlineObject1**](InlineObject1.md)|  | 

### Return type

[**InlineResponse202**](inline_response_202.md)

### Authorization

[API-Key](../README.md#API-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlaylistPostTiming

> InlineResponse202 PlaylistPostTiming(ctx, optional)

Post a playlist

Post a playlist

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PlaylistPostTimingOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PlaylistPostTimingOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data** | [**optional.Interface of InlineObject**](InlineObject.md)|  | 

### Return type

[**InlineResponse202**](inline_response_202.md)

### Authorization

[API-Key](../README.md#API-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StopCurrentItem

> Success StopCurrentItem(ctx, optional)

Stop an Item

Set a current playing or specific item on played

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***StopCurrentItemOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a StopCurrentItemOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data** | [**optional.Interface of InlineObject3**](InlineObject3.md)|  | 

### Return type

[**Success**](Success.md)

### Authorization

[API-Key](../README.md#API-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateItemById

> Success UpdateItemById(ctx, id, optional)

Update extended item details by ID.

Update item by id.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64**| ID of Item **(Required)** | [default to 0]
 **optional** | ***UpdateItemByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateItemByIdOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**optional.Interface of ItemDataInput**](ItemDataInput.md)| Data *(Optional)* | 

### Return type

[**Success**](Success.md)

### Authorization

[API-Key](../README.md#API-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

