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

type BlockResults struct {

	CurrentPage int64 `json:"current_page,omitempty"`

	From int64 `json:"from,omitempty"`

	LastPage int64 `json:"last_page,omitempty"`

	NextPageUrl string `json:"next_page_url,omitempty"`

	Path string `json:"path,omitempty"`

	PerPage int64 `json:"per_page,omitempty"`

	PrevPageUrl string `json:"prev_page_url,omitempty"`

	To int64 `json:"to,omitempty"`

	Total int64 `json:"total,omitempty"`

	Results []BlockResult `json:"results,omitempty"`
}
