/* 
 * Pluxbox Radiomanager Client
 *
 * Pluxbox RadioManager gives you the power, flexibility and speed you always wanted in a lightweight and easy-to-use web-based radio solution. With Pluxbox RadioManager you can organise your radio workflow and automate your omnichannel communication with your listeners. We offer wide range specialised services for the radio and connections like Hybrid Radio, Visual Radio, your website and social media without losing focus on your broadcast. For more information visit https://pluxbox.com
 *
 * OpenAPI spec version: 2.0
 * Contact: support@pluxbox.com
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package radiomanagerclient

import (
	"time"
)

type Broadcast struct {

	ProgramId int64 `json:"program_id,omitempty"`

	ModelTypeId int64 `json:"model_type_id,omitempty"`

	StationId int64 `json:"station_id,omitempty"`

	FieldValues BroadcastFieldValues `json:"field_values,omitempty"`

	Title string `json:"title,omitempty"`

	Start time.Time `json:"start,omitempty"`

	Stop time.Time `json:"stop,omitempty"`

	GenreId int64 `json:"genre_id,omitempty"`

	Description string `json:"description,omitempty"`

	ShortName string `json:"short_name,omitempty"`

	MediumName string `json:"medium_name,omitempty"`

	Website string `json:"website,omitempty"`

	Email string `json:"email,omitempty"`

	Recommended bool `json:"recommended,omitempty"`

	Language string `json:"language,omitempty"`

	Published bool `json:"published,omitempty"`

	RepetitionUid string `json:"repetition_uid,omitempty"`

	RepetitionType string `json:"repetition_type,omitempty"`

	RepetitionEnd time.Time `json:"repetition_end,omitempty"`

	RepetitionStart time.Time `json:"repetition_start,omitempty"`

	RepetitionDays []bool `json:"repetition_days,omitempty"`

	PtyCodeId int64 `json:"pty_code_id,omitempty"`
}
