# \GroupApi

All URIs are relative to *https://radiomanager.io/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetGroupById**](GroupApi.md#GetGroupById) | **Get** /groups/{id} | Get group by id
[**ListGroups**](GroupApi.md#ListGroups) | **Get** /groups | Get all groups.



## GetGroupById

> GroupResult GetGroupById(ctx, id).ExternalStationId(externalStationId).Execute()

Get group by id



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
    id := int64(789) // int64 | id of Group
    externalStationId := int64(789) // int64 | Query on a different (content providing) station *(Optional)* (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupApi.GetGroupById(context.Background(), id).ExternalStationId(externalStationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupApi.GetGroupById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGroupById`: GroupResult
    fmt.Fprintf(os.Stdout, "Response from `GroupApi.GetGroupById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | id of Group | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **externalStationId** | **int64** | Query on a different (content providing) station *(Optional)* | 

### Return type

[**GroupResult**](GroupResult.md)

### Authorization

[API-Key](../README.md#API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGroups

> InlineResponse2007 ListGroups(ctx).BroadcastId(broadcastId).ProgramId(programId).UserId(userId).Page(page).Limit(limit).OrderBy(orderBy).OrderDirection(orderDirection).ExternalStationId(externalStationId).Execute()

Get all groups.



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
    broadcastId := int64(789) // int64 | Search on Broadcast ID *(Optional)* `(Relation)` (optional)
    programId := int64(789) // int64 | Search on Program ID *(Optional)* `(Relation)` (optional)
    userId := int64(789) // int64 | Search on User ID *(Optional)* `(Relation)` (optional)
    page := int64(789) // int64 | Current page *(Optional)* (optional) (default to 1)
    limit := int64(789) // int64 | Results per page *(Optional)* (optional)
    orderBy := "orderBy_example" // string | Field to order the results *(Optional)* (optional)
    orderDirection := "orderDirection_example" // string | Direction of ordering *(Optional)* (optional)
    externalStationId := int64(789) // int64 | Query on a different (content providing) station *(Optional)* (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupApi.ListGroups(context.Background()).BroadcastId(broadcastId).ProgramId(programId).UserId(userId).Page(page).Limit(limit).OrderBy(orderBy).OrderDirection(orderDirection).ExternalStationId(externalStationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupApi.ListGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListGroups`: InlineResponse2007
    fmt.Fprintf(os.Stdout, "Response from `GroupApi.ListGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **broadcastId** | **int64** | Search on Broadcast ID *(Optional)* &#x60;(Relation)&#x60; | 
 **programId** | **int64** | Search on Program ID *(Optional)* &#x60;(Relation)&#x60; | 
 **userId** | **int64** | Search on User ID *(Optional)* &#x60;(Relation)&#x60; | 
 **page** | **int64** | Current page *(Optional)* | [default to 1]
 **limit** | **int64** | Results per page *(Optional)* | 
 **orderBy** | **string** | Field to order the results *(Optional)* | 
 **orderDirection** | **string** | Direction of ordering *(Optional)* | 
 **externalStationId** | **int64** | Query on a different (content providing) station *(Optional)* | 

### Return type

[**InlineResponse2007**](InlineResponse2007.md)

### Authorization

[API-Key](../README.md#API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

