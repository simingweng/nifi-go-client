/*
 * NiFi Rest Api
 *
 * The Rest Api provides programmatic access to command and control a NiFi instance in real time. Start and                                              stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.13.2
 * Contact: dev@nifi.apache.org
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nifi

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// ControllerServicesApiService ControllerServicesApi service
type ControllerServicesApiService service

type ControllerServicesApiApiClearStateRequest struct {
	ctx        _context.Context
	ApiService *ControllerServicesApiService
	id         string
}

func (r ControllerServicesApiApiClearStateRequest) Execute() (ComponentStateEntity, *_nethttp.Response, error) {
	return r.ApiService.ClearStateExecute(r)
}

/*
 * ClearState Clears the state for a controller service
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id The controller service id.
 * @return ControllerServicesApiApiClearStateRequest
 */
func (a *ControllerServicesApiService) ClearState(ctx _context.Context, id string) ControllerServicesApiApiClearStateRequest {
	return ControllerServicesApiApiClearStateRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

/*
 * Execute executes the request
 * @return ComponentStateEntity
 */
func (a *ControllerServicesApiService) ClearStateExecute(r ControllerServicesApiApiClearStateRequest) (ComponentStateEntity, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ComponentStateEntity
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ControllerServicesApiService.ClearState")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/controller-services/{id}/state/clear-requests"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ControllerServicesApiApiGetControllerServiceRequest struct {
	ctx        _context.Context
	ApiService *ControllerServicesApiService
	id         string
}

func (r ControllerServicesApiApiGetControllerServiceRequest) Execute() (ControllerServiceEntity, *_nethttp.Response, error) {
	return r.ApiService.GetControllerServiceExecute(r)
}

/*
 * GetControllerService Gets a controller service
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id The controller service id.
 * @return ControllerServicesApiApiGetControllerServiceRequest
 */
func (a *ControllerServicesApiService) GetControllerService(ctx _context.Context, id string) ControllerServicesApiApiGetControllerServiceRequest {
	return ControllerServicesApiApiGetControllerServiceRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

/*
 * Execute executes the request
 * @return ControllerServiceEntity
 */
func (a *ControllerServicesApiService) GetControllerServiceExecute(r ControllerServicesApiApiGetControllerServiceRequest) (ControllerServiceEntity, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ControllerServiceEntity
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ControllerServicesApiService.GetControllerService")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/controller-services/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ControllerServicesApiApiGetControllerServiceReferencesRequest struct {
	ctx        _context.Context
	ApiService *ControllerServicesApiService
	id         string
}

func (r ControllerServicesApiApiGetControllerServiceReferencesRequest) Execute() (ControllerServiceReferencingComponentsEntity, *_nethttp.Response, error) {
	return r.ApiService.GetControllerServiceReferencesExecute(r)
}

/*
 * GetControllerServiceReferences Gets a controller service
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id The controller service id.
 * @return ControllerServicesApiApiGetControllerServiceReferencesRequest
 */
func (a *ControllerServicesApiService) GetControllerServiceReferences(ctx _context.Context, id string) ControllerServicesApiApiGetControllerServiceReferencesRequest {
	return ControllerServicesApiApiGetControllerServiceReferencesRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

/*
 * Execute executes the request
 * @return ControllerServiceReferencingComponentsEntity
 */
func (a *ControllerServicesApiService) GetControllerServiceReferencesExecute(r ControllerServicesApiApiGetControllerServiceReferencesRequest) (ControllerServiceReferencingComponentsEntity, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ControllerServiceReferencingComponentsEntity
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ControllerServicesApiService.GetControllerServiceReferences")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/controller-services/{id}/references"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ControllerServicesApiApiGetPropertyDescriptorRequest struct {
	ctx          _context.Context
	ApiService   *ControllerServicesApiService
	id           string
	propertyName *string
}

func (r ControllerServicesApiApiGetPropertyDescriptorRequest) PropertyName(propertyName string) ControllerServicesApiApiGetPropertyDescriptorRequest {
	r.propertyName = &propertyName
	return r
}

func (r ControllerServicesApiApiGetPropertyDescriptorRequest) Execute() (PropertyDescriptorEntity, *_nethttp.Response, error) {
	return r.ApiService.GetPropertyDescriptorExecute(r)
}

/*
 * GetPropertyDescriptor Gets a controller service property descriptor
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id The controller service id.
 * @return ControllerServicesApiApiGetPropertyDescriptorRequest
 */
func (a *ControllerServicesApiService) GetPropertyDescriptor(ctx _context.Context, id string) ControllerServicesApiApiGetPropertyDescriptorRequest {
	return ControllerServicesApiApiGetPropertyDescriptorRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

/*
 * Execute executes the request
 * @return PropertyDescriptorEntity
 */
func (a *ControllerServicesApiService) GetPropertyDescriptorExecute(r ControllerServicesApiApiGetPropertyDescriptorRequest) (PropertyDescriptorEntity, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  PropertyDescriptorEntity
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ControllerServicesApiService.GetPropertyDescriptor")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/controller-services/{id}/descriptors"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.propertyName == nil {
		return localVarReturnValue, nil, reportError("propertyName is required and must be specified")
	}

	localVarQueryParams.Add("propertyName", parameterToString(*r.propertyName, ""))
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ControllerServicesApiApiGetStateRequest struct {
	ctx        _context.Context
	ApiService *ControllerServicesApiService
	id         string
}

func (r ControllerServicesApiApiGetStateRequest) Execute() (ComponentStateEntity, *_nethttp.Response, error) {
	return r.ApiService.GetStateExecute(r)
}

/*
 * GetState Gets the state for a controller service
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id The controller service id.
 * @return ControllerServicesApiApiGetStateRequest
 */
func (a *ControllerServicesApiService) GetState(ctx _context.Context, id string) ControllerServicesApiApiGetStateRequest {
	return ControllerServicesApiApiGetStateRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

/*
 * Execute executes the request
 * @return ComponentStateEntity
 */
func (a *ControllerServicesApiService) GetStateExecute(r ControllerServicesApiApiGetStateRequest) (ComponentStateEntity, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ComponentStateEntity
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ControllerServicesApiService.GetState")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/controller-services/{id}/state"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ControllerServicesApiApiRemoveControllerServiceRequest struct {
	ctx                          _context.Context
	ApiService                   *ControllerServicesApiService
	id                           string
	version                      *string
	clientId                     *string
	disconnectedNodeAcknowledged *bool
}

func (r ControllerServicesApiApiRemoveControllerServiceRequest) Version(version string) ControllerServicesApiApiRemoveControllerServiceRequest {
	r.version = &version
	return r
}
func (r ControllerServicesApiApiRemoveControllerServiceRequest) ClientId(clientId string) ControllerServicesApiApiRemoveControllerServiceRequest {
	r.clientId = &clientId
	return r
}
func (r ControllerServicesApiApiRemoveControllerServiceRequest) DisconnectedNodeAcknowledged(disconnectedNodeAcknowledged bool) ControllerServicesApiApiRemoveControllerServiceRequest {
	r.disconnectedNodeAcknowledged = &disconnectedNodeAcknowledged
	return r
}

func (r ControllerServicesApiApiRemoveControllerServiceRequest) Execute() (ControllerServiceEntity, *_nethttp.Response, error) {
	return r.ApiService.RemoveControllerServiceExecute(r)
}

/*
 * RemoveControllerService Deletes a controller service
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id The controller service id.
 * @return ControllerServicesApiApiRemoveControllerServiceRequest
 */
func (a *ControllerServicesApiService) RemoveControllerService(ctx _context.Context, id string) ControllerServicesApiApiRemoveControllerServiceRequest {
	return ControllerServicesApiApiRemoveControllerServiceRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

/*
 * Execute executes the request
 * @return ControllerServiceEntity
 */
func (a *ControllerServicesApiService) RemoveControllerServiceExecute(r ControllerServicesApiApiRemoveControllerServiceRequest) (ControllerServiceEntity, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ControllerServiceEntity
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ControllerServicesApiService.RemoveControllerService")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/controller-services/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.version != nil {
		localVarQueryParams.Add("version", parameterToString(*r.version, ""))
	}
	if r.clientId != nil {
		localVarQueryParams.Add("clientId", parameterToString(*r.clientId, ""))
	}
	if r.disconnectedNodeAcknowledged != nil {
		localVarQueryParams.Add("disconnectedNodeAcknowledged", parameterToString(*r.disconnectedNodeAcknowledged, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ControllerServicesApiApiUpdateControllerServiceRequest struct {
	ctx        _context.Context
	ApiService *ControllerServicesApiService
	id         string
	body       *ControllerServiceEntity
}

func (r ControllerServicesApiApiUpdateControllerServiceRequest) Body(body ControllerServiceEntity) ControllerServicesApiApiUpdateControllerServiceRequest {
	r.body = &body
	return r
}

func (r ControllerServicesApiApiUpdateControllerServiceRequest) Execute() (ControllerServiceEntity, *_nethttp.Response, error) {
	return r.ApiService.UpdateControllerServiceExecute(r)
}

/*
 * UpdateControllerService Updates a controller service
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id The controller service id.
 * @return ControllerServicesApiApiUpdateControllerServiceRequest
 */
func (a *ControllerServicesApiService) UpdateControllerService(ctx _context.Context, id string) ControllerServicesApiApiUpdateControllerServiceRequest {
	return ControllerServicesApiApiUpdateControllerServiceRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

/*
 * Execute executes the request
 * @return ControllerServiceEntity
 */
func (a *ControllerServicesApiService) UpdateControllerServiceExecute(r ControllerServicesApiApiUpdateControllerServiceRequest) (ControllerServiceEntity, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ControllerServiceEntity
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ControllerServicesApiService.UpdateControllerService")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/controller-services/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.body
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ControllerServicesApiApiUpdateControllerServiceReferencesRequest struct {
	ctx        _context.Context
	ApiService *ControllerServicesApiService
	id         string
	body       *UpdateControllerServiceReferenceRequestEntity
}

func (r ControllerServicesApiApiUpdateControllerServiceReferencesRequest) Body(body UpdateControllerServiceReferenceRequestEntity) ControllerServicesApiApiUpdateControllerServiceReferencesRequest {
	r.body = &body
	return r
}

func (r ControllerServicesApiApiUpdateControllerServiceReferencesRequest) Execute() (ControllerServiceReferencingComponentsEntity, *_nethttp.Response, error) {
	return r.ApiService.UpdateControllerServiceReferencesExecute(r)
}

/*
 * UpdateControllerServiceReferences Updates a controller services references
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id The controller service id.
 * @return ControllerServicesApiApiUpdateControllerServiceReferencesRequest
 */
func (a *ControllerServicesApiService) UpdateControllerServiceReferences(ctx _context.Context, id string) ControllerServicesApiApiUpdateControllerServiceReferencesRequest {
	return ControllerServicesApiApiUpdateControllerServiceReferencesRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

/*
 * Execute executes the request
 * @return ControllerServiceReferencingComponentsEntity
 */
func (a *ControllerServicesApiService) UpdateControllerServiceReferencesExecute(r ControllerServicesApiApiUpdateControllerServiceReferencesRequest) (ControllerServiceReferencingComponentsEntity, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ControllerServiceReferencingComponentsEntity
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ControllerServicesApiService.UpdateControllerServiceReferences")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/controller-services/{id}/references"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.body
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ControllerServicesApiApiUpdateRunStatusRequest struct {
	ctx        _context.Context
	ApiService *ControllerServicesApiService
	id         string
	body       *ControllerServiceRunStatusEntity
}

func (r ControllerServicesApiApiUpdateRunStatusRequest) Body(body ControllerServiceRunStatusEntity) ControllerServicesApiApiUpdateRunStatusRequest {
	r.body = &body
	return r
}

func (r ControllerServicesApiApiUpdateRunStatusRequest) Execute() (ControllerServiceEntity, *_nethttp.Response, error) {
	return r.ApiService.UpdateRunStatusExecute(r)
}

/*
 * UpdateRunStatus Updates run status of a controller service
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id The controller service id.
 * @return ControllerServicesApiApiUpdateRunStatusRequest
 */
func (a *ControllerServicesApiService) UpdateRunStatus(ctx _context.Context, id string) ControllerServicesApiApiUpdateRunStatusRequest {
	return ControllerServicesApiApiUpdateRunStatusRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

/*
 * Execute executes the request
 * @return ControllerServiceEntity
 */
func (a *ControllerServicesApiService) UpdateRunStatusExecute(r ControllerServicesApiApiUpdateRunStatusRequest) (ControllerServiceEntity, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ControllerServiceEntity
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ControllerServicesApiService.UpdateRunStatus")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/controller-services/{id}/run-status"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.body
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
