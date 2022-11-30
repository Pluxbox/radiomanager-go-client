# \UserApi

All URIs are relative to *https://radiomanager.io/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteUserById**](UserApi.md#DeleteUserById) | **Delete** /users/{id} | Remove User from station by Id
[**GetUserById**](UserApi.md#GetUserById) | **Get** /users/{id} | Get user by id
[**InviteUserByMail**](UserApi.md#InviteUserByMail) | **Post** /users/invite | Invite user by mail
[**ListUsers**](UserApi.md#ListUsers) | **Get** /users | Get all users.



## DeleteUserById

> InlineResponse202 DeleteUserById(ctx, id).Execute()

Remove User from station by Id



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
    id := int64(789) // int64 | ID of User **(Required)**

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.DeleteUserById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.DeleteUserById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteUserById`: InlineResponse202
    fmt.Fprintf(os.Stdout, "Response from `UserApi.DeleteUserById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | ID of User **(Required)** | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserByIdRequest struct via the builder pattern


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


## GetUserById

> UserResult GetUserById(ctx, id).Execute()

Get user by id



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
    id := int64(789) // int64 | id of User

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.GetUserById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.GetUserById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserById`: UserResult
    fmt.Fprintf(os.Stdout, "Response from `UserApi.GetUserById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | id of User | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserResult**](UserResult.md)

### Authorization

[API-Key](../README.md#API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InviteUserByMail

> InlineResponse202 InviteUserByMail(ctx).InviteUserData(inviteUserData).Execute()

Invite user by mail



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
    inviteUserData := *openapiclient.NewInviteUserData("info@example.com", []int32{int32(1)}) // InviteUserData | Data *(Required)*

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.InviteUserByMail(context.Background()).InviteUserData(inviteUserData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.InviteUserByMail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InviteUserByMail`: InlineResponse202
    fmt.Fprintf(os.Stdout, "Response from `UserApi.InviteUserByMail`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInviteUserByMailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inviteUserData** | [**InviteUserData**](InviteUserData.md) | Data *(Required)* | 

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


## ListUsers

> InlineResponse20013 ListUsers(ctx).RoleId(roleId).GroupId(groupId).Page(page).Limit(limit).OrderBy(orderBy).OrderDirection(orderDirection).Execute()

Get all users.



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
    roleId := int64(789) // int64 | Search on Role ID *(Optional)* (optional)
    groupId := int64(789) // int64 | Search on Group ID *(Optional)* (optional)
    page := int64(789) // int64 | Current page *(Optional)* (optional) (default to 1)
    limit := int64(789) // int64 | Results per page *(Optional)* (optional)
    orderBy := "orderBy_example" // string | Field to order the results *(Optional)* (optional)
    orderDirection := "orderDirection_example" // string | Direction of ordering *(Optional)* (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.ListUsers(context.Background()).RoleId(roleId).GroupId(groupId).Page(page).Limit(limit).OrderBy(orderBy).OrderDirection(orderDirection).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.ListUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUsers`: InlineResponse20013
    fmt.Fprintf(os.Stdout, "Response from `UserApi.ListUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **roleId** | **int64** | Search on Role ID *(Optional)* | 
 **groupId** | **int64** | Search on Group ID *(Optional)* | 
 **page** | **int64** | Current page *(Optional)* | [default to 1]
 **limit** | **int64** | Results per page *(Optional)* | 
 **orderBy** | **string** | Field to order the results *(Optional)* | 
 **orderDirection** | **string** | Direction of ordering *(Optional)* | 

### Return type

[**InlineResponse20013**](InlineResponse20013.md)

### Authorization

[API-Key](../README.md#API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

