# FlowSnippetDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProcessGroups** | Pointer to [**[]ProcessGroupDTO**](ProcessGroupDTO.md) | The process groups in this flow snippet. | [optional] 
**RemoteProcessGroups** | Pointer to [**[]RemoteProcessGroupDTO**](RemoteProcessGroupDTO.md) | The remote process groups in this flow snippet. | [optional] 
**Processors** | Pointer to [**[]ProcessorDTO**](ProcessorDTO.md) | The processors in this flow snippet. | [optional] 
**InputPorts** | Pointer to [**[]PortDTO**](PortDTO.md) | The input ports in this flow snippet. | [optional] 
**OutputPorts** | Pointer to [**[]PortDTO**](PortDTO.md) | The output ports in this flow snippet. | [optional] 
**Connections** | Pointer to [**[]ConnectionDTO**](ConnectionDTO.md) | The connections in this flow snippet. | [optional] 
**Labels** | Pointer to [**[]LabelDTO**](LabelDTO.md) | The labels in this flow snippet. | [optional] 
**Funnels** | Pointer to [**[]FunnelDTO**](FunnelDTO.md) | The funnels in this flow snippet. | [optional] 
**ControllerServices** | Pointer to [**[]ControllerServiceDTO**](ControllerServiceDTO.md) | The controller services in this flow snippet. | [optional] 

## Methods

### NewFlowSnippetDTO

`func NewFlowSnippetDTO() *FlowSnippetDTO`

NewFlowSnippetDTO instantiates a new FlowSnippetDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlowSnippetDTOWithDefaults

`func NewFlowSnippetDTOWithDefaults() *FlowSnippetDTO`

NewFlowSnippetDTOWithDefaults instantiates a new FlowSnippetDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProcessGroups

`func (o *FlowSnippetDTO) GetProcessGroups() []ProcessGroupDTO`

GetProcessGroups returns the ProcessGroups field if non-nil, zero value otherwise.

### GetProcessGroupsOk

`func (o *FlowSnippetDTO) GetProcessGroupsOk() (*[]ProcessGroupDTO, bool)`

GetProcessGroupsOk returns a tuple with the ProcessGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessGroups

`func (o *FlowSnippetDTO) SetProcessGroups(v []ProcessGroupDTO)`

SetProcessGroups sets ProcessGroups field to given value.

### HasProcessGroups

`func (o *FlowSnippetDTO) HasProcessGroups() bool`

HasProcessGroups returns a boolean if a field has been set.

### GetRemoteProcessGroups

`func (o *FlowSnippetDTO) GetRemoteProcessGroups() []RemoteProcessGroupDTO`

GetRemoteProcessGroups returns the RemoteProcessGroups field if non-nil, zero value otherwise.

### GetRemoteProcessGroupsOk

`func (o *FlowSnippetDTO) GetRemoteProcessGroupsOk() (*[]RemoteProcessGroupDTO, bool)`

GetRemoteProcessGroupsOk returns a tuple with the RemoteProcessGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteProcessGroups

`func (o *FlowSnippetDTO) SetRemoteProcessGroups(v []RemoteProcessGroupDTO)`

SetRemoteProcessGroups sets RemoteProcessGroups field to given value.

### HasRemoteProcessGroups

`func (o *FlowSnippetDTO) HasRemoteProcessGroups() bool`

HasRemoteProcessGroups returns a boolean if a field has been set.

### GetProcessors

`func (o *FlowSnippetDTO) GetProcessors() []ProcessorDTO`

GetProcessors returns the Processors field if non-nil, zero value otherwise.

### GetProcessorsOk

`func (o *FlowSnippetDTO) GetProcessorsOk() (*[]ProcessorDTO, bool)`

GetProcessorsOk returns a tuple with the Processors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessors

`func (o *FlowSnippetDTO) SetProcessors(v []ProcessorDTO)`

SetProcessors sets Processors field to given value.

### HasProcessors

`func (o *FlowSnippetDTO) HasProcessors() bool`

HasProcessors returns a boolean if a field has been set.

### GetInputPorts

`func (o *FlowSnippetDTO) GetInputPorts() []PortDTO`

GetInputPorts returns the InputPorts field if non-nil, zero value otherwise.

### GetInputPortsOk

`func (o *FlowSnippetDTO) GetInputPortsOk() (*[]PortDTO, bool)`

GetInputPortsOk returns a tuple with the InputPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputPorts

`func (o *FlowSnippetDTO) SetInputPorts(v []PortDTO)`

SetInputPorts sets InputPorts field to given value.

### HasInputPorts

`func (o *FlowSnippetDTO) HasInputPorts() bool`

HasInputPorts returns a boolean if a field has been set.

### GetOutputPorts

`func (o *FlowSnippetDTO) GetOutputPorts() []PortDTO`

GetOutputPorts returns the OutputPorts field if non-nil, zero value otherwise.

### GetOutputPortsOk

`func (o *FlowSnippetDTO) GetOutputPortsOk() (*[]PortDTO, bool)`

GetOutputPortsOk returns a tuple with the OutputPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputPorts

`func (o *FlowSnippetDTO) SetOutputPorts(v []PortDTO)`

SetOutputPorts sets OutputPorts field to given value.

### HasOutputPorts

`func (o *FlowSnippetDTO) HasOutputPorts() bool`

HasOutputPorts returns a boolean if a field has been set.

### GetConnections

`func (o *FlowSnippetDTO) GetConnections() []ConnectionDTO`

GetConnections returns the Connections field if non-nil, zero value otherwise.

### GetConnectionsOk

`func (o *FlowSnippetDTO) GetConnectionsOk() (*[]ConnectionDTO, bool)`

GetConnectionsOk returns a tuple with the Connections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnections

`func (o *FlowSnippetDTO) SetConnections(v []ConnectionDTO)`

SetConnections sets Connections field to given value.

### HasConnections

`func (o *FlowSnippetDTO) HasConnections() bool`

HasConnections returns a boolean if a field has been set.

### GetLabels

`func (o *FlowSnippetDTO) GetLabels() []LabelDTO`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *FlowSnippetDTO) GetLabelsOk() (*[]LabelDTO, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *FlowSnippetDTO) SetLabels(v []LabelDTO)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *FlowSnippetDTO) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetFunnels

`func (o *FlowSnippetDTO) GetFunnels() []FunnelDTO`

GetFunnels returns the Funnels field if non-nil, zero value otherwise.

### GetFunnelsOk

`func (o *FlowSnippetDTO) GetFunnelsOk() (*[]FunnelDTO, bool)`

GetFunnelsOk returns a tuple with the Funnels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunnels

`func (o *FlowSnippetDTO) SetFunnels(v []FunnelDTO)`

SetFunnels sets Funnels field to given value.

### HasFunnels

`func (o *FlowSnippetDTO) HasFunnels() bool`

HasFunnels returns a boolean if a field has been set.

### GetControllerServices

`func (o *FlowSnippetDTO) GetControllerServices() []ControllerServiceDTO`

GetControllerServices returns the ControllerServices field if non-nil, zero value otherwise.

### GetControllerServicesOk

`func (o *FlowSnippetDTO) GetControllerServicesOk() (*[]ControllerServiceDTO, bool)`

GetControllerServicesOk returns a tuple with the ControllerServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerServices

`func (o *FlowSnippetDTO) SetControllerServices(v []ControllerServiceDTO)`

SetControllerServices sets ControllerServices field to given value.

### HasControllerServices

`func (o *FlowSnippetDTO) HasControllerServices() bool`

HasControllerServices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


