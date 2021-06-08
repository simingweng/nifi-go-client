# UserGroupDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The id of the component. | [optional] 
**VersionedComponentId** | Pointer to **string** | The ID of the corresponding component that is under version control | [optional] 
**ParentGroupId** | Pointer to **string** | The id of parent process group of this component if applicable. | [optional] 
**Position** | Pointer to [**PositionDTO**](PositionDTO.md) |  | [optional] 
**Identity** | Pointer to **string** | The identity of the tenant. | [optional] 
**Configurable** | Pointer to **bool** | Whether this tenant is configurable. | [optional] 
**Users** | Pointer to [**[]TenantEntity**](TenantEntity.md) | The users that belong to the user group. | [optional] 
**AccessPolicies** | Pointer to [**[]AccessPolicyEntity**](AccessPolicyEntity.md) | The access policies this user group belongs to. This field was incorrectly defined as an AccessPolicyEntity. For compatibility reasons the field will remain of this type, however only the fields that are present in the AccessPolicySummaryEntity will be populated here. | [optional] [readonly] 

## Methods

### NewUserGroupDTO

`func NewUserGroupDTO() *UserGroupDTO`

NewUserGroupDTO instantiates a new UserGroupDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserGroupDTOWithDefaults

`func NewUserGroupDTOWithDefaults() *UserGroupDTO`

NewUserGroupDTOWithDefaults instantiates a new UserGroupDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserGroupDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserGroupDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserGroupDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UserGroupDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetVersionedComponentId

`func (o *UserGroupDTO) GetVersionedComponentId() string`

GetVersionedComponentId returns the VersionedComponentId field if non-nil, zero value otherwise.

### GetVersionedComponentIdOk

`func (o *UserGroupDTO) GetVersionedComponentIdOk() (*string, bool)`

GetVersionedComponentIdOk returns a tuple with the VersionedComponentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionedComponentId

`func (o *UserGroupDTO) SetVersionedComponentId(v string)`

SetVersionedComponentId sets VersionedComponentId field to given value.

### HasVersionedComponentId

`func (o *UserGroupDTO) HasVersionedComponentId() bool`

HasVersionedComponentId returns a boolean if a field has been set.

### GetParentGroupId

`func (o *UserGroupDTO) GetParentGroupId() string`

GetParentGroupId returns the ParentGroupId field if non-nil, zero value otherwise.

### GetParentGroupIdOk

`func (o *UserGroupDTO) GetParentGroupIdOk() (*string, bool)`

GetParentGroupIdOk returns a tuple with the ParentGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentGroupId

`func (o *UserGroupDTO) SetParentGroupId(v string)`

SetParentGroupId sets ParentGroupId field to given value.

### HasParentGroupId

`func (o *UserGroupDTO) HasParentGroupId() bool`

HasParentGroupId returns a boolean if a field has been set.

### GetPosition

`func (o *UserGroupDTO) GetPosition() PositionDTO`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *UserGroupDTO) GetPositionOk() (*PositionDTO, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *UserGroupDTO) SetPosition(v PositionDTO)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *UserGroupDTO) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetIdentity

`func (o *UserGroupDTO) GetIdentity() string`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *UserGroupDTO) GetIdentityOk() (*string, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *UserGroupDTO) SetIdentity(v string)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *UserGroupDTO) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### GetConfigurable

`func (o *UserGroupDTO) GetConfigurable() bool`

GetConfigurable returns the Configurable field if non-nil, zero value otherwise.

### GetConfigurableOk

`func (o *UserGroupDTO) GetConfigurableOk() (*bool, bool)`

GetConfigurableOk returns a tuple with the Configurable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurable

`func (o *UserGroupDTO) SetConfigurable(v bool)`

SetConfigurable sets Configurable field to given value.

### HasConfigurable

`func (o *UserGroupDTO) HasConfigurable() bool`

HasConfigurable returns a boolean if a field has been set.

### GetUsers

`func (o *UserGroupDTO) GetUsers() []TenantEntity`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *UserGroupDTO) GetUsersOk() (*[]TenantEntity, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *UserGroupDTO) SetUsers(v []TenantEntity)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *UserGroupDTO) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetAccessPolicies

`func (o *UserGroupDTO) GetAccessPolicies() []AccessPolicyEntity`

GetAccessPolicies returns the AccessPolicies field if non-nil, zero value otherwise.

### GetAccessPoliciesOk

`func (o *UserGroupDTO) GetAccessPoliciesOk() (*[]AccessPolicyEntity, bool)`

GetAccessPoliciesOk returns a tuple with the AccessPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessPolicies

`func (o *UserGroupDTO) SetAccessPolicies(v []AccessPolicyEntity)`

SetAccessPolicies sets AccessPolicies field to given value.

### HasAccessPolicies

`func (o *UserGroupDTO) HasAccessPolicies() bool`

HasAccessPolicies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


