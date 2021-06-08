# \FlowApi

All URIs are relative to *http://localhost/nifi-api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateControllerServices**](FlowApi.md#ActivateControllerServices) | **Put** /flow/process-groups/{id}/controller-services | Enable or disable Controller Services in the specified Process Group.
[**GenerateClientId**](FlowApi.md#GenerateClientId) | **Get** /flow/client-id | Generates a client id.
[**GetAboutInfo**](FlowApi.md#GetAboutInfo) | **Get** /flow/about | Retrieves details about this NiFi to put in the About dialog
[**GetAction**](FlowApi.md#GetAction) | **Get** /flow/history/{id} | Gets an action
[**GetBanners**](FlowApi.md#GetBanners) | **Get** /flow/banners | Retrieves the banners for this NiFi
[**GetBuckets**](FlowApi.md#GetBuckets) | **Get** /flow/registries/{id}/buckets | Gets the buckets from the specified registry for the current user
[**GetBulletinBoard**](FlowApi.md#GetBulletinBoard) | **Get** /flow/bulletin-board | Gets current bulletins
[**GetBulletins**](FlowApi.md#GetBulletins) | **Get** /flow/controller/bulletins | Retrieves Controller level bulletins
[**GetClusterSummary**](FlowApi.md#GetClusterSummary) | **Get** /flow/cluster/summary | The cluster summary for this NiFi
[**GetComponentHistory**](FlowApi.md#GetComponentHistory) | **Get** /flow/history/components/{componentId} | Gets configuration history for a component
[**GetConnectionStatistics**](FlowApi.md#GetConnectionStatistics) | **Get** /flow/connections/{id}/statistics | Gets statistics for a connection
[**GetConnectionStatus**](FlowApi.md#GetConnectionStatus) | **Get** /flow/connections/{id}/status | Gets status for a connection
[**GetConnectionStatusHistory**](FlowApi.md#GetConnectionStatusHistory) | **Get** /flow/connections/{id}/status/history | Gets the status history for a connection
[**GetControllerServiceTypes**](FlowApi.md#GetControllerServiceTypes) | **Get** /flow/controller-service-types | Retrieves the types of controller services that this NiFi supports
[**GetControllerServicesFromController**](FlowApi.md#GetControllerServicesFromController) | **Get** /flow/controller/controller-services | Gets controller services for reporting tasks
[**GetControllerServicesFromGroup**](FlowApi.md#GetControllerServicesFromGroup) | **Get** /flow/process-groups/{id}/controller-services | Gets all controller services
[**GetControllerStatus**](FlowApi.md#GetControllerStatus) | **Get** /flow/status | Gets the current status of this NiFi
[**GetCurrentUser**](FlowApi.md#GetCurrentUser) | **Get** /flow/current-user | Retrieves the user identity of the user making the request
[**GetFlow**](FlowApi.md#GetFlow) | **Get** /flow/process-groups/{id} | Gets a process group
[**GetFlowConfig**](FlowApi.md#GetFlowConfig) | **Get** /flow/config | Retrieves the configuration for this NiFi flow
[**GetFlowMetrics**](FlowApi.md#GetFlowMetrics) | **Get** /flow/metrics/{producer} | Gets all metrics for the flow from a particular node
[**GetFlows**](FlowApi.md#GetFlows) | **Get** /flow/registries/{registry-id}/buckets/{bucket-id}/flows | Gets the flows from the specified registry and bucket for the current user
[**GetInputPortStatus**](FlowApi.md#GetInputPortStatus) | **Get** /flow/input-ports/{id}/status | Gets status for an input port
[**GetOutputPortStatus**](FlowApi.md#GetOutputPortStatus) | **Get** /flow/output-ports/{id}/status | Gets status for an output port
[**GetParameterContexts**](FlowApi.md#GetParameterContexts) | **Get** /flow/parameter-contexts | Gets all Parameter Contexts
[**GetPrioritizers**](FlowApi.md#GetPrioritizers) | **Get** /flow/prioritizers | Retrieves the types of prioritizers that this NiFi supports
[**GetProcessGroupStatus**](FlowApi.md#GetProcessGroupStatus) | **Get** /flow/process-groups/{id}/status | Gets the status for a process group
[**GetProcessGroupStatusHistory**](FlowApi.md#GetProcessGroupStatusHistory) | **Get** /flow/process-groups/{id}/status/history | Gets status history for a remote process group
[**GetProcessorStatus**](FlowApi.md#GetProcessorStatus) | **Get** /flow/processors/{id}/status | Gets status for a processor
[**GetProcessorStatusHistory**](FlowApi.md#GetProcessorStatusHistory) | **Get** /flow/processors/{id}/status/history | Gets status history for a processor
[**GetProcessorTypes**](FlowApi.md#GetProcessorTypes) | **Get** /flow/processor-types | Retrieves the types of processors that this NiFi supports
[**GetRegistries**](FlowApi.md#GetRegistries) | **Get** /flow/registries | Gets the listing of available registries
[**GetRemoteProcessGroupStatus**](FlowApi.md#GetRemoteProcessGroupStatus) | **Get** /flow/remote-process-groups/{id}/status | Gets status for a remote process group
[**GetRemoteProcessGroupStatusHistory**](FlowApi.md#GetRemoteProcessGroupStatusHistory) | **Get** /flow/remote-process-groups/{id}/status/history | Gets the status history
[**GetReportingTaskTypes**](FlowApi.md#GetReportingTaskTypes) | **Get** /flow/reporting-task-types | Retrieves the types of reporting tasks that this NiFi supports
[**GetReportingTasks**](FlowApi.md#GetReportingTasks) | **Get** /flow/reporting-tasks | Gets all reporting tasks
[**GetTemplates**](FlowApi.md#GetTemplates) | **Get** /flow/templates | Gets all templates
[**GetVersions**](FlowApi.md#GetVersions) | **Get** /flow/registries/{registry-id}/buckets/{bucket-id}/flows/{flow-id}/versions | Gets the flow versions from the specified registry and bucket for the specified flow for the current user
[**QueryHistory**](FlowApi.md#QueryHistory) | **Get** /flow/history | Gets configuration history
[**ScheduleComponents**](FlowApi.md#ScheduleComponents) | **Put** /flow/process-groups/{id} | Schedule or unschedule components in the specified Process Group.
[**SearchCluster**](FlowApi.md#SearchCluster) | **Get** /flow/cluster/search-results | Searches the cluster for a node with the specified address
[**SearchFlow**](FlowApi.md#SearchFlow) | **Get** /flow/search-results | Performs a search against this NiFi using the specified search term



## ActivateControllerServices

> ActivateControllerServicesEntity ActivateControllerServices(ctx, id).Body(body).Execute()

Enable or disable Controller Services in the specified Process Group.

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
    body := *openapiclient.NewActivateControllerServicesEntity() // ActivateControllerServicesEntity | The request to schedule or unschedule. If the comopnents in the request are not specified, all authorized components will be considered.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FlowApi.ActivateControllerServices(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlowApi.ActivateControllerServices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActivateControllerServices`: ActivateControllerServicesEntity
    fmt.Fprintf(os.Stdout, "Response from `FlowApi.ActivateControllerServices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The process group id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivateControllerServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ActivateControllerServicesEntity**](ActivateControllerServicesEntity.md) | The request to schedule or unschedule. If the comopnents in the request are not specified, all authorized components will be considered. | 

### Return type

[**ActivateControllerServicesEntity**](ActivateControllerServicesEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateClientId

> string GenerateClientId(ctx).Execute()

Generates a client id.

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
    resp, r, err := api_client.FlowApi.GenerateClientId(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlowApi.GenerateClientId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateClientId`: string
    fmt.Fprintf(os.Stdout, "Response from `FlowApi.GenerateClientId`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGenerateClientIdRequest struct via the builder pattern


### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAboutInfo

> AboutEntity GetAboutInfo(ctx).Execute()

Retrieves details about this NiFi to put in the About dialog

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
    resp, r, err := api_client.FlowApi.GetAboutInfo(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlowApi.GetAboutInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAboutInfo`: AboutEntity
    fmt.Fprintf(os.Stdout, "Response from `FlowApi.GetAboutInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAboutInfoRequest struct via the builder pattern


### Return type

[**AboutEntity**](AboutEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAction

> ActionEntity GetAction(ctx, id).Execute()

Gets an action



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
    id := "id_example" // string | The action id.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FlowApi.GetAction(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlowApi.GetAction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAction`: ActionEntity
    fmt.Fprintf(os.Stdout, "Response from `FlowApi.GetAction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The action id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ActionEntity**](ActionEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBanners

> BannerEntity GetBanners(ctx).Execute()

Retrieves the banners for this NiFi

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
    resp, r, err := api_client.FlowApi.GetBanners(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlowApi.GetBanners``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBanners`: BannerEntity
    fmt.Fprintf(os.Stdout, "Response from `FlowApi.GetBanners`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBannersRequest struct via the builder pattern


### Return type

[**BannerEntity**](BannerEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBuckets

> BucketsEntity GetBuckets(ctx, id).Execute()

Gets the buckets from the specified registry for the current user

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
    resp, r, err := api_client.FlowApi.GetBuckets(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlowApi.GetBuckets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBuckets`: BucketsEntity
    fmt.Fprintf(os.Stdout, "Response from `FlowApi.GetBuckets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The registry id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBucketsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BucketsEntity**](BucketsEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBulletinBoard

> BulletinBoardEntity GetBulletinBoard(ctx).After(after).SourceName(sourceName).Message(message).SourceId(sourceId).GroupId(groupId).Limit(limit).Execute()

Gets current bulletins

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
    after := "after_example" // string | Includes bulletins with an id after this value. (optional)
    sourceName := "sourceName_example" // string | Includes bulletins originating from this sources whose name match this regular expression. (optional)
    message := "message_example" // string | Includes bulletins whose message that match this regular expression. (optional)
    sourceId := "sourceId_example" // string | Includes bulletins originating from this sources whose id match this regular expression. (optional)
    groupId := "groupId_example" // string | Includes bulletins originating from this sources whose group id match this regular expression. (optional)
    limit := "limit_example" // string | The number of bulletins to limit the response to. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FlowApi.GetBulletinBoard(context.Background()).After(after).SourceName(sourceName).Message(message).SourceId(sourceId).GroupId(groupId).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlowApi.GetBulletinBoard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBulletinBoard`: BulletinBoardEntity
    fmt.Fprintf(os.Stdout, "Response from `FlowApi.GetBulletinBoard`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBulletinBoardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **after** | **string** | Includes bulletins with an id after this value. | 
 **sourceName** | **string** | Includes bulletins originating from this sources whose name match this regular expression. | 
 **message** | **string** | Includes bulletins whose message that match this regular expression. | 
 **sourceId** | **string** | Includes bulletins originating from this sources whose id match this regular expression. | 
 **groupId** | **string** | Includes bulletins originating from this sources whose group id match this regular expression. | 
 **limit** | **string** | The number of bulletins to limit the response to. | 

### Return type

[**BulletinBoardEntity**](BulletinBoardEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBulletins

> ControllerBulletinsEntity GetBulletins(ctx).Execute()

Retrieves Controller level bulletins

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
    resp, r, err := api_client.FlowApi.GetBulletins(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlowApi.GetBulletins``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBulletins`: ControllerBulletinsEntity
    fmt.Fprintf(os.Stdout, "Response from `FlowApi.GetBulletins`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBulletinsRequest struct via the builder pattern


### Return type

[**ControllerBulletinsEntity**](ControllerBulletinsEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClusterSummary

> ClusteSummaryEntity GetClusterSummary(ctx).Execute()

The cluster summary for this NiFi

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
    resp, r, err := api_client.FlowApi.GetClusterSummary(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlowApi.GetClusterSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClusterSummary`: ClusteSummaryEntity
    fmt.Fprintf(os.Stdout, "Response from `FlowApi.GetClusterSummary`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterSummaryRequest struct via the builder pattern


### Return type

[**ClusteSummaryEntity**](ClusteSummaryEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetComponentHistory

> ComponentHistoryEntity GetComponentHistory(ctx, componentId).Execute()

Gets configuration history for a component



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
    componentId := "componentId_example" // string | The component id.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FlowApi.GetComponentHistory(context.Background(), componentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlowApi.GetComponentHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetComponentHistory`: ComponentHistoryEntity
    fmt.Fprintf(os.Stdout, "Response from `FlowApi.GetComponentHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**componentId** | **string** | The component id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetComponentHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetConnectionStatistics

> ConnectionStatisticsEntity GetConnectionStatistics(ctx, id).Nodewise(nodewise).ClusterNodeId(clusterNodeId).Execute()

Gets statistics for a connection

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
    nodewise := true // bool | Whether or not to include the breakdown per node. Optional, defaults to false (optional) (default to false)
    clusterNodeId := "clusterNodeId_example" // string | The id of the node where to get the statistics. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FlowApi.GetConnectionStatistics(context.Background(), id).Nodewise(nodewise).ClusterNodeId(clusterNodeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlowApi.GetConnectionStatistics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConnectionStatistics`: ConnectionStatisticsEntity
    fmt.Fprintf(os.Stdout, "Response from `FlowApi.GetConnectionStatistics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The connection id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectionStatisticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nodewise** | **bool** | Whether or not to include the breakdown per node. Optional, defaults to false | [default to false]
 **clusterNodeId** | **string** | The id of the node where to get the statistics. | 

### Return type

[**ConnectionStatisticsEntity**](ConnectionStatisticsEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConnectionStatus

> ConnectionStatusEntity GetConnectionStatus(ctx, id).Nodewise(nodewise).ClusterNodeId(clusterNodeId).Execute()

Gets status for a connection

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
    nodewise := true // bool | Whether or not to include the breakdown per node. Optional, defaults to false (optional) (default to false)
    clusterNodeId := "clusterNodeId_example" // string | The id of the node where to get the status. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FlowApi.GetConnectionStatus(context.Background(), id).Nodewise(nodewise).ClusterNodeId(clusterNodeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlowApi.GetConnectionStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConnectionStatus`: ConnectionStatusEntity
    fmt.Fprintf(os.Stdout, "Response from `FlowApi.GetConnectionStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The connection id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectionStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nodewise** | **bool** | Whether or not to include the breakdown per node. Optional, defaults to false | [default to false]
 **clusterNodeId** | **string** | The id of the node where to get the status. | 

### Return type

[**ConnectionStatusEntity**](ConnectionStatusEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConnectionStatusHistory

> StatusHistoryEntity GetConnectionStatusHistory(ctx, id).Execute()

Gets the status history for a connection

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
    resp, r, err := api_client.FlowApi.GetConnectionStatusHistory(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlowApi.GetConnectionStatusHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConnectionStatusHistory`: StatusHistoryEntity
    fmt.Fprintf(os.Stdout, "Response from `FlowApi.GetConnectionStatusHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The connection id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectionStatusHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StatusHistoryEntity**](StatusHistoryEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetControllerServiceTypes

> ControllerServiceTypesEntity GetControllerServiceTypes(ctx).ServiceType(serviceType).ServiceBundleGroup(serviceBundleGroup).ServiceBundleArtifact(serviceBundleArtifact).ServiceBundleVersion(serviceBundleVersion).BundleGroupFilter(bundleGroupFilter).BundleArtifactFilter(bundleArtifactFilter).TypeFilter(typeFilter).Execute()

Retrieves the types of controller services that this NiFi supports



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
    serviceType := "serviceType_example" // string | If specified, will only return controller services that are compatible with this type of service. (optional)
    serviceBundleGroup := "serviceBundleGroup_example" // string | If serviceType specified, is the bundle group of the serviceType. (optional)
    serviceBundleArtifact := "serviceBundleArtifact_example" // string | If serviceType specified, is the bundle artifact of the serviceType. (optional)
    serviceBundleVersion := "serviceBundleVersion_example" // string | If serviceType specified, is the bundle version of the serviceType. (optional)
    bundleGroupFilter := "bundleGroupFilter_example" // string | If specified, will only return types that are a member of this bundle group. (optional)
    bundleArtifactFilter := "bundleArtifactFilter_example" // string | If specified, will only return types that are a member of this bundle artifact. (optional)
    typeFilter := "typeFilter_example" // string | If specified, will only return types whose fully qualified classname matches. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FlowApi.GetControllerServiceTypes(context.Background()).ServiceType(serviceType).ServiceBundleGroup(serviceBundleGroup).ServiceBundleArtifact(serviceBundleArtifact).ServiceBundleVersion(serviceBundleVersion).BundleGroupFilter(bundleGroupFilter).BundleArtifactFilter(bundleArtifactFilter).TypeFilter(typeFilter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlowApi.GetControllerServiceTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetControllerServiceTypes`: ControllerServiceTypesEntity
    fmt.Fprintf(os.Stdout, "Response from `FlowApi.GetControllerServiceTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetControllerServiceTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serviceType** | **string** | If specified, will only return controller services that are compatible with this type of service. | 
 **serviceBundleGroup** | **string** | If serviceType specified, is the bundle group of the serviceType. | 
 **serviceBundleArtifact** | **string** | If serviceType specified, is the bundle artifact of the serviceType. | 
 **serviceBundleVersion** | **string** | If serviceType specified, is the bundle version of the serviceType. | 
 **bundleGroupFilter** | **string** | If specified, will only return types that are a member of this bundle group. | 
 **bundleArtifactFilter** | **string** | If specified, will only return types that are a member of this bundle artifact. | 
 **typeFilter** | **string** | If specified, will only return types whose fully qualified classname matches. | 

### Return type

[**ControllerServiceTypesEntity**](ControllerServiceTypesEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetControllerServicesFromController

> ControllerServicesEntity GetControllerServicesFromController(ctx).Execute()

Gets controller services for reporting tasks

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
    resp, r, err := api_client.FlowApi.GetControllerServicesFromController(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlowApi.GetControllerServicesFromController``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetControllerServicesFromController`: ControllerServicesEntity
    fmt.Fprintf(os.Stdout, "Response from `FlowApi.GetControllerServicesFromController`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetControllerServicesFromControllerRequest struct via the builder pattern


### Return type

[**ControllerServicesEntity**](ControllerServicesEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetControllerServicesFromGroup

> ControllerServicesEntity GetControllerServicesFromGroup(ctx, id).IncludeAncestorGroups(includeAncestorGroups).IncludeDescendantGroups(includeDescendantGroups).Execute()

Gets all controller services

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
    includeAncestorGroups := true // bool | Whether or not to include parent/ancestory process groups (optional) (default to true)
    includeDescendantGroups := true // bool | Whether or not to include descendant process groups (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FlowApi.GetControllerServicesFromGroup(context.Background(), id).IncludeAncestorGroups(includeAncestorGroups).IncludeDescendantGroups(includeDescendantGroups).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlowApi.GetControllerServicesFromGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetControllerServicesFromGroup`: ControllerServicesEntity
    fmt.Fprintf(os.Stdout, "Response from `FlowApi.GetControllerServicesFromGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The process group id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetControllerServicesFromGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeAncestorGroups** | **bool** | Whether or not to include parent/ancestory process groups | [default to true]
 **includeDescendantGroups** | **bool** | Whether or not to include descendant process groups | [default to false]

### Return type

[**ControllerServicesEntity**](ControllerServicesEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetControllerStatus

> ControllerStatusEntity GetControllerStatus(ctx).Execute()

Gets the current status of this NiFi

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
    resp, r, err := api_client.FlowApi.GetControllerStatus(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlowApi.GetControllerStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetControllerStatus`: ControllerStatusEntity
    fmt.Fprintf(os.Stdout, "Response from `FlowApi.GetControllerStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetControllerStatusRequest struct via the builder pattern


### Return type

[**ControllerStatusEntity**](ControllerStatusEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCurrentUser

> CurrentUserEntity GetCurrentUser(ctx).Execute()

Retrieves the user identity of the user making the request

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
    resp, r, err := api_client.FlowApi.GetCurrentUser(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlowApi.GetCurrentUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCurrentUser`: CurrentUserEntity
    fmt.Fprintf(os.Stdout, "Response from `FlowApi.GetCurrentUser`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrentUserRequest struct via the builder pattern


### Return type

[**CurrentUserEntity**](CurrentUserEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFlow

> ProcessGroupFlowEntity GetFlow(ctx, id).Execute()

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
    resp, r, err := api_client.FlowApi.GetFlow(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlowApi.GetFlow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFlow`: ProcessGroupFlowEntity
    fmt.Fprintf(os.Stdout, "Response from `FlowApi.GetFlow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The process group id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFlowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProcessGroupFlowEntity**](ProcessGroupFlowEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFlowConfig

> FlowConfigurationEntity GetFlowConfig(ctx).Execute()

Retrieves the configuration for this NiFi flow

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
    resp, r, err := api_client.FlowApi.GetFlowConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlowApi.GetFlowConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFlowConfig`: FlowConfigurationEntity
    fmt.Fprintf(os.Stdout, "Response from `FlowApi.GetFlowConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetFlowConfigRequest struct via the builder pattern


### Return type

[**FlowConfigurationEntity**](FlowConfigurationEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFlowMetrics

> map[string]interface{} GetFlowMetrics(ctx, producer).Execute()

Gets all metrics for the flow from a particular node

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
    producer := "producer_example" // string | The producer for flow file metrics. Each producer may have its own output format.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FlowApi.GetFlowMetrics(context.Background(), producer).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlowApi.GetFlowMetrics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFlowMetrics`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `FlowApi.GetFlowMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**producer** | **string** | The producer for flow file metrics. Each producer may have its own output format. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFlowMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetFlows

> VersionedFlowsEntity GetFlows(ctx, registryId, bucketId).Execute()

Gets the flows from the specified registry and bucket for the current user

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
    registryId := "registryId_example" // string | The registry id.
    bucketId := "bucketId_example" // string | The bucket id.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FlowApi.GetFlows(context.Background(), registryId, bucketId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlowApi.GetFlows``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFlows`: VersionedFlowsEntity
    fmt.Fprintf(os.Stdout, "Response from `FlowApi.GetFlows`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registryId** | **string** | The registry id. | 
**bucketId** | **string** | The bucket id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFlowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**VersionedFlowsEntity**](VersionedFlowsEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInputPortStatus

> PortStatusEntity GetInputPortStatus(ctx, id).Nodewise(nodewise).ClusterNodeId(clusterNodeId).Execute()

Gets status for an input port

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
    id := "id_example" // string | The input port id.
    nodewise := true // bool | Whether or not to include the breakdown per node. Optional, defaults to false (optional) (default to false)
    clusterNodeId := "clusterNodeId_example" // string | The id of the node where to get the status. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FlowApi.GetInputPortStatus(context.Background(), id).Nodewise(nodewise).ClusterNodeId(clusterNodeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlowApi.GetInputPortStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInputPortStatus`: PortStatusEntity
    fmt.Fprintf(os.Stdout, "Response from `FlowApi.GetInputPortStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The input port id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInputPortStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nodewise** | **bool** | Whether or not to include the breakdown per node. Optional, defaults to false | [default to false]
 **clusterNodeId** | **string** | The id of the node where to get the status. | 

### Return type

[**PortStatusEntity**](PortStatusEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOutputPortStatus

> PortStatusEntity GetOutputPortStatus(ctx, id).Nodewise(nodewise).ClusterNodeId(clusterNodeId).Execute()

Gets status for an output port

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
    id := "id_example" // string | The output port id.
    nodewise := true // bool | Whether or not to include the breakdown per node. Optional, defaults to false (optional) (default to false)
    clusterNodeId := "clusterNodeId_example" // string | The id of the node where to get the status. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FlowApi.GetOutputPortStatus(context.Background(), id).Nodewise(nodewise).ClusterNodeId(clusterNodeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlowApi.GetOutputPortStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOutputPortStatus`: PortStatusEntity
    fmt.Fprintf(os.Stdout, "Response from `FlowApi.GetOutputPortStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The output port id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOutputPortStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nodewise** | **bool** | Whether or not to include the breakdown per node. Optional, defaults to false | [default to false]
 **clusterNodeId** | **string** | The id of the node where to get the status. | 

### Return type

[**PortStatusEntity**](PortStatusEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetParameterContexts

> ParameterContextsEntity GetParameterContexts(ctx).Execute()

Gets all Parameter Contexts

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
    resp, r, err := api_client.FlowApi.GetParameterContexts(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlowApi.GetParameterContexts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetParameterContexts`: ParameterContextsEntity
    fmt.Fprintf(os.Stdout, "Response from `FlowApi.GetParameterContexts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetParameterContextsRequest struct via the builder pattern


### Return type

[**ParameterContextsEntity**](ParameterContextsEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPrioritizers

> PrioritizerTypesEntity GetPrioritizers(ctx).Execute()

Retrieves the types of prioritizers that this NiFi supports



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
    resp, r, err := api_client.FlowApi.GetPrioritizers(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlowApi.GetPrioritizers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPrioritizers`: PrioritizerTypesEntity
    fmt.Fprintf(os.Stdout, "Response from `FlowApi.GetPrioritizers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPrioritizersRequest struct via the builder pattern


### Return type

[**PrioritizerTypesEntity**](PrioritizerTypesEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProcessGroupStatus

> ProcessGroupStatusEntity GetProcessGroupStatus(ctx, id).Recursive(recursive).Nodewise(nodewise).ClusterNodeId(clusterNodeId).Execute()

Gets the status for a process group



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
    recursive := true // bool | Whether all descendant groups and the status of their content will be included. Optional, defaults to false (optional) (default to false)
    nodewise := true // bool | Whether or not to include the breakdown per node. Optional, defaults to false (optional) (default to false)
    clusterNodeId := "clusterNodeId_example" // string | The id of the node where to get the status. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FlowApi.GetProcessGroupStatus(context.Background(), id).Recursive(recursive).Nodewise(nodewise).ClusterNodeId(clusterNodeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlowApi.GetProcessGroupStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProcessGroupStatus`: ProcessGroupStatusEntity
    fmt.Fprintf(os.Stdout, "Response from `FlowApi.GetProcessGroupStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The process group id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProcessGroupStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **recursive** | **bool** | Whether all descendant groups and the status of their content will be included. Optional, defaults to false | [default to false]
 **nodewise** | **bool** | Whether or not to include the breakdown per node. Optional, defaults to false | [default to false]
 **clusterNodeId** | **string** | The id of the node where to get the status. | 

### Return type

[**ProcessGroupStatusEntity**](ProcessGroupStatusEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProcessGroupStatusHistory

> StatusHistoryEntity GetProcessGroupStatusHistory(ctx, id).Execute()

Gets status history for a remote process group

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
    resp, r, err := api_client.FlowApi.GetProcessGroupStatusHistory(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlowApi.GetProcessGroupStatusHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProcessGroupStatusHistory`: StatusHistoryEntity
    fmt.Fprintf(os.Stdout, "Response from `FlowApi.GetProcessGroupStatusHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The process group id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProcessGroupStatusHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StatusHistoryEntity**](StatusHistoryEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProcessorStatus

> ProcessorStatusEntity GetProcessorStatus(ctx, id).Nodewise(nodewise).ClusterNodeId(clusterNodeId).Execute()

Gets status for a processor

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
    nodewise := true // bool | Whether or not to include the breakdown per node. Optional, defaults to false (optional) (default to false)
    clusterNodeId := "clusterNodeId_example" // string | The id of the node where to get the status. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FlowApi.GetProcessorStatus(context.Background(), id).Nodewise(nodewise).ClusterNodeId(clusterNodeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlowApi.GetProcessorStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProcessorStatus`: ProcessorStatusEntity
    fmt.Fprintf(os.Stdout, "Response from `FlowApi.GetProcessorStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The processor id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProcessorStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nodewise** | **bool** | Whether or not to include the breakdown per node. Optional, defaults to false | [default to false]
 **clusterNodeId** | **string** | The id of the node where to get the status. | 

### Return type

[**ProcessorStatusEntity**](ProcessorStatusEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProcessorStatusHistory

> StatusHistoryEntity GetProcessorStatusHistory(ctx, id).Execute()

Gets status history for a processor

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
    resp, r, err := api_client.FlowApi.GetProcessorStatusHistory(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlowApi.GetProcessorStatusHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProcessorStatusHistory`: StatusHistoryEntity
    fmt.Fprintf(os.Stdout, "Response from `FlowApi.GetProcessorStatusHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The processor id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProcessorStatusHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StatusHistoryEntity**](StatusHistoryEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProcessorTypes

> ProcessorTypesEntity GetProcessorTypes(ctx).BundleGroupFilter(bundleGroupFilter).BundleArtifactFilter(bundleArtifactFilter).Type_(type_).Execute()

Retrieves the types of processors that this NiFi supports



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
    bundleGroupFilter := "bundleGroupFilter_example" // string | If specified, will only return types that are a member of this bundle group. (optional)
    bundleArtifactFilter := "bundleArtifactFilter_example" // string | If specified, will only return types that are a member of this bundle artifact. (optional)
    type_ := "type__example" // string | If specified, will only return types whose fully qualified classname matches. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FlowApi.GetProcessorTypes(context.Background()).BundleGroupFilter(bundleGroupFilter).BundleArtifactFilter(bundleArtifactFilter).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlowApi.GetProcessorTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProcessorTypes`: ProcessorTypesEntity
    fmt.Fprintf(os.Stdout, "Response from `FlowApi.GetProcessorTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProcessorTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bundleGroupFilter** | **string** | If specified, will only return types that are a member of this bundle group. | 
 **bundleArtifactFilter** | **string** | If specified, will only return types that are a member of this bundle artifact. | 
 **type_** | **string** | If specified, will only return types whose fully qualified classname matches. | 

### Return type

[**ProcessorTypesEntity**](ProcessorTypesEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRegistries

> RegistryClientsEntity GetRegistries(ctx).Execute()

Gets the listing of available registries

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
    resp, r, err := api_client.FlowApi.GetRegistries(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlowApi.GetRegistries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRegistries`: RegistryClientsEntity
    fmt.Fprintf(os.Stdout, "Response from `FlowApi.GetRegistries`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRegistriesRequest struct via the builder pattern


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


## GetRemoteProcessGroupStatus

> RemoteProcessGroupStatusEntity GetRemoteProcessGroupStatus(ctx, id).Nodewise(nodewise).ClusterNodeId(clusterNodeId).Execute()

Gets status for a remote process group

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
    nodewise := true // bool | Whether or not to include the breakdown per node. Optional, defaults to false (optional) (default to false)
    clusterNodeId := "clusterNodeId_example" // string | The id of the node where to get the status. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FlowApi.GetRemoteProcessGroupStatus(context.Background(), id).Nodewise(nodewise).ClusterNodeId(clusterNodeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlowApi.GetRemoteProcessGroupStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRemoteProcessGroupStatus`: RemoteProcessGroupStatusEntity
    fmt.Fprintf(os.Stdout, "Response from `FlowApi.GetRemoteProcessGroupStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The remote process group id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRemoteProcessGroupStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nodewise** | **bool** | Whether or not to include the breakdown per node. Optional, defaults to false | [default to false]
 **clusterNodeId** | **string** | The id of the node where to get the status. | 

### Return type

[**RemoteProcessGroupStatusEntity**](RemoteProcessGroupStatusEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRemoteProcessGroupStatusHistory

> StatusHistoryEntity GetRemoteProcessGroupStatusHistory(ctx, id).Execute()

Gets the status history

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
    resp, r, err := api_client.FlowApi.GetRemoteProcessGroupStatusHistory(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlowApi.GetRemoteProcessGroupStatusHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRemoteProcessGroupStatusHistory`: StatusHistoryEntity
    fmt.Fprintf(os.Stdout, "Response from `FlowApi.GetRemoteProcessGroupStatusHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The remote process group id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRemoteProcessGroupStatusHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StatusHistoryEntity**](StatusHistoryEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReportingTaskTypes

> ReportingTaskTypesEntity GetReportingTaskTypes(ctx).BundleGroupFilter(bundleGroupFilter).BundleArtifactFilter(bundleArtifactFilter).Type_(type_).Execute()

Retrieves the types of reporting tasks that this NiFi supports



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
    bundleGroupFilter := "bundleGroupFilter_example" // string | If specified, will only return types that are a member of this bundle group. (optional)
    bundleArtifactFilter := "bundleArtifactFilter_example" // string | If specified, will only return types that are a member of this bundle artifact. (optional)
    type_ := "type__example" // string | If specified, will only return types whose fully qualified classname matches. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FlowApi.GetReportingTaskTypes(context.Background()).BundleGroupFilter(bundleGroupFilter).BundleArtifactFilter(bundleArtifactFilter).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlowApi.GetReportingTaskTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReportingTaskTypes`: ReportingTaskTypesEntity
    fmt.Fprintf(os.Stdout, "Response from `FlowApi.GetReportingTaskTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetReportingTaskTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bundleGroupFilter** | **string** | If specified, will only return types that are a member of this bundle group. | 
 **bundleArtifactFilter** | **string** | If specified, will only return types that are a member of this bundle artifact. | 
 **type_** | **string** | If specified, will only return types whose fully qualified classname matches. | 

### Return type

[**ReportingTaskTypesEntity**](ReportingTaskTypesEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReportingTasks

> ReportingTasksEntity GetReportingTasks(ctx).Execute()

Gets all reporting tasks

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
    resp, r, err := api_client.FlowApi.GetReportingTasks(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlowApi.GetReportingTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReportingTasks`: ReportingTasksEntity
    fmt.Fprintf(os.Stdout, "Response from `FlowApi.GetReportingTasks`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetReportingTasksRequest struct via the builder pattern


### Return type

[**ReportingTasksEntity**](ReportingTasksEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTemplates

> TemplatesEntity GetTemplates(ctx).Execute()

Gets all templates

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
    resp, r, err := api_client.FlowApi.GetTemplates(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlowApi.GetTemplates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTemplates`: TemplatesEntity
    fmt.Fprintf(os.Stdout, "Response from `FlowApi.GetTemplates`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTemplatesRequest struct via the builder pattern


### Return type

[**TemplatesEntity**](TemplatesEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVersions

> VersionedFlowSnapshotMetadataSetEntity GetVersions(ctx, registryId, bucketId, flowId).Execute()

Gets the flow versions from the specified registry and bucket for the specified flow for the current user

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
    registryId := "registryId_example" // string | The registry id.
    bucketId := "bucketId_example" // string | The bucket id.
    flowId := "flowId_example" // string | The flow id.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FlowApi.GetVersions(context.Background(), registryId, bucketId, flowId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlowApi.GetVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVersions`: VersionedFlowSnapshotMetadataSetEntity
    fmt.Fprintf(os.Stdout, "Response from `FlowApi.GetVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registryId** | **string** | The registry id. | 
**bucketId** | **string** | The bucket id. | 
**flowId** | **string** | The flow id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**VersionedFlowSnapshotMetadataSetEntity**](VersionedFlowSnapshotMetadataSetEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryHistory

> HistoryEntity QueryHistory(ctx).Offset(offset).Count(count).SortColumn(sortColumn).SortOrder(sortOrder).StartDate(startDate).EndDate(endDate).UserIdentity(userIdentity).SourceId(sourceId).Execute()

Gets configuration history



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
    offset := "offset_example" // string | The offset into the result set.
    count := "count_example" // string | The number of actions to return.
    sortColumn := "sortColumn_example" // string | The field to sort on. (optional)
    sortOrder := "sortOrder_example" // string | The direction to sort. (optional)
    startDate := "startDate_example" // string | Include actions after this date. (optional)
    endDate := "endDate_example" // string | Include actions before this date. (optional)
    userIdentity := "userIdentity_example" // string | Include actions performed by this user. (optional)
    sourceId := "sourceId_example" // string | Include actions on this component. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FlowApi.QueryHistory(context.Background()).Offset(offset).Count(count).SortColumn(sortColumn).SortOrder(sortOrder).StartDate(startDate).EndDate(endDate).UserIdentity(userIdentity).SourceId(sourceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlowApi.QueryHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QueryHistory`: HistoryEntity
    fmt.Fprintf(os.Stdout, "Response from `FlowApi.QueryHistory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **string** | The offset into the result set. | 
 **count** | **string** | The number of actions to return. | 
 **sortColumn** | **string** | The field to sort on. | 
 **sortOrder** | **string** | The direction to sort. | 
 **startDate** | **string** | Include actions after this date. | 
 **endDate** | **string** | Include actions before this date. | 
 **userIdentity** | **string** | Include actions performed by this user. | 
 **sourceId** | **string** | Include actions on this component. | 

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


## ScheduleComponents

> ScheduleComponentsEntity ScheduleComponents(ctx, id).Body(body).Execute()

Schedule or unschedule components in the specified Process Group.

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
    body := *openapiclient.NewScheduleComponentsEntity() // ScheduleComponentsEntity | The request to schedule or unschedule. If the comopnents in the request are not specified, all authorized components will be considered.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FlowApi.ScheduleComponents(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlowApi.ScheduleComponents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ScheduleComponents`: ScheduleComponentsEntity
    fmt.Fprintf(os.Stdout, "Response from `FlowApi.ScheduleComponents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The process group id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiScheduleComponentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ScheduleComponentsEntity**](ScheduleComponentsEntity.md) | The request to schedule or unschedule. If the comopnents in the request are not specified, all authorized components will be considered. | 

### Return type

[**ScheduleComponentsEntity**](ScheduleComponentsEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchCluster

> ClusterSearchResultsEntity SearchCluster(ctx).Q(q).Execute()

Searches the cluster for a node with the specified address



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
    q := "q_example" // string | Node address to search for.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FlowApi.SearchCluster(context.Background()).Q(q).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlowApi.SearchCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchCluster`: ClusterSearchResultsEntity
    fmt.Fprintf(os.Stdout, "Response from `FlowApi.SearchCluster`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** | Node address to search for. | 

### Return type

[**ClusterSearchResultsEntity**](ClusterSearchResultsEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchFlow

> SearchResultsEntity SearchFlow(ctx).Q(q).A(a).Execute()

Performs a search against this NiFi using the specified search term



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
    q := "q_example" // string |  (optional)
    a := "a_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FlowApi.SearchFlow(context.Background()).Q(q).A(a).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlowApi.SearchFlow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchFlow`: SearchResultsEntity
    fmt.Fprintf(os.Stdout, "Response from `FlowApi.SearchFlow`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchFlowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** |  | 
 **a** | **string** |  | 

### Return type

[**SearchResultsEntity**](SearchResultsEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

