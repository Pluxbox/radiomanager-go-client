# \StringApi

All URIs are relative to *https://staging.radiomanager.io/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetStringsByName**](StringApi.md#GetStringsByName) | **Get** /strings/{name} | Get Strings (formatted)



## GetStringsByName

> TextString GetStringsByName(ctx, name).FullModel(fullModel).Execute()

Get Strings (formatted)



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of String Template **(Required)**
    fullModel := true // bool | Full model or content only **(Optional)** (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StringApi.GetStringsByName(context.Background(), name).FullModel(fullModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StringApi.GetStringsByName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStringsByName`: TextString
    fmt.Fprintf(os.Stdout, "Response from `StringApi.GetStringsByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of String Template **(Required)** | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStringsByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fullModel** | **bool** | Full model or content only **(Optional)** | 

### Return type

[**TextString**](TextString.md)

### Authorization

[API-Key](../README.md#API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

