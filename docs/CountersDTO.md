# CountersDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AggregateSnapshot** | Pointer to [**CountersSnapshotDTO**](CountersSnapshotDTO.md) |  | [optional] 
**NodeSnapshots** | Pointer to [**[]NodeCountersSnapshotDTO**](NodeCountersSnapshotDTO.md) | A Counters snapshot for each node in the cluster. If the NiFi instance is a standalone instance, rather than a cluster, this may be null. | [optional] 

## Methods

### NewCountersDTO

`func NewCountersDTO() *CountersDTO`

NewCountersDTO instantiates a new CountersDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCountersDTOWithDefaults

`func NewCountersDTOWithDefaults() *CountersDTO`

NewCountersDTOWithDefaults instantiates a new CountersDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregateSnapshot

`func (o *CountersDTO) GetAggregateSnapshot() CountersSnapshotDTO`

GetAggregateSnapshot returns the AggregateSnapshot field if non-nil, zero value otherwise.

### GetAggregateSnapshotOk

`func (o *CountersDTO) GetAggregateSnapshotOk() (*CountersSnapshotDTO, bool)`

GetAggregateSnapshotOk returns a tuple with the AggregateSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregateSnapshot

`func (o *CountersDTO) SetAggregateSnapshot(v CountersSnapshotDTO)`

SetAggregateSnapshot sets AggregateSnapshot field to given value.

### HasAggregateSnapshot

`func (o *CountersDTO) HasAggregateSnapshot() bool`

HasAggregateSnapshot returns a boolean if a field has been set.

### GetNodeSnapshots

`func (o *CountersDTO) GetNodeSnapshots() []NodeCountersSnapshotDTO`

GetNodeSnapshots returns the NodeSnapshots field if non-nil, zero value otherwise.

### GetNodeSnapshotsOk

`func (o *CountersDTO) GetNodeSnapshotsOk() (*[]NodeCountersSnapshotDTO, bool)`

GetNodeSnapshotsOk returns a tuple with the NodeSnapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeSnapshots

`func (o *CountersDTO) SetNodeSnapshots(v []NodeCountersSnapshotDTO)`

SetNodeSnapshots sets NodeSnapshots field to given value.

### HasNodeSnapshots

`func (o *CountersDTO) HasNodeSnapshots() bool`

HasNodeSnapshots returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


