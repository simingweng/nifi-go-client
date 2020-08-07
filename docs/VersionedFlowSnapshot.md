# VersionedFlowSnapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SnapshotMetadata** | [**VersionedFlowSnapshotMetadata**](VersionedFlowSnapshotMetadata.md) |  | 
**FlowContents** | [**VersionedProcessGroup**](VersionedProcessGroup.md) |  | 
**ExternalControllerServices** | [**map[string]ExternalControllerServiceReference**](ExternalControllerServiceReference.md) | The information about controller services that exist outside this versioned flow, but are referenced by components within the versioned flow. | [optional] 
**ParameterContexts** | [**map[string]VersionedParameterContext**](VersionedParameterContext.md) | The parameter contexts referenced by process groups in the flow contents. The mapping is from the name of the context to the context instance, and it is expected that any context in this map is referenced by at least one process group in this flow. | [optional] 
**FlowEncodingVersion** | **string** | The optional encoding version of the flow contents. | [optional] 
**Flow** | [**VersionedFlow**](VersionedFlow.md) |  | [optional] 
**Bucket** | [**Bucket**](Bucket.md) |  | [optional] 
**Latest** | **bool** |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


