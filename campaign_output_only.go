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

type CampaignOutputOnly struct {

	Id int64 `json:"id"`

	UpdatedAt time.Time `json:"updated_at"`

	CreatedAt time.Time `json:"created_at"`

	DeletedAt time.Time `json:"deleted_at"`

	Item *CampaignTemplateItem `json:"item,omitempty"`

	ExternalStationId int64 `json:"_external_station_id,omitempty"`
}
