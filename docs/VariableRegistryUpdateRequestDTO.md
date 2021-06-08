# VariableRegistryUpdateRequestDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | Pointer to **string** | The ID of the request | [optional] [readonly] 
**Uri** | Pointer to **string** | The URI for the request | [optional] [readonly] 
**SubmissionTime** | Pointer to **time.Time** | The timestamp of when the request was submitted | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** | The timestamp of when the request was last updated | [optional] [readonly] 
**Complete** | Pointer to **bool** | Whether or not the request is completed | [optional] [readonly] 
**FailureReason** | Pointer to **string** | The reason for the request failing, or null if the request has not failed | [optional] [readonly] 
**PercentCompleted** | Pointer to **int32** | A value between 0 and 100 (inclusive) indicating how close the request is to completion | [optional] [readonly] 
**State** | Pointer to **string** | A description of the current state of the request | [optional] [readonly] 
**UpdateSteps** | Pointer to [**[]VariableRegistryUpdateStepDTO**](VariableRegistryUpdateStepDTO.md) | The steps that are required in order to complete the request, along with the status of each | [optional] [readonly] 
**ProcessGroupId** | Pointer to **string** | The unique ID of the Process Group that the variable registry belongs to | [optional] 
**AffectedComponents** | Pointer to [**[]AffectedComponentEntity**](AffectedComponentEntity.md) | A set of all components that will be affected if the value of this variable is changed | [optional] [readonly] 

## Methods

### NewVariableRegistryUpdateRequestDTO

`func NewVariableRegistryUpdateRequestDTO() *VariableRegistryUpdateRequestDTO`

NewVariableRegistryUpdateRequestDTO instantiates a new VariableRegistryUpdateRequestDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVariableRegistryUpdateRequestDTOWithDefaults

`func NewVariableRegistryUpdateRequestDTOWithDefaults() *VariableRegistryUpdateRequestDTO`

NewVariableRegistryUpdateRequestDTOWithDefaults instantiates a new VariableRegistryUpdateRequestDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *VariableRegistryUpdateRequestDTO) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *VariableRegistryUpdateRequestDTO) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *VariableRegistryUpdateRequestDTO) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *VariableRegistryUpdateRequestDTO) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetUri

`func (o *VariableRegistryUpdateRequestDTO) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *VariableRegistryUpdateRequestDTO) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *VariableRegistryUpdateRequestDTO) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *VariableRegistryUpdateRequestDTO) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetSubmissionTime

`func (o *VariableRegistryUpdateRequestDTO) GetSubmissionTime() time.Time`

GetSubmissionTime returns the SubmissionTime field if non-nil, zero value otherwise.

### GetSubmissionTimeOk

`func (o *VariableRegistryUpdateRequestDTO) GetSubmissionTimeOk() (*time.Time, bool)`

GetSubmissionTimeOk returns a tuple with the SubmissionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmissionTime

`func (o *VariableRegistryUpdateRequestDTO) SetSubmissionTime(v time.Time)`

SetSubmissionTime sets SubmissionTime field to given value.

### HasSubmissionTime

`func (o *VariableRegistryUpdateRequestDTO) HasSubmissionTime() bool`

HasSubmissionTime returns a boolean if a field has been set.

### GetLastUpdated

`func (o *VariableRegistryUpdateRequestDTO) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *VariableRegistryUpdateRequestDTO) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *VariableRegistryUpdateRequestDTO) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *VariableRegistryUpdateRequestDTO) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetComplete

`func (o *VariableRegistryUpdateRequestDTO) GetComplete() bool`

GetComplete returns the Complete field if non-nil, zero value otherwise.

### GetCompleteOk

`func (o *VariableRegistryUpdateRequestDTO) GetCompleteOk() (*bool, bool)`

GetCompleteOk returns a tuple with the Complete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplete

`func (o *VariableRegistryUpdateRequestDTO) SetComplete(v bool)`

SetComplete sets Complete field to given value.

### HasComplete

`func (o *VariableRegistryUpdateRequestDTO) HasComplete() bool`

HasComplete returns a boolean if a field has been set.

### GetFailureReason

`func (o *VariableRegistryUpdateRequestDTO) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *VariableRegistryUpdateRequestDTO) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *VariableRegistryUpdateRequestDTO) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.

### HasFailureReason

`func (o *VariableRegistryUpdateRequestDTO) HasFailureReason() bool`

HasFailureReason returns a boolean if a field has been set.

### GetPercentCompleted

`func (o *VariableRegistryUpdateRequestDTO) GetPercentCompleted() int32`

GetPercentCompleted returns the PercentCompleted field if non-nil, zero value otherwise.

### GetPercentCompletedOk

`func (o *VariableRegistryUpdateRequestDTO) GetPercentCompletedOk() (*int32, bool)`

GetPercentCompletedOk returns a tuple with the PercentCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentCompleted

`func (o *VariableRegistryUpdateRequestDTO) SetPercentCompleted(v int32)`

SetPercentCompleted sets PercentCompleted field to given value.

### HasPercentCompleted

`func (o *VariableRegistryUpdateRequestDTO) HasPercentCompleted() bool`

HasPercentCompleted returns a boolean if a field has been set.

### GetState

`func (o *VariableRegistryUpdateRequestDTO) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *VariableRegistryUpdateRequestDTO) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *VariableRegistryUpdateRequestDTO) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *VariableRegistryUpdateRequestDTO) HasState() bool`

HasState returns a boolean if a field has been set.

### GetUpdateSteps

`func (o *VariableRegistryUpdateRequestDTO) GetUpdateSteps() []VariableRegistryUpdateStepDTO`

GetUpdateSteps returns the UpdateSteps field if non-nil, zero value otherwise.

### GetUpdateStepsOk

`func (o *VariableRegistryUpdateRequestDTO) GetUpdateStepsOk() (*[]VariableRegistryUpdateStepDTO, bool)`

GetUpdateStepsOk returns a tuple with the UpdateSteps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateSteps

`func (o *VariableRegistryUpdateRequestDTO) SetUpdateSteps(v []VariableRegistryUpdateStepDTO)`

SetUpdateSteps sets UpdateSteps field to given value.

### HasUpdateSteps

`func (o *VariableRegistryUpdateRequestDTO) HasUpdateSteps() bool`

HasUpdateSteps returns a boolean if a field has been set.

### GetProcessGroupId

`func (o *VariableRegistryUpdateRequestDTO) GetProcessGroupId() string`

GetProcessGroupId returns the ProcessGroupId field if non-nil, zero value otherwise.

### GetProcessGroupIdOk

`func (o *VariableRegistryUpdateRequestDTO) GetProcessGroupIdOk() (*string, bool)`

GetProcessGroupIdOk returns a tuple with the ProcessGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessGroupId

`func (o *VariableRegistryUpdateRequestDTO) SetProcessGroupId(v string)`

SetProcessGroupId sets ProcessGroupId field to given value.

### HasProcessGroupId

`func (o *VariableRegistryUpdateRequestDTO) HasProcessGroupId() bool`

HasProcessGroupId returns a boolean if a field has been set.

### GetAffectedComponents

`func (o *VariableRegistryUpdateRequestDTO) GetAffectedComponents() []AffectedComponentEntity`

GetAffectedComponents returns the AffectedComponents field if non-nil, zero value otherwise.

### GetAffectedComponentsOk

`func (o *VariableRegistryUpdateRequestDTO) GetAffectedComponentsOk() (*[]AffectedComponentEntity, bool)`

GetAffectedComponentsOk returns a tuple with the AffectedComponents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedComponents

`func (o *VariableRegistryUpdateRequestDTO) SetAffectedComponents(v []AffectedComponentEntity)`

SetAffectedComponents sets AffectedComponents field to given value.

### HasAffectedComponents

`func (o *VariableRegistryUpdateRequestDTO) HasAffectedComponents() bool`

HasAffectedComponents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


