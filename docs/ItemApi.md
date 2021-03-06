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


# **CreateItem**
> PostSuccess CreateItem(ctx, optional)
Create an new item.

Create item.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data** | [**ItemDataInput**](ItemDataInput.md)| Data *(Optional)* | 

### Return type

[**PostSuccess**](PostSuccess.md)

### Authorization

[API Key](../README.md#API Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CurrentItemPostStructure**
> Success CurrentItemPostStructure(ctx, optional)
Post a current playing item, keep structure

Post a current playing item, keep structure

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data** | [**ImportItem**](ImportItem.md)| Data *(Optional)* | 

### Return type

[**Success**](Success.md)

### Authorization

[API Key](../README.md#API Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CurrentItemPostTiming**
> Success CurrentItemPostTiming(ctx, optional)
Post a current playing item

Post a current playing item

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data** | [**ImportItem**](ImportItem.md)| Data *(Optional)* | 

### Return type

[**Success**](Success.md)

### Authorization

[API Key](../README.md#API Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteItemById**
> Success DeleteItemById(ctx, id)
Delete item by ID.

Delete item by id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int64**| ID of Item **(Required)** | 

### Return type

[**Success**](Success.md)

### Authorization

[API Key](../README.md#API Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCurrentItem**
> ItemResult GetCurrentItem(ctx, optional)
Get current Item

Get current Item

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lastplayed** | **bool**| Show last played item if there is no current item*(Optional)* | 

### Return type

[**ItemResult**](ItemResult.md)

### Authorization

[API Key](../README.md#API Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetItemById**
> ItemResult GetItemById(ctx, id, optional)
Get extended item details by ID.

Read item by id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int64**| ID of Item **(Required)** | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int64**| ID of Item **(Required)** | 
 **externalStationId** | **int64**| Query on a different (content providing) station *(Optional)* | 

### Return type

[**ItemResult**](ItemResult.md)

### Authorization

[API Key](../README.md#API Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListItems**
> ItemResults ListItems(ctx, optional)
Get a list of all the items currently in your station.

Get a list of all the items currently in your station. This feature supports pagination and will give a maximum results of 50 items back.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64**| Current page *(Optional)* | 
 **blockId** | **int64**| Search on Block ID *(Optional)* &#x60;(Relation)&#x60; | 
 **broadcastId** | **int64**| Search on Broadcast ID *(Optional)* &#x60;(Relation)&#x60; | 
 **modelTypeId** | **int64**| Search on ModelType ID *(Optional)* &#x60;(Relation)&#x60; | 
 **tagId** | **int64**| Search on Tag ID *(Optional)* &#x60;(Relation)&#x60; | 
 **campaignId** | **int64**| Search on Campaign ID *(Optional)* &#x60;(Relation)&#x60; | 
 **contactId** | **int64**| Search on Contact ID *(Optional)* &#x60;(Relation)&#x60; | 
 **programDraftId** | **int64**| Search on Program Draft ID *(Optional)* | 
 **userDraftId** | **int64**| Search on User Draft ID *(Optional)* | 
 **stationDraftId** | **int64**| Search on Station Draft ID *(Optional)* | 
 **programId** | **int64**| Search on Program ID *(Optional)* &#x60;(Relation)&#x60; | 
 **externalId** | **string**| Search on External ID *(Optional)* | 
 **startMin** | **time.Time**| Minimum start date *(Optional)* | 
 **startMax** | **time.Time**| Maximum start date *(Optional)* | 
 **durationMin** | **int32**| Minimum duration (seconds) *(Optional)* | 
 **durationMax** | **int32**| Maximum duration (seconds) *(Optional)* | 
 **status** | **string**| Play Status of item *(Optional)* | 
 **limit** | **int64**| Results per page *(Optional)* | 
 **orderBy** | **string**| Field to order the results *(Optional)* | 
 **orderDirection** | **string**| Direction of ordering *(Optional)* | 
 **externalStationId** | **int64**| Query on a different (content providing) station *(Optional)* | 

### Return type

[**ItemResults**](ItemResults.md)

### Authorization

[API Key](../README.md#API Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PlaylistPostMerge**
> InlineResponse202 PlaylistPostMerge(ctx, optional)
Post a playlist, do not remove previously imported items

Post a playlist, do not remove previously imported items

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data** | [**Data2**](Data2.md)| Data *(Optional)* | 

### Return type

[**InlineResponse202**](inline_response_202.md)

### Authorization

[API Key](../README.md#API Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PlaylistPostStructure**
> InlineResponse202 PlaylistPostStructure(ctx, optional)
Post a playlist, keep current structure

Post a playlist, keep current structure

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data** | [**Data1**](Data1.md)| Data *(Optional)* | 

### Return type

[**InlineResponse202**](inline_response_202.md)

### Authorization

[API Key](../README.md#API Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PlaylistPostTiming**
> InlineResponse202 PlaylistPostTiming(ctx, optional)
Post a playlist

Post a playlist

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data** | [**Data**](Data.md)| Data *(Optional)* | 

### Return type

[**InlineResponse202**](inline_response_202.md)

### Authorization

[API Key](../README.md#API Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StopCurrentItem**
> Success StopCurrentItem(ctx, optional)
Stop an Item

Set a current playing or specific item on played

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data** | [**Data3**](Data3.md)| Data *(Optional)* | 

### Return type

[**Success**](Success.md)

### Authorization

[API Key](../README.md#API Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateItemById**
> Success UpdateItemById(ctx, id, optional)
Update extended item details by ID.

Update item by id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int64**| ID of Item **(Required)** | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int64**| ID of Item **(Required)** | 
 **data** | [**ItemDataInput**](ItemDataInput.md)| Data *(Optional)* | 

### Return type

[**Success**](Success.md)

### Authorization

[API Key](../README.md#API Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

