# NodeConnectionStatisticsSnapshotDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodeId** | Pointer to **string** | The unique ID that identifies the node | [optional] 
**Address** | Pointer to **string** | The API address of the node | [optional] 
**ApiPort** | Pointer to **int32** | The API port used to communicate with the node | [optional] 
**StatisticsSnapshot** | Pointer to [**ConnectionStatisticsSnapshotDTO**](ConnectionStatisticsSnapshotDTO.md) |  | [optional] 

## Methods

### NewNodeConnectionStatisticsSnapshotDTO

`func NewNodeConnectionStatisticsSnapshotDTO() *NodeConnectionStatisticsSnapshotDTO`

NewNodeConnectionStatisticsSnapshotDTO instantiates a new NodeConnectionStatisticsSnapshotDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeConnectionStatisticsSnapshotDTOWithDefaults

`func NewNodeConnectionStatisticsSnapshotDTOWithDefaults() *NodeConnectionStatisticsSnapshotDTO`

NewNodeConnectionStatisticsSnapshotDTOWithDefaults instantiates a new NodeConnectionStatisticsSnapshotDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodeId

`func (o *NodeConnectionStatisticsSnapshotDTO) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *NodeConnectionStatisticsSnapshotDTO) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *NodeConnectionStatisticsSnapshotDTO) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *NodeConnectionStatisticsSnapshotDTO) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### GetAddress

`func (o *NodeConnectionStatisticsSnapshotDTO) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *NodeConnectionStatisticsSnapshotDTO) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *NodeConnectionStatisticsSnapshotDTO) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *NodeConnectionStatisticsSnapshotDTO) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetApiPort

`func (o *NodeConnectionStatisticsSnapshotDTO) GetApiPort() int32`

GetApiPort returns the ApiPort field if non-nil, zero value otherwise.

### GetApiPortOk

`func (o *NodeConnectionStatisticsSnapshotDTO) GetApiPortOk() (*int32, bool)`

GetApiPortOk returns a tuple with the ApiPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiPort

`func (o *NodeConnectionStatisticsSnapshotDTO) SetApiPort(v int32)`

SetApiPort sets ApiPort field to given value.

### HasApiPort

`func (o *NodeConnectionStatisticsSnapshotDTO) HasApiPort() bool`

HasApiPort returns a boolean if a field has been set.

### GetStatisticsSnapshot

`func (o *NodeConnectionStatisticsSnapshotDTO) GetStatisticsSnapshot() ConnectionStatisticsSnapshotDTO`

GetStatisticsSnapshot returns the StatisticsSnapshot field if non-nil, zero value otherwise.

### GetStatisticsSnapshotOk

`func (o *NodeConnectionStatisticsSnapshotDTO) GetStatisticsSnapshotOk() (*ConnectionStatisticsSnapshotDTO, bool)`

GetStatisticsSnapshotOk returns a tuple with the StatisticsSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatisticsSnapshot

`func (o *NodeConnectionStatisticsSnapshotDTO) SetStatisticsSnapshot(v ConnectionStatisticsSnapshotDTO)`

SetStatisticsSnapshot sets StatisticsSnapshot field to given value.

### HasStatisticsSnapshot

`func (o *NodeConnectionStatisticsSnapshotDTO) HasStatisticsSnapshot() bool`

HasStatisticsSnapshot returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


