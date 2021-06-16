# \ProvenanceApi

All URIs are relative to *http://localhost/nifi-api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteLineage**](ProvenanceApi.md#DeleteLineage) | **Delete** /provenance/lineage/{id} | Deletes a lineage query
[**DeleteProvenance**](ProvenanceApi.md#DeleteProvenance) | **Delete** /provenance/{id} | Deletes a provenance query
[**GetLineage**](ProvenanceApi.md#GetLineage) | **Get** /provenance/lineage/{id} | Gets a lineage query
[**GetProvenance**](ProvenanceApi.md#GetProvenance) | **Get** /provenance/{id} | Gets a provenance query
[**GetSearchOptions**](ProvenanceApi.md#GetSearchOptions) | **Get** /provenance/search-options | Gets the searchable attributes for provenance events
[**SubmitLineageRequest**](ProvenanceApi.md#SubmitLineageRequest) | **Post** /provenance/lineage | Submits a lineage query
[**SubmitProvenanceRequest**](ProvenanceApi.md#SubmitProvenanceRequest) | **Post** /provenance | Submits a provenance query



## DeleteLineage

> LineageEntity DeleteLineage(ctx, id).ClusterNodeId(clusterNodeId).Execute()

Deletes a lineage query

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
    id := "id_example" // string | The id of the lineage query.
    clusterNodeId := "clusterNodeId_example" // string | The id of the node where this query exists if clustered. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProvenanceApi.DeleteLineage(context.Background(), id).ClusterNodeId(clusterNodeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvenanceApi.DeleteLineage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteLineage`: LineageEntity
    fmt.Fprintf(os.Stdout, "Response from `ProvenanceApi.DeleteLineage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the lineage query. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLineageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clusterNodeId** | **string** | The id of the node where this query exists if clustered. | 

### Return type

[**LineageEntity**](LineageEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteProvenance

> ProvenanceEntity DeleteProvenance(ctx, id).ClusterNodeId(clusterNodeId).Execute()

Deletes a provenance query

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
    id := "id_example" // string | The id of the provenance query.
    clusterNodeId := "clusterNodeId_example" // string | The id of the node where this query exists if clustered. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProvenanceApi.DeleteProvenance(context.Background(), id).ClusterNodeId(clusterNodeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvenanceApi.DeleteProvenance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteProvenance`: ProvenanceEntity
    fmt.Fprintf(os.Stdout, "Response from `ProvenanceApi.DeleteProvenance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the provenance query. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProvenanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clusterNodeId** | **string** | The id of the node where this query exists if clustered. | 

### Return type

[**ProvenanceEntity**](ProvenanceEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLineage

> LineageEntity GetLineage(ctx, id).ClusterNodeId(clusterNodeId).Execute()

Gets a lineage query

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
    id := "id_example" // string | The id of the lineage query.
    clusterNodeId := "clusterNodeId_example" // string | The id of the node where this query exists if clustered. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProvenanceApi.GetLineage(context.Background(), id).ClusterNodeId(clusterNodeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvenanceApi.GetLineage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLineage`: LineageEntity
    fmt.Fprintf(os.Stdout, "Response from `ProvenanceApi.GetLineage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the lineage query. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLineageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clusterNodeId** | **string** | The id of the node where this query exists if clustered. | 

### Return type

[**LineageEntity**](LineageEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProvenance

> ProvenanceEntity GetProvenance(ctx, id).ClusterNodeId(clusterNodeId).Summarize(summarize).IncrementalResults(incrementalResults).Execute()

Gets a provenance query

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
    id := "id_example" // string | The id of the provenance query.
    clusterNodeId := "clusterNodeId_example" // string | The id of the node where this query exists if clustered. (optional)
    summarize := true // bool | Whether or not incremental results are returned. If false, provenance events are only returned once the query completes. This property is true by default. (optional) (default to false)
    incrementalResults := true // bool | Whether or not to summarize provenance events returned. This property is false by default. (optional) (default to true)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProvenanceApi.GetProvenance(context.Background(), id).ClusterNodeId(clusterNodeId).Summarize(summarize).IncrementalResults(incrementalResults).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvenanceApi.GetProvenance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProvenance`: ProvenanceEntity
    fmt.Fprintf(os.Stdout, "Response from `ProvenanceApi.GetProvenance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the provenance query. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProvenanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clusterNodeId** | **string** | The id of the node where this query exists if clustered. | 
 **summarize** | **bool** | Whether or not incremental results are returned. If false, provenance events are only returned once the query completes. This property is true by default. | [default to false]
 **incrementalResults** | **bool** | Whether or not to summarize provenance events returned. This property is false by default. | [default to true]

### Return type

[**ProvenanceEntity**](ProvenanceEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSearchOptions

> ProvenanceOptionsEntity GetSearchOptions(ctx).Execute()

Gets the searchable attributes for provenance events

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
    resp, r, err := api_client.ProvenanceApi.GetSearchOptions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvenanceApi.GetSearchOptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSearchOptions`: ProvenanceOptionsEntity
    fmt.Fprintf(os.Stdout, "Response from `ProvenanceApi.GetSearchOptions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSearchOptionsRequest struct via the builder pattern


### Return type

[**ProvenanceOptionsEntity**](ProvenanceOptionsEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitLineageRequest

> LineageEntity SubmitLineageRequest(ctx).Body(body).Execute()

Submits a lineage query



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
    body := *openapiclient.NewLineageEntity() // LineageEntity | The lineage query details.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProvenanceApi.SubmitLineageRequest(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvenanceApi.SubmitLineageRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubmitLineageRequest`: LineageEntity
    fmt.Fprintf(os.Stdout, "Response from `ProvenanceApi.SubmitLineageRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitLineageRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**LineageEntity**](LineageEntity.md) | The lineage query details. | 

### Return type

[**LineageEntity**](LineageEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitProvenanceRequest

> ProvenanceEntity SubmitProvenanceRequest(ctx).Body(body).Execute()

Submits a provenance query



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
    body := *openapiclient.NewProvenanceEntity() // ProvenanceEntity | The provenance query details.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProvenanceApi.SubmitProvenanceRequest(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvenanceApi.SubmitProvenanceRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubmitProvenanceRequest`: ProvenanceEntity
    fmt.Fprintf(os.Stdout, "Response from `ProvenanceApi.SubmitProvenanceRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitProvenanceRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ProvenanceEntity**](ProvenanceEntity.md) | The provenance query details. | 

### Return type

[**ProvenanceEntity**](ProvenanceEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

