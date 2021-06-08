# \DataTransferApi

All URIs are relative to *http://localhost/nifi-api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CommitInputPortTransaction**](DataTransferApi.md#CommitInputPortTransaction) | **Delete** /data-transfer/input-ports/{portId}/transactions/{transactionId} | Commit or cancel the specified transaction
[**CommitOutputPortTransaction**](DataTransferApi.md#CommitOutputPortTransaction) | **Delete** /data-transfer/output-ports/{portId}/transactions/{transactionId} | Commit or cancel the specified transaction
[**CreatePortTransaction**](DataTransferApi.md#CreatePortTransaction) | **Post** /data-transfer/{portType}/{portId}/transactions | Create a transaction to the specified output port or input port
[**ExtendInputPortTransactionTTL**](DataTransferApi.md#ExtendInputPortTransactionTTL) | **Put** /data-transfer/input-ports/{portId}/transactions/{transactionId} | Extend transaction TTL
[**ExtendOutputPortTransactionTTL**](DataTransferApi.md#ExtendOutputPortTransactionTTL) | **Put** /data-transfer/output-ports/{portId}/transactions/{transactionId} | Extend transaction TTL
[**ReceiveFlowFiles**](DataTransferApi.md#ReceiveFlowFiles) | **Post** /data-transfer/input-ports/{portId}/transactions/{transactionId}/flow-files | Transfer flow files to the input port
[**TransferFlowFiles**](DataTransferApi.md#TransferFlowFiles) | **Get** /data-transfer/output-ports/{portId}/transactions/{transactionId}/flow-files | Transfer flow files from the output port



## CommitInputPortTransaction

> TransactionResultEntity CommitInputPortTransaction(ctx, portId, transactionId).ResponseCode(responseCode).Execute()

Commit or cancel the specified transaction

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
    responseCode := int32(56) // int32 | The response code. Available values are BAD_CHECKSUM(19), CONFIRM_TRANSACTION(12) or CANCEL_TRANSACTION(15).
    portId := "portId_example" // string | The input port id.
    transactionId := "transactionId_example" // string | The transaction id.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DataTransferApi.CommitInputPortTransaction(context.Background(), portId, transactionId).ResponseCode(responseCode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataTransferApi.CommitInputPortTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CommitInputPortTransaction`: TransactionResultEntity
    fmt.Fprintf(os.Stdout, "Response from `DataTransferApi.CommitInputPortTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**portId** | **string** | The input port id. | 
**transactionId** | **string** | The transaction id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommitInputPortTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **responseCode** | **int32** | The response code. Available values are BAD_CHECKSUM(19), CONFIRM_TRANSACTION(12) or CANCEL_TRANSACTION(15). | 



### Return type

[**TransactionResultEntity**](TransactionResultEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommitOutputPortTransaction

> TransactionResultEntity CommitOutputPortTransaction(ctx, portId, transactionId).ResponseCode(responseCode).Checksum(checksum).Execute()

Commit or cancel the specified transaction

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
    responseCode := int32(56) // int32 | The response code. Available values are CONFIRM_TRANSACTION(12) or CANCEL_TRANSACTION(15).
    checksum := "checksum_example" // string | A checksum calculated at client side using CRC32 to check flow file content integrity. It must match with the value calculated at server side.
    portId := "portId_example" // string | The output port id.
    transactionId := "transactionId_example" // string | The transaction id.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DataTransferApi.CommitOutputPortTransaction(context.Background(), portId, transactionId).ResponseCode(responseCode).Checksum(checksum).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataTransferApi.CommitOutputPortTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CommitOutputPortTransaction`: TransactionResultEntity
    fmt.Fprintf(os.Stdout, "Response from `DataTransferApi.CommitOutputPortTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**portId** | **string** | The output port id. | 
**transactionId** | **string** | The transaction id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommitOutputPortTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **responseCode** | **int32** | The response code. Available values are CONFIRM_TRANSACTION(12) or CANCEL_TRANSACTION(15). | 
 **checksum** | **string** | A checksum calculated at client side using CRC32 to check flow file content integrity. It must match with the value calculated at server side. | 



### Return type

[**TransactionResultEntity**](TransactionResultEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePortTransaction

> TransactionResultEntity CreatePortTransaction(ctx, portType, portId).Execute()

Create a transaction to the specified output port or input port

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
    portType := "portType_example" // string | The port type.
    portId := "portId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DataTransferApi.CreatePortTransaction(context.Background(), portType, portId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataTransferApi.CreatePortTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePortTransaction`: TransactionResultEntity
    fmt.Fprintf(os.Stdout, "Response from `DataTransferApi.CreatePortTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**portType** | **string** | The port type. | 
**portId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePortTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**TransactionResultEntity**](TransactionResultEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExtendInputPortTransactionTTL

> TransactionResultEntity ExtendInputPortTransactionTTL(ctx, portId, transactionId).Execute()

Extend transaction TTL

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
    portId := "portId_example" // string | 
    transactionId := "transactionId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DataTransferApi.ExtendInputPortTransactionTTL(context.Background(), portId, transactionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataTransferApi.ExtendInputPortTransactionTTL``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExtendInputPortTransactionTTL`: TransactionResultEntity
    fmt.Fprintf(os.Stdout, "Response from `DataTransferApi.ExtendInputPortTransactionTTL`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**portId** | **string** |  | 
**transactionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExtendInputPortTransactionTTLRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**TransactionResultEntity**](TransactionResultEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExtendOutputPortTransactionTTL

> TransactionResultEntity ExtendOutputPortTransactionTTL(ctx, portId, transactionId).Execute()

Extend transaction TTL

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
    portId := "portId_example" // string | 
    transactionId := "transactionId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DataTransferApi.ExtendOutputPortTransactionTTL(context.Background(), portId, transactionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataTransferApi.ExtendOutputPortTransactionTTL``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExtendOutputPortTransactionTTL`: TransactionResultEntity
    fmt.Fprintf(os.Stdout, "Response from `DataTransferApi.ExtendOutputPortTransactionTTL`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**portId** | **string** |  | 
**transactionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExtendOutputPortTransactionTTLRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**TransactionResultEntity**](TransactionResultEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReceiveFlowFiles

> string ReceiveFlowFiles(ctx, portId, transactionId).Execute()

Transfer flow files to the input port

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
    portId := "portId_example" // string | The input port id.
    transactionId := "transactionId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DataTransferApi.ReceiveFlowFiles(context.Background(), portId, transactionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataTransferApi.ReceiveFlowFiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReceiveFlowFiles`: string
    fmt.Fprintf(os.Stdout, "Response from `DataTransferApi.ReceiveFlowFiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**portId** | **string** | The input port id. | 
**transactionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReceiveFlowFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## TransferFlowFiles

> map[string]interface{} TransferFlowFiles(ctx, portId, transactionId).Execute()

Transfer flow files from the output port

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
    portId := "portId_example" // string | The output port id.
    transactionId := "transactionId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DataTransferApi.TransferFlowFiles(context.Background(), portId, transactionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataTransferApi.TransferFlowFiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TransferFlowFiles`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DataTransferApi.TransferFlowFiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**portId** | **string** | The output port id. | 
**transactionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTransferFlowFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

