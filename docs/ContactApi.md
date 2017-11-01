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
> PostSuccess CreateContact($data)

Create contact.

Create contact.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
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
> Success DeleteContactById($id)

Delete contact by id

Delete contact by id


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
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
> ContactResult GetContactById($id, $externalStationId)

Get contact by id

Get contact by id


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int64**| ID of Contact **(Required)** | 
 **externalStationId** | **int64**| Query on a different (content providing) station *(Optional)* | [optional] 

### Return type

[**ContactResult**](ContactResult.md)

### Authorization

[API Key](../README.md#API Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListContacts**
> ContactResults ListContacts($page, $itemId, $modelTypeId, $tagId, $limit, $orderBy, $orderDirection, $externalStationId)

Get all contacts.

List all contacts.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64**| Current page *(Optional)* | [optional] [default to 1]
 **itemId** | **int64**| Search on Item ID *(Optional)* &#x60;(Relation)&#x60; | [optional] 
 **modelTypeId** | **int64**| Search on ModelType ID *(Optional)* &#x60;(Relation)&#x60; | [optional] 
 **tagId** | **int64**| Search on Tag ID *(Optional)* &#x60;(Relation)&#x60; | [optional] 
 **limit** | **int64**| Results per page *(Optional)* | [optional] 
 **orderBy** | **string**| Field to order the results *(Optional)* | [optional] 
 **orderDirection** | **string**| Direction of ordering *(Optional)* | [optional] 
 **externalStationId** | **int64**| Query on a different (content providing) station *(Optional)* | [optional] 

### Return type

[**ContactResults**](ContactResults.md)

### Authorization

[API Key](../README.md#API Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateContactByID**
> Success UpdateContactByID($id, $data)

Update contact by id

Update contact by id


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int64**| ID of Contact **(Required)** | 
 **data** | [**ContactDataInput**](ContactDataInput.md)| Data *(Optional)* | [optional] 

### Return type

[**Success**](Success.md)

### Authorization

[API Key](../README.md#API Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

