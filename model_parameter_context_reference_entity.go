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

// ParameterContextReferenceEntity struct for ParameterContextReferenceEntity
type ParameterContextReferenceEntity struct {
	// The id of the component.
	Id          *string                       `json:"id,omitempty"`
	Permissions *PermissionsDTO               `json:"permissions,omitempty"`
	Component   *ParameterContextReferenceDTO `json:"component,omitempty"`
}

// NewParameterContextReferenceEntity instantiates a new ParameterContextReferenceEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParameterContextReferenceEntity() *ParameterContextReferenceEntity {
	this := ParameterContextReferenceEntity{}
	return &this
}

// NewParameterContextReferenceEntityWithDefaults instantiates a new ParameterContextReferenceEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParameterContextReferenceEntityWithDefaults() *ParameterContextReferenceEntity {
	this := ParameterContextReferenceEntity{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ParameterContextReferenceEntity) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterContextReferenceEntity) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ParameterContextReferenceEntity) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ParameterContextReferenceEntity) SetId(v string) {
	o.Id = &v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *ParameterContextReferenceEntity) GetPermissions() PermissionsDTO {
	if o == nil || o.Permissions == nil {
		var ret PermissionsDTO
		return ret
	}
	return *o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterContextReferenceEntity) GetPermissionsOk() (*PermissionsDTO, bool) {
	if o == nil || o.Permissions == nil {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *ParameterContextReferenceEntity) HasPermissions() bool {
	if o != nil && o.Permissions != nil {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given PermissionsDTO and assigns it to the Permissions field.
func (o *ParameterContextReferenceEntity) SetPermissions(v PermissionsDTO) {
	o.Permissions = &v
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *ParameterContextReferenceEntity) GetComponent() ParameterContextReferenceDTO {
	if o == nil || o.Component == nil {
		var ret ParameterContextReferenceDTO
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterContextReferenceEntity) GetComponentOk() (*ParameterContextReferenceDTO, bool) {
	if o == nil || o.Component == nil {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *ParameterContextReferenceEntity) HasComponent() bool {
	if o != nil && o.Component != nil {
		return true
	}

	return false
}

// SetComponent gets a reference to the given ParameterContextReferenceDTO and assigns it to the Component field.
func (o *ParameterContextReferenceEntity) SetComponent(v ParameterContextReferenceDTO) {
	o.Component = &v
}

func (o ParameterContextReferenceEntity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Permissions != nil {
		toSerialize["permissions"] = o.Permissions
	}
	if o.Component != nil {
		toSerialize["component"] = o.Component
	}
	return json.Marshal(toSerialize)
}

type NullableParameterContextReferenceEntity struct {
	value *ParameterContextReferenceEntity
	isSet bool
}

func (v NullableParameterContextReferenceEntity) Get() *ParameterContextReferenceEntity {
	return v.value
}

func (v *NullableParameterContextReferenceEntity) Set(val *ParameterContextReferenceEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableParameterContextReferenceEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableParameterContextReferenceEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParameterContextReferenceEntity(val *ParameterContextReferenceEntity) *NullableParameterContextReferenceEntity {
	return &NullableParameterContextReferenceEntity{value: val, isSet: true}
}

func (v NullableParameterContextReferenceEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParameterContextReferenceEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
