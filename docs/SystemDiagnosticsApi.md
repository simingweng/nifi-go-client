# \SystemDiagnosticsApi

All URIs are relative to *http://localhost/nifi-api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSystemDiagnostics**](SystemDiagnosticsApi.md#GetSystemDiagnostics) | **Get** /system-diagnostics | Gets the diagnostics for the system NiFi is running on



## GetSystemDiagnostics

> SystemDiagnosticsEntity GetSystemDiagnostics(ctx, optional)

Gets the diagnostics for the system NiFi is running on

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetSystemDiagnosticsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetSystemDiagnosticsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nodewise** | **optional.Bool**| Whether or not to include the breakdown per node. Optional, defaults to false | [default to false]
 **clusterNodeId** | **optional.String**| The id of the node where to get the status. | 

### Return type

[**SystemDiagnosticsEntity**](SystemDiagnosticsEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

