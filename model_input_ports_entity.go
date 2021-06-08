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

// InputPortsEntity struct for InputPortsEntity
type InputPortsEntity struct {
	InputPorts *[]PortEntity `json:"inputPorts,omitempty"`
}

// NewInputPortsEntity instantiates a new InputPortsEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInputPortsEntity() *InputPortsEntity {
	this := InputPortsEntity{}
	return &this
}

// NewInputPortsEntityWithDefaults instantiates a new InputPortsEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInputPortsEntityWithDefaults() *InputPortsEntity {
	this := InputPortsEntity{}
	return &this
}

// GetInputPorts returns the InputPorts field value if set, zero value otherwise.
func (o *InputPortsEntity) GetInputPorts() []PortEntity {
	if o == nil || o.InputPorts == nil {
		var ret []PortEntity
		return ret
	}
	return *o.InputPorts
}

// GetInputPortsOk returns a tuple with the InputPorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputPortsEntity) GetInputPortsOk() (*[]PortEntity, bool) {
	if o == nil || o.InputPorts == nil {
		return nil, false
	}
	return o.InputPorts, true
}

// HasInputPorts returns a boolean if a field has been set.
func (o *InputPortsEntity) HasInputPorts() bool {
	if o != nil && o.InputPorts != nil {
		return true
	}

	return false
}

// SetInputPorts gets a reference to the given []PortEntity and assigns it to the InputPorts field.
func (o *InputPortsEntity) SetInputPorts(v []PortEntity) {
	o.InputPorts = &v
}

func (o InputPortsEntity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.InputPorts != nil {
		toSerialize["inputPorts"] = o.InputPorts
	}
	return json.Marshal(toSerialize)
}

type NullableInputPortsEntity struct {
	value *InputPortsEntity
	isSet bool
}

func (v NullableInputPortsEntity) Get() *InputPortsEntity {
	return v.value
}

func (v *NullableInputPortsEntity) Set(val *InputPortsEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableInputPortsEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableInputPortsEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInputPortsEntity(val *InputPortsEntity) *NullableInputPortsEntity {
	return &NullableInputPortsEntity{value: val, isSet: true}
}

func (v NullableInputPortsEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInputPortsEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
