/* 
 * Pluxbox Radiomanager Client
 *
 * Pluxbox RadioManager gives you the power, flexibility and speed you always wanted in a lightweight and easy-to-use web-based radio solution. With Pluxbox RadioManager you can organise your radio workflow and automate your omnichannel communication with your listeners. We offer wide range specialised services for the radio and connections like Hybrid Radio, Visual Radio, your website and social media without losing focus on your broadcast. For more information visit https://pluxbox.com
 *
 * OpenAPI spec version: 2.0
 * Contact: support@pluxbox.com
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package radiomanagerclient

import (
	"time"
)

type Block struct {

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
}
