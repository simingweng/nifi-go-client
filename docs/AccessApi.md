# \AccessApi

All URIs are relative to *http://localhost/nifi-api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccessToken**](AccessApi.md#CreateAccessToken) | **Post** /access/token | Creates a token for accessing the REST API via username/password
[**CreateAccessTokenFromTicket**](AccessApi.md#CreateAccessTokenFromTicket) | **Post** /access/kerberos | Creates a token for accessing the REST API via Kerberos ticket exchange / SPNEGO negotiation
[**CreateDownloadToken**](AccessApi.md#CreateDownloadToken) | **Post** /access/download-token | Creates a single use access token for downloading FlowFile content.
[**CreateUiExtensionToken**](AccessApi.md#CreateUiExtensionToken) | **Post** /access/ui-extension-token | Creates a single use access token for accessing a NiFi UI extension.
[**GetAccessStatus**](AccessApi.md#GetAccessStatus) | **Get** /access | Gets the status the client&#39;s access
[**GetLoginConfig**](AccessApi.md#GetLoginConfig) | **Get** /access/config | Retrieves the access configuration for this NiFi
[**KnoxCallback**](AccessApi.md#KnoxCallback) | **Get** /access/knox/callback | Redirect/callback URI for processing the result of the Apache Knox login sequence.
[**KnoxLogout**](AccessApi.md#KnoxLogout) | **Get** /access/knox/logout | Performs a logout in the Apache Knox.
[**KnoxRequest**](AccessApi.md#KnoxRequest) | **Get** /access/knox/request | Initiates a request to authenticate through Apache Knox.
[**LogOut**](AccessApi.md#LogOut) | **Delete** /access/logout | Performs a logout for other providers that have been issued a JWT.
[**OidcCallback**](AccessApi.md#OidcCallback) | **Get** /access/oidc/callback | Redirect/callback URI for processing the result of the OpenId Connect login sequence.
[**OidcExchange**](AccessApi.md#OidcExchange) | **Post** /access/oidc/exchange | Retrieves a JWT following a successful login sequence using the configured OpenId Connect provider.
[**OidcLogout**](AccessApi.md#OidcLogout) | **Get** /access/oidc/logout | Performs a logout in the OpenId Provider.
[**OidcRequest**](AccessApi.md#OidcRequest) | **Get** /access/oidc/request | Initiates a request to authenticate through the configured OpenId Connect provider.



## CreateAccessToken

> string CreateAccessToken(ctx, optional)

Creates a token for accessing the REST API via username/password

The token returned is formatted as a JSON Web Token (JWT). The token is base64 encoded and comprised of three parts. The header, the body, and the signature. The expiration of the token is a contained within the body. The token can be used in the Authorization header in the format 'Authorization: Bearer <token>'.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateAccessTokenOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateAccessTokenOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **username** | **optional.String**|  | 
 **password** | **optional.String**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAccessTokenFromTicket

> string CreateAccessTokenFromTicket(ctx, )

Creates a token for accessing the REST API via Kerberos ticket exchange / SPNEGO negotiation

The token returned is formatted as a JSON Web Token (JWT). The token is base64 encoded and comprised of three parts. The header, the body, and the signature. The expiration of the token is a contained within the body. The token can be used in the Authorization header in the format 'Authorization: Bearer <token>'.

### Required Parameters

This endpoint does not need any parameter.

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


## CreateDownloadToken

> string CreateDownloadToken(ctx, )

Creates a single use access token for downloading FlowFile content.

The token returned is a base64 encoded string. It is valid for a single request up to five minutes from being issued. It is used as a query parameter name 'access_token'.

### Required Parameters

This endpoint does not need any parameter.

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


## CreateUiExtensionToken

> string CreateUiExtensionToken(ctx, )

Creates a single use access token for accessing a NiFi UI extension.

The token returned is a base64 encoded string. It is valid for a single request up to five minutes from being issued. It is used as a query parameter name 'access_token'.

### Required Parameters

This endpoint does not need any parameter.

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


## GetAccessStatus

> AccessStatusEntity GetAccessStatus(ctx, )

Gets the status the client's access

Note: This endpoint is subject to change as NiFi and it's REST API evolve.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**AccessStatusEntity**](AccessStatusEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoginConfig

> AccessConfigurationEntity GetLoginConfig(ctx, )

Retrieves the access configuration for this NiFi

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**AccessConfigurationEntity**](AccessConfigurationEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KnoxCallback

> KnoxCallback(ctx, )

Redirect/callback URI for processing the result of the Apache Knox login sequence.

Note: This endpoint is subject to change as NiFi and it's REST API evolve.

### Required Parameters

This endpoint does not need any parameter.

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


## KnoxLogout

> KnoxLogout(ctx, )

Performs a logout in the Apache Knox.

Note: This endpoint is subject to change as NiFi and it's REST API evolve.

### Required Parameters

This endpoint does not need any parameter.

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


## KnoxRequest

> KnoxRequest(ctx, )

Initiates a request to authenticate through Apache Knox.

Note: This endpoint is subject to change as NiFi and it's REST API evolve.

### Required Parameters

This endpoint does not need any parameter.

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


## LogOut

> LogOut(ctx, )

Performs a logout for other providers that have been issued a JWT.

Note: This endpoint is subject to change as NiFi and it's REST API evolve.

### Required Parameters

This endpoint does not need any parameter.

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


## OidcCallback

> OidcCallback(ctx, )

Redirect/callback URI for processing the result of the OpenId Connect login sequence.

Note: This endpoint is subject to change as NiFi and it's REST API evolve.

### Required Parameters

This endpoint does not need any parameter.

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


## OidcExchange

> string OidcExchange(ctx, )

Retrieves a JWT following a successful login sequence using the configured OpenId Connect provider.

Note: This endpoint is subject to change as NiFi and it's REST API evolve.

### Required Parameters

This endpoint does not need any parameter.

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


## OidcLogout

> OidcLogout(ctx, )

Performs a logout in the OpenId Provider.

Note: This endpoint is subject to change as NiFi and it's REST API evolve.

### Required Parameters

This endpoint does not need any parameter.

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


## OidcRequest

> OidcRequest(ctx, )

Initiates a request to authenticate through the configured OpenId Connect provider.

Note: This endpoint is subject to change as NiFi and it's REST API evolve.

### Required Parameters

This endpoint does not need any parameter.

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

