# ConnectionStatisticsDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of the connection | [optional] 
**StatsLastRefreshed** | Pointer to **string** | The timestamp of when the stats were last refreshed | [optional] 
**AggregateSnapshot** | Pointer to [**ConnectionStatisticsSnapshotDTO**](ConnectionStatisticsSnapshotDTO.md) |  | [optional] 
**NodeSnapshots** | Pointer to [**[]NodeConnectionStatisticsSnapshotDTO**](NodeConnectionStatisticsSnapshotDTO.md) | A list of status snapshots for each node | [optional] 

## Methods

### NewConnectionStatisticsDTO

`func NewConnectionStatisticsDTO() *ConnectionStatisticsDTO`

NewConnectionStatisticsDTO instantiates a new ConnectionStatisticsDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionStatisticsDTOWithDefaults

`func NewConnectionStatisticsDTOWithDefaults() *ConnectionStatisticsDTO`

NewConnectionStatisticsDTOWithDefaults instantiates a new ConnectionStatisticsDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConnectionStatisticsDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConnectionStatisticsDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConnectionStatisticsDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ConnectionStatisticsDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatsLastRefreshed

`func (o *ConnectionStatisticsDTO) GetStatsLastRefreshed() string`

GetStatsLastRefreshed returns the StatsLastRefreshed field if non-nil, zero value otherwise.

### GetStatsLastRefreshedOk

`func (o *ConnectionStatisticsDTO) GetStatsLastRefreshedOk() (*string, bool)`

GetStatsLastRefreshedOk returns a tuple with the StatsLastRefreshed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatsLastRefreshed

`func (o *ConnectionStatisticsDTO) SetStatsLastRefreshed(v string)`

SetStatsLastRefreshed sets StatsLastRefreshed field to given value.

### HasStatsLastRefreshed

`func (o *ConnectionStatisticsDTO) HasStatsLastRefreshed() bool`

HasStatsLastRefreshed returns a boolean if a field has been set.

### GetAggregateSnapshot

`func (o *ConnectionStatisticsDTO) GetAggregateSnapshot() ConnectionStatisticsSnapshotDTO`

GetAggregateSnapshot returns the AggregateSnapshot field if non-nil, zero value otherwise.

### GetAggregateSnapshotOk

`func (o *ConnectionStatisticsDTO) GetAggregateSnapshotOk() (*ConnectionStatisticsSnapshotDTO, bool)`

GetAggregateSnapshotOk returns a tuple with the AggregateSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregateSnapshot

`func (o *ConnectionStatisticsDTO) SetAggregateSnapshot(v ConnectionStatisticsSnapshotDTO)`

SetAggregateSnapshot sets AggregateSnapshot field to given value.

### HasAggregateSnapshot

`func (o *ConnectionStatisticsDTO) HasAggregateSnapshot() bool`

HasAggregateSnapshot returns a boolean if a field has been set.

### GetNodeSnapshots

`func (o *ConnectionStatisticsDTO) GetNodeSnapshots() []NodeConnectionStatisticsSnapshotDTO`

GetNodeSnapshots returns the NodeSnapshots field if non-nil, zero value otherwise.

### GetNodeSnapshotsOk

`func (o *ConnectionStatisticsDTO) GetNodeSnapshotsOk() (*[]NodeConnectionStatisticsSnapshotDTO, bool)`

GetNodeSnapshotsOk returns a tuple with the NodeSnapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeSnapshots

`func (o *ConnectionStatisticsDTO) SetNodeSnapshots(v []NodeConnectionStatisticsSnapshotDTO)`

SetNodeSnapshots sets NodeSnapshots field to given value.

### HasNodeSnapshots

`func (o *ConnectionStatisticsDTO) HasNodeSnapshots() bool`

HasNodeSnapshots returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


