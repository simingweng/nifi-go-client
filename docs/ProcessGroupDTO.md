# ProcessGroupDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id of the component. | [optional] 
**VersionedComponentId** | **string** | The ID of the corresponding component that is under version control | [optional] 
**ParentGroupId** | **string** | The id of parent process group of this component if applicable. | [optional] 
**Position** | [**PositionDto**](PositionDTO.md) |  | [optional] 
**Name** | **string** | The name of the process group. | [optional] 
**Comments** | **string** | The comments for the process group. | [optional] 
**Variables** | **map[string]string** | The variables that are configured for the Process Group. Note that this map contains only those variables that are defined on this Process Group and not any variables that are defined in the parent Process Group, etc. I.e., this Map will not contain all variables that are accessible by components in this Process Group by rather only the variables that are defined for this Process Group itself. | [optional] [readonly] 
**VersionControlInformation** | [**VersionControlInformationDto**](VersionControlInformationDTO.md) |  | [optional] 
**ParameterContext** | [**ParameterContextReferenceEntity**](ParameterContextReferenceEntity.md) |  | [optional] 
**FlowfileConcurrency** | **string** | The FlowFile Concurrency for this Process Group. | [optional] 
**FlowfileOutboundPolicy** | **string** | The Oubound Policy that is used for determining how FlowFiles should be transferred out of the Process Group. | [optional] 
**RunningCount** | **int32** | The number of running components in this process group. | [optional] 
**StoppedCount** | **int32** | The number of stopped components in the process group. | [optional] 
**InvalidCount** | **int32** | The number of invalid components in the process group. | [optional] 
**DisabledCount** | **int32** | The number of disabled components in the process group. | [optional] 
**ActiveRemotePortCount** | **int32** | The number of active remote ports in the process group. | [optional] 
**InactiveRemotePortCount** | **int32** | The number of inactive remote ports in the process group. | [optional] 
**UpToDateCount** | **int32** | The number of up to date versioned process groups in the process group. | [optional] 
**LocallyModifiedCount** | **int32** | The number of locally modified versioned process groups in the process group. | [optional] 
**StaleCount** | **int32** | The number of stale versioned process groups in the process group. | [optional] 
**LocallyModifiedAndStaleCount** | **int32** | The number of locally modified and stale versioned process groups in the process group. | [optional] 
**SyncFailureCount** | **int32** | The number of versioned process groups in the process group that are unable to sync to a registry. | [optional] 
**LocalInputPortCount** | **int32** | The number of local input ports in the process group. | [optional] 
**LocalOutputPortCount** | **int32** | The number of local output ports in the process group. | [optional] 
**PublicInputPortCount** | **int32** | The number of public input ports in the process group. | [optional] 
**PublicOutputPortCount** | **int32** | The number of public output ports in the process group. | [optional] 
**Contents** | [**FlowSnippetDto**](FlowSnippetDTO.md) |  | [optional] 
**InputPortCount** | **int32** | The number of input ports in the process group. | [optional] [readonly] 
**OutputPortCount** | **int32** | The number of output ports in the process group. | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


