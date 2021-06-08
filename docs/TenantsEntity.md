# TenantsEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Users** | Pointer to [**[]TenantEntity**](TenantEntity.md) |  | [optional] 
**UserGroups** | Pointer to [**[]TenantEntity**](TenantEntity.md) |  | [optional] 

## Methods

### NewTenantsEntity

`func NewTenantsEntity() *TenantsEntity`

NewTenantsEntity instantiates a new TenantsEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantsEntityWithDefaults

`func NewTenantsEntityWithDefaults() *TenantsEntity`

NewTenantsEntityWithDefaults instantiates a new TenantsEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsers

`func (o *TenantsEntity) GetUsers() []TenantEntity`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *TenantsEntity) GetUsersOk() (*[]TenantEntity, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *TenantsEntity) SetUsers(v []TenantEntity)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *TenantsEntity) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetUserGroups

`func (o *TenantsEntity) GetUserGroups() []TenantEntity`

GetUserGroups returns the UserGroups field if non-nil, zero value otherwise.

### GetUserGroupsOk

`func (o *TenantsEntity) GetUserGroupsOk() (*[]TenantEntity, bool)`

GetUserGroupsOk returns a tuple with the UserGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroups

`func (o *TenantsEntity) SetUserGroups(v []TenantEntity)`

SetUserGroups sets UserGroups field to given value.

### HasUserGroups

`func (o *TenantsEntity) HasUserGroups() bool`

HasUserGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


