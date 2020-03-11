# BroadcastResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**UpdatedAt** | [**time.Time**](time.Time.md) |  | 
**CreatedAt** | [**time.Time**](time.Time.md) |  | 
**DeletedAt** | [**time.Time**](time.Time.md) |  | 
**ExternalStationId** | **int64** |  | [optional] 
**ProgramId** | **int64** |  | [optional] 
**ModelTypeId** | **int64** |  | [optional] 
**StationId** | **int64** |  | [optional] 
**FieldValues** | [**map[string]interface{}**](.md) |  | [optional] 
**Title** | **string** |  | [optional] 
**Start** | [**time.Time**](time.Time.md) |  | [optional] 
**Stop** | [**time.Time**](time.Time.md) |  | [optional] 
**GenreId** | **int64** |  | [optional] 
**Description** | **string** |  | [optional] 
**ShortName** | **string** |  | [optional] 
**MediumName** | **string** |  | [optional] 
**Website** | **string** |  | [optional] 
**Email** | **string** |  | [optional] 
**Recommended** | **bool** |  | [optional] 
**Language** | **string** |  | [optional] 
**Published** | **bool** |  | [optional] 
**RepetitionUid** | **string** |  | [optional] 
**RepetitionType** | **string** |  | [optional] 
**RepetitionEnd** | [**time.Time**](time.Time.md) |  | [optional] 
**RepetitionStart** | [**time.Time**](time.Time.md) |  | [optional] 
**RepetitionDays** | **string** |  | [optional] 
**PtyCodeId** | **int64** |  | [optional] 
**Genre** | [**BroadcastRelationsGenre**](BroadcastRelations_genre.md) |  | [optional] 
**Items** | [**BroadcastRelationsItems**](BroadcastRelations_items.md) |  | [optional] 
**Blocks** | [**BroadcastRelationsBlocks**](BroadcastRelations_blocks.md) |  | [optional] 
**Program** | [**BlockRelationsProgram**](BlockRelations_program.md) |  | [optional] 
**Tags** | [**BroadcastRelationsTags**](BroadcastRelations_tags.md) |  | [optional] 
**Presenters** | [**BroadcastRelationsPresenters**](BroadcastRelations_presenters.md) |  | [optional] 
**ModelType** | [**BroadcastRelationsModelType**](BroadcastRelations_model_type.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


