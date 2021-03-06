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

// VersionedFlowSnapshotMetadataEntity struct for VersionedFlowSnapshotMetadataEntity
type VersionedFlowSnapshotMetadataEntity struct {
	VersionedFlowSnapshotMetadata *VersionedFlowSnapshotMetadata `json:"versionedFlowSnapshotMetadata,omitempty"`
	// The ID of the Registry that this flow belongs to
	RegistryId *string `json:"registryId,omitempty"`
}

// NewVersionedFlowSnapshotMetadataEntity instantiates a new VersionedFlowSnapshotMetadataEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVersionedFlowSnapshotMetadataEntity() *VersionedFlowSnapshotMetadataEntity {
	this := VersionedFlowSnapshotMetadataEntity{}
	return &this
}

// NewVersionedFlowSnapshotMetadataEntityWithDefaults instantiates a new VersionedFlowSnapshotMetadataEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVersionedFlowSnapshotMetadataEntityWithDefaults() *VersionedFlowSnapshotMetadataEntity {
	this := VersionedFlowSnapshotMetadataEntity{}
	return &this
}

// GetVersionedFlowSnapshotMetadata returns the VersionedFlowSnapshotMetadata field value if set, zero value otherwise.
func (o *VersionedFlowSnapshotMetadataEntity) GetVersionedFlowSnapshotMetadata() VersionedFlowSnapshotMetadata {
	if o == nil || o.VersionedFlowSnapshotMetadata == nil {
		var ret VersionedFlowSnapshotMetadata
		return ret
	}
	return *o.VersionedFlowSnapshotMetadata
}

// GetVersionedFlowSnapshotMetadataOk returns a tuple with the VersionedFlowSnapshotMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionedFlowSnapshotMetadataEntity) GetVersionedFlowSnapshotMetadataOk() (*VersionedFlowSnapshotMetadata, bool) {
	if o == nil || o.VersionedFlowSnapshotMetadata == nil {
		return nil, false
	}
	return o.VersionedFlowSnapshotMetadata, true
}

// HasVersionedFlowSnapshotMetadata returns a boolean if a field has been set.
func (o *VersionedFlowSnapshotMetadataEntity) HasVersionedFlowSnapshotMetadata() bool {
	if o != nil && o.VersionedFlowSnapshotMetadata != nil {
		return true
	}

	return false
}

// SetVersionedFlowSnapshotMetadata gets a reference to the given VersionedFlowSnapshotMetadata and assigns it to the VersionedFlowSnapshotMetadata field.
func (o *VersionedFlowSnapshotMetadataEntity) SetVersionedFlowSnapshotMetadata(v VersionedFlowSnapshotMetadata) {
	o.VersionedFlowSnapshotMetadata = &v
}

// GetRegistryId returns the RegistryId field value if set, zero value otherwise.
func (o *VersionedFlowSnapshotMetadataEntity) GetRegistryId() string {
	if o == nil || o.RegistryId == nil {
		var ret string
		return ret
	}
	return *o.RegistryId
}

// GetRegistryIdOk returns a tuple with the RegistryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionedFlowSnapshotMetadataEntity) GetRegistryIdOk() (*string, bool) {
	if o == nil || o.RegistryId == nil {
		return nil, false
	}
	return o.RegistryId, true
}

// HasRegistryId returns a boolean if a field has been set.
func (o *VersionedFlowSnapshotMetadataEntity) HasRegistryId() bool {
	if o != nil && o.RegistryId != nil {
		return true
	}

	return false
}

// SetRegistryId gets a reference to the given string and assigns it to the RegistryId field.
func (o *VersionedFlowSnapshotMetadataEntity) SetRegistryId(v string) {
	o.RegistryId = &v
}

func (o VersionedFlowSnapshotMetadataEntity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.VersionedFlowSnapshotMetadata != nil {
		toSerialize["versionedFlowSnapshotMetadata"] = o.VersionedFlowSnapshotMetadata
	}
	if o.RegistryId != nil {
		toSerialize["registryId"] = o.RegistryId
	}
	return json.Marshal(toSerialize)
}

type NullableVersionedFlowSnapshotMetadataEntity struct {
	value *VersionedFlowSnapshotMetadataEntity
	isSet bool
}

func (v NullableVersionedFlowSnapshotMetadataEntity) Get() *VersionedFlowSnapshotMetadataEntity {
	return v.value
}

func (v *NullableVersionedFlowSnapshotMetadataEntity) Set(val *VersionedFlowSnapshotMetadataEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableVersionedFlowSnapshotMetadataEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableVersionedFlowSnapshotMetadataEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVersionedFlowSnapshotMetadataEntity(val *VersionedFlowSnapshotMetadataEntity) *NullableVersionedFlowSnapshotMetadataEntity {
	return &NullableVersionedFlowSnapshotMetadataEntity{value: val, isSet: true}
}

func (v NullableVersionedFlowSnapshotMetadataEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVersionedFlowSnapshotMetadataEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
