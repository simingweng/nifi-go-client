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

> LineageEntity DeleteLineage(ctx, id, optional)

Deletes a lineage query

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The id of the lineage query. | 
 **optional** | ***DeleteLineageOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteLineageOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clusterNodeId** | **optional.String**| The id of the node where this query exists if clustered. | 

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

> ProvenanceEntity DeleteProvenance(ctx, id, optional)

Deletes a provenance query

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The id of the provenance query. | 
 **optional** | ***DeleteProvenanceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteProvenanceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clusterNodeId** | **optional.String**| The id of the node where this query exists if clustered. | 

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

> LineageEntity GetLineage(ctx, id, optional)

Gets a lineage query

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The id of the lineage query. | 
 **optional** | ***GetLineageOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetLineageOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clusterNodeId** | **optional.String**| The id of the node where this query exists if clustered. | 

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

> ProvenanceEntity GetProvenance(ctx, id, optional)

Gets a provenance query

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The id of the provenance query. | 
 **optional** | ***GetProvenanceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetProvenanceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clusterNodeId** | **optional.String**| The id of the node where this query exists if clustered. | 
 **summarize** | **optional.Bool**| Whether or not incremental results are returned. If false, provenance events are only returned once the query completes. This property is true by default. | [default to false]
 **incrementalResults** | **optional.Bool**| Whether or not to summarize provenance events returned. This property is false by default. | [default to true]

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

> ProvenanceOptionsEntity GetSearchOptions(ctx, )

Gets the searchable attributes for provenance events

### Required Parameters

This endpoint does not need any parameter.

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

> LineageEntity SubmitLineageRequest(ctx, body)

Submits a lineage query

Lineage queries may be long running so this endpoint submits a request. The response will include the current state of the query. If the request is not completed the URI in the response can be used at a later time to get the updated state of the query. Once the query has completed the lineage request should be deleted by the client who originally submitted it.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**LineageEntity**](LineageEntity.md)| The lineage query details. | 

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

> ProvenanceEntity SubmitProvenanceRequest(ctx, body)

Submits a provenance query

Provenance queries may be long running so this endpoint submits a request. The response will include the current state of the query. If the request is not completed the URI in the response can be used at a later time to get the updated state of the query. Once the query has completed the provenance request should be deleted by the client who originally submitted it.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**ProvenanceEntity**](ProvenanceEntity.md)| The provenance query details. | 

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

