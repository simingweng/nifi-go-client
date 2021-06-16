# \VersionsApi

All URIs are relative to *http://localhost/nifi-api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVersionControlRequest**](VersionsApi.md#CreateVersionControlRequest) | **Post** /versions/active-requests | Create a version control request
[**DeleteRevertRequest**](VersionsApi.md#DeleteRevertRequest) | **Delete** /versions/revert-requests/{id} | Deletes the Revert Request with the given ID
[**DeleteUpdateRequest**](VersionsApi.md#DeleteUpdateRequest) | **Delete** /versions/update-requests/{id} | Deletes the Update Request with the given ID
[**DeleteVersionControlRequest**](VersionsApi.md#DeleteVersionControlRequest) | **Delete** /versions/active-requests/{id} | Deletes the version control request with the given ID
[**ExportFlowVersion**](VersionsApi.md#ExportFlowVersion) | **Get** /versions/process-groups/{id}/download | Gets the latest version of a Process Group for download
[**GetRevertRequest**](VersionsApi.md#GetRevertRequest) | **Get** /versions/revert-requests/{id} | Returns the Revert Request with the given ID
[**GetUpdateRequest**](VersionsApi.md#GetUpdateRequest) | **Get** /versions/update-requests/{id} | Returns the Update Request with the given ID
[**GetVersionInformation**](VersionsApi.md#GetVersionInformation) | **Get** /versions/process-groups/{id} | Gets the Version Control information for a process group
[**InitiateRevertFlowVersion**](VersionsApi.md#InitiateRevertFlowVersion) | **Post** /versions/revert-requests/process-groups/{id} | Initiate the Revert Request of a Process Group with the given ID
[**InitiateVersionControlUpdate**](VersionsApi.md#InitiateVersionControlUpdate) | **Post** /versions/update-requests/process-groups/{id} | Initiate the Update Request of a Process Group with the given ID
[**SaveToFlowRegistry**](VersionsApi.md#SaveToFlowRegistry) | **Post** /versions/process-groups/{id} | Save the Process Group with the given ID
[**StopVersionControl**](VersionsApi.md#StopVersionControl) | **Delete** /versions/process-groups/{id} | Stops version controlling the Process Group with the given ID
[**UpdateFlowVersion**](VersionsApi.md#UpdateFlowVersion) | **Put** /versions/process-groups/{id} | Update the version of a Process Group with the given ID
[**UpdateVersionControlRequest**](VersionsApi.md#UpdateVersionControlRequest) | **Put** /versions/active-requests/{id} | Updates the request with the given ID



## CreateVersionControlRequest

> string CreateVersionControlRequest(ctx).Body(body).Execute()

Create a version control request



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
    body := *openapiclient.NewCreateActiveRequestEntity() // CreateActiveRequestEntity | The versioned flow details.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VersionsApi.CreateVersionControlRequest(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VersionsApi.CreateVersionControlRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVersionControlRequest`: string
    fmt.Fprintf(os.Stdout, "Response from `VersionsApi.CreateVersionControlRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateVersionControlRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateActiveRequestEntity**](CreateActiveRequestEntity.md) | The versioned flow details. | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRevertRequest

> VersionedFlowUpdateRequestEntity DeleteRevertRequest(ctx, id).DisconnectedNodeAcknowledged(disconnectedNodeAcknowledged).Execute()

Deletes the Revert Request with the given ID



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
    id := "id_example" // string | The ID of the Revert Request
    disconnectedNodeAcknowledged := true // bool | Acknowledges that this node is disconnected to allow for mutable requests to proceed. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VersionsApi.DeleteRevertRequest(context.Background(), id).DisconnectedNodeAcknowledged(disconnectedNodeAcknowledged).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VersionsApi.DeleteRevertRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteRevertRequest`: VersionedFlowUpdateRequestEntity
    fmt.Fprintf(os.Stdout, "Response from `VersionsApi.DeleteRevertRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the Revert Request | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRevertRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **disconnectedNodeAcknowledged** | **bool** | Acknowledges that this node is disconnected to allow for mutable requests to proceed. | [default to false]

### Return type

[**VersionedFlowUpdateRequestEntity**](VersionedFlowUpdateRequestEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUpdateRequest

> VersionedFlowUpdateRequestEntity DeleteUpdateRequest(ctx, id).DisconnectedNodeAcknowledged(disconnectedNodeAcknowledged).Execute()

Deletes the Update Request with the given ID



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
    id := "id_example" // string | The ID of the Update Request
    disconnectedNodeAcknowledged := true // bool | Acknowledges that this node is disconnected to allow for mutable requests to proceed. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VersionsApi.DeleteUpdateRequest(context.Background(), id).DisconnectedNodeAcknowledged(disconnectedNodeAcknowledged).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VersionsApi.DeleteUpdateRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteUpdateRequest`: VersionedFlowUpdateRequestEntity
    fmt.Fprintf(os.Stdout, "Response from `VersionsApi.DeleteUpdateRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the Update Request | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUpdateRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **disconnectedNodeAcknowledged** | **bool** | Acknowledges that this node is disconnected to allow for mutable requests to proceed. | [default to false]

### Return type

[**VersionedFlowUpdateRequestEntity**](VersionedFlowUpdateRequestEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVersionControlRequest

> DeleteVersionControlRequest(ctx, id).DisconnectedNodeAcknowledged(disconnectedNodeAcknowledged).Execute()

Deletes the version control request with the given ID



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
    id := "id_example" // string | The request ID.
    disconnectedNodeAcknowledged := true // bool | Acknowledges that this node is disconnected to allow for mutable requests to proceed. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VersionsApi.DeleteVersionControlRequest(context.Background(), id).DisconnectedNodeAcknowledged(disconnectedNodeAcknowledged).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VersionsApi.DeleteVersionControlRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The request ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVersionControlRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **disconnectedNodeAcknowledged** | **bool** | Acknowledges that this node is disconnected to allow for mutable requests to proceed. | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportFlowVersion

> string ExportFlowVersion(ctx, id).Execute()

Gets the latest version of a Process Group for download

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
    id := "id_example" // string | The process group id.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VersionsApi.ExportFlowVersion(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VersionsApi.ExportFlowVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportFlowVersion`: string
    fmt.Fprintf(os.Stdout, "Response from `VersionsApi.ExportFlowVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The process group id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportFlowVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRevertRequest

> VersionedFlowUpdateRequestEntity GetRevertRequest(ctx, id).Execute()

Returns the Revert Request with the given ID



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
    id := "id_example" // string | The ID of the Revert Request

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VersionsApi.GetRevertRequest(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VersionsApi.GetRevertRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRevertRequest`: VersionedFlowUpdateRequestEntity
    fmt.Fprintf(os.Stdout, "Response from `VersionsApi.GetRevertRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the Revert Request | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRevertRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VersionedFlowUpdateRequestEntity**](VersionedFlowUpdateRequestEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUpdateRequest

> VersionedFlowUpdateRequestEntity GetUpdateRequest(ctx, id).Execute()

Returns the Update Request with the given ID



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
    id := "id_example" // string | The ID of the Update Request

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VersionsApi.GetUpdateRequest(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VersionsApi.GetUpdateRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUpdateRequest`: VersionedFlowUpdateRequestEntity
    fmt.Fprintf(os.Stdout, "Response from `VersionsApi.GetUpdateRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the Update Request | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUpdateRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VersionedFlowUpdateRequestEntity**](VersionedFlowUpdateRequestEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVersionInformation

> VersionControlInformationEntity GetVersionInformation(ctx, id).Execute()

Gets the Version Control information for a process group



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
    id := "id_example" // string | The process group id.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VersionsApi.GetVersionInformation(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VersionsApi.GetVersionInformation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVersionInformation`: VersionControlInformationEntity
    fmt.Fprintf(os.Stdout, "Response from `VersionsApi.GetVersionInformation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The process group id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVersionInformationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VersionControlInformationEntity**](VersionControlInformationEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InitiateRevertFlowVersion

> VersionedFlowUpdateRequestEntity InitiateRevertFlowVersion(ctx, id).Body(body).Execute()

Initiate the Revert Request of a Process Group with the given ID



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
    id := "id_example" // string | The process group id.
    body := *openapiclient.NewVersionControlInformationEntity() // VersionControlInformationEntity | The controller service configuration details.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VersionsApi.InitiateRevertFlowVersion(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VersionsApi.InitiateRevertFlowVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InitiateRevertFlowVersion`: VersionedFlowUpdateRequestEntity
    fmt.Fprintf(os.Stdout, "Response from `VersionsApi.InitiateRevertFlowVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The process group id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInitiateRevertFlowVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**VersionControlInformationEntity**](VersionControlInformationEntity.md) | The controller service configuration details. | 

### Return type

[**VersionedFlowUpdateRequestEntity**](VersionedFlowUpdateRequestEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InitiateVersionControlUpdate

> VersionedFlowUpdateRequestEntity InitiateVersionControlUpdate(ctx, id).Body(body).Execute()

Initiate the Update Request of a Process Group with the given ID



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
    id := "id_example" // string | The process group id.
    body := *openapiclient.NewVersionControlInformationEntity() // VersionControlInformationEntity | The controller service configuration details.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VersionsApi.InitiateVersionControlUpdate(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VersionsApi.InitiateVersionControlUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InitiateVersionControlUpdate`: VersionedFlowUpdateRequestEntity
    fmt.Fprintf(os.Stdout, "Response from `VersionsApi.InitiateVersionControlUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The process group id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInitiateVersionControlUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**VersionControlInformationEntity**](VersionControlInformationEntity.md) | The controller service configuration details. | 

### Return type

[**VersionedFlowUpdateRequestEntity**](VersionedFlowUpdateRequestEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SaveToFlowRegistry

> VersionControlInformationEntity SaveToFlowRegistry(ctx, id).Body(body).Execute()

Save the Process Group with the given ID



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
    id := "id_example" // string | The process group id.
    body := *openapiclient.NewStartVersionControlRequestEntity() // StartVersionControlRequestEntity | The versioned flow details.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VersionsApi.SaveToFlowRegistry(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VersionsApi.SaveToFlowRegistry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SaveToFlowRegistry`: VersionControlInformationEntity
    fmt.Fprintf(os.Stdout, "Response from `VersionsApi.SaveToFlowRegistry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The process group id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSaveToFlowRegistryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**StartVersionControlRequestEntity**](StartVersionControlRequestEntity.md) | The versioned flow details. | 

### Return type

[**VersionControlInformationEntity**](VersionControlInformationEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StopVersionControl

> VersionControlInformationEntity StopVersionControl(ctx, id).Version(version).ClientId(clientId).DisconnectedNodeAcknowledged(disconnectedNodeAcknowledged).Execute()

Stops version controlling the Process Group with the given ID



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
    id := "id_example" // string | The process group id.
    version := "version_example" // string | The version is used to verify the client is working with the latest version of the flow. (optional)
    clientId := "clientId_example" // string | If the client id is not specified, a new one will be generated. This value (whether specified or generated) is included in the response. (optional)
    disconnectedNodeAcknowledged := true // bool | Acknowledges that this node is disconnected to allow for mutable requests to proceed. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VersionsApi.StopVersionControl(context.Background(), id).Version(version).ClientId(clientId).DisconnectedNodeAcknowledged(disconnectedNodeAcknowledged).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VersionsApi.StopVersionControl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StopVersionControl`: VersionControlInformationEntity
    fmt.Fprintf(os.Stdout, "Response from `VersionsApi.StopVersionControl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The process group id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopVersionControlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **string** | The version is used to verify the client is working with the latest version of the flow. | 
 **clientId** | **string** | If the client id is not specified, a new one will be generated. This value (whether specified or generated) is included in the response. | 
 **disconnectedNodeAcknowledged** | **bool** | Acknowledges that this node is disconnected to allow for mutable requests to proceed. | [default to false]

### Return type

[**VersionControlInformationEntity**](VersionControlInformationEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFlowVersion

> VersionControlInformationEntity UpdateFlowVersion(ctx, id).Body(body).Execute()

Update the version of a Process Group with the given ID



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
    id := "id_example" // string | The process group id.
    body := *openapiclient.NewVersionedFlowSnapshotEntity() // VersionedFlowSnapshotEntity | The controller service configuration details.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VersionsApi.UpdateFlowVersion(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VersionsApi.UpdateFlowVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateFlowVersion`: VersionControlInformationEntity
    fmt.Fprintf(os.Stdout, "Response from `VersionsApi.UpdateFlowVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The process group id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFlowVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**VersionedFlowSnapshotEntity**](VersionedFlowSnapshotEntity.md) | The controller service configuration details. | 

### Return type

[**VersionControlInformationEntity**](VersionControlInformationEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateVersionControlRequest

> VersionControlInformationEntity UpdateVersionControlRequest(ctx, id).Body(body).Execute()

Updates the request with the given ID



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
    id := "id_example" // string | The request ID.
    body := *openapiclient.NewVersionControlComponentMappingEntity() // VersionControlComponentMappingEntity | The version control component mapping.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VersionsApi.UpdateVersionControlRequest(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VersionsApi.UpdateVersionControlRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateVersionControlRequest`: VersionControlInformationEntity
    fmt.Fprintf(os.Stdout, "Response from `VersionsApi.UpdateVersionControlRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The request ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVersionControlRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**VersionControlComponentMappingEntity**](VersionControlComponentMappingEntity.md) | The version control component mapping. | 

### Return type

[**VersionControlInformationEntity**](VersionControlInformationEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

