# \SystemDiagnosticsApi

All URIs are relative to *http://localhost/nifi-api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSystemDiagnostics**](SystemDiagnosticsApi.md#GetSystemDiagnostics) | **Get** /system-diagnostics | Gets the diagnostics for the system NiFi is running on



## GetSystemDiagnostics

> SystemDiagnosticsEntity GetSystemDiagnostics(ctx).Nodewise(nodewise).ClusterNodeId(clusterNodeId).Execute()

Gets the diagnostics for the system NiFi is running on

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    nodewise := true // bool | Whether or not to include the breakdown per node. Optional, defaults to false (optional) (default to false)
    clusterNodeId := "clusterNodeId_example" // string | The id of the node where to get the status. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SystemDiagnosticsApi.GetSystemDiagnostics(context.Background()).Nodewise(nodewise).ClusterNodeId(clusterNodeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemDiagnosticsApi.GetSystemDiagnostics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSystemDiagnostics`: SystemDiagnosticsEntity
    fmt.Fprintf(os.Stdout, "Response from `SystemDiagnosticsApi.GetSystemDiagnostics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemDiagnosticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nodewise** | **bool** | Whether or not to include the breakdown per node. Optional, defaults to false | [default to false]
 **clusterNodeId** | **string** | The id of the node where to get the status. | 

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

