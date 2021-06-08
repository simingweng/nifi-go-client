# \RemoteProcessGroupsApi

All URIs are relative to *http://localhost/nifi-api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRemoteProcessGroup**](RemoteProcessGroupsApi.md#GetRemoteProcessGroup) | **Get** /remote-process-groups/{id} | Gets a remote process group
[**GetState**](RemoteProcessGroupsApi.md#GetState) | **Get** /remote-process-groups/{id}/state | Gets the state for a RemoteProcessGroup
[**RemoveRemoteProcessGroup**](RemoteProcessGroupsApi.md#RemoveRemoteProcessGroup) | **Delete** /remote-process-groups/{id} | Deletes a remote process group
[**UpdateRemoteProcessGroup**](RemoteProcessGroupsApi.md#UpdateRemoteProcessGroup) | **Put** /remote-process-groups/{id} | Updates a remote process group
[**UpdateRemoteProcessGroupInputPort**](RemoteProcessGroupsApi.md#UpdateRemoteProcessGroupInputPort) | **Put** /remote-process-groups/{id}/input-ports/{port-id} | Updates a remote port
[**UpdateRemoteProcessGroupInputPortRunStatus**](RemoteProcessGroupsApi.md#UpdateRemoteProcessGroupInputPortRunStatus) | **Put** /remote-process-groups/{id}/input-ports/{port-id}/run-status | Updates run status of a remote port
[**UpdateRemoteProcessGroupOutputPort**](RemoteProcessGroupsApi.md#UpdateRemoteProcessGroupOutputPort) | **Put** /remote-process-groups/{id}/output-ports/{port-id} | Updates a remote port
[**UpdateRemoteProcessGroupOutputPortRunStatus**](RemoteProcessGroupsApi.md#UpdateRemoteProcessGroupOutputPortRunStatus) | **Put** /remote-process-groups/{id}/output-ports/{port-id}/run-status | Updates run status of a remote port
[**UpdateRemoteProcessGroupRunStatus**](RemoteProcessGroupsApi.md#UpdateRemoteProcessGroupRunStatus) | **Put** /remote-process-groups/{id}/run-status | Updates run status of a remote process group



## GetRemoteProcessGroup

> RemoteProcessGroupEntity GetRemoteProcessGroup(ctx, id).Execute()

Gets a remote process group

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
    id := "id_example" // string | The remote process group id.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RemoteProcessGroupsApi.GetRemoteProcessGroup(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemoteProcessGroupsApi.GetRemoteProcessGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRemoteProcessGroup`: RemoteProcessGroupEntity
    fmt.Fprintf(os.Stdout, "Response from `RemoteProcessGroupsApi.GetRemoteProcessGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The remote process group id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRemoteProcessGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RemoteProcessGroupEntity**](RemoteProcessGroupEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetState

> ComponentStateEntity GetState(ctx, id).Execute()

Gets the state for a RemoteProcessGroup

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
    id := "id_example" // string | The processor id.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RemoteProcessGroupsApi.GetState(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemoteProcessGroupsApi.GetState``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetState`: ComponentStateEntity
    fmt.Fprintf(os.Stdout, "Response from `RemoteProcessGroupsApi.GetState`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The processor id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ComponentStateEntity**](ComponentStateEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveRemoteProcessGroup

> RemoteProcessGroupEntity RemoveRemoteProcessGroup(ctx, id).Version(version).ClientId(clientId).DisconnectedNodeAcknowledged(disconnectedNodeAcknowledged).Execute()

Deletes a remote process group

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
    id := "id_example" // string | The remote process group id.
    version := "version_example" // string | The revision is used to verify the client is working with the latest version of the flow. (optional)
    clientId := "clientId_example" // string | If the client id is not specified, new one will be generated. This value (whether specified or generated) is included in the response. (optional)
    disconnectedNodeAcknowledged := true // bool | Acknowledges that this node is disconnected to allow for mutable requests to proceed. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RemoteProcessGroupsApi.RemoveRemoteProcessGroup(context.Background(), id).Version(version).ClientId(clientId).DisconnectedNodeAcknowledged(disconnectedNodeAcknowledged).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemoteProcessGroupsApi.RemoveRemoteProcessGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveRemoteProcessGroup`: RemoteProcessGroupEntity
    fmt.Fprintf(os.Stdout, "Response from `RemoteProcessGroupsApi.RemoveRemoteProcessGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The remote process group id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveRemoteProcessGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **string** | The revision is used to verify the client is working with the latest version of the flow. | 
 **clientId** | **string** | If the client id is not specified, new one will be generated. This value (whether specified or generated) is included in the response. | 
 **disconnectedNodeAcknowledged** | **bool** | Acknowledges that this node is disconnected to allow for mutable requests to proceed. | [default to false]

### Return type

[**RemoteProcessGroupEntity**](RemoteProcessGroupEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRemoteProcessGroup

> RemoteProcessGroupEntity UpdateRemoteProcessGroup(ctx, id).Body(body).Execute()

Updates a remote process group

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
    id := "id_example" // string | The remote process group id.
    body := *openapiclient.NewRemoteProcessGroupEntity() // RemoteProcessGroupEntity | The remote process group.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RemoteProcessGroupsApi.UpdateRemoteProcessGroup(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemoteProcessGroupsApi.UpdateRemoteProcessGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRemoteProcessGroup`: RemoteProcessGroupEntity
    fmt.Fprintf(os.Stdout, "Response from `RemoteProcessGroupsApi.UpdateRemoteProcessGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The remote process group id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRemoteProcessGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**RemoteProcessGroupEntity**](RemoteProcessGroupEntity.md) | The remote process group. | 

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


## UpdateRemoteProcessGroupInputPort

> RemoteProcessGroupPortEntity UpdateRemoteProcessGroupInputPort(ctx, id, portId).Body(body).Execute()

Updates a remote port



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
    id := "id_example" // string | The remote process group id.
    portId := "portId_example" // string | The remote process group port id.
    body := *openapiclient.NewRemoteProcessGroupPortEntity() // RemoteProcessGroupPortEntity | The remote process group port.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RemoteProcessGroupsApi.UpdateRemoteProcessGroupInputPort(context.Background(), id, portId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemoteProcessGroupsApi.UpdateRemoteProcessGroupInputPort``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRemoteProcessGroupInputPort`: RemoteProcessGroupPortEntity
    fmt.Fprintf(os.Stdout, "Response from `RemoteProcessGroupsApi.UpdateRemoteProcessGroupInputPort`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The remote process group id. | 
**portId** | **string** | The remote process group port id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRemoteProcessGroupInputPortRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**RemoteProcessGroupPortEntity**](RemoteProcessGroupPortEntity.md) | The remote process group port. | 

### Return type

[**RemoteProcessGroupPortEntity**](RemoteProcessGroupPortEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRemoteProcessGroupInputPortRunStatus

> RemoteProcessGroupPortEntity UpdateRemoteProcessGroupInputPortRunStatus(ctx, id, portId).Body(body).Execute()

Updates run status of a remote port



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
    id := "id_example" // string | The remote process group id.
    portId := "portId_example" // string | The remote process group port id.
    body := *openapiclient.NewRemotePortRunStatusEntity() // RemotePortRunStatusEntity | The remote process group port.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RemoteProcessGroupsApi.UpdateRemoteProcessGroupInputPortRunStatus(context.Background(), id, portId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemoteProcessGroupsApi.UpdateRemoteProcessGroupInputPortRunStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRemoteProcessGroupInputPortRunStatus`: RemoteProcessGroupPortEntity
    fmt.Fprintf(os.Stdout, "Response from `RemoteProcessGroupsApi.UpdateRemoteProcessGroupInputPortRunStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The remote process group id. | 
**portId** | **string** | The remote process group port id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRemoteProcessGroupInputPortRunStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**RemotePortRunStatusEntity**](RemotePortRunStatusEntity.md) | The remote process group port. | 

### Return type

[**RemoteProcessGroupPortEntity**](RemoteProcessGroupPortEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRemoteProcessGroupOutputPort

> RemoteProcessGroupPortEntity UpdateRemoteProcessGroupOutputPort(ctx, id, portId).Body(body).Execute()

Updates a remote port



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
    id := "id_example" // string | The remote process group id.
    portId := "portId_example" // string | The remote process group port id.
    body := *openapiclient.NewRemoteProcessGroupPortEntity() // RemoteProcessGroupPortEntity | The remote process group port.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RemoteProcessGroupsApi.UpdateRemoteProcessGroupOutputPort(context.Background(), id, portId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemoteProcessGroupsApi.UpdateRemoteProcessGroupOutputPort``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRemoteProcessGroupOutputPort`: RemoteProcessGroupPortEntity
    fmt.Fprintf(os.Stdout, "Response from `RemoteProcessGroupsApi.UpdateRemoteProcessGroupOutputPort`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The remote process group id. | 
**portId** | **string** | The remote process group port id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRemoteProcessGroupOutputPortRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**RemoteProcessGroupPortEntity**](RemoteProcessGroupPortEntity.md) | The remote process group port. | 

### Return type

[**RemoteProcessGroupPortEntity**](RemoteProcessGroupPortEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRemoteProcessGroupOutputPortRunStatus

> RemoteProcessGroupPortEntity UpdateRemoteProcessGroupOutputPortRunStatus(ctx, id, portId).Body(body).Execute()

Updates run status of a remote port



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
    id := "id_example" // string | The remote process group id.
    portId := "portId_example" // string | The remote process group port id.
    body := *openapiclient.NewRemotePortRunStatusEntity() // RemotePortRunStatusEntity | The remote process group port.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RemoteProcessGroupsApi.UpdateRemoteProcessGroupOutputPortRunStatus(context.Background(), id, portId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemoteProcessGroupsApi.UpdateRemoteProcessGroupOutputPortRunStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRemoteProcessGroupOutputPortRunStatus`: RemoteProcessGroupPortEntity
    fmt.Fprintf(os.Stdout, "Response from `RemoteProcessGroupsApi.UpdateRemoteProcessGroupOutputPortRunStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The remote process group id. | 
**portId** | **string** | The remote process group port id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRemoteProcessGroupOutputPortRunStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**RemotePortRunStatusEntity**](RemotePortRunStatusEntity.md) | The remote process group port. | 

### Return type

[**RemoteProcessGroupPortEntity**](RemoteProcessGroupPortEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRemoteProcessGroupRunStatus

> RemoteProcessGroupEntity UpdateRemoteProcessGroupRunStatus(ctx, id).Body(body).Execute()

Updates run status of a remote process group

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
    id := "id_example" // string | The remote process group id.
    body := *openapiclient.NewRemotePortRunStatusEntity() // RemotePortRunStatusEntity | The remote process group run status.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RemoteProcessGroupsApi.UpdateRemoteProcessGroupRunStatus(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemoteProcessGroupsApi.UpdateRemoteProcessGroupRunStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRemoteProcessGroupRunStatus`: RemoteProcessGroupEntity
    fmt.Fprintf(os.Stdout, "Response from `RemoteProcessGroupsApi.UpdateRemoteProcessGroupRunStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The remote process group id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRemoteProcessGroupRunStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**RemotePortRunStatusEntity**](RemotePortRunStatusEntity.md) | The remote process group run status. | 

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

