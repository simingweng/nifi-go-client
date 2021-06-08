# ParameterContextUpdateRequestDTO

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
**UpdateSteps** | Pointer to [**[]ParameterContextUpdateStepDTO**](ParameterContextUpdateStepDTO.md) | The steps that are required in order to complete the request, along with the status of each | [optional] [readonly] 
**ParameterContext** | Pointer to [**ParameterContextDTO**](ParameterContextDTO.md) |  | [optional] 
**ReferencingComponents** | Pointer to [**[]AffectedComponentEntity**](AffectedComponentEntity.md) | The components that are referenced by the update. | [optional] [readonly] 

## Methods

### NewParameterContextUpdateRequestDTO

`func NewParameterContextUpdateRequestDTO() *ParameterContextUpdateRequestDTO`

NewParameterContextUpdateRequestDTO instantiates a new ParameterContextUpdateRequestDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParameterContextUpdateRequestDTOWithDefaults

`func NewParameterContextUpdateRequestDTOWithDefaults() *ParameterContextUpdateRequestDTO`

NewParameterContextUpdateRequestDTOWithDefaults instantiates a new ParameterContextUpdateRequestDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *ParameterContextUpdateRequestDTO) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ParameterContextUpdateRequestDTO) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ParameterContextUpdateRequestDTO) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *ParameterContextUpdateRequestDTO) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetUri

`func (o *ParameterContextUpdateRequestDTO) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *ParameterContextUpdateRequestDTO) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *ParameterContextUpdateRequestDTO) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *ParameterContextUpdateRequestDTO) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetSubmissionTime

`func (o *ParameterContextUpdateRequestDTO) GetSubmissionTime() time.Time`

GetSubmissionTime returns the SubmissionTime field if non-nil, zero value otherwise.

### GetSubmissionTimeOk

`func (o *ParameterContextUpdateRequestDTO) GetSubmissionTimeOk() (*time.Time, bool)`

GetSubmissionTimeOk returns a tuple with the SubmissionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmissionTime

`func (o *ParameterContextUpdateRequestDTO) SetSubmissionTime(v time.Time)`

SetSubmissionTime sets SubmissionTime field to given value.

### HasSubmissionTime

`func (o *ParameterContextUpdateRequestDTO) HasSubmissionTime() bool`

HasSubmissionTime returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ParameterContextUpdateRequestDTO) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ParameterContextUpdateRequestDTO) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ParameterContextUpdateRequestDTO) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ParameterContextUpdateRequestDTO) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetComplete

`func (o *ParameterContextUpdateRequestDTO) GetComplete() bool`

GetComplete returns the Complete field if non-nil, zero value otherwise.

### GetCompleteOk

`func (o *ParameterContextUpdateRequestDTO) GetCompleteOk() (*bool, bool)`

GetCompleteOk returns a tuple with the Complete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplete

`func (o *ParameterContextUpdateRequestDTO) SetComplete(v bool)`

SetComplete sets Complete field to given value.

### HasComplete

`func (o *ParameterContextUpdateRequestDTO) HasComplete() bool`

HasComplete returns a boolean if a field has been set.

### GetFailureReason

`func (o *ParameterContextUpdateRequestDTO) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *ParameterContextUpdateRequestDTO) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *ParameterContextUpdateRequestDTO) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.

### HasFailureReason

`func (o *ParameterContextUpdateRequestDTO) HasFailureReason() bool`

HasFailureReason returns a boolean if a field has been set.

### GetPercentCompleted

`func (o *ParameterContextUpdateRequestDTO) GetPercentCompleted() int32`

GetPercentCompleted returns the PercentCompleted field if non-nil, zero value otherwise.

### GetPercentCompletedOk

`func (o *ParameterContextUpdateRequestDTO) GetPercentCompletedOk() (*int32, bool)`

GetPercentCompletedOk returns a tuple with the PercentCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentCompleted

`func (o *ParameterContextUpdateRequestDTO) SetPercentCompleted(v int32)`

SetPercentCompleted sets PercentCompleted field to given value.

### HasPercentCompleted

`func (o *ParameterContextUpdateRequestDTO) HasPercentCompleted() bool`

HasPercentCompleted returns a boolean if a field has been set.

### GetState

`func (o *ParameterContextUpdateRequestDTO) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ParameterContextUpdateRequestDTO) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ParameterContextUpdateRequestDTO) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ParameterContextUpdateRequestDTO) HasState() bool`

HasState returns a boolean if a field has been set.

### GetUpdateSteps

`func (o *ParameterContextUpdateRequestDTO) GetUpdateSteps() []ParameterContextUpdateStepDTO`

GetUpdateSteps returns the UpdateSteps field if non-nil, zero value otherwise.

### GetUpdateStepsOk

`func (o *ParameterContextUpdateRequestDTO) GetUpdateStepsOk() (*[]ParameterContextUpdateStepDTO, bool)`

GetUpdateStepsOk returns a tuple with the UpdateSteps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateSteps

`func (o *ParameterContextUpdateRequestDTO) SetUpdateSteps(v []ParameterContextUpdateStepDTO)`

SetUpdateSteps sets UpdateSteps field to given value.

### HasUpdateSteps

`func (o *ParameterContextUpdateRequestDTO) HasUpdateSteps() bool`

HasUpdateSteps returns a boolean if a field has been set.

### GetParameterContext

`func (o *ParameterContextUpdateRequestDTO) GetParameterContext() ParameterContextDTO`

GetParameterContext returns the ParameterContext field if non-nil, zero value otherwise.

### GetParameterContextOk

`func (o *ParameterContextUpdateRequestDTO) GetParameterContextOk() (*ParameterContextDTO, bool)`

GetParameterContextOk returns a tuple with the ParameterContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameterContext

`func (o *ParameterContextUpdateRequestDTO) SetParameterContext(v ParameterContextDTO)`

SetParameterContext sets ParameterContext field to given value.

### HasParameterContext

`func (o *ParameterContextUpdateRequestDTO) HasParameterContext() bool`

HasParameterContext returns a boolean if a field has been set.

### GetReferencingComponents

`func (o *ParameterContextUpdateRequestDTO) GetReferencingComponents() []AffectedComponentEntity`

GetReferencingComponents returns the ReferencingComponents field if non-nil, zero value otherwise.

### GetReferencingComponentsOk

`func (o *ParameterContextUpdateRequestDTO) GetReferencingComponentsOk() (*[]AffectedComponentEntity, bool)`

GetReferencingComponentsOk returns a tuple with the ReferencingComponents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencingComponents

`func (o *ParameterContextUpdateRequestDTO) SetReferencingComponents(v []AffectedComponentEntity)`

SetReferencingComponents sets ReferencingComponents field to given value.

### HasReferencingComponents

`func (o *ParameterContextUpdateRequestDTO) HasReferencingComponents() bool`

HasReferencingComponents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


