/* 
 * RadioManager
 *
 * RadioManager
 *
 * OpenAPI spec version: 2.0
 * Contact: support@pluxbox.com
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package radiomanagerclient

import (
	"net/url"
	"strings"
	"encoding/json"
	"fmt"
)

type GenreApi struct {
	Configuration *Configuration
}

func NewGenreApi() *GenreApi {
	configuration := NewConfiguration()
	return &GenreApi{
		Configuration: configuration,
	}
}

func NewGenreApiWithBasePath(basePath string) *GenreApi {
	configuration := NewConfiguration()
	configuration.BasePath = basePath

	return &GenreApi{
		Configuration: configuration,
	}
}

/**
 * Get genre by id
 * Get genre by id
 *
 * @param id ID of Genre **(Required)**
 * @param externalStationId Query on a different (content providing) station *(Optional)*
 * @return *GenreResult
 */
func (a GenreApi) GetGenreById(id int64, externalStationId int64) (*GenreResult, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/genres/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(API Key)' required
	// set key with prefix in header
	localVarHeaderParams["api-key"] = a.Configuration.GetAPIKeyWithPrefix("api-key")
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	localVarQueryParams.Add("_external_station_id", a.Configuration.APIClient.ParameterToString(externalStationId, ""))

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(GenreResult)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GetGenreById", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * List all genres.
 * List all genres.
 *
 * @param page Current page *(Optional)*
 * @param parentId Search on Parent ID of Genre *(Optional)*
 * @param programId Search on Program ID *(Optional)* &#x60;(Relation)&#x60;
 * @param broadcastId Search on Broadcast ID *(Optional)* &#x60;(Relation)&#x60;
 * @param limit Results per page *(Optional)*
 * @param orderBy Field to order the results *(Optional)*
 * @param orderDirection Direction of ordering *(Optional)*
 * @param externalStationId Query on a different (content providing) station *(Optional)*
 * @return *GenreResults
 */
func (a GenreApi) ListGenres(page int64, parentId int64, programId int64, broadcastId int64, limit int64, orderBy string, orderDirection string, externalStationId int64) (*GenreResults, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/genres"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(API Key)' required
	// set key with prefix in header
	localVarHeaderParams["api-key"] = a.Configuration.GetAPIKeyWithPrefix("api-key")
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	localVarQueryParams.Add("page", a.Configuration.APIClient.ParameterToString(page, ""))
	localVarQueryParams.Add("parent_id", a.Configuration.APIClient.ParameterToString(parentId, ""))
	localVarQueryParams.Add("program_id", a.Configuration.APIClient.ParameterToString(programId, ""))
	localVarQueryParams.Add("broadcast_id", a.Configuration.APIClient.ParameterToString(broadcastId, ""))
	localVarQueryParams.Add("limit", a.Configuration.APIClient.ParameterToString(limit, ""))
	localVarQueryParams.Add("order-by", a.Configuration.APIClient.ParameterToString(orderBy, ""))
	localVarQueryParams.Add("order-direction", a.Configuration.APIClient.ParameterToString(orderDirection, ""))
	localVarQueryParams.Add("_external_station_id", a.Configuration.APIClient.ParameterToString(externalStationId, ""))

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(GenreResults)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "ListGenres", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

