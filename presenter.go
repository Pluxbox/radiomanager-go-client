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

type Presenter struct {

	ModelTypeId int64 `json:"model_type_id"`

	FieldValues []interface{} `json:"field_values,omitempty"`

	Firstname string `json:"firstname,omitempty"`

	Lastname string `json:"lastname,omitempty"`

	Active bool `json:"active,omitempty"`

	Name string `json:"name,omitempty"`
}
