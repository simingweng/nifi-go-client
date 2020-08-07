# ProvenanceNodeDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id of the node. | [optional] 
**FlowFileUuid** | **string** | The uuid of the flowfile associated with the provenance event. | [optional] 
**ParentUuids** | **[]string** | The uuid of the parent flowfiles of the provenance event. | [optional] 
**ChildUuids** | **[]string** | The uuid of the childrent flowfiles of the provenance event. | [optional] 
**ClusterNodeIdentifier** | **string** | The identifier of the node that this event/flowfile originated from. | [optional] 
**Type** | **string** | The type of the node. | [optional] 
**EventType** | **string** | If the type is EVENT, this is the type of event. | [optional] 
**Millis** | **int64** | The timestamp of the node in milliseconds. | [optional] 
**Timestamp** | **string** | The timestamp of the node formatted. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


