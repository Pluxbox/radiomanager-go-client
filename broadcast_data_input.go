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
	"time"
)

type BroadcastDataInput struct {

	ProgramId int64 `json:"program_id,omitempty"`

	ModelTypeId int64 `json:"model_type_id,omitempty"`

	StationId int64 `json:"station_id,omitempty"`

	FieldValues *interface{} `json:"field_values,omitempty"`

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

	RepetitionDays string `json:"repetition_days,omitempty"`

	PtyCodeId int64 `json:"pty_code_id,omitempty"`

	Tags []int32 `json:"tags,omitempty"`

	Presenters []int32 `json:"presenters,omitempty"`
}
