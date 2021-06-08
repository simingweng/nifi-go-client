# UserGroupsEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserGroups** | Pointer to [**[]UserGroupEntity**](UserGroupEntity.md) |  | [optional] 

## Methods

### NewUserGroupsEntity

`func NewUserGroupsEntity() *UserGroupsEntity`

NewUserGroupsEntity instantiates a new UserGroupsEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserGroupsEntityWithDefaults

`func NewUserGroupsEntityWithDefaults() *UserGroupsEntity`

NewUserGroupsEntityWithDefaults instantiates a new UserGroupsEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserGroups

`func (o *UserGroupsEntity) GetUserGroups() []UserGroupEntity`

GetUserGroups returns the UserGroups field if non-nil, zero value otherwise.

### GetUserGroupsOk

`func (o *UserGroupsEntity) GetUserGroupsOk() (*[]UserGroupEntity, bool)`

GetUserGroupsOk returns a tuple with the UserGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroups

`func (o *UserGroupsEntity) SetUserGroups(v []UserGroupEntity)`

SetUserGroups sets UserGroups field to given value.

### HasUserGroups

`func (o *UserGroupsEntity) HasUserGroups() bool`

HasUserGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


