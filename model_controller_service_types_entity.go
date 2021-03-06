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

// ControllerServiceTypesEntity struct for ControllerServiceTypesEntity
type ControllerServiceTypesEntity struct {
	ControllerServiceTypes *[]DocumentedTypeDTO `json:"controllerServiceTypes,omitempty"`
}

// NewControllerServiceTypesEntity instantiates a new ControllerServiceTypesEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewControllerServiceTypesEntity() *ControllerServiceTypesEntity {
	this := ControllerServiceTypesEntity{}
	return &this
}

// NewControllerServiceTypesEntityWithDefaults instantiates a new ControllerServiceTypesEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewControllerServiceTypesEntityWithDefaults() *ControllerServiceTypesEntity {
	this := ControllerServiceTypesEntity{}
	return &this
}

// GetControllerServiceTypes returns the ControllerServiceTypes field value if set, zero value otherwise.
func (o *ControllerServiceTypesEntity) GetControllerServiceTypes() []DocumentedTypeDTO {
	if o == nil || o.ControllerServiceTypes == nil {
		var ret []DocumentedTypeDTO
		return ret
	}
	return *o.ControllerServiceTypes
}

// GetControllerServiceTypesOk returns a tuple with the ControllerServiceTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllerServiceTypesEntity) GetControllerServiceTypesOk() (*[]DocumentedTypeDTO, bool) {
	if o == nil || o.ControllerServiceTypes == nil {
		return nil, false
	}
	return o.ControllerServiceTypes, true
}

// HasControllerServiceTypes returns a boolean if a field has been set.
func (o *ControllerServiceTypesEntity) HasControllerServiceTypes() bool {
	if o != nil && o.ControllerServiceTypes != nil {
		return true
	}

	return false
}

// SetControllerServiceTypes gets a reference to the given []DocumentedTypeDTO and assigns it to the ControllerServiceTypes field.
func (o *ControllerServiceTypesEntity) SetControllerServiceTypes(v []DocumentedTypeDTO) {
	o.ControllerServiceTypes = &v
}

func (o ControllerServiceTypesEntity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ControllerServiceTypes != nil {
		toSerialize["controllerServiceTypes"] = o.ControllerServiceTypes
	}
	return json.Marshal(toSerialize)
}

type NullableControllerServiceTypesEntity struct {
	value *ControllerServiceTypesEntity
	isSet bool
}

func (v NullableControllerServiceTypesEntity) Get() *ControllerServiceTypesEntity {
	return v.value
}

func (v *NullableControllerServiceTypesEntity) Set(val *ControllerServiceTypesEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableControllerServiceTypesEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableControllerServiceTypesEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableControllerServiceTypesEntity(val *ControllerServiceTypesEntity) *NullableControllerServiceTypesEntity {
	return &NullableControllerServiceTypesEntity{value: val, isSet: true}
}

func (v NullableControllerServiceTypesEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableControllerServiceTypesEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
