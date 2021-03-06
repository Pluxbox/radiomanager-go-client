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

type ProgramRelations struct {

	Genre *BroadcastRelationsGenre `json:"genre,omitempty"`

	Items *ProgramRelationsItems `json:"items,omitempty"`

	Blocks *ProgramRelationsBlocks `json:"blocks,omitempty"`

	Broadcasts *ProgramRelationsBroadcasts `json:"broadcasts,omitempty"`

	Presenters *ProgramRelationsPresenters `json:"presenters,omitempty"`

	Tags *ProgramRelationsTags `json:"tags,omitempty"`

	ModelType *BroadcastRelationsModelType `json:"model_type,omitempty"`
}
