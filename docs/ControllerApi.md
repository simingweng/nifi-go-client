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
[**GetNodeStatusHistory**](ControllerApi.md#GetNodeStatusHistory) | **Get** /controller/status/history | Gets status history for the node
[**GetRegistryClient**](ControllerApi.md#GetRegistryClient) | **Get** /controller/registry-clients/{id} | Gets a registry client
[**GetRegistryClients**](ControllerApi.md#GetRegistryClients) | **Get** /controller/registry-clients | Gets the listing of available registry clients
[**UpdateControllerConfig**](ControllerApi.md#UpdateControllerConfig) | **Put** /controller/config | Retrieves the configuration for this NiFi
[**UpdateNode**](ControllerApi.md#UpdateNode) | **Put** /controller/cluster/nodes/{id} | Updates a node in the cluster
[**UpdateRegistryClient**](ControllerApi.md#UpdateRegistryClient) | **Put** /controller/registry-clients/{id} | Updates a registry client



## CreateBulletin

> BulletinEntity CreateBulletin(ctx).Body(body).Execute()

Creates a new bulletin

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
    body := *openapiclient.NewBulletinEntity() // BulletinEntity | The reporting task configuration details.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ControllerApi.CreateBulletin(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ControllerApi.CreateBulletin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBulletin`: BulletinEntity
    fmt.Fprintf(os.Stdout, "Response from `ControllerApi.CreateBulletin`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBulletinRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**BulletinEntity**](BulletinEntity.md) | The reporting task configuration details. | 

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

> ControllerServiceEntity CreateControllerService(ctx).Body(body).Execute()

Creates a new controller service

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
    body := *openapiclient.NewControllerServiceEntity() // ControllerServiceEntity | The controller service configuration details.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ControllerApi.CreateControllerService(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ControllerApi.CreateControllerService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateControllerService`: ControllerServiceEntity
    fmt.Fprintf(os.Stdout, "Response from `ControllerApi.CreateControllerService`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateControllerServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ControllerServiceEntity**](ControllerServiceEntity.md) | The controller service configuration details. | 

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

> RegistryClientEntity CreateRegistryClient(ctx).Body(body).Execute()

Creates a new registry client

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
    body := *openapiclient.NewRegistryClientEntity() // RegistryClientEntity | The registry configuration details.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ControllerApi.CreateRegistryClient(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ControllerApi.CreateRegistryClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRegistryClient`: RegistryClientEntity
    fmt.Fprintf(os.Stdout, "Response from `ControllerApi.CreateRegistryClient`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRegistryClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**RegistryClientEntity**](RegistryClientEntity.md) | The registry configuration details. | 

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

> ReportingTaskEntity CreateReportingTask(ctx).Body(body).Execute()

Creates a new reporting task

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
    body := *openapiclient.NewReportingTaskEntity() // ReportingTaskEntity | The reporting task configuration details.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ControllerApi.CreateReportingTask(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ControllerApi.CreateReportingTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateReportingTask`: ReportingTaskEntity
    fmt.Fprintf(os.Stdout, "Response from `ControllerApi.CreateReportingTask`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateReportingTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ReportingTaskEntity**](ReportingTaskEntity.md) | The reporting task configuration details. | 

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

> HistoryEntity DeleteHistory(ctx).EndDate(endDate).Execute()

Purges history

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
    endDate := "endDate_example" // string | Purge actions before this date/time.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ControllerApi.DeleteHistory(context.Background()).EndDate(endDate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ControllerApi.DeleteHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteHistory`: HistoryEntity
    fmt.Fprintf(os.Stdout, "Response from `ControllerApi.DeleteHistory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **endDate** | **string** | Purge actions before this date/time. | 

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

> NodeEntity DeleteNode(ctx, id).Execute()

Removes a node from the cluster

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
    id := "id_example" // string | The node id.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ControllerApi.DeleteNode(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ControllerApi.DeleteNode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteNode`: NodeEntity
    fmt.Fprintf(os.Stdout, "Response from `ControllerApi.DeleteNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The node id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

> RegistryClientEntity DeleteRegistryClient(ctx, id).Version(version).ClientId(clientId).DisconnectedNodeAcknowledged(disconnectedNodeAcknowledged).Execute()

Deletes a registry client

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
    id := "id_example" // string | The registry id.
    version := "version_example" // string | The revision is used to verify the client is working with the latest version of the flow. (optional)
    clientId := "clientId_example" // string | If the client id is not specified, new one will be generated. This value (whether specified or generated) is included in the response. (optional)
    disconnectedNodeAcknowledged := true // bool | Acknowledges that this node is disconnected to allow for mutable requests to proceed. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ControllerApi.DeleteRegistryClient(context.Background(), id).Version(version).ClientId(clientId).DisconnectedNodeAcknowledged(disconnectedNodeAcknowledged).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ControllerApi.DeleteRegistryClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteRegistryClient`: RegistryClientEntity
    fmt.Fprintf(os.Stdout, "Response from `ControllerApi.DeleteRegistryClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The registry id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRegistryClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **string** | The revision is used to verify the client is working with the latest version of the flow. | 
 **clientId** | **string** | If the client id is not specified, new one will be generated. This value (whether specified or generated) is included in the response. | 
 **disconnectedNodeAcknowledged** | **bool** | Acknowledges that this node is disconnected to allow for mutable requests to proceed. | [default to false]

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

> ClusterEntity GetCluster(ctx).Execute()

Gets the contents of the cluster



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ControllerApi.GetCluster(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ControllerApi.GetCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCluster`: ClusterEntity
    fmt.Fprintf(os.Stdout, "Response from `ControllerApi.GetCluster`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterRequest struct via the builder pattern


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

> ControllerConfigurationEntity GetControllerConfig(ctx).Execute()

Retrieves the configuration for this NiFi Controller

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ControllerApi.GetControllerConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ControllerApi.GetControllerConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetControllerConfig`: ControllerConfigurationEntity
    fmt.Fprintf(os.Stdout, "Response from `ControllerApi.GetControllerConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetControllerConfigRequest struct via the builder pattern


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

> NodeEntity GetNode(ctx, id).Execute()

Gets a node in the cluster

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
    id := "id_example" // string | The node id.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ControllerApi.GetNode(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ControllerApi.GetNode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNode`: NodeEntity
    fmt.Fprintf(os.Stdout, "Response from `ControllerApi.GetNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The node id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetNodeStatusHistory

> ComponentHistoryEntity GetNodeStatusHistory(ctx).Execute()

Gets status history for the node



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ControllerApi.GetNodeStatusHistory(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ControllerApi.GetNodeStatusHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNodeStatusHistory`: ComponentHistoryEntity
    fmt.Fprintf(os.Stdout, "Response from `ControllerApi.GetNodeStatusHistory`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetNodeStatusHistoryRequest struct via the builder pattern


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


## GetRegistryClient

> RegistryClientEntity GetRegistryClient(ctx, id).Execute()

Gets a registry client

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
    id := "id_example" // string | The registry id.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ControllerApi.GetRegistryClient(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ControllerApi.GetRegistryClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRegistryClient`: RegistryClientEntity
    fmt.Fprintf(os.Stdout, "Response from `ControllerApi.GetRegistryClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The registry id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRegistryClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

> RegistryClientsEntity GetRegistryClients(ctx).Execute()

Gets the listing of available registry clients

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ControllerApi.GetRegistryClients(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ControllerApi.GetRegistryClients``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRegistryClients`: RegistryClientsEntity
    fmt.Fprintf(os.Stdout, "Response from `ControllerApi.GetRegistryClients`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRegistryClientsRequest struct via the builder pattern


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

> ControllerConfigurationEntity UpdateControllerConfig(ctx).Body(body).Execute()

Retrieves the configuration for this NiFi

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
    body := *openapiclient.NewControllerConfigurationEntity() // ControllerConfigurationEntity | The controller configuration.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ControllerApi.UpdateControllerConfig(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ControllerApi.UpdateControllerConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateControllerConfig`: ControllerConfigurationEntity
    fmt.Fprintf(os.Stdout, "Response from `ControllerApi.UpdateControllerConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateControllerConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ControllerConfigurationEntity**](ControllerConfigurationEntity.md) | The controller configuration. | 

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

> NodeEntity UpdateNode(ctx, id).Body(body).Execute()

Updates a node in the cluster

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
    id := "id_example" // string | The node id.
    body := *openapiclient.NewNodeEntity() // NodeEntity | The node configuration. The only configuration that will be honored at this endpoint is the status.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ControllerApi.UpdateNode(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ControllerApi.UpdateNode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNode`: NodeEntity
    fmt.Fprintf(os.Stdout, "Response from `ControllerApi.UpdateNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The node id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**NodeEntity**](NodeEntity.md) | The node configuration. The only configuration that will be honored at this endpoint is the status. | 

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

> RegistryClientEntity UpdateRegistryClient(ctx, id).Body(body).Execute()

Updates a registry client

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
    id := "id_example" // string | The registry id.
    body := *openapiclient.NewRegistryClientEntity() // RegistryClientEntity | The registry configuration details.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ControllerApi.UpdateRegistryClient(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ControllerApi.UpdateRegistryClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRegistryClient`: RegistryClientEntity
    fmt.Fprintf(os.Stdout, "Response from `ControllerApi.UpdateRegistryClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The registry id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRegistryClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**RegistryClientEntity**](RegistryClientEntity.md) | The registry configuration details. | 

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

