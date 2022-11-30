# \TagApi

All URIs are relative to *https://radiomanager.io/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTag**](TagApi.md#CreateTag) | **Post** /tags | Create tag.
[**DeleteTagById**](TagApi.md#DeleteTagById) | **Delete** /tags/{id} | Delete tag by id
[**GetTagById**](TagApi.md#GetTagById) | **Get** /tags/{id} | Get tags by id
[**ListTags**](TagApi.md#ListTags) | **Get** /tags | Get a list of all the tags currently in your station.
[**UpdateTagByID**](TagApi.md#UpdateTagByID) | **Patch** /tags/{id} | Update tag by id



## CreateTag

> InlineResponse2002 CreateTag(ctx).TagDataInput(tagDataInput).Execute()

Create tag.



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
    tagDataInput := *openapiclient.NewTagDataInput("FooBar") // TagDataInput | Data **(Required)**

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TagApi.CreateTag(context.Background()).TagDataInput(tagDataInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagApi.CreateTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTag`: InlineResponse2002
    fmt.Fprintf(os.Stdout, "Response from `TagApi.CreateTag`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tagDataInput** | [**TagDataInput**](TagDataInput.md) | Data **(Required)** | 

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


## DeleteTagById

> InlineResponse202 DeleteTagById(ctx, id).Execute()

Delete tag by id



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
    id := int64(789) // int64 | ID of Tag **(Required)**

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TagApi.DeleteTagById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagApi.DeleteTagById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteTagById`: InlineResponse202
    fmt.Fprintf(os.Stdout, "Response from `TagApi.DeleteTagById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | ID of Tag **(Required)** | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTagByIdRequest struct via the builder pattern


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


## GetTagById

> TagResult GetTagById(ctx, id).Execute()

Get tags by id



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
    id := int64(789) // int64 | ID of Tag **(Required)**

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TagApi.GetTagById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagApi.GetTagById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTagById`: TagResult
    fmt.Fprintf(os.Stdout, "Response from `TagApi.GetTagById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | ID of Tag **(Required)** | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTagByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TagResult**](TagResult.md)

### Authorization

[API-Key](../README.md#API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTags

> InlineResponse20012 ListTags(ctx).ProgramId(programId).ItemId(itemId).BroadcastId(broadcastId).ContactId(contactId).Page(page).Limit(limit).OrderBy(orderBy).OrderDirection(orderDirection).ExternalStationId(externalStationId).Execute()

Get a list of all the tags currently in your station.



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
    itemId := int64(789) // int64 | Search on Item ID *(Optional)* `(Relation)` (optional)
    broadcastId := int64(789) // int64 | Search on Broadcast ID *(Optional)* `(Relation)` (optional)
    contactId := int64(789) // int64 | Search on Contact ID *(Optional)* `(Relation)` (optional)
    page := int64(789) // int64 | Current page *(Optional)* (optional) (default to 1)
    limit := int64(789) // int64 | Results per page *(Optional)* (optional)
    orderBy := "orderBy_example" // string | Field to order the results *(Optional)* (optional)
    orderDirection := "orderDirection_example" // string | Direction of ordering *(Optional)* (optional)
    externalStationId := int64(789) // int64 | Query on a different (content providing) station *(Optional)* (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TagApi.ListTags(context.Background()).ProgramId(programId).ItemId(itemId).BroadcastId(broadcastId).ContactId(contactId).Page(page).Limit(limit).OrderBy(orderBy).OrderDirection(orderDirection).ExternalStationId(externalStationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagApi.ListTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTags`: InlineResponse20012
    fmt.Fprintf(os.Stdout, "Response from `TagApi.ListTags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **programId** | **int64** | Search on Program ID *(Optional)* &#x60;(Relation)&#x60; | 
 **itemId** | **int64** | Search on Item ID *(Optional)* &#x60;(Relation)&#x60; | 
 **broadcastId** | **int64** | Search on Broadcast ID *(Optional)* &#x60;(Relation)&#x60; | 
 **contactId** | **int64** | Search on Contact ID *(Optional)* &#x60;(Relation)&#x60; | 
 **page** | **int64** | Current page *(Optional)* | [default to 1]
 **limit** | **int64** | Results per page *(Optional)* | 
 **orderBy** | **string** | Field to order the results *(Optional)* | 
 **orderDirection** | **string** | Direction of ordering *(Optional)* | 
 **externalStationId** | **int64** | Query on a different (content providing) station *(Optional)* | 

### Return type

[**InlineResponse20012**](InlineResponse20012.md)

### Authorization

[API-Key](../README.md#API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTagByID

> InlineResponse202 UpdateTagByID(ctx, id).TagDataInput(tagDataInput).Execute()

Update tag by id



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
    id := int64(789) // int64 | ID of Tag **(Required)**
    tagDataInput := *openapiclient.NewTagDataInput("FooBar") // TagDataInput | Data *(Optional)*

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TagApi.UpdateTagByID(context.Background(), id).TagDataInput(tagDataInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagApi.UpdateTagByID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTagByID`: InlineResponse202
    fmt.Fprintf(os.Stdout, "Response from `TagApi.UpdateTagByID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | ID of Tag **(Required)** | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTagByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tagDataInput** | [**TagDataInput**](TagDataInput.md) | Data *(Optional)* | 

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

