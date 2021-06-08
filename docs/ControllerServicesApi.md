# \ControllerServicesApi

All URIs are relative to *http://localhost/nifi-api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClearState**](ControllerServicesApi.md#ClearState) | **Post** /controller-services/{id}/state/clear-requests | Clears the state for a controller service
[**GetControllerService**](ControllerServicesApi.md#GetControllerService) | **Get** /controller-services/{id} | Gets a controller service
[**GetControllerServiceReferences**](ControllerServicesApi.md#GetControllerServiceReferences) | **Get** /controller-services/{id}/references | Gets a controller service
[**GetPropertyDescriptor**](ControllerServicesApi.md#GetPropertyDescriptor) | **Get** /controller-services/{id}/descriptors | Gets a controller service property descriptor
[**GetState**](ControllerServicesApi.md#GetState) | **Get** /controller-services/{id}/state | Gets the state for a controller service
[**RemoveControllerService**](ControllerServicesApi.md#RemoveControllerService) | **Delete** /controller-services/{id} | Deletes a controller service
[**UpdateControllerService**](ControllerServicesApi.md#UpdateControllerService) | **Put** /controller-services/{id} | Updates a controller service
[**UpdateControllerServiceReferences**](ControllerServicesApi.md#UpdateControllerServiceReferences) | **Put** /controller-services/{id}/references | Updates a controller services references
[**UpdateRunStatus**](ControllerServicesApi.md#UpdateRunStatus) | **Put** /controller-services/{id}/run-status | Updates run status of a controller service



## ClearState

> ComponentStateEntity ClearState(ctx, id).Execute()

Clears the state for a controller service

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
    id := "id_example" // string | The controller service id.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ControllerServicesApi.ClearState(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ControllerServicesApi.ClearState``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ClearState`: ComponentStateEntity
    fmt.Fprintf(os.Stdout, "Response from `ControllerServicesApi.ClearState`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The controller service id. | 

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


## GetControllerService

> ControllerServiceEntity GetControllerService(ctx, id).Execute()

Gets a controller service

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
    id := "id_example" // string | The controller service id.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ControllerServicesApi.GetControllerService(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ControllerServicesApi.GetControllerService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetControllerService`: ControllerServiceEntity
    fmt.Fprintf(os.Stdout, "Response from `ControllerServicesApi.GetControllerService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The controller service id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetControllerServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ControllerServiceEntity**](ControllerServiceEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetControllerServiceReferences

> ControllerServiceReferencingComponentsEntity GetControllerServiceReferences(ctx, id).Execute()

Gets a controller service

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
    id := "id_example" // string | The controller service id.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ControllerServicesApi.GetControllerServiceReferences(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ControllerServicesApi.GetControllerServiceReferences``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetControllerServiceReferences`: ControllerServiceReferencingComponentsEntity
    fmt.Fprintf(os.Stdout, "Response from `ControllerServicesApi.GetControllerServiceReferences`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The controller service id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetControllerServiceReferencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ControllerServiceReferencingComponentsEntity**](ControllerServiceReferencingComponentsEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPropertyDescriptor

> PropertyDescriptorEntity GetPropertyDescriptor(ctx, id).PropertyName(propertyName).Execute()

Gets a controller service property descriptor

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
    id := "id_example" // string | The controller service id.
    propertyName := "propertyName_example" // string | The property name to return the descriptor for.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ControllerServicesApi.GetPropertyDescriptor(context.Background(), id).PropertyName(propertyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ControllerServicesApi.GetPropertyDescriptor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPropertyDescriptor`: PropertyDescriptorEntity
    fmt.Fprintf(os.Stdout, "Response from `ControllerServicesApi.GetPropertyDescriptor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The controller service id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPropertyDescriptorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **propertyName** | **string** | The property name to return the descriptor for. | 

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

Gets the state for a controller service

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
    id := "id_example" // string | The controller service id.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ControllerServicesApi.GetState(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ControllerServicesApi.GetState``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetState`: ComponentStateEntity
    fmt.Fprintf(os.Stdout, "Response from `ControllerServicesApi.GetState`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The controller service id. | 

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


## RemoveControllerService

> ControllerServiceEntity RemoveControllerService(ctx, id).Version(version).ClientId(clientId).DisconnectedNodeAcknowledged(disconnectedNodeAcknowledged).Execute()

Deletes a controller service

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
    id := "id_example" // string | The controller service id.
    version := "version_example" // string | The revision is used to verify the client is working with the latest version of the flow. (optional)
    clientId := "clientId_example" // string | If the client id is not specified, new one will be generated. This value (whether specified or generated) is included in the response. (optional)
    disconnectedNodeAcknowledged := true // bool | Acknowledges that this node is disconnected to allow for mutable requests to proceed. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ControllerServicesApi.RemoveControllerService(context.Background(), id).Version(version).ClientId(clientId).DisconnectedNodeAcknowledged(disconnectedNodeAcknowledged).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ControllerServicesApi.RemoveControllerService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveControllerService`: ControllerServiceEntity
    fmt.Fprintf(os.Stdout, "Response from `ControllerServicesApi.RemoveControllerService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The controller service id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveControllerServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **string** | The revision is used to verify the client is working with the latest version of the flow. | 
 **clientId** | **string** | If the client id is not specified, new one will be generated. This value (whether specified or generated) is included in the response. | 
 **disconnectedNodeAcknowledged** | **bool** | Acknowledges that this node is disconnected to allow for mutable requests to proceed. | [default to false]

### Return type

[**ControllerServiceEntity**](ControllerServiceEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateControllerService

> ControllerServiceEntity UpdateControllerService(ctx, id).Body(body).Execute()

Updates a controller service

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
    id := "id_example" // string | The controller service id.
    body := *openapiclient.NewControllerServiceEntity() // ControllerServiceEntity | The controller service configuration details.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ControllerServicesApi.UpdateControllerService(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ControllerServicesApi.UpdateControllerService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateControllerService`: ControllerServiceEntity
    fmt.Fprintf(os.Stdout, "Response from `ControllerServicesApi.UpdateControllerService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The controller service id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateControllerServiceRequest struct via the builder pattern


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


## UpdateControllerServiceReferences

> ControllerServiceReferencingComponentsEntity UpdateControllerServiceReferences(ctx, id).Body(body).Execute()

Updates a controller services references

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
    id := "id_example" // string | The controller service id.
    body := *openapiclient.NewUpdateControllerServiceReferenceRequestEntity() // UpdateControllerServiceReferenceRequestEntity | The controller service request update request.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ControllerServicesApi.UpdateControllerServiceReferences(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ControllerServicesApi.UpdateControllerServiceReferences``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateControllerServiceReferences`: ControllerServiceReferencingComponentsEntity
    fmt.Fprintf(os.Stdout, "Response from `ControllerServicesApi.UpdateControllerServiceReferences`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The controller service id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateControllerServiceReferencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**UpdateControllerServiceReferenceRequestEntity**](UpdateControllerServiceReferenceRequestEntity.md) | The controller service request update request. | 

### Return type

[**ControllerServiceReferencingComponentsEntity**](ControllerServiceReferencingComponentsEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRunStatus

> ControllerServiceEntity UpdateRunStatus(ctx, id).Body(body).Execute()

Updates run status of a controller service

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
    id := "id_example" // string | The controller service id.
    body := *openapiclient.NewControllerServiceRunStatusEntity() // ControllerServiceRunStatusEntity | The controller service run status.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ControllerServicesApi.UpdateRunStatus(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ControllerServicesApi.UpdateRunStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRunStatus`: ControllerServiceEntity
    fmt.Fprintf(os.Stdout, "Response from `ControllerServicesApi.UpdateRunStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The controller service id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRunStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ControllerServiceRunStatusEntity**](ControllerServiceRunStatusEntity.md) | The controller service run status. | 

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

