# NodeDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodeId** | **string** | The id of the node. | [optional] [readonly] 
**Address** | **string** | The node&#39;s host/ip address. | [optional] [readonly] 
**ApiPort** | **int32** | The port the node is listening for API requests. | [optional] [readonly] 
**Status** | **string** | The node&#39;s status. | [optional] 
**Heartbeat** | **string** | the time of the nodes&#39;s last heartbeat. | [optional] [readonly] 
**ConnectionRequested** | **string** | The time of the node&#39;s last connection request. | [optional] [readonly] 
**Roles** | **[]string** | The roles of this node. | [optional] [readonly] 
**ActiveThreadCount** | **int32** | The active threads for the NiFi on the node. | [optional] [readonly] 
**Queued** | **string** | The queue the NiFi on the node. | [optional] [readonly] 
**Events** | [**[]NodeEventDto**](NodeEventDTO.md) | The node&#39;s events. | [optional] [readonly] 
**NodeStartTime** | **string** | The time at which this Node was last refreshed. | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


