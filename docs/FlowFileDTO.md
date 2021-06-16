# FlowFileDTO

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
**Attributes** | Pointer to **map[string]string** | The FlowFile attributes. | [optional] 
**ContentClaimSection** | Pointer to **string** | The section in which the content claim lives. | [optional] 
**ContentClaimContainer** | Pointer to **string** | The container in which the content claim lives. | [optional] 
**ContentClaimIdentifier** | Pointer to **string** | The identifier of the content claim. | [optional] 
**ContentClaimOffset** | Pointer to **int64** | The offset into the content claim where the flowfile&#39;s content begins. | [optional] 
**ContentClaimFileSize** | Pointer to **string** | The file size of the content claim formatted. | [optional] 
**ContentClaimFileSizeBytes** | Pointer to **int64** | The file size of the content claim in bytes. | [optional] 
**Penalized** | Pointer to **bool** | If the FlowFile is penalized. | [optional] 

## Methods

### NewFlowFileDTO

`func NewFlowFileDTO() *FlowFileDTO`

NewFlowFileDTO instantiates a new FlowFileDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlowFileDTOWithDefaults

`func NewFlowFileDTOWithDefaults() *FlowFileDTO`

NewFlowFileDTOWithDefaults instantiates a new FlowFileDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUri

`func (o *FlowFileDTO) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *FlowFileDTO) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *FlowFileDTO) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *FlowFileDTO) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetUuid

`func (o *FlowFileDTO) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *FlowFileDTO) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *FlowFileDTO) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *FlowFileDTO) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetFilename

`func (o *FlowFileDTO) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *FlowFileDTO) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *FlowFileDTO) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *FlowFileDTO) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetPosition

`func (o *FlowFileDTO) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *FlowFileDTO) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *FlowFileDTO) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *FlowFileDTO) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetSize

`func (o *FlowFileDTO) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *FlowFileDTO) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *FlowFileDTO) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *FlowFileDTO) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetQueuedDuration

`func (o *FlowFileDTO) GetQueuedDuration() int64`

GetQueuedDuration returns the QueuedDuration field if non-nil, zero value otherwise.

### GetQueuedDurationOk

`func (o *FlowFileDTO) GetQueuedDurationOk() (*int64, bool)`

GetQueuedDurationOk returns a tuple with the QueuedDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueuedDuration

`func (o *FlowFileDTO) SetQueuedDuration(v int64)`

SetQueuedDuration sets QueuedDuration field to given value.

### HasQueuedDuration

`func (o *FlowFileDTO) HasQueuedDuration() bool`

HasQueuedDuration returns a boolean if a field has been set.

### GetLineageDuration

`func (o *FlowFileDTO) GetLineageDuration() int64`

GetLineageDuration returns the LineageDuration field if non-nil, zero value otherwise.

### GetLineageDurationOk

`func (o *FlowFileDTO) GetLineageDurationOk() (*int64, bool)`

GetLineageDurationOk returns a tuple with the LineageDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineageDuration

`func (o *FlowFileDTO) SetLineageDuration(v int64)`

SetLineageDuration sets LineageDuration field to given value.

### HasLineageDuration

`func (o *FlowFileDTO) HasLineageDuration() bool`

HasLineageDuration returns a boolean if a field has been set.

### GetPenaltyExpiresIn

`func (o *FlowFileDTO) GetPenaltyExpiresIn() int64`

GetPenaltyExpiresIn returns the PenaltyExpiresIn field if non-nil, zero value otherwise.

### GetPenaltyExpiresInOk

`func (o *FlowFileDTO) GetPenaltyExpiresInOk() (*int64, bool)`

GetPenaltyExpiresInOk returns a tuple with the PenaltyExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPenaltyExpiresIn

`func (o *FlowFileDTO) SetPenaltyExpiresIn(v int64)`

SetPenaltyExpiresIn sets PenaltyExpiresIn field to given value.

### HasPenaltyExpiresIn

`func (o *FlowFileDTO) HasPenaltyExpiresIn() bool`

HasPenaltyExpiresIn returns a boolean if a field has been set.

### GetClusterNodeId

`func (o *FlowFileDTO) GetClusterNodeId() string`

GetClusterNodeId returns the ClusterNodeId field if non-nil, zero value otherwise.

### GetClusterNodeIdOk

`func (o *FlowFileDTO) GetClusterNodeIdOk() (*string, bool)`

GetClusterNodeIdOk returns a tuple with the ClusterNodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterNodeId

`func (o *FlowFileDTO) SetClusterNodeId(v string)`

SetClusterNodeId sets ClusterNodeId field to given value.

### HasClusterNodeId

`func (o *FlowFileDTO) HasClusterNodeId() bool`

HasClusterNodeId returns a boolean if a field has been set.

### GetClusterNodeAddress

`func (o *FlowFileDTO) GetClusterNodeAddress() string`

GetClusterNodeAddress returns the ClusterNodeAddress field if non-nil, zero value otherwise.

### GetClusterNodeAddressOk

`func (o *FlowFileDTO) GetClusterNodeAddressOk() (*string, bool)`

GetClusterNodeAddressOk returns a tuple with the ClusterNodeAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterNodeAddress

`func (o *FlowFileDTO) SetClusterNodeAddress(v string)`

SetClusterNodeAddress sets ClusterNodeAddress field to given value.

### HasClusterNodeAddress

`func (o *FlowFileDTO) HasClusterNodeAddress() bool`

HasClusterNodeAddress returns a boolean if a field has been set.

### GetAttributes

`func (o *FlowFileDTO) GetAttributes() map[string]string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *FlowFileDTO) GetAttributesOk() (*map[string]string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *FlowFileDTO) SetAttributes(v map[string]string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *FlowFileDTO) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetContentClaimSection

`func (o *FlowFileDTO) GetContentClaimSection() string`

GetContentClaimSection returns the ContentClaimSection field if non-nil, zero value otherwise.

### GetContentClaimSectionOk

`func (o *FlowFileDTO) GetContentClaimSectionOk() (*string, bool)`

GetContentClaimSectionOk returns a tuple with the ContentClaimSection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentClaimSection

`func (o *FlowFileDTO) SetContentClaimSection(v string)`

SetContentClaimSection sets ContentClaimSection field to given value.

### HasContentClaimSection

`func (o *FlowFileDTO) HasContentClaimSection() bool`

HasContentClaimSection returns a boolean if a field has been set.

### GetContentClaimContainer

`func (o *FlowFileDTO) GetContentClaimContainer() string`

GetContentClaimContainer returns the ContentClaimContainer field if non-nil, zero value otherwise.

### GetContentClaimContainerOk

`func (o *FlowFileDTO) GetContentClaimContainerOk() (*string, bool)`

GetContentClaimContainerOk returns a tuple with the ContentClaimContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentClaimContainer

`func (o *FlowFileDTO) SetContentClaimContainer(v string)`

SetContentClaimContainer sets ContentClaimContainer field to given value.

### HasContentClaimContainer

`func (o *FlowFileDTO) HasContentClaimContainer() bool`

HasContentClaimContainer returns a boolean if a field has been set.

### GetContentClaimIdentifier

`func (o *FlowFileDTO) GetContentClaimIdentifier() string`

GetContentClaimIdentifier returns the ContentClaimIdentifier field if non-nil, zero value otherwise.

### GetContentClaimIdentifierOk

`func (o *FlowFileDTO) GetContentClaimIdentifierOk() (*string, bool)`

GetContentClaimIdentifierOk returns a tuple with the ContentClaimIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentClaimIdentifier

`func (o *FlowFileDTO) SetContentClaimIdentifier(v string)`

SetContentClaimIdentifier sets ContentClaimIdentifier field to given value.

### HasContentClaimIdentifier

`func (o *FlowFileDTO) HasContentClaimIdentifier() bool`

HasContentClaimIdentifier returns a boolean if a field has been set.

### GetContentClaimOffset

`func (o *FlowFileDTO) GetContentClaimOffset() int64`

GetContentClaimOffset returns the ContentClaimOffset field if non-nil, zero value otherwise.

### GetContentClaimOffsetOk

`func (o *FlowFileDTO) GetContentClaimOffsetOk() (*int64, bool)`

GetContentClaimOffsetOk returns a tuple with the ContentClaimOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentClaimOffset

`func (o *FlowFileDTO) SetContentClaimOffset(v int64)`

SetContentClaimOffset sets ContentClaimOffset field to given value.

### HasContentClaimOffset

`func (o *FlowFileDTO) HasContentClaimOffset() bool`

HasContentClaimOffset returns a boolean if a field has been set.

### GetContentClaimFileSize

`func (o *FlowFileDTO) GetContentClaimFileSize() string`

GetContentClaimFileSize returns the ContentClaimFileSize field if non-nil, zero value otherwise.

### GetContentClaimFileSizeOk

`func (o *FlowFileDTO) GetContentClaimFileSizeOk() (*string, bool)`

GetContentClaimFileSizeOk returns a tuple with the ContentClaimFileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentClaimFileSize

`func (o *FlowFileDTO) SetContentClaimFileSize(v string)`

SetContentClaimFileSize sets ContentClaimFileSize field to given value.

### HasContentClaimFileSize

`func (o *FlowFileDTO) HasContentClaimFileSize() bool`

HasContentClaimFileSize returns a boolean if a field has been set.

### GetContentClaimFileSizeBytes

`func (o *FlowFileDTO) GetContentClaimFileSizeBytes() int64`

GetContentClaimFileSizeBytes returns the ContentClaimFileSizeBytes field if non-nil, zero value otherwise.

### GetContentClaimFileSizeBytesOk

`func (o *FlowFileDTO) GetContentClaimFileSizeBytesOk() (*int64, bool)`

GetContentClaimFileSizeBytesOk returns a tuple with the ContentClaimFileSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentClaimFileSizeBytes

`func (o *FlowFileDTO) SetContentClaimFileSizeBytes(v int64)`

SetContentClaimFileSizeBytes sets ContentClaimFileSizeBytes field to given value.

### HasContentClaimFileSizeBytes

`func (o *FlowFileDTO) HasContentClaimFileSizeBytes() bool`

HasContentClaimFileSizeBytes returns a boolean if a field has been set.

### GetPenalized

`func (o *FlowFileDTO) GetPenalized() bool`

GetPenalized returns the Penalized field if non-nil, zero value otherwise.

### GetPenalizedOk

`func (o *FlowFileDTO) GetPenalizedOk() (*bool, bool)`

GetPenalizedOk returns a tuple with the Penalized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPenalized

`func (o *FlowFileDTO) SetPenalized(v bool)`

SetPenalized sets Penalized field to given value.

### HasPenalized

`func (o *FlowFileDTO) HasPenalized() bool`

HasPenalized returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


