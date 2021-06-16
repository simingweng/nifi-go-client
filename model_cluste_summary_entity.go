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

// ClusteSummaryEntity struct for ClusteSummaryEntity
type ClusteSummaryEntity struct {
	ClusterSummary *ClusterSummaryDTO `json:"clusterSummary,omitempty"`
}

// NewClusteSummaryEntity instantiates a new ClusteSummaryEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusteSummaryEntity() *ClusteSummaryEntity {
	this := ClusteSummaryEntity{}
	return &this
}

// NewClusteSummaryEntityWithDefaults instantiates a new ClusteSummaryEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusteSummaryEntityWithDefaults() *ClusteSummaryEntity {
	this := ClusteSummaryEntity{}
	return &this
}

// GetClusterSummary returns the ClusterSummary field value if set, zero value otherwise.
func (o *ClusteSummaryEntity) GetClusterSummary() ClusterSummaryDTO {
	if o == nil || o.ClusterSummary == nil {
		var ret ClusterSummaryDTO
		return ret
	}
	return *o.ClusterSummary
}

// GetClusterSummaryOk returns a tuple with the ClusterSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusteSummaryEntity) GetClusterSummaryOk() (*ClusterSummaryDTO, bool) {
	if o == nil || o.ClusterSummary == nil {
		return nil, false
	}
	return o.ClusterSummary, true
}

// HasClusterSummary returns a boolean if a field has been set.
func (o *ClusteSummaryEntity) HasClusterSummary() bool {
	if o != nil && o.ClusterSummary != nil {
		return true
	}

	return false
}

// SetClusterSummary gets a reference to the given ClusterSummaryDTO and assigns it to the ClusterSummary field.
func (o *ClusteSummaryEntity) SetClusterSummary(v ClusterSummaryDTO) {
	o.ClusterSummary = &v
}

func (o ClusteSummaryEntity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClusterSummary != nil {
		toSerialize["clusterSummary"] = o.ClusterSummary
	}
	return json.Marshal(toSerialize)
}

type NullableClusteSummaryEntity struct {
	value *ClusteSummaryEntity
	isSet bool
}

func (v NullableClusteSummaryEntity) Get() *ClusteSummaryEntity {
	return v.value
}

func (v *NullableClusteSummaryEntity) Set(val *ClusteSummaryEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableClusteSummaryEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableClusteSummaryEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusteSummaryEntity(val *ClusteSummaryEntity) *NullableClusteSummaryEntity {
	return &NullableClusteSummaryEntity{value: val, isSet: true}
}

func (v NullableClusteSummaryEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusteSummaryEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
