# \ContactApi

All URIs are relative to *https://radiomanager.io/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateContact**](ContactApi.md#CreateContact) | **Post** /contacts | Create contact.
[**DeleteContactById**](ContactApi.md#DeleteContactById) | **Delete** /contacts/{id} | Delete contact by id
[**GetContactById**](ContactApi.md#GetContactById) | **Get** /contacts/{id} | Get contact by id
[**ListContacts**](ContactApi.md#ListContacts) | **Get** /contacts | Get all contacts.
[**UpdateContactByID**](ContactApi.md#UpdateContactByID) | **Patch** /contacts/{id} | Update contact by id



## CreateContact

> PostSuccess CreateContact(ctx, data)

Create contact.

Create contact.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**data** | [**ContactDataInput**](ContactDataInput.md)| Data **(Required)** | 

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


## DeleteContactById

> Success DeleteContactById(ctx, id)

Delete contact by id

Delete contact by id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64**| ID of Contact **(Required)** | [default to 0]

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


## GetContactById

> ContactResult GetContactById(ctx, id, optional)

Get contact by id

Get contact by id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64**| ID of Contact **(Required)** | [default to 0]
 **optional** | ***GetContactByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetContactByIdOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **externalStationId** | **optional.Int64**| Query on a different (content providing) station *(Optional)* | 

### Return type

[**ContactResult**](ContactResult.md)

### Authorization

[API-Key](../README.md#API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListContacts

> ContactResults ListContacts(ctx, optional)

Get all contacts.

List all contacts.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListContactsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListContactsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int64**| Current page *(Optional)* | [default to 1]
 **itemId** | **optional.Int64**| Search on Item ID *(Optional)* &#x60;(Relation)&#x60; | 
 **modelTypeId** | **optional.Int64**| Search on ModelType ID *(Optional)* &#x60;(Relation)&#x60; | 
 **tagId** | **optional.Int64**| Search on Tag ID *(Optional)* &#x60;(Relation)&#x60; | 
 **limit** | **optional.Int64**| Results per page *(Optional)* | 
 **orderBy** | **optional.String**| Field to order the results *(Optional)* | 
 **orderDirection** | **optional.String**| Direction of ordering *(Optional)* | 
 **externalStationId** | **optional.Int64**| Query on a different (content providing) station *(Optional)* | 

### Return type

[**ContactResults**](ContactResults.md)

### Authorization

[API-Key](../README.md#API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateContactByID

> Success UpdateContactByID(ctx, id, optional)

Update contact by id

Update contact by id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64**| ID of Contact **(Required)** | [default to 0]
 **optional** | ***UpdateContactByIDOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateContactByIDOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**optional.Interface of ContactDataInput**](ContactDataInput.md)| Data *(Optional)* | 

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

