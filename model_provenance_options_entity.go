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

// ProvenanceOptionsEntity struct for ProvenanceOptionsEntity
type ProvenanceOptionsEntity struct {
	ProvenanceOptions *ProvenanceOptionsDTO `json:"provenanceOptions,omitempty"`
}

// NewProvenanceOptionsEntity instantiates a new ProvenanceOptionsEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProvenanceOptionsEntity() *ProvenanceOptionsEntity {
	this := ProvenanceOptionsEntity{}
	return &this
}

// NewProvenanceOptionsEntityWithDefaults instantiates a new ProvenanceOptionsEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProvenanceOptionsEntityWithDefaults() *ProvenanceOptionsEntity {
	this := ProvenanceOptionsEntity{}
	return &this
}

// GetProvenanceOptions returns the ProvenanceOptions field value if set, zero value otherwise.
func (o *ProvenanceOptionsEntity) GetProvenanceOptions() ProvenanceOptionsDTO {
	if o == nil || o.ProvenanceOptions == nil {
		var ret ProvenanceOptionsDTO
		return ret
	}
	return *o.ProvenanceOptions
}

// GetProvenanceOptionsOk returns a tuple with the ProvenanceOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvenanceOptionsEntity) GetProvenanceOptionsOk() (*ProvenanceOptionsDTO, bool) {
	if o == nil || o.ProvenanceOptions == nil {
		return nil, false
	}
	return o.ProvenanceOptions, true
}

// HasProvenanceOptions returns a boolean if a field has been set.
func (o *ProvenanceOptionsEntity) HasProvenanceOptions() bool {
	if o != nil && o.ProvenanceOptions != nil {
		return true
	}

	return false
}

// SetProvenanceOptions gets a reference to the given ProvenanceOptionsDTO and assigns it to the ProvenanceOptions field.
func (o *ProvenanceOptionsEntity) SetProvenanceOptions(v ProvenanceOptionsDTO) {
	o.ProvenanceOptions = &v
}

func (o ProvenanceOptionsEntity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ProvenanceOptions != nil {
		toSerialize["provenanceOptions"] = o.ProvenanceOptions
	}
	return json.Marshal(toSerialize)
}

type NullableProvenanceOptionsEntity struct {
	value *ProvenanceOptionsEntity
	isSet bool
}

func (v NullableProvenanceOptionsEntity) Get() *ProvenanceOptionsEntity {
	return v.value
}

func (v *NullableProvenanceOptionsEntity) Set(val *ProvenanceOptionsEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableProvenanceOptionsEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableProvenanceOptionsEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProvenanceOptionsEntity(val *ProvenanceOptionsEntity) *NullableProvenanceOptionsEntity {
	return &NullableProvenanceOptionsEntity{value: val, isSet: true}
}

func (v NullableProvenanceOptionsEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProvenanceOptionsEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
