# \SiteToSiteApi

All URIs are relative to *http://localhost/nifi-api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPeers**](SiteToSiteApi.md#GetPeers) | **Get** /site-to-site/peers | Returns the available Peers and its status of this NiFi
[**GetSiteToSiteDetails**](SiteToSiteApi.md#GetSiteToSiteDetails) | **Get** /site-to-site | Returns the details about this NiFi necessary to communicate via site to site



## GetPeers

> PeersEntity GetPeers(ctx).Execute()

Returns the available Peers and its status of this NiFi

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
    resp, r, err := api_client.SiteToSiteApi.GetPeers(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SiteToSiteApi.GetPeers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPeers`: PeersEntity
    fmt.Fprintf(os.Stdout, "Response from `SiteToSiteApi.GetPeers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPeersRequest struct via the builder pattern


### Return type

[**PeersEntity**](PeersEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSiteToSiteDetails

> ControllerEntity GetSiteToSiteDetails(ctx).Execute()

Returns the details about this NiFi necessary to communicate via site to site

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
    resp, r, err := api_client.SiteToSiteApi.GetSiteToSiteDetails(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SiteToSiteApi.GetSiteToSiteDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSiteToSiteDetails`: ControllerEntity
    fmt.Fprintf(os.Stdout, "Response from `SiteToSiteApi.GetSiteToSiteDetails`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteToSiteDetailsRequest struct via the builder pattern


### Return type

[**ControllerEntity**](ControllerEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

