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

// ControllerServiceStatusDTO struct for ControllerServiceStatusDTO
type ControllerServiceStatusDTO struct {
	// The run status of this ControllerService
	RunStatus *string `json:"runStatus,omitempty"`
	// Indicates whether the component is valid, invalid, or still in the process of validating (i.e., it is unknown whether or not the component is valid)
	ValidationStatus *string `json:"validationStatus,omitempty"`
	// The number of active threads for the component.
	ActiveThreadCount *int32 `json:"activeThreadCount,omitempty"`
}

// NewControllerServiceStatusDTO instantiates a new ControllerServiceStatusDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewControllerServiceStatusDTO() *ControllerServiceStatusDTO {
	this := ControllerServiceStatusDTO{}
	return &this
}

// NewControllerServiceStatusDTOWithDefaults instantiates a new ControllerServiceStatusDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewControllerServiceStatusDTOWithDefaults() *ControllerServiceStatusDTO {
	this := ControllerServiceStatusDTO{}
	return &this
}

// GetRunStatus returns the RunStatus field value if set, zero value otherwise.
func (o *ControllerServiceStatusDTO) GetRunStatus() string {
	if o == nil || o.RunStatus == nil {
		var ret string
		return ret
	}
	return *o.RunStatus
}

// GetRunStatusOk returns a tuple with the RunStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllerServiceStatusDTO) GetRunStatusOk() (*string, bool) {
	if o == nil || o.RunStatus == nil {
		return nil, false
	}
	return o.RunStatus, true
}

// HasRunStatus returns a boolean if a field has been set.
func (o *ControllerServiceStatusDTO) HasRunStatus() bool {
	if o != nil && o.RunStatus != nil {
		return true
	}

	return false
}

// SetRunStatus gets a reference to the given string and assigns it to the RunStatus field.
func (o *ControllerServiceStatusDTO) SetRunStatus(v string) {
	o.RunStatus = &v
}

// GetValidationStatus returns the ValidationStatus field value if set, zero value otherwise.
func (o *ControllerServiceStatusDTO) GetValidationStatus() string {
	if o == nil || o.ValidationStatus == nil {
		var ret string
		return ret
	}
	return *o.ValidationStatus
}

// GetValidationStatusOk returns a tuple with the ValidationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllerServiceStatusDTO) GetValidationStatusOk() (*string, bool) {
	if o == nil || o.ValidationStatus == nil {
		return nil, false
	}
	return o.ValidationStatus, true
}

// HasValidationStatus returns a boolean if a field has been set.
func (o *ControllerServiceStatusDTO) HasValidationStatus() bool {
	if o != nil && o.ValidationStatus != nil {
		return true
	}

	return false
}

// SetValidationStatus gets a reference to the given string and assigns it to the ValidationStatus field.
func (o *ControllerServiceStatusDTO) SetValidationStatus(v string) {
	o.ValidationStatus = &v
}

// GetActiveThreadCount returns the ActiveThreadCount field value if set, zero value otherwise.
func (o *ControllerServiceStatusDTO) GetActiveThreadCount() int32 {
	if o == nil || o.ActiveThreadCount == nil {
		var ret int32
		return ret
	}
	return *o.ActiveThreadCount
}

// GetActiveThreadCountOk returns a tuple with the ActiveThreadCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllerServiceStatusDTO) GetActiveThreadCountOk() (*int32, bool) {
	if o == nil || o.ActiveThreadCount == nil {
		return nil, false
	}
	return o.ActiveThreadCount, true
}

// HasActiveThreadCount returns a boolean if a field has been set.
func (o *ControllerServiceStatusDTO) HasActiveThreadCount() bool {
	if o != nil && o.ActiveThreadCount != nil {
		return true
	}

	return false
}

// SetActiveThreadCount gets a reference to the given int32 and assigns it to the ActiveThreadCount field.
func (o *ControllerServiceStatusDTO) SetActiveThreadCount(v int32) {
	o.ActiveThreadCount = &v
}

func (o ControllerServiceStatusDTO) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RunStatus != nil {
		toSerialize["runStatus"] = o.RunStatus
	}
	if o.ValidationStatus != nil {
		toSerialize["validationStatus"] = o.ValidationStatus
	}
	if o.ActiveThreadCount != nil {
		toSerialize["activeThreadCount"] = o.ActiveThreadCount
	}
	return json.Marshal(toSerialize)
}

type NullableControllerServiceStatusDTO struct {
	value *ControllerServiceStatusDTO
	isSet bool
}

func (v NullableControllerServiceStatusDTO) Get() *ControllerServiceStatusDTO {
	return v.value
}

func (v *NullableControllerServiceStatusDTO) Set(val *ControllerServiceStatusDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableControllerServiceStatusDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableControllerServiceStatusDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableControllerServiceStatusDTO(val *ControllerServiceStatusDTO) *NullableControllerServiceStatusDTO {
	return &NullableControllerServiceStatusDTO{value: val, isSet: true}
}

func (v NullableControllerServiceStatusDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableControllerServiceStatusDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
