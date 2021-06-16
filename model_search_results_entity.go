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

// SearchResultsEntity struct for SearchResultsEntity
type SearchResultsEntity struct {
	SearchResultsDTO *SearchResultsDTO `json:"searchResultsDTO,omitempty"`
}

// NewSearchResultsEntity instantiates a new SearchResultsEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchResultsEntity() *SearchResultsEntity {
	this := SearchResultsEntity{}
	return &this
}

// NewSearchResultsEntityWithDefaults instantiates a new SearchResultsEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchResultsEntityWithDefaults() *SearchResultsEntity {
	this := SearchResultsEntity{}
	return &this
}

// GetSearchResultsDTO returns the SearchResultsDTO field value if set, zero value otherwise.
func (o *SearchResultsEntity) GetSearchResultsDTO() SearchResultsDTO {
	if o == nil || o.SearchResultsDTO == nil {
		var ret SearchResultsDTO
		return ret
	}
	return *o.SearchResultsDTO
}

// GetSearchResultsDTOOk returns a tuple with the SearchResultsDTO field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResultsEntity) GetSearchResultsDTOOk() (*SearchResultsDTO, bool) {
	if o == nil || o.SearchResultsDTO == nil {
		return nil, false
	}
	return o.SearchResultsDTO, true
}

// HasSearchResultsDTO returns a boolean if a field has been set.
func (o *SearchResultsEntity) HasSearchResultsDTO() bool {
	if o != nil && o.SearchResultsDTO != nil {
		return true
	}

	return false
}

// SetSearchResultsDTO gets a reference to the given SearchResultsDTO and assigns it to the SearchResultsDTO field.
func (o *SearchResultsEntity) SetSearchResultsDTO(v SearchResultsDTO) {
	o.SearchResultsDTO = &v
}

func (o SearchResultsEntity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SearchResultsDTO != nil {
		toSerialize["searchResultsDTO"] = o.SearchResultsDTO
	}
	return json.Marshal(toSerialize)
}

type NullableSearchResultsEntity struct {
	value *SearchResultsEntity
	isSet bool
}

func (v NullableSearchResultsEntity) Get() *SearchResultsEntity {
	return v.value
}

func (v *NullableSearchResultsEntity) Set(val *SearchResultsEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchResultsEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchResultsEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchResultsEntity(val *SearchResultsEntity) *NullableSearchResultsEntity {
	return &NullableSearchResultsEntity{value: val, isSet: true}
}

func (v NullableSearchResultsEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchResultsEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
