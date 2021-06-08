# ProcessGroupImportEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProcessGroupRevision** | Pointer to [**RevisionDTO**](RevisionDTO.md) |  | [optional] 
**DisconnectedNodeAcknowledged** | Pointer to **bool** | Acknowledges that this node is disconnected to allow for mutable requests to proceed. | [optional] 
**VersionedFlowSnapshot** | Pointer to [**VersionedFlowSnapshot**](VersionedFlowSnapshot.md) |  | [optional] 

## Methods

### NewProcessGroupImportEntity

`func NewProcessGroupImportEntity() *ProcessGroupImportEntity`

NewProcessGroupImportEntity instantiates a new ProcessGroupImportEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessGroupImportEntityWithDefaults

`func NewProcessGroupImportEntityWithDefaults() *ProcessGroupImportEntity`

NewProcessGroupImportEntityWithDefaults instantiates a new ProcessGroupImportEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProcessGroupRevision

`func (o *ProcessGroupImportEntity) GetProcessGroupRevision() RevisionDTO`

GetProcessGroupRevision returns the ProcessGroupRevision field if non-nil, zero value otherwise.

### GetProcessGroupRevisionOk

`func (o *ProcessGroupImportEntity) GetProcessGroupRevisionOk() (*RevisionDTO, bool)`

GetProcessGroupRevisionOk returns a tuple with the ProcessGroupRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessGroupRevision

`func (o *ProcessGroupImportEntity) SetProcessGroupRevision(v RevisionDTO)`

SetProcessGroupRevision sets ProcessGroupRevision field to given value.

### HasProcessGroupRevision

`func (o *ProcessGroupImportEntity) HasProcessGroupRevision() bool`

HasProcessGroupRevision returns a boolean if a field has been set.

### GetDisconnectedNodeAcknowledged

`func (o *ProcessGroupImportEntity) GetDisconnectedNodeAcknowledged() bool`

GetDisconnectedNodeAcknowledged returns the DisconnectedNodeAcknowledged field if non-nil, zero value otherwise.

### GetDisconnectedNodeAcknowledgedOk

`func (o *ProcessGroupImportEntity) GetDisconnectedNodeAcknowledgedOk() (*bool, bool)`

GetDisconnectedNodeAcknowledgedOk returns a tuple with the DisconnectedNodeAcknowledged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisconnectedNodeAcknowledged

`func (o *ProcessGroupImportEntity) SetDisconnectedNodeAcknowledged(v bool)`

SetDisconnectedNodeAcknowledged sets DisconnectedNodeAcknowledged field to given value.

### HasDisconnectedNodeAcknowledged

`func (o *ProcessGroupImportEntity) HasDisconnectedNodeAcknowledged() bool`

HasDisconnectedNodeAcknowledged returns a boolean if a field has been set.

### GetVersionedFlowSnapshot

`func (o *ProcessGroupImportEntity) GetVersionedFlowSnapshot() VersionedFlowSnapshot`

GetVersionedFlowSnapshot returns the VersionedFlowSnapshot field if non-nil, zero value otherwise.

### GetVersionedFlowSnapshotOk

`func (o *ProcessGroupImportEntity) GetVersionedFlowSnapshotOk() (*VersionedFlowSnapshot, bool)`

GetVersionedFlowSnapshotOk returns a tuple with the VersionedFlowSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionedFlowSnapshot

`func (o *ProcessGroupImportEntity) SetVersionedFlowSnapshot(v VersionedFlowSnapshot)`

SetVersionedFlowSnapshot sets VersionedFlowSnapshot field to given value.

### HasVersionedFlowSnapshot

`func (o *ProcessGroupImportEntity) HasVersionedFlowSnapshot() bool`

HasVersionedFlowSnapshot returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


