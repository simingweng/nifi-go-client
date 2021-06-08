# RemoteProcessGroupStatusDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | Pointer to **string** | The unique ID of the process group that the Processor belongs to | [optional] 
**Id** | Pointer to **string** | The unique ID of the Processor | [optional] 
**Name** | Pointer to **string** | The name of the remote process group. | [optional] 
**TargetUri** | Pointer to **string** | The URI of the target system. | [optional] 
**TransmissionStatus** | Pointer to **string** | The transmission status of the remote process group. | [optional] 
**StatsLastRefreshed** | Pointer to **string** | The time the status for the process group was last refreshed. | [optional] 
**ValidationStatus** | Pointer to **string** | Indicates whether the component is valid, invalid, or still in the process of validating (i.e., it is unknown whether or not the component is valid) | [optional] [readonly] 
**AggregateSnapshot** | Pointer to [**RemoteProcessGroupStatusSnapshotDTO**](RemoteProcessGroupStatusSnapshotDTO.md) |  | [optional] 
**NodeSnapshots** | Pointer to [**[]NodeRemoteProcessGroupStatusSnapshotDTO**](NodeRemoteProcessGroupStatusSnapshotDTO.md) | A status snapshot for each node in the cluster. If the NiFi instance is a standalone instance, rather than a cluster, this may be null. | [optional] 

## Methods

### NewRemoteProcessGroupStatusDTO

`func NewRemoteProcessGroupStatusDTO() *RemoteProcessGroupStatusDTO`

NewRemoteProcessGroupStatusDTO instantiates a new RemoteProcessGroupStatusDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoteProcessGroupStatusDTOWithDefaults

`func NewRemoteProcessGroupStatusDTOWithDefaults() *RemoteProcessGroupStatusDTO`

NewRemoteProcessGroupStatusDTOWithDefaults instantiates a new RemoteProcessGroupStatusDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *RemoteProcessGroupStatusDTO) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *RemoteProcessGroupStatusDTO) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *RemoteProcessGroupStatusDTO) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *RemoteProcessGroupStatusDTO) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetId

`func (o *RemoteProcessGroupStatusDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RemoteProcessGroupStatusDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RemoteProcessGroupStatusDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RemoteProcessGroupStatusDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *RemoteProcessGroupStatusDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RemoteProcessGroupStatusDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RemoteProcessGroupStatusDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RemoteProcessGroupStatusDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTargetUri

`func (o *RemoteProcessGroupStatusDTO) GetTargetUri() string`

GetTargetUri returns the TargetUri field if non-nil, zero value otherwise.

### GetTargetUriOk

`func (o *RemoteProcessGroupStatusDTO) GetTargetUriOk() (*string, bool)`

GetTargetUriOk returns a tuple with the TargetUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUri

`func (o *RemoteProcessGroupStatusDTO) SetTargetUri(v string)`

SetTargetUri sets TargetUri field to given value.

### HasTargetUri

`func (o *RemoteProcessGroupStatusDTO) HasTargetUri() bool`

HasTargetUri returns a boolean if a field has been set.

### GetTransmissionStatus

`func (o *RemoteProcessGroupStatusDTO) GetTransmissionStatus() string`

GetTransmissionStatus returns the TransmissionStatus field if non-nil, zero value otherwise.

### GetTransmissionStatusOk

`func (o *RemoteProcessGroupStatusDTO) GetTransmissionStatusOk() (*string, bool)`

GetTransmissionStatusOk returns a tuple with the TransmissionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransmissionStatus

`func (o *RemoteProcessGroupStatusDTO) SetTransmissionStatus(v string)`

SetTransmissionStatus sets TransmissionStatus field to given value.

### HasTransmissionStatus

`func (o *RemoteProcessGroupStatusDTO) HasTransmissionStatus() bool`

HasTransmissionStatus returns a boolean if a field has been set.

### GetStatsLastRefreshed

`func (o *RemoteProcessGroupStatusDTO) GetStatsLastRefreshed() string`

GetStatsLastRefreshed returns the StatsLastRefreshed field if non-nil, zero value otherwise.

### GetStatsLastRefreshedOk

`func (o *RemoteProcessGroupStatusDTO) GetStatsLastRefreshedOk() (*string, bool)`

GetStatsLastRefreshedOk returns a tuple with the StatsLastRefreshed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatsLastRefreshed

`func (o *RemoteProcessGroupStatusDTO) SetStatsLastRefreshed(v string)`

SetStatsLastRefreshed sets StatsLastRefreshed field to given value.

### HasStatsLastRefreshed

`func (o *RemoteProcessGroupStatusDTO) HasStatsLastRefreshed() bool`

HasStatsLastRefreshed returns a boolean if a field has been set.

### GetValidationStatus

`func (o *RemoteProcessGroupStatusDTO) GetValidationStatus() string`

GetValidationStatus returns the ValidationStatus field if non-nil, zero value otherwise.

### GetValidationStatusOk

`func (o *RemoteProcessGroupStatusDTO) GetValidationStatusOk() (*string, bool)`

GetValidationStatusOk returns a tuple with the ValidationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationStatus

`func (o *RemoteProcessGroupStatusDTO) SetValidationStatus(v string)`

SetValidationStatus sets ValidationStatus field to given value.

### HasValidationStatus

`func (o *RemoteProcessGroupStatusDTO) HasValidationStatus() bool`

HasValidationStatus returns a boolean if a field has been set.

### GetAggregateSnapshot

`func (o *RemoteProcessGroupStatusDTO) GetAggregateSnapshot() RemoteProcessGroupStatusSnapshotDTO`

GetAggregateSnapshot returns the AggregateSnapshot field if non-nil, zero value otherwise.

### GetAggregateSnapshotOk

`func (o *RemoteProcessGroupStatusDTO) GetAggregateSnapshotOk() (*RemoteProcessGroupStatusSnapshotDTO, bool)`

GetAggregateSnapshotOk returns a tuple with the AggregateSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregateSnapshot

`func (o *RemoteProcessGroupStatusDTO) SetAggregateSnapshot(v RemoteProcessGroupStatusSnapshotDTO)`

SetAggregateSnapshot sets AggregateSnapshot field to given value.

### HasAggregateSnapshot

`func (o *RemoteProcessGroupStatusDTO) HasAggregateSnapshot() bool`

HasAggregateSnapshot returns a boolean if a field has been set.

### GetNodeSnapshots

`func (o *RemoteProcessGroupStatusDTO) GetNodeSnapshots() []NodeRemoteProcessGroupStatusSnapshotDTO`

GetNodeSnapshots returns the NodeSnapshots field if non-nil, zero value otherwise.

### GetNodeSnapshotsOk

`func (o *RemoteProcessGroupStatusDTO) GetNodeSnapshotsOk() (*[]NodeRemoteProcessGroupStatusSnapshotDTO, bool)`

GetNodeSnapshotsOk returns a tuple with the NodeSnapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeSnapshots

`func (o *RemoteProcessGroupStatusDTO) SetNodeSnapshots(v []NodeRemoteProcessGroupStatusSnapshotDTO)`

SetNodeSnapshots sets NodeSnapshots field to given value.

### HasNodeSnapshots

`func (o *RemoteProcessGroupStatusDTO) HasNodeSnapshots() bool`

HasNodeSnapshots returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


