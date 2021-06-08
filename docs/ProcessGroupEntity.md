# ProcessGroupEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Revision** | Pointer to [**RevisionDTO**](RevisionDTO.md) |  | [optional] 
**Id** | Pointer to **string** | The id of the component. | [optional] 
**Uri** | Pointer to **string** | The URI for futures requests to the component. | [optional] 
**Position** | Pointer to [**PositionDTO**](PositionDTO.md) |  | [optional] 
**Permissions** | Pointer to [**PermissionsDTO**](PermissionsDTO.md) |  | [optional] 
**Bulletins** | Pointer to [**[]BulletinEntity**](BulletinEntity.md) | The bulletins for this component. | [optional] 
**DisconnectedNodeAcknowledged** | Pointer to **bool** | Acknowledges that this node is disconnected to allow for mutable requests to proceed. | [optional] 
**Component** | Pointer to [**ProcessGroupDTO**](ProcessGroupDTO.md) |  | [optional] 
**Status** | Pointer to [**ProcessGroupStatusDTO**](ProcessGroupStatusDTO.md) |  | [optional] 
**VersionedFlowSnapshot** | Pointer to [**VersionedFlowSnapshot**](VersionedFlowSnapshot.md) |  | [optional] 
**RunningCount** | Pointer to **int32** | The number of running components in this process group. | [optional] 
**StoppedCount** | Pointer to **int32** | The number of stopped components in the process group. | [optional] 
**InvalidCount** | Pointer to **int32** | The number of invalid components in the process group. | [optional] 
**DisabledCount** | Pointer to **int32** | The number of disabled components in the process group. | [optional] 
**ActiveRemotePortCount** | Pointer to **int32** | The number of active remote ports in the process group. | [optional] 
**InactiveRemotePortCount** | Pointer to **int32** | The number of inactive remote ports in the process group. | [optional] 
**VersionedFlowState** | Pointer to **string** | The current state of the Process Group, as it relates to the Versioned Flow | [optional] [readonly] 
**UpToDateCount** | Pointer to **int32** | The number of up to date versioned process groups in the process group. | [optional] 
**LocallyModifiedCount** | Pointer to **int32** | The number of locally modified versioned process groups in the process group. | [optional] 
**StaleCount** | Pointer to **int32** | The number of stale versioned process groups in the process group. | [optional] 
**LocallyModifiedAndStaleCount** | Pointer to **int32** | The number of locally modified and stale versioned process groups in the process group. | [optional] 
**SyncFailureCount** | Pointer to **int32** | The number of versioned process groups in the process group that are unable to sync to a registry. | [optional] 
**LocalInputPortCount** | Pointer to **int32** | The number of local input ports in the process group. | [optional] 
**LocalOutputPortCount** | Pointer to **int32** | The number of local output ports in the process group. | [optional] 
**PublicInputPortCount** | Pointer to **int32** | The number of public input ports in the process group. | [optional] 
**PublicOutputPortCount** | Pointer to **int32** | The number of public output ports in the process group. | [optional] 
**ParameterContext** | Pointer to [**ParameterContextReferenceEntity**](ParameterContextReferenceEntity.md) |  | [optional] 
**InputPortCount** | Pointer to **int32** | The number of input ports in the process group. | [optional] [readonly] 
**OutputPortCount** | Pointer to **int32** | The number of output ports in the process group. | [optional] [readonly] 

## Methods

### NewProcessGroupEntity

`func NewProcessGroupEntity() *ProcessGroupEntity`

NewProcessGroupEntity instantiates a new ProcessGroupEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessGroupEntityWithDefaults

`func NewProcessGroupEntityWithDefaults() *ProcessGroupEntity`

NewProcessGroupEntityWithDefaults instantiates a new ProcessGroupEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevision

`func (o *ProcessGroupEntity) GetRevision() RevisionDTO`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *ProcessGroupEntity) GetRevisionOk() (*RevisionDTO, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *ProcessGroupEntity) SetRevision(v RevisionDTO)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *ProcessGroupEntity) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetId

`func (o *ProcessGroupEntity) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProcessGroupEntity) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProcessGroupEntity) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProcessGroupEntity) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUri

`func (o *ProcessGroupEntity) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *ProcessGroupEntity) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *ProcessGroupEntity) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *ProcessGroupEntity) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetPosition

`func (o *ProcessGroupEntity) GetPosition() PositionDTO`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *ProcessGroupEntity) GetPositionOk() (*PositionDTO, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *ProcessGroupEntity) SetPosition(v PositionDTO)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *ProcessGroupEntity) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetPermissions

`func (o *ProcessGroupEntity) GetPermissions() PermissionsDTO`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *ProcessGroupEntity) GetPermissionsOk() (*PermissionsDTO, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *ProcessGroupEntity) SetPermissions(v PermissionsDTO)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *ProcessGroupEntity) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetBulletins

`func (o *ProcessGroupEntity) GetBulletins() []BulletinEntity`

GetBulletins returns the Bulletins field if non-nil, zero value otherwise.

### GetBulletinsOk

`func (o *ProcessGroupEntity) GetBulletinsOk() (*[]BulletinEntity, bool)`

GetBulletinsOk returns a tuple with the Bulletins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulletins

`func (o *ProcessGroupEntity) SetBulletins(v []BulletinEntity)`

SetBulletins sets Bulletins field to given value.

### HasBulletins

`func (o *ProcessGroupEntity) HasBulletins() bool`

HasBulletins returns a boolean if a field has been set.

### GetDisconnectedNodeAcknowledged

`func (o *ProcessGroupEntity) GetDisconnectedNodeAcknowledged() bool`

GetDisconnectedNodeAcknowledged returns the DisconnectedNodeAcknowledged field if non-nil, zero value otherwise.

### GetDisconnectedNodeAcknowledgedOk

`func (o *ProcessGroupEntity) GetDisconnectedNodeAcknowledgedOk() (*bool, bool)`

GetDisconnectedNodeAcknowledgedOk returns a tuple with the DisconnectedNodeAcknowledged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisconnectedNodeAcknowledged

`func (o *ProcessGroupEntity) SetDisconnectedNodeAcknowledged(v bool)`

SetDisconnectedNodeAcknowledged sets DisconnectedNodeAcknowledged field to given value.

### HasDisconnectedNodeAcknowledged

`func (o *ProcessGroupEntity) HasDisconnectedNodeAcknowledged() bool`

HasDisconnectedNodeAcknowledged returns a boolean if a field has been set.

### GetComponent

`func (o *ProcessGroupEntity) GetComponent() ProcessGroupDTO`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *ProcessGroupEntity) GetComponentOk() (*ProcessGroupDTO, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *ProcessGroupEntity) SetComponent(v ProcessGroupDTO)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *ProcessGroupEntity) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetStatus

`func (o *ProcessGroupEntity) GetStatus() ProcessGroupStatusDTO`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProcessGroupEntity) GetStatusOk() (*ProcessGroupStatusDTO, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProcessGroupEntity) SetStatus(v ProcessGroupStatusDTO)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ProcessGroupEntity) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVersionedFlowSnapshot

`func (o *ProcessGroupEntity) GetVersionedFlowSnapshot() VersionedFlowSnapshot`

GetVersionedFlowSnapshot returns the VersionedFlowSnapshot field if non-nil, zero value otherwise.

### GetVersionedFlowSnapshotOk

`func (o *ProcessGroupEntity) GetVersionedFlowSnapshotOk() (*VersionedFlowSnapshot, bool)`

GetVersionedFlowSnapshotOk returns a tuple with the VersionedFlowSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionedFlowSnapshot

`func (o *ProcessGroupEntity) SetVersionedFlowSnapshot(v VersionedFlowSnapshot)`

SetVersionedFlowSnapshot sets VersionedFlowSnapshot field to given value.

### HasVersionedFlowSnapshot

`func (o *ProcessGroupEntity) HasVersionedFlowSnapshot() bool`

HasVersionedFlowSnapshot returns a boolean if a field has been set.

### GetRunningCount

`func (o *ProcessGroupEntity) GetRunningCount() int32`

GetRunningCount returns the RunningCount field if non-nil, zero value otherwise.

### GetRunningCountOk

`func (o *ProcessGroupEntity) GetRunningCountOk() (*int32, bool)`

GetRunningCountOk returns a tuple with the RunningCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningCount

`func (o *ProcessGroupEntity) SetRunningCount(v int32)`

SetRunningCount sets RunningCount field to given value.

### HasRunningCount

`func (o *ProcessGroupEntity) HasRunningCount() bool`

HasRunningCount returns a boolean if a field has been set.

### GetStoppedCount

`func (o *ProcessGroupEntity) GetStoppedCount() int32`

GetStoppedCount returns the StoppedCount field if non-nil, zero value otherwise.

### GetStoppedCountOk

`func (o *ProcessGroupEntity) GetStoppedCountOk() (*int32, bool)`

GetStoppedCountOk returns a tuple with the StoppedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoppedCount

`func (o *ProcessGroupEntity) SetStoppedCount(v int32)`

SetStoppedCount sets StoppedCount field to given value.

### HasStoppedCount

`func (o *ProcessGroupEntity) HasStoppedCount() bool`

HasStoppedCount returns a boolean if a field has been set.

### GetInvalidCount

`func (o *ProcessGroupEntity) GetInvalidCount() int32`

GetInvalidCount returns the InvalidCount field if non-nil, zero value otherwise.

### GetInvalidCountOk

`func (o *ProcessGroupEntity) GetInvalidCountOk() (*int32, bool)`

GetInvalidCountOk returns a tuple with the InvalidCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidCount

`func (o *ProcessGroupEntity) SetInvalidCount(v int32)`

SetInvalidCount sets InvalidCount field to given value.

### HasInvalidCount

`func (o *ProcessGroupEntity) HasInvalidCount() bool`

HasInvalidCount returns a boolean if a field has been set.

### GetDisabledCount

`func (o *ProcessGroupEntity) GetDisabledCount() int32`

GetDisabledCount returns the DisabledCount field if non-nil, zero value otherwise.

### GetDisabledCountOk

`func (o *ProcessGroupEntity) GetDisabledCountOk() (*int32, bool)`

GetDisabledCountOk returns a tuple with the DisabledCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledCount

`func (o *ProcessGroupEntity) SetDisabledCount(v int32)`

SetDisabledCount sets DisabledCount field to given value.

### HasDisabledCount

`func (o *ProcessGroupEntity) HasDisabledCount() bool`

HasDisabledCount returns a boolean if a field has been set.

### GetActiveRemotePortCount

`func (o *ProcessGroupEntity) GetActiveRemotePortCount() int32`

GetActiveRemotePortCount returns the ActiveRemotePortCount field if non-nil, zero value otherwise.

### GetActiveRemotePortCountOk

`func (o *ProcessGroupEntity) GetActiveRemotePortCountOk() (*int32, bool)`

GetActiveRemotePortCountOk returns a tuple with the ActiveRemotePortCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveRemotePortCount

`func (o *ProcessGroupEntity) SetActiveRemotePortCount(v int32)`

SetActiveRemotePortCount sets ActiveRemotePortCount field to given value.

### HasActiveRemotePortCount

`func (o *ProcessGroupEntity) HasActiveRemotePortCount() bool`

HasActiveRemotePortCount returns a boolean if a field has been set.

### GetInactiveRemotePortCount

`func (o *ProcessGroupEntity) GetInactiveRemotePortCount() int32`

GetInactiveRemotePortCount returns the InactiveRemotePortCount field if non-nil, zero value otherwise.

### GetInactiveRemotePortCountOk

`func (o *ProcessGroupEntity) GetInactiveRemotePortCountOk() (*int32, bool)`

GetInactiveRemotePortCountOk returns a tuple with the InactiveRemotePortCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveRemotePortCount

`func (o *ProcessGroupEntity) SetInactiveRemotePortCount(v int32)`

SetInactiveRemotePortCount sets InactiveRemotePortCount field to given value.

### HasInactiveRemotePortCount

`func (o *ProcessGroupEntity) HasInactiveRemotePortCount() bool`

HasInactiveRemotePortCount returns a boolean if a field has been set.

### GetVersionedFlowState

`func (o *ProcessGroupEntity) GetVersionedFlowState() string`

GetVersionedFlowState returns the VersionedFlowState field if non-nil, zero value otherwise.

### GetVersionedFlowStateOk

`func (o *ProcessGroupEntity) GetVersionedFlowStateOk() (*string, bool)`

GetVersionedFlowStateOk returns a tuple with the VersionedFlowState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionedFlowState

`func (o *ProcessGroupEntity) SetVersionedFlowState(v string)`

SetVersionedFlowState sets VersionedFlowState field to given value.

### HasVersionedFlowState

`func (o *ProcessGroupEntity) HasVersionedFlowState() bool`

HasVersionedFlowState returns a boolean if a field has been set.

### GetUpToDateCount

`func (o *ProcessGroupEntity) GetUpToDateCount() int32`

GetUpToDateCount returns the UpToDateCount field if non-nil, zero value otherwise.

### GetUpToDateCountOk

`func (o *ProcessGroupEntity) GetUpToDateCountOk() (*int32, bool)`

GetUpToDateCountOk returns a tuple with the UpToDateCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpToDateCount

`func (o *ProcessGroupEntity) SetUpToDateCount(v int32)`

SetUpToDateCount sets UpToDateCount field to given value.

### HasUpToDateCount

`func (o *ProcessGroupEntity) HasUpToDateCount() bool`

HasUpToDateCount returns a boolean if a field has been set.

### GetLocallyModifiedCount

`func (o *ProcessGroupEntity) GetLocallyModifiedCount() int32`

GetLocallyModifiedCount returns the LocallyModifiedCount field if non-nil, zero value otherwise.

### GetLocallyModifiedCountOk

`func (o *ProcessGroupEntity) GetLocallyModifiedCountOk() (*int32, bool)`

GetLocallyModifiedCountOk returns a tuple with the LocallyModifiedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocallyModifiedCount

`func (o *ProcessGroupEntity) SetLocallyModifiedCount(v int32)`

SetLocallyModifiedCount sets LocallyModifiedCount field to given value.

### HasLocallyModifiedCount

`func (o *ProcessGroupEntity) HasLocallyModifiedCount() bool`

HasLocallyModifiedCount returns a boolean if a field has been set.

### GetStaleCount

`func (o *ProcessGroupEntity) GetStaleCount() int32`

GetStaleCount returns the StaleCount field if non-nil, zero value otherwise.

### GetStaleCountOk

`func (o *ProcessGroupEntity) GetStaleCountOk() (*int32, bool)`

GetStaleCountOk returns a tuple with the StaleCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaleCount

`func (o *ProcessGroupEntity) SetStaleCount(v int32)`

SetStaleCount sets StaleCount field to given value.

### HasStaleCount

`func (o *ProcessGroupEntity) HasStaleCount() bool`

HasStaleCount returns a boolean if a field has been set.

### GetLocallyModifiedAndStaleCount

`func (o *ProcessGroupEntity) GetLocallyModifiedAndStaleCount() int32`

GetLocallyModifiedAndStaleCount returns the LocallyModifiedAndStaleCount field if non-nil, zero value otherwise.

### GetLocallyModifiedAndStaleCountOk

`func (o *ProcessGroupEntity) GetLocallyModifiedAndStaleCountOk() (*int32, bool)`

GetLocallyModifiedAndStaleCountOk returns a tuple with the LocallyModifiedAndStaleCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocallyModifiedAndStaleCount

`func (o *ProcessGroupEntity) SetLocallyModifiedAndStaleCount(v int32)`

SetLocallyModifiedAndStaleCount sets LocallyModifiedAndStaleCount field to given value.

### HasLocallyModifiedAndStaleCount

`func (o *ProcessGroupEntity) HasLocallyModifiedAndStaleCount() bool`

HasLocallyModifiedAndStaleCount returns a boolean if a field has been set.

### GetSyncFailureCount

`func (o *ProcessGroupEntity) GetSyncFailureCount() int32`

GetSyncFailureCount returns the SyncFailureCount field if non-nil, zero value otherwise.

### GetSyncFailureCountOk

`func (o *ProcessGroupEntity) GetSyncFailureCountOk() (*int32, bool)`

GetSyncFailureCountOk returns a tuple with the SyncFailureCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncFailureCount

`func (o *ProcessGroupEntity) SetSyncFailureCount(v int32)`

SetSyncFailureCount sets SyncFailureCount field to given value.

### HasSyncFailureCount

`func (o *ProcessGroupEntity) HasSyncFailureCount() bool`

HasSyncFailureCount returns a boolean if a field has been set.

### GetLocalInputPortCount

`func (o *ProcessGroupEntity) GetLocalInputPortCount() int32`

GetLocalInputPortCount returns the LocalInputPortCount field if non-nil, zero value otherwise.

### GetLocalInputPortCountOk

`func (o *ProcessGroupEntity) GetLocalInputPortCountOk() (*int32, bool)`

GetLocalInputPortCountOk returns a tuple with the LocalInputPortCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalInputPortCount

`func (o *ProcessGroupEntity) SetLocalInputPortCount(v int32)`

SetLocalInputPortCount sets LocalInputPortCount field to given value.

### HasLocalInputPortCount

`func (o *ProcessGroupEntity) HasLocalInputPortCount() bool`

HasLocalInputPortCount returns a boolean if a field has been set.

### GetLocalOutputPortCount

`func (o *ProcessGroupEntity) GetLocalOutputPortCount() int32`

GetLocalOutputPortCount returns the LocalOutputPortCount field if non-nil, zero value otherwise.

### GetLocalOutputPortCountOk

`func (o *ProcessGroupEntity) GetLocalOutputPortCountOk() (*int32, bool)`

GetLocalOutputPortCountOk returns a tuple with the LocalOutputPortCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalOutputPortCount

`func (o *ProcessGroupEntity) SetLocalOutputPortCount(v int32)`

SetLocalOutputPortCount sets LocalOutputPortCount field to given value.

### HasLocalOutputPortCount

`func (o *ProcessGroupEntity) HasLocalOutputPortCount() bool`

HasLocalOutputPortCount returns a boolean if a field has been set.

### GetPublicInputPortCount

`func (o *ProcessGroupEntity) GetPublicInputPortCount() int32`

GetPublicInputPortCount returns the PublicInputPortCount field if non-nil, zero value otherwise.

### GetPublicInputPortCountOk

`func (o *ProcessGroupEntity) GetPublicInputPortCountOk() (*int32, bool)`

GetPublicInputPortCountOk returns a tuple with the PublicInputPortCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicInputPortCount

`func (o *ProcessGroupEntity) SetPublicInputPortCount(v int32)`

SetPublicInputPortCount sets PublicInputPortCount field to given value.

### HasPublicInputPortCount

`func (o *ProcessGroupEntity) HasPublicInputPortCount() bool`

HasPublicInputPortCount returns a boolean if a field has been set.

### GetPublicOutputPortCount

`func (o *ProcessGroupEntity) GetPublicOutputPortCount() int32`

GetPublicOutputPortCount returns the PublicOutputPortCount field if non-nil, zero value otherwise.

### GetPublicOutputPortCountOk

`func (o *ProcessGroupEntity) GetPublicOutputPortCountOk() (*int32, bool)`

GetPublicOutputPortCountOk returns a tuple with the PublicOutputPortCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicOutputPortCount

`func (o *ProcessGroupEntity) SetPublicOutputPortCount(v int32)`

SetPublicOutputPortCount sets PublicOutputPortCount field to given value.

### HasPublicOutputPortCount

`func (o *ProcessGroupEntity) HasPublicOutputPortCount() bool`

HasPublicOutputPortCount returns a boolean if a field has been set.

### GetParameterContext

`func (o *ProcessGroupEntity) GetParameterContext() ParameterContextReferenceEntity`

GetParameterContext returns the ParameterContext field if non-nil, zero value otherwise.

### GetParameterContextOk

`func (o *ProcessGroupEntity) GetParameterContextOk() (*ParameterContextReferenceEntity, bool)`

GetParameterContextOk returns a tuple with the ParameterContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameterContext

`func (o *ProcessGroupEntity) SetParameterContext(v ParameterContextReferenceEntity)`

SetParameterContext sets ParameterContext field to given value.

### HasParameterContext

`func (o *ProcessGroupEntity) HasParameterContext() bool`

HasParameterContext returns a boolean if a field has been set.

### GetInputPortCount

`func (o *ProcessGroupEntity) GetInputPortCount() int32`

GetInputPortCount returns the InputPortCount field if non-nil, zero value otherwise.

### GetInputPortCountOk

`func (o *ProcessGroupEntity) GetInputPortCountOk() (*int32, bool)`

GetInputPortCountOk returns a tuple with the InputPortCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputPortCount

`func (o *ProcessGroupEntity) SetInputPortCount(v int32)`

SetInputPortCount sets InputPortCount field to given value.

### HasInputPortCount

`func (o *ProcessGroupEntity) HasInputPortCount() bool`

HasInputPortCount returns a boolean if a field has been set.

### GetOutputPortCount

`func (o *ProcessGroupEntity) GetOutputPortCount() int32`

GetOutputPortCount returns the OutputPortCount field if non-nil, zero value otherwise.

### GetOutputPortCountOk

`func (o *ProcessGroupEntity) GetOutputPortCountOk() (*int32, bool)`

GetOutputPortCountOk returns a tuple with the OutputPortCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputPortCount

`func (o *ProcessGroupEntity) SetOutputPortCount(v int32)`

SetOutputPortCount sets OutputPortCount field to given value.

### HasOutputPortCount

`func (o *ProcessGroupEntity) HasOutputPortCount() bool`

HasOutputPortCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


