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

// UpdateControllerServiceReferenceRequestEntity struct for UpdateControllerServiceReferenceRequestEntity
type UpdateControllerServiceReferenceRequestEntity struct {
	// The identifier of the Controller Service.
	Id *string `json:"id,omitempty"`
	// The new state of the references for the controller service.
	State *string `json:"state,omitempty"`
	// The revisions for all referencing components.
	ReferencingComponentRevisions *map[string]RevisionDTO `json:"referencingComponentRevisions,omitempty"`
	// Acknowledges that this node is disconnected to allow for mutable requests to proceed.
	DisconnectedNodeAcknowledged *bool `json:"disconnectedNodeAcknowledged,omitempty"`
}

// NewUpdateControllerServiceReferenceRequestEntity instantiates a new UpdateControllerServiceReferenceRequestEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateControllerServiceReferenceRequestEntity() *UpdateControllerServiceReferenceRequestEntity {
	this := UpdateControllerServiceReferenceRequestEntity{}
	return &this
}

// NewUpdateControllerServiceReferenceRequestEntityWithDefaults instantiates a new UpdateControllerServiceReferenceRequestEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateControllerServiceReferenceRequestEntityWithDefaults() *UpdateControllerServiceReferenceRequestEntity {
	this := UpdateControllerServiceReferenceRequestEntity{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UpdateControllerServiceReferenceRequestEntity) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateControllerServiceReferenceRequestEntity) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UpdateControllerServiceReferenceRequestEntity) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *UpdateControllerServiceReferenceRequestEntity) SetId(v string) {
	o.Id = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *UpdateControllerServiceReferenceRequestEntity) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateControllerServiceReferenceRequestEntity) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *UpdateControllerServiceReferenceRequestEntity) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *UpdateControllerServiceReferenceRequestEntity) SetState(v string) {
	o.State = &v
}

// GetReferencingComponentRevisions returns the ReferencingComponentRevisions field value if set, zero value otherwise.
func (o *UpdateControllerServiceReferenceRequestEntity) GetReferencingComponentRevisions() map[string]RevisionDTO {
	if o == nil || o.ReferencingComponentRevisions == nil {
		var ret map[string]RevisionDTO
		return ret
	}
	return *o.ReferencingComponentRevisions
}

// GetReferencingComponentRevisionsOk returns a tuple with the ReferencingComponentRevisions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateControllerServiceReferenceRequestEntity) GetReferencingComponentRevisionsOk() (*map[string]RevisionDTO, bool) {
	if o == nil || o.ReferencingComponentRevisions == nil {
		return nil, false
	}
	return o.ReferencingComponentRevisions, true
}

// HasReferencingComponentRevisions returns a boolean if a field has been set.
func (o *UpdateControllerServiceReferenceRequestEntity) HasReferencingComponentRevisions() bool {
	if o != nil && o.ReferencingComponentRevisions != nil {
		return true
	}

	return false
}

// SetReferencingComponentRevisions gets a reference to the given map[string]RevisionDTO and assigns it to the ReferencingComponentRevisions field.
func (o *UpdateControllerServiceReferenceRequestEntity) SetReferencingComponentRevisions(v map[string]RevisionDTO) {
	o.ReferencingComponentRevisions = &v
}

// GetDisconnectedNodeAcknowledged returns the DisconnectedNodeAcknowledged field value if set, zero value otherwise.
func (o *UpdateControllerServiceReferenceRequestEntity) GetDisconnectedNodeAcknowledged() bool {
	if o == nil || o.DisconnectedNodeAcknowledged == nil {
		var ret bool
		return ret
	}
	return *o.DisconnectedNodeAcknowledged
}

// GetDisconnectedNodeAcknowledgedOk returns a tuple with the DisconnectedNodeAcknowledged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateControllerServiceReferenceRequestEntity) GetDisconnectedNodeAcknowledgedOk() (*bool, bool) {
	if o == nil || o.DisconnectedNodeAcknowledged == nil {
		return nil, false
	}
	return o.DisconnectedNodeAcknowledged, true
}

// HasDisconnectedNodeAcknowledged returns a boolean if a field has been set.
func (o *UpdateControllerServiceReferenceRequestEntity) HasDisconnectedNodeAcknowledged() bool {
	if o != nil && o.DisconnectedNodeAcknowledged != nil {
		return true
	}

	return false
}

// SetDisconnectedNodeAcknowledged gets a reference to the given bool and assigns it to the DisconnectedNodeAcknowledged field.
func (o *UpdateControllerServiceReferenceRequestEntity) SetDisconnectedNodeAcknowledged(v bool) {
	o.DisconnectedNodeAcknowledged = &v
}

func (o UpdateControllerServiceReferenceRequestEntity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.ReferencingComponentRevisions != nil {
		toSerialize["referencingComponentRevisions"] = o.ReferencingComponentRevisions
	}
	if o.DisconnectedNodeAcknowledged != nil {
		toSerialize["disconnectedNodeAcknowledged"] = o.DisconnectedNodeAcknowledged
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateControllerServiceReferenceRequestEntity struct {
	value *UpdateControllerServiceReferenceRequestEntity
	isSet bool
}

func (v NullableUpdateControllerServiceReferenceRequestEntity) Get() *UpdateControllerServiceReferenceRequestEntity {
	return v.value
}

func (v *NullableUpdateControllerServiceReferenceRequestEntity) Set(val *UpdateControllerServiceReferenceRequestEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateControllerServiceReferenceRequestEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateControllerServiceReferenceRequestEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateControllerServiceReferenceRequestEntity(val *UpdateControllerServiceReferenceRequestEntity) *NullableUpdateControllerServiceReferenceRequestEntity {
	return &NullableUpdateControllerServiceReferenceRequestEntity{value: val, isSet: true}
}

func (v NullableUpdateControllerServiceReferenceRequestEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateControllerServiceReferenceRequestEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
