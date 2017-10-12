# \StoryApi

All URIs are relative to *https://radiomanager.pb/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateStory**](StoryApi.md#CreateStory) | **Post** /stories | Create story.
[**DeleteStoryById**](StoryApi.md#DeleteStoryById) | **Delete** /stories/{id} | Delete story by id
[**GetStoryById**](StoryApi.md#GetStoryById) | **Get** /stories/{id} | Get story by id
[**ListStories**](StoryApi.md#ListStories) | **Get** /stories | Get all stories.
[**UpdateStoryByID**](StoryApi.md#UpdateStoryByID) | **Patch** /stories/{id} | Update story by id


# **CreateStory**
> PostSuccess CreateStory($data)

Create story.

Create story.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data** | [**StoryDataInput**](StoryDataInput.md)| Data **(Required)** | 

### Return type

[**PostSuccess**](PostSuccess.md)

### Authorization

[API Key](../README.md#API Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteStoryById**
> Success DeleteStoryById($id)

Delete story by id

Delete story by id


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int64**| ID of Story **(Required)** | 

### Return type

[**Success**](Success.md)

### Authorization

[API Key](../README.md#API Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStoryById**
> StoryResult GetStoryById($id, $externalStationId)

Get story by id

Get story by id


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int64**| ID of Story **(Required)** | 
 **externalStationId** | **int64**| Query on a different (content providing) station *(Optional)* | [optional] 

### Return type

[**StoryResult**](StoryResult.md)

### Authorization

[API Key](../README.md#API Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListStories**
> StoryResults ListStories($page, $itemId, $modelTypeId, $tagId, $limit, $orderBy, $orderDirection, $externalStationId)

Get all stories.

List all stories.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64**| Current page *(Optional)* | [optional] [default to 1]
 **itemId** | **int64**| Search on Item ID *(Optional)* &#x60;(Relation)&#x60; | [optional] 
 **modelTypeId** | **int64**| Search on ModelType ID *(Optional)* &#x60;(Relation)&#x60; | [optional] 
 **tagId** | **int64**| Search on Tag ID *(Optional)* &#x60;(Relation)&#x60; | [optional] 
 **limit** | **int64**| Results per page *(Optional)* | [optional] 
 **orderBy** | **string**| Field to order the results *(Optional)* | [optional] 
 **orderDirection** | **string**| Direction of ordering *(Optional)* | [optional] 
 **externalStationId** | **int64**| Query on a different (content providing) station *(Optional)* | [optional] 

### Return type

[**StoryResults**](StoryResults.md)

### Authorization

[API Key](../README.md#API Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateStoryByID**
> Success UpdateStoryByID($id, $data)

Update story by id

Update story by id


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int64**| ID of Story **(Required)** | 
 **data** | [**StoryDataInput**](StoryDataInput.md)| Data *(Optional)* | [optional] 

### Return type

[**Success**](Success.md)

### Authorization

[API Key](../README.md#API Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

