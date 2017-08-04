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

type ItemResult struct {

	Id int64 `json:"id,omitempty"`

	UpdatedAt time.Time `json:"updated_at,omitempty"`

	CreatedAt time.Time `json:"created_at,omitempty"`

	DeletedAt time.Time `json:"deleted_at,omitempty"`

	DataModified int64 `json:"data_modified,omitempty"`

	Order int64 `json:"order,omitempty"`

	TemplateblockId int64 `json:"templateblock_id,omitempty"`

	TemplateitemId int64 `json:"templateitem_id,omitempty"`

	ExternalStationId int64 `json:"_external_station_id,omitempty"`

	ModelTypeId int64 `json:"model_type_id"`

	BlockId int64 `json:"block_id,omitempty"`

	ExternalId int64 `json:"external_id"`

	FieldValues *ImportItemFieldValues `json:"field_values,omitempty"`

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

	Block *ItemRelationsBlock `json:"block,omitempty"`

	Broadcast *BlockRelationsBroadcast `json:"broadcast,omitempty"`

	Program *ItemRelationsProgram `json:"program,omitempty"`

	Contacts *ItemRelationsContacts `json:"contacts,omitempty"`

	Tags *ItemRelationsTags `json:"tags,omitempty"`

	Campaign *ItemRelationsCampaign `json:"campaign,omitempty"`

	ModelType *BroadcastRelationsModelType `json:"model_type,omitempty"`
}
