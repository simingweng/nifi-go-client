# \ProvenanceEventsApi

All URIs are relative to *http://localhost/nifi-api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetInputContent**](ProvenanceEventsApi.md#GetInputContent) | **Get** /provenance-events/{id}/content/input | Gets the input content for a provenance event
[**GetOutputContent**](ProvenanceEventsApi.md#GetOutputContent) | **Get** /provenance-events/{id}/content/output | Gets the output content for a provenance event
[**GetProvenanceEvent**](ProvenanceEventsApi.md#GetProvenanceEvent) | **Get** /provenance-events/{id} | Gets a provenance event
[**SubmitReplay**](ProvenanceEventsApi.md#SubmitReplay) | **Post** /provenance-events/replays | Replays content from a provenance event



## GetInputContent

> map[string]interface{} GetInputContent(ctx, id).ClusterNodeId(clusterNodeId).Execute()

Gets the input content for a provenance event

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
    id := "id_example" // string | The provenance event id.
    clusterNodeId := "clusterNodeId_example" // string | The id of the node where the content exists if clustered. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProvenanceEventsApi.GetInputContent(context.Background(), id).ClusterNodeId(clusterNodeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvenanceEventsApi.GetInputContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInputContent`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ProvenanceEventsApi.GetInputContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The provenance event id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInputContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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


## GetOutputContent

> map[string]interface{} GetOutputContent(ctx, id).ClusterNodeId(clusterNodeId).Execute()

Gets the output content for a provenance event

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
    id := "id_example" // string | The provenance event id.
    clusterNodeId := "clusterNodeId_example" // string | The id of the node where the content exists if clustered. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProvenanceEventsApi.GetOutputContent(context.Background(), id).ClusterNodeId(clusterNodeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvenanceEventsApi.GetOutputContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOutputContent`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ProvenanceEventsApi.GetOutputContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The provenance event id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOutputContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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


## GetProvenanceEvent

> ProvenanceEventEntity GetProvenanceEvent(ctx, id).ClusterNodeId(clusterNodeId).Execute()

Gets a provenance event

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
    id := "id_example" // string | The provenance event id.
    clusterNodeId := "clusterNodeId_example" // string | The id of the node where this event exists if clustered. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProvenanceEventsApi.GetProvenanceEvent(context.Background(), id).ClusterNodeId(clusterNodeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvenanceEventsApi.GetProvenanceEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProvenanceEvent`: ProvenanceEventEntity
    fmt.Fprintf(os.Stdout, "Response from `ProvenanceEventsApi.GetProvenanceEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The provenance event id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProvenanceEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clusterNodeId** | **string** | The id of the node where this event exists if clustered. | 

### Return type

[**ProvenanceEventEntity**](ProvenanceEventEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitReplay

> ProvenanceEventEntity SubmitReplay(ctx).Body(body).Execute()

Replays content from a provenance event

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
    body := *openapiclient.NewSubmitReplayRequestEntity() // SubmitReplayRequestEntity | The replay request.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProvenanceEventsApi.SubmitReplay(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvenanceEventsApi.SubmitReplay``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubmitReplay`: ProvenanceEventEntity
    fmt.Fprintf(os.Stdout, "Response from `ProvenanceEventsApi.SubmitReplay`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitReplayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**SubmitReplayRequestEntity**](SubmitReplayRequestEntity.md) | The replay request. | 

### Return type

[**ProvenanceEventEntity**](ProvenanceEventEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

