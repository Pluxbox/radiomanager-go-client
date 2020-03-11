# \CampaignApi

All URIs are relative to *https://radiomanager.io/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCampaign**](CampaignApi.md#CreateCampaign) | **Post** /campaigns | Create campaign.
[**DeleteCampaignById**](CampaignApi.md#DeleteCampaignById) | **Delete** /campaigns/{id} | Delete campaign by id
[**GetCampaignById**](CampaignApi.md#GetCampaignById) | **Get** /campaigns/{id} | Get campaign by id
[**ListCampaigns**](CampaignApi.md#ListCampaigns) | **Get** /campaigns | Get all campaigns.
[**UpdateCampaignByID**](CampaignApi.md#UpdateCampaignByID) | **Patch** /campaigns/{id} | Update campaign by id



## CreateCampaign

> PostSuccess CreateCampaign(ctx, data)

Create campaign.

Create campaign.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**data** | [**CampaignDataInput**](CampaignDataInput.md)| Data **(Required)** | 

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


## DeleteCampaignById

> Success DeleteCampaignById(ctx, id)

Delete campaign by id

Delete campaign by id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64**| ID of Campaign **(Required)** | [default to 0]

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


## GetCampaignById

> CampaignResult GetCampaignById(ctx, id, optional)

Get campaign by id

Get campaign by id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64**| ID of Campaign **(Required)** | [default to 0]
 **optional** | ***GetCampaignByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetCampaignByIdOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **externalStationId** | **optional.Int64**| Query on a different (content providing) station *(Optional)* | 

### Return type

[**CampaignResult**](CampaignResult.md)

### Authorization

[API-Key](../README.md#API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCampaigns

> CampaignResults ListCampaigns(ctx, optional)

Get all campaigns.

List all campaigns.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListCampaignsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListCampaignsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int64**| Current page *(Optional)* | 
 **itemId** | **optional.Int64**| Search on Item ID *(Optional)* &#x60;(Relation)&#x60; | 
 **modelTypeId** | **optional.Int64**| Search on ModelType ID *(Optional)* &#x60;(Relation)&#x60; | 
 **startMin** | **optional.Time**| Minimum start date *(Optional)* | 
 **startMax** | **optional.Time**| Maximum start date *(Optional)* | 
 **limit** | **optional.Int64**| Results per page *(Optional)* | 
 **orderBy** | **optional.String**| Field to order the results *(Optional)* | 
 **orderDirection** | **optional.String**| Direction of ordering *(Optional)* | 
 **externalStationId** | **optional.Int64**| Query on a different (content providing) station *(Optional)* | 

### Return type

[**CampaignResults**](CampaignResults.md)

### Authorization

[API-Key](../README.md#API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCampaignByID

> Success UpdateCampaignByID(ctx, id, optional)

Update campaign by id

Update campaign by id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64**| ID of Campaign **(Required)** | [default to 0]
 **optional** | ***UpdateCampaignByIDOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateCampaignByIDOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**optional.Interface of CampaignDataInput**](CampaignDataInput.md)| Data *(Optional)* | 

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

