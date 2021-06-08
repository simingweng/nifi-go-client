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

// VersionedParameter struct for VersionedParameter
type VersionedParameter struct {
	// The name of the parameter
	Name *string `json:"name,omitempty"`
	// The description of the param
	Description *string `json:"description,omitempty"`
	// Whether or not the parameter value is sensitive
	Sensitive *bool `json:"sensitive,omitempty"`
	// The value of the parameter
	Value *string `json:"value,omitempty"`
}

// NewVersionedParameter instantiates a new VersionedParameter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVersionedParameter() *VersionedParameter {
	this := VersionedParameter{}
	return &this
}

// NewVersionedParameterWithDefaults instantiates a new VersionedParameter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVersionedParameterWithDefaults() *VersionedParameter {
	this := VersionedParameter{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VersionedParameter) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionedParameter) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VersionedParameter) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VersionedParameter) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *VersionedParameter) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionedParameter) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *VersionedParameter) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *VersionedParameter) SetDescription(v string) {
	o.Description = &v
}

// GetSensitive returns the Sensitive field value if set, zero value otherwise.
func (o *VersionedParameter) GetSensitive() bool {
	if o == nil || o.Sensitive == nil {
		var ret bool
		return ret
	}
	return *o.Sensitive
}

// GetSensitiveOk returns a tuple with the Sensitive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionedParameter) GetSensitiveOk() (*bool, bool) {
	if o == nil || o.Sensitive == nil {
		return nil, false
	}
	return o.Sensitive, true
}

// HasSensitive returns a boolean if a field has been set.
func (o *VersionedParameter) HasSensitive() bool {
	if o != nil && o.Sensitive != nil {
		return true
	}

	return false
}

// SetSensitive gets a reference to the given bool and assigns it to the Sensitive field.
func (o *VersionedParameter) SetSensitive(v bool) {
	o.Sensitive = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *VersionedParameter) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionedParameter) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *VersionedParameter) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *VersionedParameter) SetValue(v string) {
	o.Value = &v
}

func (o VersionedParameter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Sensitive != nil {
		toSerialize["sensitive"] = o.Sensitive
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableVersionedParameter struct {
	value *VersionedParameter
	isSet bool
}

func (v NullableVersionedParameter) Get() *VersionedParameter {
	return v.value
}

func (v *NullableVersionedParameter) Set(val *VersionedParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableVersionedParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableVersionedParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVersionedParameter(val *VersionedParameter) *NullableVersionedParameter {
	return &NullableVersionedParameter{value: val, isSet: true}
}

func (v NullableVersionedParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVersionedParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
