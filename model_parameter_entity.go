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

// ParameterEntity struct for ParameterEntity
type ParameterEntity struct {
	// Indicates whether the user can write a given resource.
	CanWrite  *bool         `json:"canWrite,omitempty"`
	Parameter *ParameterDTO `json:"parameter,omitempty"`
}

// NewParameterEntity instantiates a new ParameterEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParameterEntity() *ParameterEntity {
	this := ParameterEntity{}
	return &this
}

// NewParameterEntityWithDefaults instantiates a new ParameterEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParameterEntityWithDefaults() *ParameterEntity {
	this := ParameterEntity{}
	return &this
}

// GetCanWrite returns the CanWrite field value if set, zero value otherwise.
func (o *ParameterEntity) GetCanWrite() bool {
	if o == nil || o.CanWrite == nil {
		var ret bool
		return ret
	}
	return *o.CanWrite
}

// GetCanWriteOk returns a tuple with the CanWrite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterEntity) GetCanWriteOk() (*bool, bool) {
	if o == nil || o.CanWrite == nil {
		return nil, false
	}
	return o.CanWrite, true
}

// HasCanWrite returns a boolean if a field has been set.
func (o *ParameterEntity) HasCanWrite() bool {
	if o != nil && o.CanWrite != nil {
		return true
	}

	return false
}

// SetCanWrite gets a reference to the given bool and assigns it to the CanWrite field.
func (o *ParameterEntity) SetCanWrite(v bool) {
	o.CanWrite = &v
}

// GetParameter returns the Parameter field value if set, zero value otherwise.
func (o *ParameterEntity) GetParameter() ParameterDTO {
	if o == nil || o.Parameter == nil {
		var ret ParameterDTO
		return ret
	}
	return *o.Parameter
}

// GetParameterOk returns a tuple with the Parameter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterEntity) GetParameterOk() (*ParameterDTO, bool) {
	if o == nil || o.Parameter == nil {
		return nil, false
	}
	return o.Parameter, true
}

// HasParameter returns a boolean if a field has been set.
func (o *ParameterEntity) HasParameter() bool {
	if o != nil && o.Parameter != nil {
		return true
	}

	return false
}

// SetParameter gets a reference to the given ParameterDTO and assigns it to the Parameter field.
func (o *ParameterEntity) SetParameter(v ParameterDTO) {
	o.Parameter = &v
}

func (o ParameterEntity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CanWrite != nil {
		toSerialize["canWrite"] = o.CanWrite
	}
	if o.Parameter != nil {
		toSerialize["parameter"] = o.Parameter
	}
	return json.Marshal(toSerialize)
}

type NullableParameterEntity struct {
	value *ParameterEntity
	isSet bool
}

func (v NullableParameterEntity) Get() *ParameterEntity {
	return v.value
}

func (v *NullableParameterEntity) Set(val *ParameterEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableParameterEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableParameterEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParameterEntity(val *ParameterEntity) *NullableParameterEntity {
	return &NullableParameterEntity{value: val, isSet: true}
}

func (v NullableParameterEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParameterEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
