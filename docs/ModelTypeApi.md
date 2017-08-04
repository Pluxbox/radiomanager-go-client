# \ModelTypeApi

All URIs are relative to *http://radiomanager.pb/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetModelTypeById**](ModelTypeApi.md#GetModelTypeById) | **Get** /model_types/{id} | Get modelType by id
[**ListModelTypes**](ModelTypeApi.md#ListModelTypes) | **Get** /model_types | Get all modelTypes.


# **GetModelTypeById**
> ModelTypeResult GetModelTypeById(ctx, id, optional)
Get modelType by id

Get modelType by id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **id** | **int64**| ID of ModelType **(Required)** | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int64**| ID of ModelType **(Required)** | 
 **externalStationId** | **int64**| Query on a different (content providing) station *(Optional)* | 

### Return type

[**ModelTypeResult**](ModelTypeResult.md)

### Authorization

[API Key](../README.md#API Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListModelTypes**
> ModelTypeResults ListModelTypes(ctx, optional)
Get all modelTypes.

List all modelTypes.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64**| Current page *(Optional)* | 
 **model** | **string**|  | 
 **programId** | **int64**| Search on Program ID *(Optional)* | 
 **broadcastId** | **int64**| Search on Broadcast ID *(Optional)* | 
 **itemId** | **int64**| Search on Item ID *(Optional)* | 
 **campaignId** | **int64**| Search on Campaign ID *(Optional)* | 
 **presenterId** | **int64**| Search on Presenter ID *(Optional)* | 
 **contactId** | **int64**| Search on Contact ID *(Optional)* | 
 **externalStationId** | **int64**| Query on a different (content providing) station *(Optional)* | 

### Return type

[**ModelTypeResults**](ModelTypeResults.md)

### Authorization

[API Key](../README.md#API Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

