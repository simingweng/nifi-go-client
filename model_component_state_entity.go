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

// ComponentStateEntity struct for ComponentStateEntity
type ComponentStateEntity struct {
	ComponentState *ComponentStateDTO `json:"componentState,omitempty"`
}

// NewComponentStateEntity instantiates a new ComponentStateEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComponentStateEntity() *ComponentStateEntity {
	this := ComponentStateEntity{}
	return &this
}

// NewComponentStateEntityWithDefaults instantiates a new ComponentStateEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComponentStateEntityWithDefaults() *ComponentStateEntity {
	this := ComponentStateEntity{}
	return &this
}

// GetComponentState returns the ComponentState field value if set, zero value otherwise.
func (o *ComponentStateEntity) GetComponentState() ComponentStateDTO {
	if o == nil || o.ComponentState == nil {
		var ret ComponentStateDTO
		return ret
	}
	return *o.ComponentState
}

// GetComponentStateOk returns a tuple with the ComponentState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentStateEntity) GetComponentStateOk() (*ComponentStateDTO, bool) {
	if o == nil || o.ComponentState == nil {
		return nil, false
	}
	return o.ComponentState, true
}

// HasComponentState returns a boolean if a field has been set.
func (o *ComponentStateEntity) HasComponentState() bool {
	if o != nil && o.ComponentState != nil {
		return true
	}

	return false
}

// SetComponentState gets a reference to the given ComponentStateDTO and assigns it to the ComponentState field.
func (o *ComponentStateEntity) SetComponentState(v ComponentStateDTO) {
	o.ComponentState = &v
}

func (o ComponentStateEntity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ComponentState != nil {
		toSerialize["componentState"] = o.ComponentState
	}
	return json.Marshal(toSerialize)
}

type NullableComponentStateEntity struct {
	value *ComponentStateEntity
	isSet bool
}

func (v NullableComponentStateEntity) Get() *ComponentStateEntity {
	return v.value
}

func (v *NullableComponentStateEntity) Set(val *ComponentStateEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableComponentStateEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableComponentStateEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComponentStateEntity(val *ComponentStateEntity) *NullableComponentStateEntity {
	return &NullableComponentStateEntity{value: val, isSet: true}
}

func (v NullableComponentStateEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComponentStateEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
