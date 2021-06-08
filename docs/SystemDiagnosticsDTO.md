# SystemDiagnosticsDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AggregateSnapshot** | Pointer to [**SystemDiagnosticsSnapshotDTO**](SystemDiagnosticsSnapshotDTO.md) |  | [optional] 
**NodeSnapshots** | Pointer to [**[]NodeSystemDiagnosticsSnapshotDTO**](NodeSystemDiagnosticsSnapshotDTO.md) | A systems diagnostics snapshot for each node in the cluster. If the NiFi instance is a standalone instance, rather than a cluster, this may be null. | [optional] 

## Methods

### NewSystemDiagnosticsDTO

`func NewSystemDiagnosticsDTO() *SystemDiagnosticsDTO`

NewSystemDiagnosticsDTO instantiates a new SystemDiagnosticsDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemDiagnosticsDTOWithDefaults

`func NewSystemDiagnosticsDTOWithDefaults() *SystemDiagnosticsDTO`

NewSystemDiagnosticsDTOWithDefaults instantiates a new SystemDiagnosticsDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregateSnapshot

`func (o *SystemDiagnosticsDTO) GetAggregateSnapshot() SystemDiagnosticsSnapshotDTO`

GetAggregateSnapshot returns the AggregateSnapshot field if non-nil, zero value otherwise.

### GetAggregateSnapshotOk

`func (o *SystemDiagnosticsDTO) GetAggregateSnapshotOk() (*SystemDiagnosticsSnapshotDTO, bool)`

GetAggregateSnapshotOk returns a tuple with the AggregateSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregateSnapshot

`func (o *SystemDiagnosticsDTO) SetAggregateSnapshot(v SystemDiagnosticsSnapshotDTO)`

SetAggregateSnapshot sets AggregateSnapshot field to given value.

### HasAggregateSnapshot

`func (o *SystemDiagnosticsDTO) HasAggregateSnapshot() bool`

HasAggregateSnapshot returns a boolean if a field has been set.

### GetNodeSnapshots

`func (o *SystemDiagnosticsDTO) GetNodeSnapshots() []NodeSystemDiagnosticsSnapshotDTO`

GetNodeSnapshots returns the NodeSnapshots field if non-nil, zero value otherwise.

### GetNodeSnapshotsOk

`func (o *SystemDiagnosticsDTO) GetNodeSnapshotsOk() (*[]NodeSystemDiagnosticsSnapshotDTO, bool)`

GetNodeSnapshotsOk returns a tuple with the NodeSnapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeSnapshots

`func (o *SystemDiagnosticsDTO) SetNodeSnapshots(v []NodeSystemDiagnosticsSnapshotDTO)`

SetNodeSnapshots sets NodeSnapshots field to given value.

### HasNodeSnapshots

`func (o *SystemDiagnosticsDTO) HasNodeSnapshots() bool`

HasNodeSnapshots returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


