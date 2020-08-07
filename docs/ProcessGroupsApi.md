# \ProcessGroupsApi

All URIs are relative to *http://localhost/nifi-api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CopySnippet**](ProcessGroupsApi.md#CopySnippet) | **Post** /process-groups/{id}/snippet-instance | Copies a snippet and discards it.
[**CreateConnection**](ProcessGroupsApi.md#CreateConnection) | **Post** /process-groups/{id}/connections | Creates a connection
[**CreateControllerService**](ProcessGroupsApi.md#CreateControllerService) | **Post** /process-groups/{id}/controller-services | Creates a new controller service
[**CreateFunnel**](ProcessGroupsApi.md#CreateFunnel) | **Post** /process-groups/{id}/funnels | Creates a funnel
[**CreateInputPort**](ProcessGroupsApi.md#CreateInputPort) | **Post** /process-groups/{id}/input-ports | Creates an input port
[**CreateLabel**](ProcessGroupsApi.md#CreateLabel) | **Post** /process-groups/{id}/labels | Creates a label
[**CreateOutputPort**](ProcessGroupsApi.md#CreateOutputPort) | **Post** /process-groups/{id}/output-ports | Creates an output port
[**CreateProcessGroup**](ProcessGroupsApi.md#CreateProcessGroup) | **Post** /process-groups/{id}/process-groups | Creates a process group
[**CreateProcessor**](ProcessGroupsApi.md#CreateProcessor) | **Post** /process-groups/{id}/processors | Creates a new processor
[**CreateRemoteProcessGroup**](ProcessGroupsApi.md#CreateRemoteProcessGroup) | **Post** /process-groups/{id}/remote-process-groups | Creates a new process group
[**CreateTemplate**](ProcessGroupsApi.md#CreateTemplate) | **Post** /process-groups/{id}/templates | Creates a template and discards the specified snippet.
[**DeleteVariableRegistryUpdateRequest**](ProcessGroupsApi.md#DeleteVariableRegistryUpdateRequest) | **Delete** /process-groups/{groupId}/variable-registry/update-requests/{updateId} | Deletes an update request for a process group&#39;s variable registry. If the request is not yet complete, it will automatically be cancelled.
[**ExportProcessGroup**](ProcessGroupsApi.md#ExportProcessGroup) | **Get** /process-groups/{id}/download | Gets a process group for download
[**GetConnections**](ProcessGroupsApi.md#GetConnections) | **Get** /process-groups/{id}/connections | Gets all connections
[**GetFunnels**](ProcessGroupsApi.md#GetFunnels) | **Get** /process-groups/{id}/funnels | Gets all funnels
[**GetInputPorts**](ProcessGroupsApi.md#GetInputPorts) | **Get** /process-groups/{id}/input-ports | Gets all input ports
[**GetLabels**](ProcessGroupsApi.md#GetLabels) | **Get** /process-groups/{id}/labels | Gets all labels
[**GetLocalModifications**](ProcessGroupsApi.md#GetLocalModifications) | **Get** /process-groups/{id}/local-modifications | Gets a list of local modifications to the Process Group since it was last synchronized with the Flow Registry
[**GetOutputPorts**](ProcessGroupsApi.md#GetOutputPorts) | **Get** /process-groups/{id}/output-ports | Gets all output ports
[**GetProcessGroup**](ProcessGroupsApi.md#GetProcessGroup) | **Get** /process-groups/{id} | Gets a process group
[**GetProcessGroups**](ProcessGroupsApi.md#GetProcessGroups) | **Get** /process-groups/{id}/process-groups | Gets all process groups
[**GetProcessors**](ProcessGroupsApi.md#GetProcessors) | **Get** /process-groups/{id}/processors | Gets all processors
[**GetRemoteProcessGroups**](ProcessGroupsApi.md#GetRemoteProcessGroups) | **Get** /process-groups/{id}/remote-process-groups | Gets all remote process groups
[**GetVariableRegistry**](ProcessGroupsApi.md#GetVariableRegistry) | **Get** /process-groups/{id}/variable-registry | Gets a process group&#39;s variable registry
[**GetVariableRegistryUpdateRequest**](ProcessGroupsApi.md#GetVariableRegistryUpdateRequest) | **Get** /process-groups/{groupId}/variable-registry/update-requests/{updateId} | Gets a process group&#39;s variable registry
[**ImportTemplate**](ProcessGroupsApi.md#ImportTemplate) | **Post** /process-groups/{id}/templates/import | Imports a template
[**InstantiateTemplate**](ProcessGroupsApi.md#InstantiateTemplate) | **Post** /process-groups/{id}/template-instance | Instantiates a template
[**RemoveProcessGroup**](ProcessGroupsApi.md#RemoveProcessGroup) | **Delete** /process-groups/{id} | Deletes a process group
[**SubmitUpdateVariableRegistryRequest**](ProcessGroupsApi.md#SubmitUpdateVariableRegistryRequest) | **Post** /process-groups/{id}/variable-registry/update-requests | Submits a request to update a process group&#39;s variable registry
[**UpdateProcessGroup**](ProcessGroupsApi.md#UpdateProcessGroup) | **Put** /process-groups/{id} | Updates a process group
[**UpdateVariableRegistry**](ProcessGroupsApi.md#UpdateVariableRegistry) | **Put** /process-groups/{id}/variable-registry | Updates the contents of a Process Group&#39;s variable Registry
[**UploadTemplate**](ProcessGroupsApi.md#UploadTemplate) | **Post** /process-groups/{id}/templates/upload | Uploads a template



## CopySnippet

> FlowEntity CopySnippet(ctx, id, body)

Copies a snippet and discards it.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The process group id. | 
**body** | [**CopySnippetRequestEntity**](CopySnippetRequestEntity.md)| The copy snippet request. | 

### Return type

[**FlowEntity**](FlowEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateConnection

> ConnectionEntity CreateConnection(ctx, id, body)

Creates a connection

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The process group id. | 
**body** | [**ConnectionEntity**](ConnectionEntity.md)| The connection configuration details. | 

### Return type

[**ConnectionEntity**](ConnectionEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateControllerService

> ControllerServiceEntity CreateControllerService(ctx, id, body)

Creates a new controller service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The process group id. | 
**body** | [**ControllerServiceEntity**](ControllerServiceEntity.md)| The controller service configuration details. | 

### Return type

[**ControllerServiceEntity**](ControllerServiceEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateFunnel

> FunnelEntity CreateFunnel(ctx, id, body)

Creates a funnel

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The process group id. | 
**body** | [**FunnelEntity**](FunnelEntity.md)| The funnel configuration details. | 

### Return type

[**FunnelEntity**](FunnelEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateInputPort

> PortEntity CreateInputPort(ctx, id, body)

Creates an input port

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The process group id. | 
**body** | [**PortEntity**](PortEntity.md)| The input port configuration details. | 

### Return type

[**PortEntity**](PortEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateLabel

> LabelEntity CreateLabel(ctx, id, body)

Creates a label

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The process group id. | 
**body** | [**LabelEntity**](LabelEntity.md)| The label configuration details. | 

### Return type

[**LabelEntity**](LabelEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOutputPort

> PortEntity CreateOutputPort(ctx, id, body)

Creates an output port

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The process group id. | 
**body** | [**PortEntity**](PortEntity.md)| The output port configuration. | 

### Return type

[**PortEntity**](PortEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateProcessGroup

> ProcessGroupEntity CreateProcessGroup(ctx, id, body)

Creates a process group

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The process group id. | 
**body** | [**ProcessGroupEntity**](ProcessGroupEntity.md)| The process group configuration details. | 

### Return type

[**ProcessGroupEntity**](ProcessGroupEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateProcessor

> ProcessorEntity CreateProcessor(ctx, id, body)

Creates a new processor

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The process group id. | 
**body** | [**ProcessorEntity**](ProcessorEntity.md)| The processor configuration details. | 

### Return type

[**ProcessorEntity**](ProcessorEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRemoteProcessGroup

> RemoteProcessGroupEntity CreateRemoteProcessGroup(ctx, id, body)

Creates a new process group

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The process group id. | 
**body** | [**RemoteProcessGroupEntity**](RemoteProcessGroupEntity.md)| The remote process group configuration details. | 

### Return type

[**RemoteProcessGroupEntity**](RemoteProcessGroupEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTemplate

> TemplateEntity CreateTemplate(ctx, id, body)

Creates a template and discards the specified snippet.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The process group id. | 
**body** | [**CreateTemplateRequestEntity**](CreateTemplateRequestEntity.md)| The create template request. | 

### Return type

[**TemplateEntity**](TemplateEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVariableRegistryUpdateRequest

> VariableRegistryUpdateRequestEntity DeleteVariableRegistryUpdateRequest(ctx, groupId, updateId, optional)

Deletes an update request for a process group's variable registry. If the request is not yet complete, it will automatically be cancelled.

Note: This endpoint is subject to change as NiFi and it's REST API evolve.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| The process group id. | 
**updateId** | **string**| The ID of the Variable Registry Update Request | 
 **optional** | ***DeleteVariableRegistryUpdateRequestOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteVariableRegistryUpdateRequestOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **disconnectedNodeAcknowledged** | **optional.Bool**| Acknowledges that this node is disconnected to allow for mutable requests to proceed. | [default to false]

### Return type

[**VariableRegistryUpdateRequestEntity**](VariableRegistryUpdateRequestEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportProcessGroup

> string ExportProcessGroup(ctx, id)

Gets a process group for download

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The process group id. | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConnections

> ConnectionsEntity GetConnections(ctx, id)

Gets all connections

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The process group id. | 

### Return type

[**ConnectionsEntity**](ConnectionsEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFunnels

> FunnelsEntity GetFunnels(ctx, id)

Gets all funnels

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The process group id. | 

### Return type

[**FunnelsEntity**](FunnelsEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInputPorts

> InputPortsEntity GetInputPorts(ctx, id)

Gets all input ports

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The process group id. | 

### Return type

[**InputPortsEntity**](InputPortsEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLabels

> LabelsEntity GetLabels(ctx, id)

Gets all labels

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The process group id. | 

### Return type

[**LabelsEntity**](LabelsEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLocalModifications

> FlowComparisonEntity GetLocalModifications(ctx, id)

Gets a list of local modifications to the Process Group since it was last synchronized with the Flow Registry

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The process group id. | 

### Return type

[**FlowComparisonEntity**](FlowComparisonEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOutputPorts

> OutputPortsEntity GetOutputPorts(ctx, id)

Gets all output ports

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The process group id. | 

### Return type

[**OutputPortsEntity**](OutputPortsEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProcessGroup

> ProcessGroupEntity GetProcessGroup(ctx, id)

Gets a process group

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The process group id. | 

### Return type

[**ProcessGroupEntity**](ProcessGroupEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProcessGroups

> ProcessGroupsEntity GetProcessGroups(ctx, id)

Gets all process groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The process group id. | 

### Return type

[**ProcessGroupsEntity**](ProcessGroupsEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProcessors

> ProcessorsEntity GetProcessors(ctx, id, optional)

Gets all processors

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The process group id. | 
 **optional** | ***GetProcessorsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetProcessorsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeDescendantGroups** | **optional.Bool**| Whether or not to include processors from descendant process groups | [default to false]

### Return type

[**ProcessorsEntity**](ProcessorsEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRemoteProcessGroups

> RemoteProcessGroupsEntity GetRemoteProcessGroups(ctx, id)

Gets all remote process groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The process group id. | 

### Return type

[**RemoteProcessGroupsEntity**](RemoteProcessGroupsEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVariableRegistry

> VariableRegistryEntity GetVariableRegistry(ctx, id, optional)

Gets a process group's variable registry

Note: This endpoint is subject to change as NiFi and it's REST API evolve.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The process group id. | 
 **optional** | ***GetVariableRegistryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetVariableRegistryOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeAncestorGroups** | **optional.Bool**| Whether or not to include ancestor groups | [default to true]

### Return type

[**VariableRegistryEntity**](VariableRegistryEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVariableRegistryUpdateRequest

> VariableRegistryUpdateRequestEntity GetVariableRegistryUpdateRequest(ctx, groupId, updateId)

Gets a process group's variable registry

Note: This endpoint is subject to change as NiFi and it's REST API evolve.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| The process group id. | 
**updateId** | **string**| The ID of the Variable Registry Update Request | 

### Return type

[**VariableRegistryUpdateRequestEntity**](VariableRegistryUpdateRequestEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportTemplate

> TemplateEntity ImportTemplate(ctx, id)

Imports a template

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The process group id. | 

### Return type

[**TemplateEntity**](TemplateEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InstantiateTemplate

> FlowEntity InstantiateTemplate(ctx, id, body)

Instantiates a template

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The process group id. | 
**body** | [**InstantiateTemplateRequestEntity**](InstantiateTemplateRequestEntity.md)| The instantiate template request. | 

### Return type

[**FlowEntity**](FlowEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveProcessGroup

> ProcessGroupEntity RemoveProcessGroup(ctx, id, optional)

Deletes a process group

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The process group id. | 
 **optional** | ***RemoveProcessGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RemoveProcessGroupOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **optional.String**| The revision is used to verify the client is working with the latest version of the flow. | 
 **clientId** | **optional.String**| If the client id is not specified, new one will be generated. This value (whether specified or generated) is included in the response. | 
 **disconnectedNodeAcknowledged** | **optional.Bool**| Acknowledges that this node is disconnected to allow for mutable requests to proceed. | [default to false]

### Return type

[**ProcessGroupEntity**](ProcessGroupEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitUpdateVariableRegistryRequest

> VariableRegistryUpdateRequestEntity SubmitUpdateVariableRegistryRequest(ctx, id, body)

Submits a request to update a process group's variable registry

Note: This endpoint is subject to change as NiFi and it's REST API evolve.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The process group id. | 
**body** | [**VariableRegistryEntity**](VariableRegistryEntity.md)| The variable registry configuration details. | 

### Return type

[**VariableRegistryUpdateRequestEntity**](VariableRegistryUpdateRequestEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProcessGroup

> ProcessGroupEntity UpdateProcessGroup(ctx, id, body)

Updates a process group

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The process group id. | 
**body** | [**ProcessGroupEntity**](ProcessGroupEntity.md)| The process group configuration details. | 

### Return type

[**ProcessGroupEntity**](ProcessGroupEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateVariableRegistry

> VariableRegistryEntity UpdateVariableRegistry(ctx, id, body)

Updates the contents of a Process Group's variable Registry

Note: This endpoint is subject to change as NiFi and it's REST API evolve.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The process group id. | 
**body** | [**VariableRegistryEntity**](VariableRegistryEntity.md)| The variable registry configuration details. | 

### Return type

[**VariableRegistryEntity**](VariableRegistryEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadTemplate

> TemplateEntity UploadTemplate(ctx, id, template)

Uploads a template

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The process group id. | 
**template** | ***os.File*****os.File**| The binary content of the template file being uploaded. | 

### Return type

[**TemplateEntity**](TemplateEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

