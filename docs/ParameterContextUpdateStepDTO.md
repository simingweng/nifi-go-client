# ParameterContextUpdateStepDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Explanation of what happens in this step | [optional] [readonly] 
**Complete** | Pointer to **bool** | Whether or not this step has completed | [optional] [readonly] 
**FailureReason** | Pointer to **string** | An explanation of why this step failed, or null if this step did not fail | [optional] [readonly] 

## Methods

### NewParameterContextUpdateStepDTO

`func NewParameterContextUpdateStepDTO() *ParameterContextUpdateStepDTO`

NewParameterContextUpdateStepDTO instantiates a new ParameterContextUpdateStepDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParameterContextUpdateStepDTOWithDefaults

`func NewParameterContextUpdateStepDTOWithDefaults() *ParameterContextUpdateStepDTO`

NewParameterContextUpdateStepDTOWithDefaults instantiates a new ParameterContextUpdateStepDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ParameterContextUpdateStepDTO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ParameterContextUpdateStepDTO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ParameterContextUpdateStepDTO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ParameterContextUpdateStepDTO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetComplete

`func (o *ParameterContextUpdateStepDTO) GetComplete() bool`

GetComplete returns the Complete field if non-nil, zero value otherwise.

### GetCompleteOk

`func (o *ParameterContextUpdateStepDTO) GetCompleteOk() (*bool, bool)`

GetCompleteOk returns a tuple with the Complete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplete

`func (o *ParameterContextUpdateStepDTO) SetComplete(v bool)`

SetComplete sets Complete field to given value.

### HasComplete

`func (o *ParameterContextUpdateStepDTO) HasComplete() bool`

HasComplete returns a boolean if a field has been set.

### GetFailureReason

`func (o *ParameterContextUpdateStepDTO) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *ParameterContextUpdateStepDTO) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *ParameterContextUpdateStepDTO) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.

### HasFailureReason

`func (o *ParameterContextUpdateStepDTO) HasFailureReason() bool`

HasFailureReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


