/*
 * NiFi Rest Api
 *
 * The Rest Api provides programmatic access to command and control a NiFi instance in real time. Start and                                              stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.13.2
 * Contact: dev@nifi.apache.org
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nifi

import (
	"encoding/json"
)

// ProcessGroupEntity struct for ProcessGroupEntity
type ProcessGroupEntity struct {
	Revision *RevisionDTO `json:"revision,omitempty"`
	// The id of the component.
	Id *string `json:"id,omitempty"`
	// The URI for futures requests to the component.
	Uri         *string         `json:"uri,omitempty"`
	Position    *PositionDTO    `json:"position,omitempty"`
	Permissions *PermissionsDTO `json:"permissions,omitempty"`
	// The bulletins for this component.
	Bulletins *[]BulletinEntity `json:"bulletins,omitempty"`
	// Acknowledges that this node is disconnected to allow for mutable requests to proceed.
	DisconnectedNodeAcknowledged *bool                  `json:"disconnectedNodeAcknowledged,omitempty"`
	Component                    *ProcessGroupDTO       `json:"component,omitempty"`
	Status                       *ProcessGroupStatusDTO `json:"status,omitempty"`
	VersionedFlowSnapshot        *VersionedFlowSnapshot `json:"versionedFlowSnapshot,omitempty"`
	// The number of running components in this process group.
	RunningCount *int32 `json:"runningCount,omitempty"`
	// The number of stopped components in the process group.
	StoppedCount *int32 `json:"stoppedCount,omitempty"`
	// The number of invalid components in the process group.
	InvalidCount *int32 `json:"invalidCount,omitempty"`
	// The number of disabled components in the process group.
	DisabledCount *int32 `json:"disabledCount,omitempty"`
	// The number of active remote ports in the process group.
	ActiveRemotePortCount *int32 `json:"activeRemotePortCount,omitempty"`
	// The number of inactive remote ports in the process group.
	InactiveRemotePortCount *int32 `json:"inactiveRemotePortCount,omitempty"`
	// The current state of the Process Group, as it relates to the Versioned Flow
	VersionedFlowState *string `json:"versionedFlowState,omitempty"`
	// The number of up to date versioned process groups in the process group.
	UpToDateCount *int32 `json:"upToDateCount,omitempty"`
	// The number of locally modified versioned process groups in the process group.
	LocallyModifiedCount *int32 `json:"locallyModifiedCount,omitempty"`
	// The number of stale versioned process groups in the process group.
	StaleCount *int32 `json:"staleCount,omitempty"`
	// The number of locally modified and stale versioned process groups in the process group.
	LocallyModifiedAndStaleCount *int32 `json:"locallyModifiedAndStaleCount,omitempty"`
	// The number of versioned process groups in the process group that are unable to sync to a registry.
	SyncFailureCount *int32 `json:"syncFailureCount,omitempty"`
	// The number of local input ports in the process group.
	LocalInputPortCount *int32 `json:"localInputPortCount,omitempty"`
	// The number of local output ports in the process group.
	LocalOutputPortCount *int32 `json:"localOutputPortCount,omitempty"`
	// The number of public input ports in the process group.
	PublicInputPortCount *int32 `json:"publicInputPortCount,omitempty"`
	// The number of public output ports in the process group.
	PublicOutputPortCount *int32                           `json:"publicOutputPortCount,omitempty"`
	ParameterContext      *ParameterContextReferenceEntity `json:"parameterContext,omitempty"`
	// The number of input ports in the process group.
	InputPortCount *int32 `json:"inputPortCount,omitempty"`
	// The number of output ports in the process group.
	OutputPortCount *int32 `json:"outputPortCount,omitempty"`
}

// NewProcessGroupEntity instantiates a new ProcessGroupEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProcessGroupEntity() *ProcessGroupEntity {
	this := ProcessGroupEntity{}
	return &this
}

// NewProcessGroupEntityWithDefaults instantiates a new ProcessGroupEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProcessGroupEntityWithDefaults() *ProcessGroupEntity {
	this := ProcessGroupEntity{}
	return &this
}

// GetRevision returns the Revision field value if set, zero value otherwise.
func (o *ProcessGroupEntity) GetRevision() RevisionDTO {
	if o == nil || o.Revision == nil {
		var ret RevisionDTO
		return ret
	}
	return *o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessGroupEntity) GetRevisionOk() (*RevisionDTO, bool) {
	if o == nil || o.Revision == nil {
		return nil, false
	}
	return o.Revision, true
}

// HasRevision returns a boolean if a field has been set.
func (o *ProcessGroupEntity) HasRevision() bool {
	if o != nil && o.Revision != nil {
		return true
	}

	return false
}

// SetRevision gets a reference to the given RevisionDTO and assigns it to the Revision field.
func (o *ProcessGroupEntity) SetRevision(v RevisionDTO) {
	o.Revision = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ProcessGroupEntity) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessGroupEntity) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ProcessGroupEntity) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ProcessGroupEntity) SetId(v string) {
	o.Id = &v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *ProcessGroupEntity) GetUri() string {
	if o == nil || o.Uri == nil {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessGroupEntity) GetUriOk() (*string, bool) {
	if o == nil || o.Uri == nil {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *ProcessGroupEntity) HasUri() bool {
	if o != nil && o.Uri != nil {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *ProcessGroupEntity) SetUri(v string) {
	o.Uri = &v
}

// GetPosition returns the Position field value if set, zero value otherwise.
func (o *ProcessGroupEntity) GetPosition() PositionDTO {
	if o == nil || o.Position == nil {
		var ret PositionDTO
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessGroupEntity) GetPositionOk() (*PositionDTO, bool) {
	if o == nil || o.Position == nil {
		return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *ProcessGroupEntity) HasPosition() bool {
	if o != nil && o.Position != nil {
		return true
	}

	return false
}

// SetPosition gets a reference to the given PositionDTO and assigns it to the Position field.
func (o *ProcessGroupEntity) SetPosition(v PositionDTO) {
	o.Position = &v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *ProcessGroupEntity) GetPermissions() PermissionsDTO {
	if o == nil || o.Permissions == nil {
		var ret PermissionsDTO
		return ret
	}
	return *o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessGroupEntity) GetPermissionsOk() (*PermissionsDTO, bool) {
	if o == nil || o.Permissions == nil {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *ProcessGroupEntity) HasPermissions() bool {
	if o != nil && o.Permissions != nil {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given PermissionsDTO and assigns it to the Permissions field.
func (o *ProcessGroupEntity) SetPermissions(v PermissionsDTO) {
	o.Permissions = &v
}

// GetBulletins returns the Bulletins field value if set, zero value otherwise.
func (o *ProcessGroupEntity) GetBulletins() []BulletinEntity {
	if o == nil || o.Bulletins == nil {
		var ret []BulletinEntity
		return ret
	}
	return *o.Bulletins
}

// GetBulletinsOk returns a tuple with the Bulletins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessGroupEntity) GetBulletinsOk() (*[]BulletinEntity, bool) {
	if o == nil || o.Bulletins == nil {
		return nil, false
	}
	return o.Bulletins, true
}

// HasBulletins returns a boolean if a field has been set.
func (o *ProcessGroupEntity) HasBulletins() bool {
	if o != nil && o.Bulletins != nil {
		return true
	}

	return false
}

// SetBulletins gets a reference to the given []BulletinEntity and assigns it to the Bulletins field.
func (o *ProcessGroupEntity) SetBulletins(v []BulletinEntity) {
	o.Bulletins = &v
}

// GetDisconnectedNodeAcknowledged returns the DisconnectedNodeAcknowledged field value if set, zero value otherwise.
func (o *ProcessGroupEntity) GetDisconnectedNodeAcknowledged() bool {
	if o == nil || o.DisconnectedNodeAcknowledged == nil {
		var ret bool
		return ret
	}
	return *o.DisconnectedNodeAcknowledged
}

// GetDisconnectedNodeAcknowledgedOk returns a tuple with the DisconnectedNodeAcknowledged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessGroupEntity) GetDisconnectedNodeAcknowledgedOk() (*bool, bool) {
	if o == nil || o.DisconnectedNodeAcknowledged == nil {
		return nil, false
	}
	return o.DisconnectedNodeAcknowledged, true
}

// HasDisconnectedNodeAcknowledged returns a boolean if a field has been set.
func (o *ProcessGroupEntity) HasDisconnectedNodeAcknowledged() bool {
	if o != nil && o.DisconnectedNodeAcknowledged != nil {
		return true
	}

	return false
}

// SetDisconnectedNodeAcknowledged gets a reference to the given bool and assigns it to the DisconnectedNodeAcknowledged field.
func (o *ProcessGroupEntity) SetDisconnectedNodeAcknowledged(v bool) {
	o.DisconnectedNodeAcknowledged = &v
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *ProcessGroupEntity) GetComponent() ProcessGroupDTO {
	if o == nil || o.Component == nil {
		var ret ProcessGroupDTO
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessGroupEntity) GetComponentOk() (*ProcessGroupDTO, bool) {
	if o == nil || o.Component == nil {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *ProcessGroupEntity) HasComponent() bool {
	if o != nil && o.Component != nil {
		return true
	}

	return false
}

// SetComponent gets a reference to the given ProcessGroupDTO and assigns it to the Component field.
func (o *ProcessGroupEntity) SetComponent(v ProcessGroupDTO) {
	o.Component = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ProcessGroupEntity) GetStatus() ProcessGroupStatusDTO {
	if o == nil || o.Status == nil {
		var ret ProcessGroupStatusDTO
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessGroupEntity) GetStatusOk() (*ProcessGroupStatusDTO, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ProcessGroupEntity) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given ProcessGroupStatusDTO and assigns it to the Status field.
func (o *ProcessGroupEntity) SetStatus(v ProcessGroupStatusDTO) {
	o.Status = &v
}

// GetVersionedFlowSnapshot returns the VersionedFlowSnapshot field value if set, zero value otherwise.
func (o *ProcessGroupEntity) GetVersionedFlowSnapshot() VersionedFlowSnapshot {
	if o == nil || o.VersionedFlowSnapshot == nil {
		var ret VersionedFlowSnapshot
		return ret
	}
	return *o.VersionedFlowSnapshot
}

// GetVersionedFlowSnapshotOk returns a tuple with the VersionedFlowSnapshot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessGroupEntity) GetVersionedFlowSnapshotOk() (*VersionedFlowSnapshot, bool) {
	if o == nil || o.VersionedFlowSnapshot == nil {
		return nil, false
	}
	return o.VersionedFlowSnapshot, true
}

// HasVersionedFlowSnapshot returns a boolean if a field has been set.
func (o *ProcessGroupEntity) HasVersionedFlowSnapshot() bool {
	if o != nil && o.VersionedFlowSnapshot != nil {
		return true
	}

	return false
}

// SetVersionedFlowSnapshot gets a reference to the given VersionedFlowSnapshot and assigns it to the VersionedFlowSnapshot field.
func (o *ProcessGroupEntity) SetVersionedFlowSnapshot(v VersionedFlowSnapshot) {
	o.VersionedFlowSnapshot = &v
}

// GetRunningCount returns the RunningCount field value if set, zero value otherwise.
func (o *ProcessGroupEntity) GetRunningCount() int32 {
	if o == nil || o.RunningCount == nil {
		var ret int32
		return ret
	}
	return *o.RunningCount
}

// GetRunningCountOk returns a tuple with the RunningCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessGroupEntity) GetRunningCountOk() (*int32, bool) {
	if o == nil || o.RunningCount == nil {
		return nil, false
	}
	return o.RunningCount, true
}

// HasRunningCount returns a boolean if a field has been set.
func (o *ProcessGroupEntity) HasRunningCount() bool {
	if o != nil && o.RunningCount != nil {
		return true
	}

	return false
}

// SetRunningCount gets a reference to the given int32 and assigns it to the RunningCount field.
func (o *ProcessGroupEntity) SetRunningCount(v int32) {
	o.RunningCount = &v
}

// GetStoppedCount returns the StoppedCount field value if set, zero value otherwise.
func (o *ProcessGroupEntity) GetStoppedCount() int32 {
	if o == nil || o.StoppedCount == nil {
		var ret int32
		return ret
	}
	return *o.StoppedCount
}

// GetStoppedCountOk returns a tuple with the StoppedCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessGroupEntity) GetStoppedCountOk() (*int32, bool) {
	if o == nil || o.StoppedCount == nil {
		return nil, false
	}
	return o.StoppedCount, true
}

// HasStoppedCount returns a boolean if a field has been set.
func (o *ProcessGroupEntity) HasStoppedCount() bool {
	if o != nil && o.StoppedCount != nil {
		return true
	}

	return false
}

// SetStoppedCount gets a reference to the given int32 and assigns it to the StoppedCount field.
func (o *ProcessGroupEntity) SetStoppedCount(v int32) {
	o.StoppedCount = &v
}

// GetInvalidCount returns the InvalidCount field value if set, zero value otherwise.
func (o *ProcessGroupEntity) GetInvalidCount() int32 {
	if o == nil || o.InvalidCount == nil {
		var ret int32
		return ret
	}
	return *o.InvalidCount
}

// GetInvalidCountOk returns a tuple with the InvalidCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessGroupEntity) GetInvalidCountOk() (*int32, bool) {
	if o == nil || o.InvalidCount == nil {
		return nil, false
	}
	return o.InvalidCount, true
}

// HasInvalidCount returns a boolean if a field has been set.
func (o *ProcessGroupEntity) HasInvalidCount() bool {
	if o != nil && o.InvalidCount != nil {
		return true
	}

	return false
}

// SetInvalidCount gets a reference to the given int32 and assigns it to the InvalidCount field.
func (o *ProcessGroupEntity) SetInvalidCount(v int32) {
	o.InvalidCount = &v
}

// GetDisabledCount returns the DisabledCount field value if set, zero value otherwise.
func (o *ProcessGroupEntity) GetDisabledCount() int32 {
	if o == nil || o.DisabledCount == nil {
		var ret int32
		return ret
	}
	return *o.DisabledCount
}

// GetDisabledCountOk returns a tuple with the DisabledCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessGroupEntity) GetDisabledCountOk() (*int32, bool) {
	if o == nil || o.DisabledCount == nil {
		return nil, false
	}
	return o.DisabledCount, true
}

// HasDisabledCount returns a boolean if a field has been set.
func (o *ProcessGroupEntity) HasDisabledCount() bool {
	if o != nil && o.DisabledCount != nil {
		return true
	}

	return false
}

// SetDisabledCount gets a reference to the given int32 and assigns it to the DisabledCount field.
func (o *ProcessGroupEntity) SetDisabledCount(v int32) {
	o.DisabledCount = &v
}

// GetActiveRemotePortCount returns the ActiveRemotePortCount field value if set, zero value otherwise.
func (o *ProcessGroupEntity) GetActiveRemotePortCount() int32 {
	if o == nil || o.ActiveRemotePortCount == nil {
		var ret int32
		return ret
	}
	return *o.ActiveRemotePortCount
}

// GetActiveRemotePortCountOk returns a tuple with the ActiveRemotePortCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessGroupEntity) GetActiveRemotePortCountOk() (*int32, bool) {
	if o == nil || o.ActiveRemotePortCount == nil {
		return nil, false
	}
	return o.ActiveRemotePortCount, true
}

// HasActiveRemotePortCount returns a boolean if a field has been set.
func (o *ProcessGroupEntity) HasActiveRemotePortCount() bool {
	if o != nil && o.ActiveRemotePortCount != nil {
		return true
	}

	return false
}

// SetActiveRemotePortCount gets a reference to the given int32 and assigns it to the ActiveRemotePortCount field.
func (o *ProcessGroupEntity) SetActiveRemotePortCount(v int32) {
	o.ActiveRemotePortCount = &v
}

// GetInactiveRemotePortCount returns the InactiveRemotePortCount field value if set, zero value otherwise.
func (o *ProcessGroupEntity) GetInactiveRemotePortCount() int32 {
	if o == nil || o.InactiveRemotePortCount == nil {
		var ret int32
		return ret
	}
	return *o.InactiveRemotePortCount
}

// GetInactiveRemotePortCountOk returns a tuple with the InactiveRemotePortCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessGroupEntity) GetInactiveRemotePortCountOk() (*int32, bool) {
	if o == nil || o.InactiveRemotePortCount == nil {
		return nil, false
	}
	return o.InactiveRemotePortCount, true
}

// HasInactiveRemotePortCount returns a boolean if a field has been set.
func (o *ProcessGroupEntity) HasInactiveRemotePortCount() bool {
	if o != nil && o.InactiveRemotePortCount != nil {
		return true
	}

	return false
}

// SetInactiveRemotePortCount gets a reference to the given int32 and assigns it to the InactiveRemotePortCount field.
func (o *ProcessGroupEntity) SetInactiveRemotePortCount(v int32) {
	o.InactiveRemotePortCount = &v
}

// GetVersionedFlowState returns the VersionedFlowState field value if set, zero value otherwise.
func (o *ProcessGroupEntity) GetVersionedFlowState() string {
	if o == nil || o.VersionedFlowState == nil {
		var ret string
		return ret
	}
	return *o.VersionedFlowState
}

// GetVersionedFlowStateOk returns a tuple with the VersionedFlowState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessGroupEntity) GetVersionedFlowStateOk() (*string, bool) {
	if o == nil || o.VersionedFlowState == nil {
		return nil, false
	}
	return o.VersionedFlowState, true
}

// HasVersionedFlowState returns a boolean if a field has been set.
func (o *ProcessGroupEntity) HasVersionedFlowState() bool {
	if o != nil && o.VersionedFlowState != nil {
		return true
	}

	return false
}

// SetVersionedFlowState gets a reference to the given string and assigns it to the VersionedFlowState field.
func (o *ProcessGroupEntity) SetVersionedFlowState(v string) {
	o.VersionedFlowState = &v
}

// GetUpToDateCount returns the UpToDateCount field value if set, zero value otherwise.
func (o *ProcessGroupEntity) GetUpToDateCount() int32 {
	if o == nil || o.UpToDateCount == nil {
		var ret int32
		return ret
	}
	return *o.UpToDateCount
}

// GetUpToDateCountOk returns a tuple with the UpToDateCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessGroupEntity) GetUpToDateCountOk() (*int32, bool) {
	if o == nil || o.UpToDateCount == nil {
		return nil, false
	}
	return o.UpToDateCount, true
}

// HasUpToDateCount returns a boolean if a field has been set.
func (o *ProcessGroupEntity) HasUpToDateCount() bool {
	if o != nil && o.UpToDateCount != nil {
		return true
	}

	return false
}

// SetUpToDateCount gets a reference to the given int32 and assigns it to the UpToDateCount field.
func (o *ProcessGroupEntity) SetUpToDateCount(v int32) {
	o.UpToDateCount = &v
}

// GetLocallyModifiedCount returns the LocallyModifiedCount field value if set, zero value otherwise.
func (o *ProcessGroupEntity) GetLocallyModifiedCount() int32 {
	if o == nil || o.LocallyModifiedCount == nil {
		var ret int32
		return ret
	}
	return *o.LocallyModifiedCount
}

// GetLocallyModifiedCountOk returns a tuple with the LocallyModifiedCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessGroupEntity) GetLocallyModifiedCountOk() (*int32, bool) {
	if o == nil || o.LocallyModifiedCount == nil {
		return nil, false
	}
	return o.LocallyModifiedCount, true
}

// HasLocallyModifiedCount returns a boolean if a field has been set.
func (o *ProcessGroupEntity) HasLocallyModifiedCount() bool {
	if o != nil && o.LocallyModifiedCount != nil {
		return true
	}

	return false
}

// SetLocallyModifiedCount gets a reference to the given int32 and assigns it to the LocallyModifiedCount field.
func (o *ProcessGroupEntity) SetLocallyModifiedCount(v int32) {
	o.LocallyModifiedCount = &v
}

// GetStaleCount returns the StaleCount field value if set, zero value otherwise.
func (o *ProcessGroupEntity) GetStaleCount() int32 {
	if o == nil || o.StaleCount == nil {
		var ret int32
		return ret
	}
	return *o.StaleCount
}

// GetStaleCountOk returns a tuple with the StaleCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessGroupEntity) GetStaleCountOk() (*int32, bool) {
	if o == nil || o.StaleCount == nil {
		return nil, false
	}
	return o.StaleCount, true
}

// HasStaleCount returns a boolean if a field has been set.
func (o *ProcessGroupEntity) HasStaleCount() bool {
	if o != nil && o.StaleCount != nil {
		return true
	}

	return false
}

// SetStaleCount gets a reference to the given int32 and assigns it to the StaleCount field.
func (o *ProcessGroupEntity) SetStaleCount(v int32) {
	o.StaleCount = &v
}

// GetLocallyModifiedAndStaleCount returns the LocallyModifiedAndStaleCount field value if set, zero value otherwise.
func (o *ProcessGroupEntity) GetLocallyModifiedAndStaleCount() int32 {
	if o == nil || o.LocallyModifiedAndStaleCount == nil {
		var ret int32
		return ret
	}
	return *o.LocallyModifiedAndStaleCount
}

// GetLocallyModifiedAndStaleCountOk returns a tuple with the LocallyModifiedAndStaleCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessGroupEntity) GetLocallyModifiedAndStaleCountOk() (*int32, bool) {
	if o == nil || o.LocallyModifiedAndStaleCount == nil {
		return nil, false
	}
	return o.LocallyModifiedAndStaleCount, true
}

// HasLocallyModifiedAndStaleCount returns a boolean if a field has been set.
func (o *ProcessGroupEntity) HasLocallyModifiedAndStaleCount() bool {
	if o != nil && o.LocallyModifiedAndStaleCount != nil {
		return true
	}

	return false
}

// SetLocallyModifiedAndStaleCount gets a reference to the given int32 and assigns it to the LocallyModifiedAndStaleCount field.
func (o *ProcessGroupEntity) SetLocallyModifiedAndStaleCount(v int32) {
	o.LocallyModifiedAndStaleCount = &v
}

// GetSyncFailureCount returns the SyncFailureCount field value if set, zero value otherwise.
func (o *ProcessGroupEntity) GetSyncFailureCount() int32 {
	if o == nil || o.SyncFailureCount == nil {
		var ret int32
		return ret
	}
	return *o.SyncFailureCount
}

// GetSyncFailureCountOk returns a tuple with the SyncFailureCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessGroupEntity) GetSyncFailureCountOk() (*int32, bool) {
	if o == nil || o.SyncFailureCount == nil {
		return nil, false
	}
	return o.SyncFailureCount, true
}

// HasSyncFailureCount returns a boolean if a field has been set.
func (o *ProcessGroupEntity) HasSyncFailureCount() bool {
	if o != nil && o.SyncFailureCount != nil {
		return true
	}

	return false
}

// SetSyncFailureCount gets a reference to the given int32 and assigns it to the SyncFailureCount field.
func (o *ProcessGroupEntity) SetSyncFailureCount(v int32) {
	o.SyncFailureCount = &v
}

// GetLocalInputPortCount returns the LocalInputPortCount field value if set, zero value otherwise.
func (o *ProcessGroupEntity) GetLocalInputPortCount() int32 {
	if o == nil || o.LocalInputPortCount == nil {
		var ret int32
		return ret
	}
	return *o.LocalInputPortCount
}

// GetLocalInputPortCountOk returns a tuple with the LocalInputPortCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessGroupEntity) GetLocalInputPortCountOk() (*int32, bool) {
	if o == nil || o.LocalInputPortCount == nil {
		return nil, false
	}
	return o.LocalInputPortCount, true
}

// HasLocalInputPortCount returns a boolean if a field has been set.
func (o *ProcessGroupEntity) HasLocalInputPortCount() bool {
	if o != nil && o.LocalInputPortCount != nil {
		return true
	}

	return false
}

// SetLocalInputPortCount gets a reference to the given int32 and assigns it to the LocalInputPortCount field.
func (o *ProcessGroupEntity) SetLocalInputPortCount(v int32) {
	o.LocalInputPortCount = &v
}

// GetLocalOutputPortCount returns the LocalOutputPortCount field value if set, zero value otherwise.
func (o *ProcessGroupEntity) GetLocalOutputPortCount() int32 {
	if o == nil || o.LocalOutputPortCount == nil {
		var ret int32
		return ret
	}
	return *o.LocalOutputPortCount
}

// GetLocalOutputPortCountOk returns a tuple with the LocalOutputPortCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessGroupEntity) GetLocalOutputPortCountOk() (*int32, bool) {
	if o == nil || o.LocalOutputPortCount == nil {
		return nil, false
	}
	return o.LocalOutputPortCount, true
}

// HasLocalOutputPortCount returns a boolean if a field has been set.
func (o *ProcessGroupEntity) HasLocalOutputPortCount() bool {
	if o != nil && o.LocalOutputPortCount != nil {
		return true
	}

	return false
}

// SetLocalOutputPortCount gets a reference to the given int32 and assigns it to the LocalOutputPortCount field.
func (o *ProcessGroupEntity) SetLocalOutputPortCount(v int32) {
	o.LocalOutputPortCount = &v
}

// GetPublicInputPortCount returns the PublicInputPortCount field value if set, zero value otherwise.
func (o *ProcessGroupEntity) GetPublicInputPortCount() int32 {
	if o == nil || o.PublicInputPortCount == nil {
		var ret int32
		return ret
	}
	return *o.PublicInputPortCount
}

// GetPublicInputPortCountOk returns a tuple with the PublicInputPortCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessGroupEntity) GetPublicInputPortCountOk() (*int32, bool) {
	if o == nil || o.PublicInputPortCount == nil {
		return nil, false
	}
	return o.PublicInputPortCount, true
}

// HasPublicInputPortCount returns a boolean if a field has been set.
func (o *ProcessGroupEntity) HasPublicInputPortCount() bool {
	if o != nil && o.PublicInputPortCount != nil {
		return true
	}

	return false
}

// SetPublicInputPortCount gets a reference to the given int32 and assigns it to the PublicInputPortCount field.
func (o *ProcessGroupEntity) SetPublicInputPortCount(v int32) {
	o.PublicInputPortCount = &v
}

// GetPublicOutputPortCount returns the PublicOutputPortCount field value if set, zero value otherwise.
func (o *ProcessGroupEntity) GetPublicOutputPortCount() int32 {
	if o == nil || o.PublicOutputPortCount == nil {
		var ret int32
		return ret
	}
	return *o.PublicOutputPortCount
}

// GetPublicOutputPortCountOk returns a tuple with the PublicOutputPortCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessGroupEntity) GetPublicOutputPortCountOk() (*int32, bool) {
	if o == nil || o.PublicOutputPortCount == nil {
		return nil, false
	}
	return o.PublicOutputPortCount, true
}

// HasPublicOutputPortCount returns a boolean if a field has been set.
func (o *ProcessGroupEntity) HasPublicOutputPortCount() bool {
	if o != nil && o.PublicOutputPortCount != nil {
		return true
	}

	return false
}

// SetPublicOutputPortCount gets a reference to the given int32 and assigns it to the PublicOutputPortCount field.
func (o *ProcessGroupEntity) SetPublicOutputPortCount(v int32) {
	o.PublicOutputPortCount = &v
}

// GetParameterContext returns the ParameterContext field value if set, zero value otherwise.
func (o *ProcessGroupEntity) GetParameterContext() ParameterContextReferenceEntity {
	if o == nil || o.ParameterContext == nil {
		var ret ParameterContextReferenceEntity
		return ret
	}
	return *o.ParameterContext
}

// GetParameterContextOk returns a tuple with the ParameterContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessGroupEntity) GetParameterContextOk() (*ParameterContextReferenceEntity, bool) {
	if o == nil || o.ParameterContext == nil {
		return nil, false
	}
	return o.ParameterContext, true
}

// HasParameterContext returns a boolean if a field has been set.
func (o *ProcessGroupEntity) HasParameterContext() bool {
	if o != nil && o.ParameterContext != nil {
		return true
	}

	return false
}

// SetParameterContext gets a reference to the given ParameterContextReferenceEntity and assigns it to the ParameterContext field.
func (o *ProcessGroupEntity) SetParameterContext(v ParameterContextReferenceEntity) {
	o.ParameterContext = &v
}

// GetInputPortCount returns the InputPortCount field value if set, zero value otherwise.
func (o *ProcessGroupEntity) GetInputPortCount() int32 {
	if o == nil || o.InputPortCount == nil {
		var ret int32
		return ret
	}
	return *o.InputPortCount
}

// GetInputPortCountOk returns a tuple with the InputPortCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessGroupEntity) GetInputPortCountOk() (*int32, bool) {
	if o == nil || o.InputPortCount == nil {
		return nil, false
	}
	return o.InputPortCount, true
}

// HasInputPortCount returns a boolean if a field has been set.
func (o *ProcessGroupEntity) HasInputPortCount() bool {
	if o != nil && o.InputPortCount != nil {
		return true
	}

	return false
}

// SetInputPortCount gets a reference to the given int32 and assigns it to the InputPortCount field.
func (o *ProcessGroupEntity) SetInputPortCount(v int32) {
	o.InputPortCount = &v
}

// GetOutputPortCount returns the OutputPortCount field value if set, zero value otherwise.
func (o *ProcessGroupEntity) GetOutputPortCount() int32 {
	if o == nil || o.OutputPortCount == nil {
		var ret int32
		return ret
	}
	return *o.OutputPortCount
}

// GetOutputPortCountOk returns a tuple with the OutputPortCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessGroupEntity) GetOutputPortCountOk() (*int32, bool) {
	if o == nil || o.OutputPortCount == nil {
		return nil, false
	}
	return o.OutputPortCount, true
}

// HasOutputPortCount returns a boolean if a field has been set.
func (o *ProcessGroupEntity) HasOutputPortCount() bool {
	if o != nil && o.OutputPortCount != nil {
		return true
	}

	return false
}

// SetOutputPortCount gets a reference to the given int32 and assigns it to the OutputPortCount field.
func (o *ProcessGroupEntity) SetOutputPortCount(v int32) {
	o.OutputPortCount = &v
}

func (o ProcessGroupEntity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Revision != nil {
		toSerialize["revision"] = o.Revision
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Uri != nil {
		toSerialize["uri"] = o.Uri
	}
	if o.Position != nil {
		toSerialize["position"] = o.Position
	}
	if o.Permissions != nil {
		toSerialize["permissions"] = o.Permissions
	}
	if o.Bulletins != nil {
		toSerialize["bulletins"] = o.Bulletins
	}
	if o.DisconnectedNodeAcknowledged != nil {
		toSerialize["disconnectedNodeAcknowledged"] = o.DisconnectedNodeAcknowledged
	}
	if o.Component != nil {
		toSerialize["component"] = o.Component
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.VersionedFlowSnapshot != nil {
		toSerialize["versionedFlowSnapshot"] = o.VersionedFlowSnapshot
	}
	if o.RunningCount != nil {
		toSerialize["runningCount"] = o.RunningCount
	}
	if o.StoppedCount != nil {
		toSerialize["stoppedCount"] = o.StoppedCount
	}
	if o.InvalidCount != nil {
		toSerialize["invalidCount"] = o.InvalidCount
	}
	if o.DisabledCount != nil {
		toSerialize["disabledCount"] = o.DisabledCount
	}
	if o.ActiveRemotePortCount != nil {
		toSerialize["activeRemotePortCount"] = o.ActiveRemotePortCount
	}
	if o.InactiveRemotePortCount != nil {
		toSerialize["inactiveRemotePortCount"] = o.InactiveRemotePortCount
	}
	if o.VersionedFlowState != nil {
		toSerialize["versionedFlowState"] = o.VersionedFlowState
	}
	if o.UpToDateCount != nil {
		toSerialize["upToDateCount"] = o.UpToDateCount
	}
	if o.LocallyModifiedCount != nil {
		toSerialize["locallyModifiedCount"] = o.LocallyModifiedCount
	}
	if o.StaleCount != nil {
		toSerialize["staleCount"] = o.StaleCount
	}
	if o.LocallyModifiedAndStaleCount != nil {
		toSerialize["locallyModifiedAndStaleCount"] = o.LocallyModifiedAndStaleCount
	}
	if o.SyncFailureCount != nil {
		toSerialize["syncFailureCount"] = o.SyncFailureCount
	}
	if o.LocalInputPortCount != nil {
		toSerialize["localInputPortCount"] = o.LocalInputPortCount
	}
	if o.LocalOutputPortCount != nil {
		toSerialize["localOutputPortCount"] = o.LocalOutputPortCount
	}
	if o.PublicInputPortCount != nil {
		toSerialize["publicInputPortCount"] = o.PublicInputPortCount
	}
	if o.PublicOutputPortCount != nil {
		toSerialize["publicOutputPortCount"] = o.PublicOutputPortCount
	}
	if o.ParameterContext != nil {
		toSerialize["parameterContext"] = o.ParameterContext
	}
	if o.InputPortCount != nil {
		toSerialize["inputPortCount"] = o.InputPortCount
	}
	if o.OutputPortCount != nil {
		toSerialize["outputPortCount"] = o.OutputPortCount
	}
	return json.Marshal(toSerialize)
}

type NullableProcessGroupEntity struct {
	value *ProcessGroupEntity
	isSet bool
}

func (v NullableProcessGroupEntity) Get() *ProcessGroupEntity {
	return v.value
}

func (v *NullableProcessGroupEntity) Set(val *ProcessGroupEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableProcessGroupEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableProcessGroupEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProcessGroupEntity(val *ProcessGroupEntity) *NullableProcessGroupEntity {
	return &NullableProcessGroupEntity{value: val, isSet: true}
}

func (v NullableProcessGroupEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProcessGroupEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
