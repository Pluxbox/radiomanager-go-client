# \ExternalMessageApi

All URIs are relative to *https://staging.radiomanager.pluxbox.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**QueueExternalMessage**](ExternalMessageApi.md#QueueExternalMessage) | **Post** /externalmessagequeue | Queue External Message.


# **QueueExternalMessage**
> interface{} QueueExternalMessage(ctx, message)
Queue External Message.

Queue External Message.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **message** | [**ExternalMessageQueueData**](ExternalMessageQueueData.md)| Data **(Required)** | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[API Key](../README.md#API Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

