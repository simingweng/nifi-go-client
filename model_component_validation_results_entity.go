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

// ComponentValidationResultsEntity struct for ComponentValidationResultsEntity
type ComponentValidationResultsEntity struct {
	// A List of ComponentValidationResultEntity, one for each component that is validated
	ValidationResults *[]ComponentValidationResultEntity `json:"validationResults,omitempty"`
}

// NewComponentValidationResultsEntity instantiates a new ComponentValidationResultsEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComponentValidationResultsEntity() *ComponentValidationResultsEntity {
	this := ComponentValidationResultsEntity{}
	return &this
}

// NewComponentValidationResultsEntityWithDefaults instantiates a new ComponentValidationResultsEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComponentValidationResultsEntityWithDefaults() *ComponentValidationResultsEntity {
	this := ComponentValidationResultsEntity{}
	return &this
}

// GetValidationResults returns the ValidationResults field value if set, zero value otherwise.
func (o *ComponentValidationResultsEntity) GetValidationResults() []ComponentValidationResultEntity {
	if o == nil || o.ValidationResults == nil {
		var ret []ComponentValidationResultEntity
		return ret
	}
	return *o.ValidationResults
}

// GetValidationResultsOk returns a tuple with the ValidationResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentValidationResultsEntity) GetValidationResultsOk() (*[]ComponentValidationResultEntity, bool) {
	if o == nil || o.ValidationResults == nil {
		return nil, false
	}
	return o.ValidationResults, true
}

// HasValidationResults returns a boolean if a field has been set.
func (o *ComponentValidationResultsEntity) HasValidationResults() bool {
	if o != nil && o.ValidationResults != nil {
		return true
	}

	return false
}

// SetValidationResults gets a reference to the given []ComponentValidationResultEntity and assigns it to the ValidationResults field.
func (o *ComponentValidationResultsEntity) SetValidationResults(v []ComponentValidationResultEntity) {
	o.ValidationResults = &v
}

func (o ComponentValidationResultsEntity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ValidationResults != nil {
		toSerialize["validationResults"] = o.ValidationResults
	}
	return json.Marshal(toSerialize)
}

type NullableComponentValidationResultsEntity struct {
	value *ComponentValidationResultsEntity
	isSet bool
}

func (v NullableComponentValidationResultsEntity) Get() *ComponentValidationResultsEntity {
	return v.value
}

func (v *NullableComponentValidationResultsEntity) Set(val *ComponentValidationResultsEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableComponentValidationResultsEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableComponentValidationResultsEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComponentValidationResultsEntity(val *ComponentValidationResultsEntity) *NullableComponentValidationResultsEntity {
	return &NullableComponentValidationResultsEntity{value: val, isSet: true}
}

func (v NullableComponentValidationResultsEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComponentValidationResultsEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
