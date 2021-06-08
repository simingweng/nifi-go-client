# ProcessorStatusSnapshotDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The id of the processor. | [optional] 
**GroupId** | Pointer to **string** | The id of the parent process group to which the processor belongs. | [optional] 
**Name** | Pointer to **string** | The name of the prcessor. | [optional] 
**Type** | Pointer to **string** | The type of the processor. | [optional] 
**RunStatus** | Pointer to **string** | The state of the processor. | [optional] 
**ExecutionNode** | Pointer to **string** | Indicates the node where the process will execute. | [optional] 
**BytesRead** | Pointer to **int64** | The number of bytes read by this Processor in the last 5 mintues | [optional] 
**BytesWritten** | Pointer to **int64** | The number of bytes written by this Processor in the last 5 minutes | [optional] 
**Read** | Pointer to **string** | The number of bytes read in the last 5 minutes. | [optional] 
**Written** | Pointer to **string** | The number of bytes written in the last 5 minutes. | [optional] 
**FlowFilesIn** | Pointer to **int32** | The number of FlowFiles that have been accepted in the last 5 minutes | [optional] 
**BytesIn** | Pointer to **int64** | The size of the FlowFiles that have been accepted in the last 5 minutes | [optional] 
**Input** | Pointer to **string** | The count/size of flowfiles that have been accepted in the last 5 minutes. | [optional] 
**FlowFilesOut** | Pointer to **int32** | The number of FlowFiles transferred to a Connection in the last 5 minutes | [optional] 
**BytesOut** | Pointer to **int64** | The size of the FlowFiles transferred to a Connection in the last 5 minutes | [optional] 
**Output** | Pointer to **string** | The count/size of flowfiles that have been processed in the last 5 minutes. | [optional] 
**TaskCount** | Pointer to **int32** | The number of times this Processor has run in the last 5 minutes | [optional] 
**TasksDurationNanos** | Pointer to **int64** | The number of nanoseconds that this Processor has spent running in the last 5 minutes | [optional] 
**Tasks** | Pointer to **string** | The total number of task this connectable has completed over the last 5 minutes. | [optional] 
**TasksDuration** | Pointer to **string** | The total duration of all tasks for this connectable over the last 5 minutes. | [optional] 
**ActiveThreadCount** | Pointer to **int32** | The number of threads currently executing in the processor. | [optional] 
**TerminatedThreadCount** | Pointer to **int32** | The number of threads currently terminated for the processor. | [optional] 

## Methods

### NewProcessorStatusSnapshotDTO

`func NewProcessorStatusSnapshotDTO() *ProcessorStatusSnapshotDTO`

NewProcessorStatusSnapshotDTO instantiates a new ProcessorStatusSnapshotDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessorStatusSnapshotDTOWithDefaults

`func NewProcessorStatusSnapshotDTOWithDefaults() *ProcessorStatusSnapshotDTO`

NewProcessorStatusSnapshotDTOWithDefaults instantiates a new ProcessorStatusSnapshotDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProcessorStatusSnapshotDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProcessorStatusSnapshotDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProcessorStatusSnapshotDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProcessorStatusSnapshotDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetGroupId

`func (o *ProcessorStatusSnapshotDTO) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ProcessorStatusSnapshotDTO) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ProcessorStatusSnapshotDTO) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *ProcessorStatusSnapshotDTO) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetName

`func (o *ProcessorStatusSnapshotDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProcessorStatusSnapshotDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProcessorStatusSnapshotDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProcessorStatusSnapshotDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *ProcessorStatusSnapshotDTO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProcessorStatusSnapshotDTO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProcessorStatusSnapshotDTO) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ProcessorStatusSnapshotDTO) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRunStatus

`func (o *ProcessorStatusSnapshotDTO) GetRunStatus() string`

GetRunStatus returns the RunStatus field if non-nil, zero value otherwise.

### GetRunStatusOk

`func (o *ProcessorStatusSnapshotDTO) GetRunStatusOk() (*string, bool)`

GetRunStatusOk returns a tuple with the RunStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunStatus

`func (o *ProcessorStatusSnapshotDTO) SetRunStatus(v string)`

SetRunStatus sets RunStatus field to given value.

### HasRunStatus

`func (o *ProcessorStatusSnapshotDTO) HasRunStatus() bool`

HasRunStatus returns a boolean if a field has been set.

### GetExecutionNode

`func (o *ProcessorStatusSnapshotDTO) GetExecutionNode() string`

GetExecutionNode returns the ExecutionNode field if non-nil, zero value otherwise.

### GetExecutionNodeOk

`func (o *ProcessorStatusSnapshotDTO) GetExecutionNodeOk() (*string, bool)`

GetExecutionNodeOk returns a tuple with the ExecutionNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionNode

`func (o *ProcessorStatusSnapshotDTO) SetExecutionNode(v string)`

SetExecutionNode sets ExecutionNode field to given value.

### HasExecutionNode

`func (o *ProcessorStatusSnapshotDTO) HasExecutionNode() bool`

HasExecutionNode returns a boolean if a field has been set.

### GetBytesRead

`func (o *ProcessorStatusSnapshotDTO) GetBytesRead() int64`

GetBytesRead returns the BytesRead field if non-nil, zero value otherwise.

### GetBytesReadOk

`func (o *ProcessorStatusSnapshotDTO) GetBytesReadOk() (*int64, bool)`

GetBytesReadOk returns a tuple with the BytesRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesRead

`func (o *ProcessorStatusSnapshotDTO) SetBytesRead(v int64)`

SetBytesRead sets BytesRead field to given value.

### HasBytesRead

`func (o *ProcessorStatusSnapshotDTO) HasBytesRead() bool`

HasBytesRead returns a boolean if a field has been set.

### GetBytesWritten

`func (o *ProcessorStatusSnapshotDTO) GetBytesWritten() int64`

GetBytesWritten returns the BytesWritten field if non-nil, zero value otherwise.

### GetBytesWrittenOk

`func (o *ProcessorStatusSnapshotDTO) GetBytesWrittenOk() (*int64, bool)`

GetBytesWrittenOk returns a tuple with the BytesWritten field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesWritten

`func (o *ProcessorStatusSnapshotDTO) SetBytesWritten(v int64)`

SetBytesWritten sets BytesWritten field to given value.

### HasBytesWritten

`func (o *ProcessorStatusSnapshotDTO) HasBytesWritten() bool`

HasBytesWritten returns a boolean if a field has been set.

### GetRead

`func (o *ProcessorStatusSnapshotDTO) GetRead() string`

GetRead returns the Read field if non-nil, zero value otherwise.

### GetReadOk

`func (o *ProcessorStatusSnapshotDTO) GetReadOk() (*string, bool)`

GetReadOk returns a tuple with the Read field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRead

`func (o *ProcessorStatusSnapshotDTO) SetRead(v string)`

SetRead sets Read field to given value.

### HasRead

`func (o *ProcessorStatusSnapshotDTO) HasRead() bool`

HasRead returns a boolean if a field has been set.

### GetWritten

`func (o *ProcessorStatusSnapshotDTO) GetWritten() string`

GetWritten returns the Written field if non-nil, zero value otherwise.

### GetWrittenOk

`func (o *ProcessorStatusSnapshotDTO) GetWrittenOk() (*string, bool)`

GetWrittenOk returns a tuple with the Written field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWritten

`func (o *ProcessorStatusSnapshotDTO) SetWritten(v string)`

SetWritten sets Written field to given value.

### HasWritten

`func (o *ProcessorStatusSnapshotDTO) HasWritten() bool`

HasWritten returns a boolean if a field has been set.

### GetFlowFilesIn

`func (o *ProcessorStatusSnapshotDTO) GetFlowFilesIn() int32`

GetFlowFilesIn returns the FlowFilesIn field if non-nil, zero value otherwise.

### GetFlowFilesInOk

`func (o *ProcessorStatusSnapshotDTO) GetFlowFilesInOk() (*int32, bool)`

GetFlowFilesInOk returns a tuple with the FlowFilesIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowFilesIn

`func (o *ProcessorStatusSnapshotDTO) SetFlowFilesIn(v int32)`

SetFlowFilesIn sets FlowFilesIn field to given value.

### HasFlowFilesIn

`func (o *ProcessorStatusSnapshotDTO) HasFlowFilesIn() bool`

HasFlowFilesIn returns a boolean if a field has been set.

### GetBytesIn

`func (o *ProcessorStatusSnapshotDTO) GetBytesIn() int64`

GetBytesIn returns the BytesIn field if non-nil, zero value otherwise.

### GetBytesInOk

`func (o *ProcessorStatusSnapshotDTO) GetBytesInOk() (*int64, bool)`

GetBytesInOk returns a tuple with the BytesIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesIn

`func (o *ProcessorStatusSnapshotDTO) SetBytesIn(v int64)`

SetBytesIn sets BytesIn field to given value.

### HasBytesIn

`func (o *ProcessorStatusSnapshotDTO) HasBytesIn() bool`

HasBytesIn returns a boolean if a field has been set.

### GetInput

`func (o *ProcessorStatusSnapshotDTO) GetInput() string`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *ProcessorStatusSnapshotDTO) GetInputOk() (*string, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *ProcessorStatusSnapshotDTO) SetInput(v string)`

SetInput sets Input field to given value.

### HasInput

`func (o *ProcessorStatusSnapshotDTO) HasInput() bool`

HasInput returns a boolean if a field has been set.

### GetFlowFilesOut

`func (o *ProcessorStatusSnapshotDTO) GetFlowFilesOut() int32`

GetFlowFilesOut returns the FlowFilesOut field if non-nil, zero value otherwise.

### GetFlowFilesOutOk

`func (o *ProcessorStatusSnapshotDTO) GetFlowFilesOutOk() (*int32, bool)`

GetFlowFilesOutOk returns a tuple with the FlowFilesOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowFilesOut

`func (o *ProcessorStatusSnapshotDTO) SetFlowFilesOut(v int32)`

SetFlowFilesOut sets FlowFilesOut field to given value.

### HasFlowFilesOut

`func (o *ProcessorStatusSnapshotDTO) HasFlowFilesOut() bool`

HasFlowFilesOut returns a boolean if a field has been set.

### GetBytesOut

`func (o *ProcessorStatusSnapshotDTO) GetBytesOut() int64`

GetBytesOut returns the BytesOut field if non-nil, zero value otherwise.

### GetBytesOutOk

`func (o *ProcessorStatusSnapshotDTO) GetBytesOutOk() (*int64, bool)`

GetBytesOutOk returns a tuple with the BytesOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesOut

`func (o *ProcessorStatusSnapshotDTO) SetBytesOut(v int64)`

SetBytesOut sets BytesOut field to given value.

### HasBytesOut

`func (o *ProcessorStatusSnapshotDTO) HasBytesOut() bool`

HasBytesOut returns a boolean if a field has been set.

### GetOutput

`func (o *ProcessorStatusSnapshotDTO) GetOutput() string`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *ProcessorStatusSnapshotDTO) GetOutputOk() (*string, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *ProcessorStatusSnapshotDTO) SetOutput(v string)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *ProcessorStatusSnapshotDTO) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### GetTaskCount

`func (o *ProcessorStatusSnapshotDTO) GetTaskCount() int32`

GetTaskCount returns the TaskCount field if non-nil, zero value otherwise.

### GetTaskCountOk

`func (o *ProcessorStatusSnapshotDTO) GetTaskCountOk() (*int32, bool)`

GetTaskCountOk returns a tuple with the TaskCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskCount

`func (o *ProcessorStatusSnapshotDTO) SetTaskCount(v int32)`

SetTaskCount sets TaskCount field to given value.

### HasTaskCount

`func (o *ProcessorStatusSnapshotDTO) HasTaskCount() bool`

HasTaskCount returns a boolean if a field has been set.

### GetTasksDurationNanos

`func (o *ProcessorStatusSnapshotDTO) GetTasksDurationNanos() int64`

GetTasksDurationNanos returns the TasksDurationNanos field if non-nil, zero value otherwise.

### GetTasksDurationNanosOk

`func (o *ProcessorStatusSnapshotDTO) GetTasksDurationNanosOk() (*int64, bool)`

GetTasksDurationNanosOk returns a tuple with the TasksDurationNanos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasksDurationNanos

`func (o *ProcessorStatusSnapshotDTO) SetTasksDurationNanos(v int64)`

SetTasksDurationNanos sets TasksDurationNanos field to given value.

### HasTasksDurationNanos

`func (o *ProcessorStatusSnapshotDTO) HasTasksDurationNanos() bool`

HasTasksDurationNanos returns a boolean if a field has been set.

### GetTasks

`func (o *ProcessorStatusSnapshotDTO) GetTasks() string`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *ProcessorStatusSnapshotDTO) GetTasksOk() (*string, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *ProcessorStatusSnapshotDTO) SetTasks(v string)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *ProcessorStatusSnapshotDTO) HasTasks() bool`

HasTasks returns a boolean if a field has been set.

### GetTasksDuration

`func (o *ProcessorStatusSnapshotDTO) GetTasksDuration() string`

GetTasksDuration returns the TasksDuration field if non-nil, zero value otherwise.

### GetTasksDurationOk

`func (o *ProcessorStatusSnapshotDTO) GetTasksDurationOk() (*string, bool)`

GetTasksDurationOk returns a tuple with the TasksDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasksDuration

`func (o *ProcessorStatusSnapshotDTO) SetTasksDuration(v string)`

SetTasksDuration sets TasksDuration field to given value.

### HasTasksDuration

`func (o *ProcessorStatusSnapshotDTO) HasTasksDuration() bool`

HasTasksDuration returns a boolean if a field has been set.

### GetActiveThreadCount

`func (o *ProcessorStatusSnapshotDTO) GetActiveThreadCount() int32`

GetActiveThreadCount returns the ActiveThreadCount field if non-nil, zero value otherwise.

### GetActiveThreadCountOk

`func (o *ProcessorStatusSnapshotDTO) GetActiveThreadCountOk() (*int32, bool)`

GetActiveThreadCountOk returns a tuple with the ActiveThreadCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveThreadCount

`func (o *ProcessorStatusSnapshotDTO) SetActiveThreadCount(v int32)`

SetActiveThreadCount sets ActiveThreadCount field to given value.

### HasActiveThreadCount

`func (o *ProcessorStatusSnapshotDTO) HasActiveThreadCount() bool`

HasActiveThreadCount returns a boolean if a field has been set.

### GetTerminatedThreadCount

`func (o *ProcessorStatusSnapshotDTO) GetTerminatedThreadCount() int32`

GetTerminatedThreadCount returns the TerminatedThreadCount field if non-nil, zero value otherwise.

### GetTerminatedThreadCountOk

`func (o *ProcessorStatusSnapshotDTO) GetTerminatedThreadCountOk() (*int32, bool)`

GetTerminatedThreadCountOk returns a tuple with the TerminatedThreadCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminatedThreadCount

`func (o *ProcessorStatusSnapshotDTO) SetTerminatedThreadCount(v int32)`

SetTerminatedThreadCount sets TerminatedThreadCount field to given value.

### HasTerminatedThreadCount

`func (o *ProcessorStatusSnapshotDTO) HasTerminatedThreadCount() bool`

HasTerminatedThreadCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


