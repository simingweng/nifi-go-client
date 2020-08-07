# VersionedFlowUpdateRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | The unique ID of this request. | [optional] [readonly] 
**ProcessGroupId** | **string** | The unique ID of the Process Group that the variable registry belongs to | [optional] 
**Uri** | **string** | The URI for future requests to this drop request. | [optional] [readonly] 
**LastUpdated** | **string** | The last time this request was updated. | [optional] [readonly] 
**Complete** | **bool** | Whether or not this request has completed | [optional] [readonly] 
**FailureReason** | **string** | An explanation of why this request failed, or null if this request has not failed | [optional] [readonly] 
**PercentCompleted** | **int32** | The percentage complete for the request, between 0 and 100 | [optional] [readonly] 
**State** | **string** | The state of the request | [optional] [readonly] 
**VersionControlInformation** | [**VersionControlInformationDto**](VersionControlInformationDTO.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


