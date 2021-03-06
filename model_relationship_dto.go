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

// RelationshipDTO struct for RelationshipDTO
type RelationshipDTO struct {
	// The relationship name.
	Name *string `json:"name,omitempty"`
	// The relationship description.
	Description *string `json:"description,omitempty"`
	// Whether or not flowfiles sent to this relationship should auto terminate.
	AutoTerminate *bool `json:"autoTerminate,omitempty"`
}

// NewRelationshipDTO instantiates a new RelationshipDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationshipDTO() *RelationshipDTO {
	this := RelationshipDTO{}
	return &this
}

// NewRelationshipDTOWithDefaults instantiates a new RelationshipDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationshipDTOWithDefaults() *RelationshipDTO {
	this := RelationshipDTO{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RelationshipDTO) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipDTO) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RelationshipDTO) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RelationshipDTO) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RelationshipDTO) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipDTO) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RelationshipDTO) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RelationshipDTO) SetDescription(v string) {
	o.Description = &v
}

// GetAutoTerminate returns the AutoTerminate field value if set, zero value otherwise.
func (o *RelationshipDTO) GetAutoTerminate() bool {
	if o == nil || o.AutoTerminate == nil {
		var ret bool
		return ret
	}
	return *o.AutoTerminate
}

// GetAutoTerminateOk returns a tuple with the AutoTerminate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipDTO) GetAutoTerminateOk() (*bool, bool) {
	if o == nil || o.AutoTerminate == nil {
		return nil, false
	}
	return o.AutoTerminate, true
}

// HasAutoTerminate returns a boolean if a field has been set.
func (o *RelationshipDTO) HasAutoTerminate() bool {
	if o != nil && o.AutoTerminate != nil {
		return true
	}

	return false
}

// SetAutoTerminate gets a reference to the given bool and assigns it to the AutoTerminate field.
func (o *RelationshipDTO) SetAutoTerminate(v bool) {
	o.AutoTerminate = &v
}

func (o RelationshipDTO) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.AutoTerminate != nil {
		toSerialize["autoTerminate"] = o.AutoTerminate
	}
	return json.Marshal(toSerialize)
}

type NullableRelationshipDTO struct {
	value *RelationshipDTO
	isSet bool
}

func (v NullableRelationshipDTO) Get() *RelationshipDTO {
	return v.value
}

func (v *NullableRelationshipDTO) Set(val *RelationshipDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableRelationshipDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableRelationshipDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRelationshipDTO(val *RelationshipDTO) *NullableRelationshipDTO {
	return &NullableRelationshipDTO{value: val, isSet: true}
}

func (v NullableRelationshipDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRelationshipDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
