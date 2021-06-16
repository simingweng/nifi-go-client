# ConnectionStatusSnapshotDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The id of the connection. | [optional] 
**GroupId** | Pointer to **string** | The id of the process group the connection belongs to. | [optional] 
**Name** | Pointer to **string** | The name of the connection. | [optional] 
**SourceId** | Pointer to **string** | The id of the source of the connection. | [optional] 
**SourceName** | Pointer to **string** | The name of the source of the connection. | [optional] 
**DestinationId** | Pointer to **string** | The id of the destination of the connection. | [optional] 
**DestinationName** | Pointer to **string** | The name of the destination of the connection. | [optional] 
**Predictions** | Pointer to [**ConnectionStatusPredictionsSnapshotDTO**](ConnectionStatusPredictionsSnapshotDTO.md) |  | [optional] 
**FlowFilesIn** | Pointer to **int32** | The number of FlowFiles that have come into the connection in the last 5 minutes. | [optional] 
**BytesIn** | Pointer to **int64** | The size of the FlowFiles that have come into the connection in the last 5 minutes. | [optional] 
**Input** | Pointer to **string** | The input count/size for the connection in the last 5 minutes, pretty printed. | [optional] 
**FlowFilesOut** | Pointer to **int32** | The number of FlowFiles that have left the connection in the last 5 minutes. | [optional] 
**BytesOut** | Pointer to **int64** | The number of bytes that have left the connection in the last 5 minutes. | [optional] 
**Output** | Pointer to **string** | The output count/sie for the connection in the last 5 minutes, pretty printed. | [optional] 
**FlowFilesQueued** | Pointer to **int32** | The number of FlowFiles that are currently queued in the connection. | [optional] 
**BytesQueued** | Pointer to **int64** | The size of the FlowFiles that are currently queued in the connection. | [optional] 
**Queued** | Pointer to **string** | The total count and size of queued flowfiles formatted. | [optional] 
**QueuedSize** | Pointer to **string** | The total size of flowfiles that are queued formatted. | [optional] 
**QueuedCount** | Pointer to **string** | The number of flowfiles that are queued, pretty printed. | [optional] 
**PercentUseCount** | Pointer to **int32** | Connection percent use regarding queued flow files count and backpressure threshold if configured. | [optional] 
**PercentUseBytes** | Pointer to **int32** | Connection percent use regarding queued flow files size and backpressure threshold if configured. | [optional] 

## Methods

### NewConnectionStatusSnapshotDTO

`func NewConnectionStatusSnapshotDTO() *ConnectionStatusSnapshotDTO`

NewConnectionStatusSnapshotDTO instantiates a new ConnectionStatusSnapshotDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionStatusSnapshotDTOWithDefaults

`func NewConnectionStatusSnapshotDTOWithDefaults() *ConnectionStatusSnapshotDTO`

NewConnectionStatusSnapshotDTOWithDefaults instantiates a new ConnectionStatusSnapshotDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConnectionStatusSnapshotDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConnectionStatusSnapshotDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConnectionStatusSnapshotDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ConnectionStatusSnapshotDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetGroupId

`func (o *ConnectionStatusSnapshotDTO) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ConnectionStatusSnapshotDTO) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ConnectionStatusSnapshotDTO) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *ConnectionStatusSnapshotDTO) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetName

`func (o *ConnectionStatusSnapshotDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConnectionStatusSnapshotDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConnectionStatusSnapshotDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConnectionStatusSnapshotDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSourceId

`func (o *ConnectionStatusSnapshotDTO) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *ConnectionStatusSnapshotDTO) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *ConnectionStatusSnapshotDTO) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *ConnectionStatusSnapshotDTO) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### GetSourceName

`func (o *ConnectionStatusSnapshotDTO) GetSourceName() string`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *ConnectionStatusSnapshotDTO) GetSourceNameOk() (*string, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *ConnectionStatusSnapshotDTO) SetSourceName(v string)`

SetSourceName sets SourceName field to given value.

### HasSourceName

`func (o *ConnectionStatusSnapshotDTO) HasSourceName() bool`

HasSourceName returns a boolean if a field has been set.

### GetDestinationId

`func (o *ConnectionStatusSnapshotDTO) GetDestinationId() string`

GetDestinationId returns the DestinationId field if non-nil, zero value otherwise.

### GetDestinationIdOk

`func (o *ConnectionStatusSnapshotDTO) GetDestinationIdOk() (*string, bool)`

GetDestinationIdOk returns a tuple with the DestinationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationId

`func (o *ConnectionStatusSnapshotDTO) SetDestinationId(v string)`

SetDestinationId sets DestinationId field to given value.

### HasDestinationId

`func (o *ConnectionStatusSnapshotDTO) HasDestinationId() bool`

HasDestinationId returns a boolean if a field has been set.

### GetDestinationName

`func (o *ConnectionStatusSnapshotDTO) GetDestinationName() string`

GetDestinationName returns the DestinationName field if non-nil, zero value otherwise.

### GetDestinationNameOk

`func (o *ConnectionStatusSnapshotDTO) GetDestinationNameOk() (*string, bool)`

GetDestinationNameOk returns a tuple with the DestinationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationName

`func (o *ConnectionStatusSnapshotDTO) SetDestinationName(v string)`

SetDestinationName sets DestinationName field to given value.

### HasDestinationName

`func (o *ConnectionStatusSnapshotDTO) HasDestinationName() bool`

HasDestinationName returns a boolean if a field has been set.

### GetPredictions

`func (o *ConnectionStatusSnapshotDTO) GetPredictions() ConnectionStatusPredictionsSnapshotDTO`

GetPredictions returns the Predictions field if non-nil, zero value otherwise.

### GetPredictionsOk

`func (o *ConnectionStatusSnapshotDTO) GetPredictionsOk() (*ConnectionStatusPredictionsSnapshotDTO, bool)`

GetPredictionsOk returns a tuple with the Predictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredictions

`func (o *ConnectionStatusSnapshotDTO) SetPredictions(v ConnectionStatusPredictionsSnapshotDTO)`

SetPredictions sets Predictions field to given value.

### HasPredictions

`func (o *ConnectionStatusSnapshotDTO) HasPredictions() bool`

HasPredictions returns a boolean if a field has been set.

### GetFlowFilesIn

`func (o *ConnectionStatusSnapshotDTO) GetFlowFilesIn() int32`

GetFlowFilesIn returns the FlowFilesIn field if non-nil, zero value otherwise.

### GetFlowFilesInOk

`func (o *ConnectionStatusSnapshotDTO) GetFlowFilesInOk() (*int32, bool)`

GetFlowFilesInOk returns a tuple with the FlowFilesIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowFilesIn

`func (o *ConnectionStatusSnapshotDTO) SetFlowFilesIn(v int32)`

SetFlowFilesIn sets FlowFilesIn field to given value.

### HasFlowFilesIn

`func (o *ConnectionStatusSnapshotDTO) HasFlowFilesIn() bool`

HasFlowFilesIn returns a boolean if a field has been set.

### GetBytesIn

`func (o *ConnectionStatusSnapshotDTO) GetBytesIn() int64`

GetBytesIn returns the BytesIn field if non-nil, zero value otherwise.

### GetBytesInOk

`func (o *ConnectionStatusSnapshotDTO) GetBytesInOk() (*int64, bool)`

GetBytesInOk returns a tuple with the BytesIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesIn

`func (o *ConnectionStatusSnapshotDTO) SetBytesIn(v int64)`

SetBytesIn sets BytesIn field to given value.

### HasBytesIn

`func (o *ConnectionStatusSnapshotDTO) HasBytesIn() bool`

HasBytesIn returns a boolean if a field has been set.

### GetInput

`func (o *ConnectionStatusSnapshotDTO) GetInput() string`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *ConnectionStatusSnapshotDTO) GetInputOk() (*string, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *ConnectionStatusSnapshotDTO) SetInput(v string)`

SetInput sets Input field to given value.

### HasInput

`func (o *ConnectionStatusSnapshotDTO) HasInput() bool`

HasInput returns a boolean if a field has been set.

### GetFlowFilesOut

`func (o *ConnectionStatusSnapshotDTO) GetFlowFilesOut() int32`

GetFlowFilesOut returns the FlowFilesOut field if non-nil, zero value otherwise.

### GetFlowFilesOutOk

`func (o *ConnectionStatusSnapshotDTO) GetFlowFilesOutOk() (*int32, bool)`

GetFlowFilesOutOk returns a tuple with the FlowFilesOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowFilesOut

`func (o *ConnectionStatusSnapshotDTO) SetFlowFilesOut(v int32)`

SetFlowFilesOut sets FlowFilesOut field to given value.

### HasFlowFilesOut

`func (o *ConnectionStatusSnapshotDTO) HasFlowFilesOut() bool`

HasFlowFilesOut returns a boolean if a field has been set.

### GetBytesOut

`func (o *ConnectionStatusSnapshotDTO) GetBytesOut() int64`

GetBytesOut returns the BytesOut field if non-nil, zero value otherwise.

### GetBytesOutOk

`func (o *ConnectionStatusSnapshotDTO) GetBytesOutOk() (*int64, bool)`

GetBytesOutOk returns a tuple with the BytesOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesOut

`func (o *ConnectionStatusSnapshotDTO) SetBytesOut(v int64)`

SetBytesOut sets BytesOut field to given value.

### HasBytesOut

`func (o *ConnectionStatusSnapshotDTO) HasBytesOut() bool`

HasBytesOut returns a boolean if a field has been set.

### GetOutput

`func (o *ConnectionStatusSnapshotDTO) GetOutput() string`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *ConnectionStatusSnapshotDTO) GetOutputOk() (*string, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *ConnectionStatusSnapshotDTO) SetOutput(v string)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *ConnectionStatusSnapshotDTO) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### GetFlowFilesQueued

`func (o *ConnectionStatusSnapshotDTO) GetFlowFilesQueued() int32`

GetFlowFilesQueued returns the FlowFilesQueued field if non-nil, zero value otherwise.

### GetFlowFilesQueuedOk

`func (o *ConnectionStatusSnapshotDTO) GetFlowFilesQueuedOk() (*int32, bool)`

GetFlowFilesQueuedOk returns a tuple with the FlowFilesQueued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowFilesQueued

`func (o *ConnectionStatusSnapshotDTO) SetFlowFilesQueued(v int32)`

SetFlowFilesQueued sets FlowFilesQueued field to given value.

### HasFlowFilesQueued

`func (o *ConnectionStatusSnapshotDTO) HasFlowFilesQueued() bool`

HasFlowFilesQueued returns a boolean if a field has been set.

### GetBytesQueued

`func (o *ConnectionStatusSnapshotDTO) GetBytesQueued() int64`

GetBytesQueued returns the BytesQueued field if non-nil, zero value otherwise.

### GetBytesQueuedOk

`func (o *ConnectionStatusSnapshotDTO) GetBytesQueuedOk() (*int64, bool)`

GetBytesQueuedOk returns a tuple with the BytesQueued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesQueued

`func (o *ConnectionStatusSnapshotDTO) SetBytesQueued(v int64)`

SetBytesQueued sets BytesQueued field to given value.

### HasBytesQueued

`func (o *ConnectionStatusSnapshotDTO) HasBytesQueued() bool`

HasBytesQueued returns a boolean if a field has been set.

### GetQueued

`func (o *ConnectionStatusSnapshotDTO) GetQueued() string`

GetQueued returns the Queued field if non-nil, zero value otherwise.

### GetQueuedOk

`func (o *ConnectionStatusSnapshotDTO) GetQueuedOk() (*string, bool)`

GetQueuedOk returns a tuple with the Queued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueued

`func (o *ConnectionStatusSnapshotDTO) SetQueued(v string)`

SetQueued sets Queued field to given value.

### HasQueued

`func (o *ConnectionStatusSnapshotDTO) HasQueued() bool`

HasQueued returns a boolean if a field has been set.

### GetQueuedSize

`func (o *ConnectionStatusSnapshotDTO) GetQueuedSize() string`

GetQueuedSize returns the QueuedSize field if non-nil, zero value otherwise.

### GetQueuedSizeOk

`func (o *ConnectionStatusSnapshotDTO) GetQueuedSizeOk() (*string, bool)`

GetQueuedSizeOk returns a tuple with the QueuedSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueuedSize

`func (o *ConnectionStatusSnapshotDTO) SetQueuedSize(v string)`

SetQueuedSize sets QueuedSize field to given value.

### HasQueuedSize

`func (o *ConnectionStatusSnapshotDTO) HasQueuedSize() bool`

HasQueuedSize returns a boolean if a field has been set.

### GetQueuedCount

`func (o *ConnectionStatusSnapshotDTO) GetQueuedCount() string`

GetQueuedCount returns the QueuedCount field if non-nil, zero value otherwise.

### GetQueuedCountOk

`func (o *ConnectionStatusSnapshotDTO) GetQueuedCountOk() (*string, bool)`

GetQueuedCountOk returns a tuple with the QueuedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueuedCount

`func (o *ConnectionStatusSnapshotDTO) SetQueuedCount(v string)`

SetQueuedCount sets QueuedCount field to given value.

### HasQueuedCount

`func (o *ConnectionStatusSnapshotDTO) HasQueuedCount() bool`

HasQueuedCount returns a boolean if a field has been set.

### GetPercentUseCount

`func (o *ConnectionStatusSnapshotDTO) GetPercentUseCount() int32`

GetPercentUseCount returns the PercentUseCount field if non-nil, zero value otherwise.

### GetPercentUseCountOk

`func (o *ConnectionStatusSnapshotDTO) GetPercentUseCountOk() (*int32, bool)`

GetPercentUseCountOk returns a tuple with the PercentUseCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentUseCount

`func (o *ConnectionStatusSnapshotDTO) SetPercentUseCount(v int32)`

SetPercentUseCount sets PercentUseCount field to given value.

### HasPercentUseCount

`func (o *ConnectionStatusSnapshotDTO) HasPercentUseCount() bool`

HasPercentUseCount returns a boolean if a field has been set.

### GetPercentUseBytes

`func (o *ConnectionStatusSnapshotDTO) GetPercentUseBytes() int32`

GetPercentUseBytes returns the PercentUseBytes field if non-nil, zero value otherwise.

### GetPercentUseBytesOk

`func (o *ConnectionStatusSnapshotDTO) GetPercentUseBytesOk() (*int32, bool)`

GetPercentUseBytesOk returns a tuple with the PercentUseBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentUseBytes

`func (o *ConnectionStatusSnapshotDTO) SetPercentUseBytes(v int32)`

SetPercentUseBytes sets PercentUseBytes field to given value.

### HasPercentUseBytes

`func (o *ConnectionStatusSnapshotDTO) HasPercentUseBytes() bool`

HasPercentUseBytes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


