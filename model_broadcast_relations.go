/*
 * RadioManager
 *
 * RadioManager
 *
 * API version: 2.0
 * Contact: support@pluxbox.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package radiomanagerclient
// BroadcastRelations struct for BroadcastRelations
type BroadcastRelations struct {
	Genre BroadcastRelationsGenre `json:"genre,omitempty"`
	Items BroadcastRelationsItems `json:"items,omitempty"`
	Blocks BroadcastRelationsBlocks `json:"blocks,omitempty"`
	Program BlockRelationsProgram `json:"program,omitempty"`
	Tags BroadcastRelationsTags `json:"tags,omitempty"`
	Presenters BroadcastRelationsPresenters `json:"presenters,omitempty"`
	ModelType BroadcastRelationsModelType `json:"model_type,omitempty"`
}
