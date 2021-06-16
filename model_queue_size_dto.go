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

// QueueSizeDTO struct for QueueSizeDTO
type QueueSizeDTO struct {
	// The size of objects in a queue.
	ByteCount *int64 `json:"byteCount,omitempty"`
	// The count of objects in a queue.
	ObjectCount *int32 `json:"objectCount,omitempty"`
}

// NewQueueSizeDTO instantiates a new QueueSizeDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueueSizeDTO() *QueueSizeDTO {
	this := QueueSizeDTO{}
	return &this
}

// NewQueueSizeDTOWithDefaults instantiates a new QueueSizeDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueueSizeDTOWithDefaults() *QueueSizeDTO {
	this := QueueSizeDTO{}
	return &this
}

// GetByteCount returns the ByteCount field value if set, zero value otherwise.
func (o *QueueSizeDTO) GetByteCount() int64 {
	if o == nil || o.ByteCount == nil {
		var ret int64
		return ret
	}
	return *o.ByteCount
}

// GetByteCountOk returns a tuple with the ByteCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueueSizeDTO) GetByteCountOk() (*int64, bool) {
	if o == nil || o.ByteCount == nil {
		return nil, false
	}
	return o.ByteCount, true
}

// HasByteCount returns a boolean if a field has been set.
func (o *QueueSizeDTO) HasByteCount() bool {
	if o != nil && o.ByteCount != nil {
		return true
	}

	return false
}

// SetByteCount gets a reference to the given int64 and assigns it to the ByteCount field.
func (o *QueueSizeDTO) SetByteCount(v int64) {
	o.ByteCount = &v
}

// GetObjectCount returns the ObjectCount field value if set, zero value otherwise.
func (o *QueueSizeDTO) GetObjectCount() int32 {
	if o == nil || o.ObjectCount == nil {
		var ret int32
		return ret
	}
	return *o.ObjectCount
}

// GetObjectCountOk returns a tuple with the ObjectCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueueSizeDTO) GetObjectCountOk() (*int32, bool) {
	if o == nil || o.ObjectCount == nil {
		return nil, false
	}
	return o.ObjectCount, true
}

// HasObjectCount returns a boolean if a field has been set.
func (o *QueueSizeDTO) HasObjectCount() bool {
	if o != nil && o.ObjectCount != nil {
		return true
	}

	return false
}

// SetObjectCount gets a reference to the given int32 and assigns it to the ObjectCount field.
func (o *QueueSizeDTO) SetObjectCount(v int32) {
	o.ObjectCount = &v
}

func (o QueueSizeDTO) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ByteCount != nil {
		toSerialize["byteCount"] = o.ByteCount
	}
	if o.ObjectCount != nil {
		toSerialize["objectCount"] = o.ObjectCount
	}
	return json.Marshal(toSerialize)
}

type NullableQueueSizeDTO struct {
	value *QueueSizeDTO
	isSet bool
}

func (v NullableQueueSizeDTO) Get() *QueueSizeDTO {
	return v.value
}

func (v *NullableQueueSizeDTO) Set(val *QueueSizeDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableQueueSizeDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableQueueSizeDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueueSizeDTO(val *QueueSizeDTO) *NullableQueueSizeDTO {
	return &NullableQueueSizeDTO{value: val, isSet: true}
}

func (v NullableQueueSizeDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueueSizeDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
