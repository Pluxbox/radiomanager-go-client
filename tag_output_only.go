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

type TagOutputOnly struct {

	Id int64 `json:"id"`

	CreatedAt time.Time `json:"created_at,omitempty"`

	UpdatedAt time.Time `json:"updated_at,omitempty"`

	DeletedAt time.Time `json:"deleted_at,omitempty"`

	ExternalStationId int64 `json:"_external_station_id,omitempty"`
}
