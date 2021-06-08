# PortStatusSnapshotDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The id of the port. | [optional] 
**GroupId** | Pointer to **string** | The id of the parent process group of the port. | [optional] 
**Name** | Pointer to **string** | The name of the port. | [optional] 
**ActiveThreadCount** | Pointer to **int32** | The active thread count for the port. | [optional] 
**FlowFilesIn** | Pointer to **int32** | The number of FlowFiles that have been accepted in the last 5 minutes. | [optional] 
**BytesIn** | Pointer to **int64** | The size of hte FlowFiles that have been accepted in the last 5 minutes. | [optional] 
**Input** | Pointer to **string** | The count/size of flowfiles that have been accepted in the last 5 minutes. | [optional] 
**FlowFilesOut** | Pointer to **int32** | The number of FlowFiles that have been processed in the last 5 minutes. | [optional] 
**BytesOut** | Pointer to **int64** | The number of bytes that have been processed in the last 5 minutes. | [optional] 
**Output** | Pointer to **string** | The count/size of flowfiles that have been processed in the last 5 minutes. | [optional] 
**Transmitting** | Pointer to **bool** | Whether the port has incoming or outgoing connections to a remote NiFi. | [optional] 
**RunStatus** | Pointer to **string** | The run status of the port. | [optional] 

## Methods

### NewPortStatusSnapshotDTO

`func NewPortStatusSnapshotDTO() *PortStatusSnapshotDTO`

NewPortStatusSnapshotDTO instantiates a new PortStatusSnapshotDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortStatusSnapshotDTOWithDefaults

`func NewPortStatusSnapshotDTOWithDefaults() *PortStatusSnapshotDTO`

NewPortStatusSnapshotDTOWithDefaults instantiates a new PortStatusSnapshotDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PortStatusSnapshotDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PortStatusSnapshotDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PortStatusSnapshotDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PortStatusSnapshotDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetGroupId

`func (o *PortStatusSnapshotDTO) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *PortStatusSnapshotDTO) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *PortStatusSnapshotDTO) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *PortStatusSnapshotDTO) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetName

`func (o *PortStatusSnapshotDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PortStatusSnapshotDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PortStatusSnapshotDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PortStatusSnapshotDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetActiveThreadCount

`func (o *PortStatusSnapshotDTO) GetActiveThreadCount() int32`

GetActiveThreadCount returns the ActiveThreadCount field if non-nil, zero value otherwise.

### GetActiveThreadCountOk

`func (o *PortStatusSnapshotDTO) GetActiveThreadCountOk() (*int32, bool)`

GetActiveThreadCountOk returns a tuple with the ActiveThreadCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveThreadCount

`func (o *PortStatusSnapshotDTO) SetActiveThreadCount(v int32)`

SetActiveThreadCount sets ActiveThreadCount field to given value.

### HasActiveThreadCount

`func (o *PortStatusSnapshotDTO) HasActiveThreadCount() bool`

HasActiveThreadCount returns a boolean if a field has been set.

### GetFlowFilesIn

`func (o *PortStatusSnapshotDTO) GetFlowFilesIn() int32`

GetFlowFilesIn returns the FlowFilesIn field if non-nil, zero value otherwise.

### GetFlowFilesInOk

`func (o *PortStatusSnapshotDTO) GetFlowFilesInOk() (*int32, bool)`

GetFlowFilesInOk returns a tuple with the FlowFilesIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowFilesIn

`func (o *PortStatusSnapshotDTO) SetFlowFilesIn(v int32)`

SetFlowFilesIn sets FlowFilesIn field to given value.

### HasFlowFilesIn

`func (o *PortStatusSnapshotDTO) HasFlowFilesIn() bool`

HasFlowFilesIn returns a boolean if a field has been set.

### GetBytesIn

`func (o *PortStatusSnapshotDTO) GetBytesIn() int64`

GetBytesIn returns the BytesIn field if non-nil, zero value otherwise.

### GetBytesInOk

`func (o *PortStatusSnapshotDTO) GetBytesInOk() (*int64, bool)`

GetBytesInOk returns a tuple with the BytesIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesIn

`func (o *PortStatusSnapshotDTO) SetBytesIn(v int64)`

SetBytesIn sets BytesIn field to given value.

### HasBytesIn

`func (o *PortStatusSnapshotDTO) HasBytesIn() bool`

HasBytesIn returns a boolean if a field has been set.

### GetInput

`func (o *PortStatusSnapshotDTO) GetInput() string`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *PortStatusSnapshotDTO) GetInputOk() (*string, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *PortStatusSnapshotDTO) SetInput(v string)`

SetInput sets Input field to given value.

### HasInput

`func (o *PortStatusSnapshotDTO) HasInput() bool`

HasInput returns a boolean if a field has been set.

### GetFlowFilesOut

`func (o *PortStatusSnapshotDTO) GetFlowFilesOut() int32`

GetFlowFilesOut returns the FlowFilesOut field if non-nil, zero value otherwise.

### GetFlowFilesOutOk

`func (o *PortStatusSnapshotDTO) GetFlowFilesOutOk() (*int32, bool)`

GetFlowFilesOutOk returns a tuple with the FlowFilesOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowFilesOut

`func (o *PortStatusSnapshotDTO) SetFlowFilesOut(v int32)`

SetFlowFilesOut sets FlowFilesOut field to given value.

### HasFlowFilesOut

`func (o *PortStatusSnapshotDTO) HasFlowFilesOut() bool`

HasFlowFilesOut returns a boolean if a field has been set.

### GetBytesOut

`func (o *PortStatusSnapshotDTO) GetBytesOut() int64`

GetBytesOut returns the BytesOut field if non-nil, zero value otherwise.

### GetBytesOutOk

`func (o *PortStatusSnapshotDTO) GetBytesOutOk() (*int64, bool)`

GetBytesOutOk returns a tuple with the BytesOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesOut

`func (o *PortStatusSnapshotDTO) SetBytesOut(v int64)`

SetBytesOut sets BytesOut field to given value.

### HasBytesOut

`func (o *PortStatusSnapshotDTO) HasBytesOut() bool`

HasBytesOut returns a boolean if a field has been set.

### GetOutput

`func (o *PortStatusSnapshotDTO) GetOutput() string`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *PortStatusSnapshotDTO) GetOutputOk() (*string, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *PortStatusSnapshotDTO) SetOutput(v string)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *PortStatusSnapshotDTO) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### GetTransmitting

`func (o *PortStatusSnapshotDTO) GetTransmitting() bool`

GetTransmitting returns the Transmitting field if non-nil, zero value otherwise.

### GetTransmittingOk

`func (o *PortStatusSnapshotDTO) GetTransmittingOk() (*bool, bool)`

GetTransmittingOk returns a tuple with the Transmitting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransmitting

`func (o *PortStatusSnapshotDTO) SetTransmitting(v bool)`

SetTransmitting sets Transmitting field to given value.

### HasTransmitting

`func (o *PortStatusSnapshotDTO) HasTransmitting() bool`

HasTransmitting returns a boolean if a field has been set.

### GetRunStatus

`func (o *PortStatusSnapshotDTO) GetRunStatus() string`

GetRunStatus returns the RunStatus field if non-nil, zero value otherwise.

### GetRunStatusOk

`func (o *PortStatusSnapshotDTO) GetRunStatusOk() (*string, bool)`

GetRunStatusOk returns a tuple with the RunStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunStatus

`func (o *PortStatusSnapshotDTO) SetRunStatus(v string)`

SetRunStatus sets RunStatus field to given value.

### HasRunStatus

`func (o *PortStatusSnapshotDTO) HasRunStatus() bool`

HasRunStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


