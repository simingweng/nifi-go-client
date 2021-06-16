# ControllerStatusDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveThreadCount** | Pointer to **int32** | The number of active threads in the NiFi. | [optional] 
**TerminatedThreadCount** | Pointer to **int32** | The number of terminated threads in the NiFi. | [optional] 
**Queued** | Pointer to **string** | The number of flowfiles queued in the NiFi. | [optional] 
**FlowFilesQueued** | Pointer to **int32** | The number of FlowFiles queued across the entire flow | [optional] 
**BytesQueued** | Pointer to **int64** | The size of the FlowFiles queued across the entire flow | [optional] 
**RunningCount** | Pointer to **int32** | The number of running components in the NiFi. | [optional] 
**StoppedCount** | Pointer to **int32** | The number of stopped components in the NiFi. | [optional] 
**InvalidCount** | Pointer to **int32** | The number of invalid components in the NiFi. | [optional] 
**DisabledCount** | Pointer to **int32** | The number of disabled components in the NiFi. | [optional] 
**ActiveRemotePortCount** | Pointer to **int32** | The number of active remote ports in the NiFi. | [optional] 
**InactiveRemotePortCount** | Pointer to **int32** | The number of inactive remote ports in the NiFi. | [optional] 
**UpToDateCount** | Pointer to **int32** | The number of up to date versioned process groups in the NiFi. | [optional] 
**LocallyModifiedCount** | Pointer to **int32** | The number of locally modified versioned process groups in the NiFi. | [optional] 
**StaleCount** | Pointer to **int32** | The number of stale versioned process groups in the NiFi. | [optional] 
**LocallyModifiedAndStaleCount** | Pointer to **int32** | The number of locally modified and stale versioned process groups in the NiFi. | [optional] 
**SyncFailureCount** | Pointer to **int32** | The number of versioned process groups in the NiFi that are unable to sync to a registry. | [optional] 

## Methods

### NewControllerStatusDTO

`func NewControllerStatusDTO() *ControllerStatusDTO`

NewControllerStatusDTO instantiates a new ControllerStatusDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControllerStatusDTOWithDefaults

`func NewControllerStatusDTOWithDefaults() *ControllerStatusDTO`

NewControllerStatusDTOWithDefaults instantiates a new ControllerStatusDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveThreadCount

`func (o *ControllerStatusDTO) GetActiveThreadCount() int32`

GetActiveThreadCount returns the ActiveThreadCount field if non-nil, zero value otherwise.

### GetActiveThreadCountOk

`func (o *ControllerStatusDTO) GetActiveThreadCountOk() (*int32, bool)`

GetActiveThreadCountOk returns a tuple with the ActiveThreadCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveThreadCount

`func (o *ControllerStatusDTO) SetActiveThreadCount(v int32)`

SetActiveThreadCount sets ActiveThreadCount field to given value.

### HasActiveThreadCount

`func (o *ControllerStatusDTO) HasActiveThreadCount() bool`

HasActiveThreadCount returns a boolean if a field has been set.

### GetTerminatedThreadCount

`func (o *ControllerStatusDTO) GetTerminatedThreadCount() int32`

GetTerminatedThreadCount returns the TerminatedThreadCount field if non-nil, zero value otherwise.

### GetTerminatedThreadCountOk

`func (o *ControllerStatusDTO) GetTerminatedThreadCountOk() (*int32, bool)`

GetTerminatedThreadCountOk returns a tuple with the TerminatedThreadCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminatedThreadCount

`func (o *ControllerStatusDTO) SetTerminatedThreadCount(v int32)`

SetTerminatedThreadCount sets TerminatedThreadCount field to given value.

### HasTerminatedThreadCount

`func (o *ControllerStatusDTO) HasTerminatedThreadCount() bool`

HasTerminatedThreadCount returns a boolean if a field has been set.

### GetQueued

`func (o *ControllerStatusDTO) GetQueued() string`

GetQueued returns the Queued field if non-nil, zero value otherwise.

### GetQueuedOk

`func (o *ControllerStatusDTO) GetQueuedOk() (*string, bool)`

GetQueuedOk returns a tuple with the Queued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueued

`func (o *ControllerStatusDTO) SetQueued(v string)`

SetQueued sets Queued field to given value.

### HasQueued

`func (o *ControllerStatusDTO) HasQueued() bool`

HasQueued returns a boolean if a field has been set.

### GetFlowFilesQueued

`func (o *ControllerStatusDTO) GetFlowFilesQueued() int32`

GetFlowFilesQueued returns the FlowFilesQueued field if non-nil, zero value otherwise.

### GetFlowFilesQueuedOk

`func (o *ControllerStatusDTO) GetFlowFilesQueuedOk() (*int32, bool)`

GetFlowFilesQueuedOk returns a tuple with the FlowFilesQueued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowFilesQueued

`func (o *ControllerStatusDTO) SetFlowFilesQueued(v int32)`

SetFlowFilesQueued sets FlowFilesQueued field to given value.

### HasFlowFilesQueued

`func (o *ControllerStatusDTO) HasFlowFilesQueued() bool`

HasFlowFilesQueued returns a boolean if a field has been set.

### GetBytesQueued

`func (o *ControllerStatusDTO) GetBytesQueued() int64`

GetBytesQueued returns the BytesQueued field if non-nil, zero value otherwise.

### GetBytesQueuedOk

`func (o *ControllerStatusDTO) GetBytesQueuedOk() (*int64, bool)`

GetBytesQueuedOk returns a tuple with the BytesQueued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesQueued

`func (o *ControllerStatusDTO) SetBytesQueued(v int64)`

SetBytesQueued sets BytesQueued field to given value.

### HasBytesQueued

`func (o *ControllerStatusDTO) HasBytesQueued() bool`

HasBytesQueued returns a boolean if a field has been set.

### GetRunningCount

`func (o *ControllerStatusDTO) GetRunningCount() int32`

GetRunningCount returns the RunningCount field if non-nil, zero value otherwise.

### GetRunningCountOk

`func (o *ControllerStatusDTO) GetRunningCountOk() (*int32, bool)`

GetRunningCountOk returns a tuple with the RunningCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningCount

`func (o *ControllerStatusDTO) SetRunningCount(v int32)`

SetRunningCount sets RunningCount field to given value.

### HasRunningCount

`func (o *ControllerStatusDTO) HasRunningCount() bool`

HasRunningCount returns a boolean if a field has been set.

### GetStoppedCount

`func (o *ControllerStatusDTO) GetStoppedCount() int32`

GetStoppedCount returns the StoppedCount field if non-nil, zero value otherwise.

### GetStoppedCountOk

`func (o *ControllerStatusDTO) GetStoppedCountOk() (*int32, bool)`

GetStoppedCountOk returns a tuple with the StoppedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoppedCount

`func (o *ControllerStatusDTO) SetStoppedCount(v int32)`

SetStoppedCount sets StoppedCount field to given value.

### HasStoppedCount

`func (o *ControllerStatusDTO) HasStoppedCount() bool`

HasStoppedCount returns a boolean if a field has been set.

### GetInvalidCount

`func (o *ControllerStatusDTO) GetInvalidCount() int32`

GetInvalidCount returns the InvalidCount field if non-nil, zero value otherwise.

### GetInvalidCountOk

`func (o *ControllerStatusDTO) GetInvalidCountOk() (*int32, bool)`

GetInvalidCountOk returns a tuple with the InvalidCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidCount

`func (o *ControllerStatusDTO) SetInvalidCount(v int32)`

SetInvalidCount sets InvalidCount field to given value.

### HasInvalidCount

`func (o *ControllerStatusDTO) HasInvalidCount() bool`

HasInvalidCount returns a boolean if a field has been set.

### GetDisabledCount

`func (o *ControllerStatusDTO) GetDisabledCount() int32`

GetDisabledCount returns the DisabledCount field if non-nil, zero value otherwise.

### GetDisabledCountOk

`func (o *ControllerStatusDTO) GetDisabledCountOk() (*int32, bool)`

GetDisabledCountOk returns a tuple with the DisabledCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledCount

`func (o *ControllerStatusDTO) SetDisabledCount(v int32)`

SetDisabledCount sets DisabledCount field to given value.

### HasDisabledCount

`func (o *ControllerStatusDTO) HasDisabledCount() bool`

HasDisabledCount returns a boolean if a field has been set.

### GetActiveRemotePortCount

`func (o *ControllerStatusDTO) GetActiveRemotePortCount() int32`

GetActiveRemotePortCount returns the ActiveRemotePortCount field if non-nil, zero value otherwise.

### GetActiveRemotePortCountOk

`func (o *ControllerStatusDTO) GetActiveRemotePortCountOk() (*int32, bool)`

GetActiveRemotePortCountOk returns a tuple with the ActiveRemotePortCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveRemotePortCount

`func (o *ControllerStatusDTO) SetActiveRemotePortCount(v int32)`

SetActiveRemotePortCount sets ActiveRemotePortCount field to given value.

### HasActiveRemotePortCount

`func (o *ControllerStatusDTO) HasActiveRemotePortCount() bool`

HasActiveRemotePortCount returns a boolean if a field has been set.

### GetInactiveRemotePortCount

`func (o *ControllerStatusDTO) GetInactiveRemotePortCount() int32`

GetInactiveRemotePortCount returns the InactiveRemotePortCount field if non-nil, zero value otherwise.

### GetInactiveRemotePortCountOk

`func (o *ControllerStatusDTO) GetInactiveRemotePortCountOk() (*int32, bool)`

GetInactiveRemotePortCountOk returns a tuple with the InactiveRemotePortCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveRemotePortCount

`func (o *ControllerStatusDTO) SetInactiveRemotePortCount(v int32)`

SetInactiveRemotePortCount sets InactiveRemotePortCount field to given value.

### HasInactiveRemotePortCount

`func (o *ControllerStatusDTO) HasInactiveRemotePortCount() bool`

HasInactiveRemotePortCount returns a boolean if a field has been set.

### GetUpToDateCount

`func (o *ControllerStatusDTO) GetUpToDateCount() int32`

GetUpToDateCount returns the UpToDateCount field if non-nil, zero value otherwise.

### GetUpToDateCountOk

`func (o *ControllerStatusDTO) GetUpToDateCountOk() (*int32, bool)`

GetUpToDateCountOk returns a tuple with the UpToDateCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpToDateCount

`func (o *ControllerStatusDTO) SetUpToDateCount(v int32)`

SetUpToDateCount sets UpToDateCount field to given value.

### HasUpToDateCount

`func (o *ControllerStatusDTO) HasUpToDateCount() bool`

HasUpToDateCount returns a boolean if a field has been set.

### GetLocallyModifiedCount

`func (o *ControllerStatusDTO) GetLocallyModifiedCount() int32`

GetLocallyModifiedCount returns the LocallyModifiedCount field if non-nil, zero value otherwise.

### GetLocallyModifiedCountOk

`func (o *ControllerStatusDTO) GetLocallyModifiedCountOk() (*int32, bool)`

GetLocallyModifiedCountOk returns a tuple with the LocallyModifiedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocallyModifiedCount

`func (o *ControllerStatusDTO) SetLocallyModifiedCount(v int32)`

SetLocallyModifiedCount sets LocallyModifiedCount field to given value.

### HasLocallyModifiedCount

`func (o *ControllerStatusDTO) HasLocallyModifiedCount() bool`

HasLocallyModifiedCount returns a boolean if a field has been set.

### GetStaleCount

`func (o *ControllerStatusDTO) GetStaleCount() int32`

GetStaleCount returns the StaleCount field if non-nil, zero value otherwise.

### GetStaleCountOk

`func (o *ControllerStatusDTO) GetStaleCountOk() (*int32, bool)`

GetStaleCountOk returns a tuple with the StaleCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaleCount

`func (o *ControllerStatusDTO) SetStaleCount(v int32)`

SetStaleCount sets StaleCount field to given value.

### HasStaleCount

`func (o *ControllerStatusDTO) HasStaleCount() bool`

HasStaleCount returns a boolean if a field has been set.

### GetLocallyModifiedAndStaleCount

`func (o *ControllerStatusDTO) GetLocallyModifiedAndStaleCount() int32`

GetLocallyModifiedAndStaleCount returns the LocallyModifiedAndStaleCount field if non-nil, zero value otherwise.

### GetLocallyModifiedAndStaleCountOk

`func (o *ControllerStatusDTO) GetLocallyModifiedAndStaleCountOk() (*int32, bool)`

GetLocallyModifiedAndStaleCountOk returns a tuple with the LocallyModifiedAndStaleCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocallyModifiedAndStaleCount

`func (o *ControllerStatusDTO) SetLocallyModifiedAndStaleCount(v int32)`

SetLocallyModifiedAndStaleCount sets LocallyModifiedAndStaleCount field to given value.

### HasLocallyModifiedAndStaleCount

`func (o *ControllerStatusDTO) HasLocallyModifiedAndStaleCount() bool`

HasLocallyModifiedAndStaleCount returns a boolean if a field has been set.

### GetSyncFailureCount

`func (o *ControllerStatusDTO) GetSyncFailureCount() int32`

GetSyncFailureCount returns the SyncFailureCount field if non-nil, zero value otherwise.

### GetSyncFailureCountOk

`func (o *ControllerStatusDTO) GetSyncFailureCountOk() (*int32, bool)`

GetSyncFailureCountOk returns a tuple with the SyncFailureCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncFailureCount

`func (o *ControllerStatusDTO) SetSyncFailureCount(v int32)`

SetSyncFailureCount sets SyncFailureCount field to given value.

### HasSyncFailureCount

`func (o *ControllerStatusDTO) HasSyncFailureCount() bool`

HasSyncFailureCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


