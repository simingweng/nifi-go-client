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

// PortStatusSnapshotEntity struct for PortStatusSnapshotEntity
type PortStatusSnapshotEntity struct {
	// The id of the port.
	Id                 *string                `json:"id,omitempty"`
	PortStatusSnapshot *PortStatusSnapshotDTO `json:"portStatusSnapshot,omitempty"`
	// Indicates whether the user can read a given resource.
	CanRead *bool `json:"canRead,omitempty"`
}

// NewPortStatusSnapshotEntity instantiates a new PortStatusSnapshotEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPortStatusSnapshotEntity() *PortStatusSnapshotEntity {
	this := PortStatusSnapshotEntity{}
	return &this
}

// NewPortStatusSnapshotEntityWithDefaults instantiates a new PortStatusSnapshotEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPortStatusSnapshotEntityWithDefaults() *PortStatusSnapshotEntity {
	this := PortStatusSnapshotEntity{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PortStatusSnapshotEntity) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortStatusSnapshotEntity) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PortStatusSnapshotEntity) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PortStatusSnapshotEntity) SetId(v string) {
	o.Id = &v
}

// GetPortStatusSnapshot returns the PortStatusSnapshot field value if set, zero value otherwise.
func (o *PortStatusSnapshotEntity) GetPortStatusSnapshot() PortStatusSnapshotDTO {
	if o == nil || o.PortStatusSnapshot == nil {
		var ret PortStatusSnapshotDTO
		return ret
	}
	return *o.PortStatusSnapshot
}

// GetPortStatusSnapshotOk returns a tuple with the PortStatusSnapshot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortStatusSnapshotEntity) GetPortStatusSnapshotOk() (*PortStatusSnapshotDTO, bool) {
	if o == nil || o.PortStatusSnapshot == nil {
		return nil, false
	}
	return o.PortStatusSnapshot, true
}

// HasPortStatusSnapshot returns a boolean if a field has been set.
func (o *PortStatusSnapshotEntity) HasPortStatusSnapshot() bool {
	if o != nil && o.PortStatusSnapshot != nil {
		return true
	}

	return false
}

// SetPortStatusSnapshot gets a reference to the given PortStatusSnapshotDTO and assigns it to the PortStatusSnapshot field.
func (o *PortStatusSnapshotEntity) SetPortStatusSnapshot(v PortStatusSnapshotDTO) {
	o.PortStatusSnapshot = &v
}

// GetCanRead returns the CanRead field value if set, zero value otherwise.
func (o *PortStatusSnapshotEntity) GetCanRead() bool {
	if o == nil || o.CanRead == nil {
		var ret bool
		return ret
	}
	return *o.CanRead
}

// GetCanReadOk returns a tuple with the CanRead field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortStatusSnapshotEntity) GetCanReadOk() (*bool, bool) {
	if o == nil || o.CanRead == nil {
		return nil, false
	}
	return o.CanRead, true
}

// HasCanRead returns a boolean if a field has been set.
func (o *PortStatusSnapshotEntity) HasCanRead() bool {
	if o != nil && o.CanRead != nil {
		return true
	}

	return false
}

// SetCanRead gets a reference to the given bool and assigns it to the CanRead field.
func (o *PortStatusSnapshotEntity) SetCanRead(v bool) {
	o.CanRead = &v
}

func (o PortStatusSnapshotEntity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.PortStatusSnapshot != nil {
		toSerialize["portStatusSnapshot"] = o.PortStatusSnapshot
	}
	if o.CanRead != nil {
		toSerialize["canRead"] = o.CanRead
	}
	return json.Marshal(toSerialize)
}

type NullablePortStatusSnapshotEntity struct {
	value *PortStatusSnapshotEntity
	isSet bool
}

func (v NullablePortStatusSnapshotEntity) Get() *PortStatusSnapshotEntity {
	return v.value
}

func (v *NullablePortStatusSnapshotEntity) Set(val *PortStatusSnapshotEntity) {
	v.value = val
	v.isSet = true
}

func (v NullablePortStatusSnapshotEntity) IsSet() bool {
	return v.isSet
}

func (v *NullablePortStatusSnapshotEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortStatusSnapshotEntity(val *PortStatusSnapshotEntity) *NullablePortStatusSnapshotEntity {
	return &NullablePortStatusSnapshotEntity{value: val, isSet: true}
}

func (v NullablePortStatusSnapshotEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePortStatusSnapshotEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
