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

// CreateActiveRequestEntity struct for CreateActiveRequestEntity
type CreateActiveRequestEntity struct {
	// The Process Group ID that this active request will update
	ProcessGroupId *string `json:"processGroupId,omitempty"`
	// Acknowledges that this node is disconnected to allow for mutable requests to proceed.
	DisconnectedNodeAcknowledged *bool `json:"disconnectedNodeAcknowledged,omitempty"`
}

// NewCreateActiveRequestEntity instantiates a new CreateActiveRequestEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateActiveRequestEntity() *CreateActiveRequestEntity {
	this := CreateActiveRequestEntity{}
	return &this
}

// NewCreateActiveRequestEntityWithDefaults instantiates a new CreateActiveRequestEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateActiveRequestEntityWithDefaults() *CreateActiveRequestEntity {
	this := CreateActiveRequestEntity{}
	return &this
}

// GetProcessGroupId returns the ProcessGroupId field value if set, zero value otherwise.
func (o *CreateActiveRequestEntity) GetProcessGroupId() string {
	if o == nil || o.ProcessGroupId == nil {
		var ret string
		return ret
	}
	return *o.ProcessGroupId
}

// GetProcessGroupIdOk returns a tuple with the ProcessGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateActiveRequestEntity) GetProcessGroupIdOk() (*string, bool) {
	if o == nil || o.ProcessGroupId == nil {
		return nil, false
	}
	return o.ProcessGroupId, true
}

// HasProcessGroupId returns a boolean if a field has been set.
func (o *CreateActiveRequestEntity) HasProcessGroupId() bool {
	if o != nil && o.ProcessGroupId != nil {
		return true
	}

	return false
}

// SetProcessGroupId gets a reference to the given string and assigns it to the ProcessGroupId field.
func (o *CreateActiveRequestEntity) SetProcessGroupId(v string) {
	o.ProcessGroupId = &v
}

// GetDisconnectedNodeAcknowledged returns the DisconnectedNodeAcknowledged field value if set, zero value otherwise.
func (o *CreateActiveRequestEntity) GetDisconnectedNodeAcknowledged() bool {
	if o == nil || o.DisconnectedNodeAcknowledged == nil {
		var ret bool
		return ret
	}
	return *o.DisconnectedNodeAcknowledged
}

// GetDisconnectedNodeAcknowledgedOk returns a tuple with the DisconnectedNodeAcknowledged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateActiveRequestEntity) GetDisconnectedNodeAcknowledgedOk() (*bool, bool) {
	if o == nil || o.DisconnectedNodeAcknowledged == nil {
		return nil, false
	}
	return o.DisconnectedNodeAcknowledged, true
}

// HasDisconnectedNodeAcknowledged returns a boolean if a field has been set.
func (o *CreateActiveRequestEntity) HasDisconnectedNodeAcknowledged() bool {
	if o != nil && o.DisconnectedNodeAcknowledged != nil {
		return true
	}

	return false
}

// SetDisconnectedNodeAcknowledged gets a reference to the given bool and assigns it to the DisconnectedNodeAcknowledged field.
func (o *CreateActiveRequestEntity) SetDisconnectedNodeAcknowledged(v bool) {
	o.DisconnectedNodeAcknowledged = &v
}

func (o CreateActiveRequestEntity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ProcessGroupId != nil {
		toSerialize["processGroupId"] = o.ProcessGroupId
	}
	if o.DisconnectedNodeAcknowledged != nil {
		toSerialize["disconnectedNodeAcknowledged"] = o.DisconnectedNodeAcknowledged
	}
	return json.Marshal(toSerialize)
}

type NullableCreateActiveRequestEntity struct {
	value *CreateActiveRequestEntity
	isSet bool
}

func (v NullableCreateActiveRequestEntity) Get() *CreateActiveRequestEntity {
	return v.value
}

func (v *NullableCreateActiveRequestEntity) Set(val *CreateActiveRequestEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateActiveRequestEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateActiveRequestEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateActiveRequestEntity(val *CreateActiveRequestEntity) *NullableCreateActiveRequestEntity {
	return &NullableCreateActiveRequestEntity{value: val, isSet: true}
}

func (v NullableCreateActiveRequestEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateActiveRequestEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
