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
	"net/http"
	"strings"
	"golang.org/x/net/context"
	"time"
	"encoding/json"
	"fmt"
)

// Linger please
var (
	_ context.Context
)

type BroadcastApiService service


/* BroadcastApiService Create broadcast.
 Create broadcast.
 * @param ctx context.Context Authentication Context 
 @param data Data **(Required)**
 @return PostSuccess*/
func (a *BroadcastApiService) CreateBroadcast(ctx context.Context, data BroadcastDataInput) (PostSuccess,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  PostSuccess
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/broadcasts"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


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
	// body params
	localVarPostBody = &data
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
		return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
	 }
	
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
	 	return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* BroadcastApiService Delete broadcast by id
 Delete broadcast by id
 * @param ctx context.Context Authentication Context 
 @param id ID of Broadcast **(Required)**
 @return Success*/
func (a *BroadcastApiService) DeleteBroadcastById(ctx context.Context, id int64) (Success,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  Success
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/broadcasts/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if id < 0 {
			return successPayload, nil, reportError("id must be greater than 0")
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
		return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
	 }
	
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
	 	return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* BroadcastApiService Get broadcast by id
 Get broadcast by id
 * @param ctx context.Context Authentication Context 
 @param id ID of Broadcast **(Required)**
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "externalStationId" (int64) Query on a different (content providing) station *(Optional)*
 @return BroadcastResult*/
func (a *BroadcastApiService) GetBroadcastById(ctx context.Context, id int64, localVarOptionals map[string]interface{}) (BroadcastResult,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  BroadcastResult
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/broadcasts/{id}"
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
		return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
	 }
	
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
	 	return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* BroadcastApiService Get current Broadcast
 Get current Broadcast
 * @param ctx context.Context Authentication Context 
 @return Broadcast*/
func (a *BroadcastApiService) GetCurrentBroadcast(ctx context.Context, ) (Broadcast,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  Broadcast
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/broadcasts/current"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


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
		return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
	 }
	
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
	 	return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* BroadcastApiService Get daily EPG
 Get current Broadcast
 * @param ctx context.Context Authentication Context 
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "date" (time.Time) Date *(Optional)*
 @return EpgBroadcast*/
func (a *BroadcastApiService) GetDailyEPG(ctx context.Context, localVarOptionals map[string]interface{}) (EpgBroadcast,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  EpgBroadcast
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/broadcasts/epg/daily"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["date"], "time.Time", "date"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["date"].(time.Time); localVarOk {
		localVarQueryParams.Add("date", parameterToString(localVarTempParam, ""))
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
		return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
	 }
	
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
	 	return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* BroadcastApiService Get EPG by date
 Get EPG by date
 * @param ctx context.Context Authentication Context 
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "date" (time.Time) Date *(Optional)*
 @return EpgBroadcast*/
func (a *BroadcastApiService) GetEPGByDate(ctx context.Context, localVarOptionals map[string]interface{}) (EpgBroadcast,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  EpgBroadcast
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/broadcasts/epg"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["date"], "time.Time", "date"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["date"].(time.Time); localVarOk {
		localVarQueryParams.Add("date", parameterToString(localVarTempParam, ""))
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
		return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
	 }
	
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
	 	return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* BroadcastApiService Get next Broadcast
 Get next Broadcast
 * @param ctx context.Context Authentication Context 
 @return Broadcast*/
func (a *BroadcastApiService) GetNextBroadcast(ctx context.Context, ) (Broadcast,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  Broadcast
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/broadcasts/next"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


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
		return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
	 }
	
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
	 	return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* BroadcastApiService Get weekly EPG
 Get weekly EPG
 * @param ctx context.Context Authentication Context 
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "date" (string) Date *(Optional)*
 @return EpgBroadcast*/
func (a *BroadcastApiService) GetWeeklyEPG(ctx context.Context, localVarOptionals map[string]interface{}) (EpgBroadcast,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  EpgBroadcast
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/broadcasts/epg/weekly"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["date"], "string", "date"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["date"].(string); localVarOk {
		localVarQueryParams.Add("date", parameterToString(localVarTempParam, ""))
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
		return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
	 }
	
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
	 	return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* BroadcastApiService Get all broadcasts.
 List all broadcasts.
 * @param ctx context.Context Authentication Context 
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "page" (int64) Current page *(Optional)*
     @param "startMin" (time.Time) Minimum start date *(Optional)*
     @param "startMax" (time.Time) Maximum start date *(Optional)*
     @param "modelTypeId" (int64) Search on ModelType ID *(Optional)*
     @param "tagId" (int64) Search on Tag ID *(Optional)* &#x60;(Relation)&#x60;
     @param "presenterId" (int64) Search on Presenter ID *(Optional)* &#x60;(Relation)&#x60;
     @param "itemId" (int64) Search on Item ID *(Optional)* &#x60;(Relation)&#x60;
     @param "blockId" (int64) Search on Block ID *(Optional)* &#x60;(Relation)&#x60;
     @param "genreId" (int64) Search on Genre ID *(Optional)* &#x60;(Relation)&#x60;
     @param "programId" (int64) Search on Program ID *(Optional)* &#x60;(Relation)&#x60;
     @param "externalStationId" (int64) Query on a different (content providing) station *(Optional)*
 @return BroadcastResults*/
func (a *BroadcastApiService) ListBroadcasts(ctx context.Context, localVarOptionals map[string]interface{}) (BroadcastResults,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  BroadcastResults
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/broadcasts"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["page"], "int64", "page"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["startMin"], "time.Time", "startMin"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["startMax"], "time.Time", "startMax"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["modelTypeId"], "int64", "modelTypeId"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["tagId"], "int64", "tagId"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["presenterId"], "int64", "presenterId"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["itemId"], "int64", "itemId"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["blockId"], "int64", "blockId"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["genreId"], "int64", "genreId"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["programId"], "int64", "programId"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["externalStationId"], "int64", "externalStationId"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["page"].(int64); localVarOk {
		localVarQueryParams.Add("page", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["startMin"].(time.Time); localVarOk {
		localVarQueryParams.Add("start-min", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["startMax"].(time.Time); localVarOk {
		localVarQueryParams.Add("start-max", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["modelTypeId"].(int64); localVarOk {
		localVarQueryParams.Add("model_type_id", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["tagId"].(int64); localVarOk {
		localVarQueryParams.Add("tag_id", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["presenterId"].(int64); localVarOk {
		localVarQueryParams.Add("presenter_id", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["itemId"].(int64); localVarOk {
		localVarQueryParams.Add("item_id", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["blockId"].(int64); localVarOk {
		localVarQueryParams.Add("block_id", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["genreId"].(int64); localVarOk {
		localVarQueryParams.Add("genre_id", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["programId"].(int64); localVarOk {
		localVarQueryParams.Add("program_id", parameterToString(localVarTempParam, ""))
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
		return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
	 }
	
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
	 	return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* BroadcastApiService Print Broadcast by id
 Print Broadcast by id
 * @param ctx context.Context Authentication Context 
 @param id ID of Broadcast **(Required)**
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "programId" (int64) Search on Program ID *(Optional)* &#x60;(Relation)&#x60;
     @param "presenterId" (int64) Search on Presenter ID *(Optional)* &#x60;(Relation)&#x60;
     @param "tagId" (int64) Search on Tag ID *(Optional)* &#x60;(Relation)&#x60;
 @return EpgBroadcast*/
func (a *BroadcastApiService) PrintBroadcastById(ctx context.Context, id int64, localVarOptionals map[string]interface{}) (EpgBroadcast,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  EpgBroadcast
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/broadcasts/print/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if id < 0 {
			return successPayload, nil, reportError("id must be greater than 0")
	}
	if err := typeCheckParameter(localVarOptionals["programId"], "int64", "programId"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["presenterId"], "int64", "presenterId"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["tagId"], "int64", "tagId"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["programId"].(int64); localVarOk {
		localVarQueryParams.Add("program_id", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["presenterId"].(int64); localVarOk {
		localVarQueryParams.Add("presenter_id", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["tagId"].(int64); localVarOk {
		localVarQueryParams.Add("tag_id", parameterToString(localVarTempParam, ""))
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
		return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
	 }
	
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
	 	return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* BroadcastApiService Update broadcast by id
 Update broadcast by id
 * @param ctx context.Context Authentication Context 
 @param id ID of Broadcast **(Required)**
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "data" (BroadcastDataInput) Data *(Optional)*
 @return Success*/
func (a *BroadcastApiService) UpdateBroadcastByID(ctx context.Context, id int64, localVarOptionals map[string]interface{}) (Success,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Patch")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  Success
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/broadcasts/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if id < 0 {
			return successPayload, nil, reportError("id must be greater than 0")
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
	// body params
	if localVarTempParam, localVarOk := localVarOptionals["data"].(BroadcastDataInput); localVarOk {
		localVarPostBody = &localVarTempParam
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
		return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
	 }
	
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
	 	return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

