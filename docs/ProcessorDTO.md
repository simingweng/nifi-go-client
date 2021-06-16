# ProcessorDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The id of the component. | [optional] 
**VersionedComponentId** | Pointer to **string** | The ID of the corresponding component that is under version control | [optional] 
**ParentGroupId** | Pointer to **string** | The id of parent process group of this component if applicable. | [optional] 
**Position** | Pointer to [**PositionDTO**](PositionDTO.md) |  | [optional] 
**Name** | Pointer to **string** | The name of the processor. | [optional] 
**Type** | Pointer to **string** | The type of the processor. | [optional] 
**Bundle** | Pointer to [**BundleDTO**](BundleDTO.md) |  | [optional] 
**State** | Pointer to **string** | The state of the processor | [optional] 
**Style** | Pointer to **map[string]string** | Styles for the processor (background-color : #eee). | [optional] 
**Relationships** | Pointer to [**[]RelationshipDTO**](RelationshipDTO.md) | The available relationships that the processor currently supports. | [optional] [readonly] 
**Description** | Pointer to **string** | The description of the processor. | [optional] 
**SupportsParallelProcessing** | Pointer to **bool** | Whether the processor supports parallel processing. | [optional] 
**SupportsEventDriven** | Pointer to **bool** | Whether the processor supports event driven scheduling. | [optional] 
**SupportsBatching** | Pointer to **bool** | Whether the processor supports batching. This makes the run duration settings available. | [optional] 
**PersistsState** | Pointer to **bool** | Whether the processor persists state. | [optional] 
**Restricted** | Pointer to **bool** | Whether the processor requires elevated privileges. | [optional] 
**Deprecated** | Pointer to **bool** | Whether the processor has been deprecated. | [optional] 
**ExecutionNodeRestricted** | Pointer to **bool** | Indicates if the execution node of a processor is restricted to run only on the primary node | [optional] 
**MultipleVersionsAvailable** | Pointer to **bool** | Whether the processor has multiple versions available. | [optional] 
**InputRequirement** | Pointer to **string** | The input requirement for this processor. | [optional] 
**Config** | Pointer to [**ProcessorConfigDTO**](ProcessorConfigDTO.md) |  | [optional] 
**ValidationErrors** | Pointer to **[]string** | The validation errors for the processor. These validation errors represent the problems with the processor that must be resolved before it can be started. | [optional] 
**ValidationStatus** | Pointer to **string** | Indicates whether the Processor is valid, invalid, or still in the process of validating (i.e., it is unknown whether or not the Processor is valid) | [optional] [readonly] 
**ExtensionMissing** | Pointer to **bool** | Whether the underlying extension is missing. | [optional] 

## Methods

### NewProcessorDTO

`func NewProcessorDTO() *ProcessorDTO`

NewProcessorDTO instantiates a new ProcessorDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessorDTOWithDefaults

`func NewProcessorDTOWithDefaults() *ProcessorDTO`

NewProcessorDTOWithDefaults instantiates a new ProcessorDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProcessorDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProcessorDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProcessorDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProcessorDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetVersionedComponentId

`func (o *ProcessorDTO) GetVersionedComponentId() string`

GetVersionedComponentId returns the VersionedComponentId field if non-nil, zero value otherwise.

### GetVersionedComponentIdOk

`func (o *ProcessorDTO) GetVersionedComponentIdOk() (*string, bool)`

GetVersionedComponentIdOk returns a tuple with the VersionedComponentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionedComponentId

`func (o *ProcessorDTO) SetVersionedComponentId(v string)`

SetVersionedComponentId sets VersionedComponentId field to given value.

### HasVersionedComponentId

`func (o *ProcessorDTO) HasVersionedComponentId() bool`

HasVersionedComponentId returns a boolean if a field has been set.

### GetParentGroupId

`func (o *ProcessorDTO) GetParentGroupId() string`

GetParentGroupId returns the ParentGroupId field if non-nil, zero value otherwise.

### GetParentGroupIdOk

`func (o *ProcessorDTO) GetParentGroupIdOk() (*string, bool)`

GetParentGroupIdOk returns a tuple with the ParentGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentGroupId

`func (o *ProcessorDTO) SetParentGroupId(v string)`

SetParentGroupId sets ParentGroupId field to given value.

### HasParentGroupId

`func (o *ProcessorDTO) HasParentGroupId() bool`

HasParentGroupId returns a boolean if a field has been set.

### GetPosition

`func (o *ProcessorDTO) GetPosition() PositionDTO`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *ProcessorDTO) GetPositionOk() (*PositionDTO, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *ProcessorDTO) SetPosition(v PositionDTO)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *ProcessorDTO) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetName

`func (o *ProcessorDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProcessorDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProcessorDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProcessorDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *ProcessorDTO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProcessorDTO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProcessorDTO) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ProcessorDTO) HasType() bool`

HasType returns a boolean if a field has been set.

### GetBundle

`func (o *ProcessorDTO) GetBundle() BundleDTO`

GetBundle returns the Bundle field if non-nil, zero value otherwise.

### GetBundleOk

`func (o *ProcessorDTO) GetBundleOk() (*BundleDTO, bool)`

GetBundleOk returns a tuple with the Bundle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundle

`func (o *ProcessorDTO) SetBundle(v BundleDTO)`

SetBundle sets Bundle field to given value.

### HasBundle

`func (o *ProcessorDTO) HasBundle() bool`

HasBundle returns a boolean if a field has been set.

### GetState

`func (o *ProcessorDTO) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ProcessorDTO) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ProcessorDTO) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ProcessorDTO) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStyle

`func (o *ProcessorDTO) GetStyle() map[string]string`

GetStyle returns the Style field if non-nil, zero value otherwise.

### GetStyleOk

`func (o *ProcessorDTO) GetStyleOk() (*map[string]string, bool)`

GetStyleOk returns a tuple with the Style field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyle

`func (o *ProcessorDTO) SetStyle(v map[string]string)`

SetStyle sets Style field to given value.

### HasStyle

`func (o *ProcessorDTO) HasStyle() bool`

HasStyle returns a boolean if a field has been set.

### GetRelationships

`func (o *ProcessorDTO) GetRelationships() []RelationshipDTO`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *ProcessorDTO) GetRelationshipsOk() (*[]RelationshipDTO, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *ProcessorDTO) SetRelationships(v []RelationshipDTO)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *ProcessorDTO) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetDescription

`func (o *ProcessorDTO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProcessorDTO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProcessorDTO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProcessorDTO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSupportsParallelProcessing

`func (o *ProcessorDTO) GetSupportsParallelProcessing() bool`

GetSupportsParallelProcessing returns the SupportsParallelProcessing field if non-nil, zero value otherwise.

### GetSupportsParallelProcessingOk

`func (o *ProcessorDTO) GetSupportsParallelProcessingOk() (*bool, bool)`

GetSupportsParallelProcessingOk returns a tuple with the SupportsParallelProcessing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsParallelProcessing

`func (o *ProcessorDTO) SetSupportsParallelProcessing(v bool)`

SetSupportsParallelProcessing sets SupportsParallelProcessing field to given value.

### HasSupportsParallelProcessing

`func (o *ProcessorDTO) HasSupportsParallelProcessing() bool`

HasSupportsParallelProcessing returns a boolean if a field has been set.

### GetSupportsEventDriven

`func (o *ProcessorDTO) GetSupportsEventDriven() bool`

GetSupportsEventDriven returns the SupportsEventDriven field if non-nil, zero value otherwise.

### GetSupportsEventDrivenOk

`func (o *ProcessorDTO) GetSupportsEventDrivenOk() (*bool, bool)`

GetSupportsEventDrivenOk returns a tuple with the SupportsEventDriven field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsEventDriven

`func (o *ProcessorDTO) SetSupportsEventDriven(v bool)`

SetSupportsEventDriven sets SupportsEventDriven field to given value.

### HasSupportsEventDriven

`func (o *ProcessorDTO) HasSupportsEventDriven() bool`

HasSupportsEventDriven returns a boolean if a field has been set.

### GetSupportsBatching

`func (o *ProcessorDTO) GetSupportsBatching() bool`

GetSupportsBatching returns the SupportsBatching field if non-nil, zero value otherwise.

### GetSupportsBatchingOk

`func (o *ProcessorDTO) GetSupportsBatchingOk() (*bool, bool)`

GetSupportsBatchingOk returns a tuple with the SupportsBatching field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsBatching

`func (o *ProcessorDTO) SetSupportsBatching(v bool)`

SetSupportsBatching sets SupportsBatching field to given value.

### HasSupportsBatching

`func (o *ProcessorDTO) HasSupportsBatching() bool`

HasSupportsBatching returns a boolean if a field has been set.

### GetPersistsState

`func (o *ProcessorDTO) GetPersistsState() bool`

GetPersistsState returns the PersistsState field if non-nil, zero value otherwise.

### GetPersistsStateOk

`func (o *ProcessorDTO) GetPersistsStateOk() (*bool, bool)`

GetPersistsStateOk returns a tuple with the PersistsState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistsState

`func (o *ProcessorDTO) SetPersistsState(v bool)`

SetPersistsState sets PersistsState field to given value.

### HasPersistsState

`func (o *ProcessorDTO) HasPersistsState() bool`

HasPersistsState returns a boolean if a field has been set.

### GetRestricted

`func (o *ProcessorDTO) GetRestricted() bool`

GetRestricted returns the Restricted field if non-nil, zero value otherwise.

### GetRestrictedOk

`func (o *ProcessorDTO) GetRestrictedOk() (*bool, bool)`

GetRestrictedOk returns a tuple with the Restricted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestricted

`func (o *ProcessorDTO) SetRestricted(v bool)`

SetRestricted sets Restricted field to given value.

### HasRestricted

`func (o *ProcessorDTO) HasRestricted() bool`

HasRestricted returns a boolean if a field has been set.

### GetDeprecated

`func (o *ProcessorDTO) GetDeprecated() bool`

GetDeprecated returns the Deprecated field if non-nil, zero value otherwise.

### GetDeprecatedOk

`func (o *ProcessorDTO) GetDeprecatedOk() (*bool, bool)`

GetDeprecatedOk returns a tuple with the Deprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecated

`func (o *ProcessorDTO) SetDeprecated(v bool)`

SetDeprecated sets Deprecated field to given value.

### HasDeprecated

`func (o *ProcessorDTO) HasDeprecated() bool`

HasDeprecated returns a boolean if a field has been set.

### GetExecutionNodeRestricted

`func (o *ProcessorDTO) GetExecutionNodeRestricted() bool`

GetExecutionNodeRestricted returns the ExecutionNodeRestricted field if non-nil, zero value otherwise.

### GetExecutionNodeRestrictedOk

`func (o *ProcessorDTO) GetExecutionNodeRestrictedOk() (*bool, bool)`

GetExecutionNodeRestrictedOk returns a tuple with the ExecutionNodeRestricted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionNodeRestricted

`func (o *ProcessorDTO) SetExecutionNodeRestricted(v bool)`

SetExecutionNodeRestricted sets ExecutionNodeRestricted field to given value.

### HasExecutionNodeRestricted

`func (o *ProcessorDTO) HasExecutionNodeRestricted() bool`

HasExecutionNodeRestricted returns a boolean if a field has been set.

### GetMultipleVersionsAvailable

`func (o *ProcessorDTO) GetMultipleVersionsAvailable() bool`

GetMultipleVersionsAvailable returns the MultipleVersionsAvailable field if non-nil, zero value otherwise.

### GetMultipleVersionsAvailableOk

`func (o *ProcessorDTO) GetMultipleVersionsAvailableOk() (*bool, bool)`

GetMultipleVersionsAvailableOk returns a tuple with the MultipleVersionsAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultipleVersionsAvailable

`func (o *ProcessorDTO) SetMultipleVersionsAvailable(v bool)`

SetMultipleVersionsAvailable sets MultipleVersionsAvailable field to given value.

### HasMultipleVersionsAvailable

`func (o *ProcessorDTO) HasMultipleVersionsAvailable() bool`

HasMultipleVersionsAvailable returns a boolean if a field has been set.

### GetInputRequirement

`func (o *ProcessorDTO) GetInputRequirement() string`

GetInputRequirement returns the InputRequirement field if non-nil, zero value otherwise.

### GetInputRequirementOk

`func (o *ProcessorDTO) GetInputRequirementOk() (*string, bool)`

GetInputRequirementOk returns a tuple with the InputRequirement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputRequirement

`func (o *ProcessorDTO) SetInputRequirement(v string)`

SetInputRequirement sets InputRequirement field to given value.

### HasInputRequirement

`func (o *ProcessorDTO) HasInputRequirement() bool`

HasInputRequirement returns a boolean if a field has been set.

### GetConfig

`func (o *ProcessorDTO) GetConfig() ProcessorConfigDTO`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ProcessorDTO) GetConfigOk() (*ProcessorConfigDTO, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ProcessorDTO) SetConfig(v ProcessorConfigDTO)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ProcessorDTO) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetValidationErrors

`func (o *ProcessorDTO) GetValidationErrors() []string`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *ProcessorDTO) GetValidationErrorsOk() (*[]string, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *ProcessorDTO) SetValidationErrors(v []string)`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrors

`func (o *ProcessorDTO) HasValidationErrors() bool`

HasValidationErrors returns a boolean if a field has been set.

### GetValidationStatus

`func (o *ProcessorDTO) GetValidationStatus() string`

GetValidationStatus returns the ValidationStatus field if non-nil, zero value otherwise.

### GetValidationStatusOk

`func (o *ProcessorDTO) GetValidationStatusOk() (*string, bool)`

GetValidationStatusOk returns a tuple with the ValidationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationStatus

`func (o *ProcessorDTO) SetValidationStatus(v string)`

SetValidationStatus sets ValidationStatus field to given value.

### HasValidationStatus

`func (o *ProcessorDTO) HasValidationStatus() bool`

HasValidationStatus returns a boolean if a field has been set.

### GetExtensionMissing

`func (o *ProcessorDTO) GetExtensionMissing() bool`

GetExtensionMissing returns the ExtensionMissing field if non-nil, zero value otherwise.

### GetExtensionMissingOk

`func (o *ProcessorDTO) GetExtensionMissingOk() (*bool, bool)`

GetExtensionMissingOk returns a tuple with the ExtensionMissing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionMissing

`func (o *ProcessorDTO) SetExtensionMissing(v bool)`

SetExtensionMissing sets ExtensionMissing field to given value.

### HasExtensionMissing

`func (o *ProcessorDTO) HasExtensionMissing() bool`

HasExtensionMissing returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


