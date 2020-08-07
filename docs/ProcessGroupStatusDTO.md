# ProcessGroupStatusDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the Process Group | [optional] 
**Name** | **string** | The name of the Process Group | [optional] 
**StatsLastRefreshed** | **string** | The time the status for the process group was last refreshed. | [optional] 
**AggregateSnapshot** | [**ProcessGroupStatusSnapshotDto**](ProcessGroupStatusSnapshotDTO.md) |  | [optional] 
**NodeSnapshots** | [**[]NodeProcessGroupStatusSnapshotDto**](NodeProcessGroupStatusSnapshotDTO.md) | The status reported by each node in the cluster. If the NiFi instance is a standalone instance, rather than a clustered instance, this value may be null. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


