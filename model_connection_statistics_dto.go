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

// ConnectionStatisticsDTO struct for ConnectionStatisticsDTO
type ConnectionStatisticsDTO struct {
	// The ID of the connection
	Id *string `json:"id,omitempty"`
	// The timestamp of when the stats were last refreshed
	StatsLastRefreshed *string                          `json:"statsLastRefreshed,omitempty"`
	AggregateSnapshot  *ConnectionStatisticsSnapshotDTO `json:"aggregateSnapshot,omitempty"`
	// A list of status snapshots for each node
	NodeSnapshots *[]NodeConnectionStatisticsSnapshotDTO `json:"nodeSnapshots,omitempty"`
}

// NewConnectionStatisticsDTO instantiates a new ConnectionStatisticsDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectionStatisticsDTO() *ConnectionStatisticsDTO {
	this := ConnectionStatisticsDTO{}
	return &this
}

// NewConnectionStatisticsDTOWithDefaults instantiates a new ConnectionStatisticsDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectionStatisticsDTOWithDefaults() *ConnectionStatisticsDTO {
	this := ConnectionStatisticsDTO{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ConnectionStatisticsDTO) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionStatisticsDTO) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ConnectionStatisticsDTO) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ConnectionStatisticsDTO) SetId(v string) {
	o.Id = &v
}

// GetStatsLastRefreshed returns the StatsLastRefreshed field value if set, zero value otherwise.
func (o *ConnectionStatisticsDTO) GetStatsLastRefreshed() string {
	if o == nil || o.StatsLastRefreshed == nil {
		var ret string
		return ret
	}
	return *o.StatsLastRefreshed
}

// GetStatsLastRefreshedOk returns a tuple with the StatsLastRefreshed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionStatisticsDTO) GetStatsLastRefreshedOk() (*string, bool) {
	if o == nil || o.StatsLastRefreshed == nil {
		return nil, false
	}
	return o.StatsLastRefreshed, true
}

// HasStatsLastRefreshed returns a boolean if a field has been set.
func (o *ConnectionStatisticsDTO) HasStatsLastRefreshed() bool {
	if o != nil && o.StatsLastRefreshed != nil {
		return true
	}

	return false
}

// SetStatsLastRefreshed gets a reference to the given string and assigns it to the StatsLastRefreshed field.
func (o *ConnectionStatisticsDTO) SetStatsLastRefreshed(v string) {
	o.StatsLastRefreshed = &v
}

// GetAggregateSnapshot returns the AggregateSnapshot field value if set, zero value otherwise.
func (o *ConnectionStatisticsDTO) GetAggregateSnapshot() ConnectionStatisticsSnapshotDTO {
	if o == nil || o.AggregateSnapshot == nil {
		var ret ConnectionStatisticsSnapshotDTO
		return ret
	}
	return *o.AggregateSnapshot
}

// GetAggregateSnapshotOk returns a tuple with the AggregateSnapshot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionStatisticsDTO) GetAggregateSnapshotOk() (*ConnectionStatisticsSnapshotDTO, bool) {
	if o == nil || o.AggregateSnapshot == nil {
		return nil, false
	}
	return o.AggregateSnapshot, true
}

// HasAggregateSnapshot returns a boolean if a field has been set.
func (o *ConnectionStatisticsDTO) HasAggregateSnapshot() bool {
	if o != nil && o.AggregateSnapshot != nil {
		return true
	}

	return false
}

// SetAggregateSnapshot gets a reference to the given ConnectionStatisticsSnapshotDTO and assigns it to the AggregateSnapshot field.
func (o *ConnectionStatisticsDTO) SetAggregateSnapshot(v ConnectionStatisticsSnapshotDTO) {
	o.AggregateSnapshot = &v
}

// GetNodeSnapshots returns the NodeSnapshots field value if set, zero value otherwise.
func (o *ConnectionStatisticsDTO) GetNodeSnapshots() []NodeConnectionStatisticsSnapshotDTO {
	if o == nil || o.NodeSnapshots == nil {
		var ret []NodeConnectionStatisticsSnapshotDTO
		return ret
	}
	return *o.NodeSnapshots
}

// GetNodeSnapshotsOk returns a tuple with the NodeSnapshots field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionStatisticsDTO) GetNodeSnapshotsOk() (*[]NodeConnectionStatisticsSnapshotDTO, bool) {
	if o == nil || o.NodeSnapshots == nil {
		return nil, false
	}
	return o.NodeSnapshots, true
}

// HasNodeSnapshots returns a boolean if a field has been set.
func (o *ConnectionStatisticsDTO) HasNodeSnapshots() bool {
	if o != nil && o.NodeSnapshots != nil {
		return true
	}

	return false
}

// SetNodeSnapshots gets a reference to the given []NodeConnectionStatisticsSnapshotDTO and assigns it to the NodeSnapshots field.
func (o *ConnectionStatisticsDTO) SetNodeSnapshots(v []NodeConnectionStatisticsSnapshotDTO) {
	o.NodeSnapshots = &v
}

func (o ConnectionStatisticsDTO) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
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

type NullableConnectionStatisticsDTO struct {
	value *ConnectionStatisticsDTO
	isSet bool
}

func (v NullableConnectionStatisticsDTO) Get() *ConnectionStatisticsDTO {
	return v.value
}

func (v *NullableConnectionStatisticsDTO) Set(val *ConnectionStatisticsDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectionStatisticsDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectionStatisticsDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectionStatisticsDTO(val *ConnectionStatisticsDTO) *NullableConnectionStatisticsDTO {
	return &NullableConnectionStatisticsDTO{value: val, isSet: true}
}

func (v NullableConnectionStatisticsDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectionStatisticsDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
