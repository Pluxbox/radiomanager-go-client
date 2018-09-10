# \GenreApi

All URIs are relative to *https://staging.radiomanager.io/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetGenreById**](GenreApi.md#GetGenreById) | **Get** /genres/{id} | Get genre by id
[**ListGenres**](GenreApi.md#ListGenres) | **Get** /genres | List all genres.


# **GetGenreById**
> GenreResult GetGenreById(ctx, id, optional)
Get genre by id

Get genre by id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int64**| ID of Genre **(Required)** | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int64**| ID of Genre **(Required)** | 
 **externalStationId** | **int64**| Query on a different (content providing) station *(Optional)* | 

### Return type

[**GenreResult**](GenreResult.md)

### Authorization

[API Key](../README.md#API Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListGenres**
> GenreResults ListGenres(ctx, optional)
List all genres.

List all genres.

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
 **parentId** | **int64**| Search on Parent ID of Genre *(Optional)* | 
 **programId** | **int64**| Search on Program ID *(Optional)* &#x60;(Relation)&#x60; | 
 **broadcastId** | **int64**| Search on Broadcast ID *(Optional)* &#x60;(Relation)&#x60; | 
 **limit** | **int64**| Results per page *(Optional)* | 
 **orderBy** | **string**| Field to order the results *(Optional)* | 
 **orderDirection** | **string**| Direction of ordering *(Optional)* | 
 **externalStationId** | **int64**| Query on a different (content providing) station *(Optional)* | 

### Return type

[**GenreResults**](GenreResults.md)

### Authorization

[API Key](../README.md#API Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

