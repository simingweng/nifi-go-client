# SystemDiagnosticsSnapshotDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalNonHeap** | Pointer to **string** | Total size of non heap. | [optional] 
**TotalNonHeapBytes** | Pointer to **int64** | Total number of bytes allocated to the JVM not used for heap | [optional] 
**UsedNonHeap** | Pointer to **string** | Amount of use non heap. | [optional] 
**UsedNonHeapBytes** | Pointer to **int64** | Total number of bytes used by the JVM not in the heap space | [optional] 
**FreeNonHeap** | Pointer to **string** | Amount of free non heap. | [optional] 
**FreeNonHeapBytes** | Pointer to **int64** | Total number of free non-heap bytes available to the JVM | [optional] 
**MaxNonHeap** | Pointer to **string** | Maximum size of non heap. | [optional] 
**MaxNonHeapBytes** | Pointer to **int64** | The maximum number of bytes that the JVM can use for non-heap purposes | [optional] 
**NonHeapUtilization** | Pointer to **string** | Utilization of non heap. | [optional] 
**TotalHeap** | Pointer to **string** | Total size of heap. | [optional] 
**TotalHeapBytes** | Pointer to **int64** | The total number of bytes that are available for the JVM heap to use | [optional] 
**UsedHeap** | Pointer to **string** | Amount of used heap. | [optional] 
**UsedHeapBytes** | Pointer to **int64** | The number of bytes of JVM heap that are currently being used | [optional] 
**FreeHeap** | Pointer to **string** | Amount of free heap. | [optional] 
**FreeHeapBytes** | Pointer to **int64** | The number of bytes that are allocated to the JVM heap but not currently being used | [optional] 
**MaxHeap** | Pointer to **string** | Maximum size of heap. | [optional] 
**MaxHeapBytes** | Pointer to **int64** | The maximum number of bytes that can be used by the JVM | [optional] 
**HeapUtilization** | Pointer to **string** | Utilization of heap. | [optional] 
**AvailableProcessors** | Pointer to **int32** | Number of available processors if supported by the underlying system. | [optional] 
**ProcessorLoadAverage** | Pointer to **float64** | The processor load average if supported by the underlying system. | [optional] 
**TotalThreads** | Pointer to **int32** | Total number of threads. | [optional] 
**DaemonThreads** | Pointer to **int32** | Number of daemon threads. | [optional] 
**Uptime** | Pointer to **string** | The uptime of the Java virtual machine | [optional] 
**FlowFileRepositoryStorageUsage** | Pointer to [**StorageUsageDTO**](StorageUsageDTO.md) |  | [optional] 
**ContentRepositoryStorageUsage** | Pointer to [**[]StorageUsageDTO**](StorageUsageDTO.md) | The content repository storage usage. | [optional] 
**ProvenanceRepositoryStorageUsage** | Pointer to [**[]StorageUsageDTO**](StorageUsageDTO.md) | The provenance repository storage usage. | [optional] 
**GarbageCollection** | Pointer to [**[]GarbageCollectionDTO**](GarbageCollectionDTO.md) | The garbage collection details. | [optional] 
**StatsLastRefreshed** | Pointer to **string** | When the diagnostics were generated. | [optional] 
**VersionInfo** | Pointer to [**VersionInfoDTO**](VersionInfoDTO.md) |  | [optional] 

## Methods

### NewSystemDiagnosticsSnapshotDTO

`func NewSystemDiagnosticsSnapshotDTO() *SystemDiagnosticsSnapshotDTO`

NewSystemDiagnosticsSnapshotDTO instantiates a new SystemDiagnosticsSnapshotDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemDiagnosticsSnapshotDTOWithDefaults

`func NewSystemDiagnosticsSnapshotDTOWithDefaults() *SystemDiagnosticsSnapshotDTO`

NewSystemDiagnosticsSnapshotDTOWithDefaults instantiates a new SystemDiagnosticsSnapshotDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalNonHeap

`func (o *SystemDiagnosticsSnapshotDTO) GetTotalNonHeap() string`

GetTotalNonHeap returns the TotalNonHeap field if non-nil, zero value otherwise.

### GetTotalNonHeapOk

`func (o *SystemDiagnosticsSnapshotDTO) GetTotalNonHeapOk() (*string, bool)`

GetTotalNonHeapOk returns a tuple with the TotalNonHeap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNonHeap

`func (o *SystemDiagnosticsSnapshotDTO) SetTotalNonHeap(v string)`

SetTotalNonHeap sets TotalNonHeap field to given value.

### HasTotalNonHeap

`func (o *SystemDiagnosticsSnapshotDTO) HasTotalNonHeap() bool`

HasTotalNonHeap returns a boolean if a field has been set.

### GetTotalNonHeapBytes

`func (o *SystemDiagnosticsSnapshotDTO) GetTotalNonHeapBytes() int64`

GetTotalNonHeapBytes returns the TotalNonHeapBytes field if non-nil, zero value otherwise.

### GetTotalNonHeapBytesOk

`func (o *SystemDiagnosticsSnapshotDTO) GetTotalNonHeapBytesOk() (*int64, bool)`

GetTotalNonHeapBytesOk returns a tuple with the TotalNonHeapBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNonHeapBytes

`func (o *SystemDiagnosticsSnapshotDTO) SetTotalNonHeapBytes(v int64)`

SetTotalNonHeapBytes sets TotalNonHeapBytes field to given value.

### HasTotalNonHeapBytes

`func (o *SystemDiagnosticsSnapshotDTO) HasTotalNonHeapBytes() bool`

HasTotalNonHeapBytes returns a boolean if a field has been set.

### GetUsedNonHeap

`func (o *SystemDiagnosticsSnapshotDTO) GetUsedNonHeap() string`

GetUsedNonHeap returns the UsedNonHeap field if non-nil, zero value otherwise.

### GetUsedNonHeapOk

`func (o *SystemDiagnosticsSnapshotDTO) GetUsedNonHeapOk() (*string, bool)`

GetUsedNonHeapOk returns a tuple with the UsedNonHeap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedNonHeap

`func (o *SystemDiagnosticsSnapshotDTO) SetUsedNonHeap(v string)`

SetUsedNonHeap sets UsedNonHeap field to given value.

### HasUsedNonHeap

`func (o *SystemDiagnosticsSnapshotDTO) HasUsedNonHeap() bool`

HasUsedNonHeap returns a boolean if a field has been set.

### GetUsedNonHeapBytes

`func (o *SystemDiagnosticsSnapshotDTO) GetUsedNonHeapBytes() int64`

GetUsedNonHeapBytes returns the UsedNonHeapBytes field if non-nil, zero value otherwise.

### GetUsedNonHeapBytesOk

`func (o *SystemDiagnosticsSnapshotDTO) GetUsedNonHeapBytesOk() (*int64, bool)`

GetUsedNonHeapBytesOk returns a tuple with the UsedNonHeapBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedNonHeapBytes

`func (o *SystemDiagnosticsSnapshotDTO) SetUsedNonHeapBytes(v int64)`

SetUsedNonHeapBytes sets UsedNonHeapBytes field to given value.

### HasUsedNonHeapBytes

`func (o *SystemDiagnosticsSnapshotDTO) HasUsedNonHeapBytes() bool`

HasUsedNonHeapBytes returns a boolean if a field has been set.

### GetFreeNonHeap

`func (o *SystemDiagnosticsSnapshotDTO) GetFreeNonHeap() string`

GetFreeNonHeap returns the FreeNonHeap field if non-nil, zero value otherwise.

### GetFreeNonHeapOk

`func (o *SystemDiagnosticsSnapshotDTO) GetFreeNonHeapOk() (*string, bool)`

GetFreeNonHeapOk returns a tuple with the FreeNonHeap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeNonHeap

`func (o *SystemDiagnosticsSnapshotDTO) SetFreeNonHeap(v string)`

SetFreeNonHeap sets FreeNonHeap field to given value.

### HasFreeNonHeap

`func (o *SystemDiagnosticsSnapshotDTO) HasFreeNonHeap() bool`

HasFreeNonHeap returns a boolean if a field has been set.

### GetFreeNonHeapBytes

`func (o *SystemDiagnosticsSnapshotDTO) GetFreeNonHeapBytes() int64`

GetFreeNonHeapBytes returns the FreeNonHeapBytes field if non-nil, zero value otherwise.

### GetFreeNonHeapBytesOk

`func (o *SystemDiagnosticsSnapshotDTO) GetFreeNonHeapBytesOk() (*int64, bool)`

GetFreeNonHeapBytesOk returns a tuple with the FreeNonHeapBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeNonHeapBytes

`func (o *SystemDiagnosticsSnapshotDTO) SetFreeNonHeapBytes(v int64)`

SetFreeNonHeapBytes sets FreeNonHeapBytes field to given value.

### HasFreeNonHeapBytes

`func (o *SystemDiagnosticsSnapshotDTO) HasFreeNonHeapBytes() bool`

HasFreeNonHeapBytes returns a boolean if a field has been set.

### GetMaxNonHeap

`func (o *SystemDiagnosticsSnapshotDTO) GetMaxNonHeap() string`

GetMaxNonHeap returns the MaxNonHeap field if non-nil, zero value otherwise.

### GetMaxNonHeapOk

`func (o *SystemDiagnosticsSnapshotDTO) GetMaxNonHeapOk() (*string, bool)`

GetMaxNonHeapOk returns a tuple with the MaxNonHeap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNonHeap

`func (o *SystemDiagnosticsSnapshotDTO) SetMaxNonHeap(v string)`

SetMaxNonHeap sets MaxNonHeap field to given value.

### HasMaxNonHeap

`func (o *SystemDiagnosticsSnapshotDTO) HasMaxNonHeap() bool`

HasMaxNonHeap returns a boolean if a field has been set.

### GetMaxNonHeapBytes

`func (o *SystemDiagnosticsSnapshotDTO) GetMaxNonHeapBytes() int64`

GetMaxNonHeapBytes returns the MaxNonHeapBytes field if non-nil, zero value otherwise.

### GetMaxNonHeapBytesOk

`func (o *SystemDiagnosticsSnapshotDTO) GetMaxNonHeapBytesOk() (*int64, bool)`

GetMaxNonHeapBytesOk returns a tuple with the MaxNonHeapBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNonHeapBytes

`func (o *SystemDiagnosticsSnapshotDTO) SetMaxNonHeapBytes(v int64)`

SetMaxNonHeapBytes sets MaxNonHeapBytes field to given value.

### HasMaxNonHeapBytes

`func (o *SystemDiagnosticsSnapshotDTO) HasMaxNonHeapBytes() bool`

HasMaxNonHeapBytes returns a boolean if a field has been set.

### GetNonHeapUtilization

`func (o *SystemDiagnosticsSnapshotDTO) GetNonHeapUtilization() string`

GetNonHeapUtilization returns the NonHeapUtilization field if non-nil, zero value otherwise.

### GetNonHeapUtilizationOk

`func (o *SystemDiagnosticsSnapshotDTO) GetNonHeapUtilizationOk() (*string, bool)`

GetNonHeapUtilizationOk returns a tuple with the NonHeapUtilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonHeapUtilization

`func (o *SystemDiagnosticsSnapshotDTO) SetNonHeapUtilization(v string)`

SetNonHeapUtilization sets NonHeapUtilization field to given value.

### HasNonHeapUtilization

`func (o *SystemDiagnosticsSnapshotDTO) HasNonHeapUtilization() bool`

HasNonHeapUtilization returns a boolean if a field has been set.

### GetTotalHeap

`func (o *SystemDiagnosticsSnapshotDTO) GetTotalHeap() string`

GetTotalHeap returns the TotalHeap field if non-nil, zero value otherwise.

### GetTotalHeapOk

`func (o *SystemDiagnosticsSnapshotDTO) GetTotalHeapOk() (*string, bool)`

GetTotalHeapOk returns a tuple with the TotalHeap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalHeap

`func (o *SystemDiagnosticsSnapshotDTO) SetTotalHeap(v string)`

SetTotalHeap sets TotalHeap field to given value.

### HasTotalHeap

`func (o *SystemDiagnosticsSnapshotDTO) HasTotalHeap() bool`

HasTotalHeap returns a boolean if a field has been set.

### GetTotalHeapBytes

`func (o *SystemDiagnosticsSnapshotDTO) GetTotalHeapBytes() int64`

GetTotalHeapBytes returns the TotalHeapBytes field if non-nil, zero value otherwise.

### GetTotalHeapBytesOk

`func (o *SystemDiagnosticsSnapshotDTO) GetTotalHeapBytesOk() (*int64, bool)`

GetTotalHeapBytesOk returns a tuple with the TotalHeapBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalHeapBytes

`func (o *SystemDiagnosticsSnapshotDTO) SetTotalHeapBytes(v int64)`

SetTotalHeapBytes sets TotalHeapBytes field to given value.

### HasTotalHeapBytes

`func (o *SystemDiagnosticsSnapshotDTO) HasTotalHeapBytes() bool`

HasTotalHeapBytes returns a boolean if a field has been set.

### GetUsedHeap

`func (o *SystemDiagnosticsSnapshotDTO) GetUsedHeap() string`

GetUsedHeap returns the UsedHeap field if non-nil, zero value otherwise.

### GetUsedHeapOk

`func (o *SystemDiagnosticsSnapshotDTO) GetUsedHeapOk() (*string, bool)`

GetUsedHeapOk returns a tuple with the UsedHeap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedHeap

`func (o *SystemDiagnosticsSnapshotDTO) SetUsedHeap(v string)`

SetUsedHeap sets UsedHeap field to given value.

### HasUsedHeap

`func (o *SystemDiagnosticsSnapshotDTO) HasUsedHeap() bool`

HasUsedHeap returns a boolean if a field has been set.

### GetUsedHeapBytes

`func (o *SystemDiagnosticsSnapshotDTO) GetUsedHeapBytes() int64`

GetUsedHeapBytes returns the UsedHeapBytes field if non-nil, zero value otherwise.

### GetUsedHeapBytesOk

`func (o *SystemDiagnosticsSnapshotDTO) GetUsedHeapBytesOk() (*int64, bool)`

GetUsedHeapBytesOk returns a tuple with the UsedHeapBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedHeapBytes

`func (o *SystemDiagnosticsSnapshotDTO) SetUsedHeapBytes(v int64)`

SetUsedHeapBytes sets UsedHeapBytes field to given value.

### HasUsedHeapBytes

`func (o *SystemDiagnosticsSnapshotDTO) HasUsedHeapBytes() bool`

HasUsedHeapBytes returns a boolean if a field has been set.

### GetFreeHeap

`func (o *SystemDiagnosticsSnapshotDTO) GetFreeHeap() string`

GetFreeHeap returns the FreeHeap field if non-nil, zero value otherwise.

### GetFreeHeapOk

`func (o *SystemDiagnosticsSnapshotDTO) GetFreeHeapOk() (*string, bool)`

GetFreeHeapOk returns a tuple with the FreeHeap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeHeap

`func (o *SystemDiagnosticsSnapshotDTO) SetFreeHeap(v string)`

SetFreeHeap sets FreeHeap field to given value.

### HasFreeHeap

`func (o *SystemDiagnosticsSnapshotDTO) HasFreeHeap() bool`

HasFreeHeap returns a boolean if a field has been set.

### GetFreeHeapBytes

`func (o *SystemDiagnosticsSnapshotDTO) GetFreeHeapBytes() int64`

GetFreeHeapBytes returns the FreeHeapBytes field if non-nil, zero value otherwise.

### GetFreeHeapBytesOk

`func (o *SystemDiagnosticsSnapshotDTO) GetFreeHeapBytesOk() (*int64, bool)`

GetFreeHeapBytesOk returns a tuple with the FreeHeapBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeHeapBytes

`func (o *SystemDiagnosticsSnapshotDTO) SetFreeHeapBytes(v int64)`

SetFreeHeapBytes sets FreeHeapBytes field to given value.

### HasFreeHeapBytes

`func (o *SystemDiagnosticsSnapshotDTO) HasFreeHeapBytes() bool`

HasFreeHeapBytes returns a boolean if a field has been set.

### GetMaxHeap

`func (o *SystemDiagnosticsSnapshotDTO) GetMaxHeap() string`

GetMaxHeap returns the MaxHeap field if non-nil, zero value otherwise.

### GetMaxHeapOk

`func (o *SystemDiagnosticsSnapshotDTO) GetMaxHeapOk() (*string, bool)`

GetMaxHeapOk returns a tuple with the MaxHeap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxHeap

`func (o *SystemDiagnosticsSnapshotDTO) SetMaxHeap(v string)`

SetMaxHeap sets MaxHeap field to given value.

### HasMaxHeap

`func (o *SystemDiagnosticsSnapshotDTO) HasMaxHeap() bool`

HasMaxHeap returns a boolean if a field has been set.

### GetMaxHeapBytes

`func (o *SystemDiagnosticsSnapshotDTO) GetMaxHeapBytes() int64`

GetMaxHeapBytes returns the MaxHeapBytes field if non-nil, zero value otherwise.

### GetMaxHeapBytesOk

`func (o *SystemDiagnosticsSnapshotDTO) GetMaxHeapBytesOk() (*int64, bool)`

GetMaxHeapBytesOk returns a tuple with the MaxHeapBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxHeapBytes

`func (o *SystemDiagnosticsSnapshotDTO) SetMaxHeapBytes(v int64)`

SetMaxHeapBytes sets MaxHeapBytes field to given value.

### HasMaxHeapBytes

`func (o *SystemDiagnosticsSnapshotDTO) HasMaxHeapBytes() bool`

HasMaxHeapBytes returns a boolean if a field has been set.

### GetHeapUtilization

`func (o *SystemDiagnosticsSnapshotDTO) GetHeapUtilization() string`

GetHeapUtilization returns the HeapUtilization field if non-nil, zero value otherwise.

### GetHeapUtilizationOk

`func (o *SystemDiagnosticsSnapshotDTO) GetHeapUtilizationOk() (*string, bool)`

GetHeapUtilizationOk returns a tuple with the HeapUtilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeapUtilization

`func (o *SystemDiagnosticsSnapshotDTO) SetHeapUtilization(v string)`

SetHeapUtilization sets HeapUtilization field to given value.

### HasHeapUtilization

`func (o *SystemDiagnosticsSnapshotDTO) HasHeapUtilization() bool`

HasHeapUtilization returns a boolean if a field has been set.

### GetAvailableProcessors

`func (o *SystemDiagnosticsSnapshotDTO) GetAvailableProcessors() int32`

GetAvailableProcessors returns the AvailableProcessors field if non-nil, zero value otherwise.

### GetAvailableProcessorsOk

`func (o *SystemDiagnosticsSnapshotDTO) GetAvailableProcessorsOk() (*int32, bool)`

GetAvailableProcessorsOk returns a tuple with the AvailableProcessors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableProcessors

`func (o *SystemDiagnosticsSnapshotDTO) SetAvailableProcessors(v int32)`

SetAvailableProcessors sets AvailableProcessors field to given value.

### HasAvailableProcessors

`func (o *SystemDiagnosticsSnapshotDTO) HasAvailableProcessors() bool`

HasAvailableProcessors returns a boolean if a field has been set.

### GetProcessorLoadAverage

`func (o *SystemDiagnosticsSnapshotDTO) GetProcessorLoadAverage() float64`

GetProcessorLoadAverage returns the ProcessorLoadAverage field if non-nil, zero value otherwise.

### GetProcessorLoadAverageOk

`func (o *SystemDiagnosticsSnapshotDTO) GetProcessorLoadAverageOk() (*float64, bool)`

GetProcessorLoadAverageOk returns a tuple with the ProcessorLoadAverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessorLoadAverage

`func (o *SystemDiagnosticsSnapshotDTO) SetProcessorLoadAverage(v float64)`

SetProcessorLoadAverage sets ProcessorLoadAverage field to given value.

### HasProcessorLoadAverage

`func (o *SystemDiagnosticsSnapshotDTO) HasProcessorLoadAverage() bool`

HasProcessorLoadAverage returns a boolean if a field has been set.

### GetTotalThreads

`func (o *SystemDiagnosticsSnapshotDTO) GetTotalThreads() int32`

GetTotalThreads returns the TotalThreads field if non-nil, zero value otherwise.

### GetTotalThreadsOk

`func (o *SystemDiagnosticsSnapshotDTO) GetTotalThreadsOk() (*int32, bool)`

GetTotalThreadsOk returns a tuple with the TotalThreads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalThreads

`func (o *SystemDiagnosticsSnapshotDTO) SetTotalThreads(v int32)`

SetTotalThreads sets TotalThreads field to given value.

### HasTotalThreads

`func (o *SystemDiagnosticsSnapshotDTO) HasTotalThreads() bool`

HasTotalThreads returns a boolean if a field has been set.

### GetDaemonThreads

`func (o *SystemDiagnosticsSnapshotDTO) GetDaemonThreads() int32`

GetDaemonThreads returns the DaemonThreads field if non-nil, zero value otherwise.

### GetDaemonThreadsOk

`func (o *SystemDiagnosticsSnapshotDTO) GetDaemonThreadsOk() (*int32, bool)`

GetDaemonThreadsOk returns a tuple with the DaemonThreads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaemonThreads

`func (o *SystemDiagnosticsSnapshotDTO) SetDaemonThreads(v int32)`

SetDaemonThreads sets DaemonThreads field to given value.

### HasDaemonThreads

`func (o *SystemDiagnosticsSnapshotDTO) HasDaemonThreads() bool`

HasDaemonThreads returns a boolean if a field has been set.

### GetUptime

`func (o *SystemDiagnosticsSnapshotDTO) GetUptime() string`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *SystemDiagnosticsSnapshotDTO) GetUptimeOk() (*string, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *SystemDiagnosticsSnapshotDTO) SetUptime(v string)`

SetUptime sets Uptime field to given value.

### HasUptime

`func (o *SystemDiagnosticsSnapshotDTO) HasUptime() bool`

HasUptime returns a boolean if a field has been set.

### GetFlowFileRepositoryStorageUsage

`func (o *SystemDiagnosticsSnapshotDTO) GetFlowFileRepositoryStorageUsage() StorageUsageDTO`

GetFlowFileRepositoryStorageUsage returns the FlowFileRepositoryStorageUsage field if non-nil, zero value otherwise.

### GetFlowFileRepositoryStorageUsageOk

`func (o *SystemDiagnosticsSnapshotDTO) GetFlowFileRepositoryStorageUsageOk() (*StorageUsageDTO, bool)`

GetFlowFileRepositoryStorageUsageOk returns a tuple with the FlowFileRepositoryStorageUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowFileRepositoryStorageUsage

`func (o *SystemDiagnosticsSnapshotDTO) SetFlowFileRepositoryStorageUsage(v StorageUsageDTO)`

SetFlowFileRepositoryStorageUsage sets FlowFileRepositoryStorageUsage field to given value.

### HasFlowFileRepositoryStorageUsage

`func (o *SystemDiagnosticsSnapshotDTO) HasFlowFileRepositoryStorageUsage() bool`

HasFlowFileRepositoryStorageUsage returns a boolean if a field has been set.

### GetContentRepositoryStorageUsage

`func (o *SystemDiagnosticsSnapshotDTO) GetContentRepositoryStorageUsage() []StorageUsageDTO`

GetContentRepositoryStorageUsage returns the ContentRepositoryStorageUsage field if non-nil, zero value otherwise.

### GetContentRepositoryStorageUsageOk

`func (o *SystemDiagnosticsSnapshotDTO) GetContentRepositoryStorageUsageOk() (*[]StorageUsageDTO, bool)`

GetContentRepositoryStorageUsageOk returns a tuple with the ContentRepositoryStorageUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentRepositoryStorageUsage

`func (o *SystemDiagnosticsSnapshotDTO) SetContentRepositoryStorageUsage(v []StorageUsageDTO)`

SetContentRepositoryStorageUsage sets ContentRepositoryStorageUsage field to given value.

### HasContentRepositoryStorageUsage

`func (o *SystemDiagnosticsSnapshotDTO) HasContentRepositoryStorageUsage() bool`

HasContentRepositoryStorageUsage returns a boolean if a field has been set.

### GetProvenanceRepositoryStorageUsage

`func (o *SystemDiagnosticsSnapshotDTO) GetProvenanceRepositoryStorageUsage() []StorageUsageDTO`

GetProvenanceRepositoryStorageUsage returns the ProvenanceRepositoryStorageUsage field if non-nil, zero value otherwise.

### GetProvenanceRepositoryStorageUsageOk

`func (o *SystemDiagnosticsSnapshotDTO) GetProvenanceRepositoryStorageUsageOk() (*[]StorageUsageDTO, bool)`

GetProvenanceRepositoryStorageUsageOk returns a tuple with the ProvenanceRepositoryStorageUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvenanceRepositoryStorageUsage

`func (o *SystemDiagnosticsSnapshotDTO) SetProvenanceRepositoryStorageUsage(v []StorageUsageDTO)`

SetProvenanceRepositoryStorageUsage sets ProvenanceRepositoryStorageUsage field to given value.

### HasProvenanceRepositoryStorageUsage

`func (o *SystemDiagnosticsSnapshotDTO) HasProvenanceRepositoryStorageUsage() bool`

HasProvenanceRepositoryStorageUsage returns a boolean if a field has been set.

### GetGarbageCollection

`func (o *SystemDiagnosticsSnapshotDTO) GetGarbageCollection() []GarbageCollectionDTO`

GetGarbageCollection returns the GarbageCollection field if non-nil, zero value otherwise.

### GetGarbageCollectionOk

`func (o *SystemDiagnosticsSnapshotDTO) GetGarbageCollectionOk() (*[]GarbageCollectionDTO, bool)`

GetGarbageCollectionOk returns a tuple with the GarbageCollection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGarbageCollection

`func (o *SystemDiagnosticsSnapshotDTO) SetGarbageCollection(v []GarbageCollectionDTO)`

SetGarbageCollection sets GarbageCollection field to given value.

### HasGarbageCollection

`func (o *SystemDiagnosticsSnapshotDTO) HasGarbageCollection() bool`

HasGarbageCollection returns a boolean if a field has been set.

### GetStatsLastRefreshed

`func (o *SystemDiagnosticsSnapshotDTO) GetStatsLastRefreshed() string`

GetStatsLastRefreshed returns the StatsLastRefreshed field if non-nil, zero value otherwise.

### GetStatsLastRefreshedOk

`func (o *SystemDiagnosticsSnapshotDTO) GetStatsLastRefreshedOk() (*string, bool)`

GetStatsLastRefreshedOk returns a tuple with the StatsLastRefreshed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatsLastRefreshed

`func (o *SystemDiagnosticsSnapshotDTO) SetStatsLastRefreshed(v string)`

SetStatsLastRefreshed sets StatsLastRefreshed field to given value.

### HasStatsLastRefreshed

`func (o *SystemDiagnosticsSnapshotDTO) HasStatsLastRefreshed() bool`

HasStatsLastRefreshed returns a boolean if a field has been set.

### GetVersionInfo

`func (o *SystemDiagnosticsSnapshotDTO) GetVersionInfo() VersionInfoDTO`

GetVersionInfo returns the VersionInfo field if non-nil, zero value otherwise.

### GetVersionInfoOk

`func (o *SystemDiagnosticsSnapshotDTO) GetVersionInfoOk() (*VersionInfoDTO, bool)`

GetVersionInfoOk returns a tuple with the VersionInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionInfo

`func (o *SystemDiagnosticsSnapshotDTO) SetVersionInfo(v VersionInfoDTO)`

SetVersionInfo sets VersionInfo field to given value.

### HasVersionInfo

`func (o *SystemDiagnosticsSnapshotDTO) HasVersionInfo() bool`

HasVersionInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


