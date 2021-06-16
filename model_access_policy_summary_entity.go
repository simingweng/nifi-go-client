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

// AccessPolicySummaryEntity struct for AccessPolicySummaryEntity
type AccessPolicySummaryEntity struct {
	Revision *RevisionDTO `json:"revision,omitempty"`
	// The id of the component.
	Id *string `json:"id,omitempty"`
	// The URI for futures requests to the component.
	Uri         *string         `json:"uri,omitempty"`
	Position    *PositionDTO    `json:"position,omitempty"`
	Permissions *PermissionsDTO `json:"permissions,omitempty"`
	// The bulletins for this component.
	Bulletins *[]BulletinEntity `json:"bulletins,omitempty"`
	// Acknowledges that this node is disconnected to allow for mutable requests to proceed.
	DisconnectedNodeAcknowledged *bool                   `json:"disconnectedNodeAcknowledged,omitempty"`
	Component                    *AccessPolicySummaryDTO `json:"component,omitempty"`
}

// NewAccessPolicySummaryEntity instantiates a new AccessPolicySummaryEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessPolicySummaryEntity() *AccessPolicySummaryEntity {
	this := AccessPolicySummaryEntity{}
	return &this
}

// NewAccessPolicySummaryEntityWithDefaults instantiates a new AccessPolicySummaryEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessPolicySummaryEntityWithDefaults() *AccessPolicySummaryEntity {
	this := AccessPolicySummaryEntity{}
	return &this
}

// GetRevision returns the Revision field value if set, zero value otherwise.
func (o *AccessPolicySummaryEntity) GetRevision() RevisionDTO {
	if o == nil || o.Revision == nil {
		var ret RevisionDTO
		return ret
	}
	return *o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessPolicySummaryEntity) GetRevisionOk() (*RevisionDTO, bool) {
	if o == nil || o.Revision == nil {
		return nil, false
	}
	return o.Revision, true
}

// HasRevision returns a boolean if a field has been set.
func (o *AccessPolicySummaryEntity) HasRevision() bool {
	if o != nil && o.Revision != nil {
		return true
	}

	return false
}

// SetRevision gets a reference to the given RevisionDTO and assigns it to the Revision field.
func (o *AccessPolicySummaryEntity) SetRevision(v RevisionDTO) {
	o.Revision = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AccessPolicySummaryEntity) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessPolicySummaryEntity) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AccessPolicySummaryEntity) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AccessPolicySummaryEntity) SetId(v string) {
	o.Id = &v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *AccessPolicySummaryEntity) GetUri() string {
	if o == nil || o.Uri == nil {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessPolicySummaryEntity) GetUriOk() (*string, bool) {
	if o == nil || o.Uri == nil {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *AccessPolicySummaryEntity) HasUri() bool {
	if o != nil && o.Uri != nil {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *AccessPolicySummaryEntity) SetUri(v string) {
	o.Uri = &v
}

// GetPosition returns the Position field value if set, zero value otherwise.
func (o *AccessPolicySummaryEntity) GetPosition() PositionDTO {
	if o == nil || o.Position == nil {
		var ret PositionDTO
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessPolicySummaryEntity) GetPositionOk() (*PositionDTO, bool) {
	if o == nil || o.Position == nil {
		return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *AccessPolicySummaryEntity) HasPosition() bool {
	if o != nil && o.Position != nil {
		return true
	}

	return false
}

// SetPosition gets a reference to the given PositionDTO and assigns it to the Position field.
func (o *AccessPolicySummaryEntity) SetPosition(v PositionDTO) {
	o.Position = &v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *AccessPolicySummaryEntity) GetPermissions() PermissionsDTO {
	if o == nil || o.Permissions == nil {
		var ret PermissionsDTO
		return ret
	}
	return *o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessPolicySummaryEntity) GetPermissionsOk() (*PermissionsDTO, bool) {
	if o == nil || o.Permissions == nil {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *AccessPolicySummaryEntity) HasPermissions() bool {
	if o != nil && o.Permissions != nil {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given PermissionsDTO and assigns it to the Permissions field.
func (o *AccessPolicySummaryEntity) SetPermissions(v PermissionsDTO) {
	o.Permissions = &v
}

// GetBulletins returns the Bulletins field value if set, zero value otherwise.
func (o *AccessPolicySummaryEntity) GetBulletins() []BulletinEntity {
	if o == nil || o.Bulletins == nil {
		var ret []BulletinEntity
		return ret
	}
	return *o.Bulletins
}

// GetBulletinsOk returns a tuple with the Bulletins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessPolicySummaryEntity) GetBulletinsOk() (*[]BulletinEntity, bool) {
	if o == nil || o.Bulletins == nil {
		return nil, false
	}
	return o.Bulletins, true
}

// HasBulletins returns a boolean if a field has been set.
func (o *AccessPolicySummaryEntity) HasBulletins() bool {
	if o != nil && o.Bulletins != nil {
		return true
	}

	return false
}

// SetBulletins gets a reference to the given []BulletinEntity and assigns it to the Bulletins field.
func (o *AccessPolicySummaryEntity) SetBulletins(v []BulletinEntity) {
	o.Bulletins = &v
}

// GetDisconnectedNodeAcknowledged returns the DisconnectedNodeAcknowledged field value if set, zero value otherwise.
func (o *AccessPolicySummaryEntity) GetDisconnectedNodeAcknowledged() bool {
	if o == nil || o.DisconnectedNodeAcknowledged == nil {
		var ret bool
		return ret
	}
	return *o.DisconnectedNodeAcknowledged
}

// GetDisconnectedNodeAcknowledgedOk returns a tuple with the DisconnectedNodeAcknowledged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessPolicySummaryEntity) GetDisconnectedNodeAcknowledgedOk() (*bool, bool) {
	if o == nil || o.DisconnectedNodeAcknowledged == nil {
		return nil, false
	}
	return o.DisconnectedNodeAcknowledged, true
}

// HasDisconnectedNodeAcknowledged returns a boolean if a field has been set.
func (o *AccessPolicySummaryEntity) HasDisconnectedNodeAcknowledged() bool {
	if o != nil && o.DisconnectedNodeAcknowledged != nil {
		return true
	}

	return false
}

// SetDisconnectedNodeAcknowledged gets a reference to the given bool and assigns it to the DisconnectedNodeAcknowledged field.
func (o *AccessPolicySummaryEntity) SetDisconnectedNodeAcknowledged(v bool) {
	o.DisconnectedNodeAcknowledged = &v
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *AccessPolicySummaryEntity) GetComponent() AccessPolicySummaryDTO {
	if o == nil || o.Component == nil {
		var ret AccessPolicySummaryDTO
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessPolicySummaryEntity) GetComponentOk() (*AccessPolicySummaryDTO, bool) {
	if o == nil || o.Component == nil {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *AccessPolicySummaryEntity) HasComponent() bool {
	if o != nil && o.Component != nil {
		return true
	}

	return false
}

// SetComponent gets a reference to the given AccessPolicySummaryDTO and assigns it to the Component field.
func (o *AccessPolicySummaryEntity) SetComponent(v AccessPolicySummaryDTO) {
	o.Component = &v
}

func (o AccessPolicySummaryEntity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Revision != nil {
		toSerialize["revision"] = o.Revision
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Uri != nil {
		toSerialize["uri"] = o.Uri
	}
	if o.Position != nil {
		toSerialize["position"] = o.Position
	}
	if o.Permissions != nil {
		toSerialize["permissions"] = o.Permissions
	}
	if o.Bulletins != nil {
		toSerialize["bulletins"] = o.Bulletins
	}
	if o.DisconnectedNodeAcknowledged != nil {
		toSerialize["disconnectedNodeAcknowledged"] = o.DisconnectedNodeAcknowledged
	}
	if o.Component != nil {
		toSerialize["component"] = o.Component
	}
	return json.Marshal(toSerialize)
}

type NullableAccessPolicySummaryEntity struct {
	value *AccessPolicySummaryEntity
	isSet bool
}

func (v NullableAccessPolicySummaryEntity) Get() *AccessPolicySummaryEntity {
	return v.value
}

func (v *NullableAccessPolicySummaryEntity) Set(val *AccessPolicySummaryEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessPolicySummaryEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessPolicySummaryEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessPolicySummaryEntity(val *AccessPolicySummaryEntity) *NullableAccessPolicySummaryEntity {
	return &NullableAccessPolicySummaryEntity{value: val, isSet: true}
}

func (v NullableAccessPolicySummaryEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessPolicySummaryEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
