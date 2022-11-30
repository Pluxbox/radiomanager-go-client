# \ModelTypeApi

All URIs are relative to *https://radiomanager.io/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetModelTypeById**](ModelTypeApi.md#GetModelTypeById) | **Get** /model_types/{id} | Get modelType by id
[**ListModelTypes**](ModelTypeApi.md#ListModelTypes) | **Get** /model_types | Get all modelTypes.



## GetModelTypeById

> ModelTypeResult GetModelTypeById(ctx, id).OrderDirection(orderDirection).ExternalStationId(externalStationId).Execute()

Get modelType by id



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
    id := int64(789) // int64 | ID of ModelType **(Required)**
    orderDirection := "orderDirection_example" // string | Direction of ordering *(Optional)* (optional)
    externalStationId := int64(789) // int64 | Query on a different (content providing) station *(Optional)* (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ModelTypeApi.GetModelTypeById(context.Background(), id).OrderDirection(orderDirection).ExternalStationId(externalStationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ModelTypeApi.GetModelTypeById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetModelTypeById`: ModelTypeResult
    fmt.Fprintf(os.Stdout, "Response from `ModelTypeApi.GetModelTypeById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | ID of ModelType **(Required)** | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetModelTypeByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orderDirection** | **string** | Direction of ordering *(Optional)* | 
 **externalStationId** | **int64** | Query on a different (content providing) station *(Optional)* | 

### Return type

[**ModelTypeResult**](ModelTypeResult.md)

### Authorization

[API-Key](../README.md#API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListModelTypes

> InlineResponse2009 ListModelTypes(ctx).ProgramId(programId).BroadcastId(broadcastId).ItemId(itemId).CampaignId(campaignId).PresenterId(presenterId).ContactId(contactId).Model(model).OrderDirection(orderDirection).ExternalStationId(externalStationId).Execute()

Get all modelTypes.



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
    programId := int64(789) // int64 | Search on Program ID *(Optional)* (optional)
    broadcastId := int64(789) // int64 | Search on Broadcast ID *(Optional)* (optional)
    itemId := int64(789) // int64 | Search on Item ID *(Optional)* (optional)
    campaignId := int64(789) // int64 | Search on Campaign ID *(Optional)* (optional)
    presenterId := int64(789) // int64 | Search on Presenter ID *(Optional)* (optional)
    contactId := int64(789) // int64 | Search on Contact ID *(Optional)* (optional)
    model := "model_example" // string | Search Modeltypes for certain Model *(Optional)* (optional)
    orderDirection := "orderDirection_example" // string | Direction of ordering *(Optional)* (optional)
    externalStationId := int64(789) // int64 | Query on a different (content providing) station *(Optional)* (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ModelTypeApi.ListModelTypes(context.Background()).ProgramId(programId).BroadcastId(broadcastId).ItemId(itemId).CampaignId(campaignId).PresenterId(presenterId).ContactId(contactId).Model(model).OrderDirection(orderDirection).ExternalStationId(externalStationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ModelTypeApi.ListModelTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListModelTypes`: InlineResponse2009
    fmt.Fprintf(os.Stdout, "Response from `ModelTypeApi.ListModelTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListModelTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **programId** | **int64** | Search on Program ID *(Optional)* | 
 **broadcastId** | **int64** | Search on Broadcast ID *(Optional)* | 
 **itemId** | **int64** | Search on Item ID *(Optional)* | 
 **campaignId** | **int64** | Search on Campaign ID *(Optional)* | 
 **presenterId** | **int64** | Search on Presenter ID *(Optional)* | 
 **contactId** | **int64** | Search on Contact ID *(Optional)* | 
 **model** | **string** | Search Modeltypes for certain Model *(Optional)* | 
 **orderDirection** | **string** | Direction of ordering *(Optional)* | 
 **externalStationId** | **int64** | Query on a different (content providing) station *(Optional)* | 

### Return type

[**InlineResponse2009**](InlineResponse2009.md)

### Authorization

[API-Key](../README.md#API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

