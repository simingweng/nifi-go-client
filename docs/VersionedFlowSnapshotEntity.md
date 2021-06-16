# VersionedFlowSnapshotEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VersionedFlowSnapshot** | Pointer to [**VersionedFlowSnapshot**](VersionedFlowSnapshot.md) |  | [optional] 
**ProcessGroupRevision** | Pointer to [**RevisionDTO**](RevisionDTO.md) |  | [optional] 
**RegistryId** | Pointer to **string** | The ID of the Registry that this flow belongs to | [optional] 
**UpdateDescendantVersionedFlows** | Pointer to **bool** | If the Process Group to be updated has a child or descendant Process Group that is also under Version Control, this specifies whether or not the contents of that child/descendant Process Group should be updated. | [optional] 
**DisconnectedNodeAcknowledged** | Pointer to **bool** | Acknowledges that this node is disconnected to allow for mutable requests to proceed. | [optional] 

## Methods

### NewVersionedFlowSnapshotEntity

`func NewVersionedFlowSnapshotEntity() *VersionedFlowSnapshotEntity`

NewVersionedFlowSnapshotEntity instantiates a new VersionedFlowSnapshotEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersionedFlowSnapshotEntityWithDefaults

`func NewVersionedFlowSnapshotEntityWithDefaults() *VersionedFlowSnapshotEntity`

NewVersionedFlowSnapshotEntityWithDefaults instantiates a new VersionedFlowSnapshotEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersionedFlowSnapshot

`func (o *VersionedFlowSnapshotEntity) GetVersionedFlowSnapshot() VersionedFlowSnapshot`

GetVersionedFlowSnapshot returns the VersionedFlowSnapshot field if non-nil, zero value otherwise.

### GetVersionedFlowSnapshotOk

`func (o *VersionedFlowSnapshotEntity) GetVersionedFlowSnapshotOk() (*VersionedFlowSnapshot, bool)`

GetVersionedFlowSnapshotOk returns a tuple with the VersionedFlowSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionedFlowSnapshot

`func (o *VersionedFlowSnapshotEntity) SetVersionedFlowSnapshot(v VersionedFlowSnapshot)`

SetVersionedFlowSnapshot sets VersionedFlowSnapshot field to given value.

### HasVersionedFlowSnapshot

`func (o *VersionedFlowSnapshotEntity) HasVersionedFlowSnapshot() bool`

HasVersionedFlowSnapshot returns a boolean if a field has been set.

### GetProcessGroupRevision

`func (o *VersionedFlowSnapshotEntity) GetProcessGroupRevision() RevisionDTO`

GetProcessGroupRevision returns the ProcessGroupRevision field if non-nil, zero value otherwise.

### GetProcessGroupRevisionOk

`func (o *VersionedFlowSnapshotEntity) GetProcessGroupRevisionOk() (*RevisionDTO, bool)`

GetProcessGroupRevisionOk returns a tuple with the ProcessGroupRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessGroupRevision

`func (o *VersionedFlowSnapshotEntity) SetProcessGroupRevision(v RevisionDTO)`

SetProcessGroupRevision sets ProcessGroupRevision field to given value.

### HasProcessGroupRevision

`func (o *VersionedFlowSnapshotEntity) HasProcessGroupRevision() bool`

HasProcessGroupRevision returns a boolean if a field has been set.

### GetRegistryId

`func (o *VersionedFlowSnapshotEntity) GetRegistryId() string`

GetRegistryId returns the RegistryId field if non-nil, zero value otherwise.

### GetRegistryIdOk

`func (o *VersionedFlowSnapshotEntity) GetRegistryIdOk() (*string, bool)`

GetRegistryIdOk returns a tuple with the RegistryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryId

`func (o *VersionedFlowSnapshotEntity) SetRegistryId(v string)`

SetRegistryId sets RegistryId field to given value.

### HasRegistryId

`func (o *VersionedFlowSnapshotEntity) HasRegistryId() bool`

HasRegistryId returns a boolean if a field has been set.

### GetUpdateDescendantVersionedFlows

`func (o *VersionedFlowSnapshotEntity) GetUpdateDescendantVersionedFlows() bool`

GetUpdateDescendantVersionedFlows returns the UpdateDescendantVersionedFlows field if non-nil, zero value otherwise.

### GetUpdateDescendantVersionedFlowsOk

`func (o *VersionedFlowSnapshotEntity) GetUpdateDescendantVersionedFlowsOk() (*bool, bool)`

GetUpdateDescendantVersionedFlowsOk returns a tuple with the UpdateDescendantVersionedFlows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateDescendantVersionedFlows

`func (o *VersionedFlowSnapshotEntity) SetUpdateDescendantVersionedFlows(v bool)`

SetUpdateDescendantVersionedFlows sets UpdateDescendantVersionedFlows field to given value.

### HasUpdateDescendantVersionedFlows

`func (o *VersionedFlowSnapshotEntity) HasUpdateDescendantVersionedFlows() bool`

HasUpdateDescendantVersionedFlows returns a boolean if a field has been set.

### GetDisconnectedNodeAcknowledged

`func (o *VersionedFlowSnapshotEntity) GetDisconnectedNodeAcknowledged() bool`

GetDisconnectedNodeAcknowledged returns the DisconnectedNodeAcknowledged field if non-nil, zero value otherwise.

### GetDisconnectedNodeAcknowledgedOk

`func (o *VersionedFlowSnapshotEntity) GetDisconnectedNodeAcknowledgedOk() (*bool, bool)`

GetDisconnectedNodeAcknowledgedOk returns a tuple with the DisconnectedNodeAcknowledged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisconnectedNodeAcknowledged

`func (o *VersionedFlowSnapshotEntity) SetDisconnectedNodeAcknowledged(v bool)`

SetDisconnectedNodeAcknowledged sets DisconnectedNodeAcknowledged field to given value.

### HasDisconnectedNodeAcknowledged

`func (o *VersionedFlowSnapshotEntity) HasDisconnectedNodeAcknowledged() bool`

HasDisconnectedNodeAcknowledged returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


