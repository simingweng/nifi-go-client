# ParameterContextValidationRequestDTO

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
**UpdateSteps** | Pointer to [**[]ParameterContextValidationStepDTO**](ParameterContextValidationStepDTO.md) | The steps that are required in order to complete the request, along with the status of each | [optional] [readonly] 
**ParameterContext** | Pointer to [**ParameterContextDTO**](ParameterContextDTO.md) |  | [optional] 
**ComponentValidationResults** | Pointer to [**ComponentValidationResultsEntity**](ComponentValidationResultsEntity.md) |  | [optional] 

## Methods

### NewParameterContextValidationRequestDTO

`func NewParameterContextValidationRequestDTO() *ParameterContextValidationRequestDTO`

NewParameterContextValidationRequestDTO instantiates a new ParameterContextValidationRequestDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParameterContextValidationRequestDTOWithDefaults

`func NewParameterContextValidationRequestDTOWithDefaults() *ParameterContextValidationRequestDTO`

NewParameterContextValidationRequestDTOWithDefaults instantiates a new ParameterContextValidationRequestDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *ParameterContextValidationRequestDTO) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ParameterContextValidationRequestDTO) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ParameterContextValidationRequestDTO) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *ParameterContextValidationRequestDTO) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetUri

`func (o *ParameterContextValidationRequestDTO) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *ParameterContextValidationRequestDTO) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *ParameterContextValidationRequestDTO) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *ParameterContextValidationRequestDTO) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetSubmissionTime

`func (o *ParameterContextValidationRequestDTO) GetSubmissionTime() time.Time`

GetSubmissionTime returns the SubmissionTime field if non-nil, zero value otherwise.

### GetSubmissionTimeOk

`func (o *ParameterContextValidationRequestDTO) GetSubmissionTimeOk() (*time.Time, bool)`

GetSubmissionTimeOk returns a tuple with the SubmissionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmissionTime

`func (o *ParameterContextValidationRequestDTO) SetSubmissionTime(v time.Time)`

SetSubmissionTime sets SubmissionTime field to given value.

### HasSubmissionTime

`func (o *ParameterContextValidationRequestDTO) HasSubmissionTime() bool`

HasSubmissionTime returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ParameterContextValidationRequestDTO) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ParameterContextValidationRequestDTO) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ParameterContextValidationRequestDTO) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ParameterContextValidationRequestDTO) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetComplete

`func (o *ParameterContextValidationRequestDTO) GetComplete() bool`

GetComplete returns the Complete field if non-nil, zero value otherwise.

### GetCompleteOk

`func (o *ParameterContextValidationRequestDTO) GetCompleteOk() (*bool, bool)`

GetCompleteOk returns a tuple with the Complete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplete

`func (o *ParameterContextValidationRequestDTO) SetComplete(v bool)`

SetComplete sets Complete field to given value.

### HasComplete

`func (o *ParameterContextValidationRequestDTO) HasComplete() bool`

HasComplete returns a boolean if a field has been set.

### GetFailureReason

`func (o *ParameterContextValidationRequestDTO) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *ParameterContextValidationRequestDTO) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *ParameterContextValidationRequestDTO) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.

### HasFailureReason

`func (o *ParameterContextValidationRequestDTO) HasFailureReason() bool`

HasFailureReason returns a boolean if a field has been set.

### GetPercentCompleted

`func (o *ParameterContextValidationRequestDTO) GetPercentCompleted() int32`

GetPercentCompleted returns the PercentCompleted field if non-nil, zero value otherwise.

### GetPercentCompletedOk

`func (o *ParameterContextValidationRequestDTO) GetPercentCompletedOk() (*int32, bool)`

GetPercentCompletedOk returns a tuple with the PercentCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentCompleted

`func (o *ParameterContextValidationRequestDTO) SetPercentCompleted(v int32)`

SetPercentCompleted sets PercentCompleted field to given value.

### HasPercentCompleted

`func (o *ParameterContextValidationRequestDTO) HasPercentCompleted() bool`

HasPercentCompleted returns a boolean if a field has been set.

### GetState

`func (o *ParameterContextValidationRequestDTO) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ParameterContextValidationRequestDTO) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ParameterContextValidationRequestDTO) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ParameterContextValidationRequestDTO) HasState() bool`

HasState returns a boolean if a field has been set.

### GetUpdateSteps

`func (o *ParameterContextValidationRequestDTO) GetUpdateSteps() []ParameterContextValidationStepDTO`

GetUpdateSteps returns the UpdateSteps field if non-nil, zero value otherwise.

### GetUpdateStepsOk

`func (o *ParameterContextValidationRequestDTO) GetUpdateStepsOk() (*[]ParameterContextValidationStepDTO, bool)`

GetUpdateStepsOk returns a tuple with the UpdateSteps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateSteps

`func (o *ParameterContextValidationRequestDTO) SetUpdateSteps(v []ParameterContextValidationStepDTO)`

SetUpdateSteps sets UpdateSteps field to given value.

### HasUpdateSteps

`func (o *ParameterContextValidationRequestDTO) HasUpdateSteps() bool`

HasUpdateSteps returns a boolean if a field has been set.

### GetParameterContext

`func (o *ParameterContextValidationRequestDTO) GetParameterContext() ParameterContextDTO`

GetParameterContext returns the ParameterContext field if non-nil, zero value otherwise.

### GetParameterContextOk

`func (o *ParameterContextValidationRequestDTO) GetParameterContextOk() (*ParameterContextDTO, bool)`

GetParameterContextOk returns a tuple with the ParameterContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameterContext

`func (o *ParameterContextValidationRequestDTO) SetParameterContext(v ParameterContextDTO)`

SetParameterContext sets ParameterContext field to given value.

### HasParameterContext

`func (o *ParameterContextValidationRequestDTO) HasParameterContext() bool`

HasParameterContext returns a boolean if a field has been set.

### GetComponentValidationResults

`func (o *ParameterContextValidationRequestDTO) GetComponentValidationResults() ComponentValidationResultsEntity`

GetComponentValidationResults returns the ComponentValidationResults field if non-nil, zero value otherwise.

### GetComponentValidationResultsOk

`func (o *ParameterContextValidationRequestDTO) GetComponentValidationResultsOk() (*ComponentValidationResultsEntity, bool)`

GetComponentValidationResultsOk returns a tuple with the ComponentValidationResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentValidationResults

`func (o *ParameterContextValidationRequestDTO) SetComponentValidationResults(v ComponentValidationResultsEntity)`

SetComponentValidationResults sets ComponentValidationResults field to given value.

### HasComponentValidationResults

`func (o *ParameterContextValidationRequestDTO) HasComponentValidationResults() bool`

HasComponentValidationResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


