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

type ContactDataInput struct {

	ModelTypeId int64 `json:"model_type_id"`

	FieldValues ContactFieldValues `json:"field_values,omitempty"`

	Email string `json:"email,omitempty"`

	Firstname string `json:"firstname"`

	Lastname string `json:"lastname"`

	Phone string `json:"phone,omitempty"`
}
