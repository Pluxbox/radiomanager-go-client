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

type StoryResult struct {

	Id int64 `json:"id,omitempty"`

	CreatedAt time.Time `json:"created_at,omitempty"`

	UpdatedAt time.Time `json:"updated_at,omitempty"`

	DeletedAt time.Time `json:"deleted_at,omitempty"`

	ExternalStationId int64 `json:"_external_station_id,omitempty"`

	ModelTypeId int64 `json:"model_type_id"`

	Recommended bool `json:"recommended,omitempty"`

	FieldValues *interface{} `json:"field_values,omitempty"`

	Name string `json:"name"`

	Description string `json:"description,omitempty"`

	Tags *StoryRelationsTags `json:"tags"`

	Items *StoryRelationsItems `json:"items,omitempty"`

	ModelType *BroadcastRelationsModelType `json:"model_type,omitempty"`
}
