# UserDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The id of the component. | [optional] 
**VersionedComponentId** | Pointer to **string** | The ID of the corresponding component that is under version control | [optional] 
**ParentGroupId** | Pointer to **string** | The id of parent process group of this component if applicable. | [optional] 
**Position** | Pointer to [**PositionDTO**](PositionDTO.md) |  | [optional] 
**Identity** | Pointer to **string** | The identity of the tenant. | [optional] 
**Configurable** | Pointer to **bool** | Whether this tenant is configurable. | [optional] 
**UserGroups** | Pointer to [**[]TenantEntity**](TenantEntity.md) | The groups to which the user belongs. This field is read only and it provided for convenience. | [optional] [readonly] 
**AccessPolicies** | Pointer to [**[]AccessPolicySummaryEntity**](AccessPolicySummaryEntity.md) | The access policies this user belongs to. | [optional] [readonly] 

## Methods

### NewUserDTO

`func NewUserDTO() *UserDTO`

NewUserDTO instantiates a new UserDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserDTOWithDefaults

`func NewUserDTOWithDefaults() *UserDTO`

NewUserDTOWithDefaults instantiates a new UserDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UserDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetVersionedComponentId

`func (o *UserDTO) GetVersionedComponentId() string`

GetVersionedComponentId returns the VersionedComponentId field if non-nil, zero value otherwise.

### GetVersionedComponentIdOk

`func (o *UserDTO) GetVersionedComponentIdOk() (*string, bool)`

GetVersionedComponentIdOk returns a tuple with the VersionedComponentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionedComponentId

`func (o *UserDTO) SetVersionedComponentId(v string)`

SetVersionedComponentId sets VersionedComponentId field to given value.

### HasVersionedComponentId

`func (o *UserDTO) HasVersionedComponentId() bool`

HasVersionedComponentId returns a boolean if a field has been set.

### GetParentGroupId

`func (o *UserDTO) GetParentGroupId() string`

GetParentGroupId returns the ParentGroupId field if non-nil, zero value otherwise.

### GetParentGroupIdOk

`func (o *UserDTO) GetParentGroupIdOk() (*string, bool)`

GetParentGroupIdOk returns a tuple with the ParentGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentGroupId

`func (o *UserDTO) SetParentGroupId(v string)`

SetParentGroupId sets ParentGroupId field to given value.

### HasParentGroupId

`func (o *UserDTO) HasParentGroupId() bool`

HasParentGroupId returns a boolean if a field has been set.

### GetPosition

`func (o *UserDTO) GetPosition() PositionDTO`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *UserDTO) GetPositionOk() (*PositionDTO, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *UserDTO) SetPosition(v PositionDTO)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *UserDTO) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetIdentity

`func (o *UserDTO) GetIdentity() string`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *UserDTO) GetIdentityOk() (*string, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *UserDTO) SetIdentity(v string)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *UserDTO) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### GetConfigurable

`func (o *UserDTO) GetConfigurable() bool`

GetConfigurable returns the Configurable field if non-nil, zero value otherwise.

### GetConfigurableOk

`func (o *UserDTO) GetConfigurableOk() (*bool, bool)`

GetConfigurableOk returns a tuple with the Configurable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurable

`func (o *UserDTO) SetConfigurable(v bool)`

SetConfigurable sets Configurable field to given value.

### HasConfigurable

`func (o *UserDTO) HasConfigurable() bool`

HasConfigurable returns a boolean if a field has been set.

### GetUserGroups

`func (o *UserDTO) GetUserGroups() []TenantEntity`

GetUserGroups returns the UserGroups field if non-nil, zero value otherwise.

### GetUserGroupsOk

`func (o *UserDTO) GetUserGroupsOk() (*[]TenantEntity, bool)`

GetUserGroupsOk returns a tuple with the UserGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroups

`func (o *UserDTO) SetUserGroups(v []TenantEntity)`

SetUserGroups sets UserGroups field to given value.

### HasUserGroups

`func (o *UserDTO) HasUserGroups() bool`

HasUserGroups returns a boolean if a field has been set.

### GetAccessPolicies

`func (o *UserDTO) GetAccessPolicies() []AccessPolicySummaryEntity`

GetAccessPolicies returns the AccessPolicies field if non-nil, zero value otherwise.

### GetAccessPoliciesOk

`func (o *UserDTO) GetAccessPoliciesOk() (*[]AccessPolicySummaryEntity, bool)`

GetAccessPoliciesOk returns a tuple with the AccessPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessPolicies

`func (o *UserDTO) SetAccessPolicies(v []AccessPolicySummaryEntity)`

SetAccessPolicies sets AccessPolicies field to given value.

### HasAccessPolicies

`func (o *UserDTO) HasAccessPolicies() bool`

HasAccessPolicies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


