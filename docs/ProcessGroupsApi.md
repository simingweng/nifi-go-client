# \ProcessGroupsApi

All URIs are relative to *http://localhost/nifi-api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CopySnippet**](ProcessGroupsApi.md#CopySnippet) | **Post** /process-groups/{id}/snippet-instance | Copies a snippet and discards it.
[**CreateConnection**](ProcessGroupsApi.md#CreateConnection) | **Post** /process-groups/{id}/connections | Creates a connection
[**CreateControllerService**](ProcessGroupsApi.md#CreateControllerService) | **Post** /process-groups/{id}/controller-services | Creates a new controller service
[**CreateEmptyAllConnectionsRequest**](ProcessGroupsApi.md#CreateEmptyAllConnectionsRequest) | **Post** /process-groups/{id}/empty-all-connections-requests | Creates a request to drop all flowfiles of all connection queues in this process group.
[**CreateFunnel**](ProcessGroupsApi.md#CreateFunnel) | **Post** /process-groups/{id}/funnels | Creates a funnel
[**CreateInputPort**](ProcessGroupsApi.md#CreateInputPort) | **Post** /process-groups/{id}/input-ports | Creates an input port
[**CreateLabel**](ProcessGroupsApi.md#CreateLabel) | **Post** /process-groups/{id}/labels | Creates a label
[**CreateOutputPort**](ProcessGroupsApi.md#CreateOutputPort) | **Post** /process-groups/{id}/output-ports | Creates an output port
[**CreateProcessGroup**](ProcessGroupsApi.md#CreateProcessGroup) | **Post** /process-groups/{id}/process-groups | Creates a process group
[**CreateProcessor**](ProcessGroupsApi.md#CreateProcessor) | **Post** /process-groups/{id}/processors | Creates a new processor
[**CreateRemoteProcessGroup**](ProcessGroupsApi.md#CreateRemoteProcessGroup) | **Post** /process-groups/{id}/remote-process-groups | Creates a new process group
[**CreateTemplate**](ProcessGroupsApi.md#CreateTemplate) | **Post** /process-groups/{id}/templates | Creates a template and discards the specified snippet.
[**DeleteReplaceProcessGroupRequest**](ProcessGroupsApi.md#DeleteReplaceProcessGroupRequest) | **Delete** /process-groups/replace-requests/{id} | Deletes the Replace Request with the given ID
[**DeleteVariableRegistryUpdateRequest**](ProcessGroupsApi.md#DeleteVariableRegistryUpdateRequest) | **Delete** /process-groups/{groupId}/variable-registry/update-requests/{updateId} | Deletes an update request for a process group&#39;s variable registry. If the request is not yet complete, it will automatically be cancelled.
[**ExportProcessGroup**](ProcessGroupsApi.md#ExportProcessGroup) | **Get** /process-groups/{id}/download | Gets a process group for download
[**GetConnections**](ProcessGroupsApi.md#GetConnections) | **Get** /process-groups/{id}/connections | Gets all connections
[**GetDropAllFlowfilesRequest**](ProcessGroupsApi.md#GetDropAllFlowfilesRequest) | **Get** /process-groups/{id}/empty-all-connections-requests/{drop-request-id} | Gets the current status of a drop all flowfiles request.
[**GetFunnels**](ProcessGroupsApi.md#GetFunnels) | **Get** /process-groups/{id}/funnels | Gets all funnels
[**GetInputPorts**](ProcessGroupsApi.md#GetInputPorts) | **Get** /process-groups/{id}/input-ports | Gets all input ports
[**GetLabels**](ProcessGroupsApi.md#GetLabels) | **Get** /process-groups/{id}/labels | Gets all labels
[**GetLocalModifications**](ProcessGroupsApi.md#GetLocalModifications) | **Get** /process-groups/{id}/local-modifications | Gets a list of local modifications to the Process Group since it was last synchronized with the Flow Registry
[**GetOutputPorts**](ProcessGroupsApi.md#GetOutputPorts) | **Get** /process-groups/{id}/output-ports | Gets all output ports
[**GetProcessGroup**](ProcessGroupsApi.md#GetProcessGroup) | **Get** /process-groups/{id} | Gets a process group
[**GetProcessGroups**](ProcessGroupsApi.md#GetProcessGroups) | **Get** /process-groups/{id}/process-groups | Gets all process groups
[**GetProcessors**](ProcessGroupsApi.md#GetProcessors) | **Get** /process-groups/{id}/processors | Gets all processors
[**GetRemoteProcessGroups**](ProcessGroupsApi.md#GetRemoteProcessGroups) | **Get** /process-groups/{id}/remote-process-groups | Gets all remote process groups
[**GetReplaceProcessGroupRequest**](ProcessGroupsApi.md#GetReplaceProcessGroupRequest) | **Get** /process-groups/replace-requests/{id} | Returns the Replace Request with the given ID
[**GetVariableRegistry**](ProcessGroupsApi.md#GetVariableRegistry) | **Get** /process-groups/{id}/variable-registry | Gets a process group&#39;s variable registry
[**GetVariableRegistryUpdateRequest**](ProcessGroupsApi.md#GetVariableRegistryUpdateRequest) | **Get** /process-groups/{groupId}/variable-registry/update-requests/{updateId} | Gets a process group&#39;s variable registry
[**ImportTemplate**](ProcessGroupsApi.md#ImportTemplate) | **Post** /process-groups/{id}/templates/import | Imports a template
[**InitiateReplaceProcessGroup**](ProcessGroupsApi.md#InitiateReplaceProcessGroup) | **Post** /process-groups/{id}/replace-requests | Initiate the Replace Request of a Process Group with the given ID
[**InstantiateTemplate**](ProcessGroupsApi.md#InstantiateTemplate) | **Post** /process-groups/{id}/template-instance | Instantiates a template
[**RemoveDropRequest**](ProcessGroupsApi.md#RemoveDropRequest) | **Delete** /process-groups/{id}/empty-all-connections-requests/{drop-request-id} | Cancels and/or removes a request to drop all flowfiles.
[**RemoveProcessGroup**](ProcessGroupsApi.md#RemoveProcessGroup) | **Delete** /process-groups/{id} | Deletes a process group
[**ReplaceProcessGroup**](ProcessGroupsApi.md#ReplaceProcessGroup) | **Put** /process-groups/{id}/flow-contents | Replace Process Group contents with the given ID with the specified Process Group contents
[**SubmitUpdateVariableRegistryRequest**](ProcessGroupsApi.md#SubmitUpdateVariableRegistryRequest) | **Post** /process-groups/{id}/variable-registry/update-requests | Submits a request to update a process group&#39;s variable registry
[**UpdateProcessGroup**](ProcessGroupsApi.md#UpdateProcessGroup) | **Put** /process-groups/{id} | Updates a process group
[**UpdateVariableRegistry**](ProcessGroupsApi.md#UpdateVariableRegistry) | **Put** /process-groups/{id}/variable-registry | Updates the contents of a Process Group&#39;s variable Registry
[**UploadTemplate**](ProcessGroupsApi.md#UploadTemplate) | **Post** /process-groups/{id}/templates/upload | Uploads a template



## CopySnippet

> FlowEntity CopySnippet(ctx, id).Body(body).Execute()

Copies a snippet and discards it.

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
    body := *openapiclient.NewCopySnippetRequestEntity() // CopySnippetRequestEntity | The copy snippet request.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProcessGroupsApi.CopySnippet(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProcessGroupsApi.CopySnippet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CopySnippet`: FlowEntity
    fmt.Fprintf(os.Stdout, "Response from `ProcessGroupsApi.CopySnippet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The process group id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCopySnippetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CopySnippetRequestEntity**](CopySnippetRequestEntity.md) | The copy snippet request. | 

### Return type

[**FlowEntity**](FlowEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateConnection

> ConnectionEntity CreateConnection(ctx, id).Body(body).Execute()

Creates a connection

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
    body := *openapiclient.NewConnectionEntity("SourceType_example", "DestinationType_example") // ConnectionEntity | The connection configuration details.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProcessGroupsApi.CreateConnection(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProcessGroupsApi.CreateConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateConnection`: ConnectionEntity
    fmt.Fprintf(os.Stdout, "Response from `ProcessGroupsApi.CreateConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The process group id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ConnectionEntity**](ConnectionEntity.md) | The connection configuration details. | 

### Return type

[**ConnectionEntity**](ConnectionEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateControllerService

> ControllerServiceEntity CreateControllerService(ctx, id).Body(body).Execute()

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
    id := "id_example" // string | The process group id.
    body := *openapiclient.NewControllerServiceEntity() // ControllerServiceEntity | The controller service configuration details.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProcessGroupsApi.CreateControllerService(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProcessGroupsApi.CreateControllerService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateControllerService`: ControllerServiceEntity
    fmt.Fprintf(os.Stdout, "Response from `ProcessGroupsApi.CreateControllerService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The process group id. | 

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


## CreateEmptyAllConnectionsRequest

> ProcessGroupEntity CreateEmptyAllConnectionsRequest(ctx, id).Execute()

Creates a request to drop all flowfiles of all connection queues in this process group.

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
    resp, r, err := api_client.ProcessGroupsApi.CreateEmptyAllConnectionsRequest(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProcessGroupsApi.CreateEmptyAllConnectionsRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateEmptyAllConnectionsRequest`: ProcessGroupEntity
    fmt.Fprintf(os.Stdout, "Response from `ProcessGroupsApi.CreateEmptyAllConnectionsRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The process group id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateEmptyAllConnectionsRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProcessGroupEntity**](ProcessGroupEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateFunnel

> FunnelEntity CreateFunnel(ctx, id).Body(body).Execute()

Creates a funnel

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
    body := *openapiclient.NewFunnelEntity() // FunnelEntity | The funnel configuration details.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProcessGroupsApi.CreateFunnel(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProcessGroupsApi.CreateFunnel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateFunnel`: FunnelEntity
    fmt.Fprintf(os.Stdout, "Response from `ProcessGroupsApi.CreateFunnel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The process group id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateFunnelRequest struct via the builder pattern


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


## CreateInputPort

> PortEntity CreateInputPort(ctx, id).Body(body).Execute()

Creates an input port

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
    body := *openapiclient.NewPortEntity() // PortEntity | The input port configuration details.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProcessGroupsApi.CreateInputPort(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProcessGroupsApi.CreateInputPort``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateInputPort`: PortEntity
    fmt.Fprintf(os.Stdout, "Response from `ProcessGroupsApi.CreateInputPort`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The process group id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateInputPortRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**PortEntity**](PortEntity.md) | The input port configuration details. | 

### Return type

[**PortEntity**](PortEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateLabel

> LabelEntity CreateLabel(ctx, id).Body(body).Execute()

Creates a label

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
    body := *openapiclient.NewLabelEntity() // LabelEntity | The label configuration details.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProcessGroupsApi.CreateLabel(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProcessGroupsApi.CreateLabel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLabel`: LabelEntity
    fmt.Fprintf(os.Stdout, "Response from `ProcessGroupsApi.CreateLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The process group id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLabelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**LabelEntity**](LabelEntity.md) | The label configuration details. | 

### Return type

[**LabelEntity**](LabelEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOutputPort

> PortEntity CreateOutputPort(ctx, id).Body(body).Execute()

Creates an output port

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
    body := *openapiclient.NewPortEntity() // PortEntity | The output port configuration.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProcessGroupsApi.CreateOutputPort(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProcessGroupsApi.CreateOutputPort``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOutputPort`: PortEntity
    fmt.Fprintf(os.Stdout, "Response from `ProcessGroupsApi.CreateOutputPort`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The process group id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOutputPortRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**PortEntity**](PortEntity.md) | The output port configuration. | 

### Return type

[**PortEntity**](PortEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateProcessGroup

> ProcessGroupEntity CreateProcessGroup(ctx, id).Body(body).Execute()

Creates a process group

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
    body := *openapiclient.NewProcessGroupEntity() // ProcessGroupEntity | The process group configuration details.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProcessGroupsApi.CreateProcessGroup(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProcessGroupsApi.CreateProcessGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateProcessGroup`: ProcessGroupEntity
    fmt.Fprintf(os.Stdout, "Response from `ProcessGroupsApi.CreateProcessGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The process group id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateProcessGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ProcessGroupEntity**](ProcessGroupEntity.md) | The process group configuration details. | 

### Return type

[**ProcessGroupEntity**](ProcessGroupEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateProcessor

> ProcessorEntity CreateProcessor(ctx, id).Body(body).Execute()

Creates a new processor

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
    body := *openapiclient.NewProcessorEntity() // ProcessorEntity | The processor configuration details.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProcessGroupsApi.CreateProcessor(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProcessGroupsApi.CreateProcessor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateProcessor`: ProcessorEntity
    fmt.Fprintf(os.Stdout, "Response from `ProcessGroupsApi.CreateProcessor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The process group id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateProcessorRequest struct via the builder pattern


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


## CreateRemoteProcessGroup

> RemoteProcessGroupEntity CreateRemoteProcessGroup(ctx, id).Body(body).Execute()

Creates a new process group

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
    body := *openapiclient.NewRemoteProcessGroupEntity() // RemoteProcessGroupEntity | The remote process group configuration details.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProcessGroupsApi.CreateRemoteProcessGroup(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProcessGroupsApi.CreateRemoteProcessGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRemoteProcessGroup`: RemoteProcessGroupEntity
    fmt.Fprintf(os.Stdout, "Response from `ProcessGroupsApi.CreateRemoteProcessGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The process group id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRemoteProcessGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**RemoteProcessGroupEntity**](RemoteProcessGroupEntity.md) | The remote process group configuration details. | 

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


## CreateTemplate

> TemplateEntity CreateTemplate(ctx, id).Body(body).Execute()

Creates a template and discards the specified snippet.

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
    body := *openapiclient.NewCreateTemplateRequestEntity() // CreateTemplateRequestEntity | The create template request.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProcessGroupsApi.CreateTemplate(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProcessGroupsApi.CreateTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTemplate`: TemplateEntity
    fmt.Fprintf(os.Stdout, "Response from `ProcessGroupsApi.CreateTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The process group id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CreateTemplateRequestEntity**](CreateTemplateRequestEntity.md) | The create template request. | 

### Return type

[**TemplateEntity**](TemplateEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteReplaceProcessGroupRequest

> ProcessGroupReplaceRequestEntity DeleteReplaceProcessGroupRequest(ctx, id).DisconnectedNodeAcknowledged(disconnectedNodeAcknowledged).Execute()

Deletes the Replace Request with the given ID



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
    resp, r, err := api_client.ProcessGroupsApi.DeleteReplaceProcessGroupRequest(context.Background(), id).DisconnectedNodeAcknowledged(disconnectedNodeAcknowledged).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProcessGroupsApi.DeleteReplaceProcessGroupRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteReplaceProcessGroupRequest`: ProcessGroupReplaceRequestEntity
    fmt.Fprintf(os.Stdout, "Response from `ProcessGroupsApi.DeleteReplaceProcessGroupRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the Update Request | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteReplaceProcessGroupRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **disconnectedNodeAcknowledged** | **bool** | Acknowledges that this node is disconnected to allow for mutable requests to proceed. | [default to false]

### Return type

[**ProcessGroupReplaceRequestEntity**](ProcessGroupReplaceRequestEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVariableRegistryUpdateRequest

> VariableRegistryUpdateRequestEntity DeleteVariableRegistryUpdateRequest(ctx, groupId, updateId).DisconnectedNodeAcknowledged(disconnectedNodeAcknowledged).Execute()

Deletes an update request for a process group's variable registry. If the request is not yet complete, it will automatically be cancelled.



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
    groupId := "groupId_example" // string | The process group id.
    updateId := "updateId_example" // string | The ID of the Variable Registry Update Request
    disconnectedNodeAcknowledged := true // bool | Acknowledges that this node is disconnected to allow for mutable requests to proceed. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProcessGroupsApi.DeleteVariableRegistryUpdateRequest(context.Background(), groupId, updateId).DisconnectedNodeAcknowledged(disconnectedNodeAcknowledged).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProcessGroupsApi.DeleteVariableRegistryUpdateRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteVariableRegistryUpdateRequest`: VariableRegistryUpdateRequestEntity
    fmt.Fprintf(os.Stdout, "Response from `ProcessGroupsApi.DeleteVariableRegistryUpdateRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The process group id. | 
**updateId** | **string** | The ID of the Variable Registry Update Request | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVariableRegistryUpdateRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **disconnectedNodeAcknowledged** | **bool** | Acknowledges that this node is disconnected to allow for mutable requests to proceed. | [default to false]

### Return type

[**VariableRegistryUpdateRequestEntity**](VariableRegistryUpdateRequestEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportProcessGroup

> string ExportProcessGroup(ctx, id).Execute()

Gets a process group for download

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
    resp, r, err := api_client.ProcessGroupsApi.ExportProcessGroup(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProcessGroupsApi.ExportProcessGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportProcessGroup`: string
    fmt.Fprintf(os.Stdout, "Response from `ProcessGroupsApi.ExportProcessGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The process group id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportProcessGroupRequest struct via the builder pattern


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


## GetConnections

> ConnectionsEntity GetConnections(ctx, id).Execute()

Gets all connections

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
    resp, r, err := api_client.ProcessGroupsApi.GetConnections(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProcessGroupsApi.GetConnections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConnections`: ConnectionsEntity
    fmt.Fprintf(os.Stdout, "Response from `ProcessGroupsApi.GetConnections`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The process group id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConnectionsEntity**](ConnectionsEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDropAllFlowfilesRequest

> DropRequestEntity GetDropAllFlowfilesRequest(ctx, id, dropRequestId).Execute()

Gets the current status of a drop all flowfiles request.

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
    dropRequestId := "dropRequestId_example" // string | The drop request id.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProcessGroupsApi.GetDropAllFlowfilesRequest(context.Background(), id, dropRequestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProcessGroupsApi.GetDropAllFlowfilesRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDropAllFlowfilesRequest`: DropRequestEntity
    fmt.Fprintf(os.Stdout, "Response from `ProcessGroupsApi.GetDropAllFlowfilesRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The process group id. | 
**dropRequestId** | **string** | The drop request id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDropAllFlowfilesRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DropRequestEntity**](DropRequestEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFunnels

> FunnelsEntity GetFunnels(ctx, id).Execute()

Gets all funnels

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
    resp, r, err := api_client.ProcessGroupsApi.GetFunnels(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProcessGroupsApi.GetFunnels``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFunnels`: FunnelsEntity
    fmt.Fprintf(os.Stdout, "Response from `ProcessGroupsApi.GetFunnels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The process group id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFunnelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FunnelsEntity**](FunnelsEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInputPorts

> InputPortsEntity GetInputPorts(ctx, id).Execute()

Gets all input ports

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
    resp, r, err := api_client.ProcessGroupsApi.GetInputPorts(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProcessGroupsApi.GetInputPorts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInputPorts`: InputPortsEntity
    fmt.Fprintf(os.Stdout, "Response from `ProcessGroupsApi.GetInputPorts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The process group id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInputPortsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InputPortsEntity**](InputPortsEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLabels

> LabelsEntity GetLabels(ctx, id).Execute()

Gets all labels

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
    resp, r, err := api_client.ProcessGroupsApi.GetLabels(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProcessGroupsApi.GetLabels``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLabels`: LabelsEntity
    fmt.Fprintf(os.Stdout, "Response from `ProcessGroupsApi.GetLabels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The process group id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLabelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LabelsEntity**](LabelsEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLocalModifications

> FlowComparisonEntity GetLocalModifications(ctx, id).Execute()

Gets a list of local modifications to the Process Group since it was last synchronized with the Flow Registry

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
    resp, r, err := api_client.ProcessGroupsApi.GetLocalModifications(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProcessGroupsApi.GetLocalModifications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLocalModifications`: FlowComparisonEntity
    fmt.Fprintf(os.Stdout, "Response from `ProcessGroupsApi.GetLocalModifications`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The process group id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLocalModificationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FlowComparisonEntity**](FlowComparisonEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOutputPorts

> OutputPortsEntity GetOutputPorts(ctx, id).Execute()

Gets all output ports

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
    resp, r, err := api_client.ProcessGroupsApi.GetOutputPorts(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProcessGroupsApi.GetOutputPorts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOutputPorts`: OutputPortsEntity
    fmt.Fprintf(os.Stdout, "Response from `ProcessGroupsApi.GetOutputPorts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The process group id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOutputPortsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OutputPortsEntity**](OutputPortsEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProcessGroup

> ProcessGroupEntity GetProcessGroup(ctx, id).Execute()

Gets a process group

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
    resp, r, err := api_client.ProcessGroupsApi.GetProcessGroup(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProcessGroupsApi.GetProcessGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProcessGroup`: ProcessGroupEntity
    fmt.Fprintf(os.Stdout, "Response from `ProcessGroupsApi.GetProcessGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The process group id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProcessGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProcessGroupEntity**](ProcessGroupEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProcessGroups

> ProcessGroupsEntity GetProcessGroups(ctx, id).Execute()

Gets all process groups

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
    resp, r, err := api_client.ProcessGroupsApi.GetProcessGroups(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProcessGroupsApi.GetProcessGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProcessGroups`: ProcessGroupsEntity
    fmt.Fprintf(os.Stdout, "Response from `ProcessGroupsApi.GetProcessGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The process group id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProcessGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProcessGroupsEntity**](ProcessGroupsEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProcessors

> ProcessorsEntity GetProcessors(ctx, id).IncludeDescendantGroups(includeDescendantGroups).Execute()

Gets all processors

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
    includeDescendantGroups := true // bool | Whether or not to include processors from descendant process groups (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProcessGroupsApi.GetProcessors(context.Background(), id).IncludeDescendantGroups(includeDescendantGroups).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProcessGroupsApi.GetProcessors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProcessors`: ProcessorsEntity
    fmt.Fprintf(os.Stdout, "Response from `ProcessGroupsApi.GetProcessors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The process group id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProcessorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeDescendantGroups** | **bool** | Whether or not to include processors from descendant process groups | [default to false]

### Return type

[**ProcessorsEntity**](ProcessorsEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRemoteProcessGroups

> RemoteProcessGroupsEntity GetRemoteProcessGroups(ctx, id).Execute()

Gets all remote process groups

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
    resp, r, err := api_client.ProcessGroupsApi.GetRemoteProcessGroups(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProcessGroupsApi.GetRemoteProcessGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRemoteProcessGroups`: RemoteProcessGroupsEntity
    fmt.Fprintf(os.Stdout, "Response from `ProcessGroupsApi.GetRemoteProcessGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The process group id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRemoteProcessGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RemoteProcessGroupsEntity**](RemoteProcessGroupsEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReplaceProcessGroupRequest

> ProcessGroupReplaceRequestEntity GetReplaceProcessGroupRequest(ctx, id).Execute()

Returns the Replace Request with the given ID



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
    id := "id_example" // string | The ID of the Replace Request

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProcessGroupsApi.GetReplaceProcessGroupRequest(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProcessGroupsApi.GetReplaceProcessGroupRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReplaceProcessGroupRequest`: ProcessGroupReplaceRequestEntity
    fmt.Fprintf(os.Stdout, "Response from `ProcessGroupsApi.GetReplaceProcessGroupRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the Replace Request | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReplaceProcessGroupRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProcessGroupReplaceRequestEntity**](ProcessGroupReplaceRequestEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVariableRegistry

> VariableRegistryEntity GetVariableRegistry(ctx, id).IncludeAncestorGroups(includeAncestorGroups).Execute()

Gets a process group's variable registry



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
    includeAncestorGroups := true // bool | Whether or not to include ancestor groups (optional) (default to true)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProcessGroupsApi.GetVariableRegistry(context.Background(), id).IncludeAncestorGroups(includeAncestorGroups).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProcessGroupsApi.GetVariableRegistry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVariableRegistry`: VariableRegistryEntity
    fmt.Fprintf(os.Stdout, "Response from `ProcessGroupsApi.GetVariableRegistry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The process group id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVariableRegistryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeAncestorGroups** | **bool** | Whether or not to include ancestor groups | [default to true]

### Return type

[**VariableRegistryEntity**](VariableRegistryEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVariableRegistryUpdateRequest

> VariableRegistryUpdateRequestEntity GetVariableRegistryUpdateRequest(ctx, groupId, updateId).Execute()

Gets a process group's variable registry



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
    groupId := "groupId_example" // string | The process group id.
    updateId := "updateId_example" // string | The ID of the Variable Registry Update Request

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProcessGroupsApi.GetVariableRegistryUpdateRequest(context.Background(), groupId, updateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProcessGroupsApi.GetVariableRegistryUpdateRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVariableRegistryUpdateRequest`: VariableRegistryUpdateRequestEntity
    fmt.Fprintf(os.Stdout, "Response from `ProcessGroupsApi.GetVariableRegistryUpdateRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The process group id. | 
**updateId** | **string** | The ID of the Variable Registry Update Request | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVariableRegistryUpdateRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**VariableRegistryUpdateRequestEntity**](VariableRegistryUpdateRequestEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportTemplate

> TemplateEntity ImportTemplate(ctx, id).Execute()

Imports a template

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
    resp, r, err := api_client.ProcessGroupsApi.ImportTemplate(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProcessGroupsApi.ImportTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportTemplate`: TemplateEntity
    fmt.Fprintf(os.Stdout, "Response from `ProcessGroupsApi.ImportTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The process group id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TemplateEntity**](TemplateEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InitiateReplaceProcessGroup

> ProcessGroupReplaceRequestEntity InitiateReplaceProcessGroup(ctx, id).Body(body).Execute()

Initiate the Replace Request of a Process Group with the given ID



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
    body := *openapiclient.NewProcessGroupImportEntity() // ProcessGroupImportEntity | The process group replace request entity

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProcessGroupsApi.InitiateReplaceProcessGroup(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProcessGroupsApi.InitiateReplaceProcessGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InitiateReplaceProcessGroup`: ProcessGroupReplaceRequestEntity
    fmt.Fprintf(os.Stdout, "Response from `ProcessGroupsApi.InitiateReplaceProcessGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The process group id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInitiateReplaceProcessGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ProcessGroupImportEntity**](ProcessGroupImportEntity.md) | The process group replace request entity | 

### Return type

[**ProcessGroupReplaceRequestEntity**](ProcessGroupReplaceRequestEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InstantiateTemplate

> FlowEntity InstantiateTemplate(ctx, id).Body(body).Execute()

Instantiates a template

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
    body := *openapiclient.NewInstantiateTemplateRequestEntity() // InstantiateTemplateRequestEntity | The instantiate template request.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProcessGroupsApi.InstantiateTemplate(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProcessGroupsApi.InstantiateTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InstantiateTemplate`: FlowEntity
    fmt.Fprintf(os.Stdout, "Response from `ProcessGroupsApi.InstantiateTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The process group id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInstantiateTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**InstantiateTemplateRequestEntity**](InstantiateTemplateRequestEntity.md) | The instantiate template request. | 

### Return type

[**FlowEntity**](FlowEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveDropRequest

> DropRequestEntity RemoveDropRequest(ctx, id, dropRequestId).Execute()

Cancels and/or removes a request to drop all flowfiles.

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
    dropRequestId := "dropRequestId_example" // string | The drop request id.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProcessGroupsApi.RemoveDropRequest(context.Background(), id, dropRequestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProcessGroupsApi.RemoveDropRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveDropRequest`: DropRequestEntity
    fmt.Fprintf(os.Stdout, "Response from `ProcessGroupsApi.RemoveDropRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The process group id. | 
**dropRequestId** | **string** | The drop request id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveDropRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DropRequestEntity**](DropRequestEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveProcessGroup

> ProcessGroupEntity RemoveProcessGroup(ctx, id).Version(version).ClientId(clientId).DisconnectedNodeAcknowledged(disconnectedNodeAcknowledged).Execute()

Deletes a process group

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
    version := "version_example" // string | The revision is used to verify the client is working with the latest version of the flow. (optional)
    clientId := "clientId_example" // string | If the client id is not specified, new one will be generated. This value (whether specified or generated) is included in the response. (optional)
    disconnectedNodeAcknowledged := true // bool | Acknowledges that this node is disconnected to allow for mutable requests to proceed. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProcessGroupsApi.RemoveProcessGroup(context.Background(), id).Version(version).ClientId(clientId).DisconnectedNodeAcknowledged(disconnectedNodeAcknowledged).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProcessGroupsApi.RemoveProcessGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveProcessGroup`: ProcessGroupEntity
    fmt.Fprintf(os.Stdout, "Response from `ProcessGroupsApi.RemoveProcessGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The process group id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveProcessGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **string** | The revision is used to verify the client is working with the latest version of the flow. | 
 **clientId** | **string** | If the client id is not specified, new one will be generated. This value (whether specified or generated) is included in the response. | 
 **disconnectedNodeAcknowledged** | **bool** | Acknowledges that this node is disconnected to allow for mutable requests to proceed. | [default to false]

### Return type

[**ProcessGroupEntity**](ProcessGroupEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceProcessGroup

> ProcessGroupImportEntity ReplaceProcessGroup(ctx, id).Body(body).Execute()

Replace Process Group contents with the given ID with the specified Process Group contents



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
    body := *openapiclient.NewProcessGroupImportEntity() // ProcessGroupImportEntity | The process group replace request entity.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProcessGroupsApi.ReplaceProcessGroup(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProcessGroupsApi.ReplaceProcessGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceProcessGroup`: ProcessGroupImportEntity
    fmt.Fprintf(os.Stdout, "Response from `ProcessGroupsApi.ReplaceProcessGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The process group id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceProcessGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ProcessGroupImportEntity**](ProcessGroupImportEntity.md) | The process group replace request entity. | 

### Return type

[**ProcessGroupImportEntity**](ProcessGroupImportEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitUpdateVariableRegistryRequest

> VariableRegistryUpdateRequestEntity SubmitUpdateVariableRegistryRequest(ctx, id).Body(body).Execute()

Submits a request to update a process group's variable registry



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
    body := *openapiclient.NewVariableRegistryEntity() // VariableRegistryEntity | The variable registry configuration details.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProcessGroupsApi.SubmitUpdateVariableRegistryRequest(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProcessGroupsApi.SubmitUpdateVariableRegistryRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubmitUpdateVariableRegistryRequest`: VariableRegistryUpdateRequestEntity
    fmt.Fprintf(os.Stdout, "Response from `ProcessGroupsApi.SubmitUpdateVariableRegistryRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The process group id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubmitUpdateVariableRegistryRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**VariableRegistryEntity**](VariableRegistryEntity.md) | The variable registry configuration details. | 

### Return type

[**VariableRegistryUpdateRequestEntity**](VariableRegistryUpdateRequestEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProcessGroup

> ProcessGroupEntity UpdateProcessGroup(ctx, id).Body(body).Execute()

Updates a process group

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
    body := *openapiclient.NewProcessGroupEntity() // ProcessGroupEntity | The process group configuration details.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProcessGroupsApi.UpdateProcessGroup(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProcessGroupsApi.UpdateProcessGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateProcessGroup`: ProcessGroupEntity
    fmt.Fprintf(os.Stdout, "Response from `ProcessGroupsApi.UpdateProcessGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The process group id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProcessGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ProcessGroupEntity**](ProcessGroupEntity.md) | The process group configuration details. | 

### Return type

[**ProcessGroupEntity**](ProcessGroupEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateVariableRegistry

> VariableRegistryEntity UpdateVariableRegistry(ctx, id).Body(body).Execute()

Updates the contents of a Process Group's variable Registry



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
    body := *openapiclient.NewVariableRegistryEntity() // VariableRegistryEntity | The variable registry configuration details.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProcessGroupsApi.UpdateVariableRegistry(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProcessGroupsApi.UpdateVariableRegistry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateVariableRegistry`: VariableRegistryEntity
    fmt.Fprintf(os.Stdout, "Response from `ProcessGroupsApi.UpdateVariableRegistry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The process group id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVariableRegistryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**VariableRegistryEntity**](VariableRegistryEntity.md) | The variable registry configuration details. | 

### Return type

[**VariableRegistryEntity**](VariableRegistryEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadTemplate

> TemplateEntity UploadTemplate(ctx, id).Template(template).Execute()

Uploads a template

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
    template := os.NewFile(1234, "some_file") // *os.File | The binary content of the template file being uploaded.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProcessGroupsApi.UploadTemplate(context.Background(), id).Template(template).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProcessGroupsApi.UploadTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadTemplate`: TemplateEntity
    fmt.Fprintf(os.Stdout, "Response from `ProcessGroupsApi.UploadTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The process group id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **template** | ***os.File** | The binary content of the template file being uploaded. | 

### Return type

[**TemplateEntity**](TemplateEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

