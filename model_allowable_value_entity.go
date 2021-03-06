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

// AllowableValueEntity struct for AllowableValueEntity
type AllowableValueEntity struct {
	AllowableValue *AllowableValueDTO `json:"allowableValue,omitempty"`
	// Indicates whether the user can read a given resource.
	CanRead *bool `json:"canRead,omitempty"`
}

// NewAllowableValueEntity instantiates a new AllowableValueEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAllowableValueEntity() *AllowableValueEntity {
	this := AllowableValueEntity{}
	return &this
}

// NewAllowableValueEntityWithDefaults instantiates a new AllowableValueEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAllowableValueEntityWithDefaults() *AllowableValueEntity {
	this := AllowableValueEntity{}
	return &this
}

// GetAllowableValue returns the AllowableValue field value if set, zero value otherwise.
func (o *AllowableValueEntity) GetAllowableValue() AllowableValueDTO {
	if o == nil || o.AllowableValue == nil {
		var ret AllowableValueDTO
		return ret
	}
	return *o.AllowableValue
}

// GetAllowableValueOk returns a tuple with the AllowableValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllowableValueEntity) GetAllowableValueOk() (*AllowableValueDTO, bool) {
	if o == nil || o.AllowableValue == nil {
		return nil, false
	}
	return o.AllowableValue, true
}

// HasAllowableValue returns a boolean if a field has been set.
func (o *AllowableValueEntity) HasAllowableValue() bool {
	if o != nil && o.AllowableValue != nil {
		return true
	}

	return false
}

// SetAllowableValue gets a reference to the given AllowableValueDTO and assigns it to the AllowableValue field.
func (o *AllowableValueEntity) SetAllowableValue(v AllowableValueDTO) {
	o.AllowableValue = &v
}

// GetCanRead returns the CanRead field value if set, zero value otherwise.
func (o *AllowableValueEntity) GetCanRead() bool {
	if o == nil || o.CanRead == nil {
		var ret bool
		return ret
	}
	return *o.CanRead
}

// GetCanReadOk returns a tuple with the CanRead field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllowableValueEntity) GetCanReadOk() (*bool, bool) {
	if o == nil || o.CanRead == nil {
		return nil, false
	}
	return o.CanRead, true
}

// HasCanRead returns a boolean if a field has been set.
func (o *AllowableValueEntity) HasCanRead() bool {
	if o != nil && o.CanRead != nil {
		return true
	}

	return false
}

// SetCanRead gets a reference to the given bool and assigns it to the CanRead field.
func (o *AllowableValueEntity) SetCanRead(v bool) {
	o.CanRead = &v
}

func (o AllowableValueEntity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AllowableValue != nil {
		toSerialize["allowableValue"] = o.AllowableValue
	}
	if o.CanRead != nil {
		toSerialize["canRead"] = o.CanRead
	}
	return json.Marshal(toSerialize)
}

type NullableAllowableValueEntity struct {
	value *AllowableValueEntity
	isSet bool
}

func (v NullableAllowableValueEntity) Get() *AllowableValueEntity {
	return v.value
}

func (v *NullableAllowableValueEntity) Set(val *AllowableValueEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableAllowableValueEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableAllowableValueEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAllowableValueEntity(val *AllowableValueEntity) *NullableAllowableValueEntity {
	return &NullableAllowableValueEntity{value: val, isSet: true}
}

func (v NullableAllowableValueEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAllowableValueEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
