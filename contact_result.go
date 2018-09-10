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

type ContactResult struct {

	Id int64 `json:"id,omitempty"`

	CreatedAt time.Time `json:"created_at,omitempty"`

	UpdatedAt time.Time `json:"updated_at,omitempty"`

	DeletedAt time.Time `json:"deleted_at,omitempty"`

	ExternalStationId int64 `json:"_external_station_id,omitempty"`

	ModelTypeId int64 `json:"model_type_id"`

	FieldValues *interface{} `json:"field_values,omitempty"`

	Email string `json:"email,omitempty"`

	Firstname string `json:"firstname"`

	Lastname string `json:"lastname"`

	Phone string `json:"phone,omitempty"`

	Tags *ContactRelationsTags `json:"tags"`

	Items *ContactRelationsItems `json:"items,omitempty"`

	ModelType *BroadcastRelationsModelType `json:"model_type,omitempty"`
}
