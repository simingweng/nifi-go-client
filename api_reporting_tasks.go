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

// ReportingTasksApiService ReportingTasksApi service
type ReportingTasksApiService service

type ReportingTasksApiApiClearStateRequest struct {
	ctx        _context.Context
	ApiService *ReportingTasksApiService
	id         string
}

func (r ReportingTasksApiApiClearStateRequest) Execute() (ComponentStateEntity, *_nethttp.Response, error) {
	return r.ApiService.ClearStateExecute(r)
}

/*
 * ClearState Clears the state for a reporting task
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id The reporting task id.
 * @return ReportingTasksApiApiClearStateRequest
 */
func (a *ReportingTasksApiService) ClearState(ctx _context.Context, id string) ReportingTasksApiApiClearStateRequest {
	return ReportingTasksApiApiClearStateRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

/*
 * Execute executes the request
 * @return ComponentStateEntity
 */
func (a *ReportingTasksApiService) ClearStateExecute(r ReportingTasksApiApiClearStateRequest) (ComponentStateEntity, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ComponentStateEntity
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReportingTasksApiService.ClearState")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/reporting-tasks/{id}/state/clear-requests"
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

type ReportingTasksApiApiGetPropertyDescriptorRequest struct {
	ctx          _context.Context
	ApiService   *ReportingTasksApiService
	id           string
	propertyName *string
}

func (r ReportingTasksApiApiGetPropertyDescriptorRequest) PropertyName(propertyName string) ReportingTasksApiApiGetPropertyDescriptorRequest {
	r.propertyName = &propertyName
	return r
}

func (r ReportingTasksApiApiGetPropertyDescriptorRequest) Execute() (PropertyDescriptorEntity, *_nethttp.Response, error) {
	return r.ApiService.GetPropertyDescriptorExecute(r)
}

/*
 * GetPropertyDescriptor Gets a reporting task property descriptor
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id The reporting task id.
 * @return ReportingTasksApiApiGetPropertyDescriptorRequest
 */
func (a *ReportingTasksApiService) GetPropertyDescriptor(ctx _context.Context, id string) ReportingTasksApiApiGetPropertyDescriptorRequest {
	return ReportingTasksApiApiGetPropertyDescriptorRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

/*
 * Execute executes the request
 * @return PropertyDescriptorEntity
 */
func (a *ReportingTasksApiService) GetPropertyDescriptorExecute(r ReportingTasksApiApiGetPropertyDescriptorRequest) (PropertyDescriptorEntity, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  PropertyDescriptorEntity
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReportingTasksApiService.GetPropertyDescriptor")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/reporting-tasks/{id}/descriptors"
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

type ReportingTasksApiApiGetReportingTaskRequest struct {
	ctx        _context.Context
	ApiService *ReportingTasksApiService
	id         string
}

func (r ReportingTasksApiApiGetReportingTaskRequest) Execute() (ReportingTaskEntity, *_nethttp.Response, error) {
	return r.ApiService.GetReportingTaskExecute(r)
}

/*
 * GetReportingTask Gets a reporting task
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id The reporting task id.
 * @return ReportingTasksApiApiGetReportingTaskRequest
 */
func (a *ReportingTasksApiService) GetReportingTask(ctx _context.Context, id string) ReportingTasksApiApiGetReportingTaskRequest {
	return ReportingTasksApiApiGetReportingTaskRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

/*
 * Execute executes the request
 * @return ReportingTaskEntity
 */
func (a *ReportingTasksApiService) GetReportingTaskExecute(r ReportingTasksApiApiGetReportingTaskRequest) (ReportingTaskEntity, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ReportingTaskEntity
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReportingTasksApiService.GetReportingTask")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/reporting-tasks/{id}"
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

type ReportingTasksApiApiGetStateRequest struct {
	ctx        _context.Context
	ApiService *ReportingTasksApiService
	id         string
}

func (r ReportingTasksApiApiGetStateRequest) Execute() (ComponentStateEntity, *_nethttp.Response, error) {
	return r.ApiService.GetStateExecute(r)
}

/*
 * GetState Gets the state for a reporting task
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id The reporting task id.
 * @return ReportingTasksApiApiGetStateRequest
 */
func (a *ReportingTasksApiService) GetState(ctx _context.Context, id string) ReportingTasksApiApiGetStateRequest {
	return ReportingTasksApiApiGetStateRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

/*
 * Execute executes the request
 * @return ComponentStateEntity
 */
func (a *ReportingTasksApiService) GetStateExecute(r ReportingTasksApiApiGetStateRequest) (ComponentStateEntity, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ComponentStateEntity
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReportingTasksApiService.GetState")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/reporting-tasks/{id}/state"
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

type ReportingTasksApiApiRemoveReportingTaskRequest struct {
	ctx                          _context.Context
	ApiService                   *ReportingTasksApiService
	id                           string
	version                      *string
	clientId                     *string
	disconnectedNodeAcknowledged *bool
}

func (r ReportingTasksApiApiRemoveReportingTaskRequest) Version(version string) ReportingTasksApiApiRemoveReportingTaskRequest {
	r.version = &version
	return r
}
func (r ReportingTasksApiApiRemoveReportingTaskRequest) ClientId(clientId string) ReportingTasksApiApiRemoveReportingTaskRequest {
	r.clientId = &clientId
	return r
}
func (r ReportingTasksApiApiRemoveReportingTaskRequest) DisconnectedNodeAcknowledged(disconnectedNodeAcknowledged bool) ReportingTasksApiApiRemoveReportingTaskRequest {
	r.disconnectedNodeAcknowledged = &disconnectedNodeAcknowledged
	return r
}

func (r ReportingTasksApiApiRemoveReportingTaskRequest) Execute() (ReportingTaskEntity, *_nethttp.Response, error) {
	return r.ApiService.RemoveReportingTaskExecute(r)
}

/*
 * RemoveReportingTask Deletes a reporting task
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id The reporting task id.
 * @return ReportingTasksApiApiRemoveReportingTaskRequest
 */
func (a *ReportingTasksApiService) RemoveReportingTask(ctx _context.Context, id string) ReportingTasksApiApiRemoveReportingTaskRequest {
	return ReportingTasksApiApiRemoveReportingTaskRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

/*
 * Execute executes the request
 * @return ReportingTaskEntity
 */
func (a *ReportingTasksApiService) RemoveReportingTaskExecute(r ReportingTasksApiApiRemoveReportingTaskRequest) (ReportingTaskEntity, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ReportingTaskEntity
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReportingTasksApiService.RemoveReportingTask")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/reporting-tasks/{id}"
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

type ReportingTasksApiApiUpdateReportingTaskRequest struct {
	ctx        _context.Context
	ApiService *ReportingTasksApiService
	id         string
	body       *ReportingTaskEntity
}

func (r ReportingTasksApiApiUpdateReportingTaskRequest) Body(body ReportingTaskEntity) ReportingTasksApiApiUpdateReportingTaskRequest {
	r.body = &body
	return r
}

func (r ReportingTasksApiApiUpdateReportingTaskRequest) Execute() (ReportingTaskEntity, *_nethttp.Response, error) {
	return r.ApiService.UpdateReportingTaskExecute(r)
}

/*
 * UpdateReportingTask Updates a reporting task
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id The reporting task id.
 * @return ReportingTasksApiApiUpdateReportingTaskRequest
 */
func (a *ReportingTasksApiService) UpdateReportingTask(ctx _context.Context, id string) ReportingTasksApiApiUpdateReportingTaskRequest {
	return ReportingTasksApiApiUpdateReportingTaskRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

/*
 * Execute executes the request
 * @return ReportingTaskEntity
 */
func (a *ReportingTasksApiService) UpdateReportingTaskExecute(r ReportingTasksApiApiUpdateReportingTaskRequest) (ReportingTaskEntity, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ReportingTaskEntity
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReportingTasksApiService.UpdateReportingTask")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/reporting-tasks/{id}"
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

type ReportingTasksApiApiUpdateRunStatusRequest struct {
	ctx        _context.Context
	ApiService *ReportingTasksApiService
	id         string
	body       *ReportingTaskRunStatusEntity
}

func (r ReportingTasksApiApiUpdateRunStatusRequest) Body(body ReportingTaskRunStatusEntity) ReportingTasksApiApiUpdateRunStatusRequest {
	r.body = &body
	return r
}

func (r ReportingTasksApiApiUpdateRunStatusRequest) Execute() (ReportingTaskEntity, *_nethttp.Response, error) {
	return r.ApiService.UpdateRunStatusExecute(r)
}

/*
 * UpdateRunStatus Updates run status of a reporting task
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id The reporting task id.
 * @return ReportingTasksApiApiUpdateRunStatusRequest
 */
func (a *ReportingTasksApiService) UpdateRunStatus(ctx _context.Context, id string) ReportingTasksApiApiUpdateRunStatusRequest {
	return ReportingTasksApiApiUpdateRunStatusRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

/*
 * Execute executes the request
 * @return ReportingTaskEntity
 */
func (a *ReportingTasksApiService) UpdateRunStatusExecute(r ReportingTasksApiApiUpdateRunStatusRequest) (ReportingTaskEntity, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ReportingTaskEntity
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReportingTasksApiService.UpdateRunStatus")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/reporting-tasks/{id}/run-status"
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
