# ParameterContextValidationStepDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Explanation of what happens in this step | [optional] [readonly] 
**Complete** | Pointer to **bool** | Whether or not this step has completed | [optional] [readonly] 
**FailureReason** | Pointer to **string** | An explanation of why this step failed, or null if this step did not fail | [optional] [readonly] 

## Methods

### NewParameterContextValidationStepDTO

`func NewParameterContextValidationStepDTO() *ParameterContextValidationStepDTO`

NewParameterContextValidationStepDTO instantiates a new ParameterContextValidationStepDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParameterContextValidationStepDTOWithDefaults

`func NewParameterContextValidationStepDTOWithDefaults() *ParameterContextValidationStepDTO`

NewParameterContextValidationStepDTOWithDefaults instantiates a new ParameterContextValidationStepDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ParameterContextValidationStepDTO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ParameterContextValidationStepDTO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ParameterContextValidationStepDTO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ParameterContextValidationStepDTO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetComplete

`func (o *ParameterContextValidationStepDTO) GetComplete() bool`

GetComplete returns the Complete field if non-nil, zero value otherwise.

### GetCompleteOk

`func (o *ParameterContextValidationStepDTO) GetCompleteOk() (*bool, bool)`

GetCompleteOk returns a tuple with the Complete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplete

`func (o *ParameterContextValidationStepDTO) SetComplete(v bool)`

SetComplete sets Complete field to given value.

### HasComplete

`func (o *ParameterContextValidationStepDTO) HasComplete() bool`

HasComplete returns a boolean if a field has been set.

### GetFailureReason

`func (o *ParameterContextValidationStepDTO) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *ParameterContextValidationStepDTO) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *ParameterContextValidationStepDTO) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.

### HasFailureReason

`func (o *ParameterContextValidationStepDTO) HasFailureReason() bool`

HasFailureReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


