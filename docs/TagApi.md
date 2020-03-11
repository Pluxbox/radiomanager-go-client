# \TagApi

All URIs are relative to *https://radiomanager.io/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTag**](TagApi.md#CreateTag) | **Post** /tags | Create tag.
[**DeleteTagById**](TagApi.md#DeleteTagById) | **Delete** /tags/{id} | Delete tag by id
[**GetTagById**](TagApi.md#GetTagById) | **Get** /tags/{id} | Get tags by id
[**ListTags**](TagApi.md#ListTags) | **Get** /tags | Get a list of all the tags currently in your station.
[**UpdateTagByID**](TagApi.md#UpdateTagByID) | **Patch** /tags/{id} | Update tag by id



## CreateTag

> PostSuccess CreateTag(ctx, data)

Create tag.

Create tag.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**data** | [**TagDataInput**](TagDataInput.md)| Data **(Required)** | 

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


## DeleteTagById

> Success DeleteTagById(ctx, id)

Delete tag by id

Delete tag by id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64**| ID of Tag **(Required)** | [default to 0]

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


## GetTagById

> TagResult GetTagById(ctx, id, optional)

Get tags by id

Get tags by id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64**| ID of Tag **(Required)** | [default to 0]
 **optional** | ***GetTagByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetTagByIdOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **externalStationId** | **optional.Int64**| Query on a different (content providing) station *(Optional)* | 

### Return type

[**TagResult**](TagResult.md)

### Authorization

[API-Key](../README.md#API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTags

> TagResults ListTags(ctx, optional)

Get a list of all the tags currently in your station.

Get a list of all the tags currently in your station. This feature supports pagination and will give a maximum results of 50 tags back.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListTagsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListTagsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int64**| Current page *(Optional)* | 
 **programId** | **optional.Int64**| Search on Program ID *(Optional)* &#x60;(Relation)&#x60; | 
 **itemId** | **optional.Int64**| Search on Item ID *(Optional)* &#x60;(Relation)&#x60; | 
 **broadcastId** | **optional.Int64**| Search on Broadcast ID *(Optional)* &#x60;(Relation)&#x60; | 
 **contactId** | **optional.Int64**| Search on Contact ID *(Optional)* &#x60;(Relation)&#x60; | 
 **limit** | **optional.Int64**| Results per page *(Optional)* | 
 **orderBy** | **optional.String**| Field to order the results *(Optional)* | 
 **orderDirection** | **optional.String**| Direction of ordering *(Optional)* | 
 **externalStationId** | **optional.Int64**| Query on a different (content providing) station *(Optional)* | 

### Return type

[**TagResults**](TagResults.md)

### Authorization

[API-Key](../README.md#API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTagByID

> Success UpdateTagByID(ctx, id, optional)

Update tag by id

Update tag by id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64**| ID of Tag **(Required)** | [default to 0]
 **optional** | ***UpdateTagByIDOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateTagByIDOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**optional.Interface of TagDataInput**](TagDataInput.md)| Data *(Optional)* | 

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

