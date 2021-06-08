# ProcessGroupStatusSnapshotDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The id of the process group. | [optional] 
**Name** | Pointer to **string** | The name of this process group. | [optional] 
**ConnectionStatusSnapshots** | Pointer to [**[]ConnectionStatusSnapshotEntity**](ConnectionStatusSnapshotEntity.md) | The status of all connections in the process group. | [optional] 
**ProcessorStatusSnapshots** | Pointer to [**[]ProcessorStatusSnapshotEntity**](ProcessorStatusSnapshotEntity.md) | The status of all processors in the process group. | [optional] 
**ProcessGroupStatusSnapshots** | Pointer to [**[]ProcessGroupStatusSnapshotEntity**](ProcessGroupStatusSnapshotEntity.md) | The status of all process groups in the process group. | [optional] 
**RemoteProcessGroupStatusSnapshots** | Pointer to [**[]RemoteProcessGroupStatusSnapshotEntity**](RemoteProcessGroupStatusSnapshotEntity.md) | The status of all remote process groups in the process group. | [optional] 
**InputPortStatusSnapshots** | Pointer to [**[]PortStatusSnapshotEntity**](PortStatusSnapshotEntity.md) | The status of all input ports in the process group. | [optional] 
**OutputPortStatusSnapshots** | Pointer to [**[]PortStatusSnapshotEntity**](PortStatusSnapshotEntity.md) | The status of all output ports in the process group. | [optional] 
**VersionedFlowState** | Pointer to **string** | The current state of the Process Group, as it relates to the Versioned Flow | [optional] [readonly] 
**FlowFilesIn** | Pointer to **int32** | The number of FlowFiles that have come into this ProcessGroup in the last 5 minutes | [optional] 
**BytesIn** | Pointer to **int64** | The number of bytes that have come into this ProcessGroup in the last 5 minutes | [optional] 
**Input** | Pointer to **string** | The input count/size for the process group in the last 5 minutes (pretty printed). | [optional] 
**FlowFilesQueued** | Pointer to **int32** | The number of FlowFiles that are queued up in this ProcessGroup right now | [optional] 
**BytesQueued** | Pointer to **int64** | The number of bytes that are queued up in this ProcessGroup right now | [optional] 
**Queued** | Pointer to **string** | The count/size that is queued in the the process group. | [optional] 
**QueuedCount** | Pointer to **string** | The count that is queued for the process group. | [optional] 
**QueuedSize** | Pointer to **string** | The size that is queued for the process group. | [optional] 
**BytesRead** | Pointer to **int64** | The number of bytes read by components in this ProcessGroup in the last 5 minutes | [optional] 
**Read** | Pointer to **string** | The number of bytes read in the last 5 minutes. | [optional] 
**BytesWritten** | Pointer to **int64** | The number of bytes written by components in this ProcessGroup in the last 5 minutes | [optional] 
**Written** | Pointer to **string** | The number of bytes written in the last 5 minutes. | [optional] 
**FlowFilesOut** | Pointer to **int32** | The number of FlowFiles transferred out of this ProcessGroup in the last 5 minutes | [optional] 
**BytesOut** | Pointer to **int64** | The number of bytes transferred out of this ProcessGroup in the last 5 minutes | [optional] 
**Output** | Pointer to **string** | The output count/size for the process group in the last 5 minutes. | [optional] 
**FlowFilesTransferred** | Pointer to **int32** | The number of FlowFiles transferred in this ProcessGroup in the last 5 minutes | [optional] 
**BytesTransferred** | Pointer to **int64** | The number of bytes transferred in this ProcessGroup in the last 5 minutes | [optional] 
**Transferred** | Pointer to **string** | The count/size transferred to/from queues in the process group in the last 5 minutes. | [optional] 
**BytesReceived** | Pointer to **int64** | The number of bytes received from external sources by components within this ProcessGroup in the last 5 minutes | [optional] 
**FlowFilesReceived** | Pointer to **int32** | The number of FlowFiles received from external sources by components within this ProcessGroup in the last 5 minutes | [optional] 
**Received** | Pointer to **string** | The count/size sent to the process group in the last 5 minutes. | [optional] 
**BytesSent** | Pointer to **int64** | The number of bytes sent to an external sink by components within this ProcessGroup in the last 5 minutes | [optional] 
**FlowFilesSent** | Pointer to **int32** | The number of FlowFiles sent to an external sink by components within this ProcessGroup in the last 5 minutes | [optional] 
**Sent** | Pointer to **string** | The count/size sent from this process group in the last 5 minutes. | [optional] 
**ActiveThreadCount** | Pointer to **int32** | The active thread count for this process group. | [optional] 
**TerminatedThreadCount** | Pointer to **int32** | The number of threads currently terminated for the process group. | [optional] 

## Methods

### NewProcessGroupStatusSnapshotDTO

`func NewProcessGroupStatusSnapshotDTO() *ProcessGroupStatusSnapshotDTO`

NewProcessGroupStatusSnapshotDTO instantiates a new ProcessGroupStatusSnapshotDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessGroupStatusSnapshotDTOWithDefaults

`func NewProcessGroupStatusSnapshotDTOWithDefaults() *ProcessGroupStatusSnapshotDTO`

NewProcessGroupStatusSnapshotDTOWithDefaults instantiates a new ProcessGroupStatusSnapshotDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProcessGroupStatusSnapshotDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProcessGroupStatusSnapshotDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProcessGroupStatusSnapshotDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProcessGroupStatusSnapshotDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ProcessGroupStatusSnapshotDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProcessGroupStatusSnapshotDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProcessGroupStatusSnapshotDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProcessGroupStatusSnapshotDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetConnectionStatusSnapshots

`func (o *ProcessGroupStatusSnapshotDTO) GetConnectionStatusSnapshots() []ConnectionStatusSnapshotEntity`

GetConnectionStatusSnapshots returns the ConnectionStatusSnapshots field if non-nil, zero value otherwise.

### GetConnectionStatusSnapshotsOk

`func (o *ProcessGroupStatusSnapshotDTO) GetConnectionStatusSnapshotsOk() (*[]ConnectionStatusSnapshotEntity, bool)`

GetConnectionStatusSnapshotsOk returns a tuple with the ConnectionStatusSnapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionStatusSnapshots

`func (o *ProcessGroupStatusSnapshotDTO) SetConnectionStatusSnapshots(v []ConnectionStatusSnapshotEntity)`

SetConnectionStatusSnapshots sets ConnectionStatusSnapshots field to given value.

### HasConnectionStatusSnapshots

`func (o *ProcessGroupStatusSnapshotDTO) HasConnectionStatusSnapshots() bool`

HasConnectionStatusSnapshots returns a boolean if a field has been set.

### GetProcessorStatusSnapshots

`func (o *ProcessGroupStatusSnapshotDTO) GetProcessorStatusSnapshots() []ProcessorStatusSnapshotEntity`

GetProcessorStatusSnapshots returns the ProcessorStatusSnapshots field if non-nil, zero value otherwise.

### GetProcessorStatusSnapshotsOk

`func (o *ProcessGroupStatusSnapshotDTO) GetProcessorStatusSnapshotsOk() (*[]ProcessorStatusSnapshotEntity, bool)`

GetProcessorStatusSnapshotsOk returns a tuple with the ProcessorStatusSnapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessorStatusSnapshots

`func (o *ProcessGroupStatusSnapshotDTO) SetProcessorStatusSnapshots(v []ProcessorStatusSnapshotEntity)`

SetProcessorStatusSnapshots sets ProcessorStatusSnapshots field to given value.

### HasProcessorStatusSnapshots

`func (o *ProcessGroupStatusSnapshotDTO) HasProcessorStatusSnapshots() bool`

HasProcessorStatusSnapshots returns a boolean if a field has been set.

### GetProcessGroupStatusSnapshots

`func (o *ProcessGroupStatusSnapshotDTO) GetProcessGroupStatusSnapshots() []ProcessGroupStatusSnapshotEntity`

GetProcessGroupStatusSnapshots returns the ProcessGroupStatusSnapshots field if non-nil, zero value otherwise.

### GetProcessGroupStatusSnapshotsOk

`func (o *ProcessGroupStatusSnapshotDTO) GetProcessGroupStatusSnapshotsOk() (*[]ProcessGroupStatusSnapshotEntity, bool)`

GetProcessGroupStatusSnapshotsOk returns a tuple with the ProcessGroupStatusSnapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessGroupStatusSnapshots

`func (o *ProcessGroupStatusSnapshotDTO) SetProcessGroupStatusSnapshots(v []ProcessGroupStatusSnapshotEntity)`

SetProcessGroupStatusSnapshots sets ProcessGroupStatusSnapshots field to given value.

### HasProcessGroupStatusSnapshots

`func (o *ProcessGroupStatusSnapshotDTO) HasProcessGroupStatusSnapshots() bool`

HasProcessGroupStatusSnapshots returns a boolean if a field has been set.

### GetRemoteProcessGroupStatusSnapshots

`func (o *ProcessGroupStatusSnapshotDTO) GetRemoteProcessGroupStatusSnapshots() []RemoteProcessGroupStatusSnapshotEntity`

GetRemoteProcessGroupStatusSnapshots returns the RemoteProcessGroupStatusSnapshots field if non-nil, zero value otherwise.

### GetRemoteProcessGroupStatusSnapshotsOk

`func (o *ProcessGroupStatusSnapshotDTO) GetRemoteProcessGroupStatusSnapshotsOk() (*[]RemoteProcessGroupStatusSnapshotEntity, bool)`

GetRemoteProcessGroupStatusSnapshotsOk returns a tuple with the RemoteProcessGroupStatusSnapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteProcessGroupStatusSnapshots

`func (o *ProcessGroupStatusSnapshotDTO) SetRemoteProcessGroupStatusSnapshots(v []RemoteProcessGroupStatusSnapshotEntity)`

SetRemoteProcessGroupStatusSnapshots sets RemoteProcessGroupStatusSnapshots field to given value.

### HasRemoteProcessGroupStatusSnapshots

`func (o *ProcessGroupStatusSnapshotDTO) HasRemoteProcessGroupStatusSnapshots() bool`

HasRemoteProcessGroupStatusSnapshots returns a boolean if a field has been set.

### GetInputPortStatusSnapshots

`func (o *ProcessGroupStatusSnapshotDTO) GetInputPortStatusSnapshots() []PortStatusSnapshotEntity`

GetInputPortStatusSnapshots returns the InputPortStatusSnapshots field if non-nil, zero value otherwise.

### GetInputPortStatusSnapshotsOk

`func (o *ProcessGroupStatusSnapshotDTO) GetInputPortStatusSnapshotsOk() (*[]PortStatusSnapshotEntity, bool)`

GetInputPortStatusSnapshotsOk returns a tuple with the InputPortStatusSnapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputPortStatusSnapshots

`func (o *ProcessGroupStatusSnapshotDTO) SetInputPortStatusSnapshots(v []PortStatusSnapshotEntity)`

SetInputPortStatusSnapshots sets InputPortStatusSnapshots field to given value.

### HasInputPortStatusSnapshots

`func (o *ProcessGroupStatusSnapshotDTO) HasInputPortStatusSnapshots() bool`

HasInputPortStatusSnapshots returns a boolean if a field has been set.

### GetOutputPortStatusSnapshots

`func (o *ProcessGroupStatusSnapshotDTO) GetOutputPortStatusSnapshots() []PortStatusSnapshotEntity`

GetOutputPortStatusSnapshots returns the OutputPortStatusSnapshots field if non-nil, zero value otherwise.

### GetOutputPortStatusSnapshotsOk

`func (o *ProcessGroupStatusSnapshotDTO) GetOutputPortStatusSnapshotsOk() (*[]PortStatusSnapshotEntity, bool)`

GetOutputPortStatusSnapshotsOk returns a tuple with the OutputPortStatusSnapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputPortStatusSnapshots

`func (o *ProcessGroupStatusSnapshotDTO) SetOutputPortStatusSnapshots(v []PortStatusSnapshotEntity)`

SetOutputPortStatusSnapshots sets OutputPortStatusSnapshots field to given value.

### HasOutputPortStatusSnapshots

`func (o *ProcessGroupStatusSnapshotDTO) HasOutputPortStatusSnapshots() bool`

HasOutputPortStatusSnapshots returns a boolean if a field has been set.

### GetVersionedFlowState

`func (o *ProcessGroupStatusSnapshotDTO) GetVersionedFlowState() string`

GetVersionedFlowState returns the VersionedFlowState field if non-nil, zero value otherwise.

### GetVersionedFlowStateOk

`func (o *ProcessGroupStatusSnapshotDTO) GetVersionedFlowStateOk() (*string, bool)`

GetVersionedFlowStateOk returns a tuple with the VersionedFlowState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionedFlowState

`func (o *ProcessGroupStatusSnapshotDTO) SetVersionedFlowState(v string)`

SetVersionedFlowState sets VersionedFlowState field to given value.

### HasVersionedFlowState

`func (o *ProcessGroupStatusSnapshotDTO) HasVersionedFlowState() bool`

HasVersionedFlowState returns a boolean if a field has been set.

### GetFlowFilesIn

`func (o *ProcessGroupStatusSnapshotDTO) GetFlowFilesIn() int32`

GetFlowFilesIn returns the FlowFilesIn field if non-nil, zero value otherwise.

### GetFlowFilesInOk

`func (o *ProcessGroupStatusSnapshotDTO) GetFlowFilesInOk() (*int32, bool)`

GetFlowFilesInOk returns a tuple with the FlowFilesIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowFilesIn

`func (o *ProcessGroupStatusSnapshotDTO) SetFlowFilesIn(v int32)`

SetFlowFilesIn sets FlowFilesIn field to given value.

### HasFlowFilesIn

`func (o *ProcessGroupStatusSnapshotDTO) HasFlowFilesIn() bool`

HasFlowFilesIn returns a boolean if a field has been set.

### GetBytesIn

`func (o *ProcessGroupStatusSnapshotDTO) GetBytesIn() int64`

GetBytesIn returns the BytesIn field if non-nil, zero value otherwise.

### GetBytesInOk

`func (o *ProcessGroupStatusSnapshotDTO) GetBytesInOk() (*int64, bool)`

GetBytesInOk returns a tuple with the BytesIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesIn

`func (o *ProcessGroupStatusSnapshotDTO) SetBytesIn(v int64)`

SetBytesIn sets BytesIn field to given value.

### HasBytesIn

`func (o *ProcessGroupStatusSnapshotDTO) HasBytesIn() bool`

HasBytesIn returns a boolean if a field has been set.

### GetInput

`func (o *ProcessGroupStatusSnapshotDTO) GetInput() string`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *ProcessGroupStatusSnapshotDTO) GetInputOk() (*string, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *ProcessGroupStatusSnapshotDTO) SetInput(v string)`

SetInput sets Input field to given value.

### HasInput

`func (o *ProcessGroupStatusSnapshotDTO) HasInput() bool`

HasInput returns a boolean if a field has been set.

### GetFlowFilesQueued

`func (o *ProcessGroupStatusSnapshotDTO) GetFlowFilesQueued() int32`

GetFlowFilesQueued returns the FlowFilesQueued field if non-nil, zero value otherwise.

### GetFlowFilesQueuedOk

`func (o *ProcessGroupStatusSnapshotDTO) GetFlowFilesQueuedOk() (*int32, bool)`

GetFlowFilesQueuedOk returns a tuple with the FlowFilesQueued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowFilesQueued

`func (o *ProcessGroupStatusSnapshotDTO) SetFlowFilesQueued(v int32)`

SetFlowFilesQueued sets FlowFilesQueued field to given value.

### HasFlowFilesQueued

`func (o *ProcessGroupStatusSnapshotDTO) HasFlowFilesQueued() bool`

HasFlowFilesQueued returns a boolean if a field has been set.

### GetBytesQueued

`func (o *ProcessGroupStatusSnapshotDTO) GetBytesQueued() int64`

GetBytesQueued returns the BytesQueued field if non-nil, zero value otherwise.

### GetBytesQueuedOk

`func (o *ProcessGroupStatusSnapshotDTO) GetBytesQueuedOk() (*int64, bool)`

GetBytesQueuedOk returns a tuple with the BytesQueued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesQueued

`func (o *ProcessGroupStatusSnapshotDTO) SetBytesQueued(v int64)`

SetBytesQueued sets BytesQueued field to given value.

### HasBytesQueued

`func (o *ProcessGroupStatusSnapshotDTO) HasBytesQueued() bool`

HasBytesQueued returns a boolean if a field has been set.

### GetQueued

`func (o *ProcessGroupStatusSnapshotDTO) GetQueued() string`

GetQueued returns the Queued field if non-nil, zero value otherwise.

### GetQueuedOk

`func (o *ProcessGroupStatusSnapshotDTO) GetQueuedOk() (*string, bool)`

GetQueuedOk returns a tuple with the Queued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueued

`func (o *ProcessGroupStatusSnapshotDTO) SetQueued(v string)`

SetQueued sets Queued field to given value.

### HasQueued

`func (o *ProcessGroupStatusSnapshotDTO) HasQueued() bool`

HasQueued returns a boolean if a field has been set.

### GetQueuedCount

`func (o *ProcessGroupStatusSnapshotDTO) GetQueuedCount() string`

GetQueuedCount returns the QueuedCount field if non-nil, zero value otherwise.

### GetQueuedCountOk

`func (o *ProcessGroupStatusSnapshotDTO) GetQueuedCountOk() (*string, bool)`

GetQueuedCountOk returns a tuple with the QueuedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueuedCount

`func (o *ProcessGroupStatusSnapshotDTO) SetQueuedCount(v string)`

SetQueuedCount sets QueuedCount field to given value.

### HasQueuedCount

`func (o *ProcessGroupStatusSnapshotDTO) HasQueuedCount() bool`

HasQueuedCount returns a boolean if a field has been set.

### GetQueuedSize

`func (o *ProcessGroupStatusSnapshotDTO) GetQueuedSize() string`

GetQueuedSize returns the QueuedSize field if non-nil, zero value otherwise.

### GetQueuedSizeOk

`func (o *ProcessGroupStatusSnapshotDTO) GetQueuedSizeOk() (*string, bool)`

GetQueuedSizeOk returns a tuple with the QueuedSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueuedSize

`func (o *ProcessGroupStatusSnapshotDTO) SetQueuedSize(v string)`

SetQueuedSize sets QueuedSize field to given value.

### HasQueuedSize

`func (o *ProcessGroupStatusSnapshotDTO) HasQueuedSize() bool`

HasQueuedSize returns a boolean if a field has been set.

### GetBytesRead

`func (o *ProcessGroupStatusSnapshotDTO) GetBytesRead() int64`

GetBytesRead returns the BytesRead field if non-nil, zero value otherwise.

### GetBytesReadOk

`func (o *ProcessGroupStatusSnapshotDTO) GetBytesReadOk() (*int64, bool)`

GetBytesReadOk returns a tuple with the BytesRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesRead

`func (o *ProcessGroupStatusSnapshotDTO) SetBytesRead(v int64)`

SetBytesRead sets BytesRead field to given value.

### HasBytesRead

`func (o *ProcessGroupStatusSnapshotDTO) HasBytesRead() bool`

HasBytesRead returns a boolean if a field has been set.

### GetRead

`func (o *ProcessGroupStatusSnapshotDTO) GetRead() string`

GetRead returns the Read field if non-nil, zero value otherwise.

### GetReadOk

`func (o *ProcessGroupStatusSnapshotDTO) GetReadOk() (*string, bool)`

GetReadOk returns a tuple with the Read field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRead

`func (o *ProcessGroupStatusSnapshotDTO) SetRead(v string)`

SetRead sets Read field to given value.

### HasRead

`func (o *ProcessGroupStatusSnapshotDTO) HasRead() bool`

HasRead returns a boolean if a field has been set.

### GetBytesWritten

`func (o *ProcessGroupStatusSnapshotDTO) GetBytesWritten() int64`

GetBytesWritten returns the BytesWritten field if non-nil, zero value otherwise.

### GetBytesWrittenOk

`func (o *ProcessGroupStatusSnapshotDTO) GetBytesWrittenOk() (*int64, bool)`

GetBytesWrittenOk returns a tuple with the BytesWritten field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesWritten

`func (o *ProcessGroupStatusSnapshotDTO) SetBytesWritten(v int64)`

SetBytesWritten sets BytesWritten field to given value.

### HasBytesWritten

`func (o *ProcessGroupStatusSnapshotDTO) HasBytesWritten() bool`

HasBytesWritten returns a boolean if a field has been set.

### GetWritten

`func (o *ProcessGroupStatusSnapshotDTO) GetWritten() string`

GetWritten returns the Written field if non-nil, zero value otherwise.

### GetWrittenOk

`func (o *ProcessGroupStatusSnapshotDTO) GetWrittenOk() (*string, bool)`

GetWrittenOk returns a tuple with the Written field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWritten

`func (o *ProcessGroupStatusSnapshotDTO) SetWritten(v string)`

SetWritten sets Written field to given value.

### HasWritten

`func (o *ProcessGroupStatusSnapshotDTO) HasWritten() bool`

HasWritten returns a boolean if a field has been set.

### GetFlowFilesOut

`func (o *ProcessGroupStatusSnapshotDTO) GetFlowFilesOut() int32`

GetFlowFilesOut returns the FlowFilesOut field if non-nil, zero value otherwise.

### GetFlowFilesOutOk

`func (o *ProcessGroupStatusSnapshotDTO) GetFlowFilesOutOk() (*int32, bool)`

GetFlowFilesOutOk returns a tuple with the FlowFilesOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowFilesOut

`func (o *ProcessGroupStatusSnapshotDTO) SetFlowFilesOut(v int32)`

SetFlowFilesOut sets FlowFilesOut field to given value.

### HasFlowFilesOut

`func (o *ProcessGroupStatusSnapshotDTO) HasFlowFilesOut() bool`

HasFlowFilesOut returns a boolean if a field has been set.

### GetBytesOut

`func (o *ProcessGroupStatusSnapshotDTO) GetBytesOut() int64`

GetBytesOut returns the BytesOut field if non-nil, zero value otherwise.

### GetBytesOutOk

`func (o *ProcessGroupStatusSnapshotDTO) GetBytesOutOk() (*int64, bool)`

GetBytesOutOk returns a tuple with the BytesOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesOut

`func (o *ProcessGroupStatusSnapshotDTO) SetBytesOut(v int64)`

SetBytesOut sets BytesOut field to given value.

### HasBytesOut

`func (o *ProcessGroupStatusSnapshotDTO) HasBytesOut() bool`

HasBytesOut returns a boolean if a field has been set.

### GetOutput

`func (o *ProcessGroupStatusSnapshotDTO) GetOutput() string`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *ProcessGroupStatusSnapshotDTO) GetOutputOk() (*string, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *ProcessGroupStatusSnapshotDTO) SetOutput(v string)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *ProcessGroupStatusSnapshotDTO) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### GetFlowFilesTransferred

`func (o *ProcessGroupStatusSnapshotDTO) GetFlowFilesTransferred() int32`

GetFlowFilesTransferred returns the FlowFilesTransferred field if non-nil, zero value otherwise.

### GetFlowFilesTransferredOk

`func (o *ProcessGroupStatusSnapshotDTO) GetFlowFilesTransferredOk() (*int32, bool)`

GetFlowFilesTransferredOk returns a tuple with the FlowFilesTransferred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowFilesTransferred

`func (o *ProcessGroupStatusSnapshotDTO) SetFlowFilesTransferred(v int32)`

SetFlowFilesTransferred sets FlowFilesTransferred field to given value.

### HasFlowFilesTransferred

`func (o *ProcessGroupStatusSnapshotDTO) HasFlowFilesTransferred() bool`

HasFlowFilesTransferred returns a boolean if a field has been set.

### GetBytesTransferred

`func (o *ProcessGroupStatusSnapshotDTO) GetBytesTransferred() int64`

GetBytesTransferred returns the BytesTransferred field if non-nil, zero value otherwise.

### GetBytesTransferredOk

`func (o *ProcessGroupStatusSnapshotDTO) GetBytesTransferredOk() (*int64, bool)`

GetBytesTransferredOk returns a tuple with the BytesTransferred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesTransferred

`func (o *ProcessGroupStatusSnapshotDTO) SetBytesTransferred(v int64)`

SetBytesTransferred sets BytesTransferred field to given value.

### HasBytesTransferred

`func (o *ProcessGroupStatusSnapshotDTO) HasBytesTransferred() bool`

HasBytesTransferred returns a boolean if a field has been set.

### GetTransferred

`func (o *ProcessGroupStatusSnapshotDTO) GetTransferred() string`

GetTransferred returns the Transferred field if non-nil, zero value otherwise.

### GetTransferredOk

`func (o *ProcessGroupStatusSnapshotDTO) GetTransferredOk() (*string, bool)`

GetTransferredOk returns a tuple with the Transferred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferred

`func (o *ProcessGroupStatusSnapshotDTO) SetTransferred(v string)`

SetTransferred sets Transferred field to given value.

### HasTransferred

`func (o *ProcessGroupStatusSnapshotDTO) HasTransferred() bool`

HasTransferred returns a boolean if a field has been set.

### GetBytesReceived

`func (o *ProcessGroupStatusSnapshotDTO) GetBytesReceived() int64`

GetBytesReceived returns the BytesReceived field if non-nil, zero value otherwise.

### GetBytesReceivedOk

`func (o *ProcessGroupStatusSnapshotDTO) GetBytesReceivedOk() (*int64, bool)`

GetBytesReceivedOk returns a tuple with the BytesReceived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesReceived

`func (o *ProcessGroupStatusSnapshotDTO) SetBytesReceived(v int64)`

SetBytesReceived sets BytesReceived field to given value.

### HasBytesReceived

`func (o *ProcessGroupStatusSnapshotDTO) HasBytesReceived() bool`

HasBytesReceived returns a boolean if a field has been set.

### GetFlowFilesReceived

`func (o *ProcessGroupStatusSnapshotDTO) GetFlowFilesReceived() int32`

GetFlowFilesReceived returns the FlowFilesReceived field if non-nil, zero value otherwise.

### GetFlowFilesReceivedOk

`func (o *ProcessGroupStatusSnapshotDTO) GetFlowFilesReceivedOk() (*int32, bool)`

GetFlowFilesReceivedOk returns a tuple with the FlowFilesReceived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowFilesReceived

`func (o *ProcessGroupStatusSnapshotDTO) SetFlowFilesReceived(v int32)`

SetFlowFilesReceived sets FlowFilesReceived field to given value.

### HasFlowFilesReceived

`func (o *ProcessGroupStatusSnapshotDTO) HasFlowFilesReceived() bool`

HasFlowFilesReceived returns a boolean if a field has been set.

### GetReceived

`func (o *ProcessGroupStatusSnapshotDTO) GetReceived() string`

GetReceived returns the Received field if non-nil, zero value otherwise.

### GetReceivedOk

`func (o *ProcessGroupStatusSnapshotDTO) GetReceivedOk() (*string, bool)`

GetReceivedOk returns a tuple with the Received field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceived

`func (o *ProcessGroupStatusSnapshotDTO) SetReceived(v string)`

SetReceived sets Received field to given value.

### HasReceived

`func (o *ProcessGroupStatusSnapshotDTO) HasReceived() bool`

HasReceived returns a boolean if a field has been set.

### GetBytesSent

`func (o *ProcessGroupStatusSnapshotDTO) GetBytesSent() int64`

GetBytesSent returns the BytesSent field if non-nil, zero value otherwise.

### GetBytesSentOk

`func (o *ProcessGroupStatusSnapshotDTO) GetBytesSentOk() (*int64, bool)`

GetBytesSentOk returns a tuple with the BytesSent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesSent

`func (o *ProcessGroupStatusSnapshotDTO) SetBytesSent(v int64)`

SetBytesSent sets BytesSent field to given value.

### HasBytesSent

`func (o *ProcessGroupStatusSnapshotDTO) HasBytesSent() bool`

HasBytesSent returns a boolean if a field has been set.

### GetFlowFilesSent

`func (o *ProcessGroupStatusSnapshotDTO) GetFlowFilesSent() int32`

GetFlowFilesSent returns the FlowFilesSent field if non-nil, zero value otherwise.

### GetFlowFilesSentOk

`func (o *ProcessGroupStatusSnapshotDTO) GetFlowFilesSentOk() (*int32, bool)`

GetFlowFilesSentOk returns a tuple with the FlowFilesSent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowFilesSent

`func (o *ProcessGroupStatusSnapshotDTO) SetFlowFilesSent(v int32)`

SetFlowFilesSent sets FlowFilesSent field to given value.

### HasFlowFilesSent

`func (o *ProcessGroupStatusSnapshotDTO) HasFlowFilesSent() bool`

HasFlowFilesSent returns a boolean if a field has been set.

### GetSent

`func (o *ProcessGroupStatusSnapshotDTO) GetSent() string`

GetSent returns the Sent field if non-nil, zero value otherwise.

### GetSentOk

`func (o *ProcessGroupStatusSnapshotDTO) GetSentOk() (*string, bool)`

GetSentOk returns a tuple with the Sent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSent

`func (o *ProcessGroupStatusSnapshotDTO) SetSent(v string)`

SetSent sets Sent field to given value.

### HasSent

`func (o *ProcessGroupStatusSnapshotDTO) HasSent() bool`

HasSent returns a boolean if a field has been set.

### GetActiveThreadCount

`func (o *ProcessGroupStatusSnapshotDTO) GetActiveThreadCount() int32`

GetActiveThreadCount returns the ActiveThreadCount field if non-nil, zero value otherwise.

### GetActiveThreadCountOk

`func (o *ProcessGroupStatusSnapshotDTO) GetActiveThreadCountOk() (*int32, bool)`

GetActiveThreadCountOk returns a tuple with the ActiveThreadCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveThreadCount

`func (o *ProcessGroupStatusSnapshotDTO) SetActiveThreadCount(v int32)`

SetActiveThreadCount sets ActiveThreadCount field to given value.

### HasActiveThreadCount

`func (o *ProcessGroupStatusSnapshotDTO) HasActiveThreadCount() bool`

HasActiveThreadCount returns a boolean if a field has been set.

### GetTerminatedThreadCount

`func (o *ProcessGroupStatusSnapshotDTO) GetTerminatedThreadCount() int32`

GetTerminatedThreadCount returns the TerminatedThreadCount field if non-nil, zero value otherwise.

### GetTerminatedThreadCountOk

`func (o *ProcessGroupStatusSnapshotDTO) GetTerminatedThreadCountOk() (*int32, bool)`

GetTerminatedThreadCountOk returns a tuple with the TerminatedThreadCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminatedThreadCount

`func (o *ProcessGroupStatusSnapshotDTO) SetTerminatedThreadCount(v int32)`

SetTerminatedThreadCount sets TerminatedThreadCount field to given value.

### HasTerminatedThreadCount

`func (o *ProcessGroupStatusSnapshotDTO) HasTerminatedThreadCount() bool`

HasTerminatedThreadCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


