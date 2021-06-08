/*
 * NiFi Rest Api
 *
 * The Rest Api provides programmatic access to command and control a NiFi instance in real time. Start and                                              stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.13.2
 * Contact: dev@nifi.apache.org
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nifi

import (
	"encoding/json"
)

// SystemDiagnosticsSnapshotDTO struct for SystemDiagnosticsSnapshotDTO
type SystemDiagnosticsSnapshotDTO struct {
	// Total size of non heap.
	TotalNonHeap *string `json:"totalNonHeap,omitempty"`
	// Total number of bytes allocated to the JVM not used for heap
	TotalNonHeapBytes *int64 `json:"totalNonHeapBytes,omitempty"`
	// Amount of use non heap.
	UsedNonHeap *string `json:"usedNonHeap,omitempty"`
	// Total number of bytes used by the JVM not in the heap space
	UsedNonHeapBytes *int64 `json:"usedNonHeapBytes,omitempty"`
	// Amount of free non heap.
	FreeNonHeap *string `json:"freeNonHeap,omitempty"`
	// Total number of free non-heap bytes available to the JVM
	FreeNonHeapBytes *int64 `json:"freeNonHeapBytes,omitempty"`
	// Maximum size of non heap.
	MaxNonHeap *string `json:"maxNonHeap,omitempty"`
	// The maximum number of bytes that the JVM can use for non-heap purposes
	MaxNonHeapBytes *int64 `json:"maxNonHeapBytes,omitempty"`
	// Utilization of non heap.
	NonHeapUtilization *string `json:"nonHeapUtilization,omitempty"`
	// Total size of heap.
	TotalHeap *string `json:"totalHeap,omitempty"`
	// The total number of bytes that are available for the JVM heap to use
	TotalHeapBytes *int64 `json:"totalHeapBytes,omitempty"`
	// Amount of used heap.
	UsedHeap *string `json:"usedHeap,omitempty"`
	// The number of bytes of JVM heap that are currently being used
	UsedHeapBytes *int64 `json:"usedHeapBytes,omitempty"`
	// Amount of free heap.
	FreeHeap *string `json:"freeHeap,omitempty"`
	// The number of bytes that are allocated to the JVM heap but not currently being used
	FreeHeapBytes *int64 `json:"freeHeapBytes,omitempty"`
	// Maximum size of heap.
	MaxHeap *string `json:"maxHeap,omitempty"`
	// The maximum number of bytes that can be used by the JVM
	MaxHeapBytes *int64 `json:"maxHeapBytes,omitempty"`
	// Utilization of heap.
	HeapUtilization *string `json:"heapUtilization,omitempty"`
	// Number of available processors if supported by the underlying system.
	AvailableProcessors *int32 `json:"availableProcessors,omitempty"`
	// The processor load average if supported by the underlying system.
	ProcessorLoadAverage *float64 `json:"processorLoadAverage,omitempty"`
	// Total number of threads.
	TotalThreads *int32 `json:"totalThreads,omitempty"`
	// Number of daemon threads.
	DaemonThreads *int32 `json:"daemonThreads,omitempty"`
	// The uptime of the Java virtual machine
	Uptime                         *string          `json:"uptime,omitempty"`
	FlowFileRepositoryStorageUsage *StorageUsageDTO `json:"flowFileRepositoryStorageUsage,omitempty"`
	// The content repository storage usage.
	ContentRepositoryStorageUsage *[]StorageUsageDTO `json:"contentRepositoryStorageUsage,omitempty"`
	// The provenance repository storage usage.
	ProvenanceRepositoryStorageUsage *[]StorageUsageDTO `json:"provenanceRepositoryStorageUsage,omitempty"`
	// The garbage collection details.
	GarbageCollection *[]GarbageCollectionDTO `json:"garbageCollection,omitempty"`
	// When the diagnostics were generated.
	StatsLastRefreshed *string         `json:"statsLastRefreshed,omitempty"`
	VersionInfo        *VersionInfoDTO `json:"versionInfo,omitempty"`
}

// NewSystemDiagnosticsSnapshotDTO instantiates a new SystemDiagnosticsSnapshotDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystemDiagnosticsSnapshotDTO() *SystemDiagnosticsSnapshotDTO {
	this := SystemDiagnosticsSnapshotDTO{}
	return &this
}

// NewSystemDiagnosticsSnapshotDTOWithDefaults instantiates a new SystemDiagnosticsSnapshotDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemDiagnosticsSnapshotDTOWithDefaults() *SystemDiagnosticsSnapshotDTO {
	this := SystemDiagnosticsSnapshotDTO{}
	return &this
}

// GetTotalNonHeap returns the TotalNonHeap field value if set, zero value otherwise.
func (o *SystemDiagnosticsSnapshotDTO) GetTotalNonHeap() string {
	if o == nil || o.TotalNonHeap == nil {
		var ret string
		return ret
	}
	return *o.TotalNonHeap
}

// GetTotalNonHeapOk returns a tuple with the TotalNonHeap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemDiagnosticsSnapshotDTO) GetTotalNonHeapOk() (*string, bool) {
	if o == nil || o.TotalNonHeap == nil {
		return nil, false
	}
	return o.TotalNonHeap, true
}

// HasTotalNonHeap returns a boolean if a field has been set.
func (o *SystemDiagnosticsSnapshotDTO) HasTotalNonHeap() bool {
	if o != nil && o.TotalNonHeap != nil {
		return true
	}

	return false
}

// SetTotalNonHeap gets a reference to the given string and assigns it to the TotalNonHeap field.
func (o *SystemDiagnosticsSnapshotDTO) SetTotalNonHeap(v string) {
	o.TotalNonHeap = &v
}

// GetTotalNonHeapBytes returns the TotalNonHeapBytes field value if set, zero value otherwise.
func (o *SystemDiagnosticsSnapshotDTO) GetTotalNonHeapBytes() int64 {
	if o == nil || o.TotalNonHeapBytes == nil {
		var ret int64
		return ret
	}
	return *o.TotalNonHeapBytes
}

// GetTotalNonHeapBytesOk returns a tuple with the TotalNonHeapBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemDiagnosticsSnapshotDTO) GetTotalNonHeapBytesOk() (*int64, bool) {
	if o == nil || o.TotalNonHeapBytes == nil {
		return nil, false
	}
	return o.TotalNonHeapBytes, true
}

// HasTotalNonHeapBytes returns a boolean if a field has been set.
func (o *SystemDiagnosticsSnapshotDTO) HasTotalNonHeapBytes() bool {
	if o != nil && o.TotalNonHeapBytes != nil {
		return true
	}

	return false
}

// SetTotalNonHeapBytes gets a reference to the given int64 and assigns it to the TotalNonHeapBytes field.
func (o *SystemDiagnosticsSnapshotDTO) SetTotalNonHeapBytes(v int64) {
	o.TotalNonHeapBytes = &v
}

// GetUsedNonHeap returns the UsedNonHeap field value if set, zero value otherwise.
func (o *SystemDiagnosticsSnapshotDTO) GetUsedNonHeap() string {
	if o == nil || o.UsedNonHeap == nil {
		var ret string
		return ret
	}
	return *o.UsedNonHeap
}

// GetUsedNonHeapOk returns a tuple with the UsedNonHeap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemDiagnosticsSnapshotDTO) GetUsedNonHeapOk() (*string, bool) {
	if o == nil || o.UsedNonHeap == nil {
		return nil, false
	}
	return o.UsedNonHeap, true
}

// HasUsedNonHeap returns a boolean if a field has been set.
func (o *SystemDiagnosticsSnapshotDTO) HasUsedNonHeap() bool {
	if o != nil && o.UsedNonHeap != nil {
		return true
	}

	return false
}

// SetUsedNonHeap gets a reference to the given string and assigns it to the UsedNonHeap field.
func (o *SystemDiagnosticsSnapshotDTO) SetUsedNonHeap(v string) {
	o.UsedNonHeap = &v
}

// GetUsedNonHeapBytes returns the UsedNonHeapBytes field value if set, zero value otherwise.
func (o *SystemDiagnosticsSnapshotDTO) GetUsedNonHeapBytes() int64 {
	if o == nil || o.UsedNonHeapBytes == nil {
		var ret int64
		return ret
	}
	return *o.UsedNonHeapBytes
}

// GetUsedNonHeapBytesOk returns a tuple with the UsedNonHeapBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemDiagnosticsSnapshotDTO) GetUsedNonHeapBytesOk() (*int64, bool) {
	if o == nil || o.UsedNonHeapBytes == nil {
		return nil, false
	}
	return o.UsedNonHeapBytes, true
}

// HasUsedNonHeapBytes returns a boolean if a field has been set.
func (o *SystemDiagnosticsSnapshotDTO) HasUsedNonHeapBytes() bool {
	if o != nil && o.UsedNonHeapBytes != nil {
		return true
	}

	return false
}

// SetUsedNonHeapBytes gets a reference to the given int64 and assigns it to the UsedNonHeapBytes field.
func (o *SystemDiagnosticsSnapshotDTO) SetUsedNonHeapBytes(v int64) {
	o.UsedNonHeapBytes = &v
}

// GetFreeNonHeap returns the FreeNonHeap field value if set, zero value otherwise.
func (o *SystemDiagnosticsSnapshotDTO) GetFreeNonHeap() string {
	if o == nil || o.FreeNonHeap == nil {
		var ret string
		return ret
	}
	return *o.FreeNonHeap
}

// GetFreeNonHeapOk returns a tuple with the FreeNonHeap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemDiagnosticsSnapshotDTO) GetFreeNonHeapOk() (*string, bool) {
	if o == nil || o.FreeNonHeap == nil {
		return nil, false
	}
	return o.FreeNonHeap, true
}

// HasFreeNonHeap returns a boolean if a field has been set.
func (o *SystemDiagnosticsSnapshotDTO) HasFreeNonHeap() bool {
	if o != nil && o.FreeNonHeap != nil {
		return true
	}

	return false
}

// SetFreeNonHeap gets a reference to the given string and assigns it to the FreeNonHeap field.
func (o *SystemDiagnosticsSnapshotDTO) SetFreeNonHeap(v string) {
	o.FreeNonHeap = &v
}

// GetFreeNonHeapBytes returns the FreeNonHeapBytes field value if set, zero value otherwise.
func (o *SystemDiagnosticsSnapshotDTO) GetFreeNonHeapBytes() int64 {
	if o == nil || o.FreeNonHeapBytes == nil {
		var ret int64
		return ret
	}
	return *o.FreeNonHeapBytes
}

// GetFreeNonHeapBytesOk returns a tuple with the FreeNonHeapBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemDiagnosticsSnapshotDTO) GetFreeNonHeapBytesOk() (*int64, bool) {
	if o == nil || o.FreeNonHeapBytes == nil {
		return nil, false
	}
	return o.FreeNonHeapBytes, true
}

// HasFreeNonHeapBytes returns a boolean if a field has been set.
func (o *SystemDiagnosticsSnapshotDTO) HasFreeNonHeapBytes() bool {
	if o != nil && o.FreeNonHeapBytes != nil {
		return true
	}

	return false
}

// SetFreeNonHeapBytes gets a reference to the given int64 and assigns it to the FreeNonHeapBytes field.
func (o *SystemDiagnosticsSnapshotDTO) SetFreeNonHeapBytes(v int64) {
	o.FreeNonHeapBytes = &v
}

// GetMaxNonHeap returns the MaxNonHeap field value if set, zero value otherwise.
func (o *SystemDiagnosticsSnapshotDTO) GetMaxNonHeap() string {
	if o == nil || o.MaxNonHeap == nil {
		var ret string
		return ret
	}
	return *o.MaxNonHeap
}

// GetMaxNonHeapOk returns a tuple with the MaxNonHeap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemDiagnosticsSnapshotDTO) GetMaxNonHeapOk() (*string, bool) {
	if o == nil || o.MaxNonHeap == nil {
		return nil, false
	}
	return o.MaxNonHeap, true
}

// HasMaxNonHeap returns a boolean if a field has been set.
func (o *SystemDiagnosticsSnapshotDTO) HasMaxNonHeap() bool {
	if o != nil && o.MaxNonHeap != nil {
		return true
	}

	return false
}

// SetMaxNonHeap gets a reference to the given string and assigns it to the MaxNonHeap field.
func (o *SystemDiagnosticsSnapshotDTO) SetMaxNonHeap(v string) {
	o.MaxNonHeap = &v
}

// GetMaxNonHeapBytes returns the MaxNonHeapBytes field value if set, zero value otherwise.
func (o *SystemDiagnosticsSnapshotDTO) GetMaxNonHeapBytes() int64 {
	if o == nil || o.MaxNonHeapBytes == nil {
		var ret int64
		return ret
	}
	return *o.MaxNonHeapBytes
}

// GetMaxNonHeapBytesOk returns a tuple with the MaxNonHeapBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemDiagnosticsSnapshotDTO) GetMaxNonHeapBytesOk() (*int64, bool) {
	if o == nil || o.MaxNonHeapBytes == nil {
		return nil, false
	}
	return o.MaxNonHeapBytes, true
}

// HasMaxNonHeapBytes returns a boolean if a field has been set.
func (o *SystemDiagnosticsSnapshotDTO) HasMaxNonHeapBytes() bool {
	if o != nil && o.MaxNonHeapBytes != nil {
		return true
	}

	return false
}

// SetMaxNonHeapBytes gets a reference to the given int64 and assigns it to the MaxNonHeapBytes field.
func (o *SystemDiagnosticsSnapshotDTO) SetMaxNonHeapBytes(v int64) {
	o.MaxNonHeapBytes = &v
}

// GetNonHeapUtilization returns the NonHeapUtilization field value if set, zero value otherwise.
func (o *SystemDiagnosticsSnapshotDTO) GetNonHeapUtilization() string {
	if o == nil || o.NonHeapUtilization == nil {
		var ret string
		return ret
	}
	return *o.NonHeapUtilization
}

// GetNonHeapUtilizationOk returns a tuple with the NonHeapUtilization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemDiagnosticsSnapshotDTO) GetNonHeapUtilizationOk() (*string, bool) {
	if o == nil || o.NonHeapUtilization == nil {
		return nil, false
	}
	return o.NonHeapUtilization, true
}

// HasNonHeapUtilization returns a boolean if a field has been set.
func (o *SystemDiagnosticsSnapshotDTO) HasNonHeapUtilization() bool {
	if o != nil && o.NonHeapUtilization != nil {
		return true
	}

	return false
}

// SetNonHeapUtilization gets a reference to the given string and assigns it to the NonHeapUtilization field.
func (o *SystemDiagnosticsSnapshotDTO) SetNonHeapUtilization(v string) {
	o.NonHeapUtilization = &v
}

// GetTotalHeap returns the TotalHeap field value if set, zero value otherwise.
func (o *SystemDiagnosticsSnapshotDTO) GetTotalHeap() string {
	if o == nil || o.TotalHeap == nil {
		var ret string
		return ret
	}
	return *o.TotalHeap
}

// GetTotalHeapOk returns a tuple with the TotalHeap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemDiagnosticsSnapshotDTO) GetTotalHeapOk() (*string, bool) {
	if o == nil || o.TotalHeap == nil {
		return nil, false
	}
	return o.TotalHeap, true
}

// HasTotalHeap returns a boolean if a field has been set.
func (o *SystemDiagnosticsSnapshotDTO) HasTotalHeap() bool {
	if o != nil && o.TotalHeap != nil {
		return true
	}

	return false
}

// SetTotalHeap gets a reference to the given string and assigns it to the TotalHeap field.
func (o *SystemDiagnosticsSnapshotDTO) SetTotalHeap(v string) {
	o.TotalHeap = &v
}

// GetTotalHeapBytes returns the TotalHeapBytes field value if set, zero value otherwise.
func (o *SystemDiagnosticsSnapshotDTO) GetTotalHeapBytes() int64 {
	if o == nil || o.TotalHeapBytes == nil {
		var ret int64
		return ret
	}
	return *o.TotalHeapBytes
}

// GetTotalHeapBytesOk returns a tuple with the TotalHeapBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemDiagnosticsSnapshotDTO) GetTotalHeapBytesOk() (*int64, bool) {
	if o == nil || o.TotalHeapBytes == nil {
		return nil, false
	}
	return o.TotalHeapBytes, true
}

// HasTotalHeapBytes returns a boolean if a field has been set.
func (o *SystemDiagnosticsSnapshotDTO) HasTotalHeapBytes() bool {
	if o != nil && o.TotalHeapBytes != nil {
		return true
	}

	return false
}

// SetTotalHeapBytes gets a reference to the given int64 and assigns it to the TotalHeapBytes field.
func (o *SystemDiagnosticsSnapshotDTO) SetTotalHeapBytes(v int64) {
	o.TotalHeapBytes = &v
}

// GetUsedHeap returns the UsedHeap field value if set, zero value otherwise.
func (o *SystemDiagnosticsSnapshotDTO) GetUsedHeap() string {
	if o == nil || o.UsedHeap == nil {
		var ret string
		return ret
	}
	return *o.UsedHeap
}

// GetUsedHeapOk returns a tuple with the UsedHeap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemDiagnosticsSnapshotDTO) GetUsedHeapOk() (*string, bool) {
	if o == nil || o.UsedHeap == nil {
		return nil, false
	}
	return o.UsedHeap, true
}

// HasUsedHeap returns a boolean if a field has been set.
func (o *SystemDiagnosticsSnapshotDTO) HasUsedHeap() bool {
	if o != nil && o.UsedHeap != nil {
		return true
	}

	return false
}

// SetUsedHeap gets a reference to the given string and assigns it to the UsedHeap field.
func (o *SystemDiagnosticsSnapshotDTO) SetUsedHeap(v string) {
	o.UsedHeap = &v
}

// GetUsedHeapBytes returns the UsedHeapBytes field value if set, zero value otherwise.
func (o *SystemDiagnosticsSnapshotDTO) GetUsedHeapBytes() int64 {
	if o == nil || o.UsedHeapBytes == nil {
		var ret int64
		return ret
	}
	return *o.UsedHeapBytes
}

// GetUsedHeapBytesOk returns a tuple with the UsedHeapBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemDiagnosticsSnapshotDTO) GetUsedHeapBytesOk() (*int64, bool) {
	if o == nil || o.UsedHeapBytes == nil {
		return nil, false
	}
	return o.UsedHeapBytes, true
}

// HasUsedHeapBytes returns a boolean if a field has been set.
func (o *SystemDiagnosticsSnapshotDTO) HasUsedHeapBytes() bool {
	if o != nil && o.UsedHeapBytes != nil {
		return true
	}

	return false
}

// SetUsedHeapBytes gets a reference to the given int64 and assigns it to the UsedHeapBytes field.
func (o *SystemDiagnosticsSnapshotDTO) SetUsedHeapBytes(v int64) {
	o.UsedHeapBytes = &v
}

// GetFreeHeap returns the FreeHeap field value if set, zero value otherwise.
func (o *SystemDiagnosticsSnapshotDTO) GetFreeHeap() string {
	if o == nil || o.FreeHeap == nil {
		var ret string
		return ret
	}
	return *o.FreeHeap
}

// GetFreeHeapOk returns a tuple with the FreeHeap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemDiagnosticsSnapshotDTO) GetFreeHeapOk() (*string, bool) {
	if o == nil || o.FreeHeap == nil {
		return nil, false
	}
	return o.FreeHeap, true
}

// HasFreeHeap returns a boolean if a field has been set.
func (o *SystemDiagnosticsSnapshotDTO) HasFreeHeap() bool {
	if o != nil && o.FreeHeap != nil {
		return true
	}

	return false
}

// SetFreeHeap gets a reference to the given string and assigns it to the FreeHeap field.
func (o *SystemDiagnosticsSnapshotDTO) SetFreeHeap(v string) {
	o.FreeHeap = &v
}

// GetFreeHeapBytes returns the FreeHeapBytes field value if set, zero value otherwise.
func (o *SystemDiagnosticsSnapshotDTO) GetFreeHeapBytes() int64 {
	if o == nil || o.FreeHeapBytes == nil {
		var ret int64
		return ret
	}
	return *o.FreeHeapBytes
}

// GetFreeHeapBytesOk returns a tuple with the FreeHeapBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemDiagnosticsSnapshotDTO) GetFreeHeapBytesOk() (*int64, bool) {
	if o == nil || o.FreeHeapBytes == nil {
		return nil, false
	}
	return o.FreeHeapBytes, true
}

// HasFreeHeapBytes returns a boolean if a field has been set.
func (o *SystemDiagnosticsSnapshotDTO) HasFreeHeapBytes() bool {
	if o != nil && o.FreeHeapBytes != nil {
		return true
	}

	return false
}

// SetFreeHeapBytes gets a reference to the given int64 and assigns it to the FreeHeapBytes field.
func (o *SystemDiagnosticsSnapshotDTO) SetFreeHeapBytes(v int64) {
	o.FreeHeapBytes = &v
}

// GetMaxHeap returns the MaxHeap field value if set, zero value otherwise.
func (o *SystemDiagnosticsSnapshotDTO) GetMaxHeap() string {
	if o == nil || o.MaxHeap == nil {
		var ret string
		return ret
	}
	return *o.MaxHeap
}

// GetMaxHeapOk returns a tuple with the MaxHeap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemDiagnosticsSnapshotDTO) GetMaxHeapOk() (*string, bool) {
	if o == nil || o.MaxHeap == nil {
		return nil, false
	}
	return o.MaxHeap, true
}

// HasMaxHeap returns a boolean if a field has been set.
func (o *SystemDiagnosticsSnapshotDTO) HasMaxHeap() bool {
	if o != nil && o.MaxHeap != nil {
		return true
	}

	return false
}

// SetMaxHeap gets a reference to the given string and assigns it to the MaxHeap field.
func (o *SystemDiagnosticsSnapshotDTO) SetMaxHeap(v string) {
	o.MaxHeap = &v
}

// GetMaxHeapBytes returns the MaxHeapBytes field value if set, zero value otherwise.
func (o *SystemDiagnosticsSnapshotDTO) GetMaxHeapBytes() int64 {
	if o == nil || o.MaxHeapBytes == nil {
		var ret int64
		return ret
	}
	return *o.MaxHeapBytes
}

// GetMaxHeapBytesOk returns a tuple with the MaxHeapBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemDiagnosticsSnapshotDTO) GetMaxHeapBytesOk() (*int64, bool) {
	if o == nil || o.MaxHeapBytes == nil {
		return nil, false
	}
	return o.MaxHeapBytes, true
}

// HasMaxHeapBytes returns a boolean if a field has been set.
func (o *SystemDiagnosticsSnapshotDTO) HasMaxHeapBytes() bool {
	if o != nil && o.MaxHeapBytes != nil {
		return true
	}

	return false
}

// SetMaxHeapBytes gets a reference to the given int64 and assigns it to the MaxHeapBytes field.
func (o *SystemDiagnosticsSnapshotDTO) SetMaxHeapBytes(v int64) {
	o.MaxHeapBytes = &v
}

// GetHeapUtilization returns the HeapUtilization field value if set, zero value otherwise.
func (o *SystemDiagnosticsSnapshotDTO) GetHeapUtilization() string {
	if o == nil || o.HeapUtilization == nil {
		var ret string
		return ret
	}
	return *o.HeapUtilization
}

// GetHeapUtilizationOk returns a tuple with the HeapUtilization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemDiagnosticsSnapshotDTO) GetHeapUtilizationOk() (*string, bool) {
	if o == nil || o.HeapUtilization == nil {
		return nil, false
	}
	return o.HeapUtilization, true
}

// HasHeapUtilization returns a boolean if a field has been set.
func (o *SystemDiagnosticsSnapshotDTO) HasHeapUtilization() bool {
	if o != nil && o.HeapUtilization != nil {
		return true
	}

	return false
}

// SetHeapUtilization gets a reference to the given string and assigns it to the HeapUtilization field.
func (o *SystemDiagnosticsSnapshotDTO) SetHeapUtilization(v string) {
	o.HeapUtilization = &v
}

// GetAvailableProcessors returns the AvailableProcessors field value if set, zero value otherwise.
func (o *SystemDiagnosticsSnapshotDTO) GetAvailableProcessors() int32 {
	if o == nil || o.AvailableProcessors == nil {
		var ret int32
		return ret
	}
	return *o.AvailableProcessors
}

// GetAvailableProcessorsOk returns a tuple with the AvailableProcessors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemDiagnosticsSnapshotDTO) GetAvailableProcessorsOk() (*int32, bool) {
	if o == nil || o.AvailableProcessors == nil {
		return nil, false
	}
	return o.AvailableProcessors, true
}

// HasAvailableProcessors returns a boolean if a field has been set.
func (o *SystemDiagnosticsSnapshotDTO) HasAvailableProcessors() bool {
	if o != nil && o.AvailableProcessors != nil {
		return true
	}

	return false
}

// SetAvailableProcessors gets a reference to the given int32 and assigns it to the AvailableProcessors field.
func (o *SystemDiagnosticsSnapshotDTO) SetAvailableProcessors(v int32) {
	o.AvailableProcessors = &v
}

// GetProcessorLoadAverage returns the ProcessorLoadAverage field value if set, zero value otherwise.
func (o *SystemDiagnosticsSnapshotDTO) GetProcessorLoadAverage() float64 {
	if o == nil || o.ProcessorLoadAverage == nil {
		var ret float64
		return ret
	}
	return *o.ProcessorLoadAverage
}

// GetProcessorLoadAverageOk returns a tuple with the ProcessorLoadAverage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemDiagnosticsSnapshotDTO) GetProcessorLoadAverageOk() (*float64, bool) {
	if o == nil || o.ProcessorLoadAverage == nil {
		return nil, false
	}
	return o.ProcessorLoadAverage, true
}

// HasProcessorLoadAverage returns a boolean if a field has been set.
func (o *SystemDiagnosticsSnapshotDTO) HasProcessorLoadAverage() bool {
	if o != nil && o.ProcessorLoadAverage != nil {
		return true
	}

	return false
}

// SetProcessorLoadAverage gets a reference to the given float64 and assigns it to the ProcessorLoadAverage field.
func (o *SystemDiagnosticsSnapshotDTO) SetProcessorLoadAverage(v float64) {
	o.ProcessorLoadAverage = &v
}

// GetTotalThreads returns the TotalThreads field value if set, zero value otherwise.
func (o *SystemDiagnosticsSnapshotDTO) GetTotalThreads() int32 {
	if o == nil || o.TotalThreads == nil {
		var ret int32
		return ret
	}
	return *o.TotalThreads
}

// GetTotalThreadsOk returns a tuple with the TotalThreads field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemDiagnosticsSnapshotDTO) GetTotalThreadsOk() (*int32, bool) {
	if o == nil || o.TotalThreads == nil {
		return nil, false
	}
	return o.TotalThreads, true
}

// HasTotalThreads returns a boolean if a field has been set.
func (o *SystemDiagnosticsSnapshotDTO) HasTotalThreads() bool {
	if o != nil && o.TotalThreads != nil {
		return true
	}

	return false
}

// SetTotalThreads gets a reference to the given int32 and assigns it to the TotalThreads field.
func (o *SystemDiagnosticsSnapshotDTO) SetTotalThreads(v int32) {
	o.TotalThreads = &v
}

// GetDaemonThreads returns the DaemonThreads field value if set, zero value otherwise.
func (o *SystemDiagnosticsSnapshotDTO) GetDaemonThreads() int32 {
	if o == nil || o.DaemonThreads == nil {
		var ret int32
		return ret
	}
	return *o.DaemonThreads
}

// GetDaemonThreadsOk returns a tuple with the DaemonThreads field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemDiagnosticsSnapshotDTO) GetDaemonThreadsOk() (*int32, bool) {
	if o == nil || o.DaemonThreads == nil {
		return nil, false
	}
	return o.DaemonThreads, true
}

// HasDaemonThreads returns a boolean if a field has been set.
func (o *SystemDiagnosticsSnapshotDTO) HasDaemonThreads() bool {
	if o != nil && o.DaemonThreads != nil {
		return true
	}

	return false
}

// SetDaemonThreads gets a reference to the given int32 and assigns it to the DaemonThreads field.
func (o *SystemDiagnosticsSnapshotDTO) SetDaemonThreads(v int32) {
	o.DaemonThreads = &v
}

// GetUptime returns the Uptime field value if set, zero value otherwise.
func (o *SystemDiagnosticsSnapshotDTO) GetUptime() string {
	if o == nil || o.Uptime == nil {
		var ret string
		return ret
	}
	return *o.Uptime
}

// GetUptimeOk returns a tuple with the Uptime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemDiagnosticsSnapshotDTO) GetUptimeOk() (*string, bool) {
	if o == nil || o.Uptime == nil {
		return nil, false
	}
	return o.Uptime, true
}

// HasUptime returns a boolean if a field has been set.
func (o *SystemDiagnosticsSnapshotDTO) HasUptime() bool {
	if o != nil && o.Uptime != nil {
		return true
	}

	return false
}

// SetUptime gets a reference to the given string and assigns it to the Uptime field.
func (o *SystemDiagnosticsSnapshotDTO) SetUptime(v string) {
	o.Uptime = &v
}

// GetFlowFileRepositoryStorageUsage returns the FlowFileRepositoryStorageUsage field value if set, zero value otherwise.
func (o *SystemDiagnosticsSnapshotDTO) GetFlowFileRepositoryStorageUsage() StorageUsageDTO {
	if o == nil || o.FlowFileRepositoryStorageUsage == nil {
		var ret StorageUsageDTO
		return ret
	}
	return *o.FlowFileRepositoryStorageUsage
}

// GetFlowFileRepositoryStorageUsageOk returns a tuple with the FlowFileRepositoryStorageUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemDiagnosticsSnapshotDTO) GetFlowFileRepositoryStorageUsageOk() (*StorageUsageDTO, bool) {
	if o == nil || o.FlowFileRepositoryStorageUsage == nil {
		return nil, false
	}
	return o.FlowFileRepositoryStorageUsage, true
}

// HasFlowFileRepositoryStorageUsage returns a boolean if a field has been set.
func (o *SystemDiagnosticsSnapshotDTO) HasFlowFileRepositoryStorageUsage() bool {
	if o != nil && o.FlowFileRepositoryStorageUsage != nil {
		return true
	}

	return false
}

// SetFlowFileRepositoryStorageUsage gets a reference to the given StorageUsageDTO and assigns it to the FlowFileRepositoryStorageUsage field.
func (o *SystemDiagnosticsSnapshotDTO) SetFlowFileRepositoryStorageUsage(v StorageUsageDTO) {
	o.FlowFileRepositoryStorageUsage = &v
}

// GetContentRepositoryStorageUsage returns the ContentRepositoryStorageUsage field value if set, zero value otherwise.
func (o *SystemDiagnosticsSnapshotDTO) GetContentRepositoryStorageUsage() []StorageUsageDTO {
	if o == nil || o.ContentRepositoryStorageUsage == nil {
		var ret []StorageUsageDTO
		return ret
	}
	return *o.ContentRepositoryStorageUsage
}

// GetContentRepositoryStorageUsageOk returns a tuple with the ContentRepositoryStorageUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemDiagnosticsSnapshotDTO) GetContentRepositoryStorageUsageOk() (*[]StorageUsageDTO, bool) {
	if o == nil || o.ContentRepositoryStorageUsage == nil {
		return nil, false
	}
	return o.ContentRepositoryStorageUsage, true
}

// HasContentRepositoryStorageUsage returns a boolean if a field has been set.
func (o *SystemDiagnosticsSnapshotDTO) HasContentRepositoryStorageUsage() bool {
	if o != nil && o.ContentRepositoryStorageUsage != nil {
		return true
	}

	return false
}

// SetContentRepositoryStorageUsage gets a reference to the given []StorageUsageDTO and assigns it to the ContentRepositoryStorageUsage field.
func (o *SystemDiagnosticsSnapshotDTO) SetContentRepositoryStorageUsage(v []StorageUsageDTO) {
	o.ContentRepositoryStorageUsage = &v
}

// GetProvenanceRepositoryStorageUsage returns the ProvenanceRepositoryStorageUsage field value if set, zero value otherwise.
func (o *SystemDiagnosticsSnapshotDTO) GetProvenanceRepositoryStorageUsage() []StorageUsageDTO {
	if o == nil || o.ProvenanceRepositoryStorageUsage == nil {
		var ret []StorageUsageDTO
		return ret
	}
	return *o.ProvenanceRepositoryStorageUsage
}

// GetProvenanceRepositoryStorageUsageOk returns a tuple with the ProvenanceRepositoryStorageUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemDiagnosticsSnapshotDTO) GetProvenanceRepositoryStorageUsageOk() (*[]StorageUsageDTO, bool) {
	if o == nil || o.ProvenanceRepositoryStorageUsage == nil {
		return nil, false
	}
	return o.ProvenanceRepositoryStorageUsage, true
}

// HasProvenanceRepositoryStorageUsage returns a boolean if a field has been set.
func (o *SystemDiagnosticsSnapshotDTO) HasProvenanceRepositoryStorageUsage() bool {
	if o != nil && o.ProvenanceRepositoryStorageUsage != nil {
		return true
	}

	return false
}

// SetProvenanceRepositoryStorageUsage gets a reference to the given []StorageUsageDTO and assigns it to the ProvenanceRepositoryStorageUsage field.
func (o *SystemDiagnosticsSnapshotDTO) SetProvenanceRepositoryStorageUsage(v []StorageUsageDTO) {
	o.ProvenanceRepositoryStorageUsage = &v
}

// GetGarbageCollection returns the GarbageCollection field value if set, zero value otherwise.
func (o *SystemDiagnosticsSnapshotDTO) GetGarbageCollection() []GarbageCollectionDTO {
	if o == nil || o.GarbageCollection == nil {
		var ret []GarbageCollectionDTO
		return ret
	}
	return *o.GarbageCollection
}

// GetGarbageCollectionOk returns a tuple with the GarbageCollection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemDiagnosticsSnapshotDTO) GetGarbageCollectionOk() (*[]GarbageCollectionDTO, bool) {
	if o == nil || o.GarbageCollection == nil {
		return nil, false
	}
	return o.GarbageCollection, true
}

// HasGarbageCollection returns a boolean if a field has been set.
func (o *SystemDiagnosticsSnapshotDTO) HasGarbageCollection() bool {
	if o != nil && o.GarbageCollection != nil {
		return true
	}

	return false
}

// SetGarbageCollection gets a reference to the given []GarbageCollectionDTO and assigns it to the GarbageCollection field.
func (o *SystemDiagnosticsSnapshotDTO) SetGarbageCollection(v []GarbageCollectionDTO) {
	o.GarbageCollection = &v
}

// GetStatsLastRefreshed returns the StatsLastRefreshed field value if set, zero value otherwise.
func (o *SystemDiagnosticsSnapshotDTO) GetStatsLastRefreshed() string {
	if o == nil || o.StatsLastRefreshed == nil {
		var ret string
		return ret
	}
	return *o.StatsLastRefreshed
}

// GetStatsLastRefreshedOk returns a tuple with the StatsLastRefreshed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemDiagnosticsSnapshotDTO) GetStatsLastRefreshedOk() (*string, bool) {
	if o == nil || o.StatsLastRefreshed == nil {
		return nil, false
	}
	return o.StatsLastRefreshed, true
}

// HasStatsLastRefreshed returns a boolean if a field has been set.
func (o *SystemDiagnosticsSnapshotDTO) HasStatsLastRefreshed() bool {
	if o != nil && o.StatsLastRefreshed != nil {
		return true
	}

	return false
}

// SetStatsLastRefreshed gets a reference to the given string and assigns it to the StatsLastRefreshed field.
func (o *SystemDiagnosticsSnapshotDTO) SetStatsLastRefreshed(v string) {
	o.StatsLastRefreshed = &v
}

// GetVersionInfo returns the VersionInfo field value if set, zero value otherwise.
func (o *SystemDiagnosticsSnapshotDTO) GetVersionInfo() VersionInfoDTO {
	if o == nil || o.VersionInfo == nil {
		var ret VersionInfoDTO
		return ret
	}
	return *o.VersionInfo
}

// GetVersionInfoOk returns a tuple with the VersionInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemDiagnosticsSnapshotDTO) GetVersionInfoOk() (*VersionInfoDTO, bool) {
	if o == nil || o.VersionInfo == nil {
		return nil, false
	}
	return o.VersionInfo, true
}

// HasVersionInfo returns a boolean if a field has been set.
func (o *SystemDiagnosticsSnapshotDTO) HasVersionInfo() bool {
	if o != nil && o.VersionInfo != nil {
		return true
	}

	return false
}

// SetVersionInfo gets a reference to the given VersionInfoDTO and assigns it to the VersionInfo field.
func (o *SystemDiagnosticsSnapshotDTO) SetVersionInfo(v VersionInfoDTO) {
	o.VersionInfo = &v
}

func (o SystemDiagnosticsSnapshotDTO) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TotalNonHeap != nil {
		toSerialize["totalNonHeap"] = o.TotalNonHeap
	}
	if o.TotalNonHeapBytes != nil {
		toSerialize["totalNonHeapBytes"] = o.TotalNonHeapBytes
	}
	if o.UsedNonHeap != nil {
		toSerialize["usedNonHeap"] = o.UsedNonHeap
	}
	if o.UsedNonHeapBytes != nil {
		toSerialize["usedNonHeapBytes"] = o.UsedNonHeapBytes
	}
	if o.FreeNonHeap != nil {
		toSerialize["freeNonHeap"] = o.FreeNonHeap
	}
	if o.FreeNonHeapBytes != nil {
		toSerialize["freeNonHeapBytes"] = o.FreeNonHeapBytes
	}
	if o.MaxNonHeap != nil {
		toSerialize["maxNonHeap"] = o.MaxNonHeap
	}
	if o.MaxNonHeapBytes != nil {
		toSerialize["maxNonHeapBytes"] = o.MaxNonHeapBytes
	}
	if o.NonHeapUtilization != nil {
		toSerialize["nonHeapUtilization"] = o.NonHeapUtilization
	}
	if o.TotalHeap != nil {
		toSerialize["totalHeap"] = o.TotalHeap
	}
	if o.TotalHeapBytes != nil {
		toSerialize["totalHeapBytes"] = o.TotalHeapBytes
	}
	if o.UsedHeap != nil {
		toSerialize["usedHeap"] = o.UsedHeap
	}
	if o.UsedHeapBytes != nil {
		toSerialize["usedHeapBytes"] = o.UsedHeapBytes
	}
	if o.FreeHeap != nil {
		toSerialize["freeHeap"] = o.FreeHeap
	}
	if o.FreeHeapBytes != nil {
		toSerialize["freeHeapBytes"] = o.FreeHeapBytes
	}
	if o.MaxHeap != nil {
		toSerialize["maxHeap"] = o.MaxHeap
	}
	if o.MaxHeapBytes != nil {
		toSerialize["maxHeapBytes"] = o.MaxHeapBytes
	}
	if o.HeapUtilization != nil {
		toSerialize["heapUtilization"] = o.HeapUtilization
	}
	if o.AvailableProcessors != nil {
		toSerialize["availableProcessors"] = o.AvailableProcessors
	}
	if o.ProcessorLoadAverage != nil {
		toSerialize["processorLoadAverage"] = o.ProcessorLoadAverage
	}
	if o.TotalThreads != nil {
		toSerialize["totalThreads"] = o.TotalThreads
	}
	if o.DaemonThreads != nil {
		toSerialize["daemonThreads"] = o.DaemonThreads
	}
	if o.Uptime != nil {
		toSerialize["uptime"] = o.Uptime
	}
	if o.FlowFileRepositoryStorageUsage != nil {
		toSerialize["flowFileRepositoryStorageUsage"] = o.FlowFileRepositoryStorageUsage
	}
	if o.ContentRepositoryStorageUsage != nil {
		toSerialize["contentRepositoryStorageUsage"] = o.ContentRepositoryStorageUsage
	}
	if o.ProvenanceRepositoryStorageUsage != nil {
		toSerialize["provenanceRepositoryStorageUsage"] = o.ProvenanceRepositoryStorageUsage
	}
	if o.GarbageCollection != nil {
		toSerialize["garbageCollection"] = o.GarbageCollection
	}
	if o.StatsLastRefreshed != nil {
		toSerialize["statsLastRefreshed"] = o.StatsLastRefreshed
	}
	if o.VersionInfo != nil {
		toSerialize["versionInfo"] = o.VersionInfo
	}
	return json.Marshal(toSerialize)
}

type NullableSystemDiagnosticsSnapshotDTO struct {
	value *SystemDiagnosticsSnapshotDTO
	isSet bool
}

func (v NullableSystemDiagnosticsSnapshotDTO) Get() *SystemDiagnosticsSnapshotDTO {
	return v.value
}

func (v *NullableSystemDiagnosticsSnapshotDTO) Set(val *SystemDiagnosticsSnapshotDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemDiagnosticsSnapshotDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemDiagnosticsSnapshotDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemDiagnosticsSnapshotDTO(val *SystemDiagnosticsSnapshotDTO) *NullableSystemDiagnosticsSnapshotDTO {
	return &NullableSystemDiagnosticsSnapshotDTO{value: val, isSet: true}
}

func (v NullableSystemDiagnosticsSnapshotDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemDiagnosticsSnapshotDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
