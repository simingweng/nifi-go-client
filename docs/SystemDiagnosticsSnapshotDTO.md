# SystemDiagnosticsSnapshotDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalNonHeap** | **string** | Total size of non heap. | [optional] 
**TotalNonHeapBytes** | **int64** | Total number of bytes allocated to the JVM not used for heap | [optional] 
**UsedNonHeap** | **string** | Amount of use non heap. | [optional] 
**UsedNonHeapBytes** | **int64** | Total number of bytes used by the JVM not in the heap space | [optional] 
**FreeNonHeap** | **string** | Amount of free non heap. | [optional] 
**FreeNonHeapBytes** | **int64** | Total number of free non-heap bytes available to the JVM | [optional] 
**MaxNonHeap** | **string** | Maximum size of non heap. | [optional] 
**MaxNonHeapBytes** | **int64** | The maximum number of bytes that the JVM can use for non-heap purposes | [optional] 
**NonHeapUtilization** | **string** | Utilization of non heap. | [optional] 
**TotalHeap** | **string** | Total size of heap. | [optional] 
**TotalHeapBytes** | **int64** | The total number of bytes that are available for the JVM heap to use | [optional] 
**UsedHeap** | **string** | Amount of used heap. | [optional] 
**UsedHeapBytes** | **int64** | The number of bytes of JVM heap that are currently being used | [optional] 
**FreeHeap** | **string** | Amount of free heap. | [optional] 
**FreeHeapBytes** | **int64** | The number of bytes that are allocated to the JVM heap but not currently being used | [optional] 
**MaxHeap** | **string** | Maximum size of heap. | [optional] 
**MaxHeapBytes** | **int64** | The maximum number of bytes that can be used by the JVM | [optional] 
**HeapUtilization** | **string** | Utilization of heap. | [optional] 
**AvailableProcessors** | **int32** | Number of available processors if supported by the underlying system. | [optional] 
**ProcessorLoadAverage** | **float64** | The processor load average if supported by the underlying system. | [optional] 
**TotalThreads** | **int32** | Total number of threads. | [optional] 
**DaemonThreads** | **int32** | Number of daemon threads. | [optional] 
**Uptime** | **string** | The uptime of the Java virtual machine | [optional] 
**FlowFileRepositoryStorageUsage** | [**StorageUsageDto**](StorageUsageDTO.md) |  | [optional] 
**ContentRepositoryStorageUsage** | [**[]StorageUsageDto**](StorageUsageDTO.md) | The content repository storage usage. | [optional] 
**ProvenanceRepositoryStorageUsage** | [**[]StorageUsageDto**](StorageUsageDTO.md) | The provenance repository storage usage. | [optional] 
**GarbageCollection** | [**[]GarbageCollectionDto**](GarbageCollectionDTO.md) | The garbage collection details. | [optional] 
**StatsLastRefreshed** | **string** | When the diagnostics were generated. | [optional] 
**VersionInfo** | [**VersionInfoDto**](VersionInfoDTO.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


