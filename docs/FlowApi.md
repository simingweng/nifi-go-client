# \FlowApi

All URIs are relative to *http://localhost/nifi-api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateControllerServices**](FlowApi.md#ActivateControllerServices) | **Put** /flow/process-groups/{id}/controller-services | Enable or disable Controller Services in the specified Process Group.
[**GenerateClientId**](FlowApi.md#GenerateClientId) | **Get** /flow/client-id | Generates a client id.
[**GetAboutInfo**](FlowApi.md#GetAboutInfo) | **Get** /flow/about | Retrieves details about this NiFi to put in the About dialog
[**GetAction**](FlowApi.md#GetAction) | **Get** /flow/history/{id} | Gets an action
[**GetBanners**](FlowApi.md#GetBanners) | **Get** /flow/banners | Retrieves the banners for this NiFi
[**GetBuckets**](FlowApi.md#GetBuckets) | **Get** /flow/registries/{id}/buckets | Gets the buckets from the specified registry for the current user
[**GetBulletinBoard**](FlowApi.md#GetBulletinBoard) | **Get** /flow/bulletin-board | Gets current bulletins
[**GetBulletins**](FlowApi.md#GetBulletins) | **Get** /flow/controller/bulletins | Retrieves Controller level bulletins
[**GetClusterSummary**](FlowApi.md#GetClusterSummary) | **Get** /flow/cluster/summary | The cluster summary for this NiFi
[**GetComponentHistory**](FlowApi.md#GetComponentHistory) | **Get** /flow/history/components/{componentId} | Gets configuration history for a component
[**GetConnectionStatistics**](FlowApi.md#GetConnectionStatistics) | **Get** /flow/connections/{id}/statistics | Gets statistics for a connection
[**GetConnectionStatus**](FlowApi.md#GetConnectionStatus) | **Get** /flow/connections/{id}/status | Gets status for a connection
[**GetConnectionStatusHistory**](FlowApi.md#GetConnectionStatusHistory) | **Get** /flow/connections/{id}/status/history | Gets the status history for a connection
[**GetControllerServiceTypes**](FlowApi.md#GetControllerServiceTypes) | **Get** /flow/controller-service-types | Retrieves the types of controller services that this NiFi supports
[**GetControllerServicesFromController**](FlowApi.md#GetControllerServicesFromController) | **Get** /flow/controller/controller-services | Gets controller services for reporting tasks
[**GetControllerServicesFromGroup**](FlowApi.md#GetControllerServicesFromGroup) | **Get** /flow/process-groups/{id}/controller-services | Gets all controller services
[**GetControllerStatus**](FlowApi.md#GetControllerStatus) | **Get** /flow/status | Gets the current status of this NiFi
[**GetCurrentUser**](FlowApi.md#GetCurrentUser) | **Get** /flow/current-user | Retrieves the user identity of the user making the request
[**GetFlow**](FlowApi.md#GetFlow) | **Get** /flow/process-groups/{id} | Gets a process group
[**GetFlowConfig**](FlowApi.md#GetFlowConfig) | **Get** /flow/config | Retrieves the configuration for this NiFi flow
[**GetFlowMetrics**](FlowApi.md#GetFlowMetrics) | **Get** /flow/metrics/{producer} | Gets all metrics for the flow from a particular node
[**GetFlows**](FlowApi.md#GetFlows) | **Get** /flow/registries/{registry-id}/buckets/{bucket-id}/flows | Gets the flows from the specified registry and bucket for the current user
[**GetInputPortStatus**](FlowApi.md#GetInputPortStatus) | **Get** /flow/input-ports/{id}/status | Gets status for an input port
[**GetOutputPortStatus**](FlowApi.md#GetOutputPortStatus) | **Get** /flow/output-ports/{id}/status | Gets status for an output port
[**GetParameterContexts**](FlowApi.md#GetParameterContexts) | **Get** /flow/parameter-contexts | Gets all Parameter Contexts
[**GetPrioritizers**](FlowApi.md#GetPrioritizers) | **Get** /flow/prioritizers | Retrieves the types of prioritizers that this NiFi supports
[**GetProcessGroupStatus**](FlowApi.md#GetProcessGroupStatus) | **Get** /flow/process-groups/{id}/status | Gets the status for a process group
[**GetProcessGroupStatusHistory**](FlowApi.md#GetProcessGroupStatusHistory) | **Get** /flow/process-groups/{id}/status/history | Gets status history for a remote process group
[**GetProcessorStatus**](FlowApi.md#GetProcessorStatus) | **Get** /flow/processors/{id}/status | Gets status for a processor
[**GetProcessorStatusHistory**](FlowApi.md#GetProcessorStatusHistory) | **Get** /flow/processors/{id}/status/history | Gets status history for a processor
[**GetProcessorTypes**](FlowApi.md#GetProcessorTypes) | **Get** /flow/processor-types | Retrieves the types of processors that this NiFi supports
[**GetRegistries**](FlowApi.md#GetRegistries) | **Get** /flow/registries | Gets the listing of available registries
[**GetRemoteProcessGroupStatus**](FlowApi.md#GetRemoteProcessGroupStatus) | **Get** /flow/remote-process-groups/{id}/status | Gets status for a remote process group
[**GetRemoteProcessGroupStatusHistory**](FlowApi.md#GetRemoteProcessGroupStatusHistory) | **Get** /flow/remote-process-groups/{id}/status/history | Gets the status history
[**GetReportingTaskTypes**](FlowApi.md#GetReportingTaskTypes) | **Get** /flow/reporting-task-types | Retrieves the types of reporting tasks that this NiFi supports
[**GetReportingTasks**](FlowApi.md#GetReportingTasks) | **Get** /flow/reporting-tasks | Gets all reporting tasks
[**GetTemplates**](FlowApi.md#GetTemplates) | **Get** /flow/templates | Gets all templates
[**GetVersions**](FlowApi.md#GetVersions) | **Get** /flow/registries/{registry-id}/buckets/{bucket-id}/flows/{flow-id}/versions | Gets the flow versions from the specified registry and bucket for the specified flow for the current user
[**QueryHistory**](FlowApi.md#QueryHistory) | **Get** /flow/history | Gets configuration history
[**ScheduleComponents**](FlowApi.md#ScheduleComponents) | **Put** /flow/process-groups/{id} | Schedule or unschedule components in the specified Process Group.
[**SearchCluster**](FlowApi.md#SearchCluster) | **Get** /flow/cluster/search-results | Searches the cluster for a node with the specified address
[**SearchFlow**](FlowApi.md#SearchFlow) | **Get** /flow/search-results | Performs a search against this NiFi using the specified search term



## ActivateControllerServices

> ActivateControllerServicesEntity ActivateControllerServices(ctx, id, body)

Enable or disable Controller Services in the specified Process Group.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The process group id. | 
**body** | [**ActivateControllerServicesEntity**](ActivateControllerServicesEntity.md)| The request to schedule or unschedule. If the comopnents in the request are not specified, all authorized components will be considered. | 

### Return type

[**ActivateControllerServicesEntity**](ActivateControllerServicesEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateClientId

> string GenerateClientId(ctx, )

Generates a client id.

### Required Parameters

This endpoint does not need any parameter.

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAboutInfo

> AboutEntity GetAboutInfo(ctx, )

Retrieves details about this NiFi to put in the About dialog

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**AboutEntity**](AboutEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAction

> ActionEntity GetAction(ctx, id)

Gets an action

Note: This endpoint is subject to change as NiFi and it's REST API evolve.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The action id. | 

### Return type

[**ActionEntity**](ActionEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBanners

> BannerEntity GetBanners(ctx, )

Retrieves the banners for this NiFi

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**BannerEntity**](BannerEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBuckets

> BucketsEntity GetBuckets(ctx, id)

Gets the buckets from the specified registry for the current user

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The registry id. | 

### Return type

[**BucketsEntity**](BucketsEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBulletinBoard

> BulletinBoardEntity GetBulletinBoard(ctx, optional)

Gets current bulletins

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetBulletinBoardOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetBulletinBoardOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **after** | **optional.String**| Includes bulletins with an id after this value. | 
 **sourceName** | **optional.String**| Includes bulletins originating from this sources whose name match this regular expression. | 
 **message** | **optional.String**| Includes bulletins whose message that match this regular expression. | 
 **sourceId** | **optional.String**| Includes bulletins originating from this sources whose id match this regular expression. | 
 **groupId** | **optional.String**| Includes bulletins originating from this sources whose group id match this regular expression. | 
 **limit** | **optional.String**| The number of bulletins to limit the response to. | 

### Return type

[**BulletinBoardEntity**](BulletinBoardEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBulletins

> ControllerBulletinsEntity GetBulletins(ctx, )

Retrieves Controller level bulletins

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**ControllerBulletinsEntity**](ControllerBulletinsEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClusterSummary

> ClusteSummaryEntity GetClusterSummary(ctx, )

The cluster summary for this NiFi

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**ClusteSummaryEntity**](ClusteSummaryEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetComponentHistory

> ComponentHistoryEntity GetComponentHistory(ctx, componentId)

Gets configuration history for a component

Note: This endpoint is subject to change as NiFi and it's REST API evolve.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**componentId** | **string**| The component id. | 

### Return type

[**ComponentHistoryEntity**](ComponentHistoryEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConnectionStatistics

> ConnectionStatisticsEntity GetConnectionStatistics(ctx, id, optional)

Gets statistics for a connection

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The connection id. | 
 **optional** | ***GetConnectionStatisticsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetConnectionStatisticsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nodewise** | **optional.Bool**| Whether or not to include the breakdown per node. Optional, defaults to false | [default to false]
 **clusterNodeId** | **optional.String**| The id of the node where to get the statistics. | 

### Return type

[**ConnectionStatisticsEntity**](ConnectionStatisticsEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConnectionStatus

> ConnectionStatusEntity GetConnectionStatus(ctx, id, optional)

Gets status for a connection

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The connection id. | 
 **optional** | ***GetConnectionStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetConnectionStatusOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nodewise** | **optional.Bool**| Whether or not to include the breakdown per node. Optional, defaults to false | [default to false]
 **clusterNodeId** | **optional.String**| The id of the node where to get the status. | 

### Return type

[**ConnectionStatusEntity**](ConnectionStatusEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConnectionStatusHistory

> StatusHistoryEntity GetConnectionStatusHistory(ctx, id)

Gets the status history for a connection

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The connection id. | 

### Return type

[**StatusHistoryEntity**](StatusHistoryEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetControllerServiceTypes

> ControllerServiceTypesEntity GetControllerServiceTypes(ctx, optional)

Retrieves the types of controller services that this NiFi supports

Note: This endpoint is subject to change as NiFi and it's REST API evolve.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetControllerServiceTypesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetControllerServiceTypesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serviceType** | **optional.String**| If specified, will only return controller services that are compatible with this type of service. | 
 **serviceBundleGroup** | **optional.String**| If serviceType specified, is the bundle group of the serviceType. | 
 **serviceBundleArtifact** | **optional.String**| If serviceType specified, is the bundle artifact of the serviceType. | 
 **serviceBundleVersion** | **optional.String**| If serviceType specified, is the bundle version of the serviceType. | 
 **bundleGroupFilter** | **optional.String**| If specified, will only return types that are a member of this bundle group. | 
 **bundleArtifactFilter** | **optional.String**| If specified, will only return types that are a member of this bundle artifact. | 
 **typeFilter** | **optional.String**| If specified, will only return types whose fully qualified classname matches. | 

### Return type

[**ControllerServiceTypesEntity**](ControllerServiceTypesEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetControllerServicesFromController

> ControllerServicesEntity GetControllerServicesFromController(ctx, )

Gets controller services for reporting tasks

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**ControllerServicesEntity**](ControllerServicesEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetControllerServicesFromGroup

> ControllerServicesEntity GetControllerServicesFromGroup(ctx, id, optional)

Gets all controller services

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The process group id. | 
 **optional** | ***GetControllerServicesFromGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetControllerServicesFromGroupOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeAncestorGroups** | **optional.Bool**| Whether or not to include parent/ancestory process groups | [default to true]
 **includeDescendantGroups** | **optional.Bool**| Whether or not to include descendant process groups | [default to false]

### Return type

[**ControllerServicesEntity**](ControllerServicesEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetControllerStatus

> ControllerStatusEntity GetControllerStatus(ctx, )

Gets the current status of this NiFi

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**ControllerStatusEntity**](ControllerStatusEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCurrentUser

> CurrentUserEntity GetCurrentUser(ctx, )

Retrieves the user identity of the user making the request

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**CurrentUserEntity**](CurrentUserEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFlow

> ProcessGroupFlowEntity GetFlow(ctx, id)

Gets a process group

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The process group id. | 

### Return type

[**ProcessGroupFlowEntity**](ProcessGroupFlowEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFlowConfig

> FlowConfigurationEntity GetFlowConfig(ctx, )

Retrieves the configuration for this NiFi flow

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**FlowConfigurationEntity**](FlowConfigurationEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFlowMetrics

> map[string]interface{} GetFlowMetrics(ctx, producer)

Gets all metrics for the flow from a particular node

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**producer** | **string**| The producer for flow file metrics. Each producer may have its own output format. | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFlows

> VersionedFlowsEntity GetFlows(ctx, registryId, bucketId)

Gets the flows from the specified registry and bucket for the current user

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registryId** | **string**| The registry id. | 
**bucketId** | **string**| The bucket id. | 

### Return type

[**VersionedFlowsEntity**](VersionedFlowsEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInputPortStatus

> PortStatusEntity GetInputPortStatus(ctx, id, optional)

Gets status for an input port

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The input port id. | 
 **optional** | ***GetInputPortStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetInputPortStatusOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nodewise** | **optional.Bool**| Whether or not to include the breakdown per node. Optional, defaults to false | [default to false]
 **clusterNodeId** | **optional.String**| The id of the node where to get the status. | 

### Return type

[**PortStatusEntity**](PortStatusEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOutputPortStatus

> PortStatusEntity GetOutputPortStatus(ctx, id, optional)

Gets status for an output port

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The output port id. | 
 **optional** | ***GetOutputPortStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetOutputPortStatusOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nodewise** | **optional.Bool**| Whether or not to include the breakdown per node. Optional, defaults to false | [default to false]
 **clusterNodeId** | **optional.String**| The id of the node where to get the status. | 

### Return type

[**PortStatusEntity**](PortStatusEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetParameterContexts

> ParameterContextsEntity GetParameterContexts(ctx, )

Gets all Parameter Contexts

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**ParameterContextsEntity**](ParameterContextsEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPrioritizers

> PrioritizerTypesEntity GetPrioritizers(ctx, )

Retrieves the types of prioritizers that this NiFi supports

Note: This endpoint is subject to change as NiFi and it's REST API evolve.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**PrioritizerTypesEntity**](PrioritizerTypesEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProcessGroupStatus

> ProcessGroupStatusEntity GetProcessGroupStatus(ctx, id, optional)

Gets the status for a process group

The status for a process group includes status for all descendent components. When invoked on the root group with recursive set to true, it will return the current status of every component in the flow.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The process group id. | 
 **optional** | ***GetProcessGroupStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetProcessGroupStatusOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **recursive** | **optional.Bool**| Whether all descendant groups and the status of their content will be included. Optional, defaults to false | [default to false]
 **nodewise** | **optional.Bool**| Whether or not to include the breakdown per node. Optional, defaults to false | [default to false]
 **clusterNodeId** | **optional.String**| The id of the node where to get the status. | 

### Return type

[**ProcessGroupStatusEntity**](ProcessGroupStatusEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProcessGroupStatusHistory

> StatusHistoryEntity GetProcessGroupStatusHistory(ctx, id)

Gets status history for a remote process group

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The process group id. | 

### Return type

[**StatusHistoryEntity**](StatusHistoryEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProcessorStatus

> ProcessorStatusEntity GetProcessorStatus(ctx, id, optional)

Gets status for a processor

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The processor id. | 
 **optional** | ***GetProcessorStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetProcessorStatusOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nodewise** | **optional.Bool**| Whether or not to include the breakdown per node. Optional, defaults to false | [default to false]
 **clusterNodeId** | **optional.String**| The id of the node where to get the status. | 

### Return type

[**ProcessorStatusEntity**](ProcessorStatusEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProcessorStatusHistory

> StatusHistoryEntity GetProcessorStatusHistory(ctx, id)

Gets status history for a processor

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The processor id. | 

### Return type

[**StatusHistoryEntity**](StatusHistoryEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProcessorTypes

> ProcessorTypesEntity GetProcessorTypes(ctx, optional)

Retrieves the types of processors that this NiFi supports

Note: This endpoint is subject to change as NiFi and it's REST API evolve.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetProcessorTypesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetProcessorTypesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bundleGroupFilter** | **optional.String**| If specified, will only return types that are a member of this bundle group. | 
 **bundleArtifactFilter** | **optional.String**| If specified, will only return types that are a member of this bundle artifact. | 
 **type_** | **optional.String**| If specified, will only return types whose fully qualified classname matches. | 

### Return type

[**ProcessorTypesEntity**](ProcessorTypesEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRegistries

> RegistryClientsEntity GetRegistries(ctx, )

Gets the listing of available registries

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


## GetRemoteProcessGroupStatus

> RemoteProcessGroupStatusEntity GetRemoteProcessGroupStatus(ctx, id, optional)

Gets status for a remote process group

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The remote process group id. | 
 **optional** | ***GetRemoteProcessGroupStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetRemoteProcessGroupStatusOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nodewise** | **optional.Bool**| Whether or not to include the breakdown per node. Optional, defaults to false | [default to false]
 **clusterNodeId** | **optional.String**| The id of the node where to get the status. | 

### Return type

[**RemoteProcessGroupStatusEntity**](RemoteProcessGroupStatusEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRemoteProcessGroupStatusHistory

> StatusHistoryEntity GetRemoteProcessGroupStatusHistory(ctx, id)

Gets the status history

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The remote process group id. | 

### Return type

[**StatusHistoryEntity**](StatusHistoryEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReportingTaskTypes

> ReportingTaskTypesEntity GetReportingTaskTypes(ctx, optional)

Retrieves the types of reporting tasks that this NiFi supports

Note: This endpoint is subject to change as NiFi and it's REST API evolve.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetReportingTaskTypesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetReportingTaskTypesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bundleGroupFilter** | **optional.String**| If specified, will only return types that are a member of this bundle group. | 
 **bundleArtifactFilter** | **optional.String**| If specified, will only return types that are a member of this bundle artifact. | 
 **type_** | **optional.String**| If specified, will only return types whose fully qualified classname matches. | 

### Return type

[**ReportingTaskTypesEntity**](ReportingTaskTypesEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReportingTasks

> ReportingTasksEntity GetReportingTasks(ctx, )

Gets all reporting tasks

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**ReportingTasksEntity**](ReportingTasksEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTemplates

> TemplatesEntity GetTemplates(ctx, )

Gets all templates

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**TemplatesEntity**](TemplatesEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVersions

> VersionedFlowSnapshotMetadataSetEntity GetVersions(ctx, registryId, bucketId, flowId)

Gets the flow versions from the specified registry and bucket for the specified flow for the current user

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registryId** | **string**| The registry id. | 
**bucketId** | **string**| The bucket id. | 
**flowId** | **string**| The flow id. | 

### Return type

[**VersionedFlowSnapshotMetadataSetEntity**](VersionedFlowSnapshotMetadataSetEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryHistory

> HistoryEntity QueryHistory(ctx, offset, count, optional)

Gets configuration history

Note: This endpoint is subject to change as NiFi and it's REST API evolve.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**offset** | **string**| The offset into the result set. | 
**count** | **string**| The number of actions to return. | 
 **optional** | ***QueryHistoryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a QueryHistoryOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sortColumn** | **optional.String**| The field to sort on. | 
 **sortOrder** | **optional.String**| The direction to sort. | 
 **startDate** | **optional.String**| Include actions after this date. | 
 **endDate** | **optional.String**| Include actions before this date. | 
 **userIdentity** | **optional.String**| Include actions performed by this user. | 
 **sourceId** | **optional.String**| Include actions on this component. | 

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


## ScheduleComponents

> ScheduleComponentsEntity ScheduleComponents(ctx, id, body)

Schedule or unschedule components in the specified Process Group.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The process group id. | 
**body** | [**ScheduleComponentsEntity**](ScheduleComponentsEntity.md)| The request to schedule or unschedule. If the comopnents in the request are not specified, all authorized components will be considered. | 

### Return type

[**ScheduleComponentsEntity**](ScheduleComponentsEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchCluster

> ClusterSearchResultsEntity SearchCluster(ctx, q)

Searches the cluster for a node with the specified address

Note: This endpoint is subject to change as NiFi and it's REST API evolve.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**q** | **string**| Node address to search for. | 

### Return type

[**ClusterSearchResultsEntity**](ClusterSearchResultsEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchFlow

> SearchResultsEntity SearchFlow(ctx, optional)

Performs a search against this NiFi using the specified search term

Only search results from authorized components will be returned.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SearchFlowOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SearchFlowOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **optional.String**|  | 
 **a** | **optional.String**|  | 

### Return type

[**SearchResultsEntity**](SearchResultsEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

