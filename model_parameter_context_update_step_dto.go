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

// ParameterContextUpdateStepDTO struct for ParameterContextUpdateStepDTO
type ParameterContextUpdateStepDTO struct {
	// Explanation of what happens in this step
	Description *string `json:"description,omitempty"`
	// Whether or not this step has completed
	Complete *bool `json:"complete,omitempty"`
	// An explanation of why this step failed, or null if this step did not fail
	FailureReason *string `json:"failureReason,omitempty"`
}

// NewParameterContextUpdateStepDTO instantiates a new ParameterContextUpdateStepDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParameterContextUpdateStepDTO() *ParameterContextUpdateStepDTO {
	this := ParameterContextUpdateStepDTO{}
	return &this
}

// NewParameterContextUpdateStepDTOWithDefaults instantiates a new ParameterContextUpdateStepDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParameterContextUpdateStepDTOWithDefaults() *ParameterContextUpdateStepDTO {
	this := ParameterContextUpdateStepDTO{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ParameterContextUpdateStepDTO) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterContextUpdateStepDTO) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ParameterContextUpdateStepDTO) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ParameterContextUpdateStepDTO) SetDescription(v string) {
	o.Description = &v
}

// GetComplete returns the Complete field value if set, zero value otherwise.
func (o *ParameterContextUpdateStepDTO) GetComplete() bool {
	if o == nil || o.Complete == nil {
		var ret bool
		return ret
	}
	return *o.Complete
}

// GetCompleteOk returns a tuple with the Complete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterContextUpdateStepDTO) GetCompleteOk() (*bool, bool) {
	if o == nil || o.Complete == nil {
		return nil, false
	}
	return o.Complete, true
}

// HasComplete returns a boolean if a field has been set.
func (o *ParameterContextUpdateStepDTO) HasComplete() bool {
	if o != nil && o.Complete != nil {
		return true
	}

	return false
}

// SetComplete gets a reference to the given bool and assigns it to the Complete field.
func (o *ParameterContextUpdateStepDTO) SetComplete(v bool) {
	o.Complete = &v
}

// GetFailureReason returns the FailureReason field value if set, zero value otherwise.
func (o *ParameterContextUpdateStepDTO) GetFailureReason() string {
	if o == nil || o.FailureReason == nil {
		var ret string
		return ret
	}
	return *o.FailureReason
}

// GetFailureReasonOk returns a tuple with the FailureReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterContextUpdateStepDTO) GetFailureReasonOk() (*string, bool) {
	if o == nil || o.FailureReason == nil {
		return nil, false
	}
	return o.FailureReason, true
}

// HasFailureReason returns a boolean if a field has been set.
func (o *ParameterContextUpdateStepDTO) HasFailureReason() bool {
	if o != nil && o.FailureReason != nil {
		return true
	}

	return false
}

// SetFailureReason gets a reference to the given string and assigns it to the FailureReason field.
func (o *ParameterContextUpdateStepDTO) SetFailureReason(v string) {
	o.FailureReason = &v
}

func (o ParameterContextUpdateStepDTO) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Complete != nil {
		toSerialize["complete"] = o.Complete
	}
	if o.FailureReason != nil {
		toSerialize["failureReason"] = o.FailureReason
	}
	return json.Marshal(toSerialize)
}

type NullableParameterContextUpdateStepDTO struct {
	value *ParameterContextUpdateStepDTO
	isSet bool
}

func (v NullableParameterContextUpdateStepDTO) Get() *ParameterContextUpdateStepDTO {
	return v.value
}

func (v *NullableParameterContextUpdateStepDTO) Set(val *ParameterContextUpdateStepDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableParameterContextUpdateStepDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableParameterContextUpdateStepDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParameterContextUpdateStepDTO(val *ParameterContextUpdateStepDTO) *NullableParameterContextUpdateStepDTO {
	return &NullableParameterContextUpdateStepDTO{value: val, isSet: true}
}

func (v NullableParameterContextUpdateStepDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParameterContextUpdateStepDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
