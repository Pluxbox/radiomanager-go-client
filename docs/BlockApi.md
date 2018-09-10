# \BlockApi

All URIs are relative to *https://staging.radiomanager.io/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBlockById**](BlockApi.md#GetBlockById) | **Get** /blocks/{id} | Get block by id
[**GetCurrentBlock**](BlockApi.md#GetCurrentBlock) | **Get** /blocks/current | Get current Block
[**GetNextBlock**](BlockApi.md#GetNextBlock) | **Get** /blocks/next | Get upcoming Block
[**ListBlocks**](BlockApi.md#ListBlocks) | **Get** /blocks | Get a list of all blocks currently in your station.


# **GetBlockById**
> BlockResult GetBlockById(ctx, id, optional)
Get block by id

Get block by id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int64**| ID of Block **(Required)** | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int64**| ID of Block **(Required)** | 
 **externalStationId** | **int64**| Query on a different (content providing) station *(Optional)* | 

### Return type

[**BlockResult**](BlockResult.md)

### Authorization

[API Key](../README.md#API Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCurrentBlock**
> BlockResult GetCurrentBlock(ctx, )
Get current Block

Get current Block

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**BlockResult**](BlockResult.md)

### Authorization

[API Key](../README.md#API Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNextBlock**
> BlockResult GetNextBlock(ctx, )
Get upcoming Block

Get upcoming Block

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**BlockResult**](BlockResult.md)

### Authorization

[API Key](../README.md#API Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListBlocks**
> BlockResults ListBlocks(ctx, optional)
Get a list of all blocks currently in your station.

Get a list of all blocks currently in your station. This feature supports pagination and will give a maximum of 50 blocks back.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64**| Current page *(Optional)* | [default to 1]
 **broadcastId** | **int64**| Search on Broadcast ID *(Optional)* &#x60;(Relation)&#x60; | 
 **itemId** | **int64**| Search on Item ID *(Optional)* &#x60;(Relation)&#x60; | 
 **programId** | **int64**| Search on Program ID *(Optional)* &#x60;(Relation)&#x60; | 
 **startMin** | **time.Time**| Minimum start date *(Optional)* | 
 **startMax** | **time.Time**| Maximum start date *(Optional)* | 
 **limit** | **int64**| Results per page *(Optional)* | 
 **orderBy** | **string**| Field to order the results *(Optional)* | 
 **orderDirection** | **string**| Direction of ordering *(Optional)* | 
 **externalStationId** | **int64**| Query on a different (content providing) station *(Optional)* | 

### Return type

[**BlockResults**](BlockResults.md)

### Authorization

[API Key](../README.md#API Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

