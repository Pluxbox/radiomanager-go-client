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

type StoryDataInput struct {

	ModelTypeId int64 `json:"model_type_id"`

	Recommended bool `json:"recommended,omitempty"`

	FieldValues *interface{} `json:"field_values,omitempty"`

	Name string `json:"name"`

	Description string `json:"description,omitempty"`

	Tags []int32 `json:"tags,omitempty"`
}
