# NodeDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodeId** | Pointer to **string** | The id of the node. | [optional] [readonly] 
**Address** | Pointer to **string** | The node&#39;s host/ip address. | [optional] [readonly] 
**ApiPort** | Pointer to **int32** | The port the node is listening for API requests. | [optional] [readonly] 
**Status** | Pointer to **string** | The node&#39;s status. | [optional] 
**Heartbeat** | Pointer to **string** | the time of the nodes&#39;s last heartbeat. | [optional] [readonly] 
**ConnectionRequested** | Pointer to **string** | The time of the node&#39;s last connection request. | [optional] [readonly] 
**Roles** | Pointer to **[]string** | The roles of this node. | [optional] [readonly] 
**ActiveThreadCount** | Pointer to **int32** | The active threads for the NiFi on the node. | [optional] [readonly] 
**Queued** | Pointer to **string** | The queue the NiFi on the node. | [optional] [readonly] 
**Events** | Pointer to [**[]NodeEventDTO**](NodeEventDTO.md) | The node&#39;s events. | [optional] [readonly] 
**NodeStartTime** | Pointer to **string** | The time at which this Node was last refreshed. | [optional] [readonly] 

## Methods

### NewNodeDTO

`func NewNodeDTO() *NodeDTO`

NewNodeDTO instantiates a new NodeDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeDTOWithDefaults

`func NewNodeDTOWithDefaults() *NodeDTO`

NewNodeDTOWithDefaults instantiates a new NodeDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodeId

`func (o *NodeDTO) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *NodeDTO) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *NodeDTO) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *NodeDTO) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### GetAddress

`func (o *NodeDTO) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *NodeDTO) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *NodeDTO) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *NodeDTO) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetApiPort

`func (o *NodeDTO) GetApiPort() int32`

GetApiPort returns the ApiPort field if non-nil, zero value otherwise.

### GetApiPortOk

`func (o *NodeDTO) GetApiPortOk() (*int32, bool)`

GetApiPortOk returns a tuple with the ApiPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiPort

`func (o *NodeDTO) SetApiPort(v int32)`

SetApiPort sets ApiPort field to given value.

### HasApiPort

`func (o *NodeDTO) HasApiPort() bool`

HasApiPort returns a boolean if a field has been set.

### GetStatus

`func (o *NodeDTO) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NodeDTO) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NodeDTO) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NodeDTO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetHeartbeat

`func (o *NodeDTO) GetHeartbeat() string`

GetHeartbeat returns the Heartbeat field if non-nil, zero value otherwise.

### GetHeartbeatOk

`func (o *NodeDTO) GetHeartbeatOk() (*string, bool)`

GetHeartbeatOk returns a tuple with the Heartbeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeartbeat

`func (o *NodeDTO) SetHeartbeat(v string)`

SetHeartbeat sets Heartbeat field to given value.

### HasHeartbeat

`func (o *NodeDTO) HasHeartbeat() bool`

HasHeartbeat returns a boolean if a field has been set.

### GetConnectionRequested

`func (o *NodeDTO) GetConnectionRequested() string`

GetConnectionRequested returns the ConnectionRequested field if non-nil, zero value otherwise.

### GetConnectionRequestedOk

`func (o *NodeDTO) GetConnectionRequestedOk() (*string, bool)`

GetConnectionRequestedOk returns a tuple with the ConnectionRequested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionRequested

`func (o *NodeDTO) SetConnectionRequested(v string)`

SetConnectionRequested sets ConnectionRequested field to given value.

### HasConnectionRequested

`func (o *NodeDTO) HasConnectionRequested() bool`

HasConnectionRequested returns a boolean if a field has been set.

### GetRoles

`func (o *NodeDTO) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *NodeDTO) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *NodeDTO) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *NodeDTO) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetActiveThreadCount

`func (o *NodeDTO) GetActiveThreadCount() int32`

GetActiveThreadCount returns the ActiveThreadCount field if non-nil, zero value otherwise.

### GetActiveThreadCountOk

`func (o *NodeDTO) GetActiveThreadCountOk() (*int32, bool)`

GetActiveThreadCountOk returns a tuple with the ActiveThreadCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveThreadCount

`func (o *NodeDTO) SetActiveThreadCount(v int32)`

SetActiveThreadCount sets ActiveThreadCount field to given value.

### HasActiveThreadCount

`func (o *NodeDTO) HasActiveThreadCount() bool`

HasActiveThreadCount returns a boolean if a field has been set.

### GetQueued

`func (o *NodeDTO) GetQueued() string`

GetQueued returns the Queued field if non-nil, zero value otherwise.

### GetQueuedOk

`func (o *NodeDTO) GetQueuedOk() (*string, bool)`

GetQueuedOk returns a tuple with the Queued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueued

`func (o *NodeDTO) SetQueued(v string)`

SetQueued sets Queued field to given value.

### HasQueued

`func (o *NodeDTO) HasQueued() bool`

HasQueued returns a boolean if a field has been set.

### GetEvents

`func (o *NodeDTO) GetEvents() []NodeEventDTO`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *NodeDTO) GetEventsOk() (*[]NodeEventDTO, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *NodeDTO) SetEvents(v []NodeEventDTO)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *NodeDTO) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetNodeStartTime

`func (o *NodeDTO) GetNodeStartTime() string`

GetNodeStartTime returns the NodeStartTime field if non-nil, zero value otherwise.

### GetNodeStartTimeOk

`func (o *NodeDTO) GetNodeStartTimeOk() (*string, bool)`

GetNodeStartTimeOk returns a tuple with the NodeStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeStartTime

`func (o *NodeDTO) SetNodeStartTime(v string)`

SetNodeStartTime sets NodeStartTime field to given value.

### HasNodeStartTime

`func (o *NodeDTO) HasNodeStartTime() bool`

HasNodeStartTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


