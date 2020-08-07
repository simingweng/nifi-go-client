# ConnectionEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Revision** | [**RevisionDto**](RevisionDTO.md) |  | [optional] 
**Id** | **string** | The id of the component. | [optional] 
**Uri** | **string** | The URI for futures requests to the component. | [optional] 
**Position** | [**PositionDto**](PositionDTO.md) |  | [optional] 
**Permissions** | [**PermissionsDto**](PermissionsDTO.md) |  | [optional] 
**Bulletins** | [**[]BulletinEntity**](BulletinEntity.md) | The bulletins for this component. | [optional] 
**DisconnectedNodeAcknowledged** | **bool** | Acknowledges that this node is disconnected to allow for mutable requests to proceed. | [optional] 
**Component** | [**ConnectionDto**](ConnectionDTO.md) |  | [optional] 
**Status** | [**ConnectionStatusDto**](ConnectionStatusDTO.md) |  | [optional] 
**Bends** | [**[]PositionDto**](PositionDTO.md) | The bend points on the connection. | [optional] 
**LabelIndex** | **int32** | The index of the bend point where to place the connection label. | [optional] 
**GetzIndex** | **int64** | The z index of the connection. | [optional] 
**SourceId** | **string** | The identifier of the source of this connection. | [optional] 
**SourceGroupId** | **string** | The identifier of the group of the source of this connection. | [optional] 
**SourceType** | **string** | The type of component the source connectable is. | 
**DestinationId** | **string** | The identifier of the destination of this connection. | [optional] 
**DestinationGroupId** | **string** | The identifier of the group of the destination of this connection. | [optional] 
**DestinationType** | **string** | The type of component the destination connectable is. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


