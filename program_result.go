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
	"time"
)

type ProgramResult struct {

	Id int64 `json:"id"`

	UpdatedAt time.Time `json:"updated_at"`

	CreatedAt time.Time `json:"created_at"`

	DeletedAt time.Time `json:"deleted_at"`

	ExternalStationId int64 `json:"_external_station_id,omitempty"`

	ModelTypeId int64 `json:"model_type_id"`

	FieldValues interface{} `json:"field_values,omitempty"`

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

	Items ProgramRelationsItems `json:"items,omitempty"`

	Blocks ProgramRelationsBlocks `json:"blocks,omitempty"`

	Broadcasts ProgramRelationsBroadcasts `json:"broadcasts,omitempty"`

	Presenters ProgramRelationsPresenters `json:"presenters,omitempty"`

	Tags ProgramRelationsTags `json:"tags,omitempty"`

	ModelType BroadcastRelationsModelType `json:"model_type,omitempty"`
}
