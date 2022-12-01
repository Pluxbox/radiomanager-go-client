# \ItemApi

All URIs are relative to *https://staging.radiomanager.io/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateItem**](ItemApi.md#CreateItem) | **Post** /items | Create an new item.
[**CurrentItemPostStructure**](ItemApi.md#CurrentItemPostStructure) | **Post** /items/current/structure | Post a current playing item, keep structure
[**CurrentItemPostTiming**](ItemApi.md#CurrentItemPostTiming) | **Post** /items/current/timing | Post current playing Item
[**DeleteItemById**](ItemApi.md#DeleteItemById) | **Delete** /items/{id} | Delete item by ID.
[**GetCurrentItem**](ItemApi.md#GetCurrentItem) | **Get** /items/current | Get current Item
[**GetItemById**](ItemApi.md#GetItemById) | **Get** /items/{id} | Get extended item details by ID.
[**ListItems**](ItemApi.md#ListItems) | **Get** /items | Get a list of all the items currently in your station.
[**PlaylistPostMerge**](ItemApi.md#PlaylistPostMerge) | **Post** /items/playlist/merge | Post a playlist, do not remove previously imported items
[**PlaylistPostStructure**](ItemApi.md#PlaylistPostStructure) | **Post** /items/playlist/structure | Post a playlist, keep current structure
[**PlaylistPostTiming**](ItemApi.md#PlaylistPostTiming) | **Post** /items/playlist/timing | Post a playlist
[**StopCurrentItem**](ItemApi.md#StopCurrentItem) | **Post** /items/stopcurrent | Stop an Item
[**UpdateItemById**](ItemApi.md#UpdateItemById) | **Patch** /items/{id} | Update extended item details by ID.



## CreateItem

> InlineResponse2002 CreateItem(ctx).ItemDataInput(itemDataInput).Execute()

Create an new item.



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
    itemDataInput := *openapiclient.NewItemDataInput(int64(14)) // ItemDataInput | Data **(Required)**

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ItemApi.CreateItem(context.Background()).ItemDataInput(itemDataInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ItemApi.CreateItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateItem`: InlineResponse2002
    fmt.Fprintf(os.Stdout, "Response from `ItemApi.CreateItem`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **itemDataInput** | [**ItemDataInput**](ItemDataInput.md) | Data **(Required)** | 

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


## CurrentItemPostStructure

> ItemResult CurrentItemPostStructure(ctx).ImportItem(importItem).Execute()

Post a current playing item, keep structure



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
    importItem := *openapiclient.NewImportItem(int64(14), "0") // ImportItem | Data **(Required)**

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ItemApi.CurrentItemPostStructure(context.Background()).ImportItem(importItem).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ItemApi.CurrentItemPostStructure``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CurrentItemPostStructure`: ItemResult
    fmt.Fprintf(os.Stdout, "Response from `ItemApi.CurrentItemPostStructure`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCurrentItemPostStructureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **importItem** | [**ImportItem**](ImportItem.md) | Data **(Required)** | 

### Return type

[**ItemResult**](ItemResult.md)

### Authorization

[API-Key](../README.md#API-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CurrentItemPostTiming

> ItemResult CurrentItemPostTiming(ctx).ImportItem(importItem).Execute()

Post current playing Item



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
    importItem := *openapiclient.NewImportItem(int64(14), "0") // ImportItem | Data **(Required)**

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ItemApi.CurrentItemPostTiming(context.Background()).ImportItem(importItem).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ItemApi.CurrentItemPostTiming``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CurrentItemPostTiming`: ItemResult
    fmt.Fprintf(os.Stdout, "Response from `ItemApi.CurrentItemPostTiming`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCurrentItemPostTimingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **importItem** | [**ImportItem**](ImportItem.md) | Data **(Required)** | 

### Return type

[**ItemResult**](ItemResult.md)

### Authorization

[API-Key](../README.md#API-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteItemById

> DeleteItemById(ctx, id).Execute()

Delete item by ID.



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
    id := int64(789) // int64 | ID of Item **(Required)**

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ItemApi.DeleteItemById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ItemApi.DeleteItemById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | ID of Item **(Required)** | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteItemByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[API-Key](../README.md#API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCurrentItem

> ItemResult GetCurrentItem(ctx).Lastplayed(lastplayed).Execute()

Get current Item



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
    lastplayed := true // bool | Show last played item if there is no current item*(Optional)* (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ItemApi.GetCurrentItem(context.Background()).Lastplayed(lastplayed).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ItemApi.GetCurrentItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCurrentItem`: ItemResult
    fmt.Fprintf(os.Stdout, "Response from `ItemApi.GetCurrentItem`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrentItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lastplayed** | **bool** | Show last played item if there is no current item*(Optional)* | 

### Return type

[**ItemResult**](ItemResult.md)

### Authorization

[API-Key](../README.md#API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItemById

> ItemResult GetItemById(ctx, id).ExternalStationId(externalStationId).Execute()

Get extended item details by ID.



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
    id := int64(789) // int64 | ID of Item **(Required)**
    externalStationId := int64(789) // int64 | Query on a different (content providing) station *(Optional)* (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ItemApi.GetItemById(context.Background(), id).ExternalStationId(externalStationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ItemApi.GetItemById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetItemById`: ItemResult
    fmt.Fprintf(os.Stdout, "Response from `ItemApi.GetItemById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | ID of Item **(Required)** | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetItemByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **externalStationId** | **int64** | Query on a different (content providing) station *(Optional)* | 

### Return type

[**ItemResult**](ItemResult.md)

### Authorization

[API-Key](../README.md#API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListItems

> InlineResponse2008 ListItems(ctx).BlockId(blockId).BroadcastId(broadcastId).ModelTypeId(modelTypeId).TagId(tagId).CampaignId(campaignId).ContactId(contactId).ProgramDraftId(programDraftId).UserDraftId(userDraftId).StationDraftId(stationDraftId).ProgramId(programId).ExternalId(externalId).DurationMin(durationMin).DurationMax(durationMax).Status(status).StartMin(startMin).StartMax(startMax).Page(page).Limit(limit).OrderBy(orderBy).OrderDirection(orderDirection).ExternalStationId(externalStationId).Execute()

Get a list of all the items currently in your station.



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
    blockId := int64(789) // int64 | Search on Block ID *(Optional)* `(Relation)` (optional)
    broadcastId := int64(789) // int64 | Search on Broadcast ID *(Optional)* `(Relation)` (optional)
    modelTypeId := int64(789) // int64 | Search on ModelType ID *(Optional)* `(Relation)` (optional)
    tagId := int64(789) // int64 | Search on Tag ID *(Optional)* `(Relation)` (optional)
    campaignId := int64(789) // int64 | Search on Campaign ID *(Optional)* `(Relation)` (optional)
    contactId := int64(789) // int64 | Search on Contact ID *(Optional)* `(Relation)` (optional)
    programDraftId := int64(789) // int64 | Search on Program Draft ID *(Optional)* (optional)
    userDraftId := int64(789) // int64 | Search on User Draft ID *(Optional)* (optional)
    stationDraftId := int64(789) // int64 | Search on Station Draft ID *(Optional)* (optional)
    programId := int64(789) // int64 | Search on Program ID *(Optional)* `(Relation)` (optional)
    externalId := "externalId_example" // string | Search on External ID *(Optional)* (optional)
    durationMin := int32(56) // int32 | Minimum duration (seconds) *(Optional)* (optional)
    durationMax := int32(56) // int32 | Maximum duration (seconds) *(Optional)* (optional)
    status := "status_example" // string | Play Status of item *(Optional)* (optional)
    startMin := time.Now() // time.Time | Minimum start date *(Optional)* (optional)
    startMax := time.Now() // time.Time | Maximum start date *(Optional)* (optional)
    page := int64(789) // int64 | Current page *(Optional)* (optional) (default to 1)
    limit := int64(789) // int64 | Results per page *(Optional)* (optional)
    orderBy := "orderBy_example" // string | Field to order the results *(Optional)* (optional)
    orderDirection := "orderDirection_example" // string | Direction of ordering *(Optional)* (optional)
    externalStationId := int64(789) // int64 | Query on a different (content providing) station *(Optional)* (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ItemApi.ListItems(context.Background()).BlockId(blockId).BroadcastId(broadcastId).ModelTypeId(modelTypeId).TagId(tagId).CampaignId(campaignId).ContactId(contactId).ProgramDraftId(programDraftId).UserDraftId(userDraftId).StationDraftId(stationDraftId).ProgramId(programId).ExternalId(externalId).DurationMin(durationMin).DurationMax(durationMax).Status(status).StartMin(startMin).StartMax(startMax).Page(page).Limit(limit).OrderBy(orderBy).OrderDirection(orderDirection).ExternalStationId(externalStationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ItemApi.ListItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListItems`: InlineResponse2008
    fmt.Fprintf(os.Stdout, "Response from `ItemApi.ListItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blockId** | **int64** | Search on Block ID *(Optional)* &#x60;(Relation)&#x60; | 
 **broadcastId** | **int64** | Search on Broadcast ID *(Optional)* &#x60;(Relation)&#x60; | 
 **modelTypeId** | **int64** | Search on ModelType ID *(Optional)* &#x60;(Relation)&#x60; | 
 **tagId** | **int64** | Search on Tag ID *(Optional)* &#x60;(Relation)&#x60; | 
 **campaignId** | **int64** | Search on Campaign ID *(Optional)* &#x60;(Relation)&#x60; | 
 **contactId** | **int64** | Search on Contact ID *(Optional)* &#x60;(Relation)&#x60; | 
 **programDraftId** | **int64** | Search on Program Draft ID *(Optional)* | 
 **userDraftId** | **int64** | Search on User Draft ID *(Optional)* | 
 **stationDraftId** | **int64** | Search on Station Draft ID *(Optional)* | 
 **programId** | **int64** | Search on Program ID *(Optional)* &#x60;(Relation)&#x60; | 
 **externalId** | **string** | Search on External ID *(Optional)* | 
 **durationMin** | **int32** | Minimum duration (seconds) *(Optional)* | 
 **durationMax** | **int32** | Maximum duration (seconds) *(Optional)* | 
 **status** | **string** | Play Status of item *(Optional)* | 
 **startMin** | **time.Time** | Minimum start date *(Optional)* | 
 **startMax** | **time.Time** | Maximum start date *(Optional)* | 
 **page** | **int64** | Current page *(Optional)* | [default to 1]
 **limit** | **int64** | Results per page *(Optional)* | 
 **orderBy** | **string** | Field to order the results *(Optional)* | 
 **orderDirection** | **string** | Direction of ordering *(Optional)* | 
 **externalStationId** | **int64** | Query on a different (content providing) station *(Optional)* | 

### Return type

[**InlineResponse2008**](InlineResponse2008.md)

### Authorization

[API-Key](../README.md#API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlaylistPostMerge

> InlineResponse2021 PlaylistPostMerge(ctx).PlaylistMergeBody(playlistMergeBody).Execute()

Post a playlist, do not remove previously imported items



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
    playlistMergeBody := *openapiclient.NewPlaylistMergeBody() // PlaylistMergeBody | Data *(Required)*

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ItemApi.PlaylistPostMerge(context.Background()).PlaylistMergeBody(playlistMergeBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ItemApi.PlaylistPostMerge``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PlaylistPostMerge`: InlineResponse2021
    fmt.Fprintf(os.Stdout, "Response from `ItemApi.PlaylistPostMerge`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlaylistPostMergeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **playlistMergeBody** | [**PlaylistMergeBody**](PlaylistMergeBody.md) | Data *(Required)* | 

### Return type

[**InlineResponse2021**](InlineResponse2021.md)

### Authorization

[API-Key](../README.md#API-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlaylistPostStructure

> InlineResponse2021 PlaylistPostStructure(ctx).PlaylistStructureBody(playlistStructureBody).Execute()

Post a playlist, keep current structure



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
    playlistStructureBody := *openapiclient.NewPlaylistStructureBody() // PlaylistStructureBody | Data *(Required)*

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ItemApi.PlaylistPostStructure(context.Background()).PlaylistStructureBody(playlistStructureBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ItemApi.PlaylistPostStructure``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PlaylistPostStructure`: InlineResponse2021
    fmt.Fprintf(os.Stdout, "Response from `ItemApi.PlaylistPostStructure`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlaylistPostStructureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **playlistStructureBody** | [**PlaylistStructureBody**](PlaylistStructureBody.md) | Data *(Required)* | 

### Return type

[**InlineResponse2021**](InlineResponse2021.md)

### Authorization

[API-Key](../README.md#API-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlaylistPostTiming

> InlineResponse2021 PlaylistPostTiming(ctx).PlaylistTimingBody(playlistTimingBody).Execute()

Post a playlist



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
    playlistTimingBody := *openapiclient.NewPlaylistTimingBody() // PlaylistTimingBody | Data *(Required)*

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ItemApi.PlaylistPostTiming(context.Background()).PlaylistTimingBody(playlistTimingBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ItemApi.PlaylistPostTiming``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PlaylistPostTiming`: InlineResponse2021
    fmt.Fprintf(os.Stdout, "Response from `ItemApi.PlaylistPostTiming`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlaylistPostTimingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **playlistTimingBody** | [**PlaylistTimingBody**](PlaylistTimingBody.md) | Data *(Required)* | 

### Return type

[**InlineResponse2021**](InlineResponse2021.md)

### Authorization

[API-Key](../README.md#API-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StopCurrentItem

> InlineResponse202 StopCurrentItem(ctx).ItemsStopcurrentBody(itemsStopcurrentBody).Execute()

Stop an Item



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
    itemsStopcurrentBody := *openapiclient.NewItemsStopcurrentBody() // ItemsStopcurrentBody | Data *(Optional)* (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ItemApi.StopCurrentItem(context.Background()).ItemsStopcurrentBody(itemsStopcurrentBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ItemApi.StopCurrentItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StopCurrentItem`: InlineResponse202
    fmt.Fprintf(os.Stdout, "Response from `ItemApi.StopCurrentItem`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStopCurrentItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **itemsStopcurrentBody** | [**ItemsStopcurrentBody**](ItemsStopcurrentBody.md) | Data *(Optional)* | 

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


## UpdateItemById

> InlineResponse202 UpdateItemById(ctx, id).ItemDataInput(itemDataInput).Execute()

Update extended item details by ID.



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
    id := int64(789) // int64 | ID of Item **(Required)**
    itemDataInput := *openapiclient.NewItemDataInput(int64(14)) // ItemDataInput | Data *(Optional)*

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ItemApi.UpdateItemById(context.Background(), id).ItemDataInput(itemDataInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ItemApi.UpdateItemById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateItemById`: InlineResponse202
    fmt.Fprintf(os.Stdout, "Response from `ItemApi.UpdateItemById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | ID of Item **(Required)** | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateItemByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **itemDataInput** | [**ItemDataInput**](ItemDataInput.md) | Data *(Optional)* | 

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

