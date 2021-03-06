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

// RemoteProcessGroupsEntity struct for RemoteProcessGroupsEntity
type RemoteProcessGroupsEntity struct {
	RemoteProcessGroups *[]RemoteProcessGroupEntity `json:"remoteProcessGroups,omitempty"`
}

// NewRemoteProcessGroupsEntity instantiates a new RemoteProcessGroupsEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemoteProcessGroupsEntity() *RemoteProcessGroupsEntity {
	this := RemoteProcessGroupsEntity{}
	return &this
}

// NewRemoteProcessGroupsEntityWithDefaults instantiates a new RemoteProcessGroupsEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemoteProcessGroupsEntityWithDefaults() *RemoteProcessGroupsEntity {
	this := RemoteProcessGroupsEntity{}
	return &this
}

// GetRemoteProcessGroups returns the RemoteProcessGroups field value if set, zero value otherwise.
func (o *RemoteProcessGroupsEntity) GetRemoteProcessGroups() []RemoteProcessGroupEntity {
	if o == nil || o.RemoteProcessGroups == nil {
		var ret []RemoteProcessGroupEntity
		return ret
	}
	return *o.RemoteProcessGroups
}

// GetRemoteProcessGroupsOk returns a tuple with the RemoteProcessGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteProcessGroupsEntity) GetRemoteProcessGroupsOk() (*[]RemoteProcessGroupEntity, bool) {
	if o == nil || o.RemoteProcessGroups == nil {
		return nil, false
	}
	return o.RemoteProcessGroups, true
}

// HasRemoteProcessGroups returns a boolean if a field has been set.
func (o *RemoteProcessGroupsEntity) HasRemoteProcessGroups() bool {
	if o != nil && o.RemoteProcessGroups != nil {
		return true
	}

	return false
}

// SetRemoteProcessGroups gets a reference to the given []RemoteProcessGroupEntity and assigns it to the RemoteProcessGroups field.
func (o *RemoteProcessGroupsEntity) SetRemoteProcessGroups(v []RemoteProcessGroupEntity) {
	o.RemoteProcessGroups = &v
}

func (o RemoteProcessGroupsEntity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RemoteProcessGroups != nil {
		toSerialize["remoteProcessGroups"] = o.RemoteProcessGroups
	}
	return json.Marshal(toSerialize)
}

type NullableRemoteProcessGroupsEntity struct {
	value *RemoteProcessGroupsEntity
	isSet bool
}

func (v NullableRemoteProcessGroupsEntity) Get() *RemoteProcessGroupsEntity {
	return v.value
}

func (v *NullableRemoteProcessGroupsEntity) Set(val *RemoteProcessGroupsEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableRemoteProcessGroupsEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableRemoteProcessGroupsEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemoteProcessGroupsEntity(val *RemoteProcessGroupsEntity) *NullableRemoteProcessGroupsEntity {
	return &NullableRemoteProcessGroupsEntity{value: val, isSet: true}
}

func (v NullableRemoteProcessGroupsEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemoteProcessGroupsEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
