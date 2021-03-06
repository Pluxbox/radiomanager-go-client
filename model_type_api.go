/*
 * RadioManager
 *
 * RadioManager
 *
 * API version: 2.0
 * Contact: support@pluxbox.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package radiomanagerclient

import (
	"io/ioutil"
	"net/url"
	"net/http"
	"strings"
	"golang.org/x/net/context"
	"encoding/json"
	"fmt"
)

// Linger please
var (
	_ context.Context
)

type ModelTypeApiService service


/* ModelTypeApiService Get modelType by id
 Get modelType by id
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param id ID of ModelType **(Required)**
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "externalStationId" (int64) Query on a different (content providing) station *(Optional)*
 @return ModelTypeResult*/
func (a *ModelTypeApiService) GetModelTypeById(ctx context.Context, id int64, localVarOptionals map[string]interface{}) (ModelTypeResult,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  ModelTypeResult
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/model_types/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if id < 0 {
		return successPayload, nil, reportError("id must be greater than 0")
	}
	if err := typeCheckParameter(localVarOptionals["externalStationId"], "int64", "externalStationId"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["externalStationId"].(int64); localVarOk {
		localVarQueryParams.Add("_external_station_id", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["api-key"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* ModelTypeApiService Get all modelTypes.
 List all modelTypes.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "page" (int64) Current page *(Optional)*
     @param "programId" (int64) Search on Program ID *(Optional)*
     @param "broadcastId" (int64) Search on Broadcast ID *(Optional)*
     @param "itemId" (int64) Search on Item ID *(Optional)*
     @param "campaignId" (int64) Search on Campaign ID *(Optional)*
     @param "presenterId" (int64) Search on Presenter ID *(Optional)*
     @param "contactId" (int64) Search on Contact ID *(Optional)*
     @param "model" (string) Search Modeltypes for certain Model *(Optional)*
     @param "limit" (int64) Results per page *(Optional)*
     @param "orderBy" (string) Field to order the results *(Optional)*
     @param "orderDirection" (string) Direction of ordering *(Optional)*
     @param "externalStationId" (int64) Query on a different (content providing) station *(Optional)*
 @return ModelTypeResults*/
func (a *ModelTypeApiService) ListModelTypes(ctx context.Context, localVarOptionals map[string]interface{}) (ModelTypeResults,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  ModelTypeResults
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/model_types"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["page"], "int64", "page"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["programId"], "int64", "programId"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["broadcastId"], "int64", "broadcastId"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["itemId"], "int64", "itemId"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["campaignId"], "int64", "campaignId"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["presenterId"], "int64", "presenterId"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["contactId"], "int64", "contactId"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["model"], "string", "model"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["limit"], "int64", "limit"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["orderBy"], "string", "orderBy"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["orderDirection"], "string", "orderDirection"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["externalStationId"], "int64", "externalStationId"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["page"].(int64); localVarOk {
		localVarQueryParams.Add("page", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["programId"].(int64); localVarOk {
		localVarQueryParams.Add("program_id", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["broadcastId"].(int64); localVarOk {
		localVarQueryParams.Add("broadcast_id", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["itemId"].(int64); localVarOk {
		localVarQueryParams.Add("item_id", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["campaignId"].(int64); localVarOk {
		localVarQueryParams.Add("campaign_id", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["presenterId"].(int64); localVarOk {
		localVarQueryParams.Add("presenter_id", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["contactId"].(int64); localVarOk {
		localVarQueryParams.Add("contact_id", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["model"].(string); localVarOk {
		localVarQueryParams.Add("model", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["limit"].(int64); localVarOk {
		localVarQueryParams.Add("limit", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["orderBy"].(string); localVarOk {
		localVarQueryParams.Add("order-by", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["orderDirection"].(string); localVarOk {
		localVarQueryParams.Add("order-direction", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["externalStationId"].(int64); localVarOk {
		localVarQueryParams.Add("_external_station_id", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["api-key"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

