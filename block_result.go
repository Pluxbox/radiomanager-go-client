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

type BlockResult struct {

	// ID of the current Block.
	Id int64 `json:"id"`

	// Currently assigned Broadcast connected to the current Block, identified by the Broadcast ID.
	BroadcastId int64 `json:"broadcast_id"`

	// Start of the Block (formatted as a DateTime object), saved with an TimeZone.
	Start time.Time `json:"start"`

	// End of the Block (formatted as a DateTime object), saved with an TimeZone.
	Stop time.Time `json:"stop"`

	// Time of the creation of the Block (formatted as a DateTime object), saved with an TimeZone.
	CreatedAt time.Time `json:"created_at"`

	// Time of the last update of the Block (formatted as a DateTime object), saved with an TimeZone.
	UpdatedAt time.Time `json:"updated_at"`

	// Moment when the Block got deleted (formatted as a DateTime object), saved with an TimeZone.
	DeletedAt time.Time `json:"deleted_at"`

	ExternalStationId int64 `json:"_external_station_id,omitempty"`

	Items *BlockRelationsItems `json:"items,omitempty"`

	Broadcast *BlockRelationsBroadcast `json:"broadcast,omitempty"`

	Program *BlockRelationsProgram `json:"program,omitempty"`
}
