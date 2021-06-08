# NodeProcessGroupStatusSnapshotDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodeId** | Pointer to **string** | The unique ID that identifies the node | [optional] 
**Address** | Pointer to **string** | The API address of the node | [optional] 
**ApiPort** | Pointer to **int32** | The API port used to communicate with the node | [optional] 
**StatusSnapshot** | Pointer to [**ProcessGroupStatusSnapshotDTO**](ProcessGroupStatusSnapshotDTO.md) |  | [optional] 

## Methods

### NewNodeProcessGroupStatusSnapshotDTO

`func NewNodeProcessGroupStatusSnapshotDTO() *NodeProcessGroupStatusSnapshotDTO`

NewNodeProcessGroupStatusSnapshotDTO instantiates a new NodeProcessGroupStatusSnapshotDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeProcessGroupStatusSnapshotDTOWithDefaults

`func NewNodeProcessGroupStatusSnapshotDTOWithDefaults() *NodeProcessGroupStatusSnapshotDTO`

NewNodeProcessGroupStatusSnapshotDTOWithDefaults instantiates a new NodeProcessGroupStatusSnapshotDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodeId

`func (o *NodeProcessGroupStatusSnapshotDTO) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *NodeProcessGroupStatusSnapshotDTO) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *NodeProcessGroupStatusSnapshotDTO) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *NodeProcessGroupStatusSnapshotDTO) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### GetAddress

`func (o *NodeProcessGroupStatusSnapshotDTO) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *NodeProcessGroupStatusSnapshotDTO) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *NodeProcessGroupStatusSnapshotDTO) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *NodeProcessGroupStatusSnapshotDTO) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetApiPort

`func (o *NodeProcessGroupStatusSnapshotDTO) GetApiPort() int32`

GetApiPort returns the ApiPort field if non-nil, zero value otherwise.

### GetApiPortOk

`func (o *NodeProcessGroupStatusSnapshotDTO) GetApiPortOk() (*int32, bool)`

GetApiPortOk returns a tuple with the ApiPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiPort

`func (o *NodeProcessGroupStatusSnapshotDTO) SetApiPort(v int32)`

SetApiPort sets ApiPort field to given value.

### HasApiPort

`func (o *NodeProcessGroupStatusSnapshotDTO) HasApiPort() bool`

HasApiPort returns a boolean if a field has been set.

### GetStatusSnapshot

`func (o *NodeProcessGroupStatusSnapshotDTO) GetStatusSnapshot() ProcessGroupStatusSnapshotDTO`

GetStatusSnapshot returns the StatusSnapshot field if non-nil, zero value otherwise.

### GetStatusSnapshotOk

`func (o *NodeProcessGroupStatusSnapshotDTO) GetStatusSnapshotOk() (*ProcessGroupStatusSnapshotDTO, bool)`

GetStatusSnapshotOk returns a tuple with the StatusSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusSnapshot

`func (o *NodeProcessGroupStatusSnapshotDTO) SetStatusSnapshot(v ProcessGroupStatusSnapshotDTO)`

SetStatusSnapshot sets StatusSnapshot field to given value.

### HasStatusSnapshot

`func (o *NodeProcessGroupStatusSnapshotDTO) HasStatusSnapshot() bool`

HasStatusSnapshot returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


