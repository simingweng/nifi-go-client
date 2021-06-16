# NodeProcessorStatusSnapshotDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodeId** | Pointer to **string** | The unique ID that identifies the node | [optional] 
**Address** | Pointer to **string** | The API address of the node | [optional] 
**ApiPort** | Pointer to **int32** | The API port used to communicate with the node | [optional] 
**StatusSnapshot** | Pointer to [**ProcessorStatusSnapshotDTO**](ProcessorStatusSnapshotDTO.md) |  | [optional] 

## Methods

### NewNodeProcessorStatusSnapshotDTO

`func NewNodeProcessorStatusSnapshotDTO() *NodeProcessorStatusSnapshotDTO`

NewNodeProcessorStatusSnapshotDTO instantiates a new NodeProcessorStatusSnapshotDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeProcessorStatusSnapshotDTOWithDefaults

`func NewNodeProcessorStatusSnapshotDTOWithDefaults() *NodeProcessorStatusSnapshotDTO`

NewNodeProcessorStatusSnapshotDTOWithDefaults instantiates a new NodeProcessorStatusSnapshotDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodeId

`func (o *NodeProcessorStatusSnapshotDTO) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *NodeProcessorStatusSnapshotDTO) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *NodeProcessorStatusSnapshotDTO) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *NodeProcessorStatusSnapshotDTO) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### GetAddress

`func (o *NodeProcessorStatusSnapshotDTO) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *NodeProcessorStatusSnapshotDTO) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *NodeProcessorStatusSnapshotDTO) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *NodeProcessorStatusSnapshotDTO) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetApiPort

`func (o *NodeProcessorStatusSnapshotDTO) GetApiPort() int32`

GetApiPort returns the ApiPort field if non-nil, zero value otherwise.

### GetApiPortOk

`func (o *NodeProcessorStatusSnapshotDTO) GetApiPortOk() (*int32, bool)`

GetApiPortOk returns a tuple with the ApiPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiPort

`func (o *NodeProcessorStatusSnapshotDTO) SetApiPort(v int32)`

SetApiPort sets ApiPort field to given value.

### HasApiPort

`func (o *NodeProcessorStatusSnapshotDTO) HasApiPort() bool`

HasApiPort returns a boolean if a field has been set.

### GetStatusSnapshot

`func (o *NodeProcessorStatusSnapshotDTO) GetStatusSnapshot() ProcessorStatusSnapshotDTO`

GetStatusSnapshot returns the StatusSnapshot field if non-nil, zero value otherwise.

### GetStatusSnapshotOk

`func (o *NodeProcessorStatusSnapshotDTO) GetStatusSnapshotOk() (*ProcessorStatusSnapshotDTO, bool)`

GetStatusSnapshotOk returns a tuple with the StatusSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusSnapshot

`func (o *NodeProcessorStatusSnapshotDTO) SetStatusSnapshot(v ProcessorStatusSnapshotDTO)`

SetStatusSnapshot sets StatusSnapshot field to given value.

### HasStatusSnapshot

`func (o *NodeProcessorStatusSnapshotDTO) HasStatusSnapshot() bool`

HasStatusSnapshot returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


