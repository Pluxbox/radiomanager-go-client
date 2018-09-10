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

type Forbidden struct {

	// Given error, describes the problem.
	Error_ string `json:"error"`

	// Assigned StatusCode, is used in order to create a relationship between Error and Response.
	StatusCode int64 `json:"statusCode"`
}
