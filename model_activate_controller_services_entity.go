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

// ActivateControllerServicesEntity struct for ActivateControllerServicesEntity
type ActivateControllerServicesEntity struct {
	// The id of the ProcessGroup
	Id *string `json:"id,omitempty"`
	// The desired state of the descendant components
	State *string `json:"state,omitempty"`
	// Optional services to schedule. If not specified, all authorized descendant controller services will be used.
	Components *map[string]RevisionDTO `json:"components,omitempty"`
	// Acknowledges that this node is disconnected to allow for mutable requests to proceed.
	DisconnectedNodeAcknowledged *bool `json:"disconnectedNodeAcknowledged,omitempty"`
}

// NewActivateControllerServicesEntity instantiates a new ActivateControllerServicesEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActivateControllerServicesEntity() *ActivateControllerServicesEntity {
	this := ActivateControllerServicesEntity{}
	return &this
}

// NewActivateControllerServicesEntityWithDefaults instantiates a new ActivateControllerServicesEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActivateControllerServicesEntityWithDefaults() *ActivateControllerServicesEntity {
	this := ActivateControllerServicesEntity{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ActivateControllerServicesEntity) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivateControllerServicesEntity) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ActivateControllerServicesEntity) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ActivateControllerServicesEntity) SetId(v string) {
	o.Id = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *ActivateControllerServicesEntity) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivateControllerServicesEntity) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *ActivateControllerServicesEntity) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *ActivateControllerServicesEntity) SetState(v string) {
	o.State = &v
}

// GetComponents returns the Components field value if set, zero value otherwise.
func (o *ActivateControllerServicesEntity) GetComponents() map[string]RevisionDTO {
	if o == nil || o.Components == nil {
		var ret map[string]RevisionDTO
		return ret
	}
	return *o.Components
}

// GetComponentsOk returns a tuple with the Components field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivateControllerServicesEntity) GetComponentsOk() (*map[string]RevisionDTO, bool) {
	if o == nil || o.Components == nil {
		return nil, false
	}
	return o.Components, true
}

// HasComponents returns a boolean if a field has been set.
func (o *ActivateControllerServicesEntity) HasComponents() bool {
	if o != nil && o.Components != nil {
		return true
	}

	return false
}

// SetComponents gets a reference to the given map[string]RevisionDTO and assigns it to the Components field.
func (o *ActivateControllerServicesEntity) SetComponents(v map[string]RevisionDTO) {
	o.Components = &v
}

// GetDisconnectedNodeAcknowledged returns the DisconnectedNodeAcknowledged field value if set, zero value otherwise.
func (o *ActivateControllerServicesEntity) GetDisconnectedNodeAcknowledged() bool {
	if o == nil || o.DisconnectedNodeAcknowledged == nil {
		var ret bool
		return ret
	}
	return *o.DisconnectedNodeAcknowledged
}

// GetDisconnectedNodeAcknowledgedOk returns a tuple with the DisconnectedNodeAcknowledged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivateControllerServicesEntity) GetDisconnectedNodeAcknowledgedOk() (*bool, bool) {
	if o == nil || o.DisconnectedNodeAcknowledged == nil {
		return nil, false
	}
	return o.DisconnectedNodeAcknowledged, true
}

// HasDisconnectedNodeAcknowledged returns a boolean if a field has been set.
func (o *ActivateControllerServicesEntity) HasDisconnectedNodeAcknowledged() bool {
	if o != nil && o.DisconnectedNodeAcknowledged != nil {
		return true
	}

	return false
}

// SetDisconnectedNodeAcknowledged gets a reference to the given bool and assigns it to the DisconnectedNodeAcknowledged field.
func (o *ActivateControllerServicesEntity) SetDisconnectedNodeAcknowledged(v bool) {
	o.DisconnectedNodeAcknowledged = &v
}

func (o ActivateControllerServicesEntity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.Components != nil {
		toSerialize["components"] = o.Components
	}
	if o.DisconnectedNodeAcknowledged != nil {
		toSerialize["disconnectedNodeAcknowledged"] = o.DisconnectedNodeAcknowledged
	}
	return json.Marshal(toSerialize)
}

type NullableActivateControllerServicesEntity struct {
	value *ActivateControllerServicesEntity
	isSet bool
}

func (v NullableActivateControllerServicesEntity) Get() *ActivateControllerServicesEntity {
	return v.value
}

func (v *NullableActivateControllerServicesEntity) Set(val *ActivateControllerServicesEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableActivateControllerServicesEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableActivateControllerServicesEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActivateControllerServicesEntity(val *ActivateControllerServicesEntity) *NullableActivateControllerServicesEntity {
	return &NullableActivateControllerServicesEntity{value: val, isSet: true}
}

func (v NullableActivateControllerServicesEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActivateControllerServicesEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
