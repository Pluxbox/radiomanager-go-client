# \ModelTypeApi

All URIs are relative to *https://radiomanager.io/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetModelTypeById**](ModelTypeApi.md#GetModelTypeById) | **Get** /model_types/{id} | Get modelType by id
[**ListModelTypes**](ModelTypeApi.md#ListModelTypes) | **Get** /model_types | Get all modelTypes.



## GetModelTypeById

> ModelTypeResult GetModelTypeById(ctx, id, optional)

Get modelType by id

Get modelType by id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64**| ID of ModelType **(Required)** | [default to 0]
 **optional** | ***GetModelTypeByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetModelTypeByIdOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **externalStationId** | **optional.Int64**| Query on a different (content providing) station *(Optional)* | 

### Return type

[**ModelTypeResult**](ModelTypeResult.md)

### Authorization

[API-Key](../README.md#API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListModelTypes

> ModelTypeResults ListModelTypes(ctx, optional)

Get all modelTypes.

List all modelTypes.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListModelTypesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListModelTypesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int64**| Current page *(Optional)* | 
 **programId** | **optional.Int64**| Search on Program ID *(Optional)* | 
 **broadcastId** | **optional.Int64**| Search on Broadcast ID *(Optional)* | 
 **itemId** | **optional.Int64**| Search on Item ID *(Optional)* | 
 **campaignId** | **optional.Int64**| Search on Campaign ID *(Optional)* | 
 **presenterId** | **optional.Int64**| Search on Presenter ID *(Optional)* | 
 **contactId** | **optional.Int64**| Search on Contact ID *(Optional)* | 
 **model** | **optional.String**| Search Modeltypes for certain Model *(Optional)* | 
 **limit** | **optional.Int64**| Results per page *(Optional)* | 
 **orderBy** | **optional.String**| Field to order the results *(Optional)* | 
 **orderDirection** | **optional.String**| Direction of ordering *(Optional)* | 
 **externalStationId** | **optional.Int64**| Query on a different (content providing) station *(Optional)* | 

### Return type

[**ModelTypeResults**](ModelTypeResults.md)

### Authorization

[API-Key](../README.md#API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

