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

type ItemOutputOnly struct {

	Id int64 `json:"id,omitempty"`

	UpdatedAt time.Time `json:"updated_at,omitempty"`

	CreatedAt time.Time `json:"created_at,omitempty"`

	DeletedAt time.Time `json:"deleted_at,omitempty"`

	DataModified int64 `json:"data_modified,omitempty"`

	Order int64 `json:"order,omitempty"`

	TemplateblockId int64 `json:"templateblock_id,omitempty"`

	TemplateitemId int64 `json:"templateitem_id,omitempty"`

	ExternalStationId int64 `json:"_external_station_id,omitempty"`
}
