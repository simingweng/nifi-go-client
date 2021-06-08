# \FlowfileQueuesApi

All URIs are relative to *http://localhost/nifi-api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDropRequest**](FlowfileQueuesApi.md#CreateDropRequest) | **Post** /flowfile-queues/{id}/drop-requests | Creates a request to drop the contents of the queue in this connection.
[**CreateFlowFileListing**](FlowfileQueuesApi.md#CreateFlowFileListing) | **Post** /flowfile-queues/{id}/listing-requests | Lists the contents of the queue in this connection.
[**DeleteListingRequest**](FlowfileQueuesApi.md#DeleteListingRequest) | **Delete** /flowfile-queues/{id}/listing-requests/{listing-request-id} | Cancels and/or removes a request to list the contents of this connection.
[**DownloadFlowFileContent**](FlowfileQueuesApi.md#DownloadFlowFileContent) | **Get** /flowfile-queues/{id}/flowfiles/{flowfile-uuid}/content | Gets the content for a FlowFile in a Connection.
[**GetDropRequest**](FlowfileQueuesApi.md#GetDropRequest) | **Get** /flowfile-queues/{id}/drop-requests/{drop-request-id} | Gets the current status of a drop request for the specified connection.
[**GetFlowFile**](FlowfileQueuesApi.md#GetFlowFile) | **Get** /flowfile-queues/{id}/flowfiles/{flowfile-uuid} | Gets a FlowFile from a Connection.
[**GetListingRequest**](FlowfileQueuesApi.md#GetListingRequest) | **Get** /flowfile-queues/{id}/listing-requests/{listing-request-id} | Gets the current status of a listing request for the specified connection.
[**RemoveDropRequest**](FlowfileQueuesApi.md#RemoveDropRequest) | **Delete** /flowfile-queues/{id}/drop-requests/{drop-request-id} | Cancels and/or removes a request to drop the contents of this connection.



## CreateDropRequest

> DropRequestEntity CreateDropRequest(ctx, id).Execute()

Creates a request to drop the contents of the queue in this connection.

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
    id := "id_example" // string | The connection id.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FlowfileQueuesApi.CreateDropRequest(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlowfileQueuesApi.CreateDropRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDropRequest`: DropRequestEntity
    fmt.Fprintf(os.Stdout, "Response from `FlowfileQueuesApi.CreateDropRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The connection id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDropRequestRequest struct via the builder pattern


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


## CreateFlowFileListing

> ListingRequestEntity CreateFlowFileListing(ctx, id).Execute()

Lists the contents of the queue in this connection.

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
    id := "id_example" // string | The connection id.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FlowfileQueuesApi.CreateFlowFileListing(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlowfileQueuesApi.CreateFlowFileListing``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateFlowFileListing`: ListingRequestEntity
    fmt.Fprintf(os.Stdout, "Response from `FlowfileQueuesApi.CreateFlowFileListing`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The connection id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateFlowFileListingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListingRequestEntity**](ListingRequestEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteListingRequest

> ListingRequestEntity DeleteListingRequest(ctx, id, listingRequestId).Execute()

Cancels and/or removes a request to list the contents of this connection.

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
    id := "id_example" // string | The connection id.
    listingRequestId := "listingRequestId_example" // string | The listing request id.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FlowfileQueuesApi.DeleteListingRequest(context.Background(), id, listingRequestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlowfileQueuesApi.DeleteListingRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteListingRequest`: ListingRequestEntity
    fmt.Fprintf(os.Stdout, "Response from `FlowfileQueuesApi.DeleteListingRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The connection id. | 
**listingRequestId** | **string** | The listing request id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteListingRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ListingRequestEntity**](ListingRequestEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadFlowFileContent

> map[string]interface{} DownloadFlowFileContent(ctx, id, flowfileUuid).ClientId(clientId).ClusterNodeId(clusterNodeId).Execute()

Gets the content for a FlowFile in a Connection.

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
    id := "id_example" // string | The connection id.
    flowfileUuid := "flowfileUuid_example" // string | The flowfile uuid.
    clientId := "clientId_example" // string | If the client id is not specified, new one will be generated. This value (whether specified or generated) is included in the response. (optional)
    clusterNodeId := "clusterNodeId_example" // string | The id of the node where the content exists if clustered. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FlowfileQueuesApi.DownloadFlowFileContent(context.Background(), id, flowfileUuid).ClientId(clientId).ClusterNodeId(clusterNodeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlowfileQueuesApi.DownloadFlowFileContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadFlowFileContent`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `FlowfileQueuesApi.DownloadFlowFileContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The connection id. | 
**flowfileUuid** | **string** | The flowfile uuid. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadFlowFileContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** | If the client id is not specified, new one will be generated. This value (whether specified or generated) is included in the response. | 
 **clusterNodeId** | **string** | The id of the node where the content exists if clustered. | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDropRequest

> DropRequestEntity GetDropRequest(ctx, id, dropRequestId).Execute()

Gets the current status of a drop request for the specified connection.

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
    id := "id_example" // string | The connection id.
    dropRequestId := "dropRequestId_example" // string | The drop request id.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FlowfileQueuesApi.GetDropRequest(context.Background(), id, dropRequestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlowfileQueuesApi.GetDropRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDropRequest`: DropRequestEntity
    fmt.Fprintf(os.Stdout, "Response from `FlowfileQueuesApi.GetDropRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The connection id. | 
**dropRequestId** | **string** | The drop request id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDropRequestRequest struct via the builder pattern


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


## GetFlowFile

> FlowFileEntity GetFlowFile(ctx, id, flowfileUuid).ClusterNodeId(clusterNodeId).Execute()

Gets a FlowFile from a Connection.

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
    id := "id_example" // string | The connection id.
    flowfileUuid := "flowfileUuid_example" // string | The flowfile uuid.
    clusterNodeId := "clusterNodeId_example" // string | The id of the node where the content exists if clustered. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FlowfileQueuesApi.GetFlowFile(context.Background(), id, flowfileUuid).ClusterNodeId(clusterNodeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlowfileQueuesApi.GetFlowFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFlowFile`: FlowFileEntity
    fmt.Fprintf(os.Stdout, "Response from `FlowfileQueuesApi.GetFlowFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The connection id. | 
**flowfileUuid** | **string** | The flowfile uuid. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFlowFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clusterNodeId** | **string** | The id of the node where the content exists if clustered. | 

### Return type

[**FlowFileEntity**](FlowFileEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetListingRequest

> ListingRequestEntity GetListingRequest(ctx, id, listingRequestId).Execute()

Gets the current status of a listing request for the specified connection.

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
    id := "id_example" // string | The connection id.
    listingRequestId := "listingRequestId_example" // string | The listing request id.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FlowfileQueuesApi.GetListingRequest(context.Background(), id, listingRequestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlowfileQueuesApi.GetListingRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetListingRequest`: ListingRequestEntity
    fmt.Fprintf(os.Stdout, "Response from `FlowfileQueuesApi.GetListingRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The connection id. | 
**listingRequestId** | **string** | The listing request id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetListingRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ListingRequestEntity**](ListingRequestEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveDropRequest

> DropRequestEntity RemoveDropRequest(ctx, id, dropRequestId).Execute()

Cancels and/or removes a request to drop the contents of this connection.

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
    id := "id_example" // string | The connection id.
    dropRequestId := "dropRequestId_example" // string | The drop request id.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FlowfileQueuesApi.RemoveDropRequest(context.Background(), id, dropRequestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlowfileQueuesApi.RemoveDropRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveDropRequest`: DropRequestEntity
    fmt.Fprintf(os.Stdout, "Response from `FlowfileQueuesApi.RemoveDropRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The connection id. | 
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

