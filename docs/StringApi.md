# \StringApi

All URIs are relative to *https://staging.radiomanager.io/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetStringsByName**](StringApi.md#GetStringsByName) | **Get** /strings/{name} | Get Strings (formatted)


# **GetStringsByName**
> TextString GetStringsByName(ctx, name, fullModel)
Get Strings (formatted)

Get Strings (formatted)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| Name of Strings **(Required)** | 
  **fullModel** | **bool**| Full model or content only **(Required)** | [default to true]

### Return type

[**TextString**](TextString.md)

### Authorization

[API Key](../README.md#API Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

