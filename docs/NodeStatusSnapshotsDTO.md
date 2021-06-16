# NodeStatusSnapshotsDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodeId** | Pointer to **string** | The id of the node. | [optional] 
**Address** | Pointer to **string** | The node&#39;s host/ip address. | [optional] 
**ApiPort** | Pointer to **int32** | The port the node is listening for API requests. | [optional] 
**StatusSnapshots** | Pointer to [**[]StatusSnapshotDTO**](StatusSnapshotDTO.md) | A list of StatusSnapshotDTO objects that provide the actual metric values for the component for this node. | [optional] 

## Methods

### NewNodeStatusSnapshotsDTO

`func NewNodeStatusSnapshotsDTO() *NodeStatusSnapshotsDTO`

NewNodeStatusSnapshotsDTO instantiates a new NodeStatusSnapshotsDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeStatusSnapshotsDTOWithDefaults

`func NewNodeStatusSnapshotsDTOWithDefaults() *NodeStatusSnapshotsDTO`

NewNodeStatusSnapshotsDTOWithDefaults instantiates a new NodeStatusSnapshotsDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodeId

`func (o *NodeStatusSnapshotsDTO) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *NodeStatusSnapshotsDTO) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *NodeStatusSnapshotsDTO) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *NodeStatusSnapshotsDTO) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### GetAddress

`func (o *NodeStatusSnapshotsDTO) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *NodeStatusSnapshotsDTO) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *NodeStatusSnapshotsDTO) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *NodeStatusSnapshotsDTO) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetApiPort

`func (o *NodeStatusSnapshotsDTO) GetApiPort() int32`

GetApiPort returns the ApiPort field if non-nil, zero value otherwise.

### GetApiPortOk

`func (o *NodeStatusSnapshotsDTO) GetApiPortOk() (*int32, bool)`

GetApiPortOk returns a tuple with the ApiPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiPort

`func (o *NodeStatusSnapshotsDTO) SetApiPort(v int32)`

SetApiPort sets ApiPort field to given value.

### HasApiPort

`func (o *NodeStatusSnapshotsDTO) HasApiPort() bool`

HasApiPort returns a boolean if a field has been set.

### GetStatusSnapshots

`func (o *NodeStatusSnapshotsDTO) GetStatusSnapshots() []StatusSnapshotDTO`

GetStatusSnapshots returns the StatusSnapshots field if non-nil, zero value otherwise.

### GetStatusSnapshotsOk

`func (o *NodeStatusSnapshotsDTO) GetStatusSnapshotsOk() (*[]StatusSnapshotDTO, bool)`

GetStatusSnapshotsOk returns a tuple with the StatusSnapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusSnapshots

`func (o *NodeStatusSnapshotsDTO) SetStatusSnapshots(v []StatusSnapshotDTO)`

SetStatusSnapshots sets StatusSnapshots field to given value.

### HasStatusSnapshots

`func (o *NodeStatusSnapshotsDTO) HasStatusSnapshots() bool`

HasStatusSnapshots returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


