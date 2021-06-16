# VersionedProcessGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | Pointer to **string** | The component&#39;s unique identifier | [optional] 
**Name** | Pointer to **string** | The component&#39;s name | [optional] 
**Comments** | Pointer to **string** | The user-supplied comments for the component | [optional] 
**Position** | Pointer to [**Position**](Position.md) |  | [optional] 
**ProcessGroups** | Pointer to [**[]VersionedProcessGroup**](VersionedProcessGroup.md) | The child Process Groups | [optional] 
**RemoteProcessGroups** | Pointer to [**[]VersionedRemoteProcessGroup**](VersionedRemoteProcessGroup.md) | The Remote Process Groups | [optional] 
**Processors** | Pointer to [**[]VersionedProcessor**](VersionedProcessor.md) | The Processors | [optional] 
**InputPorts** | Pointer to [**[]VersionedPort**](VersionedPort.md) | The Input Ports | [optional] 
**OutputPorts** | Pointer to [**[]VersionedPort**](VersionedPort.md) | The Output Ports | [optional] 
**Connections** | Pointer to [**[]VersionedConnection**](VersionedConnection.md) | The Connections | [optional] 
**Labels** | Pointer to [**[]VersionedLabel**](VersionedLabel.md) | The Labels | [optional] 
**Funnels** | Pointer to [**[]VersionedFunnel**](VersionedFunnel.md) | The Funnels | [optional] 
**ControllerServices** | Pointer to [**[]VersionedControllerService**](VersionedControllerService.md) | The Controller Services | [optional] 
**VersionedFlowCoordinates** | Pointer to [**VersionedFlowCoordinates**](VersionedFlowCoordinates.md) |  | [optional] 
**Variables** | Pointer to **map[string]string** | The Variables in the Variable Registry for this Process Group (not including any ancestor or descendant Process Groups) | [optional] 
**ParameterContextName** | Pointer to **string** | The name of the parameter context used by this process group | [optional] 
**FlowFileConcurrency** | Pointer to **string** | The configured FlowFile Concurrency for the Process Group | [optional] 
**FlowFileOutboundPolicy** | Pointer to **string** | The FlowFile Outbound Policy for the Process Group | [optional] 
**ComponentType** | Pointer to **string** |  | [optional] 
**GroupIdentifier** | Pointer to **string** | The ID of the Process Group that this component belongs to | [optional] 

## Methods

### NewVersionedProcessGroup

`func NewVersionedProcessGroup() *VersionedProcessGroup`

NewVersionedProcessGroup instantiates a new VersionedProcessGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersionedProcessGroupWithDefaults

`func NewVersionedProcessGroupWithDefaults() *VersionedProcessGroup`

NewVersionedProcessGroupWithDefaults instantiates a new VersionedProcessGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifier

`func (o *VersionedProcessGroup) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *VersionedProcessGroup) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *VersionedProcessGroup) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *VersionedProcessGroup) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetName

`func (o *VersionedProcessGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VersionedProcessGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VersionedProcessGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VersionedProcessGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetComments

`func (o *VersionedProcessGroup) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *VersionedProcessGroup) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *VersionedProcessGroup) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *VersionedProcessGroup) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetPosition

`func (o *VersionedProcessGroup) GetPosition() Position`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *VersionedProcessGroup) GetPositionOk() (*Position, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *VersionedProcessGroup) SetPosition(v Position)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *VersionedProcessGroup) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetProcessGroups

`func (o *VersionedProcessGroup) GetProcessGroups() []VersionedProcessGroup`

GetProcessGroups returns the ProcessGroups field if non-nil, zero value otherwise.

### GetProcessGroupsOk

`func (o *VersionedProcessGroup) GetProcessGroupsOk() (*[]VersionedProcessGroup, bool)`

GetProcessGroupsOk returns a tuple with the ProcessGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessGroups

`func (o *VersionedProcessGroup) SetProcessGroups(v []VersionedProcessGroup)`

SetProcessGroups sets ProcessGroups field to given value.

### HasProcessGroups

`func (o *VersionedProcessGroup) HasProcessGroups() bool`

HasProcessGroups returns a boolean if a field has been set.

### GetRemoteProcessGroups

`func (o *VersionedProcessGroup) GetRemoteProcessGroups() []VersionedRemoteProcessGroup`

GetRemoteProcessGroups returns the RemoteProcessGroups field if non-nil, zero value otherwise.

### GetRemoteProcessGroupsOk

`func (o *VersionedProcessGroup) GetRemoteProcessGroupsOk() (*[]VersionedRemoteProcessGroup, bool)`

GetRemoteProcessGroupsOk returns a tuple with the RemoteProcessGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteProcessGroups

`func (o *VersionedProcessGroup) SetRemoteProcessGroups(v []VersionedRemoteProcessGroup)`

SetRemoteProcessGroups sets RemoteProcessGroups field to given value.

### HasRemoteProcessGroups

`func (o *VersionedProcessGroup) HasRemoteProcessGroups() bool`

HasRemoteProcessGroups returns a boolean if a field has been set.

### GetProcessors

`func (o *VersionedProcessGroup) GetProcessors() []VersionedProcessor`

GetProcessors returns the Processors field if non-nil, zero value otherwise.

### GetProcessorsOk

`func (o *VersionedProcessGroup) GetProcessorsOk() (*[]VersionedProcessor, bool)`

GetProcessorsOk returns a tuple with the Processors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessors

`func (o *VersionedProcessGroup) SetProcessors(v []VersionedProcessor)`

SetProcessors sets Processors field to given value.

### HasProcessors

`func (o *VersionedProcessGroup) HasProcessors() bool`

HasProcessors returns a boolean if a field has been set.

### GetInputPorts

`func (o *VersionedProcessGroup) GetInputPorts() []VersionedPort`

GetInputPorts returns the InputPorts field if non-nil, zero value otherwise.

### GetInputPortsOk

`func (o *VersionedProcessGroup) GetInputPortsOk() (*[]VersionedPort, bool)`

GetInputPortsOk returns a tuple with the InputPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputPorts

`func (o *VersionedProcessGroup) SetInputPorts(v []VersionedPort)`

SetInputPorts sets InputPorts field to given value.

### HasInputPorts

`func (o *VersionedProcessGroup) HasInputPorts() bool`

HasInputPorts returns a boolean if a field has been set.

### GetOutputPorts

`func (o *VersionedProcessGroup) GetOutputPorts() []VersionedPort`

GetOutputPorts returns the OutputPorts field if non-nil, zero value otherwise.

### GetOutputPortsOk

`func (o *VersionedProcessGroup) GetOutputPortsOk() (*[]VersionedPort, bool)`

GetOutputPortsOk returns a tuple with the OutputPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputPorts

`func (o *VersionedProcessGroup) SetOutputPorts(v []VersionedPort)`

SetOutputPorts sets OutputPorts field to given value.

### HasOutputPorts

`func (o *VersionedProcessGroup) HasOutputPorts() bool`

HasOutputPorts returns a boolean if a field has been set.

### GetConnections

`func (o *VersionedProcessGroup) GetConnections() []VersionedConnection`

GetConnections returns the Connections field if non-nil, zero value otherwise.

### GetConnectionsOk

`func (o *VersionedProcessGroup) GetConnectionsOk() (*[]VersionedConnection, bool)`

GetConnectionsOk returns a tuple with the Connections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnections

`func (o *VersionedProcessGroup) SetConnections(v []VersionedConnection)`

SetConnections sets Connections field to given value.

### HasConnections

`func (o *VersionedProcessGroup) HasConnections() bool`

HasConnections returns a boolean if a field has been set.

### GetLabels

`func (o *VersionedProcessGroup) GetLabels() []VersionedLabel`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *VersionedProcessGroup) GetLabelsOk() (*[]VersionedLabel, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *VersionedProcessGroup) SetLabels(v []VersionedLabel)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *VersionedProcessGroup) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetFunnels

`func (o *VersionedProcessGroup) GetFunnels() []VersionedFunnel`

GetFunnels returns the Funnels field if non-nil, zero value otherwise.

### GetFunnelsOk

`func (o *VersionedProcessGroup) GetFunnelsOk() (*[]VersionedFunnel, bool)`

GetFunnelsOk returns a tuple with the Funnels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunnels

`func (o *VersionedProcessGroup) SetFunnels(v []VersionedFunnel)`

SetFunnels sets Funnels field to given value.

### HasFunnels

`func (o *VersionedProcessGroup) HasFunnels() bool`

HasFunnels returns a boolean if a field has been set.

### GetControllerServices

`func (o *VersionedProcessGroup) GetControllerServices() []VersionedControllerService`

GetControllerServices returns the ControllerServices field if non-nil, zero value otherwise.

### GetControllerServicesOk

`func (o *VersionedProcessGroup) GetControllerServicesOk() (*[]VersionedControllerService, bool)`

GetControllerServicesOk returns a tuple with the ControllerServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerServices

`func (o *VersionedProcessGroup) SetControllerServices(v []VersionedControllerService)`

SetControllerServices sets ControllerServices field to given value.

### HasControllerServices

`func (o *VersionedProcessGroup) HasControllerServices() bool`

HasControllerServices returns a boolean if a field has been set.

### GetVersionedFlowCoordinates

`func (o *VersionedProcessGroup) GetVersionedFlowCoordinates() VersionedFlowCoordinates`

GetVersionedFlowCoordinates returns the VersionedFlowCoordinates field if non-nil, zero value otherwise.

### GetVersionedFlowCoordinatesOk

`func (o *VersionedProcessGroup) GetVersionedFlowCoordinatesOk() (*VersionedFlowCoordinates, bool)`

GetVersionedFlowCoordinatesOk returns a tuple with the VersionedFlowCoordinates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionedFlowCoordinates

`func (o *VersionedProcessGroup) SetVersionedFlowCoordinates(v VersionedFlowCoordinates)`

SetVersionedFlowCoordinates sets VersionedFlowCoordinates field to given value.

### HasVersionedFlowCoordinates

`func (o *VersionedProcessGroup) HasVersionedFlowCoordinates() bool`

HasVersionedFlowCoordinates returns a boolean if a field has been set.

### GetVariables

`func (o *VersionedProcessGroup) GetVariables() map[string]string`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *VersionedProcessGroup) GetVariablesOk() (*map[string]string, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *VersionedProcessGroup) SetVariables(v map[string]string)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *VersionedProcessGroup) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### GetParameterContextName

`func (o *VersionedProcessGroup) GetParameterContextName() string`

GetParameterContextName returns the ParameterContextName field if non-nil, zero value otherwise.

### GetParameterContextNameOk

`func (o *VersionedProcessGroup) GetParameterContextNameOk() (*string, bool)`

GetParameterContextNameOk returns a tuple with the ParameterContextName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameterContextName

`func (o *VersionedProcessGroup) SetParameterContextName(v string)`

SetParameterContextName sets ParameterContextName field to given value.

### HasParameterContextName

`func (o *VersionedProcessGroup) HasParameterContextName() bool`

HasParameterContextName returns a boolean if a field has been set.

### GetFlowFileConcurrency

`func (o *VersionedProcessGroup) GetFlowFileConcurrency() string`

GetFlowFileConcurrency returns the FlowFileConcurrency field if non-nil, zero value otherwise.

### GetFlowFileConcurrencyOk

`func (o *VersionedProcessGroup) GetFlowFileConcurrencyOk() (*string, bool)`

GetFlowFileConcurrencyOk returns a tuple with the FlowFileConcurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowFileConcurrency

`func (o *VersionedProcessGroup) SetFlowFileConcurrency(v string)`

SetFlowFileConcurrency sets FlowFileConcurrency field to given value.

### HasFlowFileConcurrency

`func (o *VersionedProcessGroup) HasFlowFileConcurrency() bool`

HasFlowFileConcurrency returns a boolean if a field has been set.

### GetFlowFileOutboundPolicy

`func (o *VersionedProcessGroup) GetFlowFileOutboundPolicy() string`

GetFlowFileOutboundPolicy returns the FlowFileOutboundPolicy field if non-nil, zero value otherwise.

### GetFlowFileOutboundPolicyOk

`func (o *VersionedProcessGroup) GetFlowFileOutboundPolicyOk() (*string, bool)`

GetFlowFileOutboundPolicyOk returns a tuple with the FlowFileOutboundPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowFileOutboundPolicy

`func (o *VersionedProcessGroup) SetFlowFileOutboundPolicy(v string)`

SetFlowFileOutboundPolicy sets FlowFileOutboundPolicy field to given value.

### HasFlowFileOutboundPolicy

`func (o *VersionedProcessGroup) HasFlowFileOutboundPolicy() bool`

HasFlowFileOutboundPolicy returns a boolean if a field has been set.

### GetComponentType

`func (o *VersionedProcessGroup) GetComponentType() string`

GetComponentType returns the ComponentType field if non-nil, zero value otherwise.

### GetComponentTypeOk

`func (o *VersionedProcessGroup) GetComponentTypeOk() (*string, bool)`

GetComponentTypeOk returns a tuple with the ComponentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentType

`func (o *VersionedProcessGroup) SetComponentType(v string)`

SetComponentType sets ComponentType field to given value.

### HasComponentType

`func (o *VersionedProcessGroup) HasComponentType() bool`

HasComponentType returns a boolean if a field has been set.

### GetGroupIdentifier

`func (o *VersionedProcessGroup) GetGroupIdentifier() string`

GetGroupIdentifier returns the GroupIdentifier field if non-nil, zero value otherwise.

### GetGroupIdentifierOk

`func (o *VersionedProcessGroup) GetGroupIdentifierOk() (*string, bool)`

GetGroupIdentifierOk returns a tuple with the GroupIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupIdentifier

`func (o *VersionedProcessGroup) SetGroupIdentifier(v string)`

SetGroupIdentifier sets GroupIdentifier field to given value.

### HasGroupIdentifier

`func (o *VersionedProcessGroup) HasGroupIdentifier() bool`

HasGroupIdentifier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


