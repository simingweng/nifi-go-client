# RemoteProcessGroupEntity

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
**Component** | [**RemoteProcessGroupDto**](RemoteProcessGroupDTO.md) |  | [optional] 
**Status** | [**RemoteProcessGroupStatusDto**](RemoteProcessGroupStatusDTO.md) |  | [optional] 
**InputPortCount** | **int32** | The number of remote input ports currently available on the target. | [optional] 
**OutputPortCount** | **int32** | The number of remote output ports currently available on the target. | [optional] 
**OperatePermissions** | [**PermissionsDto**](PermissionsDTO.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


