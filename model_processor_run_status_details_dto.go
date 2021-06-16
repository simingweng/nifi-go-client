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

// ProcessorRunStatusDetailsDTO struct for ProcessorRunStatusDetailsDTO
type ProcessorRunStatusDetailsDTO struct {
	// The ID of the processor
	Id *string `json:"id,omitempty"`
	// The name of the processor
	Name *string `json:"name,omitempty"`
	// The run status of the processor
	RunStatus *string `json:"runStatus,omitempty"`
	// The processor's validation errors
	ValidationErrors *[]string `json:"validationErrors,omitempty"`
	// The current number of threads that the processor is currently using
	ActiveThreadCount *int32 `json:"activeThreadCount,omitempty"`
}

// NewProcessorRunStatusDetailsDTO instantiates a new ProcessorRunStatusDetailsDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProcessorRunStatusDetailsDTO() *ProcessorRunStatusDetailsDTO {
	this := ProcessorRunStatusDetailsDTO{}
	return &this
}

// NewProcessorRunStatusDetailsDTOWithDefaults instantiates a new ProcessorRunStatusDetailsDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProcessorRunStatusDetailsDTOWithDefaults() *ProcessorRunStatusDetailsDTO {
	this := ProcessorRunStatusDetailsDTO{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ProcessorRunStatusDetailsDTO) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorRunStatusDetailsDTO) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ProcessorRunStatusDetailsDTO) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ProcessorRunStatusDetailsDTO) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ProcessorRunStatusDetailsDTO) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorRunStatusDetailsDTO) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ProcessorRunStatusDetailsDTO) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ProcessorRunStatusDetailsDTO) SetName(v string) {
	o.Name = &v
}

// GetRunStatus returns the RunStatus field value if set, zero value otherwise.
func (o *ProcessorRunStatusDetailsDTO) GetRunStatus() string {
	if o == nil || o.RunStatus == nil {
		var ret string
		return ret
	}
	return *o.RunStatus
}

// GetRunStatusOk returns a tuple with the RunStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorRunStatusDetailsDTO) GetRunStatusOk() (*string, bool) {
	if o == nil || o.RunStatus == nil {
		return nil, false
	}
	return o.RunStatus, true
}

// HasRunStatus returns a boolean if a field has been set.
func (o *ProcessorRunStatusDetailsDTO) HasRunStatus() bool {
	if o != nil && o.RunStatus != nil {
		return true
	}

	return false
}

// SetRunStatus gets a reference to the given string and assigns it to the RunStatus field.
func (o *ProcessorRunStatusDetailsDTO) SetRunStatus(v string) {
	o.RunStatus = &v
}

// GetValidationErrors returns the ValidationErrors field value if set, zero value otherwise.
func (o *ProcessorRunStatusDetailsDTO) GetValidationErrors() []string {
	if o == nil || o.ValidationErrors == nil {
		var ret []string
		return ret
	}
	return *o.ValidationErrors
}

// GetValidationErrorsOk returns a tuple with the ValidationErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorRunStatusDetailsDTO) GetValidationErrorsOk() (*[]string, bool) {
	if o == nil || o.ValidationErrors == nil {
		return nil, false
	}
	return o.ValidationErrors, true
}

// HasValidationErrors returns a boolean if a field has been set.
func (o *ProcessorRunStatusDetailsDTO) HasValidationErrors() bool {
	if o != nil && o.ValidationErrors != nil {
		return true
	}

	return false
}

// SetValidationErrors gets a reference to the given []string and assigns it to the ValidationErrors field.
func (o *ProcessorRunStatusDetailsDTO) SetValidationErrors(v []string) {
	o.ValidationErrors = &v
}

// GetActiveThreadCount returns the ActiveThreadCount field value if set, zero value otherwise.
func (o *ProcessorRunStatusDetailsDTO) GetActiveThreadCount() int32 {
	if o == nil || o.ActiveThreadCount == nil {
		var ret int32
		return ret
	}
	return *o.ActiveThreadCount
}

// GetActiveThreadCountOk returns a tuple with the ActiveThreadCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorRunStatusDetailsDTO) GetActiveThreadCountOk() (*int32, bool) {
	if o == nil || o.ActiveThreadCount == nil {
		return nil, false
	}
	return o.ActiveThreadCount, true
}

// HasActiveThreadCount returns a boolean if a field has been set.
func (o *ProcessorRunStatusDetailsDTO) HasActiveThreadCount() bool {
	if o != nil && o.ActiveThreadCount != nil {
		return true
	}

	return false
}

// SetActiveThreadCount gets a reference to the given int32 and assigns it to the ActiveThreadCount field.
func (o *ProcessorRunStatusDetailsDTO) SetActiveThreadCount(v int32) {
	o.ActiveThreadCount = &v
}

func (o ProcessorRunStatusDetailsDTO) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.RunStatus != nil {
		toSerialize["runStatus"] = o.RunStatus
	}
	if o.ValidationErrors != nil {
		toSerialize["validationErrors"] = o.ValidationErrors
	}
	if o.ActiveThreadCount != nil {
		toSerialize["activeThreadCount"] = o.ActiveThreadCount
	}
	return json.Marshal(toSerialize)
}

type NullableProcessorRunStatusDetailsDTO struct {
	value *ProcessorRunStatusDetailsDTO
	isSet bool
}

func (v NullableProcessorRunStatusDetailsDTO) Get() *ProcessorRunStatusDetailsDTO {
	return v.value
}

func (v *NullableProcessorRunStatusDetailsDTO) Set(val *ProcessorRunStatusDetailsDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableProcessorRunStatusDetailsDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableProcessorRunStatusDetailsDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProcessorRunStatusDetailsDTO(val *ProcessorRunStatusDetailsDTO) *NullableProcessorRunStatusDetailsDTO {
	return &NullableProcessorRunStatusDetailsDTO{value: val, isSet: true}
}

func (v NullableProcessorRunStatusDetailsDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProcessorRunStatusDetailsDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
