# ConnectionStatusDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of the connection | [optional] 
**GroupId** | Pointer to **string** | The ID of the Process Group that the connection belongs to | [optional] 
**Name** | Pointer to **string** | The name of the connection | [optional] 
**StatsLastRefreshed** | Pointer to **string** | The timestamp of when the stats were last refreshed | [optional] 
**SourceId** | Pointer to **string** | The ID of the source component | [optional] 
**SourceName** | Pointer to **string** | The name of the source component | [optional] 
**DestinationId** | Pointer to **string** | The ID of the destination component | [optional] 
**DestinationName** | Pointer to **string** | The name of the destination component | [optional] 
**AggregateSnapshot** | Pointer to [**ConnectionStatusSnapshotDTO**](ConnectionStatusSnapshotDTO.md) |  | [optional] 
**NodeSnapshots** | Pointer to [**[]NodeConnectionStatusSnapshotDTO**](NodeConnectionStatusSnapshotDTO.md) | A list of status snapshots for each node | [optional] 

## Methods

### NewConnectionStatusDTO

`func NewConnectionStatusDTO() *ConnectionStatusDTO`

NewConnectionStatusDTO instantiates a new ConnectionStatusDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionStatusDTOWithDefaults

`func NewConnectionStatusDTOWithDefaults() *ConnectionStatusDTO`

NewConnectionStatusDTOWithDefaults instantiates a new ConnectionStatusDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConnectionStatusDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConnectionStatusDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConnectionStatusDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ConnectionStatusDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetGroupId

`func (o *ConnectionStatusDTO) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ConnectionStatusDTO) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ConnectionStatusDTO) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *ConnectionStatusDTO) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetName

`func (o *ConnectionStatusDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConnectionStatusDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConnectionStatusDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConnectionStatusDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatsLastRefreshed

`func (o *ConnectionStatusDTO) GetStatsLastRefreshed() string`

GetStatsLastRefreshed returns the StatsLastRefreshed field if non-nil, zero value otherwise.

### GetStatsLastRefreshedOk

`func (o *ConnectionStatusDTO) GetStatsLastRefreshedOk() (*string, bool)`

GetStatsLastRefreshedOk returns a tuple with the StatsLastRefreshed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatsLastRefreshed

`func (o *ConnectionStatusDTO) SetStatsLastRefreshed(v string)`

SetStatsLastRefreshed sets StatsLastRefreshed field to given value.

### HasStatsLastRefreshed

`func (o *ConnectionStatusDTO) HasStatsLastRefreshed() bool`

HasStatsLastRefreshed returns a boolean if a field has been set.

### GetSourceId

`func (o *ConnectionStatusDTO) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *ConnectionStatusDTO) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *ConnectionStatusDTO) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *ConnectionStatusDTO) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### GetSourceName

`func (o *ConnectionStatusDTO) GetSourceName() string`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *ConnectionStatusDTO) GetSourceNameOk() (*string, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *ConnectionStatusDTO) SetSourceName(v string)`

SetSourceName sets SourceName field to given value.

### HasSourceName

`func (o *ConnectionStatusDTO) HasSourceName() bool`

HasSourceName returns a boolean if a field has been set.

### GetDestinationId

`func (o *ConnectionStatusDTO) GetDestinationId() string`

GetDestinationId returns the DestinationId field if non-nil, zero value otherwise.

### GetDestinationIdOk

`func (o *ConnectionStatusDTO) GetDestinationIdOk() (*string, bool)`

GetDestinationIdOk returns a tuple with the DestinationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationId

`func (o *ConnectionStatusDTO) SetDestinationId(v string)`

SetDestinationId sets DestinationId field to given value.

### HasDestinationId

`func (o *ConnectionStatusDTO) HasDestinationId() bool`

HasDestinationId returns a boolean if a field has been set.

### GetDestinationName

`func (o *ConnectionStatusDTO) GetDestinationName() string`

GetDestinationName returns the DestinationName field if non-nil, zero value otherwise.

### GetDestinationNameOk

`func (o *ConnectionStatusDTO) GetDestinationNameOk() (*string, bool)`

GetDestinationNameOk returns a tuple with the DestinationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationName

`func (o *ConnectionStatusDTO) SetDestinationName(v string)`

SetDestinationName sets DestinationName field to given value.

### HasDestinationName

`func (o *ConnectionStatusDTO) HasDestinationName() bool`

HasDestinationName returns a boolean if a field has been set.

### GetAggregateSnapshot

`func (o *ConnectionStatusDTO) GetAggregateSnapshot() ConnectionStatusSnapshotDTO`

GetAggregateSnapshot returns the AggregateSnapshot field if non-nil, zero value otherwise.

### GetAggregateSnapshotOk

`func (o *ConnectionStatusDTO) GetAggregateSnapshotOk() (*ConnectionStatusSnapshotDTO, bool)`

GetAggregateSnapshotOk returns a tuple with the AggregateSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregateSnapshot

`func (o *ConnectionStatusDTO) SetAggregateSnapshot(v ConnectionStatusSnapshotDTO)`

SetAggregateSnapshot sets AggregateSnapshot field to given value.

### HasAggregateSnapshot

`func (o *ConnectionStatusDTO) HasAggregateSnapshot() bool`

HasAggregateSnapshot returns a boolean if a field has been set.

### GetNodeSnapshots

`func (o *ConnectionStatusDTO) GetNodeSnapshots() []NodeConnectionStatusSnapshotDTO`

GetNodeSnapshots returns the NodeSnapshots field if non-nil, zero value otherwise.

### GetNodeSnapshotsOk

`func (o *ConnectionStatusDTO) GetNodeSnapshotsOk() (*[]NodeConnectionStatusSnapshotDTO, bool)`

GetNodeSnapshotsOk returns a tuple with the NodeSnapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeSnapshots

`func (o *ConnectionStatusDTO) SetNodeSnapshots(v []NodeConnectionStatusSnapshotDTO)`

SetNodeSnapshots sets NodeSnapshots field to given value.

### HasNodeSnapshots

`func (o *ConnectionStatusDTO) HasNodeSnapshots() bool`

HasNodeSnapshots returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


