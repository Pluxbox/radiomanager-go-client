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

type ExternalMessageQueueData struct {

	// Message **(Required)**
	Message string `json:"message,omitempty"`

	// ID of Message **(Required)**
	MessageId int32 `json:"message_id,omitempty"`

	// Message Date *(Optional)*
	MessageDate time.Time `json:"message_date,omitempty"`

	// Sender **(Required)**
	SenderName string `json:"sender_name,omitempty"`

	// ID of Sender *(Optional)*
	SenderId int32 `json:"sender_id,omitempty"`

	// URL of Avatar *(Optional)*
	AvatarUrl string `json:"avatar_url,omitempty"`

	// Assigned Model Type ID **(Required)**
	ModelTypeId int32 `json:"model_type_id,omitempty"`
}
