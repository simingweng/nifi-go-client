# PortStatusDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The id of the port. | [optional] 
**GroupId** | Pointer to **string** | The id of the parent process group of the port. | [optional] 
**Name** | Pointer to **string** | The name of the port. | [optional] 
**Transmitting** | Pointer to **bool** | Whether the port has incoming or outgoing connections to a remote NiFi. | [optional] 
**RunStatus** | Pointer to **string** | The run status of the port. | [optional] 
**StatsLastRefreshed** | Pointer to **string** | The time the status for the process group was last refreshed. | [optional] 
**AggregateSnapshot** | Pointer to [**PortStatusSnapshotDTO**](PortStatusSnapshotDTO.md) |  | [optional] 
**NodeSnapshots** | Pointer to [**[]NodePortStatusSnapshotDTO**](NodePortStatusSnapshotDTO.md) | A status snapshot for each node in the cluster. If the NiFi instance is a standalone instance, rather than a cluster, this may be null. | [optional] 

## Methods

### NewPortStatusDTO

`func NewPortStatusDTO() *PortStatusDTO`

NewPortStatusDTO instantiates a new PortStatusDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortStatusDTOWithDefaults

`func NewPortStatusDTOWithDefaults() *PortStatusDTO`

NewPortStatusDTOWithDefaults instantiates a new PortStatusDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PortStatusDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PortStatusDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PortStatusDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PortStatusDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetGroupId

`func (o *PortStatusDTO) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *PortStatusDTO) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *PortStatusDTO) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *PortStatusDTO) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetName

`func (o *PortStatusDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PortStatusDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PortStatusDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PortStatusDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTransmitting

`func (o *PortStatusDTO) GetTransmitting() bool`

GetTransmitting returns the Transmitting field if non-nil, zero value otherwise.

### GetTransmittingOk

`func (o *PortStatusDTO) GetTransmittingOk() (*bool, bool)`

GetTransmittingOk returns a tuple with the Transmitting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransmitting

`func (o *PortStatusDTO) SetTransmitting(v bool)`

SetTransmitting sets Transmitting field to given value.

### HasTransmitting

`func (o *PortStatusDTO) HasTransmitting() bool`

HasTransmitting returns a boolean if a field has been set.

### GetRunStatus

`func (o *PortStatusDTO) GetRunStatus() string`

GetRunStatus returns the RunStatus field if non-nil, zero value otherwise.

### GetRunStatusOk

`func (o *PortStatusDTO) GetRunStatusOk() (*string, bool)`

GetRunStatusOk returns a tuple with the RunStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunStatus

`func (o *PortStatusDTO) SetRunStatus(v string)`

SetRunStatus sets RunStatus field to given value.

### HasRunStatus

`func (o *PortStatusDTO) HasRunStatus() bool`

HasRunStatus returns a boolean if a field has been set.

### GetStatsLastRefreshed

`func (o *PortStatusDTO) GetStatsLastRefreshed() string`

GetStatsLastRefreshed returns the StatsLastRefreshed field if non-nil, zero value otherwise.

### GetStatsLastRefreshedOk

`func (o *PortStatusDTO) GetStatsLastRefreshedOk() (*string, bool)`

GetStatsLastRefreshedOk returns a tuple with the StatsLastRefreshed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatsLastRefreshed

`func (o *PortStatusDTO) SetStatsLastRefreshed(v string)`

SetStatsLastRefreshed sets StatsLastRefreshed field to given value.

### HasStatsLastRefreshed

`func (o *PortStatusDTO) HasStatsLastRefreshed() bool`

HasStatsLastRefreshed returns a boolean if a field has been set.

### GetAggregateSnapshot

`func (o *PortStatusDTO) GetAggregateSnapshot() PortStatusSnapshotDTO`

GetAggregateSnapshot returns the AggregateSnapshot field if non-nil, zero value otherwise.

### GetAggregateSnapshotOk

`func (o *PortStatusDTO) GetAggregateSnapshotOk() (*PortStatusSnapshotDTO, bool)`

GetAggregateSnapshotOk returns a tuple with the AggregateSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregateSnapshot

`func (o *PortStatusDTO) SetAggregateSnapshot(v PortStatusSnapshotDTO)`

SetAggregateSnapshot sets AggregateSnapshot field to given value.

### HasAggregateSnapshot

`func (o *PortStatusDTO) HasAggregateSnapshot() bool`

HasAggregateSnapshot returns a boolean if a field has been set.

### GetNodeSnapshots

`func (o *PortStatusDTO) GetNodeSnapshots() []NodePortStatusSnapshotDTO`

GetNodeSnapshots returns the NodeSnapshots field if non-nil, zero value otherwise.

### GetNodeSnapshotsOk

`func (o *PortStatusDTO) GetNodeSnapshotsOk() (*[]NodePortStatusSnapshotDTO, bool)`

GetNodeSnapshotsOk returns a tuple with the NodeSnapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeSnapshots

`func (o *PortStatusDTO) SetNodeSnapshots(v []NodePortStatusSnapshotDTO)`

SetNodeSnapshots sets NodeSnapshots field to given value.

### HasNodeSnapshots

`func (o *PortStatusDTO) HasNodeSnapshots() bool`

HasNodeSnapshots returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


