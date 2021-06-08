# \CountersApi

All URIs are relative to *http://localhost/nifi-api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCounters**](CountersApi.md#GetCounters) | **Get** /counters | Gets the current counters for this NiFi
[**UpdateCounter**](CountersApi.md#UpdateCounter) | **Put** /counters/{id} | Updates the specified counter. This will reset the counter value to 0



## GetCounters

> CountersEntity GetCounters(ctx).Nodewise(nodewise).ClusterNodeId(clusterNodeId).Execute()

Gets the current counters for this NiFi



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
    resp, r, err := api_client.CountersApi.GetCounters(context.Background()).Nodewise(nodewise).ClusterNodeId(clusterNodeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CountersApi.GetCounters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCounters`: CountersEntity
    fmt.Fprintf(os.Stdout, "Response from `CountersApi.GetCounters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCountersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nodewise** | **bool** | Whether or not to include the breakdown per node. Optional, defaults to false | [default to false]
 **clusterNodeId** | **string** | The id of the node where to get the status. | 

### Return type

[**CountersEntity**](CountersEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCounter

> CounterEntity UpdateCounter(ctx, id).Execute()

Updates the specified counter. This will reset the counter value to 0



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
    id := "id_example" // string | The id of the counter.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CountersApi.UpdateCounter(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CountersApi.UpdateCounter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCounter`: CounterEntity
    fmt.Fprintf(os.Stdout, "Response from `CountersApi.UpdateCounter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the counter. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCounterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CounterEntity**](CounterEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

