# \ProgramApi

All URIs are relative to *https://radiomanager.io/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateProgram**](ProgramApi.md#CreateProgram) | **Post** /programs | Create program.
[**DeleteProgramById**](ProgramApi.md#DeleteProgramById) | **Delete** /programs/{id} | Delete program by id
[**GetProgramById**](ProgramApi.md#GetProgramById) | **Get** /programs/{id} | Get program by id
[**ListPrograms**](ProgramApi.md#ListPrograms) | **Get** /programs | Get all programs.
[**UpdateProgramByID**](ProgramApi.md#UpdateProgramByID) | **Patch** /programs/{id} | Update program by id



## CreateProgram

> InlineResponse2002 CreateProgram(ctx).ProgramDataInput(programDataInput).Execute()

Create program.



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
    programDataInput := *openapiclient.NewProgramDataInput(int64(1), "FooBar") // ProgramDataInput | Data **(Required)**

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProgramApi.CreateProgram(context.Background()).ProgramDataInput(programDataInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProgramApi.CreateProgram``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateProgram`: InlineResponse2002
    fmt.Fprintf(os.Stdout, "Response from `ProgramApi.CreateProgram`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateProgramRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **programDataInput** | [**ProgramDataInput**](ProgramDataInput.md) | Data **(Required)** | 

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


## DeleteProgramById

> InlineResponse202 DeleteProgramById(ctx, id).Execute()

Delete program by id



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
    id := int64(789) // int64 | ID of program **(Required)**

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProgramApi.DeleteProgramById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProgramApi.DeleteProgramById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteProgramById`: InlineResponse202
    fmt.Fprintf(os.Stdout, "Response from `ProgramApi.DeleteProgramById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | ID of program **(Required)** | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProgramByIdRequest struct via the builder pattern


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


## GetProgramById

> ProgramResult GetProgramById(ctx, id).ExternalStationId(externalStationId).Execute()

Get program by id



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
    id := int64(789) // int64 | ID of Program **(Required)**
    externalStationId := int64(789) // int64 | Query on a different (content providing) station *(Optional)* (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProgramApi.GetProgramById(context.Background(), id).ExternalStationId(externalStationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProgramApi.GetProgramById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProgramById`: ProgramResult
    fmt.Fprintf(os.Stdout, "Response from `ProgramApi.GetProgramById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | ID of Program **(Required)** | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProgramByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **externalStationId** | **int64** | Query on a different (content providing) station *(Optional)* | 

### Return type

[**ProgramResult**](ProgramResult.md)

### Authorization

[API-Key](../README.md#API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPrograms

> InlineResponse20011 ListPrograms(ctx).BroadcastId(broadcastId).ModelTypeId(modelTypeId).TagId(tagId).PresenterId(presenterId).GenreId(genreId).GroupId(groupId).BlockId(blockId).ItemId(itemId).Disabled(disabled).Page(page).Limit(limit).OrderBy(orderBy).OrderDirection(orderDirection).ExternalStationId(externalStationId).Execute()

Get all programs.



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
    modelTypeId := int64(789) // int64 | Search on ModelType ID *(Optional)* `(Relation)` (optional)
    tagId := int64(789) // int64 | Search on Tag ID *(Optional)* `(Relation)` (optional)
    presenterId := int64(789) // int64 | Search on Presenter ID *(Optional)* `(Relation)` (optional)
    genreId := int64(789) // int64 | Search on Genre ID *(Optional)* (optional)
    groupId := int64(789) // int64 | Search on Group ID *(Optional)* (optional)
    blockId := int64(789) // int64 | Search on Block ID *(Optional)* `(Relation)` (optional)
    itemId := int64(789) // int64 | Search on Item ID *(Optional)* `(Relation)` (optional)
    disabled := int32(56) // int32 | Search on Disabled status *(Optional)* (optional)
    page := int64(789) // int64 | Current page *(Optional)* (optional) (default to 1)
    limit := int64(789) // int64 | Results per page *(Optional)* (optional)
    orderBy := "orderBy_example" // string | Field to order the results *(Optional)* (optional)
    orderDirection := "orderDirection_example" // string | Direction of ordering *(Optional)* (optional)
    externalStationId := int64(789) // int64 | Query on a different (content providing) station *(Optional)* (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProgramApi.ListPrograms(context.Background()).BroadcastId(broadcastId).ModelTypeId(modelTypeId).TagId(tagId).PresenterId(presenterId).GenreId(genreId).GroupId(groupId).BlockId(blockId).ItemId(itemId).Disabled(disabled).Page(page).Limit(limit).OrderBy(orderBy).OrderDirection(orderDirection).ExternalStationId(externalStationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProgramApi.ListPrograms``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPrograms`: InlineResponse20011
    fmt.Fprintf(os.Stdout, "Response from `ProgramApi.ListPrograms`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListProgramsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **broadcastId** | **int64** | Search on Broadcast ID *(Optional)* &#x60;(Relation)&#x60; | 
 **modelTypeId** | **int64** | Search on ModelType ID *(Optional)* &#x60;(Relation)&#x60; | 
 **tagId** | **int64** | Search on Tag ID *(Optional)* &#x60;(Relation)&#x60; | 
 **presenterId** | **int64** | Search on Presenter ID *(Optional)* &#x60;(Relation)&#x60; | 
 **genreId** | **int64** | Search on Genre ID *(Optional)* | 
 **groupId** | **int64** | Search on Group ID *(Optional)* | 
 **blockId** | **int64** | Search on Block ID *(Optional)* &#x60;(Relation)&#x60; | 
 **itemId** | **int64** | Search on Item ID *(Optional)* &#x60;(Relation)&#x60; | 
 **disabled** | **int32** | Search on Disabled status *(Optional)* | 
 **page** | **int64** | Current page *(Optional)* | [default to 1]
 **limit** | **int64** | Results per page *(Optional)* | 
 **orderBy** | **string** | Field to order the results *(Optional)* | 
 **orderDirection** | **string** | Direction of ordering *(Optional)* | 
 **externalStationId** | **int64** | Query on a different (content providing) station *(Optional)* | 

### Return type

[**InlineResponse20011**](InlineResponse20011.md)

### Authorization

[API-Key](../README.md#API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProgramByID

> InlineResponse202 UpdateProgramByID(ctx, id).ProgramDataInput(programDataInput).Execute()

Update program by id



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
    id := int64(789) // int64 | ID of Program **(Required)**
    programDataInput := *openapiclient.NewProgramDataInput(int64(1), "FooBar") // ProgramDataInput | Data *(Optional)*

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProgramApi.UpdateProgramByID(context.Background(), id).ProgramDataInput(programDataInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProgramApi.UpdateProgramByID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateProgramByID`: InlineResponse202
    fmt.Fprintf(os.Stdout, "Response from `ProgramApi.UpdateProgramByID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | ID of Program **(Required)** | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProgramByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **programDataInput** | [**ProgramDataInput**](ProgramDataInput.md) | Data *(Optional)* | 

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

