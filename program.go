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

type Program struct {

	ModelTypeId int64 `json:"model_type_id"`

	FieldValues *ProgramFieldValues `json:"field_values,omitempty"`

	Title string `json:"title"`

	Disabled bool `json:"disabled,omitempty"`

	GenreId int64 `json:"genre_id,omitempty"`

	Description string `json:"description,omitempty"`

	ShortName string `json:"short_name,omitempty"`

	MediumName string `json:"medium_name,omitempty"`

	Website string `json:"website,omitempty"`

	Email string `json:"email,omitempty"`

	Recommended bool `json:"recommended,omitempty"`

	Language string `json:"language,omitempty"`

	PtyCodeId int64 `json:"pty_code_id,omitempty"`
}
