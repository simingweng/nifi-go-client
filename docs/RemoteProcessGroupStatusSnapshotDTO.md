# RemoteProcessGroupStatusSnapshotDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The id of the remote process group. | [optional] 
**GroupId** | Pointer to **string** | The id of the parent process group the remote process group resides in. | [optional] 
**Name** | Pointer to **string** | The name of the remote process group. | [optional] 
**TargetUri** | Pointer to **string** | The URI of the target system. | [optional] 
**TransmissionStatus** | Pointer to **string** | The transmission status of the remote process group. | [optional] 
**ActiveThreadCount** | Pointer to **int32** | The number of active threads for the remote process group. | [optional] 
**FlowFilesSent** | Pointer to **int32** | The number of FlowFiles sent to the remote process group in the last 5 minutes. | [optional] 
**BytesSent** | Pointer to **int64** | The size of the FlowFiles sent to the remote process group in the last 5 minutes. | [optional] 
**Sent** | Pointer to **string** | The count/size of the flowfiles sent to the remote process group in the last 5 minutes. | [optional] 
**FlowFilesReceived** | Pointer to **int32** | The number of FlowFiles received from the remote process group in the last 5 minutes. | [optional] 
**BytesReceived** | Pointer to **int64** | The size of the FlowFiles received from the remote process group in the last 5 minutes. | [optional] 
**Received** | Pointer to **string** | The count/size of the flowfiles received from the remote process group in the last 5 minutes. | [optional] 

## Methods

### NewRemoteProcessGroupStatusSnapshotDTO

`func NewRemoteProcessGroupStatusSnapshotDTO() *RemoteProcessGroupStatusSnapshotDTO`

NewRemoteProcessGroupStatusSnapshotDTO instantiates a new RemoteProcessGroupStatusSnapshotDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoteProcessGroupStatusSnapshotDTOWithDefaults

`func NewRemoteProcessGroupStatusSnapshotDTOWithDefaults() *RemoteProcessGroupStatusSnapshotDTO`

NewRemoteProcessGroupStatusSnapshotDTOWithDefaults instantiates a new RemoteProcessGroupStatusSnapshotDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RemoteProcessGroupStatusSnapshotDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RemoteProcessGroupStatusSnapshotDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RemoteProcessGroupStatusSnapshotDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RemoteProcessGroupStatusSnapshotDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetGroupId

`func (o *RemoteProcessGroupStatusSnapshotDTO) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *RemoteProcessGroupStatusSnapshotDTO) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *RemoteProcessGroupStatusSnapshotDTO) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *RemoteProcessGroupStatusSnapshotDTO) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetName

`func (o *RemoteProcessGroupStatusSnapshotDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RemoteProcessGroupStatusSnapshotDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RemoteProcessGroupStatusSnapshotDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RemoteProcessGroupStatusSnapshotDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTargetUri

`func (o *RemoteProcessGroupStatusSnapshotDTO) GetTargetUri() string`

GetTargetUri returns the TargetUri field if non-nil, zero value otherwise.

### GetTargetUriOk

`func (o *RemoteProcessGroupStatusSnapshotDTO) GetTargetUriOk() (*string, bool)`

GetTargetUriOk returns a tuple with the TargetUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUri

`func (o *RemoteProcessGroupStatusSnapshotDTO) SetTargetUri(v string)`

SetTargetUri sets TargetUri field to given value.

### HasTargetUri

`func (o *RemoteProcessGroupStatusSnapshotDTO) HasTargetUri() bool`

HasTargetUri returns a boolean if a field has been set.

### GetTransmissionStatus

`func (o *RemoteProcessGroupStatusSnapshotDTO) GetTransmissionStatus() string`

GetTransmissionStatus returns the TransmissionStatus field if non-nil, zero value otherwise.

### GetTransmissionStatusOk

`func (o *RemoteProcessGroupStatusSnapshotDTO) GetTransmissionStatusOk() (*string, bool)`

GetTransmissionStatusOk returns a tuple with the TransmissionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransmissionStatus

`func (o *RemoteProcessGroupStatusSnapshotDTO) SetTransmissionStatus(v string)`

SetTransmissionStatus sets TransmissionStatus field to given value.

### HasTransmissionStatus

`func (o *RemoteProcessGroupStatusSnapshotDTO) HasTransmissionStatus() bool`

HasTransmissionStatus returns a boolean if a field has been set.

### GetActiveThreadCount

`func (o *RemoteProcessGroupStatusSnapshotDTO) GetActiveThreadCount() int32`

GetActiveThreadCount returns the ActiveThreadCount field if non-nil, zero value otherwise.

### GetActiveThreadCountOk

`func (o *RemoteProcessGroupStatusSnapshotDTO) GetActiveThreadCountOk() (*int32, bool)`

GetActiveThreadCountOk returns a tuple with the ActiveThreadCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveThreadCount

`func (o *RemoteProcessGroupStatusSnapshotDTO) SetActiveThreadCount(v int32)`

SetActiveThreadCount sets ActiveThreadCount field to given value.

### HasActiveThreadCount

`func (o *RemoteProcessGroupStatusSnapshotDTO) HasActiveThreadCount() bool`

HasActiveThreadCount returns a boolean if a field has been set.

### GetFlowFilesSent

`func (o *RemoteProcessGroupStatusSnapshotDTO) GetFlowFilesSent() int32`

GetFlowFilesSent returns the FlowFilesSent field if non-nil, zero value otherwise.

### GetFlowFilesSentOk

`func (o *RemoteProcessGroupStatusSnapshotDTO) GetFlowFilesSentOk() (*int32, bool)`

GetFlowFilesSentOk returns a tuple with the FlowFilesSent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowFilesSent

`func (o *RemoteProcessGroupStatusSnapshotDTO) SetFlowFilesSent(v int32)`

SetFlowFilesSent sets FlowFilesSent field to given value.

### HasFlowFilesSent

`func (o *RemoteProcessGroupStatusSnapshotDTO) HasFlowFilesSent() bool`

HasFlowFilesSent returns a boolean if a field has been set.

### GetBytesSent

`func (o *RemoteProcessGroupStatusSnapshotDTO) GetBytesSent() int64`

GetBytesSent returns the BytesSent field if non-nil, zero value otherwise.

### GetBytesSentOk

`func (o *RemoteProcessGroupStatusSnapshotDTO) GetBytesSentOk() (*int64, bool)`

GetBytesSentOk returns a tuple with the BytesSent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesSent

`func (o *RemoteProcessGroupStatusSnapshotDTO) SetBytesSent(v int64)`

SetBytesSent sets BytesSent field to given value.

### HasBytesSent

`func (o *RemoteProcessGroupStatusSnapshotDTO) HasBytesSent() bool`

HasBytesSent returns a boolean if a field has been set.

### GetSent

`func (o *RemoteProcessGroupStatusSnapshotDTO) GetSent() string`

GetSent returns the Sent field if non-nil, zero value otherwise.

### GetSentOk

`func (o *RemoteProcessGroupStatusSnapshotDTO) GetSentOk() (*string, bool)`

GetSentOk returns a tuple with the Sent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSent

`func (o *RemoteProcessGroupStatusSnapshotDTO) SetSent(v string)`

SetSent sets Sent field to given value.

### HasSent

`func (o *RemoteProcessGroupStatusSnapshotDTO) HasSent() bool`

HasSent returns a boolean if a field has been set.

### GetFlowFilesReceived

`func (o *RemoteProcessGroupStatusSnapshotDTO) GetFlowFilesReceived() int32`

GetFlowFilesReceived returns the FlowFilesReceived field if non-nil, zero value otherwise.

### GetFlowFilesReceivedOk

`func (o *RemoteProcessGroupStatusSnapshotDTO) GetFlowFilesReceivedOk() (*int32, bool)`

GetFlowFilesReceivedOk returns a tuple with the FlowFilesReceived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowFilesReceived

`func (o *RemoteProcessGroupStatusSnapshotDTO) SetFlowFilesReceived(v int32)`

SetFlowFilesReceived sets FlowFilesReceived field to given value.

### HasFlowFilesReceived

`func (o *RemoteProcessGroupStatusSnapshotDTO) HasFlowFilesReceived() bool`

HasFlowFilesReceived returns a boolean if a field has been set.

### GetBytesReceived

`func (o *RemoteProcessGroupStatusSnapshotDTO) GetBytesReceived() int64`

GetBytesReceived returns the BytesReceived field if non-nil, zero value otherwise.

### GetBytesReceivedOk

`func (o *RemoteProcessGroupStatusSnapshotDTO) GetBytesReceivedOk() (*int64, bool)`

GetBytesReceivedOk returns a tuple with the BytesReceived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesReceived

`func (o *RemoteProcessGroupStatusSnapshotDTO) SetBytesReceived(v int64)`

SetBytesReceived sets BytesReceived field to given value.

### HasBytesReceived

`func (o *RemoteProcessGroupStatusSnapshotDTO) HasBytesReceived() bool`

HasBytesReceived returns a boolean if a field has been set.

### GetReceived

`func (o *RemoteProcessGroupStatusSnapshotDTO) GetReceived() string`

GetReceived returns the Received field if non-nil, zero value otherwise.

### GetReceivedOk

`func (o *RemoteProcessGroupStatusSnapshotDTO) GetReceivedOk() (*string, bool)`

GetReceivedOk returns a tuple with the Received field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceived

`func (o *RemoteProcessGroupStatusSnapshotDTO) SetReceived(v string)`

SetReceived sets Received field to given value.

### HasReceived

`func (o *RemoteProcessGroupStatusSnapshotDTO) HasReceived() bool`

HasReceived returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


