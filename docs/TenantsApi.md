# \TenantsApi

All URIs are relative to *http://localhost/nifi-api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUser**](TenantsApi.md#CreateUser) | **Post** /tenants/users | Creates a user
[**CreateUserGroup**](TenantsApi.md#CreateUserGroup) | **Post** /tenants/user-groups | Creates a user group
[**GetUser**](TenantsApi.md#GetUser) | **Get** /tenants/users/{id} | Gets a user
[**GetUserGroup**](TenantsApi.md#GetUserGroup) | **Get** /tenants/user-groups/{id} | Gets a user group
[**GetUserGroups**](TenantsApi.md#GetUserGroups) | **Get** /tenants/user-groups | Gets all user groups
[**GetUsers**](TenantsApi.md#GetUsers) | **Get** /tenants/users | Gets all users
[**RemoveUser**](TenantsApi.md#RemoveUser) | **Delete** /tenants/users/{id} | Deletes a user
[**RemoveUserGroup**](TenantsApi.md#RemoveUserGroup) | **Delete** /tenants/user-groups/{id} | Deletes a user group
[**SearchTenants**](TenantsApi.md#SearchTenants) | **Get** /tenants/search-results | Searches for a tenant with the specified identity
[**UpdateUser**](TenantsApi.md#UpdateUser) | **Put** /tenants/users/{id} | Updates a user
[**UpdateUserGroup**](TenantsApi.md#UpdateUserGroup) | **Put** /tenants/user-groups/{id} | Updates a user group



## CreateUser

> UserEntity CreateUser(ctx, body)

Creates a user

Note: This endpoint is subject to change as NiFi and it's REST API evolve.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**UserEntity**](UserEntity.md)| The user configuration details. | 

### Return type

[**UserEntity**](UserEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUserGroup

> UserGroupEntity CreateUserGroup(ctx, body)

Creates a user group

Note: This endpoint is subject to change as NiFi and it's REST API evolve.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**UserGroupEntity**](UserGroupEntity.md)| The user group configuration details. | 

### Return type

[**UserGroupEntity**](UserGroupEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUser

> UserEntity GetUser(ctx, id)

Gets a user

Note: This endpoint is subject to change as NiFi and it's REST API evolve.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The user id. | 

### Return type

[**UserEntity**](UserEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserGroup

> UserGroupEntity GetUserGroup(ctx, id)

Gets a user group

Note: This endpoint is subject to change as NiFi and it's REST API evolve.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The user group id. | 

### Return type

[**UserGroupEntity**](UserGroupEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserGroups

> UserGroupsEntity GetUserGroups(ctx, )

Gets all user groups

Note: This endpoint is subject to change as NiFi and it's REST API evolve.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**UserGroupsEntity**](UserGroupsEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsers

> UsersEntity GetUsers(ctx, )

Gets all users

Note: This endpoint is subject to change as NiFi and it's REST API evolve.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**UsersEntity**](UsersEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveUser

> UserEntity RemoveUser(ctx, id, optional)

Deletes a user

Note: This endpoint is subject to change as NiFi and it's REST API evolve.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The user id. | 
 **optional** | ***RemoveUserOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RemoveUserOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **optional.String**| The revision is used to verify the client is working with the latest version of the flow. | 
 **clientId** | **optional.String**| If the client id is not specified, new one will be generated. This value (whether specified or generated) is included in the response. | 
 **disconnectedNodeAcknowledged** | **optional.Bool**| Acknowledges that this node is disconnected to allow for mutable requests to proceed. | [default to false]

### Return type

[**UserEntity**](UserEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveUserGroup

> UserGroupEntity RemoveUserGroup(ctx, id, optional)

Deletes a user group

Note: This endpoint is subject to change as NiFi and it's REST API evolve.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The user group id. | 
 **optional** | ***RemoveUserGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RemoveUserGroupOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **optional.String**| The revision is used to verify the client is working with the latest version of the flow. | 
 **clientId** | **optional.String**| If the client id is not specified, new one will be generated. This value (whether specified or generated) is included in the response. | 
 **disconnectedNodeAcknowledged** | **optional.Bool**| Acknowledges that this node is disconnected to allow for mutable requests to proceed. | [default to false]

### Return type

[**UserGroupEntity**](UserGroupEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchTenants

> TenantsEntity SearchTenants(ctx, q)

Searches for a tenant with the specified identity

Note: This endpoint is subject to change as NiFi and it's REST API evolve.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**q** | **string**| Identity to search for. | 

### Return type

[**TenantsEntity**](TenantsEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUser

> UserEntity UpdateUser(ctx, id, body)

Updates a user

Note: This endpoint is subject to change as NiFi and it's REST API evolve.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The user id. | 
**body** | [**UserEntity**](UserEntity.md)| The user configuration details. | 

### Return type

[**UserEntity**](UserEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserGroup

> UserGroupEntity UpdateUserGroup(ctx, id, body)

Updates a user group

Note: This endpoint is subject to change as NiFi and it's REST API evolve.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The user group id. | 
**body** | [**UserGroupEntity**](UserGroupEntity.md)| The user group configuration details. | 

### Return type

[**UserGroupEntity**](UserGroupEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

