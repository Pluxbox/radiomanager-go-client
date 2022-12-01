# \VisualSlideApi

All URIs are relative to *https://staging.radiomanager.io/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetVisualSlide**](VisualSlideApi.md#GetVisualSlide) | **Get** /visual | Get Visual Slide Image



## GetVisualSlide

> VisualResult GetVisualSlide(ctx).Execute()

Get Visual Slide Image



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VisualSlideApi.GetVisualSlide(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VisualSlideApi.GetVisualSlide``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVisualSlide`: VisualResult
    fmt.Fprintf(os.Stdout, "Response from `VisualSlideApi.GetVisualSlide`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVisualSlideRequest struct via the builder pattern


### Return type

[**VisualResult**](VisualResult.md)

### Authorization

[API-Key](../README.md#API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

