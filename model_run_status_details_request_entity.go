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

// RunStatusDetailsRequestEntity struct for RunStatusDetailsRequestEntity
type RunStatusDetailsRequestEntity struct {
	// The IDs of all processors whose run status details should be provided
	ProcessorIds *[]string `json:"processorIds,omitempty"`
}

// NewRunStatusDetailsRequestEntity instantiates a new RunStatusDetailsRequestEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRunStatusDetailsRequestEntity() *RunStatusDetailsRequestEntity {
	this := RunStatusDetailsRequestEntity{}
	return &this
}

// NewRunStatusDetailsRequestEntityWithDefaults instantiates a new RunStatusDetailsRequestEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRunStatusDetailsRequestEntityWithDefaults() *RunStatusDetailsRequestEntity {
	this := RunStatusDetailsRequestEntity{}
	return &this
}

// GetProcessorIds returns the ProcessorIds field value if set, zero value otherwise.
func (o *RunStatusDetailsRequestEntity) GetProcessorIds() []string {
	if o == nil || o.ProcessorIds == nil {
		var ret []string
		return ret
	}
	return *o.ProcessorIds
}

// GetProcessorIdsOk returns a tuple with the ProcessorIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunStatusDetailsRequestEntity) GetProcessorIdsOk() (*[]string, bool) {
	if o == nil || o.ProcessorIds == nil {
		return nil, false
	}
	return o.ProcessorIds, true
}

// HasProcessorIds returns a boolean if a field has been set.
func (o *RunStatusDetailsRequestEntity) HasProcessorIds() bool {
	if o != nil && o.ProcessorIds != nil {
		return true
	}

	return false
}

// SetProcessorIds gets a reference to the given []string and assigns it to the ProcessorIds field.
func (o *RunStatusDetailsRequestEntity) SetProcessorIds(v []string) {
	o.ProcessorIds = &v
}

func (o RunStatusDetailsRequestEntity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ProcessorIds != nil {
		toSerialize["processorIds"] = o.ProcessorIds
	}
	return json.Marshal(toSerialize)
}

type NullableRunStatusDetailsRequestEntity struct {
	value *RunStatusDetailsRequestEntity
	isSet bool
}

func (v NullableRunStatusDetailsRequestEntity) Get() *RunStatusDetailsRequestEntity {
	return v.value
}

func (v *NullableRunStatusDetailsRequestEntity) Set(val *RunStatusDetailsRequestEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableRunStatusDetailsRequestEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableRunStatusDetailsRequestEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRunStatusDetailsRequestEntity(val *RunStatusDetailsRequestEntity) *NullableRunStatusDetailsRequestEntity {
	return &NullableRunStatusDetailsRequestEntity{value: val, isSet: true}
}

func (v NullableRunStatusDetailsRequestEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRunStatusDetailsRequestEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
