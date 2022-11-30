# \BlockApi

All URIs are relative to *https://radiomanager.io/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBlockById**](BlockApi.md#GetBlockById) | **Get** /blocks/{id} | Get block by id
[**GetCurrentBlock**](BlockApi.md#GetCurrentBlock) | **Get** /blocks/current | Get current Block
[**GetNextBlock**](BlockApi.md#GetNextBlock) | **Get** /blocks/next | Get upcoming Block
[**ListBlocks**](BlockApi.md#ListBlocks) | **Get** /blocks | Get a list of all blocks currently in your station.



## GetBlockById

> BlockResult GetBlockById(ctx, id).ExternalStationId(externalStationId).Execute()

Get block by id



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
    id := int64(789) // int64 | ID of Block **(Required)**
    externalStationId := int64(789) // int64 | Query on a different (content providing) station *(Optional)* (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BlockApi.GetBlockById(context.Background(), id).ExternalStationId(externalStationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlockApi.GetBlockById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBlockById`: BlockResult
    fmt.Fprintf(os.Stdout, "Response from `BlockApi.GetBlockById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | ID of Block **(Required)** | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlockByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **externalStationId** | **int64** | Query on a different (content providing) station *(Optional)* | 

### Return type

[**BlockResult**](BlockResult.md)

### Authorization

[API-Key](../README.md#API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCurrentBlock

> BlockResult GetCurrentBlock(ctx).Execute()

Get current Block



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
    resp, r, err := apiClient.BlockApi.GetCurrentBlock(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlockApi.GetCurrentBlock``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCurrentBlock`: BlockResult
    fmt.Fprintf(os.Stdout, "Response from `BlockApi.GetCurrentBlock`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrentBlockRequest struct via the builder pattern


### Return type

[**BlockResult**](BlockResult.md)

### Authorization

[API-Key](../README.md#API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNextBlock

> BlockResult GetNextBlock(ctx).Execute()

Get upcoming Block



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
    resp, r, err := apiClient.BlockApi.GetNextBlock(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlockApi.GetNextBlock``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNextBlock`: BlockResult
    fmt.Fprintf(os.Stdout, "Response from `BlockApi.GetNextBlock`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetNextBlockRequest struct via the builder pattern


### Return type

[**BlockResult**](BlockResult.md)

### Authorization

[API-Key](../README.md#API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBlocks

> InlineResponse200 ListBlocks(ctx).BroadcastId(broadcastId).ItemId(itemId).ProgramId(programId).StartMin(startMin).StartMax(startMax).Page(page).Limit(limit).OrderBy(orderBy).OrderDirection(orderDirection).ExternalStationId(externalStationId).Execute()

Get a list of all blocks currently in your station.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    broadcastId := int64(789) // int64 | Search on Broadcast ID *(Optional)* `(Relation)` (optional)
    itemId := int64(789) // int64 | Search on Item ID *(Optional)* `(Relation)` (optional)
    programId := int64(789) // int64 | Search on Program ID *(Optional)* `(Relation)` (optional)
    startMin := time.Now() // time.Time | Minimum start date *(Optional)* (optional)
    startMax := time.Now() // time.Time | Maximum start date *(Optional)* (optional)
    page := int64(789) // int64 | Current page *(Optional)* (optional) (default to 1)
    limit := int64(789) // int64 | Results per page *(Optional)* (optional)
    orderBy := "orderBy_example" // string | Field to order the results *(Optional)* (optional)
    orderDirection := "orderDirection_example" // string | Direction of ordering *(Optional)* (optional)
    externalStationId := int64(789) // int64 | Query on a different (content providing) station *(Optional)* (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BlockApi.ListBlocks(context.Background()).BroadcastId(broadcastId).ItemId(itemId).ProgramId(programId).StartMin(startMin).StartMax(startMax).Page(page).Limit(limit).OrderBy(orderBy).OrderDirection(orderDirection).ExternalStationId(externalStationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlockApi.ListBlocks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBlocks`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `BlockApi.ListBlocks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListBlocksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **broadcastId** | **int64** | Search on Broadcast ID *(Optional)* &#x60;(Relation)&#x60; | 
 **itemId** | **int64** | Search on Item ID *(Optional)* &#x60;(Relation)&#x60; | 
 **programId** | **int64** | Search on Program ID *(Optional)* &#x60;(Relation)&#x60; | 
 **startMin** | **time.Time** | Minimum start date *(Optional)* | 
 **startMax** | **time.Time** | Maximum start date *(Optional)* | 
 **page** | **int64** | Current page *(Optional)* | [default to 1]
 **limit** | **int64** | Results per page *(Optional)* | 
 **orderBy** | **string** | Field to order the results *(Optional)* | 
 **orderDirection** | **string** | Direction of ordering *(Optional)* | 
 **externalStationId** | **int64** | Query on a different (content providing) station *(Optional)* | 

### Return type

[**InlineResponse200**](InlineResponse200.md)

### Authorization

[API-Key](../README.md#API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

