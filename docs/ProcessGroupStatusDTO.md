# ProcessGroupStatusDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of the Process Group | [optional] 
**Name** | Pointer to **string** | The name of the Process Group | [optional] 
**StatsLastRefreshed** | Pointer to **string** | The time the status for the process group was last refreshed. | [optional] 
**AggregateSnapshot** | Pointer to [**ProcessGroupStatusSnapshotDTO**](ProcessGroupStatusSnapshotDTO.md) |  | [optional] 
**NodeSnapshots** | Pointer to [**[]NodeProcessGroupStatusSnapshotDTO**](NodeProcessGroupStatusSnapshotDTO.md) | The status reported by each node in the cluster. If the NiFi instance is a standalone instance, rather than a clustered instance, this value may be null. | [optional] 

## Methods

### NewProcessGroupStatusDTO

`func NewProcessGroupStatusDTO() *ProcessGroupStatusDTO`

NewProcessGroupStatusDTO instantiates a new ProcessGroupStatusDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessGroupStatusDTOWithDefaults

`func NewProcessGroupStatusDTOWithDefaults() *ProcessGroupStatusDTO`

NewProcessGroupStatusDTOWithDefaults instantiates a new ProcessGroupStatusDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProcessGroupStatusDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProcessGroupStatusDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProcessGroupStatusDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProcessGroupStatusDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ProcessGroupStatusDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProcessGroupStatusDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProcessGroupStatusDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProcessGroupStatusDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatsLastRefreshed

`func (o *ProcessGroupStatusDTO) GetStatsLastRefreshed() string`

GetStatsLastRefreshed returns the StatsLastRefreshed field if non-nil, zero value otherwise.

### GetStatsLastRefreshedOk

`func (o *ProcessGroupStatusDTO) GetStatsLastRefreshedOk() (*string, bool)`

GetStatsLastRefreshedOk returns a tuple with the StatsLastRefreshed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatsLastRefreshed

`func (o *ProcessGroupStatusDTO) SetStatsLastRefreshed(v string)`

SetStatsLastRefreshed sets StatsLastRefreshed field to given value.

### HasStatsLastRefreshed

`func (o *ProcessGroupStatusDTO) HasStatsLastRefreshed() bool`

HasStatsLastRefreshed returns a boolean if a field has been set.

### GetAggregateSnapshot

`func (o *ProcessGroupStatusDTO) GetAggregateSnapshot() ProcessGroupStatusSnapshotDTO`

GetAggregateSnapshot returns the AggregateSnapshot field if non-nil, zero value otherwise.

### GetAggregateSnapshotOk

`func (o *ProcessGroupStatusDTO) GetAggregateSnapshotOk() (*ProcessGroupStatusSnapshotDTO, bool)`

GetAggregateSnapshotOk returns a tuple with the AggregateSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregateSnapshot

`func (o *ProcessGroupStatusDTO) SetAggregateSnapshot(v ProcessGroupStatusSnapshotDTO)`

SetAggregateSnapshot sets AggregateSnapshot field to given value.

### HasAggregateSnapshot

`func (o *ProcessGroupStatusDTO) HasAggregateSnapshot() bool`

HasAggregateSnapshot returns a boolean if a field has been set.

### GetNodeSnapshots

`func (o *ProcessGroupStatusDTO) GetNodeSnapshots() []NodeProcessGroupStatusSnapshotDTO`

GetNodeSnapshots returns the NodeSnapshots field if non-nil, zero value otherwise.

### GetNodeSnapshotsOk

`func (o *ProcessGroupStatusDTO) GetNodeSnapshotsOk() (*[]NodeProcessGroupStatusSnapshotDTO, bool)`

GetNodeSnapshotsOk returns a tuple with the NodeSnapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeSnapshots

`func (o *ProcessGroupStatusDTO) SetNodeSnapshots(v []NodeProcessGroupStatusSnapshotDTO)`

SetNodeSnapshots sets NodeSnapshots field to given value.

### HasNodeSnapshots

`func (o *ProcessGroupStatusDTO) HasNodeSnapshots() bool`

HasNodeSnapshots returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


