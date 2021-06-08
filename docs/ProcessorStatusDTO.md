# ProcessorStatusDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | Pointer to **string** | The unique ID of the process group that the Processor belongs to | [optional] 
**Id** | Pointer to **string** | The unique ID of the Processor | [optional] 
**Name** | Pointer to **string** | The name of the Processor | [optional] 
**Type** | Pointer to **string** | The type of the Processor | [optional] 
**RunStatus** | Pointer to **string** | The run status of the Processor | [optional] 
**StatsLastRefreshed** | Pointer to **string** | The timestamp of when the stats were last refreshed | [optional] 
**AggregateSnapshot** | Pointer to [**ProcessorStatusSnapshotDTO**](ProcessorStatusSnapshotDTO.md) |  | [optional] 
**NodeSnapshots** | Pointer to [**[]NodeProcessorStatusSnapshotDTO**](NodeProcessorStatusSnapshotDTO.md) | A status snapshot for each node in the cluster. If the NiFi instance is a standalone instance, rather than a cluster, this may be null. | [optional] 

## Methods

### NewProcessorStatusDTO

`func NewProcessorStatusDTO() *ProcessorStatusDTO`

NewProcessorStatusDTO instantiates a new ProcessorStatusDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessorStatusDTOWithDefaults

`func NewProcessorStatusDTOWithDefaults() *ProcessorStatusDTO`

NewProcessorStatusDTOWithDefaults instantiates a new ProcessorStatusDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *ProcessorStatusDTO) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ProcessorStatusDTO) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ProcessorStatusDTO) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *ProcessorStatusDTO) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetId

`func (o *ProcessorStatusDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProcessorStatusDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProcessorStatusDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProcessorStatusDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ProcessorStatusDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProcessorStatusDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProcessorStatusDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProcessorStatusDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *ProcessorStatusDTO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProcessorStatusDTO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProcessorStatusDTO) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ProcessorStatusDTO) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRunStatus

`func (o *ProcessorStatusDTO) GetRunStatus() string`

GetRunStatus returns the RunStatus field if non-nil, zero value otherwise.

### GetRunStatusOk

`func (o *ProcessorStatusDTO) GetRunStatusOk() (*string, bool)`

GetRunStatusOk returns a tuple with the RunStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunStatus

`func (o *ProcessorStatusDTO) SetRunStatus(v string)`

SetRunStatus sets RunStatus field to given value.

### HasRunStatus

`func (o *ProcessorStatusDTO) HasRunStatus() bool`

HasRunStatus returns a boolean if a field has been set.

### GetStatsLastRefreshed

`func (o *ProcessorStatusDTO) GetStatsLastRefreshed() string`

GetStatsLastRefreshed returns the StatsLastRefreshed field if non-nil, zero value otherwise.

### GetStatsLastRefreshedOk

`func (o *ProcessorStatusDTO) GetStatsLastRefreshedOk() (*string, bool)`

GetStatsLastRefreshedOk returns a tuple with the StatsLastRefreshed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatsLastRefreshed

`func (o *ProcessorStatusDTO) SetStatsLastRefreshed(v string)`

SetStatsLastRefreshed sets StatsLastRefreshed field to given value.

### HasStatsLastRefreshed

`func (o *ProcessorStatusDTO) HasStatsLastRefreshed() bool`

HasStatsLastRefreshed returns a boolean if a field has been set.

### GetAggregateSnapshot

`func (o *ProcessorStatusDTO) GetAggregateSnapshot() ProcessorStatusSnapshotDTO`

GetAggregateSnapshot returns the AggregateSnapshot field if non-nil, zero value otherwise.

### GetAggregateSnapshotOk

`func (o *ProcessorStatusDTO) GetAggregateSnapshotOk() (*ProcessorStatusSnapshotDTO, bool)`

GetAggregateSnapshotOk returns a tuple with the AggregateSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregateSnapshot

`func (o *ProcessorStatusDTO) SetAggregateSnapshot(v ProcessorStatusSnapshotDTO)`

SetAggregateSnapshot sets AggregateSnapshot field to given value.

### HasAggregateSnapshot

`func (o *ProcessorStatusDTO) HasAggregateSnapshot() bool`

HasAggregateSnapshot returns a boolean if a field has been set.

### GetNodeSnapshots

`func (o *ProcessorStatusDTO) GetNodeSnapshots() []NodeProcessorStatusSnapshotDTO`

GetNodeSnapshots returns the NodeSnapshots field if non-nil, zero value otherwise.

### GetNodeSnapshotsOk

`func (o *ProcessorStatusDTO) GetNodeSnapshotsOk() (*[]NodeProcessorStatusSnapshotDTO, bool)`

GetNodeSnapshotsOk returns a tuple with the NodeSnapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeSnapshots

`func (o *ProcessorStatusDTO) SetNodeSnapshots(v []NodeProcessorStatusSnapshotDTO)`

SetNodeSnapshots sets NodeSnapshots field to given value.

### HasNodeSnapshots

`func (o *ProcessorStatusDTO) HasNodeSnapshots() bool`

HasNodeSnapshots returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


