# \UserApi

All URIs are relative to *https://radiomanager.io/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteUserById**](UserApi.md#DeleteUserById) | **Delete** /users/{id} | Remove user from station by Id
[**GetUserById**](UserApi.md#GetUserById) | **Get** /users/{id} | Get user by id
[**InviteUserByMail**](UserApi.md#InviteUserByMail) | **Post** /users/invite | Invite user by mail
[**ListUsers**](UserApi.md#ListUsers) | **Get** /users | Get all users.



## DeleteUserById

> Success DeleteUserById(ctx, id)

Remove user from station by Id

Remove user from station by Id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64**| id of User | [default to 0]

### Return type

[**Success**](Success.md)

### Authorization

[API-Key](../README.md#API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserById

> UserResult GetUserById(ctx, id)

Get user by id

Get user by id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64**| id of User | [default to 0]

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

> map[string]interface{} InviteUserByMail(ctx, data)

Invite user by mail

Invite user by mail

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**data** | [**InviteUserData**](InviteUserData.md)| Data **(Required)** | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[API-Key](../README.md#API-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUsers

> UserResults ListUsers(ctx, optional)

Get all users.

List all users.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListUsersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int64**| Current page *(Optional)* | [default to 1]
 **roleId** | **optional.Int64**| Search on Role ID *(Optional)* | 
 **limit** | **optional.Int64**| Results per page *(Optional)* | 
 **orderBy** | **optional.String**| Field to order the results *(Optional)* | 
 **orderDirection** | **optional.String**| Direction of ordering *(Optional)* | 

### Return type

[**UserResults**](UserResults.md)

### Authorization

[API-Key](../README.md#API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

