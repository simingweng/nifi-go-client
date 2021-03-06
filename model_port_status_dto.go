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

// PortStatusDTO struct for PortStatusDTO
type PortStatusDTO struct {
	// The id of the port.
	Id *string `json:"id,omitempty"`
	// The id of the parent process group of the port.
	GroupId *string `json:"groupId,omitempty"`
	// The name of the port.
	Name *string `json:"name,omitempty"`
	// Whether the port has incoming or outgoing connections to a remote NiFi.
	Transmitting *bool `json:"transmitting,omitempty"`
	// The run status of the port.
	RunStatus *string `json:"runStatus,omitempty"`
	// The time the status for the process group was last refreshed.
	StatsLastRefreshed *string                `json:"statsLastRefreshed,omitempty"`
	AggregateSnapshot  *PortStatusSnapshotDTO `json:"aggregateSnapshot,omitempty"`
	// A status snapshot for each node in the cluster. If the NiFi instance is a standalone instance, rather than a cluster, this may be null.
	NodeSnapshots *[]NodePortStatusSnapshotDTO `json:"nodeSnapshots,omitempty"`
}

// NewPortStatusDTO instantiates a new PortStatusDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPortStatusDTO() *PortStatusDTO {
	this := PortStatusDTO{}
	return &this
}

// NewPortStatusDTOWithDefaults instantiates a new PortStatusDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPortStatusDTOWithDefaults() *PortStatusDTO {
	this := PortStatusDTO{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PortStatusDTO) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortStatusDTO) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PortStatusDTO) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PortStatusDTO) SetId(v string) {
	o.Id = &v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *PortStatusDTO) GetGroupId() string {
	if o == nil || o.GroupId == nil {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortStatusDTO) GetGroupIdOk() (*string, bool) {
	if o == nil || o.GroupId == nil {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *PortStatusDTO) HasGroupId() bool {
	if o != nil && o.GroupId != nil {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *PortStatusDTO) SetGroupId(v string) {
	o.GroupId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PortStatusDTO) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortStatusDTO) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PortStatusDTO) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PortStatusDTO) SetName(v string) {
	o.Name = &v
}

// GetTransmitting returns the Transmitting field value if set, zero value otherwise.
func (o *PortStatusDTO) GetTransmitting() bool {
	if o == nil || o.Transmitting == nil {
		var ret bool
		return ret
	}
	return *o.Transmitting
}

// GetTransmittingOk returns a tuple with the Transmitting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortStatusDTO) GetTransmittingOk() (*bool, bool) {
	if o == nil || o.Transmitting == nil {
		return nil, false
	}
	return o.Transmitting, true
}

// HasTransmitting returns a boolean if a field has been set.
func (o *PortStatusDTO) HasTransmitting() bool {
	if o != nil && o.Transmitting != nil {
		return true
	}

	return false
}

// SetTransmitting gets a reference to the given bool and assigns it to the Transmitting field.
func (o *PortStatusDTO) SetTransmitting(v bool) {
	o.Transmitting = &v
}

// GetRunStatus returns the RunStatus field value if set, zero value otherwise.
func (o *PortStatusDTO) GetRunStatus() string {
	if o == nil || o.RunStatus == nil {
		var ret string
		return ret
	}
	return *o.RunStatus
}

// GetRunStatusOk returns a tuple with the RunStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortStatusDTO) GetRunStatusOk() (*string, bool) {
	if o == nil || o.RunStatus == nil {
		return nil, false
	}
	return o.RunStatus, true
}

// HasRunStatus returns a boolean if a field has been set.
func (o *PortStatusDTO) HasRunStatus() bool {
	if o != nil && o.RunStatus != nil {
		return true
	}

	return false
}

// SetRunStatus gets a reference to the given string and assigns it to the RunStatus field.
func (o *PortStatusDTO) SetRunStatus(v string) {
	o.RunStatus = &v
}

// GetStatsLastRefreshed returns the StatsLastRefreshed field value if set, zero value otherwise.
func (o *PortStatusDTO) GetStatsLastRefreshed() string {
	if o == nil || o.StatsLastRefreshed == nil {
		var ret string
		return ret
	}
	return *o.StatsLastRefreshed
}

// GetStatsLastRefreshedOk returns a tuple with the StatsLastRefreshed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortStatusDTO) GetStatsLastRefreshedOk() (*string, bool) {
	if o == nil || o.StatsLastRefreshed == nil {
		return nil, false
	}
	return o.StatsLastRefreshed, true
}

// HasStatsLastRefreshed returns a boolean if a field has been set.
func (o *PortStatusDTO) HasStatsLastRefreshed() bool {
	if o != nil && o.StatsLastRefreshed != nil {
		return true
	}

	return false
}

// SetStatsLastRefreshed gets a reference to the given string and assigns it to the StatsLastRefreshed field.
func (o *PortStatusDTO) SetStatsLastRefreshed(v string) {
	o.StatsLastRefreshed = &v
}

// GetAggregateSnapshot returns the AggregateSnapshot field value if set, zero value otherwise.
func (o *PortStatusDTO) GetAggregateSnapshot() PortStatusSnapshotDTO {
	if o == nil || o.AggregateSnapshot == nil {
		var ret PortStatusSnapshotDTO
		return ret
	}
	return *o.AggregateSnapshot
}

// GetAggregateSnapshotOk returns a tuple with the AggregateSnapshot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortStatusDTO) GetAggregateSnapshotOk() (*PortStatusSnapshotDTO, bool) {
	if o == nil || o.AggregateSnapshot == nil {
		return nil, false
	}
	return o.AggregateSnapshot, true
}

// HasAggregateSnapshot returns a boolean if a field has been set.
func (o *PortStatusDTO) HasAggregateSnapshot() bool {
	if o != nil && o.AggregateSnapshot != nil {
		return true
	}

	return false
}

// SetAggregateSnapshot gets a reference to the given PortStatusSnapshotDTO and assigns it to the AggregateSnapshot field.
func (o *PortStatusDTO) SetAggregateSnapshot(v PortStatusSnapshotDTO) {
	o.AggregateSnapshot = &v
}

// GetNodeSnapshots returns the NodeSnapshots field value if set, zero value otherwise.
func (o *PortStatusDTO) GetNodeSnapshots() []NodePortStatusSnapshotDTO {
	if o == nil || o.NodeSnapshots == nil {
		var ret []NodePortStatusSnapshotDTO
		return ret
	}
	return *o.NodeSnapshots
}

// GetNodeSnapshotsOk returns a tuple with the NodeSnapshots field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortStatusDTO) GetNodeSnapshotsOk() (*[]NodePortStatusSnapshotDTO, bool) {
	if o == nil || o.NodeSnapshots == nil {
		return nil, false
	}
	return o.NodeSnapshots, true
}

// HasNodeSnapshots returns a boolean if a field has been set.
func (o *PortStatusDTO) HasNodeSnapshots() bool {
	if o != nil && o.NodeSnapshots != nil {
		return true
	}

	return false
}

// SetNodeSnapshots gets a reference to the given []NodePortStatusSnapshotDTO and assigns it to the NodeSnapshots field.
func (o *PortStatusDTO) SetNodeSnapshots(v []NodePortStatusSnapshotDTO) {
	o.NodeSnapshots = &v
}

func (o PortStatusDTO) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.GroupId != nil {
		toSerialize["groupId"] = o.GroupId
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Transmitting != nil {
		toSerialize["transmitting"] = o.Transmitting
	}
	if o.RunStatus != nil {
		toSerialize["runStatus"] = o.RunStatus
	}
	if o.StatsLastRefreshed != nil {
		toSerialize["statsLastRefreshed"] = o.StatsLastRefreshed
	}
	if o.AggregateSnapshot != nil {
		toSerialize["aggregateSnapshot"] = o.AggregateSnapshot
	}
	if o.NodeSnapshots != nil {
		toSerialize["nodeSnapshots"] = o.NodeSnapshots
	}
	return json.Marshal(toSerialize)
}

type NullablePortStatusDTO struct {
	value *PortStatusDTO
	isSet bool
}

func (v NullablePortStatusDTO) Get() *PortStatusDTO {
	return v.value
}

func (v *NullablePortStatusDTO) Set(val *PortStatusDTO) {
	v.value = val
	v.isSet = true
}

func (v NullablePortStatusDTO) IsSet() bool {
	return v.isSet
}

func (v *NullablePortStatusDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortStatusDTO(val *PortStatusDTO) *NullablePortStatusDTO {
	return &NullablePortStatusDTO{value: val, isSet: true}
}

func (v NullablePortStatusDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePortStatusDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
