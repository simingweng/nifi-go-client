# AccessPolicyDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The id of the component. | [optional] 
**VersionedComponentId** | Pointer to **string** | The ID of the corresponding component that is under version control | [optional] 
**ParentGroupId** | Pointer to **string** | The id of parent process group of this component if applicable. | [optional] 
**Position** | Pointer to [**PositionDTO**](PositionDTO.md) |  | [optional] 
**Resource** | Pointer to **string** | The resource for this access policy. | [optional] 
**Action** | Pointer to **string** | The action associated with this access policy. | [optional] 
**ComponentReference** | Pointer to [**ComponentReferenceEntity**](ComponentReferenceEntity.md) |  | [optional] 
**Configurable** | Pointer to **bool** | Whether this policy is configurable. | [optional] 
**Users** | Pointer to [**[]TenantEntity**](TenantEntity.md) | The set of user IDs associated with this access policy. | [optional] 
**UserGroups** | Pointer to [**[]TenantEntity**](TenantEntity.md) | The set of user group IDs associated with this access policy. | [optional] 

## Methods

### NewAccessPolicyDTO

`func NewAccessPolicyDTO() *AccessPolicyDTO`

NewAccessPolicyDTO instantiates a new AccessPolicyDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessPolicyDTOWithDefaults

`func NewAccessPolicyDTOWithDefaults() *AccessPolicyDTO`

NewAccessPolicyDTOWithDefaults instantiates a new AccessPolicyDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AccessPolicyDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccessPolicyDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccessPolicyDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AccessPolicyDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetVersionedComponentId

`func (o *AccessPolicyDTO) GetVersionedComponentId() string`

GetVersionedComponentId returns the VersionedComponentId field if non-nil, zero value otherwise.

### GetVersionedComponentIdOk

`func (o *AccessPolicyDTO) GetVersionedComponentIdOk() (*string, bool)`

GetVersionedComponentIdOk returns a tuple with the VersionedComponentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionedComponentId

`func (o *AccessPolicyDTO) SetVersionedComponentId(v string)`

SetVersionedComponentId sets VersionedComponentId field to given value.

### HasVersionedComponentId

`func (o *AccessPolicyDTO) HasVersionedComponentId() bool`

HasVersionedComponentId returns a boolean if a field has been set.

### GetParentGroupId

`func (o *AccessPolicyDTO) GetParentGroupId() string`

GetParentGroupId returns the ParentGroupId field if non-nil, zero value otherwise.

### GetParentGroupIdOk

`func (o *AccessPolicyDTO) GetParentGroupIdOk() (*string, bool)`

GetParentGroupIdOk returns a tuple with the ParentGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentGroupId

`func (o *AccessPolicyDTO) SetParentGroupId(v string)`

SetParentGroupId sets ParentGroupId field to given value.

### HasParentGroupId

`func (o *AccessPolicyDTO) HasParentGroupId() bool`

HasParentGroupId returns a boolean if a field has been set.

### GetPosition

`func (o *AccessPolicyDTO) GetPosition() PositionDTO`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *AccessPolicyDTO) GetPositionOk() (*PositionDTO, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *AccessPolicyDTO) SetPosition(v PositionDTO)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *AccessPolicyDTO) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetResource

`func (o *AccessPolicyDTO) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *AccessPolicyDTO) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *AccessPolicyDTO) SetResource(v string)`

SetResource sets Resource field to given value.

### HasResource

`func (o *AccessPolicyDTO) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetAction

`func (o *AccessPolicyDTO) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *AccessPolicyDTO) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *AccessPolicyDTO) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *AccessPolicyDTO) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetComponentReference

`func (o *AccessPolicyDTO) GetComponentReference() ComponentReferenceEntity`

GetComponentReference returns the ComponentReference field if non-nil, zero value otherwise.

### GetComponentReferenceOk

`func (o *AccessPolicyDTO) GetComponentReferenceOk() (*ComponentReferenceEntity, bool)`

GetComponentReferenceOk returns a tuple with the ComponentReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentReference

`func (o *AccessPolicyDTO) SetComponentReference(v ComponentReferenceEntity)`

SetComponentReference sets ComponentReference field to given value.

### HasComponentReference

`func (o *AccessPolicyDTO) HasComponentReference() bool`

HasComponentReference returns a boolean if a field has been set.

### GetConfigurable

`func (o *AccessPolicyDTO) GetConfigurable() bool`

GetConfigurable returns the Configurable field if non-nil, zero value otherwise.

### GetConfigurableOk

`func (o *AccessPolicyDTO) GetConfigurableOk() (*bool, bool)`

GetConfigurableOk returns a tuple with the Configurable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurable

`func (o *AccessPolicyDTO) SetConfigurable(v bool)`

SetConfigurable sets Configurable field to given value.

### HasConfigurable

`func (o *AccessPolicyDTO) HasConfigurable() bool`

HasConfigurable returns a boolean if a field has been set.

### GetUsers

`func (o *AccessPolicyDTO) GetUsers() []TenantEntity`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *AccessPolicyDTO) GetUsersOk() (*[]TenantEntity, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *AccessPolicyDTO) SetUsers(v []TenantEntity)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *AccessPolicyDTO) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetUserGroups

`func (o *AccessPolicyDTO) GetUserGroups() []TenantEntity`

GetUserGroups returns the UserGroups field if non-nil, zero value otherwise.

### GetUserGroupsOk

`func (o *AccessPolicyDTO) GetUserGroupsOk() (*[]TenantEntity, bool)`

GetUserGroupsOk returns a tuple with the UserGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroups

`func (o *AccessPolicyDTO) SetUserGroups(v []TenantEntity)`

SetUserGroups sets UserGroups field to given value.

### HasUserGroups

`func (o *AccessPolicyDTO) HasUserGroups() bool`

HasUserGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


