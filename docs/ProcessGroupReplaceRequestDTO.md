# ProcessGroupReplaceRequestDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | Pointer to **string** | The unique ID of this request. | [optional] [readonly] 
**ProcessGroupId** | Pointer to **string** | The unique ID of the Process Group being updated | [optional] 
**Uri** | Pointer to **string** | The URI for future requests to this drop request. | [optional] [readonly] 
**LastUpdated** | Pointer to **string** | The last time this request was updated. | [optional] [readonly] 
**Complete** | Pointer to **bool** | Whether or not this request has completed | [optional] [readonly] 
**FailureReason** | Pointer to **string** | An explanation of why this request failed, or null if this request has not failed | [optional] [readonly] 
**PercentCompleted** | Pointer to **int32** | The percentage complete for the request, between 0 and 100 | [optional] [readonly] 
**State** | Pointer to **string** | The state of the request | [optional] [readonly] 

## Methods

### NewProcessGroupReplaceRequestDTO

`func NewProcessGroupReplaceRequestDTO() *ProcessGroupReplaceRequestDTO`

NewProcessGroupReplaceRequestDTO instantiates a new ProcessGroupReplaceRequestDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessGroupReplaceRequestDTOWithDefaults

`func NewProcessGroupReplaceRequestDTOWithDefaults() *ProcessGroupReplaceRequestDTO`

NewProcessGroupReplaceRequestDTOWithDefaults instantiates a new ProcessGroupReplaceRequestDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *ProcessGroupReplaceRequestDTO) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ProcessGroupReplaceRequestDTO) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ProcessGroupReplaceRequestDTO) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *ProcessGroupReplaceRequestDTO) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetProcessGroupId

`func (o *ProcessGroupReplaceRequestDTO) GetProcessGroupId() string`

GetProcessGroupId returns the ProcessGroupId field if non-nil, zero value otherwise.

### GetProcessGroupIdOk

`func (o *ProcessGroupReplaceRequestDTO) GetProcessGroupIdOk() (*string, bool)`

GetProcessGroupIdOk returns a tuple with the ProcessGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessGroupId

`func (o *ProcessGroupReplaceRequestDTO) SetProcessGroupId(v string)`

SetProcessGroupId sets ProcessGroupId field to given value.

### HasProcessGroupId

`func (o *ProcessGroupReplaceRequestDTO) HasProcessGroupId() bool`

HasProcessGroupId returns a boolean if a field has been set.

### GetUri

`func (o *ProcessGroupReplaceRequestDTO) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *ProcessGroupReplaceRequestDTO) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *ProcessGroupReplaceRequestDTO) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *ProcessGroupReplaceRequestDTO) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ProcessGroupReplaceRequestDTO) GetLastUpdated() string`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ProcessGroupReplaceRequestDTO) GetLastUpdatedOk() (*string, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ProcessGroupReplaceRequestDTO) SetLastUpdated(v string)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ProcessGroupReplaceRequestDTO) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetComplete

`func (o *ProcessGroupReplaceRequestDTO) GetComplete() bool`

GetComplete returns the Complete field if non-nil, zero value otherwise.

### GetCompleteOk

`func (o *ProcessGroupReplaceRequestDTO) GetCompleteOk() (*bool, bool)`

GetCompleteOk returns a tuple with the Complete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplete

`func (o *ProcessGroupReplaceRequestDTO) SetComplete(v bool)`

SetComplete sets Complete field to given value.

### HasComplete

`func (o *ProcessGroupReplaceRequestDTO) HasComplete() bool`

HasComplete returns a boolean if a field has been set.

### GetFailureReason

`func (o *ProcessGroupReplaceRequestDTO) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *ProcessGroupReplaceRequestDTO) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *ProcessGroupReplaceRequestDTO) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.

### HasFailureReason

`func (o *ProcessGroupReplaceRequestDTO) HasFailureReason() bool`

HasFailureReason returns a boolean if a field has been set.

### GetPercentCompleted

`func (o *ProcessGroupReplaceRequestDTO) GetPercentCompleted() int32`

GetPercentCompleted returns the PercentCompleted field if non-nil, zero value otherwise.

### GetPercentCompletedOk

`func (o *ProcessGroupReplaceRequestDTO) GetPercentCompletedOk() (*int32, bool)`

GetPercentCompletedOk returns a tuple with the PercentCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentCompleted

`func (o *ProcessGroupReplaceRequestDTO) SetPercentCompleted(v int32)`

SetPercentCompleted sets PercentCompleted field to given value.

### HasPercentCompleted

`func (o *ProcessGroupReplaceRequestDTO) HasPercentCompleted() bool`

HasPercentCompleted returns a boolean if a field has been set.

### GetState

`func (o *ProcessGroupReplaceRequestDTO) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ProcessGroupReplaceRequestDTO) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ProcessGroupReplaceRequestDTO) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ProcessGroupReplaceRequestDTO) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


