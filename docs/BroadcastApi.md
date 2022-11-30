# \BroadcastApi

All URIs are relative to *https://radiomanager.io/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBroadcast**](BroadcastApi.md#CreateBroadcast) | **Post** /broadcasts | Create broadcast.
[**DeleteBroadcastById**](BroadcastApi.md#DeleteBroadcastById) | **Delete** /broadcasts/{id} | Delete broadcast by id
[**GetBroadcastById**](BroadcastApi.md#GetBroadcastById) | **Get** /broadcasts/{id} | Get broadcast by id
[**GetCurrentBroadcast**](BroadcastApi.md#GetCurrentBroadcast) | **Get** /broadcasts/current | Get current Broadcast
[**GetDailyEPG**](BroadcastApi.md#GetDailyEPG) | **Get** /broadcasts/epg/daily | Get daily EPG
[**GetEPGByDate**](BroadcastApi.md#GetEPGByDate) | **Get** /broadcasts/epg | Get EPG by date
[**GetNextBroadcast**](BroadcastApi.md#GetNextBroadcast) | **Get** /broadcasts/next | Get next Broadcast
[**GetWeeklyEPG**](BroadcastApi.md#GetWeeklyEPG) | **Get** /broadcasts/epg/weekly | Get weekly EPG
[**ListBroadcasts**](BroadcastApi.md#ListBroadcasts) | **Get** /broadcasts | Get all broadcasts.
[**PrintBroadcastById**](BroadcastApi.md#PrintBroadcastById) | **Get** /broadcasts/print/{id} | Print broadcast by id with template
[**UpdateBroadcastByID**](BroadcastApi.md#UpdateBroadcastByID) | **Patch** /broadcasts/{id} | Update broadcast by id



## CreateBroadcast

> InlineResponse2002 CreateBroadcast(ctx).BroadcastDataInput(broadcastDataInput).Execute()

Create broadcast.



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
    broadcastDataInput := *openapiclient.NewBroadcastDataInput() // BroadcastDataInput | Data **(Required)**

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BroadcastApi.CreateBroadcast(context.Background()).BroadcastDataInput(broadcastDataInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BroadcastApi.CreateBroadcast``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBroadcast`: InlineResponse2002
    fmt.Fprintf(os.Stdout, "Response from `BroadcastApi.CreateBroadcast`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBroadcastRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **broadcastDataInput** | [**BroadcastDataInput**](BroadcastDataInput.md) | Data **(Required)** | 

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


## DeleteBroadcastById

> InlineResponse202 DeleteBroadcastById(ctx, id).Execute()

Delete broadcast by id



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
    id := int64(789) // int64 | ID of Broadcast **(Required)**

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BroadcastApi.DeleteBroadcastById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BroadcastApi.DeleteBroadcastById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteBroadcastById`: InlineResponse202
    fmt.Fprintf(os.Stdout, "Response from `BroadcastApi.DeleteBroadcastById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | ID of Broadcast **(Required)** | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBroadcastByIdRequest struct via the builder pattern


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


## GetBroadcastById

> BroadcastResult GetBroadcastById(ctx, id).ExternalStationId(externalStationId).Execute()

Get broadcast by id



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
    id := int64(789) // int64 | ID of Broadcast **(Required)**
    externalStationId := int64(789) // int64 | Query on a different (content providing) station *(Optional)* (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BroadcastApi.GetBroadcastById(context.Background(), id).ExternalStationId(externalStationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BroadcastApi.GetBroadcastById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBroadcastById`: BroadcastResult
    fmt.Fprintf(os.Stdout, "Response from `BroadcastApi.GetBroadcastById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | ID of Broadcast **(Required)** | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBroadcastByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **externalStationId** | **int64** | Query on a different (content providing) station *(Optional)* | 

### Return type

[**BroadcastResult**](BroadcastResult.md)

### Authorization

[API-Key](../README.md#API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCurrentBroadcast

> BroadcastResult GetCurrentBroadcast(ctx).Withunpublished(withunpublished).Execute()

Get current Broadcast



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
    withunpublished := true // bool | Show Unpublished *(Optional)* (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BroadcastApi.GetCurrentBroadcast(context.Background()).Withunpublished(withunpublished).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BroadcastApi.GetCurrentBroadcast``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCurrentBroadcast`: BroadcastResult
    fmt.Fprintf(os.Stdout, "Response from `BroadcastApi.GetCurrentBroadcast`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrentBroadcastRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **withunpublished** | **bool** | Show Unpublished *(Optional)* | 

### Return type

[**BroadcastResult**](BroadcastResult.md)

### Authorization

[API-Key](../README.md#API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDailyEPG

> EPGResults GetDailyEPG(ctx).Date(date).Withunpublished(withunpublished).Execute()

Get daily EPG



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
    date := time.Now() // time.Time | Date *(Optional)* (optional)
    withunpublished := true // bool | Show Unpublished *(Optional)* (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BroadcastApi.GetDailyEPG(context.Background()).Date(date).Withunpublished(withunpublished).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BroadcastApi.GetDailyEPG``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDailyEPG`: EPGResults
    fmt.Fprintf(os.Stdout, "Response from `BroadcastApi.GetDailyEPG`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDailyEPGRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **date** | **time.Time** | Date *(Optional)* | 
 **withunpublished** | **bool** | Show Unpublished *(Optional)* | 

### Return type

[**EPGResults**](EPGResults.md)

### Authorization

[API-Key](../README.md#API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEPGByDate

> EPGResults GetEPGByDate(ctx).Date(date).Withunpublished(withunpublished).Execute()

Get EPG by date



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
    date := time.Now() // time.Time | Date *(Optional)* (optional)
    withunpublished := true // bool | Show Unpublished *(Optional)* (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BroadcastApi.GetEPGByDate(context.Background()).Date(date).Withunpublished(withunpublished).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BroadcastApi.GetEPGByDate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEPGByDate`: EPGResults
    fmt.Fprintf(os.Stdout, "Response from `BroadcastApi.GetEPGByDate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEPGByDateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **date** | **time.Time** | Date *(Optional)* | 
 **withunpublished** | **bool** | Show Unpublished *(Optional)* | 

### Return type

[**EPGResults**](EPGResults.md)

### Authorization

[API-Key](../README.md#API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNextBroadcast

> BroadcastResult GetNextBroadcast(ctx).Withunpublished(withunpublished).Execute()

Get next Broadcast



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
    withunpublished := true // bool | Show Unpublished *(Optional)* (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BroadcastApi.GetNextBroadcast(context.Background()).Withunpublished(withunpublished).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BroadcastApi.GetNextBroadcast``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNextBroadcast`: BroadcastResult
    fmt.Fprintf(os.Stdout, "Response from `BroadcastApi.GetNextBroadcast`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNextBroadcastRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **withunpublished** | **bool** | Show Unpublished *(Optional)* | 

### Return type

[**BroadcastResult**](BroadcastResult.md)

### Authorization

[API-Key](../README.md#API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWeeklyEPG

> EPGResults GetWeeklyEPG(ctx).Date(date).Withunpublished(withunpublished).Execute()

Get weekly EPG



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
    date := time.Now() // string | Date *(Optional)* (optional)
    withunpublished := true // bool | Show Unpublished *(Optional)* (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BroadcastApi.GetWeeklyEPG(context.Background()).Date(date).Withunpublished(withunpublished).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BroadcastApi.GetWeeklyEPG``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWeeklyEPG`: EPGResults
    fmt.Fprintf(os.Stdout, "Response from `BroadcastApi.GetWeeklyEPG`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWeeklyEPGRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **date** | **string** | Date *(Optional)* | 
 **withunpublished** | **bool** | Show Unpublished *(Optional)* | 

### Return type

[**EPGResults**](EPGResults.md)

### Authorization

[API-Key](../README.md#API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBroadcasts

> InlineResponse2001 ListBroadcasts(ctx).ProgramId(programId).BlockId(blockId).ModelTypeId(modelTypeId).TagId(tagId).PresenterId(presenterId).GenreId(genreId).GroupId(groupId).ItemId(itemId).PlannedInEpg(plannedInEpg).StartMin(startMin).StartMax(startMax).Page(page).Limit(limit).OrderBy(orderBy).OrderDirection(orderDirection).ExternalStationId(externalStationId).Execute()

Get all broadcasts.



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
    programId := int64(789) // int64 | Search on Program ID *(Optional)* `(Relation)` (optional)
    blockId := int64(789) // int64 | Search on Block ID *(Optional)* `(Relation)` (optional)
    modelTypeId := int64(789) // int64 | Search on ModelType ID *(Optional)* `(Relation)` (optional)
    tagId := int64(789) // int64 | Search on Tag ID *(Optional)* `(Relation)` (optional)
    presenterId := int64(789) // int64 | Search on Presenter ID *(Optional)* `(Relation)` (optional)
    genreId := int64(789) // int64 | Search on Genre ID *(Optional)* `(Relation)` (optional)
    groupId := int64(789) // int64 | Search on Group ID *(Optional)* `(Relation)` (optional)
    itemId := int64(789) // int64 | Search on Item ID *(Optional)* `(Relation)` (optional)
    plannedInEpg := int64(789) // int64 | Checks if item is in EPG *(Optional)* (optional)
    startMin := time.Now() // time.Time | Minimum start date *(Optional)* (optional)
    startMax := time.Now() // time.Time | Maximum start date *(Optional)* (optional)
    page := int64(789) // int64 | Current page *(Optional)* (optional) (default to 1)
    limit := int64(789) // int64 | Results per page *(Optional)* (optional)
    orderBy := "orderBy_example" // string | Field to order the results *(Optional)* (optional)
    orderDirection := "orderDirection_example" // string | Direction of ordering *(Optional)* (optional)
    externalStationId := int64(789) // int64 | Query on a different (content providing) station *(Optional)* (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BroadcastApi.ListBroadcasts(context.Background()).ProgramId(programId).BlockId(blockId).ModelTypeId(modelTypeId).TagId(tagId).PresenterId(presenterId).GenreId(genreId).GroupId(groupId).ItemId(itemId).PlannedInEpg(plannedInEpg).StartMin(startMin).StartMax(startMax).Page(page).Limit(limit).OrderBy(orderBy).OrderDirection(orderDirection).ExternalStationId(externalStationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BroadcastApi.ListBroadcasts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBroadcasts`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `BroadcastApi.ListBroadcasts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListBroadcastsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **programId** | **int64** | Search on Program ID *(Optional)* &#x60;(Relation)&#x60; | 
 **blockId** | **int64** | Search on Block ID *(Optional)* &#x60;(Relation)&#x60; | 
 **modelTypeId** | **int64** | Search on ModelType ID *(Optional)* &#x60;(Relation)&#x60; | 
 **tagId** | **int64** | Search on Tag ID *(Optional)* &#x60;(Relation)&#x60; | 
 **presenterId** | **int64** | Search on Presenter ID *(Optional)* &#x60;(Relation)&#x60; | 
 **genreId** | **int64** | Search on Genre ID *(Optional)* &#x60;(Relation)&#x60; | 
 **groupId** | **int64** | Search on Group ID *(Optional)* &#x60;(Relation)&#x60; | 
 **itemId** | **int64** | Search on Item ID *(Optional)* &#x60;(Relation)&#x60; | 
 **plannedInEpg** | **int64** | Checks if item is in EPG *(Optional)* | 
 **startMin** | **time.Time** | Minimum start date *(Optional)* | 
 **startMax** | **time.Time** | Maximum start date *(Optional)* | 
 **page** | **int64** | Current page *(Optional)* | [default to 1]
 **limit** | **int64** | Results per page *(Optional)* | 
 **orderBy** | **string** | Field to order the results *(Optional)* | 
 **orderDirection** | **string** | Direction of ordering *(Optional)* | 
 **externalStationId** | **int64** | Query on a different (content providing) station *(Optional)* | 

### Return type

[**InlineResponse2001**](InlineResponse2001.md)

### Authorization

[API-Key](../README.md#API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrintBroadcastById

> InlineResponse2003 PrintBroadcastById(ctx, id).TemplateId(templateId).Execute()

Print broadcast by id with template



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
    id := int64(789) // int64 | ID of Broadcast **(Required)**
    templateId := int64(789) // int64 | The print template to be used *(Optional)* (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BroadcastApi.PrintBroadcastById(context.Background(), id).TemplateId(templateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BroadcastApi.PrintBroadcastById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PrintBroadcastById`: InlineResponse2003
    fmt.Fprintf(os.Stdout, "Response from `BroadcastApi.PrintBroadcastById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | ID of Broadcast **(Required)** | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrintBroadcastByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **templateId** | **int64** | The print template to be used *(Optional)* | 

### Return type

[**InlineResponse2003**](InlineResponse2003.md)

### Authorization

[API-Key](../README.md#API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBroadcastByID

> InlineResponse202 UpdateBroadcastByID(ctx, id).BroadcastDataInput(broadcastDataInput).Execute()

Update broadcast by id



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
    id := int64(789) // int64 | ID of Broadcast **(Required)**
    broadcastDataInput := *openapiclient.NewBroadcastDataInput() // BroadcastDataInput | Data *(Optional)* (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BroadcastApi.UpdateBroadcastByID(context.Background(), id).BroadcastDataInput(broadcastDataInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BroadcastApi.UpdateBroadcastByID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateBroadcastByID`: InlineResponse202
    fmt.Fprintf(os.Stdout, "Response from `BroadcastApi.UpdateBroadcastByID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | ID of Broadcast **(Required)** | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBroadcastByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **broadcastDataInput** | [**BroadcastDataInput**](BroadcastDataInput.md) | Data *(Optional)* | 

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

