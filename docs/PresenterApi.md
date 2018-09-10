# \PresenterApi

All URIs are relative to *https://staging.radiomanager.io/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePresenter**](PresenterApi.md#CreatePresenter) | **Post** /presenters | Create presenter.
[**DeletePresenterById**](PresenterApi.md#DeletePresenterById) | **Delete** /presenters/{id} | Delete presenter by id
[**GetPresenterById**](PresenterApi.md#GetPresenterById) | **Get** /presenters/{id} | Get presenter by id
[**ListPresenters**](PresenterApi.md#ListPresenters) | **Get** /presenters | Get all presenters.
[**UpdatePresenterByID**](PresenterApi.md#UpdatePresenterByID) | **Patch** /presenters/{id} | Update presenter by id


# **CreatePresenter**
> PostSuccess CreatePresenter(ctx, data)
Create presenter.

Create presenter.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **data** | [**PresenterDataInput**](PresenterDataInput.md)| Data **(Required)** | 

### Return type

[**PostSuccess**](PostSuccess.md)

### Authorization

[API Key](../README.md#API Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePresenterById**
> Success DeletePresenterById(ctx, id)
Delete presenter by id

Delete presenter by id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int64**| id of presenter | 

### Return type

[**Success**](Success.md)

### Authorization

[API Key](../README.md#API Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPresenterById**
> PresenterResult GetPresenterById(ctx, id, optional)
Get presenter by id

Get presenter by id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int64**| id of Presenter | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int64**| id of Presenter | 
 **externalStationId** | **int64**| Query on a different (content providing) station *(Optional)* | 

### Return type

[**PresenterResult**](PresenterResult.md)

### Authorization

[API Key](../README.md#API Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPresenters**
> PresenterResults ListPresenters(ctx, optional)
Get all presenters.

List all presenters.

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
 **programId** | **int64**| Search on Program ID *(Optional)* &#x60;(Relation)&#x60; | 
 **broadcastId** | **int64**| Search on Broadcast ID *(Optional)* &#x60;(Relation)&#x60; | 
 **modelTypeId** | **int64**| Search on ModelType ID (Optional) | 
 **limit** | **int64**| Results per page *(Optional)* | 
 **orderBy** | **string**| Field to order the results *(Optional)* | 
 **orderDirection** | **string**| Direction of ordering *(Optional)* | 
 **externalStationId** | **int64**| Query on a different (content providing) station *(Optional)* | 

### Return type

[**PresenterResults**](PresenterResults.md)

### Authorization

[API Key](../README.md#API Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePresenterByID**
> Success UpdatePresenterByID(ctx, id, optional)
Update presenter by id

Update presenter by id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int64**| id of Presenter | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int64**| id of Presenter | 
 **data** | [**PresenterDataInput**](PresenterDataInput.md)| Data *(Optional)* | 

### Return type

[**Success**](Success.md)

### Authorization

[API Key](../README.md#API Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

