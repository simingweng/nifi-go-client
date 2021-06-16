# NodeSystemDiagnosticsSnapshotDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodeId** | Pointer to **string** | The unique ID that identifies the node | [optional] 
**Address** | Pointer to **string** | The API address of the node | [optional] 
**ApiPort** | Pointer to **int32** | The API port used to communicate with the node | [optional] 
**Snapshot** | Pointer to [**SystemDiagnosticsSnapshotDTO**](SystemDiagnosticsSnapshotDTO.md) |  | [optional] 

## Methods

### NewNodeSystemDiagnosticsSnapshotDTO

`func NewNodeSystemDiagnosticsSnapshotDTO() *NodeSystemDiagnosticsSnapshotDTO`

NewNodeSystemDiagnosticsSnapshotDTO instantiates a new NodeSystemDiagnosticsSnapshotDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeSystemDiagnosticsSnapshotDTOWithDefaults

`func NewNodeSystemDiagnosticsSnapshotDTOWithDefaults() *NodeSystemDiagnosticsSnapshotDTO`

NewNodeSystemDiagnosticsSnapshotDTOWithDefaults instantiates a new NodeSystemDiagnosticsSnapshotDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodeId

`func (o *NodeSystemDiagnosticsSnapshotDTO) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *NodeSystemDiagnosticsSnapshotDTO) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *NodeSystemDiagnosticsSnapshotDTO) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *NodeSystemDiagnosticsSnapshotDTO) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### GetAddress

`func (o *NodeSystemDiagnosticsSnapshotDTO) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *NodeSystemDiagnosticsSnapshotDTO) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *NodeSystemDiagnosticsSnapshotDTO) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *NodeSystemDiagnosticsSnapshotDTO) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetApiPort

`func (o *NodeSystemDiagnosticsSnapshotDTO) GetApiPort() int32`

GetApiPort returns the ApiPort field if non-nil, zero value otherwise.

### GetApiPortOk

`func (o *NodeSystemDiagnosticsSnapshotDTO) GetApiPortOk() (*int32, bool)`

GetApiPortOk returns a tuple with the ApiPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiPort

`func (o *NodeSystemDiagnosticsSnapshotDTO) SetApiPort(v int32)`

SetApiPort sets ApiPort field to given value.

### HasApiPort

`func (o *NodeSystemDiagnosticsSnapshotDTO) HasApiPort() bool`

HasApiPort returns a boolean if a field has been set.

### GetSnapshot

`func (o *NodeSystemDiagnosticsSnapshotDTO) GetSnapshot() SystemDiagnosticsSnapshotDTO`

GetSnapshot returns the Snapshot field if non-nil, zero value otherwise.

### GetSnapshotOk

`func (o *NodeSystemDiagnosticsSnapshotDTO) GetSnapshotOk() (*SystemDiagnosticsSnapshotDTO, bool)`

GetSnapshotOk returns a tuple with the Snapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshot

`func (o *NodeSystemDiagnosticsSnapshotDTO) SetSnapshot(v SystemDiagnosticsSnapshotDTO)`

SetSnapshot sets Snapshot field to given value.

### HasSnapshot

`func (o *NodeSystemDiagnosticsSnapshotDTO) HasSnapshot() bool`

HasSnapshot returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


