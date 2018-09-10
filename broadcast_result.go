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

type BroadcastResult struct {

	Id int64 `json:"id"`

	UpdatedAt time.Time `json:"updated_at"`

	CreatedAt time.Time `json:"created_at"`

	DeletedAt time.Time `json:"deleted_at"`

	ExternalStationId int64 `json:"_external_station_id,omitempty"`

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

	Genre *BroadcastRelationsGenre `json:"genre,omitempty"`

	Items *BroadcastRelationsItems `json:"items,omitempty"`

	Blocks *BroadcastRelationsBlocks `json:"blocks,omitempty"`

	Program *BlockRelationsProgram `json:"program,omitempty"`

	Tags *BroadcastRelationsTags `json:"tags,omitempty"`

	Presenters *BroadcastRelationsPresenters `json:"presenters,omitempty"`

	ModelType *BroadcastRelationsModelType `json:"model_type,omitempty"`
}
