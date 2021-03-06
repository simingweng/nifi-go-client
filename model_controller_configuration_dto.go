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

// ControllerConfigurationDTO struct for ControllerConfigurationDTO
type ControllerConfigurationDTO struct {
	// The maximum number of timer driven threads the NiFi has available.
	MaxTimerDrivenThreadCount *int32 `json:"maxTimerDrivenThreadCount,omitempty"`
	// The maximum number of event driven threads the NiFi has available.
	MaxEventDrivenThreadCount *int32 `json:"maxEventDrivenThreadCount,omitempty"`
}

// NewControllerConfigurationDTO instantiates a new ControllerConfigurationDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewControllerConfigurationDTO() *ControllerConfigurationDTO {
	this := ControllerConfigurationDTO{}
	return &this
}

// NewControllerConfigurationDTOWithDefaults instantiates a new ControllerConfigurationDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewControllerConfigurationDTOWithDefaults() *ControllerConfigurationDTO {
	this := ControllerConfigurationDTO{}
	return &this
}

// GetMaxTimerDrivenThreadCount returns the MaxTimerDrivenThreadCount field value if set, zero value otherwise.
func (o *ControllerConfigurationDTO) GetMaxTimerDrivenThreadCount() int32 {
	if o == nil || o.MaxTimerDrivenThreadCount == nil {
		var ret int32
		return ret
	}
	return *o.MaxTimerDrivenThreadCount
}

// GetMaxTimerDrivenThreadCountOk returns a tuple with the MaxTimerDrivenThreadCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllerConfigurationDTO) GetMaxTimerDrivenThreadCountOk() (*int32, bool) {
	if o == nil || o.MaxTimerDrivenThreadCount == nil {
		return nil, false
	}
	return o.MaxTimerDrivenThreadCount, true
}

// HasMaxTimerDrivenThreadCount returns a boolean if a field has been set.
func (o *ControllerConfigurationDTO) HasMaxTimerDrivenThreadCount() bool {
	if o != nil && o.MaxTimerDrivenThreadCount != nil {
		return true
	}

	return false
}

// SetMaxTimerDrivenThreadCount gets a reference to the given int32 and assigns it to the MaxTimerDrivenThreadCount field.
func (o *ControllerConfigurationDTO) SetMaxTimerDrivenThreadCount(v int32) {
	o.MaxTimerDrivenThreadCount = &v
}

// GetMaxEventDrivenThreadCount returns the MaxEventDrivenThreadCount field value if set, zero value otherwise.
func (o *ControllerConfigurationDTO) GetMaxEventDrivenThreadCount() int32 {
	if o == nil || o.MaxEventDrivenThreadCount == nil {
		var ret int32
		return ret
	}
	return *o.MaxEventDrivenThreadCount
}

// GetMaxEventDrivenThreadCountOk returns a tuple with the MaxEventDrivenThreadCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllerConfigurationDTO) GetMaxEventDrivenThreadCountOk() (*int32, bool) {
	if o == nil || o.MaxEventDrivenThreadCount == nil {
		return nil, false
	}
	return o.MaxEventDrivenThreadCount, true
}

// HasMaxEventDrivenThreadCount returns a boolean if a field has been set.
func (o *ControllerConfigurationDTO) HasMaxEventDrivenThreadCount() bool {
	if o != nil && o.MaxEventDrivenThreadCount != nil {
		return true
	}

	return false
}

// SetMaxEventDrivenThreadCount gets a reference to the given int32 and assigns it to the MaxEventDrivenThreadCount field.
func (o *ControllerConfigurationDTO) SetMaxEventDrivenThreadCount(v int32) {
	o.MaxEventDrivenThreadCount = &v
}

func (o ControllerConfigurationDTO) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MaxTimerDrivenThreadCount != nil {
		toSerialize["maxTimerDrivenThreadCount"] = o.MaxTimerDrivenThreadCount
	}
	if o.MaxEventDrivenThreadCount != nil {
		toSerialize["maxEventDrivenThreadCount"] = o.MaxEventDrivenThreadCount
	}
	return json.Marshal(toSerialize)
}

type NullableControllerConfigurationDTO struct {
	value *ControllerConfigurationDTO
	isSet bool
}

func (v NullableControllerConfigurationDTO) Get() *ControllerConfigurationDTO {
	return v.value
}

func (v *NullableControllerConfigurationDTO) Set(val *ControllerConfigurationDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableControllerConfigurationDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableControllerConfigurationDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableControllerConfigurationDTO(val *ControllerConfigurationDTO) *NullableControllerConfigurationDTO {
	return &NullableControllerConfigurationDTO{value: val, isSet: true}
}

func (v NullableControllerConfigurationDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableControllerConfigurationDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
