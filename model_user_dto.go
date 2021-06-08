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

// UserDTO struct for UserDTO
type UserDTO struct {
	// The id of the component.
	Id *string `json:"id,omitempty"`
	// The ID of the corresponding component that is under version control
	VersionedComponentId *string `json:"versionedComponentId,omitempty"`
	// The id of parent process group of this component if applicable.
	ParentGroupId *string      `json:"parentGroupId,omitempty"`
	Position      *PositionDTO `json:"position,omitempty"`
	// The identity of the tenant.
	Identity *string `json:"identity,omitempty"`
	// Whether this tenant is configurable.
	Configurable *bool `json:"configurable,omitempty"`
	// The groups to which the user belongs. This field is read only and it provided for convenience.
	UserGroups *[]TenantEntity `json:"userGroups,omitempty"`
	// The access policies this user belongs to.
	AccessPolicies *[]AccessPolicySummaryEntity `json:"accessPolicies,omitempty"`
}

// NewUserDTO instantiates a new UserDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserDTO() *UserDTO {
	this := UserDTO{}
	return &this
}

// NewUserDTOWithDefaults instantiates a new UserDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserDTOWithDefaults() *UserDTO {
	this := UserDTO{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UserDTO) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDTO) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UserDTO) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *UserDTO) SetId(v string) {
	o.Id = &v
}

// GetVersionedComponentId returns the VersionedComponentId field value if set, zero value otherwise.
func (o *UserDTO) GetVersionedComponentId() string {
	if o == nil || o.VersionedComponentId == nil {
		var ret string
		return ret
	}
	return *o.VersionedComponentId
}

// GetVersionedComponentIdOk returns a tuple with the VersionedComponentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDTO) GetVersionedComponentIdOk() (*string, bool) {
	if o == nil || o.VersionedComponentId == nil {
		return nil, false
	}
	return o.VersionedComponentId, true
}

// HasVersionedComponentId returns a boolean if a field has been set.
func (o *UserDTO) HasVersionedComponentId() bool {
	if o != nil && o.VersionedComponentId != nil {
		return true
	}

	return false
}

// SetVersionedComponentId gets a reference to the given string and assigns it to the VersionedComponentId field.
func (o *UserDTO) SetVersionedComponentId(v string) {
	o.VersionedComponentId = &v
}

// GetParentGroupId returns the ParentGroupId field value if set, zero value otherwise.
func (o *UserDTO) GetParentGroupId() string {
	if o == nil || o.ParentGroupId == nil {
		var ret string
		return ret
	}
	return *o.ParentGroupId
}

// GetParentGroupIdOk returns a tuple with the ParentGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDTO) GetParentGroupIdOk() (*string, bool) {
	if o == nil || o.ParentGroupId == nil {
		return nil, false
	}
	return o.ParentGroupId, true
}

// HasParentGroupId returns a boolean if a field has been set.
func (o *UserDTO) HasParentGroupId() bool {
	if o != nil && o.ParentGroupId != nil {
		return true
	}

	return false
}

// SetParentGroupId gets a reference to the given string and assigns it to the ParentGroupId field.
func (o *UserDTO) SetParentGroupId(v string) {
	o.ParentGroupId = &v
}

// GetPosition returns the Position field value if set, zero value otherwise.
func (o *UserDTO) GetPosition() PositionDTO {
	if o == nil || o.Position == nil {
		var ret PositionDTO
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDTO) GetPositionOk() (*PositionDTO, bool) {
	if o == nil || o.Position == nil {
		return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *UserDTO) HasPosition() bool {
	if o != nil && o.Position != nil {
		return true
	}

	return false
}

// SetPosition gets a reference to the given PositionDTO and assigns it to the Position field.
func (o *UserDTO) SetPosition(v PositionDTO) {
	o.Position = &v
}

// GetIdentity returns the Identity field value if set, zero value otherwise.
func (o *UserDTO) GetIdentity() string {
	if o == nil || o.Identity == nil {
		var ret string
		return ret
	}
	return *o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDTO) GetIdentityOk() (*string, bool) {
	if o == nil || o.Identity == nil {
		return nil, false
	}
	return o.Identity, true
}

// HasIdentity returns a boolean if a field has been set.
func (o *UserDTO) HasIdentity() bool {
	if o != nil && o.Identity != nil {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given string and assigns it to the Identity field.
func (o *UserDTO) SetIdentity(v string) {
	o.Identity = &v
}

// GetConfigurable returns the Configurable field value if set, zero value otherwise.
func (o *UserDTO) GetConfigurable() bool {
	if o == nil || o.Configurable == nil {
		var ret bool
		return ret
	}
	return *o.Configurable
}

// GetConfigurableOk returns a tuple with the Configurable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDTO) GetConfigurableOk() (*bool, bool) {
	if o == nil || o.Configurable == nil {
		return nil, false
	}
	return o.Configurable, true
}

// HasConfigurable returns a boolean if a field has been set.
func (o *UserDTO) HasConfigurable() bool {
	if o != nil && o.Configurable != nil {
		return true
	}

	return false
}

// SetConfigurable gets a reference to the given bool and assigns it to the Configurable field.
func (o *UserDTO) SetConfigurable(v bool) {
	o.Configurable = &v
}

// GetUserGroups returns the UserGroups field value if set, zero value otherwise.
func (o *UserDTO) GetUserGroups() []TenantEntity {
	if o == nil || o.UserGroups == nil {
		var ret []TenantEntity
		return ret
	}
	return *o.UserGroups
}

// GetUserGroupsOk returns a tuple with the UserGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDTO) GetUserGroupsOk() (*[]TenantEntity, bool) {
	if o == nil || o.UserGroups == nil {
		return nil, false
	}
	return o.UserGroups, true
}

// HasUserGroups returns a boolean if a field has been set.
func (o *UserDTO) HasUserGroups() bool {
	if o != nil && o.UserGroups != nil {
		return true
	}

	return false
}

// SetUserGroups gets a reference to the given []TenantEntity and assigns it to the UserGroups field.
func (o *UserDTO) SetUserGroups(v []TenantEntity) {
	o.UserGroups = &v
}

// GetAccessPolicies returns the AccessPolicies field value if set, zero value otherwise.
func (o *UserDTO) GetAccessPolicies() []AccessPolicySummaryEntity {
	if o == nil || o.AccessPolicies == nil {
		var ret []AccessPolicySummaryEntity
		return ret
	}
	return *o.AccessPolicies
}

// GetAccessPoliciesOk returns a tuple with the AccessPolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDTO) GetAccessPoliciesOk() (*[]AccessPolicySummaryEntity, bool) {
	if o == nil || o.AccessPolicies == nil {
		return nil, false
	}
	return o.AccessPolicies, true
}

// HasAccessPolicies returns a boolean if a field has been set.
func (o *UserDTO) HasAccessPolicies() bool {
	if o != nil && o.AccessPolicies != nil {
		return true
	}

	return false
}

// SetAccessPolicies gets a reference to the given []AccessPolicySummaryEntity and assigns it to the AccessPolicies field.
func (o *UserDTO) SetAccessPolicies(v []AccessPolicySummaryEntity) {
	o.AccessPolicies = &v
}

func (o UserDTO) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.VersionedComponentId != nil {
		toSerialize["versionedComponentId"] = o.VersionedComponentId
	}
	if o.ParentGroupId != nil {
		toSerialize["parentGroupId"] = o.ParentGroupId
	}
	if o.Position != nil {
		toSerialize["position"] = o.Position
	}
	if o.Identity != nil {
		toSerialize["identity"] = o.Identity
	}
	if o.Configurable != nil {
		toSerialize["configurable"] = o.Configurable
	}
	if o.UserGroups != nil {
		toSerialize["userGroups"] = o.UserGroups
	}
	if o.AccessPolicies != nil {
		toSerialize["accessPolicies"] = o.AccessPolicies
	}
	return json.Marshal(toSerialize)
}

type NullableUserDTO struct {
	value *UserDTO
	isSet bool
}

func (v NullableUserDTO) Get() *UserDTO {
	return v.value
}

func (v *NullableUserDTO) Set(val *UserDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableUserDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableUserDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserDTO(val *UserDTO) *NullableUserDTO {
	return &NullableUserDTO{value: val, isSet: true}
}

func (v NullableUserDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
