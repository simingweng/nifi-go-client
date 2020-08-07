# ProcessGroupEntity

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
**Component** | [**ProcessGroupDto**](ProcessGroupDTO.md) |  | [optional] 
**Status** | [**ProcessGroupStatusDto**](ProcessGroupStatusDTO.md) |  | [optional] 
**VersionedFlowSnapshot** | [**VersionedFlowSnapshot**](VersionedFlowSnapshot.md) |  | [optional] 
**RunningCount** | **int32** | The number of running components in this process group. | [optional] 
**StoppedCount** | **int32** | The number of stopped components in the process group. | [optional] 
**InvalidCount** | **int32** | The number of invalid components in the process group. | [optional] 
**DisabledCount** | **int32** | The number of disabled components in the process group. | [optional] 
**ActiveRemotePortCount** | **int32** | The number of active remote ports in the process group. | [optional] 
**InactiveRemotePortCount** | **int32** | The number of inactive remote ports in the process group. | [optional] 
**VersionedFlowState** | **string** | The current state of the Process Group, as it relates to the Versioned Flow | [optional] [readonly] 
**UpToDateCount** | **int32** | The number of up to date versioned process groups in the process group. | [optional] 
**LocallyModifiedCount** | **int32** | The number of locally modified versioned process groups in the process group. | [optional] 
**StaleCount** | **int32** | The number of stale versioned process groups in the process group. | [optional] 
**LocallyModifiedAndStaleCount** | **int32** | The number of locally modified and stale versioned process groups in the process group. | [optional] 
**SyncFailureCount** | **int32** | The number of versioned process groups in the process group that are unable to sync to a registry. | [optional] 
**LocalInputPortCount** | **int32** | The number of local input ports in the process group. | [optional] 
**LocalOutputPortCount** | **int32** | The number of local output ports in the process group. | [optional] 
**PublicInputPortCount** | **int32** | The number of public input ports in the process group. | [optional] 
**PublicOutputPortCount** | **int32** | The number of public output ports in the process group. | [optional] 
**ParameterContext** | [**ParameterContextReferenceEntity**](ParameterContextReferenceEntity.md) |  | [optional] 
**InputPortCount** | **int32** | The number of input ports in the process group. | [optional] [readonly] 
**OutputPortCount** | **int32** | The number of output ports in the process group. | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


