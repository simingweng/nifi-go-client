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

// ProvenanceOptionsDTO struct for ProvenanceOptionsDTO
type ProvenanceOptionsDTO struct {
	// The available searchable field for the NiFi.
	SearchableFields *[]ProvenanceSearchableFieldDTO `json:"searchableFields,omitempty"`
}

// NewProvenanceOptionsDTO instantiates a new ProvenanceOptionsDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProvenanceOptionsDTO() *ProvenanceOptionsDTO {
	this := ProvenanceOptionsDTO{}
	return &this
}

// NewProvenanceOptionsDTOWithDefaults instantiates a new ProvenanceOptionsDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProvenanceOptionsDTOWithDefaults() *ProvenanceOptionsDTO {
	this := ProvenanceOptionsDTO{}
	return &this
}

// GetSearchableFields returns the SearchableFields field value if set, zero value otherwise.
func (o *ProvenanceOptionsDTO) GetSearchableFields() []ProvenanceSearchableFieldDTO {
	if o == nil || o.SearchableFields == nil {
		var ret []ProvenanceSearchableFieldDTO
		return ret
	}
	return *o.SearchableFields
}

// GetSearchableFieldsOk returns a tuple with the SearchableFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvenanceOptionsDTO) GetSearchableFieldsOk() (*[]ProvenanceSearchableFieldDTO, bool) {
	if o == nil || o.SearchableFields == nil {
		return nil, false
	}
	return o.SearchableFields, true
}

// HasSearchableFields returns a boolean if a field has been set.
func (o *ProvenanceOptionsDTO) HasSearchableFields() bool {
	if o != nil && o.SearchableFields != nil {
		return true
	}

	return false
}

// SetSearchableFields gets a reference to the given []ProvenanceSearchableFieldDTO and assigns it to the SearchableFields field.
func (o *ProvenanceOptionsDTO) SetSearchableFields(v []ProvenanceSearchableFieldDTO) {
	o.SearchableFields = &v
}

func (o ProvenanceOptionsDTO) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SearchableFields != nil {
		toSerialize["searchableFields"] = o.SearchableFields
	}
	return json.Marshal(toSerialize)
}

type NullableProvenanceOptionsDTO struct {
	value *ProvenanceOptionsDTO
	isSet bool
}

func (v NullableProvenanceOptionsDTO) Get() *ProvenanceOptionsDTO {
	return v.value
}

func (v *NullableProvenanceOptionsDTO) Set(val *ProvenanceOptionsDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableProvenanceOptionsDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableProvenanceOptionsDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProvenanceOptionsDTO(val *ProvenanceOptionsDTO) *NullableProvenanceOptionsDTO {
	return &NullableProvenanceOptionsDTO{value: val, isSet: true}
}

func (v NullableProvenanceOptionsDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProvenanceOptionsDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
