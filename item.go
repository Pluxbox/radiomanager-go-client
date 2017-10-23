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

type Item struct {

	ModelTypeId int64 `json:"model_type_id"`

	BlockId int64 `json:"block_id,omitempty"`

	ExternalId string `json:"external_id,omitempty"`

	FieldValues interface{} `json:"field_values,omitempty"`

	Title string `json:"title,omitempty"`

	Duration int64 `json:"duration,omitempty"`

	Start time.Time `json:"start,omitempty"`

	Status string `json:"status,omitempty"`

	Import_ int64 `json:"import,omitempty"`

	CampaignId int64 `json:"campaign_id,omitempty"`

	Recommended bool `json:"recommended,omitempty"`

	StationDraftId int64 `json:"station_draft_id,omitempty"`

	ProgramDraftId int64 `json:"program_draft_id,omitempty"`

	UserDraftId int64 `json:"user_draft_id,omitempty"`

	StaticStart bool `json:"static_start,omitempty"`

	Details string `json:"details,omitempty"`
}
