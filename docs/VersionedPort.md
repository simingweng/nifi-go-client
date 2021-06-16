# VersionedPort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | Pointer to **string** | The component&#39;s unique identifier | [optional] 
**Name** | Pointer to **string** | The component&#39;s name | [optional] 
**Comments** | Pointer to **string** | The user-supplied comments for the component | [optional] 
**Position** | Pointer to [**Position**](Position.md) |  | [optional] 
**Type** | Pointer to **string** | The type of port. | [optional] 
**ConcurrentlySchedulableTaskCount** | Pointer to **int32** | The number of tasks that should be concurrently scheduled for the port. | [optional] 
**ScheduledState** | Pointer to **string** | The scheduled state of the component | [optional] 
**AllowRemoteAccess** | Pointer to **bool** | Whether or not this port allows remote access for site-to-site | [optional] 
**ComponentType** | Pointer to **string** |  | [optional] 
**GroupIdentifier** | Pointer to **string** | The ID of the Process Group that this component belongs to | [optional] 

## Methods

### NewVersionedPort

`func NewVersionedPort() *VersionedPort`

NewVersionedPort instantiates a new VersionedPort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersionedPortWithDefaults

`func NewVersionedPortWithDefaults() *VersionedPort`

NewVersionedPortWithDefaults instantiates a new VersionedPort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifier

`func (o *VersionedPort) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *VersionedPort) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *VersionedPort) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *VersionedPort) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetName

`func (o *VersionedPort) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VersionedPort) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VersionedPort) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VersionedPort) HasName() bool`

HasName returns a boolean if a field has been set.

### GetComments

`func (o *VersionedPort) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *VersionedPort) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *VersionedPort) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *VersionedPort) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetPosition

`func (o *VersionedPort) GetPosition() Position`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *VersionedPort) GetPositionOk() (*Position, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *VersionedPort) SetPosition(v Position)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *VersionedPort) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetType

`func (o *VersionedPort) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VersionedPort) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VersionedPort) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *VersionedPort) HasType() bool`

HasType returns a boolean if a field has been set.

### GetConcurrentlySchedulableTaskCount

`func (o *VersionedPort) GetConcurrentlySchedulableTaskCount() int32`

GetConcurrentlySchedulableTaskCount returns the ConcurrentlySchedulableTaskCount field if non-nil, zero value otherwise.

### GetConcurrentlySchedulableTaskCountOk

`func (o *VersionedPort) GetConcurrentlySchedulableTaskCountOk() (*int32, bool)`

GetConcurrentlySchedulableTaskCountOk returns a tuple with the ConcurrentlySchedulableTaskCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurrentlySchedulableTaskCount

`func (o *VersionedPort) SetConcurrentlySchedulableTaskCount(v int32)`

SetConcurrentlySchedulableTaskCount sets ConcurrentlySchedulableTaskCount field to given value.

### HasConcurrentlySchedulableTaskCount

`func (o *VersionedPort) HasConcurrentlySchedulableTaskCount() bool`

HasConcurrentlySchedulableTaskCount returns a boolean if a field has been set.

### GetScheduledState

`func (o *VersionedPort) GetScheduledState() string`

GetScheduledState returns the ScheduledState field if non-nil, zero value otherwise.

### GetScheduledStateOk

`func (o *VersionedPort) GetScheduledStateOk() (*string, bool)`

GetScheduledStateOk returns a tuple with the ScheduledState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledState

`func (o *VersionedPort) SetScheduledState(v string)`

SetScheduledState sets ScheduledState field to given value.

### HasScheduledState

`func (o *VersionedPort) HasScheduledState() bool`

HasScheduledState returns a boolean if a field has been set.

### GetAllowRemoteAccess

`func (o *VersionedPort) GetAllowRemoteAccess() bool`

GetAllowRemoteAccess returns the AllowRemoteAccess field if non-nil, zero value otherwise.

### GetAllowRemoteAccessOk

`func (o *VersionedPort) GetAllowRemoteAccessOk() (*bool, bool)`

GetAllowRemoteAccessOk returns a tuple with the AllowRemoteAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowRemoteAccess

`func (o *VersionedPort) SetAllowRemoteAccess(v bool)`

SetAllowRemoteAccess sets AllowRemoteAccess field to given value.

### HasAllowRemoteAccess

`func (o *VersionedPort) HasAllowRemoteAccess() bool`

HasAllowRemoteAccess returns a boolean if a field has been set.

### GetComponentType

`func (o *VersionedPort) GetComponentType() string`

GetComponentType returns the ComponentType field if non-nil, zero value otherwise.

### GetComponentTypeOk

`func (o *VersionedPort) GetComponentTypeOk() (*string, bool)`

GetComponentTypeOk returns a tuple with the ComponentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentType

`func (o *VersionedPort) SetComponentType(v string)`

SetComponentType sets ComponentType field to given value.

### HasComponentType

`func (o *VersionedPort) HasComponentType() bool`

HasComponentType returns a boolean if a field has been set.

### GetGroupIdentifier

`func (o *VersionedPort) GetGroupIdentifier() string`

GetGroupIdentifier returns the GroupIdentifier field if non-nil, zero value otherwise.

### GetGroupIdentifierOk

`func (o *VersionedPort) GetGroupIdentifierOk() (*string, bool)`

GetGroupIdentifierOk returns a tuple with the GroupIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupIdentifier

`func (o *VersionedPort) SetGroupIdentifier(v string)`

SetGroupIdentifier sets GroupIdentifier field to given value.

### HasGroupIdentifier

`func (o *VersionedPort) HasGroupIdentifier() bool`

HasGroupIdentifier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


