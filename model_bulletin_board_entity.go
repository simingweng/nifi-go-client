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

// BulletinBoardEntity struct for BulletinBoardEntity
type BulletinBoardEntity struct {
	BulletinBoard *BulletinBoardDTO `json:"bulletinBoard,omitempty"`
}

// NewBulletinBoardEntity instantiates a new BulletinBoardEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulletinBoardEntity() *BulletinBoardEntity {
	this := BulletinBoardEntity{}
	return &this
}

// NewBulletinBoardEntityWithDefaults instantiates a new BulletinBoardEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulletinBoardEntityWithDefaults() *BulletinBoardEntity {
	this := BulletinBoardEntity{}
	return &this
}

// GetBulletinBoard returns the BulletinBoard field value if set, zero value otherwise.
func (o *BulletinBoardEntity) GetBulletinBoard() BulletinBoardDTO {
	if o == nil || o.BulletinBoard == nil {
		var ret BulletinBoardDTO
		return ret
	}
	return *o.BulletinBoard
}

// GetBulletinBoardOk returns a tuple with the BulletinBoard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulletinBoardEntity) GetBulletinBoardOk() (*BulletinBoardDTO, bool) {
	if o == nil || o.BulletinBoard == nil {
		return nil, false
	}
	return o.BulletinBoard, true
}

// HasBulletinBoard returns a boolean if a field has been set.
func (o *BulletinBoardEntity) HasBulletinBoard() bool {
	if o != nil && o.BulletinBoard != nil {
		return true
	}

	return false
}

// SetBulletinBoard gets a reference to the given BulletinBoardDTO and assigns it to the BulletinBoard field.
func (o *BulletinBoardEntity) SetBulletinBoard(v BulletinBoardDTO) {
	o.BulletinBoard = &v
}

func (o BulletinBoardEntity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BulletinBoard != nil {
		toSerialize["bulletinBoard"] = o.BulletinBoard
	}
	return json.Marshal(toSerialize)
}

type NullableBulletinBoardEntity struct {
	value *BulletinBoardEntity
	isSet bool
}

func (v NullableBulletinBoardEntity) Get() *BulletinBoardEntity {
	return v.value
}

func (v *NullableBulletinBoardEntity) Set(val *BulletinBoardEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableBulletinBoardEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableBulletinBoardEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulletinBoardEntity(val *BulletinBoardEntity) *NullableBulletinBoardEntity {
	return &NullableBulletinBoardEntity{value: val, isSet: true}
}

func (v NullableBulletinBoardEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulletinBoardEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
