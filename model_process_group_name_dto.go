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

// ProcessGroupNameDTO struct for ProcessGroupNameDTO
type ProcessGroupNameDTO struct {
	// The ID of the Process Group
	Id *string `json:"id,omitempty"`
	// The name of the Process Group, or the ID of the Process Group if the user does not have the READ policy for the Process Group
	Name *string `json:"name,omitempty"`
}

// NewProcessGroupNameDTO instantiates a new ProcessGroupNameDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProcessGroupNameDTO() *ProcessGroupNameDTO {
	this := ProcessGroupNameDTO{}
	return &this
}

// NewProcessGroupNameDTOWithDefaults instantiates a new ProcessGroupNameDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProcessGroupNameDTOWithDefaults() *ProcessGroupNameDTO {
	this := ProcessGroupNameDTO{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ProcessGroupNameDTO) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessGroupNameDTO) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ProcessGroupNameDTO) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ProcessGroupNameDTO) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ProcessGroupNameDTO) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessGroupNameDTO) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ProcessGroupNameDTO) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ProcessGroupNameDTO) SetName(v string) {
	o.Name = &v
}

func (o ProcessGroupNameDTO) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableProcessGroupNameDTO struct {
	value *ProcessGroupNameDTO
	isSet bool
}

func (v NullableProcessGroupNameDTO) Get() *ProcessGroupNameDTO {
	return v.value
}

func (v *NullableProcessGroupNameDTO) Set(val *ProcessGroupNameDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableProcessGroupNameDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableProcessGroupNameDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProcessGroupNameDTO(val *ProcessGroupNameDTO) *NullableProcessGroupNameDTO {
	return &NullableProcessGroupNameDTO{value: val, isSet: true}
}

func (v NullableProcessGroupNameDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProcessGroupNameDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
