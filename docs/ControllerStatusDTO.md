# ControllerStatusDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveThreadCount** | **int32** | The number of active threads in the NiFi. | [optional] 
**TerminatedThreadCount** | **int32** | The number of terminated threads in the NiFi. | [optional] 
**Queued** | **string** | The number of flowfiles queued in the NiFi. | [optional] 
**FlowFilesQueued** | **int32** | The number of FlowFiles queued across the entire flow | [optional] 
**BytesQueued** | **int64** | The size of the FlowFiles queued across the entire flow | [optional] 
**RunningCount** | **int32** | The number of running components in the NiFi. | [optional] 
**StoppedCount** | **int32** | The number of stopped components in the NiFi. | [optional] 
**InvalidCount** | **int32** | The number of invalid components in the NiFi. | [optional] 
**DisabledCount** | **int32** | The number of disabled components in the NiFi. | [optional] 
**ActiveRemotePortCount** | **int32** | The number of active remote ports in the NiFi. | [optional] 
**InactiveRemotePortCount** | **int32** | The number of inactive remote ports in the NiFi. | [optional] 
**UpToDateCount** | **int32** | The number of up to date versioned process groups in the NiFi. | [optional] 
**LocallyModifiedCount** | **int32** | The number of locally modified versioned process groups in the NiFi. | [optional] 
**StaleCount** | **int32** | The number of stale versioned process groups in the NiFi. | [optional] 
**LocallyModifiedAndStaleCount** | **int32** | The number of locally modified and stale versioned process groups in the NiFi. | [optional] 
**SyncFailureCount** | **int32** | The number of versioned process groups in the NiFi that are unable to sync to a registry. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


