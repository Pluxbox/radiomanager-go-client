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

type PresenterOutputOnly struct {

	Id int64 `json:"id"`

	UpdatedAt time.Time `json:"updated_at"`

	CreatedAt time.Time `json:"created_at"`

	DeletedAt time.Time `json:"deleted_at"`

	ExternalStationId int64 `json:"_external_station_id,omitempty"`
}
