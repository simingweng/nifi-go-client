# VersionedFlowSnapshotMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Link** | [**JaxbLink**](JaxbLink.md) |  | [optional] 
**BucketIdentifier** | **string** | The identifier of the bucket this snapshot belongs to. | 
**FlowIdentifier** | **string** | The identifier of the flow this snapshot belongs to. | 
**Version** | **int32** | The version of this snapshot of the flow. | 
**Timestamp** | **int64** | The timestamp when the flow was saved, as milliseconds since epoch. | [optional] [readonly] 
**Author** | **string** | The user that created this snapshot of the flow. | [optional] [readonly] 
**Comments** | **string** | The comments provided by the user when creating the snapshot. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


