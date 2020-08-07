# ParameterContextUpdateRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | The ID of the request | [optional] [readonly] 
**Uri** | **string** | The URI for the request | [optional] [readonly] 
**SubmissionTime** | [**time.Time**](time.Time.md) | The timestamp of when the request was submitted | [optional] [readonly] 
**LastUpdated** | [**time.Time**](time.Time.md) | The timestamp of when the request was last updated | [optional] [readonly] 
**Complete** | **bool** | Whether or not the request is completed | [optional] [readonly] 
**FailureReason** | **string** | The reason for the request failing, or null if the request has not failed | [optional] [readonly] 
**PercentCompleted** | **int32** | A value between 0 and 100 (inclusive) indicating how close the request is to completion | [optional] [readonly] 
**State** | **string** | A description of the current state of the request | [optional] [readonly] 
**UpdateSteps** | [**[]ParameterContextUpdateStepDto**](ParameterContextUpdateStepDTO.md) | The steps that are required in order to complete the request, along with the status of each | [optional] [readonly] 
**ParameterContext** | [**ParameterContextDto**](ParameterContextDTO.md) |  | [optional] 
**ReferencingComponents** | [**[]AffectedComponentEntity**](AffectedComponentEntity.md) | The components that are referenced by the update. | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


