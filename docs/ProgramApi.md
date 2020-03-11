# \ProgramApi

All URIs are relative to *https://radiomanager.io/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateProgram**](ProgramApi.md#CreateProgram) | **Post** /programs | Create program.
[**DeleteProgramById**](ProgramApi.md#DeleteProgramById) | **Delete** /programs/{id} | Delete program by id
[**GetProgramById**](ProgramApi.md#GetProgramById) | **Get** /programs/{id} | Get program by id
[**ListPrograms**](ProgramApi.md#ListPrograms) | **Get** /programs | Get all programs.
[**UpdateProgramByID**](ProgramApi.md#UpdateProgramByID) | **Patch** /programs/{id} | Update program by id



## CreateProgram

> PostSuccess CreateProgram(ctx, data)

Create program.

Create program.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**data** | [**ProgramDataInput**](ProgramDataInput.md)| Data **(Required)** | 

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


## DeleteProgramById

> Success DeleteProgramById(ctx, id)

Delete program by id

Delete program by id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64**| ID of program **(Required)** | [default to 0]

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


## GetProgramById

> ProgramResult GetProgramById(ctx, id, optional)

Get program by id

Get program by id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64**| ID of Program **(Required)** | [default to 0]
 **optional** | ***GetProgramByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetProgramByIdOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **externalStationId** | **optional.Int64**| Query on a different (content providing) station *(Optional)* | 

### Return type

[**ProgramResult**](ProgramResult.md)

### Authorization

[API-Key](../README.md#API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPrograms

> ProgramResults ListPrograms(ctx, optional)

Get all programs.

List all programs.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListProgramsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListProgramsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int64**| Current page *(Optional)* | 
 **broadcastId** | **optional.Int64**| Search on Broadcast ID *(Optional)* &#x60;(Relation)&#x60; | 
 **modelTypeId** | **optional.Int64**| Search on ModelType ID *(Optional)* &#x60;(Relation)&#x60; | 
 **tagId** | **optional.Int64**| Search on Tag ID *(Optional)* &#x60;(Relation)&#x60; | 
 **presenterId** | **optional.Int64**| Search on Presenter ID *(Optional)* &#x60;(Relation)&#x60; | 
 **genreId** | **optional.Int64**| Search on Genre ID *(Optional)* | 
 **blockId** | **optional.Int64**| Search on Block ID *(Optional)* &#x60;(Relation)&#x60; | 
 **itemId** | **optional.Int64**| Search on Item ID *(Optional)* &#x60;(Relation)&#x60; | 
 **disabled** | **optional.Int32**| Search on Disabled status *(Optional)* | 
 **limit** | **optional.Int64**| Results per page *(Optional)* | 
 **orderBy** | **optional.String**| Field to order the results *(Optional)* | 
 **orderDirection** | **optional.String**| Direction of ordering *(Optional)* | 
 **externalStationId** | **optional.Int64**| Query on a different (content providing) station *(Optional)* | 

### Return type

[**ProgramResults**](ProgramResults.md)

### Authorization

[API-Key](../README.md#API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProgramByID

> Success UpdateProgramByID(ctx, id, optional)

Update program by id

Update program by id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64**| ID of Program **(Required)** | [default to 0]
 **optional** | ***UpdateProgramByIDOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateProgramByIDOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**optional.Interface of ProgramDataInput**](ProgramDataInput.md)| Data *(Optional)* | 

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

