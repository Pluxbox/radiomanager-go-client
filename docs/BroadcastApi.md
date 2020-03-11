# \BroadcastApi

All URIs are relative to *https://radiomanager.io/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBroadcast**](BroadcastApi.md#CreateBroadcast) | **Post** /broadcasts | Create broadcast.
[**DeleteBroadcastById**](BroadcastApi.md#DeleteBroadcastById) | **Delete** /broadcasts/{id} | Delete broadcast by id
[**GetBroadcastById**](BroadcastApi.md#GetBroadcastById) | **Get** /broadcasts/{id} | Get broadcast by id
[**GetCurrentBroadcast**](BroadcastApi.md#GetCurrentBroadcast) | **Get** /broadcasts/current | Get current Broadcast
[**GetDailyEPG**](BroadcastApi.md#GetDailyEPG) | **Get** /broadcasts/epg/daily | Get daily EPG
[**GetEPGByDate**](BroadcastApi.md#GetEPGByDate) | **Get** /broadcasts/epg | Get EPG by date
[**GetNextBroadcast**](BroadcastApi.md#GetNextBroadcast) | **Get** /broadcasts/next | Get next Broadcast
[**GetWeeklyEPG**](BroadcastApi.md#GetWeeklyEPG) | **Get** /broadcasts/epg/weekly | Get weekly EPG
[**ListBroadcasts**](BroadcastApi.md#ListBroadcasts) | **Get** /broadcasts | Get all broadcasts.
[**PrintBroadcastById**](BroadcastApi.md#PrintBroadcastById) | **Get** /broadcasts/print/{id} | Print broadcast by id with template
[**UpdateBroadcastByID**](BroadcastApi.md#UpdateBroadcastByID) | **Patch** /broadcasts/{id} | Update broadcast by id



## CreateBroadcast

> PostSuccess CreateBroadcast(ctx, data)

Create broadcast.

Create broadcast.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**data** | [**BroadcastDataInput**](BroadcastDataInput.md)| Data **(Required)** | 

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


## DeleteBroadcastById

> Success DeleteBroadcastById(ctx, id)

Delete broadcast by id

Delete broadcast by id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64**| ID of Broadcast **(Required)** | [default to 0]

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


## GetBroadcastById

> BroadcastResult GetBroadcastById(ctx, id, optional)

Get broadcast by id

Get broadcast by id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64**| ID of Broadcast **(Required)** | [default to 0]
 **optional** | ***GetBroadcastByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetBroadcastByIdOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **externalStationId** | **optional.Int64**| Query on a different (content providing) station *(Optional)* | 

### Return type

[**BroadcastResult**](BroadcastResult.md)

### Authorization

[API-Key](../README.md#API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCurrentBroadcast

> BroadcastResult GetCurrentBroadcast(ctx, optional)

Get current Broadcast

Get current Broadcast

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetCurrentBroadcastOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetCurrentBroadcastOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **withunpublished** | **optional.Bool**| Show Unpublished *(Optional)* | 

### Return type

[**BroadcastResult**](BroadcastResult.md)

### Authorization

[API-Key](../README.md#API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDailyEPG

> EpgResults GetDailyEPG(ctx, optional)

Get daily EPG

Get current Broadcast

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetDailyEPGOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetDailyEPGOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **date** | **optional.Time**| Date *(Optional)* | 
 **withunpublished** | **optional.Bool**| Show Unpublished *(Optional)* | 

### Return type

[**EpgResults**](EPGResults.md)

### Authorization

[API-Key](../README.md#API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEPGByDate

> EpgResults GetEPGByDate(ctx, optional)

Get EPG by date

Get EPG by date

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetEPGByDateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetEPGByDateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **date** | **optional.Time**| Date *(Optional)* | 
 **withunpublished** | **optional.Bool**| Show Unpublished *(Optional)* | 

### Return type

[**EpgResults**](EPGResults.md)

### Authorization

[API-Key](../README.md#API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNextBroadcast

> BroadcastResult GetNextBroadcast(ctx, optional)

Get next Broadcast

Get next Broadcast

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetNextBroadcastOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetNextBroadcastOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **withunpublished** | **optional.Bool**| Show Unpublished *(Optional)* | 

### Return type

[**BroadcastResult**](BroadcastResult.md)

### Authorization

[API-Key](../README.md#API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWeeklyEPG

> EpgResults GetWeeklyEPG(ctx, optional)

Get weekly EPG

Get weekly EPG

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetWeeklyEPGOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetWeeklyEPGOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **date** | **optional.String**| Date *(Optional)* | 
 **withunpublished** | **optional.Bool**| Show Unpublished *(Optional)* | 

### Return type

[**EpgResults**](EPGResults.md)

### Authorization

[API-Key](../README.md#API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBroadcasts

> BroadcastResults ListBroadcasts(ctx, optional)

Get all broadcasts.

List all broadcasts.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListBroadcastsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListBroadcastsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int64**| Current page *(Optional)* | [default to 1]
 **programId** | **optional.Int64**| Search on Program ID *(Optional)* &#x60;(Relation)&#x60; | 
 **blockId** | **optional.Int64**| Search on Block ID *(Optional)* &#x60;(Relation)&#x60; | 
 **modelTypeId** | **optional.Int64**| Search on ModelType ID *(Optional)* &#x60;(Relation)&#x60; | 
 **tagId** | **optional.Int64**| Search on Tag ID *(Optional)* &#x60;(Relation)&#x60; | 
 **presenterId** | **optional.Int64**| Search on Presenter ID *(Optional)* &#x60;(Relation)&#x60; | 
 **genreId** | **optional.Int64**| Search on Genre ID *(Optional)* &#x60;(Relation)&#x60; | 
 **itemId** | **optional.Int64**| Search on Item ID *(Optional)* &#x60;(Relation)&#x60; | 
 **startMin** | **optional.Time**| Minimum start date *(Optional)* | 
 **startMax** | **optional.Time**| Maximum start date *(Optional)* | 
 **limit** | **optional.Int64**| Results per page *(Optional)* | 
 **orderBy** | **optional.String**| Field to order the results *(Optional)* | 
 **orderDirection** | **optional.String**| Direction of ordering *(Optional)* | 
 **externalStationId** | **optional.Int64**| Query on a different (content providing) station *(Optional)* | 

### Return type

[**BroadcastResults**](BroadcastResults.md)

### Authorization

[API-Key](../README.md#API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrintBroadcastById

> string PrintBroadcastById(ctx, id, optional)

Print broadcast by id with template

Print broadcast by id with template

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64**| ID of Broadcast **(Required)** | [default to 0]
 **optional** | ***PrintBroadcastByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PrintBroadcastByIdOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **templateId** | **optional.Int64**| Search on template ID *(Optional)* | 

### Return type

**string**

### Authorization

[API-Key](../README.md#API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBroadcastByID

> Success UpdateBroadcastByID(ctx, id, optional)

Update broadcast by id

Update broadcast by id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64**| ID of Broadcast **(Required)** | [default to 0]
 **optional** | ***UpdateBroadcastByIDOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateBroadcastByIDOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**optional.Interface of BroadcastDataInput**](BroadcastDataInput.md)| Data *(Optional)* | 

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

