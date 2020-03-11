# ProgramResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**UpdatedAt** | [**time.Time**](time.Time.md) |  | 
**CreatedAt** | [**time.Time**](time.Time.md) |  | 
**DeletedAt** | [**time.Time**](time.Time.md) |  | 
**ExternalStationId** | **int64** |  | [optional] 
**ModelTypeId** | **int64** |  | 
**FieldValues** | [**map[string]interface{}**](.md) |  | [optional] 
**Title** | **string** |  | 
**Disabled** | **bool** |  | [optional] 
**GenreId** | **int64** |  | [optional] 
**Description** | **string** |  | [optional] 
**ShortName** | **string** |  | [optional] 
**MediumName** | **string** |  | [optional] 
**Website** | **string** |  | [optional] 
**Email** | **string** |  | [optional] 
**Recommended** | **bool** |  | [optional] 
**Language** | **string** |  | [optional] 
**PtyCodeId** | **int64** |  | [optional] 
**Genre** | [**BroadcastRelationsGenre**](BroadcastRelations_genre.md) |  | [optional] 
**Items** | [**ProgramRelationsItems**](ProgramRelations_items.md) |  | [optional] 
**Blocks** | [**ProgramRelationsBlocks**](ProgramRelations_blocks.md) |  | [optional] 
**Broadcasts** | [**ProgramRelationsBroadcasts**](ProgramRelations_broadcasts.md) |  | [optional] 
**Presenters** | [**ProgramRelationsPresenters**](ProgramRelations_presenters.md) |  | [optional] 
**Tags** | [**ProgramRelationsTags**](ProgramRelations_tags.md) |  | [optional] 
**ModelType** | [**BroadcastRelationsModelType**](BroadcastRelations_model_type.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


