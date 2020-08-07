# ProcessorDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id of the component. | [optional] 
**VersionedComponentId** | **string** | The ID of the corresponding component that is under version control | [optional] 
**ParentGroupId** | **string** | The id of parent process group of this component if applicable. | [optional] 
**Position** | [**PositionDto**](PositionDTO.md) |  | [optional] 
**Name** | **string** | The name of the processor. | [optional] 
**Type** | **string** | The type of the processor. | [optional] 
**Bundle** | [**BundleDto**](BundleDTO.md) |  | [optional] 
**State** | **string** | The state of the processor | [optional] 
**Style** | **map[string]string** | Styles for the processor (background-color : #eee). | [optional] 
**Relationships** | [**[]RelationshipDto**](RelationshipDTO.md) | The available relationships that the processor currently supports. | [optional] [readonly] 
**Description** | **string** | The description of the processor. | [optional] 
**SupportsParallelProcessing** | **bool** | Whether the processor supports parallel processing. | [optional] 
**SupportsEventDriven** | **bool** | Whether the processor supports event driven scheduling. | [optional] 
**SupportsBatching** | **bool** | Whether the processor supports batching. This makes the run duration settings available. | [optional] 
**PersistsState** | **bool** | Whether the processor persists state. | [optional] 
**Restricted** | **bool** | Whether the processor requires elevated privileges. | [optional] 
**Deprecated** | **bool** | Whether the processor has been deprecated. | [optional] 
**ExecutionNodeRestricted** | **bool** | Indicates if the execution node of a processor is restricted to run only on the primary node | [optional] 
**MultipleVersionsAvailable** | **bool** | Whether the processor has multiple versions available. | [optional] 
**InputRequirement** | **string** | The input requirement for this processor. | [optional] 
**Config** | [**ProcessorConfigDto**](ProcessorConfigDTO.md) |  | [optional] 
**ValidationErrors** | **[]string** | The validation errors for the processor. These validation errors represent the problems with the processor that must be resolved before it can be started. | [optional] 
**ValidationStatus** | **string** | Indicates whether the Processor is valid, invalid, or still in the process of validating (i.e., it is unknown whether or not the Processor is valid) | [optional] [readonly] 
**ExtensionMissing** | **bool** | Whether the underlying extension is missing. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


