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

type StationResultStation struct {

	Id int32 `json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	CreatedAt time.Time `json:"created_at,omitempty"`

	UpdatedAt time.Time `json:"updated_at,omitempty"`

	SystemName string `json:"system_name,omitempty"`

	ShortName string `json:"short_name,omitempty"`

	MediumName string `json:"medium_name,omitempty"`

	Website string `json:"website,omitempty"`

	Email string `json:"email,omitempty"`

	Keywords []string `json:"keywords,omitempty"`

	Description string `json:"description,omitempty"`

	Sms string `json:"sms,omitempty"`

	Telephone string `json:"telephone,omitempty"`

	GenreId int32 `json:"genre_id,omitempty"`

	Language string `json:"language,omitempty"`

	Active bool `json:"active,omitempty"`

	LogoRectangle string `json:"logo_rectangle,omitempty"`

	Logo128x128 string `json:"logo_128x128,omitempty"`

	Logo320x320 string `json:"logo_320x320,omitempty"`

	Logo600x600 string `json:"logo_600x600,omitempty"`

	PayOff string `json:"pay_off,omitempty"`

	PtyCode int32 `json:"pty_code,omitempty"`

	PtyType string `json:"pty_type,omitempty"`

	StationKey string `json:"station_key,omitempty"`

	Timezone string `json:"timezone,omitempty"`

	MetadataradioOrganisation string `json:"metadataradio_organisation,omitempty"`

	MetadataradioStationId string `json:"metadataradio_station_id,omitempty"`

	StartDays *StationResultStationStartDays `json:"start_days,omitempty"`
}
