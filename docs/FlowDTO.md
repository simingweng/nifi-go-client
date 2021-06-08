# FlowDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProcessGroups** | Pointer to [**[]ProcessGroupEntity**](ProcessGroupEntity.md) | The process groups in this flow. | [optional] 
**RemoteProcessGroups** | Pointer to [**[]RemoteProcessGroupEntity**](RemoteProcessGroupEntity.md) | The remote process groups in this flow. | [optional] 
**Processors** | Pointer to [**[]ProcessorEntity**](ProcessorEntity.md) | The processors in this flow. | [optional] 
**InputPorts** | Pointer to [**[]PortEntity**](PortEntity.md) | The input ports in this flow. | [optional] 
**OutputPorts** | Pointer to [**[]PortEntity**](PortEntity.md) | The output ports in this flow. | [optional] 
**Connections** | Pointer to [**[]ConnectionEntity**](ConnectionEntity.md) | The connections in this flow. | [optional] 
**Labels** | Pointer to [**[]LabelEntity**](LabelEntity.md) | The labels in this flow. | [optional] 
**Funnels** | Pointer to [**[]FunnelEntity**](FunnelEntity.md) | The funnels in this flow. | [optional] 

## Methods

### NewFlowDTO

`func NewFlowDTO() *FlowDTO`

NewFlowDTO instantiates a new FlowDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlowDTOWithDefaults

`func NewFlowDTOWithDefaults() *FlowDTO`

NewFlowDTOWithDefaults instantiates a new FlowDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProcessGroups

`func (o *FlowDTO) GetProcessGroups() []ProcessGroupEntity`

GetProcessGroups returns the ProcessGroups field if non-nil, zero value otherwise.

### GetProcessGroupsOk

`func (o *FlowDTO) GetProcessGroupsOk() (*[]ProcessGroupEntity, bool)`

GetProcessGroupsOk returns a tuple with the ProcessGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessGroups

`func (o *FlowDTO) SetProcessGroups(v []ProcessGroupEntity)`

SetProcessGroups sets ProcessGroups field to given value.

### HasProcessGroups

`func (o *FlowDTO) HasProcessGroups() bool`

HasProcessGroups returns a boolean if a field has been set.

### GetRemoteProcessGroups

`func (o *FlowDTO) GetRemoteProcessGroups() []RemoteProcessGroupEntity`

GetRemoteProcessGroups returns the RemoteProcessGroups field if non-nil, zero value otherwise.

### GetRemoteProcessGroupsOk

`func (o *FlowDTO) GetRemoteProcessGroupsOk() (*[]RemoteProcessGroupEntity, bool)`

GetRemoteProcessGroupsOk returns a tuple with the RemoteProcessGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteProcessGroups

`func (o *FlowDTO) SetRemoteProcessGroups(v []RemoteProcessGroupEntity)`

SetRemoteProcessGroups sets RemoteProcessGroups field to given value.

### HasRemoteProcessGroups

`func (o *FlowDTO) HasRemoteProcessGroups() bool`

HasRemoteProcessGroups returns a boolean if a field has been set.

### GetProcessors

`func (o *FlowDTO) GetProcessors() []ProcessorEntity`

GetProcessors returns the Processors field if non-nil, zero value otherwise.

### GetProcessorsOk

`func (o *FlowDTO) GetProcessorsOk() (*[]ProcessorEntity, bool)`

GetProcessorsOk returns a tuple with the Processors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessors

`func (o *FlowDTO) SetProcessors(v []ProcessorEntity)`

SetProcessors sets Processors field to given value.

### HasProcessors

`func (o *FlowDTO) HasProcessors() bool`

HasProcessors returns a boolean if a field has been set.

### GetInputPorts

`func (o *FlowDTO) GetInputPorts() []PortEntity`

GetInputPorts returns the InputPorts field if non-nil, zero value otherwise.

### GetInputPortsOk

`func (o *FlowDTO) GetInputPortsOk() (*[]PortEntity, bool)`

GetInputPortsOk returns a tuple with the InputPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputPorts

`func (o *FlowDTO) SetInputPorts(v []PortEntity)`

SetInputPorts sets InputPorts field to given value.

### HasInputPorts

`func (o *FlowDTO) HasInputPorts() bool`

HasInputPorts returns a boolean if a field has been set.

### GetOutputPorts

`func (o *FlowDTO) GetOutputPorts() []PortEntity`

GetOutputPorts returns the OutputPorts field if non-nil, zero value otherwise.

### GetOutputPortsOk

`func (o *FlowDTO) GetOutputPortsOk() (*[]PortEntity, bool)`

GetOutputPortsOk returns a tuple with the OutputPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputPorts

`func (o *FlowDTO) SetOutputPorts(v []PortEntity)`

SetOutputPorts sets OutputPorts field to given value.

### HasOutputPorts

`func (o *FlowDTO) HasOutputPorts() bool`

HasOutputPorts returns a boolean if a field has been set.

### GetConnections

`func (o *FlowDTO) GetConnections() []ConnectionEntity`

GetConnections returns the Connections field if non-nil, zero value otherwise.

### GetConnectionsOk

`func (o *FlowDTO) GetConnectionsOk() (*[]ConnectionEntity, bool)`

GetConnectionsOk returns a tuple with the Connections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnections

`func (o *FlowDTO) SetConnections(v []ConnectionEntity)`

SetConnections sets Connections field to given value.

### HasConnections

`func (o *FlowDTO) HasConnections() bool`

HasConnections returns a boolean if a field has been set.

### GetLabels

`func (o *FlowDTO) GetLabels() []LabelEntity`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *FlowDTO) GetLabelsOk() (*[]LabelEntity, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *FlowDTO) SetLabels(v []LabelEntity)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *FlowDTO) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetFunnels

`func (o *FlowDTO) GetFunnels() []FunnelEntity`

GetFunnels returns the Funnels field if non-nil, zero value otherwise.

### GetFunnelsOk

`func (o *FlowDTO) GetFunnelsOk() (*[]FunnelEntity, bool)`

GetFunnelsOk returns a tuple with the Funnels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunnels

`func (o *FlowDTO) SetFunnels(v []FunnelEntity)`

SetFunnels sets Funnels field to given value.

### HasFunnels

`func (o *FlowDTO) HasFunnels() bool`

HasFunnels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


