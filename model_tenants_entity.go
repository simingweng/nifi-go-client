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

// TenantsEntity struct for TenantsEntity
type TenantsEntity struct {
	Users      *[]TenantEntity `json:"users,omitempty"`
	UserGroups *[]TenantEntity `json:"userGroups,omitempty"`
}

// NewTenantsEntity instantiates a new TenantsEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTenantsEntity() *TenantsEntity {
	this := TenantsEntity{}
	return &this
}

// NewTenantsEntityWithDefaults instantiates a new TenantsEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTenantsEntityWithDefaults() *TenantsEntity {
	this := TenantsEntity{}
	return &this
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *TenantsEntity) GetUsers() []TenantEntity {
	if o == nil || o.Users == nil {
		var ret []TenantEntity
		return ret
	}
	return *o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantsEntity) GetUsersOk() (*[]TenantEntity, bool) {
	if o == nil || o.Users == nil {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *TenantsEntity) HasUsers() bool {
	if o != nil && o.Users != nil {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []TenantEntity and assigns it to the Users field.
func (o *TenantsEntity) SetUsers(v []TenantEntity) {
	o.Users = &v
}

// GetUserGroups returns the UserGroups field value if set, zero value otherwise.
func (o *TenantsEntity) GetUserGroups() []TenantEntity {
	if o == nil || o.UserGroups == nil {
		var ret []TenantEntity
		return ret
	}
	return *o.UserGroups
}

// GetUserGroupsOk returns a tuple with the UserGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantsEntity) GetUserGroupsOk() (*[]TenantEntity, bool) {
	if o == nil || o.UserGroups == nil {
		return nil, false
	}
	return o.UserGroups, true
}

// HasUserGroups returns a boolean if a field has been set.
func (o *TenantsEntity) HasUserGroups() bool {
	if o != nil && o.UserGroups != nil {
		return true
	}

	return false
}

// SetUserGroups gets a reference to the given []TenantEntity and assigns it to the UserGroups field.
func (o *TenantsEntity) SetUserGroups(v []TenantEntity) {
	o.UserGroups = &v
}

func (o TenantsEntity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Users != nil {
		toSerialize["users"] = o.Users
	}
	if o.UserGroups != nil {
		toSerialize["userGroups"] = o.UserGroups
	}
	return json.Marshal(toSerialize)
}

type NullableTenantsEntity struct {
	value *TenantsEntity
	isSet bool
}

func (v NullableTenantsEntity) Get() *TenantsEntity {
	return v.value
}

func (v *NullableTenantsEntity) Set(val *TenantsEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableTenantsEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableTenantsEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTenantsEntity(val *TenantsEntity) *NullableTenantsEntity {
	return &NullableTenantsEntity{value: val, isSet: true}
}

func (v NullableTenantsEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTenantsEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
