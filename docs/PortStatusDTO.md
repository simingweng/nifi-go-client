# PortStatusDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id of the port. | [optional] 
**GroupId** | **string** | The id of the parent process group of the port. | [optional] 
**Name** | **string** | The name of the port. | [optional] 
**Transmitting** | **bool** | Whether the port has incoming or outgoing connections to a remote NiFi. | [optional] 
**RunStatus** | **string** | The run status of the port. | [optional] 
**StatsLastRefreshed** | **string** | The time the status for the process group was last refreshed. | [optional] 
**AggregateSnapshot** | [**PortStatusSnapshotDto**](PortStatusSnapshotDTO.md) |  | [optional] 
**NodeSnapshots** | [**[]NodePortStatusSnapshotDto**](NodePortStatusSnapshotDTO.md) | A status snapshot for each node in the cluster. If the NiFi instance is a standalone instance, rather than a cluster, this may be null. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


