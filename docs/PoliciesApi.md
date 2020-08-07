# \PoliciesApi

All URIs are relative to *http://localhost/nifi-api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccessPolicy**](PoliciesApi.md#CreateAccessPolicy) | **Post** /policies | Creates an access policy
[**GetAccessPolicy**](PoliciesApi.md#GetAccessPolicy) | **Get** /policies/{id} | Gets an access policy
[**GetAccessPolicyForResource**](PoliciesApi.md#GetAccessPolicyForResource) | **Get** /policies/{action}/{resource} | Gets an access policy for the specified action and resource
[**RemoveAccessPolicy**](PoliciesApi.md#RemoveAccessPolicy) | **Delete** /policies/{id} | Deletes an access policy
[**UpdateAccessPolicy**](PoliciesApi.md#UpdateAccessPolicy) | **Put** /policies/{id} | Updates a access policy



## CreateAccessPolicy

> AccessPolicyEntity CreateAccessPolicy(ctx, body)

Creates an access policy

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**AccessPolicyEntity**](AccessPolicyEntity.md)| The access policy configuration details. | 

### Return type

[**AccessPolicyEntity**](AccessPolicyEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccessPolicy

> AccessPolicyEntity GetAccessPolicy(ctx, id)

Gets an access policy

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The access policy id. | 

### Return type

[**AccessPolicyEntity**](AccessPolicyEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccessPolicyForResource

> AccessPolicyEntity GetAccessPolicyForResource(ctx, action, resource)

Gets an access policy for the specified action and resource

Will return the effective policy if no component specific policy exists for the specified action and resource. Must have Read permissions to the policy with the desired action and resource. Permissions for the policy that is returned will be indicated in the response. This means the client could be authorized to get the policy for a given component but the effective policy may be inherited from an ancestor Process Group. If the client does not have permissions to that policy, the response will not include the policy and the permissions in the response will be marked accordingly. If the client does not have permissions to the policy of the desired action and resource a 403 response will be returned.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**action** | **string**| The request action. | 
**resource** | **string**| The resource of the policy. | 

### Return type

[**AccessPolicyEntity**](AccessPolicyEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveAccessPolicy

> AccessPolicyEntity RemoveAccessPolicy(ctx, id, optional)

Deletes an access policy

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The access policy id. | 
 **optional** | ***RemoveAccessPolicyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RemoveAccessPolicyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **optional.String**| The revision is used to verify the client is working with the latest version of the flow. | 
 **clientId** | **optional.String**| If the client id is not specified, new one will be generated. This value (whether specified or generated) is included in the response. | 
 **disconnectedNodeAcknowledged** | **optional.Bool**| Acknowledges that this node is disconnected to allow for mutable requests to proceed. | [default to false]

### Return type

[**AccessPolicyEntity**](AccessPolicyEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAccessPolicy

> AccessPolicyEntity UpdateAccessPolicy(ctx, id, body)

Updates a access policy

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The access policy id. | 
**body** | [**AccessPolicyEntity**](AccessPolicyEntity.md)| The access policy configuration details. | 

### Return type

[**AccessPolicyEntity**](AccessPolicyEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

