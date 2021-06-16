# \ParameterContextsApi

All URIs are relative to *http://localhost/nifi-api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateParameterContext**](ParameterContextsApi.md#CreateParameterContext) | **Post** /parameter-contexts | Create a Parameter Context
[**DeleteParameterContext**](ParameterContextsApi.md#DeleteParameterContext) | **Delete** /parameter-contexts/{id} | Deletes the Parameter Context with the given ID
[**DeleteUpdateRequest**](ParameterContextsApi.md#DeleteUpdateRequest) | **Delete** /parameter-contexts/{contextId}/update-requests/{requestId} | Deletes the Update Request with the given ID
[**DeleteValidationRequest**](ParameterContextsApi.md#DeleteValidationRequest) | **Delete** /parameter-contexts/{contextId}/validation-requests/{id} | Deletes the Validation Request with the given ID
[**GetParameterContext**](ParameterContextsApi.md#GetParameterContext) | **Get** /parameter-contexts/{id} | Returns the Parameter Context with the given ID
[**GetParameterContextUpdate**](ParameterContextsApi.md#GetParameterContextUpdate) | **Get** /parameter-contexts/{contextId}/update-requests/{requestId} | Returns the Update Request with the given ID
[**GetValidationRequest**](ParameterContextsApi.md#GetValidationRequest) | **Get** /parameter-contexts/{contextId}/validation-requests/{id} | Returns the Validation Request with the given ID
[**SubmitParameterContextUpdate**](ParameterContextsApi.md#SubmitParameterContextUpdate) | **Post** /parameter-contexts/{contextId}/update-requests | Initiate the Update Request of a Parameter Context
[**SubmitValidationRequest**](ParameterContextsApi.md#SubmitValidationRequest) | **Post** /parameter-contexts/{contextId}/validation-requests | Initiate a Validation Request to determine how the validity of components will change if a Parameter Context were to be updated
[**UpdateParameterContext**](ParameterContextsApi.md#UpdateParameterContext) | **Put** /parameter-contexts/{id} | Modifies a Parameter Context



## CreateParameterContext

> ParameterContextEntity CreateParameterContext(ctx).Body(body).Execute()

Create a Parameter Context

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
    body := *openapiclient.NewParameterContextEntity() // ParameterContextEntity | The Parameter Context.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ParameterContextsApi.CreateParameterContext(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ParameterContextsApi.CreateParameterContext``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateParameterContext`: ParameterContextEntity
    fmt.Fprintf(os.Stdout, "Response from `ParameterContextsApi.CreateParameterContext`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateParameterContextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ParameterContextEntity**](ParameterContextEntity.md) | The Parameter Context. | 

### Return type

[**ParameterContextEntity**](ParameterContextEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteParameterContext

> ParameterContextEntity DeleteParameterContext(ctx, id).Version(version).ClientId(clientId).DisconnectedNodeAcknowledged(disconnectedNodeAcknowledged).Execute()

Deletes the Parameter Context with the given ID



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
    id := "id_example" // string | The Parameter Context ID.
    version := "version_example" // string | The version is used to verify the client is working with the latest version of the flow. (optional)
    clientId := "clientId_example" // string | If the client id is not specified, a new one will be generated. This value (whether specified or generated) is included in the response. (optional)
    disconnectedNodeAcknowledged := true // bool | Acknowledges that this node is disconnected to allow for mutable requests to proceed. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ParameterContextsApi.DeleteParameterContext(context.Background(), id).Version(version).ClientId(clientId).DisconnectedNodeAcknowledged(disconnectedNodeAcknowledged).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ParameterContextsApi.DeleteParameterContext``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteParameterContext`: ParameterContextEntity
    fmt.Fprintf(os.Stdout, "Response from `ParameterContextsApi.DeleteParameterContext`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Parameter Context ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteParameterContextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **string** | The version is used to verify the client is working with the latest version of the flow. | 
 **clientId** | **string** | If the client id is not specified, a new one will be generated. This value (whether specified or generated) is included in the response. | 
 **disconnectedNodeAcknowledged** | **bool** | Acknowledges that this node is disconnected to allow for mutable requests to proceed. | [default to false]

### Return type

[**ParameterContextEntity**](ParameterContextEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUpdateRequest

> ParameterContextUpdateRequestEntity DeleteUpdateRequest(ctx, contextId, requestId).DisconnectedNodeAcknowledged(disconnectedNodeAcknowledged).Execute()

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
    contextId := "contextId_example" // string | The ID of the ParameterContext
    requestId := "requestId_example" // string | The ID of the Update Request
    disconnectedNodeAcknowledged := true // bool | Acknowledges that this node is disconnected to allow for mutable requests to proceed. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ParameterContextsApi.DeleteUpdateRequest(context.Background(), contextId, requestId).DisconnectedNodeAcknowledged(disconnectedNodeAcknowledged).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ParameterContextsApi.DeleteUpdateRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteUpdateRequest`: ParameterContextUpdateRequestEntity
    fmt.Fprintf(os.Stdout, "Response from `ParameterContextsApi.DeleteUpdateRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contextId** | **string** | The ID of the ParameterContext | 
**requestId** | **string** | The ID of the Update Request | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUpdateRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **disconnectedNodeAcknowledged** | **bool** | Acknowledges that this node is disconnected to allow for mutable requests to proceed. | [default to false]

### Return type

[**ParameterContextUpdateRequestEntity**](ParameterContextUpdateRequestEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteValidationRequest

> ParameterContextValidationRequestEntity DeleteValidationRequest(ctx, contextId, id).DisconnectedNodeAcknowledged(disconnectedNodeAcknowledged).Execute()

Deletes the Validation Request with the given ID



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
    contextId := "contextId_example" // string | The ID of the Parameter Context
    id := "id_example" // string | The ID of the Update Request
    disconnectedNodeAcknowledged := true // bool | Acknowledges that this node is disconnected to allow for mutable requests to proceed. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ParameterContextsApi.DeleteValidationRequest(context.Background(), contextId, id).DisconnectedNodeAcknowledged(disconnectedNodeAcknowledged).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ParameterContextsApi.DeleteValidationRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteValidationRequest`: ParameterContextValidationRequestEntity
    fmt.Fprintf(os.Stdout, "Response from `ParameterContextsApi.DeleteValidationRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contextId** | **string** | The ID of the Parameter Context | 
**id** | **string** | The ID of the Update Request | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteValidationRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **disconnectedNodeAcknowledged** | **bool** | Acknowledges that this node is disconnected to allow for mutable requests to proceed. | [default to false]

### Return type

[**ParameterContextValidationRequestEntity**](ParameterContextValidationRequestEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetParameterContext

> ParameterContextEntity GetParameterContext(ctx, id).Execute()

Returns the Parameter Context with the given ID



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
    id := "id_example" // string | The ID of the Parameter Context

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ParameterContextsApi.GetParameterContext(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ParameterContextsApi.GetParameterContext``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetParameterContext`: ParameterContextEntity
    fmt.Fprintf(os.Stdout, "Response from `ParameterContextsApi.GetParameterContext`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the Parameter Context | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetParameterContextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ParameterContextEntity**](ParameterContextEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetParameterContextUpdate

> ParameterContextUpdateRequestEntity GetParameterContextUpdate(ctx, contextId, requestId).Execute()

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
    contextId := "contextId_example" // string | The ID of the Parameter Context
    requestId := "requestId_example" // string | The ID of the Update Request

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ParameterContextsApi.GetParameterContextUpdate(context.Background(), contextId, requestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ParameterContextsApi.GetParameterContextUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetParameterContextUpdate`: ParameterContextUpdateRequestEntity
    fmt.Fprintf(os.Stdout, "Response from `ParameterContextsApi.GetParameterContextUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contextId** | **string** | The ID of the Parameter Context | 
**requestId** | **string** | The ID of the Update Request | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetParameterContextUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ParameterContextUpdateRequestEntity**](ParameterContextUpdateRequestEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetValidationRequest

> ParameterContextValidationRequestEntity GetValidationRequest(ctx, contextId, id).Execute()

Returns the Validation Request with the given ID



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
    contextId := "contextId_example" // string | The ID of the Parameter Context
    id := "id_example" // string | The ID of the Validation Request

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ParameterContextsApi.GetValidationRequest(context.Background(), contextId, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ParameterContextsApi.GetValidationRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetValidationRequest`: ParameterContextValidationRequestEntity
    fmt.Fprintf(os.Stdout, "Response from `ParameterContextsApi.GetValidationRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contextId** | **string** | The ID of the Parameter Context | 
**id** | **string** | The ID of the Validation Request | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetValidationRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ParameterContextValidationRequestEntity**](ParameterContextValidationRequestEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitParameterContextUpdate

> ParameterContextUpdateRequestEntity SubmitParameterContextUpdate(ctx, contextId).Body(body).Execute()

Initiate the Update Request of a Parameter Context



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
    contextId := "contextId_example" // string | 
    body := *openapiclient.NewParameterContextEntity() // ParameterContextEntity | The updated version of the parameter context.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ParameterContextsApi.SubmitParameterContextUpdate(context.Background(), contextId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ParameterContextsApi.SubmitParameterContextUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubmitParameterContextUpdate`: ParameterContextUpdateRequestEntity
    fmt.Fprintf(os.Stdout, "Response from `ParameterContextsApi.SubmitParameterContextUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contextId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubmitParameterContextUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ParameterContextEntity**](ParameterContextEntity.md) | The updated version of the parameter context. | 

### Return type

[**ParameterContextUpdateRequestEntity**](ParameterContextUpdateRequestEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitValidationRequest

> ParameterContextValidationRequestEntity SubmitValidationRequest(ctx, contextId).Body(body).Execute()

Initiate a Validation Request to determine how the validity of components will change if a Parameter Context were to be updated



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
    contextId := "contextId_example" // string | 
    body := *openapiclient.NewParameterContextValidationRequestEntity() // ParameterContextValidationRequestEntity | The validation request

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ParameterContextsApi.SubmitValidationRequest(context.Background(), contextId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ParameterContextsApi.SubmitValidationRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubmitValidationRequest`: ParameterContextValidationRequestEntity
    fmt.Fprintf(os.Stdout, "Response from `ParameterContextsApi.SubmitValidationRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contextId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubmitValidationRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ParameterContextValidationRequestEntity**](ParameterContextValidationRequestEntity.md) | The validation request | 

### Return type

[**ParameterContextValidationRequestEntity**](ParameterContextValidationRequestEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateParameterContext

> ParameterContextEntity UpdateParameterContext(ctx, id).Body(body).Execute()

Modifies a Parameter Context



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
    id := "id_example" // string | 
    body := *openapiclient.NewParameterContextEntity() // ParameterContextEntity | The updated Parameter Context

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ParameterContextsApi.UpdateParameterContext(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ParameterContextsApi.UpdateParameterContext``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateParameterContext`: ParameterContextEntity
    fmt.Fprintf(os.Stdout, "Response from `ParameterContextsApi.UpdateParameterContext`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateParameterContextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ParameterContextEntity**](ParameterContextEntity.md) | The updated Parameter Context | 

### Return type

[**ParameterContextEntity**](ParameterContextEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

