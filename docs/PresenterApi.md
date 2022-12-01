# \PresenterApi

All URIs are relative to *https://staging.radiomanager.io/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePresenter**](PresenterApi.md#CreatePresenter) | **Post** /presenters | Create presenter.
[**DeletePresenterById**](PresenterApi.md#DeletePresenterById) | **Delete** /presenters/{id} | Delete presenter by id
[**GetPresenterById**](PresenterApi.md#GetPresenterById) | **Get** /presenters/{id} | Get presenter by id
[**ListPresenters**](PresenterApi.md#ListPresenters) | **Get** /presenters | Get all presenters.
[**UpdatePresenterByID**](PresenterApi.md#UpdatePresenterByID) | **Patch** /presenters/{id} | Update presenter by id



## CreatePresenter

> InlineResponse2002 CreatePresenter(ctx).PresenterDataInput(presenterDataInput).Execute()

Create presenter.



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
    presenterDataInput := *openapiclient.NewPresenterDataInput(int64(1)) // PresenterDataInput | Data **(Required)**

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PresenterApi.CreatePresenter(context.Background()).PresenterDataInput(presenterDataInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PresenterApi.CreatePresenter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePresenter`: InlineResponse2002
    fmt.Fprintf(os.Stdout, "Response from `PresenterApi.CreatePresenter`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePresenterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **presenterDataInput** | [**PresenterDataInput**](PresenterDataInput.md) | Data **(Required)** | 

### Return type

[**InlineResponse2002**](InlineResponse2002.md)

### Authorization

[API-Key](../README.md#API-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePresenterById

> InlineResponse202 DeletePresenterById(ctx, id).Execute()

Delete presenter by id



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
    id := int64(789) // int64 | id of presenter

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PresenterApi.DeletePresenterById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PresenterApi.DeletePresenterById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletePresenterById`: InlineResponse202
    fmt.Fprintf(os.Stdout, "Response from `PresenterApi.DeletePresenterById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | id of presenter | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePresenterByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse202**](InlineResponse202.md)

### Authorization

[API-Key](../README.md#API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPresenterById

> PresenterResult GetPresenterById(ctx, id).Execute()

Get presenter by id



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
    id := int64(789) // int64 | id of Presenter

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PresenterApi.GetPresenterById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PresenterApi.GetPresenterById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPresenterById`: PresenterResult
    fmt.Fprintf(os.Stdout, "Response from `PresenterApi.GetPresenterById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | id of Presenter | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPresenterByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PresenterResult**](PresenterResult.md)

### Authorization

[API-Key](../README.md#API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPresenters

> InlineResponse20010 ListPresenters(ctx).ProgramId(programId).BroadcastId(broadcastId).ModelTypeId(modelTypeId).Page(page).Limit(limit).OrderBy(orderBy).OrderDirection(orderDirection).ExternalStationId(externalStationId).Execute()

Get all presenters.



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
    programId := int64(789) // int64 | Search on Program ID *(Optional)* `(Relation)` (optional)
    broadcastId := int64(789) // int64 | Search on Broadcast ID *(Optional)* `(Relation)` (optional)
    modelTypeId := int64(789) // int64 | Search on ModelType ID (Optional) (optional)
    page := int64(789) // int64 | Current page *(Optional)* (optional) (default to 1)
    limit := int64(789) // int64 | Results per page *(Optional)* (optional)
    orderBy := "orderBy_example" // string | Field to order the results *(Optional)* (optional)
    orderDirection := "orderDirection_example" // string | Direction of ordering *(Optional)* (optional)
    externalStationId := int64(789) // int64 | Query on a different (content providing) station *(Optional)* (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PresenterApi.ListPresenters(context.Background()).ProgramId(programId).BroadcastId(broadcastId).ModelTypeId(modelTypeId).Page(page).Limit(limit).OrderBy(orderBy).OrderDirection(orderDirection).ExternalStationId(externalStationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PresenterApi.ListPresenters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPresenters`: InlineResponse20010
    fmt.Fprintf(os.Stdout, "Response from `PresenterApi.ListPresenters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPresentersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **programId** | **int64** | Search on Program ID *(Optional)* &#x60;(Relation)&#x60; | 
 **broadcastId** | **int64** | Search on Broadcast ID *(Optional)* &#x60;(Relation)&#x60; | 
 **modelTypeId** | **int64** | Search on ModelType ID (Optional) | 
 **page** | **int64** | Current page *(Optional)* | [default to 1]
 **limit** | **int64** | Results per page *(Optional)* | 
 **orderBy** | **string** | Field to order the results *(Optional)* | 
 **orderDirection** | **string** | Direction of ordering *(Optional)* | 
 **externalStationId** | **int64** | Query on a different (content providing) station *(Optional)* | 

### Return type

[**InlineResponse20010**](InlineResponse20010.md)

### Authorization

[API-Key](../README.md#API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePresenterByID

> InlineResponse202 UpdatePresenterByID(ctx, id).PresenterDataInput(presenterDataInput).Execute()

Update presenter by id



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
    id := int64(789) // int64 | id of Presenter
    presenterDataInput := *openapiclient.NewPresenterDataInput(int64(1)) // PresenterDataInput | Data *(Optional)*

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PresenterApi.UpdatePresenterByID(context.Background(), id).PresenterDataInput(presenterDataInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PresenterApi.UpdatePresenterByID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePresenterByID`: InlineResponse202
    fmt.Fprintf(os.Stdout, "Response from `PresenterApi.UpdatePresenterByID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | id of Presenter | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePresenterByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **presenterDataInput** | [**PresenterDataInput**](PresenterDataInput.md) | Data *(Optional)* | 

### Return type

[**InlineResponse202**](InlineResponse202.md)

### Authorization

[API-Key](../README.md#API-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

