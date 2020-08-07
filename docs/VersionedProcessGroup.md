# VersionedProcessGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | **string** | The component&#39;s unique identifier | [optional] 
**Name** | **string** | The component&#39;s name | [optional] 
**Comments** | **string** | The user-supplied comments for the component | [optional] 
**Position** | [**Position**](Position.md) |  | [optional] 
**ProcessGroups** | [**[]VersionedProcessGroup**](VersionedProcessGroup.md) | The child Process Groups | [optional] 
**RemoteProcessGroups** | [**[]VersionedRemoteProcessGroup**](VersionedRemoteProcessGroup.md) | The Remote Process Groups | [optional] 
**Processors** | [**[]VersionedProcessor**](VersionedProcessor.md) | The Processors | [optional] 
**InputPorts** | [**[]VersionedPort**](VersionedPort.md) | The Input Ports | [optional] 
**OutputPorts** | [**[]VersionedPort**](VersionedPort.md) | The Output Ports | [optional] 
**Connections** | [**[]VersionedConnection**](VersionedConnection.md) | The Connections | [optional] 
**Labels** | [**[]VersionedLabel**](VersionedLabel.md) | The Labels | [optional] 
**Funnels** | [**[]VersionedFunnel**](VersionedFunnel.md) | The Funnels | [optional] 
**ControllerServices** | [**[]VersionedControllerService**](VersionedControllerService.md) | The Controller Services | [optional] 
**VersionedFlowCoordinates** | [**VersionedFlowCoordinates**](VersionedFlowCoordinates.md) |  | [optional] 
**Variables** | **map[string]string** | The Variables in the Variable Registry for this Process Group (not including any ancestor or descendant Process Groups) | [optional] 
**ParameterContextName** | **string** | The name of the parameter context used by this process group | [optional] 
**ComponentType** | **string** |  | [optional] 
**GroupIdentifier** | **string** | The ID of the Process Group that this component belongs to | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


