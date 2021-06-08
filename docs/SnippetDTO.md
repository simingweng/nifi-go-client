# SnippetDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The id of the snippet. | [optional] 
**Uri** | Pointer to **string** | The URI of the snippet. | [optional] 
**ParentGroupId** | Pointer to **string** | The group id for the components in the snippet. | [optional] 
**ProcessGroups** | Pointer to [**map[string]RevisionDTO**](RevisionDTO.md) | The ids of the process groups in this snippet. These ids will be populated within each response. They can be specified when creating a snippet. However, once a snippet has been created its contents cannot be modified (these ids are ignored during update requests). | [optional] 
**RemoteProcessGroups** | Pointer to [**map[string]RevisionDTO**](RevisionDTO.md) | The ids of the remote process groups in this snippet. These ids will be populated within each response. They can be specified when creating a snippet. However, once a snippet has been created its contents cannot be modified (these ids are ignored during update requests). | [optional] 
**Processors** | Pointer to [**map[string]RevisionDTO**](RevisionDTO.md) | The ids of the processors in this snippet. These ids will be populated within each response. They can be specified when creating a snippet. However, once a snippet has been created its contents cannot be modified (these ids are ignored during update requests). | [optional] 
**InputPorts** | Pointer to [**map[string]RevisionDTO**](RevisionDTO.md) | The ids of the input ports in this snippet. These ids will be populated within each response. They can be specified when creating a snippet. However, once a snippet has been created its contents cannot be modified (these ids are ignored during update requests). | [optional] 
**OutputPorts** | Pointer to [**map[string]RevisionDTO**](RevisionDTO.md) | The ids of the output ports in this snippet. These ids will be populated within each response. They can be specified when creating a snippet. However, once a snippet has been created its contents cannot be modified (these ids are ignored during update requests). | [optional] 
**Connections** | Pointer to [**map[string]RevisionDTO**](RevisionDTO.md) | The ids of the connections in this snippet. These ids will be populated within each response. They can be specified when creating a snippet. However, once a snippet has been created its contents cannot be modified (these ids are ignored during update requests). | [optional] 
**Labels** | Pointer to [**map[string]RevisionDTO**](RevisionDTO.md) | The ids of the labels in this snippet. These ids will be populated within each response. They can be specified when creating a snippet. However, once a snippet has been created its contents cannot be modified (these ids are ignored during update requests). | [optional] 
**Funnels** | Pointer to [**map[string]RevisionDTO**](RevisionDTO.md) | The ids of the funnels in this snippet. These ids will be populated within each response. They can be specified when creating a snippet. However, once a snippet has been created its contents cannot be modified (these ids are ignored during update requests). | [optional] 

## Methods

### NewSnippetDTO

`func NewSnippetDTO() *SnippetDTO`

NewSnippetDTO instantiates a new SnippetDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnippetDTOWithDefaults

`func NewSnippetDTOWithDefaults() *SnippetDTO`

NewSnippetDTOWithDefaults instantiates a new SnippetDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SnippetDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SnippetDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SnippetDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SnippetDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUri

`func (o *SnippetDTO) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *SnippetDTO) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *SnippetDTO) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *SnippetDTO) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetParentGroupId

`func (o *SnippetDTO) GetParentGroupId() string`

GetParentGroupId returns the ParentGroupId field if non-nil, zero value otherwise.

### GetParentGroupIdOk

`func (o *SnippetDTO) GetParentGroupIdOk() (*string, bool)`

GetParentGroupIdOk returns a tuple with the ParentGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentGroupId

`func (o *SnippetDTO) SetParentGroupId(v string)`

SetParentGroupId sets ParentGroupId field to given value.

### HasParentGroupId

`func (o *SnippetDTO) HasParentGroupId() bool`

HasParentGroupId returns a boolean if a field has been set.

### GetProcessGroups

`func (o *SnippetDTO) GetProcessGroups() map[string]RevisionDTO`

GetProcessGroups returns the ProcessGroups field if non-nil, zero value otherwise.

### GetProcessGroupsOk

`func (o *SnippetDTO) GetProcessGroupsOk() (*map[string]RevisionDTO, bool)`

GetProcessGroupsOk returns a tuple with the ProcessGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessGroups

`func (o *SnippetDTO) SetProcessGroups(v map[string]RevisionDTO)`

SetProcessGroups sets ProcessGroups field to given value.

### HasProcessGroups

`func (o *SnippetDTO) HasProcessGroups() bool`

HasProcessGroups returns a boolean if a field has been set.

### GetRemoteProcessGroups

`func (o *SnippetDTO) GetRemoteProcessGroups() map[string]RevisionDTO`

GetRemoteProcessGroups returns the RemoteProcessGroups field if non-nil, zero value otherwise.

### GetRemoteProcessGroupsOk

`func (o *SnippetDTO) GetRemoteProcessGroupsOk() (*map[string]RevisionDTO, bool)`

GetRemoteProcessGroupsOk returns a tuple with the RemoteProcessGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteProcessGroups

`func (o *SnippetDTO) SetRemoteProcessGroups(v map[string]RevisionDTO)`

SetRemoteProcessGroups sets RemoteProcessGroups field to given value.

### HasRemoteProcessGroups

`func (o *SnippetDTO) HasRemoteProcessGroups() bool`

HasRemoteProcessGroups returns a boolean if a field has been set.

### GetProcessors

`func (o *SnippetDTO) GetProcessors() map[string]RevisionDTO`

GetProcessors returns the Processors field if non-nil, zero value otherwise.

### GetProcessorsOk

`func (o *SnippetDTO) GetProcessorsOk() (*map[string]RevisionDTO, bool)`

GetProcessorsOk returns a tuple with the Processors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessors

`func (o *SnippetDTO) SetProcessors(v map[string]RevisionDTO)`

SetProcessors sets Processors field to given value.

### HasProcessors

`func (o *SnippetDTO) HasProcessors() bool`

HasProcessors returns a boolean if a field has been set.

### GetInputPorts

`func (o *SnippetDTO) GetInputPorts() map[string]RevisionDTO`

GetInputPorts returns the InputPorts field if non-nil, zero value otherwise.

### GetInputPortsOk

`func (o *SnippetDTO) GetInputPortsOk() (*map[string]RevisionDTO, bool)`

GetInputPortsOk returns a tuple with the InputPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputPorts

`func (o *SnippetDTO) SetInputPorts(v map[string]RevisionDTO)`

SetInputPorts sets InputPorts field to given value.

### HasInputPorts

`func (o *SnippetDTO) HasInputPorts() bool`

HasInputPorts returns a boolean if a field has been set.

### GetOutputPorts

`func (o *SnippetDTO) GetOutputPorts() map[string]RevisionDTO`

GetOutputPorts returns the OutputPorts field if non-nil, zero value otherwise.

### GetOutputPortsOk

`func (o *SnippetDTO) GetOutputPortsOk() (*map[string]RevisionDTO, bool)`

GetOutputPortsOk returns a tuple with the OutputPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputPorts

`func (o *SnippetDTO) SetOutputPorts(v map[string]RevisionDTO)`

SetOutputPorts sets OutputPorts field to given value.

### HasOutputPorts

`func (o *SnippetDTO) HasOutputPorts() bool`

HasOutputPorts returns a boolean if a field has been set.

### GetConnections

`func (o *SnippetDTO) GetConnections() map[string]RevisionDTO`

GetConnections returns the Connections field if non-nil, zero value otherwise.

### GetConnectionsOk

`func (o *SnippetDTO) GetConnectionsOk() (*map[string]RevisionDTO, bool)`

GetConnectionsOk returns a tuple with the Connections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnections

`func (o *SnippetDTO) SetConnections(v map[string]RevisionDTO)`

SetConnections sets Connections field to given value.

### HasConnections

`func (o *SnippetDTO) HasConnections() bool`

HasConnections returns a boolean if a field has been set.

### GetLabels

`func (o *SnippetDTO) GetLabels() map[string]RevisionDTO`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *SnippetDTO) GetLabelsOk() (*map[string]RevisionDTO, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *SnippetDTO) SetLabels(v map[string]RevisionDTO)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *SnippetDTO) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetFunnels

`func (o *SnippetDTO) GetFunnels() map[string]RevisionDTO`

GetFunnels returns the Funnels field if non-nil, zero value otherwise.

### GetFunnelsOk

`func (o *SnippetDTO) GetFunnelsOk() (*map[string]RevisionDTO, bool)`

GetFunnelsOk returns a tuple with the Funnels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunnels

`func (o *SnippetDTO) SetFunnels(v map[string]RevisionDTO)`

SetFunnels sets Funnels field to given value.

### HasFunnels

`func (o *SnippetDTO) HasFunnels() bool`

HasFunnels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


