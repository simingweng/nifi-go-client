# \ParameterContextsApi

All URIs are relative to *http://localhost/nifi-api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateParameterContext**](ParameterContextsApi.md#CreateParameterContext) | **Post** /parameter-contexts | Create a Parameter Context
[**DeleteParameterContext**](ParameterContextsApi.md#DeleteParameterContext) | **Delete** /parameter-contexts/{id} | Deletes the Parameter Context with the given ID
[**DeleteUpdateRequest**](ParameterContextsApi.md#DeleteUpdateRequest) | **Delete** /parameter-contexts/{contextId}/update-requests/{requestId} | Deletes the Update Request with the given ID
[**DeleteValidationRequest**](ParameterContextsApi.md#DeleteValidationRequest) | **Delete** /parameter-contexts/{contextId}/validation-requests/{id} | Deletes the Validation Request with the given ID
[**GetParameterContext**](ParameterContextsApi.md#GetParameterContext) | **Get** /parameter-contexts/{id} | Returns the Parameter Context with the given ID
[**GetParameterContextUpdate**](ParameterContextsApi.md#GetParameterContextUpdate) | **Get** /parameter-contexts/{contextId}/update-requests/{requestId} | Returns the Update Request with the given ID
[**GetValidationRequest**](ParameterContextsApi.md#GetValidationRequest) | **Get** /parameter-contexts/{contextId}/validation-requests/{id} | Returns the Validation Request with the given ID
[**SubmitParameterContextUpdate**](ParameterContextsApi.md#SubmitParameterContextUpdate) | **Post** /parameter-contexts/{contextId}/update-requests | Initiate the Update Request of a Parameter Context
[**SubmitValidationRequest**](ParameterContextsApi.md#SubmitValidationRequest) | **Post** /parameter-contexts/{contextId}/validation-requests | Initiate a Validation Request to determine how the validity of components will change if a Parameter Context were to be updated
[**UpdateParameterContext**](ParameterContextsApi.md#UpdateParameterContext) | **Put** /parameter-contexts/{id} | Modifies a Parameter Context



## CreateParameterContext

> ParameterContextEntity CreateParameterContext(ctx, body)

Create a Parameter Context

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**ParameterContextEntity**](ParameterContextEntity.md)| The Parameter Context. | 

### Return type

[**ParameterContextEntity**](ParameterContextEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteParameterContext

> ParameterContextEntity DeleteParameterContext(ctx, id, optional)

Deletes the Parameter Context with the given ID

Deletes the Parameter Context with the given ID.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The Parameter Context ID. | 
 **optional** | ***DeleteParameterContextOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteParameterContextOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **optional.String**| The version is used to verify the client is working with the latest version of the flow. | 
 **clientId** | **optional.String**| If the client id is not specified, a new one will be generated. This value (whether specified or generated) is included in the response. | 
 **disconnectedNodeAcknowledged** | **optional.Bool**| Acknowledges that this node is disconnected to allow for mutable requests to proceed. | [default to false]

### Return type

[**ParameterContextEntity**](ParameterContextEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUpdateRequest

> ParameterContextUpdateRequestEntity DeleteUpdateRequest(ctx, contextId, requestId, optional)

Deletes the Update Request with the given ID

Deletes the Update Request with the given ID. After a request is created via a POST to /nifi-api/parameter-contexts/update-requests, it is expected that the client will properly clean up the request by DELETE'ing it, once the Update process has completed. If the request is deleted before the request completes, then the Update request will finish the step that it is currently performing and then will cancel any subsequent steps.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contextId** | **string**| The ID of the ParameterContext | 
**requestId** | **string**| The ID of the Update Request | 
 **optional** | ***DeleteUpdateRequestOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteUpdateRequestOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **disconnectedNodeAcknowledged** | **optional.Bool**| Acknowledges that this node is disconnected to allow for mutable requests to proceed. | [default to false]

### Return type

[**ParameterContextUpdateRequestEntity**](ParameterContextUpdateRequestEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteValidationRequest

> ParameterContextValidationRequestEntity DeleteValidationRequest(ctx, contextId, id, optional)

Deletes the Validation Request with the given ID

Deletes the Validation Request with the given ID. After a request is created via a POST to /nifi-api/validation-contexts, it is expected that the client will properly clean up the request by DELETE'ing it, once the validation process has completed. If the request is deleted before the request completes, then the Validation request will finish the step that it is currently performing and then will cancel any subsequent steps.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contextId** | **string**| The ID of the Parameter Context | 
**id** | **string**| The ID of the Update Request | 
 **optional** | ***DeleteValidationRequestOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteValidationRequestOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **disconnectedNodeAcknowledged** | **optional.Bool**| Acknowledges that this node is disconnected to allow for mutable requests to proceed. | [default to false]

### Return type

[**ParameterContextValidationRequestEntity**](ParameterContextValidationRequestEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetParameterContext

> ParameterContextEntity GetParameterContext(ctx, id)

Returns the Parameter Context with the given ID

Returns the Parameter Context with the given ID.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of the Parameter Context | 

### Return type

[**ParameterContextEntity**](ParameterContextEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetParameterContextUpdate

> ParameterContextUpdateRequestEntity GetParameterContextUpdate(ctx, contextId, requestId)

Returns the Update Request with the given ID

Returns the Update Request with the given ID. Once an Update Request has been created by performing a POST to /nifi-api/parameter-contexts, that request can subsequently be retrieved via this endpoint, and the request that is fetched will contain the updated state, such as percent complete, the current state of the request, and any failures. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contextId** | **string**| The ID of the Parameter Context | 
**requestId** | **string**| The ID of the Update Request | 

### Return type

[**ParameterContextUpdateRequestEntity**](ParameterContextUpdateRequestEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetValidationRequest

> ParameterContextValidationRequestEntity GetValidationRequest(ctx, contextId, id)

Returns the Validation Request with the given ID

Returns the Validation Request with the given ID. Once a Validation Request has been created by performing a POST to /nifi-api/validation-contexts, that request can subsequently be retrieved via this endpoint, and the request that is fetched will contain the updated state, such as percent complete, the current state of the request, and any failures. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contextId** | **string**| The ID of the Parameter Context | 
**id** | **string**| The ID of the Validation Request | 

### Return type

[**ParameterContextValidationRequestEntity**](ParameterContextValidationRequestEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitParameterContextUpdate

> ParameterContextUpdateRequestEntity SubmitParameterContextUpdate(ctx, contextId, body)

Initiate the Update Request of a Parameter Context

This will initiate the process of updating a Parameter Context. Changing the value of a Parameter may require that one or more components be stopped and restarted, so this acttion may take significantly more time than many other REST API actions. As a result, this endpoint will immediately return a ParameterContextUpdateRequestEntity, and the process of updating the necessary components will occur asynchronously in the background. The client may then periodically poll the status of the request by issuing a GET request to /parameter-contexts/update-requests/{requestId}. Once the request is completed, the client is expected to issue a DELETE request to /parameter-contexts/update-requests/{requestId}.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contextId** | **string**|  | 
**body** | [**ParameterContextEntity**](ParameterContextEntity.md)| The updated version of the parameter context. | 

### Return type

[**ParameterContextUpdateRequestEntity**](ParameterContextUpdateRequestEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitValidationRequest

> ParameterContextValidationRequestEntity SubmitValidationRequest(ctx, contextId, body)

Initiate a Validation Request to determine how the validity of components will change if a Parameter Context were to be updated

This will initiate the process of validating all components whose Process Group is bound to the specified Parameter Context. Performing validation against an arbitrary number of components may be expect and take significantly more time than many other REST API actions. As a result, this endpoint will immediately return a ParameterContextValidationRequestEntity, and the process of validating the necessary components will occur asynchronously in the background. The client may then periodically poll the status of the request by issuing a GET request to /parameter-contexts/validation-requests/{requestId}. Once the request is completed, the client is expected to issue a DELETE request to /parameter-contexts/validation-requests/{requestId}.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contextId** | **string**|  | 
**body** | [**ParameterContextValidationRequestEntity**](ParameterContextValidationRequestEntity.md)| The validation request | 

### Return type

[**ParameterContextValidationRequestEntity**](ParameterContextValidationRequestEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateParameterContext

> ParameterContextEntity UpdateParameterContext(ctx, id, body)

Modifies a Parameter Context

This endpoint will update a Parameter Context to match the provided entity. However, this request will fail if any component is running and is referencing a Parameter in the Parameter Context. Generally, this endpoint is not called directly. Instead, an update request should be submitted by making a POST to the /parameter-contexts/update-requests endpoint. That endpoint will, in turn, call this endpoint.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
**body** | [**ParameterContextEntity**](ParameterContextEntity.md)| The updated Parameter Context | 

### Return type

[**ParameterContextEntity**](ParameterContextEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

