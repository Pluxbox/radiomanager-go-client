# \ContactApi

All URIs are relative to *https://staging.radiomanager.io/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateContact**](ContactApi.md#CreateContact) | **Post** /contacts | Create contact.
[**DeleteContactById**](ContactApi.md#DeleteContactById) | **Delete** /contacts/{id} | Delete contact by id
[**GetContactById**](ContactApi.md#GetContactById) | **Get** /contacts/{id} | Get contact by id
[**ListContacts**](ContactApi.md#ListContacts) | **Get** /contacts | Get all contacts.
[**UpdateContactByID**](ContactApi.md#UpdateContactByID) | **Patch** /contacts/{id} | Update contact by id


# **CreateContact**
> PostSuccess CreateContact(ctx, data)
Create contact.

Create contact.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **data** | [**ContactDataInput**](ContactDataInput.md)| Data **(Required)** | 

### Return type

[**PostSuccess**](PostSuccess.md)

### Authorization

[API Key](../README.md#API Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteContactById**
> Success DeleteContactById(ctx, id)
Delete contact by id

Delete contact by id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int64**| ID of Contact **(Required)** | 

### Return type

[**Success**](Success.md)

### Authorization

[API Key](../README.md#API Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetContactById**
> ContactResult GetContactById(ctx, id, optional)
Get contact by id

Get contact by id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int64**| ID of Contact **(Required)** | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int64**| ID of Contact **(Required)** | 
 **externalStationId** | **int64**| Query on a different (content providing) station *(Optional)* | 

### Return type

[**ContactResult**](ContactResult.md)

### Authorization

[API Key](../README.md#API Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListContacts**
> ContactResults ListContacts(ctx, optional)
Get all contacts.

List all contacts.

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
 **itemId** | **int64**| Search on Item ID *(Optional)* &#x60;(Relation)&#x60; | 
 **modelTypeId** | **int64**| Search on ModelType ID *(Optional)* &#x60;(Relation)&#x60; | 
 **tagId** | **int64**| Search on Tag ID *(Optional)* &#x60;(Relation)&#x60; | 
 **limit** | **int64**| Results per page *(Optional)* | 
 **orderBy** | **string**| Field to order the results *(Optional)* | 
 **orderDirection** | **string**| Direction of ordering *(Optional)* | 
 **externalStationId** | **int64**| Query on a different (content providing) station *(Optional)* | 

### Return type

[**ContactResults**](ContactResults.md)

### Authorization

[API Key](../README.md#API Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateContactByID**
> Success UpdateContactByID(ctx, id, optional)
Update contact by id

Update contact by id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int64**| ID of Contact **(Required)** | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int64**| ID of Contact **(Required)** | 
 **data** | [**ContactDataInput**](ContactDataInput.md)| Data *(Optional)* | 

### Return type

[**Success**](Success.md)

### Authorization

[API Key](../README.md#API Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

