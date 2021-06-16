# ProcessGroupDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The id of the component. | [optional] 
**VersionedComponentId** | Pointer to **string** | The ID of the corresponding component that is under version control | [optional] 
**ParentGroupId** | Pointer to **string** | The id of parent process group of this component if applicable. | [optional] 
**Position** | Pointer to [**PositionDTO**](PositionDTO.md) |  | [optional] 
**Name** | Pointer to **string** | The name of the process group. | [optional] 
**Comments** | Pointer to **string** | The comments for the process group. | [optional] 
**Variables** | Pointer to **map[string]string** | The variables that are configured for the Process Group. Note that this map contains only those variables that are defined on this Process Group and not any variables that are defined in the parent Process Group, etc. I.e., this Map will not contain all variables that are accessible by components in this Process Group by rather only the variables that are defined for this Process Group itself. | [optional] [readonly] 
**VersionControlInformation** | Pointer to [**VersionControlInformationDTO**](VersionControlInformationDTO.md) |  | [optional] 
**ParameterContext** | Pointer to [**ParameterContextReferenceEntity**](ParameterContextReferenceEntity.md) |  | [optional] 
**FlowfileConcurrency** | Pointer to **string** | The FlowFile Concurrency for this Process Group. | [optional] 
**FlowfileOutboundPolicy** | Pointer to **string** | The Oubound Policy that is used for determining how FlowFiles should be transferred out of the Process Group. | [optional] 
**RunningCount** | Pointer to **int32** | The number of running components in this process group. | [optional] 
**StoppedCount** | Pointer to **int32** | The number of stopped components in the process group. | [optional] 
**InvalidCount** | Pointer to **int32** | The number of invalid components in the process group. | [optional] 
**DisabledCount** | Pointer to **int32** | The number of disabled components in the process group. | [optional] 
**ActiveRemotePortCount** | Pointer to **int32** | The number of active remote ports in the process group. | [optional] 
**InactiveRemotePortCount** | Pointer to **int32** | The number of inactive remote ports in the process group. | [optional] 
**UpToDateCount** | Pointer to **int32** | The number of up to date versioned process groups in the process group. | [optional] 
**LocallyModifiedCount** | Pointer to **int32** | The number of locally modified versioned process groups in the process group. | [optional] 
**StaleCount** | Pointer to **int32** | The number of stale versioned process groups in the process group. | [optional] 
**LocallyModifiedAndStaleCount** | Pointer to **int32** | The number of locally modified and stale versioned process groups in the process group. | [optional] 
**SyncFailureCount** | Pointer to **int32** | The number of versioned process groups in the process group that are unable to sync to a registry. | [optional] 
**LocalInputPortCount** | Pointer to **int32** | The number of local input ports in the process group. | [optional] 
**LocalOutputPortCount** | Pointer to **int32** | The number of local output ports in the process group. | [optional] 
**PublicInputPortCount** | Pointer to **int32** | The number of public input ports in the process group. | [optional] 
**PublicOutputPortCount** | Pointer to **int32** | The number of public output ports in the process group. | [optional] 
**Contents** | Pointer to [**FlowSnippetDTO**](FlowSnippetDTO.md) |  | [optional] 
**InputPortCount** | Pointer to **int32** | The number of input ports in the process group. | [optional] [readonly] 
**OutputPortCount** | Pointer to **int32** | The number of output ports in the process group. | [optional] [readonly] 

## Methods

### NewProcessGroupDTO

`func NewProcessGroupDTO() *ProcessGroupDTO`

NewProcessGroupDTO instantiates a new ProcessGroupDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessGroupDTOWithDefaults

`func NewProcessGroupDTOWithDefaults() *ProcessGroupDTO`

NewProcessGroupDTOWithDefaults instantiates a new ProcessGroupDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProcessGroupDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProcessGroupDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProcessGroupDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProcessGroupDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetVersionedComponentId

`func (o *ProcessGroupDTO) GetVersionedComponentId() string`

GetVersionedComponentId returns the VersionedComponentId field if non-nil, zero value otherwise.

### GetVersionedComponentIdOk

`func (o *ProcessGroupDTO) GetVersionedComponentIdOk() (*string, bool)`

GetVersionedComponentIdOk returns a tuple with the VersionedComponentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionedComponentId

`func (o *ProcessGroupDTO) SetVersionedComponentId(v string)`

SetVersionedComponentId sets VersionedComponentId field to given value.

### HasVersionedComponentId

`func (o *ProcessGroupDTO) HasVersionedComponentId() bool`

HasVersionedComponentId returns a boolean if a field has been set.

### GetParentGroupId

`func (o *ProcessGroupDTO) GetParentGroupId() string`

GetParentGroupId returns the ParentGroupId field if non-nil, zero value otherwise.

### GetParentGroupIdOk

`func (o *ProcessGroupDTO) GetParentGroupIdOk() (*string, bool)`

GetParentGroupIdOk returns a tuple with the ParentGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentGroupId

`func (o *ProcessGroupDTO) SetParentGroupId(v string)`

SetParentGroupId sets ParentGroupId field to given value.

### HasParentGroupId

`func (o *ProcessGroupDTO) HasParentGroupId() bool`

HasParentGroupId returns a boolean if a field has been set.

### GetPosition

`func (o *ProcessGroupDTO) GetPosition() PositionDTO`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *ProcessGroupDTO) GetPositionOk() (*PositionDTO, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *ProcessGroupDTO) SetPosition(v PositionDTO)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *ProcessGroupDTO) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetName

`func (o *ProcessGroupDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProcessGroupDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProcessGroupDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProcessGroupDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetComments

`func (o *ProcessGroupDTO) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *ProcessGroupDTO) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *ProcessGroupDTO) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *ProcessGroupDTO) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetVariables

`func (o *ProcessGroupDTO) GetVariables() map[string]string`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *ProcessGroupDTO) GetVariablesOk() (*map[string]string, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *ProcessGroupDTO) SetVariables(v map[string]string)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *ProcessGroupDTO) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### GetVersionControlInformation

`func (o *ProcessGroupDTO) GetVersionControlInformation() VersionControlInformationDTO`

GetVersionControlInformation returns the VersionControlInformation field if non-nil, zero value otherwise.

### GetVersionControlInformationOk

`func (o *ProcessGroupDTO) GetVersionControlInformationOk() (*VersionControlInformationDTO, bool)`

GetVersionControlInformationOk returns a tuple with the VersionControlInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionControlInformation

`func (o *ProcessGroupDTO) SetVersionControlInformation(v VersionControlInformationDTO)`

SetVersionControlInformation sets VersionControlInformation field to given value.

### HasVersionControlInformation

`func (o *ProcessGroupDTO) HasVersionControlInformation() bool`

HasVersionControlInformation returns a boolean if a field has been set.

### GetParameterContext

`func (o *ProcessGroupDTO) GetParameterContext() ParameterContextReferenceEntity`

GetParameterContext returns the ParameterContext field if non-nil, zero value otherwise.

### GetParameterContextOk

`func (o *ProcessGroupDTO) GetParameterContextOk() (*ParameterContextReferenceEntity, bool)`

GetParameterContextOk returns a tuple with the ParameterContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameterContext

`func (o *ProcessGroupDTO) SetParameterContext(v ParameterContextReferenceEntity)`

SetParameterContext sets ParameterContext field to given value.

### HasParameterContext

`func (o *ProcessGroupDTO) HasParameterContext() bool`

HasParameterContext returns a boolean if a field has been set.

### GetFlowfileConcurrency

`func (o *ProcessGroupDTO) GetFlowfileConcurrency() string`

GetFlowfileConcurrency returns the FlowfileConcurrency field if non-nil, zero value otherwise.

### GetFlowfileConcurrencyOk

`func (o *ProcessGroupDTO) GetFlowfileConcurrencyOk() (*string, bool)`

GetFlowfileConcurrencyOk returns a tuple with the FlowfileConcurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowfileConcurrency

`func (o *ProcessGroupDTO) SetFlowfileConcurrency(v string)`

SetFlowfileConcurrency sets FlowfileConcurrency field to given value.

### HasFlowfileConcurrency

`func (o *ProcessGroupDTO) HasFlowfileConcurrency() bool`

HasFlowfileConcurrency returns a boolean if a field has been set.

### GetFlowfileOutboundPolicy

`func (o *ProcessGroupDTO) GetFlowfileOutboundPolicy() string`

GetFlowfileOutboundPolicy returns the FlowfileOutboundPolicy field if non-nil, zero value otherwise.

### GetFlowfileOutboundPolicyOk

`func (o *ProcessGroupDTO) GetFlowfileOutboundPolicyOk() (*string, bool)`

GetFlowfileOutboundPolicyOk returns a tuple with the FlowfileOutboundPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowfileOutboundPolicy

`func (o *ProcessGroupDTO) SetFlowfileOutboundPolicy(v string)`

SetFlowfileOutboundPolicy sets FlowfileOutboundPolicy field to given value.

### HasFlowfileOutboundPolicy

`func (o *ProcessGroupDTO) HasFlowfileOutboundPolicy() bool`

HasFlowfileOutboundPolicy returns a boolean if a field has been set.

### GetRunningCount

`func (o *ProcessGroupDTO) GetRunningCount() int32`

GetRunningCount returns the RunningCount field if non-nil, zero value otherwise.

### GetRunningCountOk

`func (o *ProcessGroupDTO) GetRunningCountOk() (*int32, bool)`

GetRunningCountOk returns a tuple with the RunningCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningCount

`func (o *ProcessGroupDTO) SetRunningCount(v int32)`

SetRunningCount sets RunningCount field to given value.

### HasRunningCount

`func (o *ProcessGroupDTO) HasRunningCount() bool`

HasRunningCount returns a boolean if a field has been set.

### GetStoppedCount

`func (o *ProcessGroupDTO) GetStoppedCount() int32`

GetStoppedCount returns the StoppedCount field if non-nil, zero value otherwise.

### GetStoppedCountOk

`func (o *ProcessGroupDTO) GetStoppedCountOk() (*int32, bool)`

GetStoppedCountOk returns a tuple with the StoppedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoppedCount

`func (o *ProcessGroupDTO) SetStoppedCount(v int32)`

SetStoppedCount sets StoppedCount field to given value.

### HasStoppedCount

`func (o *ProcessGroupDTO) HasStoppedCount() bool`

HasStoppedCount returns a boolean if a field has been set.

### GetInvalidCount

`func (o *ProcessGroupDTO) GetInvalidCount() int32`

GetInvalidCount returns the InvalidCount field if non-nil, zero value otherwise.

### GetInvalidCountOk

`func (o *ProcessGroupDTO) GetInvalidCountOk() (*int32, bool)`

GetInvalidCountOk returns a tuple with the InvalidCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidCount

`func (o *ProcessGroupDTO) SetInvalidCount(v int32)`

SetInvalidCount sets InvalidCount field to given value.

### HasInvalidCount

`func (o *ProcessGroupDTO) HasInvalidCount() bool`

HasInvalidCount returns a boolean if a field has been set.

### GetDisabledCount

`func (o *ProcessGroupDTO) GetDisabledCount() int32`

GetDisabledCount returns the DisabledCount field if non-nil, zero value otherwise.

### GetDisabledCountOk

`func (o *ProcessGroupDTO) GetDisabledCountOk() (*int32, bool)`

GetDisabledCountOk returns a tuple with the DisabledCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledCount

`func (o *ProcessGroupDTO) SetDisabledCount(v int32)`

SetDisabledCount sets DisabledCount field to given value.

### HasDisabledCount

`func (o *ProcessGroupDTO) HasDisabledCount() bool`

HasDisabledCount returns a boolean if a field has been set.

### GetActiveRemotePortCount

`func (o *ProcessGroupDTO) GetActiveRemotePortCount() int32`

GetActiveRemotePortCount returns the ActiveRemotePortCount field if non-nil, zero value otherwise.

### GetActiveRemotePortCountOk

`func (o *ProcessGroupDTO) GetActiveRemotePortCountOk() (*int32, bool)`

GetActiveRemotePortCountOk returns a tuple with the ActiveRemotePortCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveRemotePortCount

`func (o *ProcessGroupDTO) SetActiveRemotePortCount(v int32)`

SetActiveRemotePortCount sets ActiveRemotePortCount field to given value.

### HasActiveRemotePortCount

`func (o *ProcessGroupDTO) HasActiveRemotePortCount() bool`

HasActiveRemotePortCount returns a boolean if a field has been set.

### GetInactiveRemotePortCount

`func (o *ProcessGroupDTO) GetInactiveRemotePortCount() int32`

GetInactiveRemotePortCount returns the InactiveRemotePortCount field if non-nil, zero value otherwise.

### GetInactiveRemotePortCountOk

`func (o *ProcessGroupDTO) GetInactiveRemotePortCountOk() (*int32, bool)`

GetInactiveRemotePortCountOk returns a tuple with the InactiveRemotePortCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveRemotePortCount

`func (o *ProcessGroupDTO) SetInactiveRemotePortCount(v int32)`

SetInactiveRemotePortCount sets InactiveRemotePortCount field to given value.

### HasInactiveRemotePortCount

`func (o *ProcessGroupDTO) HasInactiveRemotePortCount() bool`

HasInactiveRemotePortCount returns a boolean if a field has been set.

### GetUpToDateCount

`func (o *ProcessGroupDTO) GetUpToDateCount() int32`

GetUpToDateCount returns the UpToDateCount field if non-nil, zero value otherwise.

### GetUpToDateCountOk

`func (o *ProcessGroupDTO) GetUpToDateCountOk() (*int32, bool)`

GetUpToDateCountOk returns a tuple with the UpToDateCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpToDateCount

`func (o *ProcessGroupDTO) SetUpToDateCount(v int32)`

SetUpToDateCount sets UpToDateCount field to given value.

### HasUpToDateCount

`func (o *ProcessGroupDTO) HasUpToDateCount() bool`

HasUpToDateCount returns a boolean if a field has been set.

### GetLocallyModifiedCount

`func (o *ProcessGroupDTO) GetLocallyModifiedCount() int32`

GetLocallyModifiedCount returns the LocallyModifiedCount field if non-nil, zero value otherwise.

### GetLocallyModifiedCountOk

`func (o *ProcessGroupDTO) GetLocallyModifiedCountOk() (*int32, bool)`

GetLocallyModifiedCountOk returns a tuple with the LocallyModifiedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocallyModifiedCount

`func (o *ProcessGroupDTO) SetLocallyModifiedCount(v int32)`

SetLocallyModifiedCount sets LocallyModifiedCount field to given value.

### HasLocallyModifiedCount

`func (o *ProcessGroupDTO) HasLocallyModifiedCount() bool`

HasLocallyModifiedCount returns a boolean if a field has been set.

### GetStaleCount

`func (o *ProcessGroupDTO) GetStaleCount() int32`

GetStaleCount returns the StaleCount field if non-nil, zero value otherwise.

### GetStaleCountOk

`func (o *ProcessGroupDTO) GetStaleCountOk() (*int32, bool)`

GetStaleCountOk returns a tuple with the StaleCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaleCount

`func (o *ProcessGroupDTO) SetStaleCount(v int32)`

SetStaleCount sets StaleCount field to given value.

### HasStaleCount

`func (o *ProcessGroupDTO) HasStaleCount() bool`

HasStaleCount returns a boolean if a field has been set.

### GetLocallyModifiedAndStaleCount

`func (o *ProcessGroupDTO) GetLocallyModifiedAndStaleCount() int32`

GetLocallyModifiedAndStaleCount returns the LocallyModifiedAndStaleCount field if non-nil, zero value otherwise.

### GetLocallyModifiedAndStaleCountOk

`func (o *ProcessGroupDTO) GetLocallyModifiedAndStaleCountOk() (*int32, bool)`

GetLocallyModifiedAndStaleCountOk returns a tuple with the LocallyModifiedAndStaleCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocallyModifiedAndStaleCount

`func (o *ProcessGroupDTO) SetLocallyModifiedAndStaleCount(v int32)`

SetLocallyModifiedAndStaleCount sets LocallyModifiedAndStaleCount field to given value.

### HasLocallyModifiedAndStaleCount

`func (o *ProcessGroupDTO) HasLocallyModifiedAndStaleCount() bool`

HasLocallyModifiedAndStaleCount returns a boolean if a field has been set.

### GetSyncFailureCount

`func (o *ProcessGroupDTO) GetSyncFailureCount() int32`

GetSyncFailureCount returns the SyncFailureCount field if non-nil, zero value otherwise.

### GetSyncFailureCountOk

`func (o *ProcessGroupDTO) GetSyncFailureCountOk() (*int32, bool)`

GetSyncFailureCountOk returns a tuple with the SyncFailureCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncFailureCount

`func (o *ProcessGroupDTO) SetSyncFailureCount(v int32)`

SetSyncFailureCount sets SyncFailureCount field to given value.

### HasSyncFailureCount

`func (o *ProcessGroupDTO) HasSyncFailureCount() bool`

HasSyncFailureCount returns a boolean if a field has been set.

### GetLocalInputPortCount

`func (o *ProcessGroupDTO) GetLocalInputPortCount() int32`

GetLocalInputPortCount returns the LocalInputPortCount field if non-nil, zero value otherwise.

### GetLocalInputPortCountOk

`func (o *ProcessGroupDTO) GetLocalInputPortCountOk() (*int32, bool)`

GetLocalInputPortCountOk returns a tuple with the LocalInputPortCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalInputPortCount

`func (o *ProcessGroupDTO) SetLocalInputPortCount(v int32)`

SetLocalInputPortCount sets LocalInputPortCount field to given value.

### HasLocalInputPortCount

`func (o *ProcessGroupDTO) HasLocalInputPortCount() bool`

HasLocalInputPortCount returns a boolean if a field has been set.

### GetLocalOutputPortCount

`func (o *ProcessGroupDTO) GetLocalOutputPortCount() int32`

GetLocalOutputPortCount returns the LocalOutputPortCount field if non-nil, zero value otherwise.

### GetLocalOutputPortCountOk

`func (o *ProcessGroupDTO) GetLocalOutputPortCountOk() (*int32, bool)`

GetLocalOutputPortCountOk returns a tuple with the LocalOutputPortCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalOutputPortCount

`func (o *ProcessGroupDTO) SetLocalOutputPortCount(v int32)`

SetLocalOutputPortCount sets LocalOutputPortCount field to given value.

### HasLocalOutputPortCount

`func (o *ProcessGroupDTO) HasLocalOutputPortCount() bool`

HasLocalOutputPortCount returns a boolean if a field has been set.

### GetPublicInputPortCount

`func (o *ProcessGroupDTO) GetPublicInputPortCount() int32`

GetPublicInputPortCount returns the PublicInputPortCount field if non-nil, zero value otherwise.

### GetPublicInputPortCountOk

`func (o *ProcessGroupDTO) GetPublicInputPortCountOk() (*int32, bool)`

GetPublicInputPortCountOk returns a tuple with the PublicInputPortCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicInputPortCount

`func (o *ProcessGroupDTO) SetPublicInputPortCount(v int32)`

SetPublicInputPortCount sets PublicInputPortCount field to given value.

### HasPublicInputPortCount

`func (o *ProcessGroupDTO) HasPublicInputPortCount() bool`

HasPublicInputPortCount returns a boolean if a field has been set.

### GetPublicOutputPortCount

`func (o *ProcessGroupDTO) GetPublicOutputPortCount() int32`

GetPublicOutputPortCount returns the PublicOutputPortCount field if non-nil, zero value otherwise.

### GetPublicOutputPortCountOk

`func (o *ProcessGroupDTO) GetPublicOutputPortCountOk() (*int32, bool)`

GetPublicOutputPortCountOk returns a tuple with the PublicOutputPortCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicOutputPortCount

`func (o *ProcessGroupDTO) SetPublicOutputPortCount(v int32)`

SetPublicOutputPortCount sets PublicOutputPortCount field to given value.

### HasPublicOutputPortCount

`func (o *ProcessGroupDTO) HasPublicOutputPortCount() bool`

HasPublicOutputPortCount returns a boolean if a field has been set.

### GetContents

`func (o *ProcessGroupDTO) GetContents() FlowSnippetDTO`

GetContents returns the Contents field if non-nil, zero value otherwise.

### GetContentsOk

`func (o *ProcessGroupDTO) GetContentsOk() (*FlowSnippetDTO, bool)`

GetContentsOk returns a tuple with the Contents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContents

`func (o *ProcessGroupDTO) SetContents(v FlowSnippetDTO)`

SetContents sets Contents field to given value.

### HasContents

`func (o *ProcessGroupDTO) HasContents() bool`

HasContents returns a boolean if a field has been set.

### GetInputPortCount

`func (o *ProcessGroupDTO) GetInputPortCount() int32`

GetInputPortCount returns the InputPortCount field if non-nil, zero value otherwise.

### GetInputPortCountOk

`func (o *ProcessGroupDTO) GetInputPortCountOk() (*int32, bool)`

GetInputPortCountOk returns a tuple with the InputPortCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputPortCount

`func (o *ProcessGroupDTO) SetInputPortCount(v int32)`

SetInputPortCount sets InputPortCount field to given value.

### HasInputPortCount

`func (o *ProcessGroupDTO) HasInputPortCount() bool`

HasInputPortCount returns a boolean if a field has been set.

### GetOutputPortCount

`func (o *ProcessGroupDTO) GetOutputPortCount() int32`

GetOutputPortCount returns the OutputPortCount field if non-nil, zero value otherwise.

### GetOutputPortCountOk

`func (o *ProcessGroupDTO) GetOutputPortCountOk() (*int32, bool)`

GetOutputPortCountOk returns a tuple with the OutputPortCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputPortCount

`func (o *ProcessGroupDTO) SetOutputPortCount(v int32)`

SetOutputPortCount sets OutputPortCount field to given value.

### HasOutputPortCount

`func (o *ProcessGroupDTO) HasOutputPortCount() bool`

HasOutputPortCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


