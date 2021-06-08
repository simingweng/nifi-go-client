# VariableRegistryUpdateStepDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Explanation of what happens in this step | [optional] [readonly] 
**Complete** | Pointer to **bool** | Whether or not this step has completed | [optional] [readonly] 
**FailureReason** | Pointer to **string** | An explanation of why this step failed, or null if this step did not fail | [optional] [readonly] 

## Methods

### NewVariableRegistryUpdateStepDTO

`func NewVariableRegistryUpdateStepDTO() *VariableRegistryUpdateStepDTO`

NewVariableRegistryUpdateStepDTO instantiates a new VariableRegistryUpdateStepDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVariableRegistryUpdateStepDTOWithDefaults

`func NewVariableRegistryUpdateStepDTOWithDefaults() *VariableRegistryUpdateStepDTO`

NewVariableRegistryUpdateStepDTOWithDefaults instantiates a new VariableRegistryUpdateStepDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *VariableRegistryUpdateStepDTO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VariableRegistryUpdateStepDTO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VariableRegistryUpdateStepDTO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VariableRegistryUpdateStepDTO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetComplete

`func (o *VariableRegistryUpdateStepDTO) GetComplete() bool`

GetComplete returns the Complete field if non-nil, zero value otherwise.

### GetCompleteOk

`func (o *VariableRegistryUpdateStepDTO) GetCompleteOk() (*bool, bool)`

GetCompleteOk returns a tuple with the Complete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplete

`func (o *VariableRegistryUpdateStepDTO) SetComplete(v bool)`

SetComplete sets Complete field to given value.

### HasComplete

`func (o *VariableRegistryUpdateStepDTO) HasComplete() bool`

HasComplete returns a boolean if a field has been set.

### GetFailureReason

`func (o *VariableRegistryUpdateStepDTO) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *VariableRegistryUpdateStepDTO) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *VariableRegistryUpdateStepDTO) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.

### HasFailureReason

`func (o *VariableRegistryUpdateStepDTO) HasFailureReason() bool`

HasFailureReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


