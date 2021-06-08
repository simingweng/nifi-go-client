# \ProcessorsApi

All URIs are relative to *http://localhost/nifi-api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClearState**](ProcessorsApi.md#ClearState) | **Post** /processors/{id}/state/clear-requests | Clears the state for a processor
[**DeleteProcessor**](ProcessorsApi.md#DeleteProcessor) | **Delete** /processors/{id} | Deletes a processor
[**GetProcessor**](ProcessorsApi.md#GetProcessor) | **Get** /processors/{id} | Gets a processor
[**GetProcessorDiagnostics**](ProcessorsApi.md#GetProcessorDiagnostics) | **Get** /processors/{id}/diagnostics | Gets diagnostics information about a processor
[**GetProcessorRunStatusDetails**](ProcessorsApi.md#GetProcessorRunStatusDetails) | **Post** /processors/run-status-details/queries | Submits a query to retrieve the run status details of all processors that are in the given list of Processor IDs
[**GetPropertyDescriptor**](ProcessorsApi.md#GetPropertyDescriptor) | **Get** /processors/{id}/descriptors | Gets the descriptor for a processor property
[**GetState**](ProcessorsApi.md#GetState) | **Get** /processors/{id}/state | Gets the state for a processor
[**TerminateProcessor**](ProcessorsApi.md#TerminateProcessor) | **Delete** /processors/{id}/threads | Terminates a processor, essentially \&quot;deleting\&quot; its threads and any active tasks
[**UpdateProcessor**](ProcessorsApi.md#UpdateProcessor) | **Put** /processors/{id} | Updates a processor
[**UpdateRunStatus**](ProcessorsApi.md#UpdateRunStatus) | **Put** /processors/{id}/run-status | Updates run status of a processor



## ClearState

> ComponentStateEntity ClearState(ctx, id).Execute()

Clears the state for a processor

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
    resp, r, err := api_client.ProcessorsApi.ClearState(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProcessorsApi.ClearState``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ClearState`: ComponentStateEntity
    fmt.Fprintf(os.Stdout, "Response from `ProcessorsApi.ClearState`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The processor id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiClearStateRequest struct via the builder pattern


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


## DeleteProcessor

> ProcessorEntity DeleteProcessor(ctx, id).Version(version).ClientId(clientId).DisconnectedNodeAcknowledged(disconnectedNodeAcknowledged).Execute()

Deletes a processor

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
    version := "version_example" // string | The revision is used to verify the client is working with the latest version of the flow. (optional)
    clientId := "clientId_example" // string | If the client id is not specified, new one will be generated. This value (whether specified or generated) is included in the response. (optional)
    disconnectedNodeAcknowledged := true // bool | Acknowledges that this node is disconnected to allow for mutable requests to proceed. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProcessorsApi.DeleteProcessor(context.Background(), id).Version(version).ClientId(clientId).DisconnectedNodeAcknowledged(disconnectedNodeAcknowledged).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProcessorsApi.DeleteProcessor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteProcessor`: ProcessorEntity
    fmt.Fprintf(os.Stdout, "Response from `ProcessorsApi.DeleteProcessor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The processor id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProcessorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **string** | The revision is used to verify the client is working with the latest version of the flow. | 
 **clientId** | **string** | If the client id is not specified, new one will be generated. This value (whether specified or generated) is included in the response. | 
 **disconnectedNodeAcknowledged** | **bool** | Acknowledges that this node is disconnected to allow for mutable requests to proceed. | [default to false]

### Return type

[**ProcessorEntity**](ProcessorEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProcessor

> ProcessorEntity GetProcessor(ctx, id).Execute()

Gets a processor

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
    resp, r, err := api_client.ProcessorsApi.GetProcessor(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProcessorsApi.GetProcessor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProcessor`: ProcessorEntity
    fmt.Fprintf(os.Stdout, "Response from `ProcessorsApi.GetProcessor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The processor id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProcessorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProcessorEntity**](ProcessorEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProcessorDiagnostics

> ProcessorEntity GetProcessorDiagnostics(ctx, id).Execute()

Gets diagnostics information about a processor



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
    resp, r, err := api_client.ProcessorsApi.GetProcessorDiagnostics(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProcessorsApi.GetProcessorDiagnostics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProcessorDiagnostics`: ProcessorEntity
    fmt.Fprintf(os.Stdout, "Response from `ProcessorsApi.GetProcessorDiagnostics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The processor id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProcessorDiagnosticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProcessorEntity**](ProcessorEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProcessorRunStatusDetails

> ProcessorsRunStatusDetailsEntity GetProcessorRunStatusDetails(ctx).Body(body).Execute()

Submits a query to retrieve the run status details of all processors that are in the given list of Processor IDs

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
    body := *openapiclient.NewRunStatusDetailsRequestEntity() // RunStatusDetailsRequestEntity | The request for the processors that should be included in the results (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProcessorsApi.GetProcessorRunStatusDetails(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProcessorsApi.GetProcessorRunStatusDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProcessorRunStatusDetails`: ProcessorsRunStatusDetailsEntity
    fmt.Fprintf(os.Stdout, "Response from `ProcessorsApi.GetProcessorRunStatusDetails`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProcessorRunStatusDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**RunStatusDetailsRequestEntity**](RunStatusDetailsRequestEntity.md) | The request for the processors that should be included in the results | 

### Return type

[**ProcessorsRunStatusDetailsEntity**](ProcessorsRunStatusDetailsEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPropertyDescriptor

> PropertyDescriptorEntity GetPropertyDescriptor(ctx, id).PropertyName(propertyName).ClientId(clientId).Execute()

Gets the descriptor for a processor property

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
    propertyName := "propertyName_example" // string | The property name.
    clientId := "clientId_example" // string | If the client id is not specified, new one will be generated. This value (whether specified or generated) is included in the response. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProcessorsApi.GetPropertyDescriptor(context.Background(), id).PropertyName(propertyName).ClientId(clientId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProcessorsApi.GetPropertyDescriptor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPropertyDescriptor`: PropertyDescriptorEntity
    fmt.Fprintf(os.Stdout, "Response from `ProcessorsApi.GetPropertyDescriptor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The processor id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPropertyDescriptorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **propertyName** | **string** | The property name. | 
 **clientId** | **string** | If the client id is not specified, new one will be generated. This value (whether specified or generated) is included in the response. | 

### Return type

[**PropertyDescriptorEntity**](PropertyDescriptorEntity.md)

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

Gets the state for a processor

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
    resp, r, err := api_client.ProcessorsApi.GetState(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProcessorsApi.GetState``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetState`: ComponentStateEntity
    fmt.Fprintf(os.Stdout, "Response from `ProcessorsApi.GetState`: %v\n", resp)
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


## TerminateProcessor

> ProcessorEntity TerminateProcessor(ctx, id).Execute()

Terminates a processor, essentially \"deleting\" its threads and any active tasks

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
    resp, r, err := api_client.ProcessorsApi.TerminateProcessor(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProcessorsApi.TerminateProcessor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TerminateProcessor`: ProcessorEntity
    fmt.Fprintf(os.Stdout, "Response from `ProcessorsApi.TerminateProcessor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The processor id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTerminateProcessorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProcessorEntity**](ProcessorEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProcessor

> ProcessorEntity UpdateProcessor(ctx, id).Body(body).Execute()

Updates a processor

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
    body := *openapiclient.NewProcessorEntity() // ProcessorEntity | The processor configuration details.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProcessorsApi.UpdateProcessor(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProcessorsApi.UpdateProcessor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateProcessor`: ProcessorEntity
    fmt.Fprintf(os.Stdout, "Response from `ProcessorsApi.UpdateProcessor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The processor id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProcessorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ProcessorEntity**](ProcessorEntity.md) | The processor configuration details. | 

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


## UpdateRunStatus

> ProcessorEntity UpdateRunStatus(ctx, id).Body(body).Execute()

Updates run status of a processor

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
    body := *openapiclient.NewProcessorRunStatusEntity() // ProcessorRunStatusEntity | The processor run status.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProcessorsApi.UpdateRunStatus(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProcessorsApi.UpdateRunStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRunStatus`: ProcessorEntity
    fmt.Fprintf(os.Stdout, "Response from `ProcessorsApi.UpdateRunStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The processor id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRunStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ProcessorRunStatusEntity**](ProcessorRunStatusEntity.md) | The processor run status. | 

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

