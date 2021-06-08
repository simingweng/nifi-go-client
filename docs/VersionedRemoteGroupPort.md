# VersionedRemoteGroupPort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | Pointer to **string** | The component&#39;s unique identifier | [optional] 
**Name** | Pointer to **string** | The component&#39;s name | [optional] 
**Comments** | Pointer to **string** | The user-supplied comments for the component | [optional] 
**Position** | Pointer to [**Position**](Position.md) |  | [optional] 
**RemoteGroupId** | Pointer to **string** | The id of the remote process group that the port resides in. | [optional] 
**ConcurrentlySchedulableTaskCount** | Pointer to **int32** | The number of task that may transmit flowfiles to the target port concurrently. | [optional] 
**UseCompression** | Pointer to **bool** | Whether the flowfiles are compressed when sent to the target port. | [optional] 
**BatchSize** | Pointer to [**BatchSize**](BatchSize.md) |  | [optional] 
**ComponentType** | Pointer to **string** |  | [optional] 
**TargetId** | Pointer to **string** | The ID of the port on the target NiFi instance | [optional] 
**ScheduledState** | Pointer to **string** | The scheduled state of the component | [optional] 
**GroupIdentifier** | Pointer to **string** | The ID of the Process Group that this component belongs to | [optional] 

## Methods

### NewVersionedRemoteGroupPort

`func NewVersionedRemoteGroupPort() *VersionedRemoteGroupPort`

NewVersionedRemoteGroupPort instantiates a new VersionedRemoteGroupPort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersionedRemoteGroupPortWithDefaults

`func NewVersionedRemoteGroupPortWithDefaults() *VersionedRemoteGroupPort`

NewVersionedRemoteGroupPortWithDefaults instantiates a new VersionedRemoteGroupPort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifier

`func (o *VersionedRemoteGroupPort) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *VersionedRemoteGroupPort) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *VersionedRemoteGroupPort) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *VersionedRemoteGroupPort) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetName

`func (o *VersionedRemoteGroupPort) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VersionedRemoteGroupPort) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VersionedRemoteGroupPort) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VersionedRemoteGroupPort) HasName() bool`

HasName returns a boolean if a field has been set.

### GetComments

`func (o *VersionedRemoteGroupPort) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *VersionedRemoteGroupPort) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *VersionedRemoteGroupPort) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *VersionedRemoteGroupPort) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetPosition

`func (o *VersionedRemoteGroupPort) GetPosition() Position`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *VersionedRemoteGroupPort) GetPositionOk() (*Position, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *VersionedRemoteGroupPort) SetPosition(v Position)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *VersionedRemoteGroupPort) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetRemoteGroupId

`func (o *VersionedRemoteGroupPort) GetRemoteGroupId() string`

GetRemoteGroupId returns the RemoteGroupId field if non-nil, zero value otherwise.

### GetRemoteGroupIdOk

`func (o *VersionedRemoteGroupPort) GetRemoteGroupIdOk() (*string, bool)`

GetRemoteGroupIdOk returns a tuple with the RemoteGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteGroupId

`func (o *VersionedRemoteGroupPort) SetRemoteGroupId(v string)`

SetRemoteGroupId sets RemoteGroupId field to given value.

### HasRemoteGroupId

`func (o *VersionedRemoteGroupPort) HasRemoteGroupId() bool`

HasRemoteGroupId returns a boolean if a field has been set.

### GetConcurrentlySchedulableTaskCount

`func (o *VersionedRemoteGroupPort) GetConcurrentlySchedulableTaskCount() int32`

GetConcurrentlySchedulableTaskCount returns the ConcurrentlySchedulableTaskCount field if non-nil, zero value otherwise.

### GetConcurrentlySchedulableTaskCountOk

`func (o *VersionedRemoteGroupPort) GetConcurrentlySchedulableTaskCountOk() (*int32, bool)`

GetConcurrentlySchedulableTaskCountOk returns a tuple with the ConcurrentlySchedulableTaskCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurrentlySchedulableTaskCount

`func (o *VersionedRemoteGroupPort) SetConcurrentlySchedulableTaskCount(v int32)`

SetConcurrentlySchedulableTaskCount sets ConcurrentlySchedulableTaskCount field to given value.

### HasConcurrentlySchedulableTaskCount

`func (o *VersionedRemoteGroupPort) HasConcurrentlySchedulableTaskCount() bool`

HasConcurrentlySchedulableTaskCount returns a boolean if a field has been set.

### GetUseCompression

`func (o *VersionedRemoteGroupPort) GetUseCompression() bool`

GetUseCompression returns the UseCompression field if non-nil, zero value otherwise.

### GetUseCompressionOk

`func (o *VersionedRemoteGroupPort) GetUseCompressionOk() (*bool, bool)`

GetUseCompressionOk returns a tuple with the UseCompression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCompression

`func (o *VersionedRemoteGroupPort) SetUseCompression(v bool)`

SetUseCompression sets UseCompression field to given value.

### HasUseCompression

`func (o *VersionedRemoteGroupPort) HasUseCompression() bool`

HasUseCompression returns a boolean if a field has been set.

### GetBatchSize

`func (o *VersionedRemoteGroupPort) GetBatchSize() BatchSize`

GetBatchSize returns the BatchSize field if non-nil, zero value otherwise.

### GetBatchSizeOk

`func (o *VersionedRemoteGroupPort) GetBatchSizeOk() (*BatchSize, bool)`

GetBatchSizeOk returns a tuple with the BatchSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchSize

`func (o *VersionedRemoteGroupPort) SetBatchSize(v BatchSize)`

SetBatchSize sets BatchSize field to given value.

### HasBatchSize

`func (o *VersionedRemoteGroupPort) HasBatchSize() bool`

HasBatchSize returns a boolean if a field has been set.

### GetComponentType

`func (o *VersionedRemoteGroupPort) GetComponentType() string`

GetComponentType returns the ComponentType field if non-nil, zero value otherwise.

### GetComponentTypeOk

`func (o *VersionedRemoteGroupPort) GetComponentTypeOk() (*string, bool)`

GetComponentTypeOk returns a tuple with the ComponentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentType

`func (o *VersionedRemoteGroupPort) SetComponentType(v string)`

SetComponentType sets ComponentType field to given value.

### HasComponentType

`func (o *VersionedRemoteGroupPort) HasComponentType() bool`

HasComponentType returns a boolean if a field has been set.

### GetTargetId

`func (o *VersionedRemoteGroupPort) GetTargetId() string`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *VersionedRemoteGroupPort) GetTargetIdOk() (*string, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *VersionedRemoteGroupPort) SetTargetId(v string)`

SetTargetId sets TargetId field to given value.

### HasTargetId

`func (o *VersionedRemoteGroupPort) HasTargetId() bool`

HasTargetId returns a boolean if a field has been set.

### GetScheduledState

`func (o *VersionedRemoteGroupPort) GetScheduledState() string`

GetScheduledState returns the ScheduledState field if non-nil, zero value otherwise.

### GetScheduledStateOk

`func (o *VersionedRemoteGroupPort) GetScheduledStateOk() (*string, bool)`

GetScheduledStateOk returns a tuple with the ScheduledState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledState

`func (o *VersionedRemoteGroupPort) SetScheduledState(v string)`

SetScheduledState sets ScheduledState field to given value.

### HasScheduledState

`func (o *VersionedRemoteGroupPort) HasScheduledState() bool`

HasScheduledState returns a boolean if a field has been set.

### GetGroupIdentifier

`func (o *VersionedRemoteGroupPort) GetGroupIdentifier() string`

GetGroupIdentifier returns the GroupIdentifier field if non-nil, zero value otherwise.

### GetGroupIdentifierOk

`func (o *VersionedRemoteGroupPort) GetGroupIdentifierOk() (*string, bool)`

GetGroupIdentifierOk returns a tuple with the GroupIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupIdentifier

`func (o *VersionedRemoteGroupPort) SetGroupIdentifier(v string)`

SetGroupIdentifier sets GroupIdentifier field to given value.

### HasGroupIdentifier

`func (o *VersionedRemoteGroupPort) HasGroupIdentifier() bool`

HasGroupIdentifier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


