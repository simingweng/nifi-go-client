# \ControllerApi

All URIs are relative to *http://localhost/nifi-api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBulletin**](ControllerApi.md#CreateBulletin) | **Post** /controller/bulletin | Creates a new bulletin
[**CreateControllerService**](ControllerApi.md#CreateControllerService) | **Post** /controller/controller-services | Creates a new controller service
[**CreateRegistryClient**](ControllerApi.md#CreateRegistryClient) | **Post** /controller/registry-clients | Creates a new registry client
[**CreateReportingTask**](ControllerApi.md#CreateReportingTask) | **Post** /controller/reporting-tasks | Creates a new reporting task
[**DeleteHistory**](ControllerApi.md#DeleteHistory) | **Delete** /controller/history | Purges history
[**DeleteNode**](ControllerApi.md#DeleteNode) | **Delete** /controller/cluster/nodes/{id} | Removes a node from the cluster
[**DeleteRegistryClient**](ControllerApi.md#DeleteRegistryClient) | **Delete** /controller/registry-clients/{id} | Deletes a registry client
[**GetCluster**](ControllerApi.md#GetCluster) | **Get** /controller/cluster | Gets the contents of the cluster
[**GetControllerConfig**](ControllerApi.md#GetControllerConfig) | **Get** /controller/config | Retrieves the configuration for this NiFi Controller
[**GetNode**](ControllerApi.md#GetNode) | **Get** /controller/cluster/nodes/{id} | Gets a node in the cluster
[**GetRegistryClient**](ControllerApi.md#GetRegistryClient) | **Get** /controller/registry-clients/{id} | Gets a registry client
[**GetRegistryClients**](ControllerApi.md#GetRegistryClients) | **Get** /controller/registry-clients | Gets the listing of available registry clients
[**UpdateControllerConfig**](ControllerApi.md#UpdateControllerConfig) | **Put** /controller/config | Retrieves the configuration for this NiFi
[**UpdateNode**](ControllerApi.md#UpdateNode) | **Put** /controller/cluster/nodes/{id} | Updates a node in the cluster
[**UpdateRegistryClient**](ControllerApi.md#UpdateRegistryClient) | **Put** /controller/registry-clients/{id} | Updates a registry client



## CreateBulletin

> BulletinEntity CreateBulletin(ctx, body)

Creates a new bulletin

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**BulletinEntity**](BulletinEntity.md)| The reporting task configuration details. | 

### Return type

[**BulletinEntity**](BulletinEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateControllerService

> ControllerServiceEntity CreateControllerService(ctx, body)

Creates a new controller service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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


## CreateRegistryClient

> RegistryClientEntity CreateRegistryClient(ctx, body)

Creates a new registry client

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**RegistryClientEntity**](RegistryClientEntity.md)| The registry configuration details. | 

### Return type

[**RegistryClientEntity**](RegistryClientEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateReportingTask

> ReportingTaskEntity CreateReportingTask(ctx, body)

Creates a new reporting task

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**ReportingTaskEntity**](ReportingTaskEntity.md)| The reporting task configuration details. | 

### Return type

[**ReportingTaskEntity**](ReportingTaskEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteHistory

> HistoryEntity DeleteHistory(ctx, endDate)

Purges history

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**endDate** | **string**| Purge actions before this date/time. | 

### Return type

[**HistoryEntity**](HistoryEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNode

> NodeEntity DeleteNode(ctx, id)

Removes a node from the cluster

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The node id. | 

### Return type

[**NodeEntity**](NodeEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRegistryClient

> RegistryClientEntity DeleteRegistryClient(ctx, id, optional)

Deletes a registry client

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The registry id. | 
 **optional** | ***DeleteRegistryClientOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteRegistryClientOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **optional.String**| The revision is used to verify the client is working with the latest version of the flow. | 
 **clientId** | **optional.String**| If the client id is not specified, new one will be generated. This value (whether specified or generated) is included in the response. | 
 **disconnectedNodeAcknowledged** | **optional.Bool**| Acknowledges that this node is disconnected to allow for mutable requests to proceed. | [default to false]

### Return type

[**RegistryClientEntity**](RegistryClientEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCluster

> ClusterEntity GetCluster(ctx, )

Gets the contents of the cluster

Returns the contents of the cluster including all nodes and their status.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**ClusterEntity**](ClusterEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetControllerConfig

> ControllerConfigurationEntity GetControllerConfig(ctx, )

Retrieves the configuration for this NiFi Controller

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**ControllerConfigurationEntity**](ControllerConfigurationEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNode

> NodeEntity GetNode(ctx, id)

Gets a node in the cluster

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The node id. | 

### Return type

[**NodeEntity**](NodeEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRegistryClient

> RegistryClientEntity GetRegistryClient(ctx, id)

Gets a registry client

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The registry id. | 

### Return type

[**RegistryClientEntity**](RegistryClientEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRegistryClients

> RegistryClientsEntity GetRegistryClients(ctx, )

Gets the listing of available registry clients

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**RegistryClientsEntity**](RegistryClientsEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateControllerConfig

> ControllerConfigurationEntity UpdateControllerConfig(ctx, body)

Retrieves the configuration for this NiFi

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**ControllerConfigurationEntity**](ControllerConfigurationEntity.md)| The controller configuration. | 

### Return type

[**ControllerConfigurationEntity**](ControllerConfigurationEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNode

> NodeEntity UpdateNode(ctx, id, body)

Updates a node in the cluster

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The node id. | 
**body** | [**NodeEntity**](NodeEntity.md)| The node configuration. The only configuration that will be honored at this endpoint is the status. | 

### Return type

[**NodeEntity**](NodeEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRegistryClient

> RegistryClientEntity UpdateRegistryClient(ctx, id, body)

Updates a registry client

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The registry id. | 
**body** | [**RegistryClientEntity**](RegistryClientEntity.md)| The registry configuration details. | 

### Return type

[**RegistryClientEntity**](RegistryClientEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

