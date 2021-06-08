# \FunnelApi

All URIs are relative to *http://localhost/nifi-api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFunnel**](FunnelApi.md#GetFunnel) | **Get** /funnels/{id} | Gets a funnel
[**RemoveFunnel**](FunnelApi.md#RemoveFunnel) | **Delete** /funnels/{id} | Deletes a funnel
[**UpdateFunnel**](FunnelApi.md#UpdateFunnel) | **Put** /funnels/{id} | Updates a funnel



## GetFunnel

> FunnelEntity GetFunnel(ctx, id).Execute()

Gets a funnel

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
    id := "id_example" // string | The funnel id.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FunnelApi.GetFunnel(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FunnelApi.GetFunnel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFunnel`: FunnelEntity
    fmt.Fprintf(os.Stdout, "Response from `FunnelApi.GetFunnel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The funnel id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFunnelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FunnelEntity**](FunnelEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveFunnel

> FunnelEntity RemoveFunnel(ctx, id).Version(version).ClientId(clientId).DisconnectedNodeAcknowledged(disconnectedNodeAcknowledged).Execute()

Deletes a funnel

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
    id := "id_example" // string | The funnel id.
    version := "version_example" // string | The revision is used to verify the client is working with the latest version of the flow. (optional)
    clientId := "clientId_example" // string | If the client id is not specified, new one will be generated. This value (whether specified or generated) is included in the response. (optional)
    disconnectedNodeAcknowledged := true // bool | Acknowledges that this node is disconnected to allow for mutable requests to proceed. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FunnelApi.RemoveFunnel(context.Background(), id).Version(version).ClientId(clientId).DisconnectedNodeAcknowledged(disconnectedNodeAcknowledged).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FunnelApi.RemoveFunnel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveFunnel`: FunnelEntity
    fmt.Fprintf(os.Stdout, "Response from `FunnelApi.RemoveFunnel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The funnel id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveFunnelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **string** | The revision is used to verify the client is working with the latest version of the flow. | 
 **clientId** | **string** | If the client id is not specified, new one will be generated. This value (whether specified or generated) is included in the response. | 
 **disconnectedNodeAcknowledged** | **bool** | Acknowledges that this node is disconnected to allow for mutable requests to proceed. | [default to false]

### Return type

[**FunnelEntity**](FunnelEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFunnel

> FunnelEntity UpdateFunnel(ctx, id).Body(body).Execute()

Updates a funnel

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
    id := "id_example" // string | The funnel id.
    body := *openapiclient.NewFunnelEntity() // FunnelEntity | The funnel configuration details.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FunnelApi.UpdateFunnel(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FunnelApi.UpdateFunnel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateFunnel`: FunnelEntity
    fmt.Fprintf(os.Stdout, "Response from `FunnelApi.UpdateFunnel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The funnel id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFunnelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**FunnelEntity**](FunnelEntity.md) | The funnel configuration details. | 

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

