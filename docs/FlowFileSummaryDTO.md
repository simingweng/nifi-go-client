# FlowFileSummaryDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uri** | Pointer to **string** | The URI that can be used to access this FlowFile. | [optional] 
**Uuid** | Pointer to **string** | The FlowFile UUID. | [optional] 
**Filename** | Pointer to **string** | The FlowFile filename. | [optional] 
**Position** | Pointer to **int32** | The FlowFile&#39;s position in the queue. | [optional] 
**Size** | Pointer to **int64** | The FlowFile file size. | [optional] 
**QueuedDuration** | Pointer to **int64** | How long this FlowFile has been enqueued. | [optional] 
**LineageDuration** | Pointer to **int64** | Duration since the FlowFile&#39;s greatest ancestor entered the flow. | [optional] 
**PenaltyExpiresIn** | Pointer to **int64** | How long in milliseconds until the FlowFile penalty expires. | [optional] 
**ClusterNodeId** | Pointer to **string** | The id of the node where this FlowFile resides. | [optional] 
**ClusterNodeAddress** | Pointer to **string** | The label for the node where this FlowFile resides. | [optional] 
**Penalized** | Pointer to **bool** | If the FlowFile is penalized. | [optional] 

## Methods

### NewFlowFileSummaryDTO

`func NewFlowFileSummaryDTO() *FlowFileSummaryDTO`

NewFlowFileSummaryDTO instantiates a new FlowFileSummaryDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlowFileSummaryDTOWithDefaults

`func NewFlowFileSummaryDTOWithDefaults() *FlowFileSummaryDTO`

NewFlowFileSummaryDTOWithDefaults instantiates a new FlowFileSummaryDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUri

`func (o *FlowFileSummaryDTO) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *FlowFileSummaryDTO) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *FlowFileSummaryDTO) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *FlowFileSummaryDTO) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetUuid

`func (o *FlowFileSummaryDTO) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *FlowFileSummaryDTO) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *FlowFileSummaryDTO) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *FlowFileSummaryDTO) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetFilename

`func (o *FlowFileSummaryDTO) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *FlowFileSummaryDTO) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *FlowFileSummaryDTO) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *FlowFileSummaryDTO) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetPosition

`func (o *FlowFileSummaryDTO) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *FlowFileSummaryDTO) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *FlowFileSummaryDTO) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *FlowFileSummaryDTO) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetSize

`func (o *FlowFileSummaryDTO) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *FlowFileSummaryDTO) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *FlowFileSummaryDTO) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *FlowFileSummaryDTO) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetQueuedDuration

`func (o *FlowFileSummaryDTO) GetQueuedDuration() int64`

GetQueuedDuration returns the QueuedDuration field if non-nil, zero value otherwise.

### GetQueuedDurationOk

`func (o *FlowFileSummaryDTO) GetQueuedDurationOk() (*int64, bool)`

GetQueuedDurationOk returns a tuple with the QueuedDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueuedDuration

`func (o *FlowFileSummaryDTO) SetQueuedDuration(v int64)`

SetQueuedDuration sets QueuedDuration field to given value.

### HasQueuedDuration

`func (o *FlowFileSummaryDTO) HasQueuedDuration() bool`

HasQueuedDuration returns a boolean if a field has been set.

### GetLineageDuration

`func (o *FlowFileSummaryDTO) GetLineageDuration() int64`

GetLineageDuration returns the LineageDuration field if non-nil, zero value otherwise.

### GetLineageDurationOk

`func (o *FlowFileSummaryDTO) GetLineageDurationOk() (*int64, bool)`

GetLineageDurationOk returns a tuple with the LineageDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineageDuration

`func (o *FlowFileSummaryDTO) SetLineageDuration(v int64)`

SetLineageDuration sets LineageDuration field to given value.

### HasLineageDuration

`func (o *FlowFileSummaryDTO) HasLineageDuration() bool`

HasLineageDuration returns a boolean if a field has been set.

### GetPenaltyExpiresIn

`func (o *FlowFileSummaryDTO) GetPenaltyExpiresIn() int64`

GetPenaltyExpiresIn returns the PenaltyExpiresIn field if non-nil, zero value otherwise.

### GetPenaltyExpiresInOk

`func (o *FlowFileSummaryDTO) GetPenaltyExpiresInOk() (*int64, bool)`

GetPenaltyExpiresInOk returns a tuple with the PenaltyExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPenaltyExpiresIn

`func (o *FlowFileSummaryDTO) SetPenaltyExpiresIn(v int64)`

SetPenaltyExpiresIn sets PenaltyExpiresIn field to given value.

### HasPenaltyExpiresIn

`func (o *FlowFileSummaryDTO) HasPenaltyExpiresIn() bool`

HasPenaltyExpiresIn returns a boolean if a field has been set.

### GetClusterNodeId

`func (o *FlowFileSummaryDTO) GetClusterNodeId() string`

GetClusterNodeId returns the ClusterNodeId field if non-nil, zero value otherwise.

### GetClusterNodeIdOk

`func (o *FlowFileSummaryDTO) GetClusterNodeIdOk() (*string, bool)`

GetClusterNodeIdOk returns a tuple with the ClusterNodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterNodeId

`func (o *FlowFileSummaryDTO) SetClusterNodeId(v string)`

SetClusterNodeId sets ClusterNodeId field to given value.

### HasClusterNodeId

`func (o *FlowFileSummaryDTO) HasClusterNodeId() bool`

HasClusterNodeId returns a boolean if a field has been set.

### GetClusterNodeAddress

`func (o *FlowFileSummaryDTO) GetClusterNodeAddress() string`

GetClusterNodeAddress returns the ClusterNodeAddress field if non-nil, zero value otherwise.

### GetClusterNodeAddressOk

`func (o *FlowFileSummaryDTO) GetClusterNodeAddressOk() (*string, bool)`

GetClusterNodeAddressOk returns a tuple with the ClusterNodeAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterNodeAddress

`func (o *FlowFileSummaryDTO) SetClusterNodeAddress(v string)`

SetClusterNodeAddress sets ClusterNodeAddress field to given value.

### HasClusterNodeAddress

`func (o *FlowFileSummaryDTO) HasClusterNodeAddress() bool`

HasClusterNodeAddress returns a boolean if a field has been set.

### GetPenalized

`func (o *FlowFileSummaryDTO) GetPenalized() bool`

GetPenalized returns the Penalized field if non-nil, zero value otherwise.

### GetPenalizedOk

`func (o *FlowFileSummaryDTO) GetPenalizedOk() (*bool, bool)`

GetPenalizedOk returns a tuple with the Penalized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPenalized

`func (o *FlowFileSummaryDTO) SetPenalized(v bool)`

SetPenalized sets Penalized field to given value.

### HasPenalized

`func (o *FlowFileSummaryDTO) HasPenalized() bool`

HasPenalized returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


