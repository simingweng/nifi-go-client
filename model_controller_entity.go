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

// ControllerEntity struct for ControllerEntity
type ControllerEntity struct {
	Controller *ControllerDTO `json:"controller,omitempty"`
}

// NewControllerEntity instantiates a new ControllerEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewControllerEntity() *ControllerEntity {
	this := ControllerEntity{}
	return &this
}

// NewControllerEntityWithDefaults instantiates a new ControllerEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewControllerEntityWithDefaults() *ControllerEntity {
	this := ControllerEntity{}
	return &this
}

// GetController returns the Controller field value if set, zero value otherwise.
func (o *ControllerEntity) GetController() ControllerDTO {
	if o == nil || o.Controller == nil {
		var ret ControllerDTO
		return ret
	}
	return *o.Controller
}

// GetControllerOk returns a tuple with the Controller field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllerEntity) GetControllerOk() (*ControllerDTO, bool) {
	if o == nil || o.Controller == nil {
		return nil, false
	}
	return o.Controller, true
}

// HasController returns a boolean if a field has been set.
func (o *ControllerEntity) HasController() bool {
	if o != nil && o.Controller != nil {
		return true
	}

	return false
}

// SetController gets a reference to the given ControllerDTO and assigns it to the Controller field.
func (o *ControllerEntity) SetController(v ControllerDTO) {
	o.Controller = &v
}

func (o ControllerEntity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Controller != nil {
		toSerialize["controller"] = o.Controller
	}
	return json.Marshal(toSerialize)
}

type NullableControllerEntity struct {
	value *ControllerEntity
	isSet bool
}

func (v NullableControllerEntity) Get() *ControllerEntity {
	return v.value
}

func (v *NullableControllerEntity) Set(val *ControllerEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableControllerEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableControllerEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableControllerEntity(val *ControllerEntity) *NullableControllerEntity {
	return &NullableControllerEntity{value: val, isSet: true}
}

func (v NullableControllerEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableControllerEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
