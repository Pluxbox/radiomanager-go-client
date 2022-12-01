# \CampaignApi

All URIs are relative to *https://staging.radiomanager.io/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCampaign**](CampaignApi.md#CreateCampaign) | **Post** /campaigns | Create campaign.
[**DeleteCampaignById**](CampaignApi.md#DeleteCampaignById) | **Delete** /campaigns/{id} | Delete campaign by id
[**GetCampaignById**](CampaignApi.md#GetCampaignById) | **Get** /campaigns/{id} | Get campaign by id
[**ListCampaigns**](CampaignApi.md#ListCampaigns) | **Get** /campaigns | Get all campaigns.
[**UpdateCampaignByID**](CampaignApi.md#UpdateCampaignByID) | **Patch** /campaigns/{id} | Update campaign by id



## CreateCampaign

> InlineResponse2002 CreateCampaign(ctx).CampaignDataInput(campaignDataInput).Execute()

Create campaign.



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
    campaignDataInput := *openapiclient.NewCampaignDataInput(int64(1), time.Now(), time.Now()) // CampaignDataInput | Data **(Required)**

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CampaignApi.CreateCampaign(context.Background()).CampaignDataInput(campaignDataInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CampaignApi.CreateCampaign``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCampaign`: InlineResponse2002
    fmt.Fprintf(os.Stdout, "Response from `CampaignApi.CreateCampaign`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCampaignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **campaignDataInput** | [**CampaignDataInput**](CampaignDataInput.md) | Data **(Required)** | 

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


## DeleteCampaignById

> InlineResponse202 DeleteCampaignById(ctx, id).Execute()

Delete campaign by id



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
    id := int64(789) // int64 | ID of Campaign **(Required)**

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CampaignApi.DeleteCampaignById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CampaignApi.DeleteCampaignById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteCampaignById`: InlineResponse202
    fmt.Fprintf(os.Stdout, "Response from `CampaignApi.DeleteCampaignById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | ID of Campaign **(Required)** | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCampaignByIdRequest struct via the builder pattern


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


## GetCampaignById

> CampaignResult GetCampaignById(ctx, id).Execute()

Get campaign by id



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
    id := int64(789) // int64 | ID of Campaign **(Required)**

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CampaignApi.GetCampaignById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CampaignApi.GetCampaignById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCampaignById`: CampaignResult
    fmt.Fprintf(os.Stdout, "Response from `CampaignApi.GetCampaignById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | ID of Campaign **(Required)** | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CampaignResult**](CampaignResult.md)

### Authorization

[API-Key](../README.md#API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCampaigns

> InlineResponse2004 ListCampaigns(ctx).ItemId(itemId).ModelTypeId(modelTypeId).StartMin(startMin).StartMax(startMax).Page(page).Limit(limit).OrderBy(orderBy).OrderDirection(orderDirection).ExternalStationId(externalStationId).Execute()

Get all campaigns.



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
    itemId := int64(789) // int64 | Search on Item ID *(Optional)* `(Relation)` (optional)
    modelTypeId := int64(789) // int64 | Search on ModelType ID *(Optional)* `(Relation)` (optional)
    startMin := time.Now() // time.Time | Minimum start date *(Optional)* (optional)
    startMax := time.Now() // time.Time | Maximum start date *(Optional)* (optional)
    page := int64(789) // int64 | Current page *(Optional)* (optional) (default to 1)
    limit := int64(789) // int64 | Results per page *(Optional)* (optional)
    orderBy := "orderBy_example" // string | Field to order the results *(Optional)* (optional)
    orderDirection := "orderDirection_example" // string | Direction of ordering *(Optional)* (optional)
    externalStationId := int64(789) // int64 | Query on a different (content providing) station *(Optional)* (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CampaignApi.ListCampaigns(context.Background()).ItemId(itemId).ModelTypeId(modelTypeId).StartMin(startMin).StartMax(startMax).Page(page).Limit(limit).OrderBy(orderBy).OrderDirection(orderDirection).ExternalStationId(externalStationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CampaignApi.ListCampaigns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCampaigns`: InlineResponse2004
    fmt.Fprintf(os.Stdout, "Response from `CampaignApi.ListCampaigns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCampaignsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **itemId** | **int64** | Search on Item ID *(Optional)* &#x60;(Relation)&#x60; | 
 **modelTypeId** | **int64** | Search on ModelType ID *(Optional)* &#x60;(Relation)&#x60; | 
 **startMin** | **time.Time** | Minimum start date *(Optional)* | 
 **startMax** | **time.Time** | Maximum start date *(Optional)* | 
 **page** | **int64** | Current page *(Optional)* | [default to 1]
 **limit** | **int64** | Results per page *(Optional)* | 
 **orderBy** | **string** | Field to order the results *(Optional)* | 
 **orderDirection** | **string** | Direction of ordering *(Optional)* | 
 **externalStationId** | **int64** | Query on a different (content providing) station *(Optional)* | 

### Return type

[**InlineResponse2004**](InlineResponse2004.md)

### Authorization

[API-Key](../README.md#API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCampaignByID

> InlineResponse202 UpdateCampaignByID(ctx, id).CampaignDataInput(campaignDataInput).Execute()

Update campaign by id



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
    id := int64(789) // int64 | ID of Campaign **(Required)**
    campaignDataInput := *openapiclient.NewCampaignDataInput(int64(1), time.Now(), time.Now()) // CampaignDataInput | Data **(Optional)**

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CampaignApi.UpdateCampaignByID(context.Background(), id).CampaignDataInput(campaignDataInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CampaignApi.UpdateCampaignByID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCampaignByID`: InlineResponse202
    fmt.Fprintf(os.Stdout, "Response from `CampaignApi.UpdateCampaignByID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | ID of Campaign **(Required)** | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCampaignByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **campaignDataInput** | [**CampaignDataInput**](CampaignDataInput.md) | Data **(Optional)** | 

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

