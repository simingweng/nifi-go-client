# RemoteProcessGroupContentsDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InputPorts** | Pointer to [**[]RemoteProcessGroupPortDTO**](RemoteProcessGroupPortDTO.md) | The input ports to which data can be sent. | [optional] 
**OutputPorts** | Pointer to [**[]RemoteProcessGroupPortDTO**](RemoteProcessGroupPortDTO.md) | The output ports from which data can be retrieved. | [optional] 

## Methods

### NewRemoteProcessGroupContentsDTO

`func NewRemoteProcessGroupContentsDTO() *RemoteProcessGroupContentsDTO`

NewRemoteProcessGroupContentsDTO instantiates a new RemoteProcessGroupContentsDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoteProcessGroupContentsDTOWithDefaults

`func NewRemoteProcessGroupContentsDTOWithDefaults() *RemoteProcessGroupContentsDTO`

NewRemoteProcessGroupContentsDTOWithDefaults instantiates a new RemoteProcessGroupContentsDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInputPorts

`func (o *RemoteProcessGroupContentsDTO) GetInputPorts() []RemoteProcessGroupPortDTO`

GetInputPorts returns the InputPorts field if non-nil, zero value otherwise.

### GetInputPortsOk

`func (o *RemoteProcessGroupContentsDTO) GetInputPortsOk() (*[]RemoteProcessGroupPortDTO, bool)`

GetInputPortsOk returns a tuple with the InputPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputPorts

`func (o *RemoteProcessGroupContentsDTO) SetInputPorts(v []RemoteProcessGroupPortDTO)`

SetInputPorts sets InputPorts field to given value.

### HasInputPorts

`func (o *RemoteProcessGroupContentsDTO) HasInputPorts() bool`

HasInputPorts returns a boolean if a field has been set.

### GetOutputPorts

`func (o *RemoteProcessGroupContentsDTO) GetOutputPorts() []RemoteProcessGroupPortDTO`

GetOutputPorts returns the OutputPorts field if non-nil, zero value otherwise.

### GetOutputPortsOk

`func (o *RemoteProcessGroupContentsDTO) GetOutputPortsOk() (*[]RemoteProcessGroupPortDTO, bool)`

GetOutputPortsOk returns a tuple with the OutputPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputPorts

`func (o *RemoteProcessGroupContentsDTO) SetOutputPorts(v []RemoteProcessGroupPortDTO)`

SetOutputPorts sets OutputPorts field to given value.

### HasOutputPorts

`func (o *RemoteProcessGroupContentsDTO) HasOutputPorts() bool`

HasOutputPorts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


